package utils

import (
	"errors"
	"strconv"

	"github.com/sirupsen/logrus"
)

// RunTimesUntilSuccess 函数运行失败则重试，直到times次后返回错误信息
func RunTimesUntilSuccess(f func() error, times int) error {
	err := errors.New("")
	for i := 0; i < times; i++ {
		if err = f(); err == nil {
			return nil
		}
		logrus.Warning("Function error with " + strconv.FormatInt(int64(i), 10) + " times: " + err.Error())
	}
	return errors.New("Function fail after " + strconv.FormatInt(int64(times), 10) + " times: " + err.Error())
}
