package parser

import (
	"fmt"
	"strings"
)

var traceLevel int = 0

const traceIdentPlaceholder string = "\t"

func trace(msg string) string {
	incIndent()
	tracePrint("BEGIN " + msg)
	return msg
}

func untrace(msg string) {
	tracePrint("END " + msg)
	decIndent()
}

func tracePrint(fs string) {
	fmt.Printf("%s%s\n", indentLevel(), fs)
}

func indentLevel() string {
	return strings.Repeat(traceIdentPlaceholder, traceLevel-1)
}

func incIndent() {
	traceLevel = traceLevel + 1
}

func decIndent() {
	traceLevel = traceLevel - 1
}
