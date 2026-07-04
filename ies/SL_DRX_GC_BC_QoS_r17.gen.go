// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SL-DRX-GC-BC-QoS-r17 (line 27139).

var sLDRXGCBCQoSR17Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-DRX-GC-BC-MappedQoS-FlowList-r17", Optional: true},
		{Name: "sl-DRX-GC-BC-OnDurationTimer-r17"},
		{Name: "sl-DRX-GC-InactivityTimer-r17"},
		{Name: "sl-DRX-GC-BC-Cycle-r17"},
	},
}

var sLDRXGCBCQoSR17SlDRXGCBCMappedQoSFlowListR17Constraints = per.SizeRange(1, common.MaxNrofSL_QFIs_r16)

var sL_DRX_GC_BC_QoS_r17SlDRXGCBCOnDurationTimerR17Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "subMilliSeconds"},
		{Name: "milliSeconds"},
	},
}

const (
	SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_BC_OnDurationTimer_r17_SubMilliSeconds = 0
	SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_BC_OnDurationTimer_r17_MilliSeconds    = 1
)

const (
	SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_InactivityTimer_r17_Ms0    = 0
	SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_InactivityTimer_r17_Ms1    = 1
	SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_InactivityTimer_r17_Ms2    = 2
	SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_InactivityTimer_r17_Ms3    = 3
	SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_InactivityTimer_r17_Ms4    = 4
	SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_InactivityTimer_r17_Ms5    = 5
	SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_InactivityTimer_r17_Ms6    = 6
	SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_InactivityTimer_r17_Ms8    = 7
	SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_InactivityTimer_r17_Ms10   = 8
	SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_InactivityTimer_r17_Ms20   = 9
	SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_InactivityTimer_r17_Ms30   = 10
	SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_InactivityTimer_r17_Ms40   = 11
	SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_InactivityTimer_r17_Ms50   = 12
	SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_InactivityTimer_r17_Ms60   = 13
	SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_InactivityTimer_r17_Ms80   = 14
	SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_InactivityTimer_r17_Ms100  = 15
	SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_InactivityTimer_r17_Ms200  = 16
	SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_InactivityTimer_r17_Ms300  = 17
	SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_InactivityTimer_r17_Ms500  = 18
	SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_InactivityTimer_r17_Ms750  = 19
	SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_InactivityTimer_r17_Ms1280 = 20
	SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_InactivityTimer_r17_Ms1920 = 21
	SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_InactivityTimer_r17_Ms2560 = 22
	SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_InactivityTimer_r17_Spare9 = 23
	SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_InactivityTimer_r17_Spare8 = 24
	SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_InactivityTimer_r17_Spare7 = 25
	SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_InactivityTimer_r17_Spare6 = 26
	SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_InactivityTimer_r17_Spare5 = 27
	SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_InactivityTimer_r17_Spare4 = 28
	SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_InactivityTimer_r17_Spare3 = 29
	SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_InactivityTimer_r17_Spare2 = 30
	SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_InactivityTimer_r17_Spare1 = 31
)

var sLDRXGCBCQoSR17SlDRXGCInactivityTimerR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_InactivityTimer_r17_Ms0, SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_InactivityTimer_r17_Ms1, SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_InactivityTimer_r17_Ms2, SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_InactivityTimer_r17_Ms3, SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_InactivityTimer_r17_Ms4, SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_InactivityTimer_r17_Ms5, SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_InactivityTimer_r17_Ms6, SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_InactivityTimer_r17_Ms8, SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_InactivityTimer_r17_Ms10, SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_InactivityTimer_r17_Ms20, SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_InactivityTimer_r17_Ms30, SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_InactivityTimer_r17_Ms40, SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_InactivityTimer_r17_Ms50, SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_InactivityTimer_r17_Ms60, SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_InactivityTimer_r17_Ms80, SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_InactivityTimer_r17_Ms100, SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_InactivityTimer_r17_Ms200, SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_InactivityTimer_r17_Ms300, SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_InactivityTimer_r17_Ms500, SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_InactivityTimer_r17_Ms750, SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_InactivityTimer_r17_Ms1280, SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_InactivityTimer_r17_Ms1920, SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_InactivityTimer_r17_Ms2560, SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_InactivityTimer_r17_Spare9, SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_InactivityTimer_r17_Spare8, SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_InactivityTimer_r17_Spare7, SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_InactivityTimer_r17_Spare6, SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_InactivityTimer_r17_Spare5, SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_InactivityTimer_r17_Spare4, SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_InactivityTimer_r17_Spare3, SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_InactivityTimer_r17_Spare2, SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_InactivityTimer_r17_Spare1},
}

const (
	SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_BC_Cycle_r17_Ms10    = 0
	SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_BC_Cycle_r17_Ms20    = 1
	SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_BC_Cycle_r17_Ms32    = 2
	SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_BC_Cycle_r17_Ms40    = 3
	SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_BC_Cycle_r17_Ms60    = 4
	SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_BC_Cycle_r17_Ms64    = 5
	SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_BC_Cycle_r17_Ms70    = 6
	SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_BC_Cycle_r17_Ms80    = 7
	SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_BC_Cycle_r17_Ms128   = 8
	SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_BC_Cycle_r17_Ms160   = 9
	SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_BC_Cycle_r17_Ms256   = 10
	SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_BC_Cycle_r17_Ms320   = 11
	SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_BC_Cycle_r17_Ms512   = 12
	SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_BC_Cycle_r17_Ms640   = 13
	SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_BC_Cycle_r17_Ms1024  = 14
	SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_BC_Cycle_r17_Ms1280  = 15
	SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_BC_Cycle_r17_Ms2048  = 16
	SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_BC_Cycle_r17_Ms2560  = 17
	SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_BC_Cycle_r17_Ms5120  = 18
	SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_BC_Cycle_r17_Ms10240 = 19
	SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_BC_Cycle_r17_Spare12 = 20
	SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_BC_Cycle_r17_Spare11 = 21
	SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_BC_Cycle_r17_Spare10 = 22
	SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_BC_Cycle_r17_Spare9  = 23
	SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_BC_Cycle_r17_Spare8  = 24
	SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_BC_Cycle_r17_Spare7  = 25
	SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_BC_Cycle_r17_Spare6  = 26
	SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_BC_Cycle_r17_Spare5  = 27
	SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_BC_Cycle_r17_Spare4  = 28
	SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_BC_Cycle_r17_Spare3  = 29
	SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_BC_Cycle_r17_Spare2  = 30
	SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_BC_Cycle_r17_Spare1  = 31
)

var sLDRXGCBCQoSR17SlDRXGCBCCycleR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_BC_Cycle_r17_Ms10, SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_BC_Cycle_r17_Ms20, SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_BC_Cycle_r17_Ms32, SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_BC_Cycle_r17_Ms40, SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_BC_Cycle_r17_Ms60, SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_BC_Cycle_r17_Ms64, SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_BC_Cycle_r17_Ms70, SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_BC_Cycle_r17_Ms80, SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_BC_Cycle_r17_Ms128, SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_BC_Cycle_r17_Ms160, SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_BC_Cycle_r17_Ms256, SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_BC_Cycle_r17_Ms320, SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_BC_Cycle_r17_Ms512, SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_BC_Cycle_r17_Ms640, SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_BC_Cycle_r17_Ms1024, SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_BC_Cycle_r17_Ms1280, SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_BC_Cycle_r17_Ms2048, SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_BC_Cycle_r17_Ms2560, SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_BC_Cycle_r17_Ms5120, SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_BC_Cycle_r17_Ms10240, SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_BC_Cycle_r17_Spare12, SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_BC_Cycle_r17_Spare11, SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_BC_Cycle_r17_Spare10, SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_BC_Cycle_r17_Spare9, SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_BC_Cycle_r17_Spare8, SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_BC_Cycle_r17_Spare7, SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_BC_Cycle_r17_Spare6, SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_BC_Cycle_r17_Spare5, SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_BC_Cycle_r17_Spare4, SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_BC_Cycle_r17_Spare3, SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_BC_Cycle_r17_Spare2, SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_BC_Cycle_r17_Spare1},
}

const (
	SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_BC_OnDurationTimer_r17_MilliSeconds_Ms1    = 0
	SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_BC_OnDurationTimer_r17_MilliSeconds_Ms2    = 1
	SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_BC_OnDurationTimer_r17_MilliSeconds_Ms3    = 2
	SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_BC_OnDurationTimer_r17_MilliSeconds_Ms4    = 3
	SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_BC_OnDurationTimer_r17_MilliSeconds_Ms5    = 4
	SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_BC_OnDurationTimer_r17_MilliSeconds_Ms6    = 5
	SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_BC_OnDurationTimer_r17_MilliSeconds_Ms8    = 6
	SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_BC_OnDurationTimer_r17_MilliSeconds_Ms10   = 7
	SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_BC_OnDurationTimer_r17_MilliSeconds_Ms20   = 8
	SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_BC_OnDurationTimer_r17_MilliSeconds_Ms30   = 9
	SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_BC_OnDurationTimer_r17_MilliSeconds_Ms40   = 10
	SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_BC_OnDurationTimer_r17_MilliSeconds_Ms50   = 11
	SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_BC_OnDurationTimer_r17_MilliSeconds_Ms60   = 12
	SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_BC_OnDurationTimer_r17_MilliSeconds_Ms80   = 13
	SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_BC_OnDurationTimer_r17_MilliSeconds_Ms100  = 14
	SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_BC_OnDurationTimer_r17_MilliSeconds_Ms200  = 15
	SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_BC_OnDurationTimer_r17_MilliSeconds_Ms300  = 16
	SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_BC_OnDurationTimer_r17_MilliSeconds_Ms400  = 17
	SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_BC_OnDurationTimer_r17_MilliSeconds_Ms500  = 18
	SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_BC_OnDurationTimer_r17_MilliSeconds_Ms600  = 19
	SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_BC_OnDurationTimer_r17_MilliSeconds_Ms800  = 20
	SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_BC_OnDurationTimer_r17_MilliSeconds_Ms1000 = 21
	SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_BC_OnDurationTimer_r17_MilliSeconds_Ms1200 = 22
	SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_BC_OnDurationTimer_r17_MilliSeconds_Ms1600 = 23
	SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_BC_OnDurationTimer_r17_MilliSeconds_Spare8 = 24
	SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_BC_OnDurationTimer_r17_MilliSeconds_Spare7 = 25
	SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_BC_OnDurationTimer_r17_MilliSeconds_Spare6 = 26
	SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_BC_OnDurationTimer_r17_MilliSeconds_Spare5 = 27
	SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_BC_OnDurationTimer_r17_MilliSeconds_Spare4 = 28
	SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_BC_OnDurationTimer_r17_MilliSeconds_Spare3 = 29
	SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_BC_OnDurationTimer_r17_MilliSeconds_Spare2 = 30
	SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_BC_OnDurationTimer_r17_MilliSeconds_Spare1 = 31
)

var sLDRXGCBCQoSR17SlDRXGCBCOnDurationTimerR17MilliSecondsConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_BC_OnDurationTimer_r17_MilliSeconds_Ms1, SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_BC_OnDurationTimer_r17_MilliSeconds_Ms2, SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_BC_OnDurationTimer_r17_MilliSeconds_Ms3, SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_BC_OnDurationTimer_r17_MilliSeconds_Ms4, SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_BC_OnDurationTimer_r17_MilliSeconds_Ms5, SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_BC_OnDurationTimer_r17_MilliSeconds_Ms6, SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_BC_OnDurationTimer_r17_MilliSeconds_Ms8, SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_BC_OnDurationTimer_r17_MilliSeconds_Ms10, SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_BC_OnDurationTimer_r17_MilliSeconds_Ms20, SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_BC_OnDurationTimer_r17_MilliSeconds_Ms30, SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_BC_OnDurationTimer_r17_MilliSeconds_Ms40, SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_BC_OnDurationTimer_r17_MilliSeconds_Ms50, SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_BC_OnDurationTimer_r17_MilliSeconds_Ms60, SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_BC_OnDurationTimer_r17_MilliSeconds_Ms80, SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_BC_OnDurationTimer_r17_MilliSeconds_Ms100, SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_BC_OnDurationTimer_r17_MilliSeconds_Ms200, SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_BC_OnDurationTimer_r17_MilliSeconds_Ms300, SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_BC_OnDurationTimer_r17_MilliSeconds_Ms400, SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_BC_OnDurationTimer_r17_MilliSeconds_Ms500, SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_BC_OnDurationTimer_r17_MilliSeconds_Ms600, SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_BC_OnDurationTimer_r17_MilliSeconds_Ms800, SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_BC_OnDurationTimer_r17_MilliSeconds_Ms1000, SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_BC_OnDurationTimer_r17_MilliSeconds_Ms1200, SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_BC_OnDurationTimer_r17_MilliSeconds_Ms1600, SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_BC_OnDurationTimer_r17_MilliSeconds_Spare8, SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_BC_OnDurationTimer_r17_MilliSeconds_Spare7, SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_BC_OnDurationTimer_r17_MilliSeconds_Spare6, SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_BC_OnDurationTimer_r17_MilliSeconds_Spare5, SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_BC_OnDurationTimer_r17_MilliSeconds_Spare4, SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_BC_OnDurationTimer_r17_MilliSeconds_Spare3, SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_BC_OnDurationTimer_r17_MilliSeconds_Spare2, SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_BC_OnDurationTimer_r17_MilliSeconds_Spare1},
}

type SL_DRX_GC_BC_QoS_r17 struct {
	Sl_DRX_GC_BC_MappedQoS_FlowList_r17 []SL_QoS_Profile_r16
	Sl_DRX_GC_BC_OnDurationTimer_r17    struct {
		Choice          int
		SubMilliSeconds *int64
		MilliSeconds    *int64
	}
	Sl_DRX_GC_InactivityTimer_r17 int64
	Sl_DRX_GC_BC_Cycle_r17        int64
}

func (ie *SL_DRX_GC_BC_QoS_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sLDRXGCBCQoSR17Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Sl_DRX_GC_BC_MappedQoS_FlowList_r17 != nil}); err != nil {
		return err
	}

	// 3. sl-DRX-GC-BC-MappedQoS-FlowList-r17: sequence-of
	{
		if ie.Sl_DRX_GC_BC_MappedQoS_FlowList_r17 != nil {
			seqOf := e.NewSequenceOfEncoder(sLDRXGCBCQoSR17SlDRXGCBCMappedQoSFlowListR17Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Sl_DRX_GC_BC_MappedQoS_FlowList_r17))); err != nil {
				return err
			}
			for i := range ie.Sl_DRX_GC_BC_MappedQoS_FlowList_r17 {
				if err := ie.Sl_DRX_GC_BC_MappedQoS_FlowList_r17[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 4. sl-DRX-GC-BC-OnDurationTimer-r17: choice
	{
		choiceEnc := e.NewChoiceEncoder(sL_DRX_GC_BC_QoS_r17SlDRXGCBCOnDurationTimerR17Constraints)
		if err := choiceEnc.EncodeChoice(int64(ie.Sl_DRX_GC_BC_OnDurationTimer_r17.Choice), false, nil); err != nil {
			return err
		}
		switch ie.Sl_DRX_GC_BC_OnDurationTimer_r17.Choice {
		case SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_BC_OnDurationTimer_r17_SubMilliSeconds:
			if err := e.EncodeInteger((*ie.Sl_DRX_GC_BC_OnDurationTimer_r17.SubMilliSeconds), per.Constrained(1, 31)); err != nil {
				return err
			}
		case SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_BC_OnDurationTimer_r17_MilliSeconds:
			if err := e.EncodeEnumerated((*ie.Sl_DRX_GC_BC_OnDurationTimer_r17.MilliSeconds), sLDRXGCBCQoSR17SlDRXGCBCOnDurationTimerR17MilliSecondsConstraints); err != nil {
				return err
			}
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.Sl_DRX_GC_BC_OnDurationTimer_r17.Choice), Constraint: "index out of range"}
		}
	}

	// 5. sl-DRX-GC-InactivityTimer-r17: enumerated
	{
		if err := e.EncodeEnumerated(ie.Sl_DRX_GC_InactivityTimer_r17, sLDRXGCBCQoSR17SlDRXGCInactivityTimerR17Constraints); err != nil {
			return err
		}
	}

	// 6. sl-DRX-GC-BC-Cycle-r17: enumerated
	{
		if err := e.EncodeEnumerated(ie.Sl_DRX_GC_BC_Cycle_r17, sLDRXGCBCQoSR17SlDRXGCBCCycleR17Constraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *SL_DRX_GC_BC_QoS_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sLDRXGCBCQoSR17Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. sl-DRX-GC-BC-MappedQoS-FlowList-r17: sequence-of
	{
		if seq.IsComponentPresent(0) {
			seqOf := d.NewSequenceOfDecoder(sLDRXGCBCQoSR17SlDRXGCBCMappedQoSFlowListR17Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Sl_DRX_GC_BC_MappedQoS_FlowList_r17 = make([]SL_QoS_Profile_r16, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Sl_DRX_GC_BC_MappedQoS_FlowList_r17[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 4. sl-DRX-GC-BC-OnDurationTimer-r17: choice
	{
		choiceDec := d.NewChoiceDecoder(sL_DRX_GC_BC_QoS_r17SlDRXGCBCOnDurationTimerR17Constraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.Sl_DRX_GC_BC_OnDurationTimer_r17.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_BC_OnDurationTimer_r17_SubMilliSeconds:
			ie.Sl_DRX_GC_BC_OnDurationTimer_r17.SubMilliSeconds = new(int64)
			v, err := d.DecodeInteger(per.Constrained(1, 31))
			if err != nil {
				return err
			}
			(*ie.Sl_DRX_GC_BC_OnDurationTimer_r17.SubMilliSeconds) = v
		case SL_DRX_GC_BC_QoS_r17_Sl_DRX_GC_BC_OnDurationTimer_r17_MilliSeconds:
			ie.Sl_DRX_GC_BC_OnDurationTimer_r17.MilliSeconds = new(int64)
			v, err := d.DecodeEnumerated(sLDRXGCBCQoSR17SlDRXGCBCOnDurationTimerR17MilliSecondsConstraints)
			if err != nil {
				return err
			}
			(*ie.Sl_DRX_GC_BC_OnDurationTimer_r17.MilliSeconds) = v
		}
	}

	// 5. sl-DRX-GC-InactivityTimer-r17: enumerated
	{
		v2, err := d.DecodeEnumerated(sLDRXGCBCQoSR17SlDRXGCInactivityTimerR17Constraints)
		if err != nil {
			return err
		}
		ie.Sl_DRX_GC_InactivityTimer_r17 = v2
	}

	// 6. sl-DRX-GC-BC-Cycle-r17: enumerated
	{
		v3, err := d.DecodeEnumerated(sLDRXGCBCQoSR17SlDRXGCBCCycleR17Constraints)
		if err != nil {
			return err
		}
		ie.Sl_DRX_GC_BC_Cycle_r17 = v3
	}

	return nil
}
