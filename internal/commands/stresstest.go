package commands

import (
	"fmt"

	"github.com/gblcarvalho/go-expert-stress-test/internal/usecase"
	"github.com/spf13/cobra"
)

type StressTestCMD struct {
	cmd *cobra.Command
	makeRequestUC *usecase.MakeRequestUC
}

func NewStressTestCMD() *StressTestCMD {
	cmd := &cobra.Command{
		Use:   "go-expert-stress-test",
		Short: "Run simultaneous requests for stress test",
	}
	stCMD := &StressTestCMD{
		cmd: cmd,
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

	report, err := st.makeRequestUC.Execute(url, requests, concurrency)
	if err != nil {
		return err
	}

	fmt.Println("*** REPORT ***")
	fmt.Printf("Report: %v \n", report)

	return nil
}
