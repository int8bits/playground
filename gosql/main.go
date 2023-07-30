package main

import (
	"flag"
	"fmt"
)

var (
	host      *string
	port      *int
	user      *string
	password  *string
	db        *string
	sshPath   *string
	sshServer *string
	sshPort   *int
	queryStr  *string
	queryFile *string
	ouputPath *string
)

func init() {
	host = flag.String("host", "127.0.0.1", "host to connect")
	port = flag.Int("port", 1234, "port to connect")
	user = flag.String("user", "root", "user to connect")
	password = flag.String("password", string(""), "password to connect")
	db = flag.String("db", string(""), "database to connect")
	sshPath = flag.String("sshPath", string(""), "path of key ssh")
	sshServer = flag.String("sshServer", string(""), "ssh server to connect")
	sshPort = flag.Int("sshPort", -1, "ssh port to connect")
	queryStr = flag.String("query", string(""), "query to execute")
	queryFile = flag.String("queryFile", string(""), "file with query to execute")
	ouputPath = flag.String("output", string(""), "Outut of query")
}

func main() {
	flag.Parse()

	fmt.Println(*host)
	fmt.Println(*port)
	fmt.Println(*user)
	fmt.Println(*password)
	fmt.Println(*db)
	fmt.Println(*sshPath)
	fmt.Println(*sshServer)
	fmt.Println(*sshPort)
	fmt.Println(*queryStr)
	fmt.Println(*queryFile)
	fmt.Println(*ouputPath)
}
