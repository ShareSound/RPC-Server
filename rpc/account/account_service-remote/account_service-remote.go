// Autogenerated by Thrift Compiler (0.9.3)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package main

import (
	"flag"
	"fmt"
	"github.com/ShareSound/RPC-Server/rpc/account"
	"github.com/ShareSound/RPC-Server/rpc/shared"
	"github.com/mshockwave/thrift-go/thrift"
	"math"
	"net"
	"net/url"
	"os"
	"strconv"
	"strings"
)

func Usage() {
	fmt.Fprintln(os.Stderr, "Usage of ", os.Args[0], " [-h host:port] [-u url] [-f[ramed]] function [arg1 [arg2...]]:")
	flag.PrintDefaults()
	fmt.Fprintln(os.Stderr, "\nFunctions:")
	fmt.Fprintln(os.Stderr, "  Session registerAccount(string email, string username, string password)")
	fmt.Fprintln(os.Stderr, "  Session login(string email, string password)")
	fmt.Fprintln(os.Stderr, "  void logout(Session ctx)")
	fmt.Fprintln(os.Stderr, "  ProfileResult getProfile(Session ctx)")
	fmt.Fprintln(os.Stderr)
	os.Exit(0)
}

func main() {
	flag.Usage = Usage
	var host string
	var port int
	var protocol string
	var urlString string
	var framed bool
	var useHttp bool
	var parsedUrl url.URL
	var trans thrift.TTransport
	_ = strconv.Atoi
	_ = math.Abs
	flag.Usage = Usage
	flag.StringVar(&host, "h", "localhost", "Specify host and port")
	flag.IntVar(&port, "p", 9090, "Specify port")
	flag.StringVar(&protocol, "P", "binary", "Specify the protocol (binary, compact, simplejson, json)")
	flag.StringVar(&urlString, "u", "", "Specify the url")
	flag.BoolVar(&framed, "framed", false, "Use framed transport")
	flag.BoolVar(&useHttp, "http", false, "Use http")
	flag.Parse()

	if len(urlString) > 0 {
		parsedUrl, err := url.Parse(urlString)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
			flag.Usage()
		}
		host = parsedUrl.Host
		useHttp = len(parsedUrl.Scheme) <= 0 || parsedUrl.Scheme == "http"
	} else if useHttp {
		_, err := url.Parse(fmt.Sprint("http://", host, ":", port))
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
			flag.Usage()
		}
	}

	cmd := flag.Arg(0)
	var err error
	if useHttp {
		trans, err = thrift.NewTHttpClient(parsedUrl.String())
	} else {
		portStr := fmt.Sprint(port)
		if strings.Contains(host, ":") {
			host, portStr, err = net.SplitHostPort(host)
			if err != nil {
				fmt.Fprintln(os.Stderr, "error with host:", err)
				os.Exit(1)
			}
		}
		trans, err = thrift.NewTSocket(net.JoinHostPort(host, portStr))
		if err != nil {
			fmt.Fprintln(os.Stderr, "error resolving address:", err)
			os.Exit(1)
		}
		if framed {
			trans = thrift.NewTFramedTransport(trans)
		}
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error creating transport", err)
		os.Exit(1)
	}
	defer trans.Close()
	var protocolFactory thrift.TProtocolFactory
	switch protocol {
	case "compact":
		protocolFactory = thrift.NewTCompactProtocolFactory()
		break
	case "simplejson":
		protocolFactory = thrift.NewTSimpleJSONProtocolFactory()
		break
	case "json":
		protocolFactory = thrift.NewTJSONProtocolFactory()
		break
	case "binary", "":
		protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()
		break
	default:
		fmt.Fprintln(os.Stderr, "Invalid protocol specified: ", protocol)
		Usage()
		os.Exit(1)
	}
	client := account.NewAccountServiceClientFactory(trans, protocolFactory)
	if err := trans.Open(); err != nil {
		fmt.Fprintln(os.Stderr, "Error opening socket to ", host, ":", port, " ", err)
		os.Exit(1)
	}

	switch cmd {
	case "registerAccount":
		if flag.NArg()-1 != 3 {
			fmt.Fprintln(os.Stderr, "RegisterAccount requires 3 args")
			flag.Usage()
		}
		argvalue0 := flag.Arg(1)
		value0 := argvalue0
		argvalue1 := flag.Arg(2)
		value1 := argvalue1
		argvalue2 := flag.Arg(3)
		value2 := argvalue2
		fmt.Print(client.RegisterAccount(value0, value1, value2))
		fmt.Print("\n")
		break
	case "login":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "Login requires 2 args")
			flag.Usage()
		}
		argvalue0 := flag.Arg(1)
		value0 := argvalue0
		argvalue1 := flag.Arg(2)
		value1 := argvalue1
		fmt.Print(client.Login(value0, value1))
		fmt.Print("\n")
		break
	case "logout":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "Logout requires 1 args")
			flag.Usage()
		}
		arg15 := flag.Arg(1)
		mbTrans16 := thrift.NewTMemoryBufferLen(len(arg15))
		defer mbTrans16.Close()
		_, err17 := mbTrans16.WriteString(arg15)
		if err17 != nil {
			Usage()
			return
		}
		factory18 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt19 := factory18.GetProtocol(mbTrans16)
		argvalue0 := shared.NewSession()
		err20 := argvalue0.Read(jsProt19)
		if err20 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.Logout(value0))
		fmt.Print("\n")
		break
	case "getProfile":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "GetProfile requires 1 args")
			flag.Usage()
		}
		arg21 := flag.Arg(1)
		mbTrans22 := thrift.NewTMemoryBufferLen(len(arg21))
		defer mbTrans22.Close()
		_, err23 := mbTrans22.WriteString(arg21)
		if err23 != nil {
			Usage()
			return
		}
		factory24 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt25 := factory24.GetProtocol(mbTrans22)
		argvalue0 := shared.NewSession()
		err26 := argvalue0.Read(jsProt25)
		if err26 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.GetProfile(value0))
		fmt.Print("\n")
		break
	case "":
		Usage()
		break
	default:
		fmt.Fprintln(os.Stderr, "Invalid function ", cmd)
	}
}
