package etl

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"
	"strconv"
)

type Product struct {
	PartNumber string
	UnitCost float64
	UnitPrice float64
}

type Order struct {
	CustomerNumber int
	PartNumber int
	Quantity int

	UnitCost float64
	UnitPrice float64
}

func InitETL() {
	start := time.Now()
	extractChan := make(chan *Order)
	transformChan := make(chan *Order)
	doneChannel := make(chan bool)

	go extract(extractChan)
	go transform(extractChan, transformChan)
	go load(transformChan, doneChannel)

	<- doneChannel
	fmt.Println(time.Since(start))
}

func extract(ch chan *Order) {
	//result := []*Order{}
	file, _ := os.Open("./orders.txt")
	//wd, errorWd := os.Getwd()
	//os.OpenFile()
	defer file.Close()
	reader := csv.NewReader(file)
	record, err := reader.Read()

	for record, err = reader.Read(); err == nil; {
		order := new(Order)
		rec0, _ := strconv.Atoi(record[0])
		rec1, _ := strconv.Atoi(record[1])
		rec2, _ := strconv.Atoi(record[2])
		order.CustomerNumber = rec0
		order.PartNumber = rec1
		order.Quantity = rec2
		//result = append(result, order)
		fmt.Println("record: ", record)
		ch <- order
	}

	close(ch)
	//return result
}

//func transform(orders[]*Order) []*Order {
func transform(extractChan, transformChan chan *Order) {
	file, _ := os.Open("./product.txt")
	defer file.Close()
	reader := csv.NewReader(file)
	productList, _ := reader.ReadAll()
	products := make(map[int]*Product)

	for _, record := range productList {
		product := new(Product)
		product.PartNumber = record[0]
		product.UnitCost, _ = strconv.ParseFloat(record[1], 64)
		product.UnitPrice, _ = strconv.ParseFloat(record[2], 64)
		products[product.PartNumber] = product
	}

	//for id, _ := range orders {
	for o := range extractChan {
		time.Sleep(30 * time.Millisecond)
		//o := orders[id]
		o.UnitCost = products[o.PartNumber].UnitCost
		o.UnitPrice = products[o.PartNumber].UnitPrice
		transformChan <- o
	}
	close(transformChan)
}

//func load(orders[]*Order) {
func load(transformChannel chan * Order, doneChannel chan bool) {
	f, _ := os.Create("./dest.txt")
    defer f.Close()

	fmt.Fprintf(f, "%20s%15s%12s%12s%15s%15s\n", "Part number", "Quantity", "Unit Cost",
	"Unit Price", "Total Cost", "Total Price")

	//for _, o := range orders {
	for o := range transformChannel {
		time.Sleep(1 * time.Millisecond)
		fmt.Fprintf(f, "%20s %15d %12.2f %15.2f %15.2f\n",
			o.PartNumber, o.Quantity, o.UnitCost, o.UnitPrice,
			o.UnitCost * float64(o.Quantity),
			o.UnitPrice * float64(o.Quantity))
	}

	doneChannel <- true
}


