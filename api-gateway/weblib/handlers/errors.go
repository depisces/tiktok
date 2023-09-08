package handlers

import (
	"api-gateway/pkg/logging"
	"errors"
)

func PanicIfUserError(err error) {
	if err != nil {
		err = errors.New("userService--" + err.Error())
		logging.Info(err)
		panic(err)
	}
}

func PanicIfFeedError(err error) {
	if err != nil {
		err = errors.New("feedService--" + err.Error())
		logging.Info(err)
		panic(err)
	}
}

func PanicIfPublishError(err error) {
	if err != nil {
		err = errors.New("publishService--" + err.Error())
		logging.Info(err)
		panic(err)
	}
}
