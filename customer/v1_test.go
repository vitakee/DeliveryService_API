package customer

import (
	"testing"
	"fmt"
	"context"
	pb "vita.com/grpc"
	"time"
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

func TestV1_OpenCart(t *testing.T) {
	setup()
	atc := &pb.AddToCartReq{1,1}
	_,e :=api.OpenCart(context.Background(),atc)
	if e!=nil{
		fmt.Println("failed to open cat  = ",e)
		t.FailNow()
	}

}
func TestV1_AddToCart(t *testing.T) {
	setup()
	atc := &pb.AddToCartReq{1,3}
	_,e := api.AddToCart(context.Background(),atc)
	if e != nil{
		fmt.Println("AHHH FAIL e = ", e)
		t.FailNow()
	}
	fmt.Println("YEAH ITS WOKING:  ::::::::: e is <nil>")
}
func TestV1_AddToOrder(t *testing.T) {
	setup()
	o := &pb.Order{ProductId: 1, OrderId: 2, Comment: "PLEASE with alot of SPICE", Price: 200,}
	_,e := api.AddToOrder(context.Background(),o)
	if e != nil{
		fmt.Println("Failed to add to ode poduct with id 1 e =",e)
		t.FailNow()
	}
}
func TestV1_RemoveFromOrder(t *testing.T) {
	setup()
	_ , e :=api.RemoveFromOrder(context.Background(),&pb.Order{1,2,"blah blah blah",4,3})
	if e != nil{
		fmt.Println("FAILED TO REMOVE FROM ORDER E = ",e)
		t.FailNow()
	}
}
func TestV1_ListCafes(t *testing.T) {
	setup()
	_ , e := api.ListCafes(context.Background(),&pb.Nil{})
	if e != nil{
		fmt.Println("FAILED to list cafes e = ", e)
		t.FailNow()
	}
}
func TestV1_ListProducts(t *testing.T) {
	setup()
	var tags = []string{"qwe","asd"}
	var location *pb.Location
	location = &pb.Location{123.1,456.2}
	_ , e := api.ListProducts(context.Background(),&pb.Cafe{1,"Olya's Food",tags,location,"https://google.com.ru",true,1})
	if e != nil{
		fmt.Println("FAILED to list Products e = ", e)
		t.FailNow()
	}
}
func TestV1_CloseCart(t *testing.T) {
	setup()
	_,e := api.CloseCart(context.Background(),&pb.Cart{1,true,&pb.Location{123.12,456.45},100,false,false})
	if e != nil{
		fmt.Println("FAILED to Close Cart e = ",e)
		t.FailNow()
	}
}
func TestV1_RemoveFromCart(t *testing.T) {
	setup()
	_,e := api.RemoveFromCart(context.Background(),&pb.Order{1,2,"SUPER DUPER SPICE",400,2})
	if e != nil{
		fmt.Println("FAILED to remove from cart e = ",e)
		t.FailNow()
	}
}
func ()  {
	
}

func TestV1_Test(a *testing.T) {
	t := time.Now()
	fmt.Println("TEST WEEKDAY:", t.Weekday())
	fmt.Println("TEST YEAR:", t.Year())
	fmt.Println("TEST MONTH:", t.Month())
	fmt.Println("TEST DAY:", t.Day())
	fmt.Println("TEST HOUR:", t.Hour())
	fmt.Println("TEST MINUTE:", t.Minute())
	fmt.Println("TEST SECONDS:", t.Second())

	err := api.Test(2)
	fmt.Println("TESTING :",err)
}
