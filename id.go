package fake

func RandGuid() string {
	return RandHexString(8) + `-` +
		RandHexString(4) + `-` +
		RandHexString(4) + `-` +
		RandHexString(4) + `-` +
		RandHexString(12)
}
