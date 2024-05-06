package utils

import (
	"strconv"
	"time"
)

func AddPrefixToFilename(filename string) string {

	unix := time.Now().Unix()

	return strconv.Itoa(int(unix)) + "_" + filename
}
