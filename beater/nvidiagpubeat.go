package beater

import (
	"fmt"
	"time"

	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/logp"
	"github.com/yoonlee95/nvidiagpubeat/nvidia"

	"github.com/yoonlee95/nvidiagpubeat/config"
)

type Nvidiagpubeat struct {
	done   chan struct{}
	config config.Config
	client beat.Client
}

// Creates beater
func New(b *beat.Beat, cfg *common.Config) (beat.Beater, error) {
	c := config.DefaultConfig
	if err := cfg.Unpack(&c); err != nil {
		return nil, fmt.Errorf("Error reading config file: %v", err)
	}

	bt := &Nvidiagpubeat{
		done:   make(chan struct{}),
		config: c,
	}
	return bt, nil
}

func (bt *Nvidiagpubeat) Run(b *beat.Beat) error {
	logp.Info("nvidiagpubeat is running! Hit CTRL-C to stop it.")

	var err error
	bt.client, err = b.Publisher.Connect()
	if err != nil {
		return err
	}

	ticker := time.NewTicker(bt.config.Period)
	for {
		select {
		case <-bt.done:
			return nil
		case <-ticker.C:
		}
		metrics := nvidia.NewMetrics()
		gpuMetrics, err := metrics.Get(bt.config.Env, bt.config.Query)
		if err != nil {
			logp.Err("Event not generated, error: %s", err.Error())
		} else {
			logp.Info("Event generated, Attempting to publish to configured output.")
			for _, gpuMetric := range gpuMetrics {

				event := beat.Event{
					Timestamp: time.Now(),
					Fields:    gpuMetric,
				}
				bt.client.Publish(event)
				logp.Info("Event sent")
			}
		}
	}
}

func (bt *Nvidiagpubeat) Stop() {
	bt.client.Close()
	close(bt.done)
}
