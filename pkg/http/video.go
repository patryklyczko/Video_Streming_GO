package http

import (
	"encoding/json"
	"io"

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
	i.log.Debugf("Upload video %s", VideoUpdate.Name)
	ctx.Response.SetStatusCode(fasthttp.StatusCreated)
}

func (i *HTTPInstanceAPI) watchVideo(ctx *fasthttp.RequestCtx) {
	ctx.Response.Header.Set("Content-Type", "video/mp4")
	i.log.Debugf("Got request to extract video")
	var video *gridfs.DownloadStream
	var err error

	args := ctx.QueryArgs()
	VideoRequest := db.VideoRequest{
		Filename: string(args.Peek("filename")),
	}
	if video, err = i.api.WatchVideo(VideoRequest); err != nil {
		ctx.Response.SetStatusCode(fasthttp.StatusBadRequest)
		i.log.Errorf("error while getting video %v", err)
		return
	}
	_, _ = io.Copy(ctx.Response.BodyWriter(), video)
}

func (i *HTTPInstanceAPI) uploadVideo(ctx *fasthttp.RequestCtx) {
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

func (i *HTTPInstanceAPI) videos(ctx *fasthttp.RequestCtx) {
	i.log.Debugf("Got request show videos")
	var videoFilter db.VideoFilter
	var videos []db.VideoInfo
	var err error
	var body []byte

	args := ctx.QueryArgs()
	videoFilter = db.VideoFilter{
		Filename: string(args.Peek("filename")),
	}

	if videos, err = i.api.Videos(videoFilter); err != nil {
		ctx.Response.SetStatusCode(fasthttp.StatusBadRequest)
		i.log.Errorf("error getting videos %v", err)
		return
	}

	if body, err = json.Marshal(videos); err != nil {
		ctx.Response.SetStatusCode(fasthttp.StatusBadRequest)
		i.log.Errorf("error marshaling data %v", err)
		return
	}

	ctx.Response.SetStatusCode(fasthttp.StatusOK)
	ctx.Response.SetBody(body)
}

func (i *HTTPInstanceAPI) deleteVideo(ctx *fasthttp.RequestCtx) {
	var videoName db.VideoRequest
	body := ctx.Request.Body()

	if err := json.Unmarshal(body, &videoName); err != nil {
		ctx.Response.SetStatusCode(fasthttp.StatusBadRequest)
		i.log.Errorf("error while unmarshaling %v", err)
		return
	}

	if err := i.api.DeleteVideo(videoName.Filename); err != nil {
		ctx.Response.SetStatusCode(fasthttp.StatusBadRequest)
		i.log.Errorf("error deleting video %v", err)
		return
	}
	i.log.Debugf("Succesfully deleted video %s", videoName.Filename)
	ctx.Response.SetStatusCode(fasthttp.StatusOK)
}
