// // // // package main

// // // // import (
// // // // 	"fmt"
// // // // 	"time"
// // // // )

// // // // // import (
// // // // // 	"fmt"
// // // // // 	"log"
// // // // // )

// // // // // // func main() {
// // // // // //     a := []int{1, 3, 4, 5, 6, 7, 10, 10}

// // // // // //     countMap := make(map[int]int)

// // // // // //     for _, n := range a {
// // // // // //         countMap[n]++
// // // // // //     }

// // // // // //     log.Println(countMap)

// // // // // //     uniqueNumbers := []int{}
// // // // // //     for key := range countMap {
// // // // // //         uniqueNumbers = append(uniqueNumbers, key)
// // // // // //     }

// // // // // //     log.Println("Unique numbers:", uniqueNumbers)
// // // // // // 	go UniqueCheck()
// // // // // // }

// // // // // func checkNumbers(ch <-chan int, n int) {
// // // // // 	log.Println("N>>>>>>>>.", n)
// // // // // 	for value := range ch {
// // // // // 		fmt.Println("Received:", value)
// // // // // 	}

// // // // // }
// // // // // func main() {
// // // // // 	ch := make(chan int)
// // // // // 	slice := []int{2, 3, 45, 5, 6}
// // // // // 	for _, n := range slice {
// // // // // 		go checkNumbers(ch, n)
// // // // // 	}

// // // // // }

// // // // func main() {
// // // // 	c := make(chan int)
// // // // 	i := 0
// // // // 	slice := []int{2,3,4,5,5,7}
// // // // 	for _, n := range slice {
// // // // 		go goroutine(n)
// // // // 	}

// // // // 	c <- i

// // // // }

// // // // func goroutine(c chan int) {
// // // // 	for {
// // // // 		num := <-c
// // // // 		fmt.Println(num)
// // // // 		num++
// // // // 		time.Sleep(1 * time.Second)
// // // // 		c <- num
// // // // 	}
// // // // }

// // // package main

// // // import (
// // // 	"fmt"
// // // 	"log"
// // // 	"time"
// // // )

// // // func main() {
// // // 	c := make(chan int)
// // // 	slice := []int{2, 3, 4, 5, 6, 7}

// // // 	for _, n := range slice {
// // // 		_ = n
// // // 		go goroutine(c, n)
// // // 	}

// // // 	c <- 0

// // // 	time.Sleep(10 * time.Second)
// // // }

// // // func goroutine(c chan int, n int) {
// // // 	for {
// // // 		log.Println("N>>>>>>", n)
// // // 		num := <-c
// // // 		fmt.Println(num)
// // // 		num++
// // // 		time.Sleep(1 * time.Second)
// // // 		c <- num
// // // 	}
// // // }
// // package main

// // import (
// // 	"fmt"
// // 	"time"
// // )

// // func main() {
// // 	slice := []int{2, 3, 4, 5, 6, 7}
// // 	for _, n := range slice {
// // 		c := make(chan int)
// // 		go goroutine(c, n)
// // 		c <- 0
// // 	}

// // 	// time.Sleep(10 * time.Second)
// // }

// // func goroutine(c chan int, n int) {
// // 	for {
// // 		num := <-c
// // 		countMap := make(map[int]bool)
// // 		countMap[num]= false
// // 		uniqueNumbers := make(map[int]int)
// // 		if countMap[num] == false{
// // 			uniqueNumbers = append(uniqueNumbers, num)
// // 		}
// // 		fmt.Printf("Goroutine %d received: %d\n", n, num)
// // 		num++
// // 		time.Sleep(1 * time.Second)
// // 		c <- num
// // 	}
// // }

// package main

// import (
// 	"fmt"
// 	"log"
// 	"time"
// )

// func main() {
// 	slice := []int{2, 3, 4, 5, 6, 7}
// 	for _, n := range slice {
// 		c := make(chan int)
// 		go goroutine(c, n)
// 		c <- 0
// 	}

// 	time.Sleep(10 * time.Second)
// }

// func goroutine(c chan int, n int) {
// 	uniqueNumbers := make(map[int]bool)
// 	for {
// 		num := <-c
// 		if !uniqueNumbers[num] {
// 			uniqueNumbers[num] = true
// 			fmt.Printf("Goroutine %d received a unique number: %d\n", n, num)
// 		} else {
// 			fmt.Printf("Goroutine %d received a duplicate number: %d\n", n, num)
// 		}
// 		log.Println("", uniqueNumbers)
// 		num++
// 		time.Sleep(1 * time.Second)
// 		c <- num
// 	}
// }

package main

import (
	"fmt"
	"log"
	"os"
)

func UniqueConc(nums <-chan int, result chan<- int) {
	m := make(map[int]struct{})

	for v := range nums {
		m[v] = struct{}{}
		log.Println("m>>", m)
	}

	for k := range m {
		result <- k
	}

	close(result)
}

func main() {
	fmt.Println(os.Args)
	ch := make(chan int)

	resultCh := make(chan int)

	nums := []int{1, 2, 4, 5, 6, 76, 7, 7, 8}

	go UniqueConc(ch, resultCh)

	for i := 0; i < len(nums); i++ {
		ch <- nums[i]
	}
	close(ch)

	slice := []int{}
	for n := range resultCh {
		slice = append(slice, n)
	}
	log.Println("Slice>>>>>>>>", slice)
	// close(resultCh)
}
