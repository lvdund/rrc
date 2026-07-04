// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: LPWUS-Config-r19 (line 11767).

var lPWUSConfigR19Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "lpwus-MvalueAndSeqConfigFR1-r19", Optional: true},
		{Name: "lpwus-MvalueAndSeqConfigFR2-r19", Optional: true},
		{Name: "lpwus-StartRB-r19", Optional: true},
		{Name: "lpwus-TimeOffset1-1-r19", Optional: true},
		{Name: "lpwus-TimeOffset1-2-r19", Optional: true},
		{Name: "lpwus-MO-1-1-r19", Optional: true},
		{Name: "lpwus-MO-1-2-r19", Optional: true},
		{Name: "lpwus-NumOfMO-1-1-r19", Optional: true},
		{Name: "lpwus-NumOfMO-1-2-r19", Optional: true},
		{Name: "lpwus-TCI-States-r19", Optional: true},
		{Name: "lpwus-NominalMoDuration-r19", Optional: true},
		{Name: "lpwus-ActualDuration-r19", Optional: true},
		{Name: "lpwus-AvailableSlot-r19", Optional: true},
		{Name: "lpwus-AvailableSymbol-r19", Optional: true},
		{Name: "lpwus-TransmitOtherPeriodicCSI-r19", Optional: true},
		{Name: "lpwus-TransmitPeriodicL1-RSRP-r19", Optional: true},
		{Name: "lpwus-Codepoint-r19"},
		{Name: "lpwus-NumOfBits-r19", Optional: true},
		{Name: "lpwus-PDCCH-MonitoringTimer-r19", Optional: true},
	},
}

var lPWUS_Config_r19LpwusMvalueAndSeqConfigFR1R19Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "nOne"},
		{Name: "nTwo"},
		{Name: "nFour"},
	},
}

const (
	LPWUS_Config_r19_Lpwus_MvalueAndSeqConfigFR1_r19_NOne  = 0
	LPWUS_Config_r19_Lpwus_MvalueAndSeqConfigFR1_r19_NTwo  = 1
	LPWUS_Config_r19_Lpwus_MvalueAndSeqConfigFR1_r19_NFour = 2
)

var lPWUS_Config_r19LpwusMvalueAndSeqConfigFR2R19Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "nOne"},
		{Name: "nTwo"},
	},
}

const (
	LPWUS_Config_r19_Lpwus_MvalueAndSeqConfigFR2_r19_NOne = 0
	LPWUS_Config_r19_Lpwus_MvalueAndSeqConfigFR2_r19_NTwo = 1
)

var lPWUSConfigR19LpwusStartRBR19Constraints = per.Constrained(0, 263)

var lPWUSConfigR19LpwusTimeOffset11R19Constraints = per.Constrained(41, 592)

var lPWUSConfigR19LpwusTimeOffset12R19Constraints = per.Constrained(41, 592)

var lPWUS_Config_r19LpwusMO11R19Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "sl1"},
		{Name: "sl2"},
		{Name: "sl4"},
		{Name: "sl5"},
		{Name: "sl8"},
		{Name: "sl10"},
		{Name: "sl16"},
		{Name: "sl20"},
		{Name: "sl40"},
		{Name: "sl80"},
		{Name: "sl160"},
		{Name: "sl320"},
		{Name: "sl640"},
		{Name: "sl1280"},
		{Name: "sl2560"},
		{Name: "sl5120"},
		{Name: "sl10240"},
	},
}

const (
	LPWUS_Config_r19_Lpwus_MO_1_1_r19_Sl1     = 0
	LPWUS_Config_r19_Lpwus_MO_1_1_r19_Sl2     = 1
	LPWUS_Config_r19_Lpwus_MO_1_1_r19_Sl4     = 2
	LPWUS_Config_r19_Lpwus_MO_1_1_r19_Sl5     = 3
	LPWUS_Config_r19_Lpwus_MO_1_1_r19_Sl8     = 4
	LPWUS_Config_r19_Lpwus_MO_1_1_r19_Sl10    = 5
	LPWUS_Config_r19_Lpwus_MO_1_1_r19_Sl16    = 6
	LPWUS_Config_r19_Lpwus_MO_1_1_r19_Sl20    = 7
	LPWUS_Config_r19_Lpwus_MO_1_1_r19_Sl40    = 8
	LPWUS_Config_r19_Lpwus_MO_1_1_r19_Sl80    = 9
	LPWUS_Config_r19_Lpwus_MO_1_1_r19_Sl160   = 10
	LPWUS_Config_r19_Lpwus_MO_1_1_r19_Sl320   = 11
	LPWUS_Config_r19_Lpwus_MO_1_1_r19_Sl640   = 12
	LPWUS_Config_r19_Lpwus_MO_1_1_r19_Sl1280  = 13
	LPWUS_Config_r19_Lpwus_MO_1_1_r19_Sl2560  = 14
	LPWUS_Config_r19_Lpwus_MO_1_1_r19_Sl5120  = 15
	LPWUS_Config_r19_Lpwus_MO_1_1_r19_Sl10240 = 16
)

var lPWUS_Config_r19LpwusMO12R19Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "sl1"},
		{Name: "sl2"},
		{Name: "sl4"},
		{Name: "sl5"},
		{Name: "sl8"},
		{Name: "sl10"},
		{Name: "sl16"},
		{Name: "sl20"},
		{Name: "sl40"},
		{Name: "sl80"},
	},
}

const (
	LPWUS_Config_r19_Lpwus_MO_1_2_r19_Sl1  = 0
	LPWUS_Config_r19_Lpwus_MO_1_2_r19_Sl2  = 1
	LPWUS_Config_r19_Lpwus_MO_1_2_r19_Sl4  = 2
	LPWUS_Config_r19_Lpwus_MO_1_2_r19_Sl5  = 3
	LPWUS_Config_r19_Lpwus_MO_1_2_r19_Sl8  = 4
	LPWUS_Config_r19_Lpwus_MO_1_2_r19_Sl10 = 5
	LPWUS_Config_r19_Lpwus_MO_1_2_r19_Sl16 = 6
	LPWUS_Config_r19_Lpwus_MO_1_2_r19_Sl20 = 7
	LPWUS_Config_r19_Lpwus_MO_1_2_r19_Sl40 = 8
	LPWUS_Config_r19_Lpwus_MO_1_2_r19_Sl80 = 9
)

var lPWUSConfigR19LpwusNumOfMO11R19Constraints = per.Constrained(1, 4)

var lPWUSConfigR19LpwusNumOfMO12R19Constraints = per.Constrained(1, 4)

var lPWUSConfigR19LpwusTCIStatesR19Constraints = per.SizeRange(1, common.MaxNrofBWPs)

var lPWUSConfigR19LpwusNominalMoDurationR19Constraints = per.Constrained(1, 98)

var lPWUS_Config_r19LpwusActualDurationR19Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "mValue1"},
		{Name: "mValue2"},
		{Name: "mValue4"},
	},
}

const (
	LPWUS_Config_r19_Lpwus_ActualDuration_r19_MValue1 = 0
	LPWUS_Config_r19_Lpwus_ActualDuration_r19_MValue2 = 1
	LPWUS_Config_r19_Lpwus_ActualDuration_r19_MValue4 = 2
)

var lPWUS_Config_r19LpwusAvailableSlotR19Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "n10"},
		{Name: "n20"},
		{Name: "n40"},
	},
}

const (
	LPWUS_Config_r19_Lpwus_AvailableSlot_r19_N10 = 0
	LPWUS_Config_r19_Lpwus_AvailableSlot_r19_N20 = 1
	LPWUS_Config_r19_Lpwus_AvailableSlot_r19_N40 = 2
)

var lPWUS_Config_r19LpwusAvailableSymbolR19Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "oneSlot"},
		{Name: "twoSlots"},
	},
}

const (
	LPWUS_Config_r19_Lpwus_AvailableSymbol_r19_OneSlot  = 0
	LPWUS_Config_r19_Lpwus_AvailableSymbol_r19_TwoSlots = 1
)

const (
	LPWUS_Config_r19_Lpwus_TransmitOtherPeriodicCSI_r19_True = 0
)

var lPWUSConfigR19LpwusTransmitOtherPeriodicCSIR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{LPWUS_Config_r19_Lpwus_TransmitOtherPeriodicCSI_r19_True},
}

const (
	LPWUS_Config_r19_Lpwus_TransmitPeriodicL1_RSRP_r19_True = 0
)

var lPWUSConfigR19LpwusTransmitPeriodicL1RSRPR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{LPWUS_Config_r19_Lpwus_TransmitPeriodicL1_RSRP_r19_True},
}

var lPWUSConfigR19LpwusCodepointR19Constraints = per.SizeRange(1, 8)

var lPWUSConfigR19LpwusNumOfBitsR19Constraints = per.Constrained(1, 5)

var lPWUSConfigR19LpwusMvalueAndSeqConfigFR1R19NOneConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "lpwus-OverlaidSeqRoot-r19"},
		{Name: "lpwus-OverlaidSeqNum-r19", Optional: true},
	},
}

var lPWUSConfigR19LpwusMvalueAndSeqConfigFR1R19NOneLpwusOverlaidSeqRootR19Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "root1-r19"},
		{Name: "root2-r19", Optional: true},
	},
}

const (
	LPWUS_Config_r19_Lpwus_MvalueAndSeqConfigFR1_r19_NOne_Lpwus_OverlaidSeqNum_r19_N1  = 0
	LPWUS_Config_r19_Lpwus_MvalueAndSeqConfigFR1_r19_NOne_Lpwus_OverlaidSeqNum_r19_N2  = 1
	LPWUS_Config_r19_Lpwus_MvalueAndSeqConfigFR1_r19_NOne_Lpwus_OverlaidSeqNum_r19_N4  = 2
	LPWUS_Config_r19_Lpwus_MvalueAndSeqConfigFR1_r19_NOne_Lpwus_OverlaidSeqNum_r19_N8  = 3
	LPWUS_Config_r19_Lpwus_MvalueAndSeqConfigFR1_r19_NOne_Lpwus_OverlaidSeqNum_r19_N16 = 4
)

var lPWUSConfigR19LpwusMvalueAndSeqConfigFR1R19NOneLpwusOverlaidSeqNumR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{LPWUS_Config_r19_Lpwus_MvalueAndSeqConfigFR1_r19_NOne_Lpwus_OverlaidSeqNum_r19_N1, LPWUS_Config_r19_Lpwus_MvalueAndSeqConfigFR1_r19_NOne_Lpwus_OverlaidSeqNum_r19_N2, LPWUS_Config_r19_Lpwus_MvalueAndSeqConfigFR1_r19_NOne_Lpwus_OverlaidSeqNum_r19_N4, LPWUS_Config_r19_Lpwus_MvalueAndSeqConfigFR1_r19_NOne_Lpwus_OverlaidSeqNum_r19_N8, LPWUS_Config_r19_Lpwus_MvalueAndSeqConfigFR1_r19_NOne_Lpwus_OverlaidSeqNum_r19_N16},
}

var lPWUSConfigR19LpwusMvalueAndSeqConfigFR1R19NTwoConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "lpwus-OverlaidSeqRoot-r19"},
		{Name: "lpwus-OverlaidSeqNum-r19", Optional: true},
	},
}

var lPWUSConfigR19LpwusMvalueAndSeqConfigFR1R19NTwoLpwusOverlaidSeqRootR19Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "root1-r19"},
		{Name: "root2-r19", Optional: true},
	},
}

const (
	LPWUS_Config_r19_Lpwus_MvalueAndSeqConfigFR1_r19_NTwo_Lpwus_OverlaidSeqNum_r19_N1 = 0
	LPWUS_Config_r19_Lpwus_MvalueAndSeqConfigFR1_r19_NTwo_Lpwus_OverlaidSeqNum_r19_N2 = 1
	LPWUS_Config_r19_Lpwus_MvalueAndSeqConfigFR1_r19_NTwo_Lpwus_OverlaidSeqNum_r19_N4 = 2
	LPWUS_Config_r19_Lpwus_MvalueAndSeqConfigFR1_r19_NTwo_Lpwus_OverlaidSeqNum_r19_N8 = 3
)

var lPWUSConfigR19LpwusMvalueAndSeqConfigFR1R19NTwoLpwusOverlaidSeqNumR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{LPWUS_Config_r19_Lpwus_MvalueAndSeqConfigFR1_r19_NTwo_Lpwus_OverlaidSeqNum_r19_N1, LPWUS_Config_r19_Lpwus_MvalueAndSeqConfigFR1_r19_NTwo_Lpwus_OverlaidSeqNum_r19_N2, LPWUS_Config_r19_Lpwus_MvalueAndSeqConfigFR1_r19_NTwo_Lpwus_OverlaidSeqNum_r19_N4, LPWUS_Config_r19_Lpwus_MvalueAndSeqConfigFR1_r19_NTwo_Lpwus_OverlaidSeqNum_r19_N8},
}

var lPWUSConfigR19LpwusMvalueAndSeqConfigFR1R19NFourConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "lpwus-OverlaidSeqRoot-r19"},
		{Name: "lpwus-OverlaidSeqNum-r19", Optional: true},
	},
}

var lPWUSConfigR19LpwusMvalueAndSeqConfigFR1R19NFourLpwusOverlaidSeqRootR19Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "root1-r19"},
		{Name: "root2-r19", Optional: true},
	},
}

const (
	LPWUS_Config_r19_Lpwus_MvalueAndSeqConfigFR1_r19_NFour_Lpwus_OverlaidSeqNum_r19_N1 = 0
	LPWUS_Config_r19_Lpwus_MvalueAndSeqConfigFR1_r19_NFour_Lpwus_OverlaidSeqNum_r19_N2 = 1
	LPWUS_Config_r19_Lpwus_MvalueAndSeqConfigFR1_r19_NFour_Lpwus_OverlaidSeqNum_r19_N4 = 2
)

var lPWUSConfigR19LpwusMvalueAndSeqConfigFR1R19NFourLpwusOverlaidSeqNumR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{LPWUS_Config_r19_Lpwus_MvalueAndSeqConfigFR1_r19_NFour_Lpwus_OverlaidSeqNum_r19_N1, LPWUS_Config_r19_Lpwus_MvalueAndSeqConfigFR1_r19_NFour_Lpwus_OverlaidSeqNum_r19_N2, LPWUS_Config_r19_Lpwus_MvalueAndSeqConfigFR1_r19_NFour_Lpwus_OverlaidSeqNum_r19_N4},
}

var lPWUSConfigR19LpwusMvalueAndSeqConfigFR2R19NOneConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "lpwus-OverlaidSeqRoot-r19"},
		{Name: "lpwus-OverlaidSeqNum-SCS-120kHz-r19", Optional: true},
		{Name: "lpwus-OverlaidSeqNum-SCS-60kHz-r19", Optional: true},
	},
}

var lPWUSConfigR19LpwusMvalueAndSeqConfigFR2R19NOneLpwusOverlaidSeqRootR19Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "root1-r19"},
		{Name: "root2-r19", Optional: true},
	},
}

const (
	LPWUS_Config_r19_Lpwus_MvalueAndSeqConfigFR2_r19_NOne_Lpwus_OverlaidSeqNum_SCS_120kHz_r19_N1 = 0
	LPWUS_Config_r19_Lpwus_MvalueAndSeqConfigFR2_r19_NOne_Lpwus_OverlaidSeqNum_SCS_120kHz_r19_N2 = 1
)

var lPWUSConfigR19LpwusMvalueAndSeqConfigFR2R19NOneLpwusOverlaidSeqNumSCS120kHzR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{LPWUS_Config_r19_Lpwus_MvalueAndSeqConfigFR2_r19_NOne_Lpwus_OverlaidSeqNum_SCS_120kHz_r19_N1, LPWUS_Config_r19_Lpwus_MvalueAndSeqConfigFR2_r19_NOne_Lpwus_OverlaidSeqNum_SCS_120kHz_r19_N2},
}

const (
	LPWUS_Config_r19_Lpwus_MvalueAndSeqConfigFR2_r19_NOne_Lpwus_OverlaidSeqNum_SCS_60kHz_r19_N1 = 0
	LPWUS_Config_r19_Lpwus_MvalueAndSeqConfigFR2_r19_NOne_Lpwus_OverlaidSeqNum_SCS_60kHz_r19_N2 = 1
	LPWUS_Config_r19_Lpwus_MvalueAndSeqConfigFR2_r19_NOne_Lpwus_OverlaidSeqNum_SCS_60kHz_r19_N4 = 2
)

var lPWUSConfigR19LpwusMvalueAndSeqConfigFR2R19NOneLpwusOverlaidSeqNumSCS60kHzR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{LPWUS_Config_r19_Lpwus_MvalueAndSeqConfigFR2_r19_NOne_Lpwus_OverlaidSeqNum_SCS_60kHz_r19_N1, LPWUS_Config_r19_Lpwus_MvalueAndSeqConfigFR2_r19_NOne_Lpwus_OverlaidSeqNum_SCS_60kHz_r19_N2, LPWUS_Config_r19_Lpwus_MvalueAndSeqConfigFR2_r19_NOne_Lpwus_OverlaidSeqNum_SCS_60kHz_r19_N4},
}

var lPWUSConfigR19LpwusMvalueAndSeqConfigFR2R19NTwoConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "lpwus-OverlaidSeqRoot-r19"},
		{Name: "lpwus-OverlaidSeqNum-SCS-60kHz-r19", Optional: true},
	},
}

var lPWUSConfigR19LpwusMvalueAndSeqConfigFR2R19NTwoLpwusOverlaidSeqRootR19Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "root1-r19"},
		{Name: "root2-r19", Optional: true},
	},
}

const (
	LPWUS_Config_r19_Lpwus_MvalueAndSeqConfigFR2_r19_NTwo_Lpwus_OverlaidSeqNum_SCS_60kHz_r19_N1 = 0
	LPWUS_Config_r19_Lpwus_MvalueAndSeqConfigFR2_r19_NTwo_Lpwus_OverlaidSeqNum_SCS_60kHz_r19_N2 = 1
)

var lPWUSConfigR19LpwusMvalueAndSeqConfigFR2R19NTwoLpwusOverlaidSeqNumSCS60kHzR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{LPWUS_Config_r19_Lpwus_MvalueAndSeqConfigFR2_r19_NTwo_Lpwus_OverlaidSeqNum_SCS_60kHz_r19_N1, LPWUS_Config_r19_Lpwus_MvalueAndSeqConfigFR2_r19_NTwo_Lpwus_OverlaidSeqNum_SCS_60kHz_r19_N2},
}

var lPWUSConfigR19LpwusPDCCHMonitoringTimerR19Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "timer1-r19"},
		{Name: "timer2-r19", Optional: true},
	},
}

var lPWUSConfigR19LpwusPDCCHMonitoringTimerR19Timer1R19Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "subMilliSeconds"},
		{Name: "milliSeconds"},
	},
}

const (
	LPWUS_Config_r19_Lpwus_PDCCH_MonitoringTimer_r19_Timer1_r19_SubMilliSeconds = 0
	LPWUS_Config_r19_Lpwus_PDCCH_MonitoringTimer_r19_Timer1_r19_MilliSeconds    = 1
)

const (
	LPWUS_Config_r19_Lpwus_PDCCH_MonitoringTimer_r19_Timer1_r19_MilliSeconds_Ms1    = 0
	LPWUS_Config_r19_Lpwus_PDCCH_MonitoringTimer_r19_Timer1_r19_MilliSeconds_Ms2    = 1
	LPWUS_Config_r19_Lpwus_PDCCH_MonitoringTimer_r19_Timer1_r19_MilliSeconds_Ms3    = 2
	LPWUS_Config_r19_Lpwus_PDCCH_MonitoringTimer_r19_Timer1_r19_MilliSeconds_Ms4    = 3
	LPWUS_Config_r19_Lpwus_PDCCH_MonitoringTimer_r19_Timer1_r19_MilliSeconds_Ms5    = 4
	LPWUS_Config_r19_Lpwus_PDCCH_MonitoringTimer_r19_Timer1_r19_MilliSeconds_Ms6    = 5
	LPWUS_Config_r19_Lpwus_PDCCH_MonitoringTimer_r19_Timer1_r19_MilliSeconds_Ms8    = 6
	LPWUS_Config_r19_Lpwus_PDCCH_MonitoringTimer_r19_Timer1_r19_MilliSeconds_Ms10   = 7
	LPWUS_Config_r19_Lpwus_PDCCH_MonitoringTimer_r19_Timer1_r19_MilliSeconds_Ms20   = 8
	LPWUS_Config_r19_Lpwus_PDCCH_MonitoringTimer_r19_Timer1_r19_MilliSeconds_Ms30   = 9
	LPWUS_Config_r19_Lpwus_PDCCH_MonitoringTimer_r19_Timer1_r19_MilliSeconds_Ms40   = 10
	LPWUS_Config_r19_Lpwus_PDCCH_MonitoringTimer_r19_Timer1_r19_MilliSeconds_Ms50   = 11
	LPWUS_Config_r19_Lpwus_PDCCH_MonitoringTimer_r19_Timer1_r19_MilliSeconds_Ms60   = 12
	LPWUS_Config_r19_Lpwus_PDCCH_MonitoringTimer_r19_Timer1_r19_MilliSeconds_Spare3 = 13
	LPWUS_Config_r19_Lpwus_PDCCH_MonitoringTimer_r19_Timer1_r19_MilliSeconds_Spare2 = 14
	LPWUS_Config_r19_Lpwus_PDCCH_MonitoringTimer_r19_Timer1_r19_MilliSeconds_Spare1 = 15
)

var lPWUSConfigR19LpwusPDCCHMonitoringTimerR19Timer1R19MilliSecondsConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{LPWUS_Config_r19_Lpwus_PDCCH_MonitoringTimer_r19_Timer1_r19_MilliSeconds_Ms1, LPWUS_Config_r19_Lpwus_PDCCH_MonitoringTimer_r19_Timer1_r19_MilliSeconds_Ms2, LPWUS_Config_r19_Lpwus_PDCCH_MonitoringTimer_r19_Timer1_r19_MilliSeconds_Ms3, LPWUS_Config_r19_Lpwus_PDCCH_MonitoringTimer_r19_Timer1_r19_MilliSeconds_Ms4, LPWUS_Config_r19_Lpwus_PDCCH_MonitoringTimer_r19_Timer1_r19_MilliSeconds_Ms5, LPWUS_Config_r19_Lpwus_PDCCH_MonitoringTimer_r19_Timer1_r19_MilliSeconds_Ms6, LPWUS_Config_r19_Lpwus_PDCCH_MonitoringTimer_r19_Timer1_r19_MilliSeconds_Ms8, LPWUS_Config_r19_Lpwus_PDCCH_MonitoringTimer_r19_Timer1_r19_MilliSeconds_Ms10, LPWUS_Config_r19_Lpwus_PDCCH_MonitoringTimer_r19_Timer1_r19_MilliSeconds_Ms20, LPWUS_Config_r19_Lpwus_PDCCH_MonitoringTimer_r19_Timer1_r19_MilliSeconds_Ms30, LPWUS_Config_r19_Lpwus_PDCCH_MonitoringTimer_r19_Timer1_r19_MilliSeconds_Ms40, LPWUS_Config_r19_Lpwus_PDCCH_MonitoringTimer_r19_Timer1_r19_MilliSeconds_Ms50, LPWUS_Config_r19_Lpwus_PDCCH_MonitoringTimer_r19_Timer1_r19_MilliSeconds_Ms60, LPWUS_Config_r19_Lpwus_PDCCH_MonitoringTimer_r19_Timer1_r19_MilliSeconds_Spare3, LPWUS_Config_r19_Lpwus_PDCCH_MonitoringTimer_r19_Timer1_r19_MilliSeconds_Spare2, LPWUS_Config_r19_Lpwus_PDCCH_MonitoringTimer_r19_Timer1_r19_MilliSeconds_Spare1},
}

var lPWUSConfigR19LpwusPDCCHMonitoringTimerR19Timer2R19Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "subMilliSeconds"},
		{Name: "milliSeconds"},
	},
}

const (
	LPWUS_Config_r19_Lpwus_PDCCH_MonitoringTimer_r19_Timer2_r19_SubMilliSeconds = 0
	LPWUS_Config_r19_Lpwus_PDCCH_MonitoringTimer_r19_Timer2_r19_MilliSeconds    = 1
)

const (
	LPWUS_Config_r19_Lpwus_PDCCH_MonitoringTimer_r19_Timer2_r19_MilliSeconds_Ms1    = 0
	LPWUS_Config_r19_Lpwus_PDCCH_MonitoringTimer_r19_Timer2_r19_MilliSeconds_Ms2    = 1
	LPWUS_Config_r19_Lpwus_PDCCH_MonitoringTimer_r19_Timer2_r19_MilliSeconds_Ms3    = 2
	LPWUS_Config_r19_Lpwus_PDCCH_MonitoringTimer_r19_Timer2_r19_MilliSeconds_Ms4    = 3
	LPWUS_Config_r19_Lpwus_PDCCH_MonitoringTimer_r19_Timer2_r19_MilliSeconds_Ms5    = 4
	LPWUS_Config_r19_Lpwus_PDCCH_MonitoringTimer_r19_Timer2_r19_MilliSeconds_Ms6    = 5
	LPWUS_Config_r19_Lpwus_PDCCH_MonitoringTimer_r19_Timer2_r19_MilliSeconds_Ms8    = 6
	LPWUS_Config_r19_Lpwus_PDCCH_MonitoringTimer_r19_Timer2_r19_MilliSeconds_Ms10   = 7
	LPWUS_Config_r19_Lpwus_PDCCH_MonitoringTimer_r19_Timer2_r19_MilliSeconds_Ms20   = 8
	LPWUS_Config_r19_Lpwus_PDCCH_MonitoringTimer_r19_Timer2_r19_MilliSeconds_Ms30   = 9
	LPWUS_Config_r19_Lpwus_PDCCH_MonitoringTimer_r19_Timer2_r19_MilliSeconds_Ms40   = 10
	LPWUS_Config_r19_Lpwus_PDCCH_MonitoringTimer_r19_Timer2_r19_MilliSeconds_Ms50   = 11
	LPWUS_Config_r19_Lpwus_PDCCH_MonitoringTimer_r19_Timer2_r19_MilliSeconds_Ms60   = 12
	LPWUS_Config_r19_Lpwus_PDCCH_MonitoringTimer_r19_Timer2_r19_MilliSeconds_Spare3 = 13
	LPWUS_Config_r19_Lpwus_PDCCH_MonitoringTimer_r19_Timer2_r19_MilliSeconds_Spare2 = 14
	LPWUS_Config_r19_Lpwus_PDCCH_MonitoringTimer_r19_Timer2_r19_MilliSeconds_Spare1 = 15
)

var lPWUSConfigR19LpwusPDCCHMonitoringTimerR19Timer2R19MilliSecondsConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{LPWUS_Config_r19_Lpwus_PDCCH_MonitoringTimer_r19_Timer2_r19_MilliSeconds_Ms1, LPWUS_Config_r19_Lpwus_PDCCH_MonitoringTimer_r19_Timer2_r19_MilliSeconds_Ms2, LPWUS_Config_r19_Lpwus_PDCCH_MonitoringTimer_r19_Timer2_r19_MilliSeconds_Ms3, LPWUS_Config_r19_Lpwus_PDCCH_MonitoringTimer_r19_Timer2_r19_MilliSeconds_Ms4, LPWUS_Config_r19_Lpwus_PDCCH_MonitoringTimer_r19_Timer2_r19_MilliSeconds_Ms5, LPWUS_Config_r19_Lpwus_PDCCH_MonitoringTimer_r19_Timer2_r19_MilliSeconds_Ms6, LPWUS_Config_r19_Lpwus_PDCCH_MonitoringTimer_r19_Timer2_r19_MilliSeconds_Ms8, LPWUS_Config_r19_Lpwus_PDCCH_MonitoringTimer_r19_Timer2_r19_MilliSeconds_Ms10, LPWUS_Config_r19_Lpwus_PDCCH_MonitoringTimer_r19_Timer2_r19_MilliSeconds_Ms20, LPWUS_Config_r19_Lpwus_PDCCH_MonitoringTimer_r19_Timer2_r19_MilliSeconds_Ms30, LPWUS_Config_r19_Lpwus_PDCCH_MonitoringTimer_r19_Timer2_r19_MilliSeconds_Ms40, LPWUS_Config_r19_Lpwus_PDCCH_MonitoringTimer_r19_Timer2_r19_MilliSeconds_Ms50, LPWUS_Config_r19_Lpwus_PDCCH_MonitoringTimer_r19_Timer2_r19_MilliSeconds_Ms60, LPWUS_Config_r19_Lpwus_PDCCH_MonitoringTimer_r19_Timer2_r19_MilliSeconds_Spare3, LPWUS_Config_r19_Lpwus_PDCCH_MonitoringTimer_r19_Timer2_r19_MilliSeconds_Spare2, LPWUS_Config_r19_Lpwus_PDCCH_MonitoringTimer_r19_Timer2_r19_MilliSeconds_Spare1},
}

type LPWUS_Config_r19 struct {
	Lpwus_MvalueAndSeqConfigFR1_r19 *struct {
		Choice int
		NOne   *struct {
			Lpwus_OverlaidSeqRoot_r19 struct {
				Root1_r19 int64
				Root2_r19 *int64
			}
			Lpwus_OverlaidSeqNum_r19 *int64
		}
		NTwo *struct {
			Lpwus_OverlaidSeqRoot_r19 struct {
				Root1_r19 int64
				Root2_r19 *int64
			}
			Lpwus_OverlaidSeqNum_r19 *int64
		}
		NFour *struct {
			Lpwus_OverlaidSeqRoot_r19 struct {
				Root1_r19 int64
				Root2_r19 *int64
			}
			Lpwus_OverlaidSeqNum_r19 *int64
		}
	}
	Lpwus_MvalueAndSeqConfigFR2_r19 *struct {
		Choice int
		NOne   *struct {
			Lpwus_OverlaidSeqRoot_r19 struct {
				Root1_r19 int64
				Root2_r19 *int64
			}
			Lpwus_OverlaidSeqNum_SCS_120kHz_r19 *int64
			Lpwus_OverlaidSeqNum_SCS_60kHz_r19  *int64
		}
		NTwo *struct {
			Lpwus_OverlaidSeqRoot_r19 struct {
				Root1_r19 int64
				Root2_r19 *int64
			}
			Lpwus_OverlaidSeqNum_SCS_60kHz_r19 *int64
		}
	}
	Lpwus_StartRB_r19       *int64
	Lpwus_TimeOffset1_1_r19 *int64
	Lpwus_TimeOffset1_2_r19 *int64
	Lpwus_MO_1_1_r19        *struct {
		Choice  int
		Sl1     *struct{}
		Sl2     *int64
		Sl4     *int64
		Sl5     *int64
		Sl8     *int64
		Sl10    *int64
		Sl16    *int64
		Sl20    *int64
		Sl40    *int64
		Sl80    *int64
		Sl160   *int64
		Sl320   *int64
		Sl640   *int64
		Sl1280  *int64
		Sl2560  *int64
		Sl5120  *int64
		Sl10240 *int64
	}
	Lpwus_MO_1_2_r19 *struct {
		Choice int
		Sl1    *struct{}
		Sl2    *int64
		Sl4    *int64
		Sl5    *int64
		Sl8    *int64
		Sl10   *int64
		Sl16   *int64
		Sl20   *int64
		Sl40   *int64
		Sl80   *int64
	}
	Lpwus_NumOfMO_1_1_r19       *int64
	Lpwus_NumOfMO_1_2_r19       *int64
	Lpwus_TCI_States_r19        []int64
	Lpwus_NominalMoDuration_r19 *int64
	Lpwus_ActualDuration_r19    *struct {
		Choice  int
		MValue1 *int64
		MValue2 *int64
		MValue4 *int64
	}
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
	Lpwus_TransmitOtherPeriodicCSI_r19 *int64
	Lpwus_TransmitPeriodicL1_RSRP_r19  *int64
	Lpwus_Codepoint_r19                []per.BitString
	Lpwus_NumOfBits_r19                *int64
	Lpwus_PDCCH_MonitoringTimer_r19    *struct {
		Timer1_r19 struct {
			Choice          int
			SubMilliSeconds *int64
			MilliSeconds    *int64
		}
		Timer2_r19 *struct {
			Choice          int
			SubMilliSeconds *int64
			MilliSeconds    *int64
		}
	}
}

func (ie *LPWUS_Config_r19) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(lPWUSConfigR19Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Lpwus_MvalueAndSeqConfigFR1_r19 != nil, ie.Lpwus_MvalueAndSeqConfigFR2_r19 != nil, ie.Lpwus_StartRB_r19 != nil, ie.Lpwus_TimeOffset1_1_r19 != nil, ie.Lpwus_TimeOffset1_2_r19 != nil, ie.Lpwus_MO_1_1_r19 != nil, ie.Lpwus_MO_1_2_r19 != nil, ie.Lpwus_NumOfMO_1_1_r19 != nil, ie.Lpwus_NumOfMO_1_2_r19 != nil, ie.Lpwus_TCI_States_r19 != nil, ie.Lpwus_NominalMoDuration_r19 != nil, ie.Lpwus_ActualDuration_r19 != nil, ie.Lpwus_AvailableSlot_r19 != nil, ie.Lpwus_AvailableSymbol_r19 != nil, ie.Lpwus_TransmitOtherPeriodicCSI_r19 != nil, ie.Lpwus_TransmitPeriodicL1_RSRP_r19 != nil, ie.Lpwus_NumOfBits_r19 != nil, ie.Lpwus_PDCCH_MonitoringTimer_r19 != nil}); err != nil {
		return err
	}

	// 2. lpwus-MvalueAndSeqConfigFR1-r19: choice
	{
		if ie.Lpwus_MvalueAndSeqConfigFR1_r19 != nil {
			choiceEnc := e.NewChoiceEncoder(lPWUS_Config_r19LpwusMvalueAndSeqConfigFR1R19Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.Lpwus_MvalueAndSeqConfigFR1_r19).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.Lpwus_MvalueAndSeqConfigFR1_r19).Choice {
			case LPWUS_Config_r19_Lpwus_MvalueAndSeqConfigFR1_r19_NOne:
				c := (*(*ie.Lpwus_MvalueAndSeqConfigFR1_r19).NOne)
				lPWUSConfigR19LpwusMvalueAndSeqConfigFR1R19NOneSeq := e.NewSequenceEncoder(lPWUSConfigR19LpwusMvalueAndSeqConfigFR1R19NOneConstraints)
				if err := lPWUSConfigR19LpwusMvalueAndSeqConfigFR1R19NOneSeq.EncodePreamble([]bool{c.Lpwus_OverlaidSeqNum_r19 != nil}); err != nil {
					return err
				}
				{
					c := &c.Lpwus_OverlaidSeqRoot_r19
					lPWUSConfigR19LpwusMvalueAndSeqConfigFR1R19NOneLpwusOverlaidSeqRootR19Seq := e.NewSequenceEncoder(lPWUSConfigR19LpwusMvalueAndSeqConfigFR1R19NOneLpwusOverlaidSeqRootR19Constraints)
					if err := lPWUSConfigR19LpwusMvalueAndSeqConfigFR1R19NOneLpwusOverlaidSeqRootR19Seq.EncodePreamble([]bool{c.Root2_r19 != nil}); err != nil {
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
					if err := e.EncodeEnumerated((*c.Lpwus_OverlaidSeqNum_r19), lPWUSConfigR19LpwusMvalueAndSeqConfigFR1R19NOneLpwusOverlaidSeqNumR19Constraints); err != nil {
						return err
					}
				}
			case LPWUS_Config_r19_Lpwus_MvalueAndSeqConfigFR1_r19_NTwo:
				c := (*(*ie.Lpwus_MvalueAndSeqConfigFR1_r19).NTwo)
				lPWUSConfigR19LpwusMvalueAndSeqConfigFR1R19NTwoSeq := e.NewSequenceEncoder(lPWUSConfigR19LpwusMvalueAndSeqConfigFR1R19NTwoConstraints)
				if err := lPWUSConfigR19LpwusMvalueAndSeqConfigFR1R19NTwoSeq.EncodePreamble([]bool{c.Lpwus_OverlaidSeqNum_r19 != nil}); err != nil {
					return err
				}
				{
					c := &c.Lpwus_OverlaidSeqRoot_r19
					lPWUSConfigR19LpwusMvalueAndSeqConfigFR1R19NTwoLpwusOverlaidSeqRootR19Seq := e.NewSequenceEncoder(lPWUSConfigR19LpwusMvalueAndSeqConfigFR1R19NTwoLpwusOverlaidSeqRootR19Constraints)
					if err := lPWUSConfigR19LpwusMvalueAndSeqConfigFR1R19NTwoLpwusOverlaidSeqRootR19Seq.EncodePreamble([]bool{c.Root2_r19 != nil}); err != nil {
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
					if err := e.EncodeEnumerated((*c.Lpwus_OverlaidSeqNum_r19), lPWUSConfigR19LpwusMvalueAndSeqConfigFR1R19NTwoLpwusOverlaidSeqNumR19Constraints); err != nil {
						return err
					}
				}
			case LPWUS_Config_r19_Lpwus_MvalueAndSeqConfigFR1_r19_NFour:
				c := (*(*ie.Lpwus_MvalueAndSeqConfigFR1_r19).NFour)
				lPWUSConfigR19LpwusMvalueAndSeqConfigFR1R19NFourSeq := e.NewSequenceEncoder(lPWUSConfigR19LpwusMvalueAndSeqConfigFR1R19NFourConstraints)
				if err := lPWUSConfigR19LpwusMvalueAndSeqConfigFR1R19NFourSeq.EncodePreamble([]bool{c.Lpwus_OverlaidSeqNum_r19 != nil}); err != nil {
					return err
				}
				{
					c := &c.Lpwus_OverlaidSeqRoot_r19
					lPWUSConfigR19LpwusMvalueAndSeqConfigFR1R19NFourLpwusOverlaidSeqRootR19Seq := e.NewSequenceEncoder(lPWUSConfigR19LpwusMvalueAndSeqConfigFR1R19NFourLpwusOverlaidSeqRootR19Constraints)
					if err := lPWUSConfigR19LpwusMvalueAndSeqConfigFR1R19NFourLpwusOverlaidSeqRootR19Seq.EncodePreamble([]bool{c.Root2_r19 != nil}); err != nil {
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
					if err := e.EncodeEnumerated((*c.Lpwus_OverlaidSeqNum_r19), lPWUSConfigR19LpwusMvalueAndSeqConfigFR1R19NFourLpwusOverlaidSeqNumR19Constraints); err != nil {
						return err
					}
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.Lpwus_MvalueAndSeqConfigFR1_r19).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 3. lpwus-MvalueAndSeqConfigFR2-r19: choice
	{
		if ie.Lpwus_MvalueAndSeqConfigFR2_r19 != nil {
			choiceEnc := e.NewChoiceEncoder(lPWUS_Config_r19LpwusMvalueAndSeqConfigFR2R19Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.Lpwus_MvalueAndSeqConfigFR2_r19).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.Lpwus_MvalueAndSeqConfigFR2_r19).Choice {
			case LPWUS_Config_r19_Lpwus_MvalueAndSeqConfigFR2_r19_NOne:
				c := (*(*ie.Lpwus_MvalueAndSeqConfigFR2_r19).NOne)
				lPWUSConfigR19LpwusMvalueAndSeqConfigFR2R19NOneSeq := e.NewSequenceEncoder(lPWUSConfigR19LpwusMvalueAndSeqConfigFR2R19NOneConstraints)
				if err := lPWUSConfigR19LpwusMvalueAndSeqConfigFR2R19NOneSeq.EncodePreamble([]bool{c.Lpwus_OverlaidSeqNum_SCS_120kHz_r19 != nil, c.Lpwus_OverlaidSeqNum_SCS_60kHz_r19 != nil}); err != nil {
					return err
				}
				{
					c := &c.Lpwus_OverlaidSeqRoot_r19
					lPWUSConfigR19LpwusMvalueAndSeqConfigFR2R19NOneLpwusOverlaidSeqRootR19Seq := e.NewSequenceEncoder(lPWUSConfigR19LpwusMvalueAndSeqConfigFR2R19NOneLpwusOverlaidSeqRootR19Constraints)
					if err := lPWUSConfigR19LpwusMvalueAndSeqConfigFR2R19NOneLpwusOverlaidSeqRootR19Seq.EncodePreamble([]bool{c.Root2_r19 != nil}); err != nil {
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
					if err := e.EncodeEnumerated((*c.Lpwus_OverlaidSeqNum_SCS_120kHz_r19), lPWUSConfigR19LpwusMvalueAndSeqConfigFR2R19NOneLpwusOverlaidSeqNumSCS120kHzR19Constraints); err != nil {
						return err
					}
				}
				if c.Lpwus_OverlaidSeqNum_SCS_60kHz_r19 != nil {
					if err := e.EncodeEnumerated((*c.Lpwus_OverlaidSeqNum_SCS_60kHz_r19), lPWUSConfigR19LpwusMvalueAndSeqConfigFR2R19NOneLpwusOverlaidSeqNumSCS60kHzR19Constraints); err != nil {
						return err
					}
				}
			case LPWUS_Config_r19_Lpwus_MvalueAndSeqConfigFR2_r19_NTwo:
				c := (*(*ie.Lpwus_MvalueAndSeqConfigFR2_r19).NTwo)
				lPWUSConfigR19LpwusMvalueAndSeqConfigFR2R19NTwoSeq := e.NewSequenceEncoder(lPWUSConfigR19LpwusMvalueAndSeqConfigFR2R19NTwoConstraints)
				if err := lPWUSConfigR19LpwusMvalueAndSeqConfigFR2R19NTwoSeq.EncodePreamble([]bool{c.Lpwus_OverlaidSeqNum_SCS_60kHz_r19 != nil}); err != nil {
					return err
				}
				{
					c := &c.Lpwus_OverlaidSeqRoot_r19
					lPWUSConfigR19LpwusMvalueAndSeqConfigFR2R19NTwoLpwusOverlaidSeqRootR19Seq := e.NewSequenceEncoder(lPWUSConfigR19LpwusMvalueAndSeqConfigFR2R19NTwoLpwusOverlaidSeqRootR19Constraints)
					if err := lPWUSConfigR19LpwusMvalueAndSeqConfigFR2R19NTwoLpwusOverlaidSeqRootR19Seq.EncodePreamble([]bool{c.Root2_r19 != nil}); err != nil {
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
				if c.Lpwus_OverlaidSeqNum_SCS_60kHz_r19 != nil {
					if err := e.EncodeEnumerated((*c.Lpwus_OverlaidSeqNum_SCS_60kHz_r19), lPWUSConfigR19LpwusMvalueAndSeqConfigFR2R19NTwoLpwusOverlaidSeqNumSCS60kHzR19Constraints); err != nil {
						return err
					}
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.Lpwus_MvalueAndSeqConfigFR2_r19).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 4. lpwus-StartRB-r19: integer
	{
		if ie.Lpwus_StartRB_r19 != nil {
			if err := e.EncodeInteger(*ie.Lpwus_StartRB_r19, lPWUSConfigR19LpwusStartRBR19Constraints); err != nil {
				return err
			}
		}
	}

	// 5. lpwus-TimeOffset1-1-r19: integer
	{
		if ie.Lpwus_TimeOffset1_1_r19 != nil {
			if err := e.EncodeInteger(*ie.Lpwus_TimeOffset1_1_r19, lPWUSConfigR19LpwusTimeOffset11R19Constraints); err != nil {
				return err
			}
		}
	}

	// 6. lpwus-TimeOffset1-2-r19: integer
	{
		if ie.Lpwus_TimeOffset1_2_r19 != nil {
			if err := e.EncodeInteger(*ie.Lpwus_TimeOffset1_2_r19, lPWUSConfigR19LpwusTimeOffset12R19Constraints); err != nil {
				return err
			}
		}
	}

	// 7. lpwus-MO-1-1-r19: choice
	{
		if ie.Lpwus_MO_1_1_r19 != nil {
			choiceEnc := e.NewChoiceEncoder(lPWUS_Config_r19LpwusMO11R19Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.Lpwus_MO_1_1_r19).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.Lpwus_MO_1_1_r19).Choice {
			case LPWUS_Config_r19_Lpwus_MO_1_1_r19_Sl1:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case LPWUS_Config_r19_Lpwus_MO_1_1_r19_Sl2:
				if err := e.EncodeInteger((*(*ie.Lpwus_MO_1_1_r19).Sl2), per.Constrained(0, 1)); err != nil {
					return err
				}
			case LPWUS_Config_r19_Lpwus_MO_1_1_r19_Sl4:
				if err := e.EncodeInteger((*(*ie.Lpwus_MO_1_1_r19).Sl4), per.Constrained(0, 3)); err != nil {
					return err
				}
			case LPWUS_Config_r19_Lpwus_MO_1_1_r19_Sl5:
				if err := e.EncodeInteger((*(*ie.Lpwus_MO_1_1_r19).Sl5), per.Constrained(0, 4)); err != nil {
					return err
				}
			case LPWUS_Config_r19_Lpwus_MO_1_1_r19_Sl8:
				if err := e.EncodeInteger((*(*ie.Lpwus_MO_1_1_r19).Sl8), per.Constrained(0, 7)); err != nil {
					return err
				}
			case LPWUS_Config_r19_Lpwus_MO_1_1_r19_Sl10:
				if err := e.EncodeInteger((*(*ie.Lpwus_MO_1_1_r19).Sl10), per.Constrained(0, 9)); err != nil {
					return err
				}
			case LPWUS_Config_r19_Lpwus_MO_1_1_r19_Sl16:
				if err := e.EncodeInteger((*(*ie.Lpwus_MO_1_1_r19).Sl16), per.Constrained(0, 15)); err != nil {
					return err
				}
			case LPWUS_Config_r19_Lpwus_MO_1_1_r19_Sl20:
				if err := e.EncodeInteger((*(*ie.Lpwus_MO_1_1_r19).Sl20), per.Constrained(0, 19)); err != nil {
					return err
				}
			case LPWUS_Config_r19_Lpwus_MO_1_1_r19_Sl40:
				if err := e.EncodeInteger((*(*ie.Lpwus_MO_1_1_r19).Sl40), per.Constrained(0, 39)); err != nil {
					return err
				}
			case LPWUS_Config_r19_Lpwus_MO_1_1_r19_Sl80:
				if err := e.EncodeInteger((*(*ie.Lpwus_MO_1_1_r19).Sl80), per.Constrained(0, 79)); err != nil {
					return err
				}
			case LPWUS_Config_r19_Lpwus_MO_1_1_r19_Sl160:
				if err := e.EncodeInteger((*(*ie.Lpwus_MO_1_1_r19).Sl160), per.Constrained(0, 159)); err != nil {
					return err
				}
			case LPWUS_Config_r19_Lpwus_MO_1_1_r19_Sl320:
				if err := e.EncodeInteger((*(*ie.Lpwus_MO_1_1_r19).Sl320), per.Constrained(0, 319)); err != nil {
					return err
				}
			case LPWUS_Config_r19_Lpwus_MO_1_1_r19_Sl640:
				if err := e.EncodeInteger((*(*ie.Lpwus_MO_1_1_r19).Sl640), per.Constrained(0, 639)); err != nil {
					return err
				}
			case LPWUS_Config_r19_Lpwus_MO_1_1_r19_Sl1280:
				if err := e.EncodeInteger((*(*ie.Lpwus_MO_1_1_r19).Sl1280), per.Constrained(0, 1279)); err != nil {
					return err
				}
			case LPWUS_Config_r19_Lpwus_MO_1_1_r19_Sl2560:
				if err := e.EncodeInteger((*(*ie.Lpwus_MO_1_1_r19).Sl2560), per.Constrained(0, 2559)); err != nil {
					return err
				}
			case LPWUS_Config_r19_Lpwus_MO_1_1_r19_Sl5120:
				if err := e.EncodeInteger((*(*ie.Lpwus_MO_1_1_r19).Sl5120), per.Constrained(0, 5119)); err != nil {
					return err
				}
			case LPWUS_Config_r19_Lpwus_MO_1_1_r19_Sl10240:
				if err := e.EncodeInteger((*(*ie.Lpwus_MO_1_1_r19).Sl10240), per.Constrained(0, 10239)); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.Lpwus_MO_1_1_r19).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 8. lpwus-MO-1-2-r19: choice
	{
		if ie.Lpwus_MO_1_2_r19 != nil {
			choiceEnc := e.NewChoiceEncoder(lPWUS_Config_r19LpwusMO12R19Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.Lpwus_MO_1_2_r19).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.Lpwus_MO_1_2_r19).Choice {
			case LPWUS_Config_r19_Lpwus_MO_1_2_r19_Sl1:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case LPWUS_Config_r19_Lpwus_MO_1_2_r19_Sl2:
				if err := e.EncodeInteger((*(*ie.Lpwus_MO_1_2_r19).Sl2), per.Constrained(0, 1)); err != nil {
					return err
				}
			case LPWUS_Config_r19_Lpwus_MO_1_2_r19_Sl4:
				if err := e.EncodeInteger((*(*ie.Lpwus_MO_1_2_r19).Sl4), per.Constrained(0, 3)); err != nil {
					return err
				}
			case LPWUS_Config_r19_Lpwus_MO_1_2_r19_Sl5:
				if err := e.EncodeInteger((*(*ie.Lpwus_MO_1_2_r19).Sl5), per.Constrained(0, 4)); err != nil {
					return err
				}
			case LPWUS_Config_r19_Lpwus_MO_1_2_r19_Sl8:
				if err := e.EncodeInteger((*(*ie.Lpwus_MO_1_2_r19).Sl8), per.Constrained(0, 7)); err != nil {
					return err
				}
			case LPWUS_Config_r19_Lpwus_MO_1_2_r19_Sl10:
				if err := e.EncodeInteger((*(*ie.Lpwus_MO_1_2_r19).Sl10), per.Constrained(0, 9)); err != nil {
					return err
				}
			case LPWUS_Config_r19_Lpwus_MO_1_2_r19_Sl16:
				if err := e.EncodeInteger((*(*ie.Lpwus_MO_1_2_r19).Sl16), per.Constrained(0, 15)); err != nil {
					return err
				}
			case LPWUS_Config_r19_Lpwus_MO_1_2_r19_Sl20:
				if err := e.EncodeInteger((*(*ie.Lpwus_MO_1_2_r19).Sl20), per.Constrained(0, 19)); err != nil {
					return err
				}
			case LPWUS_Config_r19_Lpwus_MO_1_2_r19_Sl40:
				if err := e.EncodeInteger((*(*ie.Lpwus_MO_1_2_r19).Sl40), per.Constrained(0, 39)); err != nil {
					return err
				}
			case LPWUS_Config_r19_Lpwus_MO_1_2_r19_Sl80:
				if err := e.EncodeInteger((*(*ie.Lpwus_MO_1_2_r19).Sl80), per.Constrained(0, 79)); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.Lpwus_MO_1_2_r19).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 9. lpwus-NumOfMO-1-1-r19: integer
	{
		if ie.Lpwus_NumOfMO_1_1_r19 != nil {
			if err := e.EncodeInteger(*ie.Lpwus_NumOfMO_1_1_r19, lPWUSConfigR19LpwusNumOfMO11R19Constraints); err != nil {
				return err
			}
		}
	}

	// 10. lpwus-NumOfMO-1-2-r19: integer
	{
		if ie.Lpwus_NumOfMO_1_2_r19 != nil {
			if err := e.EncodeInteger(*ie.Lpwus_NumOfMO_1_2_r19, lPWUSConfigR19LpwusNumOfMO12R19Constraints); err != nil {
				return err
			}
		}
	}

	// 11. lpwus-TCI-States-r19: sequence-of
	{
		if ie.Lpwus_TCI_States_r19 != nil {
			seqOf := e.NewSequenceOfEncoder(lPWUSConfigR19LpwusTCIStatesR19Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Lpwus_TCI_States_r19))); err != nil {
				return err
			}
			for i := range ie.Lpwus_TCI_States_r19 {
				if err := e.EncodeInteger(ie.Lpwus_TCI_States_r19[i], per.Constrained(0, 15)); err != nil {
					return err
				}
			}
		}
	}

	// 12. lpwus-NominalMoDuration-r19: integer
	{
		if ie.Lpwus_NominalMoDuration_r19 != nil {
			if err := e.EncodeInteger(*ie.Lpwus_NominalMoDuration_r19, lPWUSConfigR19LpwusNominalMoDurationR19Constraints); err != nil {
				return err
			}
		}
	}

	// 13. lpwus-ActualDuration-r19: choice
	{
		if ie.Lpwus_ActualDuration_r19 != nil {
			choiceEnc := e.NewChoiceEncoder(lPWUS_Config_r19LpwusActualDurationR19Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.Lpwus_ActualDuration_r19).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.Lpwus_ActualDuration_r19).Choice {
			case LPWUS_Config_r19_Lpwus_ActualDuration_r19_MValue1:
				if err := e.EncodeInteger((*(*ie.Lpwus_ActualDuration_r19).MValue1), per.Constrained(2, 64)); err != nil {
					return err
				}
			case LPWUS_Config_r19_Lpwus_ActualDuration_r19_MValue2:
				if err := e.EncodeInteger((*(*ie.Lpwus_ActualDuration_r19).MValue2), per.Constrained(1, 32)); err != nil {
					return err
				}
			case LPWUS_Config_r19_Lpwus_ActualDuration_r19_MValue4:
				if err := e.EncodeInteger((*(*ie.Lpwus_ActualDuration_r19).MValue4), per.Constrained(1, 16)); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.Lpwus_ActualDuration_r19).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 14. lpwus-AvailableSlot-r19: choice
	{
		if ie.Lpwus_AvailableSlot_r19 != nil {
			choiceEnc := e.NewChoiceEncoder(lPWUS_Config_r19LpwusAvailableSlotR19Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.Lpwus_AvailableSlot_r19).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.Lpwus_AvailableSlot_r19).Choice {
			case LPWUS_Config_r19_Lpwus_AvailableSlot_r19_N10:
				if err := e.EncodeBitString((*(*ie.Lpwus_AvailableSlot_r19).N10), per.FixedSize(10)); err != nil {
					return err
				}
			case LPWUS_Config_r19_Lpwus_AvailableSlot_r19_N20:
				if err := e.EncodeBitString((*(*ie.Lpwus_AvailableSlot_r19).N20), per.FixedSize(20)); err != nil {
					return err
				}
			case LPWUS_Config_r19_Lpwus_AvailableSlot_r19_N40:
				if err := e.EncodeBitString((*(*ie.Lpwus_AvailableSlot_r19).N40), per.FixedSize(40)); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.Lpwus_AvailableSlot_r19).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 15. lpwus-AvailableSymbol-r19: choice
	{
		if ie.Lpwus_AvailableSymbol_r19 != nil {
			choiceEnc := e.NewChoiceEncoder(lPWUS_Config_r19LpwusAvailableSymbolR19Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.Lpwus_AvailableSymbol_r19).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.Lpwus_AvailableSymbol_r19).Choice {
			case LPWUS_Config_r19_Lpwus_AvailableSymbol_r19_OneSlot:
				if err := e.EncodeBitString((*(*ie.Lpwus_AvailableSymbol_r19).OneSlot), per.FixedSize(14)); err != nil {
					return err
				}
			case LPWUS_Config_r19_Lpwus_AvailableSymbol_r19_TwoSlots:
				if err := e.EncodeBitString((*(*ie.Lpwus_AvailableSymbol_r19).TwoSlots), per.FixedSize(28)); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.Lpwus_AvailableSymbol_r19).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 16. lpwus-TransmitOtherPeriodicCSI-r19: enumerated
	{
		if ie.Lpwus_TransmitOtherPeriodicCSI_r19 != nil {
			if err := e.EncodeEnumerated(*ie.Lpwus_TransmitOtherPeriodicCSI_r19, lPWUSConfigR19LpwusTransmitOtherPeriodicCSIR19Constraints); err != nil {
				return err
			}
		}
	}

	// 17. lpwus-TransmitPeriodicL1-RSRP-r19: enumerated
	{
		if ie.Lpwus_TransmitPeriodicL1_RSRP_r19 != nil {
			if err := e.EncodeEnumerated(*ie.Lpwus_TransmitPeriodicL1_RSRP_r19, lPWUSConfigR19LpwusTransmitPeriodicL1RSRPR19Constraints); err != nil {
				return err
			}
		}
	}

	// 18. lpwus-Codepoint-r19: sequence-of
	{
		seqOf := e.NewSequenceOfEncoder(lPWUSConfigR19LpwusCodepointR19Constraints)
		if err := seqOf.EncodeLength(int64(len(ie.Lpwus_Codepoint_r19))); err != nil {
			return err
		}
		for i := range ie.Lpwus_Codepoint_r19 {
			if err := e.EncodeBitString(ie.Lpwus_Codepoint_r19[i], per.SizeRange(1, 5)); err != nil {
				return err
			}
		}
	}

	// 19. lpwus-NumOfBits-r19: integer
	{
		if ie.Lpwus_NumOfBits_r19 != nil {
			if err := e.EncodeInteger(*ie.Lpwus_NumOfBits_r19, lPWUSConfigR19LpwusNumOfBitsR19Constraints); err != nil {
				return err
			}
		}
	}

	// 20. lpwus-PDCCH-MonitoringTimer-r19: sequence
	{
		if ie.Lpwus_PDCCH_MonitoringTimer_r19 != nil {
			c := ie.Lpwus_PDCCH_MonitoringTimer_r19
			lPWUSConfigR19LpwusPDCCHMonitoringTimerR19Seq := e.NewSequenceEncoder(lPWUSConfigR19LpwusPDCCHMonitoringTimerR19Constraints)
			if err := lPWUSConfigR19LpwusPDCCHMonitoringTimerR19Seq.EncodePreamble([]bool{c.Timer2_r19 != nil}); err != nil {
				return err
			}
			{
				choiceEnc := e.NewChoiceEncoder(lPWUSConfigR19LpwusPDCCHMonitoringTimerR19Timer1R19Constraints)
				if err := choiceEnc.EncodeChoice(int64(c.Timer1_r19.Choice), false, nil); err != nil {
					return err
				}
				switch c.Timer1_r19.Choice {
				case LPWUS_Config_r19_Lpwus_PDCCH_MonitoringTimer_r19_Timer1_r19_SubMilliSeconds:
					if err := e.EncodeInteger((*c.Timer1_r19.SubMilliSeconds), per.Constrained(1, 31)); err != nil {
						return err
					}
				case LPWUS_Config_r19_Lpwus_PDCCH_MonitoringTimer_r19_Timer1_r19_MilliSeconds:
					if err := e.EncodeEnumerated((*c.Timer1_r19.MilliSeconds), lPWUSConfigR19LpwusPDCCHMonitoringTimerR19Timer1R19MilliSecondsConstraints); err != nil {
						return err
					}
				}
			}
			if c.Timer2_r19 != nil {
				choiceEnc := e.NewChoiceEncoder(lPWUSConfigR19LpwusPDCCHMonitoringTimerR19Timer2R19Constraints)
				if err := choiceEnc.EncodeChoice(int64((*c.Timer2_r19).Choice), false, nil); err != nil {
					return err
				}
				switch (*c.Timer2_r19).Choice {
				case LPWUS_Config_r19_Lpwus_PDCCH_MonitoringTimer_r19_Timer2_r19_SubMilliSeconds:
					if err := e.EncodeInteger((*(*c.Timer2_r19).SubMilliSeconds), per.Constrained(1, 31)); err != nil {
						return err
					}
				case LPWUS_Config_r19_Lpwus_PDCCH_MonitoringTimer_r19_Timer2_r19_MilliSeconds:
					if err := e.EncodeEnumerated((*(*c.Timer2_r19).MilliSeconds), lPWUSConfigR19LpwusPDCCHMonitoringTimerR19Timer2R19MilliSecondsConstraints); err != nil {
						return err
					}
				}
			}
		}
	}

	return nil
}

func (ie *LPWUS_Config_r19) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(lPWUSConfigR19Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. lpwus-MvalueAndSeqConfigFR1-r19: choice
	{
		if seq.IsComponentPresent(0) {
			ie.Lpwus_MvalueAndSeqConfigFR1_r19 = &struct {
				Choice int
				NOne   *struct {
					Lpwus_OverlaidSeqRoot_r19 struct {
						Root1_r19 int64
						Root2_r19 *int64
					}
					Lpwus_OverlaidSeqNum_r19 *int64
				}
				NTwo *struct {
					Lpwus_OverlaidSeqRoot_r19 struct {
						Root1_r19 int64
						Root2_r19 *int64
					}
					Lpwus_OverlaidSeqNum_r19 *int64
				}
				NFour *struct {
					Lpwus_OverlaidSeqRoot_r19 struct {
						Root1_r19 int64
						Root2_r19 *int64
					}
					Lpwus_OverlaidSeqNum_r19 *int64
				}
			}{}
			choiceDec := d.NewChoiceDecoder(lPWUS_Config_r19LpwusMvalueAndSeqConfigFR1R19Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Lpwus_MvalueAndSeqConfigFR1_r19).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case LPWUS_Config_r19_Lpwus_MvalueAndSeqConfigFR1_r19_NOne:
				(*ie.Lpwus_MvalueAndSeqConfigFR1_r19).NOne = &struct {
					Lpwus_OverlaidSeqRoot_r19 struct {
						Root1_r19 int64
						Root2_r19 *int64
					}
					Lpwus_OverlaidSeqNum_r19 *int64
				}{}
				c := (*(*ie.Lpwus_MvalueAndSeqConfigFR1_r19).NOne)
				lPWUSConfigR19LpwusMvalueAndSeqConfigFR1R19NOneSeq := d.NewSequenceDecoder(lPWUSConfigR19LpwusMvalueAndSeqConfigFR1R19NOneConstraints)
				if err := lPWUSConfigR19LpwusMvalueAndSeqConfigFR1R19NOneSeq.DecodePreamble(); err != nil {
					return err
				}
				{
					c := &c.Lpwus_OverlaidSeqRoot_r19
					lPWUSConfigR19LpwusMvalueAndSeqConfigFR1R19NOneLpwusOverlaidSeqRootR19Seq := d.NewSequenceDecoder(lPWUSConfigR19LpwusMvalueAndSeqConfigFR1R19NOneLpwusOverlaidSeqRootR19Constraints)
					if err := lPWUSConfigR19LpwusMvalueAndSeqConfigFR1R19NOneLpwusOverlaidSeqRootR19Seq.DecodePreamble(); err != nil {
						return err
					}
					{
						v, err := d.DecodeInteger(per.Constrained(1, 131))
						if err != nil {
							return err
						}
						c.Root1_r19 = v
					}
					if lPWUSConfigR19LpwusMvalueAndSeqConfigFR1R19NOneLpwusOverlaidSeqRootR19Seq.IsComponentPresent(1) {
						c.Root2_r19 = new(int64)
						v, err := d.DecodeInteger(per.Constrained(1, 131))
						if err != nil {
							return err
						}
						(*c.Root2_r19) = v
					}
				}
				if lPWUSConfigR19LpwusMvalueAndSeqConfigFR1R19NOneSeq.IsComponentPresent(1) {
					c.Lpwus_OverlaidSeqNum_r19 = new(int64)
					v, err := d.DecodeEnumerated(lPWUSConfigR19LpwusMvalueAndSeqConfigFR1R19NOneLpwusOverlaidSeqNumR19Constraints)
					if err != nil {
						return err
					}
					(*c.Lpwus_OverlaidSeqNum_r19) = v
				}
			case LPWUS_Config_r19_Lpwus_MvalueAndSeqConfigFR1_r19_NTwo:
				(*ie.Lpwus_MvalueAndSeqConfigFR1_r19).NTwo = &struct {
					Lpwus_OverlaidSeqRoot_r19 struct {
						Root1_r19 int64
						Root2_r19 *int64
					}
					Lpwus_OverlaidSeqNum_r19 *int64
				}{}
				c := (*(*ie.Lpwus_MvalueAndSeqConfigFR1_r19).NTwo)
				lPWUSConfigR19LpwusMvalueAndSeqConfigFR1R19NTwoSeq := d.NewSequenceDecoder(lPWUSConfigR19LpwusMvalueAndSeqConfigFR1R19NTwoConstraints)
				if err := lPWUSConfigR19LpwusMvalueAndSeqConfigFR1R19NTwoSeq.DecodePreamble(); err != nil {
					return err
				}
				{
					c := &c.Lpwus_OverlaidSeqRoot_r19
					lPWUSConfigR19LpwusMvalueAndSeqConfigFR1R19NTwoLpwusOverlaidSeqRootR19Seq := d.NewSequenceDecoder(lPWUSConfigR19LpwusMvalueAndSeqConfigFR1R19NTwoLpwusOverlaidSeqRootR19Constraints)
					if err := lPWUSConfigR19LpwusMvalueAndSeqConfigFR1R19NTwoLpwusOverlaidSeqRootR19Seq.DecodePreamble(); err != nil {
						return err
					}
					{
						v, err := d.DecodeInteger(per.Constrained(1, 61))
						if err != nil {
							return err
						}
						c.Root1_r19 = v
					}
					if lPWUSConfigR19LpwusMvalueAndSeqConfigFR1R19NTwoLpwusOverlaidSeqRootR19Seq.IsComponentPresent(1) {
						c.Root2_r19 = new(int64)
						v, err := d.DecodeInteger(per.Constrained(1, 61))
						if err != nil {
							return err
						}
						(*c.Root2_r19) = v
					}
				}
				if lPWUSConfigR19LpwusMvalueAndSeqConfigFR1R19NTwoSeq.IsComponentPresent(1) {
					c.Lpwus_OverlaidSeqNum_r19 = new(int64)
					v, err := d.DecodeEnumerated(lPWUSConfigR19LpwusMvalueAndSeqConfigFR1R19NTwoLpwusOverlaidSeqNumR19Constraints)
					if err != nil {
						return err
					}
					(*c.Lpwus_OverlaidSeqNum_r19) = v
				}
			case LPWUS_Config_r19_Lpwus_MvalueAndSeqConfigFR1_r19_NFour:
				(*ie.Lpwus_MvalueAndSeqConfigFR1_r19).NFour = &struct {
					Lpwus_OverlaidSeqRoot_r19 struct {
						Root1_r19 int64
						Root2_r19 *int64
					}
					Lpwus_OverlaidSeqNum_r19 *int64
				}{}
				c := (*(*ie.Lpwus_MvalueAndSeqConfigFR1_r19).NFour)
				lPWUSConfigR19LpwusMvalueAndSeqConfigFR1R19NFourSeq := d.NewSequenceDecoder(lPWUSConfigR19LpwusMvalueAndSeqConfigFR1R19NFourConstraints)
				if err := lPWUSConfigR19LpwusMvalueAndSeqConfigFR1R19NFourSeq.DecodePreamble(); err != nil {
					return err
				}
				{
					c := &c.Lpwus_OverlaidSeqRoot_r19
					lPWUSConfigR19LpwusMvalueAndSeqConfigFR1R19NFourLpwusOverlaidSeqRootR19Seq := d.NewSequenceDecoder(lPWUSConfigR19LpwusMvalueAndSeqConfigFR1R19NFourLpwusOverlaidSeqRootR19Constraints)
					if err := lPWUSConfigR19LpwusMvalueAndSeqConfigFR1R19NFourLpwusOverlaidSeqRootR19Seq.DecodePreamble(); err != nil {
						return err
					}
					{
						v, err := d.DecodeInteger(per.Constrained(1, 31))
						if err != nil {
							return err
						}
						c.Root1_r19 = v
					}
					if lPWUSConfigR19LpwusMvalueAndSeqConfigFR1R19NFourLpwusOverlaidSeqRootR19Seq.IsComponentPresent(1) {
						c.Root2_r19 = new(int64)
						v, err := d.DecodeInteger(per.Constrained(1, 31))
						if err != nil {
							return err
						}
						(*c.Root2_r19) = v
					}
				}
				if lPWUSConfigR19LpwusMvalueAndSeqConfigFR1R19NFourSeq.IsComponentPresent(1) {
					c.Lpwus_OverlaidSeqNum_r19 = new(int64)
					v, err := d.DecodeEnumerated(lPWUSConfigR19LpwusMvalueAndSeqConfigFR1R19NFourLpwusOverlaidSeqNumR19Constraints)
					if err != nil {
						return err
					}
					(*c.Lpwus_OverlaidSeqNum_r19) = v
				}
			}
		}
	}

	// 3. lpwus-MvalueAndSeqConfigFR2-r19: choice
	{
		if seq.IsComponentPresent(1) {
			ie.Lpwus_MvalueAndSeqConfigFR2_r19 = &struct {
				Choice int
				NOne   *struct {
					Lpwus_OverlaidSeqRoot_r19 struct {
						Root1_r19 int64
						Root2_r19 *int64
					}
					Lpwus_OverlaidSeqNum_SCS_120kHz_r19 *int64
					Lpwus_OverlaidSeqNum_SCS_60kHz_r19  *int64
				}
				NTwo *struct {
					Lpwus_OverlaidSeqRoot_r19 struct {
						Root1_r19 int64
						Root2_r19 *int64
					}
					Lpwus_OverlaidSeqNum_SCS_60kHz_r19 *int64
				}
			}{}
			choiceDec := d.NewChoiceDecoder(lPWUS_Config_r19LpwusMvalueAndSeqConfigFR2R19Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Lpwus_MvalueAndSeqConfigFR2_r19).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case LPWUS_Config_r19_Lpwus_MvalueAndSeqConfigFR2_r19_NOne:
				(*ie.Lpwus_MvalueAndSeqConfigFR2_r19).NOne = &struct {
					Lpwus_OverlaidSeqRoot_r19 struct {
						Root1_r19 int64
						Root2_r19 *int64
					}
					Lpwus_OverlaidSeqNum_SCS_120kHz_r19 *int64
					Lpwus_OverlaidSeqNum_SCS_60kHz_r19  *int64
				}{}
				c := (*(*ie.Lpwus_MvalueAndSeqConfigFR2_r19).NOne)
				lPWUSConfigR19LpwusMvalueAndSeqConfigFR2R19NOneSeq := d.NewSequenceDecoder(lPWUSConfigR19LpwusMvalueAndSeqConfigFR2R19NOneConstraints)
				if err := lPWUSConfigR19LpwusMvalueAndSeqConfigFR2R19NOneSeq.DecodePreamble(); err != nil {
					return err
				}
				{
					c := &c.Lpwus_OverlaidSeqRoot_r19
					lPWUSConfigR19LpwusMvalueAndSeqConfigFR2R19NOneLpwusOverlaidSeqRootR19Seq := d.NewSequenceDecoder(lPWUSConfigR19LpwusMvalueAndSeqConfigFR2R19NOneLpwusOverlaidSeqRootR19Constraints)
					if err := lPWUSConfigR19LpwusMvalueAndSeqConfigFR2R19NOneLpwusOverlaidSeqRootR19Seq.DecodePreamble(); err != nil {
						return err
					}
					{
						v, err := d.DecodeInteger(per.Constrained(1, 131))
						if err != nil {
							return err
						}
						c.Root1_r19 = v
					}
					if lPWUSConfigR19LpwusMvalueAndSeqConfigFR2R19NOneLpwusOverlaidSeqRootR19Seq.IsComponentPresent(1) {
						c.Root2_r19 = new(int64)
						v, err := d.DecodeInteger(per.Constrained(1, 131))
						if err != nil {
							return err
						}
						(*c.Root2_r19) = v
					}
				}
				if lPWUSConfigR19LpwusMvalueAndSeqConfigFR2R19NOneSeq.IsComponentPresent(1) {
					c.Lpwus_OverlaidSeqNum_SCS_120kHz_r19 = new(int64)
					v, err := d.DecodeEnumerated(lPWUSConfigR19LpwusMvalueAndSeqConfigFR2R19NOneLpwusOverlaidSeqNumSCS120kHzR19Constraints)
					if err != nil {
						return err
					}
					(*c.Lpwus_OverlaidSeqNum_SCS_120kHz_r19) = v
				}
				if lPWUSConfigR19LpwusMvalueAndSeqConfigFR2R19NOneSeq.IsComponentPresent(2) {
					c.Lpwus_OverlaidSeqNum_SCS_60kHz_r19 = new(int64)
					v, err := d.DecodeEnumerated(lPWUSConfigR19LpwusMvalueAndSeqConfigFR2R19NOneLpwusOverlaidSeqNumSCS60kHzR19Constraints)
					if err != nil {
						return err
					}
					(*c.Lpwus_OverlaidSeqNum_SCS_60kHz_r19) = v
				}
			case LPWUS_Config_r19_Lpwus_MvalueAndSeqConfigFR2_r19_NTwo:
				(*ie.Lpwus_MvalueAndSeqConfigFR2_r19).NTwo = &struct {
					Lpwus_OverlaidSeqRoot_r19 struct {
						Root1_r19 int64
						Root2_r19 *int64
					}
					Lpwus_OverlaidSeqNum_SCS_60kHz_r19 *int64
				}{}
				c := (*(*ie.Lpwus_MvalueAndSeqConfigFR2_r19).NTwo)
				lPWUSConfigR19LpwusMvalueAndSeqConfigFR2R19NTwoSeq := d.NewSequenceDecoder(lPWUSConfigR19LpwusMvalueAndSeqConfigFR2R19NTwoConstraints)
				if err := lPWUSConfigR19LpwusMvalueAndSeqConfigFR2R19NTwoSeq.DecodePreamble(); err != nil {
					return err
				}
				{
					c := &c.Lpwus_OverlaidSeqRoot_r19
					lPWUSConfigR19LpwusMvalueAndSeqConfigFR2R19NTwoLpwusOverlaidSeqRootR19Seq := d.NewSequenceDecoder(lPWUSConfigR19LpwusMvalueAndSeqConfigFR2R19NTwoLpwusOverlaidSeqRootR19Constraints)
					if err := lPWUSConfigR19LpwusMvalueAndSeqConfigFR2R19NTwoLpwusOverlaidSeqRootR19Seq.DecodePreamble(); err != nil {
						return err
					}
					{
						v, err := d.DecodeInteger(per.Constrained(1, 61))
						if err != nil {
							return err
						}
						c.Root1_r19 = v
					}
					if lPWUSConfigR19LpwusMvalueAndSeqConfigFR2R19NTwoLpwusOverlaidSeqRootR19Seq.IsComponentPresent(1) {
						c.Root2_r19 = new(int64)
						v, err := d.DecodeInteger(per.Constrained(1, 61))
						if err != nil {
							return err
						}
						(*c.Root2_r19) = v
					}
				}
				if lPWUSConfigR19LpwusMvalueAndSeqConfigFR2R19NTwoSeq.IsComponentPresent(1) {
					c.Lpwus_OverlaidSeqNum_SCS_60kHz_r19 = new(int64)
					v, err := d.DecodeEnumerated(lPWUSConfigR19LpwusMvalueAndSeqConfigFR2R19NTwoLpwusOverlaidSeqNumSCS60kHzR19Constraints)
					if err != nil {
						return err
					}
					(*c.Lpwus_OverlaidSeqNum_SCS_60kHz_r19) = v
				}
			}
		}
	}

	// 4. lpwus-StartRB-r19: integer
	{
		if seq.IsComponentPresent(2) {
			v, err := d.DecodeInteger(lPWUSConfigR19LpwusStartRBR19Constraints)
			if err != nil {
				return err
			}
			ie.Lpwus_StartRB_r19 = &v
		}
	}

	// 5. lpwus-TimeOffset1-1-r19: integer
	{
		if seq.IsComponentPresent(3) {
			v, err := d.DecodeInteger(lPWUSConfigR19LpwusTimeOffset11R19Constraints)
			if err != nil {
				return err
			}
			ie.Lpwus_TimeOffset1_1_r19 = &v
		}
	}

	// 6. lpwus-TimeOffset1-2-r19: integer
	{
		if seq.IsComponentPresent(4) {
			v, err := d.DecodeInteger(lPWUSConfigR19LpwusTimeOffset12R19Constraints)
			if err != nil {
				return err
			}
			ie.Lpwus_TimeOffset1_2_r19 = &v
		}
	}

	// 7. lpwus-MO-1-1-r19: choice
	{
		if seq.IsComponentPresent(5) {
			ie.Lpwus_MO_1_1_r19 = &struct {
				Choice  int
				Sl1     *struct{}
				Sl2     *int64
				Sl4     *int64
				Sl5     *int64
				Sl8     *int64
				Sl10    *int64
				Sl16    *int64
				Sl20    *int64
				Sl40    *int64
				Sl80    *int64
				Sl160   *int64
				Sl320   *int64
				Sl640   *int64
				Sl1280  *int64
				Sl2560  *int64
				Sl5120  *int64
				Sl10240 *int64
			}{}
			choiceDec := d.NewChoiceDecoder(lPWUS_Config_r19LpwusMO11R19Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Lpwus_MO_1_1_r19).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case LPWUS_Config_r19_Lpwus_MO_1_1_r19_Sl1:
				(*ie.Lpwus_MO_1_1_r19).Sl1 = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case LPWUS_Config_r19_Lpwus_MO_1_1_r19_Sl2:
				(*ie.Lpwus_MO_1_1_r19).Sl2 = new(int64)
				v, err := d.DecodeInteger(per.Constrained(0, 1))
				if err != nil {
					return err
				}
				(*(*ie.Lpwus_MO_1_1_r19).Sl2) = v
			case LPWUS_Config_r19_Lpwus_MO_1_1_r19_Sl4:
				(*ie.Lpwus_MO_1_1_r19).Sl4 = new(int64)
				v, err := d.DecodeInteger(per.Constrained(0, 3))
				if err != nil {
					return err
				}
				(*(*ie.Lpwus_MO_1_1_r19).Sl4) = v
			case LPWUS_Config_r19_Lpwus_MO_1_1_r19_Sl5:
				(*ie.Lpwus_MO_1_1_r19).Sl5 = new(int64)
				v, err := d.DecodeInteger(per.Constrained(0, 4))
				if err != nil {
					return err
				}
				(*(*ie.Lpwus_MO_1_1_r19).Sl5) = v
			case LPWUS_Config_r19_Lpwus_MO_1_1_r19_Sl8:
				(*ie.Lpwus_MO_1_1_r19).Sl8 = new(int64)
				v, err := d.DecodeInteger(per.Constrained(0, 7))
				if err != nil {
					return err
				}
				(*(*ie.Lpwus_MO_1_1_r19).Sl8) = v
			case LPWUS_Config_r19_Lpwus_MO_1_1_r19_Sl10:
				(*ie.Lpwus_MO_1_1_r19).Sl10 = new(int64)
				v, err := d.DecodeInteger(per.Constrained(0, 9))
				if err != nil {
					return err
				}
				(*(*ie.Lpwus_MO_1_1_r19).Sl10) = v
			case LPWUS_Config_r19_Lpwus_MO_1_1_r19_Sl16:
				(*ie.Lpwus_MO_1_1_r19).Sl16 = new(int64)
				v, err := d.DecodeInteger(per.Constrained(0, 15))
				if err != nil {
					return err
				}
				(*(*ie.Lpwus_MO_1_1_r19).Sl16) = v
			case LPWUS_Config_r19_Lpwus_MO_1_1_r19_Sl20:
				(*ie.Lpwus_MO_1_1_r19).Sl20 = new(int64)
				v, err := d.DecodeInteger(per.Constrained(0, 19))
				if err != nil {
					return err
				}
				(*(*ie.Lpwus_MO_1_1_r19).Sl20) = v
			case LPWUS_Config_r19_Lpwus_MO_1_1_r19_Sl40:
				(*ie.Lpwus_MO_1_1_r19).Sl40 = new(int64)
				v, err := d.DecodeInteger(per.Constrained(0, 39))
				if err != nil {
					return err
				}
				(*(*ie.Lpwus_MO_1_1_r19).Sl40) = v
			case LPWUS_Config_r19_Lpwus_MO_1_1_r19_Sl80:
				(*ie.Lpwus_MO_1_1_r19).Sl80 = new(int64)
				v, err := d.DecodeInteger(per.Constrained(0, 79))
				if err != nil {
					return err
				}
				(*(*ie.Lpwus_MO_1_1_r19).Sl80) = v
			case LPWUS_Config_r19_Lpwus_MO_1_1_r19_Sl160:
				(*ie.Lpwus_MO_1_1_r19).Sl160 = new(int64)
				v, err := d.DecodeInteger(per.Constrained(0, 159))
				if err != nil {
					return err
				}
				(*(*ie.Lpwus_MO_1_1_r19).Sl160) = v
			case LPWUS_Config_r19_Lpwus_MO_1_1_r19_Sl320:
				(*ie.Lpwus_MO_1_1_r19).Sl320 = new(int64)
				v, err := d.DecodeInteger(per.Constrained(0, 319))
				if err != nil {
					return err
				}
				(*(*ie.Lpwus_MO_1_1_r19).Sl320) = v
			case LPWUS_Config_r19_Lpwus_MO_1_1_r19_Sl640:
				(*ie.Lpwus_MO_1_1_r19).Sl640 = new(int64)
				v, err := d.DecodeInteger(per.Constrained(0, 639))
				if err != nil {
					return err
				}
				(*(*ie.Lpwus_MO_1_1_r19).Sl640) = v
			case LPWUS_Config_r19_Lpwus_MO_1_1_r19_Sl1280:
				(*ie.Lpwus_MO_1_1_r19).Sl1280 = new(int64)
				v, err := d.DecodeInteger(per.Constrained(0, 1279))
				if err != nil {
					return err
				}
				(*(*ie.Lpwus_MO_1_1_r19).Sl1280) = v
			case LPWUS_Config_r19_Lpwus_MO_1_1_r19_Sl2560:
				(*ie.Lpwus_MO_1_1_r19).Sl2560 = new(int64)
				v, err := d.DecodeInteger(per.Constrained(0, 2559))
				if err != nil {
					return err
				}
				(*(*ie.Lpwus_MO_1_1_r19).Sl2560) = v
			case LPWUS_Config_r19_Lpwus_MO_1_1_r19_Sl5120:
				(*ie.Lpwus_MO_1_1_r19).Sl5120 = new(int64)
				v, err := d.DecodeInteger(per.Constrained(0, 5119))
				if err != nil {
					return err
				}
				(*(*ie.Lpwus_MO_1_1_r19).Sl5120) = v
			case LPWUS_Config_r19_Lpwus_MO_1_1_r19_Sl10240:
				(*ie.Lpwus_MO_1_1_r19).Sl10240 = new(int64)
				v, err := d.DecodeInteger(per.Constrained(0, 10239))
				if err != nil {
					return err
				}
				(*(*ie.Lpwus_MO_1_1_r19).Sl10240) = v
			}
		}
	}

	// 8. lpwus-MO-1-2-r19: choice
	{
		if seq.IsComponentPresent(6) {
			ie.Lpwus_MO_1_2_r19 = &struct {
				Choice int
				Sl1    *struct{}
				Sl2    *int64
				Sl4    *int64
				Sl5    *int64
				Sl8    *int64
				Sl10   *int64
				Sl16   *int64
				Sl20   *int64
				Sl40   *int64
				Sl80   *int64
			}{}
			choiceDec := d.NewChoiceDecoder(lPWUS_Config_r19LpwusMO12R19Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Lpwus_MO_1_2_r19).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case LPWUS_Config_r19_Lpwus_MO_1_2_r19_Sl1:
				(*ie.Lpwus_MO_1_2_r19).Sl1 = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case LPWUS_Config_r19_Lpwus_MO_1_2_r19_Sl2:
				(*ie.Lpwus_MO_1_2_r19).Sl2 = new(int64)
				v, err := d.DecodeInteger(per.Constrained(0, 1))
				if err != nil {
					return err
				}
				(*(*ie.Lpwus_MO_1_2_r19).Sl2) = v
			case LPWUS_Config_r19_Lpwus_MO_1_2_r19_Sl4:
				(*ie.Lpwus_MO_1_2_r19).Sl4 = new(int64)
				v, err := d.DecodeInteger(per.Constrained(0, 3))
				if err != nil {
					return err
				}
				(*(*ie.Lpwus_MO_1_2_r19).Sl4) = v
			case LPWUS_Config_r19_Lpwus_MO_1_2_r19_Sl5:
				(*ie.Lpwus_MO_1_2_r19).Sl5 = new(int64)
				v, err := d.DecodeInteger(per.Constrained(0, 4))
				if err != nil {
					return err
				}
				(*(*ie.Lpwus_MO_1_2_r19).Sl5) = v
			case LPWUS_Config_r19_Lpwus_MO_1_2_r19_Sl8:
				(*ie.Lpwus_MO_1_2_r19).Sl8 = new(int64)
				v, err := d.DecodeInteger(per.Constrained(0, 7))
				if err != nil {
					return err
				}
				(*(*ie.Lpwus_MO_1_2_r19).Sl8) = v
			case LPWUS_Config_r19_Lpwus_MO_1_2_r19_Sl10:
				(*ie.Lpwus_MO_1_2_r19).Sl10 = new(int64)
				v, err := d.DecodeInteger(per.Constrained(0, 9))
				if err != nil {
					return err
				}
				(*(*ie.Lpwus_MO_1_2_r19).Sl10) = v
			case LPWUS_Config_r19_Lpwus_MO_1_2_r19_Sl16:
				(*ie.Lpwus_MO_1_2_r19).Sl16 = new(int64)
				v, err := d.DecodeInteger(per.Constrained(0, 15))
				if err != nil {
					return err
				}
				(*(*ie.Lpwus_MO_1_2_r19).Sl16) = v
			case LPWUS_Config_r19_Lpwus_MO_1_2_r19_Sl20:
				(*ie.Lpwus_MO_1_2_r19).Sl20 = new(int64)
				v, err := d.DecodeInteger(per.Constrained(0, 19))
				if err != nil {
					return err
				}
				(*(*ie.Lpwus_MO_1_2_r19).Sl20) = v
			case LPWUS_Config_r19_Lpwus_MO_1_2_r19_Sl40:
				(*ie.Lpwus_MO_1_2_r19).Sl40 = new(int64)
				v, err := d.DecodeInteger(per.Constrained(0, 39))
				if err != nil {
					return err
				}
				(*(*ie.Lpwus_MO_1_2_r19).Sl40) = v
			case LPWUS_Config_r19_Lpwus_MO_1_2_r19_Sl80:
				(*ie.Lpwus_MO_1_2_r19).Sl80 = new(int64)
				v, err := d.DecodeInteger(per.Constrained(0, 79))
				if err != nil {
					return err
				}
				(*(*ie.Lpwus_MO_1_2_r19).Sl80) = v
			}
		}
	}

	// 9. lpwus-NumOfMO-1-1-r19: integer
	{
		if seq.IsComponentPresent(7) {
			v, err := d.DecodeInteger(lPWUSConfigR19LpwusNumOfMO11R19Constraints)
			if err != nil {
				return err
			}
			ie.Lpwus_NumOfMO_1_1_r19 = &v
		}
	}

	// 10. lpwus-NumOfMO-1-2-r19: integer
	{
		if seq.IsComponentPresent(8) {
			v, err := d.DecodeInteger(lPWUSConfigR19LpwusNumOfMO12R19Constraints)
			if err != nil {
				return err
			}
			ie.Lpwus_NumOfMO_1_2_r19 = &v
		}
	}

	// 11. lpwus-TCI-States-r19: sequence-of
	{
		if seq.IsComponentPresent(9) {
			seqOf := d.NewSequenceOfDecoder(lPWUSConfigR19LpwusTCIStatesR19Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Lpwus_TCI_States_r19 = make([]int64, n)
			for i := int64(0); i < n; i++ {
				v, err := d.DecodeInteger(per.Constrained(0, 15))
				if err != nil {
					return err
				}
				ie.Lpwus_TCI_States_r19[i] = v
			}
		}
	}

	// 12. lpwus-NominalMoDuration-r19: integer
	{
		if seq.IsComponentPresent(10) {
			v, err := d.DecodeInteger(lPWUSConfigR19LpwusNominalMoDurationR19Constraints)
			if err != nil {
				return err
			}
			ie.Lpwus_NominalMoDuration_r19 = &v
		}
	}

	// 13. lpwus-ActualDuration-r19: choice
	{
		if seq.IsComponentPresent(11) {
			ie.Lpwus_ActualDuration_r19 = &struct {
				Choice  int
				MValue1 *int64
				MValue2 *int64
				MValue4 *int64
			}{}
			choiceDec := d.NewChoiceDecoder(lPWUS_Config_r19LpwusActualDurationR19Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Lpwus_ActualDuration_r19).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case LPWUS_Config_r19_Lpwus_ActualDuration_r19_MValue1:
				(*ie.Lpwus_ActualDuration_r19).MValue1 = new(int64)
				v, err := d.DecodeInteger(per.Constrained(2, 64))
				if err != nil {
					return err
				}
				(*(*ie.Lpwus_ActualDuration_r19).MValue1) = v
			case LPWUS_Config_r19_Lpwus_ActualDuration_r19_MValue2:
				(*ie.Lpwus_ActualDuration_r19).MValue2 = new(int64)
				v, err := d.DecodeInteger(per.Constrained(1, 32))
				if err != nil {
					return err
				}
				(*(*ie.Lpwus_ActualDuration_r19).MValue2) = v
			case LPWUS_Config_r19_Lpwus_ActualDuration_r19_MValue4:
				(*ie.Lpwus_ActualDuration_r19).MValue4 = new(int64)
				v, err := d.DecodeInteger(per.Constrained(1, 16))
				if err != nil {
					return err
				}
				(*(*ie.Lpwus_ActualDuration_r19).MValue4) = v
			}
		}
	}

	// 14. lpwus-AvailableSlot-r19: choice
	{
		if seq.IsComponentPresent(12) {
			ie.Lpwus_AvailableSlot_r19 = &struct {
				Choice int
				N10    *per.BitString
				N20    *per.BitString
				N40    *per.BitString
			}{}
			choiceDec := d.NewChoiceDecoder(lPWUS_Config_r19LpwusAvailableSlotR19Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Lpwus_AvailableSlot_r19).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case LPWUS_Config_r19_Lpwus_AvailableSlot_r19_N10:
				(*ie.Lpwus_AvailableSlot_r19).N10 = new(per.BitString)
				v, err := d.DecodeBitString(per.FixedSize(10))
				if err != nil {
					return err
				}
				(*(*ie.Lpwus_AvailableSlot_r19).N10) = v
			case LPWUS_Config_r19_Lpwus_AvailableSlot_r19_N20:
				(*ie.Lpwus_AvailableSlot_r19).N20 = new(per.BitString)
				v, err := d.DecodeBitString(per.FixedSize(20))
				if err != nil {
					return err
				}
				(*(*ie.Lpwus_AvailableSlot_r19).N20) = v
			case LPWUS_Config_r19_Lpwus_AvailableSlot_r19_N40:
				(*ie.Lpwus_AvailableSlot_r19).N40 = new(per.BitString)
				v, err := d.DecodeBitString(per.FixedSize(40))
				if err != nil {
					return err
				}
				(*(*ie.Lpwus_AvailableSlot_r19).N40) = v
			}
		}
	}

	// 15. lpwus-AvailableSymbol-r19: choice
	{
		if seq.IsComponentPresent(13) {
			ie.Lpwus_AvailableSymbol_r19 = &struct {
				Choice   int
				OneSlot  *per.BitString
				TwoSlots *per.BitString
			}{}
			choiceDec := d.NewChoiceDecoder(lPWUS_Config_r19LpwusAvailableSymbolR19Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Lpwus_AvailableSymbol_r19).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case LPWUS_Config_r19_Lpwus_AvailableSymbol_r19_OneSlot:
				(*ie.Lpwus_AvailableSymbol_r19).OneSlot = new(per.BitString)
				v, err := d.DecodeBitString(per.FixedSize(14))
				if err != nil {
					return err
				}
				(*(*ie.Lpwus_AvailableSymbol_r19).OneSlot) = v
			case LPWUS_Config_r19_Lpwus_AvailableSymbol_r19_TwoSlots:
				(*ie.Lpwus_AvailableSymbol_r19).TwoSlots = new(per.BitString)
				v, err := d.DecodeBitString(per.FixedSize(28))
				if err != nil {
					return err
				}
				(*(*ie.Lpwus_AvailableSymbol_r19).TwoSlots) = v
			}
		}
	}

	// 16. lpwus-TransmitOtherPeriodicCSI-r19: enumerated
	{
		if seq.IsComponentPresent(14) {
			idx, err := d.DecodeEnumerated(lPWUSConfigR19LpwusTransmitOtherPeriodicCSIR19Constraints)
			if err != nil {
				return err
			}
			ie.Lpwus_TransmitOtherPeriodicCSI_r19 = &idx
		}
	}

	// 17. lpwus-TransmitPeriodicL1-RSRP-r19: enumerated
	{
		if seq.IsComponentPresent(15) {
			idx, err := d.DecodeEnumerated(lPWUSConfigR19LpwusTransmitPeriodicL1RSRPR19Constraints)
			if err != nil {
				return err
			}
			ie.Lpwus_TransmitPeriodicL1_RSRP_r19 = &idx
		}
	}

	// 18. lpwus-Codepoint-r19: sequence-of
	{
		seqOf := d.NewSequenceOfDecoder(lPWUSConfigR19LpwusCodepointR19Constraints)
		n, err := seqOf.DecodeLength()
		if err != nil {
			return err
		}
		ie.Lpwus_Codepoint_r19 = make([]per.BitString, n)
		for i := int64(0); i < n; i++ {
			v, err := d.DecodeBitString(per.SizeRange(1, 5))
			if err != nil {
				return err
			}
			ie.Lpwus_Codepoint_r19[i] = v
		}
	}

	// 19. lpwus-NumOfBits-r19: integer
	{
		if seq.IsComponentPresent(17) {
			v, err := d.DecodeInteger(lPWUSConfigR19LpwusNumOfBitsR19Constraints)
			if err != nil {
				return err
			}
			ie.Lpwus_NumOfBits_r19 = &v
		}
	}

	// 20. lpwus-PDCCH-MonitoringTimer-r19: sequence
	{
		if seq.IsComponentPresent(18) {
			ie.Lpwus_PDCCH_MonitoringTimer_r19 = &struct {
				Timer1_r19 struct {
					Choice          int
					SubMilliSeconds *int64
					MilliSeconds    *int64
				}
				Timer2_r19 *struct {
					Choice          int
					SubMilliSeconds *int64
					MilliSeconds    *int64
				}
			}{}
			c := ie.Lpwus_PDCCH_MonitoringTimer_r19
			lPWUSConfigR19LpwusPDCCHMonitoringTimerR19Seq := d.NewSequenceDecoder(lPWUSConfigR19LpwusPDCCHMonitoringTimerR19Constraints)
			if err := lPWUSConfigR19LpwusPDCCHMonitoringTimerR19Seq.DecodePreamble(); err != nil {
				return err
			}
			{
				choiceDec := d.NewChoiceDecoder(lPWUSConfigR19LpwusPDCCHMonitoringTimerR19Timer1R19Constraints)
				index, _, _, err := choiceDec.DecodeChoice()
				if err != nil {
					return err
				}
				c.Timer1_r19.Choice = int(index)
				switch index {
				case LPWUS_Config_r19_Lpwus_PDCCH_MonitoringTimer_r19_Timer1_r19_SubMilliSeconds:
					c.Timer1_r19.SubMilliSeconds = new(int64)
					v, err := d.DecodeInteger(per.Constrained(1, 31))
					if err != nil {
						return err
					}
					(*c.Timer1_r19.SubMilliSeconds) = v
				case LPWUS_Config_r19_Lpwus_PDCCH_MonitoringTimer_r19_Timer1_r19_MilliSeconds:
					c.Timer1_r19.MilliSeconds = new(int64)
					v, err := d.DecodeEnumerated(lPWUSConfigR19LpwusPDCCHMonitoringTimerR19Timer1R19MilliSecondsConstraints)
					if err != nil {
						return err
					}
					(*c.Timer1_r19.MilliSeconds) = v
				}
			}
			if lPWUSConfigR19LpwusPDCCHMonitoringTimerR19Seq.IsComponentPresent(1) {
				c.Timer2_r19 = &struct {
					Choice          int
					SubMilliSeconds *int64
					MilliSeconds    *int64
				}{}
				choiceDec := d.NewChoiceDecoder(lPWUSConfigR19LpwusPDCCHMonitoringTimerR19Timer2R19Constraints)
				index, _, _, err := choiceDec.DecodeChoice()
				if err != nil {
					return err
				}
				(*c.Timer2_r19).Choice = int(index)
				switch index {
				case LPWUS_Config_r19_Lpwus_PDCCH_MonitoringTimer_r19_Timer2_r19_SubMilliSeconds:
					(*c.Timer2_r19).SubMilliSeconds = new(int64)
					v, err := d.DecodeInteger(per.Constrained(1, 31))
					if err != nil {
						return err
					}
					(*(*c.Timer2_r19).SubMilliSeconds) = v
				case LPWUS_Config_r19_Lpwus_PDCCH_MonitoringTimer_r19_Timer2_r19_MilliSeconds:
					(*c.Timer2_r19).MilliSeconds = new(int64)
					v, err := d.DecodeEnumerated(lPWUSConfigR19LpwusPDCCHMonitoringTimerR19Timer2R19MilliSecondsConstraints)
					if err != nil {
						return err
					}
					(*(*c.Timer2_r19).MilliSeconds) = v
				}
			}
		}
	}

	return nil
}
