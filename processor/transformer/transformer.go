package transformer

import (
	"github.com/BaritoLog/unclog-worker/filter"
	"github.com/BaritoLog/unclog-worker/processor"
	"github.com/Shopify/sarama"
	"github.com/prometheus/common/log"
)

type (
	Transformer struct {
		handler processor.StreamHandler
	}
)

func NewTransformer() *Transformer {
	return &Transformer{
		handler: func(s *filter.StreamProcessorService, m *sarama.ConsumerMessage) {
			timberWolf, err := filter.ConvertKafkaMessageToTimberWolf(m)
			if err != nil {
				log.Warn(err.Error())
				return
			}

			err = s.SendLogs("unclog_logs_processed", timberWolf)
			if err != nil {
				log.Info(err)
				return
			}
			s.LogTimberWolf(timberWolf)
		},
	}
}
func (t *Transformer) GetHandler() processor.StreamHandler {
	return t.handler
}
func start() {

}
func stop() {

}
