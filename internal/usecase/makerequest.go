package usecase

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/gblcarvalho/go-expert-stress-test/internal/utils"
)

type MakeRequestUC struct {
}

type MakeRequestUCOutput struct {
	TotalDuration   time.Duration
	TotalRequests   int
	TotalStatusOk   int
	TotalStatusCode map[int]int
}

type RequestStats struct {
	mu              sync.Mutex
	initAt          time.Time
	endAt           time.Time
	total           int
	statusOkCounter int
	statusCounter   map[int]int
}

func newRequestStats() *RequestStats {
	return &RequestStats{
		statusCounter: make(map[int]int),
	}
}

func (s *RequestStats) Init() {
	s.initAt = time.Now()
}

func (s *RequestStats) Finish() {
	s.endAt = time.Now()
}

func (s *RequestStats) Duration() time.Duration {
	return s.endAt.Sub(s.initAt)
}

func (s *RequestStats) Update(status int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.total++
	if status == 200 {
		s.statusOkCounter++
		return
	}
	s.statusCounter[status]++
}

func NewMakeRequestUC() *MakeRequestUC {
	return &MakeRequestUC{}
}

func (uc *MakeRequestUC) Execute(url string, nRequests int, nConcurr int) (*MakeRequestUCOutput, error) {
	if err := uc.validateInput(url, nRequests, nConcurr); err != nil {
		return nil, err
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	listenSignal(cancel)
	stats := makeRequests(ctx, url, nRequests, nConcurr)
	return &MakeRequestUCOutput{
		TotalDuration:   stats.Duration(),
		TotalRequests:   stats.total,
		TotalStatusOk:   stats.statusOkCounter,
		TotalStatusCode: stats.statusCounter,
	}, nil
}

func (uc *MakeRequestUC) validateInput(url string, nRequests int, nConcurr int) error {
	if err := utils.AssertNotEmpty(url, "URL cannot be empty"); err != nil {
		return err
	}
	if err := utils.AssertPositive(nRequests, "number of requests need to be positive"); err != nil {
		return err
	}
	if err := utils.AssertPositive(nConcurr, "number of concurrency need to be positive"); err != nil {
		return err
	}
	return nil
}

func listenSignal(cancel func()) {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sig
		cancel()
	}()
}

func makeRequests(ctx context.Context, url string, nRequests int, nConcurr int) *RequestStats {
	jobs := make(chan int)
	var wg sync.WaitGroup
	stats := newRequestStats()

	for i := 0; i < nConcurr; i++ {
		wg.Add(1)
		go requestWorker(ctx, &wg, jobs, url, stats)
	}

	stats.Init()
	go func() {
		defer close(jobs)
		for i := 0; i < nRequests; i++ {
			select {
			case <-ctx.Done():
				//Canceling received
				return
			case jobs <- i:
			}
		}
	}()
	wg.Wait()
	stats.Finish()
	return stats
}

func requestWorker(ctx context.Context, wg *sync.WaitGroup, jobs <-chan int, url string, stats *RequestStats) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		case _, ok := <-jobs:
			if !ok {
				return
			}
			req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
			if err != nil {
				fmt.Printf("create request error: %v \n", err)
				continue
			}

			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				fmt.Printf("send request error: %v \n", err)
				continue
			}

			stats.Update(resp.StatusCode)
		}
	}
}
