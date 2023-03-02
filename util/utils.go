package util

import (
	"fmt"
	"github.com/liuzl/da"
	"strings"
)

type is func(rune) bool

func StringIs(s string, f is) bool {
	for _, c