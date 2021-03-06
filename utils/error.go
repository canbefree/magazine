package utils

import (
	"fmt"

	jww "github.com/spf13/jwalterweatherman"
)

func FatalIfError(err error) {
	jww.ERROR.Fatalf(err.Error())
}

func PanicIfErr(err error) {
	if err == nil {
		return
	}
	jww.INFO.Fatalf(err.Error())
	panic(fmt.Sprintf("err:%v", err.Error()))
}
