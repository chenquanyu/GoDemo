// Echo1 prints its command-line arguments.
package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	var s, sep string
	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)

	times := 1000000
	s = ""
	start := time.Now()
	// for i := 0; i < times; i++ {
	// 	s += string(i)
	// }
	// fmt.Println(time.Now().Sub(start).Milliseconds())

	fmt.Println(start)

	start = time.Now()
	var buffer bytes.Buffer
	for i := 0; i < times; i++ {
		buffer.WriteString(strconv.Itoa(i))
	}
	fmt.Printf("length: %d buffer: %s \n", len(buffer.String()), time.Since(start))

	start = time.Now()
	byteArray := []byte{}
	for i := 0; i < times+1; i++ {
		byteArray = append(byteArray, strconv.Itoa(i)...)
	}
	fmt.Printf("length: %d append: %s \n", len(string(byteArray)), time.Since(start))

	start = time.Now()
	var builder strings.Builder
	for i := 0; i < times+2; i++ {
		builder.WriteString(strconv.Itoa(i))
	}
	fmt.Printf("length: %d builder: %s \n", len(builder.String()), time.Since(start))

}
