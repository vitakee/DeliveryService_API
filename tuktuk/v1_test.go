package tuktuk

import (
	"testing"
	"fmt"
	"context"
	pb "vita.com/grpc"
)
const dbPort=32770

var api *V1

func setup() {
	if api==nil {
		var e error
		api,e = NewV1(dbPort)

		if e!=nil{
			fmt.Println("e = ",e)
		}

	}
}

func TestV1_OnIt(t *testing.T) {
	setup()
	_,e := api.OnIt(context.Background(),&pb.OnItReq{&pb.Location{123.12,456.45},1})
	if e!=nil{
		fmt.Println("Failed to approve tuktuk , e is :",e)

		t.FailNow()
	}
}
func TestV1_ListOrdersToDeliver(t *testing.T) {
	setup()
	_,e := api.ListOrdersToDeliver(context.Background())
	if e!=nil{
		fmt.Println("Failed to ListOrdersToDeliver , e is :",e)

		t.FailNow()
	}
}
