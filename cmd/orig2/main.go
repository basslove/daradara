package main

import (
	"fmt"
	"regexp"
	"time"
)

func main() {
	fmt.Println("start")

	// 時間
	t := time.Now()
	fmt.Println(t.Format(time.RFC3339))

	// 正規表現
	match, _ := regexp.MatchString("a([a-z]+)e", "apple")
	fmt.Println(match)

	r := regexp.MustCompile("a([a-z]+)e")
	ms := r.MatchString("apple")
	fmt.Println(ms)

	r2 := regexp.MustCompile("^/(edit|save|view)/")
	fs := r2.FindString("dsgfw/test")
	fmt.Println(fs)

	// context
	// semaphore
}
