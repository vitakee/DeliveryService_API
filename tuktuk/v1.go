package tuktuk

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

/*DONE*/ func (v *V1)OnIt(ctx context.Context, oir *pb.OnItReq) (*pb.ApproveAnswer, error){


       rows,err := v.db.Query("UPDATE test.cart SET tuktuk_is_on_it = 'true' WHERE cartid = $1",oir.OrderId)
	if err!=nil {
		fmt.Println("cannot  set on it true:",err)
	}
	fmt.Println("",rows)
	approve_answer := true
	return &pb.ApproveAnswer{approve_answer},err
}
/*DONE*/ func (v *V1)ListOrdersToDeliver(context.Context) (*pb.OrdersToDeliver, error){
	rows,err := v.db.Query("SELECT FROM test.cart WHERE approved_by_cafe = 'true'")
	if err != nil {
		fmt.Println("Failed to get orders to deliver", err)
		return &pb.OrdersToDeliver{}, status.New(codes.Internal, err.Error()).Err()
	}
	fmt.Println("Orders to deliver are ", rows)
	defer rows.Close()

	var ordersToDeliver= make([]*pb.OrderToDeliver, 0, 0)

	for rows.Next() {
		var id  int32
	 	var order *pb.Order
	    var aprooved_by_cafe bool
		rows.Scan(&id, &order, &aprooved_by_cafe)

		ordersToDeliver = append(ordersToDeliver, &pb.OrderToDeliver{Id: id,AproovedByCafe: aprooved_by_cafe ,Order:order})
	}
	return &pb.OrdersToDeliver{ordersToDeliver},nil
}

// DONE 2/2 DONE OnIt,ListOrdersToDeliver