package commands

func min(a uint64, b uint64) uint64 {
	if a > b {
		return b
	}
	return a
}

func max(a uint64, b uint64) uint64 {
	if a > b {
		return a
	}
	return b
}
