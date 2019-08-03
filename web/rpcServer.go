package web

import (
	"errors"
	"net/http"
	"net/rpc"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

type Arith int

func (t *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func (t *Arith) Divide(args *Args, quotient *Quotient) error {
	if args.B == 0 {
		return errors.New("divide by zero")
	}
	quotient.Quo = args.A / args.B
	quotient.Rem = args.A % args.B
	return nil
}

func RPCServer() {
	arith := new(Arith)
	err := rpc.Register(arith)
	checkErr(err)
	rpc.HandleHTTP()
	err1 := http.ListenAndServe(":9099", nil)
	checkErr(err1)
}
