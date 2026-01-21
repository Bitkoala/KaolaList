package ocr

import (
	"context"
	"io"

	"github.com/go-resty/resty/v2"
	jsoniter "github.com/json-iterator/go"
)

type PaddleEngine struct {
	client *resty.Client
}

func (e *PaddleEngine) Name() string {
	return "paddle"
}

func (e *PaddleEngine) Recognize(ctx context.Context, image io.Reader) (*Result, error) {
	// Paddle Sidecar usually runs at localhost:9898
	api := "http://localhost:9898/ocr/file/json"
	res, err := e.client.R().
		SetContext(ctx).
		SetMultipartField("image", "cap.png", "image/png", image).
		Post(api)
	if err != nil {
		return nil, err
	}

	return &Result{
		Text:   jsoniter.Get(res.Body(), "result").ToString(),
		Status: jsoniter.Get(res.Body(), "status").ToInt(),
		Msg:    jsoniter.Get(res.Body(), "msg").ToString(),
	}, nil
}

func init() {
	Register(&PaddleEngine{
		client: resty.New(),
	})
}
