package template

// 待实现逻辑
type BehaviorInterface interface{
	// 其他参数校验
	checkParams(ctx *Context) error

	getPrizesByNode(ctx *Context) error
}