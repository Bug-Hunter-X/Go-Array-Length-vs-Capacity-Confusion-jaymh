func main() {
  var a [10]int
  fmt.Println(len(a)) // Output: 10
  fmt.Println(cap(a)) // Output: 10

  // Correct use case of a slice for dynamic size:
  b := make([]int, 0, 10)
  fmt.Println(len(b)) // Output: 0
  fmt.Println(cap(b)) // Output: 10
  b = append(b, 1, 2, 3) 
  fmt.Println(len(b)) // Output: 3
  fmt.Println(cap(b)) // Output: 10
}