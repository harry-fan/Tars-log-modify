package main

import (
	"context"
	"fmt"
)

// SayHelloImp servant implementation
type SayHelloImp struct {
}

// Init servant init
func (imp *SayHelloImp) Init() error {
	//initialize servant here:
	//...
	return nil
}

// Destroy servant destroy
func (imp *SayHelloImp) Destroy() {
	//destroy servant here:
	//...
}

func (imp *SayHelloImp) Add(ctx context.Context, a int32, b int32, c *int32) (int32, error) {
	*c = a + b
	fmt.Println("req calc add a:", a, "|b:", b)
	return 0, nil
}
func (imp *SayHelloImp) Sub(ctx context.Context, a int32, b int32, c *int32) (int32, error) {
	*c = b - a
	return 0, nil
}
func (imp *SayHelloImp) ShowGoReq(ctx context.Context, input string, output *string) (int32, error) {
	*output = "hello " + input;
	fmt.Println("recv msg:", input)
	return 0, nil
}
