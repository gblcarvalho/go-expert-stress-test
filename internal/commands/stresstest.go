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
	cmd := &cobra.Command{
		Use:   "go-expert-stress-test --url <URL> --requests <total> --concurrency <concurrency>",
		Short: "Perform simultaneous requests for stress test",
		Long:  `The program performs multiple simultaneous HTTP requests based on the provided parameters and generates a detailed report with the results.`,
	}
	stCMD := &StressTestCMD{
		cmd:           cmd,
		makeRequestUC: usecase.NewMakeRequestUC(),
	}
	stCMD.initCMD()
	return stCMD
}

func (st *StressTestCMD) initCMD() {
	st.cmd.Flags().String("url", "", "uso")
	st.cmd.MarkFlagRequired("url")
	st.cmd.Flags().Int("requests", 1, "uso")
	st.cmd.Flags().Int("concurrency", 1, "uso")
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
	err := st.cmd.Execute()
	if err != nil {
		return err
	}
	url, err := st.cmd.Flags().GetString("url")
	if err != nil {
		return err
	}
	requests, err := st.cmd.Flags().GetInt("requests")
	if err != nil {
		return err
	}
	concurrency, err := st.cmd.Flags().GetInt("concurrency")
	if err != nil {
		return err
	}

	fmt.Printf("Executing ....")
	report, err := st.makeRequestUC.Execute(url, requests, concurrency)
	if err != nil {
		return err
	}

	st.printReport(report)
	return nil
}
