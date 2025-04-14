package utils

import (
	"github.com/alibaba/sentinel-golang/api"
	"github.com/alibaba/sentinel-golang/core/flow"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

func InitSentinel() {
	rules := []*flow.Rule{
		{
			Resource:          "POST:/upload",
			ControlBehavior:   flow.Reject,
			StatIntervalInMs:  1000,
			MaxQueueingTimeMs: 100,
			RefResource:       "",
		},
	}

	_, err := flow.LoadRules(rules)
	if err != nil {
		hlog.Fatalf("Failed to load flow rules: %v", err)
	}
	api.InitDefault()
}
