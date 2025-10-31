package ies

import "rrc/utils"

// ServingCellConfig ::= SEQUENCE
// Extensible
type Servingcellconfig struct {
	TddUlDlConfigurationdedicated               *TddUlDlConfigdedicated
	Initialdownlinkbwp                          *BwpDownlinkdedicated
	DownlinkbwpToreleaselist                    *[]BwpId       `lb:1,ub:maxNrofBWPs`
	DownlinkbwpToaddmodlist                     *[]BwpDownlink `lb:1,ub:maxNrofBWPs`
	FirstactivedownlinkbwpId                    *BwpId
	BwpInactivitytimer                          *ServingcellconfigBwpInactivitytimer
	DefaultdownlinkbwpId                        *BwpId
	Uplinkconfig                                *Uplinkconfig
	Supplementaryuplink                         *Uplinkconfig
	PdcchServingcellconfig                      *utils.Setuprelease[PdcchServingcellconfig]
	PdschServingcellconfig                      *utils.Setuprelease[PdschServingcellconfig]
	CsiMeasconfig                               *utils.Setuprelease[CsiMeasconfig]
	Scelldeactivationtimer                      *ServingcellconfigScelldeactivationtimer
	Crosscarrierschedulingconfig                *Crosscarrierschedulingconfig
	TagId                                       TagId
	Dummy1                                      *ServingcellconfigDummy1
	Pathlossreferencelinking                    *ServingcellconfigPathlossreferencelinking
	Servingcellmo                               *Measobjectid
	LteCrsTomatcharound                         *utils.Setuprelease[RatematchpatternlteCrs]                   `ext`
	Ratematchpatterntoaddmodlist                *[]Ratematchpattern                                           `lb:1,ub:maxNrofRateMatchPatterns,ext`
	Ratematchpatterntoreleaselist               *[]Ratematchpatternid                                         `lb:1,ub:maxNrofRateMatchPatterns,ext`
	DownlinkchannelbwPerscsList                 *[]ScsSpecificcarrier                                         `lb:1,ub:maxSCSs,ext`
	SupplementaryuplinkreleaseR16               *utils.BOOLEAN                                                `ext`
	TddUlDlConfigurationdedicatedIabMtR16       *TddUlDlConfigdedicatedIabMtR16                               `ext`
	DormantbwpConfigR16                         *utils.Setuprelease[DormantbwpConfigR16]                      `ext`
	CaSlotoffsetR16                             *ServingcellconfigCaSlotoffsetR16                             `ext`
	Dummy2                                      *utils.Setuprelease[Dummyj]                                   `ext`
	IntracellguardbandsdlListR16                *[]IntracellguardbandsperscsR16                               `lb:1,ub:maxSCSs,ext`
	IntracellguardbandsulListR16                *[]IntracellguardbandsperscsR16                               `lb:1,ub:maxSCSs,ext`
	CsiRsValidationwithdciR16                   *ServingcellconfigCsiRsValidationwithdciR16                   `ext`
	LteCrsPatternlist1R16                       *utils.Setuprelease[LteCrsPatternlistR16]                     `ext`
	LteCrsPatternlist2R16                       *utils.Setuprelease[LteCrsPatternlistR16]                     `ext`
	CrsRatematchPercoresetpoolindexR16          *ServingcellconfigCrsRatematchPercoresetpoolindexR16          `ext`
	EnabletwodefaulttciStatesR16                *ServingcellconfigEnabletwodefaulttciStatesR16                `ext`
	EnabledefaulttciStatepercoresetpoolindexR16 *ServingcellconfigEnabledefaulttciStatepercoresetpoolindexR16 `ext`
	EnablebeamswitchtimingR16                   *utils.BOOLEAN                                                `ext`
	CbgTxdifftbsprocessingtype1R16              *ServingcellconfigCbgTxdifftbsprocessingtype1R16              `ext`
	CbgTxdifftbsprocessingtype2R16              *ServingcellconfigCbgTxdifftbsprocessingtype2R16              `ext`
	DirectionalcollisionhandlingR16             *ServingcellconfigDirectionalcollisionhandlingR16             `ext`
	ChannelaccessconfigR16                      *utils.Setuprelease[ChannelaccessconfigR16]                   `ext`
	NrDlPrsPdcInfoR17                           *utils.Setuprelease[NrDlPrsPdcInfoR17]                        `ext`
	SemistaticchannelaccessconfigueR17          *utils.Setuprelease[SemistaticchannelaccessconfigueR17]       `ext`
	MimoparamR17                                *utils.Setuprelease[MimoparamR17]                             `ext`
	Channelaccessmode2R17                       *ServingcellconfigChannelaccessmode2R17                       `ext`
	TimedomainharqBundlingtype1R17              *ServingcellconfigTimedomainharqBundlingtype1R17              `ext`
	NrofharqBundlinggroupsR17                   *ServingcellconfigNrofharqBundlinggroupsR17                   `ext`
	FdmedReceptionmulticastR17                  *utils.BOOLEAN                                                `ext`
	MorethanonenackonlymodeR17                  *ServingcellconfigMorethanonenackonlymodeR17                  `ext`
	TciActivatedconfigR17                       *TciActivatedconfigR17                                        `ext`
	DirectionalcollisionhandlingDcR17           *ServingcellconfigDirectionalcollisionhandlingDcR17           `ext`
	LteNeighcellscrsAssistinfolistR17           *utils.Setuprelease[LteNeighcellscrsAssistinfolistR17]        `ext`
	LteNeighcellscrsAssumptionsR17              *ServingcellconfigLteNeighcellscrsAssumptionsR17              `ext`
}
