// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: LowPowerConfig-r19 (line 7941).

var lowPowerConfigR19Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "lpwus-MvalueAndSeqConfigFR1-r19", Optional: true},
		{Name: "lpwus-MvalueAndSeqConfigFR2-r19", Optional: true},
		{Name: "lpwus-LO-FrameOffsetList-r19"},
		{Name: "lpwus-LO-FrameOffsetListForPagingAdapt-r19"},
		{Name: "lpwus-MO-NumPerLO-r19"},
		{Name: "lpwus-PO-NumPerLO-r19", Optional: true},
		{Name: "lpwus-EPRE-Ratio-r19", Optional: true},
		{Name: "lpwus-AvailableSlot-r19", Optional: true},
		{Name: "lpwus-AvailableSymbol-r19", Optional: true},
		{Name: "lpwus-OffsetFirstMoWithinLo-r19", Optional: true},
		{Name: "lpwus-NominalMoDuration-r19", Optional: true},
		{Name: "lpwus-ActualDuration-r19", Optional: true},
		{Name: "lpwus-LPSS-StartRB-r19", Optional: true},
		{Name: "lpwus-LPSS-BeamSubset-r19", Optional: true},
		{Name: "lpss-EPRE-Ratio-r19", Optional: true},
		{Name: "lpss-BinarySeqIndex-r19", Optional: true},
		{Name: "lpss-MvalueAndSeqConfig-r19", Optional: true},
		{Name: "lpss-PeriodicityAndOffset-r19", Optional: true},
		{Name: "lpss-StartSymbol-r19", Optional: true},
		{Name: "lp-SubgroupConfig-r19"},
		{Name: "entryCondition-r19"},
		{Name: "exitCondition-r19"},
	},
}

var lowPowerConfig_r19LpwusMvalueAndSeqConfigFR1R19Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "nOne"},
		{Name: "nTwo"},
		{Name: "nFour"},
	},
}

const (
	LowPowerConfig_r19_Lpwus_MvalueAndSeqConfigFR1_r19_NOne  = 0
	LowPowerConfig_r19_Lpwus_MvalueAndSeqConfigFR1_r19_NTwo  = 1
	LowPowerConfig_r19_Lpwus_MvalueAndSeqConfigFR1_r19_NFour = 2
)

var lowPowerConfig_r19LpwusMvalueAndSeqConfigFR2R19Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "nOne"},
	},
}

const (
	LowPowerConfig_r19_Lpwus_MvalueAndSeqConfigFR2_r19_NOne = 0
)

const (
	LowPowerConfig_r19_Lpwus_MO_NumPerLO_r19_N1 = 0
	LowPowerConfig_r19_Lpwus_MO_NumPerLO_r19_N2 = 1
	LowPowerConfig_r19_Lpwus_MO_NumPerLO_r19_N3 = 2
	LowPowerConfig_r19_Lpwus_MO_NumPerLO_r19_N4 = 3
)

var lowPowerConfigR19LpwusMONumPerLOR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{LowPowerConfig_r19_Lpwus_MO_NumPerLO_r19_N1, LowPowerConfig_r19_Lpwus_MO_NumPerLO_r19_N2, LowPowerConfig_r19_Lpwus_MO_NumPerLO_r19_N3, LowPowerConfig_r19_Lpwus_MO_NumPerLO_r19_N4},
}

const (
	LowPowerConfig_r19_Lpwus_PO_NumPerLO_r19_Po1 = 0
	LowPowerConfig_r19_Lpwus_PO_NumPerLO_r19_Po2 = 1
	LowPowerConfig_r19_Lpwus_PO_NumPerLO_r19_Po4 = 2
)

var lowPowerConfigR19LpwusPONumPerLOR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{LowPowerConfig_r19_Lpwus_PO_NumPerLO_r19_Po1, LowPowerConfig_r19_Lpwus_PO_NumPerLO_r19_Po2, LowPowerConfig_r19_Lpwus_PO_NumPerLO_r19_Po4},
}

const (
	LowPowerConfig_r19_Lpwus_EPRE_Ratio_r19_DBminus3 = 0
	LowPowerConfig_r19_Lpwus_EPRE_Ratio_r19_DB0      = 1
	LowPowerConfig_r19_Lpwus_EPRE_Ratio_r19_DB3      = 2
	LowPowerConfig_r19_Lpwus_EPRE_Ratio_r19_DB6      = 3
)

var lowPowerConfigR19LpwusEPRERatioR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{LowPowerConfig_r19_Lpwus_EPRE_Ratio_r19_DBminus3, LowPowerConfig_r19_Lpwus_EPRE_Ratio_r19_DB0, LowPowerConfig_r19_Lpwus_EPRE_Ratio_r19_DB3, LowPowerConfig_r19_Lpwus_EPRE_Ratio_r19_DB6},
}

var lowPowerConfig_r19LpwusAvailableSlotR19Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "n10"},
		{Name: "n20"},
		{Name: "n40"},
	},
}

const (
	LowPowerConfig_r19_Lpwus_AvailableSlot_r19_N10 = 0
	LowPowerConfig_r19_Lpwus_AvailableSlot_r19_N20 = 1
	LowPowerConfig_r19_Lpwus_AvailableSlot_r19_N40 = 2
)

var lowPowerConfig_r19LpwusAvailableSymbolR19Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "oneSlot"},
		{Name: "twoSlots"},
	},
}

const (
	LowPowerConfig_r19_Lpwus_AvailableSymbol_r19_OneSlot  = 0
	LowPowerConfig_r19_Lpwus_AvailableSymbol_r19_TwoSlots = 1
)

var lowPowerConfig_r19LpwusOffsetFirstMoWithinLoR19Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "offsetForSCS-15kHz"},
		{Name: "offsetForSCS-30kHz"},
		{Name: "offsetForSCS-120kHz"},
	},
}

const (
	LowPowerConfig_r19_Lpwus_OffsetFirstMoWithinLo_r19_OffsetForSCS_15kHz  = 0
	LowPowerConfig_r19_Lpwus_OffsetFirstMoWithinLo_r19_OffsetForSCS_30kHz  = 1
	LowPowerConfig_r19_Lpwus_OffsetFirstMoWithinLo_r19_OffsetForSCS_120kHz = 2
)

var lowPowerConfigR19LpwusNominalMoDurationR19Constraints = per.Constrained(1, 98)

var lowPowerConfig_r19LpwusActualDurationR19Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "mValue1"},
		{Name: "mValue2"},
		{Name: "mValue4"},
	},
}

const (
	LowPowerConfig_r19_Lpwus_ActualDuration_r19_MValue1 = 0
	LowPowerConfig_r19_Lpwus_ActualDuration_r19_MValue2 = 1
	LowPowerConfig_r19_Lpwus_ActualDuration_r19_MValue4 = 2
)

var lowPowerConfigR19LpwusLPSSStartRBR19Constraints = per.Constrained(0, 263)

var lowPowerConfig_r19LpwusLPSSBeamSubsetR19Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "shortBitmap"},
		{Name: "mediumBitmap"},
		{Name: "longBitmap"},
	},
}

const (
	LowPowerConfig_r19_Lpwus_LPSS_BeamSubset_r19_ShortBitmap  = 0
	LowPowerConfig_r19_Lpwus_LPSS_BeamSubset_r19_MediumBitmap = 1
	LowPowerConfig_r19_Lpwus_LPSS_BeamSubset_r19_LongBitmap   = 2
)

const (
	LowPowerConfig_r19_Lpss_EPRE_Ratio_r19_DBminus3 = 0
	LowPowerConfig_r19_Lpss_EPRE_Ratio_r19_DB0      = 1
	LowPowerConfig_r19_Lpss_EPRE_Ratio_r19_DB3      = 2
	LowPowerConfig_r19_Lpss_EPRE_Ratio_r19_DB6      = 3
)

var lowPowerConfigR19LpssEPRERatioR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{LowPowerConfig_r19_Lpss_EPRE_Ratio_r19_DBminus3, LowPowerConfig_r19_Lpss_EPRE_Ratio_r19_DB0, LowPowerConfig_r19_Lpss_EPRE_Ratio_r19_DB3, LowPowerConfig_r19_Lpss_EPRE_Ratio_r19_DB6},
}

var lowPowerConfigR19LpssBinarySeqIndexR19Constraints = per.Constrained(0, 3)

var lowPowerConfig_r19LpssMvalueAndSeqConfigR19Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "nOne"},
		{Name: "nTwo"},
		{Name: "nFour"},
	},
}

const (
	LowPowerConfig_r19_Lpss_MvalueAndSeqConfig_r19_NOne  = 0
	LowPowerConfig_r19_Lpss_MvalueAndSeqConfig_r19_NTwo  = 1
	LowPowerConfig_r19_Lpss_MvalueAndSeqConfig_r19_NFour = 2
)

var lowPowerConfig_r19LpssPeriodicityAndOffsetR19Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "ms160"},
		{Name: "ms320"},
	},
}

const (
	LowPowerConfig_r19_Lpss_PeriodicityAndOffset_r19_Ms160 = 0
	LowPowerConfig_r19_Lpss_PeriodicityAndOffset_r19_Ms320 = 1
)

var lowPowerConfigR19LpwusMvalueAndSeqConfigFR1R19NOneConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "lpwus-OverlaidSeqRoot-r19", Optional: true},
		{Name: "lpwus-OverlaidSeqNum-r19", Optional: true},
	},
}

var lowPowerConfigR19LpwusMvalueAndSeqConfigFR1R19NOneLpwusOverlaidSeqRootR19Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "root1-r19"},
		{Name: "root2-r19", Optional: true},
	},
}

const (
	LowPowerConfig_r19_Lpwus_MvalueAndSeqConfigFR1_r19_NOne_Lpwus_OverlaidSeqNum_r19_N1  = 0
	LowPowerConfig_r19_Lpwus_MvalueAndSeqConfigFR1_r19_NOne_Lpwus_OverlaidSeqNum_r19_N2  = 1
	LowPowerConfig_r19_Lpwus_MvalueAndSeqConfigFR1_r19_NOne_Lpwus_OverlaidSeqNum_r19_N4  = 2
	LowPowerConfig_r19_Lpwus_MvalueAndSeqConfigFR1_r19_NOne_Lpwus_OverlaidSeqNum_r19_N8  = 3
	LowPowerConfig_r19_Lpwus_MvalueAndSeqConfigFR1_r19_NOne_Lpwus_OverlaidSeqNum_r19_N16 = 4
)

var lowPowerConfigR19LpwusMvalueAndSeqConfigFR1R19NOneLpwusOverlaidSeqNumR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{LowPowerConfig_r19_Lpwus_MvalueAndSeqConfigFR1_r19_NOne_Lpwus_OverlaidSeqNum_r19_N1, LowPowerConfig_r19_Lpwus_MvalueAndSeqConfigFR1_r19_NOne_Lpwus_OverlaidSeqNum_r19_N2, LowPowerConfig_r19_Lpwus_MvalueAndSeqConfigFR1_r19_NOne_Lpwus_OverlaidSeqNum_r19_N4, LowPowerConfig_r19_Lpwus_MvalueAndSeqConfigFR1_r19_NOne_Lpwus_OverlaidSeqNum_r19_N8, LowPowerConfig_r19_Lpwus_MvalueAndSeqConfigFR1_r19_NOne_Lpwus_OverlaidSeqNum_r19_N16},
}

var lowPowerConfigR19LpwusMvalueAndSeqConfigFR1R19NTwoConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "lpwus-OverlaidSeqRoot-r19", Optional: true},
		{Name: "lpwus-OverlaidSeqNum-r19", Optional: true},
	},
}

var lowPowerConfigR19LpwusMvalueAndSeqConfigFR1R19NTwoLpwusOverlaidSeqRootR19Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "root1-r19"},
		{Name: "root2-r19", Optional: true},
	},
}

const (
	LowPowerConfig_r19_Lpwus_MvalueAndSeqConfigFR1_r19_NTwo_Lpwus_OverlaidSeqNum_r19_N1 = 0
	LowPowerConfig_r19_Lpwus_MvalueAndSeqConfigFR1_r19_NTwo_Lpwus_OverlaidSeqNum_r19_N2 = 1
	LowPowerConfig_r19_Lpwus_MvalueAndSeqConfigFR1_r19_NTwo_Lpwus_OverlaidSeqNum_r19_N4 = 2
	LowPowerConfig_r19_Lpwus_MvalueAndSeqConfigFR1_r19_NTwo_Lpwus_OverlaidSeqNum_r19_N8 = 3
)

var lowPowerConfigR19LpwusMvalueAndSeqConfigFR1R19NTwoLpwusOverlaidSeqNumR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{LowPowerConfig_r19_Lpwus_MvalueAndSeqConfigFR1_r19_NTwo_Lpwus_OverlaidSeqNum_r19_N1, LowPowerConfig_r19_Lpwus_MvalueAndSeqConfigFR1_r19_NTwo_Lpwus_OverlaidSeqNum_r19_N2, LowPowerConfig_r19_Lpwus_MvalueAndSeqConfigFR1_r19_NTwo_Lpwus_OverlaidSeqNum_r19_N4, LowPowerConfig_r19_Lpwus_MvalueAndSeqConfigFR1_r19_NTwo_Lpwus_OverlaidSeqNum_r19_N8},
}

var lowPowerConfigR19LpwusMvalueAndSeqConfigFR1R19NFourConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "lpwus-OverlaidSeqRoot-r19", Optional: true},
		{Name: "lpwus-OverlaidSeqNum-r19", Optional: true},
	},
}

var lowPowerConfigR19LpwusMvalueAndSeqConfigFR1R19NFourLpwusOverlaidSeqRootR19Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "root1-r19"},
		{Name: "root2-r19", Optional: true},
	},
}

const (
	LowPowerConfig_r19_Lpwus_MvalueAndSeqConfigFR1_r19_NFour_Lpwus_OverlaidSeqNum_r19_N1 = 0
	LowPowerConfig_r19_Lpwus_MvalueAndSeqConfigFR1_r19_NFour_Lpwus_OverlaidSeqNum_r19_N2 = 1
	LowPowerConfig_r19_Lpwus_MvalueAndSeqConfigFR1_r19_NFour_Lpwus_OverlaidSeqNum_r19_N4 = 2
)

var lowPowerConfigR19LpwusMvalueAndSeqConfigFR1R19NFourLpwusOverlaidSeqNumR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{LowPowerConfig_r19_Lpwus_MvalueAndSeqConfigFR1_r19_NFour_Lpwus_OverlaidSeqNum_r19_N1, LowPowerConfig_r19_Lpwus_MvalueAndSeqConfigFR1_r19_NFour_Lpwus_OverlaidSeqNum_r19_N2, LowPowerConfig_r19_Lpwus_MvalueAndSeqConfigFR1_r19_NFour_Lpwus_OverlaidSeqNum_r19_N4},
}

var lowPowerConfigR19LpwusMvalueAndSeqConfigFR2R19NOneConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "lpwus-OverlaidSeqRoot-r19", Optional: true},
		{Name: "lpwus-OverlaidSeqNum-SCS-120kHz-r19", Optional: true},
	},
}

var lowPowerConfigR19LpwusMvalueAndSeqConfigFR2R19NOneLpwusOverlaidSeqRootR19Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "root1-r19"},
		{Name: "root2-r19", Optional: true},
	},
}

const (
	LowPowerConfig_r19_Lpwus_MvalueAndSeqConfigFR2_r19_NOne_Lpwus_OverlaidSeqNum_SCS_120kHz_r19_N1 = 0
	LowPowerConfig_r19_Lpwus_MvalueAndSeqConfigFR2_r19_NOne_Lpwus_OverlaidSeqNum_SCS_120kHz_r19_N2 = 1
)

var lowPowerConfigR19LpwusMvalueAndSeqConfigFR2R19NOneLpwusOverlaidSeqNumSCS120kHzR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{LowPowerConfig_r19_Lpwus_MvalueAndSeqConfigFR2_r19_NOne_Lpwus_OverlaidSeqNum_SCS_120kHz_r19_N1, LowPowerConfig_r19_Lpwus_MvalueAndSeqConfigFR2_r19_NOne_Lpwus_OverlaidSeqNum_SCS_120kHz_r19_N2},
}

var lowPowerConfigR19LpwusLOFrameOffsetListR19Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "offsetForLongerWakeUpDelay-r19", Optional: true},
		{Name: "offsetForShorterWakeUpDelay-r19", Optional: true},
	},
}

var lowPowerConfigR19LpwusLOFrameOffsetListR19OffsetForLongerWakeUpDelayR19Constraints = per.SizeRange(1, 4)

var lowPowerConfigR19LpwusLOFrameOffsetListR19OffsetForShorterWakeUpDelayR19Constraints = per.SizeRange(1, 4)

var lowPowerConfigR19LpwusLOFrameOffsetListForPagingAdaptR19Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "offsetForLongerWakeUpDelay-r19", Optional: true},
		{Name: "offsetForShorterWakeUpDelay-r19", Optional: true},
	},
}

var lowPowerConfigR19LpwusLOFrameOffsetListForPagingAdaptR19OffsetForLongerWakeUpDelayR19Constraints = per.SizeRange(1, 8)

var lowPowerConfigR19LpwusLOFrameOffsetListForPagingAdaptR19OffsetForShorterWakeUpDelayR19Constraints = per.SizeRange(1, 8)

var lowPowerConfigR19LpssMvalueAndSeqConfigR19NOneConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "lpss-BinarySeqLen-r19", Optional: true},
		{Name: "lpss-OverlaidSeqRoot-r19", Optional: true},
	},
}

const (
	LowPowerConfig_r19_Lpss_MvalueAndSeqConfig_r19_NOne_Lpss_BinarySeqLen_r19_N6 = 0
	LowPowerConfig_r19_Lpss_MvalueAndSeqConfig_r19_NOne_Lpss_BinarySeqLen_r19_N8 = 1
)

var lowPowerConfigR19LpssMvalueAndSeqConfigR19NOneLpssBinarySeqLenR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{LowPowerConfig_r19_Lpss_MvalueAndSeqConfig_r19_NOne_Lpss_BinarySeqLen_r19_N6, LowPowerConfig_r19_Lpss_MvalueAndSeqConfig_r19_NOne_Lpss_BinarySeqLen_r19_N8},
}

var lowPowerConfigR19LpssMvalueAndSeqConfigR19NTwoConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "lpss-BinarySeqLen-r19", Optional: true},
		{Name: "lpss-OverlaidSeqRoot-r19", Optional: true},
	},
}

const (
	LowPowerConfig_r19_Lpss_MvalueAndSeqConfig_r19_NTwo_Lpss_BinarySeqLen_r19_N12 = 0
	LowPowerConfig_r19_Lpss_MvalueAndSeqConfig_r19_NTwo_Lpss_BinarySeqLen_r19_N16 = 1
)

var lowPowerConfigR19LpssMvalueAndSeqConfigR19NTwoLpssBinarySeqLenR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{LowPowerConfig_r19_Lpss_MvalueAndSeqConfig_r19_NTwo_Lpss_BinarySeqLen_r19_N12, LowPowerConfig_r19_Lpss_MvalueAndSeqConfig_r19_NTwo_Lpss_BinarySeqLen_r19_N16},
}

var lowPowerConfigR19LpssMvalueAndSeqConfigR19NFourConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "lpss-BinarySeqLen-r19", Optional: true},
		{Name: "lpss-OverlaidSeqRoot-r19", Optional: true},
	},
}

const (
	LowPowerConfig_r19_Lpss_MvalueAndSeqConfig_r19_NFour_Lpss_BinarySeqLen_r19_N16 = 0
	LowPowerConfig_r19_Lpss_MvalueAndSeqConfig_r19_NFour_Lpss_BinarySeqLen_r19_N32 = 1
)

var lowPowerConfigR19LpssMvalueAndSeqConfigR19NFourLpssBinarySeqLenR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{LowPowerConfig_r19_Lpss_MvalueAndSeqConfig_r19_NFour_Lpss_BinarySeqLen_r19_N16, LowPowerConfig_r19_Lpss_MvalueAndSeqConfig_r19_NFour_Lpss_BinarySeqLen_r19_N32},
}

var lowPowerConfigR19LpssStartSymbolR19Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "startSymbol1-r19"},
		{Name: "startSymbol2-r19", Optional: true},
	},
}

type LowPowerConfig_r19 struct {
	Lpwus_MvalueAndSeqConfigFR1_r19 *struct {
		Choice int
		NOne   *struct {
			Lpwus_OverlaidSeqRoot_r19 *struct {
				Root1_r19 int64
				Root2_r19 *int64
			}
			Lpwus_OverlaidSeqNum_r19 *int64
		}
		NTwo *struct {
			Lpwus_OverlaidSeqRoot_r19 *struct {
				Root1_r19 int64
				Root2_r19 *int64
			}
			Lpwus_OverlaidSeqNum_r19 *int64
		}
		NFour *struct {
			Lpwus_OverlaidSeqRoot_r19 *struct {
				Root1_r19 int64
				Root2_r19 *int64
			}
			Lpwus_OverlaidSeqNum_r19 *int64
		}
	}
	Lpwus_MvalueAndSeqConfigFR2_r19 *struct {
		Choice int
		NOne   *struct {
			Lpwus_OverlaidSeqRoot_r19 *struct {
				Root1_r19 int64
				Root2_r19 *int64
			}
			Lpwus_OverlaidSeqNum_SCS_120kHz_r19 *int64
		}
	}
	Lpwus_LO_FrameOffsetList_r19 struct {
		OffsetForLongerWakeUpDelay_r19  []int64
		OffsetForShorterWakeUpDelay_r19 []int64
	}
	Lpwus_LO_FrameOffsetListForPagingAdapt_r19 struct {
		OffsetForLongerWakeUpDelay_r19  []int64
		OffsetForShorterWakeUpDelay_r19 []int64
	}
	Lpwus_MO_NumPerLO_r19   int64
	Lpwus_PO_NumPerLO_r19   *int64
	Lpwus_EPRE_Ratio_r19    *int64
	Lpwus_AvailableSlot_r19 *struct {
		Choice int
		N10    *per.BitString
		N20    *per.BitString
		N40    *per.BitString
	}
	Lpwus_AvailableSymbol_r19 *struct {
		Choice   int
		OneSlot  *per.BitString
		TwoSlots *per.BitString
	}
	Lpwus_OffsetFirstMoWithinLo_r19 *struct {
		Choice              int
		OffsetForSCS_15kHz  *int64
		OffsetForSCS_30kHz  *int64
		OffsetForSCS_120kHz *int64
	}
	Lpwus_NominalMoDuration_r19 *int64
	Lpwus_ActualDuration_r19    *struct {
		Choice  int
		MValue1 *int64
		MValue2 *int64
		MValue4 *int64
	}
	Lpwus_LPSS_StartRB_r19    *int64
	Lpwus_LPSS_BeamSubset_r19 *struct {
		Choice       int
		ShortBitmap  *per.BitString
		MediumBitmap *per.BitString
		LongBitmap   *per.BitString
	}
	Lpss_EPRE_Ratio_r19         *int64
	Lpss_BinarySeqIndex_r19     *int64
	Lpss_MvalueAndSeqConfig_r19 *struct {
		Choice int
		NOne   *struct {
			Lpss_BinarySeqLen_r19    *int64
			Lpss_OverlaidSeqRoot_r19 *int64
		}
		NTwo *struct {
			Lpss_BinarySeqLen_r19    *int64
			Lpss_OverlaidSeqRoot_r19 *int64
		}
		NFour *struct {
			Lpss_BinarySeqLen_r19    *int64
			Lpss_OverlaidSeqRoot_r19 *int64
		}
	}
	Lpss_PeriodicityAndOffset_r19 *struct {
		Choice int
		Ms160  *int64
		Ms320  *int64
	}
	Lpss_StartSymbol_r19 *struct {
		StartSymbol1_r19 int64
		StartSymbol2_r19 *int64
	}
	Lp_SubgroupConfig_r19 LP_SubgroupConfig_r19
	EntryCondition_r19    EntryCondition_r19
	ExitCondition_r19     ExitCondition_r19
}

func (ie *LowPowerConfig_r19) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(lowPowerConfigR19Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Lpwus_MvalueAndSeqConfigFR1_r19 != nil, ie.Lpwus_MvalueAndSeqConfigFR2_r19 != nil, ie.Lpwus_PO_NumPerLO_r19 != nil, ie.Lpwus_EPRE_Ratio_r19 != nil, ie.Lpwus_AvailableSlot_r19 != nil, ie.Lpwus_AvailableSymbol_r19 != nil, ie.Lpwus_OffsetFirstMoWithinLo_r19 != nil, ie.Lpwus_NominalMoDuration_r19 != nil, ie.Lpwus_ActualDuration_r19 != nil, ie.Lpwus_LPSS_StartRB_r19 != nil, ie.Lpwus_LPSS_BeamSubset_r19 != nil, ie.Lpss_EPRE_Ratio_r19 != nil, ie.Lpss_BinarySeqIndex_r19 != nil, ie.Lpss_MvalueAndSeqConfig_r19 != nil, ie.Lpss_PeriodicityAndOffset_r19 != nil, ie.Lpss_StartSymbol_r19 != nil}); err != nil {
		return err
	}

	// 3. lpwus-MvalueAndSeqConfigFR1-r19: choice
	{
		if ie.Lpwus_MvalueAndSeqConfigFR1_r19 != nil {
			choiceEnc := e.NewChoiceEncoder(lowPowerConfig_r19LpwusMvalueAndSeqConfigFR1R19Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.Lpwus_MvalueAndSeqConfigFR1_r19).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.Lpwus_MvalueAndSeqConfigFR1_r19).Choice {
			case LowPowerConfig_r19_Lpwus_MvalueAndSeqConfigFR1_r19_NOne:
				c := (*(*ie.Lpwus_MvalueAndSeqConfigFR1_r19).NOne)
				lowPowerConfigR19LpwusMvalueAndSeqConfigFR1R19NOneSeq := e.NewSequenceEncoder(lowPowerConfigR19LpwusMvalueAndSeqConfigFR1R19NOneConstraints)
				if err := lowPowerConfigR19LpwusMvalueAndSeqConfigFR1R19NOneSeq.EncodePreamble([]bool{c.Lpwus_OverlaidSeqRoot_r19 != nil, c.Lpwus_OverlaidSeqNum_r19 != nil}); err != nil {
					return err
				}
				if c.Lpwus_OverlaidSeqRoot_r19 != nil {
					c := (*c.Lpwus_OverlaidSeqRoot_r19)
					lowPowerConfigR19LpwusMvalueAndSeqConfigFR1R19NOneLpwusOverlaidSeqRootR19Seq := e.NewSequenceEncoder(lowPowerConfigR19LpwusMvalueAndSeqConfigFR1R19NOneLpwusOverlaidSeqRootR19Constraints)
					if err := lowPowerConfigR19LpwusMvalueAndSeqConfigFR1R19NOneLpwusOverlaidSeqRootR19Seq.EncodePreamble([]bool{c.Root2_r19 != nil}); err != nil {
						return err
					}
					if err := e.EncodeInteger(c.Root1_r19, per.Constrained(1, 131)); err != nil {
						return err
					}
					if c.Root2_r19 != nil {
						if err := e.EncodeInteger((*c.Root2_r19), per.Constrained(1, 131)); err != nil {
							return err
						}
					}
				}
				if c.Lpwus_OverlaidSeqNum_r19 != nil {
					if err := e.EncodeEnumerated((*c.Lpwus_OverlaidSeqNum_r19), lowPowerConfigR19LpwusMvalueAndSeqConfigFR1R19NOneLpwusOverlaidSeqNumR19Constraints); err != nil {
						return err
					}
				}
			case LowPowerConfig_r19_Lpwus_MvalueAndSeqConfigFR1_r19_NTwo:
				c := (*(*ie.Lpwus_MvalueAndSeqConfigFR1_r19).NTwo)
				lowPowerConfigR19LpwusMvalueAndSeqConfigFR1R19NTwoSeq := e.NewSequenceEncoder(lowPowerConfigR19LpwusMvalueAndSeqConfigFR1R19NTwoConstraints)
				if err := lowPowerConfigR19LpwusMvalueAndSeqConfigFR1R19NTwoSeq.EncodePreamble([]bool{c.Lpwus_OverlaidSeqRoot_r19 != nil, c.Lpwus_OverlaidSeqNum_r19 != nil}); err != nil {
					return err
				}
				if c.Lpwus_OverlaidSeqRoot_r19 != nil {
					c := (*c.Lpwus_OverlaidSeqRoot_r19)
					lowPowerConfigR19LpwusMvalueAndSeqConfigFR1R19NTwoLpwusOverlaidSeqRootR19Seq := e.NewSequenceEncoder(lowPowerConfigR19LpwusMvalueAndSeqConfigFR1R19NTwoLpwusOverlaidSeqRootR19Constraints)
					if err := lowPowerConfigR19LpwusMvalueAndSeqConfigFR1R19NTwoLpwusOverlaidSeqRootR19Seq.EncodePreamble([]bool{c.Root2_r19 != nil}); err != nil {
						return err
					}
					if err := e.EncodeInteger(c.Root1_r19, per.Constrained(1, 61)); err != nil {
						return err
					}
					if c.Root2_r19 != nil {
						if err := e.EncodeInteger((*c.Root2_r19), per.Constrained(1, 61)); err != nil {
							return err
						}
					}
				}
				if c.Lpwus_OverlaidSeqNum_r19 != nil {
					if err := e.EncodeEnumerated((*c.Lpwus_OverlaidSeqNum_r19), lowPowerConfigR19LpwusMvalueAndSeqConfigFR1R19NTwoLpwusOverlaidSeqNumR19Constraints); err != nil {
						return err
					}
				}
			case LowPowerConfig_r19_Lpwus_MvalueAndSeqConfigFR1_r19_NFour:
				c := (*(*ie.Lpwus_MvalueAndSeqConfigFR1_r19).NFour)
				lowPowerConfigR19LpwusMvalueAndSeqConfigFR1R19NFourSeq := e.NewSequenceEncoder(lowPowerConfigR19LpwusMvalueAndSeqConfigFR1R19NFourConstraints)
				if err := lowPowerConfigR19LpwusMvalueAndSeqConfigFR1R19NFourSeq.EncodePreamble([]bool{c.Lpwus_OverlaidSeqRoot_r19 != nil, c.Lpwus_OverlaidSeqNum_r19 != nil}); err != nil {
					return err
				}
				if c.Lpwus_OverlaidSeqRoot_r19 != nil {
					c := (*c.Lpwus_OverlaidSeqRoot_r19)
					lowPowerConfigR19LpwusMvalueAndSeqConfigFR1R19NFourLpwusOverlaidSeqRootR19Seq := e.NewSequenceEncoder(lowPowerConfigR19LpwusMvalueAndSeqConfigFR1R19NFourLpwusOverlaidSeqRootR19Constraints)
					if err := lowPowerConfigR19LpwusMvalueAndSeqConfigFR1R19NFourLpwusOverlaidSeqRootR19Seq.EncodePreamble([]bool{c.Root2_r19 != nil}); err != nil {
						return err
					}
					if err := e.EncodeInteger(c.Root1_r19, per.Constrained(1, 31)); err != nil {
						return err
					}
					if c.Root2_r19 != nil {
						if err := e.EncodeInteger((*c.Root2_r19), per.Constrained(1, 31)); err != nil {
							return err
						}
					}
				}
				if c.Lpwus_OverlaidSeqNum_r19 != nil {
					if err := e.EncodeEnumerated((*c.Lpwus_OverlaidSeqNum_r19), lowPowerConfigR19LpwusMvalueAndSeqConfigFR1R19NFourLpwusOverlaidSeqNumR19Constraints); err != nil {
						return err
					}
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.Lpwus_MvalueAndSeqConfigFR1_r19).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 4. lpwus-MvalueAndSeqConfigFR2-r19: choice
	{
		if ie.Lpwus_MvalueAndSeqConfigFR2_r19 != nil {
			choiceEnc := e.NewChoiceEncoder(lowPowerConfig_r19LpwusMvalueAndSeqConfigFR2R19Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.Lpwus_MvalueAndSeqConfigFR2_r19).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.Lpwus_MvalueAndSeqConfigFR2_r19).Choice {
			case LowPowerConfig_r19_Lpwus_MvalueAndSeqConfigFR2_r19_NOne:
				c := (*(*ie.Lpwus_MvalueAndSeqConfigFR2_r19).NOne)
				lowPowerConfigR19LpwusMvalueAndSeqConfigFR2R19NOneSeq := e.NewSequenceEncoder(lowPowerConfigR19LpwusMvalueAndSeqConfigFR2R19NOneConstraints)
				if err := lowPowerConfigR19LpwusMvalueAndSeqConfigFR2R19NOneSeq.EncodePreamble([]bool{c.Lpwus_OverlaidSeqRoot_r19 != nil, c.Lpwus_OverlaidSeqNum_SCS_120kHz_r19 != nil}); err != nil {
					return err
				}
				if c.Lpwus_OverlaidSeqRoot_r19 != nil {
					c := (*c.Lpwus_OverlaidSeqRoot_r19)
					lowPowerConfigR19LpwusMvalueAndSeqConfigFR2R19NOneLpwusOverlaidSeqRootR19Seq := e.NewSequenceEncoder(lowPowerConfigR19LpwusMvalueAndSeqConfigFR2R19NOneLpwusOverlaidSeqRootR19Constraints)
					if err := lowPowerConfigR19LpwusMvalueAndSeqConfigFR2R19NOneLpwusOverlaidSeqRootR19Seq.EncodePreamble([]bool{c.Root2_r19 != nil}); err != nil {
						return err
					}
					if err := e.EncodeInteger(c.Root1_r19, per.Constrained(1, 131)); err != nil {
						return err
					}
					if c.Root2_r19 != nil {
						if err := e.EncodeInteger((*c.Root2_r19), per.Constrained(1, 131)); err != nil {
							return err
						}
					}
				}
				if c.Lpwus_OverlaidSeqNum_SCS_120kHz_r19 != nil {
					if err := e.EncodeEnumerated((*c.Lpwus_OverlaidSeqNum_SCS_120kHz_r19), lowPowerConfigR19LpwusMvalueAndSeqConfigFR2R19NOneLpwusOverlaidSeqNumSCS120kHzR19Constraints); err != nil {
						return err
					}
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.Lpwus_MvalueAndSeqConfigFR2_r19).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 5. lpwus-LO-FrameOffsetList-r19: sequence
	{
		{
			c := &ie.Lpwus_LO_FrameOffsetList_r19
			lowPowerConfigR19LpwusLOFrameOffsetListR19Seq := e.NewSequenceEncoder(lowPowerConfigR19LpwusLOFrameOffsetListR19Constraints)
			if err := lowPowerConfigR19LpwusLOFrameOffsetListR19Seq.EncodePreamble([]bool{c.OffsetForLongerWakeUpDelay_r19 != nil, c.OffsetForShorterWakeUpDelay_r19 != nil}); err != nil {
				return err
			}
			if c.OffsetForLongerWakeUpDelay_r19 != nil {
				seqOf := e.NewSequenceOfEncoder(lowPowerConfigR19LpwusLOFrameOffsetListR19OffsetForLongerWakeUpDelayR19Constraints)
				if err := seqOf.EncodeLength(int64(len(c.OffsetForLongerWakeUpDelay_r19))); err != nil {
					return err
				}
				for i := range c.OffsetForLongerWakeUpDelay_r19 {
					if err := e.EncodeInteger(c.OffsetForLongerWakeUpDelay_r19[i], per.Constrained(8, 200)); err != nil {
						return err
					}
				}
			}
			if c.OffsetForShorterWakeUpDelay_r19 != nil {
				seqOf := e.NewSequenceOfEncoder(lowPowerConfigR19LpwusLOFrameOffsetListR19OffsetForShorterWakeUpDelayR19Constraints)
				if err := seqOf.EncodeLength(int64(len(c.OffsetForShorterWakeUpDelay_r19))); err != nil {
					return err
				}
				for i := range c.OffsetForShorterWakeUpDelay_r19 {
					if err := e.EncodeInteger(c.OffsetForShorterWakeUpDelay_r19[i], per.Constrained(8, 200)); err != nil {
						return err
					}
				}
			}
		}
	}

	// 6. lpwus-LO-FrameOffsetListForPagingAdapt-r19: sequence
	{
		{
			c := &ie.Lpwus_LO_FrameOffsetListForPagingAdapt_r19
			lowPowerConfigR19LpwusLOFrameOffsetListForPagingAdaptR19Seq := e.NewSequenceEncoder(lowPowerConfigR19LpwusLOFrameOffsetListForPagingAdaptR19Constraints)
			if err := lowPowerConfigR19LpwusLOFrameOffsetListForPagingAdaptR19Seq.EncodePreamble([]bool{c.OffsetForLongerWakeUpDelay_r19 != nil, c.OffsetForShorterWakeUpDelay_r19 != nil}); err != nil {
				return err
			}
			if c.OffsetForLongerWakeUpDelay_r19 != nil {
				seqOf := e.NewSequenceOfEncoder(lowPowerConfigR19LpwusLOFrameOffsetListForPagingAdaptR19OffsetForLongerWakeUpDelayR19Constraints)
				if err := seqOf.EncodeLength(int64(len(c.OffsetForLongerWakeUpDelay_r19))); err != nil {
					return err
				}
				for i := range c.OffsetForLongerWakeUpDelay_r19 {
					if err := e.EncodeInteger(c.OffsetForLongerWakeUpDelay_r19[i], per.Constrained(8, 200)); err != nil {
						return err
					}
				}
			}
			if c.OffsetForShorterWakeUpDelay_r19 != nil {
				seqOf := e.NewSequenceOfEncoder(lowPowerConfigR19LpwusLOFrameOffsetListForPagingAdaptR19OffsetForShorterWakeUpDelayR19Constraints)
				if err := seqOf.EncodeLength(int64(len(c.OffsetForShorterWakeUpDelay_r19))); err != nil {
					return err
				}
				for i := range c.OffsetForShorterWakeUpDelay_r19 {
					if err := e.EncodeInteger(c.OffsetForShorterWakeUpDelay_r19[i], per.Constrained(8, 200)); err != nil {
						return err
					}
				}
			}
		}
	}

	// 7. lpwus-MO-NumPerLO-r19: enumerated
	{
		if err := e.EncodeEnumerated(ie.Lpwus_MO_NumPerLO_r19, lowPowerConfigR19LpwusMONumPerLOR19Constraints); err != nil {
			return err
		}
	}

	// 8. lpwus-PO-NumPerLO-r19: enumerated
	{
		if ie.Lpwus_PO_NumPerLO_r19 != nil {
			if err := e.EncodeEnumerated(*ie.Lpwus_PO_NumPerLO_r19, lowPowerConfigR19LpwusPONumPerLOR19Constraints); err != nil {
				return err
			}
		}
	}

	// 9. lpwus-EPRE-Ratio-r19: enumerated
	{
		if ie.Lpwus_EPRE_Ratio_r19 != nil {
			if err := e.EncodeEnumerated(*ie.Lpwus_EPRE_Ratio_r19, lowPowerConfigR19LpwusEPRERatioR19Constraints); err != nil {
				return err
			}
		}
	}

	// 10. lpwus-AvailableSlot-r19: choice
	{
		if ie.Lpwus_AvailableSlot_r19 != nil {
			choiceEnc := e.NewChoiceEncoder(lowPowerConfig_r19LpwusAvailableSlotR19Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.Lpwus_AvailableSlot_r19).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.Lpwus_AvailableSlot_r19).Choice {
			case LowPowerConfig_r19_Lpwus_AvailableSlot_r19_N10:
				if err := e.EncodeBitString((*(*ie.Lpwus_AvailableSlot_r19).N10), per.FixedSize(10)); err != nil {
					return err
				}
			case LowPowerConfig_r19_Lpwus_AvailableSlot_r19_N20:
				if err := e.EncodeBitString((*(*ie.Lpwus_AvailableSlot_r19).N20), per.FixedSize(20)); err != nil {
					return err
				}
			case LowPowerConfig_r19_Lpwus_AvailableSlot_r19_N40:
				if err := e.EncodeBitString((*(*ie.Lpwus_AvailableSlot_r19).N40), per.FixedSize(40)); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.Lpwus_AvailableSlot_r19).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 11. lpwus-AvailableSymbol-r19: choice
	{
		if ie.Lpwus_AvailableSymbol_r19 != nil {
			choiceEnc := e.NewChoiceEncoder(lowPowerConfig_r19LpwusAvailableSymbolR19Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.Lpwus_AvailableSymbol_r19).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.Lpwus_AvailableSymbol_r19).Choice {
			case LowPowerConfig_r19_Lpwus_AvailableSymbol_r19_OneSlot:
				if err := e.EncodeBitString((*(*ie.Lpwus_AvailableSymbol_r19).OneSlot), per.FixedSize(14)); err != nil {
					return err
				}
			case LowPowerConfig_r19_Lpwus_AvailableSymbol_r19_TwoSlots:
				if err := e.EncodeBitString((*(*ie.Lpwus_AvailableSymbol_r19).TwoSlots), per.FixedSize(28)); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.Lpwus_AvailableSymbol_r19).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 12. lpwus-OffsetFirstMoWithinLo-r19: choice
	{
		if ie.Lpwus_OffsetFirstMoWithinLo_r19 != nil {
			choiceEnc := e.NewChoiceEncoder(lowPowerConfig_r19LpwusOffsetFirstMoWithinLoR19Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.Lpwus_OffsetFirstMoWithinLo_r19).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.Lpwus_OffsetFirstMoWithinLo_r19).Choice {
			case LowPowerConfig_r19_Lpwus_OffsetFirstMoWithinLo_r19_OffsetForSCS_15kHz:
				if err := e.EncodeInteger((*(*ie.Lpwus_OffsetFirstMoWithinLo_r19).OffsetForSCS_15kHz), per.Constrained(0, 139)); err != nil {
					return err
				}
			case LowPowerConfig_r19_Lpwus_OffsetFirstMoWithinLo_r19_OffsetForSCS_30kHz:
				if err := e.EncodeInteger((*(*ie.Lpwus_OffsetFirstMoWithinLo_r19).OffsetForSCS_30kHz), per.Constrained(0, 279)); err != nil {
					return err
				}
			case LowPowerConfig_r19_Lpwus_OffsetFirstMoWithinLo_r19_OffsetForSCS_120kHz:
				if err := e.EncodeInteger((*(*ie.Lpwus_OffsetFirstMoWithinLo_r19).OffsetForSCS_120kHz), per.Constrained(0, 1119)); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.Lpwus_OffsetFirstMoWithinLo_r19).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 13. lpwus-NominalMoDuration-r19: integer
	{
		if ie.Lpwus_NominalMoDuration_r19 != nil {
			if err := e.EncodeInteger(*ie.Lpwus_NominalMoDuration_r19, lowPowerConfigR19LpwusNominalMoDurationR19Constraints); err != nil {
				return err
			}
		}
	}

	// 14. lpwus-ActualDuration-r19: choice
	{
		if ie.Lpwus_ActualDuration_r19 != nil {
			choiceEnc := e.NewChoiceEncoder(lowPowerConfig_r19LpwusActualDurationR19Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.Lpwus_ActualDuration_r19).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.Lpwus_ActualDuration_r19).Choice {
			case LowPowerConfig_r19_Lpwus_ActualDuration_r19_MValue1:
				if err := e.EncodeInteger((*(*ie.Lpwus_ActualDuration_r19).MValue1), per.Constrained(2, 64)); err != nil {
					return err
				}
			case LowPowerConfig_r19_Lpwus_ActualDuration_r19_MValue2:
				if err := e.EncodeInteger((*(*ie.Lpwus_ActualDuration_r19).MValue2), per.Constrained(1, 32)); err != nil {
					return err
				}
			case LowPowerConfig_r19_Lpwus_ActualDuration_r19_MValue4:
				if err := e.EncodeInteger((*(*ie.Lpwus_ActualDuration_r19).MValue4), per.Constrained(1, 16)); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.Lpwus_ActualDuration_r19).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 15. lpwus-LPSS-StartRB-r19: integer
	{
		if ie.Lpwus_LPSS_StartRB_r19 != nil {
			if err := e.EncodeInteger(*ie.Lpwus_LPSS_StartRB_r19, lowPowerConfigR19LpwusLPSSStartRBR19Constraints); err != nil {
				return err
			}
		}
	}

	// 16. lpwus-LPSS-BeamSubset-r19: choice
	{
		if ie.Lpwus_LPSS_BeamSubset_r19 != nil {
			choiceEnc := e.NewChoiceEncoder(lowPowerConfig_r19LpwusLPSSBeamSubsetR19Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.Lpwus_LPSS_BeamSubset_r19).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.Lpwus_LPSS_BeamSubset_r19).Choice {
			case LowPowerConfig_r19_Lpwus_LPSS_BeamSubset_r19_ShortBitmap:
				if err := e.EncodeBitString((*(*ie.Lpwus_LPSS_BeamSubset_r19).ShortBitmap), per.FixedSize(4)); err != nil {
					return err
				}
			case LowPowerConfig_r19_Lpwus_LPSS_BeamSubset_r19_MediumBitmap:
				if err := e.EncodeBitString((*(*ie.Lpwus_LPSS_BeamSubset_r19).MediumBitmap), per.FixedSize(8)); err != nil {
					return err
				}
			case LowPowerConfig_r19_Lpwus_LPSS_BeamSubset_r19_LongBitmap:
				if err := e.EncodeBitString((*(*ie.Lpwus_LPSS_BeamSubset_r19).LongBitmap), per.FixedSize(64)); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.Lpwus_LPSS_BeamSubset_r19).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 17. lpss-EPRE-Ratio-r19: enumerated
	{
		if ie.Lpss_EPRE_Ratio_r19 != nil {
			if err := e.EncodeEnumerated(*ie.Lpss_EPRE_Ratio_r19, lowPowerConfigR19LpssEPRERatioR19Constraints); err != nil {
				return err
			}
		}
	}

	// 18. lpss-BinarySeqIndex-r19: integer
	{
		if ie.Lpss_BinarySeqIndex_r19 != nil {
			if err := e.EncodeInteger(*ie.Lpss_BinarySeqIndex_r19, lowPowerConfigR19LpssBinarySeqIndexR19Constraints); err != nil {
				return err
			}
		}
	}

	// 19. lpss-MvalueAndSeqConfig-r19: choice
	{
		if ie.Lpss_MvalueAndSeqConfig_r19 != nil {
			choiceEnc := e.NewChoiceEncoder(lowPowerConfig_r19LpssMvalueAndSeqConfigR19Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.Lpss_MvalueAndSeqConfig_r19).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.Lpss_MvalueAndSeqConfig_r19).Choice {
			case LowPowerConfig_r19_Lpss_MvalueAndSeqConfig_r19_NOne:
				c := (*(*ie.Lpss_MvalueAndSeqConfig_r19).NOne)
				lowPowerConfigR19LpssMvalueAndSeqConfigR19NOneSeq := e.NewSequenceEncoder(lowPowerConfigR19LpssMvalueAndSeqConfigR19NOneConstraints)
				if err := lowPowerConfigR19LpssMvalueAndSeqConfigR19NOneSeq.EncodePreamble([]bool{c.Lpss_BinarySeqLen_r19 != nil, c.Lpss_OverlaidSeqRoot_r19 != nil}); err != nil {
					return err
				}
				if c.Lpss_BinarySeqLen_r19 != nil {
					if err := e.EncodeEnumerated((*c.Lpss_BinarySeqLen_r19), lowPowerConfigR19LpssMvalueAndSeqConfigR19NOneLpssBinarySeqLenR19Constraints); err != nil {
						return err
					}
				}
				if c.Lpss_OverlaidSeqRoot_r19 != nil {
					if err := e.EncodeInteger((*c.Lpss_OverlaidSeqRoot_r19), per.Constrained(1, 131)); err != nil {
						return err
					}
				}
			case LowPowerConfig_r19_Lpss_MvalueAndSeqConfig_r19_NTwo:
				c := (*(*ie.Lpss_MvalueAndSeqConfig_r19).NTwo)
				lowPowerConfigR19LpssMvalueAndSeqConfigR19NTwoSeq := e.NewSequenceEncoder(lowPowerConfigR19LpssMvalueAndSeqConfigR19NTwoConstraints)
				if err := lowPowerConfigR19LpssMvalueAndSeqConfigR19NTwoSeq.EncodePreamble([]bool{c.Lpss_BinarySeqLen_r19 != nil, c.Lpss_OverlaidSeqRoot_r19 != nil}); err != nil {
					return err
				}
				if c.Lpss_BinarySeqLen_r19 != nil {
					if err := e.EncodeEnumerated((*c.Lpss_BinarySeqLen_r19), lowPowerConfigR19LpssMvalueAndSeqConfigR19NTwoLpssBinarySeqLenR19Constraints); err != nil {
						return err
					}
				}
				if c.Lpss_OverlaidSeqRoot_r19 != nil {
					if err := e.EncodeInteger((*c.Lpss_OverlaidSeqRoot_r19), per.Constrained(1, 61)); err != nil {
						return err
					}
				}
			case LowPowerConfig_r19_Lpss_MvalueAndSeqConfig_r19_NFour:
				c := (*(*ie.Lpss_MvalueAndSeqConfig_r19).NFour)
				lowPowerConfigR19LpssMvalueAndSeqConfigR19NFourSeq := e.NewSequenceEncoder(lowPowerConfigR19LpssMvalueAndSeqConfigR19NFourConstraints)
				if err := lowPowerConfigR19LpssMvalueAndSeqConfigR19NFourSeq.EncodePreamble([]bool{c.Lpss_BinarySeqLen_r19 != nil, c.Lpss_OverlaidSeqRoot_r19 != nil}); err != nil {
					return err
				}
				if c.Lpss_BinarySeqLen_r19 != nil {
					if err := e.EncodeEnumerated((*c.Lpss_BinarySeqLen_r19), lowPowerConfigR19LpssMvalueAndSeqConfigR19NFourLpssBinarySeqLenR19Constraints); err != nil {
						return err
					}
				}
				if c.Lpss_OverlaidSeqRoot_r19 != nil {
					if err := e.EncodeInteger((*c.Lpss_OverlaidSeqRoot_r19), per.Constrained(1, 31)); err != nil {
						return err
					}
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.Lpss_MvalueAndSeqConfig_r19).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 20. lpss-PeriodicityAndOffset-r19: choice
	{
		if ie.Lpss_PeriodicityAndOffset_r19 != nil {
			choiceEnc := e.NewChoiceEncoder(lowPowerConfig_r19LpssPeriodicityAndOffsetR19Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.Lpss_PeriodicityAndOffset_r19).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.Lpss_PeriodicityAndOffset_r19).Choice {
			case LowPowerConfig_r19_Lpss_PeriodicityAndOffset_r19_Ms160:
				if err := e.EncodeInteger((*(*ie.Lpss_PeriodicityAndOffset_r19).Ms160), per.Constrained(0, 159)); err != nil {
					return err
				}
			case LowPowerConfig_r19_Lpss_PeriodicityAndOffset_r19_Ms320:
				if err := e.EncodeInteger((*(*ie.Lpss_PeriodicityAndOffset_r19).Ms320), per.Constrained(0, 319)); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.Lpss_PeriodicityAndOffset_r19).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 21. lpss-StartSymbol-r19: sequence
	{
		if ie.Lpss_StartSymbol_r19 != nil {
			c := ie.Lpss_StartSymbol_r19
			lowPowerConfigR19LpssStartSymbolR19Seq := e.NewSequenceEncoder(lowPowerConfigR19LpssStartSymbolR19Constraints)
			if err := lowPowerConfigR19LpssStartSymbolR19Seq.EncodePreamble([]bool{c.StartSymbol2_r19 != nil}); err != nil {
				return err
			}
			if err := e.EncodeInteger(c.StartSymbol1_r19, per.Constrained(0, 10)); err != nil {
				return err
			}
			if c.StartSymbol2_r19 != nil {
				if err := e.EncodeInteger((*c.StartSymbol2_r19), per.Constrained(0, 10)); err != nil {
					return err
				}
			}
		}
	}

	// 22. lp-SubgroupConfig-r19: ref
	{
		if err := ie.Lp_SubgroupConfig_r19.Encode(e); err != nil {
			return err
		}
	}

	// 23. entryCondition-r19: ref
	{
		if err := ie.EntryCondition_r19.Encode(e); err != nil {
			return err
		}
	}

	// 24. exitCondition-r19: ref
	{
		if err := ie.ExitCondition_r19.Encode(e); err != nil {
			return err
		}
	}

	return nil
}

func (ie *LowPowerConfig_r19) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(lowPowerConfigR19Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. lpwus-MvalueAndSeqConfigFR1-r19: choice
	{
		if seq.IsComponentPresent(0) {
			ie.Lpwus_MvalueAndSeqConfigFR1_r19 = &struct {
				Choice int
				NOne   *struct {
					Lpwus_OverlaidSeqRoot_r19 *struct {
						Root1_r19 int64
						Root2_r19 *int64
					}
					Lpwus_OverlaidSeqNum_r19 *int64
				}
				NTwo *struct {
					Lpwus_OverlaidSeqRoot_r19 *struct {
						Root1_r19 int64
						Root2_r19 *int64
					}
					Lpwus_OverlaidSeqNum_r19 *int64
				}
				NFour *struct {
					Lpwus_OverlaidSeqRoot_r19 *struct {
						Root1_r19 int64
						Root2_r19 *int64
					}
					Lpwus_OverlaidSeqNum_r19 *int64
				}
			}{}
			choiceDec := d.NewChoiceDecoder(lowPowerConfig_r19LpwusMvalueAndSeqConfigFR1R19Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Lpwus_MvalueAndSeqConfigFR1_r19).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case LowPowerConfig_r19_Lpwus_MvalueAndSeqConfigFR1_r19_NOne:
				(*ie.Lpwus_MvalueAndSeqConfigFR1_r19).NOne = &struct {
					Lpwus_OverlaidSeqRoot_r19 *struct {
						Root1_r19 int64
						Root2_r19 *int64
					}
					Lpwus_OverlaidSeqNum_r19 *int64
				}{}
				c := (*(*ie.Lpwus_MvalueAndSeqConfigFR1_r19).NOne)
				lowPowerConfigR19LpwusMvalueAndSeqConfigFR1R19NOneSeq := d.NewSequenceDecoder(lowPowerConfigR19LpwusMvalueAndSeqConfigFR1R19NOneConstraints)
				if err := lowPowerConfigR19LpwusMvalueAndSeqConfigFR1R19NOneSeq.DecodePreamble(); err != nil {
					return err
				}
				if lowPowerConfigR19LpwusMvalueAndSeqConfigFR1R19NOneSeq.IsComponentPresent(0) {
					c.Lpwus_OverlaidSeqRoot_r19 = &struct {
						Root1_r19 int64
						Root2_r19 *int64
					}{}
					c := (*c.Lpwus_OverlaidSeqRoot_r19)
					lowPowerConfigR19LpwusMvalueAndSeqConfigFR1R19NOneLpwusOverlaidSeqRootR19Seq := d.NewSequenceDecoder(lowPowerConfigR19LpwusMvalueAndSeqConfigFR1R19NOneLpwusOverlaidSeqRootR19Constraints)
					if err := lowPowerConfigR19LpwusMvalueAndSeqConfigFR1R19NOneLpwusOverlaidSeqRootR19Seq.DecodePreamble(); err != nil {
						return err
					}
					{
						v, err := d.DecodeInteger(per.Constrained(1, 131))
						if err != nil {
							return err
						}
						c.Root1_r19 = v
					}
					if lowPowerConfigR19LpwusMvalueAndSeqConfigFR1R19NOneLpwusOverlaidSeqRootR19Seq.IsComponentPresent(1) {
						c.Root2_r19 = new(int64)
						v, err := d.DecodeInteger(per.Constrained(1, 131))
						if err != nil {
							return err
						}
						(*c.Root2_r19) = v
					}
				}
				if lowPowerConfigR19LpwusMvalueAndSeqConfigFR1R19NOneSeq.IsComponentPresent(1) {
					c.Lpwus_OverlaidSeqNum_r19 = new(int64)
					v, err := d.DecodeEnumerated(lowPowerConfigR19LpwusMvalueAndSeqConfigFR1R19NOneLpwusOverlaidSeqNumR19Constraints)
					if err != nil {
						return err
					}
					(*c.Lpwus_OverlaidSeqNum_r19) = v
				}
			case LowPowerConfig_r19_Lpwus_MvalueAndSeqConfigFR1_r19_NTwo:
				(*ie.Lpwus_MvalueAndSeqConfigFR1_r19).NTwo = &struct {
					Lpwus_OverlaidSeqRoot_r19 *struct {
						Root1_r19 int64
						Root2_r19 *int64
					}
					Lpwus_OverlaidSeqNum_r19 *int64
				}{}
				c := (*(*ie.Lpwus_MvalueAndSeqConfigFR1_r19).NTwo)
				lowPowerConfigR19LpwusMvalueAndSeqConfigFR1R19NTwoSeq := d.NewSequenceDecoder(lowPowerConfigR19LpwusMvalueAndSeqConfigFR1R19NTwoConstraints)
				if err := lowPowerConfigR19LpwusMvalueAndSeqConfigFR1R19NTwoSeq.DecodePreamble(); err != nil {
					return err
				}
				if lowPowerConfigR19LpwusMvalueAndSeqConfigFR1R19NTwoSeq.IsComponentPresent(0) {
					c.Lpwus_OverlaidSeqRoot_r19 = &struct {
						Root1_r19 int64
						Root2_r19 *int64
					}{}
					c := (*c.Lpwus_OverlaidSeqRoot_r19)
					lowPowerConfigR19LpwusMvalueAndSeqConfigFR1R19NTwoLpwusOverlaidSeqRootR19Seq := d.NewSequenceDecoder(lowPowerConfigR19LpwusMvalueAndSeqConfigFR1R19NTwoLpwusOverlaidSeqRootR19Constraints)
					if err := lowPowerConfigR19LpwusMvalueAndSeqConfigFR1R19NTwoLpwusOverlaidSeqRootR19Seq.DecodePreamble(); err != nil {
						return err
					}
					{
						v, err := d.DecodeInteger(per.Constrained(1, 61))
						if err != nil {
							return err
						}
						c.Root1_r19 = v
					}
					if lowPowerConfigR19LpwusMvalueAndSeqConfigFR1R19NTwoLpwusOverlaidSeqRootR19Seq.IsComponentPresent(1) {
						c.Root2_r19 = new(int64)
						v, err := d.DecodeInteger(per.Constrained(1, 61))
						if err != nil {
							return err
						}
						(*c.Root2_r19) = v
					}
				}
				if lowPowerConfigR19LpwusMvalueAndSeqConfigFR1R19NTwoSeq.IsComponentPresent(1) {
					c.Lpwus_OverlaidSeqNum_r19 = new(int64)
					v, err := d.DecodeEnumerated(lowPowerConfigR19LpwusMvalueAndSeqConfigFR1R19NTwoLpwusOverlaidSeqNumR19Constraints)
					if err != nil {
						return err
					}
					(*c.Lpwus_OverlaidSeqNum_r19) = v
				}
			case LowPowerConfig_r19_Lpwus_MvalueAndSeqConfigFR1_r19_NFour:
				(*ie.Lpwus_MvalueAndSeqConfigFR1_r19).NFour = &struct {
					Lpwus_OverlaidSeqRoot_r19 *struct {
						Root1_r19 int64
						Root2_r19 *int64
					}
					Lpwus_OverlaidSeqNum_r19 *int64
				}{}
				c := (*(*ie.Lpwus_MvalueAndSeqConfigFR1_r19).NFour)
				lowPowerConfigR19LpwusMvalueAndSeqConfigFR1R19NFourSeq := d.NewSequenceDecoder(lowPowerConfigR19LpwusMvalueAndSeqConfigFR1R19NFourConstraints)
				if err := lowPowerConfigR19LpwusMvalueAndSeqConfigFR1R19NFourSeq.DecodePreamble(); err != nil {
					return err
				}
				if lowPowerConfigR19LpwusMvalueAndSeqConfigFR1R19NFourSeq.IsComponentPresent(0) {
					c.Lpwus_OverlaidSeqRoot_r19 = &struct {
						Root1_r19 int64
						Root2_r19 *int64
					}{}
					c := (*c.Lpwus_OverlaidSeqRoot_r19)
					lowPowerConfigR19LpwusMvalueAndSeqConfigFR1R19NFourLpwusOverlaidSeqRootR19Seq := d.NewSequenceDecoder(lowPowerConfigR19LpwusMvalueAndSeqConfigFR1R19NFourLpwusOverlaidSeqRootR19Constraints)
					if err := lowPowerConfigR19LpwusMvalueAndSeqConfigFR1R19NFourLpwusOverlaidSeqRootR19Seq.DecodePreamble(); err != nil {
						return err
					}
					{
						v, err := d.DecodeInteger(per.Constrained(1, 31))
						if err != nil {
							return err
						}
						c.Root1_r19 = v
					}
					if lowPowerConfigR19LpwusMvalueAndSeqConfigFR1R19NFourLpwusOverlaidSeqRootR19Seq.IsComponentPresent(1) {
						c.Root2_r19 = new(int64)
						v, err := d.DecodeInteger(per.Constrained(1, 31))
						if err != nil {
							return err
						}
						(*c.Root2_r19) = v
					}
				}
				if lowPowerConfigR19LpwusMvalueAndSeqConfigFR1R19NFourSeq.IsComponentPresent(1) {
					c.Lpwus_OverlaidSeqNum_r19 = new(int64)
					v, err := d.DecodeEnumerated(lowPowerConfigR19LpwusMvalueAndSeqConfigFR1R19NFourLpwusOverlaidSeqNumR19Constraints)
					if err != nil {
						return err
					}
					(*c.Lpwus_OverlaidSeqNum_r19) = v
				}
			}
		}
	}

	// 4. lpwus-MvalueAndSeqConfigFR2-r19: choice
	{
		if seq.IsComponentPresent(1) {
			ie.Lpwus_MvalueAndSeqConfigFR2_r19 = &struct {
				Choice int
				NOne   *struct {
					Lpwus_OverlaidSeqRoot_r19 *struct {
						Root1_r19 int64
						Root2_r19 *int64
					}
					Lpwus_OverlaidSeqNum_SCS_120kHz_r19 *int64
				}
			}{}
			choiceDec := d.NewChoiceDecoder(lowPowerConfig_r19LpwusMvalueAndSeqConfigFR2R19Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Lpwus_MvalueAndSeqConfigFR2_r19).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case LowPowerConfig_r19_Lpwus_MvalueAndSeqConfigFR2_r19_NOne:
				(*ie.Lpwus_MvalueAndSeqConfigFR2_r19).NOne = &struct {
					Lpwus_OverlaidSeqRoot_r19 *struct {
						Root1_r19 int64
						Root2_r19 *int64
					}
					Lpwus_OverlaidSeqNum_SCS_120kHz_r19 *int64
				}{}
				c := (*(*ie.Lpwus_MvalueAndSeqConfigFR2_r19).NOne)
				lowPowerConfigR19LpwusMvalueAndSeqConfigFR2R19NOneSeq := d.NewSequenceDecoder(lowPowerConfigR19LpwusMvalueAndSeqConfigFR2R19NOneConstraints)
				if err := lowPowerConfigR19LpwusMvalueAndSeqConfigFR2R19NOneSeq.DecodePreamble(); err != nil {
					return err
				}
				if lowPowerConfigR19LpwusMvalueAndSeqConfigFR2R19NOneSeq.IsComponentPresent(0) {
					c.Lpwus_OverlaidSeqRoot_r19 = &struct {
						Root1_r19 int64
						Root2_r19 *int64
					}{}
					c := (*c.Lpwus_OverlaidSeqRoot_r19)
					lowPowerConfigR19LpwusMvalueAndSeqConfigFR2R19NOneLpwusOverlaidSeqRootR19Seq := d.NewSequenceDecoder(lowPowerConfigR19LpwusMvalueAndSeqConfigFR2R19NOneLpwusOverlaidSeqRootR19Constraints)
					if err := lowPowerConfigR19LpwusMvalueAndSeqConfigFR2R19NOneLpwusOverlaidSeqRootR19Seq.DecodePreamble(); err != nil {
						return err
					}
					{
						v, err := d.DecodeInteger(per.Constrained(1, 131))
						if err != nil {
							return err
						}
						c.Root1_r19 = v
					}
					if lowPowerConfigR19LpwusMvalueAndSeqConfigFR2R19NOneLpwusOverlaidSeqRootR19Seq.IsComponentPresent(1) {
						c.Root2_r19 = new(int64)
						v, err := d.DecodeInteger(per.Constrained(1, 131))
						if err != nil {
							return err
						}
						(*c.Root2_r19) = v
					}
				}
				if lowPowerConfigR19LpwusMvalueAndSeqConfigFR2R19NOneSeq.IsComponentPresent(1) {
					c.Lpwus_OverlaidSeqNum_SCS_120kHz_r19 = new(int64)
					v, err := d.DecodeEnumerated(lowPowerConfigR19LpwusMvalueAndSeqConfigFR2R19NOneLpwusOverlaidSeqNumSCS120kHzR19Constraints)
					if err != nil {
						return err
					}
					(*c.Lpwus_OverlaidSeqNum_SCS_120kHz_r19) = v
				}
			}
		}
	}

	// 5. lpwus-LO-FrameOffsetList-r19: sequence
	{
		{
			c := &ie.Lpwus_LO_FrameOffsetList_r19
			lowPowerConfigR19LpwusLOFrameOffsetListR19Seq := d.NewSequenceDecoder(lowPowerConfigR19LpwusLOFrameOffsetListR19Constraints)
			if err := lowPowerConfigR19LpwusLOFrameOffsetListR19Seq.DecodePreamble(); err != nil {
				return err
			}
			if lowPowerConfigR19LpwusLOFrameOffsetListR19Seq.IsComponentPresent(0) {
				seqOf := d.NewSequenceOfDecoder(lowPowerConfigR19LpwusLOFrameOffsetListR19OffsetForLongerWakeUpDelayR19Constraints)
				n, err := seqOf.DecodeLength()
				if err != nil {
					return err
				}
				c.OffsetForLongerWakeUpDelay_r19 = make([]int64, n)
				for i := int64(0); i < n; i++ {
					v, err := d.DecodeInteger(per.Constrained(8, 200))
					if err != nil {
						return err
					}
					c.OffsetForLongerWakeUpDelay_r19[i] = v
				}
			}
			if lowPowerConfigR19LpwusLOFrameOffsetListR19Seq.IsComponentPresent(1) {
				seqOf := d.NewSequenceOfDecoder(lowPowerConfigR19LpwusLOFrameOffsetListR19OffsetForShorterWakeUpDelayR19Constraints)
				n, err := seqOf.DecodeLength()
				if err != nil {
					return err
				}
				c.OffsetForShorterWakeUpDelay_r19 = make([]int64, n)
				for i := int64(0); i < n; i++ {
					v, err := d.DecodeInteger(per.Constrained(8, 200))
					if err != nil {
						return err
					}
					c.OffsetForShorterWakeUpDelay_r19[i] = v
				}
			}
		}
	}

	// 6. lpwus-LO-FrameOffsetListForPagingAdapt-r19: sequence
	{
		{
			c := &ie.Lpwus_LO_FrameOffsetListForPagingAdapt_r19
			lowPowerConfigR19LpwusLOFrameOffsetListForPagingAdaptR19Seq := d.NewSequenceDecoder(lowPowerConfigR19LpwusLOFrameOffsetListForPagingAdaptR19Constraints)
			if err := lowPowerConfigR19LpwusLOFrameOffsetListForPagingAdaptR19Seq.DecodePreamble(); err != nil {
				return err
			}
			if lowPowerConfigR19LpwusLOFrameOffsetListForPagingAdaptR19Seq.IsComponentPresent(0) {
				seqOf := d.NewSequenceOfDecoder(lowPowerConfigR19LpwusLOFrameOffsetListForPagingAdaptR19OffsetForLongerWakeUpDelayR19Constraints)
				n, err := seqOf.DecodeLength()
				if err != nil {
					return err
				}
				c.OffsetForLongerWakeUpDelay_r19 = make([]int64, n)
				for i := int64(0); i < n; i++ {
					v, err := d.DecodeInteger(per.Constrained(8, 200))
					if err != nil {
						return err
					}
					c.OffsetForLongerWakeUpDelay_r19[i] = v
				}
			}
			if lowPowerConfigR19LpwusLOFrameOffsetListForPagingAdaptR19Seq.IsComponentPresent(1) {
				seqOf := d.NewSequenceOfDecoder(lowPowerConfigR19LpwusLOFrameOffsetListForPagingAdaptR19OffsetForShorterWakeUpDelayR19Constraints)
				n, err := seqOf.DecodeLength()
				if err != nil {
					return err
				}
				c.OffsetForShorterWakeUpDelay_r19 = make([]int64, n)
				for i := int64(0); i < n; i++ {
					v, err := d.DecodeInteger(per.Constrained(8, 200))
					if err != nil {
						return err
					}
					c.OffsetForShorterWakeUpDelay_r19[i] = v
				}
			}
		}
	}

	// 7. lpwus-MO-NumPerLO-r19: enumerated
	{
		v4, err := d.DecodeEnumerated(lowPowerConfigR19LpwusMONumPerLOR19Constraints)
		if err != nil {
			return err
		}
		ie.Lpwus_MO_NumPerLO_r19 = v4
	}

	// 8. lpwus-PO-NumPerLO-r19: enumerated
	{
		if seq.IsComponentPresent(5) {
			idx, err := d.DecodeEnumerated(lowPowerConfigR19LpwusPONumPerLOR19Constraints)
			if err != nil {
				return err
			}
			ie.Lpwus_PO_NumPerLO_r19 = &idx
		}
	}

	// 9. lpwus-EPRE-Ratio-r19: enumerated
	{
		if seq.IsComponentPresent(6) {
			idx, err := d.DecodeEnumerated(lowPowerConfigR19LpwusEPRERatioR19Constraints)
			if err != nil {
				return err
			}
			ie.Lpwus_EPRE_Ratio_r19 = &idx
		}
	}

	// 10. lpwus-AvailableSlot-r19: choice
	{
		if seq.IsComponentPresent(7) {
			ie.Lpwus_AvailableSlot_r19 = &struct {
				Choice int
				N10    *per.BitString
				N20    *per.BitString
				N40    *per.BitString
			}{}
			choiceDec := d.NewChoiceDecoder(lowPowerConfig_r19LpwusAvailableSlotR19Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Lpwus_AvailableSlot_r19).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case LowPowerConfig_r19_Lpwus_AvailableSlot_r19_N10:
				(*ie.Lpwus_AvailableSlot_r19).N10 = new(per.BitString)
				v, err := d.DecodeBitString(per.FixedSize(10))
				if err != nil {
					return err
				}
				(*(*ie.Lpwus_AvailableSlot_r19).N10) = v
			case LowPowerConfig_r19_Lpwus_AvailableSlot_r19_N20:
				(*ie.Lpwus_AvailableSlot_r19).N20 = new(per.BitString)
				v, err := d.DecodeBitString(per.FixedSize(20))
				if err != nil {
					return err
				}
				(*(*ie.Lpwus_AvailableSlot_r19).N20) = v
			case LowPowerConfig_r19_Lpwus_AvailableSlot_r19_N40:
				(*ie.Lpwus_AvailableSlot_r19).N40 = new(per.BitString)
				v, err := d.DecodeBitString(per.FixedSize(40))
				if err != nil {
					return err
				}
				(*(*ie.Lpwus_AvailableSlot_r19).N40) = v
			}
		}
	}

	// 11. lpwus-AvailableSymbol-r19: choice
	{
		if seq.IsComponentPresent(8) {
			ie.Lpwus_AvailableSymbol_r19 = &struct {
				Choice   int
				OneSlot  *per.BitString
				TwoSlots *per.BitString
			}{}
			choiceDec := d.NewChoiceDecoder(lowPowerConfig_r19LpwusAvailableSymbolR19Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Lpwus_AvailableSymbol_r19).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case LowPowerConfig_r19_Lpwus_AvailableSymbol_r19_OneSlot:
				(*ie.Lpwus_AvailableSymbol_r19).OneSlot = new(per.BitString)
				v, err := d.DecodeBitString(per.FixedSize(14))
				if err != nil {
					return err
				}
				(*(*ie.Lpwus_AvailableSymbol_r19).OneSlot) = v
			case LowPowerConfig_r19_Lpwus_AvailableSymbol_r19_TwoSlots:
				(*ie.Lpwus_AvailableSymbol_r19).TwoSlots = new(per.BitString)
				v, err := d.DecodeBitString(per.FixedSize(28))
				if err != nil {
					return err
				}
				(*(*ie.Lpwus_AvailableSymbol_r19).TwoSlots) = v
			}
		}
	}

	// 12. lpwus-OffsetFirstMoWithinLo-r19: choice
	{
		if seq.IsComponentPresent(9) {
			ie.Lpwus_OffsetFirstMoWithinLo_r19 = &struct {
				Choice              int
				OffsetForSCS_15kHz  *int64
				OffsetForSCS_30kHz  *int64
				OffsetForSCS_120kHz *int64
			}{}
			choiceDec := d.NewChoiceDecoder(lowPowerConfig_r19LpwusOffsetFirstMoWithinLoR19Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Lpwus_OffsetFirstMoWithinLo_r19).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case LowPowerConfig_r19_Lpwus_OffsetFirstMoWithinLo_r19_OffsetForSCS_15kHz:
				(*ie.Lpwus_OffsetFirstMoWithinLo_r19).OffsetForSCS_15kHz = new(int64)
				v, err := d.DecodeInteger(per.Constrained(0, 139))
				if err != nil {
					return err
				}
				(*(*ie.Lpwus_OffsetFirstMoWithinLo_r19).OffsetForSCS_15kHz) = v
			case LowPowerConfig_r19_Lpwus_OffsetFirstMoWithinLo_r19_OffsetForSCS_30kHz:
				(*ie.Lpwus_OffsetFirstMoWithinLo_r19).OffsetForSCS_30kHz = new(int64)
				v, err := d.DecodeInteger(per.Constrained(0, 279))
				if err != nil {
					return err
				}
				(*(*ie.Lpwus_OffsetFirstMoWithinLo_r19).OffsetForSCS_30kHz) = v
			case LowPowerConfig_r19_Lpwus_OffsetFirstMoWithinLo_r19_OffsetForSCS_120kHz:
				(*ie.Lpwus_OffsetFirstMoWithinLo_r19).OffsetForSCS_120kHz = new(int64)
				v, err := d.DecodeInteger(per.Constrained(0, 1119))
				if err != nil {
					return err
				}
				(*(*ie.Lpwus_OffsetFirstMoWithinLo_r19).OffsetForSCS_120kHz) = v
			}
		}
	}

	// 13. lpwus-NominalMoDuration-r19: integer
	{
		if seq.IsComponentPresent(10) {
			v, err := d.DecodeInteger(lowPowerConfigR19LpwusNominalMoDurationR19Constraints)
			if err != nil {
				return err
			}
			ie.Lpwus_NominalMoDuration_r19 = &v
		}
	}

	// 14. lpwus-ActualDuration-r19: choice
	{
		if seq.IsComponentPresent(11) {
			ie.Lpwus_ActualDuration_r19 = &struct {
				Choice  int
				MValue1 *int64
				MValue2 *int64
				MValue4 *int64
			}{}
			choiceDec := d.NewChoiceDecoder(lowPowerConfig_r19LpwusActualDurationR19Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Lpwus_ActualDuration_r19).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case LowPowerConfig_r19_Lpwus_ActualDuration_r19_MValue1:
				(*ie.Lpwus_ActualDuration_r19).MValue1 = new(int64)
				v, err := d.DecodeInteger(per.Constrained(2, 64))
				if err != nil {
					return err
				}
				(*(*ie.Lpwus_ActualDuration_r19).MValue1) = v
			case LowPowerConfig_r19_Lpwus_ActualDuration_r19_MValue2:
				(*ie.Lpwus_ActualDuration_r19).MValue2 = new(int64)
				v, err := d.DecodeInteger(per.Constrained(1, 32))
				if err != nil {
					return err
				}
				(*(*ie.Lpwus_ActualDuration_r19).MValue2) = v
			case LowPowerConfig_r19_Lpwus_ActualDuration_r19_MValue4:
				(*ie.Lpwus_ActualDuration_r19).MValue4 = new(int64)
				v, err := d.DecodeInteger(per.Constrained(1, 16))
				if err != nil {
					return err
				}
				(*(*ie.Lpwus_ActualDuration_r19).MValue4) = v
			}
		}
	}

	// 15. lpwus-LPSS-StartRB-r19: integer
	{
		if seq.IsComponentPresent(12) {
			v, err := d.DecodeInteger(lowPowerConfigR19LpwusLPSSStartRBR19Constraints)
			if err != nil {
				return err
			}
			ie.Lpwus_LPSS_StartRB_r19 = &v
		}
	}

	// 16. lpwus-LPSS-BeamSubset-r19: choice
	{
		if seq.IsComponentPresent(13) {
			ie.Lpwus_LPSS_BeamSubset_r19 = &struct {
				Choice       int
				ShortBitmap  *per.BitString
				MediumBitmap *per.BitString
				LongBitmap   *per.BitString
			}{}
			choiceDec := d.NewChoiceDecoder(lowPowerConfig_r19LpwusLPSSBeamSubsetR19Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Lpwus_LPSS_BeamSubset_r19).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case LowPowerConfig_r19_Lpwus_LPSS_BeamSubset_r19_ShortBitmap:
				(*ie.Lpwus_LPSS_BeamSubset_r19).ShortBitmap = new(per.BitString)
				v, err := d.DecodeBitString(per.FixedSize(4))
				if err != nil {
					return err
				}
				(*(*ie.Lpwus_LPSS_BeamSubset_r19).ShortBitmap) = v
			case LowPowerConfig_r19_Lpwus_LPSS_BeamSubset_r19_MediumBitmap:
				(*ie.Lpwus_LPSS_BeamSubset_r19).MediumBitmap = new(per.BitString)
				v, err := d.DecodeBitString(per.FixedSize(8))
				if err != nil {
					return err
				}
				(*(*ie.Lpwus_LPSS_BeamSubset_r19).MediumBitmap) = v
			case LowPowerConfig_r19_Lpwus_LPSS_BeamSubset_r19_LongBitmap:
				(*ie.Lpwus_LPSS_BeamSubset_r19).LongBitmap = new(per.BitString)
				v, err := d.DecodeBitString(per.FixedSize(64))
				if err != nil {
					return err
				}
				(*(*ie.Lpwus_LPSS_BeamSubset_r19).LongBitmap) = v
			}
		}
	}

	// 17. lpss-EPRE-Ratio-r19: enumerated
	{
		if seq.IsComponentPresent(14) {
			idx, err := d.DecodeEnumerated(lowPowerConfigR19LpssEPRERatioR19Constraints)
			if err != nil {
				return err
			}
			ie.Lpss_EPRE_Ratio_r19 = &idx
		}
	}

	// 18. lpss-BinarySeqIndex-r19: integer
	{
		if seq.IsComponentPresent(15) {
			v, err := d.DecodeInteger(lowPowerConfigR19LpssBinarySeqIndexR19Constraints)
			if err != nil {
				return err
			}
			ie.Lpss_BinarySeqIndex_r19 = &v
		}
	}

	// 19. lpss-MvalueAndSeqConfig-r19: choice
	{
		if seq.IsComponentPresent(16) {
			ie.Lpss_MvalueAndSeqConfig_r19 = &struct {
				Choice int
				NOne   *struct {
					Lpss_BinarySeqLen_r19    *int64
					Lpss_OverlaidSeqRoot_r19 *int64
				}
				NTwo *struct {
					Lpss_BinarySeqLen_r19    *int64
					Lpss_OverlaidSeqRoot_r19 *int64
				}
				NFour *struct {
					Lpss_BinarySeqLen_r19    *int64
					Lpss_OverlaidSeqRoot_r19 *int64
				}
			}{}
			choiceDec := d.NewChoiceDecoder(lowPowerConfig_r19LpssMvalueAndSeqConfigR19Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Lpss_MvalueAndSeqConfig_r19).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case LowPowerConfig_r19_Lpss_MvalueAndSeqConfig_r19_NOne:
				(*ie.Lpss_MvalueAndSeqConfig_r19).NOne = &struct {
					Lpss_BinarySeqLen_r19    *int64
					Lpss_OverlaidSeqRoot_r19 *int64
				}{}
				c := (*(*ie.Lpss_MvalueAndSeqConfig_r19).NOne)
				lowPowerConfigR19LpssMvalueAndSeqConfigR19NOneSeq := d.NewSequenceDecoder(lowPowerConfigR19LpssMvalueAndSeqConfigR19NOneConstraints)
				if err := lowPowerConfigR19LpssMvalueAndSeqConfigR19NOneSeq.DecodePreamble(); err != nil {
					return err
				}
				if lowPowerConfigR19LpssMvalueAndSeqConfigR19NOneSeq.IsComponentPresent(0) {
					c.Lpss_BinarySeqLen_r19 = new(int64)
					v, err := d.DecodeEnumerated(lowPowerConfigR19LpssMvalueAndSeqConfigR19NOneLpssBinarySeqLenR19Constraints)
					if err != nil {
						return err
					}
					(*c.Lpss_BinarySeqLen_r19) = v
				}
				if lowPowerConfigR19LpssMvalueAndSeqConfigR19NOneSeq.IsComponentPresent(1) {
					c.Lpss_OverlaidSeqRoot_r19 = new(int64)
					v, err := d.DecodeInteger(per.Constrained(1, 131))
					if err != nil {
						return err
					}
					(*c.Lpss_OverlaidSeqRoot_r19) = v
				}
			case LowPowerConfig_r19_Lpss_MvalueAndSeqConfig_r19_NTwo:
				(*ie.Lpss_MvalueAndSeqConfig_r19).NTwo = &struct {
					Lpss_BinarySeqLen_r19    *int64
					Lpss_OverlaidSeqRoot_r19 *int64
				}{}
				c := (*(*ie.Lpss_MvalueAndSeqConfig_r19).NTwo)
				lowPowerConfigR19LpssMvalueAndSeqConfigR19NTwoSeq := d.NewSequenceDecoder(lowPowerConfigR19LpssMvalueAndSeqConfigR19NTwoConstraints)
				if err := lowPowerConfigR19LpssMvalueAndSeqConfigR19NTwoSeq.DecodePreamble(); err != nil {
					return err
				}
				if lowPowerConfigR19LpssMvalueAndSeqConfigR19NTwoSeq.IsComponentPresent(0) {
					c.Lpss_BinarySeqLen_r19 = new(int64)
					v, err := d.DecodeEnumerated(lowPowerConfigR19LpssMvalueAndSeqConfigR19NTwoLpssBinarySeqLenR19Constraints)
					if err != nil {
						return err
					}
					(*c.Lpss_BinarySeqLen_r19) = v
				}
				if lowPowerConfigR19LpssMvalueAndSeqConfigR19NTwoSeq.IsComponentPresent(1) {
					c.Lpss_OverlaidSeqRoot_r19 = new(int64)
					v, err := d.DecodeInteger(per.Constrained(1, 61))
					if err != nil {
						return err
					}
					(*c.Lpss_OverlaidSeqRoot_r19) = v
				}
			case LowPowerConfig_r19_Lpss_MvalueAndSeqConfig_r19_NFour:
				(*ie.Lpss_MvalueAndSeqConfig_r19).NFour = &struct {
					Lpss_BinarySeqLen_r19    *int64
					Lpss_OverlaidSeqRoot_r19 *int64
				}{}
				c := (*(*ie.Lpss_MvalueAndSeqConfig_r19).NFour)
				lowPowerConfigR19LpssMvalueAndSeqConfigR19NFourSeq := d.NewSequenceDecoder(lowPowerConfigR19LpssMvalueAndSeqConfigR19NFourConstraints)
				if err := lowPowerConfigR19LpssMvalueAndSeqConfigR19NFourSeq.DecodePreamble(); err != nil {
					return err
				}
				if lowPowerConfigR19LpssMvalueAndSeqConfigR19NFourSeq.IsComponentPresent(0) {
					c.Lpss_BinarySeqLen_r19 = new(int64)
					v, err := d.DecodeEnumerated(lowPowerConfigR19LpssMvalueAndSeqConfigR19NFourLpssBinarySeqLenR19Constraints)
					if err != nil {
						return err
					}
					(*c.Lpss_BinarySeqLen_r19) = v
				}
				if lowPowerConfigR19LpssMvalueAndSeqConfigR19NFourSeq.IsComponentPresent(1) {
					c.Lpss_OverlaidSeqRoot_r19 = new(int64)
					v, err := d.DecodeInteger(per.Constrained(1, 31))
					if err != nil {
						return err
					}
					(*c.Lpss_OverlaidSeqRoot_r19) = v
				}
			}
		}
	}

	// 20. lpss-PeriodicityAndOffset-r19: choice
	{
		if seq.IsComponentPresent(17) {
			ie.Lpss_PeriodicityAndOffset_r19 = &struct {
				Choice int
				Ms160  *int64
				Ms320  *int64
			}{}
			choiceDec := d.NewChoiceDecoder(lowPowerConfig_r19LpssPeriodicityAndOffsetR19Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Lpss_PeriodicityAndOffset_r19).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case LowPowerConfig_r19_Lpss_PeriodicityAndOffset_r19_Ms160:
				(*ie.Lpss_PeriodicityAndOffset_r19).Ms160 = new(int64)
				v, err := d.DecodeInteger(per.Constrained(0, 159))
				if err != nil {
					return err
				}
				(*(*ie.Lpss_PeriodicityAndOffset_r19).Ms160) = v
			case LowPowerConfig_r19_Lpss_PeriodicityAndOffset_r19_Ms320:
				(*ie.Lpss_PeriodicityAndOffset_r19).Ms320 = new(int64)
				v, err := d.DecodeInteger(per.Constrained(0, 319))
				if err != nil {
					return err
				}
				(*(*ie.Lpss_PeriodicityAndOffset_r19).Ms320) = v
			}
		}
	}

	// 21. lpss-StartSymbol-r19: sequence
	{
		if seq.IsComponentPresent(18) {
			ie.Lpss_StartSymbol_r19 = &struct {
				StartSymbol1_r19 int64
				StartSymbol2_r19 *int64
			}{}
			c := ie.Lpss_StartSymbol_r19
			lowPowerConfigR19LpssStartSymbolR19Seq := d.NewSequenceDecoder(lowPowerConfigR19LpssStartSymbolR19Constraints)
			if err := lowPowerConfigR19LpssStartSymbolR19Seq.DecodePreamble(); err != nil {
				return err
			}
			{
				v, err := d.DecodeInteger(per.Constrained(0, 10))
				if err != nil {
					return err
				}
				c.StartSymbol1_r19 = v
			}
			if lowPowerConfigR19LpssStartSymbolR19Seq.IsComponentPresent(1) {
				c.StartSymbol2_r19 = new(int64)
				v, err := d.DecodeInteger(per.Constrained(0, 10))
				if err != nil {
					return err
				}
				(*c.StartSymbol2_r19) = v
			}
		}
	}

	// 22. lp-SubgroupConfig-r19: ref
	{
		if err := ie.Lp_SubgroupConfig_r19.Decode(d); err != nil {
			return err
		}
	}

	// 23. entryCondition-r19: ref
	{
		if err := ie.EntryCondition_r19.Decode(d); err != nil {
			return err
		}
	}

	// 24. exitCondition-r19: ref
	{
		if err := ie.ExitCondition_r19.Decode(d); err != nil {
			return err
		}
	}

	return nil
}
