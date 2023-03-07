package http

import (
	"encoding/json"

	"github.com/patryklyczko/Video_Streming_GO/pkg/db"
	"github.com/valyala/fasthttp"
)

func (i *HTTPInstanceAPI) videoInfo(ctx *fasthttp.RequestCtx) {
	i.log.Debugf("Got request to retrive video info")
	var videoInfo *db.VideoInformation
	var body []byte
	var err error

	args := ctx.QueryArgs()
	name := string(args.Peek("filename"))
	if videoInfo, err = i.api.VideoInfo(name); err != nil {
		ctx.Response.SetStatusCode(fasthttp.StatusBadRequest)
		i.log.Errorf("error getting video info %v", err)
		return
	}

	if body, err = json.Marshal(videoInfo); err != nil {
		ctx.Response.SetStatusCode(fasthttp.StatusBadRequest)
		i.log.Errorf("error marshaling data %v", err)
		return
	}

	ctx.Response.SetStatusCode(fasthttp.StatusOK)
	ctx.Response.SetBody(body)
}
