package main

import "strings"
import "fmt"
import "strconv"
import "bufio"
import "os"

const TEST_MODE = false

func AssertEqual(expected int, actual int) {
  if expected == actual {
    fmt.Printf(".")
  } else {
    fmt.Printf("F\nExpected: %d\nActual: %d\n", expected, actual)
  }
}

func AssertEqualStr(expected string, actual string) {
  if expected == actual {
    fmt.Printf(".")
  } else {
    fmt.Printf("F\nExpected: %s\nActual: %s\n", expected, actual)
  }
}

// ================================

func TestMinuteDeduction() {
  fmt.Println("TestMinuteDeduction")
  AssertEqual(0, DeductFromMinutes(0, 0))
  AssertEqual(50, DeductFromMinutes(50, 0))
  AssertEqual(49, DeductFromMinutes(50, 1))
  AssertEqual(5, DeductFromMinutes(30, 25))
  AssertEqual(0, DeductFromMinutes(30, 30))
  AssertEqual(0, DeductFromMinutes(30, 30))
  AssertEqual(50, DeductFromMinutes(30, 40))
  AssertEqual(22, DeductFromMinutes(30, 68))
  AssertEqual(22, DeductFromMinutes(30, 128))
  AssertEqual(5, DeductFromMinutes(50, 45))
  AssertEqual(0, DeductFromMinutes(45, 45))
  AssertEqual(59, DeductFromMinutes(45, 46))
  fmt.Println()
}

func TestHourDeduction() {
  fmt.Println("TestHourDeduction")
  AssertEqual(0, DeductFromHours(0, 0, 0))
  AssertEqual(15, DeductFromHours(15, 0, 0))
  AssertEqual(15, DeductFromHours(15, 3, 3))
  AssertEqual(14, DeductFromHours(15, 3, 4))
  AssertEqual(14, DeductFromHours(15, 3, 60))
  AssertEqual(14, DeductFromHours(15, 3, 61))
  AssertEqual(22, DeductFromHours(23, 3, 63))
  AssertEqual(21, DeductFromHours(23, 3, 64))
  AssertEqual(22, DeductFromHours(23, 40, 45))
  AssertEqual(23, DeductFromHours(0, 40, 45))
  AssertEqual(9, DeductFromHours(10, 10, 45))
  fmt.Println()
}

func TestProgram() {
  fmt.Println("TestProgram")
  AssertEqualStr("9 25", DeductTime("10 10"))
  AssertEqualStr("23 45", DeductTime("0 30"))
  AssertEqualStr("22 55", DeductTime("23 40"))
  fmt.Println()
}

func RunTests() {
  TestMinuteDeduction()
  TestHourDeduction()
  TestProgram()
}

// ================================

func DeductTime(input string) string {
  string_split := strings.Split(input, " ")
  // fmt.Println(string_split)
  oldHours, _ := strconv.Atoi(string_split[0])
  oldMinutes, _ := strconv.Atoi(string_split[1])
  newHours := DeductFromHours(oldHours, oldMinutes, 45)
  newMinutes := DeductFromMinutes(oldMinutes, 45)
  // fmt.Printf("%d %d, %d %d\n", oldHours, oldMinutes, newHours, newMinutes)
  return strconv.Itoa(newHours) + " " + strconv.Itoa(newMinutes)
}

func DeductFromMinutes(minutes int, amount int) int {
  potentialMinutes := minutes - amount
  for potentialMinutes < 0 {
    potentialMinutes += 60
  }
  return potentialMinutes
}

func DeductFromHours(hours int, minutes int, amount int) int {
  potentialHours := 0
  minuteDifference := -(minutes - amount)
  if minuteDifference > 120 {
    potentialHours = hours - 3
  } else if minuteDifference > 60 {
    potentialHours = hours - 2
  } else if minuteDifference > 0 {
    potentialHours = hours - 1
  } else {
    potentialHours = hours
  }
  for potentialHours < 0 {
    potentialHours += 24
  }
  return potentialHours
}

func Run() {
  buf := bufio.NewReader(os.Stdin)
  input, _ := buf.ReadString('\n')
  input = strings.Trim(input, "\n")
  fmt.Println(DeductTime(input))
}

// ================================

func main() {
  if TEST_MODE {
    RunTests()
  } else {
    Run()
  }
}