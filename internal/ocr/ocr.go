package ocr

import (
	"context"
	"io"
)

type Result struct {
	Text   string `json:"text"`
	Status int    `json:"status"` // 200 for success
	Msg    string `json:"msg"`
}

type Engine interface {
	Name() string
	Recognize(ctx context.Context, image io.Reader) (*Result, error)
}

var engines = make(map[string]Engine)
var engineGetter func() string

func Register(engine Engine) {
	engines[engine.Name()] = engine
}

func Get(name string) Engine {
	if e, ok := engines[name]; ok {
		return e
	}
	return engines["remote"]
}

func SetEngineGetter(getter func() string) {
	engineGetter = getter
}

func Recognize(ctx context.Context, image io.Reader) (*Result, error) {
	engineName := "remote"
	if engineGetter != nil {
		engineName = engineGetter()
	}
	return Get(engineName).Recognize(ctx, image)
}
