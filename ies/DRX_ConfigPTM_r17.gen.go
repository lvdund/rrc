// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: DRX-ConfigPTM-r17 (line 28406).

var dRXConfigPTMR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "drx-onDurationTimerPTM-r17"},
		{Name: "drx-InactivityTimerPTM-r17"},
		{Name: "drx-HARQ-RTT-TimerDL-PTM-r17", Optional: true},
		{Name: "drx-RetransmissionTimerDL-PTM-r17", Optional: true},
		{Name: "drx-LongCycleStartOffsetPTM-r17"},
		{Name: "drx-SlotOffsetPTM-r17"},
	},
}

var dRX_ConfigPTM_r17DrxOnDurationTimerPTMR17Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "subMilliSeconds"},
		{Name: "milliSeconds"},
	},
}

const (
	DRX_ConfigPTM_r17_Drx_OnDurationTimerPTM_r17_SubMilliSeconds = 0
	DRX_ConfigPTM_r17_Drx_OnDurationTimerPTM_r17_MilliSeconds    = 1
)

const (
	DRX_ConfigPTM_r17_Drx_InactivityTimerPTM_r17_Ms0    = 0
	DRX_ConfigPTM_r17_Drx_InactivityTimerPTM_r17_Ms1    = 1
	DRX_ConfigPTM_r17_Drx_InactivityTimerPTM_r17_Ms2    = 2
	DRX_ConfigPTM_r17_Drx_InactivityTimerPTM_r17_Ms3    = 3
	DRX_ConfigPTM_r17_Drx_InactivityTimerPTM_r17_Ms4    = 4
	DRX_ConfigPTM_r17_Drx_InactivityTimerPTM_r17_Ms5    = 5
	DRX_ConfigPTM_r17_Drx_InactivityTimerPTM_r17_Ms6    = 6
	DRX_ConfigPTM_r17_Drx_InactivityTimerPTM_r17_Ms8    = 7
	DRX_ConfigPTM_r17_Drx_InactivityTimerPTM_r17_Ms10   = 8
	DRX_ConfigPTM_r17_Drx_InactivityTimerPTM_r17_Ms20   = 9
	DRX_ConfigPTM_r17_Drx_InactivityTimerPTM_r17_Ms30   = 10
	DRX_ConfigPTM_r17_Drx_InactivityTimerPTM_r17_Ms40   = 11
	DRX_ConfigPTM_r17_Drx_InactivityTimerPTM_r17_Ms50   = 12
	DRX_ConfigPTM_r17_Drx_InactivityTimerPTM_r17_Ms60   = 13
	DRX_ConfigPTM_r17_Drx_InactivityTimerPTM_r17_Ms80   = 14
	DRX_ConfigPTM_r17_Drx_InactivityTimerPTM_r17_Ms100  = 15
	DRX_ConfigPTM_r17_Drx_InactivityTimerPTM_r17_Ms200  = 16
	DRX_ConfigPTM_r17_Drx_InactivityTimerPTM_r17_Ms300  = 17
	DRX_ConfigPTM_r17_Drx_InactivityTimerPTM_r17_Ms500  = 18
	DRX_ConfigPTM_r17_Drx_InactivityTimerPTM_r17_Ms750  = 19
	DRX_ConfigPTM_r17_Drx_InactivityTimerPTM_r17_Ms1280 = 20
	DRX_ConfigPTM_r17_Drx_InactivityTimerPTM_r17_Ms1920 = 21
	DRX_ConfigPTM_r17_Drx_InactivityTimerPTM_r17_Ms2560 = 22
	DRX_ConfigPTM_r17_Drx_InactivityTimerPTM_r17_Spare9 = 23
	DRX_ConfigPTM_r17_Drx_InactivityTimerPTM_r17_Spare8 = 24
	DRX_ConfigPTM_r17_Drx_InactivityTimerPTM_r17_Spare7 = 25
	DRX_ConfigPTM_r17_Drx_InactivityTimerPTM_r17_Spare6 = 26
	DRX_ConfigPTM_r17_Drx_InactivityTimerPTM_r17_Spare5 = 27
	DRX_ConfigPTM_r17_Drx_InactivityTimerPTM_r17_Spare4 = 28
	DRX_ConfigPTM_r17_Drx_InactivityTimerPTM_r17_Spare3 = 29
	DRX_ConfigPTM_r17_Drx_InactivityTimerPTM_r17_Spare2 = 30
	DRX_ConfigPTM_r17_Drx_InactivityTimerPTM_r17_Spare1 = 31
)

var dRXConfigPTMR17DrxInactivityTimerPTMR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{DRX_ConfigPTM_r17_Drx_InactivityTimerPTM_r17_Ms0, DRX_ConfigPTM_r17_Drx_InactivityTimerPTM_r17_Ms1, DRX_ConfigPTM_r17_Drx_InactivityTimerPTM_r17_Ms2, DRX_ConfigPTM_r17_Drx_InactivityTimerPTM_r17_Ms3, DRX_ConfigPTM_r17_Drx_InactivityTimerPTM_r17_Ms4, DRX_ConfigPTM_r17_Drx_InactivityTimerPTM_r17_Ms5, DRX_ConfigPTM_r17_Drx_InactivityTimerPTM_r17_Ms6, DRX_ConfigPTM_r17_Drx_InactivityTimerPTM_r17_Ms8, DRX_ConfigPTM_r17_Drx_InactivityTimerPTM_r17_Ms10, DRX_ConfigPTM_r17_Drx_InactivityTimerPTM_r17_Ms20, DRX_ConfigPTM_r17_Drx_InactivityTimerPTM_r17_Ms30, DRX_ConfigPTM_r17_Drx_InactivityTimerPTM_r17_Ms40, DRX_ConfigPTM_r17_Drx_InactivityTimerPTM_r17_Ms50, DRX_ConfigPTM_r17_Drx_InactivityTimerPTM_r17_Ms60, DRX_ConfigPTM_r17_Drx_InactivityTimerPTM_r17_Ms80, DRX_ConfigPTM_r17_Drx_InactivityTimerPTM_r17_Ms100, DRX_ConfigPTM_r17_Drx_InactivityTimerPTM_r17_Ms200, DRX_ConfigPTM_r17_Drx_InactivityTimerPTM_r17_Ms300, DRX_ConfigPTM_r17_Drx_InactivityTimerPTM_r17_Ms500, DRX_ConfigPTM_r17_Drx_InactivityTimerPTM_r17_Ms750, DRX_ConfigPTM_r17_Drx_InactivityTimerPTM_r17_Ms1280, DRX_ConfigPTM_r17_Drx_InactivityTimerPTM_r17_Ms1920, DRX_ConfigPTM_r17_Drx_InactivityTimerPTM_r17_Ms2560, DRX_ConfigPTM_r17_Drx_InactivityTimerPTM_r17_Spare9, DRX_ConfigPTM_r17_Drx_InactivityTimerPTM_r17_Spare8, DRX_ConfigPTM_r17_Drx_InactivityTimerPTM_r17_Spare7, DRX_ConfigPTM_r17_Drx_InactivityTimerPTM_r17_Spare6, DRX_ConfigPTM_r17_Drx_InactivityTimerPTM_r17_Spare5, DRX_ConfigPTM_r17_Drx_InactivityTimerPTM_r17_Spare4, DRX_ConfigPTM_r17_Drx_InactivityTimerPTM_r17_Spare3, DRX_ConfigPTM_r17_Drx_InactivityTimerPTM_r17_Spare2, DRX_ConfigPTM_r17_Drx_InactivityTimerPTM_r17_Spare1},
}

var dRXConfigPTMR17DrxHARQRTTTimerDLPTMR17Constraints = per.Constrained(0, 56)

const (
	DRX_ConfigPTM_r17_Drx_RetransmissionTimerDL_PTM_r17_Sl0     = 0
	DRX_ConfigPTM_r17_Drx_RetransmissionTimerDL_PTM_r17_Sl1     = 1
	DRX_ConfigPTM_r17_Drx_RetransmissionTimerDL_PTM_r17_Sl2     = 2
	DRX_ConfigPTM_r17_Drx_RetransmissionTimerDL_PTM_r17_Sl4     = 3
	DRX_ConfigPTM_r17_Drx_RetransmissionTimerDL_PTM_r17_Sl6     = 4
	DRX_ConfigPTM_r17_Drx_RetransmissionTimerDL_PTM_r17_Sl8     = 5
	DRX_ConfigPTM_r17_Drx_RetransmissionTimerDL_PTM_r17_Sl16    = 6
	DRX_ConfigPTM_r17_Drx_RetransmissionTimerDL_PTM_r17_Sl24    = 7
	DRX_ConfigPTM_r17_Drx_RetransmissionTimerDL_PTM_r17_Sl33    = 8
	DRX_ConfigPTM_r17_Drx_RetransmissionTimerDL_PTM_r17_Sl40    = 9
	DRX_ConfigPTM_r17_Drx_RetransmissionTimerDL_PTM_r17_Sl64    = 10
	DRX_ConfigPTM_r17_Drx_RetransmissionTimerDL_PTM_r17_Sl80    = 11
	DRX_ConfigPTM_r17_Drx_RetransmissionTimerDL_PTM_r17_Sl96    = 12
	DRX_ConfigPTM_r17_Drx_RetransmissionTimerDL_PTM_r17_Sl112   = 13
	DRX_ConfigPTM_r17_Drx_RetransmissionTimerDL_PTM_r17_Sl128   = 14
	DRX_ConfigPTM_r17_Drx_RetransmissionTimerDL_PTM_r17_Sl160   = 15
	DRX_ConfigPTM_r17_Drx_RetransmissionTimerDL_PTM_r17_Sl320   = 16
	DRX_ConfigPTM_r17_Drx_RetransmissionTimerDL_PTM_r17_Spare15 = 17
	DRX_ConfigPTM_r17_Drx_RetransmissionTimerDL_PTM_r17_Spare14 = 18
	DRX_ConfigPTM_r17_Drx_RetransmissionTimerDL_PTM_r17_Spare13 = 19
	DRX_ConfigPTM_r17_Drx_RetransmissionTimerDL_PTM_r17_Spare12 = 20
	DRX_ConfigPTM_r17_Drx_RetransmissionTimerDL_PTM_r17_Spare11 = 21
	DRX_ConfigPTM_r17_Drx_RetransmissionTimerDL_PTM_r17_Spare10 = 22
	DRX_ConfigPTM_r17_Drx_RetransmissionTimerDL_PTM_r17_Spare9  = 23
	DRX_ConfigPTM_r17_Drx_RetransmissionTimerDL_PTM_r17_Spare8  = 24
	DRX_ConfigPTM_r17_Drx_RetransmissionTimerDL_PTM_r17_Spare7  = 25
	DRX_ConfigPTM_r17_Drx_RetransmissionTimerDL_PTM_r17_Spare6  = 26
	DRX_ConfigPTM_r17_Drx_RetransmissionTimerDL_PTM_r17_Spare5  = 27
	DRX_ConfigPTM_r17_Drx_RetransmissionTimerDL_PTM_r17_Spare4  = 28
	DRX_ConfigPTM_r17_Drx_RetransmissionTimerDL_PTM_r17_Spare3  = 29
	DRX_ConfigPTM_r17_Drx_RetransmissionTimerDL_PTM_r17_Spare2  = 30
	DRX_ConfigPTM_r17_Drx_RetransmissionTimerDL_PTM_r17_Spare1  = 31
)

var dRXConfigPTMR17DrxRetransmissionTimerDLPTMR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{DRX_ConfigPTM_r17_Drx_RetransmissionTimerDL_PTM_r17_Sl0, DRX_ConfigPTM_r17_Drx_RetransmissionTimerDL_PTM_r17_Sl1, DRX_ConfigPTM_r17_Drx_RetransmissionTimerDL_PTM_r17_Sl2, DRX_ConfigPTM_r17_Drx_RetransmissionTimerDL_PTM_r17_Sl4, DRX_ConfigPTM_r17_Drx_RetransmissionTimerDL_PTM_r17_Sl6, DRX_ConfigPTM_r17_Drx_RetransmissionTimerDL_PTM_r17_Sl8, DRX_ConfigPTM_r17_Drx_RetransmissionTimerDL_PTM_r17_Sl16, DRX_ConfigPTM_r17_Drx_RetransmissionTimerDL_PTM_r17_Sl24, DRX_ConfigPTM_r17_Drx_RetransmissionTimerDL_PTM_r17_Sl33, DRX_ConfigPTM_r17_Drx_RetransmissionTimerDL_PTM_r17_Sl40, DRX_ConfigPTM_r17_Drx_RetransmissionTimerDL_PTM_r17_Sl64, DRX_ConfigPTM_r17_Drx_RetransmissionTimerDL_PTM_r17_Sl80, DRX_ConfigPTM_r17_Drx_RetransmissionTimerDL_PTM_r17_Sl96, DRX_ConfigPTM_r17_Drx_RetransmissionTimerDL_PTM_r17_Sl112, DRX_ConfigPTM_r17_Drx_RetransmissionTimerDL_PTM_r17_Sl128, DRX_ConfigPTM_r17_Drx_RetransmissionTimerDL_PTM_r17_Sl160, DRX_ConfigPTM_r17_Drx_RetransmissionTimerDL_PTM_r17_Sl320, DRX_ConfigPTM_r17_Drx_RetransmissionTimerDL_PTM_r17_Spare15, DRX_ConfigPTM_r17_Drx_RetransmissionTimerDL_PTM_r17_Spare14, DRX_ConfigPTM_r17_Drx_RetransmissionTimerDL_PTM_r17_Spare13, DRX_ConfigPTM_r17_Drx_RetransmissionTimerDL_PTM_r17_Spare12, DRX_ConfigPTM_r17_Drx_RetransmissionTimerDL_PTM_r17_Spare11, DRX_ConfigPTM_r17_Drx_RetransmissionTimerDL_PTM_r17_Spare10, DRX_ConfigPTM_r17_Drx_RetransmissionTimerDL_PTM_r17_Spare9, DRX_ConfigPTM_r17_Drx_RetransmissionTimerDL_PTM_r17_Spare8, DRX_ConfigPTM_r17_Drx_RetransmissionTimerDL_PTM_r17_Spare7, DRX_ConfigPTM_r17_Drx_RetransmissionTimerDL_PTM_r17_Spare6, DRX_ConfigPTM_r17_Drx_RetransmissionTimerDL_PTM_r17_Spare5, DRX_ConfigPTM_r17_Drx_RetransmissionTimerDL_PTM_r17_Spare4, DRX_ConfigPTM_r17_Drx_RetransmissionTimerDL_PTM_r17_Spare3, DRX_ConfigPTM_r17_Drx_RetransmissionTimerDL_PTM_r17_Spare2, DRX_ConfigPTM_r17_Drx_RetransmissionTimerDL_PTM_r17_Spare1},
}

var dRX_ConfigPTM_r17DrxLongCycleStartOffsetPTMR17Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "ms10"},
		{Name: "ms20"},
		{Name: "ms32"},
		{Name: "ms40"},
		{Name: "ms60"},
		{Name: "ms64"},
		{Name: "ms70"},
		{Name: "ms80"},
		{Name: "ms128"},
		{Name: "ms160"},
		{Name: "ms256"},
		{Name: "ms320"},
		{Name: "ms512"},
		{Name: "ms640"},
		{Name: "ms1024"},
		{Name: "ms1280"},
		{Name: "ms2048"},
		{Name: "ms2560"},
		{Name: "ms5120"},
		{Name: "ms10240"},
	},
}

const (
	DRX_ConfigPTM_r17_Drx_LongCycleStartOffsetPTM_r17_Ms10    = 0
	DRX_ConfigPTM_r17_Drx_LongCycleStartOffsetPTM_r17_Ms20    = 1
	DRX_ConfigPTM_r17_Drx_LongCycleStartOffsetPTM_r17_Ms32    = 2
	DRX_ConfigPTM_r17_Drx_LongCycleStartOffsetPTM_r17_Ms40    = 3
	DRX_ConfigPTM_r17_Drx_LongCycleStartOffsetPTM_r17_Ms60    = 4
	DRX_ConfigPTM_r17_Drx_LongCycleStartOffsetPTM_r17_Ms64    = 5
	DRX_ConfigPTM_r17_Drx_LongCycleStartOffsetPTM_r17_Ms70    = 6
	DRX_ConfigPTM_r17_Drx_LongCycleStartOffsetPTM_r17_Ms80    = 7
	DRX_ConfigPTM_r17_Drx_LongCycleStartOffsetPTM_r17_Ms128   = 8
	DRX_ConfigPTM_r17_Drx_LongCycleStartOffsetPTM_r17_Ms160   = 9
	DRX_ConfigPTM_r17_Drx_LongCycleStartOffsetPTM_r17_Ms256   = 10
	DRX_ConfigPTM_r17_Drx_LongCycleStartOffsetPTM_r17_Ms320   = 11
	DRX_ConfigPTM_r17_Drx_LongCycleStartOffsetPTM_r17_Ms512   = 12
	DRX_ConfigPTM_r17_Drx_LongCycleStartOffsetPTM_r17_Ms640   = 13
	DRX_ConfigPTM_r17_Drx_LongCycleStartOffsetPTM_r17_Ms1024  = 14
	DRX_ConfigPTM_r17_Drx_LongCycleStartOffsetPTM_r17_Ms1280  = 15
	DRX_ConfigPTM_r17_Drx_LongCycleStartOffsetPTM_r17_Ms2048  = 16
	DRX_ConfigPTM_r17_Drx_LongCycleStartOffsetPTM_r17_Ms2560  = 17
	DRX_ConfigPTM_r17_Drx_LongCycleStartOffsetPTM_r17_Ms5120  = 18
	DRX_ConfigPTM_r17_Drx_LongCycleStartOffsetPTM_r17_Ms10240 = 19
)

var dRXConfigPTMR17DrxSlotOffsetPTMR17Constraints = per.Constrained(0, 31)

const (
	DRX_ConfigPTM_r17_Drx_OnDurationTimerPTM_r17_MilliSeconds_Ms1    = 0
	DRX_ConfigPTM_r17_Drx_OnDurationTimerPTM_r17_MilliSeconds_Ms2    = 1
	DRX_ConfigPTM_r17_Drx_OnDurationTimerPTM_r17_MilliSeconds_Ms3    = 2
	DRX_ConfigPTM_r17_Drx_OnDurationTimerPTM_r17_MilliSeconds_Ms4    = 3
	DRX_ConfigPTM_r17_Drx_OnDurationTimerPTM_r17_MilliSeconds_Ms5    = 4
	DRX_ConfigPTM_r17_Drx_OnDurationTimerPTM_r17_MilliSeconds_Ms6    = 5
	DRX_ConfigPTM_r17_Drx_OnDurationTimerPTM_r17_MilliSeconds_Ms8    = 6
	DRX_ConfigPTM_r17_Drx_OnDurationTimerPTM_r17_MilliSeconds_Ms10   = 7
	DRX_ConfigPTM_r17_Drx_OnDurationTimerPTM_r17_MilliSeconds_Ms20   = 8
	DRX_ConfigPTM_r17_Drx_OnDurationTimerPTM_r17_MilliSeconds_Ms30   = 9
	DRX_ConfigPTM_r17_Drx_OnDurationTimerPTM_r17_MilliSeconds_Ms40   = 10
	DRX_ConfigPTM_r17_Drx_OnDurationTimerPTM_r17_MilliSeconds_Ms50   = 11
	DRX_ConfigPTM_r17_Drx_OnDurationTimerPTM_r17_MilliSeconds_Ms60   = 12
	DRX_ConfigPTM_r17_Drx_OnDurationTimerPTM_r17_MilliSeconds_Ms80   = 13
	DRX_ConfigPTM_r17_Drx_OnDurationTimerPTM_r17_MilliSeconds_Ms100  = 14
	DRX_ConfigPTM_r17_Drx_OnDurationTimerPTM_r17_MilliSeconds_Ms200  = 15
	DRX_ConfigPTM_r17_Drx_OnDurationTimerPTM_r17_MilliSeconds_Ms300  = 16
	DRX_ConfigPTM_r17_Drx_OnDurationTimerPTM_r17_MilliSeconds_Ms400  = 17
	DRX_ConfigPTM_r17_Drx_OnDurationTimerPTM_r17_MilliSeconds_Ms500  = 18
	DRX_ConfigPTM_r17_Drx_OnDurationTimerPTM_r17_MilliSeconds_Ms600  = 19
	DRX_ConfigPTM_r17_Drx_OnDurationTimerPTM_r17_MilliSeconds_Ms800  = 20
	DRX_ConfigPTM_r17_Drx_OnDurationTimerPTM_r17_MilliSeconds_Ms1000 = 21
	DRX_ConfigPTM_r17_Drx_OnDurationTimerPTM_r17_MilliSeconds_Ms1200 = 22
	DRX_ConfigPTM_r17_Drx_OnDurationTimerPTM_r17_MilliSeconds_Ms1600 = 23
	DRX_ConfigPTM_r17_Drx_OnDurationTimerPTM_r17_MilliSeconds_Spare8 = 24
	DRX_ConfigPTM_r17_Drx_OnDurationTimerPTM_r17_MilliSeconds_Spare7 = 25
	DRX_ConfigPTM_r17_Drx_OnDurationTimerPTM_r17_MilliSeconds_Spare6 = 26
	DRX_ConfigPTM_r17_Drx_OnDurationTimerPTM_r17_MilliSeconds_Spare5 = 27
	DRX_ConfigPTM_r17_Drx_OnDurationTimerPTM_r17_MilliSeconds_Spare4 = 28
	DRX_ConfigPTM_r17_Drx_OnDurationTimerPTM_r17_MilliSeconds_Spare3 = 29
	DRX_ConfigPTM_r17_Drx_OnDurationTimerPTM_r17_MilliSeconds_Spare2 = 30
	DRX_ConfigPTM_r17_Drx_OnDurationTimerPTM_r17_MilliSeconds_Spare1 = 31
)

var dRXConfigPTMR17DrxOnDurationTimerPTMR17MilliSecondsConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{DRX_ConfigPTM_r17_Drx_OnDurationTimerPTM_r17_MilliSeconds_Ms1, DRX_ConfigPTM_r17_Drx_OnDurationTimerPTM_r17_MilliSeconds_Ms2, DRX_ConfigPTM_r17_Drx_OnDurationTimerPTM_r17_MilliSeconds_Ms3, DRX_ConfigPTM_r17_Drx_OnDurationTimerPTM_r17_MilliSeconds_Ms4, DRX_ConfigPTM_r17_Drx_OnDurationTimerPTM_r17_MilliSeconds_Ms5, DRX_ConfigPTM_r17_Drx_OnDurationTimerPTM_r17_MilliSeconds_Ms6, DRX_ConfigPTM_r17_Drx_OnDurationTimerPTM_r17_MilliSeconds_Ms8, DRX_ConfigPTM_r17_Drx_OnDurationTimerPTM_r17_MilliSeconds_Ms10, DRX_ConfigPTM_r17_Drx_OnDurationTimerPTM_r17_MilliSeconds_Ms20, DRX_ConfigPTM_r17_Drx_OnDurationTimerPTM_r17_MilliSeconds_Ms30, DRX_ConfigPTM_r17_Drx_OnDurationTimerPTM_r17_MilliSeconds_Ms40, DRX_ConfigPTM_r17_Drx_OnDurationTimerPTM_r17_MilliSeconds_Ms50, DRX_ConfigPTM_r17_Drx_OnDurationTimerPTM_r17_MilliSeconds_Ms60, DRX_ConfigPTM_r17_Drx_OnDurationTimerPTM_r17_MilliSeconds_Ms80, DRX_ConfigPTM_r17_Drx_OnDurationTimerPTM_r17_MilliSeconds_Ms100, DRX_ConfigPTM_r17_Drx_OnDurationTimerPTM_r17_MilliSeconds_Ms200, DRX_ConfigPTM_r17_Drx_OnDurationTimerPTM_r17_MilliSeconds_Ms300, DRX_ConfigPTM_r17_Drx_OnDurationTimerPTM_r17_MilliSeconds_Ms400, DRX_ConfigPTM_r17_Drx_OnDurationTimerPTM_r17_MilliSeconds_Ms500, DRX_ConfigPTM_r17_Drx_OnDurationTimerPTM_r17_MilliSeconds_Ms600, DRX_ConfigPTM_r17_Drx_OnDurationTimerPTM_r17_MilliSeconds_Ms800, DRX_ConfigPTM_r17_Drx_OnDurationTimerPTM_r17_MilliSeconds_Ms1000, DRX_ConfigPTM_r17_Drx_OnDurationTimerPTM_r17_MilliSeconds_Ms1200, DRX_ConfigPTM_r17_Drx_OnDurationTimerPTM_r17_MilliSeconds_Ms1600, DRX_ConfigPTM_r17_Drx_OnDurationTimerPTM_r17_MilliSeconds_Spare8, DRX_ConfigPTM_r17_Drx_OnDurationTimerPTM_r17_MilliSeconds_Spare7, DRX_ConfigPTM_r17_Drx_OnDurationTimerPTM_r17_MilliSeconds_Spare6, DRX_ConfigPTM_r17_Drx_OnDurationTimerPTM_r17_MilliSeconds_Spare5, DRX_ConfigPTM_r17_Drx_OnDurationTimerPTM_r17_MilliSeconds_Spare4, DRX_ConfigPTM_r17_Drx_OnDurationTimerPTM_r17_MilliSeconds_Spare3, DRX_ConfigPTM_r17_Drx_OnDurationTimerPTM_r17_MilliSeconds_Spare2, DRX_ConfigPTM_r17_Drx_OnDurationTimerPTM_r17_MilliSeconds_Spare1},
}

type DRX_ConfigPTM_r17 struct {
	Drx_OnDurationTimerPTM_r17 struct {
		Choice          int
		SubMilliSeconds *int64
		MilliSeconds    *int64
	}
	Drx_InactivityTimerPTM_r17        int64
	Drx_HARQ_RTT_TimerDL_PTM_r17      *int64
	Drx_RetransmissionTimerDL_PTM_r17 *int64
	Drx_LongCycleStartOffsetPTM_r17   struct {
		Choice  int
		Ms10    *int64
		Ms20    *int64
		Ms32    *int64
		Ms40    *int64
		Ms60    *int64
		Ms64    *int64
		Ms70    *int64
		Ms80    *int64
		Ms128   *int64
		Ms160   *int64
		Ms256   *int64
		Ms320   *int64
		Ms512   *int64
		Ms640   *int64
		Ms1024  *int64
		Ms1280  *int64
		Ms2048  *int64
		Ms2560  *int64
		Ms5120  *int64
		Ms10240 *int64
	}
	Drx_SlotOffsetPTM_r17 int64
}

func (ie *DRX_ConfigPTM_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(dRXConfigPTMR17Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Drx_HARQ_RTT_TimerDL_PTM_r17 != nil, ie.Drx_RetransmissionTimerDL_PTM_r17 != nil}); err != nil {
		return err
	}

	// 2. drx-onDurationTimerPTM-r17: choice
	{
		choiceEnc := e.NewChoiceEncoder(dRX_ConfigPTM_r17DrxOnDurationTimerPTMR17Constraints)
		if err := choiceEnc.EncodeChoice(int64(ie.Drx_OnDurationTimerPTM_r17.Choice), false, nil); err != nil {
			return err
		}
		switch ie.Drx_OnDurationTimerPTM_r17.Choice {
		case DRX_ConfigPTM_r17_Drx_OnDurationTimerPTM_r17_SubMilliSeconds:
			if err := e.EncodeInteger((*ie.Drx_OnDurationTimerPTM_r17.SubMilliSeconds), per.Constrained(1, 31)); err != nil {
				return err
			}
		case DRX_ConfigPTM_r17_Drx_OnDurationTimerPTM_r17_MilliSeconds:
			if err := e.EncodeEnumerated((*ie.Drx_OnDurationTimerPTM_r17.MilliSeconds), dRXConfigPTMR17DrxOnDurationTimerPTMR17MilliSecondsConstraints); err != nil {
				return err
			}
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.Drx_OnDurationTimerPTM_r17.Choice), Constraint: "index out of range"}
		}
	}

	// 3. drx-InactivityTimerPTM-r17: enumerated
	{
		if err := e.EncodeEnumerated(ie.Drx_InactivityTimerPTM_r17, dRXConfigPTMR17DrxInactivityTimerPTMR17Constraints); err != nil {
			return err
		}
	}

	// 4. drx-HARQ-RTT-TimerDL-PTM-r17: integer
	{
		if ie.Drx_HARQ_RTT_TimerDL_PTM_r17 != nil {
			if err := e.EncodeInteger(*ie.Drx_HARQ_RTT_TimerDL_PTM_r17, dRXConfigPTMR17DrxHARQRTTTimerDLPTMR17Constraints); err != nil {
				return err
			}
		}
	}

	// 5. drx-RetransmissionTimerDL-PTM-r17: enumerated
	{
		if ie.Drx_RetransmissionTimerDL_PTM_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Drx_RetransmissionTimerDL_PTM_r17, dRXConfigPTMR17DrxRetransmissionTimerDLPTMR17Constraints); err != nil {
				return err
			}
		}
	}

	// 6. drx-LongCycleStartOffsetPTM-r17: choice
	{
		choiceEnc := e.NewChoiceEncoder(dRX_ConfigPTM_r17DrxLongCycleStartOffsetPTMR17Constraints)
		if err := choiceEnc.EncodeChoice(int64(ie.Drx_LongCycleStartOffsetPTM_r17.Choice), false, nil); err != nil {
			return err
		}
		switch ie.Drx_LongCycleStartOffsetPTM_r17.Choice {
		case DRX_ConfigPTM_r17_Drx_LongCycleStartOffsetPTM_r17_Ms10:
			if err := e.EncodeInteger((*ie.Drx_LongCycleStartOffsetPTM_r17.Ms10), per.Constrained(0, 9)); err != nil {
				return err
			}
		case DRX_ConfigPTM_r17_Drx_LongCycleStartOffsetPTM_r17_Ms20:
			if err := e.EncodeInteger((*ie.Drx_LongCycleStartOffsetPTM_r17.Ms20), per.Constrained(0, 19)); err != nil {
				return err
			}
		case DRX_ConfigPTM_r17_Drx_LongCycleStartOffsetPTM_r17_Ms32:
			if err := e.EncodeInteger((*ie.Drx_LongCycleStartOffsetPTM_r17.Ms32), per.Constrained(0, 31)); err != nil {
				return err
			}
		case DRX_ConfigPTM_r17_Drx_LongCycleStartOffsetPTM_r17_Ms40:
			if err := e.EncodeInteger((*ie.Drx_LongCycleStartOffsetPTM_r17.Ms40), per.Constrained(0, 39)); err != nil {
				return err
			}
		case DRX_ConfigPTM_r17_Drx_LongCycleStartOffsetPTM_r17_Ms60:
			if err := e.EncodeInteger((*ie.Drx_LongCycleStartOffsetPTM_r17.Ms60), per.Constrained(0, 59)); err != nil {
				return err
			}
		case DRX_ConfigPTM_r17_Drx_LongCycleStartOffsetPTM_r17_Ms64:
			if err := e.EncodeInteger((*ie.Drx_LongCycleStartOffsetPTM_r17.Ms64), per.Constrained(0, 63)); err != nil {
				return err
			}
		case DRX_ConfigPTM_r17_Drx_LongCycleStartOffsetPTM_r17_Ms70:
			if err := e.EncodeInteger((*ie.Drx_LongCycleStartOffsetPTM_r17.Ms70), per.Constrained(0, 69)); err != nil {
				return err
			}
		case DRX_ConfigPTM_r17_Drx_LongCycleStartOffsetPTM_r17_Ms80:
			if err := e.EncodeInteger((*ie.Drx_LongCycleStartOffsetPTM_r17.Ms80), per.Constrained(0, 79)); err != nil {
				return err
			}
		case DRX_ConfigPTM_r17_Drx_LongCycleStartOffsetPTM_r17_Ms128:
			if err := e.EncodeInteger((*ie.Drx_LongCycleStartOffsetPTM_r17.Ms128), per.Constrained(0, 127)); err != nil {
				return err
			}
		case DRX_ConfigPTM_r17_Drx_LongCycleStartOffsetPTM_r17_Ms160:
			if err := e.EncodeInteger((*ie.Drx_LongCycleStartOffsetPTM_r17.Ms160), per.Constrained(0, 159)); err != nil {
				return err
			}
		case DRX_ConfigPTM_r17_Drx_LongCycleStartOffsetPTM_r17_Ms256:
			if err := e.EncodeInteger((*ie.Drx_LongCycleStartOffsetPTM_r17.Ms256), per.Constrained(0, 255)); err != nil {
				return err
			}
		case DRX_ConfigPTM_r17_Drx_LongCycleStartOffsetPTM_r17_Ms320:
			if err := e.EncodeInteger((*ie.Drx_LongCycleStartOffsetPTM_r17.Ms320), per.Constrained(0, 319)); err != nil {
				return err
			}
		case DRX_ConfigPTM_r17_Drx_LongCycleStartOffsetPTM_r17_Ms512:
			if err := e.EncodeInteger((*ie.Drx_LongCycleStartOffsetPTM_r17.Ms512), per.Constrained(0, 511)); err != nil {
				return err
			}
		case DRX_ConfigPTM_r17_Drx_LongCycleStartOffsetPTM_r17_Ms640:
			if err := e.EncodeInteger((*ie.Drx_LongCycleStartOffsetPTM_r17.Ms640), per.Constrained(0, 639)); err != nil {
				return err
			}
		case DRX_ConfigPTM_r17_Drx_LongCycleStartOffsetPTM_r17_Ms1024:
			if err := e.EncodeInteger((*ie.Drx_LongCycleStartOffsetPTM_r17.Ms1024), per.Constrained(0, 1023)); err != nil {
				return err
			}
		case DRX_ConfigPTM_r17_Drx_LongCycleStartOffsetPTM_r17_Ms1280:
			if err := e.EncodeInteger((*ie.Drx_LongCycleStartOffsetPTM_r17.Ms1280), per.Constrained(0, 1279)); err != nil {
				return err
			}
		case DRX_ConfigPTM_r17_Drx_LongCycleStartOffsetPTM_r17_Ms2048:
			if err := e.EncodeInteger((*ie.Drx_LongCycleStartOffsetPTM_r17.Ms2048), per.Constrained(0, 2047)); err != nil {
				return err
			}
		case DRX_ConfigPTM_r17_Drx_LongCycleStartOffsetPTM_r17_Ms2560:
			if err := e.EncodeInteger((*ie.Drx_LongCycleStartOffsetPTM_r17.Ms2560), per.Constrained(0, 2559)); err != nil {
				return err
			}
		case DRX_ConfigPTM_r17_Drx_LongCycleStartOffsetPTM_r17_Ms5120:
			if err := e.EncodeInteger((*ie.Drx_LongCycleStartOffsetPTM_r17.Ms5120), per.Constrained(0, 5119)); err != nil {
				return err
			}
		case DRX_ConfigPTM_r17_Drx_LongCycleStartOffsetPTM_r17_Ms10240:
			if err := e.EncodeInteger((*ie.Drx_LongCycleStartOffsetPTM_r17.Ms10240), per.Constrained(0, 10239)); err != nil {
				return err
			}
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.Drx_LongCycleStartOffsetPTM_r17.Choice), Constraint: "index out of range"}
		}
	}

	// 7. drx-SlotOffsetPTM-r17: integer
	{
		if err := e.EncodeInteger(ie.Drx_SlotOffsetPTM_r17, dRXConfigPTMR17DrxSlotOffsetPTMR17Constraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *DRX_ConfigPTM_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(dRXConfigPTMR17Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. drx-onDurationTimerPTM-r17: choice
	{
		choiceDec := d.NewChoiceDecoder(dRX_ConfigPTM_r17DrxOnDurationTimerPTMR17Constraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.Drx_OnDurationTimerPTM_r17.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case DRX_ConfigPTM_r17_Drx_OnDurationTimerPTM_r17_SubMilliSeconds:
			ie.Drx_OnDurationTimerPTM_r17.SubMilliSeconds = new(int64)
			v, err := d.DecodeInteger(per.Constrained(1, 31))
			if err != nil {
				return err
			}
			(*ie.Drx_OnDurationTimerPTM_r17.SubMilliSeconds) = v
		case DRX_ConfigPTM_r17_Drx_OnDurationTimerPTM_r17_MilliSeconds:
			ie.Drx_OnDurationTimerPTM_r17.MilliSeconds = new(int64)
			v, err := d.DecodeEnumerated(dRXConfigPTMR17DrxOnDurationTimerPTMR17MilliSecondsConstraints)
			if err != nil {
				return err
			}
			(*ie.Drx_OnDurationTimerPTM_r17.MilliSeconds) = v
		}
	}

	// 3. drx-InactivityTimerPTM-r17: enumerated
	{
		v1, err := d.DecodeEnumerated(dRXConfigPTMR17DrxInactivityTimerPTMR17Constraints)
		if err != nil {
			return err
		}
		ie.Drx_InactivityTimerPTM_r17 = v1
	}

	// 4. drx-HARQ-RTT-TimerDL-PTM-r17: integer
	{
		if seq.IsComponentPresent(2) {
			v, err := d.DecodeInteger(dRXConfigPTMR17DrxHARQRTTTimerDLPTMR17Constraints)
			if err != nil {
				return err
			}
			ie.Drx_HARQ_RTT_TimerDL_PTM_r17 = &v
		}
	}

	// 5. drx-RetransmissionTimerDL-PTM-r17: enumerated
	{
		if seq.IsComponentPresent(3) {
			idx, err := d.DecodeEnumerated(dRXConfigPTMR17DrxRetransmissionTimerDLPTMR17Constraints)
			if err != nil {
				return err
			}
			ie.Drx_RetransmissionTimerDL_PTM_r17 = &idx
		}
	}

	// 6. drx-LongCycleStartOffsetPTM-r17: choice
	{
		choiceDec := d.NewChoiceDecoder(dRX_ConfigPTM_r17DrxLongCycleStartOffsetPTMR17Constraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.Drx_LongCycleStartOffsetPTM_r17.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case DRX_ConfigPTM_r17_Drx_LongCycleStartOffsetPTM_r17_Ms10:
			ie.Drx_LongCycleStartOffsetPTM_r17.Ms10 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 9))
			if err != nil {
				return err
			}
			(*ie.Drx_LongCycleStartOffsetPTM_r17.Ms10) = v
		case DRX_ConfigPTM_r17_Drx_LongCycleStartOffsetPTM_r17_Ms20:
			ie.Drx_LongCycleStartOffsetPTM_r17.Ms20 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 19))
			if err != nil {
				return err
			}
			(*ie.Drx_LongCycleStartOffsetPTM_r17.Ms20) = v
		case DRX_ConfigPTM_r17_Drx_LongCycleStartOffsetPTM_r17_Ms32:
			ie.Drx_LongCycleStartOffsetPTM_r17.Ms32 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 31))
			if err != nil {
				return err
			}
			(*ie.Drx_LongCycleStartOffsetPTM_r17.Ms32) = v
		case DRX_ConfigPTM_r17_Drx_LongCycleStartOffsetPTM_r17_Ms40:
			ie.Drx_LongCycleStartOffsetPTM_r17.Ms40 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 39))
			if err != nil {
				return err
			}
			(*ie.Drx_LongCycleStartOffsetPTM_r17.Ms40) = v
		case DRX_ConfigPTM_r17_Drx_LongCycleStartOffsetPTM_r17_Ms60:
			ie.Drx_LongCycleStartOffsetPTM_r17.Ms60 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 59))
			if err != nil {
				return err
			}
			(*ie.Drx_LongCycleStartOffsetPTM_r17.Ms60) = v
		case DRX_ConfigPTM_r17_Drx_LongCycleStartOffsetPTM_r17_Ms64:
			ie.Drx_LongCycleStartOffsetPTM_r17.Ms64 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 63))
			if err != nil {
				return err
			}
			(*ie.Drx_LongCycleStartOffsetPTM_r17.Ms64) = v
		case DRX_ConfigPTM_r17_Drx_LongCycleStartOffsetPTM_r17_Ms70:
			ie.Drx_LongCycleStartOffsetPTM_r17.Ms70 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 69))
			if err != nil {
				return err
			}
			(*ie.Drx_LongCycleStartOffsetPTM_r17.Ms70) = v
		case DRX_ConfigPTM_r17_Drx_LongCycleStartOffsetPTM_r17_Ms80:
			ie.Drx_LongCycleStartOffsetPTM_r17.Ms80 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 79))
			if err != nil {
				return err
			}
			(*ie.Drx_LongCycleStartOffsetPTM_r17.Ms80) = v
		case DRX_ConfigPTM_r17_Drx_LongCycleStartOffsetPTM_r17_Ms128:
			ie.Drx_LongCycleStartOffsetPTM_r17.Ms128 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 127))
			if err != nil {
				return err
			}
			(*ie.Drx_LongCycleStartOffsetPTM_r17.Ms128) = v
		case DRX_ConfigPTM_r17_Drx_LongCycleStartOffsetPTM_r17_Ms160:
			ie.Drx_LongCycleStartOffsetPTM_r17.Ms160 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 159))
			if err != nil {
				return err
			}
			(*ie.Drx_LongCycleStartOffsetPTM_r17.Ms160) = v
		case DRX_ConfigPTM_r17_Drx_LongCycleStartOffsetPTM_r17_Ms256:
			ie.Drx_LongCycleStartOffsetPTM_r17.Ms256 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 255))
			if err != nil {
				return err
			}
			(*ie.Drx_LongCycleStartOffsetPTM_r17.Ms256) = v
		case DRX_ConfigPTM_r17_Drx_LongCycleStartOffsetPTM_r17_Ms320:
			ie.Drx_LongCycleStartOffsetPTM_r17.Ms320 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 319))
			if err != nil {
				return err
			}
			(*ie.Drx_LongCycleStartOffsetPTM_r17.Ms320) = v
		case DRX_ConfigPTM_r17_Drx_LongCycleStartOffsetPTM_r17_Ms512:
			ie.Drx_LongCycleStartOffsetPTM_r17.Ms512 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 511))
			if err != nil {
				return err
			}
			(*ie.Drx_LongCycleStartOffsetPTM_r17.Ms512) = v
		case DRX_ConfigPTM_r17_Drx_LongCycleStartOffsetPTM_r17_Ms640:
			ie.Drx_LongCycleStartOffsetPTM_r17.Ms640 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 639))
			if err != nil {
				return err
			}
			(*ie.Drx_LongCycleStartOffsetPTM_r17.Ms640) = v
		case DRX_ConfigPTM_r17_Drx_LongCycleStartOffsetPTM_r17_Ms1024:
			ie.Drx_LongCycleStartOffsetPTM_r17.Ms1024 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 1023))
			if err != nil {
				return err
			}
			(*ie.Drx_LongCycleStartOffsetPTM_r17.Ms1024) = v
		case DRX_ConfigPTM_r17_Drx_LongCycleStartOffsetPTM_r17_Ms1280:
			ie.Drx_LongCycleStartOffsetPTM_r17.Ms1280 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 1279))
			if err != nil {
				return err
			}
			(*ie.Drx_LongCycleStartOffsetPTM_r17.Ms1280) = v
		case DRX_ConfigPTM_r17_Drx_LongCycleStartOffsetPTM_r17_Ms2048:
			ie.Drx_LongCycleStartOffsetPTM_r17.Ms2048 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 2047))
			if err != nil {
				return err
			}
			(*ie.Drx_LongCycleStartOffsetPTM_r17.Ms2048) = v
		case DRX_ConfigPTM_r17_Drx_LongCycleStartOffsetPTM_r17_Ms2560:
			ie.Drx_LongCycleStartOffsetPTM_r17.Ms2560 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 2559))
			if err != nil {
				return err
			}
			(*ie.Drx_LongCycleStartOffsetPTM_r17.Ms2560) = v
		case DRX_ConfigPTM_r17_Drx_LongCycleStartOffsetPTM_r17_Ms5120:
			ie.Drx_LongCycleStartOffsetPTM_r17.Ms5120 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 5119))
			if err != nil {
				return err
			}
			(*ie.Drx_LongCycleStartOffsetPTM_r17.Ms5120) = v
		case DRX_ConfigPTM_r17_Drx_LongCycleStartOffsetPTM_r17_Ms10240:
			ie.Drx_LongCycleStartOffsetPTM_r17.Ms10240 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 10239))
			if err != nil {
				return err
			}
			(*ie.Drx_LongCycleStartOffsetPTM_r17.Ms10240) = v
		}
	}

	// 7. drx-SlotOffsetPTM-r17: integer
	{
		v5, err := d.DecodeInteger(dRXConfigPTMR17DrxSlotOffsetPTMR17Constraints)
		if err != nil {
			return err
		}
		ie.Drx_SlotOffsetPTM_r17 = v5
	}

	return nil
}
