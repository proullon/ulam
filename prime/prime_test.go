package prime

import (
	"testing"
	"fmt"
)

func TestIsPrime(t *testing.T) {
	if IsPrime(5) == false {
		t.Errorf("5 is prime")
		return
	}

	if IsPrime(19) == false {
		t.Errorf("19 is prime")
		return
	}

	if IsPrime(15) == true {
		t.Errorf("15 is not prime")
		return
	}
}

func TestUlam(t *testing.T) {
	spiral, err := Ulam(20)
	if err == nil {
		t.Errorf("Ulam(20) must return an error")
	}

	spiral, _ = Ulam(31)
	for i := range spiral {
		for y := range spiral[i] {
			if spiral[i][y] == true {
				fmt.Printf("x")
			} else {
				fmt.Printf("o")
			}
		}
		fmt.Printf("\n")
	}
}

func TestUlamParallel(t *testing.T) {

	spiral, _ := UlamParallel(5)
	for i := range spiral {
		for y := range spiral[i] {
			if spiral[i][y] == true {
				fmt.Printf("x")
			} else {
				fmt.Printf("o")
			}
		}
		fmt.Printf("\n")
	}
}

func BenchmarkBigIsPrime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPrime(234323423)
	}
}

func BenchmarkBigUlam(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Ulam(1001)
	}
}

func BenchmarkBigUlamParallel(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Ulam(1001)
	}
}
