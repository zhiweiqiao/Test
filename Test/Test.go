package main

import (
	"errors"
	"fmt"
	"net/smtp"
	"strings"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	//"log"

	//"github.com/emersion/go-imap"
	//"github.com/emersion/go-imap/client"
)

// Chaincode Prototype implementation
type ChaincodePrototype struct {
}

func (t *ChaincodePrototype) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Printf("Init called, initializing chaincode")

	return nil, nil
}

// Transaction makes payment of X units from A to B
func (t *ChaincodePrototype) transaction(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	fmt.Printf("Running invoke")

	user := "hyperledgertest1@gmail.com"
	password := "George2017"
	host := "smtp.gmail.com:587"
	to1 := "hyperledgertest2@gmail.com"
	to2 := "hyperledgertest3@gmail.com"

	subject := "Indentification"

	body := `
		<html>
		<body>
		<h1>
		"Have you finished your work yet?"
		</h1>
		<button name="yes" type="submit" style="height:50px;width:100px;font-size:30px">Yes</button>
		&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
		<button name="no" type="submit" style="height:50px;width:100px;font-size:30px">No</button>
		</body>
		</html>
		`
	fmt.Println("Sending email")

	err1 := SendToMail(user, password, host, to1, subject, body, "html")
	if err1 != nil {
		fmt.Println("Send email1 error!")
		fmt.Println(err1)
	} else {
		fmt.Println("Send email1 success!")
	}

	err2 := SendToMail(user, password, host, to2, subject, body, "html")
	if err2 != nil {
		fmt.Println("Send email2 error!")
		fmt.Println(err2)
	} else {
		fmt.Println("Send email2 success!")
	}
	/*
	log.Println("Connecting to server...")

	// Connect to server
	c, err := client.DialTLS("imap.gmail.com:993", nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected")

	// Don't forget to logout
	defer c.Logout()

	// Login
	if err := c.Login("hyperledgertest1@gmail.com", "George2017"); err != nil {
		log.Fatal(err)
	}
	log.Println("Logged in")

	// List mailboxes
	mailboxes := make(chan *imap.MailboxInfo, 10)
	done := make(chan error, 1)
	go func () {
		done <- c.List("", "*", mailboxes)
	}()

	log.Println("Mailboxes:")
	for m := range mailboxes {
		log.Println("* " + m.Name)
	}

	if err := <-done; err != nil {
		log.Fatal(err)
	}

	// Select INBOX
	mbox, err := c.Select("INBOX", false)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Flags for INBOX:", mbox.Flags)

	// Get the last 4 messages
	from := uint32(1)
	to := mbox.Messages
	if mbox.Messages > 3 {
		// We're using unsigned integers here, only substract if the result is > 0
		from = mbox.Messages - 3
	}
	seqset := new(imap.SeqSet)
	seqset.AddRange(from, to)

	messages := make(chan *imap.Message, 10)
	done = make(chan error, 1)
	go func() {
		done <- c.Fetch(seqset, []string{imap.EnvelopeMsgAttr}, messages)
	}()

	log.Println("Last 4 messages:")
	for msg := range messages {
		log.Println("* " + msg.Envelope.Subject)
	}

	if err := <-done; err != nil {
		log.Fatal(err)
	}

	log.Println("Done!")*/
	
	
	return nil, nil
}

// Invoke callback representing the invocation of a chaincode
// This chaincode will manage two accounts A and B and will transfer X units from A to B upon invoke
func (t *ChaincodePrototype) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Printf("Invoke called, determining function")
	
	// Handle different functions
	if function == "invoke" {
		// Transaction makes payment of X units from A to B
		fmt.Printf("Function is invoke")
		return t.transaction(stub, args)
	} else if function == "init" {
		fmt.Printf("Function is init")
		return t.Init(stub, function, args)
	}

	return nil, errors.New("Received unknown function invocation")
}

// Query callback representing the query of a chaincode
func (t *ChaincodePrototype) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Printf("Query called, determining function")
	
	return nil, nil
}

func SendToMail(user, password, host, to, subject, body, mailtype string) error {
	hp := strings.Split(host, ":")
	auth := smtp.PlainAuth("", user, password, hp[0])
	var content_type string
	if mailtype == "html" {
		content_type = "Content-Type: text/" + mailtype + "; charset=UTF-8"
	} else {
		content_type = "Content-Type: text/plain" + "; charset=UTF-8"
	}

	msg := []byte("To: " + to + "\r\nFrom: " + user + "\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
	send_to := strings.Split(to, ";")
	err := smtp.SendMail(host, auth, user, send_to, msg)
	return err
}

func main() {
	err := shim.Start(new(ChaincodePrototype))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}
