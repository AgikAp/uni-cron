package helpers

import (
	"admin_backend/app/exceptions"
)

func Panics(msg string, err interface{}) {
	if err != nil {
		panic(exceptions.NewBusinnessException(err, msg))
	}
}

func PanicsConditional(msg string, err interface{}, condition bool) {
	if condition {
		panic(exceptions.NewBusinnessException(err, msg))
	}
}
