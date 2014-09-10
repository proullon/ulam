package prime

import (
"math"
"errors"
"fmt"
)


func IsPrime(n int64) bool {

	if n == 1 || n == 3 {
		return true
	}

	if n % 2 == 0 {
		return false
	}

	sqrt := math.Sqrt(float64(n))

	var i int64
	for i = 3; i <= int64(sqrt); i+=2 {
		if n % i == 0 {
			return false
		}
	}

	return true
}

// Size of square
func Ulam(size int64) (spiral [][]bool, err error) {
	var i int64

	if size % 2 != 1 {
		err = errors.New("Size must be not even")
		return
	}

	// Initialize spiral
	spiral = make([][]bool, size)
	for it := range spiral {
		spiral[it] = make([]bool, size)
	}

	// Find spiral center
	y := size / 2
	x := size / 2

	offset_x := 1 
	offset_y := 1
	total := size * size

	// fmt.Printf("x:%d y:%d\n", x, y)

	// Until all square is calculated
	for i = 1; i < total;{

		// Go up
		// fmt.Printf("up!\n")
		for it_y := 0; it_y < offset_y; it_y++ {
			if i == total {
				break
			}

			y++
			// fmt.Printf("x:%d y:%d i:%d\n", x, y, i)
			spiral[y][x] = IsPrime(i)
			i++
		}
		offset_y++

		// Go right
		// fmt.Printf("right!\n")
		for it_x := 0; it_x < offset_x; it_x++ {
			if i == total {
				break
			}

			x++
			spiral[y][x] = IsPrime(i)
			// fmt.Printf("x:%d y:%d i:%d\n", x, y, i)
			i++
		}
		offset_x++

		// Go down
		// fmt.Printf("down!\n")
		for it_y := offset_y; it_y > 0; it_y-- {
						if i == total {
				break
			}

			y--
			spiral[y][x] = IsPrime(i)
			// fmt.Printf("x:%d y:%d i:%d\n", x, y, i)
			i++
		}		
		offset_y++

		// Go left
		// fmt.Printf("left!\n")
		for it_x := offset_x; it_x > 0; it_x-- {
			if i == total {
				break
			}

			x--
			spiral[y][x] = IsPrime(i)
			// fmt.Printf("x:%d y:%d i:%d\n", x, y, i)
			i++
		}
		offset_x++
	} 


	return
}

type Point struct {
	x int64
	y int64
	prime bool

}

func _isprime(n int64, x int64, y int64, res chan Point) {
	p := Point{x:x, y:y, prime:IsPrime(n)}

	res <- p
}

// Size of square
func UlamParallel(size int64) (spiral [][]bool, err error) {
	var i int64
	res := make(chan Point)

	if size % 2 != 1 {
		err = errors.New("Size must be not even")
		return
	}

	// Initialize spiral
	spiral = make([][]bool, size)
	for it := range spiral {
		spiral[it] = make([]bool, size)
	}

	// Find spiral center
	y := size / 2
	x := size / 2

	offset_x := 1 
	offset_y := 1
	total := size * size

	// fmt.Printf("x:%d y:%d\n", x, y)
	go _isprime(i, x, y, res)

	// Until all square is calculated
	for i = 1; i < total;{

		// Go up
		// fmt.Printf("up!\n")
		for it_y := 0; it_y < offset_y; it_y++ {
			if i == total {
				break
			}

			y++
			// fmt.Printf("x:%d y:%d i:%d\n", x, y, i)
			go _isprime(i, x, y, res)
			i++
		}
		offset_y++

		// Go right
		// fmt.Printf("right!\n")
		for it_x := 0; it_x < offset_x; it_x++ {
			if i == total {
				break
			}

			x++
			go _isprime(i, x, y, res)
			// fmt.Printf("x:%d y:%d i:%d\n", x, y, i)
			i++
		}
		offset_x++

		// Go down
		// fmt.Printf("down!\n")
		for it_y := offset_y; it_y > 0; it_y-- {
						if i == total {
				break
			}

			y--
			go _isprime(i, x, y, res)
			// fmt.Printf("x:%d y:%d i:%d\n", x, y, i)
			i++
		}		
		offset_y++

		// Go left
		// fmt.Printf("left!\n")
		for it_x := offset_x; it_x > 0; it_x-- {
			if i == total {
				break
			}

			x--
			go _isprime(i, x, y, res)
			// fmt.Printf("x:%d y:%d i:%d\n", x, y, i)
			i++
		}
		offset_x++
	}


	var p Point
	for i = 0; i < total; i++{
		p = <- res
		fmt.Printf("%d %d : %v\n", p.x, p.y, p.prime)
		spiral[p.y][p.x] = p.prime
	}

	return
}