package novichok

import (
	"fmt"
	"net/url"
	"time"
)

func urlFromString(raw string) (res *url.URL) {
	res, _ = url.Parse(raw)
	return
}

func stringMonthFromTime(target time.Time) string {
	var month string
	if target.Month() < 10 {
		month = fmt.Sprintf("0%d", target.Month())
	} else {
		month = fmt.Sprintf("%d", target.Month())
	}
	return month
}
