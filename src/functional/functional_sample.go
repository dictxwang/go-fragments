package functional

import (
	"crypto/tls"
	"fmt"
	"time"
)

type Server struct {
	Addr     string
	Port     int
	Protocol string
	Timeout  time.Duration
	MaxConns int
	TLS      *tls.Config
}

type Option func(*Server)

func Timeout(timeout time.Duration) Option {
	return func(server *Server) {
		server.Timeout = timeout
	}
}

func Protocol(protocol string) Option {
	return func(server *Server) {
		server.Protocol = protocol
	}
}

func MaxConns(maxconns int) Option {
	return func(server *Server) {
		server.MaxConns = maxconns;
	}
}

func TLS(tls *tls.Config) Option {
	return func(server *Server) {
		server.TLS = tls
	}
}

func NewServer(addr string, port int, options... func(*Server)) (*Server, error) {

	srv := Server{
		Addr: addr,
		Port: port,
		Protocol: "tcp",
		Timeout: 30 * time.Second,
		MaxConns: 1000,
		TLS: nil,
	}

	for _, option := range options {
		option(&srv)
	}

	return &srv, nil
}

func SampleMain() {

	fmt.Println("\n[functional_sample]")

	// 函数式编程方式，优化server创建过程，比构造器模式更优雅
	s1, _ := NewServer("localhost", 1001)
	s2, _ := NewServer("localhost", 1002, Protocol("udp"))
	s3, _ := NewServer("0.0.0.0", 8080, Timeout(1000 * time.Millisecond), MaxConns(100))
	fmt.Println(s1, s2, s3)
}