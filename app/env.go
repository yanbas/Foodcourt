package app

import (
	"fmt"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var APPLICATION_NAME = "FOODCOURT SERVICE"
var LOGIN_EXPIRATION_DURATION = time.Duration(1) * time.Hour
var JWT_SIGNING_METHOD = jwt.SigningMethodHS256
var JWT_SIGNATURE_KEY = []byte("secret key")

var RESULT_ERROR = ResultData{
	Code:    333,
	Success: false,
	Data:    "Error Find Data",
}

var BINDING_JSON_ERROR = ResultData{
	Code:    310,
	Success: false,
	Data:    "Error Binding Request Data",
}

var RESULTS_ERROR = ResultData{
	Code:    315,
	Success: false,
	Data:    "Error GET All Data",
}

var INSERT_DATA_ERROR = ResultData{
	Code:    330,
	Success: false,
	Data:    "Error Insert Data",
}

var UPDATE_DATA_ERROR = ResultData{
	Code:    370,
	Success: false,
	Data:    "Error Update Data",
}

var DELET_DATA_ERROR = ResultData{
	Code:    370,
	Success: false,
	Data:    "Error Update Data",
}

var INVALID_ATHORIZATION = ResultData{
	Code:    999,
	Success: false,
	Data:    "Invalid Authorization",
}

func (a *App) Logger(data string) {
	msg := fmt.Sprintf("%s %s \n", time.Now(), data)
	log.Println(msg)
}
