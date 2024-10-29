package billing

import "math/big"

func calculate(val *big.Int) (c int) {
	multipliers := []int{2, 3, 4, 5, 6, 7}
	sum := big.NewInt(0)
	position := 0

	tempNumber := new(big.Int).Set(val)
	mod := big.NewInt(10)
	zero := big.NewInt(0)

	for tempNumber.Cmp(zero) > 0 {
		digit := new(big.Int).Mod(tempNumber, mod)
		multiplier := int64(multipliers[position%len(multipliers)])
		digit.Mul(digit, big.NewInt(multiplier))
		sum.Add(sum, digit)
		tempNumber.Div(tempNumber, mod)
		position++
	}
	remainder := new(big.Int).Mod(sum, big.NewInt(11))

	if remainder.Cmp(big.NewInt(0)) == 0 || remainder.Cmp(big.NewInt(1)) == 0 {
		return 0
	}

	control := new(big.Int).Sub(big.NewInt(11), remainder)
	return int(control.Int64())
}

func countDigits(num *big.Int) int {
	count := 0
	tmp := new(big.Int).Set(num)

	for tmp.Sign() != 0 {
		tmp.Div(tmp, big.NewInt(10))
		count++
	}
	return count
}

func joinNumbers(num1, num2 *big.Int) *big.Int {
	numDigits := countDigits(num2)

	multiplier := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(numDigits)), nil)

	result := new(big.Int).Mul(num1, multiplier)
	result.Add(result, num2)

	return result
}

func splitLastDigit(number *big.Int) (*big.Int, int) {
	remainder := new(big.Int)

	quotient, _ := new(big.Int).DivMod(number, big.NewInt(10), remainder)

	return quotient, int(remainder.Int64())
}

func checkBillID(billID *big.Int, control int) bool {
	return calculate(billID) == control
}

func checkPayID(payID *big.Int, control int) bool {
	return calculate(payID) == control
}

func checkBill(billID, payID *big.Int) bool {
	var control int

	billIDQuotient, billIDControl := splitLastDigit(billID)
	billIdIsValid := checkBillID(billIDQuotient, billIDControl)

	payIDQuotient, control := splitLastDigit(payID)
	payIdIsValid := checkPayID(splitLastDigit(payIDQuotient))

	bill := joinNumbers(billID, payIDQuotient)

	return billIdIsValid && payIdIsValid && calculate(bill) == control
}
