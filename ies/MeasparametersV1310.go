package ies

import "rrc/utils"

// MeasParameters-v1310 ::= SEQUENCE
type MeasparametersV1310 struct {
	RsSinrMeasR13                       *utils.ENUMERATED
	WhitecelllistR13                    *utils.ENUMERATED
	ExtendedmaxobjectidR13              *utils.ENUMERATED
	UlPdcpDelayR13                      *utils.ENUMERATED
	ExtendedfreqprioritiesR13           *utils.ENUMERATED
	MultibandinforeportR13              *utils.ENUMERATED
	RssiAndchanneloccupancyreportingR13 *utils.ENUMERATED
}
