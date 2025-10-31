package ies

// ConnEstFailReport-r11-measResultFailedCell-r11 ::= SEQUENCE
type ConnestfailreportR11MeasresultfailedcellR11 struct {
	RsrpresultR11 RsrpRange
	RsrqresultR11 *RsrqRange
}
