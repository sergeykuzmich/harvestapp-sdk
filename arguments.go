package hrvst

import (
	"net/url"
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
		v.Set(key, value)
	}
	return v
}
