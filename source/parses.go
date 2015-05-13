package source
import (
	"strconv"
	"fmt"
	"net/url"
	"net"
)

func parseNumber() {
	f, _ := strconv.ParseFloat("1.234", 64)
	fmt.Println(f)

	i, _ := strconv.ParseInt("f", 0, 64)
	fmt.Println(i)

	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Println(d)

	u, _ := strconv.ParseUint("789", 0, 64)
	fmt.Println(u)

	k, _ := strconv.Atoi("135")
	fmt.Println(k)

	_, e := strconv.Atoi("wat")
	fmt.Println(e)

}

func parseUrl() {
	s := "postgres://user:pass@host.com:5432/path?k=v#f"
	fmt.Println(s)

	u, e := url.Parse(s)
	if e!=nil {
		panic(e)
	}

	fmt.Println(u.Scheme)
	fmt.Println(u.User)
	fmt.Println(u.User.Username())
	pwd, _ := u.User.Password()
	fmt.Println(pwd)
	fmt.Println(u.Host)
	host, port, _ := net.SplitHostPort(u.Host);
	fmt.Println(host)
	fmt.Println(port)
	fmt.Println(u.Path)
	fmt.Println(u.Fragment)
	fmt.Println(u.RawQuery)
	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println(m)
	v:=m["k"][0]
	fmt.Println(v)



}

func Parses() {
	parseNumber()
	parseUrl()

}
