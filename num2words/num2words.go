package num2words

var units = [20]string{"", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine",
                       "ten", "eleven", "twelve", "thirteen", "forteen", "fifteen","sixteen","seventeen", "eighteen","nineteen"}
var twenties = [10]string{"","","twenty","thirty","forty", "fifty","sixty","seventy","eighty","ninety"}
var thousands = [3]string{"thousand", "lakh", "crore"}

/*
var thousands = [22]string{"hundred","thousand","million", "billion", "trillion", "quadrillion", "quintillion", "sextillion",
						   "septillion","octillion", "nonillion", "decillion", "undecillion", "duodecillion", "tredecillion", 
						   "quattuordecillion", "quindecillion", "sexdecillion", "septendecillion", "octodecillion", 
						   "novemdecillion", "vigintillion"}

func ToWords(num int) (result string) {
	if num>=0 && num<10 {
		result = units[num]
	} else {
		result = "Unknown Number"
	}
	return 
}
*/
func wordsTill100(num int) string {
	d := num % 100
	if d < 20 {
		return units[d]
	}
	d1 := num % 10
	d2 := num / 10
	d2 = d2 % 10
	if d1 != 0 {
		return twenties[d2] + "-" + units[d1]
	}
	return twenties[d2]
}

func ToWords(number int) string {
	result := ""
	num := number
	if num == 0 {
		return "zero"
	}
	if num > 0 && num < 1000000000 {
		result = wordsTill100(num)
		num = num / 100
		d := num % 10
		if d != 0 {
			result = units[d] + " hundred " + result
		} 
		num = num / 10
		for i := 0; num != 0; i++ {
			if num % 100 != 0 {
				result = wordsTill100(num) + " " + thousands[i] + " " + result
			}
			num = num / 100
			if num == 0 {
				break
			}
		}
		return result
	} 
	return "Unknown Number"
}
