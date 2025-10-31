package ies

import "rrc/utils"

// CSI-RS-ConfigNZP-r11-qcl-CRS-Info-r11 ::= SEQUENCE
type CsiRsConfignzpR11QclCrsInfoR11 struct {
	QclScramblingidentityR11   utils.INTEGER `lb:0,ub:503`
	CrsPortscountR11           CsiRsConfignzpR11QclCrsInfoR11CrsPortscountR11
	MbsfnSubframeconfiglistR11 *CsiRsConfignzpR11QclCrsInfoR11MbsfnSubframeconfiglistR11
}
