package util

func CheckNonNullString(target string) bool {
	return target != ""
}

func CheckNonNullInt64(target int64) bool {
	return target != 0
}