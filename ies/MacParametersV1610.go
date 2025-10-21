package ies

import "rrc/utils"

// MAC-Parameters-v1610 ::= SEQUENCE
type MacParametersV1610 struct {
	DirectmcgScellactivationresumeR16 *utils.ENUMERATED
	DirectscgScellactivationresumeR16 *utils.ENUMERATED
	EarlydataUp5gcR16                 *utils.ENUMERATED
	RaiSupportenhR16                  *utils.ENUMERATED
}
