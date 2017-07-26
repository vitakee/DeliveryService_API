package customer

import ("golang.org/x/net/context"
 pb "vita.com/grpc"
	_ "github.com/lib/pq"
	"database/sql"
	"fmt"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/codes"
	"time"
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
         func (v *V1)Test(id int)(error)  {

	row := v.db.QueryRow(`SELECT COUNT(open) as N FROM test.cart WHERE cartid = $1 AND open = true `,id)
	var n int
	err := row.Scan(&n)
	return err
}
         func (v *V1)AddToOrder(ctx context.Context, req *pb.Order) (*pb.Nil, error){

t := time.Now()
			rows, err := v.db.Query("INSERT INTO test.orders (price,comm,productid,created_at,updated_at) VALUES ($1,$2,$3,$4,$5)",req.Price, req.Comment, req.ProductId,t,t)
			if err != nil {
				fmt.Println("Failed to add product to cart:", err)
				return &pb.Nil{}, status.New(codes.Internal, err.Error()).Err()
			}
			defer rows.Close()

		return &pb.Nil{},nil
  }
         func (v *V1)RemoveFromOrder(ctx context.Context, req  *pb.Order) (*pb.Nil, error){

	  rows,err := v.db.Query("DELETE FROM test.orders WHERE id = $1",req.OrderId)
	  if err!=nil{
		  fmt.Println("Failed to Remove product from the order:", err)
		  return &pb.Nil{}, status.New(codes.Internal, err.Error()).Err()
	  }
	  defer rows.Close()
	return &pb.Nil{},nil
  }
         func (v *V1)ListCafes(context.Context, *pb.Nil) (*pb.Cafes, error) {

	rows, err := v.db.Query("SELECT id,name,location FROM test.cafes")
	if err != nil {
		fmt.Println("Failed to get cafes", err)
		return &pb.Cafes{}, status.New(codes.Internal, err.Error()).Err()
	}
	fmt.Println("Cafes are ", rows)
	defer rows.Close()

	var cafes= make([]*pb.Cafe, 0, 0)

	for rows.Next() {
		var id int32
		var  latitude, longitude float32
		var name, link string
		var location *pb.Location
		rows.Scan(&id, &name, &location, &link)
		location = &pb.Location{Latitude: latitude, Longitude: longitude}
		cafes = append(cafes, &pb.Cafe{Id: id, Name: name, Location: location, Link: link})
	}

	return &pb.Cafes{cafes}, nil

}
         func (v *V1)ListProducts(ctx context.Context, cafe *pb.Cafe) (*pb.Products, error){
	// ---------------------------------------PRODUCTS ---------------------------------------------------------------
	// ---------------------------------------PRODUCTS ---------------------------------------------------------------
	// ---------------------------------------PRODUCTS ---------------------------------------------------------------
	// ---------------------------------------PRODUCTS ---------------------------------------------------------------
	// ---------------------------------------PRODUCTS ---------------------------------------------------------------
	rows,err := v.db.Query("SELECT * FROM test.products WHERE cafeid = $1 ",cafe.Id)
	if err != nil{
		fmt.Println("error to get products :",err)
	}
	fmt.Println("Products are ",rows)
	defer rows.Close()
	var products = make ([]*pb.Product,0,0)

	for rows.Next() {
		var id, price int32
		var name,description,link string
		var catId int32
		var cafe int32
		var tags = []string{"food","olya's"}
		var category bool
		rows.Scan(&id,&name,&cafe,&description,&price,&link,&category)

		products = append(products, &pb.Product{name,price,description,cafe,tags,id,link,catId,category})
	}

	return &pb.Products{products},nil
	// ---------------------------------------PRODUCTS ---------------------------------------------------------------
	// ---------------------------------------PRODUCTS ---------------------------------------------------------------
	// ---------------------------------------PRODUCTS ---------------------------------------------------------------
	// ---------------------------------------PRODUCTS ---------------------------------------------------------------
	// ---------------------------------------PRODUCTS ---------------------------------------------------------------
}
         func (v *V1)OpenCart(ctx context.Context,cId *pb.AddToCartReq) (*pb.Nil,error){
			 rows,err := v.db.Query("UPDATE test.cart SET Open ='true' WHERE cartid = $1",cId.CartId)
			 if err != nil {
				 fmt.Println("Failed to open Cart", err)
				 return &pb.Nil{}, status.New(codes.Internal, err.Error()).Err()
			 }
			 defer rows.Close()

			 return &pb.Nil{},nil
}
         func (v *V1)CloseCart(ctx context.Context,cId *pb.Cart) (*pb.Nil,error){
	rows,err := v.db.Query("UPDATE test.cart SET Open ='false' WHERE cartid = $1",cId.Cartid)
	if err != nil {
		fmt.Println("Failed to close Cart", err)
		return &pb.Nil{}, status.New(codes.Internal, err.Error()).Err()
	}
	defer rows.Close()

	return &pb.Nil{},nil
}
         func (v *V1)AddToCart(ctx context.Context,req *pb.AddToCartReq) (*pb.Nil, error){

	row := v.db.QueryRow(`SELECT COUNT(*) as N FROM test.cart WHERE cartid = $1 AND open = true `,req.CartId)
	var n int
	e := row.Scan(&n)
	if e == nil {
		rows, err := v.db.Exec("UPDATE test.orders SET cartid = $1 WHERE id = $2", req.CartId,req.OrderId)
		if err != nil {
			fmt.Println("Failed to add order to cart:", err)
			fmt.Println(rows)
			return &pb.Nil{}, status.New(codes.Internal, err.Error()).Err()
		}
		return &pb.Nil{},nil
	}
	return &pb.Nil{},e
}
         func (v *V1)RemoveFromCart(ctx context.Context, req *pb.Order) (*pb.Nil, error) {

	rows,err := v.db.Exec("UPDATE test.orders SET cartid = null WHERE id=$1",req.CartId)
	if err!=nil{
		fmt.Println("Failed to DELETE order FROM cart:", err)
		fmt.Println(rows)
		return &pb.Nil{}, status.New(codes.Internal, err.Error()).Err()
	}



	return &pb.Nil{},nil
}

	     func (v *V1)Checkout(ctx context.Context,c *pb.Cart) (*pb.Cart, error) {
			_ , e := v.db.Exec("UPDATE test.cart SET approved_by_cafe = false WHERE id = $1",c.Cartid)
			 if e != nil{
			fmt.Println("ERROR : FAILED TO CHECKOUT ERROR MESSAGE IS :::",e)
				 return &pb.Cart{},e
			 }
	     return &pb.Cart{},nil
         }
         func (v *V1)ListCarts(context.Context, *pb.Nil) (*pb.Carts, error) {

			 rows, err := v.db.Query("SELECT * FROM test.cafes")
			 if err != nil {
				 fmt.Println("Failed to get cafes", err)
				 return &pb.Carts{}, status.New(codes.Internal, err.Error()).Err()
			 }
			 fmt.Println("Carts are ", rows)
			 defer rows.Close()

			 var carts = make([]*pb.Cart, 0, 0)

			 for rows.Next() {
				 var id int32
				 var ttioi, open, abc bool
				 var p int32
				 var ptd *pb.Location
				 rows.Scan(&id, &ttioi, &p, &ptd, &open, &abc)
				 carts = append(carts, &pb.Cart{id, open, ptd, p, abc, ttioi})
			 }

			 return &pb.Carts{carts}, nil
		 }
         func (v *V1)ListCategories(context.Context, *pb.Nil) (*pb.Products, error){

			 rows ,e := v.db.Query("SELECT * FROM test.products WHERE category = true")
			 if e != nil {
				 fmt.Println("Failed to get cafes", e)
				 return &pb.Products{}, status.New(codes.Internal, e.Error()).Err()
			 }
			 fmt.Println("Categories are ", rows)
			 defer rows.Close()

			 var products = make([]*pb.Product, 0, 0)

			 for rows.Next() {
				 var id, cafeId, categoryid, price int32
				 var name, description, link string
				 var cat bool
				 var tags []string
				 rows.Scan(&id, &cafeId, &categoryid, &price, &name, &description, &link, &cat, &tags)
				 products = append(products, &pb.Product{name, price, description, cafeId, tags, id,link,categoryid,cat})
			 }


			 return &pb.Products{products},nil
         }
         func (v *V1)ListCategory(c context.Context,cat *pb.Product) (*pb.Products, error){

			 rows ,e := v.db.Query("SELECT * FROM test.products WHERE id = $1",cat.Id)
			 if e != nil {
				 fmt.Println("Failed to get category :", e)
				 return &pb.Products{}, status.New(codes.Internal, e.Error()).Err()
			 }
			 fmt.Println("Products in category are : ", rows)
			 defer rows.Close()

			 var products = make([]*pb.Product, 0, 0)

			 for rows.Next() {
				 var id, cafeId, categoryid, price int32
				 var name, description, link string
				 var cat bool
				 var tags []string
				 rows.Scan(&id, &cafeId, &categoryid, &price, &name, &description, &link, &cat, &tags)
				 products = append(products, &pb.Product{name, price, description, cafeId, tags, id,link,categoryid,cat})
			 }


			 return &pb.Products{products},nil
         }

         func (v *V1)Declined(context.Context, *pb.WhyNot) (*pb.Order, error){
	return &pb.Order{},nil
  }

// DONE 10/11 DONE AddToOrder,RemoveFromOrder,ListCafes,ListProducts,OpenCart,CloseCart,AddToCart,RemoveFromCart,ListCategories,ListCategory,
// UNDONE 1/11 Declined
