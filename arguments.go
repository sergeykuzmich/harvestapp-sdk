package hrvst

import (
	"net/url"
	"strings"

	"github.com/sergeykuzmich/harvestapp-sdk/flags"
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
		if !strings.HasPrefix(key, flags.Prefix) {
			v.Set(key, value)
		}
	}
	return v
}
