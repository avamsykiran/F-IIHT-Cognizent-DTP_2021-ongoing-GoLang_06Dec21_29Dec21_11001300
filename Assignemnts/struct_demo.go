package main

import "fmt"

type bankAccount struct {
	accNum int
	bal    float32
	branch string
	ifsc   string
}

type address struct {
	doorNum string
	street  string
	city    string
}

type vendor struct {
	vid         int
	title       string
	account     bankAccount
	contactAddr address
}
type customer struct {
	cid          int
	title        string
	deliveryAddr address
}

func main() {
	v1 := vendor{
		vid:   1,
		title: "vendor1",
		account: bankAccount{
			accNum: 1234,
			bal:    5678.0,
			branch: "DWK",
			ifsc:   "SBI00090",
		},
		contactAddr: address{
			"4/1",
			"street1",
			"cbe",
		},
	}
	c1 := customer{
		cid:   45,
		title: "customer45",
		deliveryAddr: address{
			"67-2",
			"dummystreet",
			"chn",
		},
	}
	fmt.Println(v1)
	fmt.Println(c1)
}
