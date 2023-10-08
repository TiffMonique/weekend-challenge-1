package physics

func Ohm(v, r, i float32) float32 {
	if v == 0 {
		if r != 0 && i != 0 {
			return i * r
		}
	} else if r == 0 {
		if v != 0 && i != 0 {
			return v / i
		}
	} else if i == 0 {
		if v != 0 && r != 0 {
			return v / r
		}
	}
	return 0
}
