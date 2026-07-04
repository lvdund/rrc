// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: PDCP-Config (line 11174).

var pDCPConfigConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "drb", Optional: true},
		{Name: "moreThanOneRLC", Optional: true},
		{Name: "t-Reordering", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
		{Name: "extension-addition-1", Optional: true},
		{Name: "extension-addition-2", Optional: true},
		{Name: "extension-addition-3", Optional: true},
		{Name: "extension-addition-4", Optional: true},
	},
}

const (
	PDCP_Config_T_Reordering_Ms0     = 0
	PDCP_Config_T_Reordering_Ms1     = 1
	PDCP_Config_T_Reordering_Ms2     = 2
	PDCP_Config_T_Reordering_Ms4     = 3
	PDCP_Config_T_Reordering_Ms5     = 4
	PDCP_Config_T_Reordering_Ms8     = 5
	PDCP_Config_T_Reordering_Ms10    = 6
	PDCP_Config_T_Reordering_Ms15    = 7
	PDCP_Config_T_Reordering_Ms20    = 8
	PDCP_Config_T_Reordering_Ms30    = 9
	PDCP_Config_T_Reordering_Ms40    = 10
	PDCP_Config_T_Reordering_Ms50    = 11
	PDCP_Config_T_Reordering_Ms60    = 12
	PDCP_Config_T_Reordering_Ms80    = 13
	PDCP_Config_T_Reordering_Ms100   = 14
	PDCP_Config_T_Reordering_Ms120   = 15
	PDCP_Config_T_Reordering_Ms140   = 16
	PDCP_Config_T_Reordering_Ms160   = 17
	PDCP_Config_T_Reordering_Ms180   = 18
	PDCP_Config_T_Reordering_Ms200   = 19
	PDCP_Config_T_Reordering_Ms220   = 20
	PDCP_Config_T_Reordering_Ms240   = 21
	PDCP_Config_T_Reordering_Ms260   = 22
	PDCP_Config_T_Reordering_Ms280   = 23
	PDCP_Config_T_Reordering_Ms300   = 24
	PDCP_Config_T_Reordering_Ms500   = 25
	PDCP_Config_T_Reordering_Ms750   = 26
	PDCP_Config_T_Reordering_Ms1000  = 27
	PDCP_Config_T_Reordering_Ms1250  = 28
	PDCP_Config_T_Reordering_Ms1500  = 29
	PDCP_Config_T_Reordering_Ms1750  = 30
	PDCP_Config_T_Reordering_Ms2000  = 31
	PDCP_Config_T_Reordering_Ms2250  = 32
	PDCP_Config_T_Reordering_Ms2500  = 33
	PDCP_Config_T_Reordering_Ms2750  = 34
	PDCP_Config_T_Reordering_Ms3000  = 35
	PDCP_Config_T_Reordering_Spare28 = 36
	PDCP_Config_T_Reordering_Spare27 = 37
	PDCP_Config_T_Reordering_Spare26 = 38
	PDCP_Config_T_Reordering_Spare25 = 39
	PDCP_Config_T_Reordering_Spare24 = 40
	PDCP_Config_T_Reordering_Spare23 = 41
	PDCP_Config_T_Reordering_Spare22 = 42
	PDCP_Config_T_Reordering_Spare21 = 43
	PDCP_Config_T_Reordering_Spare20 = 44
	PDCP_Config_T_Reordering_Spare19 = 45
	PDCP_Config_T_Reordering_Spare18 = 46
	PDCP_Config_T_Reordering_Spare17 = 47
	PDCP_Config_T_Reordering_Spare16 = 48
	PDCP_Config_T_Reordering_Spare15 = 49
	PDCP_Config_T_Reordering_Spare14 = 50
	PDCP_Config_T_Reordering_Spare13 = 51
	PDCP_Config_T_Reordering_Spare12 = 52
	PDCP_Config_T_Reordering_Spare11 = 53
	PDCP_Config_T_Reordering_Spare10 = 54
	PDCP_Config_T_Reordering_Spare09 = 55
	PDCP_Config_T_Reordering_Spare08 = 56
	PDCP_Config_T_Reordering_Spare07 = 57
	PDCP_Config_T_Reordering_Spare06 = 58
	PDCP_Config_T_Reordering_Spare05 = 59
	PDCP_Config_T_Reordering_Spare04 = 60
	PDCP_Config_T_Reordering_Spare03 = 61
	PDCP_Config_T_Reordering_Spare02 = 62
	PDCP_Config_T_Reordering_Spare01 = 63
)

var pDCPConfigTReorderingConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PDCP_Config_T_Reordering_Ms0, PDCP_Config_T_Reordering_Ms1, PDCP_Config_T_Reordering_Ms2, PDCP_Config_T_Reordering_Ms4, PDCP_Config_T_Reordering_Ms5, PDCP_Config_T_Reordering_Ms8, PDCP_Config_T_Reordering_Ms10, PDCP_Config_T_Reordering_Ms15, PDCP_Config_T_Reordering_Ms20, PDCP_Config_T_Reordering_Ms30, PDCP_Config_T_Reordering_Ms40, PDCP_Config_T_Reordering_Ms50, PDCP_Config_T_Reordering_Ms60, PDCP_Config_T_Reordering_Ms80, PDCP_Config_T_Reordering_Ms100, PDCP_Config_T_Reordering_Ms120, PDCP_Config_T_Reordering_Ms140, PDCP_Config_T_Reordering_Ms160, PDCP_Config_T_Reordering_Ms180, PDCP_Config_T_Reordering_Ms200, PDCP_Config_T_Reordering_Ms220, PDCP_Config_T_Reordering_Ms240, PDCP_Config_T_Reordering_Ms260, PDCP_Config_T_Reordering_Ms280, PDCP_Config_T_Reordering_Ms300, PDCP_Config_T_Reordering_Ms500, PDCP_Config_T_Reordering_Ms750, PDCP_Config_T_Reordering_Ms1000, PDCP_Config_T_Reordering_Ms1250, PDCP_Config_T_Reordering_Ms1500, PDCP_Config_T_Reordering_Ms1750, PDCP_Config_T_Reordering_Ms2000, PDCP_Config_T_Reordering_Ms2250, PDCP_Config_T_Reordering_Ms2500, PDCP_Config_T_Reordering_Ms2750, PDCP_Config_T_Reordering_Ms3000, PDCP_Config_T_Reordering_Spare28, PDCP_Config_T_Reordering_Spare27, PDCP_Config_T_Reordering_Spare26, PDCP_Config_T_Reordering_Spare25, PDCP_Config_T_Reordering_Spare24, PDCP_Config_T_Reordering_Spare23, PDCP_Config_T_Reordering_Spare22, PDCP_Config_T_Reordering_Spare21, PDCP_Config_T_Reordering_Spare20, PDCP_Config_T_Reordering_Spare19, PDCP_Config_T_Reordering_Spare18, PDCP_Config_T_Reordering_Spare17, PDCP_Config_T_Reordering_Spare16, PDCP_Config_T_Reordering_Spare15, PDCP_Config_T_Reordering_Spare14, PDCP_Config_T_Reordering_Spare13, PDCP_Config_T_Reordering_Spare12, PDCP_Config_T_Reordering_Spare11, PDCP_Config_T_Reordering_Spare10, PDCP_Config_T_Reordering_Spare09, PDCP_Config_T_Reordering_Spare08, PDCP_Config_T_Reordering_Spare07, PDCP_Config_T_Reordering_Spare06, PDCP_Config_T_Reordering_Spare05, PDCP_Config_T_Reordering_Spare04, PDCP_Config_T_Reordering_Spare03, PDCP_Config_T_Reordering_Spare02, PDCP_Config_T_Reordering_Spare01},
}

const (
	PDCP_Config_Ext_CipheringDisabled_True = 0
)

var pDCPConfigExtCipheringDisabledConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PDCP_Config_Ext_CipheringDisabled_True},
}

var pDCPConfigExtDiscardTimerExtR16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	PDCP_Config_Ext_DiscardTimerExt_r16_Release = 0
	PDCP_Config_Ext_DiscardTimerExt_r16_Setup   = 1
)

var pDCPConfigExtEthernetHeaderCompressionR16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	PDCP_Config_Ext_EthernetHeaderCompression_r16_Release = 0
	PDCP_Config_Ext_EthernetHeaderCompression_r16_Setup   = 1
)

const (
	PDCP_Config_Ext_SurvivalTimeStateSupport_r17_True = 0
)

var pDCPConfigExtSurvivalTimeStateSupportR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PDCP_Config_Ext_SurvivalTimeStateSupport_r17_True},
}

var pDCPConfigExtUplinkDataCompressionR17Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	PDCP_Config_Ext_UplinkDataCompression_r17_Release = 0
	PDCP_Config_Ext_UplinkDataCompression_r17_Setup   = 1
)

var pDCPConfigExtDiscardTimerExt2R17Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	PDCP_Config_Ext_DiscardTimerExt2_r17_Release = 0
	PDCP_Config_Ext_DiscardTimerExt2_r17_Setup   = 1
)

var pDCPConfigInitialRXDELIVR17Constraints = per.FixedSize(32)

const (
	PDCP_Config_Ext_Pdu_SetDiscard_r18_True = 0
)

var pDCPConfigExtPduSetDiscardR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PDCP_Config_Ext_Pdu_SetDiscard_r18_True},
}

var pDCPConfigExtDiscardTimerForLowImportanceR18Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	PDCP_Config_Ext_DiscardTimerForLowImportance_r18_Release = 0
	PDCP_Config_Ext_DiscardTimerForLowImportance_r18_Setup   = 1
)

const (
	PDCP_Config_Ext_PrimaryPathOnIndirectPath_r18_True = 0
)

var pDCPConfigExtPrimaryPathOnIndirectPathR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PDCP_Config_Ext_PrimaryPathOnIndirectPath_r18_True},
}

const (
	PDCP_Config_Ext_Sn_GapReport_r18_True = 0
)

var pDCPConfigExtSnGapReportR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PDCP_Config_Ext_Sn_GapReport_r18_True},
}

var pDCPConfigDrbConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "discardTimer", Optional: true},
		{Name: "pdcp-SN-SizeUL", Optional: true},
		{Name: "pdcp-SN-SizeDL", Optional: true},
		{Name: "headerCompression"},
		{Name: "integrityProtection", Optional: true},
		{Name: "statusReportRequired", Optional: true},
		{Name: "outOfOrderDelivery", Optional: true},
	},
}

const (
	PDCP_Config_Drb_DiscardTimer_Ms10     = 0
	PDCP_Config_Drb_DiscardTimer_Ms20     = 1
	PDCP_Config_Drb_DiscardTimer_Ms30     = 2
	PDCP_Config_Drb_DiscardTimer_Ms40     = 3
	PDCP_Config_Drb_DiscardTimer_Ms50     = 4
	PDCP_Config_Drb_DiscardTimer_Ms60     = 5
	PDCP_Config_Drb_DiscardTimer_Ms75     = 6
	PDCP_Config_Drb_DiscardTimer_Ms100    = 7
	PDCP_Config_Drb_DiscardTimer_Ms150    = 8
	PDCP_Config_Drb_DiscardTimer_Ms200    = 9
	PDCP_Config_Drb_DiscardTimer_Ms250    = 10
	PDCP_Config_Drb_DiscardTimer_Ms300    = 11
	PDCP_Config_Drb_DiscardTimer_Ms500    = 12
	PDCP_Config_Drb_DiscardTimer_Ms750    = 13
	PDCP_Config_Drb_DiscardTimer_Ms1500   = 14
	PDCP_Config_Drb_DiscardTimer_Infinity = 15
)

var pDCPConfigDrbDiscardTimerConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PDCP_Config_Drb_DiscardTimer_Ms10, PDCP_Config_Drb_DiscardTimer_Ms20, PDCP_Config_Drb_DiscardTimer_Ms30, PDCP_Config_Drb_DiscardTimer_Ms40, PDCP_Config_Drb_DiscardTimer_Ms50, PDCP_Config_Drb_DiscardTimer_Ms60, PDCP_Config_Drb_DiscardTimer_Ms75, PDCP_Config_Drb_DiscardTimer_Ms100, PDCP_Config_Drb_DiscardTimer_Ms150, PDCP_Config_Drb_DiscardTimer_Ms200, PDCP_Config_Drb_DiscardTimer_Ms250, PDCP_Config_Drb_DiscardTimer_Ms300, PDCP_Config_Drb_DiscardTimer_Ms500, PDCP_Config_Drb_DiscardTimer_Ms750, PDCP_Config_Drb_DiscardTimer_Ms1500, PDCP_Config_Drb_DiscardTimer_Infinity},
}

const (
	PDCP_Config_Drb_Pdcp_SN_SizeUL_Len12bits = 0
	PDCP_Config_Drb_Pdcp_SN_SizeUL_Len18bits = 1
)

var pDCPConfigDrbPdcpSNSizeULConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PDCP_Config_Drb_Pdcp_SN_SizeUL_Len12bits, PDCP_Config_Drb_Pdcp_SN_SizeUL_Len18bits},
}

const (
	PDCP_Config_Drb_Pdcp_SN_SizeDL_Len12bits = 0
	PDCP_Config_Drb_Pdcp_SN_SizeDL_Len18bits = 1
)

var pDCPConfigDrbPdcpSNSizeDLConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PDCP_Config_Drb_Pdcp_SN_SizeDL_Len12bits, PDCP_Config_Drb_Pdcp_SN_SizeDL_Len18bits},
}

var pDCPConfigDrbHeaderCompressionConstraints = per.ChoiceConstraints{
	Extensible: true,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "notUsed"},
		{Name: "rohc"},
		{Name: "uplinkOnlyROHC"},
	},
}

const (
	PDCP_Config_Drb_HeaderCompression_NotUsed        = 0
	PDCP_Config_Drb_HeaderCompression_Rohc           = 1
	PDCP_Config_Drb_HeaderCompression_UplinkOnlyROHC = 2
)

var pDCPConfigDrbHeaderCompressionRohcConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "maxCID"},
		{Name: "profiles"},
		{Name: "drb-ContinueROHC", Optional: true},
	},
}

const (
	PDCP_Config_Drb_HeaderCompression_Rohc_Drb_ContinueROHC_True = 0
)

var pDCPConfigDrbHeaderCompressionRohcDrbContinueROHCConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PDCP_Config_Drb_HeaderCompression_Rohc_Drb_ContinueROHC_True},
}

var pDCPConfigDrbHeaderCompressionUplinkOnlyROHCConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "maxCID"},
		{Name: "profiles"},
		{Name: "drb-ContinueROHC", Optional: true},
	},
}

const (
	PDCP_Config_Drb_HeaderCompression_UplinkOnlyROHC_Drb_ContinueROHC_True = 0
)

var pDCPConfigDrbHeaderCompressionUplinkOnlyROHCDrbContinueROHCConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PDCP_Config_Drb_HeaderCompression_UplinkOnlyROHC_Drb_ContinueROHC_True},
}

const (
	PDCP_Config_Drb_IntegrityProtection_Enabled = 0
)

var pDCPConfigDrbIntegrityProtectionConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PDCP_Config_Drb_IntegrityProtection_Enabled},
}

const (
	PDCP_Config_Drb_StatusReportRequired_True = 0
)

var pDCPConfigDrbStatusReportRequiredConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PDCP_Config_Drb_StatusReportRequired_True},
}

const (
	PDCP_Config_Drb_OutOfOrderDelivery_True = 0
)

var pDCPConfigDrbOutOfOrderDeliveryConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PDCP_Config_Drb_OutOfOrderDelivery_True},
}

var pDCPConfigMoreThanOneRLCConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "primaryPath"},
		{Name: "ul-DataSplitThreshold", Optional: true},
		{Name: "pdcp-Duplication", Optional: true},
	},
}

var pDCPConfigMoreThanOneRLCPrimaryPathConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "cellGroup", Optional: true},
		{Name: "logicalChannel", Optional: true},
	},
}

var pDCPConfigExtMoreThanTwoRLCDRBR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "splitSecondaryPath-r16", Optional: true},
		{Name: "duplicationState-r16", Optional: true},
	},
}

var pDCPConfigExtMoreThanTwoRLCDRBR16DuplicationStateR16Constraints = per.FixedSize(3)

type PDCP_Config struct {
	Drb *struct {
		DiscardTimer      *int64
		Pdcp_SN_SizeUL    *int64
		Pdcp_SN_SizeDL    *int64
		HeaderCompression struct {
			Choice  int
			NotUsed *struct{}
			Rohc    *struct {
				MaxCID   int64
				Profiles struct {
					Profile0x0001 bool
					Profile0x0002 bool
					Profile0x0003 bool
					Profile0x0004 bool
					Profile0x0006 bool
					Profile0x0101 bool
					Profile0x0102 bool
					Profile0x0103 bool
					Profile0x0104 bool
				}
				Drb_ContinueROHC *int64
			}
			UplinkOnlyROHC *struct {
				MaxCID           int64
				Profiles         struct{ Profile0x0006 bool }
				Drb_ContinueROHC *int64
			}
		}
		IntegrityProtection  *int64
		StatusReportRequired *int64
		OutOfOrderDelivery   *int64
	}
	MoreThanOneRLC *struct {
		PrimaryPath struct {
			CellGroup      *CellGroupId
			LogicalChannel *LogicalChannelIdentity
		}
		Ul_DataSplitThreshold *UL_DataSplitThreshold
		Pdcp_Duplication      *bool
	}
	T_Reordering        *int64
	CipheringDisabled   *int64
	DiscardTimerExt_r16 *struct {
		Choice  int
		Release *struct{}
		Setup   *DiscardTimerExt_r16
	}
	MoreThanTwoRLC_DRB_r16 *struct {
		SplitSecondaryPath_r16 *LogicalChannelIdentity
		DuplicationState_r16   []bool
	}
	EthernetHeaderCompression_r16 *struct {
		Choice  int
		Release *struct{}
		Setup   *EthernetHeaderCompression_r16
	}
	SurvivalTimeStateSupport_r17 *int64
	UplinkDataCompression_r17    *struct {
		Choice  int
		Release *struct{}
		Setup   *UplinkDataCompression_r17
	}
	DiscardTimerExt2_r17 *struct {
		Choice  int
		Release *struct{}
		Setup   *DiscardTimerExt2_r17
	}
	InitialRX_DELIV_r17              *per.BitString
	Pdu_SetDiscard_r18               *int64
	DiscardTimerForLowImportance_r18 *struct {
		Choice  int
		Release *struct{}
		Setup   *DiscardTimerForLowImportance_r18
	}
	PrimaryPathOnIndirectPath_r18         *int64
	Sn_GapReport_r18                      *int64
	RemainingTimeThresholdRLC_ReTx_r19    *RLC_AM_RemainingTimeThreshold_r19
	RemainingTimeThresholdRLC_Polling_r19 *RLC_AM_RemainingTimeThreshold_r19
}

func (ie *PDCP_Config) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(pDCPConfigConstraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.CipheringDisabled != nil
	hasExtGroup1 := ie.DiscardTimerExt_r16 != nil || ie.MoreThanTwoRLC_DRB_r16 != nil || ie.EthernetHeaderCompression_r16 != nil
	hasExtGroup2 := ie.SurvivalTimeStateSupport_r17 != nil || ie.UplinkDataCompression_r17 != nil || ie.DiscardTimerExt2_r17 != nil || ie.InitialRX_DELIV_r17 != nil
	hasExtGroup3 := ie.Pdu_SetDiscard_r18 != nil || ie.DiscardTimerForLowImportance_r18 != nil || ie.PrimaryPathOnIndirectPath_r18 != nil || ie.Sn_GapReport_r18 != nil
	hasExtGroup4 := ie.RemainingTimeThresholdRLC_ReTx_r19 != nil || ie.RemainingTimeThresholdRLC_Polling_r19 != nil
	hasExtensions := hasExtGroup0 || hasExtGroup1 || hasExtGroup2 || hasExtGroup3 || hasExtGroup4

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Drb != nil, ie.MoreThanOneRLC != nil, ie.T_Reordering != nil}); err != nil {
		return err
	}

	// 3. drb: sequence
	{
		if ie.Drb != nil {
			c := ie.Drb
			pDCPConfigDrbSeq := e.NewSequenceEncoder(pDCPConfigDrbConstraints)
			if err := pDCPConfigDrbSeq.EncodePreamble([]bool{c.DiscardTimer != nil, c.Pdcp_SN_SizeUL != nil, c.Pdcp_SN_SizeDL != nil, c.IntegrityProtection != nil, c.StatusReportRequired != nil, c.OutOfOrderDelivery != nil}); err != nil {
				return err
			}
			if c.DiscardTimer != nil {
				if err := e.EncodeEnumerated((*c.DiscardTimer), pDCPConfigDrbDiscardTimerConstraints); err != nil {
					return err
				}
			}
			if c.Pdcp_SN_SizeUL != nil {
				if err := e.EncodeEnumerated((*c.Pdcp_SN_SizeUL), pDCPConfigDrbPdcpSNSizeULConstraints); err != nil {
					return err
				}
			}
			if c.Pdcp_SN_SizeDL != nil {
				if err := e.EncodeEnumerated((*c.Pdcp_SN_SizeDL), pDCPConfigDrbPdcpSNSizeDLConstraints); err != nil {
					return err
				}
			}
			{
				choiceEnc := e.NewChoiceEncoder(pDCPConfigDrbHeaderCompressionConstraints)
				if err := choiceEnc.EncodeChoice(int64(c.HeaderCompression.Choice), false, nil); err != nil {
					return err
				}
				switch c.HeaderCompression.Choice {
				case PDCP_Config_Drb_HeaderCompression_NotUsed:
					if err := e.EncodeNull(); err != nil {
						return err
					}
				case PDCP_Config_Drb_HeaderCompression_Rohc:
					c := (*c.HeaderCompression.Rohc)
					pDCPConfigDrbHeaderCompressionRohcSeq := e.NewSequenceEncoder(pDCPConfigDrbHeaderCompressionRohcConstraints)
					if err := pDCPConfigDrbHeaderCompressionRohcSeq.EncodePreamble([]bool{c.Drb_ContinueROHC != nil}); err != nil {
						return err
					}
					if err := e.EncodeInteger(c.MaxCID, per.Constrained(1, 16383)); err != nil {
						return err
					}
					{
						c := &c.Profiles
						if err := e.EncodeBoolean(c.Profile0x0001); err != nil {
							return err
						}
						if err := e.EncodeBoolean(c.Profile0x0002); err != nil {
							return err
						}
						if err := e.EncodeBoolean(c.Profile0x0003); err != nil {
							return err
						}
						if err := e.EncodeBoolean(c.Profile0x0004); err != nil {
							return err
						}
						if err := e.EncodeBoolean(c.Profile0x0006); err != nil {
							return err
						}
						if err := e.EncodeBoolean(c.Profile0x0101); err != nil {
							return err
						}
						if err := e.EncodeBoolean(c.Profile0x0102); err != nil {
							return err
						}
						if err := e.EncodeBoolean(c.Profile0x0103); err != nil {
							return err
						}
						if err := e.EncodeBoolean(c.Profile0x0104); err != nil {
							return err
						}
					}
					if c.Drb_ContinueROHC != nil {
						if err := e.EncodeEnumerated((*c.Drb_ContinueROHC), pDCPConfigDrbHeaderCompressionRohcDrbContinueROHCConstraints); err != nil {
							return err
						}
					}
				case PDCP_Config_Drb_HeaderCompression_UplinkOnlyROHC:
					c := (*c.HeaderCompression.UplinkOnlyROHC)
					pDCPConfigDrbHeaderCompressionUplinkOnlyROHCSeq := e.NewSequenceEncoder(pDCPConfigDrbHeaderCompressionUplinkOnlyROHCConstraints)
					if err := pDCPConfigDrbHeaderCompressionUplinkOnlyROHCSeq.EncodePreamble([]bool{c.Drb_ContinueROHC != nil}); err != nil {
						return err
					}
					if err := e.EncodeInteger(c.MaxCID, per.Constrained(1, 16383)); err != nil {
						return err
					}
					{
						c := &c.Profiles
						if err := e.EncodeBoolean(c.Profile0x0006); err != nil {
							return err
						}
					}
					if c.Drb_ContinueROHC != nil {
						if err := e.EncodeEnumerated((*c.Drb_ContinueROHC), pDCPConfigDrbHeaderCompressionUplinkOnlyROHCDrbContinueROHCConstraints); err != nil {
							return err
						}
					}
				}
			}
			if c.IntegrityProtection != nil {
				if err := e.EncodeEnumerated((*c.IntegrityProtection), pDCPConfigDrbIntegrityProtectionConstraints); err != nil {
					return err
				}
			}
			if c.StatusReportRequired != nil {
				if err := e.EncodeEnumerated((*c.StatusReportRequired), pDCPConfigDrbStatusReportRequiredConstraints); err != nil {
					return err
				}
			}
			if c.OutOfOrderDelivery != nil {
				if err := e.EncodeEnumerated((*c.OutOfOrderDelivery), pDCPConfigDrbOutOfOrderDeliveryConstraints); err != nil {
					return err
				}
			}
		}
	}

	// 4. moreThanOneRLC: sequence
	{
		if ie.MoreThanOneRLC != nil {
			c := ie.MoreThanOneRLC
			pDCPConfigMoreThanOneRLCSeq := e.NewSequenceEncoder(pDCPConfigMoreThanOneRLCConstraints)
			if err := pDCPConfigMoreThanOneRLCSeq.EncodePreamble([]bool{c.Ul_DataSplitThreshold != nil, c.Pdcp_Duplication != nil}); err != nil {
				return err
			}
			{
				c := &c.PrimaryPath
				pDCPConfigMoreThanOneRLCPrimaryPathSeq := e.NewSequenceEncoder(pDCPConfigMoreThanOneRLCPrimaryPathConstraints)
				if err := pDCPConfigMoreThanOneRLCPrimaryPathSeq.EncodePreamble([]bool{c.CellGroup != nil, c.LogicalChannel != nil}); err != nil {
					return err
				}
				if c.CellGroup != nil {
					if err := c.CellGroup.Encode(e); err != nil {
						return err
					}
				}
				if c.LogicalChannel != nil {
					if err := c.LogicalChannel.Encode(e); err != nil {
						return err
					}
				}
			}
			if c.Ul_DataSplitThreshold != nil {
				if err := c.Ul_DataSplitThreshold.Encode(e); err != nil {
					return err
				}
			}
			if c.Pdcp_Duplication != nil {
				if err := e.EncodeBoolean((*c.Pdcp_Duplication)); err != nil {
					return err
				}
			}
		}
	}

	// 5. t-Reordering: enumerated
	{
		if ie.T_Reordering != nil {
			if err := e.EncodeEnumerated(*ie.T_Reordering, pDCPConfigTReorderingConstraints); err != nil {
				return err
			}
		}
	}

	if hasExtensions {
		extPresence := []bool{hasExtGroup0, hasExtGroup1, hasExtGroup2, hasExtGroup3, hasExtGroup4}
		var extValues [][]byte

		// Extension group 0:
		if hasExtGroup0 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "cipheringDisabled", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.CipheringDisabled != nil}); err != nil {
				return err
			}

			if ie.CipheringDisabled != nil {
				if err := ex.EncodeEnumerated(*ie.CipheringDisabled, pDCPConfigExtCipheringDisabledConstraints); err != nil {
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
					{Name: "discardTimerExt-r16", Optional: true},
					{Name: "moreThanTwoRLC-DRB-r16", Optional: true},
					{Name: "ethernetHeaderCompression-r16", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.DiscardTimerExt_r16 != nil, ie.MoreThanTwoRLC_DRB_r16 != nil, ie.EthernetHeaderCompression_r16 != nil}); err != nil {
				return err
			}

			if ie.DiscardTimerExt_r16 != nil {
				choiceEnc := ex.NewChoiceEncoder(pDCPConfigExtDiscardTimerExtR16Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.DiscardTimerExt_r16).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.DiscardTimerExt_r16).Choice {
				case PDCP_Config_Ext_DiscardTimerExt_r16_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case PDCP_Config_Ext_DiscardTimerExt_r16_Setup:
					if err := (*ie.DiscardTimerExt_r16).Setup.Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.MoreThanTwoRLC_DRB_r16 != nil {
				c := ie.MoreThanTwoRLC_DRB_r16
				pDCPConfigExtMoreThanTwoRLCDRBR16Seq := ex.NewSequenceEncoder(pDCPConfigExtMoreThanTwoRLCDRBR16Constraints)
				if err := pDCPConfigExtMoreThanTwoRLCDRBR16Seq.EncodePreamble([]bool{c.SplitSecondaryPath_r16 != nil, c.DuplicationState_r16 != nil}); err != nil {
					return err
				}
				if c.SplitSecondaryPath_r16 != nil {
					if err := c.SplitSecondaryPath_r16.Encode(ex); err != nil {
						return err
					}
				}
				if c.DuplicationState_r16 != nil {
					seqOf := ex.NewSequenceOfEncoder(pDCPConfigExtMoreThanTwoRLCDRBR16DuplicationStateR16Constraints)
					if err := seqOf.EncodeLength(int64(len(c.DuplicationState_r16))); err != nil {
						return err
					}
					for i := range c.DuplicationState_r16 {
						if err := ex.EncodeBoolean(c.DuplicationState_r16[i]); err != nil {
							return err
						}
					}
				}
			}

			if ie.EthernetHeaderCompression_r16 != nil {
				choiceEnc := ex.NewChoiceEncoder(pDCPConfigExtEthernetHeaderCompressionR16Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.EthernetHeaderCompression_r16).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.EthernetHeaderCompression_r16).Choice {
				case PDCP_Config_Ext_EthernetHeaderCompression_r16_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case PDCP_Config_Ext_EthernetHeaderCompression_r16_Setup:
					if err := (*ie.EthernetHeaderCompression_r16).Setup.Encode(ex); err != nil {
						return err
					}
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
					{Name: "survivalTimeStateSupport-r17", Optional: true},
					{Name: "uplinkDataCompression-r17", Optional: true},
					{Name: "discardTimerExt2-r17", Optional: true},
					{Name: "initialRX-DELIV-r17", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.SurvivalTimeStateSupport_r17 != nil, ie.UplinkDataCompression_r17 != nil, ie.DiscardTimerExt2_r17 != nil, ie.InitialRX_DELIV_r17 != nil}); err != nil {
				return err
			}

			if ie.SurvivalTimeStateSupport_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.SurvivalTimeStateSupport_r17, pDCPConfigExtSurvivalTimeStateSupportR17Constraints); err != nil {
					return err
				}
			}

			if ie.UplinkDataCompression_r17 != nil {
				choiceEnc := ex.NewChoiceEncoder(pDCPConfigExtUplinkDataCompressionR17Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.UplinkDataCompression_r17).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.UplinkDataCompression_r17).Choice {
				case PDCP_Config_Ext_UplinkDataCompression_r17_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case PDCP_Config_Ext_UplinkDataCompression_r17_Setup:
					if err := (*ie.UplinkDataCompression_r17).Setup.Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.DiscardTimerExt2_r17 != nil {
				choiceEnc := ex.NewChoiceEncoder(pDCPConfigExtDiscardTimerExt2R17Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.DiscardTimerExt2_r17).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.DiscardTimerExt2_r17).Choice {
				case PDCP_Config_Ext_DiscardTimerExt2_r17_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case PDCP_Config_Ext_DiscardTimerExt2_r17_Setup:
					if err := (*ie.DiscardTimerExt2_r17).Setup.Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.InitialRX_DELIV_r17 != nil {
				if err := ex.EncodeBitString(*ie.InitialRX_DELIV_r17, pDCPConfigInitialRXDELIVR17Constraints); err != nil {
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
					{Name: "pdu-SetDiscard-r18", Optional: true},
					{Name: "discardTimerForLowImportance-r18", Optional: true},
					{Name: "primaryPathOnIndirectPath-r18", Optional: true},
					{Name: "sn-GapReport-r18", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Pdu_SetDiscard_r18 != nil, ie.DiscardTimerForLowImportance_r18 != nil, ie.PrimaryPathOnIndirectPath_r18 != nil, ie.Sn_GapReport_r18 != nil}); err != nil {
				return err
			}

			if ie.Pdu_SetDiscard_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.Pdu_SetDiscard_r18, pDCPConfigExtPduSetDiscardR18Constraints); err != nil {
					return err
				}
			}

			if ie.DiscardTimerForLowImportance_r18 != nil {
				choiceEnc := ex.NewChoiceEncoder(pDCPConfigExtDiscardTimerForLowImportanceR18Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.DiscardTimerForLowImportance_r18).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.DiscardTimerForLowImportance_r18).Choice {
				case PDCP_Config_Ext_DiscardTimerForLowImportance_r18_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case PDCP_Config_Ext_DiscardTimerForLowImportance_r18_Setup:
					if err := (*ie.DiscardTimerForLowImportance_r18).Setup.Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.PrimaryPathOnIndirectPath_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.PrimaryPathOnIndirectPath_r18, pDCPConfigExtPrimaryPathOnIndirectPathR18Constraints); err != nil {
					return err
				}
			}

			if ie.Sn_GapReport_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.Sn_GapReport_r18, pDCPConfigExtSnGapReportR18Constraints); err != nil {
					return err
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
					{Name: "remainingTimeThresholdRLC-ReTx-r19", Optional: true},
					{Name: "remainingTimeThresholdRLC-Polling-r19", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.RemainingTimeThresholdRLC_ReTx_r19 != nil, ie.RemainingTimeThresholdRLC_Polling_r19 != nil}); err != nil {
				return err
			}

			if ie.RemainingTimeThresholdRLC_ReTx_r19 != nil {
				if err := ie.RemainingTimeThresholdRLC_ReTx_r19.Encode(ex); err != nil {
					return err
				}
			}

			if ie.RemainingTimeThresholdRLC_Polling_r19 != nil {
				if err := ie.RemainingTimeThresholdRLC_Polling_r19.Encode(ex); err != nil {
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

func (ie *PDCP_Config) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(pDCPConfigConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. drb: sequence
	{
		if seq.IsComponentPresent(0) {
			ie.Drb = &struct {
				DiscardTimer      *int64
				Pdcp_SN_SizeUL    *int64
				Pdcp_SN_SizeDL    *int64
				HeaderCompression struct {
					Choice  int
					NotUsed *struct{}
					Rohc    *struct {
						MaxCID   int64
						Profiles struct {
							Profile0x0001 bool
							Profile0x0002 bool
							Profile0x0003 bool
							Profile0x0004 bool
							Profile0x0006 bool
							Profile0x0101 bool
							Profile0x0102 bool
							Profile0x0103 bool
							Profile0x0104 bool
						}
						Drb_ContinueROHC *int64
					}
					UplinkOnlyROHC *struct {
						MaxCID           int64
						Profiles         struct{ Profile0x0006 bool }
						Drb_ContinueROHC *int64
					}
				}
				IntegrityProtection  *int64
				StatusReportRequired *int64
				OutOfOrderDelivery   *int64
			}{}
			c := ie.Drb
			pDCPConfigDrbSeq := d.NewSequenceDecoder(pDCPConfigDrbConstraints)
			if err := pDCPConfigDrbSeq.DecodePreamble(); err != nil {
				return err
			}
			if pDCPConfigDrbSeq.IsComponentPresent(0) {
				c.DiscardTimer = new(int64)
				v, err := d.DecodeEnumerated(pDCPConfigDrbDiscardTimerConstraints)
				if err != nil {
					return err
				}
				(*c.DiscardTimer) = v
			}
			if pDCPConfigDrbSeq.IsComponentPresent(1) {
				c.Pdcp_SN_SizeUL = new(int64)
				v, err := d.DecodeEnumerated(pDCPConfigDrbPdcpSNSizeULConstraints)
				if err != nil {
					return err
				}
				(*c.Pdcp_SN_SizeUL) = v
			}
			if pDCPConfigDrbSeq.IsComponentPresent(2) {
				c.Pdcp_SN_SizeDL = new(int64)
				v, err := d.DecodeEnumerated(pDCPConfigDrbPdcpSNSizeDLConstraints)
				if err != nil {
					return err
				}
				(*c.Pdcp_SN_SizeDL) = v
			}
			{
				choiceDec := d.NewChoiceDecoder(pDCPConfigDrbHeaderCompressionConstraints)
				index, _, _, err := choiceDec.DecodeChoice()
				if err != nil {
					return err
				}
				c.HeaderCompression.Choice = int(index)
				switch index {
				case PDCP_Config_Drb_HeaderCompression_NotUsed:
					c.HeaderCompression.NotUsed = &struct{}{}
					if err := d.DecodeNull(); err != nil {
						return err
					}
				case PDCP_Config_Drb_HeaderCompression_Rohc:
					c.HeaderCompression.Rohc = &struct {
						MaxCID   int64
						Profiles struct {
							Profile0x0001 bool
							Profile0x0002 bool
							Profile0x0003 bool
							Profile0x0004 bool
							Profile0x0006 bool
							Profile0x0101 bool
							Profile0x0102 bool
							Profile0x0103 bool
							Profile0x0104 bool
						}
						Drb_ContinueROHC *int64
					}{}
					c := (*c.HeaderCompression.Rohc)
					pDCPConfigDrbHeaderCompressionRohcSeq := d.NewSequenceDecoder(pDCPConfigDrbHeaderCompressionRohcConstraints)
					if err := pDCPConfigDrbHeaderCompressionRohcSeq.DecodePreamble(); err != nil {
						return err
					}
					{
						v, err := d.DecodeInteger(per.Constrained(1, 16383))
						if err != nil {
							return err
						}
						c.MaxCID = v
					}
					{
						c := &c.Profiles
						{
							v, err := d.DecodeBoolean()
							if err != nil {
								return err
							}
							c.Profile0x0001 = v
						}
						{
							v, err := d.DecodeBoolean()
							if err != nil {
								return err
							}
							c.Profile0x0002 = v
						}
						{
							v, err := d.DecodeBoolean()
							if err != nil {
								return err
							}
							c.Profile0x0003 = v
						}
						{
							v, err := d.DecodeBoolean()
							if err != nil {
								return err
							}
							c.Profile0x0004 = v
						}
						{
							v, err := d.DecodeBoolean()
							if err != nil {
								return err
							}
							c.Profile0x0006 = v
						}
						{
							v, err := d.DecodeBoolean()
							if err != nil {
								return err
							}
							c.Profile0x0101 = v
						}
						{
							v, err := d.DecodeBoolean()
							if err != nil {
								return err
							}
							c.Profile0x0102 = v
						}
						{
							v, err := d.DecodeBoolean()
							if err != nil {
								return err
							}
							c.Profile0x0103 = v
						}
						{
							v, err := d.DecodeBoolean()
							if err != nil {
								return err
							}
							c.Profile0x0104 = v
						}
					}
					if pDCPConfigDrbHeaderCompressionRohcSeq.IsComponentPresent(2) {
						c.Drb_ContinueROHC = new(int64)
						v, err := d.DecodeEnumerated(pDCPConfigDrbHeaderCompressionRohcDrbContinueROHCConstraints)
						if err != nil {
							return err
						}
						(*c.Drb_ContinueROHC) = v
					}
				case PDCP_Config_Drb_HeaderCompression_UplinkOnlyROHC:
					c.HeaderCompression.UplinkOnlyROHC = &struct {
						MaxCID           int64
						Profiles         struct{ Profile0x0006 bool }
						Drb_ContinueROHC *int64
					}{}
					c := (*c.HeaderCompression.UplinkOnlyROHC)
					pDCPConfigDrbHeaderCompressionUplinkOnlyROHCSeq := d.NewSequenceDecoder(pDCPConfigDrbHeaderCompressionUplinkOnlyROHCConstraints)
					if err := pDCPConfigDrbHeaderCompressionUplinkOnlyROHCSeq.DecodePreamble(); err != nil {
						return err
					}
					{
						v, err := d.DecodeInteger(per.Constrained(1, 16383))
						if err != nil {
							return err
						}
						c.MaxCID = v
					}
					{
						c := &c.Profiles
						{
							v, err := d.DecodeBoolean()
							if err != nil {
								return err
							}
							c.Profile0x0006 = v
						}
					}
					if pDCPConfigDrbHeaderCompressionUplinkOnlyROHCSeq.IsComponentPresent(2) {
						c.Drb_ContinueROHC = new(int64)
						v, err := d.DecodeEnumerated(pDCPConfigDrbHeaderCompressionUplinkOnlyROHCDrbContinueROHCConstraints)
						if err != nil {
							return err
						}
						(*c.Drb_ContinueROHC) = v
					}
				}
			}
			if pDCPConfigDrbSeq.IsComponentPresent(4) {
				c.IntegrityProtection = new(int64)
				v, err := d.DecodeEnumerated(pDCPConfigDrbIntegrityProtectionConstraints)
				if err != nil {
					return err
				}
				(*c.IntegrityProtection) = v
			}
			if pDCPConfigDrbSeq.IsComponentPresent(5) {
				c.StatusReportRequired = new(int64)
				v, err := d.DecodeEnumerated(pDCPConfigDrbStatusReportRequiredConstraints)
				if err != nil {
					return err
				}
				(*c.StatusReportRequired) = v
			}
			if pDCPConfigDrbSeq.IsComponentPresent(6) {
				c.OutOfOrderDelivery = new(int64)
				v, err := d.DecodeEnumerated(pDCPConfigDrbOutOfOrderDeliveryConstraints)
				if err != nil {
					return err
				}
				(*c.OutOfOrderDelivery) = v
			}
		}
	}

	// 4. moreThanOneRLC: sequence
	{
		if seq.IsComponentPresent(1) {
			ie.MoreThanOneRLC = &struct {
				PrimaryPath struct {
					CellGroup      *CellGroupId
					LogicalChannel *LogicalChannelIdentity
				}
				Ul_DataSplitThreshold *UL_DataSplitThreshold
				Pdcp_Duplication      *bool
			}{}
			c := ie.MoreThanOneRLC
			pDCPConfigMoreThanOneRLCSeq := d.NewSequenceDecoder(pDCPConfigMoreThanOneRLCConstraints)
			if err := pDCPConfigMoreThanOneRLCSeq.DecodePreamble(); err != nil {
				return err
			}
			{
				c := &c.PrimaryPath
				pDCPConfigMoreThanOneRLCPrimaryPathSeq := d.NewSequenceDecoder(pDCPConfigMoreThanOneRLCPrimaryPathConstraints)
				if err := pDCPConfigMoreThanOneRLCPrimaryPathSeq.DecodePreamble(); err != nil {
					return err
				}
				if pDCPConfigMoreThanOneRLCPrimaryPathSeq.IsComponentPresent(0) {
					c.CellGroup = new(CellGroupId)
					if err := (*c.CellGroup).Decode(d); err != nil {
						return err
					}
				}
				if pDCPConfigMoreThanOneRLCPrimaryPathSeq.IsComponentPresent(1) {
					c.LogicalChannel = new(LogicalChannelIdentity)
					if err := (*c.LogicalChannel).Decode(d); err != nil {
						return err
					}
				}
			}
			if pDCPConfigMoreThanOneRLCSeq.IsComponentPresent(1) {
				c.Ul_DataSplitThreshold = new(UL_DataSplitThreshold)
				if err := (*c.Ul_DataSplitThreshold).Decode(d); err != nil {
					return err
				}
			}
			if pDCPConfigMoreThanOneRLCSeq.IsComponentPresent(2) {
				c.Pdcp_Duplication = new(bool)
				v, err := d.DecodeBoolean()
				if err != nil {
					return err
				}
				(*c.Pdcp_Duplication) = v
			}
		}
	}

	// 5. t-Reordering: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(pDCPConfigTReorderingConstraints)
			if err != nil {
				return err
			}
			ie.T_Reordering = &idx
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
				{Name: "cipheringDisabled", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(pDCPConfigExtCipheringDisabledConstraints)
			if err != nil {
				return err
			}
			ie.CipheringDisabled = &v
		}
	}

	// Extension group 1:
	if seq.IsExtensionPresent(1) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "discardTimerExt-r16", Optional: true},
				{Name: "moreThanTwoRLC-DRB-r16", Optional: true},
				{Name: "ethernetHeaderCompression-r16", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.DiscardTimerExt_r16 = &struct {
				Choice  int
				Release *struct{}
				Setup   *DiscardTimerExt_r16
			}{}
			choiceDec := dx.NewChoiceDecoder(pDCPConfigExtDiscardTimerExtR16Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.DiscardTimerExt_r16).Choice = int(index)
			switch index {
			case PDCP_Config_Ext_DiscardTimerExt_r16_Release:
				(*ie.DiscardTimerExt_r16).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case PDCP_Config_Ext_DiscardTimerExt_r16_Setup:
				(*ie.DiscardTimerExt_r16).Setup = new(DiscardTimerExt_r16)
				if err := (*ie.DiscardTimerExt_r16).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(1) {
			ie.MoreThanTwoRLC_DRB_r16 = &struct {
				SplitSecondaryPath_r16 *LogicalChannelIdentity
				DuplicationState_r16   []bool
			}{}
			c := ie.MoreThanTwoRLC_DRB_r16
			pDCPConfigExtMoreThanTwoRLCDRBR16Seq := dx.NewSequenceDecoder(pDCPConfigExtMoreThanTwoRLCDRBR16Constraints)
			if err := pDCPConfigExtMoreThanTwoRLCDRBR16Seq.DecodePreamble(); err != nil {
				return err
			}
			if pDCPConfigExtMoreThanTwoRLCDRBR16Seq.IsComponentPresent(0) {
				c.SplitSecondaryPath_r16 = new(LogicalChannelIdentity)
				if err := (*c.SplitSecondaryPath_r16).Decode(dx); err != nil {
					return err
				}
			}
			if pDCPConfigExtMoreThanTwoRLCDRBR16Seq.IsComponentPresent(1) {
				seqOf := dx.NewSequenceOfDecoder(pDCPConfigExtMoreThanTwoRLCDRBR16DuplicationStateR16Constraints)
				n, err := seqOf.DecodeLength()
				if err != nil {
					return err
				}
				c.DuplicationState_r16 = make([]bool, n)
				for i := int64(0); i < n; i++ {
					v, err := dx.DecodeBoolean()
					if err != nil {
						return err
					}
					c.DuplicationState_r16[i] = v
				}
			}
		}

		if groupSeq.IsComponentPresent(2) {
			ie.EthernetHeaderCompression_r16 = &struct {
				Choice  int
				Release *struct{}
				Setup   *EthernetHeaderCompression_r16
			}{}
			choiceDec := dx.NewChoiceDecoder(pDCPConfigExtEthernetHeaderCompressionR16Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.EthernetHeaderCompression_r16).Choice = int(index)
			switch index {
			case PDCP_Config_Ext_EthernetHeaderCompression_r16_Release:
				(*ie.EthernetHeaderCompression_r16).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case PDCP_Config_Ext_EthernetHeaderCompression_r16_Setup:
				(*ie.EthernetHeaderCompression_r16).Setup = new(EthernetHeaderCompression_r16)
				if err := (*ie.EthernetHeaderCompression_r16).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}
	}

	// Extension group 2:
	if seq.IsExtensionPresent(2) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "survivalTimeStateSupport-r17", Optional: true},
				{Name: "uplinkDataCompression-r17", Optional: true},
				{Name: "discardTimerExt2-r17", Optional: true},
				{Name: "initialRX-DELIV-r17", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(pDCPConfigExtSurvivalTimeStateSupportR17Constraints)
			if err != nil {
				return err
			}
			ie.SurvivalTimeStateSupport_r17 = &v
		}

		if groupSeq.IsComponentPresent(1) {
			ie.UplinkDataCompression_r17 = &struct {
				Choice  int
				Release *struct{}
				Setup   *UplinkDataCompression_r17
			}{}
			choiceDec := dx.NewChoiceDecoder(pDCPConfigExtUplinkDataCompressionR17Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.UplinkDataCompression_r17).Choice = int(index)
			switch index {
			case PDCP_Config_Ext_UplinkDataCompression_r17_Release:
				(*ie.UplinkDataCompression_r17).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case PDCP_Config_Ext_UplinkDataCompression_r17_Setup:
				(*ie.UplinkDataCompression_r17).Setup = new(UplinkDataCompression_r17)
				if err := (*ie.UplinkDataCompression_r17).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(2) {
			ie.DiscardTimerExt2_r17 = &struct {
				Choice  int
				Release *struct{}
				Setup   *DiscardTimerExt2_r17
			}{}
			choiceDec := dx.NewChoiceDecoder(pDCPConfigExtDiscardTimerExt2R17Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.DiscardTimerExt2_r17).Choice = int(index)
			switch index {
			case PDCP_Config_Ext_DiscardTimerExt2_r17_Release:
				(*ie.DiscardTimerExt2_r17).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case PDCP_Config_Ext_DiscardTimerExt2_r17_Setup:
				(*ie.DiscardTimerExt2_r17).Setup = new(DiscardTimerExt2_r17)
				if err := (*ie.DiscardTimerExt2_r17).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(3) {
			v, err := dx.DecodeBitString(pDCPConfigInitialRXDELIVR17Constraints)
			if err != nil {
				return err
			}
			ie.InitialRX_DELIV_r17 = &v
		}
	}

	// Extension group 3:
	if seq.IsExtensionPresent(3) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "pdu-SetDiscard-r18", Optional: true},
				{Name: "discardTimerForLowImportance-r18", Optional: true},
				{Name: "primaryPathOnIndirectPath-r18", Optional: true},
				{Name: "sn-GapReport-r18", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(pDCPConfigExtPduSetDiscardR18Constraints)
			if err != nil {
				return err
			}
			ie.Pdu_SetDiscard_r18 = &v
		}

		if groupSeq.IsComponentPresent(1) {
			ie.DiscardTimerForLowImportance_r18 = &struct {
				Choice  int
				Release *struct{}
				Setup   *DiscardTimerForLowImportance_r18
			}{}
			choiceDec := dx.NewChoiceDecoder(pDCPConfigExtDiscardTimerForLowImportanceR18Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.DiscardTimerForLowImportance_r18).Choice = int(index)
			switch index {
			case PDCP_Config_Ext_DiscardTimerForLowImportance_r18_Release:
				(*ie.DiscardTimerForLowImportance_r18).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case PDCP_Config_Ext_DiscardTimerForLowImportance_r18_Setup:
				(*ie.DiscardTimerForLowImportance_r18).Setup = new(DiscardTimerForLowImportance_r18)
				if err := (*ie.DiscardTimerForLowImportance_r18).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(2) {
			v, err := dx.DecodeEnumerated(pDCPConfigExtPrimaryPathOnIndirectPathR18Constraints)
			if err != nil {
				return err
			}
			ie.PrimaryPathOnIndirectPath_r18 = &v
		}

		if groupSeq.IsComponentPresent(3) {
			v, err := dx.DecodeEnumerated(pDCPConfigExtSnGapReportR18Constraints)
			if err != nil {
				return err
			}
			ie.Sn_GapReport_r18 = &v
		}
	}

	// Extension group 4:
	if seq.IsExtensionPresent(4) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "remainingTimeThresholdRLC-ReTx-r19", Optional: true},
				{Name: "remainingTimeThresholdRLC-Polling-r19", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.RemainingTimeThresholdRLC_ReTx_r19 = new(RLC_AM_RemainingTimeThreshold_r19)
			if err := ie.RemainingTimeThresholdRLC_ReTx_r19.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(1) {
			ie.RemainingTimeThresholdRLC_Polling_r19 = new(RLC_AM_RemainingTimeThreshold_r19)
			if err := ie.RemainingTimeThresholdRLC_Polling_r19.Decode(dx); err != nil {
				return err
			}
		}
	}

	return nil
}
