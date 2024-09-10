package exchangetimes

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
)

type LogMiddleware struct {
	next ExchangeTimesRepository
}

func NewLogMiddleware(next ExchangeTimesRepository) ExchangeTimesRepository {
	return &LogMiddleware{next: next}
}

func (lm *LogMiddleware) GetExchangeTimes(ctx context.Context) (results *ExchangeTimes, err error) {
	defer func(start time.Time) {
		logrus.WithFields(logrus.Fields{
			"took": time.Since(start),
			"err":  err,
		}).Info("GetExchangeTimes")
	}(time.Now())
	results, _ = lm.next.GetExchangeTimes(ctx)
	return
}
