package app

import (
	"fmt"
	"time"
)

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

func (a *App) Logger(data string) {
	log := fmt.Sprintf("%s %s \n", time.Now(), data)
	a.Log.WriteString(log)
}
