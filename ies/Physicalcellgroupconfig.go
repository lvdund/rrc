package ies

import "rrc/utils"

// PhysicalCellGroupConfig ::= SEQUENCE
// Extensible
type Physicalcellgroupconfig struct {
	HarqAckSpatialbundlingpucch                        *utils.BOOLEAN
	HarqAckSpatialbundlingpusch                        *utils.BOOLEAN
	PNrFr1                                             *PMax
	PdschHarqAckCodebook                               PhysicalcellgroupconfigPdschHarqAckCodebook
	TpcSrsRnti                                         *RntiValue
	TpcPucchRnti                                       *RntiValue
	TpcPuschRnti                                       *RntiValue
	SpCsiRnti                                          *RntiValue
	CsRnti                                             *utils.Setuprelease[RntiValue]
	McsCRnti                                           *RntiValue                                                                 `ext`
	PUeFr1                                             *PMax                                                                      `ext`
	Xscale                                             *PhysicalcellgroupconfigXscale                                             `ext`
	PdcchBlinddetection                                *utils.Setuprelease[PdcchBlinddetection]                                   `ext`
	DcpConfigR16                                       *utils.Setuprelease[DcpConfigR16]                                          `ext`
	HarqAckSpatialbundlingpucchSecondarypucchgroupR16  *PhysicalcellgroupconfigHarqAckSpatialbundlingpucchSecondarypucchgroupR16  `ext`
	HarqAckSpatialbundlingpuschSecondarypucchgroupR16  *PhysicalcellgroupconfigHarqAckSpatialbundlingpuschSecondarypucchgroupR16  `ext`
	PdschHarqAckCodebookSecondarypucchgroupR16         *PhysicalcellgroupconfigPdschHarqAckCodebookSecondarypucchgroupR16         `ext`
	PNrFr2R16                                          *PMax                                                                      `ext`
	PUeFr2R16                                          *PMax                                                                      `ext`
	NrdcPcmodeFr1R16                                   *PhysicalcellgroupconfigNrdcPcmodeFr1R16                                   `ext`
	NrdcPcmodeFr2R16                                   *PhysicalcellgroupconfigNrdcPcmodeFr2R16                                   `ext`
	PdschHarqAckCodebookR16                            *PhysicalcellgroupconfigPdschHarqAckCodebookR16                            `ext`
	NfiTotaldaiIncludedR16                             *utils.BOOLEAN                                                             `ext`
	UlTotaldaiIncludedR16                              *utils.BOOLEAN                                                             `ext`
	PdschHarqAckOneshotfeedbackR16                     *utils.BOOLEAN                                                             `ext`
	PdschHarqAckOneshotfeedbackndiR16                  *utils.BOOLEAN                                                             `ext`
	PdschHarqAckOneshotfeedbackcbgR16                  *utils.BOOLEAN                                                             `ext`
	Downlinkassignmentindexdci02R16                    *PhysicalcellgroupconfigDownlinkassignmentindexdci02R16                    `ext`
	Downlinkassignmentindexdci12R16                    *PhysicalcellgroupconfigDownlinkassignmentindexdci12R16                    `ext`
	PdschHarqAckCodebooklistR16                        *utils.Setuprelease[PdschHarqAckCodebooklistR16]                           `ext`
	AcknackfeedbackmodeR16                             *PhysicalcellgroupconfigAcknackfeedbackmodeR16                             `ext`
	PdcchBlinddetectioncaCombindicatorR16              *utils.Setuprelease[PdcchBlinddetectioncaCombindicatorR16]                 `ext`
	PdcchBlinddetection2R16                            *utils.Setuprelease[PdcchBlinddetection2R16]                               `ext`
	PdcchBlinddetection3R16                            *utils.Setuprelease[PdcchBlinddetection3R16]                               `ext`
	BdfactorrR16                                       *PhysicalcellgroupconfigBdfactorrR16                                       `ext`
	PdschHarqAckEnhtype3toaddmodlistR17                *[]PdschHarqAckEnhtype3R17                                                 `lb:1,ub:maxNrofEnhType3HARQAckR17,ext`
	PdschHarqAckEnhtype3toreleaselistR17               *[]PdschHarqAckEnhtype3indexR17                                            `lb:1,ub:maxNrofEnhType3HARQAckR17,ext`
	PdschHarqAckEnhtype3secondarytoaddmodlistR17       *[]PdschHarqAckEnhtype3R17                                                 `lb:1,ub:maxNrofEnhType3HARQAckR17,ext`
	PdschHarqAckEnhtype3secondarytoreleaselistR17      *[]PdschHarqAckEnhtype3indexR17                                            `lb:1,ub:maxNrofEnhType3HARQAckR17,ext`
	PdschHarqAckEnhtype3dciFieldsecondarypucchgroupR17 *PhysicalcellgroupconfigPdschHarqAckEnhtype3dciFieldsecondarypucchgroupR17 `ext`
	PdschHarqAckEnhtype3dciFieldR17                    *PhysicalcellgroupconfigPdschHarqAckEnhtype3dciFieldR17                    `ext`
	PdschHarqAckRetxR17                                *PhysicalcellgroupconfigPdschHarqAckRetxR17                                `ext`
	PdschHarqAckRetxsecondarypucchgroupR17             *PhysicalcellgroupconfigPdschHarqAckRetxsecondarypucchgroupR17             `ext`
	PucchSscellR17                                     *Scellindex                                                                `ext`
	PucchSscellsecondarypucchgroupR17                  *Scellindex                                                                `ext`
	PucchSscelldynR17                                  *PhysicalcellgroupconfigPucchSscelldynR17                                  `ext`
	PucchSscelldynsecondarypucchgroupR17               *PhysicalcellgroupconfigPucchSscelldynsecondarypucchgroupR17               `ext`
	PucchSscellpatternR17                              *[]utils.INTEGER                                                           `lb:1,ub:maxNrofSlots,ext`
	PucchSscellpatternsecondarypucchgroupR17           *[]utils.INTEGER                                                           `lb:1,ub:maxNrofSlots,ext`
	UciMuxwithdiffprioR17                              *PhysicalcellgroupconfigUciMuxwithdiffprioR17                              `ext`
	UciMuxwithdiffpriosecondarypucchgroupR17           *PhysicalcellgroupconfigUciMuxwithdiffpriosecondarypucchgroupR17           `ext`
	SimultaneouspucchPuschR17                          *PhysicalcellgroupconfigSimultaneouspucchPuschR17                          `ext`
	SimultaneouspucchPuschSecondarypucchgroupR17       *PhysicalcellgroupconfigSimultaneouspucchPuschSecondarypucchgroupR17       `ext`
	PriolowdgHighcgR17                                 *PhysicalcellgroupconfigPriolowdgHighcgR17                                 `ext`
	PriohighdgLowcgR17                                 *PhysicalcellgroupconfigPriohighdgLowcgR17                                 `ext`
	TwoqcltypedforpdcchrepetitionR17                   *PhysicalcellgroupconfigTwoqcltypedforpdcchrepetitionR17                   `ext`
	MulticastconfigR17                                 *utils.Setuprelease[MulticastconfigR17]                                    `ext`
	PdcchBlinddetectioncaCombindicatorR17              *utils.Setuprelease[PdcchBlinddetectioncaCombindicatorR17]                 `ext`
	SimultaneoussrPuschDiffpucchGroupsR17              *PhysicalcellgroupconfigSimultaneoussrPuschDiffpucchGroupsR17              `ext`
	IntrabandncPrachSimultxR17                         *PhysicalcellgroupconfigIntrabandncPrachSimultxR17                         `ext`
}
