package ies

import "rrc/utils"

// CA-ParametersNR-v1610-pdcch-MonitoringCA-r16 ::= SEQUENCE
type CaParametersnrV1610PdcchMonitoringcaR16 struct {
	MaxnumberofmonitoringccR16  utils.INTEGER `lb:0,ub:16`
	SupportedspanarrangementR16 CaParametersnrV1610PdcchMonitoringcaR16SupportedspanarrangementR16
}
