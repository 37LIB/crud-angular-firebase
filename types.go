package ling

import (
	"bytes"
)

type Document struct {
	Text   string   `json:"text"`
	Tokens []*Token `json:"tokens"`
	Spans  []*Span  `json:"spans"`
	Lang   string   `json:"