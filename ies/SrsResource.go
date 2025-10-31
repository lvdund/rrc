package ies

import "rrc/utils"

// SRS-Resource ::= SEQUENCE
// Extensible
type SrsResource struct {
	SrsResourceid             SrsResourceid
	NrofsrsPorts              SrsResourceNrofsrsPorts
	PtrsPortindex             *SrsResourcePtrsPortindex
	Transmissioncomb          SrsResourceTransmissioncomb
	Resourcemapping           SrsResourceResourcemapping
	Freqdomainposition        utils.INTEGER `lb:0,ub:67`
	Freqdomainshift           utils.INTEGER `lb:0,ub:268`
	Freqhopping               SrsResourceFreqhopping
	Grouporsequencehopping    SrsResourceGrouporsequencehopping
	Resourcetype              SrsResourceResourcetype
	Sequenceid                utils.INTEGER `lb:0,ub:1023`
	Spatialrelationinfo       *SrsSpatialrelationinfo
	ResourcemappingR16        *SrsResourceResourcemappingR16                 `ext`
	SpatialrelationinfoPdcR17 *utils.Setuprelease[SpatialrelationinfoPdcR17] `ext`
	ResourcemappingR17        *SrsResourceResourcemappingR17                 `ext`
	PartialfreqsoundingR17    *SrsResourcePartialfreqsoundingR17             `ext`
	TransmissioncombN8R17     *SrsResourceTransmissioncombN8R17              `ext`
	SrsTciStateR17            *SrsResourceSrsTciStateR17                     `ext`
	RepetitionfactorV1730     *SrsResourceRepetitionfactorV1730              `ext`
	SrsDlorjointtciStateV1730 *SrsResourceSrsDlorjointtciStateV1730          `ext`
}
