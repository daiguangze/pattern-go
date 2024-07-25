package template


type ActivityInfo struct{
	// 活动类型
	ActivityType int32
}

type Context struct{
	ActivityInfo *ActivityInfo
}