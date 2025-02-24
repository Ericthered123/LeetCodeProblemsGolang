package problems

func TwoSumMap(nums []int, target int) []int {

	m := make(map[int]int)

	for i, num := range nums {
		complement := target - num //Calcula el complemento

		if index, found := m[complement]; found { //Busca el complemento en el mapa, porque si lo encuentra eso significa
			// que complement + num = target
			return []int{index, i} //retorna resultado

		}
		m[num] = i // guarda en el mapa
	}
	return nil
}
