// Copyright 2017 The toolkit Authors. All rights reserved.
// Use of this source code is governed by a MIT License
// license that can be found in the LICENSE file.

package schedule_test

import (
	"context"
	"testing"
	"time"

	"github.com/donutloop/toolkit/schedule"
)

func TestFIFOSchedule(t *testing.T) {
	s := schedule.NewFIFOScheduler()
	defer s.Stop()

	next := 0
	jobCreator := func(i int) schedule.Job {
		return func(ctx context.Context) {
			if next != i {
				t.Fatalf("job#%d (Actual: %d, Expected: %d)", i, next, i)
			}
			next = i + 1
		}
	}

	var jobs []schedule.Job
	for i := 0; i < 100; i++ {
		jobs = append(jobs, jobCreator(i))
	}

	for _, j := range jobs {
		s.Schedule(j)
	}

	s.WaitFinish(100)
	expectedJobCount := 100
	if s.Scheduled() != expectedJobCount {
		t.Fatalf("scheduled (actual: %d, expected: %d)", s.Scheduled(), expectedJobCount)
	}
}

func TestFIFOScheduleCtx(t *testing.T) {
	s := schedule.NewFIFOScheduler()

	s.Schedule(func(ctx context.Context) {
		<-time.After(250 * time.Millisecond)
		err := ctx.Err()
		if err == nil {
			t.Fatal("unexpected nil error")
		}

		expectedErrorMessage := "context canceled"
		if err.Error() != expectedErrorMessage {
			t.Fatal(err)
		}
	})

	s.Stop()

	expectedJobCount := 1
	if s.Scheduled() != expectedJobCount {
		t.Fatalf("scheduled (actual: %d, expected: %d)", s.Scheduled(), expectedJobCount)
	}
}

func BenchmarkFIFOSchedule(b *testing.B) {
	for n := 0; n < b.N; n++ {
		s := schedule.NewFIFOScheduler()

		jobCreator := func() schedule.Job {
			return func(ctx context.Context) {}
		}

		var jobs []schedule.Job
		for i := 0; i < 100; i++ {
			jobs = append(jobs, jobCreator())
		}

		for _, j := range jobs {
			if err := s.Schedule(j); err != nil {
				b.Fatal(err)
				s.Stop()
			}
		}

		s.WaitFinish(100)

		expectedJobCount := 100
		if s.Scheduled() != expectedJobCount {
			b.Fatalf("scheduled (actual: %d, expected: %d)", s.Scheduled(), expectedJobCount)
			s.Stop()
		}
		s.Stop()
	}
}
