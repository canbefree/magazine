package log

import (
	"context"

	jww "github.com/spf13/jwalterweatherman"
)

func Tracef(ctx context.Context, format string, v ...interface{}) {
	jww.TRACE.Printf(format, v...)
}

func Infof(ctx context.Context, format string, v ...interface{}) {
	jww.INFO.Printf(format, v...)
}

func Errorf(ctx context.Context, format string, v ...interface{}) {
	jww.ERROR.Printf(format, v...)
}

func init() {
	jww.SetStdoutThreshold(jww.LevelTrace)
}
