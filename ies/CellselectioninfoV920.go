package ies

import "rrc/utils"

// CellSelectionInfo-v920 ::= SEQUENCE
type CellselectioninfoV920 struct {
	QQualminR9       QQualminR9
	QQualminoffsetR9 *utils.INTEGER `lb:0,ub:8`
}
