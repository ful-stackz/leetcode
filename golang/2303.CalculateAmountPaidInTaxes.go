package leetcode

func calculateTax(brackets [][]int, income int) float64 {
	if income == 0 {
		return 0
	}

	taxes := 0.0
	remainingIncome := float64(income)

	for i, tax := range brackets {
		bracket := float64(tax[0])
		rate := float64(tax[1]) / 100

		maxTaxable := bracket
		if i != 0 {
			maxTaxable -= float64(brackets[i-1][0])
		}

		taxable := min(maxTaxable, remainingIncome)
		tax := taxable * rate
		taxes += tax
		remainingIncome -= taxable
		if remainingIncome <= 0 {
			break
		}
	}

	return taxes
}

func min(a, b float64) float64 {
	if a <= b {
		return a
	}
	return b
}
