package cpf

import (
	"bytes"
	"math"
	"regexp"
	"strconv"
	"unicode/utf8"
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

	return fullCPF(cpf) == clean(cpf)
}

func fullCPF(cpf string) string {
	fd := calcFirstDigit(cpf)
	sd := calcSecondDigit(cpf)

	var b bytes.Buffer
	b.WriteString(clean(cpf)[:9])
	b.WriteString(strconv.Itoa(fd))
	b.WriteString(strconv.Itoa(sd))

	return b.String()
}

func calcDigit(cpf string, size int) int {
	var sum float64
	for i, v := range reverse(clean(cpf)[:size]) {
		buf := make([]byte, 1)
		_ = utf8.EncodeRune(buf, v)
		value, _ := strconv.ParseFloat(string(buf), 64)
		sum += value * (float64(i) + 2.0)
	}

	r := math.Mod(sum, 11)
	var d float64
	if r < 2 {
		d = 0
	} else {
		d = 11.0 - r
	}
	return int(d)
}

func calcFirstDigit(cpf string) int {
	return calcDigit(cpf, 9)
}

func calcSecondDigit(cpf string) int {
	return calcDigit(cpf, 10)
}

func clean(s string) string {
	reg := regexp.MustCompile(`[^\d]+`)
	c := reg.ReplaceAllString(s, "")
	return c
}

func reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}
