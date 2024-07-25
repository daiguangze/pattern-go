package template

import "fmt"


type Lottery struct{
	// 需要实现
	ConcreteBehavior BehaviorInterface
}


func(lottery *Lottery) Run(context *Context) error {
	if err := lottery.checkParams(context);err!=nil{
		return err
	}

	return nil
}

func(lottery *Lottery) checkStatus(context *Context) error{
	fmt.Println("common check status ~~")
	return nil
}

func(lottery *Lottery) deductStock(context *Context) error{
	fmt.Println("common deduct Stock ~~")
	return nil
}

func(lottery *Lottery) checkParams(context *Context) error{
	return lottery.ConcreteBehavior.checkParams(context)
}


func(lottery *Lottery) getPrizesByNode(context *Context) error{
	return lottery.getPrizesByNode(context)
}