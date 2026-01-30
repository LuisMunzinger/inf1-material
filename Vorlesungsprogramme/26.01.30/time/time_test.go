package time

import (
	"fmt"
	"time"
)

func Example() {
	now := time.Now().Unix()

	year := now/60/60/24/365 + 1970
	minutes := (now / 60) % 60
	houers := (now / 60 / 60) % 24

	fmt.Println(now)
	fmt.Println(year)
	fmt.Println(minutes)
	fmt.Println(houers)

	// Output:
}
