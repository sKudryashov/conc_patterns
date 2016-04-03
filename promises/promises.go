package promises

import (
	"fmt"
	"errors"
	"time"
)

type Promise struct {
	successChannel chan interface{}
	failureChannel chan error
}

type PurchaseOrder struct {
	Number int
	Value float64
}

func InitOrders()  {
	po := new(PurchaseOrder)
	po.Number = 1
	po.Value = 42.54

	saveOrder(po, false).Then(func(obj interface{}) error {
		po := obj.(*PurchaseOrder)
		fmt.Printf("Order saved with id: %d \n", po.Number)

		return nil
		///return errors.New("First promise failed")
	}, func(err error) {
		fmt.Printf("Failed to save order:" + err.Error() + "\n")
	}).Then(func(obj interface{}) error {
		fmt.Println("Second promise success")

		return nil
	}, func(err error) {
		fmt.Println("Second promise failed" + err.Error())
	})

	fmt.Scanln()
}

func saveOrder(po *PurchaseOrder, isFailed bool) *Promise {
	result := new(Promise)
	result.successChannel = make(chan interface{}, 1)
	result.failureChannel = make(chan error)

	go func() {
		time.Sleep(2 * time.Second)
		if shouldFail := isFailed; shouldFail {
			result.failureChannel <- errors.New("Failed to save an order")
		} else {
			po.Number = 123;
			result.successChannel <- po
		}
	}()

	return result
}

func (this *Promise) Then (success func(interface{}) error, failure func(error)) *Promise {
	result := new(Promise)
	result.successChannel = make(chan interface{}, 1)
	result.failureChannel = make(chan error, 1)

	timeout := time.After(1 * time.Second)

	go func() {
		select {
		case obj := <- this.successChannel:
			newErr := success(obj)
			if newErr == nil {
				result.successChannel <- obj
			} else {
				result.failureChannel <- newErr
			}
		case errObj := <- this.failureChannel:
			failure(errObj)
			result.failureChannel <- errObj
		case <- timeout:
			failure(errors.New("Promise timed out"))
		}
	}()

	return result
}