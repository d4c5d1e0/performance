package performance

import (
	"strconv"
	"time"
)

type Browser struct {
	hasDecimal bool
}

type Performance struct {
	start      int64
	offset     int64
	hasDecimal bool
	Mode       LengthMode
}

var (
	// Chromium is the Browser value corresponding to Chromium browsers, and try to mimic
	// their Performance implementation.
	Chromium = NewBrowser(true)
	// Gecko is the Browser value corresponding to the firefox browser, and try to mimic
	// its Performance implementation. Firefox's performance doesn't return floats with
	// decimal precision.
	Gecko = NewBrowser(false)
)

// NewBrowser creates a new instance of the Browser struct with the specified decimal precision flag.
func NewBrowser(hasDecimal bool) *Browser {
	return &Browser{hasDecimal: hasDecimal}
}

// TimeOrigin returns a string representation of the current time in
// milliseconds, with or without decimal precision depending on the associated Browser instance.
func (b *Browser) TimeOrigin() float64 {
	t := time.Now().UnixMilli()
	if !b.hasDecimal {
		return float64(t)
	}

	return float64(t) + modes[SingleDigitMode].float()
}

// TimeOriginString returns a string representation of the current time in
// milliseconds, with or without decimal precision depending on the associated Browser instance.
func (b *Browser) TimeOriginString() string {
	t := time.Now().UnixMilli()
	if !b.hasDecimal {
		return strconv.FormatInt(t, 10)
	}

	f := float64(t) + modes[SingleDigitMode].float()
	return strconv.FormatFloat(f, 'f', SingleDigitMode.Index(), 64)
}

// NewPerformance creates a new instance of the Performance struct with the
// specified length mode and the starting time set to the current time.
func (b *Browser) NewPerformance(mode LengthMode) *Performance {
	return &Performance{
		start:      time.Now().UnixMilli(),
		hasDecimal: b.hasDecimal,
		Mode:       mode,
	}
}

// Now returns the elapsed time since the start of the Performance instance,
// with or without decimal precision depending on the Browser.
func (p *Performance) Now() float64 {
	t := time.Now().UnixMilli()
	if !p.hasDecimal {
		return float64(t - p.start)
	}

	return float64(t-p.start) + modes[p.Mode].float()
}

// Add adds the specified offset value to the underlying offset.
func (p *Performance) Add(offset int64) {
	p.offset += offset
}

// Min subtracts the specified offset value from the underlying offset.
func (p *Performance) Min(offset int64) {
	p.offset -= offset
}

// Current returns a float of the current offset with the random decimal value added to it, to match the Browser implementation.
func (p *Performance) Current() float64 {
	if !p.hasDecimal {
		return float64(p.offset)
	}

	return float64(p.offset) + modes[p.Mode].float()
}

// CurrentString returns a string representation of the current offset value,
// including the random decimal value generated to match the Browser implementation.
func (p *Performance) CurrentString() string {
	if !p.hasDecimal {
		return strconv.FormatInt(p.offset, 10)
	}

	f := float64(p.offset) + modes[p.Mode].float()
	return strconv.FormatFloat(f, 'f', p.Mode.Index(), 64)
}
