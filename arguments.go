package hrvst

import (
	"net/url"
)

// [`key` => `value`] object of items should be passed as GET arguments
type Arguments map[string]string

// Returns `empty` Arguments object
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
