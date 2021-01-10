package logger

import (
	"fmt"
	"log"
	"log/syslog"
	"os"
	"path/filepath"
)

type Logger interface {
	Log()
	LogFatal()
	LogPanic()
	CustomLog()
}

type Repository struct {}

func NewRepo() *Repository  {
	return &Repository{}
}

/**
grep LOG_ /var/log/syslog
grep LOG_LOCAL7 /var/log/cisco
grep LOG_MAIL /var/log/mail.log
 */
func(Repository) Log() {
	pn := filepath.Base(os.Args[0])
	sl, err := syslog.New(syslog.LOG_INFO|syslog.LOG_LOCAL7, pn)
	if err != nil {
		log.Fatal(err)
	} else {
		log.SetOutput(sl)
	}
	log.Println("LOG_INFO + LOG_LOCAL7: Logging in.")

	sl, err = syslog.New(syslog.LOG_MAIL, pn)
	if err != nil {
		log.Fatal(err)
	} else {
		log.SetOutput(sl)
	}

	log.Println("LOG_MAIL: logging in")
	log.Println("some text")
}

func(Repository) LogFatal() {
	sl, err := syslog.New(syslog.LOG_ALERT|syslog.LOG_MAIL, "Fatal example")
	if err != nil {
		log.Fatal(err)
	} else {
		log.SetOutput(sl)
	}

	log.Fatal(sl)
	fmt.Println("Some text")
}

func(Repository) LogPanic() {
	sl, err := syslog.New(syslog.LOG_ALERT|syslog.LOG_MAIL, "Panic example")
	if err != nil {
		log.Fatal(err)
	} else {
		log.SetOutput(sl)
	}

	log.Panic(sl)
	log.Println("some text")
}

//before use, create /tmp/custom_log.log
var LOGFILE = "/tmp/custom_log.log"
func(Repository) CustomLog(){
	f, err := os.OpenFile(LOGFILE, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Panic(err)
	}
	defer f.Close()
	il := log.New(f, "customLogLineNumber", log.LstdFlags)

	// log.LstdFlags date time
	// log.Lshortfile line number and file name
	il.SetFlags(log.LstdFlags|log.Lshortfile)
	il.Println("Message one")
	il.Println("Another log entry, message two")
}