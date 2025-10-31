package ies

import "rrc/utils"

// PUSCH-Config ::= SEQUENCE
// Extensible
type PuschConfig struct {
	Datascramblingidentitypusch                   *utils.INTEGER `lb:0,ub:1023`
	Txconfig                                      *PuschConfigTxconfig
	DmrsUplinkforpuschMappingtypea                *utils.Setuprelease[DmrsUplinkconfig]
	DmrsUplinkforpuschMappingtypeb                *utils.Setuprelease[DmrsUplinkconfig]
	PuschPowercontrol                             *PuschPowercontrol
	Frequencyhopping                              *PuschConfigFrequencyhopping
	Frequencyhoppingoffsetlists                   *[]utils.INTEGER `lb:1,ub:4`
	Resourceallocation                            PuschConfigResourceallocation
	PuschTimedomainallocationlist                 *utils.Setuprelease[PuschTimedomainresourceallocationlist]
	PuschAggregationfactor                        *PuschConfigPuschAggregationfactor
	McsTable                                      *PuschConfigMcsTable
	McsTabletransformprecoder                     *PuschConfigMcsTabletransformprecoder
	Transformprecoder                             *PuschConfigTransformprecoder
	Codebooksubset                                *PuschConfigCodebooksubset
	Maxrank                                       *utils.INTEGER `lb:0,ub:4`
	RbgSize                                       *PuschConfigRbgSize
	UciOnpusch                                    *utils.Setuprelease[UciOnpusch]
	TpPi2bpsk                                     *PuschConfigTpPi2bpsk
	Minimumschedulingoffsetk2R16                  *utils.Setuprelease[Minschedulingoffsetk2ValuesR16]           `ext`
	UlAccessconfiglistdci01R16                    *utils.Setuprelease[UlAccessconfiglistdci01R16]               `ext`
	HarqProcessnumbersizedci02R16                 *utils.INTEGER                                                `lb:0,ub:4,ext`
	DmrsSequenceinitializationdci02R16            *PuschConfigDmrsSequenceinitializationdci02R16                `ext`
	NumberofbitsforrvDci02R16                     *utils.INTEGER                                                `lb:0,ub:2,ext`
	Antennaportsfieldpresencedci02R16             *PuschConfigAntennaportsfieldpresencedci02R16                 `ext`
	DmrsUplinkforpuschMappingtypeaDci02R16        *utils.Setuprelease[DmrsUplinkconfig]                         `ext`
	DmrsUplinkforpuschMappingtypebDci02R16        *utils.Setuprelease[DmrsUplinkconfig]                         `ext`
	Frequencyhoppingdci02R16                      *PuschConfigFrequencyhoppingdci02R16                          `ext`
	Frequencyhoppingoffsetlistsdci02R16           *utils.Setuprelease[Frequencyhoppingoffsetlistsdci02R16]      `ext`
	Codebooksubsetdci02R16                        *PuschConfigCodebooksubsetdci02R16                            `ext`
	Invalidsymbolpatternindicatordci02R16         *PuschConfigInvalidsymbolpatternindicatordci02R16             `ext`
	Maxrankdci02R16                               *utils.INTEGER                                                `lb:0,ub:4,ext`
	McsTabledci02R16                              *PuschConfigMcsTabledci02R16                                  `ext`
	McsTabletransformprecoderdci02R16             *PuschConfigMcsTabletransformprecoderdci02R16                 `ext`
	Priorityindicatordci02R16                     *PuschConfigPriorityindicatordci02R16                         `ext`
	PuschReptypeindicatordci02R16                 *PuschConfigPuschReptypeindicatordci02R16                     `ext`
	Resourceallocationdci02R16                    *PuschConfigResourceallocationdci02R16                        `ext`
	Resourceallocationtype1granularitydci02R16    *PuschConfigResourceallocationtype1granularitydci02R16        `ext`
	UciOnpuschListdci02R16                        *utils.Setuprelease[UciOnpuschListdci02R16]                   `ext`
	PuschTimedomainallocationlistdci02R16         *utils.Setuprelease[PuschTimedomainresourceallocationlistR16] `ext`
	PuschTimedomainallocationlistdci01R16         *utils.Setuprelease[PuschTimedomainresourceallocationlistR16] `ext`
	Invalidsymbolpatternindicatordci01R16         *PuschConfigInvalidsymbolpatternindicatordci01R16             `ext`
	Priorityindicatordci01R16                     *PuschConfigPriorityindicatordci01R16                         `ext`
	PuschReptypeindicatordci01R16                 *PuschConfigPuschReptypeindicatordci01R16                     `ext`
	Frequencyhoppingdci01R16                      *PuschConfigFrequencyhoppingdci01R16                          `ext`
	UciOnpuschListdci01R16                        *utils.Setuprelease[UciOnpuschListdci01R16]                   `ext`
	InvalidsymbolpatternR16                       *InvalidsymbolpatternR16                                      `ext`
	PuschPowercontrolV1610                        *utils.Setuprelease[PuschPowercontrolV1610]                   `ext`
	UlFullpowertransmissionR16                    *PuschConfigUlFullpowertransmissionR16                        `ext`
	PuschTimedomainallocationlistformultipuschR16 *utils.Setuprelease[PuschTimedomainresourceallocationlistR16] `ext`
	NumberofinvalidsymbolsfordlUlSwitchingR16     *utils.INTEGER                                                `lb:0,ub:4,ext`
	UlAccessconfiglistdci02R17                    *utils.Setuprelease[UlAccessconfiglistdci02R17]               `ext`
	Betaoffsetscrosspri0R17                       *utils.Setuprelease[BetaoffsetscrosspriselR17]                `ext`
	Betaoffsetscrosspri1R17                       *utils.Setuprelease[BetaoffsetscrosspriselR17]                `ext`
	Betaoffsetscrosspri0dci02R17                  *utils.Setuprelease[Betaoffsetscrosspriseldci02R17]           `ext`
	Betaoffsetscrosspri1dci02R17                  *utils.Setuprelease[Betaoffsetscrosspriseldci02R17]           `ext`
	MappingpatternR17                             *PuschConfigMappingpatternR17                                 `ext`
	Secondtpcfielddci01R17                        *PuschConfigSecondtpcfielddci01R17                            `ext`
	Secondtpcfielddci02R17                        *PuschConfigSecondtpcfielddci02R17                            `ext`
	SequenceoffsetforrvR17                        *utils.INTEGER                                                `lb:0,ub:3,ext`
	UlAccessconfiglistdci01R17                    *utils.Setuprelease[UlAccessconfiglistdci01R17]               `ext`
	Minimumschedulingoffsetk2R17                  *utils.Setuprelease[Minschedulingoffsetk2ValuesR17]           `ext`
	AvailableslotcountingR17                      *PuschConfigAvailableslotcountingR17                          `ext`
	DmrsBundlingpuschConfigR17                    *utils.Setuprelease[DmrsBundlingpuschConfigR17]               `ext`
	HarqProcessnumbersizedci02V1700               *utils.INTEGER                                                `lb:0,ub:5,ext`
	HarqProcessnumbersizedci01R17                 *utils.INTEGER                                                `lb:0,ub:5,ext`
	MpeResourcepooltoaddmodlistR17                *[]MpeResourceR17                                             `lb:1,ub:maxMPEResourcesR17,ext`
	MpeResourcepooltoreleaselistR17               *[]MpeResourceidR17                                           `lb:1,ub:maxMPEResourcesR17,ext`
}
