package ies

import "rrc/utils"

// LWA-Parameters-v1430 ::= SEQUENCE
type LwaParametersV1430 struct {
	LwaHoWithoutwtChangeR14  *LwaParametersV1430LwaHoWithoutwtChangeR14
	LwaUlR14                 *LwaParametersV1430LwaUlR14
	WlanPeriodicmeasR14      *LwaParametersV1430WlanPeriodicmeasR14
	WlanReportanywlanR14     *LwaParametersV1430WlanReportanywlanR14
	WlanSupporteddatarateR14 *utils.INTEGER `lb:0,ub:2048`
}
