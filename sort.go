package banner

type ByExpiry []Banner

func (b ByExpiry) Len() int {
	return len(b)
}

func (b ByExpiry) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]

}

func (b ByExpiry) Less(i, j int) bool {
	return b[i].ExpireAt < b[j].ExpireAt
}
