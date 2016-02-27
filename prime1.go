package main

// Peter wants to generate some prime numbers for his cryptosystem. Help him!
// Your task is to generate all prime numbers between two given numbers!
//
// Input
// The input begins with the number t of test cases in a single line (t<=10).
// In each of the next t lines there are two numbers m and
// n (1 <= m <= n <= 1000000000, n-m<=100000) separated by a space.
//
// Output
//
// For every test case print all prime numbers p such that m <= p <= n,
// one number per line, test cases separated by an empty line.

import(
  "fmt"
  "bufio"
  "os"
  "strings"
  "strconv"
  "math"
)

func main() {
  reader := bufio.NewReader(os.Stdin)
  text, _ := reader.ReadString('\n')
  text = strings.TrimSpace(text)

  num_tests, _ := strconv.Atoi(text)

  for num_tests > 0 {
    num_tests = num_tests - 1

    // get boundaries
    text, _ = reader.ReadString('\n')
    text = strings.TrimSpace(text)
    nums := strings.Split(text, " ")

    lower, _ := strconv.Atoi(nums[0])
    if lower == 1 {
      lower = 2
    }
    higher, _ := strconv.Atoi(nums[1])

    for lower <= higher {
      current := 2
      half := int(math.Sqrt(float64(lower)))
      prime := true
      for current <= half {
        if lower % current == 0 {
          prime = false
          break
        }
        current = current + 1
      }
      if prime {
        fmt.Println(lower)
      }
      lower = lower + 1
    }

    fmt.Println()
  }
}
