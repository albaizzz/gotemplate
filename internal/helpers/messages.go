package helpers

import (
	"fmt"
	"net/http"

	"gotemplate/internal/models"

	"gotemplate/internal/consts"

	"github.com/gin-gonic/gin"
)

var Message = map[int]models.APIResponse{
	consts.APINotFoundCode:         models.APIResponse{Data: "Data not found", Code: consts.APINotFoundCode, HttpCode: http.StatusNotFound},
	consts.APIInvalidParameterCode: models.APIResponse{Data: "Invalid parameters", Code: consts.APIInvalidParameterCode, HttpCode: http.StatusBadRequest},
	consts.APISuccessCode:          models.APIResponse{Data: "Success", Code: consts.APISuccessCode, HttpCode: http.StatusAccepted},
	consts.APIUnknownCode:          models.APIResponse{Data: "Internal Server error", Code: consts.APIUnknownCode, HttpCode: http.StatusInternalServerError},
	consts.APIConflictCode:         models.APIResponse{Data: "Data already exist", Code: consts.APIConflictCode, HttpCode: http.StatusConflict},
}

func ServiceResp(apiStatusCode int, data ...interface{}) (resp models.APIResponse) {
	if apiStatusCode != consts.APISuccessCode {
		if data != nil {
			errData, ok := data[0].(models.ErrorDetails)
			if ok {
				resp = Message[apiStatusCode]
				resp.Data = fmt.Sprintf("%s, %s", resp.Data, errData.ErrorMsg)
			} else {
				resp = Message[apiStatusCode]
			}
		} else {
			resp = Message[apiStatusCode]
		}
	} else {
		if data == nil {
			resp = Message[apiStatusCode]
		} else {
			resp.Code = apiStatusCode
			resp.Data = data[0]
			resp.HttpCode = http.StatusOK
		}
	}
	return
}

func APIResponse(ctx *gin.Context, resp models.APIResponse) {
	apidata := map[string]interface{}{
		"code": resp.Code,
		"data": resp.Data,
	}
	ctx.JSON(resp.HttpCode, apidata)
}
