package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/mskrha/gosmtp"
)

var (
	version string
)

func main() {
	var s, f, r, p string

	flag.StringVar(&s, "s", "", "Message subject")
	flag.StringVar(&f, "f", "", "Sender address")
	flag.StringVar(&r, "r", "", "Recipient address")
	flag.StringVar(&p, "p", "127.0.0.1", "SMTP proxy host")
	flag.Parse()

	smtp, err := gosmtp.NewServer(p, fmt.Sprintf("GoMail %s", version))
	if err != nil {
		fmt.Println(err)
		return
	}

	in, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Println(err)
		return
	}

	msg, err := gosmtp.NewMessage(f, r, s, string(in))
	if err != nil {
		fmt.Println(err)
		return
	}

	id, err := smtp.Send(msg)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Message queued on the SMTP proxy with ID: %s\n", id)
}
