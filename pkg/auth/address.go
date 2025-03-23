package auth

func IsValidAddress(address string) bool {
	return !(address == "" || address[:2] != "0x" || len(address) != 42)
}
