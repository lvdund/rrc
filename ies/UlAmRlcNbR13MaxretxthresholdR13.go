package ies

import "rrc/utils"

// UL-AM-RLC-NB-r13-maxRetxThreshold-r13 ::= ENUMERATED
type UlAmRlcNbR13MaxretxthresholdR13 struct {
	Value utils.ENUMERATED
}

const (
	UlAmRlcNbR13MaxretxthresholdR13EnumeratedNothing = iota
	UlAmRlcNbR13MaxretxthresholdR13EnumeratedT1
	UlAmRlcNbR13MaxretxthresholdR13EnumeratedT2
	UlAmRlcNbR13MaxretxthresholdR13EnumeratedT3
	UlAmRlcNbR13MaxretxthresholdR13EnumeratedT4
	UlAmRlcNbR13MaxretxthresholdR13EnumeratedT6
	UlAmRlcNbR13MaxretxthresholdR13EnumeratedT8
	UlAmRlcNbR13MaxretxthresholdR13EnumeratedT16
	UlAmRlcNbR13MaxretxthresholdR13EnumeratedT32
)
