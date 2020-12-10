package server

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/hashicorp/go-hclog"
	protos "github.com/manureddy7143/sql/project/protos"
)

// Course is
type Course struct {
	log hclog.Logger
}
var(
	name string 
	value string 
	count int32
	hour string
	repeat int32
	id int
	date string
	result string
)

var db *sql.DB
var err error




// NewCourse is
func NewCourse(l hclog.Logger) *Course {
	return &Course{l}
}
var a string


// "Getcourse is a"
func (c *Course) Getcourse(ctx context.Context, rr *protos.Request) (*protos.Response,error) {
	/*c.log.Info( "Handle Getcourse","key", rr.GetKey())*/
	a := rr.GetKey()
	b, _, _ := time.Now().Clock()
	date = time.Now().Format("01-02-2006")
	fmt.Println("a,b values",a,b,date)
	
	db ,err := sql.Open("mysql", "root:143114@tcp(127.0.0.1:3306)/courser")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	err = db.QueryRow("SELECT name from courses where name = ? ", a ).Scan(&result)
	fmt.Println("Result values",result)
	if result == a {
		var dhour int
		dcounter := 0
		repeat := 0
		err := db.QueryRow("select time from time").Scan(&dhour)
		err = db.QueryRow("select repeated from details where name=? ORDER BY id DESC",a).Scan(&repeat)

		err = db.QueryRow("select count  from details where hour= ? AND name= ? AND date= ?",dhour,a,date).Scan(&dcounter,)
		fmt.Println("dhour,dcounter values",dhour,dcounter)
		fmt.Println("error",err)

		if (b == dhour && err != nil) {
			c:=1
			repeat++
			_,err=db.Query("INSERT INTO details (name,count,hour,repeated,value,date) values(?,?,?,?,'available',?)" ,a,c,b,repeat,date)
			fmt.Println("c,repeat,err",c,repeat,err)

		} else if (b == dhour && err == nil){
			dcounter++
			_,err=db.Query("UPDATE details SET count = ? where hour= ? AND name= ? ",dcounter,dhour,a)
			fmt.Println("database updated",err,dcounter)

		} else {
            fmt.Println("finalelse")
			_,err=db.Query("TRUNCATE TABLE time")
			_,err = db.Query("INSERT INTO time values(?)",b)
		}

	}else{
		fmt.Println("the course you requseted is not in database")
	}
	
	


	err = db.QueryRow("SELECT * from details where name = ? ORDER BY id DESC", a ).Scan( &name, &value, &count, &hour,&repeat,&id,&date)
	if err != nil {
		panic(err.Error())
	  
		}
	  
	return &protos.Response{Value:value,Count:count,Hour:hour,Repeat:repeat}, nil
}

