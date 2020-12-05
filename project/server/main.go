package main

import (
	"net"
	"os"
	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc"
	protos "github.com/manureddy7143/sql/project/protos"
	"github.com/manureddy7143/sql/project/protos/server"
	"google.golang.org/grpc/reflection"



)

func main(){

	log := hclog.Default()

	gs := grpc.NewServer()
	cs := server.NewCourse(log)

	protos.RegisterCourseService(gs,cs)
	reflection.Register(gs)
	l,err := net.Listen("tcp",":8080")
	if err !=nil {
		log.Error{"unable to listen","error",err}
		os.Exit(1)
	}
	gs.Serve(l)
}