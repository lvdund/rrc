package ies

import "rrc/utils"

// RNReconfiguration-r10-IEs ::= SEQUENCE
type RnreconfigurationR10 struct {
	RnSysteminfoR10          *RnSysteminfoR10
	RnSubframeconfigR10      *RnSubframeconfigR10
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *RnreconfigurationR10IesNoncriticalextension
}
