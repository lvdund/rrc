package ies

import "rrc/utils"

// MeasGapInfoNR-r16 ::= SEQUENCE
type MeasgapinfonrR16 struct {
	InterratBandlistnrEnDcR16 *InterratBandlistnrR16
	InterratBandlistnrSaR16   *InterratBandlistnrR16
}
