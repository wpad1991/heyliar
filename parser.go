package heyliar

import (
	"math"
	"strconv"
	"strings"
)

// for expressing equality of variables
const epsilon = 1e-14

// Totalparsing is handles many expression and operators.
// example : Totalparsing("{10<5}or{1>0}") => return true
// example : Totalparsing("{10<5}and{1>0}") => return false
func Totalparsing(s string) bool {

	s = strings.Replace(s, " ", "", -1)
	s = strings.Replace(s, "\t", "", -1)

	split := sliceexpression(s)

	postsplit := intopostfixexpression(split)

	return calcpostexpression(postsplit)
}

// Calcliar is processes a single comparison expression
// and returns a boolean value.
// example : Calcliar("3*5==15") => return true
// example : Calcliar("3*5!=15") => return true
func Calcliar(s string) bool {

	var cword string

	if strings.Contains(s, ">=") {
		cword = ">="
	} else if strings.Contains(s, "<=") {
		cword = "<="
	} else if strings.Contains(s, "==") {
		cword = "=="
	} else if strings.Contains(s, "!=") {
		cword = "!="
	} else if strings.Contains(s, ">") {
		cword = ">"
	} else if strings.Contains(s, "<") {
		cword = "<"
	}

	split := strings.Split(s, cword)

	if len(split) != 2 {
		panic("Calcliar invalid argument : " + s)
	}

	return compareexpression(Calculate(split[0]), Calculate(split[1]), cword)
}

// Calculate is claculate expression
// example : Calculate("(3*2)+3-2*(2/2)") return => 7
func Calculate(s string) float64 {

	split := sliceoper(s)
	splitpost := intopostfixcalc(split)
	return calcpostfix(splitpost)
}

func compareexpression(val1 float64, val2 float64, cword string) bool {

	if cword == ">=" {
		if math.Abs(val1-val2) <= epsilon {
			return true
		}
		if val1 > val2 {
			return true
		}
	} else if cword == "<=" {
		if math.Abs(val1-val2) <= epsilon {
			return true
		}
		if val1 < val2 {
			return true
		}
	} else if cword == "==" {
		if math.Abs(val1-val2) <= epsilon {
			return true
		}
	} else if cword == "!=" {
		if math.Abs(val1-val2) > epsilon {
			return true
		}
	} else if cword == ">" {
		if math.Abs(val1-val2) <= epsilon {
			return false
		}
		if val1 > val2 {
			return true
		}
	} else if cword == "<" {
		if math.Abs(val1-val2) <= epsilon {
			return false
		}
		if val1 < val2 {
			return true
		}
	}
	return false
}

func sliceexpression(s string) []string {

	var rs []string

	s = strings.ToLower(s)

	sSize := len(s)

	for i := 0; i < sSize; i++ {

		sValue := string(s[i])
		switch sValue {
		case "(", ")":
			rs = append(rs, sValue)
		case "{":
			index := strings.Index(s[i:sSize], "}")
			slice := s[i+1 : i+index]
			rs = append(rs, slice)
			i = i + index
		case " ":
		case "a":
			if s[i:i+3] == "and" {
				rs = append(rs, s[i:i+3])
				i = i + 2
			} else {
				panic("sliceexpression argument invalid : " + s)
			}
		case "o":
			if s[i:i+2] == "or" {
				rs = append(rs, s[i:i+2])
				i = i + 1
			} else {
				panic("sliceexpression argument invalid : " + s)
			}
		default:
			panic("sliceexpression argument invalid : " + s)
		}

	}

	return rs
}

func calcpostexpression(split []string) bool {
	if split == nil {
		panic("calcpostexpression argument is null")
	}

	sSize := len(split)

	if sSize == 0 {
		panic("calcpostexpression argument size is null")
	}

	st := Stackstring{}
	var val1 string
	var val2 string

	for _, value := range split {

		switch value {
		case "":
		case "and":
			val2 = st.Pop()
			val1 = st.Pop()

			var b1 bool
			var b2 bool
			var err error
			if b1, err = strconv.ParseBool(val1); err != nil {
				b1 = Calcliar(val1)
			}

			if b2, err = strconv.ParseBool(val2); err != nil {
				b2 = Calcliar(val2)
			}

			st.Push(strconv.FormatBool(b1 && b2))

		case "or":
			val2 = st.Pop()
			val1 = st.Pop()

			var b1 bool
			var b2 bool
			var err error
			if b1, err = strconv.ParseBool(val1); err != nil {
				b1 = Calcliar(val1)
			}

			if b2, err = strconv.ParseBool(val2); err != nil {
				b2 = Calcliar(val2)
			}

			st.Push(strconv.FormatBool(b1 || b2))

		case " ":
		default:
			st.Push(value)
		}
	}

	if st.GetSize() != 1 {
		panic("calcpostexpression invalid calc : " + strings.Join(split, ""))
	}

	rs, err := strconv.ParseBool(st.Pop())

	if err != nil {
		panic(err)
	}

	return rs
}

func intopostfixexpression(split []string) []string {
	if split == nil {
		panic("intopostfixexpression argument is null")
	}

	sSize := len(split)

	if sSize == 0 {
		panic("intopostfixexpression argument size is null")
	}

	var pf []string
	st := Stackstring{}

	for _, value := range split {
		switch value {
		case "(":
			st.Push(value)
		case ")":
			if st.GetSize() != 0 {
				sz := st.GetSize()
				for i := 0; i < sz; i++ {
					if st.Top() == "(" {
						st.Pop()
						break
					}
					pf = append(pf, st.Pop())
				}
			}
		case "or":
			if st.GetSize() != 0 {
				sz := st.GetSize()
				for i := 0; i < sz; i++ {
					if st.Top() == "(" {
						break
					}
					pf = append(pf, st.Pop())
				}
			}

			st.Push(value)
		case "and":
			if st.GetSize() != 0 {
				sz := st.GetSize()
				for i := 0; i < sz; i++ {
					if st.Top() == "and" {
						if st.Top() == "(" {
							break
						}
						pf = append(pf, st.Pop())
					} else {
						break
					}

				}
			}
			st.Push(value)
		default:
			pf = append(pf, value)
		}
	}

	stSize := st.GetSize()

	if stSize > 0 {
		for i := 0; i < stSize; i++ {
			pf = append(pf, st.Pop())
		}
	}

	return pf
}

func calcpostfix(split []string) float64 {
	if split == nil {
		panic("intopostfix argument is null")
	}

	sSize := len(split)

	if sSize == 0 {
		panic("intopostfix argument size is null")
	}

	st := Stackfloat64{}
	var val1 float64
	var val2 float64

	for _, value := range split {

		switch value {
		case "*":
			val2 = st.Pop()
			val1 = st.Pop()
			st.Push(val1 * val2)

		case "/":
			val2 = st.Pop()
			val1 = st.Pop()
			st.Push(val1 / val2)

		case "+":
			val2 = st.Pop()
			val1 = st.Pop()
			st.Push(val1 + val2)

		case "-":
			val2 = st.Pop()
			val1 = st.Pop()
			st.Push(val1 - val2)

		case " ":
		default:
			if num, err := strconv.ParseFloat(value, 64); err != nil {
				panic("clacpostfix invalid argument : " + strings.Join(split, ""))
			} else {
				st.Push(num)
			}
		}

	}

	if st.GetSize() != 1 {
		panic("clacpostfix invalid calc : " + strings.Join(split, ""))
	}

	return st.Pop()
}

func sliceoper(s string) []string {

	var rs []string
	var num string

	s = strings.Trim(s, " ")
	sSize := len(s)

	for i := 0; i < sSize; i++ {
		sValue := string(s[i])
		switch sValue {
		case "(", ")", "*", "/", "+", "-":
			if num != "" {
				rs = append(rs, num)
				num = ""
			}

			if sValue == "(" {
				if string(s[i+1]) == "-" {

					index := strings.Index(s[i+1:sSize], ")")

					rs = append(rs, s[i+1:i+1+index])

					i = i + index + 1

					continue
				}
			}

			rs = append(rs, sValue)
		default:
			if sValue == "." {
				if num == "" {
					panic("sliceoper argument invalid : " + s)
				}
				if strings.Contains(num, sValue) {
					panic("sliceoper argument invalid : " + s)
				}

				num += sValue
			} else if _, err := strconv.ParseFloat(sValue, 64); err != nil {
				panic(err)
			} else {
				num += sValue
			}
		}
	}

	if num != "" {
		rs = append(rs, num)
	}

	return rs
}

func intopostfixcalc(split []string) []string {
	if split == nil {
		panic("intopostfixcalc argument is null")
	}

	sSize := len(split)

	if sSize == 0 {
		panic("intopostfixcalc argument size is null")
	}

	var pf []string
	st := Stackstring{}

	for _, value := range split {
		switch value {
		case "(":
			st.Push(value)
		case ")":
			if st.GetSize() != 0 {
				sz := st.GetSize()
				for i := 0; i < sz; i++ {
					if st.Top() == "(" {
						st.Pop()
						break
					}
					pf = append(pf, st.Pop())
				}
			}
		case "+", "-":
			if st.GetSize() != 0 {
				sz := st.GetSize()
				for i := 0; i < sz; i++ {
					if st.Top() == "(" {
						break
					}
					pf = append(pf, st.Pop())
				}
			}

			st.Push(value)
		case "*", "/":
			if st.GetSize() != 0 {
				sz := st.GetSize()
				for i := 0; i < sz; i++ {
					if st.Top() == "*" || st.Top() == "/" {
						if st.Top() == "(" {
							break
						}
						pf = append(pf, st.Pop())
					} else {
						break
					}

				}
			}

			st.Push(value)
		default:
			if _, err := strconv.ParseFloat(value, 64); err != nil {
				panic("intopostfixcalc argument invalid : " + value + ", " + err.Error())
			}
			pf = append(pf, value)
		}
	}

	stSize := st.GetSize()

	if stSize > 0 {
		for i := 0; i < stSize; i++ {
			pf = append(pf, st.Pop())
		}
	}

	return pf
}
