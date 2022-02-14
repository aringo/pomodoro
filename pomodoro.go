// pormodoro timer
package main

import (
	"flag"
	"fmt"
	"os"
	"time"
	"log"

)

func main() {
	start := time.Now()


	

	finish, err := waitDuration(start)
	if err != nil {
		flag.Usage()
		os.Exit(2)
	}
	wait := finish.Sub(start)

	formatter := formatSeconds
	switch {
	case wait >= 24*time.Hour:
		formatter = formatDays
	case wait >= time.Hour:
		formatter = formatHours
	case wait >= time.Minute:
		formatter = formatMinutes
	}

	//Set up file for logging 
    dirname, err := os.UserHomeDir()
	if err != nil {
		log.Fatal( err )
	}

	logfile := dirname + "/pomodoro.log"
	fmt.Println( logfile )
	f, err := os.OpenFile(logfile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	_, err2 := f.WriteString(start.Format("2006-01-02 15:04:05") + " " + *task + "\n")
	if err2 != nil {
        log.Fatal(err2)
    }

	fmt.Printf("Start timer for %s.\n\n", wait)

	if *simple {
		simpleCountdown(finish, formatter)
	} else {
		task := *task
		fullscreenCountdown(start, finish, formatter, task)
	}

	if !*silence {
		fmt.Println("\a") // \a is the bell literal.
	} else {
		fmt.Println()
	}
}
