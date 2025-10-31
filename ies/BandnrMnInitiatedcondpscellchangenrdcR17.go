package ies

import "rrc/utils"

// BandNR-mn-InitiatedCondPSCellChangeNRDC-r17 ::= ENUMERATED
type BandnrMnInitiatedcondpscellchangenrdcR17 struct {
	Value utils.ENUMERATED
}

const (
	BandnrMnInitiatedcondpscellchangenrdcR17EnumeratedNothing = iota
	BandnrMnInitiatedcondpscellchangenrdcR17EnumeratedSupported
)
