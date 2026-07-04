// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: DRX-Config (line 8110).

var dRXConfigConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "drx-onDurationTimer"},
		{Name: "drx-InactivityTimer"},
		{Name: "drx-HARQ-RTT-TimerDL"},
		{Name: "drx-HARQ-RTT-TimerUL"},
		{Name: "drx-RetransmissionTimerDL"},
		{Name: "drx-RetransmissionTimerUL"},
		{Name: "drx-LongCycleStartOffset"},
		{Name: "shortDRX", Optional: true},
		{Name: "drx-SlotOffset"},
	},
}

var dRX_ConfigDrxOnDurationTimerConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "subMilliSeconds"},
		{Name: "milliSeconds"},
	},
}

const (
	DRX_Config_Drx_OnDurationTimer_SubMilliSeconds = 0
	DRX_Config_Drx_OnDurationTimer_MilliSeconds    = 1
)

const (
	DRX_Config_Drx_InactivityTimer_Ms0    = 0
	DRX_Config_Drx_InactivityTimer_Ms1    = 1
	DRX_Config_Drx_InactivityTimer_Ms2    = 2
	DRX_Config_Drx_InactivityTimer_Ms3    = 3
	DRX_Config_Drx_InactivityTimer_Ms4    = 4
	DRX_Config_Drx_InactivityTimer_Ms5    = 5
	DRX_Config_Drx_InactivityTimer_Ms6    = 6
	DRX_Config_Drx_InactivityTimer_Ms8    = 7
	DRX_Config_Drx_InactivityTimer_Ms10   = 8
	DRX_Config_Drx_InactivityTimer_Ms20   = 9
	DRX_Config_Drx_InactivityTimer_Ms30   = 10
	DRX_Config_Drx_InactivityTimer_Ms40   = 11
	DRX_Config_Drx_InactivityTimer_Ms50   = 12
	DRX_Config_Drx_InactivityTimer_Ms60   = 13
	DRX_Config_Drx_InactivityTimer_Ms80   = 14
	DRX_Config_Drx_InactivityTimer_Ms100  = 15
	DRX_Config_Drx_InactivityTimer_Ms200  = 16
	DRX_Config_Drx_InactivityTimer_Ms300  = 17
	DRX_Config_Drx_InactivityTimer_Ms500  = 18
	DRX_Config_Drx_InactivityTimer_Ms750  = 19
	DRX_Config_Drx_InactivityTimer_Ms1280 = 20
	DRX_Config_Drx_InactivityTimer_Ms1920 = 21
	DRX_Config_Drx_InactivityTimer_Ms2560 = 22
	DRX_Config_Drx_InactivityTimer_Spare9 = 23
	DRX_Config_Drx_InactivityTimer_Spare8 = 24
	DRX_Config_Drx_InactivityTimer_Spare7 = 25
	DRX_Config_Drx_InactivityTimer_Spare6 = 26
	DRX_Config_Drx_InactivityTimer_Spare5 = 27
	DRX_Config_Drx_InactivityTimer_Spare4 = 28
	DRX_Config_Drx_InactivityTimer_Spare3 = 29
	DRX_Config_Drx_InactivityTimer_Spare2 = 30
	DRX_Config_Drx_InactivityTimer_Spare1 = 31
)

var dRXConfigDrxInactivityTimerConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{DRX_Config_Drx_InactivityTimer_Ms0, DRX_Config_Drx_InactivityTimer_Ms1, DRX_Config_Drx_InactivityTimer_Ms2, DRX_Config_Drx_InactivityTimer_Ms3, DRX_Config_Drx_InactivityTimer_Ms4, DRX_Config_Drx_InactivityTimer_Ms5, DRX_Config_Drx_InactivityTimer_Ms6, DRX_Config_Drx_InactivityTimer_Ms8, DRX_Config_Drx_InactivityTimer_Ms10, DRX_Config_Drx_InactivityTimer_Ms20, DRX_Config_Drx_InactivityTimer_Ms30, DRX_Config_Drx_InactivityTimer_Ms40, DRX_Config_Drx_InactivityTimer_Ms50, DRX_Config_Drx_InactivityTimer_Ms60, DRX_Config_Drx_InactivityTimer_Ms80, DRX_Config_Drx_InactivityTimer_Ms100, DRX_Config_Drx_InactivityTimer_Ms200, DRX_Config_Drx_InactivityTimer_Ms300, DRX_Config_Drx_InactivityTimer_Ms500, DRX_Config_Drx_InactivityTimer_Ms750, DRX_Config_Drx_InactivityTimer_Ms1280, DRX_Config_Drx_InactivityTimer_Ms1920, DRX_Config_Drx_InactivityTimer_Ms2560, DRX_Config_Drx_InactivityTimer_Spare9, DRX_Config_Drx_InactivityTimer_Spare8, DRX_Config_Drx_InactivityTimer_Spare7, DRX_Config_Drx_InactivityTimer_Spare6, DRX_Config_Drx_InactivityTimer_Spare5, DRX_Config_Drx_InactivityTimer_Spare4, DRX_Config_Drx_InactivityTimer_Spare3, DRX_Config_Drx_InactivityTimer_Spare2, DRX_Config_Drx_InactivityTimer_Spare1},
}

var dRXConfigDrxHARQRTTTimerDLConstraints = per.Constrained(0, 56)

var dRXConfigDrxHARQRTTTimerULConstraints = per.Constrained(0, 56)

const (
	DRX_Config_Drx_RetransmissionTimerDL_Sl0     = 0
	DRX_Config_Drx_RetransmissionTimerDL_Sl1     = 1
	DRX_Config_Drx_RetransmissionTimerDL_Sl2     = 2
	DRX_Config_Drx_RetransmissionTimerDL_Sl4     = 3
	DRX_Config_Drx_RetransmissionTimerDL_Sl6     = 4
	DRX_Config_Drx_RetransmissionTimerDL_Sl8     = 5
	DRX_Config_Drx_RetransmissionTimerDL_Sl16    = 6
	DRX_Config_Drx_RetransmissionTimerDL_Sl24    = 7
	DRX_Config_Drx_RetransmissionTimerDL_Sl33    = 8
	DRX_Config_Drx_RetransmissionTimerDL_Sl40    = 9
	DRX_Config_Drx_RetransmissionTimerDL_Sl64    = 10
	DRX_Config_Drx_RetransmissionTimerDL_Sl80    = 11
	DRX_Config_Drx_RetransmissionTimerDL_Sl96    = 12
	DRX_Config_Drx_RetransmissionTimerDL_Sl112   = 13
	DRX_Config_Drx_RetransmissionTimerDL_Sl128   = 14
	DRX_Config_Drx_RetransmissionTimerDL_Sl160   = 15
	DRX_Config_Drx_RetransmissionTimerDL_Sl320   = 16
	DRX_Config_Drx_RetransmissionTimerDL_Spare15 = 17
	DRX_Config_Drx_RetransmissionTimerDL_Spare14 = 18
	DRX_Config_Drx_RetransmissionTimerDL_Spare13 = 19
	DRX_Config_Drx_RetransmissionTimerDL_Spare12 = 20
	DRX_Config_Drx_RetransmissionTimerDL_Spare11 = 21
	DRX_Config_Drx_RetransmissionTimerDL_Spare10 = 22
	DRX_Config_Drx_RetransmissionTimerDL_Spare9  = 23
	DRX_Config_Drx_RetransmissionTimerDL_Spare8  = 24
	DRX_Config_Drx_RetransmissionTimerDL_Spare7  = 25
	DRX_Config_Drx_RetransmissionTimerDL_Spare6  = 26
	DRX_Config_Drx_RetransmissionTimerDL_Spare5  = 27
	DRX_Config_Drx_RetransmissionTimerDL_Spare4  = 28
	DRX_Config_Drx_RetransmissionTimerDL_Spare3  = 29
	DRX_Config_Drx_RetransmissionTimerDL_Spare2  = 30
	DRX_Config_Drx_RetransmissionTimerDL_Spare1  = 31
)

var dRXConfigDrxRetransmissionTimerDLConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{DRX_Config_Drx_RetransmissionTimerDL_Sl0, DRX_Config_Drx_RetransmissionTimerDL_Sl1, DRX_Config_Drx_RetransmissionTimerDL_Sl2, DRX_Config_Drx_RetransmissionTimerDL_Sl4, DRX_Config_Drx_RetransmissionTimerDL_Sl6, DRX_Config_Drx_RetransmissionTimerDL_Sl8, DRX_Config_Drx_RetransmissionTimerDL_Sl16, DRX_Config_Drx_RetransmissionTimerDL_Sl24, DRX_Config_Drx_RetransmissionTimerDL_Sl33, DRX_Config_Drx_RetransmissionTimerDL_Sl40, DRX_Config_Drx_RetransmissionTimerDL_Sl64, DRX_Config_Drx_RetransmissionTimerDL_Sl80, DRX_Config_Drx_RetransmissionTimerDL_Sl96, DRX_Config_Drx_RetransmissionTimerDL_Sl112, DRX_Config_Drx_RetransmissionTimerDL_Sl128, DRX_Config_Drx_RetransmissionTimerDL_Sl160, DRX_Config_Drx_RetransmissionTimerDL_Sl320, DRX_Config_Drx_RetransmissionTimerDL_Spare15, DRX_Config_Drx_RetransmissionTimerDL_Spare14, DRX_Config_Drx_RetransmissionTimerDL_Spare13, DRX_Config_Drx_RetransmissionTimerDL_Spare12, DRX_Config_Drx_RetransmissionTimerDL_Spare11, DRX_Config_Drx_RetransmissionTimerDL_Spare10, DRX_Config_Drx_RetransmissionTimerDL_Spare9, DRX_Config_Drx_RetransmissionTimerDL_Spare8, DRX_Config_Drx_RetransmissionTimerDL_Spare7, DRX_Config_Drx_RetransmissionTimerDL_Spare6, DRX_Config_Drx_RetransmissionTimerDL_Spare5, DRX_Config_Drx_RetransmissionTimerDL_Spare4, DRX_Config_Drx_RetransmissionTimerDL_Spare3, DRX_Config_Drx_RetransmissionTimerDL_Spare2, DRX_Config_Drx_RetransmissionTimerDL_Spare1},
}

const (
	DRX_Config_Drx_RetransmissionTimerUL_Sl0     = 0
	DRX_Config_Drx_RetransmissionTimerUL_Sl1     = 1
	DRX_Config_Drx_RetransmissionTimerUL_Sl2     = 2
	DRX_Config_Drx_RetransmissionTimerUL_Sl4     = 3
	DRX_Config_Drx_RetransmissionTimerUL_Sl6     = 4
	DRX_Config_Drx_RetransmissionTimerUL_Sl8     = 5
	DRX_Config_Drx_RetransmissionTimerUL_Sl16    = 6
	DRX_Config_Drx_RetransmissionTimerUL_Sl24    = 7
	DRX_Config_Drx_RetransmissionTimerUL_Sl33    = 8
	DRX_Config_Drx_RetransmissionTimerUL_Sl40    = 9
	DRX_Config_Drx_RetransmissionTimerUL_Sl64    = 10
	DRX_Config_Drx_RetransmissionTimerUL_Sl80    = 11
	DRX_Config_Drx_RetransmissionTimerUL_Sl96    = 12
	DRX_Config_Drx_RetransmissionTimerUL_Sl112   = 13
	DRX_Config_Drx_RetransmissionTimerUL_Sl128   = 14
	DRX_Config_Drx_RetransmissionTimerUL_Sl160   = 15
	DRX_Config_Drx_RetransmissionTimerUL_Sl320   = 16
	DRX_Config_Drx_RetransmissionTimerUL_Spare15 = 17
	DRX_Config_Drx_RetransmissionTimerUL_Spare14 = 18
	DRX_Config_Drx_RetransmissionTimerUL_Spare13 = 19
	DRX_Config_Drx_RetransmissionTimerUL_Spare12 = 20
	DRX_Config_Drx_RetransmissionTimerUL_Spare11 = 21
	DRX_Config_Drx_RetransmissionTimerUL_Spare10 = 22
	DRX_Config_Drx_RetransmissionTimerUL_Spare9  = 23
	DRX_Config_Drx_RetransmissionTimerUL_Spare8  = 24
	DRX_Config_Drx_RetransmissionTimerUL_Spare7  = 25
	DRX_Config_Drx_RetransmissionTimerUL_Spare6  = 26
	DRX_Config_Drx_RetransmissionTimerUL_Spare5  = 27
	DRX_Config_Drx_RetransmissionTimerUL_Spare4  = 28
	DRX_Config_Drx_RetransmissionTimerUL_Spare3  = 29
	DRX_Config_Drx_RetransmissionTimerUL_Spare2  = 30
	DRX_Config_Drx_RetransmissionTimerUL_Spare1  = 31
)

var dRXConfigDrxRetransmissionTimerULConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{DRX_Config_Drx_RetransmissionTimerUL_Sl0, DRX_Config_Drx_RetransmissionTimerUL_Sl1, DRX_Config_Drx_RetransmissionTimerUL_Sl2, DRX_Config_Drx_RetransmissionTimerUL_Sl4, DRX_Config_Drx_RetransmissionTimerUL_Sl6, DRX_Config_Drx_RetransmissionTimerUL_Sl8, DRX_Config_Drx_RetransmissionTimerUL_Sl16, DRX_Config_Drx_RetransmissionTimerUL_Sl24, DRX_Config_Drx_RetransmissionTimerUL_Sl33, DRX_Config_Drx_RetransmissionTimerUL_Sl40, DRX_Config_Drx_RetransmissionTimerUL_Sl64, DRX_Config_Drx_RetransmissionTimerUL_Sl80, DRX_Config_Drx_RetransmissionTimerUL_Sl96, DRX_Config_Drx_RetransmissionTimerUL_Sl112, DRX_Config_Drx_RetransmissionTimerUL_Sl128, DRX_Config_Drx_RetransmissionTimerUL_Sl160, DRX_Config_Drx_RetransmissionTimerUL_Sl320, DRX_Config_Drx_RetransmissionTimerUL_Spare15, DRX_Config_Drx_RetransmissionTimerUL_Spare14, DRX_Config_Drx_RetransmissionTimerUL_Spare13, DRX_Config_Drx_RetransmissionTimerUL_Spare12, DRX_Config_Drx_RetransmissionTimerUL_Spare11, DRX_Config_Drx_RetransmissionTimerUL_Spare10, DRX_Config_Drx_RetransmissionTimerUL_Spare9, DRX_Config_Drx_RetransmissionTimerUL_Spare8, DRX_Config_Drx_RetransmissionTimerUL_Spare7, DRX_Config_Drx_RetransmissionTimerUL_Spare6, DRX_Config_Drx_RetransmissionTimerUL_Spare5, DRX_Config_Drx_RetransmissionTimerUL_Spare4, DRX_Config_Drx_RetransmissionTimerUL_Spare3, DRX_Config_Drx_RetransmissionTimerUL_Spare2, DRX_Config_Drx_RetransmissionTimerUL_Spare1},
}

var dRX_ConfigDrxLongCycleStartOffsetConstraints = per.ChoiceConstraints{
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
	DRX_Config_Drx_LongCycleStartOffset_Ms10    = 0
	DRX_Config_Drx_LongCycleStartOffset_Ms20    = 1
	DRX_Config_Drx_LongCycleStartOffset_Ms32    = 2
	DRX_Config_Drx_LongCycleStartOffset_Ms40    = 3
	DRX_Config_Drx_LongCycleStartOffset_Ms60    = 4
	DRX_Config_Drx_LongCycleStartOffset_Ms64    = 5
	DRX_Config_Drx_LongCycleStartOffset_Ms70    = 6
	DRX_Config_Drx_LongCycleStartOffset_Ms80    = 7
	DRX_Config_Drx_LongCycleStartOffset_Ms128   = 8
	DRX_Config_Drx_LongCycleStartOffset_Ms160   = 9
	DRX_Config_Drx_LongCycleStartOffset_Ms256   = 10
	DRX_Config_Drx_LongCycleStartOffset_Ms320   = 11
	DRX_Config_Drx_LongCycleStartOffset_Ms512   = 12
	DRX_Config_Drx_LongCycleStartOffset_Ms640   = 13
	DRX_Config_Drx_LongCycleStartOffset_Ms1024  = 14
	DRX_Config_Drx_LongCycleStartOffset_Ms1280  = 15
	DRX_Config_Drx_LongCycleStartOffset_Ms2048  = 16
	DRX_Config_Drx_LongCycleStartOffset_Ms2560  = 17
	DRX_Config_Drx_LongCycleStartOffset_Ms5120  = 18
	DRX_Config_Drx_LongCycleStartOffset_Ms10240 = 19
)

var dRXConfigDrxSlotOffsetConstraints = per.Constrained(0, 31)

const (
	DRX_Config_Drx_OnDurationTimer_MilliSeconds_Ms1    = 0
	DRX_Config_Drx_OnDurationTimer_MilliSeconds_Ms2    = 1
	DRX_Config_Drx_OnDurationTimer_MilliSeconds_Ms3    = 2
	DRX_Config_Drx_OnDurationTimer_MilliSeconds_Ms4    = 3
	DRX_Config_Drx_OnDurationTimer_MilliSeconds_Ms5    = 4
	DRX_Config_Drx_OnDurationTimer_MilliSeconds_Ms6    = 5
	DRX_Config_Drx_OnDurationTimer_MilliSeconds_Ms8    = 6
	DRX_Config_Drx_OnDurationTimer_MilliSeconds_Ms10   = 7
	DRX_Config_Drx_OnDurationTimer_MilliSeconds_Ms20   = 8
	DRX_Config_Drx_OnDurationTimer_MilliSeconds_Ms30   = 9
	DRX_Config_Drx_OnDurationTimer_MilliSeconds_Ms40   = 10
	DRX_Config_Drx_OnDurationTimer_MilliSeconds_Ms50   = 11
	DRX_Config_Drx_OnDurationTimer_MilliSeconds_Ms60   = 12
	DRX_Config_Drx_OnDurationTimer_MilliSeconds_Ms80   = 13
	DRX_Config_Drx_OnDurationTimer_MilliSeconds_Ms100  = 14
	DRX_Config_Drx_OnDurationTimer_MilliSeconds_Ms200  = 15
	DRX_Config_Drx_OnDurationTimer_MilliSeconds_Ms300  = 16
	DRX_Config_Drx_OnDurationTimer_MilliSeconds_Ms400  = 17
	DRX_Config_Drx_OnDurationTimer_MilliSeconds_Ms500  = 18
	DRX_Config_Drx_OnDurationTimer_MilliSeconds_Ms600  = 19
	DRX_Config_Drx_OnDurationTimer_MilliSeconds_Ms800  = 20
	DRX_Config_Drx_OnDurationTimer_MilliSeconds_Ms1000 = 21
	DRX_Config_Drx_OnDurationTimer_MilliSeconds_Ms1200 = 22
	DRX_Config_Drx_OnDurationTimer_MilliSeconds_Ms1600 = 23
	DRX_Config_Drx_OnDurationTimer_MilliSeconds_Spare8 = 24
	DRX_Config_Drx_OnDurationTimer_MilliSeconds_Spare7 = 25
	DRX_Config_Drx_OnDurationTimer_MilliSeconds_Spare6 = 26
	DRX_Config_Drx_OnDurationTimer_MilliSeconds_Spare5 = 27
	DRX_Config_Drx_OnDurationTimer_MilliSeconds_Spare4 = 28
	DRX_Config_Drx_OnDurationTimer_MilliSeconds_Spare3 = 29
	DRX_Config_Drx_OnDurationTimer_MilliSeconds_Spare2 = 30
	DRX_Config_Drx_OnDurationTimer_MilliSeconds_Spare1 = 31
)

var dRXConfigDrxOnDurationTimerMilliSecondsConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{DRX_Config_Drx_OnDurationTimer_MilliSeconds_Ms1, DRX_Config_Drx_OnDurationTimer_MilliSeconds_Ms2, DRX_Config_Drx_OnDurationTimer_MilliSeconds_Ms3, DRX_Config_Drx_OnDurationTimer_MilliSeconds_Ms4, DRX_Config_Drx_OnDurationTimer_MilliSeconds_Ms5, DRX_Config_Drx_OnDurationTimer_MilliSeconds_Ms6, DRX_Config_Drx_OnDurationTimer_MilliSeconds_Ms8, DRX_Config_Drx_OnDurationTimer_MilliSeconds_Ms10, DRX_Config_Drx_OnDurationTimer_MilliSeconds_Ms20, DRX_Config_Drx_OnDurationTimer_MilliSeconds_Ms30, DRX_Config_Drx_OnDurationTimer_MilliSeconds_Ms40, DRX_Config_Drx_OnDurationTimer_MilliSeconds_Ms50, DRX_Config_Drx_OnDurationTimer_MilliSeconds_Ms60, DRX_Config_Drx_OnDurationTimer_MilliSeconds_Ms80, DRX_Config_Drx_OnDurationTimer_MilliSeconds_Ms100, DRX_Config_Drx_OnDurationTimer_MilliSeconds_Ms200, DRX_Config_Drx_OnDurationTimer_MilliSeconds_Ms300, DRX_Config_Drx_OnDurationTimer_MilliSeconds_Ms400, DRX_Config_Drx_OnDurationTimer_MilliSeconds_Ms500, DRX_Config_Drx_OnDurationTimer_MilliSeconds_Ms600, DRX_Config_Drx_OnDurationTimer_MilliSeconds_Ms800, DRX_Config_Drx_OnDurationTimer_MilliSeconds_Ms1000, DRX_Config_Drx_OnDurationTimer_MilliSeconds_Ms1200, DRX_Config_Drx_OnDurationTimer_MilliSeconds_Ms1600, DRX_Config_Drx_OnDurationTimer_MilliSeconds_Spare8, DRX_Config_Drx_OnDurationTimer_MilliSeconds_Spare7, DRX_Config_Drx_OnDurationTimer_MilliSeconds_Spare6, DRX_Config_Drx_OnDurationTimer_MilliSeconds_Spare5, DRX_Config_Drx_OnDurationTimer_MilliSeconds_Spare4, DRX_Config_Drx_OnDurationTimer_MilliSeconds_Spare3, DRX_Config_Drx_OnDurationTimer_MilliSeconds_Spare2, DRX_Config_Drx_OnDurationTimer_MilliSeconds_Spare1},
}

const (
	DRX_Config_ShortDRX_Drx_ShortCycle_Ms2    = 0
	DRX_Config_ShortDRX_Drx_ShortCycle_Ms3    = 1
	DRX_Config_ShortDRX_Drx_ShortCycle_Ms4    = 2
	DRX_Config_ShortDRX_Drx_ShortCycle_Ms5    = 3
	DRX_Config_ShortDRX_Drx_ShortCycle_Ms6    = 4
	DRX_Config_ShortDRX_Drx_ShortCycle_Ms7    = 5
	DRX_Config_ShortDRX_Drx_ShortCycle_Ms8    = 6
	DRX_Config_ShortDRX_Drx_ShortCycle_Ms10   = 7
	DRX_Config_ShortDRX_Drx_ShortCycle_Ms14   = 8
	DRX_Config_ShortDRX_Drx_ShortCycle_Ms16   = 9
	DRX_Config_ShortDRX_Drx_ShortCycle_Ms20   = 10
	DRX_Config_ShortDRX_Drx_ShortCycle_Ms30   = 11
	DRX_Config_ShortDRX_Drx_ShortCycle_Ms32   = 12
	DRX_Config_ShortDRX_Drx_ShortCycle_Ms35   = 13
	DRX_Config_ShortDRX_Drx_ShortCycle_Ms40   = 14
	DRX_Config_ShortDRX_Drx_ShortCycle_Ms64   = 15
	DRX_Config_ShortDRX_Drx_ShortCycle_Ms80   = 16
	DRX_Config_ShortDRX_Drx_ShortCycle_Ms128  = 17
	DRX_Config_ShortDRX_Drx_ShortCycle_Ms160  = 18
	DRX_Config_ShortDRX_Drx_ShortCycle_Ms256  = 19
	DRX_Config_ShortDRX_Drx_ShortCycle_Ms320  = 20
	DRX_Config_ShortDRX_Drx_ShortCycle_Ms512  = 21
	DRX_Config_ShortDRX_Drx_ShortCycle_Ms640  = 22
	DRX_Config_ShortDRX_Drx_ShortCycle_Spare9 = 23
	DRX_Config_ShortDRX_Drx_ShortCycle_Spare8 = 24
	DRX_Config_ShortDRX_Drx_ShortCycle_Spare7 = 25
	DRX_Config_ShortDRX_Drx_ShortCycle_Spare6 = 26
	DRX_Config_ShortDRX_Drx_ShortCycle_Spare5 = 27
	DRX_Config_ShortDRX_Drx_ShortCycle_Spare4 = 28
	DRX_Config_ShortDRX_Drx_ShortCycle_Spare3 = 29
	DRX_Config_ShortDRX_Drx_ShortCycle_Spare2 = 30
	DRX_Config_ShortDRX_Drx_ShortCycle_Spare1 = 31
)

var dRXConfigShortDRXDrxShortCycleConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{DRX_Config_ShortDRX_Drx_ShortCycle_Ms2, DRX_Config_ShortDRX_Drx_ShortCycle_Ms3, DRX_Config_ShortDRX_Drx_ShortCycle_Ms4, DRX_Config_ShortDRX_Drx_ShortCycle_Ms5, DRX_Config_ShortDRX_Drx_ShortCycle_Ms6, DRX_Config_ShortDRX_Drx_ShortCycle_Ms7, DRX_Config_ShortDRX_Drx_ShortCycle_Ms8, DRX_Config_ShortDRX_Drx_ShortCycle_Ms10, DRX_Config_ShortDRX_Drx_ShortCycle_Ms14, DRX_Config_ShortDRX_Drx_ShortCycle_Ms16, DRX_Config_ShortDRX_Drx_ShortCycle_Ms20, DRX_Config_ShortDRX_Drx_ShortCycle_Ms30, DRX_Config_ShortDRX_Drx_ShortCycle_Ms32, DRX_Config_ShortDRX_Drx_ShortCycle_Ms35, DRX_Config_ShortDRX_Drx_ShortCycle_Ms40, DRX_Config_ShortDRX_Drx_ShortCycle_Ms64, DRX_Config_ShortDRX_Drx_ShortCycle_Ms80, DRX_Config_ShortDRX_Drx_ShortCycle_Ms128, DRX_Config_ShortDRX_Drx_ShortCycle_Ms160, DRX_Config_ShortDRX_Drx_ShortCycle_Ms256, DRX_Config_ShortDRX_Drx_ShortCycle_Ms320, DRX_Config_ShortDRX_Drx_ShortCycle_Ms512, DRX_Config_ShortDRX_Drx_ShortCycle_Ms640, DRX_Config_ShortDRX_Drx_ShortCycle_Spare9, DRX_Config_ShortDRX_Drx_ShortCycle_Spare8, DRX_Config_ShortDRX_Drx_ShortCycle_Spare7, DRX_Config_ShortDRX_Drx_ShortCycle_Spare6, DRX_Config_ShortDRX_Drx_ShortCycle_Spare5, DRX_Config_ShortDRX_Drx_ShortCycle_Spare4, DRX_Config_ShortDRX_Drx_ShortCycle_Spare3, DRX_Config_ShortDRX_Drx_ShortCycle_Spare2, DRX_Config_ShortDRX_Drx_ShortCycle_Spare1},
}

type DRX_Config struct {
	Drx_OnDurationTimer struct {
		Choice          int
		SubMilliSeconds *int64
		MilliSeconds    *int64
	}
	Drx_InactivityTimer       int64
	Drx_HARQ_RTT_TimerDL      int64
	Drx_HARQ_RTT_TimerUL      int64
	Drx_RetransmissionTimerDL int64
	Drx_RetransmissionTimerUL int64
	Drx_LongCycleStartOffset  struct {
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
	ShortDRX *struct {
		Drx_ShortCycle      int64
		Drx_ShortCycleTimer int64
	}
	Drx_SlotOffset int64
}

func (ie *DRX_Config) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(dRXConfigConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.ShortDRX != nil}); err != nil {
		return err
	}

	// 2. drx-onDurationTimer: choice
	{
		choiceEnc := e.NewChoiceEncoder(dRX_ConfigDrxOnDurationTimerConstraints)
		if err := choiceEnc.EncodeChoice(int64(ie.Drx_OnDurationTimer.Choice), false, nil); err != nil {
			return err
		}
		switch ie.Drx_OnDurationTimer.Choice {
		case DRX_Config_Drx_OnDurationTimer_SubMilliSeconds:
			if err := e.EncodeInteger((*ie.Drx_OnDurationTimer.SubMilliSeconds), per.Constrained(1, 31)); err != nil {
				return err
			}
		case DRX_Config_Drx_OnDurationTimer_MilliSeconds:
			if err := e.EncodeEnumerated((*ie.Drx_OnDurationTimer.MilliSeconds), dRXConfigDrxOnDurationTimerMilliSecondsConstraints); err != nil {
				return err
			}
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.Drx_OnDurationTimer.Choice), Constraint: "index out of range"}
		}
	}

	// 3. drx-InactivityTimer: enumerated
	{
		if err := e.EncodeEnumerated(ie.Drx_InactivityTimer, dRXConfigDrxInactivityTimerConstraints); err != nil {
			return err
		}
	}

	// 4. drx-HARQ-RTT-TimerDL: integer
	{
		if err := e.EncodeInteger(ie.Drx_HARQ_RTT_TimerDL, dRXConfigDrxHARQRTTTimerDLConstraints); err != nil {
			return err
		}
	}

	// 5. drx-HARQ-RTT-TimerUL: integer
	{
		if err := e.EncodeInteger(ie.Drx_HARQ_RTT_TimerUL, dRXConfigDrxHARQRTTTimerULConstraints); err != nil {
			return err
		}
	}

	// 6. drx-RetransmissionTimerDL: enumerated
	{
		if err := e.EncodeEnumerated(ie.Drx_RetransmissionTimerDL, dRXConfigDrxRetransmissionTimerDLConstraints); err != nil {
			return err
		}
	}

	// 7. drx-RetransmissionTimerUL: enumerated
	{
		if err := e.EncodeEnumerated(ie.Drx_RetransmissionTimerUL, dRXConfigDrxRetransmissionTimerULConstraints); err != nil {
			return err
		}
	}

	// 8. drx-LongCycleStartOffset: choice
	{
		choiceEnc := e.NewChoiceEncoder(dRX_ConfigDrxLongCycleStartOffsetConstraints)
		if err := choiceEnc.EncodeChoice(int64(ie.Drx_LongCycleStartOffset.Choice), false, nil); err != nil {
			return err
		}
		switch ie.Drx_LongCycleStartOffset.Choice {
		case DRX_Config_Drx_LongCycleStartOffset_Ms10:
			if err := e.EncodeInteger((*ie.Drx_LongCycleStartOffset.Ms10), per.Constrained(0, 9)); err != nil {
				return err
			}
		case DRX_Config_Drx_LongCycleStartOffset_Ms20:
			if err := e.EncodeInteger((*ie.Drx_LongCycleStartOffset.Ms20), per.Constrained(0, 19)); err != nil {
				return err
			}
		case DRX_Config_Drx_LongCycleStartOffset_Ms32:
			if err := e.EncodeInteger((*ie.Drx_LongCycleStartOffset.Ms32), per.Constrained(0, 31)); err != nil {
				return err
			}
		case DRX_Config_Drx_LongCycleStartOffset_Ms40:
			if err := e.EncodeInteger((*ie.Drx_LongCycleStartOffset.Ms40), per.Constrained(0, 39)); err != nil {
				return err
			}
		case DRX_Config_Drx_LongCycleStartOffset_Ms60:
			if err := e.EncodeInteger((*ie.Drx_LongCycleStartOffset.Ms60), per.Constrained(0, 59)); err != nil {
				return err
			}
		case DRX_Config_Drx_LongCycleStartOffset_Ms64:
			if err := e.EncodeInteger((*ie.Drx_LongCycleStartOffset.Ms64), per.Constrained(0, 63)); err != nil {
				return err
			}
		case DRX_Config_Drx_LongCycleStartOffset_Ms70:
			if err := e.EncodeInteger((*ie.Drx_LongCycleStartOffset.Ms70), per.Constrained(0, 69)); err != nil {
				return err
			}
		case DRX_Config_Drx_LongCycleStartOffset_Ms80:
			if err := e.EncodeInteger((*ie.Drx_LongCycleStartOffset.Ms80), per.Constrained(0, 79)); err != nil {
				return err
			}
		case DRX_Config_Drx_LongCycleStartOffset_Ms128:
			if err := e.EncodeInteger((*ie.Drx_LongCycleStartOffset.Ms128), per.Constrained(0, 127)); err != nil {
				return err
			}
		case DRX_Config_Drx_LongCycleStartOffset_Ms160:
			if err := e.EncodeInteger((*ie.Drx_LongCycleStartOffset.Ms160), per.Constrained(0, 159)); err != nil {
				return err
			}
		case DRX_Config_Drx_LongCycleStartOffset_Ms256:
			if err := e.EncodeInteger((*ie.Drx_LongCycleStartOffset.Ms256), per.Constrained(0, 255)); err != nil {
				return err
			}
		case DRX_Config_Drx_LongCycleStartOffset_Ms320:
			if err := e.EncodeInteger((*ie.Drx_LongCycleStartOffset.Ms320), per.Constrained(0, 319)); err != nil {
				return err
			}
		case DRX_Config_Drx_LongCycleStartOffset_Ms512:
			if err := e.EncodeInteger((*ie.Drx_LongCycleStartOffset.Ms512), per.Constrained(0, 511)); err != nil {
				return err
			}
		case DRX_Config_Drx_LongCycleStartOffset_Ms640:
			if err := e.EncodeInteger((*ie.Drx_LongCycleStartOffset.Ms640), per.Constrained(0, 639)); err != nil {
				return err
			}
		case DRX_Config_Drx_LongCycleStartOffset_Ms1024:
			if err := e.EncodeInteger((*ie.Drx_LongCycleStartOffset.Ms1024), per.Constrained(0, 1023)); err != nil {
				return err
			}
		case DRX_Config_Drx_LongCycleStartOffset_Ms1280:
			if err := e.EncodeInteger((*ie.Drx_LongCycleStartOffset.Ms1280), per.Constrained(0, 1279)); err != nil {
				return err
			}
		case DRX_Config_Drx_LongCycleStartOffset_Ms2048:
			if err := e.EncodeInteger((*ie.Drx_LongCycleStartOffset.Ms2048), per.Constrained(0, 2047)); err != nil {
				return err
			}
		case DRX_Config_Drx_LongCycleStartOffset_Ms2560:
			if err := e.EncodeInteger((*ie.Drx_LongCycleStartOffset.Ms2560), per.Constrained(0, 2559)); err != nil {
				return err
			}
		case DRX_Config_Drx_LongCycleStartOffset_Ms5120:
			if err := e.EncodeInteger((*ie.Drx_LongCycleStartOffset.Ms5120), per.Constrained(0, 5119)); err != nil {
				return err
			}
		case DRX_Config_Drx_LongCycleStartOffset_Ms10240:
			if err := e.EncodeInteger((*ie.Drx_LongCycleStartOffset.Ms10240), per.Constrained(0, 10239)); err != nil {
				return err
			}
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.Drx_LongCycleStartOffset.Choice), Constraint: "index out of range"}
		}
	}

	// 9. shortDRX: sequence
	{
		if ie.ShortDRX != nil {
			c := ie.ShortDRX
			if err := e.EncodeEnumerated(c.Drx_ShortCycle, dRXConfigShortDRXDrxShortCycleConstraints); err != nil {
				return err
			}
			if err := e.EncodeInteger(c.Drx_ShortCycleTimer, per.Constrained(1, 16)); err != nil {
				return err
			}
		}
	}

	// 10. drx-SlotOffset: integer
	{
		if err := e.EncodeInteger(ie.Drx_SlotOffset, dRXConfigDrxSlotOffsetConstraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *DRX_Config) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(dRXConfigConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. drx-onDurationTimer: choice
	{
		choiceDec := d.NewChoiceDecoder(dRX_ConfigDrxOnDurationTimerConstraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.Drx_OnDurationTimer.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case DRX_Config_Drx_OnDurationTimer_SubMilliSeconds:
			ie.Drx_OnDurationTimer.SubMilliSeconds = new(int64)
			v, err := d.DecodeInteger(per.Constrained(1, 31))
			if err != nil {
				return err
			}
			(*ie.Drx_OnDurationTimer.SubMilliSeconds) = v
		case DRX_Config_Drx_OnDurationTimer_MilliSeconds:
			ie.Drx_OnDurationTimer.MilliSeconds = new(int64)
			v, err := d.DecodeEnumerated(dRXConfigDrxOnDurationTimerMilliSecondsConstraints)
			if err != nil {
				return err
			}
			(*ie.Drx_OnDurationTimer.MilliSeconds) = v
		}
	}

	// 3. drx-InactivityTimer: enumerated
	{
		v1, err := d.DecodeEnumerated(dRXConfigDrxInactivityTimerConstraints)
		if err != nil {
			return err
		}
		ie.Drx_InactivityTimer = v1
	}

	// 4. drx-HARQ-RTT-TimerDL: integer
	{
		v2, err := d.DecodeInteger(dRXConfigDrxHARQRTTTimerDLConstraints)
		if err != nil {
			return err
		}
		ie.Drx_HARQ_RTT_TimerDL = v2
	}

	// 5. drx-HARQ-RTT-TimerUL: integer
	{
		v3, err := d.DecodeInteger(dRXConfigDrxHARQRTTTimerULConstraints)
		if err != nil {
			return err
		}
		ie.Drx_HARQ_RTT_TimerUL = v3
	}

	// 6. drx-RetransmissionTimerDL: enumerated
	{
		v4, err := d.DecodeEnumerated(dRXConfigDrxRetransmissionTimerDLConstraints)
		if err != nil {
			return err
		}
		ie.Drx_RetransmissionTimerDL = v4
	}

	// 7. drx-RetransmissionTimerUL: enumerated
	{
		v5, err := d.DecodeEnumerated(dRXConfigDrxRetransmissionTimerULConstraints)
		if err != nil {
			return err
		}
		ie.Drx_RetransmissionTimerUL = v5
	}

	// 8. drx-LongCycleStartOffset: choice
	{
		choiceDec := d.NewChoiceDecoder(dRX_ConfigDrxLongCycleStartOffsetConstraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.Drx_LongCycleStartOffset.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case DRX_Config_Drx_LongCycleStartOffset_Ms10:
			ie.Drx_LongCycleStartOffset.Ms10 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 9))
			if err != nil {
				return err
			}
			(*ie.Drx_LongCycleStartOffset.Ms10) = v
		case DRX_Config_Drx_LongCycleStartOffset_Ms20:
			ie.Drx_LongCycleStartOffset.Ms20 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 19))
			if err != nil {
				return err
			}
			(*ie.Drx_LongCycleStartOffset.Ms20) = v
		case DRX_Config_Drx_LongCycleStartOffset_Ms32:
			ie.Drx_LongCycleStartOffset.Ms32 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 31))
			if err != nil {
				return err
			}
			(*ie.Drx_LongCycleStartOffset.Ms32) = v
		case DRX_Config_Drx_LongCycleStartOffset_Ms40:
			ie.Drx_LongCycleStartOffset.Ms40 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 39))
			if err != nil {
				return err
			}
			(*ie.Drx_LongCycleStartOffset.Ms40) = v
		case DRX_Config_Drx_LongCycleStartOffset_Ms60:
			ie.Drx_LongCycleStartOffset.Ms60 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 59))
			if err != nil {
				return err
			}
			(*ie.Drx_LongCycleStartOffset.Ms60) = v
		case DRX_Config_Drx_LongCycleStartOffset_Ms64:
			ie.Drx_LongCycleStartOffset.Ms64 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 63))
			if err != nil {
				return err
			}
			(*ie.Drx_LongCycleStartOffset.Ms64) = v
		case DRX_Config_Drx_LongCycleStartOffset_Ms70:
			ie.Drx_LongCycleStartOffset.Ms70 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 69))
			if err != nil {
				return err
			}
			(*ie.Drx_LongCycleStartOffset.Ms70) = v
		case DRX_Config_Drx_LongCycleStartOffset_Ms80:
			ie.Drx_LongCycleStartOffset.Ms80 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 79))
			if err != nil {
				return err
			}
			(*ie.Drx_LongCycleStartOffset.Ms80) = v
		case DRX_Config_Drx_LongCycleStartOffset_Ms128:
			ie.Drx_LongCycleStartOffset.Ms128 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 127))
			if err != nil {
				return err
			}
			(*ie.Drx_LongCycleStartOffset.Ms128) = v
		case DRX_Config_Drx_LongCycleStartOffset_Ms160:
			ie.Drx_LongCycleStartOffset.Ms160 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 159))
			if err != nil {
				return err
			}
			(*ie.Drx_LongCycleStartOffset.Ms160) = v
		case DRX_Config_Drx_LongCycleStartOffset_Ms256:
			ie.Drx_LongCycleStartOffset.Ms256 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 255))
			if err != nil {
				return err
			}
			(*ie.Drx_LongCycleStartOffset.Ms256) = v
		case DRX_Config_Drx_LongCycleStartOffset_Ms320:
			ie.Drx_LongCycleStartOffset.Ms320 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 319))
			if err != nil {
				return err
			}
			(*ie.Drx_LongCycleStartOffset.Ms320) = v
		case DRX_Config_Drx_LongCycleStartOffset_Ms512:
			ie.Drx_LongCycleStartOffset.Ms512 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 511))
			if err != nil {
				return err
			}
			(*ie.Drx_LongCycleStartOffset.Ms512) = v
		case DRX_Config_Drx_LongCycleStartOffset_Ms640:
			ie.Drx_LongCycleStartOffset.Ms640 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 639))
			if err != nil {
				return err
			}
			(*ie.Drx_LongCycleStartOffset.Ms640) = v
		case DRX_Config_Drx_LongCycleStartOffset_Ms1024:
			ie.Drx_LongCycleStartOffset.Ms1024 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 1023))
			if err != nil {
				return err
			}
			(*ie.Drx_LongCycleStartOffset.Ms1024) = v
		case DRX_Config_Drx_LongCycleStartOffset_Ms1280:
			ie.Drx_LongCycleStartOffset.Ms1280 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 1279))
			if err != nil {
				return err
			}
			(*ie.Drx_LongCycleStartOffset.Ms1280) = v
		case DRX_Config_Drx_LongCycleStartOffset_Ms2048:
			ie.Drx_LongCycleStartOffset.Ms2048 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 2047))
			if err != nil {
				return err
			}
			(*ie.Drx_LongCycleStartOffset.Ms2048) = v
		case DRX_Config_Drx_LongCycleStartOffset_Ms2560:
			ie.Drx_LongCycleStartOffset.Ms2560 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 2559))
			if err != nil {
				return err
			}
			(*ie.Drx_LongCycleStartOffset.Ms2560) = v
		case DRX_Config_Drx_LongCycleStartOffset_Ms5120:
			ie.Drx_LongCycleStartOffset.Ms5120 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 5119))
			if err != nil {
				return err
			}
			(*ie.Drx_LongCycleStartOffset.Ms5120) = v
		case DRX_Config_Drx_LongCycleStartOffset_Ms10240:
			ie.Drx_LongCycleStartOffset.Ms10240 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 10239))
			if err != nil {
				return err
			}
			(*ie.Drx_LongCycleStartOffset.Ms10240) = v
		}
	}

	// 9. shortDRX: sequence
	{
		if seq.IsComponentPresent(7) {
			ie.ShortDRX = &struct {
				Drx_ShortCycle      int64
				Drx_ShortCycleTimer int64
			}{}
			c := ie.ShortDRX
			{
				v, err := d.DecodeEnumerated(dRXConfigShortDRXDrxShortCycleConstraints)
				if err != nil {
					return err
				}
				c.Drx_ShortCycle = v
			}
			{
				v, err := d.DecodeInteger(per.Constrained(1, 16))
				if err != nil {
					return err
				}
				c.Drx_ShortCycleTimer = v
			}
		}
	}

	// 10. drx-SlotOffset: integer
	{
		v8, err := d.DecodeInteger(dRXConfigDrxSlotOffsetConstraints)
		if err != nil {
			return err
		}
		ie.Drx_SlotOffset = v8
	}

	return nil
}
