package ies

import "rrc/utils"

// BWP-DownlinkDedicated ::= SEQUENCE
// Extensible
type BwpDownlinkdedicated struct {
	PdcchConfig                         *utils.Setuprelease[PdcchConfig]
	PdschConfig                         *utils.Setuprelease[PdschConfig]
	SpsConfig                           *utils.Setuprelease[SpsConfig]
	Radiolinkmonitoringconfig           *utils.Setuprelease[Radiolinkmonitoringconfig]
	SpsConfigtoaddmodlistR16            *SpsConfigtoaddmodlistR16                           `ext`
	SpsConfigtoreleaselistR16           *SpsConfigtoreleaselistR16                          `ext`
	SpsConfigdeactivationstatelistR16   *SpsConfigdeactivationstatelistR16                  `ext`
	BeamfailurerecoveryscellconfigR16   *utils.Setuprelease[BeamfailurerecoveryrsconfigR16] `ext`
	SlPdcchConfigR16                    *utils.Setuprelease[PdcchConfig]                    `ext`
	SlV2xPdcchConfigR16                 *utils.Setuprelease[PdcchConfig]                    `ext`
	PreconfgapstatusR17                 *utils.BITSTRING                                    `lb:maxNrofGapIdR17,ub:maxNrofGapIdR17,ext`
	BeamfailurerecoveryspcellconfigR17  *utils.Setuprelease[BeamfailurerecoveryrsconfigR16] `ext`
	HarqFeedbackenablingforspsactiveR17 *utils.BOOLEAN                                      `ext`
	CfrConfigmulticastR17               *utils.Setuprelease[CfrConfigmulticastR17]          `ext`
	DlPpwPreconfigtoaddmodlistR17       *DlPpwPreconfigtoaddmodlistR17                      `ext`
	DlPpwPreconfigtoreleaselistR17      *DlPpwPreconfigtoreleaselistR17                     `ext`
	NoncelldefiningssbR17               *NoncelldefiningssbR17                              `ext`
	ServingcellmoR17                    *Measobjectid                                       `ext`
}
