package ies

import "rrc/utils"

// UE-NR-Capability-v1700-sliceInfoforCellReselection-r17 ::= ENUMERATED
type UeNrCapabilityV1700SliceinfoforcellreselectionR17 struct {
	Value utils.ENUMERATED
}

const (
	UeNrCapabilityV1700SliceinfoforcellreselectionR17EnumeratedNothing = iota
	UeNrCapabilityV1700SliceinfoforcellreselectionR17EnumeratedSupported
)
