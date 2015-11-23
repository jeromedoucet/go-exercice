// white box test of messaging pkg
package messaging
import (
	"testing"
	"github.com/jeromedoucet/go-concurrency/drunker/message"
	"encoding/json"
)


// test the message.Order unMarshalling func
func TestUnMarshallJson(t *testing.T) {
	// given
	var expectedOrder message.Order
	expectedOrder.Type = message.Whisky
	expectedOrder.Id = 1
	bytes, _ := json.Marshal(&expectedOrder)
	// when
	actualOrder := unmarshall(bytes)
	// then
	if actualOrder.Id != expectedOrder.Id || actualOrder.Type != expectedOrder.Type {
		t.Errorf("the atual order: %v , is not the expected: %v", actualOrder, expectedOrder)
	}
}