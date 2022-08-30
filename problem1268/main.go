package problem1268

func suggestedProducts(products []string, searchWord string) [][]string {
	sortByDictOrder(products)
	result := make([][]string, len(searchWord))
	for i := 0; i < len(searchWord); i++ {
		result[i] = make([]string, 0)
	}

	for _, product := range products {
		if len(product) >= 1 && product[0] == searchWord[0] {
			result[0] = append(result[0], product)
		}
	}

	for i := 1; i < len(searchWord); i++ {
		for _, product := range result[i-1] {
			if len(product) > i && product[i] == searchWord[i] {
				result[i] = append(result[i], product)
			}
		}
	}

	for i := 0; i < len(result); i++ {
		if len(result[i]) > 3 {
			result[i] = result[i][:3]
		}
	}
	return result
}

func sortByDictOrder(products []string) {
	for i := 0; i < len(products); i++ {
		for j := i + 1; j < len(products); j++ {
			if products[i] > products[j] {
				products[i], products[j] = products[j], products[i]
			}
		}
	}
}
