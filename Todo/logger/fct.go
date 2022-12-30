package logger

import (
	"context"

	"github.com/soreing/trex"
	"go.uber.org/zap"
)

const (
	TRACE_INFO_KEY = "tinfo"
)

func NewLoggerFactory() (*LoggerFactory, error) {
	lgr, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}
	return &LoggerFactory{
		lgr: lgr,
	}, nil
}

type LoggerFactory struct {
	lgr *zap.Logger
}

func (lf *LoggerFactory) NewLogger(
	ctx context.Context,
) *zap.Logger {
	if ctx == nil {
		return lf.lgr
	}
	raw := ctx.Value(TRACE_INFO_KEY)
	if raw == nil {
		return lf.lgr
	}
	trace, ok := raw.(trex.TxModel)
	if !ok {
		return lf.lgr
	}
	return lf.lgr.With(
		zap.String("tid", trace.Tid),
		zap.String("pid", trace.Pid),
		zap.String("rid", trace.Rid),
	)
}
