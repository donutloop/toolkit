package schedule_test

import (
	"context"
	"fmt"

	"github.com/donutloop/toolkit/schedule"
)

func Example() {
	s := schedule.NewFIFOScheduler()
	defer s.Stop()

	job := func(ctx context.Context) {
		fmt.Println("create db entry")
	}

	if err := s.Schedule(job); err != nil {
		fmt.Printf("error: (%v) \n", err)
	}

	s.WaitFinish(1)

	// Output: create db entry
}
