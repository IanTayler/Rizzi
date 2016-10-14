package mijn

func Exp (x, y int) int { 
	if y < 0 {
		x = 1/x
		y = -y
	} else if y == 0 {
		return 1
	}
	
	z := 1
	
	for y > 1 {
		if y % 2 == 0 {
			x = x*x
			y = y/2
		} else {
			z = x * z
			x = x * x
			y = (y - 1) / 2
		}
		
	}
	
	return x * z
}

func IsBlank(b byte) bool {
	switch b {
		case ' ', '\n', '\t':
			return true
			
		default:
			return false
	}
}

func IsOp(b byte) bool {
	switch b {
		case '+', '-', '*', '/', 'n', 'r', 'e':
			return true
			
		default:
			return false
	}
}

func OpToFunc(s string) func(int, int) int {
	switch s {
		case "+": 	return func(x, y int) int { return x + y }
		case "-": 	return func(x, y int) int { return x - y }
		case "*":	return func(x, y int) int { return x * y }
		case "/":	return func(x, y int) int { return x / y }
		case "rem":	return func(x, y int) int { return x % y }
		
		case "exp": return Exp
		
		default:	return func(x, y int) int { return 1	 }
	}
}

func UnOpToFunc(s string) func(int) int {
	switch s {
		case "neg":	return func(x int) int { return -x }
		default:	return func(x int) int { return 0  }
	}
}
