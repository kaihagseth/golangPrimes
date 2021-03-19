/*
 * Attempt to find all primes up to 500k by splitting the workload between different cores
 * Instead of only testing the 
 */
package main
import (
    "fmt"
    "time"
    "sync"
)

const max_test = 100000000

// Set all indexes is map which is
func removeCandidates(n int, prime_map *[max_test+1]bool, wg *sync.WaitGroup){
    for i := n*2; i < max_test+1; i+=n {
        prime_map[i] = true
    }
    wg.Done()
}

func countTrue(in [max_test+1]bool) (int) {
    count := 0
    for _, n := range in {
       if (!n) {
           //fmt.Printf("%d is prime!\n", i)
           count++
       }
    }
    return count
}

func setConventions(in *[max_test+1]bool) {
    in[0] = true
    in[1] = true
}



func main() {
    start := time.Now()
    var dividables [max_test+1]bool
    setConventions(&dividables)
    var wg sync.WaitGroup
    for i := 2; i <= max_test; i+=1 {
        wg.Add(1)
        go removeCandidates(i, &dividables, &wg)
    }
    wg.Wait()
    count := countTrue(dividables) 
    elapsed := time.Since(start)
    fmt.Printf("Found %d primes in %s \n", count + 1, elapsed)
}
