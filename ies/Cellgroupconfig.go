package ies

import "rrc/utils"

// CellGroupConfig ::= SEQUENCE
// Extensible
type Cellgroupconfig struct {
	Cellgroupid                               Cellgroupid
	RlcBearertoaddmodlist                     *[]RlcBearerconfig        `lb:1,ub:maxLCId`
	RlcBearertoreleaselist                    *[]Logicalchannelidentity `lb:1,ub:maxLCId`
	MacCellgroupconfig                        *MacCellgroupconfig
	Physicalcellgroupconfig                   *Physicalcellgroupconfig
	Spcellconfig                              *Spcellconfig
	Scelltoaddmodlist                         *[]Scellconfig                                    `lb:1,ub:maxNrofSCells`
	Scelltoreleaselist                        *[]Scellindex                                     `lb:1,ub:maxNrofSCells`
	Reportuplinktxdirectcurrent               *utils.BOOLEAN                                    `ext`
	BapAddressR16                             *utils.BITSTRING                                  `lb:10,ub:10,ext`
	BhRlcChanneltoaddmodlistR16               *[]BhRlcChannelconfigR16                          `lb:1,ub:maxBHRlcChannelidR16,ext`
	BhRlcChanneltoreleaselistR16              *[]BhRlcChannelidR16                              `lb:1,ub:maxBHRlcChannelidR16,ext`
	F1cTransferpathR16                        *CellgroupconfigF1cTransferpathR16                `ext`
	SimultaneoustciUpdatelist1R16             *[]Servcellindex                                  `lb:1,ub:maxNrofServingCellsTCIR16,ext`
	SimultaneoustciUpdatelist2R16             *[]Servcellindex                                  `lb:1,ub:maxNrofServingCellsTCIR16,ext`
	SimultaneousspatialUpdatedlist1R16        *[]Servcellindex                                  `lb:1,ub:maxNrofServingCellsTCIR16,ext`
	SimultaneousspatialUpdatedlist2R16        *[]Servcellindex                                  `lb:1,ub:maxNrofServingCellsTCIR16,ext`
	UplinktxswitchingoptionR16                *CellgroupconfigUplinktxswitchingoptionR16        `ext`
	UplinktxswitchingpowerboostingR16         *CellgroupconfigUplinktxswitchingpowerboostingR16 `ext`
	ReportuplinktxdirectcurrenttwocarrierR16  *utils.BOOLEAN                                    `ext`
	F1cTransferpathnrdcR17                    *CellgroupconfigF1cTransferpathnrdcR17            `ext`
	Uplinktxswitching2tModeR17                *CellgroupconfigUplinktxswitching2tModeR17        `ext`
	UplinktxswitchingDualulTxstateR17         *CellgroupconfigUplinktxswitchingDualulTxstateR17 `ext`
	UuRelayrlcChanneltoaddmodlistR17          *[]UuRelayrlcChannelconfigR17                     `lb:1,ub:maxUuRelayrlcChannelidR17,ext`
	UuRelayrlcChanneltoreleaselistR17         *[]UuRelayrlcChannelidR17                         `lb:1,ub:maxUuRelayrlcChannelidR17,ext`
	SimultaneousuTciUpdatelist1R17            *[]Servcellindex                                  `lb:1,ub:maxNrofServingCellsTCIR16,ext`
	SimultaneousuTciUpdatelist2R17            *[]Servcellindex                                  `lb:1,ub:maxNrofServingCellsTCIR16,ext`
	SimultaneousuTciUpdatelist3R17            *[]Servcellindex                                  `lb:1,ub:maxNrofServingCellsTCIR16,ext`
	SimultaneousuTciUpdatelist4R17            *[]Servcellindex                                  `lb:1,ub:maxNrofServingCellsTCIR16,ext`
	RlcBearertoreleaselistextR17              *[]LogicalchannelidentityextR17                   `lb:1,ub:maxLCId,ext`
	IabResourceconfigtoaddmodlistR17          *[]IabResourceconfigR17                           `lb:1,ub:maxNrofIABResourceConfigR17,ext`
	IabResourceconfigtoreleaselistR17         *[]IabResourceconfigidR17                         `lb:1,ub:maxNrofIABResourceConfigR17,ext`
	ReportuplinktxdirectcurrentmorecarrierR17 *ReportuplinktxdirectcurrentmorecarrierR17        `ext`
}
