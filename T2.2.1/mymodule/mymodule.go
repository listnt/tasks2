package mymodule

import (
	"fmt"
	"github.com/beevik/ntp"
	"os"
)

func Time() int{
	time, err := ntp.Time("time.nist.gov")
	if err != nil {
		os.Stderr.Write([]byte(err.Error()))
		return 2
	}
	const layout = "3:04:05 PM (MST) on Monday, January _2, 2006"
	fmt.Println("Current Local Time:")
	fmt.Println(time.Local().Format(layout))
	return 1
}