package main

import (
	bu "bufio"
	"fmt"
	"os"
	"strconv"
	stg "strings"
	"sort"
	"sync"
)

func subSort(n []int, c chan []int, wg *sync.WaitGroup) {
	fmt.Println(n)
	sort.Ints(n)
	c <- n
	wg.Done()
}

func main()  {
	reader := bu.NewReader(os.Stdin)

	fmt.Print("Enter integers separated by spaces:")

	input, _ := reader.ReadString('\n')
	input = stg.TrimSuffix(input, "\n")
	input = stg.ToLower(input)

	numbers :=  stg.Split(input, " ")

	var num []int
	for _, s := range(numbers) {
		n, _ := strconv.Atoi(s)
		num = append(num, n)
	}
	size := len(num) / 4

	fmt.Println("Num:", num,"Size:", size)

	channel := make(chan []int, 4)
	var wg sync.WaitGroup

	wg.Add(4)

	go subSort(num[:size], channel, &wg)
	go subSort(num[size:size*2], channel, &wg)
	go subSort(num[size*2:size*3], channel, &wg)
	go subSort(num[size*3:], channel, &wg)

	wg.Wait()
	part1 := <- channel
	fmt.Println("Part 1:", part1)
	part2 := <- channel
	fmt.Println("Part 2:", part2)
	part3 := <- channel
	fmt.Println("Part 3:", part3)
	part4 := <- channel
	fmt.Println("Part 4:", part4)

	var sorted []int
	for i := 0; i < len(part1); i++ {
		sorted = append(sorted, part1[i])
	}
	for i := 0; i < len(part2); i++ {
		sorted = append(sorted, part2[i])
	}
	for i := 0; i < len(part3); i++ {
		sorted = append(sorted, part3[i])
	}
	for i := 0; i < len(part4); i++ {
		sorted = append(sorted, part4[i])
	}

	fmt.Println("No sorted:", sorted)
	sort.Ints(sorted)
	fmt.Println("Sorted:", sorted)
}