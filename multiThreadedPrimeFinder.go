package main
import (
    "fmt"
    "time"
)

const max_test = 1000000
const workers =  4

// tests if an integer is dividable by earlies integers
func isPrime(ps []int, n int) (bool) {
    //fmt.Printf("Testing if %d is prime.\n", n)
    found := make(chan bool)
    done := make(chan bool)
    for i := 1; i <= workers; i++ {
        go isPrimeJob(ps, n, i, found, done)
    }

    workersDone := 0
    for workersDone < workers {
        select {
        case <-found:
        return false
        case <-done:
//        fmt.Println("Worker done")
        workersDone++
        }
    }
    return true
}

//test
func isPrimeJob(ps []int, n int, step int, found chan bool, done chan bool) {
    //fmt.Println("Starting job")
    for i := step-1; i < len(ps); i+=step {
       if (n % ps[i]) == 0 {
          found <- true // This is not a prime! Abort!
       }
    }
    done <- true
    //fmt.Println("Job done")
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
