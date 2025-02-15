package module01

// NumInList will return true if num is in the
// list slice, and false otherwise.
func NumInList(list []int, num int) bool {
	if list != nil && len(list) != 0 {
		for _, val := range list {
			if val == num {
				return true
			}
		}
	}
	return false
}