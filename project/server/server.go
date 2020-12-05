package server
import(
	"github.com/hashicorp/go-hclog"
	protos "github.com/manureddy7143/sql/project/protos"
	"context"
)

type Course struct{
	log hclog.Logger
}
func NewCourse(l hclog.Logger) *Currency {
	return &Course{l}

}  
	


func (c *Course) GetKey(ctx context.Context, rr *protos.Request) (*protos.Response,error) {
 
 return &protos.Response{Value:"available",Count=1,Hour:"7",Repeat=5},nil
}

