package main
import "strings"
func RemoveEven(input []int) (result []int) {
  i := 0
  result = make([]int, 0)
  for i < len(input) {
    if input[i] % 2==1 {
      result = append(result, input[i])
    }
    i++
  }
  return
}
func PowerGenerator(a int) (result int) {
  deg := 1;
  deg = deg * a
  result = deg
  return
}
func Prin(x rune) bool {
  if x >= 'a' && x <= 'z' {
    return false
  } else {
    return true
  }
}
func DifferentWordsCount(c string) (counter int) {
  dif := strings.ToLower(c)
  text:= strings.FieldsFunc(dif, Prin)
  counter = 0;
  i:= 0;
  for i < len(text) {
    if text[i] != "1" {
      counter++
      j:=i+1
      for j < len (text) {
        if text[i] == text[j] {
          text[j] = "1"
        }
        j++
      }
      text[i] = "1"
    }
    i++
  }
  return
}
