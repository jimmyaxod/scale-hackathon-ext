package fetchext

import (
	"errors"
	"fmt"
	"signature"
)

func Scale(ctx *signature.Context) (*signature.Context, error) {

	// Go get something
	data1, err := fetch(ctx, "https://ifconfig.me")
	if err != nil {
		return ctx, nil
	}

	// Go get something
	data2, err := fetch(ctx, "https://api.ipify.org?format=json")
	if err != nil {
		return ctx, nil
	}

	// Show it off!
	ctx.Output = fmt.Sprintf("Input=%s We did a fetch to ifconfig [%s] then ipify [%s]", ctx.Input, data1, data2)

	return signature.Next(ctx)
}

// Simple cache so we keep state
var cache map[string][]byte = make(map[string][]byte)

/**
 * Fetch extension!
 *
 */
func fetch(ctx *signature.Context, url string) ([]byte, error) {

	// If there's a response ready, accept it
	if len(ctx.Fetchresponse) > 0 {
		cache[ctx.Fetchrequest] = ctx.Fetchresponse
	}

	ctx.Fetchrequest = ""
	ctx.Fetchresponse = make([]byte, 0)

	// Now check the one asked for
	d, ok := cache[url]
	if ok {
		return d, nil
	}
	ctx.Fetchrequest = url
	return nil, errors.New("Return this to host")
}
