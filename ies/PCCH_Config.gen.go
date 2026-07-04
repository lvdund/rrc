// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: PCCH-Config (line 7859).

var pCCHConfigConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "defaultPagingCycle"},
		{Name: "nAndPagingFrameOffset"},
		{Name: "ns"},
		{Name: "firstPDCCH-MonitoringOccasionOfPO", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
		{Name: "extension-addition-1", Optional: true},
		{Name: "extension-addition-2", Optional: true},
	},
}

var pCCH_ConfigNAndPagingFrameOffsetConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "oneT"},
		{Name: "halfT"},
		{Name: "quarterT"},
		{Name: "oneEighthT"},
		{Name: "oneSixteenthT"},
	},
}

const (
	PCCH_Config_NAndPagingFrameOffset_OneT          = 0
	PCCH_Config_NAndPagingFrameOffset_HalfT         = 1
	PCCH_Config_NAndPagingFrameOffset_QuarterT      = 2
	PCCH_Config_NAndPagingFrameOffset_OneEighthT    = 3
	PCCH_Config_NAndPagingFrameOffset_OneSixteenthT = 4
)

const (
	PCCH_Config_Ns_Four = 0
	PCCH_Config_Ns_Two  = 1
	PCCH_Config_Ns_One  = 2
)

var pCCHConfigNsConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PCCH_Config_Ns_Four, PCCH_Config_Ns_Two, PCCH_Config_Ns_One},
}

var pCCH_ConfigFirstPDCCHMonitoringOccasionOfPOConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "sCS15KHZoneT"},
		{Name: "sCS30KHZoneT-SCS15KHZhalfT"},
		{Name: "sCS60KHZoneT-SCS30KHZhalfT-SCS15KHZquarterT"},
		{Name: "sCS120KHZoneT-SCS60KHZhalfT-SCS30KHZquarterT-SCS15KHZoneEighthT"},
		{Name: "sCS120KHZhalfT-SCS60KHZquarterT-SCS30KHZoneEighthT-SCS15KHZoneSixteenthT"},
		{Name: "sCS480KHZoneT-SCS120KHZquarterT-SCS60KHZoneEighthT-SCS30KHZoneSixteenthT"},
		{Name: "sCS480KHZhalfT-SCS120KHZoneEighthT-SCS60KHZoneSixteenthT"},
		{Name: "sCS480KHZquarterT-SCS120KHZoneSixteenthT"},
	},
}

const (
	PCCH_Config_FirstPDCCH_MonitoringOccasionOfPO_SCS15KHZoneT                                                             = 0
	PCCH_Config_FirstPDCCH_MonitoringOccasionOfPO_SCS30KHZoneT_SCS15KHZhalfT                                               = 1
	PCCH_Config_FirstPDCCH_MonitoringOccasionOfPO_SCS60KHZoneT_SCS30KHZhalfT_SCS15KHZquarterT                              = 2
	PCCH_Config_FirstPDCCH_MonitoringOccasionOfPO_SCS120KHZoneT_SCS60KHZhalfT_SCS30KHZquarterT_SCS15KHZoneEighthT          = 3
	PCCH_Config_FirstPDCCH_MonitoringOccasionOfPO_SCS120KHZhalfT_SCS60KHZquarterT_SCS30KHZoneEighthT_SCS15KHZoneSixteenthT = 4
	PCCH_Config_FirstPDCCH_MonitoringOccasionOfPO_SCS480KHZoneT_SCS120KHZquarterT_SCS60KHZoneEighthT_SCS30KHZoneSixteenthT = 5
	PCCH_Config_FirstPDCCH_MonitoringOccasionOfPO_SCS480KHZhalfT_SCS120KHZoneEighthT_SCS60KHZoneSixteenthT                 = 6
	PCCH_Config_FirstPDCCH_MonitoringOccasionOfPO_SCS480KHZquarterT_SCS120KHZoneSixteenthT                                 = 7
)

var pCCHConfigNrofPDCCHMonitoringOccasionPerSSBInPOR16Constraints = per.Constrained(2, 4)

const (
	PCCH_Config_Ext_RanPagingInIdlePO_r17_True = 0
)

var pCCHConfigExtRanPagingInIdlePOR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PCCH_Config_Ext_RanPagingInIdlePO_r17_True},
}

var pCCHConfigExtFirstPDCCHMonitoringOccasionOfPOV1710Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "sCS480KHZoneEighthT"},
		{Name: "sCS480KHZoneSixteenthT"},
	},
}

const (
	PCCH_Config_Ext_FirstPDCCH_MonitoringOccasionOfPO_v1710_SCS480KHZoneEighthT    = 0
	PCCH_Config_Ext_FirstPDCCH_MonitoringOccasionOfPO_v1710_SCS480KHZoneSixteenthT = 1
)

var pCCHConfigFirstPDCCHMonitoringOccasionOfPOSCS15KHZoneTConstraints = per.SizeRange(1, common.MaxPO_PerPF)

var pCCHConfigFirstPDCCHMonitoringOccasionOfPOSCS30KHZoneTSCS15KHZhalfTConstraints = per.SizeRange(1, common.MaxPO_PerPF)

var pCCHConfigFirstPDCCHMonitoringOccasionOfPOSCS60KHZoneTSCS30KHZhalfTSCS15KHZquarterTConstraints = per.SizeRange(1, common.MaxPO_PerPF)

var pCCHConfigFirstPDCCHMonitoringOccasionOfPOSCS120KHZoneTSCS60KHZhalfTSCS30KHZquarterTSCS15KHZoneEighthTConstraints = per.SizeRange(1, common.MaxPO_PerPF)

var pCCHConfigFirstPDCCHMonitoringOccasionOfPOSCS120KHZhalfTSCS60KHZquarterTSCS30KHZoneEighthTSCS15KHZoneSixteenthTConstraints = per.SizeRange(1, common.MaxPO_PerPF)

var pCCHConfigFirstPDCCHMonitoringOccasionOfPOSCS480KHZoneTSCS120KHZquarterTSCS60KHZoneEighthTSCS30KHZoneSixteenthTConstraints = per.SizeRange(1, common.MaxPO_PerPF)

var pCCHConfigFirstPDCCHMonitoringOccasionOfPOSCS480KHZhalfTSCS120KHZoneEighthTSCS60KHZoneSixteenthTConstraints = per.SizeRange(1, common.MaxPO_PerPF)

var pCCHConfigFirstPDCCHMonitoringOccasionOfPOSCS480KHZquarterTSCS120KHZoneSixteenthTConstraints = per.SizeRange(1, common.MaxPO_PerPF)

var pCCHConfigExtFirstPDCCHMonitoringOccasionOfPOV1710SCS480KHZoneEighthTConstraints = per.SizeRange(1, common.MaxPO_PerPF)

var pCCHConfigExtFirstPDCCHMonitoringOccasionOfPOV1710SCS480KHZoneSixteenthTConstraints = per.SizeRange(1, common.MaxPO_PerPF)

var pCCHConfigExtPagingAdaptationR19Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "pagingAdapt-NS-r19"},
		{Name: "pagingAdaptNAndPagingFrameOffset-r19"},
		{Name: "pagingAdaptFirstPDCCH-MonitoringOccasionOfPO-r19", Optional: true},
	},
}

const (
	PCCH_Config_Ext_PagingAdaptation_r19_PagingAdapt_NS_r19_Eight = 0
	PCCH_Config_Ext_PagingAdaptation_r19_PagingAdapt_NS_r19_Four  = 1
	PCCH_Config_Ext_PagingAdaptation_r19_PagingAdapt_NS_r19_Two   = 2
	PCCH_Config_Ext_PagingAdaptation_r19_PagingAdapt_NS_r19_One   = 3
)

var pCCHConfigExtPagingAdaptationR19PagingAdaptNSR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PCCH_Config_Ext_PagingAdaptation_r19_PagingAdapt_NS_r19_Eight, PCCH_Config_Ext_PagingAdaptation_r19_PagingAdapt_NS_r19_Four, PCCH_Config_Ext_PagingAdaptation_r19_PagingAdapt_NS_r19_Two, PCCH_Config_Ext_PagingAdaptation_r19_PagingAdapt_NS_r19_One},
}

var pCCHConfigExtPagingAdaptationR19PagingAdaptNAndPagingFrameOffsetR19Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "oneT"},
		{Name: "halfT"},
		{Name: "quarterT"},
		{Name: "oneEighthT"},
		{Name: "oneSixteenthT"},
		{Name: "oneThirtySecondT"},
	},
}

const (
	PCCH_Config_Ext_PagingAdaptation_r19_PagingAdaptNAndPagingFrameOffset_r19_OneT             = 0
	PCCH_Config_Ext_PagingAdaptation_r19_PagingAdaptNAndPagingFrameOffset_r19_HalfT            = 1
	PCCH_Config_Ext_PagingAdaptation_r19_PagingAdaptNAndPagingFrameOffset_r19_QuarterT         = 2
	PCCH_Config_Ext_PagingAdaptation_r19_PagingAdaptNAndPagingFrameOffset_r19_OneEighthT       = 3
	PCCH_Config_Ext_PagingAdaptation_r19_PagingAdaptNAndPagingFrameOffset_r19_OneSixteenthT    = 4
	PCCH_Config_Ext_PagingAdaptation_r19_PagingAdaptNAndPagingFrameOffset_r19_OneThirtySecondT = 5
)

var pCCHConfigExtPagingAdaptationR19PagingAdaptFirstPDCCHMonitoringOccasionOfPOR19Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "sCS15KHZoneT"},
		{Name: "sCS30KHZoneT-SCS15KHZhalfT"},
		{Name: "sCS60KHZoneT-SCS30KHZhalfT-SCS15KHZquarterT"},
		{Name: "sCS120KHZoneT-SCS60KHZhalfT-SCS30KHZquarterT-SCS15KHZoneEighthT"},
		{Name: "sCS120KHZhalfT-SCS60KHZquarterT-SCS30KHZoneEighthT-SCS15KHZoneSixteenthT"},
		{Name: "sCS480KHZoneT-SCS120KHZquarterT-SCS60KHZoneEighthT-SCS30KHZoneSixteenthT-SCS15KHZoneThirtySecondT"},
		{Name: "sCS480KHZhalfT-SCS120KHZoneEighthT-SCS60KHZoneSixteenthT-SCS30KHZoneThirtySecondT"},
		{Name: "sCS480KHZquarterT-SCS120KHZoneSixteenthT-SCS60KHZoneThirtySecondT"},
		{Name: "sCS480KHZoneEighthT-sCS120KHZoneThirtySecondT"},
		{Name: "sCS480KHZoneSixteenthT"},
		{Name: "sCS480KHZoneThirtySecondT"},
	},
}

const (
	PCCH_Config_Ext_PagingAdaptation_r19_PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19_SCS15KHZoneT                                                                                      = 0
	PCCH_Config_Ext_PagingAdaptation_r19_PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19_SCS30KHZoneT_SCS15KHZhalfT                                                                        = 1
	PCCH_Config_Ext_PagingAdaptation_r19_PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19_SCS60KHZoneT_SCS30KHZhalfT_SCS15KHZquarterT                                                       = 2
	PCCH_Config_Ext_PagingAdaptation_r19_PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19_SCS120KHZoneT_SCS60KHZhalfT_SCS30KHZquarterT_SCS15KHZoneEighthT                                   = 3
	PCCH_Config_Ext_PagingAdaptation_r19_PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19_SCS120KHZhalfT_SCS60KHZquarterT_SCS30KHZoneEighthT_SCS15KHZoneSixteenthT                          = 4
	PCCH_Config_Ext_PagingAdaptation_r19_PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19_SCS480KHZoneT_SCS120KHZquarterT_SCS60KHZoneEighthT_SCS30KHZoneSixteenthT_SCS15KHZoneThirtySecondT = 5
	PCCH_Config_Ext_PagingAdaptation_r19_PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19_SCS480KHZhalfT_SCS120KHZoneEighthT_SCS60KHZoneSixteenthT_SCS30KHZoneThirtySecondT                 = 6
	PCCH_Config_Ext_PagingAdaptation_r19_PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19_SCS480KHZquarterT_SCS120KHZoneSixteenthT_SCS60KHZoneThirtySecondT                                 = 7
	PCCH_Config_Ext_PagingAdaptation_r19_PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19_SCS480KHZoneEighthT_SCS120KHZoneThirtySecondT                                                     = 8
	PCCH_Config_Ext_PagingAdaptation_r19_PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19_SCS480KHZoneSixteenthT                                                                            = 9
	PCCH_Config_Ext_PagingAdaptation_r19_PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19_SCS480KHZoneThirtySecondT                                                                         = 10
)

var pCCHConfigExtPagingAdaptationR19PagingAdaptFirstPDCCHMonitoringOccasionOfPOR19SCS15KHZoneTConstraints = per.SizeRange(1, common.MaxNrofPO_PerPF_r19)

var pCCHConfigExtPagingAdaptationR19PagingAdaptFirstPDCCHMonitoringOccasionOfPOR19SCS30KHZoneTSCS15KHZhalfTConstraints = per.SizeRange(1, common.MaxNrofPO_PerPF_r19)

var pCCHConfigExtPagingAdaptationR19PagingAdaptFirstPDCCHMonitoringOccasionOfPOR19SCS60KHZoneTSCS30KHZhalfTSCS15KHZquarterTConstraints = per.SizeRange(1, common.MaxNrofPO_PerPF_r19)

var pCCHConfigExtPagingAdaptationR19PagingAdaptFirstPDCCHMonitoringOccasionOfPOR19SCS120KHZoneTSCS60KHZhalfTSCS30KHZquarterTSCS15KHZoneEighthTConstraints = per.SizeRange(1, common.MaxNrofPO_PerPF_r19)

var pCCHConfigExtPagingAdaptationR19PagingAdaptFirstPDCCHMonitoringOccasionOfPOR19SCS120KHZhalfTSCS60KHZquarterTSCS30KHZoneEighthTSCS15KHZoneSixteenthTConstraints = per.SizeRange(1, common.MaxNrofPO_PerPF_r19)

var pCCHConfigExtPagingAdaptationR19PagingAdaptFirstPDCCHMonitoringOccasionOfPOR19SCS480KHZoneTSCS120KHZquarterTSCS60KHZoneEighthTSCS30KHZoneSixteenthTSCS15KHZoneThirtySecondTConstraints = per.SizeRange(1, common.MaxNrofPO_PerPF_r19)

var pCCHConfigExtPagingAdaptationR19PagingAdaptFirstPDCCHMonitoringOccasionOfPOR19SCS480KHZhalfTSCS120KHZoneEighthTSCS60KHZoneSixteenthTSCS30KHZoneThirtySecondTConstraints = per.SizeRange(1, common.MaxNrofPO_PerPF_r19)

var pCCHConfigExtPagingAdaptationR19PagingAdaptFirstPDCCHMonitoringOccasionOfPOR19SCS480KHZquarterTSCS120KHZoneSixteenthTSCS60KHZoneThirtySecondTConstraints = per.SizeRange(1, common.MaxNrofPO_PerPF_r19)

var pCCHConfigExtPagingAdaptationR19PagingAdaptFirstPDCCHMonitoringOccasionOfPOR19SCS480KHZoneEighthTSCS120KHZoneThirtySecondTConstraints = per.SizeRange(1, common.MaxNrofPO_PerPF_r19)

var pCCHConfigExtPagingAdaptationR19PagingAdaptFirstPDCCHMonitoringOccasionOfPOR19SCS480KHZoneSixteenthTConstraints = per.SizeRange(1, common.MaxNrofPO_PerPF_r19)

var pCCHConfigExtPagingAdaptationR19PagingAdaptFirstPDCCHMonitoringOccasionOfPOR19SCS480KHZoneThirtySecondTConstraints = per.SizeRange(1, common.MaxNrofPO_PerPF_r19)

type PCCH_Config struct {
	DefaultPagingCycle    PagingCycle
	NAndPagingFrameOffset struct {
		Choice        int
		OneT          *struct{}
		HalfT         *int64
		QuarterT      *int64
		OneEighthT    *int64
		OneSixteenthT *int64
	}
	Ns                                int64
	FirstPDCCH_MonitoringOccasionOfPO *struct {
		Choice                                                                   int
		SCS15KHZoneT                                                             *[]int64
		SCS30KHZoneT_SCS15KHZhalfT                                               *[]int64
		SCS60KHZoneT_SCS30KHZhalfT_SCS15KHZquarterT                              *[]int64
		SCS120KHZoneT_SCS60KHZhalfT_SCS30KHZquarterT_SCS15KHZoneEighthT          *[]int64
		SCS120KHZhalfT_SCS60KHZquarterT_SCS30KHZoneEighthT_SCS15KHZoneSixteenthT *[]int64
		SCS480KHZoneT_SCS120KHZquarterT_SCS60KHZoneEighthT_SCS30KHZoneSixteenthT *[]int64
		SCS480KHZhalfT_SCS120KHZoneEighthT_SCS60KHZoneSixteenthT                 *[]int64
		SCS480KHZquarterT_SCS120KHZoneSixteenthT                                 *[]int64
	}
	NrofPDCCH_MonitoringOccasionPerSSB_InPO_r16 *int64
	RanPagingInIdlePO_r17                       *int64
	FirstPDCCH_MonitoringOccasionOfPO_v1710     *struct {
		Choice                 int
		SCS480KHZoneEighthT    *[]int64
		SCS480KHZoneSixteenthT *[]int64
	}
	PagingAdaptation_r19 *struct {
		PagingAdapt_NS_r19                   int64
		PagingAdaptNAndPagingFrameOffset_r19 struct {
			Choice           int
			OneT             *struct{}
			HalfT            *int64
			QuarterT         *int64
			OneEighthT       *int64
			OneSixteenthT    *int64
			OneThirtySecondT *int64
		}
		PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19 *struct {
			Choice                                                                                            int
			SCS15KHZoneT                                                                                      *[]int64
			SCS30KHZoneT_SCS15KHZhalfT                                                                        *[]int64
			SCS60KHZoneT_SCS30KHZhalfT_SCS15KHZquarterT                                                       *[]int64
			SCS120KHZoneT_SCS60KHZhalfT_SCS30KHZquarterT_SCS15KHZoneEighthT                                   *[]int64
			SCS120KHZhalfT_SCS60KHZquarterT_SCS30KHZoneEighthT_SCS15KHZoneSixteenthT                          *[]int64
			SCS480KHZoneT_SCS120KHZquarterT_SCS60KHZoneEighthT_SCS30KHZoneSixteenthT_SCS15KHZoneThirtySecondT *[]int64
			SCS480KHZhalfT_SCS120KHZoneEighthT_SCS60KHZoneSixteenthT_SCS30KHZoneThirtySecondT                 *[]int64
			SCS480KHZquarterT_SCS120KHZoneSixteenthT_SCS60KHZoneThirtySecondT                                 *[]int64
			SCS480KHZoneEighthT_SCS120KHZoneThirtySecondT                                                     *[]int64
			SCS480KHZoneSixteenthT                                                                            *[]int64
			SCS480KHZoneThirtySecondT                                                                         *[]int64
		}
	}
}

func (ie *PCCH_Config) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(pCCHConfigConstraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.NrofPDCCH_MonitoringOccasionPerSSB_InPO_r16 != nil
	hasExtGroup1 := ie.RanPagingInIdlePO_r17 != nil || ie.FirstPDCCH_MonitoringOccasionOfPO_v1710 != nil
	hasExtGroup2 := ie.PagingAdaptation_r19 != nil
	hasExtensions := hasExtGroup0 || hasExtGroup1 || hasExtGroup2

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.FirstPDCCH_MonitoringOccasionOfPO != nil}); err != nil {
		return err
	}

	// 3. defaultPagingCycle: ref
	{
		if err := ie.DefaultPagingCycle.Encode(e); err != nil {
			return err
		}
	}

	// 4. nAndPagingFrameOffset: choice
	{
		choiceEnc := e.NewChoiceEncoder(pCCH_ConfigNAndPagingFrameOffsetConstraints)
		if err := choiceEnc.EncodeChoice(int64(ie.NAndPagingFrameOffset.Choice), false, nil); err != nil {
			return err
		}
		switch ie.NAndPagingFrameOffset.Choice {
		case PCCH_Config_NAndPagingFrameOffset_OneT:
			if err := e.EncodeNull(); err != nil {
				return err
			}
		case PCCH_Config_NAndPagingFrameOffset_HalfT:
			if err := e.EncodeInteger((*ie.NAndPagingFrameOffset.HalfT), per.Constrained(0, 1)); err != nil {
				return err
			}
		case PCCH_Config_NAndPagingFrameOffset_QuarterT:
			if err := e.EncodeInteger((*ie.NAndPagingFrameOffset.QuarterT), per.Constrained(0, 3)); err != nil {
				return err
			}
		case PCCH_Config_NAndPagingFrameOffset_OneEighthT:
			if err := e.EncodeInteger((*ie.NAndPagingFrameOffset.OneEighthT), per.Constrained(0, 7)); err != nil {
				return err
			}
		case PCCH_Config_NAndPagingFrameOffset_OneSixteenthT:
			if err := e.EncodeInteger((*ie.NAndPagingFrameOffset.OneSixteenthT), per.Constrained(0, 15)); err != nil {
				return err
			}
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.NAndPagingFrameOffset.Choice), Constraint: "index out of range"}
		}
	}

	// 5. ns: enumerated
	{
		if err := e.EncodeEnumerated(ie.Ns, pCCHConfigNsConstraints); err != nil {
			return err
		}
	}

	// 6. firstPDCCH-MonitoringOccasionOfPO: choice
	{
		if ie.FirstPDCCH_MonitoringOccasionOfPO != nil {
			choiceEnc := e.NewChoiceEncoder(pCCH_ConfigFirstPDCCHMonitoringOccasionOfPOConstraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.FirstPDCCH_MonitoringOccasionOfPO).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.FirstPDCCH_MonitoringOccasionOfPO).Choice {
			case PCCH_Config_FirstPDCCH_MonitoringOccasionOfPO_SCS15KHZoneT:
				seqOf := e.NewSequenceOfEncoder(pCCHConfigFirstPDCCHMonitoringOccasionOfPOSCS15KHZoneTConstraints)
				if err := seqOf.EncodeLength(int64(len((*(*ie.FirstPDCCH_MonitoringOccasionOfPO).SCS15KHZoneT)))); err != nil {
					return err
				}
				for i := range *(*ie.FirstPDCCH_MonitoringOccasionOfPO).SCS15KHZoneT {
					if err := e.EncodeInteger((*(*ie.FirstPDCCH_MonitoringOccasionOfPO).SCS15KHZoneT)[i], per.Constrained(0, 139)); err != nil {
						return err
					}
				}
			case PCCH_Config_FirstPDCCH_MonitoringOccasionOfPO_SCS30KHZoneT_SCS15KHZhalfT:
				seqOf := e.NewSequenceOfEncoder(pCCHConfigFirstPDCCHMonitoringOccasionOfPOSCS30KHZoneTSCS15KHZhalfTConstraints)
				if err := seqOf.EncodeLength(int64(len((*(*ie.FirstPDCCH_MonitoringOccasionOfPO).SCS30KHZoneT_SCS15KHZhalfT)))); err != nil {
					return err
				}
				for i := range *(*ie.FirstPDCCH_MonitoringOccasionOfPO).SCS30KHZoneT_SCS15KHZhalfT {
					if err := e.EncodeInteger((*(*ie.FirstPDCCH_MonitoringOccasionOfPO).SCS30KHZoneT_SCS15KHZhalfT)[i], per.Constrained(0, 279)); err != nil {
						return err
					}
				}
			case PCCH_Config_FirstPDCCH_MonitoringOccasionOfPO_SCS60KHZoneT_SCS30KHZhalfT_SCS15KHZquarterT:
				seqOf := e.NewSequenceOfEncoder(pCCHConfigFirstPDCCHMonitoringOccasionOfPOSCS60KHZoneTSCS30KHZhalfTSCS15KHZquarterTConstraints)
				if err := seqOf.EncodeLength(int64(len((*(*ie.FirstPDCCH_MonitoringOccasionOfPO).SCS60KHZoneT_SCS30KHZhalfT_SCS15KHZquarterT)))); err != nil {
					return err
				}
				for i := range *(*ie.FirstPDCCH_MonitoringOccasionOfPO).SCS60KHZoneT_SCS30KHZhalfT_SCS15KHZquarterT {
					if err := e.EncodeInteger((*(*ie.FirstPDCCH_MonitoringOccasionOfPO).SCS60KHZoneT_SCS30KHZhalfT_SCS15KHZquarterT)[i], per.Constrained(0, 559)); err != nil {
						return err
					}
				}
			case PCCH_Config_FirstPDCCH_MonitoringOccasionOfPO_SCS120KHZoneT_SCS60KHZhalfT_SCS30KHZquarterT_SCS15KHZoneEighthT:
				seqOf := e.NewSequenceOfEncoder(pCCHConfigFirstPDCCHMonitoringOccasionOfPOSCS120KHZoneTSCS60KHZhalfTSCS30KHZquarterTSCS15KHZoneEighthTConstraints)
				if err := seqOf.EncodeLength(int64(len((*(*ie.FirstPDCCH_MonitoringOccasionOfPO).SCS120KHZoneT_SCS60KHZhalfT_SCS30KHZquarterT_SCS15KHZoneEighthT)))); err != nil {
					return err
				}
				for i := range *(*ie.FirstPDCCH_MonitoringOccasionOfPO).SCS120KHZoneT_SCS60KHZhalfT_SCS30KHZquarterT_SCS15KHZoneEighthT {
					if err := e.EncodeInteger((*(*ie.FirstPDCCH_MonitoringOccasionOfPO).SCS120KHZoneT_SCS60KHZhalfT_SCS30KHZquarterT_SCS15KHZoneEighthT)[i], per.Constrained(0, 1119)); err != nil {
						return err
					}
				}
			case PCCH_Config_FirstPDCCH_MonitoringOccasionOfPO_SCS120KHZhalfT_SCS60KHZquarterT_SCS30KHZoneEighthT_SCS15KHZoneSixteenthT:
				seqOf := e.NewSequenceOfEncoder(pCCHConfigFirstPDCCHMonitoringOccasionOfPOSCS120KHZhalfTSCS60KHZquarterTSCS30KHZoneEighthTSCS15KHZoneSixteenthTConstraints)
				if err := seqOf.EncodeLength(int64(len((*(*ie.FirstPDCCH_MonitoringOccasionOfPO).SCS120KHZhalfT_SCS60KHZquarterT_SCS30KHZoneEighthT_SCS15KHZoneSixteenthT)))); err != nil {
					return err
				}
				for i := range *(*ie.FirstPDCCH_MonitoringOccasionOfPO).SCS120KHZhalfT_SCS60KHZquarterT_SCS30KHZoneEighthT_SCS15KHZoneSixteenthT {
					if err := e.EncodeInteger((*(*ie.FirstPDCCH_MonitoringOccasionOfPO).SCS120KHZhalfT_SCS60KHZquarterT_SCS30KHZoneEighthT_SCS15KHZoneSixteenthT)[i], per.Constrained(0, 2239)); err != nil {
						return err
					}
				}
			case PCCH_Config_FirstPDCCH_MonitoringOccasionOfPO_SCS480KHZoneT_SCS120KHZquarterT_SCS60KHZoneEighthT_SCS30KHZoneSixteenthT:
				seqOf := e.NewSequenceOfEncoder(pCCHConfigFirstPDCCHMonitoringOccasionOfPOSCS480KHZoneTSCS120KHZquarterTSCS60KHZoneEighthTSCS30KHZoneSixteenthTConstraints)
				if err := seqOf.EncodeLength(int64(len((*(*ie.FirstPDCCH_MonitoringOccasionOfPO).SCS480KHZoneT_SCS120KHZquarterT_SCS60KHZoneEighthT_SCS30KHZoneSixteenthT)))); err != nil {
					return err
				}
				for i := range *(*ie.FirstPDCCH_MonitoringOccasionOfPO).SCS480KHZoneT_SCS120KHZquarterT_SCS60KHZoneEighthT_SCS30KHZoneSixteenthT {
					if err := e.EncodeInteger((*(*ie.FirstPDCCH_MonitoringOccasionOfPO).SCS480KHZoneT_SCS120KHZquarterT_SCS60KHZoneEighthT_SCS30KHZoneSixteenthT)[i], per.Constrained(0, 4479)); err != nil {
						return err
					}
				}
			case PCCH_Config_FirstPDCCH_MonitoringOccasionOfPO_SCS480KHZhalfT_SCS120KHZoneEighthT_SCS60KHZoneSixteenthT:
				seqOf := e.NewSequenceOfEncoder(pCCHConfigFirstPDCCHMonitoringOccasionOfPOSCS480KHZhalfTSCS120KHZoneEighthTSCS60KHZoneSixteenthTConstraints)
				if err := seqOf.EncodeLength(int64(len((*(*ie.FirstPDCCH_MonitoringOccasionOfPO).SCS480KHZhalfT_SCS120KHZoneEighthT_SCS60KHZoneSixteenthT)))); err != nil {
					return err
				}
				for i := range *(*ie.FirstPDCCH_MonitoringOccasionOfPO).SCS480KHZhalfT_SCS120KHZoneEighthT_SCS60KHZoneSixteenthT {
					if err := e.EncodeInteger((*(*ie.FirstPDCCH_MonitoringOccasionOfPO).SCS480KHZhalfT_SCS120KHZoneEighthT_SCS60KHZoneSixteenthT)[i], per.Constrained(0, 8959)); err != nil {
						return err
					}
				}
			case PCCH_Config_FirstPDCCH_MonitoringOccasionOfPO_SCS480KHZquarterT_SCS120KHZoneSixteenthT:
				seqOf := e.NewSequenceOfEncoder(pCCHConfigFirstPDCCHMonitoringOccasionOfPOSCS480KHZquarterTSCS120KHZoneSixteenthTConstraints)
				if err := seqOf.EncodeLength(int64(len((*(*ie.FirstPDCCH_MonitoringOccasionOfPO).SCS480KHZquarterT_SCS120KHZoneSixteenthT)))); err != nil {
					return err
				}
				for i := range *(*ie.FirstPDCCH_MonitoringOccasionOfPO).SCS480KHZquarterT_SCS120KHZoneSixteenthT {
					if err := e.EncodeInteger((*(*ie.FirstPDCCH_MonitoringOccasionOfPO).SCS480KHZquarterT_SCS120KHZoneSixteenthT)[i], per.Constrained(0, 17919)); err != nil {
						return err
					}
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.FirstPDCCH_MonitoringOccasionOfPO).Choice), Constraint: "index out of range"}
			}
		}
	}

	if hasExtensions {
		extPresence := []bool{hasExtGroup0, hasExtGroup1, hasExtGroup2}
		var extValues [][]byte

		// Extension group 0:
		if hasExtGroup0 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "nrofPDCCH-MonitoringOccasionPerSSB-InPO-r16", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.NrofPDCCH_MonitoringOccasionPerSSB_InPO_r16 != nil}); err != nil {
				return err
			}

			if ie.NrofPDCCH_MonitoringOccasionPerSSB_InPO_r16 != nil {
				if err := ex.EncodeInteger(*ie.NrofPDCCH_MonitoringOccasionPerSSB_InPO_r16, pCCHConfigNrofPDCCHMonitoringOccasionPerSSBInPOR16Constraints); err != nil {
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
					{Name: "ranPagingInIdlePO-r17", Optional: true},
					{Name: "firstPDCCH-MonitoringOccasionOfPO-v1710", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.RanPagingInIdlePO_r17 != nil, ie.FirstPDCCH_MonitoringOccasionOfPO_v1710 != nil}); err != nil {
				return err
			}

			if ie.RanPagingInIdlePO_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.RanPagingInIdlePO_r17, pCCHConfigExtRanPagingInIdlePOR17Constraints); err != nil {
					return err
				}
			}

			if ie.FirstPDCCH_MonitoringOccasionOfPO_v1710 != nil {
				choiceEnc := ex.NewChoiceEncoder(pCCHConfigExtFirstPDCCHMonitoringOccasionOfPOV1710Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.FirstPDCCH_MonitoringOccasionOfPO_v1710).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.FirstPDCCH_MonitoringOccasionOfPO_v1710).Choice {
				case PCCH_Config_Ext_FirstPDCCH_MonitoringOccasionOfPO_v1710_SCS480KHZoneEighthT:
					seqOf := ex.NewSequenceOfEncoder(pCCHConfigExtFirstPDCCHMonitoringOccasionOfPOV1710SCS480KHZoneEighthTConstraints)
					if err := seqOf.EncodeLength(int64(len((*(*ie.FirstPDCCH_MonitoringOccasionOfPO_v1710).SCS480KHZoneEighthT)))); err != nil {
						return err
					}
					for i := range *(*ie.FirstPDCCH_MonitoringOccasionOfPO_v1710).SCS480KHZoneEighthT {
						if err := ex.EncodeInteger((*(*ie.FirstPDCCH_MonitoringOccasionOfPO_v1710).SCS480KHZoneEighthT)[i], per.Constrained(0, 35839)); err != nil {
							return err
						}
					}
				case PCCH_Config_Ext_FirstPDCCH_MonitoringOccasionOfPO_v1710_SCS480KHZoneSixteenthT:
					seqOf := ex.NewSequenceOfEncoder(pCCHConfigExtFirstPDCCHMonitoringOccasionOfPOV1710SCS480KHZoneSixteenthTConstraints)
					if err := seqOf.EncodeLength(int64(len((*(*ie.FirstPDCCH_MonitoringOccasionOfPO_v1710).SCS480KHZoneSixteenthT)))); err != nil {
						return err
					}
					for i := range *(*ie.FirstPDCCH_MonitoringOccasionOfPO_v1710).SCS480KHZoneSixteenthT {
						if err := ex.EncodeInteger((*(*ie.FirstPDCCH_MonitoringOccasionOfPO_v1710).SCS480KHZoneSixteenthT)[i], per.Constrained(0, 71679)); err != nil {
							return err
						}
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
					{Name: "pagingAdaptation-r19", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.PagingAdaptation_r19 != nil}); err != nil {
				return err
			}

			if ie.PagingAdaptation_r19 != nil {
				c := ie.PagingAdaptation_r19
				pCCHConfigExtPagingAdaptationR19Seq := ex.NewSequenceEncoder(pCCHConfigExtPagingAdaptationR19Constraints)
				if err := pCCHConfigExtPagingAdaptationR19Seq.EncodePreamble([]bool{c.PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19 != nil}); err != nil {
					return err
				}
				if err := ex.EncodeEnumerated(c.PagingAdapt_NS_r19, pCCHConfigExtPagingAdaptationR19PagingAdaptNSR19Constraints); err != nil {
					return err
				}
				{
					choiceEnc := ex.NewChoiceEncoder(pCCHConfigExtPagingAdaptationR19PagingAdaptNAndPagingFrameOffsetR19Constraints)
					if err := choiceEnc.EncodeChoice(int64(c.PagingAdaptNAndPagingFrameOffset_r19.Choice), false, nil); err != nil {
						return err
					}
					switch c.PagingAdaptNAndPagingFrameOffset_r19.Choice {
					case PCCH_Config_Ext_PagingAdaptation_r19_PagingAdaptNAndPagingFrameOffset_r19_OneT:
						if err := ex.EncodeNull(); err != nil {
							return err
						}
					case PCCH_Config_Ext_PagingAdaptation_r19_PagingAdaptNAndPagingFrameOffset_r19_HalfT:
						if err := ex.EncodeInteger((*c.PagingAdaptNAndPagingFrameOffset_r19.HalfT), per.Constrained(0, 1)); err != nil {
							return err
						}
					case PCCH_Config_Ext_PagingAdaptation_r19_PagingAdaptNAndPagingFrameOffset_r19_QuarterT:
						if err := ex.EncodeInteger((*c.PagingAdaptNAndPagingFrameOffset_r19.QuarterT), per.Constrained(0, 3)); err != nil {
							return err
						}
					case PCCH_Config_Ext_PagingAdaptation_r19_PagingAdaptNAndPagingFrameOffset_r19_OneEighthT:
						if err := ex.EncodeInteger((*c.PagingAdaptNAndPagingFrameOffset_r19.OneEighthT), per.Constrained(0, 7)); err != nil {
							return err
						}
					case PCCH_Config_Ext_PagingAdaptation_r19_PagingAdaptNAndPagingFrameOffset_r19_OneSixteenthT:
						if err := ex.EncodeInteger((*c.PagingAdaptNAndPagingFrameOffset_r19.OneSixteenthT), per.Constrained(0, 15)); err != nil {
							return err
						}
					case PCCH_Config_Ext_PagingAdaptation_r19_PagingAdaptNAndPagingFrameOffset_r19_OneThirtySecondT:
						if err := ex.EncodeInteger((*c.PagingAdaptNAndPagingFrameOffset_r19.OneThirtySecondT), per.Constrained(0, 31)); err != nil {
							return err
						}
					}
				}
				if c.PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19 != nil {
					choiceEnc := ex.NewChoiceEncoder(pCCHConfigExtPagingAdaptationR19PagingAdaptFirstPDCCHMonitoringOccasionOfPOR19Constraints)
					if err := choiceEnc.EncodeChoice(int64((*c.PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19).Choice), false, nil); err != nil {
						return err
					}
					switch (*c.PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19).Choice {
					case PCCH_Config_Ext_PagingAdaptation_r19_PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19_SCS15KHZoneT:
						seqOf := ex.NewSequenceOfEncoder(pCCHConfigExtPagingAdaptationR19PagingAdaptFirstPDCCHMonitoringOccasionOfPOR19SCS15KHZoneTConstraints)
						if err := seqOf.EncodeLength(int64(len((*(*c.PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19).SCS15KHZoneT)))); err != nil {
							return err
						}
						for i := range *(*c.PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19).SCS15KHZoneT {
							if err := ex.EncodeInteger((*(*c.PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19).SCS15KHZoneT)[i], per.Constrained(0, 139)); err != nil {
								return err
							}
						}
					case PCCH_Config_Ext_PagingAdaptation_r19_PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19_SCS30KHZoneT_SCS15KHZhalfT:
						seqOf := ex.NewSequenceOfEncoder(pCCHConfigExtPagingAdaptationR19PagingAdaptFirstPDCCHMonitoringOccasionOfPOR19SCS30KHZoneTSCS15KHZhalfTConstraints)
						if err := seqOf.EncodeLength(int64(len((*(*c.PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19).SCS30KHZoneT_SCS15KHZhalfT)))); err != nil {
							return err
						}
						for i := range *(*c.PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19).SCS30KHZoneT_SCS15KHZhalfT {
							if err := ex.EncodeInteger((*(*c.PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19).SCS30KHZoneT_SCS15KHZhalfT)[i], per.Constrained(0, 279)); err != nil {
								return err
							}
						}
					case PCCH_Config_Ext_PagingAdaptation_r19_PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19_SCS60KHZoneT_SCS30KHZhalfT_SCS15KHZquarterT:
						seqOf := ex.NewSequenceOfEncoder(pCCHConfigExtPagingAdaptationR19PagingAdaptFirstPDCCHMonitoringOccasionOfPOR19SCS60KHZoneTSCS30KHZhalfTSCS15KHZquarterTConstraints)
						if err := seqOf.EncodeLength(int64(len((*(*c.PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19).SCS60KHZoneT_SCS30KHZhalfT_SCS15KHZquarterT)))); err != nil {
							return err
						}
						for i := range *(*c.PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19).SCS60KHZoneT_SCS30KHZhalfT_SCS15KHZquarterT {
							if err := ex.EncodeInteger((*(*c.PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19).SCS60KHZoneT_SCS30KHZhalfT_SCS15KHZquarterT)[i], per.Constrained(0, 559)); err != nil {
								return err
							}
						}
					case PCCH_Config_Ext_PagingAdaptation_r19_PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19_SCS120KHZoneT_SCS60KHZhalfT_SCS30KHZquarterT_SCS15KHZoneEighthT:
						seqOf := ex.NewSequenceOfEncoder(pCCHConfigExtPagingAdaptationR19PagingAdaptFirstPDCCHMonitoringOccasionOfPOR19SCS120KHZoneTSCS60KHZhalfTSCS30KHZquarterTSCS15KHZoneEighthTConstraints)
						if err := seqOf.EncodeLength(int64(len((*(*c.PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19).SCS120KHZoneT_SCS60KHZhalfT_SCS30KHZquarterT_SCS15KHZoneEighthT)))); err != nil {
							return err
						}
						for i := range *(*c.PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19).SCS120KHZoneT_SCS60KHZhalfT_SCS30KHZquarterT_SCS15KHZoneEighthT {
							if err := ex.EncodeInteger((*(*c.PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19).SCS120KHZoneT_SCS60KHZhalfT_SCS30KHZquarterT_SCS15KHZoneEighthT)[i], per.Constrained(0, 1119)); err != nil {
								return err
							}
						}
					case PCCH_Config_Ext_PagingAdaptation_r19_PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19_SCS120KHZhalfT_SCS60KHZquarterT_SCS30KHZoneEighthT_SCS15KHZoneSixteenthT:
						seqOf := ex.NewSequenceOfEncoder(pCCHConfigExtPagingAdaptationR19PagingAdaptFirstPDCCHMonitoringOccasionOfPOR19SCS120KHZhalfTSCS60KHZquarterTSCS30KHZoneEighthTSCS15KHZoneSixteenthTConstraints)
						if err := seqOf.EncodeLength(int64(len((*(*c.PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19).SCS120KHZhalfT_SCS60KHZquarterT_SCS30KHZoneEighthT_SCS15KHZoneSixteenthT)))); err != nil {
							return err
						}
						for i := range *(*c.PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19).SCS120KHZhalfT_SCS60KHZquarterT_SCS30KHZoneEighthT_SCS15KHZoneSixteenthT {
							if err := ex.EncodeInteger((*(*c.PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19).SCS120KHZhalfT_SCS60KHZquarterT_SCS30KHZoneEighthT_SCS15KHZoneSixteenthT)[i], per.Constrained(0, 2239)); err != nil {
								return err
							}
						}
					case PCCH_Config_Ext_PagingAdaptation_r19_PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19_SCS480KHZoneT_SCS120KHZquarterT_SCS60KHZoneEighthT_SCS30KHZoneSixteenthT_SCS15KHZoneThirtySecondT:
						seqOf := ex.NewSequenceOfEncoder(pCCHConfigExtPagingAdaptationR19PagingAdaptFirstPDCCHMonitoringOccasionOfPOR19SCS480KHZoneTSCS120KHZquarterTSCS60KHZoneEighthTSCS30KHZoneSixteenthTSCS15KHZoneThirtySecondTConstraints)
						if err := seqOf.EncodeLength(int64(len((*(*c.PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19).SCS480KHZoneT_SCS120KHZquarterT_SCS60KHZoneEighthT_SCS30KHZoneSixteenthT_SCS15KHZoneThirtySecondT)))); err != nil {
							return err
						}
						for i := range *(*c.PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19).SCS480KHZoneT_SCS120KHZquarterT_SCS60KHZoneEighthT_SCS30KHZoneSixteenthT_SCS15KHZoneThirtySecondT {
							if err := ex.EncodeInteger((*(*c.PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19).SCS480KHZoneT_SCS120KHZquarterT_SCS60KHZoneEighthT_SCS30KHZoneSixteenthT_SCS15KHZoneThirtySecondT)[i], per.Constrained(0, 4479)); err != nil {
								return err
							}
						}
					case PCCH_Config_Ext_PagingAdaptation_r19_PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19_SCS480KHZhalfT_SCS120KHZoneEighthT_SCS60KHZoneSixteenthT_SCS30KHZoneThirtySecondT:
						seqOf := ex.NewSequenceOfEncoder(pCCHConfigExtPagingAdaptationR19PagingAdaptFirstPDCCHMonitoringOccasionOfPOR19SCS480KHZhalfTSCS120KHZoneEighthTSCS60KHZoneSixteenthTSCS30KHZoneThirtySecondTConstraints)
						if err := seqOf.EncodeLength(int64(len((*(*c.PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19).SCS480KHZhalfT_SCS120KHZoneEighthT_SCS60KHZoneSixteenthT_SCS30KHZoneThirtySecondT)))); err != nil {
							return err
						}
						for i := range *(*c.PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19).SCS480KHZhalfT_SCS120KHZoneEighthT_SCS60KHZoneSixteenthT_SCS30KHZoneThirtySecondT {
							if err := ex.EncodeInteger((*(*c.PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19).SCS480KHZhalfT_SCS120KHZoneEighthT_SCS60KHZoneSixteenthT_SCS30KHZoneThirtySecondT)[i], per.Constrained(0, 8959)); err != nil {
								return err
							}
						}
					case PCCH_Config_Ext_PagingAdaptation_r19_PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19_SCS480KHZquarterT_SCS120KHZoneSixteenthT_SCS60KHZoneThirtySecondT:
						seqOf := ex.NewSequenceOfEncoder(pCCHConfigExtPagingAdaptationR19PagingAdaptFirstPDCCHMonitoringOccasionOfPOR19SCS480KHZquarterTSCS120KHZoneSixteenthTSCS60KHZoneThirtySecondTConstraints)
						if err := seqOf.EncodeLength(int64(len((*(*c.PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19).SCS480KHZquarterT_SCS120KHZoneSixteenthT_SCS60KHZoneThirtySecondT)))); err != nil {
							return err
						}
						for i := range *(*c.PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19).SCS480KHZquarterT_SCS120KHZoneSixteenthT_SCS60KHZoneThirtySecondT {
							if err := ex.EncodeInteger((*(*c.PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19).SCS480KHZquarterT_SCS120KHZoneSixteenthT_SCS60KHZoneThirtySecondT)[i], per.Constrained(0, 17919)); err != nil {
								return err
							}
						}
					case PCCH_Config_Ext_PagingAdaptation_r19_PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19_SCS480KHZoneEighthT_SCS120KHZoneThirtySecondT:
						seqOf := ex.NewSequenceOfEncoder(pCCHConfigExtPagingAdaptationR19PagingAdaptFirstPDCCHMonitoringOccasionOfPOR19SCS480KHZoneEighthTSCS120KHZoneThirtySecondTConstraints)
						if err := seqOf.EncodeLength(int64(len((*(*c.PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19).SCS480KHZoneEighthT_SCS120KHZoneThirtySecondT)))); err != nil {
							return err
						}
						for i := range *(*c.PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19).SCS480KHZoneEighthT_SCS120KHZoneThirtySecondT {
							if err := ex.EncodeInteger((*(*c.PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19).SCS480KHZoneEighthT_SCS120KHZoneThirtySecondT)[i], per.Constrained(0, 35839)); err != nil {
								return err
							}
						}
					case PCCH_Config_Ext_PagingAdaptation_r19_PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19_SCS480KHZoneSixteenthT:
						seqOf := ex.NewSequenceOfEncoder(pCCHConfigExtPagingAdaptationR19PagingAdaptFirstPDCCHMonitoringOccasionOfPOR19SCS480KHZoneSixteenthTConstraints)
						if err := seqOf.EncodeLength(int64(len((*(*c.PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19).SCS480KHZoneSixteenthT)))); err != nil {
							return err
						}
						for i := range *(*c.PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19).SCS480KHZoneSixteenthT {
							if err := ex.EncodeInteger((*(*c.PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19).SCS480KHZoneSixteenthT)[i], per.Constrained(0, 71679)); err != nil {
								return err
							}
						}
					case PCCH_Config_Ext_PagingAdaptation_r19_PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19_SCS480KHZoneThirtySecondT:
						seqOf := ex.NewSequenceOfEncoder(pCCHConfigExtPagingAdaptationR19PagingAdaptFirstPDCCHMonitoringOccasionOfPOR19SCS480KHZoneThirtySecondTConstraints)
						if err := seqOf.EncodeLength(int64(len((*(*c.PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19).SCS480KHZoneThirtySecondT)))); err != nil {
							return err
						}
						for i := range *(*c.PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19).SCS480KHZoneThirtySecondT {
							if err := ex.EncodeInteger((*(*c.PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19).SCS480KHZoneThirtySecondT)[i], per.Constrained(0, 143359)); err != nil {
								return err
							}
						}
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

func (ie *PCCH_Config) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(pCCHConfigConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. defaultPagingCycle: ref
	{
		if err := ie.DefaultPagingCycle.Decode(d); err != nil {
			return err
		}
	}

	// 4. nAndPagingFrameOffset: choice
	{
		choiceDec := d.NewChoiceDecoder(pCCH_ConfigNAndPagingFrameOffsetConstraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.NAndPagingFrameOffset.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case PCCH_Config_NAndPagingFrameOffset_OneT:
			ie.NAndPagingFrameOffset.OneT = &struct{}{}
			if err := d.DecodeNull(); err != nil {
				return err
			}
		case PCCH_Config_NAndPagingFrameOffset_HalfT:
			ie.NAndPagingFrameOffset.HalfT = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 1))
			if err != nil {
				return err
			}
			(*ie.NAndPagingFrameOffset.HalfT) = v
		case PCCH_Config_NAndPagingFrameOffset_QuarterT:
			ie.NAndPagingFrameOffset.QuarterT = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 3))
			if err != nil {
				return err
			}
			(*ie.NAndPagingFrameOffset.QuarterT) = v
		case PCCH_Config_NAndPagingFrameOffset_OneEighthT:
			ie.NAndPagingFrameOffset.OneEighthT = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 7))
			if err != nil {
				return err
			}
			(*ie.NAndPagingFrameOffset.OneEighthT) = v
		case PCCH_Config_NAndPagingFrameOffset_OneSixteenthT:
			ie.NAndPagingFrameOffset.OneSixteenthT = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 15))
			if err != nil {
				return err
			}
			(*ie.NAndPagingFrameOffset.OneSixteenthT) = v
		}
	}

	// 5. ns: enumerated
	{
		v2, err := d.DecodeEnumerated(pCCHConfigNsConstraints)
		if err != nil {
			return err
		}
		ie.Ns = v2
	}

	// 6. firstPDCCH-MonitoringOccasionOfPO: choice
	{
		if seq.IsComponentPresent(3) {
			ie.FirstPDCCH_MonitoringOccasionOfPO = &struct {
				Choice                                                                   int
				SCS15KHZoneT                                                             *[]int64
				SCS30KHZoneT_SCS15KHZhalfT                                               *[]int64
				SCS60KHZoneT_SCS30KHZhalfT_SCS15KHZquarterT                              *[]int64
				SCS120KHZoneT_SCS60KHZhalfT_SCS30KHZquarterT_SCS15KHZoneEighthT          *[]int64
				SCS120KHZhalfT_SCS60KHZquarterT_SCS30KHZoneEighthT_SCS15KHZoneSixteenthT *[]int64
				SCS480KHZoneT_SCS120KHZquarterT_SCS60KHZoneEighthT_SCS30KHZoneSixteenthT *[]int64
				SCS480KHZhalfT_SCS120KHZoneEighthT_SCS60KHZoneSixteenthT                 *[]int64
				SCS480KHZquarterT_SCS120KHZoneSixteenthT                                 *[]int64
			}{}
			choiceDec := d.NewChoiceDecoder(pCCH_ConfigFirstPDCCHMonitoringOccasionOfPOConstraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.FirstPDCCH_MonitoringOccasionOfPO).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case PCCH_Config_FirstPDCCH_MonitoringOccasionOfPO_SCS15KHZoneT:
				(*ie.FirstPDCCH_MonitoringOccasionOfPO).SCS15KHZoneT = new([]int64)
				seqOf := d.NewSequenceOfDecoder(pCCHConfigFirstPDCCHMonitoringOccasionOfPOSCS15KHZoneTConstraints)
				n, err := seqOf.DecodeLength()
				if err != nil {
					return err
				}
				(*(*ie.FirstPDCCH_MonitoringOccasionOfPO).SCS15KHZoneT) = make([]int64, n)
				for i := int64(0); i < n; i++ {
					v, err := d.DecodeInteger(per.Constrained(0, 139))
					if err != nil {
						return err
					}
					(*(*ie.FirstPDCCH_MonitoringOccasionOfPO).SCS15KHZoneT)[i] = v
				}
			case PCCH_Config_FirstPDCCH_MonitoringOccasionOfPO_SCS30KHZoneT_SCS15KHZhalfT:
				(*ie.FirstPDCCH_MonitoringOccasionOfPO).SCS30KHZoneT_SCS15KHZhalfT = new([]int64)
				seqOf := d.NewSequenceOfDecoder(pCCHConfigFirstPDCCHMonitoringOccasionOfPOSCS30KHZoneTSCS15KHZhalfTConstraints)
				n, err := seqOf.DecodeLength()
				if err != nil {
					return err
				}
				(*(*ie.FirstPDCCH_MonitoringOccasionOfPO).SCS30KHZoneT_SCS15KHZhalfT) = make([]int64, n)
				for i := int64(0); i < n; i++ {
					v, err := d.DecodeInteger(per.Constrained(0, 279))
					if err != nil {
						return err
					}
					(*(*ie.FirstPDCCH_MonitoringOccasionOfPO).SCS30KHZoneT_SCS15KHZhalfT)[i] = v
				}
			case PCCH_Config_FirstPDCCH_MonitoringOccasionOfPO_SCS60KHZoneT_SCS30KHZhalfT_SCS15KHZquarterT:
				(*ie.FirstPDCCH_MonitoringOccasionOfPO).SCS60KHZoneT_SCS30KHZhalfT_SCS15KHZquarterT = new([]int64)
				seqOf := d.NewSequenceOfDecoder(pCCHConfigFirstPDCCHMonitoringOccasionOfPOSCS60KHZoneTSCS30KHZhalfTSCS15KHZquarterTConstraints)
				n, err := seqOf.DecodeLength()
				if err != nil {
					return err
				}
				(*(*ie.FirstPDCCH_MonitoringOccasionOfPO).SCS60KHZoneT_SCS30KHZhalfT_SCS15KHZquarterT) = make([]int64, n)
				for i := int64(0); i < n; i++ {
					v, err := d.DecodeInteger(per.Constrained(0, 559))
					if err != nil {
						return err
					}
					(*(*ie.FirstPDCCH_MonitoringOccasionOfPO).SCS60KHZoneT_SCS30KHZhalfT_SCS15KHZquarterT)[i] = v
				}
			case PCCH_Config_FirstPDCCH_MonitoringOccasionOfPO_SCS120KHZoneT_SCS60KHZhalfT_SCS30KHZquarterT_SCS15KHZoneEighthT:
				(*ie.FirstPDCCH_MonitoringOccasionOfPO).SCS120KHZoneT_SCS60KHZhalfT_SCS30KHZquarterT_SCS15KHZoneEighthT = new([]int64)
				seqOf := d.NewSequenceOfDecoder(pCCHConfigFirstPDCCHMonitoringOccasionOfPOSCS120KHZoneTSCS60KHZhalfTSCS30KHZquarterTSCS15KHZoneEighthTConstraints)
				n, err := seqOf.DecodeLength()
				if err != nil {
					return err
				}
				(*(*ie.FirstPDCCH_MonitoringOccasionOfPO).SCS120KHZoneT_SCS60KHZhalfT_SCS30KHZquarterT_SCS15KHZoneEighthT) = make([]int64, n)
				for i := int64(0); i < n; i++ {
					v, err := d.DecodeInteger(per.Constrained(0, 1119))
					if err != nil {
						return err
					}
					(*(*ie.FirstPDCCH_MonitoringOccasionOfPO).SCS120KHZoneT_SCS60KHZhalfT_SCS30KHZquarterT_SCS15KHZoneEighthT)[i] = v
				}
			case PCCH_Config_FirstPDCCH_MonitoringOccasionOfPO_SCS120KHZhalfT_SCS60KHZquarterT_SCS30KHZoneEighthT_SCS15KHZoneSixteenthT:
				(*ie.FirstPDCCH_MonitoringOccasionOfPO).SCS120KHZhalfT_SCS60KHZquarterT_SCS30KHZoneEighthT_SCS15KHZoneSixteenthT = new([]int64)
				seqOf := d.NewSequenceOfDecoder(pCCHConfigFirstPDCCHMonitoringOccasionOfPOSCS120KHZhalfTSCS60KHZquarterTSCS30KHZoneEighthTSCS15KHZoneSixteenthTConstraints)
				n, err := seqOf.DecodeLength()
				if err != nil {
					return err
				}
				(*(*ie.FirstPDCCH_MonitoringOccasionOfPO).SCS120KHZhalfT_SCS60KHZquarterT_SCS30KHZoneEighthT_SCS15KHZoneSixteenthT) = make([]int64, n)
				for i := int64(0); i < n; i++ {
					v, err := d.DecodeInteger(per.Constrained(0, 2239))
					if err != nil {
						return err
					}
					(*(*ie.FirstPDCCH_MonitoringOccasionOfPO).SCS120KHZhalfT_SCS60KHZquarterT_SCS30KHZoneEighthT_SCS15KHZoneSixteenthT)[i] = v
				}
			case PCCH_Config_FirstPDCCH_MonitoringOccasionOfPO_SCS480KHZoneT_SCS120KHZquarterT_SCS60KHZoneEighthT_SCS30KHZoneSixteenthT:
				(*ie.FirstPDCCH_MonitoringOccasionOfPO).SCS480KHZoneT_SCS120KHZquarterT_SCS60KHZoneEighthT_SCS30KHZoneSixteenthT = new([]int64)
				seqOf := d.NewSequenceOfDecoder(pCCHConfigFirstPDCCHMonitoringOccasionOfPOSCS480KHZoneTSCS120KHZquarterTSCS60KHZoneEighthTSCS30KHZoneSixteenthTConstraints)
				n, err := seqOf.DecodeLength()
				if err != nil {
					return err
				}
				(*(*ie.FirstPDCCH_MonitoringOccasionOfPO).SCS480KHZoneT_SCS120KHZquarterT_SCS60KHZoneEighthT_SCS30KHZoneSixteenthT) = make([]int64, n)
				for i := int64(0); i < n; i++ {
					v, err := d.DecodeInteger(per.Constrained(0, 4479))
					if err != nil {
						return err
					}
					(*(*ie.FirstPDCCH_MonitoringOccasionOfPO).SCS480KHZoneT_SCS120KHZquarterT_SCS60KHZoneEighthT_SCS30KHZoneSixteenthT)[i] = v
				}
			case PCCH_Config_FirstPDCCH_MonitoringOccasionOfPO_SCS480KHZhalfT_SCS120KHZoneEighthT_SCS60KHZoneSixteenthT:
				(*ie.FirstPDCCH_MonitoringOccasionOfPO).SCS480KHZhalfT_SCS120KHZoneEighthT_SCS60KHZoneSixteenthT = new([]int64)
				seqOf := d.NewSequenceOfDecoder(pCCHConfigFirstPDCCHMonitoringOccasionOfPOSCS480KHZhalfTSCS120KHZoneEighthTSCS60KHZoneSixteenthTConstraints)
				n, err := seqOf.DecodeLength()
				if err != nil {
					return err
				}
				(*(*ie.FirstPDCCH_MonitoringOccasionOfPO).SCS480KHZhalfT_SCS120KHZoneEighthT_SCS60KHZoneSixteenthT) = make([]int64, n)
				for i := int64(0); i < n; i++ {
					v, err := d.DecodeInteger(per.Constrained(0, 8959))
					if err != nil {
						return err
					}
					(*(*ie.FirstPDCCH_MonitoringOccasionOfPO).SCS480KHZhalfT_SCS120KHZoneEighthT_SCS60KHZoneSixteenthT)[i] = v
				}
			case PCCH_Config_FirstPDCCH_MonitoringOccasionOfPO_SCS480KHZquarterT_SCS120KHZoneSixteenthT:
				(*ie.FirstPDCCH_MonitoringOccasionOfPO).SCS480KHZquarterT_SCS120KHZoneSixteenthT = new([]int64)
				seqOf := d.NewSequenceOfDecoder(pCCHConfigFirstPDCCHMonitoringOccasionOfPOSCS480KHZquarterTSCS120KHZoneSixteenthTConstraints)
				n, err := seqOf.DecodeLength()
				if err != nil {
					return err
				}
				(*(*ie.FirstPDCCH_MonitoringOccasionOfPO).SCS480KHZquarterT_SCS120KHZoneSixteenthT) = make([]int64, n)
				for i := int64(0); i < n; i++ {
					v, err := d.DecodeInteger(per.Constrained(0, 17919))
					if err != nil {
						return err
					}
					(*(*ie.FirstPDCCH_MonitoringOccasionOfPO).SCS480KHZquarterT_SCS120KHZoneSixteenthT)[i] = v
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
				{Name: "nrofPDCCH-MonitoringOccasionPerSSB-InPO-r16", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeInteger(pCCHConfigNrofPDCCHMonitoringOccasionPerSSBInPOR16Constraints)
			if err != nil {
				return err
			}
			ie.NrofPDCCH_MonitoringOccasionPerSSB_InPO_r16 = &v
		}
	}

	// Extension group 1:
	if seq.IsExtensionPresent(1) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "ranPagingInIdlePO-r17", Optional: true},
				{Name: "firstPDCCH-MonitoringOccasionOfPO-v1710", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(pCCHConfigExtRanPagingInIdlePOR17Constraints)
			if err != nil {
				return err
			}
			ie.RanPagingInIdlePO_r17 = &v
		}

		if groupSeq.IsComponentPresent(1) {
			ie.FirstPDCCH_MonitoringOccasionOfPO_v1710 = &struct {
				Choice                 int
				SCS480KHZoneEighthT    *[]int64
				SCS480KHZoneSixteenthT *[]int64
			}{}
			choiceDec := dx.NewChoiceDecoder(pCCHConfigExtFirstPDCCHMonitoringOccasionOfPOV1710Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.FirstPDCCH_MonitoringOccasionOfPO_v1710).Choice = int(index)
			switch index {
			case PCCH_Config_Ext_FirstPDCCH_MonitoringOccasionOfPO_v1710_SCS480KHZoneEighthT:
				(*ie.FirstPDCCH_MonitoringOccasionOfPO_v1710).SCS480KHZoneEighthT = new([]int64)
				seqOf := dx.NewSequenceOfDecoder(pCCHConfigExtFirstPDCCHMonitoringOccasionOfPOV1710SCS480KHZoneEighthTConstraints)
				n, err := seqOf.DecodeLength()
				if err != nil {
					return err
				}
				(*(*ie.FirstPDCCH_MonitoringOccasionOfPO_v1710).SCS480KHZoneEighthT) = make([]int64, n)
				for i := int64(0); i < n; i++ {
					v, err := dx.DecodeInteger(per.Constrained(0, 35839))
					if err != nil {
						return err
					}
					(*(*ie.FirstPDCCH_MonitoringOccasionOfPO_v1710).SCS480KHZoneEighthT)[i] = v
				}
			case PCCH_Config_Ext_FirstPDCCH_MonitoringOccasionOfPO_v1710_SCS480KHZoneSixteenthT:
				(*ie.FirstPDCCH_MonitoringOccasionOfPO_v1710).SCS480KHZoneSixteenthT = new([]int64)
				seqOf := dx.NewSequenceOfDecoder(pCCHConfigExtFirstPDCCHMonitoringOccasionOfPOV1710SCS480KHZoneSixteenthTConstraints)
				n, err := seqOf.DecodeLength()
				if err != nil {
					return err
				}
				(*(*ie.FirstPDCCH_MonitoringOccasionOfPO_v1710).SCS480KHZoneSixteenthT) = make([]int64, n)
				for i := int64(0); i < n; i++ {
					v, err := dx.DecodeInteger(per.Constrained(0, 71679))
					if err != nil {
						return err
					}
					(*(*ie.FirstPDCCH_MonitoringOccasionOfPO_v1710).SCS480KHZoneSixteenthT)[i] = v
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
				{Name: "pagingAdaptation-r19", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.PagingAdaptation_r19 = &struct {
				PagingAdapt_NS_r19                   int64
				PagingAdaptNAndPagingFrameOffset_r19 struct {
					Choice           int
					OneT             *struct{}
					HalfT            *int64
					QuarterT         *int64
					OneEighthT       *int64
					OneSixteenthT    *int64
					OneThirtySecondT *int64
				}
				PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19 *struct {
					Choice                                                                                            int
					SCS15KHZoneT                                                                                      *[]int64
					SCS30KHZoneT_SCS15KHZhalfT                                                                        *[]int64
					SCS60KHZoneT_SCS30KHZhalfT_SCS15KHZquarterT                                                       *[]int64
					SCS120KHZoneT_SCS60KHZhalfT_SCS30KHZquarterT_SCS15KHZoneEighthT                                   *[]int64
					SCS120KHZhalfT_SCS60KHZquarterT_SCS30KHZoneEighthT_SCS15KHZoneSixteenthT                          *[]int64
					SCS480KHZoneT_SCS120KHZquarterT_SCS60KHZoneEighthT_SCS30KHZoneSixteenthT_SCS15KHZoneThirtySecondT *[]int64
					SCS480KHZhalfT_SCS120KHZoneEighthT_SCS60KHZoneSixteenthT_SCS30KHZoneThirtySecondT                 *[]int64
					SCS480KHZquarterT_SCS120KHZoneSixteenthT_SCS60KHZoneThirtySecondT                                 *[]int64
					SCS480KHZoneEighthT_SCS120KHZoneThirtySecondT                                                     *[]int64
					SCS480KHZoneSixteenthT                                                                            *[]int64
					SCS480KHZoneThirtySecondT                                                                         *[]int64
				}
			}{}
			c := ie.PagingAdaptation_r19
			pCCHConfigExtPagingAdaptationR19Seq := dx.NewSequenceDecoder(pCCHConfigExtPagingAdaptationR19Constraints)
			if err := pCCHConfigExtPagingAdaptationR19Seq.DecodePreamble(); err != nil {
				return err
			}
			{
				v, err := dx.DecodeEnumerated(pCCHConfigExtPagingAdaptationR19PagingAdaptNSR19Constraints)
				if err != nil {
					return err
				}
				c.PagingAdapt_NS_r19 = v
			}
			{
				choiceDec := dx.NewChoiceDecoder(pCCHConfigExtPagingAdaptationR19PagingAdaptNAndPagingFrameOffsetR19Constraints)
				index, _, _, err := choiceDec.DecodeChoice()
				if err != nil {
					return err
				}
				c.PagingAdaptNAndPagingFrameOffset_r19.Choice = int(index)
				switch index {
				case PCCH_Config_Ext_PagingAdaptation_r19_PagingAdaptNAndPagingFrameOffset_r19_OneT:
					c.PagingAdaptNAndPagingFrameOffset_r19.OneT = &struct{}{}
					if err := dx.DecodeNull(); err != nil {
						return err
					}
				case PCCH_Config_Ext_PagingAdaptation_r19_PagingAdaptNAndPagingFrameOffset_r19_HalfT:
					c.PagingAdaptNAndPagingFrameOffset_r19.HalfT = new(int64)
					v, err := dx.DecodeInteger(per.Constrained(0, 1))
					if err != nil {
						return err
					}
					(*c.PagingAdaptNAndPagingFrameOffset_r19.HalfT) = v
				case PCCH_Config_Ext_PagingAdaptation_r19_PagingAdaptNAndPagingFrameOffset_r19_QuarterT:
					c.PagingAdaptNAndPagingFrameOffset_r19.QuarterT = new(int64)
					v, err := dx.DecodeInteger(per.Constrained(0, 3))
					if err != nil {
						return err
					}
					(*c.PagingAdaptNAndPagingFrameOffset_r19.QuarterT) = v
				case PCCH_Config_Ext_PagingAdaptation_r19_PagingAdaptNAndPagingFrameOffset_r19_OneEighthT:
					c.PagingAdaptNAndPagingFrameOffset_r19.OneEighthT = new(int64)
					v, err := dx.DecodeInteger(per.Constrained(0, 7))
					if err != nil {
						return err
					}
					(*c.PagingAdaptNAndPagingFrameOffset_r19.OneEighthT) = v
				case PCCH_Config_Ext_PagingAdaptation_r19_PagingAdaptNAndPagingFrameOffset_r19_OneSixteenthT:
					c.PagingAdaptNAndPagingFrameOffset_r19.OneSixteenthT = new(int64)
					v, err := dx.DecodeInteger(per.Constrained(0, 15))
					if err != nil {
						return err
					}
					(*c.PagingAdaptNAndPagingFrameOffset_r19.OneSixteenthT) = v
				case PCCH_Config_Ext_PagingAdaptation_r19_PagingAdaptNAndPagingFrameOffset_r19_OneThirtySecondT:
					c.PagingAdaptNAndPagingFrameOffset_r19.OneThirtySecondT = new(int64)
					v, err := dx.DecodeInteger(per.Constrained(0, 31))
					if err != nil {
						return err
					}
					(*c.PagingAdaptNAndPagingFrameOffset_r19.OneThirtySecondT) = v
				}
			}
			if pCCHConfigExtPagingAdaptationR19Seq.IsComponentPresent(2) {
				c.PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19 = &struct {
					Choice                                                                                            int
					SCS15KHZoneT                                                                                      *[]int64
					SCS30KHZoneT_SCS15KHZhalfT                                                                        *[]int64
					SCS60KHZoneT_SCS30KHZhalfT_SCS15KHZquarterT                                                       *[]int64
					SCS120KHZoneT_SCS60KHZhalfT_SCS30KHZquarterT_SCS15KHZoneEighthT                                   *[]int64
					SCS120KHZhalfT_SCS60KHZquarterT_SCS30KHZoneEighthT_SCS15KHZoneSixteenthT                          *[]int64
					SCS480KHZoneT_SCS120KHZquarterT_SCS60KHZoneEighthT_SCS30KHZoneSixteenthT_SCS15KHZoneThirtySecondT *[]int64
					SCS480KHZhalfT_SCS120KHZoneEighthT_SCS60KHZoneSixteenthT_SCS30KHZoneThirtySecondT                 *[]int64
					SCS480KHZquarterT_SCS120KHZoneSixteenthT_SCS60KHZoneThirtySecondT                                 *[]int64
					SCS480KHZoneEighthT_SCS120KHZoneThirtySecondT                                                     *[]int64
					SCS480KHZoneSixteenthT                                                                            *[]int64
					SCS480KHZoneThirtySecondT                                                                         *[]int64
				}{}
				choiceDec := dx.NewChoiceDecoder(pCCHConfigExtPagingAdaptationR19PagingAdaptFirstPDCCHMonitoringOccasionOfPOR19Constraints)
				index, _, _, err := choiceDec.DecodeChoice()
				if err != nil {
					return err
				}
				(*c.PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19).Choice = int(index)
				switch index {
				case PCCH_Config_Ext_PagingAdaptation_r19_PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19_SCS15KHZoneT:
					(*c.PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19).SCS15KHZoneT = new([]int64)
					seqOf := dx.NewSequenceOfDecoder(pCCHConfigExtPagingAdaptationR19PagingAdaptFirstPDCCHMonitoringOccasionOfPOR19SCS15KHZoneTConstraints)
					n, err := seqOf.DecodeLength()
					if err != nil {
						return err
					}
					(*(*c.PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19).SCS15KHZoneT) = make([]int64, n)
					for i := int64(0); i < n; i++ {
						v, err := dx.DecodeInteger(per.Constrained(0, 139))
						if err != nil {
							return err
						}
						(*(*c.PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19).SCS15KHZoneT)[i] = v
					}
				case PCCH_Config_Ext_PagingAdaptation_r19_PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19_SCS30KHZoneT_SCS15KHZhalfT:
					(*c.PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19).SCS30KHZoneT_SCS15KHZhalfT = new([]int64)
					seqOf := dx.NewSequenceOfDecoder(pCCHConfigExtPagingAdaptationR19PagingAdaptFirstPDCCHMonitoringOccasionOfPOR19SCS30KHZoneTSCS15KHZhalfTConstraints)
					n, err := seqOf.DecodeLength()
					if err != nil {
						return err
					}
					(*(*c.PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19).SCS30KHZoneT_SCS15KHZhalfT) = make([]int64, n)
					for i := int64(0); i < n; i++ {
						v, err := dx.DecodeInteger(per.Constrained(0, 279))
						if err != nil {
							return err
						}
						(*(*c.PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19).SCS30KHZoneT_SCS15KHZhalfT)[i] = v
					}
				case PCCH_Config_Ext_PagingAdaptation_r19_PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19_SCS60KHZoneT_SCS30KHZhalfT_SCS15KHZquarterT:
					(*c.PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19).SCS60KHZoneT_SCS30KHZhalfT_SCS15KHZquarterT = new([]int64)
					seqOf := dx.NewSequenceOfDecoder(pCCHConfigExtPagingAdaptationR19PagingAdaptFirstPDCCHMonitoringOccasionOfPOR19SCS60KHZoneTSCS30KHZhalfTSCS15KHZquarterTConstraints)
					n, err := seqOf.DecodeLength()
					if err != nil {
						return err
					}
					(*(*c.PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19).SCS60KHZoneT_SCS30KHZhalfT_SCS15KHZquarterT) = make([]int64, n)
					for i := int64(0); i < n; i++ {
						v, err := dx.DecodeInteger(per.Constrained(0, 559))
						if err != nil {
							return err
						}
						(*(*c.PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19).SCS60KHZoneT_SCS30KHZhalfT_SCS15KHZquarterT)[i] = v
					}
				case PCCH_Config_Ext_PagingAdaptation_r19_PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19_SCS120KHZoneT_SCS60KHZhalfT_SCS30KHZquarterT_SCS15KHZoneEighthT:
					(*c.PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19).SCS120KHZoneT_SCS60KHZhalfT_SCS30KHZquarterT_SCS15KHZoneEighthT = new([]int64)
					seqOf := dx.NewSequenceOfDecoder(pCCHConfigExtPagingAdaptationR19PagingAdaptFirstPDCCHMonitoringOccasionOfPOR19SCS120KHZoneTSCS60KHZhalfTSCS30KHZquarterTSCS15KHZoneEighthTConstraints)
					n, err := seqOf.DecodeLength()
					if err != nil {
						return err
					}
					(*(*c.PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19).SCS120KHZoneT_SCS60KHZhalfT_SCS30KHZquarterT_SCS15KHZoneEighthT) = make([]int64, n)
					for i := int64(0); i < n; i++ {
						v, err := dx.DecodeInteger(per.Constrained(0, 1119))
						if err != nil {
							return err
						}
						(*(*c.PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19).SCS120KHZoneT_SCS60KHZhalfT_SCS30KHZquarterT_SCS15KHZoneEighthT)[i] = v
					}
				case PCCH_Config_Ext_PagingAdaptation_r19_PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19_SCS120KHZhalfT_SCS60KHZquarterT_SCS30KHZoneEighthT_SCS15KHZoneSixteenthT:
					(*c.PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19).SCS120KHZhalfT_SCS60KHZquarterT_SCS30KHZoneEighthT_SCS15KHZoneSixteenthT = new([]int64)
					seqOf := dx.NewSequenceOfDecoder(pCCHConfigExtPagingAdaptationR19PagingAdaptFirstPDCCHMonitoringOccasionOfPOR19SCS120KHZhalfTSCS60KHZquarterTSCS30KHZoneEighthTSCS15KHZoneSixteenthTConstraints)
					n, err := seqOf.DecodeLength()
					if err != nil {
						return err
					}
					(*(*c.PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19).SCS120KHZhalfT_SCS60KHZquarterT_SCS30KHZoneEighthT_SCS15KHZoneSixteenthT) = make([]int64, n)
					for i := int64(0); i < n; i++ {
						v, err := dx.DecodeInteger(per.Constrained(0, 2239))
						if err != nil {
							return err
						}
						(*(*c.PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19).SCS120KHZhalfT_SCS60KHZquarterT_SCS30KHZoneEighthT_SCS15KHZoneSixteenthT)[i] = v
					}
				case PCCH_Config_Ext_PagingAdaptation_r19_PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19_SCS480KHZoneT_SCS120KHZquarterT_SCS60KHZoneEighthT_SCS30KHZoneSixteenthT_SCS15KHZoneThirtySecondT:
					(*c.PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19).SCS480KHZoneT_SCS120KHZquarterT_SCS60KHZoneEighthT_SCS30KHZoneSixteenthT_SCS15KHZoneThirtySecondT = new([]int64)
					seqOf := dx.NewSequenceOfDecoder(pCCHConfigExtPagingAdaptationR19PagingAdaptFirstPDCCHMonitoringOccasionOfPOR19SCS480KHZoneTSCS120KHZquarterTSCS60KHZoneEighthTSCS30KHZoneSixteenthTSCS15KHZoneThirtySecondTConstraints)
					n, err := seqOf.DecodeLength()
					if err != nil {
						return err
					}
					(*(*c.PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19).SCS480KHZoneT_SCS120KHZquarterT_SCS60KHZoneEighthT_SCS30KHZoneSixteenthT_SCS15KHZoneThirtySecondT) = make([]int64, n)
					for i := int64(0); i < n; i++ {
						v, err := dx.DecodeInteger(per.Constrained(0, 4479))
						if err != nil {
							return err
						}
						(*(*c.PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19).SCS480KHZoneT_SCS120KHZquarterT_SCS60KHZoneEighthT_SCS30KHZoneSixteenthT_SCS15KHZoneThirtySecondT)[i] = v
					}
				case PCCH_Config_Ext_PagingAdaptation_r19_PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19_SCS480KHZhalfT_SCS120KHZoneEighthT_SCS60KHZoneSixteenthT_SCS30KHZoneThirtySecondT:
					(*c.PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19).SCS480KHZhalfT_SCS120KHZoneEighthT_SCS60KHZoneSixteenthT_SCS30KHZoneThirtySecondT = new([]int64)
					seqOf := dx.NewSequenceOfDecoder(pCCHConfigExtPagingAdaptationR19PagingAdaptFirstPDCCHMonitoringOccasionOfPOR19SCS480KHZhalfTSCS120KHZoneEighthTSCS60KHZoneSixteenthTSCS30KHZoneThirtySecondTConstraints)
					n, err := seqOf.DecodeLength()
					if err != nil {
						return err
					}
					(*(*c.PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19).SCS480KHZhalfT_SCS120KHZoneEighthT_SCS60KHZoneSixteenthT_SCS30KHZoneThirtySecondT) = make([]int64, n)
					for i := int64(0); i < n; i++ {
						v, err := dx.DecodeInteger(per.Constrained(0, 8959))
						if err != nil {
							return err
						}
						(*(*c.PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19).SCS480KHZhalfT_SCS120KHZoneEighthT_SCS60KHZoneSixteenthT_SCS30KHZoneThirtySecondT)[i] = v
					}
				case PCCH_Config_Ext_PagingAdaptation_r19_PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19_SCS480KHZquarterT_SCS120KHZoneSixteenthT_SCS60KHZoneThirtySecondT:
					(*c.PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19).SCS480KHZquarterT_SCS120KHZoneSixteenthT_SCS60KHZoneThirtySecondT = new([]int64)
					seqOf := dx.NewSequenceOfDecoder(pCCHConfigExtPagingAdaptationR19PagingAdaptFirstPDCCHMonitoringOccasionOfPOR19SCS480KHZquarterTSCS120KHZoneSixteenthTSCS60KHZoneThirtySecondTConstraints)
					n, err := seqOf.DecodeLength()
					if err != nil {
						return err
					}
					(*(*c.PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19).SCS480KHZquarterT_SCS120KHZoneSixteenthT_SCS60KHZoneThirtySecondT) = make([]int64, n)
					for i := int64(0); i < n; i++ {
						v, err := dx.DecodeInteger(per.Constrained(0, 17919))
						if err != nil {
							return err
						}
						(*(*c.PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19).SCS480KHZquarterT_SCS120KHZoneSixteenthT_SCS60KHZoneThirtySecondT)[i] = v
					}
				case PCCH_Config_Ext_PagingAdaptation_r19_PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19_SCS480KHZoneEighthT_SCS120KHZoneThirtySecondT:
					(*c.PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19).SCS480KHZoneEighthT_SCS120KHZoneThirtySecondT = new([]int64)
					seqOf := dx.NewSequenceOfDecoder(pCCHConfigExtPagingAdaptationR19PagingAdaptFirstPDCCHMonitoringOccasionOfPOR19SCS480KHZoneEighthTSCS120KHZoneThirtySecondTConstraints)
					n, err := seqOf.DecodeLength()
					if err != nil {
						return err
					}
					(*(*c.PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19).SCS480KHZoneEighthT_SCS120KHZoneThirtySecondT) = make([]int64, n)
					for i := int64(0); i < n; i++ {
						v, err := dx.DecodeInteger(per.Constrained(0, 35839))
						if err != nil {
							return err
						}
						(*(*c.PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19).SCS480KHZoneEighthT_SCS120KHZoneThirtySecondT)[i] = v
					}
				case PCCH_Config_Ext_PagingAdaptation_r19_PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19_SCS480KHZoneSixteenthT:
					(*c.PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19).SCS480KHZoneSixteenthT = new([]int64)
					seqOf := dx.NewSequenceOfDecoder(pCCHConfigExtPagingAdaptationR19PagingAdaptFirstPDCCHMonitoringOccasionOfPOR19SCS480KHZoneSixteenthTConstraints)
					n, err := seqOf.DecodeLength()
					if err != nil {
						return err
					}
					(*(*c.PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19).SCS480KHZoneSixteenthT) = make([]int64, n)
					for i := int64(0); i < n; i++ {
						v, err := dx.DecodeInteger(per.Constrained(0, 71679))
						if err != nil {
							return err
						}
						(*(*c.PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19).SCS480KHZoneSixteenthT)[i] = v
					}
				case PCCH_Config_Ext_PagingAdaptation_r19_PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19_SCS480KHZoneThirtySecondT:
					(*c.PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19).SCS480KHZoneThirtySecondT = new([]int64)
					seqOf := dx.NewSequenceOfDecoder(pCCHConfigExtPagingAdaptationR19PagingAdaptFirstPDCCHMonitoringOccasionOfPOR19SCS480KHZoneThirtySecondTConstraints)
					n, err := seqOf.DecodeLength()
					if err != nil {
						return err
					}
					(*(*c.PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19).SCS480KHZoneThirtySecondT) = make([]int64, n)
					for i := int64(0); i < n; i++ {
						v, err := dx.DecodeInteger(per.Constrained(0, 143359))
						if err != nil {
							return err
						}
						(*(*c.PagingAdaptFirstPDCCH_MonitoringOccasionOfPO_r19).SCS480KHZoneThirtySecondT)[i] = v
					}
				}
			}
		}
	}

	return nil
}
