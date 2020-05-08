package hrvst

import (
	"net/url"
	"strings"
)

// Arguments [`key`=>`value`] dictionary should be passed as a GET query string
type Arguments map[string]string

// Defaults is just an `empty` Arguments object
func Defaults() Arguments {
	return make(Arguments)
}

// toURLValues converts [`key`=>`value`] dictionary to a GET query string
func (args Arguments) toURLValues() url.Values {
	v := url.Values{}
	for key, value := range args {
		// Skip all internal flags while making query string
		if !strings.HasPrefix(key, PREFIX) {
			v.Set(key, value)
		}
	}
	return v
}
