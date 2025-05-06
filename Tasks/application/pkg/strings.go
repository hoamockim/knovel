package pkg

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"math/big"
	"regexp"
	"strings"
)

const (
	pwdPattern = "^(?=.*[a-z])(?=.*[A-Z])(?=.*[0-9])(?=.*[!@#\\$%\\^&\\*]).{8,}$"
	emlPattern = "^(([^<>()\\[\\]\\\\.,;:\\s@\"]+(\\.[^<>()\\[\\]\\\\.,;:\\s@\"]+)*)|(\".+\"))@((\\[[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}])|(([a-zA-Z\\-0-9]+\\.)+[a-zA-Z]{2,}))$"
)

var (
	mySecret = "e0345e58ede1c4eb060b4403a6685c4e" //os.Getenv("APP_SECRET")
	bytes    = []byte{35, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 05}
)

func Sha256(text string) string {
	encode := sha256.Sum256([]byte(text))
	return fmt.Sprintf("%x", encode)
}

func IsEmail(email string) bool {
	matched, _ := regexp.MatchString(emlPattern, email)
	return matched
}

func IsPassword(password string) bool {
	matched, _ := regexp.MatchString(pwdPattern, password)
	return matched
}

func StripAccents(name string) string {
	source := []string{"À", "Á", "Â", "Ã", "È", "É",
		"Ê", "Ì", "Í", "Ò", "Ó", "Ô", "Õ", "Ù", "Ú", "Ý", "à", "á", "â",
		"ã", "è", "é", "ê", "ì", "í", "ò", "ó", "ô", "õ", "ù", "ú", "ý",
		"Ă", "ă", "Đ", "đ", "Ĩ", "ĩ", "Ũ", "ũ", "Ơ", "ơ", "Ư", "ư", "Ạ",
		"ạ", "Ả", "ả", "Ấ", "ấ", "Ầ", "ầ", "Ẩ", "ẩ", "Ẫ", "ẫ", "Ậ", "ậ",
		"Ắ", "ắ", "Ằ", "ằ", "Ẳ", "ẳ", "Ẵ", "ẵ", "Ặ", "ặ", "Ẹ", "ẹ", "Ẻ",
		"ẻ", "Ẽ", "ẽ", "Ế", "ế", "Ề", "ề", "Ể", "ể", "Ễ", "ễ", "Ệ", "ệ",
		"Ỉ", "ỉ", "Ị", "ị", "Ọ", "ọ", "Ỏ", "ỏ", "Ố", "ố", "Ồ", "ồ", "Ổ",
		"ổ", "Ỗ", "ỗ", "Ộ", "ộ", "Ớ", "ớ", "Ờ", "ờ", "Ở", "ở", "Ỡ", "ỡ",
		"Ợ", "ợ", "Ụ", "ụ", "Ủ", "ủ", "Ứ", "ứ", "Ừ", "ừ", "Ử", "ử", "Ữ",
		"ữ", "Ự", "ự", "ý", "ỳ", "ỷ", "ỹ", "ỵ", "Ý", "Ỳ", "Ỷ", "Ỹ", "Ỵ"}

	dist := []string{"A", "A", "A", "A", "E",
		"E", "E", "I", "I", "O", "O", "O", "O", "U", "U", "Y", "a", "a",
		"a", "a", "e", "e", "e", "i", "i", "o", "o", "o", "o", "u", "u",
		"y", "A", "a", "D", "d", "I", "i", "U", "u", "O", "o", "U", "u",
		"A", "a", "A", "a", "A", "a", "A", "a", "A", "a", "A", "a", "A",
		"a", "A", "a", "A", "a", "A", "a", "A", "a", "A", "a", "E", "e",
		"E", "e", "E", "e", "E", "e", "E", "e", "E", "e", "E", "e", "E",
		"e", "I", "i", "I", "i", "O", "o", "O", "o", "O", "o", "O", "o",
		"O", "o", "O", "o", "O", "o", "O", "o", "O", "o", "O", "o", "O",
		"o", "O", "o", "U", "u", "U", "u", "U", "u", "U", "u", "U", "u",
		"U", "u", "U", "u", "y", "y", "y", "y", "y", "Y", "Y", "Y", "Y", "Y"}

	for index, char := range source {
		name = strings.Replace(name, char, dist[index], -1)
	}

	name = strings.Replace(name, "'", " ", 10)

	return strings.ToUpper(name)
}

func GenerateRandom(n int) string {
	const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-_"
	rs := make([]byte, n)
	for i := 0; i < n; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			//log error
			return ""
		}
		rs[i] = letters[num.Int64()]
	}

	return string(rs)
}

func EncryptBase64(text string) (string, error) {
	block, err := aes.NewCipher([]byte(mySecret))
	if err != nil {
		return "", err
	}
	plainText := []byte(text)
	cfb := cipher.NewCFBEncrypter(block, bytes)
	cipherText := make([]byte, len(plainText))
	cfb.XORKeyStream(cipherText, plainText)
	return EncodeBase64(cipherText), nil
}

// Decrypt method is to extract back the encrypted text
func DecryptBase64(text string) (string, error) {
	block, err := aes.NewCipher([]byte(mySecret))
	if err != nil {
		return "", err
	}
	cipherText := DecodeBase64(text)
	cfb := cipher.NewCFBDecrypter(block, bytes)
	plainText := make([]byte, len(cipherText))
	cfb.XORKeyStream(plainText, cipherText)
	return string(plainText), nil
}

func CompareArrays(source, destination []string) bool {
	set := make(map[string]struct{})
	for _, item := range source {
		set[item] = struct{}{}
	}

	for _, item := range destination {
		if _, exists := set[item]; !exists {
			return false
		}
	}
	return true
}
