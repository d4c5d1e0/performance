package performance

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBrowser_TimeOrigin(t *testing.T) {
	firefox := Gecko.TimeOriginString()
	chrome := Chromium.TimeOriginString()

	assert.Equal(t, 0, strings.Count(firefox, "."), "bad firefox result")
	assert.Equal(t, 1, strings.Count(chrome, "."), "bad chromium result")
}

func TestPerformance_Add(t *testing.T) {
	chrome := Chromium.NewPerformance(SingleDigitMode)
	firefox := Gecko.NewPerformance(SingleDigitMode)

	firefox.Add(22)
	chrome.Add(22)

	assert.True(t, chrome.Current() >= 22.0, "performance: add: wrong result")
	assert.True(t, firefox.Current() == float64(22), "performance: add: wrong result")
}

func TestPerformance_Min(t *testing.T) {
	chrome := Chromium.NewPerformance(SingleDigitMode)
	firefox := Gecko.NewPerformance(SingleDigitMode)

	firefox.Add(22)
	chrome.Add(22)

	firefox.Min(12)
	chrome.Min(12)

	assert.True(t, chrome.Current() >= 10.0, "performance: min: wrong result")
	assert.True(t, firefox.Current() == float64(10), "performance: min: wrong result")
}
