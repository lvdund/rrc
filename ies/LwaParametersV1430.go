package ies

import "rrc/utils"

// LWA-Parameters-v1430 ::= SEQUENCE
type LwaParametersV1430 struct {
	LwaHoWithoutwtChangeR14  *utils.ENUMERATED
	LwaUlR14                 *utils.ENUMERATED
	WlanPeriodicmeasR14      *utils.ENUMERATED
	WlanReportanywlanR14     *utils.ENUMERATED
	WlanSupporteddatarateR14 *utils.INTEGER
}
