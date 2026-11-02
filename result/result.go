package result

func CalculateTotal(marks []int) int {
	total := 0

	for _, mark := range marks {
		total += mark
	}

	return total
}

func CalculateAverage(total int, count int) float64 {
	return float64(total) / float64(count)
}

func CalculateGrade(average float64) string {
	if average >= 90 {
		return "A+"
	} else if average >= 80 {
		return "A"
	} else if average >= 70 {
		return "B"
	} else if average >= 60 {
		return "C"
	} else if average >= 50 {
		return "D"
	}

	return "F"
}

func GetResult(marks []int) string {
	for _, mark := range marks {
		if mark < 40 {
			return "FAIL"
		}
	}

	return "PASS"
}