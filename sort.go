package banner

// SortByExpiry sort banners using expire time
type SortByExpiry []*Banner

// Len returns number of items in the slice
func (s SortByExpiry) Len() int {
	return len(s)
}

// Swap items from the slice
func (s SortByExpiry) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]

}

// Less compares two items and returns bool
func (s SortByExpiry) Less(i, j int) bool {
	return s[i].ExpireAt < s[j].ExpireAt
}
