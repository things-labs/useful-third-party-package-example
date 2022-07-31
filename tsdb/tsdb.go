package main

import (
	"context"
	"fmt"
	"math"
	"time"

	"github.com/prometheus/prometheus/model/labels"
	"github.com/prometheus/prometheus/tsdb"
)

func main() {
	// Create a random dir to work in.  Open() doesn't require a pre-existing dir, but
	// we want to make sure not to make a mess where we shouldn't.
	// dir, err := os.MkdirTemp("", "tsdb-test")
	// noErr(err)
	dir := "./tsdb-test"
	// Open a TSDB for reading and/or writing.
	db, err := tsdb.Open(dir, nil, nil, tsdb.DefaultOptions(), nil)
	noErr(err)

	// Open an appender for writing.
	app := db.Appender(context.Background())

	series := labels.FromStrings("foo", "bar")

	// Ref is 0 for the first append since we don't know the reference for the series.
	ref, err := app.Append(0, series, time.Now().Unix(), 123)
	noErr(err)

	// Another append for a second later.
	// Re-using the ref from above since it's the same series, makes append faster.
	time.Sleep(time.Second)
	_, err = app.Append(ref, series, time.Now().Unix(), 124)
	noErr(err)

	// Commit to storage.
	err = app.Commit()
	noErr(err)

	// In case you want to do more appends after app.Commit(),
	// you need a new appender.
	app = db.Appender(context.Background())
	// ... adding more samples.

	// Open a querier for reading.
	querier, err := db.Querier(context.Background(), math.MinInt64, math.MaxInt64)
	noErr(err)
	ss := querier.Select(false, nil, labels.MustNewMatcher(labels.MatchEqual, "foo", "bar"))

	for ss.Next() {
		series := ss.At()
		fmt.Println("series:", series.Labels().String())

		it := series.Iterator()
		for it.Next() {
			_, v := it.At() // We ignore the timestamp here, only to have a predictable output we can test against (below)
			fmt.Println("sample", v)
		}

		fmt.Println("it.Err():", it.Err())
	}
	fmt.Println("ss.Err():", ss.Err())
	ws := ss.Warnings()
	if len(ws) > 0 {
		fmt.Println("warnings:", ws)
	}
	err = querier.Close()
	noErr(err)

	// Clean up any last resources when done.
	err = db.Close()
	noErr(err)
	// err = os.RemoveAll(dir)
	// noErr(err)
}

func noErr(err error) {
	if err != nil {
		panic(err)
	}
}
