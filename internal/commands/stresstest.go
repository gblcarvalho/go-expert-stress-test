package commands

import (
	"fmt"

	"github.com/gblcarvalho/go-expert-stress-test/internal/usecase"
	"github.com/spf13/cobra"
)

type StressTestCMD struct {
	cmd           *cobra.Command
	makeRequestUC *usecase.MakeRequestUC
}

func NewStressTestCMD() *StressTestCMD {
	stCMD := &StressTestCMD{
		makeRequestUC: usecase.NewMakeRequestUC(),
	}

	cmd := &cobra.Command{
		Use:   "go-expert-stress-test",
		Short: "Perform simultaneous requests for stress test",
		Long:  `The program performs multiple simultaneous HTTP requests based on the provided parameters and generates a detailed report with the results.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			url, _ := cmd.Flags().GetString("url")
			requests, _ := cmd.Flags().GetInt("requests")
			concurrency, _ := cmd.Flags().GetInt("concurrency")

			fmt.Println("Executing...")
			report, err := stCMD.makeRequestUC.Execute(url, requests, concurrency)
			if err != nil {
				return err
			}
			stCMD.printReport(report)
			return nil
		},
	}

	cmd.Flags().String("url", "", "URL to send requests to")
	cmd.Flags().Int("requests", 1, "Total number of requests to make")
	cmd.Flags().Int("concurrency", 1, "Number of simultaneous requests")
	cmd.MarkFlagRequired("url")

	stCMD.cmd = cmd
	return stCMD
}

func (st *StressTestCMD) printReport(report *usecase.MakeRequestUCOutput) {
	fmt.Println("*** REPORT ***")
	fmt.Printf("Total time (in seconds) .: %f\n", report.TotalDuration.Seconds())
	fmt.Printf("Total requests ..........: %d\n", report.TotalRequests)
	fmt.Printf("Total Status 200 OK .....: %d\n", report.TotalStatusOk)
	if len(report.TotalStatusCode) > 0 {
		fmt.Printf("Others Status Code ......: \n")
		for code, count := range report.TotalStatusCode {
			fmt.Printf("	%d .: %d\n", code, count)
		}
	} else {
	}
	fmt.Println("*** ------ ***")
}

func (st *StressTestCMD) Execute() error {
	return st.cmd.Execute()
}
