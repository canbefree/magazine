package log

import (
	"context"
	"fmt"
)

//  待完善的日志器 TODO

func Info(ctx context.Context, msg string) {
	fmt.Println(msg)
}

func Error(ctx context.Context, msg string) {
	fmt.Println(msg)
}
