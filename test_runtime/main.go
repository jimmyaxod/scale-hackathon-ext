package main

import (
	"context"
	"fmt"
	_ "image/jpeg"
	"io"
	"net/http"
	sig "signature"

	scale "github.com/loopholelabs/scale"
	"github.com/loopholelabs/scale/scalefunc"
)

func main() {
	fn, err := scalefunc.Read("./local-fetchext-latest.scale")
	if err != nil {
		panic(err)
	}

	r, err := scale.New(context.Background(), sig.New, []*scalefunc.Schema{fn})
	if err != nil {
		panic(err)
	}

	i, err := r.Instance(nil)
	if err != nil {
		panic(err)
	}

	s := sig.New()
	s.Context.Input = "Hello there! "
	for {
		err = i.Run(context.Background(), s)
		if err != nil {
			panic(err)
		}

		if s.Context.Fetchrequest == "" {
			break
		} else {
			// Extension logic goes here...
			resp, err := http.Get(s.Context.Fetchrequest)
			if err != nil {
				panic(err)
			}
			body, err := io.ReadAll(resp.Body)
			s.Context.Fetchresponse = body
		}
	}

	fmt.Printf("Scale function returned with %s\n", s.Context.Output)

}
