package prime

func CheckPrime(num int) bool {
	return pmap[num]
}

var pmap map[int]bool

func InitPmap(top int) {
	pmap = make(map[int]bool)
	primes := sieveOfEratosthenes(top)
	for _, v := range primes {
		pmap[v] = true
	}
}

func sieveOfEratosthenes(N int) (primes []int) {
	b := make([]bool, N)
	for i := 2; i < N; i++ {
		if b[i] == true {
			continue
		}
		primes = append(primes, i)
		for k := i * i; k < N; k += i {
			b[k] = true
		}
	}
	return
}
