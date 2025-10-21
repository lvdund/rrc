package ies

import "rrc/utils"

// MeasParameters-v1610 ::= SEQUENCE
type MeasparametersV1610 struct {
	BandinfonrV1610                  *interface{}
	AltfreqpriorityR16               *utils.ENUMERATED
	CeDlChannelqualityreportingR16   *utils.ENUMERATED
	CeMeasrssDedicatedR16            *utils.ENUMERATED
	EutraIdleinactivemeasurementsR16 *utils.ENUMERATED
	NrIdleinactivemeasfr1R16         *utils.ENUMERATED
	NrIdleinactivemeasfr2R16         *utils.ENUMERATED
	IdleinactivevalidityarealistR16  *utils.ENUMERATED
	MeasgappatternsNronlyR16         *utils.ENUMERATED
	MeasgappatternsNronlyEndcR16     *utils.ENUMERATED
}
