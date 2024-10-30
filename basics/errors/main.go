package main

import (
	"errors"
	"fmt"
	"log/slog"
)

var (
	ErrorNetwork = errors.New("network error")
	ErrorInput   = errors.New("input error")
)

type ComplexErrorNetwork struct {
	Code int
	Msg  string
}

func (c ComplexErrorNetwork) Error() string {
	return fmt.Sprintf("Network Error %s: code %d", c.Msg, c.Code)
}

func networkCall() error {
	return ErrorNetwork
}

func networkCallWithCmplexError() error {
	return ComplexErrorNetwork{
		Code: 404,
		Msg:  "Something went wrong",
	}
}
func main() {
	err := networkCall()

	if errors.Is(err, ErrorNetwork) {
		slog.Error("Received network error")
	}

	err = networkCallWithCmplexError()

	if err != nil {
		fmt.Println(err)
	}

}
