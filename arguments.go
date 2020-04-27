package hrvst

import (
	"net/url"
)

// Arguments [`key` => `value`] object of items should be passed as GET query string
type Arguments map[string]string

// Defaults is just `empty` Arguments object
func Defaults() Arguments {
	return make(Arguments)
}

func (args Arguments) toURLValues() url.Values {
	v := url.Values{}
	for key, value := range args {
		v.Set(key, value)
	}
	return v
}
