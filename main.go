package main

import (
	"flag"
	"fmt"
	"strings"
	"time"

	"github.com/araddon/dateparse"
)

// TODO:
// Output Formats:
//
// Standard output
// Jira output

var (
	timezone = ""
	date     = time.Now()
)

func useZone(datestring string, zone string) (date time.Time) {

	tz, err := time.LoadLocation(zone)
	if err != nil {
		panic(err.Error())
	}
	time.Local = tz

	date, err = dateparse.ParseIn(datestring, time.Local)
	if err != nil {
		panic(err.Error())
	}
	return
}

func useUTC(datestring string) (date time.Time) {
	date, err := dateparse.ParseAny(datestring)
	if err != nil {
		panic(err.Error())
	}
	return
}

func main() {

	flag.StringVar(&timezone, "timezone", "UTC", "Timezone aka `America/Los_Angeles` formatted time-zone")
	convertToLocal := flag.Bool("convertToLocal", false, "Parse output to your local time")
	flag.Parse()

	args := flag.Args()
	fullstr := strings.Join(args, " ")
	fmt.Println("Input:")
	fmt.Println(args)

	if timezone != "UTC" {
		date = useZone(fullstr, timezone).UTC()
	} else {
		date = useUTC(fullstr)
	}

	if *convertToLocal {
		fmt.Println(date.Local().String())
	} else {
		fmt.Println(date.String())
	}
}
