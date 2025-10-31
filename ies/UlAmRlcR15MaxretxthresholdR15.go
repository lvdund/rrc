package ies

import "rrc/utils"

// UL-AM-RLC-r15-maxRetxThreshold-r15 ::= ENUMERATED
type UlAmRlcR15MaxretxthresholdR15 struct {
	Value utils.ENUMERATED
}

const (
	UlAmRlcR15MaxretxthresholdR15EnumeratedNothing = iota
	UlAmRlcR15MaxretxthresholdR15EnumeratedT1
	UlAmRlcR15MaxretxthresholdR15EnumeratedT2
	UlAmRlcR15MaxretxthresholdR15EnumeratedT3
	UlAmRlcR15MaxretxthresholdR15EnumeratedT4
	UlAmRlcR15MaxretxthresholdR15EnumeratedT6
	UlAmRlcR15MaxretxthresholdR15EnumeratedT8
	UlAmRlcR15MaxretxthresholdR15EnumeratedT16
	UlAmRlcR15MaxretxthresholdR15EnumeratedT32
)
