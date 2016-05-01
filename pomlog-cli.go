package main

import (
	"encoding/json"
	"fmt"
	"gopkg.in/cheggaaa/pb.v1"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"
)

func progressbar(minutes int, message string) {
	count := minutes * 60
	bar := pb.New(count)
	bar.ShowTimeLeft = false
	bar.ShowCounters = false
	bar.SetMaxWidth(80)
	bar.Start()
	for i := 0; i < count; i++ {
		bar.Increment()
		time.Sleep(time.Second)
	}
	bar.FinishPrint(message)
}

func main() {
	file, e := ioutil.ReadFile("./config.json")
	if e != nil {
		log.Fatal(e)
	}

	logfile, e := os.OpenFile("./pomodoro.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if e != nil {
		log.Fatal(e)
	}

	var items []string
	json.Unmarshal(file, &items)

	for pomodoro := 1; ; pomodoro++ {
		for {
			fmt.Println("What number are you doing?")
			for i, v := range items {
				fmt.Printf("[%d: %s]", i, v)
			}
			fmt.Printf("[999: exit]\n")
			var s string
			var number int
			fmt.Scanln(&s)
			n, _ := fmt.Sscanf(s, "%d", &number)
			if n == 1 && number < len(items) {
				fmt.Fprintf(logfile, "%s: %s start\n", time.Now().Format(time.RFC3339), items[number])
				break
			}
			if number == 999 {
				logfile.Close()
				os.Exit(0)
			}
		}
		progressbar(25, "Take 5 min rest!")
		progressbar(5, strconv.Itoa(pomodoro)+" pomodoro")
	}
}
