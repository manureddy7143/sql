package server
import(
	"github.com/hashicorp/go-hclog"
	


)

type Course struct{
	log hclog.Logger
}

func (c*Course) GetKey(ctx context.context, *Request) (*proto.Response,error) {

}
