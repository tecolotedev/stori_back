package utils

import (
	"strconv"
	"time"

	"github.com/fatih/color"
)

func AddPrefixToFilename(filename string) string {

	unix := time.Now().Unix()

	return strconv.Itoa(int(unix)) + "_" + filename
}

func ErrorLog(err error) {
	color.Red("ERROR %s | %v \n", time.Now().String(), err)
}
