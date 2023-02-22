package normalize

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"github.com/golang/glog"
	"github.com/liuzl/da"
	"github.com/liuzl/ling/util"
	"github.com/liuzl/tokenizer"
	"strings"
)

var Funcs = make(map[string]util.ConvertFunc)

func genFuncs(lang string, contents ...*string) error {
	var dicts []*da.Dict
	for _, c := range contents {
		data, err := base64.StdEncoding.DecodeString(*c)
		if err != nil {
			glog