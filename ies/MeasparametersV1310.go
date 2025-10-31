package ies

// MeasParameters-v1310 ::= SEQUENCE
type MeasparametersV1310 struct {
	RsSinrMeasR13                       *MeasparametersV1310RsSinrMeasR13
	WhitecelllistR13                    *MeasparametersV1310WhitecelllistR13
	ExtendedmaxobjectidR13              *MeasparametersV1310ExtendedmaxobjectidR13
	UlPdcpDelayR13                      *MeasparametersV1310UlPdcpDelayR13
	ExtendedfreqprioritiesR13           *MeasparametersV1310ExtendedfreqprioritiesR13
	MultibandinforeportR13              *MeasparametersV1310MultibandinforeportR13
	RssiAndchanneloccupancyreportingR13 *MeasparametersV1310RssiAndchanneloccupancyreportingR13
}
