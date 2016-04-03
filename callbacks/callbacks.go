package callbacks

import (
	"math/rand"
	"fmt"
)

type purchaseOrder struct {
	Number int
	Value float64
	Comment string
}

func InitPurchaseCallback()  {
	po := new(purchaseOrder)
	orderChannel := make(chan *purchaseOrder)

	go saveOrder(po, orderChannel)

	cbval := <- orderChannel;
	fmt.Printf("Order type: %v", cbval)

}

func saveOrder(po *purchaseOrder, callback chan *purchaseOrder) {
	po.Number = rand.Intn(5)
	po.Value = rand.Float64()
	po.Comment = "Orange bucket: "
	callback <- po
}

