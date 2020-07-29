package tools

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

func RandInt64(min, max int64) int64 {
	if min >= max || min == 0 || max == 0 {
		return max
	}
	return rand.Int63n(max-min) + min
}

func TimeToString(time string) string {
	str := ""
	t := strings.Split(time, " ")
	//fmt.Println(len(t))
	if len(t) == 2 {
		day := t[0]
		clock := t[1]
		t1 := strings.Split(day, "-")
		t2 := strings.Split(clock, ":")
		for _, v := range t1 {
			str += v
		}
		for _, v := range t2 {
			str += v
		}
	} else {
		day := t[0]
		t1 := strings.Split(day, "-")
		for _, v := range t1 {
			str += v
		}
	}
	return str
}

func WriteToFile(filepath, str string) {
	file, err := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	writer.WriteString(str)
	writer.Flush()
}
