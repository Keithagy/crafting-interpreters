package lex

import (
	"bufio"
)

var Tokenizer bufio.SplitFunc = bufio.ScanWords
