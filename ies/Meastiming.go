package ies

// MeasTiming ::= SEQUENCE
// Extensible
type Meastiming struct {
	Frequencyandtiming *MeastimingFrequencyandtiming
	SsbTomeasure       *SsbTomeasure `ext`
	Physcellid         *Physcellid   `ext`
}
