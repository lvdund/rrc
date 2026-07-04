// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SL-DRX-ConfigUC-r17 (line 27171).

var sLDRXConfigUCR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-drx-onDurationTimer-r17"},
		{Name: "sl-drx-InactivityTimer-r17"},
		{Name: "sl-drx-HARQ-RTT-Timer1-r17", Optional: true},
		{Name: "sl-drx-HARQ-RTT-Timer2-r17", Optional: true},
		{Name: "sl-drx-RetransmissionTimer-r17"},
		{Name: "sl-drx-CycleStartOffset-r17"},
		{Name: "sl-drx-SlotOffset"},
	},
}

var sL_DRX_ConfigUC_r17SlDrxOnDurationTimerR17Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "subMilliSeconds"},
		{Name: "milliSeconds"},
	},
}

const (
	SL_DRX_ConfigUC_r17_Sl_Drx_OnDurationTimer_r17_SubMilliSeconds = 0
	SL_DRX_ConfigUC_r17_Sl_Drx_OnDurationTimer_r17_MilliSeconds    = 1
)

const (
	SL_DRX_ConfigUC_r17_Sl_Drx_InactivityTimer_r17_Ms0    = 0
	SL_DRX_ConfigUC_r17_Sl_Drx_InactivityTimer_r17_Ms1    = 1
	SL_DRX_ConfigUC_r17_Sl_Drx_InactivityTimer_r17_Ms2    = 2
	SL_DRX_ConfigUC_r17_Sl_Drx_InactivityTimer_r17_Ms3    = 3
	SL_DRX_ConfigUC_r17_Sl_Drx_InactivityTimer_r17_Ms4    = 4
	SL_DRX_ConfigUC_r17_Sl_Drx_InactivityTimer_r17_Ms5    = 5
	SL_DRX_ConfigUC_r17_Sl_Drx_InactivityTimer_r17_Ms6    = 6
	SL_DRX_ConfigUC_r17_Sl_Drx_InactivityTimer_r17_Ms8    = 7
	SL_DRX_ConfigUC_r17_Sl_Drx_InactivityTimer_r17_Ms10   = 8
	SL_DRX_ConfigUC_r17_Sl_Drx_InactivityTimer_r17_Ms20   = 9
	SL_DRX_ConfigUC_r17_Sl_Drx_InactivityTimer_r17_Ms30   = 10
	SL_DRX_ConfigUC_r17_Sl_Drx_InactivityTimer_r17_Ms40   = 11
	SL_DRX_ConfigUC_r17_Sl_Drx_InactivityTimer_r17_Ms50   = 12
	SL_DRX_ConfigUC_r17_Sl_Drx_InactivityTimer_r17_Ms60   = 13
	SL_DRX_ConfigUC_r17_Sl_Drx_InactivityTimer_r17_Ms80   = 14
	SL_DRX_ConfigUC_r17_Sl_Drx_InactivityTimer_r17_Ms100  = 15
	SL_DRX_ConfigUC_r17_Sl_Drx_InactivityTimer_r17_Ms200  = 16
	SL_DRX_ConfigUC_r17_Sl_Drx_InactivityTimer_r17_Ms300  = 17
	SL_DRX_ConfigUC_r17_Sl_Drx_InactivityTimer_r17_Ms500  = 18
	SL_DRX_ConfigUC_r17_Sl_Drx_InactivityTimer_r17_Ms750  = 19
	SL_DRX_ConfigUC_r17_Sl_Drx_InactivityTimer_r17_Ms1280 = 20
	SL_DRX_ConfigUC_r17_Sl_Drx_InactivityTimer_r17_Ms1920 = 21
	SL_DRX_ConfigUC_r17_Sl_Drx_InactivityTimer_r17_Ms2560 = 22
	SL_DRX_ConfigUC_r17_Sl_Drx_InactivityTimer_r17_Spare9 = 23
	SL_DRX_ConfigUC_r17_Sl_Drx_InactivityTimer_r17_Spare8 = 24
	SL_DRX_ConfigUC_r17_Sl_Drx_InactivityTimer_r17_Spare7 = 25
	SL_DRX_ConfigUC_r17_Sl_Drx_InactivityTimer_r17_Spare6 = 26
	SL_DRX_ConfigUC_r17_Sl_Drx_InactivityTimer_r17_Spare5 = 27
	SL_DRX_ConfigUC_r17_Sl_Drx_InactivityTimer_r17_Spare4 = 28
	SL_DRX_ConfigUC_r17_Sl_Drx_InactivityTimer_r17_Spare3 = 29
	SL_DRX_ConfigUC_r17_Sl_Drx_InactivityTimer_r17_Spare2 = 30
	SL_DRX_ConfigUC_r17_Sl_Drx_InactivityTimer_r17_Spare1 = 31
)

var sLDRXConfigUCR17SlDrxInactivityTimerR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SL_DRX_ConfigUC_r17_Sl_Drx_InactivityTimer_r17_Ms0, SL_DRX_ConfigUC_r17_Sl_Drx_InactivityTimer_r17_Ms1, SL_DRX_ConfigUC_r17_Sl_Drx_InactivityTimer_r17_Ms2, SL_DRX_ConfigUC_r17_Sl_Drx_InactivityTimer_r17_Ms3, SL_DRX_ConfigUC_r17_Sl_Drx_InactivityTimer_r17_Ms4, SL_DRX_ConfigUC_r17_Sl_Drx_InactivityTimer_r17_Ms5, SL_DRX_ConfigUC_r17_Sl_Drx_InactivityTimer_r17_Ms6, SL_DRX_ConfigUC_r17_Sl_Drx_InactivityTimer_r17_Ms8, SL_DRX_ConfigUC_r17_Sl_Drx_InactivityTimer_r17_Ms10, SL_DRX_ConfigUC_r17_Sl_Drx_InactivityTimer_r17_Ms20, SL_DRX_ConfigUC_r17_Sl_Drx_InactivityTimer_r17_Ms30, SL_DRX_ConfigUC_r17_Sl_Drx_InactivityTimer_r17_Ms40, SL_DRX_ConfigUC_r17_Sl_Drx_InactivityTimer_r17_Ms50, SL_DRX_ConfigUC_r17_Sl_Drx_InactivityTimer_r17_Ms60, SL_DRX_ConfigUC_r17_Sl_Drx_InactivityTimer_r17_Ms80, SL_DRX_ConfigUC_r17_Sl_Drx_InactivityTimer_r17_Ms100, SL_DRX_ConfigUC_r17_Sl_Drx_InactivityTimer_r17_Ms200, SL_DRX_ConfigUC_r17_Sl_Drx_InactivityTimer_r17_Ms300, SL_DRX_ConfigUC_r17_Sl_Drx_InactivityTimer_r17_Ms500, SL_DRX_ConfigUC_r17_Sl_Drx_InactivityTimer_r17_Ms750, SL_DRX_ConfigUC_r17_Sl_Drx_InactivityTimer_r17_Ms1280, SL_DRX_ConfigUC_r17_Sl_Drx_InactivityTimer_r17_Ms1920, SL_DRX_ConfigUC_r17_Sl_Drx_InactivityTimer_r17_Ms2560, SL_DRX_ConfigUC_r17_Sl_Drx_InactivityTimer_r17_Spare9, SL_DRX_ConfigUC_r17_Sl_Drx_InactivityTimer_r17_Spare8, SL_DRX_ConfigUC_r17_Sl_Drx_InactivityTimer_r17_Spare7, SL_DRX_ConfigUC_r17_Sl_Drx_InactivityTimer_r17_Spare6, SL_DRX_ConfigUC_r17_Sl_Drx_InactivityTimer_r17_Spare5, SL_DRX_ConfigUC_r17_Sl_Drx_InactivityTimer_r17_Spare4, SL_DRX_ConfigUC_r17_Sl_Drx_InactivityTimer_r17_Spare3, SL_DRX_ConfigUC_r17_Sl_Drx_InactivityTimer_r17_Spare2, SL_DRX_ConfigUC_r17_Sl_Drx_InactivityTimer_r17_Spare1},
}

const (
	SL_DRX_ConfigUC_r17_Sl_Drx_HARQ_RTT_Timer1_r17_Sl0    = 0
	SL_DRX_ConfigUC_r17_Sl_Drx_HARQ_RTT_Timer1_r17_Sl1    = 1
	SL_DRX_ConfigUC_r17_Sl_Drx_HARQ_RTT_Timer1_r17_Sl2    = 2
	SL_DRX_ConfigUC_r17_Sl_Drx_HARQ_RTT_Timer1_r17_Sl4    = 3
	SL_DRX_ConfigUC_r17_Sl_Drx_HARQ_RTT_Timer1_r17_Spare4 = 4
	SL_DRX_ConfigUC_r17_Sl_Drx_HARQ_RTT_Timer1_r17_Spare3 = 5
	SL_DRX_ConfigUC_r17_Sl_Drx_HARQ_RTT_Timer1_r17_Spare2 = 6
	SL_DRX_ConfigUC_r17_Sl_Drx_HARQ_RTT_Timer1_r17_Spare1 = 7
)

var sLDRXConfigUCR17SlDrxHARQRTTTimer1R17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SL_DRX_ConfigUC_r17_Sl_Drx_HARQ_RTT_Timer1_r17_Sl0, SL_DRX_ConfigUC_r17_Sl_Drx_HARQ_RTT_Timer1_r17_Sl1, SL_DRX_ConfigUC_r17_Sl_Drx_HARQ_RTT_Timer1_r17_Sl2, SL_DRX_ConfigUC_r17_Sl_Drx_HARQ_RTT_Timer1_r17_Sl4, SL_DRX_ConfigUC_r17_Sl_Drx_HARQ_RTT_Timer1_r17_Spare4, SL_DRX_ConfigUC_r17_Sl_Drx_HARQ_RTT_Timer1_r17_Spare3, SL_DRX_ConfigUC_r17_Sl_Drx_HARQ_RTT_Timer1_r17_Spare2, SL_DRX_ConfigUC_r17_Sl_Drx_HARQ_RTT_Timer1_r17_Spare1},
}

const (
	SL_DRX_ConfigUC_r17_Sl_Drx_HARQ_RTT_Timer2_r17_Sl0    = 0
	SL_DRX_ConfigUC_r17_Sl_Drx_HARQ_RTT_Timer2_r17_Sl1    = 1
	SL_DRX_ConfigUC_r17_Sl_Drx_HARQ_RTT_Timer2_r17_Sl2    = 2
	SL_DRX_ConfigUC_r17_Sl_Drx_HARQ_RTT_Timer2_r17_Sl4    = 3
	SL_DRX_ConfigUC_r17_Sl_Drx_HARQ_RTT_Timer2_r17_Spare4 = 4
	SL_DRX_ConfigUC_r17_Sl_Drx_HARQ_RTT_Timer2_r17_Spare3 = 5
	SL_DRX_ConfigUC_r17_Sl_Drx_HARQ_RTT_Timer2_r17_Spare2 = 6
	SL_DRX_ConfigUC_r17_Sl_Drx_HARQ_RTT_Timer2_r17_Spare1 = 7
)

var sLDRXConfigUCR17SlDrxHARQRTTTimer2R17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SL_DRX_ConfigUC_r17_Sl_Drx_HARQ_RTT_Timer2_r17_Sl0, SL_DRX_ConfigUC_r17_Sl_Drx_HARQ_RTT_Timer2_r17_Sl1, SL_DRX_ConfigUC_r17_Sl_Drx_HARQ_RTT_Timer2_r17_Sl2, SL_DRX_ConfigUC_r17_Sl_Drx_HARQ_RTT_Timer2_r17_Sl4, SL_DRX_ConfigUC_r17_Sl_Drx_HARQ_RTT_Timer2_r17_Spare4, SL_DRX_ConfigUC_r17_Sl_Drx_HARQ_RTT_Timer2_r17_Spare3, SL_DRX_ConfigUC_r17_Sl_Drx_HARQ_RTT_Timer2_r17_Spare2, SL_DRX_ConfigUC_r17_Sl_Drx_HARQ_RTT_Timer2_r17_Spare1},
}

const (
	SL_DRX_ConfigUC_r17_Sl_Drx_RetransmissionTimer_r17_Sl0     = 0
	SL_DRX_ConfigUC_r17_Sl_Drx_RetransmissionTimer_r17_Sl1     = 1
	SL_DRX_ConfigUC_r17_Sl_Drx_RetransmissionTimer_r17_Sl2     = 2
	SL_DRX_ConfigUC_r17_Sl_Drx_RetransmissionTimer_r17_Sl4     = 3
	SL_DRX_ConfigUC_r17_Sl_Drx_RetransmissionTimer_r17_Sl6     = 4
	SL_DRX_ConfigUC_r17_Sl_Drx_RetransmissionTimer_r17_Sl8     = 5
	SL_DRX_ConfigUC_r17_Sl_Drx_RetransmissionTimer_r17_Sl16    = 6
	SL_DRX_ConfigUC_r17_Sl_Drx_RetransmissionTimer_r17_Sl24    = 7
	SL_DRX_ConfigUC_r17_Sl_Drx_RetransmissionTimer_r17_Sl33    = 8
	SL_DRX_ConfigUC_r17_Sl_Drx_RetransmissionTimer_r17_Sl40    = 9
	SL_DRX_ConfigUC_r17_Sl_Drx_RetransmissionTimer_r17_Sl64    = 10
	SL_DRX_ConfigUC_r17_Sl_Drx_RetransmissionTimer_r17_Sl80    = 11
	SL_DRX_ConfigUC_r17_Sl_Drx_RetransmissionTimer_r17_Sl96    = 12
	SL_DRX_ConfigUC_r17_Sl_Drx_RetransmissionTimer_r17_Sl112   = 13
	SL_DRX_ConfigUC_r17_Sl_Drx_RetransmissionTimer_r17_Sl128   = 14
	SL_DRX_ConfigUC_r17_Sl_Drx_RetransmissionTimer_r17_Sl160   = 15
	SL_DRX_ConfigUC_r17_Sl_Drx_RetransmissionTimer_r17_Sl320   = 16
	SL_DRX_ConfigUC_r17_Sl_Drx_RetransmissionTimer_r17_Spare15 = 17
	SL_DRX_ConfigUC_r17_Sl_Drx_RetransmissionTimer_r17_Spare14 = 18
	SL_DRX_ConfigUC_r17_Sl_Drx_RetransmissionTimer_r17_Spare13 = 19
	SL_DRX_ConfigUC_r17_Sl_Drx_RetransmissionTimer_r17_Spare12 = 20
	SL_DRX_ConfigUC_r17_Sl_Drx_RetransmissionTimer_r17_Spare11 = 21
	SL_DRX_ConfigUC_r17_Sl_Drx_RetransmissionTimer_r17_Spare10 = 22
	SL_DRX_ConfigUC_r17_Sl_Drx_RetransmissionTimer_r17_Spare9  = 23
	SL_DRX_ConfigUC_r17_Sl_Drx_RetransmissionTimer_r17_Spare8  = 24
	SL_DRX_ConfigUC_r17_Sl_Drx_RetransmissionTimer_r17_Spare7  = 25
	SL_DRX_ConfigUC_r17_Sl_Drx_RetransmissionTimer_r17_Spare6  = 26
	SL_DRX_ConfigUC_r17_Sl_Drx_RetransmissionTimer_r17_Spare5  = 27
	SL_DRX_ConfigUC_r17_Sl_Drx_RetransmissionTimer_r17_Spare4  = 28
	SL_DRX_ConfigUC_r17_Sl_Drx_RetransmissionTimer_r17_Spare3  = 29
	SL_DRX_ConfigUC_r17_Sl_Drx_RetransmissionTimer_r17_Spare2  = 30
	SL_DRX_ConfigUC_r17_Sl_Drx_RetransmissionTimer_r17_Spare1  = 31
)

var sLDRXConfigUCR17SlDrxRetransmissionTimerR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SL_DRX_ConfigUC_r17_Sl_Drx_RetransmissionTimer_r17_Sl0, SL_DRX_ConfigUC_r17_Sl_Drx_RetransmissionTimer_r17_Sl1, SL_DRX_ConfigUC_r17_Sl_Drx_RetransmissionTimer_r17_Sl2, SL_DRX_ConfigUC_r17_Sl_Drx_RetransmissionTimer_r17_Sl4, SL_DRX_ConfigUC_r17_Sl_Drx_RetransmissionTimer_r17_Sl6, SL_DRX_ConfigUC_r17_Sl_Drx_RetransmissionTimer_r17_Sl8, SL_DRX_ConfigUC_r17_Sl_Drx_RetransmissionTimer_r17_Sl16, SL_DRX_ConfigUC_r17_Sl_Drx_RetransmissionTimer_r17_Sl24, SL_DRX_ConfigUC_r17_Sl_Drx_RetransmissionTimer_r17_Sl33, SL_DRX_ConfigUC_r17_Sl_Drx_RetransmissionTimer_r17_Sl40, SL_DRX_ConfigUC_r17_Sl_Drx_RetransmissionTimer_r17_Sl64, SL_DRX_ConfigUC_r17_Sl_Drx_RetransmissionTimer_r17_Sl80, SL_DRX_ConfigUC_r17_Sl_Drx_RetransmissionTimer_r17_Sl96, SL_DRX_ConfigUC_r17_Sl_Drx_RetransmissionTimer_r17_Sl112, SL_DRX_ConfigUC_r17_Sl_Drx_RetransmissionTimer_r17_Sl128, SL_DRX_ConfigUC_r17_Sl_Drx_RetransmissionTimer_r17_Sl160, SL_DRX_ConfigUC_r17_Sl_Drx_RetransmissionTimer_r17_Sl320, SL_DRX_ConfigUC_r17_Sl_Drx_RetransmissionTimer_r17_Spare15, SL_DRX_ConfigUC_r17_Sl_Drx_RetransmissionTimer_r17_Spare14, SL_DRX_ConfigUC_r17_Sl_Drx_RetransmissionTimer_r17_Spare13, SL_DRX_ConfigUC_r17_Sl_Drx_RetransmissionTimer_r17_Spare12, SL_DRX_ConfigUC_r17_Sl_Drx_RetransmissionTimer_r17_Spare11, SL_DRX_ConfigUC_r17_Sl_Drx_RetransmissionTimer_r17_Spare10, SL_DRX_ConfigUC_r17_Sl_Drx_RetransmissionTimer_r17_Spare9, SL_DRX_ConfigUC_r17_Sl_Drx_RetransmissionTimer_r17_Spare8, SL_DRX_ConfigUC_r17_Sl_Drx_RetransmissionTimer_r17_Spare7, SL_DRX_ConfigUC_r17_Sl_Drx_RetransmissionTimer_r17_Spare6, SL_DRX_ConfigUC_r17_Sl_Drx_RetransmissionTimer_r17_Spare5, SL_DRX_ConfigUC_r17_Sl_Drx_RetransmissionTimer_r17_Spare4, SL_DRX_ConfigUC_r17_Sl_Drx_RetransmissionTimer_r17_Spare3, SL_DRX_ConfigUC_r17_Sl_Drx_RetransmissionTimer_r17_Spare2, SL_DRX_ConfigUC_r17_Sl_Drx_RetransmissionTimer_r17_Spare1},
}

var sL_DRX_ConfigUC_r17SlDrxCycleStartOffsetR17Constraints = per.ChoiceConstraints{
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
	SL_DRX_ConfigUC_r17_Sl_Drx_CycleStartOffset_r17_Ms10    = 0
	SL_DRX_ConfigUC_r17_Sl_Drx_CycleStartOffset_r17_Ms20    = 1
	SL_DRX_ConfigUC_r17_Sl_Drx_CycleStartOffset_r17_Ms32    = 2
	SL_DRX_ConfigUC_r17_Sl_Drx_CycleStartOffset_r17_Ms40    = 3
	SL_DRX_ConfigUC_r17_Sl_Drx_CycleStartOffset_r17_Ms60    = 4
	SL_DRX_ConfigUC_r17_Sl_Drx_CycleStartOffset_r17_Ms64    = 5
	SL_DRX_ConfigUC_r17_Sl_Drx_CycleStartOffset_r17_Ms70    = 6
	SL_DRX_ConfigUC_r17_Sl_Drx_CycleStartOffset_r17_Ms80    = 7
	SL_DRX_ConfigUC_r17_Sl_Drx_CycleStartOffset_r17_Ms128   = 8
	SL_DRX_ConfigUC_r17_Sl_Drx_CycleStartOffset_r17_Ms160   = 9
	SL_DRX_ConfigUC_r17_Sl_Drx_CycleStartOffset_r17_Ms256   = 10
	SL_DRX_ConfigUC_r17_Sl_Drx_CycleStartOffset_r17_Ms320   = 11
	SL_DRX_ConfigUC_r17_Sl_Drx_CycleStartOffset_r17_Ms512   = 12
	SL_DRX_ConfigUC_r17_Sl_Drx_CycleStartOffset_r17_Ms640   = 13
	SL_DRX_ConfigUC_r17_Sl_Drx_CycleStartOffset_r17_Ms1024  = 14
	SL_DRX_ConfigUC_r17_Sl_Drx_CycleStartOffset_r17_Ms1280  = 15
	SL_DRX_ConfigUC_r17_Sl_Drx_CycleStartOffset_r17_Ms2048  = 16
	SL_DRX_ConfigUC_r17_Sl_Drx_CycleStartOffset_r17_Ms2560  = 17
	SL_DRX_ConfigUC_r17_Sl_Drx_CycleStartOffset_r17_Ms5120  = 18
	SL_DRX_ConfigUC_r17_Sl_Drx_CycleStartOffset_r17_Ms10240 = 19
)

var sLDRXConfigUCR17SlDrxSlotOffsetConstraints = per.Constrained(0, 31)

const (
	SL_DRX_ConfigUC_r17_Sl_Drx_OnDurationTimer_r17_MilliSeconds_Ms1    = 0
	SL_DRX_ConfigUC_r17_Sl_Drx_OnDurationTimer_r17_MilliSeconds_Ms2    = 1
	SL_DRX_ConfigUC_r17_Sl_Drx_OnDurationTimer_r17_MilliSeconds_Ms3    = 2
	SL_DRX_ConfigUC_r17_Sl_Drx_OnDurationTimer_r17_MilliSeconds_Ms4    = 3
	SL_DRX_ConfigUC_r17_Sl_Drx_OnDurationTimer_r17_MilliSeconds_Ms5    = 4
	SL_DRX_ConfigUC_r17_Sl_Drx_OnDurationTimer_r17_MilliSeconds_Ms6    = 5
	SL_DRX_ConfigUC_r17_Sl_Drx_OnDurationTimer_r17_MilliSeconds_Ms8    = 6
	SL_DRX_ConfigUC_r17_Sl_Drx_OnDurationTimer_r17_MilliSeconds_Ms10   = 7
	SL_DRX_ConfigUC_r17_Sl_Drx_OnDurationTimer_r17_MilliSeconds_Ms20   = 8
	SL_DRX_ConfigUC_r17_Sl_Drx_OnDurationTimer_r17_MilliSeconds_Ms30   = 9
	SL_DRX_ConfigUC_r17_Sl_Drx_OnDurationTimer_r17_MilliSeconds_Ms40   = 10
	SL_DRX_ConfigUC_r17_Sl_Drx_OnDurationTimer_r17_MilliSeconds_Ms50   = 11
	SL_DRX_ConfigUC_r17_Sl_Drx_OnDurationTimer_r17_MilliSeconds_Ms60   = 12
	SL_DRX_ConfigUC_r17_Sl_Drx_OnDurationTimer_r17_MilliSeconds_Ms80   = 13
	SL_DRX_ConfigUC_r17_Sl_Drx_OnDurationTimer_r17_MilliSeconds_Ms100  = 14
	SL_DRX_ConfigUC_r17_Sl_Drx_OnDurationTimer_r17_MilliSeconds_Ms200  = 15
	SL_DRX_ConfigUC_r17_Sl_Drx_OnDurationTimer_r17_MilliSeconds_Ms300  = 16
	SL_DRX_ConfigUC_r17_Sl_Drx_OnDurationTimer_r17_MilliSeconds_Ms400  = 17
	SL_DRX_ConfigUC_r17_Sl_Drx_OnDurationTimer_r17_MilliSeconds_Ms500  = 18
	SL_DRX_ConfigUC_r17_Sl_Drx_OnDurationTimer_r17_MilliSeconds_Ms600  = 19
	SL_DRX_ConfigUC_r17_Sl_Drx_OnDurationTimer_r17_MilliSeconds_Ms800  = 20
	SL_DRX_ConfigUC_r17_Sl_Drx_OnDurationTimer_r17_MilliSeconds_Ms1000 = 21
	SL_DRX_ConfigUC_r17_Sl_Drx_OnDurationTimer_r17_MilliSeconds_Ms1200 = 22
	SL_DRX_ConfigUC_r17_Sl_Drx_OnDurationTimer_r17_MilliSeconds_Ms1600 = 23
	SL_DRX_ConfigUC_r17_Sl_Drx_OnDurationTimer_r17_MilliSeconds_Spare8 = 24
	SL_DRX_ConfigUC_r17_Sl_Drx_OnDurationTimer_r17_MilliSeconds_Spare7 = 25
	SL_DRX_ConfigUC_r17_Sl_Drx_OnDurationTimer_r17_MilliSeconds_Spare6 = 26
	SL_DRX_ConfigUC_r17_Sl_Drx_OnDurationTimer_r17_MilliSeconds_Spare5 = 27
	SL_DRX_ConfigUC_r17_Sl_Drx_OnDurationTimer_r17_MilliSeconds_Spare4 = 28
	SL_DRX_ConfigUC_r17_Sl_Drx_OnDurationTimer_r17_MilliSeconds_Spare3 = 29
	SL_DRX_ConfigUC_r17_Sl_Drx_OnDurationTimer_r17_MilliSeconds_Spare2 = 30
	SL_DRX_ConfigUC_r17_Sl_Drx_OnDurationTimer_r17_MilliSeconds_Spare1 = 31
)

var sLDRXConfigUCR17SlDrxOnDurationTimerR17MilliSecondsConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SL_DRX_ConfigUC_r17_Sl_Drx_OnDurationTimer_r17_MilliSeconds_Ms1, SL_DRX_ConfigUC_r17_Sl_Drx_OnDurationTimer_r17_MilliSeconds_Ms2, SL_DRX_ConfigUC_r17_Sl_Drx_OnDurationTimer_r17_MilliSeconds_Ms3, SL_DRX_ConfigUC_r17_Sl_Drx_OnDurationTimer_r17_MilliSeconds_Ms4, SL_DRX_ConfigUC_r17_Sl_Drx_OnDurationTimer_r17_MilliSeconds_Ms5, SL_DRX_ConfigUC_r17_Sl_Drx_OnDurationTimer_r17_MilliSeconds_Ms6, SL_DRX_ConfigUC_r17_Sl_Drx_OnDurationTimer_r17_MilliSeconds_Ms8, SL_DRX_ConfigUC_r17_Sl_Drx_OnDurationTimer_r17_MilliSeconds_Ms10, SL_DRX_ConfigUC_r17_Sl_Drx_OnDurationTimer_r17_MilliSeconds_Ms20, SL_DRX_ConfigUC_r17_Sl_Drx_OnDurationTimer_r17_MilliSeconds_Ms30, SL_DRX_ConfigUC_r17_Sl_Drx_OnDurationTimer_r17_MilliSeconds_Ms40, SL_DRX_ConfigUC_r17_Sl_Drx_OnDurationTimer_r17_MilliSeconds_Ms50, SL_DRX_ConfigUC_r17_Sl_Drx_OnDurationTimer_r17_MilliSeconds_Ms60, SL_DRX_ConfigUC_r17_Sl_Drx_OnDurationTimer_r17_MilliSeconds_Ms80, SL_DRX_ConfigUC_r17_Sl_Drx_OnDurationTimer_r17_MilliSeconds_Ms100, SL_DRX_ConfigUC_r17_Sl_Drx_OnDurationTimer_r17_MilliSeconds_Ms200, SL_DRX_ConfigUC_r17_Sl_Drx_OnDurationTimer_r17_MilliSeconds_Ms300, SL_DRX_ConfigUC_r17_Sl_Drx_OnDurationTimer_r17_MilliSeconds_Ms400, SL_DRX_ConfigUC_r17_Sl_Drx_OnDurationTimer_r17_MilliSeconds_Ms500, SL_DRX_ConfigUC_r17_Sl_Drx_OnDurationTimer_r17_MilliSeconds_Ms600, SL_DRX_ConfigUC_r17_Sl_Drx_OnDurationTimer_r17_MilliSeconds_Ms800, SL_DRX_ConfigUC_r17_Sl_Drx_OnDurationTimer_r17_MilliSeconds_Ms1000, SL_DRX_ConfigUC_r17_Sl_Drx_OnDurationTimer_r17_MilliSeconds_Ms1200, SL_DRX_ConfigUC_r17_Sl_Drx_OnDurationTimer_r17_MilliSeconds_Ms1600, SL_DRX_ConfigUC_r17_Sl_Drx_OnDurationTimer_r17_MilliSeconds_Spare8, SL_DRX_ConfigUC_r17_Sl_Drx_OnDurationTimer_r17_MilliSeconds_Spare7, SL_DRX_ConfigUC_r17_Sl_Drx_OnDurationTimer_r17_MilliSeconds_Spare6, SL_DRX_ConfigUC_r17_Sl_Drx_OnDurationTimer_r17_MilliSeconds_Spare5, SL_DRX_ConfigUC_r17_Sl_Drx_OnDurationTimer_r17_MilliSeconds_Spare4, SL_DRX_ConfigUC_r17_Sl_Drx_OnDurationTimer_r17_MilliSeconds_Spare3, SL_DRX_ConfigUC_r17_Sl_Drx_OnDurationTimer_r17_MilliSeconds_Spare2, SL_DRX_ConfigUC_r17_Sl_Drx_OnDurationTimer_r17_MilliSeconds_Spare1},
}

type SL_DRX_ConfigUC_r17 struct {
	Sl_Drx_OnDurationTimer_r17 struct {
		Choice          int
		SubMilliSeconds *int64
		MilliSeconds    *int64
	}
	Sl_Drx_InactivityTimer_r17     int64
	Sl_Drx_HARQ_RTT_Timer1_r17     *int64
	Sl_Drx_HARQ_RTT_Timer2_r17     *int64
	Sl_Drx_RetransmissionTimer_r17 int64
	Sl_Drx_CycleStartOffset_r17    struct {
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
	Sl_Drx_SlotOffset int64
}

func (ie *SL_DRX_ConfigUC_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sLDRXConfigUCR17Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Sl_Drx_HARQ_RTT_Timer1_r17 != nil, ie.Sl_Drx_HARQ_RTT_Timer2_r17 != nil}); err != nil {
		return err
	}

	// 2. sl-drx-onDurationTimer-r17: choice
	{
		choiceEnc := e.NewChoiceEncoder(sL_DRX_ConfigUC_r17SlDrxOnDurationTimerR17Constraints)
		if err := choiceEnc.EncodeChoice(int64(ie.Sl_Drx_OnDurationTimer_r17.Choice), false, nil); err != nil {
			return err
		}
		switch ie.Sl_Drx_OnDurationTimer_r17.Choice {
		case SL_DRX_ConfigUC_r17_Sl_Drx_OnDurationTimer_r17_SubMilliSeconds:
			if err := e.EncodeInteger((*ie.Sl_Drx_OnDurationTimer_r17.SubMilliSeconds), per.Constrained(1, 31)); err != nil {
				return err
			}
		case SL_DRX_ConfigUC_r17_Sl_Drx_OnDurationTimer_r17_MilliSeconds:
			if err := e.EncodeEnumerated((*ie.Sl_Drx_OnDurationTimer_r17.MilliSeconds), sLDRXConfigUCR17SlDrxOnDurationTimerR17MilliSecondsConstraints); err != nil {
				return err
			}
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.Sl_Drx_OnDurationTimer_r17.Choice), Constraint: "index out of range"}
		}
	}

	// 3. sl-drx-InactivityTimer-r17: enumerated
	{
		if err := e.EncodeEnumerated(ie.Sl_Drx_InactivityTimer_r17, sLDRXConfigUCR17SlDrxInactivityTimerR17Constraints); err != nil {
			return err
		}
	}

	// 4. sl-drx-HARQ-RTT-Timer1-r17: enumerated
	{
		if ie.Sl_Drx_HARQ_RTT_Timer1_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Sl_Drx_HARQ_RTT_Timer1_r17, sLDRXConfigUCR17SlDrxHARQRTTTimer1R17Constraints); err != nil {
				return err
			}
		}
	}

	// 5. sl-drx-HARQ-RTT-Timer2-r17: enumerated
	{
		if ie.Sl_Drx_HARQ_RTT_Timer2_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Sl_Drx_HARQ_RTT_Timer2_r17, sLDRXConfigUCR17SlDrxHARQRTTTimer2R17Constraints); err != nil {
				return err
			}
		}
	}

	// 6. sl-drx-RetransmissionTimer-r17: enumerated
	{
		if err := e.EncodeEnumerated(ie.Sl_Drx_RetransmissionTimer_r17, sLDRXConfigUCR17SlDrxRetransmissionTimerR17Constraints); err != nil {
			return err
		}
	}

	// 7. sl-drx-CycleStartOffset-r17: choice
	{
		choiceEnc := e.NewChoiceEncoder(sL_DRX_ConfigUC_r17SlDrxCycleStartOffsetR17Constraints)
		if err := choiceEnc.EncodeChoice(int64(ie.Sl_Drx_CycleStartOffset_r17.Choice), false, nil); err != nil {
			return err
		}
		switch ie.Sl_Drx_CycleStartOffset_r17.Choice {
		case SL_DRX_ConfigUC_r17_Sl_Drx_CycleStartOffset_r17_Ms10:
			if err := e.EncodeInteger((*ie.Sl_Drx_CycleStartOffset_r17.Ms10), per.Constrained(0, 9)); err != nil {
				return err
			}
		case SL_DRX_ConfigUC_r17_Sl_Drx_CycleStartOffset_r17_Ms20:
			if err := e.EncodeInteger((*ie.Sl_Drx_CycleStartOffset_r17.Ms20), per.Constrained(0, 19)); err != nil {
				return err
			}
		case SL_DRX_ConfigUC_r17_Sl_Drx_CycleStartOffset_r17_Ms32:
			if err := e.EncodeInteger((*ie.Sl_Drx_CycleStartOffset_r17.Ms32), per.Constrained(0, 31)); err != nil {
				return err
			}
		case SL_DRX_ConfigUC_r17_Sl_Drx_CycleStartOffset_r17_Ms40:
			if err := e.EncodeInteger((*ie.Sl_Drx_CycleStartOffset_r17.Ms40), per.Constrained(0, 39)); err != nil {
				return err
			}
		case SL_DRX_ConfigUC_r17_Sl_Drx_CycleStartOffset_r17_Ms60:
			if err := e.EncodeInteger((*ie.Sl_Drx_CycleStartOffset_r17.Ms60), per.Constrained(0, 59)); err != nil {
				return err
			}
		case SL_DRX_ConfigUC_r17_Sl_Drx_CycleStartOffset_r17_Ms64:
			if err := e.EncodeInteger((*ie.Sl_Drx_CycleStartOffset_r17.Ms64), per.Constrained(0, 63)); err != nil {
				return err
			}
		case SL_DRX_ConfigUC_r17_Sl_Drx_CycleStartOffset_r17_Ms70:
			if err := e.EncodeInteger((*ie.Sl_Drx_CycleStartOffset_r17.Ms70), per.Constrained(0, 69)); err != nil {
				return err
			}
		case SL_DRX_ConfigUC_r17_Sl_Drx_CycleStartOffset_r17_Ms80:
			if err := e.EncodeInteger((*ie.Sl_Drx_CycleStartOffset_r17.Ms80), per.Constrained(0, 79)); err != nil {
				return err
			}
		case SL_DRX_ConfigUC_r17_Sl_Drx_CycleStartOffset_r17_Ms128:
			if err := e.EncodeInteger((*ie.Sl_Drx_CycleStartOffset_r17.Ms128), per.Constrained(0, 127)); err != nil {
				return err
			}
		case SL_DRX_ConfigUC_r17_Sl_Drx_CycleStartOffset_r17_Ms160:
			if err := e.EncodeInteger((*ie.Sl_Drx_CycleStartOffset_r17.Ms160), per.Constrained(0, 159)); err != nil {
				return err
			}
		case SL_DRX_ConfigUC_r17_Sl_Drx_CycleStartOffset_r17_Ms256:
			if err := e.EncodeInteger((*ie.Sl_Drx_CycleStartOffset_r17.Ms256), per.Constrained(0, 255)); err != nil {
				return err
			}
		case SL_DRX_ConfigUC_r17_Sl_Drx_CycleStartOffset_r17_Ms320:
			if err := e.EncodeInteger((*ie.Sl_Drx_CycleStartOffset_r17.Ms320), per.Constrained(0, 319)); err != nil {
				return err
			}
		case SL_DRX_ConfigUC_r17_Sl_Drx_CycleStartOffset_r17_Ms512:
			if err := e.EncodeInteger((*ie.Sl_Drx_CycleStartOffset_r17.Ms512), per.Constrained(0, 511)); err != nil {
				return err
			}
		case SL_DRX_ConfigUC_r17_Sl_Drx_CycleStartOffset_r17_Ms640:
			if err := e.EncodeInteger((*ie.Sl_Drx_CycleStartOffset_r17.Ms640), per.Constrained(0, 639)); err != nil {
				return err
			}
		case SL_DRX_ConfigUC_r17_Sl_Drx_CycleStartOffset_r17_Ms1024:
			if err := e.EncodeInteger((*ie.Sl_Drx_CycleStartOffset_r17.Ms1024), per.Constrained(0, 1023)); err != nil {
				return err
			}
		case SL_DRX_ConfigUC_r17_Sl_Drx_CycleStartOffset_r17_Ms1280:
			if err := e.EncodeInteger((*ie.Sl_Drx_CycleStartOffset_r17.Ms1280), per.Constrained(0, 1279)); err != nil {
				return err
			}
		case SL_DRX_ConfigUC_r17_Sl_Drx_CycleStartOffset_r17_Ms2048:
			if err := e.EncodeInteger((*ie.Sl_Drx_CycleStartOffset_r17.Ms2048), per.Constrained(0, 2047)); err != nil {
				return err
			}
		case SL_DRX_ConfigUC_r17_Sl_Drx_CycleStartOffset_r17_Ms2560:
			if err := e.EncodeInteger((*ie.Sl_Drx_CycleStartOffset_r17.Ms2560), per.Constrained(0, 2559)); err != nil {
				return err
			}
		case SL_DRX_ConfigUC_r17_Sl_Drx_CycleStartOffset_r17_Ms5120:
			if err := e.EncodeInteger((*ie.Sl_Drx_CycleStartOffset_r17.Ms5120), per.Constrained(0, 5119)); err != nil {
				return err
			}
		case SL_DRX_ConfigUC_r17_Sl_Drx_CycleStartOffset_r17_Ms10240:
			if err := e.EncodeInteger((*ie.Sl_Drx_CycleStartOffset_r17.Ms10240), per.Constrained(0, 10239)); err != nil {
				return err
			}
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.Sl_Drx_CycleStartOffset_r17.Choice), Constraint: "index out of range"}
		}
	}

	// 8. sl-drx-SlotOffset: integer
	{
		if err := e.EncodeInteger(ie.Sl_Drx_SlotOffset, sLDRXConfigUCR17SlDrxSlotOffsetConstraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *SL_DRX_ConfigUC_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sLDRXConfigUCR17Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. sl-drx-onDurationTimer-r17: choice
	{
		choiceDec := d.NewChoiceDecoder(sL_DRX_ConfigUC_r17SlDrxOnDurationTimerR17Constraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.Sl_Drx_OnDurationTimer_r17.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case SL_DRX_ConfigUC_r17_Sl_Drx_OnDurationTimer_r17_SubMilliSeconds:
			ie.Sl_Drx_OnDurationTimer_r17.SubMilliSeconds = new(int64)
			v, err := d.DecodeInteger(per.Constrained(1, 31))
			if err != nil {
				return err
			}
			(*ie.Sl_Drx_OnDurationTimer_r17.SubMilliSeconds) = v
		case SL_DRX_ConfigUC_r17_Sl_Drx_OnDurationTimer_r17_MilliSeconds:
			ie.Sl_Drx_OnDurationTimer_r17.MilliSeconds = new(int64)
			v, err := d.DecodeEnumerated(sLDRXConfigUCR17SlDrxOnDurationTimerR17MilliSecondsConstraints)
			if err != nil {
				return err
			}
			(*ie.Sl_Drx_OnDurationTimer_r17.MilliSeconds) = v
		}
	}

	// 3. sl-drx-InactivityTimer-r17: enumerated
	{
		v1, err := d.DecodeEnumerated(sLDRXConfigUCR17SlDrxInactivityTimerR17Constraints)
		if err != nil {
			return err
		}
		ie.Sl_Drx_InactivityTimer_r17 = v1
	}

	// 4. sl-drx-HARQ-RTT-Timer1-r17: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(sLDRXConfigUCR17SlDrxHARQRTTTimer1R17Constraints)
			if err != nil {
				return err
			}
			ie.Sl_Drx_HARQ_RTT_Timer1_r17 = &idx
		}
	}

	// 5. sl-drx-HARQ-RTT-Timer2-r17: enumerated
	{
		if seq.IsComponentPresent(3) {
			idx, err := d.DecodeEnumerated(sLDRXConfigUCR17SlDrxHARQRTTTimer2R17Constraints)
			if err != nil {
				return err
			}
			ie.Sl_Drx_HARQ_RTT_Timer2_r17 = &idx
		}
	}

	// 6. sl-drx-RetransmissionTimer-r17: enumerated
	{
		v4, err := d.DecodeEnumerated(sLDRXConfigUCR17SlDrxRetransmissionTimerR17Constraints)
		if err != nil {
			return err
		}
		ie.Sl_Drx_RetransmissionTimer_r17 = v4
	}

	// 7. sl-drx-CycleStartOffset-r17: choice
	{
		choiceDec := d.NewChoiceDecoder(sL_DRX_ConfigUC_r17SlDrxCycleStartOffsetR17Constraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.Sl_Drx_CycleStartOffset_r17.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case SL_DRX_ConfigUC_r17_Sl_Drx_CycleStartOffset_r17_Ms10:
			ie.Sl_Drx_CycleStartOffset_r17.Ms10 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 9))
			if err != nil {
				return err
			}
			(*ie.Sl_Drx_CycleStartOffset_r17.Ms10) = v
		case SL_DRX_ConfigUC_r17_Sl_Drx_CycleStartOffset_r17_Ms20:
			ie.Sl_Drx_CycleStartOffset_r17.Ms20 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 19))
			if err != nil {
				return err
			}
			(*ie.Sl_Drx_CycleStartOffset_r17.Ms20) = v
		case SL_DRX_ConfigUC_r17_Sl_Drx_CycleStartOffset_r17_Ms32:
			ie.Sl_Drx_CycleStartOffset_r17.Ms32 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 31))
			if err != nil {
				return err
			}
			(*ie.Sl_Drx_CycleStartOffset_r17.Ms32) = v
		case SL_DRX_ConfigUC_r17_Sl_Drx_CycleStartOffset_r17_Ms40:
			ie.Sl_Drx_CycleStartOffset_r17.Ms40 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 39))
			if err != nil {
				return err
			}
			(*ie.Sl_Drx_CycleStartOffset_r17.Ms40) = v
		case SL_DRX_ConfigUC_r17_Sl_Drx_CycleStartOffset_r17_Ms60:
			ie.Sl_Drx_CycleStartOffset_r17.Ms60 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 59))
			if err != nil {
				return err
			}
			(*ie.Sl_Drx_CycleStartOffset_r17.Ms60) = v
		case SL_DRX_ConfigUC_r17_Sl_Drx_CycleStartOffset_r17_Ms64:
			ie.Sl_Drx_CycleStartOffset_r17.Ms64 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 63))
			if err != nil {
				return err
			}
			(*ie.Sl_Drx_CycleStartOffset_r17.Ms64) = v
		case SL_DRX_ConfigUC_r17_Sl_Drx_CycleStartOffset_r17_Ms70:
			ie.Sl_Drx_CycleStartOffset_r17.Ms70 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 69))
			if err != nil {
				return err
			}
			(*ie.Sl_Drx_CycleStartOffset_r17.Ms70) = v
		case SL_DRX_ConfigUC_r17_Sl_Drx_CycleStartOffset_r17_Ms80:
			ie.Sl_Drx_CycleStartOffset_r17.Ms80 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 79))
			if err != nil {
				return err
			}
			(*ie.Sl_Drx_CycleStartOffset_r17.Ms80) = v
		case SL_DRX_ConfigUC_r17_Sl_Drx_CycleStartOffset_r17_Ms128:
			ie.Sl_Drx_CycleStartOffset_r17.Ms128 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 127))
			if err != nil {
				return err
			}
			(*ie.Sl_Drx_CycleStartOffset_r17.Ms128) = v
		case SL_DRX_ConfigUC_r17_Sl_Drx_CycleStartOffset_r17_Ms160:
			ie.Sl_Drx_CycleStartOffset_r17.Ms160 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 159))
			if err != nil {
				return err
			}
			(*ie.Sl_Drx_CycleStartOffset_r17.Ms160) = v
		case SL_DRX_ConfigUC_r17_Sl_Drx_CycleStartOffset_r17_Ms256:
			ie.Sl_Drx_CycleStartOffset_r17.Ms256 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 255))
			if err != nil {
				return err
			}
			(*ie.Sl_Drx_CycleStartOffset_r17.Ms256) = v
		case SL_DRX_ConfigUC_r17_Sl_Drx_CycleStartOffset_r17_Ms320:
			ie.Sl_Drx_CycleStartOffset_r17.Ms320 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 319))
			if err != nil {
				return err
			}
			(*ie.Sl_Drx_CycleStartOffset_r17.Ms320) = v
		case SL_DRX_ConfigUC_r17_Sl_Drx_CycleStartOffset_r17_Ms512:
			ie.Sl_Drx_CycleStartOffset_r17.Ms512 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 511))
			if err != nil {
				return err
			}
			(*ie.Sl_Drx_CycleStartOffset_r17.Ms512) = v
		case SL_DRX_ConfigUC_r17_Sl_Drx_CycleStartOffset_r17_Ms640:
			ie.Sl_Drx_CycleStartOffset_r17.Ms640 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 639))
			if err != nil {
				return err
			}
			(*ie.Sl_Drx_CycleStartOffset_r17.Ms640) = v
		case SL_DRX_ConfigUC_r17_Sl_Drx_CycleStartOffset_r17_Ms1024:
			ie.Sl_Drx_CycleStartOffset_r17.Ms1024 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 1023))
			if err != nil {
				return err
			}
			(*ie.Sl_Drx_CycleStartOffset_r17.Ms1024) = v
		case SL_DRX_ConfigUC_r17_Sl_Drx_CycleStartOffset_r17_Ms1280:
			ie.Sl_Drx_CycleStartOffset_r17.Ms1280 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 1279))
			if err != nil {
				return err
			}
			(*ie.Sl_Drx_CycleStartOffset_r17.Ms1280) = v
		case SL_DRX_ConfigUC_r17_Sl_Drx_CycleStartOffset_r17_Ms2048:
			ie.Sl_Drx_CycleStartOffset_r17.Ms2048 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 2047))
			if err != nil {
				return err
			}
			(*ie.Sl_Drx_CycleStartOffset_r17.Ms2048) = v
		case SL_DRX_ConfigUC_r17_Sl_Drx_CycleStartOffset_r17_Ms2560:
			ie.Sl_Drx_CycleStartOffset_r17.Ms2560 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 2559))
			if err != nil {
				return err
			}
			(*ie.Sl_Drx_CycleStartOffset_r17.Ms2560) = v
		case SL_DRX_ConfigUC_r17_Sl_Drx_CycleStartOffset_r17_Ms5120:
			ie.Sl_Drx_CycleStartOffset_r17.Ms5120 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 5119))
			if err != nil {
				return err
			}
			(*ie.Sl_Drx_CycleStartOffset_r17.Ms5120) = v
		case SL_DRX_ConfigUC_r17_Sl_Drx_CycleStartOffset_r17_Ms10240:
			ie.Sl_Drx_CycleStartOffset_r17.Ms10240 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 10239))
			if err != nil {
				return err
			}
			(*ie.Sl_Drx_CycleStartOffset_r17.Ms10240) = v
		}
	}

	// 8. sl-drx-SlotOffset: integer
	{
		v6, err := d.DecodeInteger(sLDRXConfigUCR17SlDrxSlotOffsetConstraints)
		if err != nil {
			return err
		}
		ie.Sl_Drx_SlotOffset = v6
	}

	return nil
}
