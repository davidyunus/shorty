package generator

import "crypto/rand"

// RandomStringSet ...
func RandomStringSet(n int, set string) (string, error) {
	bytes, err := RandomBytes(n)
	if err != nil {
		return "", err
	}
	for i, b := range bytes {
		bytes[i] = set[b%byte(len(set))]
	}
	return string(bytes), nil
}

// RandomBytes ...
func RandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)

	if err != nil {
		return nil, err
	}

	return b, nil
}

func RandomNumericString(n int) (string, error) {
	const letters = "0123456789"
	return RandomStringSet(n, letters)
}
