package logger

import "log"

/**
存在错误记录日志
*/
func LogError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
