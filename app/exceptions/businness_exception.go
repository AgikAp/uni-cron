package exceptions

import "fmt"

type BusinnessException struct {
	Message string
	Err     interface{}
}

func NewBusinnessException(err interface{}, message string) *BusinnessException {
	return &BusinnessException{Err: err, Message: message}
}

func (e *BusinnessException) Error() string {
	return fmt.Sprintf("%v, %v", e.Message, e.Err)
}
