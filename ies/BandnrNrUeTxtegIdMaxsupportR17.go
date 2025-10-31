package ies

import "rrc/utils"

// BandNR-nr-UE-TxTEG-ID-MaxSupport-r17 ::= ENUMERATED
type BandnrNrUeTxtegIdMaxsupportR17 struct {
	Value utils.ENUMERATED
}

const (
	BandnrNrUeTxtegIdMaxsupportR17EnumeratedNothing = iota
	BandnrNrUeTxtegIdMaxsupportR17EnumeratedN1
	BandnrNrUeTxtegIdMaxsupportR17EnumeratedN2
	BandnrNrUeTxtegIdMaxsupportR17EnumeratedN3
	BandnrNrUeTxtegIdMaxsupportR17EnumeratedN4
	BandnrNrUeTxtegIdMaxsupportR17EnumeratedN6
	BandnrNrUeTxtegIdMaxsupportR17EnumeratedN8
)
