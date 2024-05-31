package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func primes(n int) []int {
	is_prime := make([]bool, n+1)
	for i := range is_prime {
		is_prime[i] = true
	}
	primes := make([]int, 0)
	for p := 2; p*p <= n; p++ {
		if is_prime[p] {
			for i := p * p; i <= n; i += p {
				is_prime[i] = false
			}
		}
	}
	for p := 2; p <= n; p++ {
		if is_prime[p] {
			primes = append(primes, p)
		}
	}
	return primes
}

func buildNewFromOld(prime []int, bit int) int {
	rand.Seed(time.Now().UnixNano())
	var q, maxIndex, p int
	for maxIndex = 0; prime[maxIndex] < int(math.Pow(2, float64(bit/2))); maxIndex++ {
	}
	for {
		q = prime[rand.Intn(maxIndex+1)]
		if q > int(math.Pow(2, float64(bit/2)-1)) && q <= int(math.Pow(2, float64(bit/2))-1) {
			break
		}
	}
	for {
		n := (math.Pow(2, float64(bit-1))/float64(q) + math.Pow(2, float64(bit-1))*rand.Float64()/float64(q))
		if int(n)%2 == 1 {
			n += 1
		}
		for u := 0; true; u += 2 {
			p = int((n+float64(u))*float64(q) + 1)
			if p > int(math.Pow(2, float64(bit))) {
				break
			}
			if powerMod(2, p-1, p) == 1 && powerMod(2, int(n+float64(u)), p) != 1 {
				return p
			}
		}
	}
}

func powerMod(a, b, n int) int {
	result := 1
	for b > 0 {
		if b%2 == 1 {
			result = (result * a) % n
		}
		a = (a * a) % n
		b /= 2
	}
	return result
}

func rnInt(a, b int) int {
	return rand.Intn(b-a+1) + a
}

func rnDouble(a, b int) float64 {
	return float64(rand.Intn(b-a+1)) + float64(a) + rand.Float64()
}

func printRes(res []int) {
	for i, v := range res {
		fmt.Printf("%d\t\t|\t\t%d\n", i+1, v)
	}
}

func main() {
	prime := primes(500)
	var bit int
	fmt.Scan(&bit)
	res := make([]int, 0, 10)
	for len(res) != 10 {
		p := buildNewFromOld(prime, bit)
		if !contains(res, p) {
			res = append(res, p)
		}
	}
	printRes(res)
}

func contains(slice []int, value int) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}
