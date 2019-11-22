package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var zeroTime = time.Time{}

type options struct {
	layout string
	utc    bool
}

func main() {
	opts := options{
		layout: "2/01/2006 15:04:05 MST",
		utc:    true,
	}

	cmd := &cobra.Command{
		Use:   "ut <timestamp>",
		Short: "ut - command-line UNIX timestamp converter",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			t, err := convertTimestamp(args[0])
			if err != nil {
				return err
			}

			fmt.Fprintln(os.Stdout, formatTime(t, opts))

			return nil
		},
	}

	cmd.Flags().BoolVarP(&opts.utc, "utc", "", opts.utc, "output time in UTC")

	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func convertTimestamp(timestamp string) (time.Time, error) {
	sec, err := strconv.ParseInt(timestamp, 10, 64)
	if err != nil {
		return zeroTime, errors.Wrap(err, "failed to format timestamp")
	}

	if sec < 0 {
		return zeroTime, errors.New("invalid timestamp")
	}

	return time.Unix(sec, 0), nil
}

func formatTime(t time.Time, opts options) string {
	if opts.utc {
		return t.UTC().Format(opts.layout)
	}

	return t.Format(opts.layout)
}
