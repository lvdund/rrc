package ies

import "rrc/utils"

// SL-PSFCH-Config-r16-sl-MinTimeGapPSFCH-r16 ::= ENUMERATED
type SlPsfchConfigR16SlMintimegappsfchR16 struct {
	Value utils.ENUMERATED
}

const (
	SlPsfchConfigR16SlMintimegappsfchR16EnumeratedNothing = iota
	SlPsfchConfigR16SlMintimegappsfchR16EnumeratedSl2
	SlPsfchConfigR16SlMintimegappsfchR16EnumeratedSl3
)
