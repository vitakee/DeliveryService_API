package cafeshop

import ("golang.org/x/net/context"
	pb "vita.com/grpc"
	_ "github.com/lib/pq"
	"database/sql"
	"fmt"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/codes"
)

type V1 struct {
	db *sql.DB
}

func (v *V1)DBTest(query string)(*sql.Rows ,error)  {
	rows,e := v.db.Query(query)


	if e != nil {
		fmt.Println("Failed to blah blah blah:", e)
		return nil ,e
	}

	return rows,nil

}
func NewV1(dbPort int)(*V1,error){


	var dbConnString=fmt.Sprintf("postgres://postgres@localhost:%v?sslmode=disable",dbPort)
	var db,e = sql.Open("postgres",dbConnString)

	if e!=nil{
		return nil,e
	}
	if e=db.Ping();e!=nil{
		return nil,e
	}

	return &V1{db:db},e
}



/*DONE*/ func (v *V1) Approve(ctx context.Context, cartID int) (*pb.ApproveAnswer, error){


          _,e := v.db.Exec("UPDATE test.cart SET approved_by_cafe = true WHERE cartId = $1",cartID)

	if e != nil {
		fmt.Println("Failed to approve cart:", e)
		return &pb.ApproveAnswer{}, status.New(codes.Internal, e.Error()).Err()
	}

       var approved_answer bool
	approved_answer = true
	return &pb.ApproveAnswer{ approved_answer},nil


}
/*DONE*/ func (v *V1)Decline(ctx context.Context, wn *pb.WhyNot) (*pb.WhyNot, error){

	_,err := v.db.Exec("UPDATE test.cart SET approved_by_cafe = false WHERE cartId = $1", wn.CartId)
	if err != nil {
		fmt.Println("Failed to decline order:", err)
		return &pb.WhyNot{}, status.New(codes.Internal, err.Error()).Err()
	}

	return &pb.WhyNot{wn.WhyNot,wn.CartId},nil
}
/*DONE*/func (v *V1)ListCarts(context.Context) (*pb.Carts, error){
	rows,err := v.db.Query("SELECT * FROM test.cart WHERE approved_by_cafe = false ")
	if err!=nil{
		return &pb.Carts{}, status.New(codes.Internal, err.Error()).Err()
	}
	fmt.Println("Carts are ",rows)
	defer rows.Close()
	var carts = make ([]*pb.Cart,0,0)

	for rows.Next() {
		var products []*pb.Product
		var orderId,price int32
		var place_to_deliver *pb.Location

		rows.Scan(&orderId,&products,&place_to_deliver)
		carts = append(carts, &pb.Cart{Price:price,Cartid:orderId,PlaceToDeliver:place_to_deliver})
	}

	return &pb.Carts{carts},nil
}
