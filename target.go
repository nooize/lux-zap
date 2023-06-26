package lux_zap

import (
	"fmt"
	"github.com/nooize/lux"
	"go.uber.org/zap"
)

type zapTarget struct {
	logger zap.Logger
}

func (t zapTarget) Handle(event lux.Event) error {
	fields := make([]zap.Field, 0)
	event.Tags().ForEach(func(key string, val interface{}) {
		fields = append(fields, zap.Any(key, val))
	})
	switch event.Level() {
	case lux.Nop, lux.Trace:
		return nil
	case lux.Debug:
		t.logger.Debug(event.Message(), fields...)
	case lux.Info:
		t.logger.Info(event.Message(), fields...)
	case lux.Warning:
		t.logger.Warn(event.Message(), fields...)
	case lux.Error:
		t.logger.Error(event.Message(), fields...)
	case lux.Fatal:
		t.logger.Fatal(event.Message(), fields...)
	default:
		return fmt.Errorf("unknown log level: %d", event.Level())
	}
	return nil
}
