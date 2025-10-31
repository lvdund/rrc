package ies

import "rrc/utils"

// CSI-RS-ConfigNZP-r11 ::= SEQUENCE
// Extensible
type CsiRsConfignzpR11 struct {
	CsiRsConfignzpidR11   CsiRsConfignzpidR11
	AntennaportscountR11  CsiRsConfignzpR11AntennaportscountR11
	ResourceconfigR11     utils.INTEGER `lb:0,ub:31`
	SubframeconfigR11     utils.INTEGER `lb:0,ub:154`
	ScramblingidentityR11 utils.INTEGER `lb:0,ub:503`
	QclCrsInfoR11         *CsiRsConfignzpR11QclCrsInfoR11
}
