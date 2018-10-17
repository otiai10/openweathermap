package openweathermap

import "net/url"

// Option ...
type Option struct {
	Units string
}

// DefaultOption ...
var DefaultOption = Option{
	Units: "metrics",
}

// Query ...
func (opt *Option) Query() url.Values {
	v := url.Values{
		"units": {opt.Units},
	}
	return v
}
