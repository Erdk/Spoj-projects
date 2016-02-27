// Your program is to use the brute-force approach in order to find 
// the Answer to Life, the Universe, and Everything. More precisely...
// rewrite small numbers from input to output. Stop processing input after
// reading in the number 42. All numbers at input are integers of one or
// two digits.

package main

import (
  "os"
  "io"
  "bufio"
  "fmt"
  "strings"
)

func main() {
  reader := bufio.NewReader(os.Stdin)
  text, err := reader.ReadString('\n')
  text = strings.TrimSpace(text)

  for err != io.EOF && text != "42" {
    if text != "\n" {
      fmt.Println(text)
    }
    text, err = reader.ReadString('\n')
    text = strings.TrimSpace(text)
  }
}
