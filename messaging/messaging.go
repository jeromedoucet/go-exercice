package messaging
import (
	"github.com/jeromedoucet/go-concurrency/drunker/message"
	"encoding/json"
)

type Handler interface {
	HandleMessage(data []byte) error
}


func StartListening() {

}


func AddHandler(handler Handler) {

}

func unmarshall(data []byte) message.Order {
	var order message.Order
	json.Unmarshal(data, &order) // TODO enlever pour l'exercice
	return order
}
