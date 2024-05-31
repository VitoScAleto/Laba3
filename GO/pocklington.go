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

func builderTest(prime []int, bit int) (int, []int) {
	rand.Seed(time.Now().UnixNano())
	maxIndex, maxPow := 0, 1
	for ; maxIndex < len(prime) && prime[maxIndex] < int(math.Pow(2, float64((bit/2)+1))); maxIndex++ {
	}
	if maxIndex == len(prime) {
		return 0, nil
	}
	for ; int(math.Pow(2, float64(maxPow))) < int(math.Pow(2, float64((bit/2)+1))); maxPow++ {
	}
	f := 1
	q := make([]int, 0)
	for {
		num, power := rand.Intn(maxIndex+1), rand.Intn(maxPow)+1
		if int(math.Pow(float64(prime[num]), float64(power))) < math.MaxInt32 {
			f *= int(math.Pow(float64(prime[num]), float64(power)))
			q = append(q, prime[num])
		}
		if f > int(math.Pow(2, float64(bit/2))) {
			if f >= int(math.Pow(2, float64((bit/2)+1))) {
				f = 1
				q = q[:0]
			} else {
				break
			}
		}
	}
	R := rn(int(math.Pow(2, float64((bit/2)-1)))+1, int(math.Pow(2, float64(bit/2))))
	for R%2 != 0 {
		R = rn(int(math.Pow(2, float64((bit/2)-1)))+1, int(math.Pow(2, float64(bit/2))))
	}
	n := R*f + 1
	return n, q
}

func testPoklin(n, t int, q []int) int {
	a := make([]int, 0, t)
	for len(a) != t {
		aj := rn(2, n-1)
		if !contains(a, aj) {
			a = append(a, aj)
		}
	}
	for _, aj := range a {
		if powerMod(aj, n-1, n) != 1 {
			return 0
		}
	}
	isValidALLResultOne := true
	i := 0
	for _, aj := range a {
		if i < len(q) && q[i] != 0 && powerMod(aj, (n-1)/q[i], n) == 1 {
			isValidALLResultOne = false
			return 0
		}
		i++
	}
	if isValidALLResultOne {
		return 1
	}
	return 1
}

func powerMod(a, b, n int) int {
	var result int64 = 1
	for b > 0 {
		if b%2 == 1 {
			result = (result * int64(a)) % int64(n)
		}
		a = (a * a) % n
		b /= 2
	}
	return int(result)
}

func rn(a, b int) int {
	return rand.Intn(b-a+1) + a
}

func contains(slice []int, value int) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

func printResults(res []int, resVerTest []string, otvegnutie []int) {
	fmt.Println("Prime Numbers\tTest Results\tOccurrences")
	fmt.Println("----------------------------------------------")
	for i, v := range res {
		fmt.Printf("%d\t\t%s\t\t%d\n", v, resVerTest[i], otvegnutie[i])
	}
}

func main() {
	prime := primes(500)
	var bit int
	fmt.Scan(&bit)
	if bit > 2*len(prime) {
		fmt.Println("Error: bit value is too large for the current prime list size.")
		return
	}
	res := make([]int, 0, 10)
	resVerTest := make([]string, 0, 10)
	otvegnutie := make([]int, 0, 10)
	k := 0
	for len(res) != 10 {
		n, q := builderTest(prime, bit)
		if n == 0 {
			fmt.Printf("Error: unable to generate suitable number with the given bit value (%d).\n", bit)
			return
		}
		probability := testPoklin(n, 10, q)
		if probability == 1 {
			if !contains(res, n) {
				res = append(res, n)
				probability = testPoklin(n, 1, q)
				if probability == 1 {
					resVerTest = append(resVerTest, "+")
				} else {
					resVerTest = append(resVerTest, "-")
				}
				otvegnutie = append(otvegnutie, k)
				k = 0
			}
		} else {
			k++
		}
	}
	printResults(res, resVerTest, otvegnutie)
}
