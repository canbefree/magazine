package utils

import (
	jww "github.com/spf13/jwalterweatherman"
)

func FatalIfError(err error) {
	jww.ERROR.Fatalf(err.Error())
}
