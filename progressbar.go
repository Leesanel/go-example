package main

import (
	"fmt"
	"github.com/k0kubun/go-ansi"
	"time"

	"github.com/schollz/progressbar"
)

func main() {
	doneCh := make(chan struct{})
	desc := ""

	bar := progressbar.NewOptions(100,
		progressbar.OptionEnableColorCodes(true),

		progressbar.OptionSetDescription(desc),
		progressbar.OptionSetWriter(ansi.NewAnsiStdout()),
		progressbar.OptionSetWidth(10),
		progressbar.OptionThrottle(65*time.Millisecond),
		progressbar.OptionShowCount(),
		progressbar.OptionShowIts(),
		progressbar.OptionSpinnerType(14),
		progressbar.OptionFullWidth(),
		//progressbar.OptionSetRenderBlankState(true),
		progressbar.OptionSetDescription("load..."),
		progressbar.OptionSetTheme(progressbar.Theme{
			Saucer: 		"[blue]█[reset]",
			SaucerHead:    	"[blue]^[reset]",
			SaucerPadding: 	"-",
			BarStart: 		"||",
			BarEnd: 		"||",
		}),

		progressbar.OptionOnCompletion(func() {
			doneCh <- struct{}{}
		}),
	)

	go func() {
		for i := 0; i < 100; i++ {
			_ = bar.Add(1)
			time.Sleep(20 * time.Millisecond)
			bar.RenderBlank()
		}
	}()

	// got notified that progress bar is complete.
	<-doneCh
	fmt.Println("\n ======= progress bar completed ==========\n")

}
