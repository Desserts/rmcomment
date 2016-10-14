package rmcomment

import (
	"fmt"
)

var lineComment []rune
var startBlockComment []rune
var endBlockComment []rune

const (
	INIT               = 0
	LINECOMMENT0       = 1
	INLINECOMMENT      = 4
	STARTBLOCKCOMMENT0 = -1
	ENDBLOCKCOMMENT0   = 3
	INBLOCKCOMMENT     = 2
	ESCAPE             = 9 //转义
)

func InitRm(linecomment, blockcommentstart, blockcommentend string) {
	lineComment = []rune(linecomment)
	startBlockComment = []rune(blockcommentstart)
	endBlockComment = []rune(blockcommentend)
}

func Rm(content string) string {
	chars := []rune(content)
	res := dfaRemove(chars)
	return res
}

func dfaRemove(chars []rune) string {
	var c rune
	res := make([]rune, 0)

	state := INIT
	for _, c = range chars {
		switch state {
		case INIT:
			if c == lineComment[0] {
				state = LINECOMMENT0
			} else if c == startBlockComment[0] {
				state = STARTBLOCKCOMMENT0
			} else {
				res = append(res, c)
			}
		case LINECOMMENT0:
			if c == lineComment[1] {
				state = INLINECOMMENT
			} else {
				res = append(res, c)
				state = INIT
			}
		case STARTBLOCKCOMMENT0:
			if c == startBlockComment[1] {
				state = INBLOCKCOMMENT
			} else {
				res = append(res, c)
				state = INIT
			}
		case INBLOCKCOMMENT:
			if c == endBlockComment[0] {
				state = ENDBLOCKCOMMENT0
			}
		case ENDBLOCKCOMMENT0:
			if c == endBlockComment[1] {
				state = INIT
			} else {
				state = INBLOCKCOMMENT
			}
		case INLINECOMMENT:
			if c == '\\' {
				state = ESCAPE
			} else if c == '\n' {
				state = INIT
			}
		case ESCAPE:
			if c == '\\' {
				state = ESCAPE
			} else {
				state = INIT
			}
		}
	}

	return string(res)
}
