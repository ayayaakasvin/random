package logger

import (
	"fmt"
	"os"
)

func ErrorLog(msg string, args ...any) {
	msg = fmt.Sprintf("[ERR]%s\n", msg)
	fmt.Fprintf(os.Stderr, msg, args...)
}