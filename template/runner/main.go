package main

import (
	"fmt"
	"pattern-go/template"
)

func main() {
	context := &template.Context{
		ActivityInfo: &template.ActivityInfo{
			ActivityType: template.ConstActivityTypeTime,
		},
	}

	switch context.ActivityInfo.ActivityType {
	case template.ConstActivityTypeTime:
		lottery := &template.Lottery{
			ConcreteBehavior: template.TimeDraw{},
		}
		lottery.Run(context)
	default:
		fmt.Println("not implement")

	}

}
