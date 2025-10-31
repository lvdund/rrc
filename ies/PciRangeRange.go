package ies

import "rrc/utils"

// PCI-Range-range ::= ENUMERATED
type PciRangeRange struct {
	Value utils.ENUMERATED
}

const (
	PciRangeRangeEnumeratedNothing = iota
	PciRangeRangeEnumeratedN4
	PciRangeRangeEnumeratedN8
	PciRangeRangeEnumeratedN12
	PciRangeRangeEnumeratedN16
	PciRangeRangeEnumeratedN24
	PciRangeRangeEnumeratedN32
	PciRangeRangeEnumeratedN48
	PciRangeRangeEnumeratedN64
	PciRangeRangeEnumeratedN84
	PciRangeRangeEnumeratedN96
	PciRangeRangeEnumeratedN128
	PciRangeRangeEnumeratedN168
	PciRangeRangeEnumeratedN252
	PciRangeRangeEnumeratedN504
	PciRangeRangeEnumeratedN1008
	PciRangeRangeEnumeratedSpare1
)
