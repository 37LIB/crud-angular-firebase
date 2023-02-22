package ling

import (
	"strings"
	"unicode"

	"github.com/liuzl/ling/util"
	"github.com/liuzl/tokenizer"
)

const Lower = "lower"

func init() {
	Processors["_"] = &Tokenizer{}
}

func Type(text string) TokenType {
	switch {
	case util.StringIs