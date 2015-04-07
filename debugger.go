package debugger

import "runtime"

func Where() string {
	var _, file, line, _ = runtime.Caller(1)
	return file + ":" + line
}
