package main

import (
  "fmt"
  "math"
)

func VerificaPrimo(number int) bool {
    for i := 2; i <= int(math.Floor(float64(number) / 2)); i++ {
        if number%i == 0 {
            return false
        }
    }
    return number > 1
}

func main() {
  var num int
  fmt.Println("Digite um número: ")
  fmt.Scan(&num)
  res := VerificaPrimo(num)
  if res == true {
    fmt.Println("[+] É primo!")
  }else{
    fmt.Println("[-] Não é primo!")
  }
}
