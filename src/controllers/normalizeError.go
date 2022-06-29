package controllers

import (
	"errors"
	"fmt"
	"strings"

	"gorm.io/gorm"
)

func NormalizerErrorDB(e error) (statusCode int, err error) {

	statusCode = 400
	err = e

	stringERR := e.Error()
	if strings.Contains(stringERR, "ERROR") {
		statusCode = 500
		err = errors.New("Internal error!")
		err = e
	}

	if errors.Is(e, gorm.ErrRecordNotFound) {
		statusCode = 404
		err = errors.New(fmt.Sprintf("Register not found!"))
	}

	return
}
