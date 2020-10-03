package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/smtp"
	"os"
	"time"
)

const (
	RFC2822 = "Mon, _2 Jan 2006 15:04:05 -0700"
)

type Message struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
}

var (
	version string
)

func main() {
	var msg Message
	var srv string

	flag.StringVar(&msg.Subject, "s", "", "Message subject")
	flag.StringVar(&msg.From, "f", "", "Sender address")
	flag.StringVar(&msg.To, "r", "", "Recipient address")
	flag.StringVar(&srv, "p", "127.0.0.1", "SMTP proxy host")
	flag.Parse()

	if in, err := ioutil.ReadAll(os.Stdin); err != nil {
		panic(err)
	} else {
		msg.Body = string(in)
	}

	if len(msg.Subject) == 0 {
		fmt.Println("No subject specified!")
		return
	}

	if len(msg.From) == 0 {
		fmt.Println("No sender specified!")
		return
	}

	if len(msg.To) == 0 {
		fmt.Println("No recipient specified!")
		return
	}

	fmt.Println(msg.Send(srv))
}

func (m *Message) Send(srv string) (err error) {
	host, err := os.Hostname()
	if err != nil {
		panic(err)
	}

	var x string
	t := time.Now()
	x += fmt.Sprintf("From: %s\r\n", m.From)
	x += fmt.Sprintf("To: %s\r\n", m.To)
	x += fmt.Sprintf("Subject: %s\r\n", m.Subject)
	x += fmt.Sprintf("Date: %s\r\n", t.Format(RFC2822))
	x += fmt.Sprintf("Message-ID: %d@%s\r\n", t.UnixNano(), host)
	x += fmt.Sprintf("User-Agent: GoMail %s\r\n", version)
	x += "\r\n"
	x += m.Body
	x += "\r\n"

	c, err := smtp.Dial(srv + ":25")
	if err != nil {
		return
	}

	err = c.Hello(host)
	if err != nil {
		return
	}

	err = c.Mail(m.From)
	if err != nil {
		return
	}

	err = c.Rcpt(m.To)
	if err != nil {
		return
	}

	d, err := c.Data()
	if err != nil {
		return
	}

	_, err = d.Write([]byte(x))
	if err != nil {
		return
	}

	return c.Quit()
}
