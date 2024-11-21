package main

import (
	"crypto/sha256"
	"encoding/binary"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"
)

// modExp calculates (base^exp) % mod
func modExp(base, exp, mod int) int {
	result := 1
	base = base % mod
	for exp > 0 {
		if exp%2 == 1 { // If exp is odd
			result = (result * base) % mod
		}
		exp = exp / 2
		base = (base * base) % mod
	}
	return result
}

// generateHashE generates a hash value for r and message
func generateHashE(r int, message string) int {
	hasher := sha256.New()
	hasher.Write([]byte(fmt.Sprintf("%d%s", r, message)))
	hash := hasher.Sum(nil)

	// Convert the first 4 bytes of the hash to an integer
	return int(binary.BigEndian.Uint32(hash))
}

// findGenerators finds valid generators for given p and q
func findGenerators(p, q int) []int {
	validGenerators := []int{}
	for g := 2; g < p; g++ {
		if modExp(g, q, p) == 1 {
			validGenerators = append(validGenerators, g)
		}
	}
	return validGenerators
}

// validate checks if q divides (p-1) and prints valid generators
func validate(p, q int) {
	if (p-1)%q != 0 {
		fmt.Println("Error: q does not divide p-1")
		os.Exit(1)
	}
	validGenerators := findGenerators(p, q)
	fmt.Printf("Valid generators for p = %d, q = %d: %v\n", p, q, validGenerators)
}

func main() {
	maxK := flag.Int("maxK", 10000, "Maximum value of K")
	p := flag.Int("p", 48731, "Value of p")
	q := flag.Int("q", 443, "Value of q")
	g := flag.Int("g", 11444, "Value of g")
	message := flag.String("m", "TRALALA", "Message to be processed")
	shouldValide := flag.Bool("v", false, "Validate")
	flag.Parse()

	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	if *shouldValide {
		validate(*p, *q)
	}

	fmt.Println("GENERATE KEYS")
	// private key x: random [1, q-1]
	x := r.Intn(*q-1) + 1

	// public key y = g^(q-x) mod p
	y := modExp(*g, *q-x, *p)
	fmt.Printf("(p=%d, q=%d, g=%d, x=%d, y=%d)\n", *p, *q, *g, x, y)

	fmt.Println("AUTHENTICATE")
	// random number k
	k := r.Intn(*maxK)

	// rA = g^k mod p
	rA := modExp(*g, k, *p)
	fmt.Printf("rA = %d\n", rA)

	// e = H(rA, message)
	e := generateHashE(rA, *message)
	fmt.Printf("e = %d\n", e)

	// s = k + x*eA
	s := (k + x*e)
	fmt.Printf("s = %d\n", s)

	fmt.Println("VERIFY AUTHENTICATION")
	// rB = (g^s * y^eA) mod p
	rB := (modExp(*g, s, *p) * modExp(y, e, *p)) % *p
	fmt.Printf("rB = %d\n", rB)

	// eB = H(rB, message)
	eB := generateHashE(rB, *message)
	fmt.Printf("eB = %d\n", eB)

	// verify authentication
	if e == eB {
		fmt.Println("Bob verifies Alice's identity\ne == eB")
	} else {
		fmt.Println("Bob cannot verify Alice's identity\ne != eB")
	}
}
