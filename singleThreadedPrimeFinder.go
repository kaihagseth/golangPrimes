package main
import (
    "fmt"
    "time"
)

const max_test = 1000000

// tests if an integer is dividable by earlies integers
func isPrime(ps []int, n int) (bool) {
    for _, p := range ps {
       if (n % p) == 0 {
         return false 
       }
    }
    return true 
}

func main() {
    start := time.Now()
    primes := []int{2}
    for i := 3; i < max_test; i+=2 {
        if isPrime(primes, i){
            primes = append(primes, i)
        }
    }
    elapsed := time.Since(start)
    fmt.Printf("Found %d primes in %s \n", len(primes), elapsed)
}
