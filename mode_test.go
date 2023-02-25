package performance

import (
	"testing"
)

func BenchmarkModeSingleDigitString(b *testing.B) {
	mode := modes[SingleDigitMode]
	for n := 0; n < b.N; n++ {
		mode.string()
	}
}

func BenchmarkModeMediumLengthModeString(b *testing.B) {
	mode := modes[MediumLengthMode]
	for n := 0; n < b.N; n++ {
		mode.string()
	}
}

func BenchmarkModeLargeLengthModeString(b *testing.B) {
	mode := modes[LargeLengthMode]
	for n := 0; n < b.N; n++ {
		mode.string()
	}
}

func BenchmarkRandomNum12(b *testing.B) {
	for n := 0; n < b.N; n++ {
		randNum(12)
	}
}
