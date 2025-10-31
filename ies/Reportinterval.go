package ies

import "rrc/utils"

// ReportInterval ::= ENUMERATED
type Reportinterval struct {
	Value utils.ENUMERATED
}

const (
	ReportintervalEnumeratedNothing = iota
	ReportintervalEnumeratedMs120
	ReportintervalEnumeratedMs240
	ReportintervalEnumeratedMs480
	ReportintervalEnumeratedMs640
	ReportintervalEnumeratedMs1024
	ReportintervalEnumeratedMs2048
	ReportintervalEnumeratedMs5120
	ReportintervalEnumeratedMs10240
	ReportintervalEnumeratedMin1
	ReportintervalEnumeratedMin6
	ReportintervalEnumeratedMin12
	ReportintervalEnumeratedMin30
	ReportintervalEnumeratedMin60
	ReportintervalEnumeratedSpare3
	ReportintervalEnumeratedSpare2
	ReportintervalEnumeratedSpare1
)
