package ies

import "rrc/utils"

// SL-DiscSysInfoReport-r13-discCellSelectionInfo-r13 ::= SEQUENCE
type SlDiscsysinforeportR13DisccellselectioninfoR13 struct {
	QRxlevminR13       QRxlevmin
	QRxlevminoffsetR13 *utils.INTEGER `lb:0,ub:8`
}
