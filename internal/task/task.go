package task

import (
	"heatpump/internal/cli"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var collectInterval time.Duration

func init() {
	collectInterval = cli.Options().CollectInterval
}

func Start() {
	collectTicker := schedule(collect, collectInterval)
	replayTicker := schedule(replay, time.Minute*2)

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs

	collectTicker.Stop()
	replayTicker.Stop()
}

func schedule(f func(), interval time.Duration) *time.Ticker {
	ticker := time.NewTicker(interval)
	go func() {
		for range ticker.C {
			f()
		}
	}()
	return ticker
}
