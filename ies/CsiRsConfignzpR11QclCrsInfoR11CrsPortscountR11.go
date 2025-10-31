package ies

import "rrc/utils"

// CSI-RS-ConfigNZP-r11-qcl-CRS-Info-r11-crs-PortsCount-r11 ::= ENUMERATED
type CsiRsConfignzpR11QclCrsInfoR11CrsPortscountR11 struct {
	Value utils.ENUMERATED
}

const (
	CsiRsConfignzpR11QclCrsInfoR11CrsPortscountR11EnumeratedNothing = iota
	CsiRsConfignzpR11QclCrsInfoR11CrsPortscountR11EnumeratedN1
	CsiRsConfignzpR11QclCrsInfoR11CrsPortscountR11EnumeratedN2
	CsiRsConfignzpR11QclCrsInfoR11CrsPortscountR11EnumeratedN4
	CsiRsConfignzpR11QclCrsInfoR11CrsPortscountR11EnumeratedSpare1
)
