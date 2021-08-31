package logging

import (
	"context"
	"github.com/sirupsen/logrus"
)

var (
	std *PBLogger = &PBLogger{Logger: logrus.StandardLogger()}
)

func WithTraceID(traceID string) *logrus.Entry {
	return std.WithTraceID(traceID)
}

func WithTraceContext(ctx context.Context) *logrus.Entry {
	return std.WithTraceContext(ctx)
}

type PBLogger struct {
	*logrus.Logger
}

func (p *PBLogger) WithTraceID(traceID string) *logrus.Entry {
	return p.WithField("traceID", traceID)
}

func (p *PBLogger) WithTraceContext(ctx context.Context) *logrus.Entry {
	traceID := strext.ExtractTraceID(ctx)
	return p.WithTraceID(traceID)
}
