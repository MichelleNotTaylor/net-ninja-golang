package main

import "fmt"

func main() {
	myBill := newBill("Michelle's bill")
	myBill.addItem("onion soup", 4.50)
	myBill.addItem("chicken salad sandwich", 6.99)
	myBill.addItem("summer salad", 5.00)
	myBill.updateTip(10)
	fmt.Println(myBill.format())
}
