package http

import (
	"encoding/json"
	"io"
	"io/ioutil"

	"github.com/patryklyczko/Video_Streming_GO/pkg/db"
	"github.com/valyala/fasthttp"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
)

func (i *HTTPInstanceAPI) postVideoFolder(ctx *fasthttp.RequestCtx) {
	i.log.Debugf("Got request to add video")
	var VideoUpdate db.VideoRequestFolder
	body := ctx.Request.Body()

	if err := json.Unmarshal(body, &VideoUpdate); err != nil {
		ctx.Response.SetStatusCode(fasthttp.StatusBadRequest)
		i.log.Errorf("error while unmarshaling %v", err)
		return
	}

	if err := i.api.PostVideoFolder(VideoUpdate); err != nil {
		i.log.Debugf("Error while updating video %v", err)
		return
	}
	i.log.Debugf("Upload video")
	ctx.Response.SetStatusCode(fasthttp.StatusCreated)
}

func (i *HTTPInstanceAPI) watchVideo(ctx *fasthttp.RequestCtx) {
	ctx.Response.Header.Set("Content-Type", "video/mp4")
	i.log.Debugf("Got request to extract video")
	var video *gridfs.DownloadStream
	var err error

	args := ctx.QueryArgs()
	VideoRequest := db.VideoRequest{
		Name: string(args.Peek("name")),
	}
	if video, err = i.api.WatchVideo(VideoRequest); err != nil {
		ctx.Response.SetStatusCode(fasthttp.StatusBadRequest)
		i.log.Errorf("error while getting video %v", err)
		return
	}
	_, _ = io.Copy(ctx.Response.BodyWriter(), video)
}

func (i *HTTPInstanceAPI) uploadVideo(ctx *fasthttp.RequestCtx) {
	indexData, err := ioutil.ReadFile("./static/index.html")
	ctx.Write(indexData)
	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.Response.SetStatusCode(fasthttp.StatusBadRequest)
		i.log.Errorf("error while getting video %v", err)
		return
	}
	filename := file.Filename

	fileOpened, err := file.Open()
	if err != nil {
		ctx.Response.SetStatusCode(fasthttp.StatusBadRequest)
		i.log.Errorf("error while openning video %v", err)
		return
	}

	fileBytes, err := io.ReadAll(fileOpened)
	if err != nil {
		ctx.Response.SetStatusCode(fasthttp.StatusBadRequest)
		i.log.Errorf("error while converting video %v", err)
		return
	}

	if err := i.api.AddVideo(fileBytes, filename); err != nil {
		i.log.Debugf("Error while updating video %v", err)
		return
	}
	i.log.Debugf("Upload video")
	ctx.Response.SetStatusCode(fasthttp.StatusCreated)

}

// func (i *HTTPInstanceAPI) videos(ctx *fasthttp.RequestCtx) {
// 	i.log.Debugf("Got request show videos")
// 	var videoFilter db.VideoFilter
// 	body := ctx.Request.Body()

// 	if err := json.Unmarshal(body, &videoFilter); err != nil {
// 		ctx.Response.SetStatusCode(fasthttp.StatusBadRequest)
// 		i.log.Errorf("error while unmarshaling %v", err)
// 		return
// 	}

// 	if
// 	i.log.Debugf("Upload video")
// 	ctx.Response.SetStatusCode(fasthttp.StatusCreated)
// }
