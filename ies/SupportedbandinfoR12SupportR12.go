package ies

import "rrc/utils"

// SupportedBandInfo-r12-support-r12 ::= ENUMERATED
type SupportedbandinfoR12SupportR12 struct {
	Value utils.ENUMERATED
}

const (
	SupportedbandinfoR12SupportR12EnumeratedNothing = iota
	SupportedbandinfoR12SupportR12EnumeratedSupported
)
