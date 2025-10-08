package main

import "fmt"

func denseRanking(scores []int, gits []int) []int {
	// Buat daftar skor unik (tanpa duplikat)
	unique := []int{}
	for i, s := range scores {
		if i == 0 || s != scores[i-1] {
			unique = append(unique, s)
		}
	}

	n := len(unique)
	result := []int{}
	i := n - 1

	// Hitung ranking GITS
	for _, g := range gits {
		for i >= 0 && g >= unique[i] {
			i--
		}
		result = append(result, i+2)
	}
	return result
}

func printFormatted(scores []int, gits []int) {
	fmt.Println("Sampel Input:")
	fmt.Println(len(scores))
	for _, s := range scores {
		fmt.Printf("%d ", s)
	}
	fmt.Println("\n", len(gits))
	for _, g := range gits {
		fmt.Printf("%d ", g)
	}
	fmt.Println("\n\nSampel Output:")
	for _, rank := range denseRanking(scores, gits) {
		fmt.Printf("%d ", rank)
	}
}

func main() {
	scores1 := []int{100, 100, 50, 40, 40, 20, 10}
	gits1 := []int{5, 25, 50, 120}
	printFormatted(scores1, gits1)

	scores2 := []int{90, 80, 80, 75, 60}
	gits2 := []int{50, 65, 77, 90, 102}
	printFormatted(scores2, gits2)

	scores3 := []int{200, 150, 150, 100, 90}
	gits3 := []int{70, 120, 150, 210}
	printFormatted(scores3, gits3)
}