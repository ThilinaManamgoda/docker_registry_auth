

package main

import (
	"fmt"
	"net/url"
)

func main() {


	u, _ := url.Parse("https://google.com/")


	//
	//
	//u1:=&url.URL{
	//	Host:u.Host,
	//	Scheme:u.Scheme,
	//	RawPath:url.QueryEscape("ss@ss.com"),
	//	Path:"ss@ss.com",
	//}

	u.Path="ss@ss.com"
	u.RawPath=url.QueryEscape("ss@ss.com")
	fmt.Println(u.String())
}
