package lux_zap

import (
	"github.com/nooize/lux"
	"go.uber.org/zap"
)

func NewZapTarget(logger zap.Logger) (lux.Target, error) {
	return &zapTarget{
		logger: logger,
	}, nil
}
