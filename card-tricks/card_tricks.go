package cards

// FavoriteCards returns a slice with the cards 2, 6 and 9 in that order.
func FavoriteCards() []int {
	favSlice := []int{2, 6, 9}
	return favSlice
}

// GetItem retrieves an item from a slice at given position.
// If the index is out of range, we want it to return -1.
func GetItem(slice []int, index int) int {
	if index >= len(slice) || index < 0 {
		return -1
	}
	x := slice[index]
	return x
}

// SetItem writes an item to a slice at given position overwriting an existing value.
// If the index is out of range the value needs to be appended.
func SetItem(slice []int, index, value int) []int {
	x := slice[:]
	if index >= len(slice) || index < 0 {
		x = append(x, value)
		return x
	} else {
		x[index] = value
		return x
	}
}

// PrependItems adds an arbitrary number of values at the front of a slice.
func PrependItems(slice []int, values ...int) []int {
	ss := slice[:]
	if len(values) == 0 {
		return ss
	} else {
		ss = append(values, slice...)
		return ss
	}
}

// RemoveItem removes an item from a slice by modifying the existing slice.
func RemoveItem(slice []int, index int) []int {
	if index >= 0 && index <= len(slice) {
		slice = append(slice[:index], slice[index+1:]...)
		return slice
	} else {
		return slice
	}
}
