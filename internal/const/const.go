package consts

//Pkg default package name
// const Pkg = "bitbucket.org/kudoindonesia/ovo_payments"

//Region environment constant
const (
	//EnvKey as keys to get environment variable
	EnvKey = "ENV"
	//EnvProduction as environment variable for production
	EnvProduction = "production"
	//EnvStaging as environment variable for staging
	EnvStaging = "staging"
	//EnvDevelopment as environment variable for development
	EnvDevelopment = "development"
)

//FilesDevelPaths collection of possible file location
// var FilesDevelPaths = []string{
// 	"./files/",
// 	fmt.Sprintf("%s/src/%s/files/", env.Get("GOPATH"), Pkg),
// }

//Region static files
const (
	//FilesDevelPath configuration file path for development environment
	FilesDevelPath = "./files/"
	//FilesStagProdPath configuration file path for staging environment
	FilesStagProdPath = "/var/www/microservice/files/"
)

//Region Application Constant
const (
	//HTTPPort default http port
	HTTPPort = 9090
	//GRPCPort default grpc port
	GRPCPort = 9091
)

const (
	//ServiceTypeHTTP marks the usecase service type for HTTP operation
	ServiceTypeHTTP = "HTTP"
	//ServiceTypeMQ marks the usecasee service type for MQ Operation
	ServiceTypeMQ = "MQ"
)

const (
	//RunTimeHTTP indicator for HTTP Runtime
	RunTimeHTTP = "HTTP"
	//RunTimeMQ indicator for MQ Runtime
	RunTimeMQ = "MQ"
)
