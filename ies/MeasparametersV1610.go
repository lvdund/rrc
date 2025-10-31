package ies

// MeasParameters-v1610 ::= SEQUENCE
type MeasparametersV1610 struct {
	BandinfonrV1610                  *[]MeasgapinfonrR16 `lb:1,ub:maxBands`
	AltfreqpriorityR16               *MeasparametersV1610AltfreqpriorityR16
	CeDlChannelqualityreportingR16   *MeasparametersV1610CeDlChannelqualityreportingR16
	CeMeasrssDedicatedR16            *MeasparametersV1610CeMeasrssDedicatedR16
	EutraIdleinactivemeasurementsR16 *MeasparametersV1610EutraIdleinactivemeasurementsR16
	NrIdleinactivemeasfr1R16         *MeasparametersV1610NrIdleinactivemeasfr1R16
	NrIdleinactivemeasfr2R16         *MeasparametersV1610NrIdleinactivemeasfr2R16
	IdleinactivevalidityarealistR16  *MeasparametersV1610IdleinactivevalidityarealistR16
	MeasgappatternsNronlyR16         *MeasparametersV1610MeasgappatternsNronlyR16
	MeasgappatternsNronlyEndcR16     *MeasparametersV1610MeasgappatternsNronlyEndcR16
}
