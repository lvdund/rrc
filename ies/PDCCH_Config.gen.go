// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: PDCCH-Config (line 10989).

var pDCCHConfigConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "controlResourceSetToAddModList", Optional: true},
		{Name: "controlResourceSetToReleaseList", Optional: true},
		{Name: "searchSpacesToAddModList", Optional: true},
		{Name: "searchSpacesToReleaseList", Optional: true},
		{Name: "downlinkPreemption", Optional: true},
		{Name: "tpc-PUSCH", Optional: true},
		{Name: "tpc-PUCCH", Optional: true},
		{Name: "tpc-SRS", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
		{Name: "extension-addition-1", Optional: true},
		{Name: "extension-addition-2", Optional: true},
		{Name: "extension-addition-3", Optional: true},
	},
}

var pDCCHConfigControlResourceSetToAddModListConstraints = per.SizeRange(1, 3)

var pDCCHConfigControlResourceSetToReleaseListConstraints = per.SizeRange(1, 3)

var pDCCHConfigSearchSpacesToAddModListConstraints = per.SizeRange(1, 10)

var pDCCHConfigSearchSpacesToReleaseListConstraints = per.SizeRange(1, 10)

var pDCCH_ConfigDownlinkPreemptionConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	PDCCH_Config_DownlinkPreemption_Release = 0
	PDCCH_Config_DownlinkPreemption_Setup   = 1
)

var pDCCH_ConfigTpcPUSCHConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	PDCCH_Config_Tpc_PUSCH_Release = 0
	PDCCH_Config_Tpc_PUSCH_Setup   = 1
)

var pDCCH_ConfigTpcPUCCHConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	PDCCH_Config_Tpc_PUCCH_Release = 0
	PDCCH_Config_Tpc_PUCCH_Setup   = 1
)

var pDCCH_ConfigTpcSRSConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	PDCCH_Config_Tpc_SRS_Release = 0
	PDCCH_Config_Tpc_SRS_Setup   = 1
)

var pDCCHConfigExtControlResourceSetToAddModListSizeExtV1610Constraints = per.SizeRange(1, 2)

var pDCCHConfigExtControlResourceSetToReleaseListSizeExtR16Constraints = per.SizeRange(1, 5)

var pDCCHConfigExtSearchSpacesToAddModListExtR16Constraints = per.SizeRange(1, 10)

var pDCCHConfigExtUplinkCancellationR16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	PDCCH_Config_Ext_UplinkCancellation_r16_Release = 0
	PDCCH_Config_Ext_UplinkCancellation_r16_Setup   = 1
)

const (
	PDCCH_Config_Ext_MonitoringCapabilityConfig_r16_R15monitoringcapability = 0
	PDCCH_Config_Ext_MonitoringCapabilityConfig_r16_R16monitoringcapability = 1
)

var pDCCHConfigExtMonitoringCapabilityConfigR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PDCCH_Config_Ext_MonitoringCapabilityConfig_r16_R15monitoringcapability, PDCCH_Config_Ext_MonitoringCapabilityConfig_r16_R16monitoringcapability},
}

var pDCCHConfigExtSearchSpacesToAddModListExtV1700Constraints = per.SizeRange(1, 10)

const (
	PDCCH_Config_Ext_MonitoringCapabilityConfig_v1710_R17monitoringcapability = 0
)

var pDCCHConfigExtMonitoringCapabilityConfigV1710Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PDCCH_Config_Ext_MonitoringCapabilityConfig_v1710_R17monitoringcapability},
}

var pDCCHConfigExtPdcchSkippingDurationListR17Constraints = per.SizeRange(1, 3)

const (
	PDCCH_Config_Ext_Pdcch_MonitoringResumptionAfterNack_r18_True = 0
)

var pDCCHConfigExtPdcchMonitoringResumptionAfterNackR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PDCCH_Config_Ext_Pdcch_MonitoringResumptionAfterNack_r18_True},
}

var pDCCHConfigExtSearchSpacesToAddModListExtV1800Constraints = per.SizeRange(1, 10)

var pDCCHConfigExtSearchSpacesToAddModListExtV1900Constraints = per.SizeRange(1, 10)

type PDCCH_Config struct {
	ControlResourceSetToAddModList  []ControlResourceSet
	ControlResourceSetToReleaseList []ControlResourceSetId
	SearchSpacesToAddModList        []SearchSpace
	SearchSpacesToReleaseList       []SearchSpaceId
	DownlinkPreemption              *struct {
		Choice  int
		Release *struct{}
		Setup   *DownlinkPreemption
	}
	Tpc_PUSCH *struct {
		Choice  int
		Release *struct{}
		Setup   *PUSCH_TPC_CommandConfig
	}
	Tpc_PUCCH *struct {
		Choice  int
		Release *struct{}
		Setup   *PUCCH_TPC_CommandConfig
	}
	Tpc_SRS *struct {
		Choice  int
		Release *struct{}
		Setup   *SRS_TPC_CommandConfig
	}
	ControlResourceSetToAddModListSizeExt_v1610 []ControlResourceSet
	ControlResourceSetToReleaseListSizeExt_r16  []ControlResourceSetId_r16
	SearchSpacesToAddModListExt_r16             []SearchSpaceExt_r16
	UplinkCancellation_r16                      *struct {
		Choice  int
		Release *struct{}
		Setup   *UplinkCancellation_r16
	}
	MonitoringCapabilityConfig_r16          *int64
	SearchSpaceSwitchConfig_r16             *SearchSpaceSwitchConfig_r16
	SearchSpacesToAddModListExt_v1700       []SearchSpaceExt_v1700
	MonitoringCapabilityConfig_v1710        *int64
	SearchSpaceSwitchConfig_r17             *SearchSpaceSwitchConfig_r17
	Pdcch_SkippingDurationList_r17          []SCS_SpecificDuration_r17
	Pdcch_MonitoringResumptionAfterNack_r18 *int64
	SearchSpacesToAddModListExt_v1800       []SearchSpaceExt_v1800
	SearchSpaceSwitchConfig_r19             *SearchSpaceSwitchConfig_r19
	SearchSpacesToAddModListExt_v1900       []SearchSpaceExt_v1900
}

func (ie *PDCCH_Config) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(pDCCHConfigConstraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.ControlResourceSetToAddModListSizeExt_v1610 != nil || ie.ControlResourceSetToReleaseListSizeExt_r16 != nil || ie.SearchSpacesToAddModListExt_r16 != nil || ie.UplinkCancellation_r16 != nil || ie.MonitoringCapabilityConfig_r16 != nil || ie.SearchSpaceSwitchConfig_r16 != nil
	hasExtGroup1 := ie.SearchSpacesToAddModListExt_v1700 != nil || ie.MonitoringCapabilityConfig_v1710 != nil || ie.SearchSpaceSwitchConfig_r17 != nil || ie.Pdcch_SkippingDurationList_r17 != nil
	hasExtGroup2 := ie.Pdcch_MonitoringResumptionAfterNack_r18 != nil || ie.SearchSpacesToAddModListExt_v1800 != nil
	hasExtGroup3 := ie.SearchSpaceSwitchConfig_r19 != nil || ie.SearchSpacesToAddModListExt_v1900 != nil
	hasExtensions := hasExtGroup0 || hasExtGroup1 || hasExtGroup2 || hasExtGroup3

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.ControlResourceSetToAddModList != nil, ie.ControlResourceSetToReleaseList != nil, ie.SearchSpacesToAddModList != nil, ie.SearchSpacesToReleaseList != nil, ie.DownlinkPreemption != nil, ie.Tpc_PUSCH != nil, ie.Tpc_PUCCH != nil, ie.Tpc_SRS != nil}); err != nil {
		return err
	}

	// 3. controlResourceSetToAddModList: sequence-of
	{
		if ie.ControlResourceSetToAddModList != nil {
			seqOf := e.NewSequenceOfEncoder(pDCCHConfigControlResourceSetToAddModListConstraints)
			if err := seqOf.EncodeLength(int64(len(ie.ControlResourceSetToAddModList))); err != nil {
				return err
			}
			for i := range ie.ControlResourceSetToAddModList {
				if err := ie.ControlResourceSetToAddModList[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 4. controlResourceSetToReleaseList: sequence-of
	{
		if ie.ControlResourceSetToReleaseList != nil {
			seqOf := e.NewSequenceOfEncoder(pDCCHConfigControlResourceSetToReleaseListConstraints)
			if err := seqOf.EncodeLength(int64(len(ie.ControlResourceSetToReleaseList))); err != nil {
				return err
			}
			for i := range ie.ControlResourceSetToReleaseList {
				if err := ie.ControlResourceSetToReleaseList[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 5. searchSpacesToAddModList: sequence-of
	{
		if ie.SearchSpacesToAddModList != nil {
			seqOf := e.NewSequenceOfEncoder(pDCCHConfigSearchSpacesToAddModListConstraints)
			if err := seqOf.EncodeLength(int64(len(ie.SearchSpacesToAddModList))); err != nil {
				return err
			}
			for i := range ie.SearchSpacesToAddModList {
				if err := ie.SearchSpacesToAddModList[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 6. searchSpacesToReleaseList: sequence-of
	{
		if ie.SearchSpacesToReleaseList != nil {
			seqOf := e.NewSequenceOfEncoder(pDCCHConfigSearchSpacesToReleaseListConstraints)
			if err := seqOf.EncodeLength(int64(len(ie.SearchSpacesToReleaseList))); err != nil {
				return err
			}
			for i := range ie.SearchSpacesToReleaseList {
				if err := ie.SearchSpacesToReleaseList[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 7. downlinkPreemption: choice
	{
		if ie.DownlinkPreemption != nil {
			choiceEnc := e.NewChoiceEncoder(pDCCH_ConfigDownlinkPreemptionConstraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.DownlinkPreemption).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.DownlinkPreemption).Choice {
			case PDCCH_Config_DownlinkPreemption_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case PDCCH_Config_DownlinkPreemption_Setup:
				if err := (*ie.DownlinkPreemption).Setup.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.DownlinkPreemption).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 8. tpc-PUSCH: choice
	{
		if ie.Tpc_PUSCH != nil {
			choiceEnc := e.NewChoiceEncoder(pDCCH_ConfigTpcPUSCHConstraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.Tpc_PUSCH).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.Tpc_PUSCH).Choice {
			case PDCCH_Config_Tpc_PUSCH_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case PDCCH_Config_Tpc_PUSCH_Setup:
				if err := (*ie.Tpc_PUSCH).Setup.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.Tpc_PUSCH).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 9. tpc-PUCCH: choice
	{
		if ie.Tpc_PUCCH != nil {
			choiceEnc := e.NewChoiceEncoder(pDCCH_ConfigTpcPUCCHConstraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.Tpc_PUCCH).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.Tpc_PUCCH).Choice {
			case PDCCH_Config_Tpc_PUCCH_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case PDCCH_Config_Tpc_PUCCH_Setup:
				if err := (*ie.Tpc_PUCCH).Setup.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.Tpc_PUCCH).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 10. tpc-SRS: choice
	{
		if ie.Tpc_SRS != nil {
			choiceEnc := e.NewChoiceEncoder(pDCCH_ConfigTpcSRSConstraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.Tpc_SRS).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.Tpc_SRS).Choice {
			case PDCCH_Config_Tpc_SRS_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case PDCCH_Config_Tpc_SRS_Setup:
				if err := (*ie.Tpc_SRS).Setup.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.Tpc_SRS).Choice), Constraint: "index out of range"}
			}
		}
	}

	if hasExtensions {
		extPresence := []bool{hasExtGroup0, hasExtGroup1, hasExtGroup2, hasExtGroup3}
		var extValues [][]byte

		// Extension group 0:
		if hasExtGroup0 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "controlResourceSetToAddModListSizeExt-v1610", Optional: true},
					{Name: "controlResourceSetToReleaseListSizeExt-r16", Optional: true},
					{Name: "searchSpacesToAddModListExt-r16", Optional: true},
					{Name: "uplinkCancellation-r16", Optional: true},
					{Name: "monitoringCapabilityConfig-r16", Optional: true},
					{Name: "searchSpaceSwitchConfig-r16", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.ControlResourceSetToAddModListSizeExt_v1610 != nil, ie.ControlResourceSetToReleaseListSizeExt_r16 != nil, ie.SearchSpacesToAddModListExt_r16 != nil, ie.UplinkCancellation_r16 != nil, ie.MonitoringCapabilityConfig_r16 != nil, ie.SearchSpaceSwitchConfig_r16 != nil}); err != nil {
				return err
			}

			if ie.ControlResourceSetToAddModListSizeExt_v1610 != nil {
				seqOf := ex.NewSequenceOfEncoder(pDCCHConfigExtControlResourceSetToAddModListSizeExtV1610Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.ControlResourceSetToAddModListSizeExt_v1610))); err != nil {
					return err
				}
				for i := range ie.ControlResourceSetToAddModListSizeExt_v1610 {
					if err := ie.ControlResourceSetToAddModListSizeExt_v1610[i].Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.ControlResourceSetToReleaseListSizeExt_r16 != nil {
				seqOf := ex.NewSequenceOfEncoder(pDCCHConfigExtControlResourceSetToReleaseListSizeExtR16Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.ControlResourceSetToReleaseListSizeExt_r16))); err != nil {
					return err
				}
				for i := range ie.ControlResourceSetToReleaseListSizeExt_r16 {
					if err := ie.ControlResourceSetToReleaseListSizeExt_r16[i].Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.SearchSpacesToAddModListExt_r16 != nil {
				seqOf := ex.NewSequenceOfEncoder(pDCCHConfigExtSearchSpacesToAddModListExtR16Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.SearchSpacesToAddModListExt_r16))); err != nil {
					return err
				}
				for i := range ie.SearchSpacesToAddModListExt_r16 {
					if err := ie.SearchSpacesToAddModListExt_r16[i].Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.UplinkCancellation_r16 != nil {
				choiceEnc := ex.NewChoiceEncoder(pDCCHConfigExtUplinkCancellationR16Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.UplinkCancellation_r16).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.UplinkCancellation_r16).Choice {
				case PDCCH_Config_Ext_UplinkCancellation_r16_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case PDCCH_Config_Ext_UplinkCancellation_r16_Setup:
					if err := (*ie.UplinkCancellation_r16).Setup.Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.MonitoringCapabilityConfig_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.MonitoringCapabilityConfig_r16, pDCCHConfigExtMonitoringCapabilityConfigR16Constraints); err != nil {
					return err
				}
			}

			if ie.SearchSpaceSwitchConfig_r16 != nil {
				if err := ie.SearchSpaceSwitchConfig_r16.Encode(ex); err != nil {
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
					{Name: "searchSpacesToAddModListExt-v1700", Optional: true},
					{Name: "monitoringCapabilityConfig-v1710", Optional: true},
					{Name: "searchSpaceSwitchConfig-r17", Optional: true},
					{Name: "pdcch-SkippingDurationList-r17", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.SearchSpacesToAddModListExt_v1700 != nil, ie.MonitoringCapabilityConfig_v1710 != nil, ie.SearchSpaceSwitchConfig_r17 != nil, ie.Pdcch_SkippingDurationList_r17 != nil}); err != nil {
				return err
			}

			if ie.SearchSpacesToAddModListExt_v1700 != nil {
				seqOf := ex.NewSequenceOfEncoder(pDCCHConfigExtSearchSpacesToAddModListExtV1700Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.SearchSpacesToAddModListExt_v1700))); err != nil {
					return err
				}
				for i := range ie.SearchSpacesToAddModListExt_v1700 {
					if err := ie.SearchSpacesToAddModListExt_v1700[i].Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.MonitoringCapabilityConfig_v1710 != nil {
				if err := ex.EncodeEnumerated(*ie.MonitoringCapabilityConfig_v1710, pDCCHConfigExtMonitoringCapabilityConfigV1710Constraints); err != nil {
					return err
				}
			}

			if ie.SearchSpaceSwitchConfig_r17 != nil {
				if err := ie.SearchSpaceSwitchConfig_r17.Encode(ex); err != nil {
					return err
				}
			}

			if ie.Pdcch_SkippingDurationList_r17 != nil {
				seqOf := ex.NewSequenceOfEncoder(pDCCHConfigExtPdcchSkippingDurationListR17Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.Pdcch_SkippingDurationList_r17))); err != nil {
					return err
				}
				for i := range ie.Pdcch_SkippingDurationList_r17 {
					if err := ie.Pdcch_SkippingDurationList_r17[i].Encode(ex); err != nil {
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
					{Name: "pdcch-MonitoringResumptionAfterNack-r18", Optional: true},
					{Name: "searchSpacesToAddModListExt-v1800", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Pdcch_MonitoringResumptionAfterNack_r18 != nil, ie.SearchSpacesToAddModListExt_v1800 != nil}); err != nil {
				return err
			}

			if ie.Pdcch_MonitoringResumptionAfterNack_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.Pdcch_MonitoringResumptionAfterNack_r18, pDCCHConfigExtPdcchMonitoringResumptionAfterNackR18Constraints); err != nil {
					return err
				}
			}

			if ie.SearchSpacesToAddModListExt_v1800 != nil {
				seqOf := ex.NewSequenceOfEncoder(pDCCHConfigExtSearchSpacesToAddModListExtV1800Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.SearchSpacesToAddModListExt_v1800))); err != nil {
					return err
				}
				for i := range ie.SearchSpacesToAddModListExt_v1800 {
					if err := ie.SearchSpacesToAddModListExt_v1800[i].Encode(ex); err != nil {
						return err
					}
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
					{Name: "searchSpaceSwitchConfig-r19", Optional: true},
					{Name: "searchSpacesToAddModListExt-v1900", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.SearchSpaceSwitchConfig_r19 != nil, ie.SearchSpacesToAddModListExt_v1900 != nil}); err != nil {
				return err
			}

			if ie.SearchSpaceSwitchConfig_r19 != nil {
				if err := ie.SearchSpaceSwitchConfig_r19.Encode(ex); err != nil {
					return err
				}
			}

			if ie.SearchSpacesToAddModListExt_v1900 != nil {
				seqOf := ex.NewSequenceOfEncoder(pDCCHConfigExtSearchSpacesToAddModListExtV1900Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.SearchSpacesToAddModListExt_v1900))); err != nil {
					return err
				}
				for i := range ie.SearchSpacesToAddModListExt_v1900 {
					if err := ie.SearchSpacesToAddModListExt_v1900[i].Encode(ex); err != nil {
						return err
					}
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

func (ie *PDCCH_Config) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(pDCCHConfigConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. controlResourceSetToAddModList: sequence-of
	{
		if seq.IsComponentPresent(0) {
			seqOf := d.NewSequenceOfDecoder(pDCCHConfigControlResourceSetToAddModListConstraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.ControlResourceSetToAddModList = make([]ControlResourceSet, n)
			for i := int64(0); i < n; i++ {
				if err := ie.ControlResourceSetToAddModList[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 4. controlResourceSetToReleaseList: sequence-of
	{
		if seq.IsComponentPresent(1) {
			seqOf := d.NewSequenceOfDecoder(pDCCHConfigControlResourceSetToReleaseListConstraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.ControlResourceSetToReleaseList = make([]ControlResourceSetId, n)
			for i := int64(0); i < n; i++ {
				if err := ie.ControlResourceSetToReleaseList[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 5. searchSpacesToAddModList: sequence-of
	{
		if seq.IsComponentPresent(2) {
			seqOf := d.NewSequenceOfDecoder(pDCCHConfigSearchSpacesToAddModListConstraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.SearchSpacesToAddModList = make([]SearchSpace, n)
			for i := int64(0); i < n; i++ {
				if err := ie.SearchSpacesToAddModList[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 6. searchSpacesToReleaseList: sequence-of
	{
		if seq.IsComponentPresent(3) {
			seqOf := d.NewSequenceOfDecoder(pDCCHConfigSearchSpacesToReleaseListConstraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.SearchSpacesToReleaseList = make([]SearchSpaceId, n)
			for i := int64(0); i < n; i++ {
				if err := ie.SearchSpacesToReleaseList[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 7. downlinkPreemption: choice
	{
		if seq.IsComponentPresent(4) {
			ie.DownlinkPreemption = &struct {
				Choice  int
				Release *struct{}
				Setup   *DownlinkPreemption
			}{}
			choiceDec := d.NewChoiceDecoder(pDCCH_ConfigDownlinkPreemptionConstraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.DownlinkPreemption).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case PDCCH_Config_DownlinkPreemption_Release:
				(*ie.DownlinkPreemption).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case PDCCH_Config_DownlinkPreemption_Setup:
				(*ie.DownlinkPreemption).Setup = new(DownlinkPreemption)
				if err := (*ie.DownlinkPreemption).Setup.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 8. tpc-PUSCH: choice
	{
		if seq.IsComponentPresent(5) {
			ie.Tpc_PUSCH = &struct {
				Choice  int
				Release *struct{}
				Setup   *PUSCH_TPC_CommandConfig
			}{}
			choiceDec := d.NewChoiceDecoder(pDCCH_ConfigTpcPUSCHConstraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Tpc_PUSCH).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case PDCCH_Config_Tpc_PUSCH_Release:
				(*ie.Tpc_PUSCH).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case PDCCH_Config_Tpc_PUSCH_Setup:
				(*ie.Tpc_PUSCH).Setup = new(PUSCH_TPC_CommandConfig)
				if err := (*ie.Tpc_PUSCH).Setup.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 9. tpc-PUCCH: choice
	{
		if seq.IsComponentPresent(6) {
			ie.Tpc_PUCCH = &struct {
				Choice  int
				Release *struct{}
				Setup   *PUCCH_TPC_CommandConfig
			}{}
			choiceDec := d.NewChoiceDecoder(pDCCH_ConfigTpcPUCCHConstraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Tpc_PUCCH).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case PDCCH_Config_Tpc_PUCCH_Release:
				(*ie.Tpc_PUCCH).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case PDCCH_Config_Tpc_PUCCH_Setup:
				(*ie.Tpc_PUCCH).Setup = new(PUCCH_TPC_CommandConfig)
				if err := (*ie.Tpc_PUCCH).Setup.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 10. tpc-SRS: choice
	{
		if seq.IsComponentPresent(7) {
			ie.Tpc_SRS = &struct {
				Choice  int
				Release *struct{}
				Setup   *SRS_TPC_CommandConfig
			}{}
			choiceDec := d.NewChoiceDecoder(pDCCH_ConfigTpcSRSConstraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Tpc_SRS).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case PDCCH_Config_Tpc_SRS_Release:
				(*ie.Tpc_SRS).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case PDCCH_Config_Tpc_SRS_Setup:
				(*ie.Tpc_SRS).Setup = new(SRS_TPC_CommandConfig)
				if err := (*ie.Tpc_SRS).Setup.Decode(d); err != nil {
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
				{Name: "controlResourceSetToAddModListSizeExt-v1610", Optional: true},
				{Name: "controlResourceSetToReleaseListSizeExt-r16", Optional: true},
				{Name: "searchSpacesToAddModListExt-r16", Optional: true},
				{Name: "uplinkCancellation-r16", Optional: true},
				{Name: "monitoringCapabilityConfig-r16", Optional: true},
				{Name: "searchSpaceSwitchConfig-r16", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			seqOf := dx.NewSequenceOfDecoder(pDCCHConfigExtControlResourceSetToAddModListSizeExtV1610Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.ControlResourceSetToAddModListSizeExt_v1610 = make([]ControlResourceSet, n)
			for i := int64(0); i < n; i++ {
				if err := ie.ControlResourceSetToAddModListSizeExt_v1610[i].Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(1) {
			seqOf := dx.NewSequenceOfDecoder(pDCCHConfigExtControlResourceSetToReleaseListSizeExtR16Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.ControlResourceSetToReleaseListSizeExt_r16 = make([]ControlResourceSetId_r16, n)
			for i := int64(0); i < n; i++ {
				if err := ie.ControlResourceSetToReleaseListSizeExt_r16[i].Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(2) {
			seqOf := dx.NewSequenceOfDecoder(pDCCHConfigExtSearchSpacesToAddModListExtR16Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.SearchSpacesToAddModListExt_r16 = make([]SearchSpaceExt_r16, n)
			for i := int64(0); i < n; i++ {
				if err := ie.SearchSpacesToAddModListExt_r16[i].Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(3) {
			ie.UplinkCancellation_r16 = &struct {
				Choice  int
				Release *struct{}
				Setup   *UplinkCancellation_r16
			}{}
			choiceDec := dx.NewChoiceDecoder(pDCCHConfigExtUplinkCancellationR16Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.UplinkCancellation_r16).Choice = int(index)
			switch index {
			case PDCCH_Config_Ext_UplinkCancellation_r16_Release:
				(*ie.UplinkCancellation_r16).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case PDCCH_Config_Ext_UplinkCancellation_r16_Setup:
				(*ie.UplinkCancellation_r16).Setup = new(UplinkCancellation_r16)
				if err := (*ie.UplinkCancellation_r16).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(4) {
			v, err := dx.DecodeEnumerated(pDCCHConfigExtMonitoringCapabilityConfigR16Constraints)
			if err != nil {
				return err
			}
			ie.MonitoringCapabilityConfig_r16 = &v
		}

		if groupSeq.IsComponentPresent(5) {
			ie.SearchSpaceSwitchConfig_r16 = new(SearchSpaceSwitchConfig_r16)
			if err := ie.SearchSpaceSwitchConfig_r16.Decode(dx); err != nil {
				return err
			}
		}
	}

	// Extension group 1:
	if seq.IsExtensionPresent(1) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "searchSpacesToAddModListExt-v1700", Optional: true},
				{Name: "monitoringCapabilityConfig-v1710", Optional: true},
				{Name: "searchSpaceSwitchConfig-r17", Optional: true},
				{Name: "pdcch-SkippingDurationList-r17", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			seqOf := dx.NewSequenceOfDecoder(pDCCHConfigExtSearchSpacesToAddModListExtV1700Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.SearchSpacesToAddModListExt_v1700 = make([]SearchSpaceExt_v1700, n)
			for i := int64(0); i < n; i++ {
				if err := ie.SearchSpacesToAddModListExt_v1700[i].Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(1) {
			v, err := dx.DecodeEnumerated(pDCCHConfigExtMonitoringCapabilityConfigV1710Constraints)
			if err != nil {
				return err
			}
			ie.MonitoringCapabilityConfig_v1710 = &v
		}

		if groupSeq.IsComponentPresent(2) {
			ie.SearchSpaceSwitchConfig_r17 = new(SearchSpaceSwitchConfig_r17)
			if err := ie.SearchSpaceSwitchConfig_r17.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(3) {
			seqOf := dx.NewSequenceOfDecoder(pDCCHConfigExtPdcchSkippingDurationListR17Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Pdcch_SkippingDurationList_r17 = make([]SCS_SpecificDuration_r17, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Pdcch_SkippingDurationList_r17[i].Decode(dx); err != nil {
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
				{Name: "pdcch-MonitoringResumptionAfterNack-r18", Optional: true},
				{Name: "searchSpacesToAddModListExt-v1800", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(pDCCHConfigExtPdcchMonitoringResumptionAfterNackR18Constraints)
			if err != nil {
				return err
			}
			ie.Pdcch_MonitoringResumptionAfterNack_r18 = &v
		}

		if groupSeq.IsComponentPresent(1) {
			seqOf := dx.NewSequenceOfDecoder(pDCCHConfigExtSearchSpacesToAddModListExtV1800Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.SearchSpacesToAddModListExt_v1800 = make([]SearchSpaceExt_v1800, n)
			for i := int64(0); i < n; i++ {
				if err := ie.SearchSpacesToAddModListExt_v1800[i].Decode(dx); err != nil {
					return err
				}
			}
		}
	}

	// Extension group 3:
	if seq.IsExtensionPresent(3) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "searchSpaceSwitchConfig-r19", Optional: true},
				{Name: "searchSpacesToAddModListExt-v1900", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.SearchSpaceSwitchConfig_r19 = new(SearchSpaceSwitchConfig_r19)
			if err := ie.SearchSpaceSwitchConfig_r19.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(1) {
			seqOf := dx.NewSequenceOfDecoder(pDCCHConfigExtSearchSpacesToAddModListExtV1900Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.SearchSpacesToAddModListExt_v1900 = make([]SearchSpaceExt_v1900, n)
			for i := int64(0); i < n; i++ {
				if err := ie.SearchSpacesToAddModListExt_v1900[i].Decode(dx); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
