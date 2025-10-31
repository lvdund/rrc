package ies

import "rrc/utils"

// UEPositioningAssistanceInfo-r17-IEs ::= SEQUENCE
type UepositioningassistanceinfoR17 struct {
	UeTxtegAssociationlistR17 *UeTxtegAssociationlistR17
	Latenoncriticalextension  *utils.OCTETSTRING
	Noncriticalextension      *UepositioningassistanceinfoV1720
}
