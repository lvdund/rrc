package ies

import "rrc/utils"

// ReportInterval ::= ENUMERATED
type Reportinterval struct {
	Value utils.ENUMERATED
}

const (
	ReportintervalMs120   = 0
	ReportintervalMs240   = 1
	ReportintervalMs480   = 2
	ReportintervalMs640   = 3
	ReportintervalMs1024  = 4
	ReportintervalMs2048  = 5
	ReportintervalMs5120  = 6
	ReportintervalMs10240 = 7
	ReportintervalMin1    = 8
	ReportintervalMin6    = 9
	ReportintervalMin12   = 10
	ReportintervalMin30   = 11
	ReportintervalMin60   = 12
	ReportintervalSpare3  = 13
	ReportintervalSpare2  = 14
	ReportintervalSpare1  = 15
)
