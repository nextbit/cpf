package cpf

import (
	"bytes"
	"math"
	"regexp"
	"strconv"
	"unicode/utf8"
)

func Valid(cpf string) bool {
	return fullCPF(cpf) == clean(cpf)
}

func fullCPF(cpf string) string {
	fd := calcDigit(cpf, 1)
	sd := calcDigit(cpf, 2)

	var b bytes.Buffer
	b.WriteString(clean(cpf)[:9])
	b.WriteString(strconv.Itoa(fd))
	b.WriteString(strconv.Itoa(sd))

	return b.String()
}

func calcDigit(cpf string, digit int) int {
	var size int
	if digit == 1 {
		size = 9
	}
	if digit == 2 {
		size = 10
	}

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
