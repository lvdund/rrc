package ies

import "rrc/utils"

// PerRACSI-RSInfo-v1660 ::= SEQUENCE
type PerracsiRsinfoV1660 struct {
	CsiRsIndexV1660 *utils.INTEGER `lb:0,ub:96`
}
