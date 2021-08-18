package main

import "github.com/hlhgogo/athena/gateway"

func main() {
	svr := gateway.NewGateWay(
		gateway.Name("Athena"),
	)

	svr.Init()

	if err := svr.Run(); err != nil {
		panic(err)
	}
}
