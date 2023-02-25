package performance

import (
	"strconv"
	"strings"

	"pgregory.net/rand"
)

type LengthMode int

const (
	SingleDigitMode LengthMode = iota
	MediumLengthMode
	LargeLengthMode
)

var (
	mediumLength = [2]int{4, 5}
	largeLength  = [2]int{9, 10}
	dotByte      = byte('.')

	num      = [...]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	numBytes = [...]byte{0x30, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39}

	modes = [...]*mode{
		{func() string {
			return "." + num[rand.Intn(len(num))]
		}},
		{func() string {
			zeroes := mediumLength[rand.Intn(2)]
			trailing := numBytes[(rand.Intn(2)+9)%10]

			buf := new(strings.Builder)
			cache := make([]byte, zeroes)
			buf.Grow(zeroes * 2)

			for i := 0; i < zeroes; i++ {
				cache[i] = trailing
			}

			buf.Write([]byte{dotByte, numBytes[rand.Intn(len(numBytes))]})
			buf.Write(cache)
			buf.WriteString(randNum(zeroes))
			return buf.String()
		}},
		{func() string {
			zeroes := largeLength[rand.Intn(2)]
			endNums := mediumLength[rand.Intn(2)]
			trailing := numBytes[(rand.Intn(2)+9)%10]

			buf := new(strings.Builder)
			cache := make([]byte, zeroes)
			buf.Grow(endNums + zeroes)

			for i := 0; i < zeroes; i++ {
				cache[i] = trailing
			}

			buf.Write([]byte{dotByte, numBytes[rand.Intn(len(numBytes))]})
			buf.Write(cache)
			buf.WriteString(randNum(endNums))
			return buf.String()
		}},
	}
)

func (l LengthMode) Index() int {
	return [...]int{1, 11, 16}[l]
}

type mode struct {
	f func() string
}

func (m *mode) float() float64 {
	f, _ := strconv.ParseFloat(m.f(), 64)
	return f
}
func (m *mode) string() string {
	return m.f()
}

func randNum(length int) string {
	buf := make([]byte, length)
	for i := 0; i < length; i++ {
		buf[i] = numBytes[rand.Intn(len(numBytes))]
	}
	return string(buf)
}
