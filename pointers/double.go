package pointers

func Double(input *int) {
	if input != nil {
		*input *= 2
	}
}
