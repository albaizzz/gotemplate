package svc

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/DataDog/dd-trace-go/tracer"
)

//router performs routing
func (h *httpServer) route() {
	mt := muxtrace.NewMuxTracer("ovo_payments", tracer.DefaultTracer)
	r := h.ui.Router

	ovo := r.PathPrefix("/msovo").Subrouter()
	ovoV1 := ovo.PathPrefix("/v1/").Subrouter()
	springV1 := ovoV1.PathPrefix("/spring").Subrouter()

	mt.HandleFunc(springV1, "/get-saldo", h.legacy(
		handler.GetBalanceHTTPHandler,
		ucase.GetOVOCashBalanceScenario,
		middleware.AuthMiddleware,
	)).Methods(http.MethodPost)

	mt.HandleFunc(springV1, "/fund-histories", h.legacy(
		handler.GetFundHistoriesHTTPHandler,
		ucase.GetFundHistoriesScenario,
		middleware.AuthMiddleware,
	)).Methods(http.MethodPost)

	mt.HandleFunc(springV1, "/wallet-type", h.legacy(
		handler.GetWalletTypeHTTPHandler,
		ucase.GetWalletType,
		middleware.AuthMiddleware,
	)).Methods(http.MethodPost)
}

//handle orchestrates http handles with middleware
func (h *httpServer) handle(hfunc httpHandlerFunc, scenario ucase.ScenarioFunc, mfs ...middleware.ExecFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var rerr error
		var d interface{}
		var s string
		st := time.Now()
		// defer func() {
		// 	go ddog.PublishDatadog(r.URL.Host+r.URL.Path, ddog.TypeHandler, st, rerr)
		// }()
		if err := middleware.Filters(h.ui.MiddlewareContext, r, mfs); err != nil {
			rerr = err
			h.print(w, nil, "", err)
			return
		}
		d, s, rerr = hfunc(r, h.ui, scenario)
		h.print(w, d, s, rerr)
		return
	}
}

//legacy orchestrates http handles with middleware for deposit grab pay legacy
func (h *httpServer) legacy(hfunc httpHandlerFunc, scenario ucase.ScenarioFunc, mfs ...middleware.ExecFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var rerr error
		var d interface{}
		var s string
		st := time.Now()
		// defer func() {
		// 	go ddog.PublishDatadog(r.URL.Host+r.URL.Path, ddog.TypeHandler, st, rerr)
		// }()
		if err := middleware.Filters(h.ui.MiddlewareContext, r, mfs); err != nil {
			rerr = err
			h.printlg(w, nil, "", err)
			return
		}
		d, s, rerr = hfunc(r, h.ui, scenario)
		h.printlg(w, d, s, rerr)
		return
	}
}

//print prints as a json and formatted string
func (h *httpServer) print(w http.ResponseWriter, data interface{}, msg string, err error) {
	var res HTTPResponse
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		m, c := derrors.TranslateHTTPCode(err)
		res.Message = m
		w.WriteHeader(c)
	} else {
		res.Message = msg
		res.Data = data
	}
	json.NewEncoder(w).Encode(res)
}

//printlg prints as a json and formatted string for DGP legacy
func (h *httpServer) printlg(w http.ResponseWriter, data interface{}, msg string, err error) {
	var sdata interface{}
	var scode int
	var smsg string
	w.Header().Set("Content-Type", "application/json")
	defer func() {
		res := LegacyHTTPResponse{
			Data:          sdata,
			APIStatus:     smsg,
			APIStatusCode: scode,
		}
		// olog.InfoLn(res.Data)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(res)
	}()
	if err != nil {
		if err.Error() == errors.New(consts.LegacySpringGenericError).Error() {
			sdata, scode, smsg = springResp(data)
			return
		}
		smsg, scode = derrors.TranslateHTTPCode(err)
		return
	}
	smsg, scode = derrors.TranslateHTTPCode(errors.New(consts.LegacyAPIGeneralSuccess))
	sdata = data
	return
}
