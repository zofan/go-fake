package fake

func Guid() string {
	return HexString(8) + `-` +
		HexString(4) + `-` +
		HexString(4) + `-` +
		HexString(4) + `-` +
		HexString(12)
}
