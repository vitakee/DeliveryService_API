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

const (
	DB_USER     = "postgres"
	DB_PASSWORD = "postgres"
	DB_NAME     = "postgres"
	DB_PORT		= 32768
)
func NewV1()(*V1,error){
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s port=%v sslmode=disable",
		DB_USER, DB_PASSWORD, DB_NAME, DB_PORT)

	db,err := sql.Open("postgres", dbinfo)
	if err != nil {
		fmt.Println("ERROR: ",err)
		return nil,err
	}
	if err=db.Ping();err!=nil{
		fmt.Println("Connect to db failed",err)
		return nil,err
	}

	return &V1{db:db},nil
	return &V1{},nil
}

/*DONE*/ func (v *V1)OnIt(ctx context.Context, oir *pb.OnItReq) (*pb.AprooveAnswer, error){


       rows,err := v.db.Query("UPDATE test.orders SET tuktuk_is_on_it = 'true' WHERE id = %v",oir.OrderId)
	if err!=nil {
		fmt.Println("cannot  set on it true:",err)
	}
	fmt.Println("",rows)
	aproove_answer := true
	return &pb.AprooveAnswer{aproove_answer},nil
}
/*DONE*/ func (v *V1)ListOrdersToDeliver(context.Context, *pb.Location) (*pb.OrdersToDeliver, error){
	rows,err := v.db.Query("SELECT FROM test.orders WHERE aprooved_by_cafe = 'true'")
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