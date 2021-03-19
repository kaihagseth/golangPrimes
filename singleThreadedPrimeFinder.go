package main
import (
    "fmt"
    "time"
)

const max_test = 500000

// tests if an integer is dividable by any of the integers in the passed array
func isPrime(ps []int, n int) (bool) {
    for _, p := range ps {
       if (n % p) == 0 {
         return false 
       }
    }
    return true 
}

// returns all primes up to a limit with the exception of 2
func getPrimes(max int) ([]int) {
    primes := []int{}
    for i := 3; i < max_test; i+=2 {
        if isPrime(primes, i){
            primes = append(primes, i)
        }
    }
    return primes
}

func main() {
    start := time.Now()
    primes := getPrimes(max_test)
    elapsed := time.Since(start)
    fmt.Printf("Found %d primes in %s \n", len(primes) + 1, elapsed)
}
