package utils

import (
	"fmt"
	"runtime"
	"strings"
)

func WhereAmI(skipList ...int) string {
	var skip int
	if skipList == nil {
		skip = 1
	} else {
		skip = skipList[0]
	}
	function, file, line, _ := runtime.Caller(skip)
	return fmt.Sprintf("Function: %s \nFile: %s:%d", chopPath(runtime.FuncForPC(function).Name(), "."), file, line)
}

// return the source filename after the last slash
func chopPath(original string, pathChar string) string {
	i := strings.LastIndex(original, pathChar)
	if i == -1 {
		return original
	}
	return original[i+1:]
}
