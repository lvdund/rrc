package ies

import "rrc/utils"

// BandNR-sn-InitiatedCondPSCellChangeNRDC-r17 ::= ENUMERATED
type BandnrSnInitiatedcondpscellchangenrdcR17 struct {
	Value utils.ENUMERATED
}

const (
	BandnrSnInitiatedcondpscellchangenrdcR17EnumeratedNothing = iota
	BandnrSnInitiatedcondpscellchangenrdcR17EnumeratedSupported
)
