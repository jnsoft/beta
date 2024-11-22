package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"

	"github.com/jnsoft/beta/util/httputil"

	"github.com/spf13/cobra"
)

var resultType string
var repitions string
var parallel string

//var resultType string
//var resultType string
//var resultType string
//var resultType string

func httpCmd() *cobra.Command {
	var httpCmd = &cobra.Command{
		Use:   "http",
		Short: "Http helper.",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return IncorrectUsageErr()
		},
		Run: func(cmd *cobra.Command, args []string) {
		},
	}

	httpCmd.AddCommand(getCmd())

	return httpCmd
}

func getCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "get [url]",
		Short: "Get http request",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			url := args[0]

			reps, err := strconv.Atoi(repitions)
			if err != nil {
				fmt.Printf("Error converting repitions to int: %v\n", err)
				os.Exit(1)
			}

			pars, err := strconv.Atoi(parallel)
			if err != nil {
				fmt.Printf("Error converting repitions to int: %v\n", err)
				os.Exit(1)
			}

			var wg sync.WaitGroup
			sem := make(chan struct{}, pars) // Buffered channel to limit concurrency

			for i := 0; i < reps; i++ {
				wg.Add(1)
				sem <- struct{}{} // Acquire a slot

				go func(i int) {
					defer wg.Done()
					defer func() { <-sem }() // Release the slot

					var res string
					var err error
					var code int
					var duration int64

					switch resultType {
					case "string":
						res, code, err = httputil.GetString(url, proxyUrl)
					case "json":
						res, code, err = httputil.GetJSON(url, proxyUrl)
					case "time":
						_, code, duration, err = httputil.MeasureTime(httputil.GetString, url, proxyUrl)
						res = fmt.Sprintf("Request %d - Time: %v ms", i, duration)
					default:
						res, code, err = httputil.GetString(url, proxyUrl)
					}

					if err != nil {
						fmt.Printf("Error getting url: %v\n", err)
						//fmt.Printf("Request %d - Result: %s\n", i, res)
					} else {
						fmt.Printf("[%d] : %s\n", code, res)
					}
				}(i)
			}
			wg.Wait()

		},
	}
	addProxyFlag(cmd)
	cmd.Flags().StringVarP(&resultType, "outformat", "o", "string", "format of result to return (string, json, time)")
	cmd.Flags().StringVarP(&repitions, "repitions", "n", "1", "number of calls to make")
	cmd.Flags().StringVarP(&parallel, "parallell", "m", "1", "number of parallell calls to make")
	return cmd
}
