package schedule_test

import (
	"github.com/donutloop/toolkit/schedule"
	"context"
	"fmt"
)

func ExampleFIFOScheduler() {

	s := schedule.NewFIFOScheduler()
	defer s.Stop()

	job := func(ctx context.Context) {
		fmt.Println("create db entry")
	}

	if err := s.Schedule(job); err != nil {
		fmt.Println(fmt.Sprintf("error: (%v)", err))
	}

	s.WaitFinish(1)

	// Output: create db entry
}
