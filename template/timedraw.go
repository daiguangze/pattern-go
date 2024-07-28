package template

import "fmt"

type TimeDraw struct {
}

// 实现behaviorInterface接口

func (t TimeDraw) checkParams(ctx *Context) error {
	fmt.Println("TimeDraw checkParams invoke!")
	return nil
}

func (t TimeDraw) getPrizesByNode(ctx *Context) error {
	fmt.Println("TimeDraw getPrizesByNode invoke!")
	return nil
}
