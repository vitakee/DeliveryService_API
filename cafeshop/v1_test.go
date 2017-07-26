package cafeshop

import (
	"testing"
	"fmt"
	"context"
	pb "vita.com/grpc"
)
const dbPort=32770

var api *V1

func setup(){
	if api==nil {
		var e error
		api,e = NewV1(dbPort)
		if e!=nil{
			fmt.Println("e = ",e)
		}
	}
}

func TestV1_DBTest(t *testing.T) {
	setup()
	_,e :=api.DBTest("SELECT * FROM test.cafes")
	if e !=nil{
		fmt.Println("failed to select * from cafes e = ",e)
		t.FailNow()
	}
}

func TestV1_Approve(t *testing.T) {
setup()
	_,e := api.Approve(context.Background(),0)
	if e!=nil{
		fmt.Println("Failed to aproove error is :",e)

		t.FailNow()
	}
}

func TestV1_Decline(t *testing.T) {
	setup()
	whynot := &pb.WhyNot{"Beacause,we dont have anymore sousage",1}
	catId := string(whynot.CartId)
	_,e := api.Decline(context.Background(),whynot)
	if e!=nil{
		fmt.Println("Failed to decline cart with id :"+ catId +" error is :",e)

		t.FailNow()
	}

}

func TestV1_ListOrders(t *testing.T) {
setup()
	_,e :=api.ListCarts(context.Background())

	if e !=nil {
		fmt.Println("failed to list carts e = ",e)
		t.FailNow()
	}
}
