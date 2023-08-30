package fetchext

import (
	"fmt"
	"signature"
)

func Scale(ctx *signature.Context) (*signature.Context, error) {

	if len(ctx.Fetchresponse) == 0 {
		// Do some things maybe... store some state maybe

		// Ask for a fetch
		ctx.Fetchrequest = "https://ifconfig.me"
		return ctx, nil
	} else {
		ctx.Fetchrequest = ""
		ctx.Output = fmt.Sprintf("Input=%s We did a fetch %s", ctx.Input, string(ctx.Fetchresponse))

		return signature.Next(ctx)
	}

}
