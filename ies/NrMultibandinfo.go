package ies

// NR-MultiBandInfo ::= SEQUENCE
type NrMultibandinfo struct {
	Freqbandindicatornr *Freqbandindicatornr
	NrNsPmaxlist        *NrNsPmaxlist
}
