package log

import (
	"context"
	"fmt"
	"github.com/go-errors/errors"
	athCtx "github.com/hlhgogo/athena/pkg/context"
	"github.com/sirupsen/logrus"
)

// TraceWithTrace trace增加traceId
func TraceWithTrace(ctx context.Context, format string, args ...interface{}) {
	log.WithFields(getTraceField(ctx)).Tracef(format, args...)
}

// DebugWithTrace debug增加traceId
func DebugWithTrace(ctx context.Context, format string, args ...interface{}) {
	log.WithFields(getTraceField(ctx)).Debugf(format, args...)
}

// InfoWithTrace Info增加traceId
func InfoWithTrace(ctx context.Context, format string, args ...interface{}) {
	log.WithFields(getTraceField(ctx)).Infof(format, args...)
}

// InfoMapWithTrace info增加map信息到日志
func InfoMapWithTrace(ctx context.Context, infos map[string]interface{}, format string, args ...interface{}) {
	fields := getTraceField(ctx)
	for k, v := range infos {
		fields[k] = v
	}
	log.WithFields(fields).Infof(format, args...)
}

// WarnWithTrace warn增加traceId
func WarnWithTrace(ctx context.Context, format string, args ...interface{}) {
	log.WithFields(getTraceField(ctx)).Warnf(format, args...)
}

// WarnMapWithTrace warn增加map信息到日志
func WarnMapWithTrace(ctx context.Context, infos map[string]interface{}, format string, args ...interface{}) {
	fields := getTraceField(ctx)
	for k, v := range infos {
		fields[k] = v
	}
	log.WithFields(fields).Warnf(format, args...)
}

// ErrorWithTrace Error增加traceId
func ErrorWithTrace(ctx context.Context, err error, format string, args ...interface{}) {
	fields := getTraceField(ctx)
	switch err := err.(type) {
	case *errors.Error:
		fields["stack"] = err.ErrorStack()
	default:
		newErr := errors.Wrap(fmt.Sprintf(format, args...), 1)
		fields["stack"] = newErr.ErrorStack()
	}

	log.WithFields(fields).Errorf(format, args...)
}

// ErrorMapWithTrace error增加map信息到日志
func ErrorMapWithTrace(ctx context.Context, infos map[string]interface{}, err error, format string, args ...interface{}) {
	fields := getTraceField(ctx)
	for k, v := range infos {
		fields[k] = v
	}
	switch err := err.(type) {
	case *errors.Error:
		fields["stack"] = err.ErrorStack()
	default:
		newErr := errors.Wrap(fmt.Sprintf(format, args...), 1)
		fields["stack"] = newErr.ErrorStack()
	}
	log.WithFields(fields).Infof(format, args...)
}

// getTraceField 获取loggerField
func getTraceField(ctx context.Context) logrus.Fields {
	fields := logrus.Fields{
		"type": Type,
	}

	requestId := athCtx.GetTraceId(ctx)
	fields["traceId"] = requestId

	return fields
}
