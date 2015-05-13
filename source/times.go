package source
import (
	"time"
	"fmt"
)

func Times() {

	now := time.Now()
	fmt.Println(now)

	fmt.Println(now.Weekday())

	then := now.Add(-time.Hour*48)
	fmt.Println(then)

	diff := now.Sub(then)
	fmt.Println(diff)

	diff=then.Sub(now)
	fmt.Println(diff)

	sec := now.Unix()
	fmt.Println(sec)

	nanos := now.UnixNano()
	fmt.Println(nanos)

	millis := nanos/1000000
	fmt.Println(millis)

	fmt.Println(time.Unix(sec, 0))
	fmt.Println(time.Unix(0, nanos))
	fmt.Println(time.Unix(sec, nanos))

	fmt.Println(now.Format(time.RFC3339))

	form := "3 04 PM"
	t2, _ := time.Parse(form, "8 41 PM")
	fmt.Println(t2)

	fmt.Println(now.Local())



}
