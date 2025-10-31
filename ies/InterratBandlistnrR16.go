package ies

// InterRAT-BandListNR-r16 ::= SEQUENCE OF InterRAT-BandInfoNR-r16
// SIZE (1..maxBandsNR-r15)
type InterratBandlistnrR16 struct {
	Value []InterratBandinfonrR16 `lb:1,ub:maxBandsNRR15`
}
