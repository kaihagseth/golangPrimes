/*
 * Attempt to find all primes up to 500k by splitting the workload between different cores
 * Instead of only testing the 
 */
package main
import (
    "fmt"
    "time"
)

const max_test = 500000
const batch_size = 5000
const workers =  4

// tests if an integer is dividable by earlies integers
func isPrime(n int) (bool) {
    for i := 3; i < n; i+=2 {
        if(n % i) == 0 {
           return false
       }
    }
    return true
}

func countPrimesBetween(begin int, end int, c chan int) {
    count := 0
    for i := begin; i < end; i+=2 {
        if isPrime(i) {
            count++
        }
    }
    //fmt.Printf("found %d primes between %d and %d\n", count, begin, end)
    c <- count
}


func main() {
    start := time.Now()
    nprimes := 0
    c := make(chan int)

    jobs := 0
    for i := 3; i < max_test; i+=batch_size {
        go countPrimesBetween(i, i + batch_size, c)
        jobs++
    }
    for jobs > 0 {
        nprimes += <-c
        jobs--
        //fmt.Printf("Jobs left: %d\n", jobs)
    }

    elapsed := time.Since(start)
    fmt.Printf("Found %d primes in %s \n", nprimes + 1, elapsed)
}
