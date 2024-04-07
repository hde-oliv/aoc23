package main

import "fmt"

type Race struct {
	time     int
	distance int
}

func check_victory(distance int, charge int, time int) bool {
	r_time := time - charge
	progress := r_time * charge

	if progress > distance {
		return true
	}
	return false
}

func calculate_wins(r *Race) int {
	var total int

	for i := 0; i <= r.time; i++ {
		result := check_victory(r.distance, i, r.time)

		if result {
			total++
		}
	}

	return total
}

func main() {
	// Part 01
	r01 := Race{time: 3, distance: 4}
	r02 := Race{time: 8, distance: 1}
	r03 := Race{time: 4, distance: 2}
	r04 := Race{time: 8, distance: 1}

	w01 := calculate_wins(&r01)
	w02 := calculate_wins(&r02)
	w03 := calculate_wins(&r03)
	w04 := calculate_wins(&r04)

	fmt.Println(w01 * w02 * w03 * w04)

	// Part 02
	r_bonus := Race{time: 6, distance: 4}
	w_bonus := calculate_wins(&r_bonus)

	fmt.Println(w_bonus)
}
