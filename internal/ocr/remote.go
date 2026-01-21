package ocr

import (
	"context"
	"io"

	"github.com/OpenListTeam/OpenList/v4/internal/conf"
	"github.com/OpenListTeam/OpenList/v4/internal/setting"
	"github.com/go-resty/resty/v2"
	jsoniter "github.com/json-iterator/go"
)

type RemoteEngine struct {
	client *resty.Client
}

func (e *RemoteEngine) Name() string {
	return "remote"
}

func (e *RemoteEngine) Recognize(ctx context.Context, image io.Reader) (*Result, error) {
	api := setting.GetStr(conf.OcrApi)
	res, err := e.client.R().
		SetContext(ctx).
		SetMultipartField("image", "cap.png", "image/png", image).
		Post(api)
	if err != nil {
		return nil, err
	}

	text := jsoniter.Get(res.Body(), "result").ToString()
	status := jsoniter.Get(res.Body(), "status").ToInt()
	msg := jsoniter.Get(res.Body(), "msg").ToString()

	if status == 0 && text != "" {
		status = 200 // Normalize success
	}

	return &Result{
		Text:   text,
		Status: status,
		Msg:    msg,
	}, nil
}

func init() {
	Register(&RemoteEngine{
		client: resty.New(),
	})
}
