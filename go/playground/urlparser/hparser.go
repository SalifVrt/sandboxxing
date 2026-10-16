package urlparser

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type ParsedURL struct {
	protocol string
	host     string
	port     int
	user     string
	password string
	query    string
	fragment string
	path     string
}

//Structure of the complete URL : https://user:pass@sub.domain.tld:8080/path/to/resource?query=1#fragment

func ParseURL(raw string) (ParsedURL, error) {
	//parse protocol
	protocol, rest, hasProtocol := strings.Cut(raw, "://")
	if !hasProtocol {
		return ParsedURL{}, errors.New("Invalid or missing URL scheme")
	}
	if protocol != "http" && protocol != "https" {
		return ParsedURL{}, errors.New("Invalid Protocol, use either http or https")
	}

	//parse fragment, then query
	r2, fragment, _ := strings.Cut(rest, "#")
	fullPath, query, _ := strings.Cut(r2, "?")

	//parse user, pass, host and port

	fullAuth, restPath, found := strings.Cut(fullPath, "/")
	var path string
	if !found {
		path = "/"
	} else {
		path = "/" + restPath
	}

	var (
		host, user, passwd string
		portStr            string
		port               int
	)

	userPass, hostPort, hasAuth := strings.Cut(fullAuth, "@")
	if hasAuth {
		//fmt.Println(userPass, hostPort)
		user, passwd, _ = strings.Cut(userPass, ":")
		host, portStr, _ = strings.Cut(hostPort, ":")

	} else {
		host, portStr, _ = strings.Cut(fullAuth, ":")
	}
	if portStr == "" {
		if protocol == "https" {
			port = 443
		} else {
			port = 80
		}
	} else {
		port, err := strconv.Atoi(portStr)

		if err != nil || port <= 0 || port > 65535 {
			//conversion error on port
			return ParsedURL{}, fmt.Errorf("Invalid port : %s", portStr)
		}
	}

	return ParsedURL{
		protocol: protocol,
		host:     host,
		port:     port,
		user:     user,
		password: passwd,
		query:    query,
		fragment: fragment,
		path:     path}, nil
}

/*
func main() {
	fmt.Println(ParseURL("https://archlinux.org"))
	fmt.Println(ParseURL("http://localhost:3000/api/v1?page=2#top"))
	fmt.Println(ParseURL("http://uuu:1234@localhost:3000/api/v1?page=2#top"))
	fmt.Println(ParseURL("ftp://bad-scheme.com"))
	fmt.Println(ParseURL("https://domain.com:abc/"))
}
*/
