package ies

import "rrc/utils"

// RA-InformationCommon-r16 ::= SEQUENCE
// Extensible
type RaInformationcommonR16 struct {
	AbsolutefrequencypointaR16                ArfcnValuenr
	LocationandbandwidthR16                   utils.INTEGER `lb:0,ub:37949`
	SubcarrierspacingR16                      Subcarrierspacing
	Msg1FrequencystartR16                     *utils.INTEGER `lb:0,ub:maxNrofPhysicalResourceBlocks1`
	Msg1FrequencystartcfraR16                 *utils.INTEGER `lb:0,ub:maxNrofPhysicalResourceBlocks1`
	Msg1SubcarrierspacingR16                  *Subcarrierspacing
	Msg1SubcarrierspacingcfraR16              *Subcarrierspacing
	Msg1FdmR16                                *RaInformationcommonR16Msg1FdmR16
	Msg1FdmcfraR16                            *RaInformationcommonR16Msg1FdmcfraR16
	PerrainfolistR16                          PerrainfolistR16
	PerrainfolistV1660                        *PerrainfolistV1660                                              `ext`
	Msg1ScsFromPrachConfigurationindexR16     *RaInformationcommonR16Msg1ScsFromPrachConfigurationindexR16     `ext`
	Msg1ScsFromPrachConfigurationindexcfraR16 *RaInformationcommonR16Msg1ScsFromPrachConfigurationindexcfraR16 `ext`
	MsgaRoFrequencystartR17                   *utils.INTEGER                                                   `lb:0,ub:maxNrofPhysicalResourceBlocks1,ext`
	MsgaRoFrequencystartcfraR17               *utils.INTEGER                                                   `lb:0,ub:maxNrofPhysicalResourceBlocks1,ext`
	MsgaSubcarrierspacingR17                  *Subcarrierspacing                                               `ext`
	MsgaRoFdmR17                              *RaInformationcommonR16MsgaRoFdmR17                              `ext`
	MsgaRoFdmcfraR17                          *RaInformationcommonR16MsgaRoFdmcfraR17                          `ext`
	MsgaScsFromPrachConfigurationindexR17     *RaInformationcommonR16MsgaScsFromPrachConfigurationindexR17     `ext`
	MsgaTransmaxR17                           *RaInformationcommonR16MsgaTransmaxR17                           `ext`
	MsgaMcsR17                                *utils.INTEGER                                                   `lb:0,ub:15,ext`
	NrofprbsPermsgaPoR17                      *utils.INTEGER                                                   `lb:0,ub:32,ext`
	MsgaPuschTimedomainallocationR17          *utils.INTEGER                                                   `lb:0,ub:maxNrofULAllocations,ext`
	FrequencystartmsgaPuschR17                *utils.INTEGER                                                   `lb:0,ub:maxNrofPhysicalResourceBlocks1,ext`
	NrofmsgaPoFdmR17                          *RaInformationcommonR16NrofmsgaPoFdmR17                          `ext`
	DlpathlossrsrpR17                         *RsrpRange                                                       `ext`
	IntendedsibsR17                           *[]SibTypeR17                                                    `lb:1,ub:maxSIB,ext`
	SsbsforsiAcquisitionR17                   *[]SsbIndex                                                      `lb:1,ub:maxNrofSSBsR16,ext`
	MsgaPuschPayloadsizeR17                   *utils.BITSTRING                                                 `lb:5,ub:5,ext`
	OndemandsisuccessR17                      *utils.BOOLEAN                                                   `ext`
}
