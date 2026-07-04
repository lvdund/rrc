// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: CellDTX-DRX-Config-r18 (line 5557).

var cellDTXDRXConfigR18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "cellDTX-DRX-onDurationTimer-r18"},
		{Name: "cellDTX-DRX-CycleStartOffset-r18"},
		{Name: "cellDTX-DRX-SlotOffset-r18"},
		{Name: "cellDTX-DRX-ConfigType-r18"},
		{Name: "cellDTX-DRX-ActivationStatus-r18", Optional: true},
	},
}

var cellDTX_DRX_Config_r18CellDTXDRXOnDurationTimerR18Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "subMilliSeconds"},
		{Name: "milliSeconds"},
	},
}

const (
	CellDTX_DRX_Config_r18_CellDTX_DRX_OnDurationTimer_r18_SubMilliSeconds = 0
	CellDTX_DRX_Config_r18_CellDTX_DRX_OnDurationTimer_r18_MilliSeconds    = 1
)

var cellDTX_DRX_Config_r18CellDTXDRXCycleStartOffsetR18Constraints = per.ChoiceConstraints{
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
	CellDTX_DRX_Config_r18_CellDTX_DRX_CycleStartOffset_r18_Ms10    = 0
	CellDTX_DRX_Config_r18_CellDTX_DRX_CycleStartOffset_r18_Ms20    = 1
	CellDTX_DRX_Config_r18_CellDTX_DRX_CycleStartOffset_r18_Ms32    = 2
	CellDTX_DRX_Config_r18_CellDTX_DRX_CycleStartOffset_r18_Ms40    = 3
	CellDTX_DRX_Config_r18_CellDTX_DRX_CycleStartOffset_r18_Ms60    = 4
	CellDTX_DRX_Config_r18_CellDTX_DRX_CycleStartOffset_r18_Ms64    = 5
	CellDTX_DRX_Config_r18_CellDTX_DRX_CycleStartOffset_r18_Ms70    = 6
	CellDTX_DRX_Config_r18_CellDTX_DRX_CycleStartOffset_r18_Ms80    = 7
	CellDTX_DRX_Config_r18_CellDTX_DRX_CycleStartOffset_r18_Ms128   = 8
	CellDTX_DRX_Config_r18_CellDTX_DRX_CycleStartOffset_r18_Ms160   = 9
	CellDTX_DRX_Config_r18_CellDTX_DRX_CycleStartOffset_r18_Ms256   = 10
	CellDTX_DRX_Config_r18_CellDTX_DRX_CycleStartOffset_r18_Ms320   = 11
	CellDTX_DRX_Config_r18_CellDTX_DRX_CycleStartOffset_r18_Ms512   = 12
	CellDTX_DRX_Config_r18_CellDTX_DRX_CycleStartOffset_r18_Ms640   = 13
	CellDTX_DRX_Config_r18_CellDTX_DRX_CycleStartOffset_r18_Ms1024  = 14
	CellDTX_DRX_Config_r18_CellDTX_DRX_CycleStartOffset_r18_Ms1280  = 15
	CellDTX_DRX_Config_r18_CellDTX_DRX_CycleStartOffset_r18_Ms2048  = 16
	CellDTX_DRX_Config_r18_CellDTX_DRX_CycleStartOffset_r18_Ms2560  = 17
	CellDTX_DRX_Config_r18_CellDTX_DRX_CycleStartOffset_r18_Ms5120  = 18
	CellDTX_DRX_Config_r18_CellDTX_DRX_CycleStartOffset_r18_Ms10240 = 19
)

var cellDTXDRXConfigR18CellDTXDRXSlotOffsetR18Constraints = per.Constrained(0, 31)

const (
	CellDTX_DRX_Config_r18_CellDTX_DRX_ConfigType_r18_Dtx    = 0
	CellDTX_DRX_Config_r18_CellDTX_DRX_ConfigType_r18_Drx    = 1
	CellDTX_DRX_Config_r18_CellDTX_DRX_ConfigType_r18_Dtxdrx = 2
)

var cellDTXDRXConfigR18CellDTXDRXConfigTypeR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CellDTX_DRX_Config_r18_CellDTX_DRX_ConfigType_r18_Dtx, CellDTX_DRX_Config_r18_CellDTX_DRX_ConfigType_r18_Drx, CellDTX_DRX_Config_r18_CellDTX_DRX_ConfigType_r18_Dtxdrx},
}

const (
	CellDTX_DRX_Config_r18_CellDTX_DRX_ActivationStatus_r18_Activated   = 0
	CellDTX_DRX_Config_r18_CellDTX_DRX_ActivationStatus_r18_Deactivated = 1
)

var cellDTXDRXConfigR18CellDTXDRXActivationStatusR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CellDTX_DRX_Config_r18_CellDTX_DRX_ActivationStatus_r18_Activated, CellDTX_DRX_Config_r18_CellDTX_DRX_ActivationStatus_r18_Deactivated},
}

const (
	CellDTX_DRX_Config_r18_CellDTX_DRX_OnDurationTimer_r18_MilliSeconds_Ms1    = 0
	CellDTX_DRX_Config_r18_CellDTX_DRX_OnDurationTimer_r18_MilliSeconds_Ms2    = 1
	CellDTX_DRX_Config_r18_CellDTX_DRX_OnDurationTimer_r18_MilliSeconds_Ms3    = 2
	CellDTX_DRX_Config_r18_CellDTX_DRX_OnDurationTimer_r18_MilliSeconds_Ms4    = 3
	CellDTX_DRX_Config_r18_CellDTX_DRX_OnDurationTimer_r18_MilliSeconds_Ms5    = 4
	CellDTX_DRX_Config_r18_CellDTX_DRX_OnDurationTimer_r18_MilliSeconds_Ms6    = 5
	CellDTX_DRX_Config_r18_CellDTX_DRX_OnDurationTimer_r18_MilliSeconds_Ms8    = 6
	CellDTX_DRX_Config_r18_CellDTX_DRX_OnDurationTimer_r18_MilliSeconds_Ms10   = 7
	CellDTX_DRX_Config_r18_CellDTX_DRX_OnDurationTimer_r18_MilliSeconds_Ms20   = 8
	CellDTX_DRX_Config_r18_CellDTX_DRX_OnDurationTimer_r18_MilliSeconds_Ms30   = 9
	CellDTX_DRX_Config_r18_CellDTX_DRX_OnDurationTimer_r18_MilliSeconds_Ms40   = 10
	CellDTX_DRX_Config_r18_CellDTX_DRX_OnDurationTimer_r18_MilliSeconds_Ms50   = 11
	CellDTX_DRX_Config_r18_CellDTX_DRX_OnDurationTimer_r18_MilliSeconds_Ms60   = 12
	CellDTX_DRX_Config_r18_CellDTX_DRX_OnDurationTimer_r18_MilliSeconds_Ms80   = 13
	CellDTX_DRX_Config_r18_CellDTX_DRX_OnDurationTimer_r18_MilliSeconds_Ms100  = 14
	CellDTX_DRX_Config_r18_CellDTX_DRX_OnDurationTimer_r18_MilliSeconds_Ms200  = 15
	CellDTX_DRX_Config_r18_CellDTX_DRX_OnDurationTimer_r18_MilliSeconds_Ms300  = 16
	CellDTX_DRX_Config_r18_CellDTX_DRX_OnDurationTimer_r18_MilliSeconds_Ms400  = 17
	CellDTX_DRX_Config_r18_CellDTX_DRX_OnDurationTimer_r18_MilliSeconds_Ms500  = 18
	CellDTX_DRX_Config_r18_CellDTX_DRX_OnDurationTimer_r18_MilliSeconds_Ms600  = 19
	CellDTX_DRX_Config_r18_CellDTX_DRX_OnDurationTimer_r18_MilliSeconds_Ms800  = 20
	CellDTX_DRX_Config_r18_CellDTX_DRX_OnDurationTimer_r18_MilliSeconds_Ms1000 = 21
	CellDTX_DRX_Config_r18_CellDTX_DRX_OnDurationTimer_r18_MilliSeconds_Ms1200 = 22
	CellDTX_DRX_Config_r18_CellDTX_DRX_OnDurationTimer_r18_MilliSeconds_Ms1600 = 23
	CellDTX_DRX_Config_r18_CellDTX_DRX_OnDurationTimer_r18_MilliSeconds_Spare8 = 24
	CellDTX_DRX_Config_r18_CellDTX_DRX_OnDurationTimer_r18_MilliSeconds_Spare7 = 25
	CellDTX_DRX_Config_r18_CellDTX_DRX_OnDurationTimer_r18_MilliSeconds_Spare6 = 26
	CellDTX_DRX_Config_r18_CellDTX_DRX_OnDurationTimer_r18_MilliSeconds_Spare5 = 27
	CellDTX_DRX_Config_r18_CellDTX_DRX_OnDurationTimer_r18_MilliSeconds_Spare4 = 28
	CellDTX_DRX_Config_r18_CellDTX_DRX_OnDurationTimer_r18_MilliSeconds_Spare3 = 29
	CellDTX_DRX_Config_r18_CellDTX_DRX_OnDurationTimer_r18_MilliSeconds_Spare2 = 30
	CellDTX_DRX_Config_r18_CellDTX_DRX_OnDurationTimer_r18_MilliSeconds_Spare1 = 31
)

var cellDTXDRXConfigR18CellDTXDRXOnDurationTimerR18MilliSecondsConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CellDTX_DRX_Config_r18_CellDTX_DRX_OnDurationTimer_r18_MilliSeconds_Ms1, CellDTX_DRX_Config_r18_CellDTX_DRX_OnDurationTimer_r18_MilliSeconds_Ms2, CellDTX_DRX_Config_r18_CellDTX_DRX_OnDurationTimer_r18_MilliSeconds_Ms3, CellDTX_DRX_Config_r18_CellDTX_DRX_OnDurationTimer_r18_MilliSeconds_Ms4, CellDTX_DRX_Config_r18_CellDTX_DRX_OnDurationTimer_r18_MilliSeconds_Ms5, CellDTX_DRX_Config_r18_CellDTX_DRX_OnDurationTimer_r18_MilliSeconds_Ms6, CellDTX_DRX_Config_r18_CellDTX_DRX_OnDurationTimer_r18_MilliSeconds_Ms8, CellDTX_DRX_Config_r18_CellDTX_DRX_OnDurationTimer_r18_MilliSeconds_Ms10, CellDTX_DRX_Config_r18_CellDTX_DRX_OnDurationTimer_r18_MilliSeconds_Ms20, CellDTX_DRX_Config_r18_CellDTX_DRX_OnDurationTimer_r18_MilliSeconds_Ms30, CellDTX_DRX_Config_r18_CellDTX_DRX_OnDurationTimer_r18_MilliSeconds_Ms40, CellDTX_DRX_Config_r18_CellDTX_DRX_OnDurationTimer_r18_MilliSeconds_Ms50, CellDTX_DRX_Config_r18_CellDTX_DRX_OnDurationTimer_r18_MilliSeconds_Ms60, CellDTX_DRX_Config_r18_CellDTX_DRX_OnDurationTimer_r18_MilliSeconds_Ms80, CellDTX_DRX_Config_r18_CellDTX_DRX_OnDurationTimer_r18_MilliSeconds_Ms100, CellDTX_DRX_Config_r18_CellDTX_DRX_OnDurationTimer_r18_MilliSeconds_Ms200, CellDTX_DRX_Config_r18_CellDTX_DRX_OnDurationTimer_r18_MilliSeconds_Ms300, CellDTX_DRX_Config_r18_CellDTX_DRX_OnDurationTimer_r18_MilliSeconds_Ms400, CellDTX_DRX_Config_r18_CellDTX_DRX_OnDurationTimer_r18_MilliSeconds_Ms500, CellDTX_DRX_Config_r18_CellDTX_DRX_OnDurationTimer_r18_MilliSeconds_Ms600, CellDTX_DRX_Config_r18_CellDTX_DRX_OnDurationTimer_r18_MilliSeconds_Ms800, CellDTX_DRX_Config_r18_CellDTX_DRX_OnDurationTimer_r18_MilliSeconds_Ms1000, CellDTX_DRX_Config_r18_CellDTX_DRX_OnDurationTimer_r18_MilliSeconds_Ms1200, CellDTX_DRX_Config_r18_CellDTX_DRX_OnDurationTimer_r18_MilliSeconds_Ms1600, CellDTX_DRX_Config_r18_CellDTX_DRX_OnDurationTimer_r18_MilliSeconds_Spare8, CellDTX_DRX_Config_r18_CellDTX_DRX_OnDurationTimer_r18_MilliSeconds_Spare7, CellDTX_DRX_Config_r18_CellDTX_DRX_OnDurationTimer_r18_MilliSeconds_Spare6, CellDTX_DRX_Config_r18_CellDTX_DRX_OnDurationTimer_r18_MilliSeconds_Spare5, CellDTX_DRX_Config_r18_CellDTX_DRX_OnDurationTimer_r18_MilliSeconds_Spare4, CellDTX_DRX_Config_r18_CellDTX_DRX_OnDurationTimer_r18_MilliSeconds_Spare3, CellDTX_DRX_Config_r18_CellDTX_DRX_OnDurationTimer_r18_MilliSeconds_Spare2, CellDTX_DRX_Config_r18_CellDTX_DRX_OnDurationTimer_r18_MilliSeconds_Spare1},
}

type CellDTX_DRX_Config_r18 struct {
	CellDTX_DRX_OnDurationTimer_r18 struct {
		Choice          int
		SubMilliSeconds *int64
		MilliSeconds    *int64
	}
	CellDTX_DRX_CycleStartOffset_r18 struct {
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
	CellDTX_DRX_SlotOffset_r18       int64
	CellDTX_DRX_ConfigType_r18       int64
	CellDTX_DRX_ActivationStatus_r18 *int64
}

func (ie *CellDTX_DRX_Config_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(cellDTXDRXConfigR18Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.CellDTX_DRX_ActivationStatus_r18 != nil}); err != nil {
		return err
	}

	// 2. cellDTX-DRX-onDurationTimer-r18: choice
	{
		choiceEnc := e.NewChoiceEncoder(cellDTX_DRX_Config_r18CellDTXDRXOnDurationTimerR18Constraints)
		if err := choiceEnc.EncodeChoice(int64(ie.CellDTX_DRX_OnDurationTimer_r18.Choice), false, nil); err != nil {
			return err
		}
		switch ie.CellDTX_DRX_OnDurationTimer_r18.Choice {
		case CellDTX_DRX_Config_r18_CellDTX_DRX_OnDurationTimer_r18_SubMilliSeconds:
			if err := e.EncodeInteger((*ie.CellDTX_DRX_OnDurationTimer_r18.SubMilliSeconds), per.Constrained(1, 31)); err != nil {
				return err
			}
		case CellDTX_DRX_Config_r18_CellDTX_DRX_OnDurationTimer_r18_MilliSeconds:
			if err := e.EncodeEnumerated((*ie.CellDTX_DRX_OnDurationTimer_r18.MilliSeconds), cellDTXDRXConfigR18CellDTXDRXOnDurationTimerR18MilliSecondsConstraints); err != nil {
				return err
			}
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.CellDTX_DRX_OnDurationTimer_r18.Choice), Constraint: "index out of range"}
		}
	}

	// 3. cellDTX-DRX-CycleStartOffset-r18: choice
	{
		choiceEnc := e.NewChoiceEncoder(cellDTX_DRX_Config_r18CellDTXDRXCycleStartOffsetR18Constraints)
		if err := choiceEnc.EncodeChoice(int64(ie.CellDTX_DRX_CycleStartOffset_r18.Choice), false, nil); err != nil {
			return err
		}
		switch ie.CellDTX_DRX_CycleStartOffset_r18.Choice {
		case CellDTX_DRX_Config_r18_CellDTX_DRX_CycleStartOffset_r18_Ms10:
			if err := e.EncodeInteger((*ie.CellDTX_DRX_CycleStartOffset_r18.Ms10), per.Constrained(0, 9)); err != nil {
				return err
			}
		case CellDTX_DRX_Config_r18_CellDTX_DRX_CycleStartOffset_r18_Ms20:
			if err := e.EncodeInteger((*ie.CellDTX_DRX_CycleStartOffset_r18.Ms20), per.Constrained(0, 19)); err != nil {
				return err
			}
		case CellDTX_DRX_Config_r18_CellDTX_DRX_CycleStartOffset_r18_Ms32:
			if err := e.EncodeInteger((*ie.CellDTX_DRX_CycleStartOffset_r18.Ms32), per.Constrained(0, 31)); err != nil {
				return err
			}
		case CellDTX_DRX_Config_r18_CellDTX_DRX_CycleStartOffset_r18_Ms40:
			if err := e.EncodeInteger((*ie.CellDTX_DRX_CycleStartOffset_r18.Ms40), per.Constrained(0, 39)); err != nil {
				return err
			}
		case CellDTX_DRX_Config_r18_CellDTX_DRX_CycleStartOffset_r18_Ms60:
			if err := e.EncodeInteger((*ie.CellDTX_DRX_CycleStartOffset_r18.Ms60), per.Constrained(0, 59)); err != nil {
				return err
			}
		case CellDTX_DRX_Config_r18_CellDTX_DRX_CycleStartOffset_r18_Ms64:
			if err := e.EncodeInteger((*ie.CellDTX_DRX_CycleStartOffset_r18.Ms64), per.Constrained(0, 63)); err != nil {
				return err
			}
		case CellDTX_DRX_Config_r18_CellDTX_DRX_CycleStartOffset_r18_Ms70:
			if err := e.EncodeInteger((*ie.CellDTX_DRX_CycleStartOffset_r18.Ms70), per.Constrained(0, 69)); err != nil {
				return err
			}
		case CellDTX_DRX_Config_r18_CellDTX_DRX_CycleStartOffset_r18_Ms80:
			if err := e.EncodeInteger((*ie.CellDTX_DRX_CycleStartOffset_r18.Ms80), per.Constrained(0, 79)); err != nil {
				return err
			}
		case CellDTX_DRX_Config_r18_CellDTX_DRX_CycleStartOffset_r18_Ms128:
			if err := e.EncodeInteger((*ie.CellDTX_DRX_CycleStartOffset_r18.Ms128), per.Constrained(0, 127)); err != nil {
				return err
			}
		case CellDTX_DRX_Config_r18_CellDTX_DRX_CycleStartOffset_r18_Ms160:
			if err := e.EncodeInteger((*ie.CellDTX_DRX_CycleStartOffset_r18.Ms160), per.Constrained(0, 159)); err != nil {
				return err
			}
		case CellDTX_DRX_Config_r18_CellDTX_DRX_CycleStartOffset_r18_Ms256:
			if err := e.EncodeInteger((*ie.CellDTX_DRX_CycleStartOffset_r18.Ms256), per.Constrained(0, 255)); err != nil {
				return err
			}
		case CellDTX_DRX_Config_r18_CellDTX_DRX_CycleStartOffset_r18_Ms320:
			if err := e.EncodeInteger((*ie.CellDTX_DRX_CycleStartOffset_r18.Ms320), per.Constrained(0, 319)); err != nil {
				return err
			}
		case CellDTX_DRX_Config_r18_CellDTX_DRX_CycleStartOffset_r18_Ms512:
			if err := e.EncodeInteger((*ie.CellDTX_DRX_CycleStartOffset_r18.Ms512), per.Constrained(0, 511)); err != nil {
				return err
			}
		case CellDTX_DRX_Config_r18_CellDTX_DRX_CycleStartOffset_r18_Ms640:
			if err := e.EncodeInteger((*ie.CellDTX_DRX_CycleStartOffset_r18.Ms640), per.Constrained(0, 639)); err != nil {
				return err
			}
		case CellDTX_DRX_Config_r18_CellDTX_DRX_CycleStartOffset_r18_Ms1024:
			if err := e.EncodeInteger((*ie.CellDTX_DRX_CycleStartOffset_r18.Ms1024), per.Constrained(0, 1023)); err != nil {
				return err
			}
		case CellDTX_DRX_Config_r18_CellDTX_DRX_CycleStartOffset_r18_Ms1280:
			if err := e.EncodeInteger((*ie.CellDTX_DRX_CycleStartOffset_r18.Ms1280), per.Constrained(0, 1279)); err != nil {
				return err
			}
		case CellDTX_DRX_Config_r18_CellDTX_DRX_CycleStartOffset_r18_Ms2048:
			if err := e.EncodeInteger((*ie.CellDTX_DRX_CycleStartOffset_r18.Ms2048), per.Constrained(0, 2047)); err != nil {
				return err
			}
		case CellDTX_DRX_Config_r18_CellDTX_DRX_CycleStartOffset_r18_Ms2560:
			if err := e.EncodeInteger((*ie.CellDTX_DRX_CycleStartOffset_r18.Ms2560), per.Constrained(0, 2559)); err != nil {
				return err
			}
		case CellDTX_DRX_Config_r18_CellDTX_DRX_CycleStartOffset_r18_Ms5120:
			if err := e.EncodeInteger((*ie.CellDTX_DRX_CycleStartOffset_r18.Ms5120), per.Constrained(0, 5119)); err != nil {
				return err
			}
		case CellDTX_DRX_Config_r18_CellDTX_DRX_CycleStartOffset_r18_Ms10240:
			if err := e.EncodeInteger((*ie.CellDTX_DRX_CycleStartOffset_r18.Ms10240), per.Constrained(0, 10239)); err != nil {
				return err
			}
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.CellDTX_DRX_CycleStartOffset_r18.Choice), Constraint: "index out of range"}
		}
	}

	// 4. cellDTX-DRX-SlotOffset-r18: integer
	{
		if err := e.EncodeInteger(ie.CellDTX_DRX_SlotOffset_r18, cellDTXDRXConfigR18CellDTXDRXSlotOffsetR18Constraints); err != nil {
			return err
		}
	}

	// 5. cellDTX-DRX-ConfigType-r18: enumerated
	{
		if err := e.EncodeEnumerated(ie.CellDTX_DRX_ConfigType_r18, cellDTXDRXConfigR18CellDTXDRXConfigTypeR18Constraints); err != nil {
			return err
		}
	}

	// 6. cellDTX-DRX-ActivationStatus-r18: enumerated
	{
		if ie.CellDTX_DRX_ActivationStatus_r18 != nil {
			if err := e.EncodeEnumerated(*ie.CellDTX_DRX_ActivationStatus_r18, cellDTXDRXConfigR18CellDTXDRXActivationStatusR18Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *CellDTX_DRX_Config_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(cellDTXDRXConfigR18Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. cellDTX-DRX-onDurationTimer-r18: choice
	{
		choiceDec := d.NewChoiceDecoder(cellDTX_DRX_Config_r18CellDTXDRXOnDurationTimerR18Constraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.CellDTX_DRX_OnDurationTimer_r18.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case CellDTX_DRX_Config_r18_CellDTX_DRX_OnDurationTimer_r18_SubMilliSeconds:
			ie.CellDTX_DRX_OnDurationTimer_r18.SubMilliSeconds = new(int64)
			v, err := d.DecodeInteger(per.Constrained(1, 31))
			if err != nil {
				return err
			}
			(*ie.CellDTX_DRX_OnDurationTimer_r18.SubMilliSeconds) = v
		case CellDTX_DRX_Config_r18_CellDTX_DRX_OnDurationTimer_r18_MilliSeconds:
			ie.CellDTX_DRX_OnDurationTimer_r18.MilliSeconds = new(int64)
			v, err := d.DecodeEnumerated(cellDTXDRXConfigR18CellDTXDRXOnDurationTimerR18MilliSecondsConstraints)
			if err != nil {
				return err
			}
			(*ie.CellDTX_DRX_OnDurationTimer_r18.MilliSeconds) = v
		}
	}

	// 3. cellDTX-DRX-CycleStartOffset-r18: choice
	{
		choiceDec := d.NewChoiceDecoder(cellDTX_DRX_Config_r18CellDTXDRXCycleStartOffsetR18Constraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.CellDTX_DRX_CycleStartOffset_r18.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case CellDTX_DRX_Config_r18_CellDTX_DRX_CycleStartOffset_r18_Ms10:
			ie.CellDTX_DRX_CycleStartOffset_r18.Ms10 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 9))
			if err != nil {
				return err
			}
			(*ie.CellDTX_DRX_CycleStartOffset_r18.Ms10) = v
		case CellDTX_DRX_Config_r18_CellDTX_DRX_CycleStartOffset_r18_Ms20:
			ie.CellDTX_DRX_CycleStartOffset_r18.Ms20 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 19))
			if err != nil {
				return err
			}
			(*ie.CellDTX_DRX_CycleStartOffset_r18.Ms20) = v
		case CellDTX_DRX_Config_r18_CellDTX_DRX_CycleStartOffset_r18_Ms32:
			ie.CellDTX_DRX_CycleStartOffset_r18.Ms32 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 31))
			if err != nil {
				return err
			}
			(*ie.CellDTX_DRX_CycleStartOffset_r18.Ms32) = v
		case CellDTX_DRX_Config_r18_CellDTX_DRX_CycleStartOffset_r18_Ms40:
			ie.CellDTX_DRX_CycleStartOffset_r18.Ms40 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 39))
			if err != nil {
				return err
			}
			(*ie.CellDTX_DRX_CycleStartOffset_r18.Ms40) = v
		case CellDTX_DRX_Config_r18_CellDTX_DRX_CycleStartOffset_r18_Ms60:
			ie.CellDTX_DRX_CycleStartOffset_r18.Ms60 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 59))
			if err != nil {
				return err
			}
			(*ie.CellDTX_DRX_CycleStartOffset_r18.Ms60) = v
		case CellDTX_DRX_Config_r18_CellDTX_DRX_CycleStartOffset_r18_Ms64:
			ie.CellDTX_DRX_CycleStartOffset_r18.Ms64 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 63))
			if err != nil {
				return err
			}
			(*ie.CellDTX_DRX_CycleStartOffset_r18.Ms64) = v
		case CellDTX_DRX_Config_r18_CellDTX_DRX_CycleStartOffset_r18_Ms70:
			ie.CellDTX_DRX_CycleStartOffset_r18.Ms70 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 69))
			if err != nil {
				return err
			}
			(*ie.CellDTX_DRX_CycleStartOffset_r18.Ms70) = v
		case CellDTX_DRX_Config_r18_CellDTX_DRX_CycleStartOffset_r18_Ms80:
			ie.CellDTX_DRX_CycleStartOffset_r18.Ms80 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 79))
			if err != nil {
				return err
			}
			(*ie.CellDTX_DRX_CycleStartOffset_r18.Ms80) = v
		case CellDTX_DRX_Config_r18_CellDTX_DRX_CycleStartOffset_r18_Ms128:
			ie.CellDTX_DRX_CycleStartOffset_r18.Ms128 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 127))
			if err != nil {
				return err
			}
			(*ie.CellDTX_DRX_CycleStartOffset_r18.Ms128) = v
		case CellDTX_DRX_Config_r18_CellDTX_DRX_CycleStartOffset_r18_Ms160:
			ie.CellDTX_DRX_CycleStartOffset_r18.Ms160 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 159))
			if err != nil {
				return err
			}
			(*ie.CellDTX_DRX_CycleStartOffset_r18.Ms160) = v
		case CellDTX_DRX_Config_r18_CellDTX_DRX_CycleStartOffset_r18_Ms256:
			ie.CellDTX_DRX_CycleStartOffset_r18.Ms256 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 255))
			if err != nil {
				return err
			}
			(*ie.CellDTX_DRX_CycleStartOffset_r18.Ms256) = v
		case CellDTX_DRX_Config_r18_CellDTX_DRX_CycleStartOffset_r18_Ms320:
			ie.CellDTX_DRX_CycleStartOffset_r18.Ms320 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 319))
			if err != nil {
				return err
			}
			(*ie.CellDTX_DRX_CycleStartOffset_r18.Ms320) = v
		case CellDTX_DRX_Config_r18_CellDTX_DRX_CycleStartOffset_r18_Ms512:
			ie.CellDTX_DRX_CycleStartOffset_r18.Ms512 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 511))
			if err != nil {
				return err
			}
			(*ie.CellDTX_DRX_CycleStartOffset_r18.Ms512) = v
		case CellDTX_DRX_Config_r18_CellDTX_DRX_CycleStartOffset_r18_Ms640:
			ie.CellDTX_DRX_CycleStartOffset_r18.Ms640 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 639))
			if err != nil {
				return err
			}
			(*ie.CellDTX_DRX_CycleStartOffset_r18.Ms640) = v
		case CellDTX_DRX_Config_r18_CellDTX_DRX_CycleStartOffset_r18_Ms1024:
			ie.CellDTX_DRX_CycleStartOffset_r18.Ms1024 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 1023))
			if err != nil {
				return err
			}
			(*ie.CellDTX_DRX_CycleStartOffset_r18.Ms1024) = v
		case CellDTX_DRX_Config_r18_CellDTX_DRX_CycleStartOffset_r18_Ms1280:
			ie.CellDTX_DRX_CycleStartOffset_r18.Ms1280 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 1279))
			if err != nil {
				return err
			}
			(*ie.CellDTX_DRX_CycleStartOffset_r18.Ms1280) = v
		case CellDTX_DRX_Config_r18_CellDTX_DRX_CycleStartOffset_r18_Ms2048:
			ie.CellDTX_DRX_CycleStartOffset_r18.Ms2048 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 2047))
			if err != nil {
				return err
			}
			(*ie.CellDTX_DRX_CycleStartOffset_r18.Ms2048) = v
		case CellDTX_DRX_Config_r18_CellDTX_DRX_CycleStartOffset_r18_Ms2560:
			ie.CellDTX_DRX_CycleStartOffset_r18.Ms2560 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 2559))
			if err != nil {
				return err
			}
			(*ie.CellDTX_DRX_CycleStartOffset_r18.Ms2560) = v
		case CellDTX_DRX_Config_r18_CellDTX_DRX_CycleStartOffset_r18_Ms5120:
			ie.CellDTX_DRX_CycleStartOffset_r18.Ms5120 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 5119))
			if err != nil {
				return err
			}
			(*ie.CellDTX_DRX_CycleStartOffset_r18.Ms5120) = v
		case CellDTX_DRX_Config_r18_CellDTX_DRX_CycleStartOffset_r18_Ms10240:
			ie.CellDTX_DRX_CycleStartOffset_r18.Ms10240 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 10239))
			if err != nil {
				return err
			}
			(*ie.CellDTX_DRX_CycleStartOffset_r18.Ms10240) = v
		}
	}

	// 4. cellDTX-DRX-SlotOffset-r18: integer
	{
		v2, err := d.DecodeInteger(cellDTXDRXConfigR18CellDTXDRXSlotOffsetR18Constraints)
		if err != nil {
			return err
		}
		ie.CellDTX_DRX_SlotOffset_r18 = v2
	}

	// 5. cellDTX-DRX-ConfigType-r18: enumerated
	{
		v3, err := d.DecodeEnumerated(cellDTXDRXConfigR18CellDTXDRXConfigTypeR18Constraints)
		if err != nil {
			return err
		}
		ie.CellDTX_DRX_ConfigType_r18 = v3
	}

	// 6. cellDTX-DRX-ActivationStatus-r18: enumerated
	{
		if seq.IsComponentPresent(4) {
			idx, err := d.DecodeEnumerated(cellDTXDRXConfigR18CellDTXDRXActivationStatusR18Constraints)
			if err != nil {
				return err
			}
			ie.CellDTX_DRX_ActivationStatus_r18 = &idx
		}
	}

	return nil
}
