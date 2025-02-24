package problems

func IntToRoman(num int) string {
	if num >= 1 && num <= 3999 {

		ints := map[int]string{
			1: "I", 4: "IV", 5: "V", 9: "IX", 10: "X",
			40: "XL", 50: "L", 90: "XC", 100: "C",
			400: "CD", 500: "D", 900: "CM", 1000: "M",
		}

		values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}

		result := ""

		for i := 0; i < len(values); i++ {
			for num >= values[i] {
				num -= values[i]
				result += ints[values[i]]
			}

		}
		return result
	}
	return ""
}
