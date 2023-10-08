package loops

import "strconv"

func parseNum(num int) string {
	if num%3 == 0 && num%5 == 0 {
		return "FooBar"
	} else if num%5 == 0 {
		return "Bar"
	} else if num%3 == 0 {
		return "Foo"
	} else {
		return strconv.Itoa(num)
	}
}

func Foobar(num int) string {
	var chainstr string
	for i := 1; i <= num; i++ {
		if i == num {
			chainstr += parseNum(i)
		} else {
			chainstr += parseNum(i) + "->"
		}
	}
	return chainstr
}
