package util

import (
	_const "awesomeProject/src/const"
	"path"
	"runtime"
	"strconv"
	"strings"
)

var SuccessStatus = &_const.Status{
	Code:    0,
	Message: "success",
	Data:    nil,
}

func NewError(err error) *_const.Status {
	return &_const.Status{
		Code:    400,
		Message: err.Error(),
		Stack:   stack,
		Data:    err.Error(),
	}
}

var stack = []string{Caller()}

func Caller() string {
	pc, file, line, ok := runtime.Caller(2)
	if !ok {
		file = "???"
		line = 0
	}
	funcName := "???"
	f := runtime.FuncForPC(pc)
	if f != nil {
		funcName = f.Name()
	}
	_, filename := path.Split(file)
	flist := strings.Split(funcName, ".")
	funcName = flist[len(flist)-1]
	format := "filename:" + filename + ",func:" + funcName + ",line:" + strconv.FormatInt(int64(line), 10)
	return format
}
