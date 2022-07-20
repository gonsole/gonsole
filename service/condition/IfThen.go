package condition

func IfThen(condition bool, then func()) {
	if condition {
		then()
	}
}
