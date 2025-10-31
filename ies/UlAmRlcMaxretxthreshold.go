package ies

import "rrc/utils"

// UL-AM-RLC-maxRetxThreshold ::= ENUMERATED
type UlAmRlcMaxretxthreshold struct {
	Value utils.ENUMERATED
}

const (
	UlAmRlcMaxretxthresholdEnumeratedNothing = iota
	UlAmRlcMaxretxthresholdEnumeratedT1
	UlAmRlcMaxretxthresholdEnumeratedT2
	UlAmRlcMaxretxthresholdEnumeratedT3
	UlAmRlcMaxretxthresholdEnumeratedT4
	UlAmRlcMaxretxthresholdEnumeratedT6
	UlAmRlcMaxretxthresholdEnumeratedT8
	UlAmRlcMaxretxthresholdEnumeratedT16
	UlAmRlcMaxretxthresholdEnumeratedT32
)
