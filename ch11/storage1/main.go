package storage1

import (
	"fmt"
	"log"
	"net/smtp"
)

var usage = make(map[string]int64)

func bytesInUse(username string) int64 { return usage[username] }

const sender = "a@example.com"
const password = "this is obviously unsafe"
const hostname = "smtp.example.com"

const template = `Warning: you are using %d bytes of storage, %d%% of your quota`

func CheckQuota(username string) {
	used := bytesInUse(username)
	const quota = 100000000
	percent := 100 * used / quota
	if percent < 90 {
		return
	}
	msg := fmt.Sprintf(template, used, percent)
	auth := smtp.PlainAuth("", sender, password, hostname)
	err := smtp.SendMail(hostname+":587", auth, sender, []string{username}, []byte(msg))
	if err != nil {
		log.Printf("Send main failed for user (%s): %s", username, err)
	}
}
