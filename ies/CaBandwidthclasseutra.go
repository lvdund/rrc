package ies

import "rrc/utils"

// CA-BandwidthClassEUTRA ::= utils.ENUMERATED // Extensible
type CaBandwidthclasseutra struct {
	Value utils.ENUMERATED
}

const (
	CaBandwidthclasseutraEnumeratedNothing = iota
	CaBandwidthclasseutraEnumeratedA
	CaBandwidthclasseutraEnumeratedB
	CaBandwidthclasseutraEnumeratedC
	CaBandwidthclasseutraEnumeratedD
	CaBandwidthclasseutraEnumeratedE
	CaBandwidthclasseutraEnumeratedF
)
