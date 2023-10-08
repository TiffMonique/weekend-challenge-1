package bmi

func Bmi(weight, height float32) float32 {
	return weight / (height * height)
}
