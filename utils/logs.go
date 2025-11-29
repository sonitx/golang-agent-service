package utils

import (
	"encoding/json"
	"fmt"
	"path/filepath"
	"runtime"

	"github.com/sirupsen/logrus"
)

func ShowErrorLogs(errData error) {
	if errData != nil {
		file, line := Caller(2)
		logrus.SetFormatter(&logrus.TextFormatter{})
		logrus.Errorf("%s[%d] %s", filepath.Base(file), line, errData.Error())
	}
}

func ShowInfoLogs(format string, a ...any) {
	file, line := Caller(2)
	logrus.SetFormatter(&logrus.TextFormatter{})
	if len(a) == 0 {
		logrus.Infof("%s[%d] %s", filepath.Base(file), line, format)
	} else {
		logrus.Infof("%s[%d] %s", filepath.Base(file), line, fmt.Sprintf(format, a...))
	}
}

func ShowJsonLog(data interface{}) {
	if data != nil {
		file, line := Caller(2)
		logrus.SetFormatter(&logrus.TextFormatter{})
		jData := MarshalToJson(data)
		logrus.Infof("%s[%d] %v", filepath.Base(file), line, jData)
	}
}

func Caller(level int) (string, int) {
	_, file, line, _ := runtime.Caller(level)
	return file, line
}

func ToJson(data interface{}) string {
	s, ok := data.(string)

	if ok {
		return s
	}

	result, err := json.Marshal(data)

	if err != nil {
		return ""
	}

	return string(result)
}
