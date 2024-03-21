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


func ShowUserMessage(isPrime bool){
  if isPrime {
    fmt.Println("[+] É primo!")
  }else{
    fmt.Println("[-] Não é primo!")
  }
}

func main() {
  ShowUserMessage(VerificaPrimo(0))
  ShowUserMessage(VerificaPrimo(1))
  ShowUserMessage(VerificaPrimo(2))
  ShowUserMessage(VerificaPrimo(3))
  ShowUserMessage(VerificaPrimo(7))
  ShowUserMessage(VerificaPrimo(83))
  ShowUserMessage(VerificaPrimo(100))
  ShowUserMessage(VerificaPrimo(991))
  ShowUserMessage(VerificaPrimo(104729))
  ShowUserMessage(VerificaPrimo(14348907))
}
