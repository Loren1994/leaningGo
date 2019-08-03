package web

import (
	"fmt"
	"net/rpc"
)

func RPCClient() {
	client, err := rpc.DialHTTP("tcp", ":9099")
	checkErr(err)
	args := Args{17, 8}
	var reply int
	err = client.Call("Arith.Multiply", args, &reply)
	checkErr(err)
	fmt.Printf("Arith: %d*%d=%d\n", args.A, args.B, reply)
}
