package conf

import (
	"impl"
)

type api struct {
	ServiceUri string
	InterClass interface{}
}

func GetConf() map[string]api {
	var ApiConfig map[string]api
	ApiConfig = make(map[string]api)

	ApiConfig["test"] = api{"/service/test", impl.Test{}}
	return ApiConfig
}
