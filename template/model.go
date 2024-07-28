package template

const (
	ConstActivityTypeTime int32 = 1

	ConstActivityTypeTimes int32 = 2

	ConstActivityTypeAmount int32 = 3
)

type ActivityInfo struct {
	// 活动类型
	ActivityType int32
}

type Context struct {
	ActivityInfo *ActivityInfo
}
