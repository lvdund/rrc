package ies

import "rrc/utils"

// PhyLayerParameters-v1450-crs-LessDwPTS-r14 ::= ENUMERATED
type PhylayerparametersV1450CrsLessdwptsR14 struct {
	Value utils.ENUMERATED
}

const (
	PhylayerparametersV1450CrsLessdwptsR14EnumeratedNothing = iota
	PhylayerparametersV1450CrsLessdwptsR14EnumeratedSupported
)
