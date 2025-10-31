package ies

import "rrc/utils"

// ServingCellConfig-csi-RS-ValidationWithDCI-r16 ::= ENUMERATED
type ServingcellconfigCsiRsValidationwithdciR16 struct {
	Value utils.ENUMERATED
}

const (
	ServingcellconfigCsiRsValidationwithdciR16EnumeratedNothing = iota
	ServingcellconfigCsiRsValidationwithdciR16EnumeratedEnabled
)
