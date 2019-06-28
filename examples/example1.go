package main

import (
	"context"
	"flag"

	"github.com/nikandfor/tlog"
	"github.com/nikandfor/tlog/examples/sub"
)

var (
	f   = flag.Int("f", 1, "int flag")
	str = flag.String("str", "two", "string flag")
)

func main() {
	flag.Parse()

	tlog.DefaultLabels.Set("mylabel", "value")
	tlog.DefaultLabels.Set("myflag", "")

	tlog.Printf("main: %d %q", *f, *str)

	sub.Func1(tlog.FullID{}, 5)

	tr := tlog.Start()
	defer tr.Finish()

	ctx := tlog.WithFullID(context.Background(), tr.ID)

	sub.Func2(ctx, 9)
}
