package consumer

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/HeRedBo/pkg/es"
	"github.com/HeRedBo/pkg/mq"
	"github.com/HeRedBo/pkg/strutil"
	"github.com/IBM/sarama"
	"order-consumer/global"
)

var orderConsumer *mq.Consumer

func StartOrderConsumer() {
	var err error
	orderConsumer, err = mq.StartKafkaConsumer(global.CONFIG.Kafka.Hosts, []string{global.CONFIG.Kafka.OrderTopic}, "order-consumer", nil, OrderMsgHandler)
	if err != nil {
		panic(fmt.Sprintf("err %v,host %v", err, global.CONFIG.Kafka.Hosts))
	}
}

func OrderMsgHandler(msg *sarama.ConsumerMessage) (bool, error) {
	mq.KafkaStdLogger.Printf("partion: %d ; offset : %d; msg : %s",
		msg.Partition, msg.Offset, string(msg.Value))
	orderMsg := OrderMsg{}
	err := json.Unmarshal(msg.Value, &orderMsg)
	if err != nil {
		//格式异常的数据，回到队列也不会解析成功
		global.LOG.Error("Unmarshal error", err, string(msg.Value))
		return true, nil
	}
	mq.KafkaStdLogger.Printf("product: %+v", orderMsg)
	orderIndex := orderMsg.OrderIndex
	orderIndex.OrderStatus = orderMsg.Status
	names := make([]string, 0)
	productIDs := make([]int64, 0)
	orderSuffixLen := 4

	if len(orderMsg.OrderId) >= orderSuffixLen {
		orderIndex.OrderIdSuffix = orderMsg.OrderId[len(orderMsg.OrderId)-4:]
	}

	for _, cart := range orderMsg.CartInfo {
		names = append(names, cart.ProductInfo.StoreName)
		productIDs = append(productIDs, cart.ProductInfo.Id)
	}
	orderIndex.Names = names
	orderIndex.ProductIds = productIDs

	esClient := es.GetClient(es.DefaultClient)
	switch orderMsg.Operation {
	case global.OperationCreate, global.OperationUpdate:
		esClient.BulkCreate(global.IndexName, orderIndex.OrderId,
			strutil.Int64ToString(orderIndex.Uid), orderIndex)

	case global.OperationDelete:
		err := esClient.Delete(context.Background(), global.IndexName, orderIndex.OrderId, strutil.Int64ToString(orderIndex.Uid))
		if err != nil {
			global.LOG.Error("DeleteRefresh error", err, "id", orderIndex.OrderId)
		}
	}

	return true, nil
}

func CloseOrderConsumer() {
	if orderConsumer != nil {
		orderConsumer.Close()
	}
}
