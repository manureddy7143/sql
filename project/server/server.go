package server
import(
	"github.com/hashicorp/go-hclog"
	protos "github.com/manureddy7143/sql/project/protos"
	"context"
)

type Course struct{
	log hclog.Logger
}
func NewCourse(l hclog.Logger) *Course {
	return &Course{l}

}  
	


func (c *Course) Getcourse(ctx context.Context, rr *protos.Request) (*protos.Response,error) {
	c.log.Info( "Handle Getcourse","key", rr.GetKey())
/* fgg*/
 
 return &protos.Response{Value:"available",Count:1,Hour:"7",Repeat:5}, nil
}

