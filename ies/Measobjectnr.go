package ies

import "rrc/utils"

// MeasObjectNR ::= SEQUENCE
// Extensible
type Measobjectnr struct {
	Ssbfrequency                   *ArfcnValuenr
	Ssbsubcarrierspacing           *Subcarrierspacing
	Smtc1                          *SsbMtc
	Smtc2                          *SsbMtc2
	ReffreqcsiRs                   *ArfcnValuenr
	Referencesignalconfig          Referencesignalconfig
	AbsthreshssBlocksconsolidation *Thresholdnr
	AbsthreshcsiRsConsolidation    *Thresholdnr
	NrofssBlockstoaverage          *utils.INTEGER `lb:0,ub:maxNrofSSBlockstoaverage`
	NrofcsiRsResourcestoaverage    *utils.INTEGER `lb:0,ub:maxNrofCSIRsResourcestoaverage`
	Quantityconfigindex            utils.INTEGER  `lb:0,ub:maxNrofQuantityConfig`
	Offsetmo                       QOffsetrangelist
	Cellstoremovelist              *PciList
	Cellstoaddmodlist              *Cellstoaddmodlist
	Excludedcellstoremovelist      *PciRangeindexlist
	Excludedcellstoaddmodlist      *[]PciRangeelement `lb:1,ub:maxNrofPCIRanges`
	Allowedcellstoremovelist       *PciRangeindexlist
	Allowedcellstoaddmodlist       *[]PciRangeelement                 `lb:1,ub:maxNrofPCIRanges`
	Freqbandindicatornr            *Freqbandindicatornr               `ext`
	Meascyclescell                 *MeasobjectnrMeascyclescell        `ext`
	Smtc3listR16                   *SsbMtc3listR16                    `ext`
	RmtcConfigR16                  *utils.Setuprelease[RmtcConfigR16] `ext`
	T312R16                        *utils.Setuprelease[T312R16]       `ext`
	AssociatedmeasgapssbR17        *MeasgapidR17                      `ext`
	AssociatedmeasgapcsirsR17      *MeasgapidR17                      `ext`
	Smtc4listR17                   *SsbMtc4listR17                    `ext`
	MeascyclepscellR17             *MeasobjectnrMeascyclepscellR17    `ext`
	CellstoaddmodlistextV1710      *CellstoaddmodlistextV1710         `ext`
	Associatedmeasgapssb2V1720     *MeasgapidR17                      `ext`
	Associatedmeasgapcsirs2V1720   *MeasgapidR17                      `ext`
}
