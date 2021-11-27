package utils

import (
	"fmt"
	"log"

	jww "github.com/spf13/jwalterweatherman"
)

func FatalIfError(err error) {
	jww.ERROR.Fatalf(err.Error())
}

func PanicIfErr(err error) {
	if err != nil {
		log.Fatalf("err:%v", err)
	}
	panic(fmt.Sprintf("err:%v", err.Error()))
}
