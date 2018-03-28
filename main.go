package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func usage() {
	fmt.Println("Usage: sleep NUMBER[SUFFIX]")
	fmt.Println("    SUFFIX ns - nanoseconds")
	fmt.Println("           us - microseconds")
	fmt.Println("           ms - milliseconds")
	fmt.Println("            s - seconds")
	fmt.Println("            m - minutes")
	fmt.Println("            h - hours")
}

func main() {
	if len(os.Args) < 2 {
		usage()
		os.Exit(1)
	}

	tm := os.Args[1]

	_, err := strconv.Atoi(tm)

	if err == nil {
		tm = fmt.Sprintf("%ss", tm)
	}

	dur, err := time.ParseDuration(tm)

	if err != nil {
		fmt.Printf("Could not parse duration %s\n", tm)
		os.Exit(1)
	}

	time.Sleep(dur)
}
