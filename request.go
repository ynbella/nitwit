package nitwit

import (
	"net/url"
	"strconv"
)

type Request url.URL

type RequestOption func(*Request)

func WithQueryParams(key string, values ...string) RequestOption {
	return func(req *Request) {
		params := new(url.Values)
		for _, value := range values {
			params.Add(key, value)
		}
		req.RawQuery = params.Encode()
	}
}

func NewRequest(endpoint string, opts ...RequestOption) (*Request, error) {
	req, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}

	// Loop through each option
	for _, opt := range opts {
		opt((*Request)(req))
	}

	return (*Request)(req), nil
}

func WithID(id string) RequestOption {
	return WithQueryParams("id", id)
}

func WithSince(since string) RequestOption {
	return WithQueryParams("since", since)
}

func WithMax(max string) RequestOption {
	return WithQueryParams("max", max)
}

func WithLimit(limit int) RequestOption {
	return WithQueryParams("id", strconv.FormatInt(int64(limit), 10))
}

func WithFilter(filter string) RequestOption {
	return WithQueryParams("filter", filter)
}
