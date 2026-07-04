// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: CellGroupConfig (line 5596).

var cellGroupConfigConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "cellGroupId"},
		{Name: "rlc-BearerToAddModList", Optional: true},
		{Name: "rlc-BearerToReleaseList", Optional: true},
		{Name: "mac-CellGroupConfig", Optional: true},
		{Name: "physicalCellGroupConfig", Optional: true},
		{Name: "spCellConfig", Optional: true},
		{Name: "sCellToAddModList", Optional: true},
		{Name: "sCellToReleaseList", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
		{Name: "extension-addition-1", Optional: true},
		{Name: "extension-addition-2", Optional: true},
		{Name: "extension-addition-3", Optional: true},
		{Name: "extension-addition-4", Optional: true},
		{Name: "extension-addition-5", Optional: true},
		{Name: "extension-addition-6", Optional: true},
		{Name: "extension-addition-7", Optional: true},
	},
}

var cellGroupConfigRlcBearerToAddModListConstraints = per.SizeRange(1, common.MaxLC_ID)

var cellGroupConfigRlcBearerToReleaseListConstraints = per.SizeRange(1, common.MaxLC_ID)

var cellGroupConfigSCellToAddModListConstraints = per.SizeRange(1, common.MaxNrofSCells)

var cellGroupConfigSCellToReleaseListConstraints = per.SizeRange(1, common.MaxNrofSCells)

const (
	CellGroupConfig_Ext_ReportUplinkTxDirectCurrent_True = 0
)

var cellGroupConfigExtReportUplinkTxDirectCurrentConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CellGroupConfig_Ext_ReportUplinkTxDirectCurrent_True},
}

var cellGroupConfigBapAddressR16Constraints = per.FixedSize(10)

var cellGroupConfigExtBhRLCChannelToAddModListR16Constraints = per.SizeRange(1, common.MaxBH_RLC_ChannelID_r16)

var cellGroupConfigExtBhRLCChannelToReleaseListR16Constraints = per.SizeRange(1, common.MaxBH_RLC_ChannelID_r16)

const (
	CellGroupConfig_Ext_F1c_TransferPath_r16_Lte  = 0
	CellGroupConfig_Ext_F1c_TransferPath_r16_Nr   = 1
	CellGroupConfig_Ext_F1c_TransferPath_r16_Both = 2
)

var cellGroupConfigExtF1cTransferPathR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CellGroupConfig_Ext_F1c_TransferPath_r16_Lte, CellGroupConfig_Ext_F1c_TransferPath_r16_Nr, CellGroupConfig_Ext_F1c_TransferPath_r16_Both},
}

var cellGroupConfigExtSimultaneousTCIUpdateList1R16Constraints = per.SizeRange(1, common.MaxNrofServingCellsTCI_r16)

var cellGroupConfigExtSimultaneousTCIUpdateList2R16Constraints = per.SizeRange(1, common.MaxNrofServingCellsTCI_r16)

var cellGroupConfigExtSimultaneousSpatialUpdatedList1R16Constraints = per.SizeRange(1, common.MaxNrofServingCellsTCI_r16)

var cellGroupConfigExtSimultaneousSpatialUpdatedList2R16Constraints = per.SizeRange(1, common.MaxNrofServingCellsTCI_r16)

const (
	CellGroupConfig_Ext_UplinkTxSwitchingOption_r16_SwitchedUL = 0
	CellGroupConfig_Ext_UplinkTxSwitchingOption_r16_DualUL     = 1
)

var cellGroupConfigExtUplinkTxSwitchingOptionR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CellGroupConfig_Ext_UplinkTxSwitchingOption_r16_SwitchedUL, CellGroupConfig_Ext_UplinkTxSwitchingOption_r16_DualUL},
}

const (
	CellGroupConfig_Ext_UplinkTxSwitchingPowerBoosting_r16_Enabled = 0
)

var cellGroupConfigExtUplinkTxSwitchingPowerBoostingR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CellGroupConfig_Ext_UplinkTxSwitchingPowerBoosting_r16_Enabled},
}

const (
	CellGroupConfig_Ext_ReportUplinkTxDirectCurrentTwoCarrier_r16_True = 0
)

var cellGroupConfigExtReportUplinkTxDirectCurrentTwoCarrierR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CellGroupConfig_Ext_ReportUplinkTxDirectCurrentTwoCarrier_r16_True},
}

const (
	CellGroupConfig_Ext_F1c_TransferPathNRDC_r17_Mcg  = 0
	CellGroupConfig_Ext_F1c_TransferPathNRDC_r17_Scg  = 1
	CellGroupConfig_Ext_F1c_TransferPathNRDC_r17_Both = 2
)

var cellGroupConfigExtF1cTransferPathNRDCR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CellGroupConfig_Ext_F1c_TransferPathNRDC_r17_Mcg, CellGroupConfig_Ext_F1c_TransferPathNRDC_r17_Scg, CellGroupConfig_Ext_F1c_TransferPathNRDC_r17_Both},
}

const (
	CellGroupConfig_Ext_UplinkTxSwitching_2T_Mode_r17_Enabled = 0
)

var cellGroupConfigExtUplinkTxSwitching2TModeR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CellGroupConfig_Ext_UplinkTxSwitching_2T_Mode_r17_Enabled},
}

const (
	CellGroupConfig_Ext_UplinkTxSwitching_DualUL_TxState_r17_OneT = 0
	CellGroupConfig_Ext_UplinkTxSwitching_DualUL_TxState_r17_TwoT = 1
)

var cellGroupConfigExtUplinkTxSwitchingDualULTxStateR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CellGroupConfig_Ext_UplinkTxSwitching_DualUL_TxState_r17_OneT, CellGroupConfig_Ext_UplinkTxSwitching_DualUL_TxState_r17_TwoT},
}

var cellGroupConfigExtUuRelayRLCChannelToAddModListR17Constraints = per.SizeRange(1, common.MaxUu_RelayRLC_ChannelID_r17)

var cellGroupConfigExtUuRelayRLCChannelToReleaseListR17Constraints = per.SizeRange(1, common.MaxUu_RelayRLC_ChannelID_r17)

var cellGroupConfigExtSimultaneousUTCIUpdateList1R17Constraints = per.SizeRange(1, common.MaxNrofServingCellsTCI_r16)

var cellGroupConfigExtSimultaneousUTCIUpdateList2R17Constraints = per.SizeRange(1, common.MaxNrofServingCellsTCI_r16)

var cellGroupConfigExtSimultaneousUTCIUpdateList3R17Constraints = per.SizeRange(1, common.MaxNrofServingCellsTCI_r16)

var cellGroupConfigExtSimultaneousUTCIUpdateList4R17Constraints = per.SizeRange(1, common.MaxNrofServingCellsTCI_r16)

var cellGroupConfigExtRlcBearerToReleaseListExtR17Constraints = per.SizeRange(1, common.MaxLC_ID)

var cellGroupConfigExtIabResourceConfigToAddModListR17Constraints = per.SizeRange(1, common.MaxNrofIABResourceConfig_r17)

var cellGroupConfigExtIabResourceConfigToReleaseListR17Constraints = per.SizeRange(1, common.MaxNrofIABResourceConfig_r17)

const (
	CellGroupConfig_Ext_PrioSCellPRACH_OverSP_PeriodicSRS_r17_Enabled = 0
)

var cellGroupConfigExtPrioSCellPRACHOverSPPeriodicSRSR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CellGroupConfig_Ext_PrioSCellPRACH_OverSP_PeriodicSRS_r17_Enabled},
}

var cellGroupConfigExtNcrFwdConfigR18Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	CellGroupConfig_Ext_Ncr_FwdConfig_r18_Release = 0
	CellGroupConfig_Ext_Ncr_FwdConfig_r18_Setup   = 1
)

var cellGroupConfigExtAutonomousDenialParametersR18Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	CellGroupConfig_Ext_AutonomousDenialParameters_r18_Release = 0
	CellGroupConfig_Ext_AutonomousDenialParameters_r18_Setup   = 1
)

const (
	CellGroupConfig_Ext_NonCollocatedTypeMRDC_r18_True = 0
)

var cellGroupConfigExtNonCollocatedTypeMRDCR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CellGroupConfig_Ext_NonCollocatedTypeMRDC_r18_True},
}

const (
	CellGroupConfig_Ext_NonCollocatedTypeNR_CA_r18_True = 0
)

var cellGroupConfigExtNonCollocatedTypeNRCAR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CellGroupConfig_Ext_NonCollocatedTypeNR_CA_r18_True},
}

var cellGroupConfigExtUplinkTxSwitchingMoreBandsR18Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	CellGroupConfig_Ext_UplinkTxSwitchingMoreBands_r18_Release = 0
	CellGroupConfig_Ext_UplinkTxSwitchingMoreBands_r18_Setup   = 1
)

const (
	CellGroupConfig_Ext_UplinkTxSwitching3Tx_r19_True = 0
)

var cellGroupConfigExtUplinkTxSwitching3TxR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CellGroupConfig_Ext_UplinkTxSwitching3Tx_r19_True},
}

var cellGroupConfigExtLowBandCASwitchingR19Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	CellGroupConfig_Ext_LowBandCA_Switching_r19_Release = 0
	CellGroupConfig_Ext_LowBandCA_Switching_r19_Setup   = 1
)

const (
	CellGroupConfig_Ext_NonCollocatedTypeMRDC_v1900_Type1 = 0
	CellGroupConfig_Ext_NonCollocatedTypeMRDC_v1900_Type4 = 1
)

var cellGroupConfigExtNonCollocatedTypeMRDCV1900Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CellGroupConfig_Ext_NonCollocatedTypeMRDC_v1900_Type1, CellGroupConfig_Ext_NonCollocatedTypeMRDC_v1900_Type4},
}

const (
	CellGroupConfig_Ext_NonCollocatedTypeNR_CA_v1900_Type1 = 0
	CellGroupConfig_Ext_NonCollocatedTypeNR_CA_v1900_Type4 = 1
)

var cellGroupConfigExtNonCollocatedTypeNRCAV1900Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CellGroupConfig_Ext_NonCollocatedTypeNR_CA_v1900_Type1, CellGroupConfig_Ext_NonCollocatedTypeNR_CA_v1900_Type4},
}

const (
	CellGroupConfig_Ext_MprReductionExtensionRatio_r19_Ratio1 = 0
	CellGroupConfig_Ext_MprReductionExtensionRatio_r19_Ratio2 = 1
	CellGroupConfig_Ext_MprReductionExtensionRatio_r19_Ratio3 = 2
	CellGroupConfig_Ext_MprReductionExtensionRatio_r19_Ratio4 = 3
	CellGroupConfig_Ext_MprReductionExtensionRatio_r19_Ratio5 = 4
	CellGroupConfig_Ext_MprReductionExtensionRatio_r19_Ratio6 = 5
	CellGroupConfig_Ext_MprReductionExtensionRatio_r19_Spare2 = 6
	CellGroupConfig_Ext_MprReductionExtensionRatio_r19_Spare1 = 7
)

var cellGroupConfigExtMprReductionExtensionRatioR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CellGroupConfig_Ext_MprReductionExtensionRatio_r19_Ratio1, CellGroupConfig_Ext_MprReductionExtensionRatio_r19_Ratio2, CellGroupConfig_Ext_MprReductionExtensionRatio_r19_Ratio3, CellGroupConfig_Ext_MprReductionExtensionRatio_r19_Ratio4, CellGroupConfig_Ext_MprReductionExtensionRatio_r19_Ratio5, CellGroupConfig_Ext_MprReductionExtensionRatio_r19_Ratio6, CellGroupConfig_Ext_MprReductionExtensionRatio_r19_Spare2, CellGroupConfig_Ext_MprReductionExtensionRatio_r19_Spare1},
}

type CellGroupConfig struct {
	CellGroupId                                CellGroupId
	Rlc_BearerToAddModList                     []RLC_BearerConfig
	Rlc_BearerToReleaseList                    []LogicalChannelIdentity
	Mac_CellGroupConfig                        *MAC_CellGroupConfig
	PhysicalCellGroupConfig                    *PhysicalCellGroupConfig
	SpCellConfig                               *SpCellConfig
	SCellToAddModList                          []SCellConfig
	SCellToReleaseList                         []SCellIndex
	ReportUplinkTxDirectCurrent                *int64
	Bap_Address_r16                            *per.BitString
	Bh_RLC_ChannelToAddModList_r16             []BH_RLC_ChannelConfig_r16
	Bh_RLC_ChannelToReleaseList_r16            []BH_RLC_ChannelID_r16
	F1c_TransferPath_r16                       *int64
	SimultaneousTCI_UpdateList1_r16            []ServCellIndex
	SimultaneousTCI_UpdateList2_r16            []ServCellIndex
	SimultaneousSpatial_UpdatedList1_r16       []ServCellIndex
	SimultaneousSpatial_UpdatedList2_r16       []ServCellIndex
	UplinkTxSwitchingOption_r16                *int64
	UplinkTxSwitchingPowerBoosting_r16         *int64
	ReportUplinkTxDirectCurrentTwoCarrier_r16  *int64
	F1c_TransferPathNRDC_r17                   *int64
	UplinkTxSwitching_2T_Mode_r17              *int64
	UplinkTxSwitching_DualUL_TxState_r17       *int64
	Uu_RelayRLC_ChannelToAddModList_r17        []Uu_RelayRLC_ChannelConfig_r17
	Uu_RelayRLC_ChannelToReleaseList_r17       []Uu_RelayRLC_ChannelID_r17
	SimultaneousU_TCI_UpdateList1_r17          []ServCellIndex
	SimultaneousU_TCI_UpdateList2_r17          []ServCellIndex
	SimultaneousU_TCI_UpdateList3_r17          []ServCellIndex
	SimultaneousU_TCI_UpdateList4_r17          []ServCellIndex
	Rlc_BearerToReleaseListExt_r17             []LogicalChannelIdentityExt_r17
	Iab_ResourceConfigToAddModList_r17         []IAB_ResourceConfig_r17
	Iab_ResourceConfigToReleaseList_r17        []IAB_ResourceConfigID_r17
	ReportUplinkTxDirectCurrentMoreCarrier_r17 *ReportUplinkTxDirectCurrentMoreCarrier_r17
	PrioSCellPRACH_OverSP_PeriodicSRS_r17      *int64
	Ncr_FwdConfig_r18                          *struct {
		Choice  int
		Release *struct{}
		Setup   *NCR_FwdConfig_r18
	}
	AutonomousDenialParameters_r18 *struct {
		Choice  int
		Release *struct{}
		Setup   *AutonomousDenialParameters_r18
	}
	NonCollocatedTypeMRDC_r18      *int64
	NonCollocatedTypeNR_CA_r18     *int64
	UplinkTxSwitchingMoreBands_r18 *struct {
		Choice  int
		Release *struct{}
		Setup   *UplinkTxSwitchingMoreBands_r18
	}
	UplinkTxSwitching3Tx_r19 *int64
	LowBandCA_Switching_r19  *struct {
		Choice  int
		Release *struct{}
		Setup   *LowBandCA_Switching_r19
	}
	NonCollocatedTypeMRDC_v1900    *int64
	NonCollocatedTypeNR_CA_v1900   *int64
	MprReductionExtensionRatio_r19 *int64
}

func (ie *CellGroupConfig) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(cellGroupConfigConstraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.ReportUplinkTxDirectCurrent != nil
	hasExtGroup1 := ie.Bap_Address_r16 != nil || ie.Bh_RLC_ChannelToAddModList_r16 != nil || ie.Bh_RLC_ChannelToReleaseList_r16 != nil || ie.F1c_TransferPath_r16 != nil || ie.SimultaneousTCI_UpdateList1_r16 != nil || ie.SimultaneousTCI_UpdateList2_r16 != nil || ie.SimultaneousSpatial_UpdatedList1_r16 != nil || ie.SimultaneousSpatial_UpdatedList2_r16 != nil || ie.UplinkTxSwitchingOption_r16 != nil || ie.UplinkTxSwitchingPowerBoosting_r16 != nil
	hasExtGroup2 := ie.ReportUplinkTxDirectCurrentTwoCarrier_r16 != nil
	hasExtGroup3 := ie.F1c_TransferPathNRDC_r17 != nil || ie.UplinkTxSwitching_2T_Mode_r17 != nil || ie.UplinkTxSwitching_DualUL_TxState_r17 != nil || ie.Uu_RelayRLC_ChannelToAddModList_r17 != nil || ie.Uu_RelayRLC_ChannelToReleaseList_r17 != nil || ie.SimultaneousU_TCI_UpdateList1_r17 != nil || ie.SimultaneousU_TCI_UpdateList2_r17 != nil || ie.SimultaneousU_TCI_UpdateList3_r17 != nil || ie.SimultaneousU_TCI_UpdateList4_r17 != nil || ie.Rlc_BearerToReleaseListExt_r17 != nil || ie.Iab_ResourceConfigToAddModList_r17 != nil || ie.Iab_ResourceConfigToReleaseList_r17 != nil
	hasExtGroup4 := ie.ReportUplinkTxDirectCurrentMoreCarrier_r17 != nil
	hasExtGroup5 := ie.PrioSCellPRACH_OverSP_PeriodicSRS_r17 != nil
	hasExtGroup6 := ie.Ncr_FwdConfig_r18 != nil || ie.AutonomousDenialParameters_r18 != nil || ie.NonCollocatedTypeMRDC_r18 != nil || ie.NonCollocatedTypeNR_CA_r18 != nil || ie.UplinkTxSwitchingMoreBands_r18 != nil
	hasExtGroup7 := ie.UplinkTxSwitching3Tx_r19 != nil || ie.LowBandCA_Switching_r19 != nil || ie.NonCollocatedTypeMRDC_v1900 != nil || ie.NonCollocatedTypeNR_CA_v1900 != nil || ie.MprReductionExtensionRatio_r19 != nil
	hasExtensions := hasExtGroup0 || hasExtGroup1 || hasExtGroup2 || hasExtGroup3 || hasExtGroup4 || hasExtGroup5 || hasExtGroup6 || hasExtGroup7

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Rlc_BearerToAddModList != nil, ie.Rlc_BearerToReleaseList != nil, ie.Mac_CellGroupConfig != nil, ie.PhysicalCellGroupConfig != nil, ie.SpCellConfig != nil, ie.SCellToAddModList != nil, ie.SCellToReleaseList != nil}); err != nil {
		return err
	}

	// 3. cellGroupId: ref
	{
		if err := ie.CellGroupId.Encode(e); err != nil {
			return err
		}
	}

	// 4. rlc-BearerToAddModList: sequence-of
	{
		if ie.Rlc_BearerToAddModList != nil {
			seqOf := e.NewSequenceOfEncoder(cellGroupConfigRlcBearerToAddModListConstraints)
			if err := seqOf.EncodeLength(int64(len(ie.Rlc_BearerToAddModList))); err != nil {
				return err
			}
			for i := range ie.Rlc_BearerToAddModList {
				if err := ie.Rlc_BearerToAddModList[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 5. rlc-BearerToReleaseList: sequence-of
	{
		if ie.Rlc_BearerToReleaseList != nil {
			seqOf := e.NewSequenceOfEncoder(cellGroupConfigRlcBearerToReleaseListConstraints)
			if err := seqOf.EncodeLength(int64(len(ie.Rlc_BearerToReleaseList))); err != nil {
				return err
			}
			for i := range ie.Rlc_BearerToReleaseList {
				if err := ie.Rlc_BearerToReleaseList[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 6. mac-CellGroupConfig: ref
	{
		if ie.Mac_CellGroupConfig != nil {
			if err := ie.Mac_CellGroupConfig.Encode(e); err != nil {
				return err
			}
		}
	}

	// 7. physicalCellGroupConfig: ref
	{
		if ie.PhysicalCellGroupConfig != nil {
			if err := ie.PhysicalCellGroupConfig.Encode(e); err != nil {
				return err
			}
		}
	}

	// 8. spCellConfig: ref
	{
		if ie.SpCellConfig != nil {
			if err := ie.SpCellConfig.Encode(e); err != nil {
				return err
			}
		}
	}

	// 9. sCellToAddModList: sequence-of
	{
		if ie.SCellToAddModList != nil {
			seqOf := e.NewSequenceOfEncoder(cellGroupConfigSCellToAddModListConstraints)
			if err := seqOf.EncodeLength(int64(len(ie.SCellToAddModList))); err != nil {
				return err
			}
			for i := range ie.SCellToAddModList {
				if err := ie.SCellToAddModList[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 10. sCellToReleaseList: sequence-of
	{
		if ie.SCellToReleaseList != nil {
			seqOf := e.NewSequenceOfEncoder(cellGroupConfigSCellToReleaseListConstraints)
			if err := seqOf.EncodeLength(int64(len(ie.SCellToReleaseList))); err != nil {
				return err
			}
			for i := range ie.SCellToReleaseList {
				if err := ie.SCellToReleaseList[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	if hasExtensions {
		extPresence := []bool{hasExtGroup0, hasExtGroup1, hasExtGroup2, hasExtGroup3, hasExtGroup4, hasExtGroup5, hasExtGroup6, hasExtGroup7}
		var extValues [][]byte

		// Extension group 0:
		if hasExtGroup0 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "reportUplinkTxDirectCurrent", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.ReportUplinkTxDirectCurrent != nil}); err != nil {
				return err
			}

			if ie.ReportUplinkTxDirectCurrent != nil {
				if err := ex.EncodeEnumerated(*ie.ReportUplinkTxDirectCurrent, cellGroupConfigExtReportUplinkTxDirectCurrentConstraints); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 1:
		if hasExtGroup1 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "bap-Address-r16", Optional: true},
					{Name: "bh-RLC-ChannelToAddModList-r16", Optional: true},
					{Name: "bh-RLC-ChannelToReleaseList-r16", Optional: true},
					{Name: "f1c-TransferPath-r16", Optional: true},
					{Name: "simultaneousTCI-UpdateList1-r16", Optional: true},
					{Name: "simultaneousTCI-UpdateList2-r16", Optional: true},
					{Name: "simultaneousSpatial-UpdatedList1-r16", Optional: true},
					{Name: "simultaneousSpatial-UpdatedList2-r16", Optional: true},
					{Name: "uplinkTxSwitchingOption-r16", Optional: true},
					{Name: "uplinkTxSwitchingPowerBoosting-r16", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Bap_Address_r16 != nil, ie.Bh_RLC_ChannelToAddModList_r16 != nil, ie.Bh_RLC_ChannelToReleaseList_r16 != nil, ie.F1c_TransferPath_r16 != nil, ie.SimultaneousTCI_UpdateList1_r16 != nil, ie.SimultaneousTCI_UpdateList2_r16 != nil, ie.SimultaneousSpatial_UpdatedList1_r16 != nil, ie.SimultaneousSpatial_UpdatedList2_r16 != nil, ie.UplinkTxSwitchingOption_r16 != nil, ie.UplinkTxSwitchingPowerBoosting_r16 != nil}); err != nil {
				return err
			}

			if ie.Bap_Address_r16 != nil {
				if err := ex.EncodeBitString(*ie.Bap_Address_r16, cellGroupConfigBapAddressR16Constraints); err != nil {
					return err
				}
			}

			if ie.Bh_RLC_ChannelToAddModList_r16 != nil {
				seqOf := ex.NewSequenceOfEncoder(cellGroupConfigExtBhRLCChannelToAddModListR16Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.Bh_RLC_ChannelToAddModList_r16))); err != nil {
					return err
				}
				for i := range ie.Bh_RLC_ChannelToAddModList_r16 {
					if err := ie.Bh_RLC_ChannelToAddModList_r16[i].Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.Bh_RLC_ChannelToReleaseList_r16 != nil {
				seqOf := ex.NewSequenceOfEncoder(cellGroupConfigExtBhRLCChannelToReleaseListR16Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.Bh_RLC_ChannelToReleaseList_r16))); err != nil {
					return err
				}
				for i := range ie.Bh_RLC_ChannelToReleaseList_r16 {
					if err := ie.Bh_RLC_ChannelToReleaseList_r16[i].Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.F1c_TransferPath_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.F1c_TransferPath_r16, cellGroupConfigExtF1cTransferPathR16Constraints); err != nil {
					return err
				}
			}

			if ie.SimultaneousTCI_UpdateList1_r16 != nil {
				seqOf := ex.NewSequenceOfEncoder(cellGroupConfigExtSimultaneousTCIUpdateList1R16Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.SimultaneousTCI_UpdateList1_r16))); err != nil {
					return err
				}
				for i := range ie.SimultaneousTCI_UpdateList1_r16 {
					if err := ie.SimultaneousTCI_UpdateList1_r16[i].Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.SimultaneousTCI_UpdateList2_r16 != nil {
				seqOf := ex.NewSequenceOfEncoder(cellGroupConfigExtSimultaneousTCIUpdateList2R16Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.SimultaneousTCI_UpdateList2_r16))); err != nil {
					return err
				}
				for i := range ie.SimultaneousTCI_UpdateList2_r16 {
					if err := ie.SimultaneousTCI_UpdateList2_r16[i].Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.SimultaneousSpatial_UpdatedList1_r16 != nil {
				seqOf := ex.NewSequenceOfEncoder(cellGroupConfigExtSimultaneousSpatialUpdatedList1R16Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.SimultaneousSpatial_UpdatedList1_r16))); err != nil {
					return err
				}
				for i := range ie.SimultaneousSpatial_UpdatedList1_r16 {
					if err := ie.SimultaneousSpatial_UpdatedList1_r16[i].Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.SimultaneousSpatial_UpdatedList2_r16 != nil {
				seqOf := ex.NewSequenceOfEncoder(cellGroupConfigExtSimultaneousSpatialUpdatedList2R16Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.SimultaneousSpatial_UpdatedList2_r16))); err != nil {
					return err
				}
				for i := range ie.SimultaneousSpatial_UpdatedList2_r16 {
					if err := ie.SimultaneousSpatial_UpdatedList2_r16[i].Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.UplinkTxSwitchingOption_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.UplinkTxSwitchingOption_r16, cellGroupConfigExtUplinkTxSwitchingOptionR16Constraints); err != nil {
					return err
				}
			}

			if ie.UplinkTxSwitchingPowerBoosting_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.UplinkTxSwitchingPowerBoosting_r16, cellGroupConfigExtUplinkTxSwitchingPowerBoostingR16Constraints); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 2:
		if hasExtGroup2 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "reportUplinkTxDirectCurrentTwoCarrier-r16", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.ReportUplinkTxDirectCurrentTwoCarrier_r16 != nil}); err != nil {
				return err
			}

			if ie.ReportUplinkTxDirectCurrentTwoCarrier_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.ReportUplinkTxDirectCurrentTwoCarrier_r16, cellGroupConfigExtReportUplinkTxDirectCurrentTwoCarrierR16Constraints); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 3:
		if hasExtGroup3 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "f1c-TransferPathNRDC-r17", Optional: true},
					{Name: "uplinkTxSwitching-2T-Mode-r17", Optional: true},
					{Name: "uplinkTxSwitching-DualUL-TxState-r17", Optional: true},
					{Name: "uu-RelayRLC-ChannelToAddModList-r17", Optional: true},
					{Name: "uu-RelayRLC-ChannelToReleaseList-r17", Optional: true},
					{Name: "simultaneousU-TCI-UpdateList1-r17", Optional: true},
					{Name: "simultaneousU-TCI-UpdateList2-r17", Optional: true},
					{Name: "simultaneousU-TCI-UpdateList3-r17", Optional: true},
					{Name: "simultaneousU-TCI-UpdateList4-r17", Optional: true},
					{Name: "rlc-BearerToReleaseListExt-r17", Optional: true},
					{Name: "iab-ResourceConfigToAddModList-r17", Optional: true},
					{Name: "iab-ResourceConfigToReleaseList-r17", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.F1c_TransferPathNRDC_r17 != nil, ie.UplinkTxSwitching_2T_Mode_r17 != nil, ie.UplinkTxSwitching_DualUL_TxState_r17 != nil, ie.Uu_RelayRLC_ChannelToAddModList_r17 != nil, ie.Uu_RelayRLC_ChannelToReleaseList_r17 != nil, ie.SimultaneousU_TCI_UpdateList1_r17 != nil, ie.SimultaneousU_TCI_UpdateList2_r17 != nil, ie.SimultaneousU_TCI_UpdateList3_r17 != nil, ie.SimultaneousU_TCI_UpdateList4_r17 != nil, ie.Rlc_BearerToReleaseListExt_r17 != nil, ie.Iab_ResourceConfigToAddModList_r17 != nil, ie.Iab_ResourceConfigToReleaseList_r17 != nil}); err != nil {
				return err
			}

			if ie.F1c_TransferPathNRDC_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.F1c_TransferPathNRDC_r17, cellGroupConfigExtF1cTransferPathNRDCR17Constraints); err != nil {
					return err
				}
			}

			if ie.UplinkTxSwitching_2T_Mode_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.UplinkTxSwitching_2T_Mode_r17, cellGroupConfigExtUplinkTxSwitching2TModeR17Constraints); err != nil {
					return err
				}
			}

			if ie.UplinkTxSwitching_DualUL_TxState_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.UplinkTxSwitching_DualUL_TxState_r17, cellGroupConfigExtUplinkTxSwitchingDualULTxStateR17Constraints); err != nil {
					return err
				}
			}

			if ie.Uu_RelayRLC_ChannelToAddModList_r17 != nil {
				seqOf := ex.NewSequenceOfEncoder(cellGroupConfigExtUuRelayRLCChannelToAddModListR17Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.Uu_RelayRLC_ChannelToAddModList_r17))); err != nil {
					return err
				}
				for i := range ie.Uu_RelayRLC_ChannelToAddModList_r17 {
					if err := ie.Uu_RelayRLC_ChannelToAddModList_r17[i].Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.Uu_RelayRLC_ChannelToReleaseList_r17 != nil {
				seqOf := ex.NewSequenceOfEncoder(cellGroupConfigExtUuRelayRLCChannelToReleaseListR17Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.Uu_RelayRLC_ChannelToReleaseList_r17))); err != nil {
					return err
				}
				for i := range ie.Uu_RelayRLC_ChannelToReleaseList_r17 {
					if err := ie.Uu_RelayRLC_ChannelToReleaseList_r17[i].Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.SimultaneousU_TCI_UpdateList1_r17 != nil {
				seqOf := ex.NewSequenceOfEncoder(cellGroupConfigExtSimultaneousUTCIUpdateList1R17Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.SimultaneousU_TCI_UpdateList1_r17))); err != nil {
					return err
				}
				for i := range ie.SimultaneousU_TCI_UpdateList1_r17 {
					if err := ie.SimultaneousU_TCI_UpdateList1_r17[i].Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.SimultaneousU_TCI_UpdateList2_r17 != nil {
				seqOf := ex.NewSequenceOfEncoder(cellGroupConfigExtSimultaneousUTCIUpdateList2R17Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.SimultaneousU_TCI_UpdateList2_r17))); err != nil {
					return err
				}
				for i := range ie.SimultaneousU_TCI_UpdateList2_r17 {
					if err := ie.SimultaneousU_TCI_UpdateList2_r17[i].Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.SimultaneousU_TCI_UpdateList3_r17 != nil {
				seqOf := ex.NewSequenceOfEncoder(cellGroupConfigExtSimultaneousUTCIUpdateList3R17Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.SimultaneousU_TCI_UpdateList3_r17))); err != nil {
					return err
				}
				for i := range ie.SimultaneousU_TCI_UpdateList3_r17 {
					if err := ie.SimultaneousU_TCI_UpdateList3_r17[i].Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.SimultaneousU_TCI_UpdateList4_r17 != nil {
				seqOf := ex.NewSequenceOfEncoder(cellGroupConfigExtSimultaneousUTCIUpdateList4R17Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.SimultaneousU_TCI_UpdateList4_r17))); err != nil {
					return err
				}
				for i := range ie.SimultaneousU_TCI_UpdateList4_r17 {
					if err := ie.SimultaneousU_TCI_UpdateList4_r17[i].Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.Rlc_BearerToReleaseListExt_r17 != nil {
				seqOf := ex.NewSequenceOfEncoder(cellGroupConfigExtRlcBearerToReleaseListExtR17Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.Rlc_BearerToReleaseListExt_r17))); err != nil {
					return err
				}
				for i := range ie.Rlc_BearerToReleaseListExt_r17 {
					if err := ie.Rlc_BearerToReleaseListExt_r17[i].Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.Iab_ResourceConfigToAddModList_r17 != nil {
				seqOf := ex.NewSequenceOfEncoder(cellGroupConfigExtIabResourceConfigToAddModListR17Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.Iab_ResourceConfigToAddModList_r17))); err != nil {
					return err
				}
				for i := range ie.Iab_ResourceConfigToAddModList_r17 {
					if err := ie.Iab_ResourceConfigToAddModList_r17[i].Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.Iab_ResourceConfigToReleaseList_r17 != nil {
				seqOf := ex.NewSequenceOfEncoder(cellGroupConfigExtIabResourceConfigToReleaseListR17Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.Iab_ResourceConfigToReleaseList_r17))); err != nil {
					return err
				}
				for i := range ie.Iab_ResourceConfigToReleaseList_r17 {
					if err := ie.Iab_ResourceConfigToReleaseList_r17[i].Encode(ex); err != nil {
						return err
					}
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 4:
		if hasExtGroup4 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "reportUplinkTxDirectCurrentMoreCarrier-r17", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.ReportUplinkTxDirectCurrentMoreCarrier_r17 != nil}); err != nil {
				return err
			}

			if ie.ReportUplinkTxDirectCurrentMoreCarrier_r17 != nil {
				if err := ie.ReportUplinkTxDirectCurrentMoreCarrier_r17.Encode(ex); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 5:
		if hasExtGroup5 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "prioSCellPRACH-OverSP-PeriodicSRS-r17", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.PrioSCellPRACH_OverSP_PeriodicSRS_r17 != nil}); err != nil {
				return err
			}

			if ie.PrioSCellPRACH_OverSP_PeriodicSRS_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.PrioSCellPRACH_OverSP_PeriodicSRS_r17, cellGroupConfigExtPrioSCellPRACHOverSPPeriodicSRSR17Constraints); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 6:
		if hasExtGroup6 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "ncr-FwdConfig-r18", Optional: true},
					{Name: "autonomousDenialParameters-r18", Optional: true},
					{Name: "nonCollocatedTypeMRDC-r18", Optional: true},
					{Name: "nonCollocatedTypeNR-CA-r18", Optional: true},
					{Name: "uplinkTxSwitchingMoreBands-r18", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Ncr_FwdConfig_r18 != nil, ie.AutonomousDenialParameters_r18 != nil, ie.NonCollocatedTypeMRDC_r18 != nil, ie.NonCollocatedTypeNR_CA_r18 != nil, ie.UplinkTxSwitchingMoreBands_r18 != nil}); err != nil {
				return err
			}

			if ie.Ncr_FwdConfig_r18 != nil {
				choiceEnc := ex.NewChoiceEncoder(cellGroupConfigExtNcrFwdConfigR18Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.Ncr_FwdConfig_r18).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.Ncr_FwdConfig_r18).Choice {
				case CellGroupConfig_Ext_Ncr_FwdConfig_r18_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case CellGroupConfig_Ext_Ncr_FwdConfig_r18_Setup:
					if err := (*ie.Ncr_FwdConfig_r18).Setup.Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.AutonomousDenialParameters_r18 != nil {
				choiceEnc := ex.NewChoiceEncoder(cellGroupConfigExtAutonomousDenialParametersR18Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.AutonomousDenialParameters_r18).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.AutonomousDenialParameters_r18).Choice {
				case CellGroupConfig_Ext_AutonomousDenialParameters_r18_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case CellGroupConfig_Ext_AutonomousDenialParameters_r18_Setup:
					if err := (*ie.AutonomousDenialParameters_r18).Setup.Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.NonCollocatedTypeMRDC_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.NonCollocatedTypeMRDC_r18, cellGroupConfigExtNonCollocatedTypeMRDCR18Constraints); err != nil {
					return err
				}
			}

			if ie.NonCollocatedTypeNR_CA_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.NonCollocatedTypeNR_CA_r18, cellGroupConfigExtNonCollocatedTypeNRCAR18Constraints); err != nil {
					return err
				}
			}

			if ie.UplinkTxSwitchingMoreBands_r18 != nil {
				choiceEnc := ex.NewChoiceEncoder(cellGroupConfigExtUplinkTxSwitchingMoreBandsR18Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.UplinkTxSwitchingMoreBands_r18).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.UplinkTxSwitchingMoreBands_r18).Choice {
				case CellGroupConfig_Ext_UplinkTxSwitchingMoreBands_r18_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case CellGroupConfig_Ext_UplinkTxSwitchingMoreBands_r18_Setup:
					if err := (*ie.UplinkTxSwitchingMoreBands_r18).Setup.Encode(ex); err != nil {
						return err
					}
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 7:
		if hasExtGroup7 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "uplinkTxSwitching3Tx-r19", Optional: true},
					{Name: "lowBandCA-Switching-r19", Optional: true},
					{Name: "nonCollocatedTypeMRDC-v1900", Optional: true},
					{Name: "nonCollocatedTypeNR-CA-v1900", Optional: true},
					{Name: "mprReductionExtensionRatio-r19", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.UplinkTxSwitching3Tx_r19 != nil, ie.LowBandCA_Switching_r19 != nil, ie.NonCollocatedTypeMRDC_v1900 != nil, ie.NonCollocatedTypeNR_CA_v1900 != nil, ie.MprReductionExtensionRatio_r19 != nil}); err != nil {
				return err
			}

			if ie.UplinkTxSwitching3Tx_r19 != nil {
				if err := ex.EncodeEnumerated(*ie.UplinkTxSwitching3Tx_r19, cellGroupConfigExtUplinkTxSwitching3TxR19Constraints); err != nil {
					return err
				}
			}

			if ie.LowBandCA_Switching_r19 != nil {
				choiceEnc := ex.NewChoiceEncoder(cellGroupConfigExtLowBandCASwitchingR19Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.LowBandCA_Switching_r19).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.LowBandCA_Switching_r19).Choice {
				case CellGroupConfig_Ext_LowBandCA_Switching_r19_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case CellGroupConfig_Ext_LowBandCA_Switching_r19_Setup:
					if err := (*ie.LowBandCA_Switching_r19).Setup.Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.NonCollocatedTypeMRDC_v1900 != nil {
				if err := ex.EncodeEnumerated(*ie.NonCollocatedTypeMRDC_v1900, cellGroupConfigExtNonCollocatedTypeMRDCV1900Constraints); err != nil {
					return err
				}
			}

			if ie.NonCollocatedTypeNR_CA_v1900 != nil {
				if err := ex.EncodeEnumerated(*ie.NonCollocatedTypeNR_CA_v1900, cellGroupConfigExtNonCollocatedTypeNRCAV1900Constraints); err != nil {
					return err
				}
			}

			if ie.MprReductionExtensionRatio_r19 != nil {
				if err := ex.EncodeEnumerated(*ie.MprReductionExtensionRatio_r19, cellGroupConfigExtMprReductionExtensionRatioR19Constraints); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		if err := seq.EncodeExtensionAdditions(extPresence, extValues); err != nil {
			return err
		}
	}

	return nil
}

func (ie *CellGroupConfig) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(cellGroupConfigConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. cellGroupId: ref
	{
		if err := ie.CellGroupId.Decode(d); err != nil {
			return err
		}
	}

	// 4. rlc-BearerToAddModList: sequence-of
	{
		if seq.IsComponentPresent(1) {
			seqOf := d.NewSequenceOfDecoder(cellGroupConfigRlcBearerToAddModListConstraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Rlc_BearerToAddModList = make([]RLC_BearerConfig, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Rlc_BearerToAddModList[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 5. rlc-BearerToReleaseList: sequence-of
	{
		if seq.IsComponentPresent(2) {
			seqOf := d.NewSequenceOfDecoder(cellGroupConfigRlcBearerToReleaseListConstraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Rlc_BearerToReleaseList = make([]LogicalChannelIdentity, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Rlc_BearerToReleaseList[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 6. mac-CellGroupConfig: ref
	{
		if seq.IsComponentPresent(3) {
			ie.Mac_CellGroupConfig = new(MAC_CellGroupConfig)
			if err := ie.Mac_CellGroupConfig.Decode(d); err != nil {
				return err
			}
		}
	}

	// 7. physicalCellGroupConfig: ref
	{
		if seq.IsComponentPresent(4) {
			ie.PhysicalCellGroupConfig = new(PhysicalCellGroupConfig)
			if err := ie.PhysicalCellGroupConfig.Decode(d); err != nil {
				return err
			}
		}
	}

	// 8. spCellConfig: ref
	{
		if seq.IsComponentPresent(5) {
			ie.SpCellConfig = new(SpCellConfig)
			if err := ie.SpCellConfig.Decode(d); err != nil {
				return err
			}
		}
	}

	// 9. sCellToAddModList: sequence-of
	{
		if seq.IsComponentPresent(6) {
			seqOf := d.NewSequenceOfDecoder(cellGroupConfigSCellToAddModListConstraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.SCellToAddModList = make([]SCellConfig, n)
			for i := int64(0); i < n; i++ {
				if err := ie.SCellToAddModList[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 10. sCellToReleaseList: sequence-of
	{
		if seq.IsComponentPresent(7) {
			seqOf := d.NewSequenceOfDecoder(cellGroupConfigSCellToReleaseListConstraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.SCellToReleaseList = make([]SCellIndex, n)
			for i := int64(0); i < n; i++ {
				if err := ie.SCellToReleaseList[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// Extension additions as open types.
	extValues, err := seq.DecodeExtensionAdditions()
	if err != nil {
		return err
	}
	extValueIndex := 0

	// Extension group 0:
	if seq.IsExtensionPresent(0) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "reportUplinkTxDirectCurrent", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(cellGroupConfigExtReportUplinkTxDirectCurrentConstraints)
			if err != nil {
				return err
			}
			ie.ReportUplinkTxDirectCurrent = &v
		}
	}

	// Extension group 1:
	if seq.IsExtensionPresent(1) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "bap-Address-r16", Optional: true},
				{Name: "bh-RLC-ChannelToAddModList-r16", Optional: true},
				{Name: "bh-RLC-ChannelToReleaseList-r16", Optional: true},
				{Name: "f1c-TransferPath-r16", Optional: true},
				{Name: "simultaneousTCI-UpdateList1-r16", Optional: true},
				{Name: "simultaneousTCI-UpdateList2-r16", Optional: true},
				{Name: "simultaneousSpatial-UpdatedList1-r16", Optional: true},
				{Name: "simultaneousSpatial-UpdatedList2-r16", Optional: true},
				{Name: "uplinkTxSwitchingOption-r16", Optional: true},
				{Name: "uplinkTxSwitchingPowerBoosting-r16", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeBitString(cellGroupConfigBapAddressR16Constraints)
			if err != nil {
				return err
			}
			ie.Bap_Address_r16 = &v
		}

		if groupSeq.IsComponentPresent(1) {
			seqOf := dx.NewSequenceOfDecoder(cellGroupConfigExtBhRLCChannelToAddModListR16Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Bh_RLC_ChannelToAddModList_r16 = make([]BH_RLC_ChannelConfig_r16, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Bh_RLC_ChannelToAddModList_r16[i].Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(2) {
			seqOf := dx.NewSequenceOfDecoder(cellGroupConfigExtBhRLCChannelToReleaseListR16Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Bh_RLC_ChannelToReleaseList_r16 = make([]BH_RLC_ChannelID_r16, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Bh_RLC_ChannelToReleaseList_r16[i].Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(3) {
			v, err := dx.DecodeEnumerated(cellGroupConfigExtF1cTransferPathR16Constraints)
			if err != nil {
				return err
			}
			ie.F1c_TransferPath_r16 = &v
		}

		if groupSeq.IsComponentPresent(4) {
			seqOf := dx.NewSequenceOfDecoder(cellGroupConfigExtSimultaneousTCIUpdateList1R16Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.SimultaneousTCI_UpdateList1_r16 = make([]ServCellIndex, n)
			for i := int64(0); i < n; i++ {
				if err := ie.SimultaneousTCI_UpdateList1_r16[i].Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(5) {
			seqOf := dx.NewSequenceOfDecoder(cellGroupConfigExtSimultaneousTCIUpdateList2R16Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.SimultaneousTCI_UpdateList2_r16 = make([]ServCellIndex, n)
			for i := int64(0); i < n; i++ {
				if err := ie.SimultaneousTCI_UpdateList2_r16[i].Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(6) {
			seqOf := dx.NewSequenceOfDecoder(cellGroupConfigExtSimultaneousSpatialUpdatedList1R16Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.SimultaneousSpatial_UpdatedList1_r16 = make([]ServCellIndex, n)
			for i := int64(0); i < n; i++ {
				if err := ie.SimultaneousSpatial_UpdatedList1_r16[i].Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(7) {
			seqOf := dx.NewSequenceOfDecoder(cellGroupConfigExtSimultaneousSpatialUpdatedList2R16Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.SimultaneousSpatial_UpdatedList2_r16 = make([]ServCellIndex, n)
			for i := int64(0); i < n; i++ {
				if err := ie.SimultaneousSpatial_UpdatedList2_r16[i].Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(8) {
			v, err := dx.DecodeEnumerated(cellGroupConfigExtUplinkTxSwitchingOptionR16Constraints)
			if err != nil {
				return err
			}
			ie.UplinkTxSwitchingOption_r16 = &v
		}

		if groupSeq.IsComponentPresent(9) {
			v, err := dx.DecodeEnumerated(cellGroupConfigExtUplinkTxSwitchingPowerBoostingR16Constraints)
			if err != nil {
				return err
			}
			ie.UplinkTxSwitchingPowerBoosting_r16 = &v
		}
	}

	// Extension group 2:
	if seq.IsExtensionPresent(2) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "reportUplinkTxDirectCurrentTwoCarrier-r16", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(cellGroupConfigExtReportUplinkTxDirectCurrentTwoCarrierR16Constraints)
			if err != nil {
				return err
			}
			ie.ReportUplinkTxDirectCurrentTwoCarrier_r16 = &v
		}
	}

	// Extension group 3:
	if seq.IsExtensionPresent(3) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "f1c-TransferPathNRDC-r17", Optional: true},
				{Name: "uplinkTxSwitching-2T-Mode-r17", Optional: true},
				{Name: "uplinkTxSwitching-DualUL-TxState-r17", Optional: true},
				{Name: "uu-RelayRLC-ChannelToAddModList-r17", Optional: true},
				{Name: "uu-RelayRLC-ChannelToReleaseList-r17", Optional: true},
				{Name: "simultaneousU-TCI-UpdateList1-r17", Optional: true},
				{Name: "simultaneousU-TCI-UpdateList2-r17", Optional: true},
				{Name: "simultaneousU-TCI-UpdateList3-r17", Optional: true},
				{Name: "simultaneousU-TCI-UpdateList4-r17", Optional: true},
				{Name: "rlc-BearerToReleaseListExt-r17", Optional: true},
				{Name: "iab-ResourceConfigToAddModList-r17", Optional: true},
				{Name: "iab-ResourceConfigToReleaseList-r17", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(cellGroupConfigExtF1cTransferPathNRDCR17Constraints)
			if err != nil {
				return err
			}
			ie.F1c_TransferPathNRDC_r17 = &v
		}

		if groupSeq.IsComponentPresent(1) {
			v, err := dx.DecodeEnumerated(cellGroupConfigExtUplinkTxSwitching2TModeR17Constraints)
			if err != nil {
				return err
			}
			ie.UplinkTxSwitching_2T_Mode_r17 = &v
		}

		if groupSeq.IsComponentPresent(2) {
			v, err := dx.DecodeEnumerated(cellGroupConfigExtUplinkTxSwitchingDualULTxStateR17Constraints)
			if err != nil {
				return err
			}
			ie.UplinkTxSwitching_DualUL_TxState_r17 = &v
		}

		if groupSeq.IsComponentPresent(3) {
			seqOf := dx.NewSequenceOfDecoder(cellGroupConfigExtUuRelayRLCChannelToAddModListR17Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Uu_RelayRLC_ChannelToAddModList_r17 = make([]Uu_RelayRLC_ChannelConfig_r17, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Uu_RelayRLC_ChannelToAddModList_r17[i].Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(4) {
			seqOf := dx.NewSequenceOfDecoder(cellGroupConfigExtUuRelayRLCChannelToReleaseListR17Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Uu_RelayRLC_ChannelToReleaseList_r17 = make([]Uu_RelayRLC_ChannelID_r17, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Uu_RelayRLC_ChannelToReleaseList_r17[i].Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(5) {
			seqOf := dx.NewSequenceOfDecoder(cellGroupConfigExtSimultaneousUTCIUpdateList1R17Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.SimultaneousU_TCI_UpdateList1_r17 = make([]ServCellIndex, n)
			for i := int64(0); i < n; i++ {
				if err := ie.SimultaneousU_TCI_UpdateList1_r17[i].Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(6) {
			seqOf := dx.NewSequenceOfDecoder(cellGroupConfigExtSimultaneousUTCIUpdateList2R17Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.SimultaneousU_TCI_UpdateList2_r17 = make([]ServCellIndex, n)
			for i := int64(0); i < n; i++ {
				if err := ie.SimultaneousU_TCI_UpdateList2_r17[i].Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(7) {
			seqOf := dx.NewSequenceOfDecoder(cellGroupConfigExtSimultaneousUTCIUpdateList3R17Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.SimultaneousU_TCI_UpdateList3_r17 = make([]ServCellIndex, n)
			for i := int64(0); i < n; i++ {
				if err := ie.SimultaneousU_TCI_UpdateList3_r17[i].Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(8) {
			seqOf := dx.NewSequenceOfDecoder(cellGroupConfigExtSimultaneousUTCIUpdateList4R17Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.SimultaneousU_TCI_UpdateList4_r17 = make([]ServCellIndex, n)
			for i := int64(0); i < n; i++ {
				if err := ie.SimultaneousU_TCI_UpdateList4_r17[i].Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(9) {
			seqOf := dx.NewSequenceOfDecoder(cellGroupConfigExtRlcBearerToReleaseListExtR17Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Rlc_BearerToReleaseListExt_r17 = make([]LogicalChannelIdentityExt_r17, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Rlc_BearerToReleaseListExt_r17[i].Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(10) {
			seqOf := dx.NewSequenceOfDecoder(cellGroupConfigExtIabResourceConfigToAddModListR17Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Iab_ResourceConfigToAddModList_r17 = make([]IAB_ResourceConfig_r17, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Iab_ResourceConfigToAddModList_r17[i].Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(11) {
			seqOf := dx.NewSequenceOfDecoder(cellGroupConfigExtIabResourceConfigToReleaseListR17Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Iab_ResourceConfigToReleaseList_r17 = make([]IAB_ResourceConfigID_r17, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Iab_ResourceConfigToReleaseList_r17[i].Decode(dx); err != nil {
					return err
				}
			}
		}
	}

	// Extension group 4:
	if seq.IsExtensionPresent(4) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "reportUplinkTxDirectCurrentMoreCarrier-r17", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.ReportUplinkTxDirectCurrentMoreCarrier_r17 = new(ReportUplinkTxDirectCurrentMoreCarrier_r17)
			if err := ie.ReportUplinkTxDirectCurrentMoreCarrier_r17.Decode(dx); err != nil {
				return err
			}
		}
	}

	// Extension group 5:
	if seq.IsExtensionPresent(5) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "prioSCellPRACH-OverSP-PeriodicSRS-r17", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(cellGroupConfigExtPrioSCellPRACHOverSPPeriodicSRSR17Constraints)
			if err != nil {
				return err
			}
			ie.PrioSCellPRACH_OverSP_PeriodicSRS_r17 = &v
		}
	}

	// Extension group 6:
	if seq.IsExtensionPresent(6) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "ncr-FwdConfig-r18", Optional: true},
				{Name: "autonomousDenialParameters-r18", Optional: true},
				{Name: "nonCollocatedTypeMRDC-r18", Optional: true},
				{Name: "nonCollocatedTypeNR-CA-r18", Optional: true},
				{Name: "uplinkTxSwitchingMoreBands-r18", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.Ncr_FwdConfig_r18 = &struct {
				Choice  int
				Release *struct{}
				Setup   *NCR_FwdConfig_r18
			}{}
			choiceDec := dx.NewChoiceDecoder(cellGroupConfigExtNcrFwdConfigR18Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Ncr_FwdConfig_r18).Choice = int(index)
			switch index {
			case CellGroupConfig_Ext_Ncr_FwdConfig_r18_Release:
				(*ie.Ncr_FwdConfig_r18).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case CellGroupConfig_Ext_Ncr_FwdConfig_r18_Setup:
				(*ie.Ncr_FwdConfig_r18).Setup = new(NCR_FwdConfig_r18)
				if err := (*ie.Ncr_FwdConfig_r18).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(1) {
			ie.AutonomousDenialParameters_r18 = &struct {
				Choice  int
				Release *struct{}
				Setup   *AutonomousDenialParameters_r18
			}{}
			choiceDec := dx.NewChoiceDecoder(cellGroupConfigExtAutonomousDenialParametersR18Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.AutonomousDenialParameters_r18).Choice = int(index)
			switch index {
			case CellGroupConfig_Ext_AutonomousDenialParameters_r18_Release:
				(*ie.AutonomousDenialParameters_r18).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case CellGroupConfig_Ext_AutonomousDenialParameters_r18_Setup:
				(*ie.AutonomousDenialParameters_r18).Setup = new(AutonomousDenialParameters_r18)
				if err := (*ie.AutonomousDenialParameters_r18).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(2) {
			v, err := dx.DecodeEnumerated(cellGroupConfigExtNonCollocatedTypeMRDCR18Constraints)
			if err != nil {
				return err
			}
			ie.NonCollocatedTypeMRDC_r18 = &v
		}

		if groupSeq.IsComponentPresent(3) {
			v, err := dx.DecodeEnumerated(cellGroupConfigExtNonCollocatedTypeNRCAR18Constraints)
			if err != nil {
				return err
			}
			ie.NonCollocatedTypeNR_CA_r18 = &v
		}

		if groupSeq.IsComponentPresent(4) {
			ie.UplinkTxSwitchingMoreBands_r18 = &struct {
				Choice  int
				Release *struct{}
				Setup   *UplinkTxSwitchingMoreBands_r18
			}{}
			choiceDec := dx.NewChoiceDecoder(cellGroupConfigExtUplinkTxSwitchingMoreBandsR18Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.UplinkTxSwitchingMoreBands_r18).Choice = int(index)
			switch index {
			case CellGroupConfig_Ext_UplinkTxSwitchingMoreBands_r18_Release:
				(*ie.UplinkTxSwitchingMoreBands_r18).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case CellGroupConfig_Ext_UplinkTxSwitchingMoreBands_r18_Setup:
				(*ie.UplinkTxSwitchingMoreBands_r18).Setup = new(UplinkTxSwitchingMoreBands_r18)
				if err := (*ie.UplinkTxSwitchingMoreBands_r18).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}
	}

	// Extension group 7:
	if seq.IsExtensionPresent(7) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "uplinkTxSwitching3Tx-r19", Optional: true},
				{Name: "lowBandCA-Switching-r19", Optional: true},
				{Name: "nonCollocatedTypeMRDC-v1900", Optional: true},
				{Name: "nonCollocatedTypeNR-CA-v1900", Optional: true},
				{Name: "mprReductionExtensionRatio-r19", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(cellGroupConfigExtUplinkTxSwitching3TxR19Constraints)
			if err != nil {
				return err
			}
			ie.UplinkTxSwitching3Tx_r19 = &v
		}

		if groupSeq.IsComponentPresent(1) {
			ie.LowBandCA_Switching_r19 = &struct {
				Choice  int
				Release *struct{}
				Setup   *LowBandCA_Switching_r19
			}{}
			choiceDec := dx.NewChoiceDecoder(cellGroupConfigExtLowBandCASwitchingR19Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.LowBandCA_Switching_r19).Choice = int(index)
			switch index {
			case CellGroupConfig_Ext_LowBandCA_Switching_r19_Release:
				(*ie.LowBandCA_Switching_r19).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case CellGroupConfig_Ext_LowBandCA_Switching_r19_Setup:
				(*ie.LowBandCA_Switching_r19).Setup = new(LowBandCA_Switching_r19)
				if err := (*ie.LowBandCA_Switching_r19).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(2) {
			v, err := dx.DecodeEnumerated(cellGroupConfigExtNonCollocatedTypeMRDCV1900Constraints)
			if err != nil {
				return err
			}
			ie.NonCollocatedTypeMRDC_v1900 = &v
		}

		if groupSeq.IsComponentPresent(3) {
			v, err := dx.DecodeEnumerated(cellGroupConfigExtNonCollocatedTypeNRCAV1900Constraints)
			if err != nil {
				return err
			}
			ie.NonCollocatedTypeNR_CA_v1900 = &v
		}

		if groupSeq.IsComponentPresent(4) {
			v, err := dx.DecodeEnumerated(cellGroupConfigExtMprReductionExtensionRatioR19Constraints)
			if err != nil {
				return err
			}
			ie.MprReductionExtensionRatio_r19 = &v
		}
	}

	return nil
}
