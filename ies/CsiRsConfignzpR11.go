package ies

import "rrc/utils"

// CSI-RS-ConfigNZP-r11 ::= SEQUENCE
// Extensible
type CsiRsConfignzpR11 struct {
	CsiRsConfignzpidR11   CsiRsConfignzpidR11
	AntennaportscountR11  utils.ENUMERATED
	ResourceconfigR11     utils.INTEGER
	SubframeconfigR11     utils.INTEGER
	ScramblingidentityR11 utils.INTEGER
	QclCrsInfoR11         *interface{}
}
