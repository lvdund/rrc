package ies

import "rrc/utils"

// CSI-IM-ConfigExt-r12 ::= SEQUENCE
// Extensible
type CsiImConfigextR12 struct {
	CsiImConfigidV1250 CsiImConfigidV1250
	ResourceconfigR12  utils.INTEGER
	SubframeconfigR12  utils.INTEGER
}
