package datatype

func copyData(b []byte) []byte {
	bCopy := make([]byte, len(b))
	copy(bCopy, b)
	return bCopy
}
