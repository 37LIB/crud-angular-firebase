
// generated by jsonenums -type=TokenType; DO NOT EDIT

package ling

import (
	"encoding/json"
	"fmt"
)

var (
	_TokenTypeNameToValue = map[string]TokenType{
		"EOF":     EOF,
		"Space":   Space,
		"Symbol":  Symbol,
		"Number":  Number,
		"Letters": Letters,
		"Punct":   Punct,
		"Word":    Word,
	}

	_TokenTypeValueToName = map[TokenType]string{
		EOF:     "EOF",
		Space:   "Space",
		Symbol:  "Symbol",
		Number:  "Number",
		Letters: "Letters",
		Punct:   "Punct",
		Word:    "Word",
	}
)

func init() {
	var v TokenType
	if _, ok := interface{}(v).(fmt.Stringer); ok {
		_TokenTypeNameToValue = map[string]TokenType{
			interface{}(EOF).(fmt.Stringer).String():     EOF,
			interface{}(Space).(fmt.Stringer).String():   Space,
			interface{}(Symbol).(fmt.Stringer).String():  Symbol,
			interface{}(Number).(fmt.Stringer).String():  Number,
			interface{}(Letters).(fmt.Stringer).String(): Letters,
			interface{}(Punct).(fmt.Stringer).String():   Punct,
			interface{}(Word).(fmt.Stringer).String():    Word,
		}
	}
}

// MarshalJSON is generated so TokenType satisfies json.Marshaler.
func (r TokenType) MarshalJSON() ([]byte, error) {
	if s, ok := interface{}(r).(fmt.Stringer); ok {
		return json.Marshal(s.String())
	}
	s, ok := _TokenTypeValueToName[r]
	if !ok {
		return nil, fmt.Errorf("invalid TokenType: %d", r)
	}
	return json.Marshal(s)
}

// UnmarshalJSON is generated so TokenType satisfies json.Unmarshaler.
func (r *TokenType) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("TokenType should be a string, got %s", data)
	}
	v, ok := _TokenTypeNameToValue[s]
	if !ok {
		return fmt.Errorf("invalid TokenType %q", s)
	}
	*r = v
	return nil
}