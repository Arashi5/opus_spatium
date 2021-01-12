package logger

import (
	"errors"
	"fmt"
	"log"
	"log/syslog"
	"os"
	"path/filepath"
	"work_space/pkg/messages"
)

type Exec struct {
	repo repository
}

type repository struct{}

func NewRepo() *Exec {
	return &Exec{repo: repository{}}
}

func (e Exec) Exec(args []string) *error {
	var err error
	switch args[0] {
	case "log":
		e.repo.log()
	case "fatal":
		e.repo.logFatal()
	case "panic":
		e.repo.logPanic()
	case "custom":
		e.repo.customLog()
	default:
		err = errors.New(messages.ArgErrorMessage("Logger"))
	}

	if err != nil {
		return &err
	} else {
		return nil
	}
}

/**
grep LOG_ /var/log/syslog
grep LOG_LOCAL7 /var/log/cisco
grep LOG_MAIL /var/log/mail.log
*/
func (repository) log() {
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

func (repository) logFatal() {
	sl, err := syslog.New(syslog.LOG_ALERT|syslog.LOG_MAIL, "Fatal example")
	if err != nil {
		log.Fatal(err)
	} else {
		log.SetOutput(sl)
	}

	log.Fatal(sl)
	fmt.Println("Some text")
}

func (repository) logPanic() {
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

func (repository) customLog() {
	f, err := os.OpenFile(LOGFILE, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Panic(err)
	}
	defer f.Close()
	il := log.New(f, "customLogLineNumber", log.LstdFlags)

	// log.LstdFlags date time
	// log.Lshortfile line number and file name
	il.SetFlags(log.LstdFlags | log.Lshortfile)
	il.Println("Message one")
	il.Println("Another log entry, message two")
}
