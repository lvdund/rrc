package ies

import "rrc/utils"

// SRS-PosResource-r16 ::= SEQUENCE
// Extensible
type SrsPosresourceR16 struct {
	SrsPosresourceidR16       SrsPosresourceidR16
	TransmissioncombR16       SrsPosresourceR16TransmissioncombR16
	ResourcemappingR16        SrsPosresourceR16ResourcemappingR16
	FreqdomainshiftR16        utils.INTEGER `lb:0,ub:268`
	FreqhoppingR16            SrsPosresourceR16FreqhoppingR16
	GrouporsequencehoppingR16 SrsPosresourceR16GrouporsequencehoppingR16
	ResourcetypeR16           *SrsPosresourceR16ResourcetypeR16 `ext`
	SequenceidR16             utils.INTEGER                     `lb:0,ub:65535`
	SpatialrelationinfoposR16 *SrsSpatialrelationinfoposR16
}
