package cpf

import (
	"bytes"
	"math"
	"regexp"
	"strconv"
	"strings"
)

// Valid checks if the passed cpf is valid
func Valid(cpf string) bool {
	if cpf = clean(cpf); len(cpf) != 11 {
		return false
	}

	reg := regexp.MustCompile(`(1|2|3|4|5|6|7|8|9){11}`)
	if reg.MatchString(clean(cpf)) {
		return false
	}

	verifiedCPF, err := validateCPF(cpf)
	if err != nil {
		return false
	}

	return verifiedCPF == clean(cpf)
}

func validateCPF(cpf string) (string, error) {
	fd, err := calcFirstDigit(cpf[:9])
	if err != nil {
		return "", err
	}
	sd, err := calcSecondDigit(cpf[:9] + strconv.Itoa(fd))
	if err != nil {
		return "", err
	}

	var b bytes.Buffer
	b.WriteString(clean(cpf)[:9])
	b.WriteString(strconv.Itoa(fd))
	b.WriteString(strconv.Itoa(sd))

	return b.String(), nil
}

func calcFirstDigit(cpf string) (int, error) {
	verifiers := []int{10, 9, 8, 7, 6, 5, 4, 3, 2}
	digits := strings.Split(cpf, "")
	return calcDigit(digits, verifiers, 9)
}

func calcSecondDigit(cpf string) (int, error) {
	verifiers := []int{11, 10, 9, 8, 7, 6, 5, 4, 3, 2}
	digits := strings.Split(cpf, "")
	return calcDigit(digits, verifiers, 10)
}

func calcDigit(digits []string, verifiers []int, size int) (int, error) {
	var sum int

	for i := 0; i < size; i++ {
		dig, err := strconv.Atoi(digits[i])
		if err != nil {
			return sum, err
		}
		sum = sum + verifiers[i]*dig
	}

	if r := math.Mod(float64(sum), 11); r >= 2 {
		return int(11 - r), nil
	}

	return 0, nil
}

func clean(s string) string {
	reg := regexp.MustCompile(`[^\d]+`)
	c := reg.ReplaceAllString(s, "")
	return c
}
