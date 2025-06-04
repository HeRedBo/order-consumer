package consumer

import (
	"github.com/HeRedBo/pkg/mq"
	"github.com/IBM/sarama"
	"order-consumer/global"
)

var orderConsumer *mq.Consumer

func StartOrderConsumer() {
	var err error
	orderConsumer, err = mq.StartKafkaConsumer(global.CONFIG.Kafka.Hosts,
		[]string{global.CONFIG.Kafka.OrderTopic}, "order-consumer",
		nil, OrderMsgHandler)
	if err != nil {
		panic(err)
	}
}

func OrderMsgHandler(msg *sarama.ConsumerMessage) (bool, error) {
	mq.KafkaStdLogger.Println("partion : %d ; offset : %d ; msg %s", msg.Partition, msg.Offset, string(msg.Value))
	//es.GetClient(es.DefaultClient).BulkCreate(global.IndexName)
	return true, nil
}

func CloseOrderConsumer() {
	if orderConsumer != nil {
		orderConsumer.Close()
	}
}
