package hasher

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"
	"unicode"

	"github.com/michalsabaj/pldlchecker/config"
)

func HashDane(firstName, lastName, driverLicenseNumber string) string {
	toHash := fmt.Sprintf("%s%s%s", firstName, lastName, driverLicenseNumber)
	//normalization
	//fmt.Printf("Before normalization: '%s'\n", toHash)
	toHash = normalizeInput(toHash)
	cfg := config.GetConfig()
	if !cfg.Debug {
		fmt.Printf("[DEBUG] To hash: (before normalized) '%s'\n", toHash)
	}
	//hashing
	normalized := NormalizeForHash(toHash)
	if !cfg.Debug {
		fmt.Printf("[DEBUG] To hash: (normalized) '%s'\n", normalized)
	}
	hash := md5.Sum([]byte(normalized))
	formatedHash := strings.ToUpper(hex.EncodeToString(hash[:]))
	if !cfg.Debug {
		fmt.Printf("[DEBUG] Hash: '%s'\n", formatedHash)

	}
	return formatedHash
}

func NormalizeForHash(t string) string {
	if t == "" {
		t = string(t)
	}

	// Truncate if length exceeds 500 characters
	if len(t) > 500 {
		t = t[:500]
	}

	// Convert to uppercase
	t = strings.ToUpper(t)
	//fmt.Printf("Uppercase: '%s'\n", t)
	// Remove characters not in A-Z, 0-9, or specific Polish characters
	var b strings.Builder
	for _, r := range t {
		if unicode.IsLetter(r) || unicode.IsDigit(r) || containsPolishChar(r) {
			b.WriteRune(r)
		}
	}
	t = b.String()

	// Replace specific Polish characters with their Latin equivalents
	t = strings.ReplaceAll(t, "Ż", "Z")
	t = strings.ReplaceAll(t, "Ó", "O")
	t = strings.ReplaceAll(t, "Ł", "L")
	t = strings.ReplaceAll(t, "Ć", "C")
	t = strings.ReplaceAll(t, "Ę", "E")
	t = strings.ReplaceAll(t, "Ś", "S")
	t = strings.ReplaceAll(t, "Ą", "A")
	t = strings.ReplaceAll(t, "Ź", "Z")
	t = strings.ReplaceAll(t, "Ń", "N")

	return t
}

// containsPolishChar checks if the rune is one of the special Polish characters to keep before normalization.
func containsPolishChar(r rune) bool {
	switch r {
	case 'Ż', 'Ó', 'Ł', 'Ć', 'Ę', 'Ś', 'Ą', 'Ź', 'Ń':
		return true
	default:
		return false
	}
}

// GetNormalizedHexMD5 computes the MD5 hash of the normalized input string.
func GetNormalizedHexMD5(t string) string {
	normalized := NormalizeForHash(t)
	hash := md5.Sum([]byte(normalized))
	return fmt.Sprintf("%X", hash)
}
func normalizeInput(input string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(input)), " ")
}
