package ies

import "rrc/utils"

// DC-Parameters-v1310 ::= SEQUENCE
type DcParametersV1310 struct {
	PdcpTransfersplitulR13 *utils.ENUMERATED
	UeSstdMeasR13          *utils.ENUMERATED
}
