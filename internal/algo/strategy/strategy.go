package strategy

import (
	"fmt"

	"github.com/elijahelrod/vespene/internal/algo/signal"
	"github.com/elijahelrod/vespene/pkg/logger"
	"github.com/elijahelrod/vespene/pkg/model"
)

type Strategy struct {
	name    string
	signals []signal.Signal
	pnl     float64
	logger  logger.Logger
}

func newEmptyStrategy(name string, logger logger.Logger) *Strategy {

	if name == "" {
		name = "default"
	}

	return &Strategy{
		name:    name,
		signals: make([]signal.Signal, 1, 10),
		pnl:     0.0,
		logger:  logger,
	}
}

func NewStrategy(name string, logger logger.Logger, signals ...signal.Signal) *Strategy {
	if len(signals) == 0 {
		return newEmptyStrategy(name, logger)
	}

	return &Strategy{
		name:    name,
		signals: signals,
		pnl:     0.0,
		logger:  logger,
	}
}

func (s *Strategy) AddSignal(sig signal.Signal) {
	s.logger.Info("Added %s signal to strategy", sig.Name())
	s.signals = append(s.signals, sig)
}

func (s *Strategy) UpdateSignals(tick model.Tick) {

	for _, sig := range s.signals {

		sig.Update(tick)
		if !sig.SignalActive() {
			s.logger.Info(fmt.Sprintf("Building %s signal", sig.Name()))
		} else {
			s.logger.Info(sig.Details())
		}
	}
}

func (s *Strategy) EvaluateSignals(tick model.Tick) signal.Side {

	buySignals := 0
	sellSignals := 0

	for _, sig := range s.signals {
		if sig.SignalActive() {

			switch sig.Evaluate(tick) {
			case signal.SELL:
				sellSignals += 1
			case signal.BUY:
				buySignals += 1
			}

		}
	}

	if buySignals > sellSignals {
		return signal.BUY
	}

	if sellSignals > buySignals {
		return signal.SELL
	}

	return signal.NONE
}
