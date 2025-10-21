package ies

import "rrc/utils"

// CA-BandwidthClass-r10 ::= ENUMERATED
// Extensible
type CaBandwidthclassR10 struct {
	Value utils.ENUMERATED
}

const (
	CaBandwidthclassR10A = 0
	CaBandwidthclassR10B = 1
	CaBandwidthclassR10C = 2
	CaBandwidthclassR10D = 3
	CaBandwidthclassR10E = 4
	CaBandwidthclassR10F = 5
)
