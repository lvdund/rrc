package ies

import "rrc/utils"

// CA-BandwidthClass-r10 ::= utils.ENUMERATED // Extensible
type CaBandwidthclassR10 struct {
	Value utils.ENUMERATED
}

const (
	CaBandwidthclassR10EnumeratedNothing = iota
	CaBandwidthclassR10EnumeratedA
	CaBandwidthclassR10EnumeratedB
	CaBandwidthclassR10EnumeratedC
	CaBandwidthclassR10EnumeratedD
	CaBandwidthclassR10EnumeratedE
	CaBandwidthclassR10EnumeratedF
)
