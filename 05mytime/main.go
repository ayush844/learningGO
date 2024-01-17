package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("hello")

	presentTime := time.Now()
	fmt.Println(presentTime)
	// 2024-01-17 16:21:39.486733323 +0530 IST m=+0.000052713

	// Basic full date  "Mon Jan 2 15:04:05 MST 2006"

	fmt.Println(presentTime.Format("02-01-2006"))
	// 17-01-2024   => 17 jan 2006

	fmt.Println(presentTime.Format("15:04:05"))
	// 16:29:54   => 16hr 29min 54sec

	fmt.Println(presentTime.Format("Monday"))
	// Wednesday

	fmt.Println(presentTime.Format("02-01-2006 Monday 15:04:05"))
	// 17-01-2024 Wednesday 16:31:43

	// creating date //
	createdDate := time.Date(2004, time.April, 8, 23, 30, 0, 0, time.Local)

	fmt.Println(createdDate)
	// 2004-04-08 23:30:00 +0530 IST

	fmt.Println(createdDate.Format("02-01-2006 Monday 15:04:05"))
	// 08-04-2004 Thursday 23:30:00

}
