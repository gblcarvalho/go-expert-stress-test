package usecase

import (
	"fmt"
	"time"

	"github.com/gblcarvalho/go-expert-stress-test/internal/utils"
)

type MakeRequestUC struct {
}

type MakeRequestUCOutput struct {
	TotalDuration time.Duration
	TotalRequests int
	TotalSuccessRequests int
	//TODO: OtherCodesDists
}

func NewMakeRequestUC() *MakeRequestUC {
	return &MakeRequestUC{}
}

func (uc *MakeRequestUC) Execute(url string, requests int, concurrency int) (*MakeRequestUCOutput, error) {
	if err := utils.AssertNotEmpty(url, "URL cannot be empty"); err != nil {
		return nil, err
	}
	if err := utils.AssertPositive(requests, "number of requests need to be positive"); err != nil {
		return nil, err
	}
	if err := utils.AssertPositive(requests, "number of concurrency need to be positive"); err != nil {
		return nil, err
	}

	fmt.Println("*** Executando UC ***")
	fmt.Printf("URL ..........: %s \n", url)
	fmt.Printf("Requests .....: %d \n", requests)
	fmt.Printf("Concurrency ..: %d \n", concurrency)

	return &MakeRequestUCOutput{}, nil
}
