// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: IDC-TDM-Assistance-r18 (line 2721).

var iDCTDMAssistanceR18Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "cycleLength-r18"},
		{Name: "startOffset-r18"},
		{Name: "slotOffset-r18"},
		{Name: "activeDuration-r18"},
	},
}

const (
	IDC_TDM_Assistance_r18_CycleLength_r18_Ms2     = 0
	IDC_TDM_Assistance_r18_CycleLength_r18_Ms3     = 1
	IDC_TDM_Assistance_r18_CycleLength_r18_Ms4     = 2
	IDC_TDM_Assistance_r18_CycleLength_r18_Ms5     = 3
	IDC_TDM_Assistance_r18_CycleLength_r18_Ms6     = 4
	IDC_TDM_Assistance_r18_CycleLength_r18_Ms7     = 5
	IDC_TDM_Assistance_r18_CycleLength_r18_Ms8     = 6
	IDC_TDM_Assistance_r18_CycleLength_r18_Ms10    = 7
	IDC_TDM_Assistance_r18_CycleLength_r18_Ms14    = 8
	IDC_TDM_Assistance_r18_CycleLength_r18_Ms16    = 9
	IDC_TDM_Assistance_r18_CycleLength_r18_Ms20    = 10
	IDC_TDM_Assistance_r18_CycleLength_r18_Ms30    = 11
	IDC_TDM_Assistance_r18_CycleLength_r18_Ms32    = 12
	IDC_TDM_Assistance_r18_CycleLength_r18_Ms35    = 13
	IDC_TDM_Assistance_r18_CycleLength_r18_Ms40    = 14
	IDC_TDM_Assistance_r18_CycleLength_r18_Ms60    = 15
	IDC_TDM_Assistance_r18_CycleLength_r18_Ms64    = 16
	IDC_TDM_Assistance_r18_CycleLength_r18_Ms70    = 17
	IDC_TDM_Assistance_r18_CycleLength_r18_Ms80    = 18
	IDC_TDM_Assistance_r18_CycleLength_r18_Ms96    = 19
	IDC_TDM_Assistance_r18_CycleLength_r18_Ms100   = 20
	IDC_TDM_Assistance_r18_CycleLength_r18_Ms128   = 21
	IDC_TDM_Assistance_r18_CycleLength_r18_Ms160   = 22
	IDC_TDM_Assistance_r18_CycleLength_r18_Ms256   = 23
	IDC_TDM_Assistance_r18_CycleLength_r18_Ms320   = 24
	IDC_TDM_Assistance_r18_CycleLength_r18_Ms512   = 25
	IDC_TDM_Assistance_r18_CycleLength_r18_Ms640   = 26
	IDC_TDM_Assistance_r18_CycleLength_r18_Ms1024  = 27
	IDC_TDM_Assistance_r18_CycleLength_r18_Ms1280  = 28
	IDC_TDM_Assistance_r18_CycleLength_r18_Ms2048  = 29
	IDC_TDM_Assistance_r18_CycleLength_r18_Ms2560  = 30
	IDC_TDM_Assistance_r18_CycleLength_r18_Ms5120  = 31
	IDC_TDM_Assistance_r18_CycleLength_r18_Ms10240 = 32
)

var iDCTDMAssistanceR18CycleLengthR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{IDC_TDM_Assistance_r18_CycleLength_r18_Ms2, IDC_TDM_Assistance_r18_CycleLength_r18_Ms3, IDC_TDM_Assistance_r18_CycleLength_r18_Ms4, IDC_TDM_Assistance_r18_CycleLength_r18_Ms5, IDC_TDM_Assistance_r18_CycleLength_r18_Ms6, IDC_TDM_Assistance_r18_CycleLength_r18_Ms7, IDC_TDM_Assistance_r18_CycleLength_r18_Ms8, IDC_TDM_Assistance_r18_CycleLength_r18_Ms10, IDC_TDM_Assistance_r18_CycleLength_r18_Ms14, IDC_TDM_Assistance_r18_CycleLength_r18_Ms16, IDC_TDM_Assistance_r18_CycleLength_r18_Ms20, IDC_TDM_Assistance_r18_CycleLength_r18_Ms30, IDC_TDM_Assistance_r18_CycleLength_r18_Ms32, IDC_TDM_Assistance_r18_CycleLength_r18_Ms35, IDC_TDM_Assistance_r18_CycleLength_r18_Ms40, IDC_TDM_Assistance_r18_CycleLength_r18_Ms60, IDC_TDM_Assistance_r18_CycleLength_r18_Ms64, IDC_TDM_Assistance_r18_CycleLength_r18_Ms70, IDC_TDM_Assistance_r18_CycleLength_r18_Ms80, IDC_TDM_Assistance_r18_CycleLength_r18_Ms96, IDC_TDM_Assistance_r18_CycleLength_r18_Ms100, IDC_TDM_Assistance_r18_CycleLength_r18_Ms128, IDC_TDM_Assistance_r18_CycleLength_r18_Ms160, IDC_TDM_Assistance_r18_CycleLength_r18_Ms256, IDC_TDM_Assistance_r18_CycleLength_r18_Ms320, IDC_TDM_Assistance_r18_CycleLength_r18_Ms512, IDC_TDM_Assistance_r18_CycleLength_r18_Ms640, IDC_TDM_Assistance_r18_CycleLength_r18_Ms1024, IDC_TDM_Assistance_r18_CycleLength_r18_Ms1280, IDC_TDM_Assistance_r18_CycleLength_r18_Ms2048, IDC_TDM_Assistance_r18_CycleLength_r18_Ms2560, IDC_TDM_Assistance_r18_CycleLength_r18_Ms5120, IDC_TDM_Assistance_r18_CycleLength_r18_Ms10240},
}

var iDCTDMAssistanceR18StartOffsetR18Constraints = per.Constrained(0, 10239)

var iDCTDMAssistanceR18SlotOffsetR18Constraints = per.Constrained(0, 31)

var iDC_TDM_Assistance_r18ActiveDurationR18Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "subMilliSeconds-r18"},
		{Name: "milliSeconds-r18"},
	},
}

const (
	IDC_TDM_Assistance_r18_ActiveDuration_r18_SubMilliSeconds_r18 = 0
	IDC_TDM_Assistance_r18_ActiveDuration_r18_MilliSeconds_r18    = 1
)

const (
	IDC_TDM_Assistance_r18_ActiveDuration_r18_MilliSeconds_r18_Ms1    = 0
	IDC_TDM_Assistance_r18_ActiveDuration_r18_MilliSeconds_r18_Ms2    = 1
	IDC_TDM_Assistance_r18_ActiveDuration_r18_MilliSeconds_r18_Ms3    = 2
	IDC_TDM_Assistance_r18_ActiveDuration_r18_MilliSeconds_r18_Ms4    = 3
	IDC_TDM_Assistance_r18_ActiveDuration_r18_MilliSeconds_r18_Ms5    = 4
	IDC_TDM_Assistance_r18_ActiveDuration_r18_MilliSeconds_r18_Ms6    = 5
	IDC_TDM_Assistance_r18_ActiveDuration_r18_MilliSeconds_r18_Ms8    = 6
	IDC_TDM_Assistance_r18_ActiveDuration_r18_MilliSeconds_r18_Ms10   = 7
	IDC_TDM_Assistance_r18_ActiveDuration_r18_MilliSeconds_r18_Ms20   = 8
	IDC_TDM_Assistance_r18_ActiveDuration_r18_MilliSeconds_r18_Ms30   = 9
	IDC_TDM_Assistance_r18_ActiveDuration_r18_MilliSeconds_r18_Ms40   = 10
	IDC_TDM_Assistance_r18_ActiveDuration_r18_MilliSeconds_r18_Ms50   = 11
	IDC_TDM_Assistance_r18_ActiveDuration_r18_MilliSeconds_r18_Ms60   = 12
	IDC_TDM_Assistance_r18_ActiveDuration_r18_MilliSeconds_r18_Ms80   = 13
	IDC_TDM_Assistance_r18_ActiveDuration_r18_MilliSeconds_r18_Ms100  = 14
	IDC_TDM_Assistance_r18_ActiveDuration_r18_MilliSeconds_r18_Ms200  = 15
	IDC_TDM_Assistance_r18_ActiveDuration_r18_MilliSeconds_r18_Ms300  = 16
	IDC_TDM_Assistance_r18_ActiveDuration_r18_MilliSeconds_r18_Ms400  = 17
	IDC_TDM_Assistance_r18_ActiveDuration_r18_MilliSeconds_r18_Ms500  = 18
	IDC_TDM_Assistance_r18_ActiveDuration_r18_MilliSeconds_r18_Ms600  = 19
	IDC_TDM_Assistance_r18_ActiveDuration_r18_MilliSeconds_r18_Ms800  = 20
	IDC_TDM_Assistance_r18_ActiveDuration_r18_MilliSeconds_r18_Ms1000 = 21
	IDC_TDM_Assistance_r18_ActiveDuration_r18_MilliSeconds_r18_Ms1200 = 22
	IDC_TDM_Assistance_r18_ActiveDuration_r18_MilliSeconds_r18_Ms1600 = 23
	IDC_TDM_Assistance_r18_ActiveDuration_r18_MilliSeconds_r18_Spare8 = 24
	IDC_TDM_Assistance_r18_ActiveDuration_r18_MilliSeconds_r18_Spare7 = 25
	IDC_TDM_Assistance_r18_ActiveDuration_r18_MilliSeconds_r18_Spare6 = 26
	IDC_TDM_Assistance_r18_ActiveDuration_r18_MilliSeconds_r18_Spare5 = 27
	IDC_TDM_Assistance_r18_ActiveDuration_r18_MilliSeconds_r18_Spare4 = 28
	IDC_TDM_Assistance_r18_ActiveDuration_r18_MilliSeconds_r18_Spare3 = 29
	IDC_TDM_Assistance_r18_ActiveDuration_r18_MilliSeconds_r18_Spare2 = 30
	IDC_TDM_Assistance_r18_ActiveDuration_r18_MilliSeconds_r18_Spare1 = 31
)

var iDCTDMAssistanceR18ActiveDurationR18MilliSecondsR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{IDC_TDM_Assistance_r18_ActiveDuration_r18_MilliSeconds_r18_Ms1, IDC_TDM_Assistance_r18_ActiveDuration_r18_MilliSeconds_r18_Ms2, IDC_TDM_Assistance_r18_ActiveDuration_r18_MilliSeconds_r18_Ms3, IDC_TDM_Assistance_r18_ActiveDuration_r18_MilliSeconds_r18_Ms4, IDC_TDM_Assistance_r18_ActiveDuration_r18_MilliSeconds_r18_Ms5, IDC_TDM_Assistance_r18_ActiveDuration_r18_MilliSeconds_r18_Ms6, IDC_TDM_Assistance_r18_ActiveDuration_r18_MilliSeconds_r18_Ms8, IDC_TDM_Assistance_r18_ActiveDuration_r18_MilliSeconds_r18_Ms10, IDC_TDM_Assistance_r18_ActiveDuration_r18_MilliSeconds_r18_Ms20, IDC_TDM_Assistance_r18_ActiveDuration_r18_MilliSeconds_r18_Ms30, IDC_TDM_Assistance_r18_ActiveDuration_r18_MilliSeconds_r18_Ms40, IDC_TDM_Assistance_r18_ActiveDuration_r18_MilliSeconds_r18_Ms50, IDC_TDM_Assistance_r18_ActiveDuration_r18_MilliSeconds_r18_Ms60, IDC_TDM_Assistance_r18_ActiveDuration_r18_MilliSeconds_r18_Ms80, IDC_TDM_Assistance_r18_ActiveDuration_r18_MilliSeconds_r18_Ms100, IDC_TDM_Assistance_r18_ActiveDuration_r18_MilliSeconds_r18_Ms200, IDC_TDM_Assistance_r18_ActiveDuration_r18_MilliSeconds_r18_Ms300, IDC_TDM_Assistance_r18_ActiveDuration_r18_MilliSeconds_r18_Ms400, IDC_TDM_Assistance_r18_ActiveDuration_r18_MilliSeconds_r18_Ms500, IDC_TDM_Assistance_r18_ActiveDuration_r18_MilliSeconds_r18_Ms600, IDC_TDM_Assistance_r18_ActiveDuration_r18_MilliSeconds_r18_Ms800, IDC_TDM_Assistance_r18_ActiveDuration_r18_MilliSeconds_r18_Ms1000, IDC_TDM_Assistance_r18_ActiveDuration_r18_MilliSeconds_r18_Ms1200, IDC_TDM_Assistance_r18_ActiveDuration_r18_MilliSeconds_r18_Ms1600, IDC_TDM_Assistance_r18_ActiveDuration_r18_MilliSeconds_r18_Spare8, IDC_TDM_Assistance_r18_ActiveDuration_r18_MilliSeconds_r18_Spare7, IDC_TDM_Assistance_r18_ActiveDuration_r18_MilliSeconds_r18_Spare6, IDC_TDM_Assistance_r18_ActiveDuration_r18_MilliSeconds_r18_Spare5, IDC_TDM_Assistance_r18_ActiveDuration_r18_MilliSeconds_r18_Spare4, IDC_TDM_Assistance_r18_ActiveDuration_r18_MilliSeconds_r18_Spare3, IDC_TDM_Assistance_r18_ActiveDuration_r18_MilliSeconds_r18_Spare2, IDC_TDM_Assistance_r18_ActiveDuration_r18_MilliSeconds_r18_Spare1},
}

type IDC_TDM_Assistance_r18 struct {
	CycleLength_r18    int64
	StartOffset_r18    int64
	SlotOffset_r18     int64
	ActiveDuration_r18 struct {
		Choice              int
		SubMilliSeconds_r18 *int64
		MilliSeconds_r18    *int64
	}
}

func (ie *IDC_TDM_Assistance_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(iDCTDMAssistanceR18Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. cycleLength-r18: enumerated
	{
		if err := e.EncodeEnumerated(ie.CycleLength_r18, iDCTDMAssistanceR18CycleLengthR18Constraints); err != nil {
			return err
		}
	}

	// 3. startOffset-r18: integer
	{
		if err := e.EncodeInteger(ie.StartOffset_r18, iDCTDMAssistanceR18StartOffsetR18Constraints); err != nil {
			return err
		}
	}

	// 4. slotOffset-r18: integer
	{
		if err := e.EncodeInteger(ie.SlotOffset_r18, iDCTDMAssistanceR18SlotOffsetR18Constraints); err != nil {
			return err
		}
	}

	// 5. activeDuration-r18: choice
	{
		choiceEnc := e.NewChoiceEncoder(iDC_TDM_Assistance_r18ActiveDurationR18Constraints)
		if err := choiceEnc.EncodeChoice(int64(ie.ActiveDuration_r18.Choice), false, nil); err != nil {
			return err
		}
		switch ie.ActiveDuration_r18.Choice {
		case IDC_TDM_Assistance_r18_ActiveDuration_r18_SubMilliSeconds_r18:
			if err := e.EncodeInteger((*ie.ActiveDuration_r18.SubMilliSeconds_r18), per.Constrained(1, 31)); err != nil {
				return err
			}
		case IDC_TDM_Assistance_r18_ActiveDuration_r18_MilliSeconds_r18:
			if err := e.EncodeEnumerated((*ie.ActiveDuration_r18.MilliSeconds_r18), iDCTDMAssistanceR18ActiveDurationR18MilliSecondsR18Constraints); err != nil {
				return err
			}
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.ActiveDuration_r18.Choice), Constraint: "index out of range"}
		}
	}

	return nil
}

func (ie *IDC_TDM_Assistance_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(iDCTDMAssistanceR18Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. cycleLength-r18: enumerated
	{
		v0, err := d.DecodeEnumerated(iDCTDMAssistanceR18CycleLengthR18Constraints)
		if err != nil {
			return err
		}
		ie.CycleLength_r18 = v0
	}

	// 3. startOffset-r18: integer
	{
		v1, err := d.DecodeInteger(iDCTDMAssistanceR18StartOffsetR18Constraints)
		if err != nil {
			return err
		}
		ie.StartOffset_r18 = v1
	}

	// 4. slotOffset-r18: integer
	{
		v2, err := d.DecodeInteger(iDCTDMAssistanceR18SlotOffsetR18Constraints)
		if err != nil {
			return err
		}
		ie.SlotOffset_r18 = v2
	}

	// 5. activeDuration-r18: choice
	{
		choiceDec := d.NewChoiceDecoder(iDC_TDM_Assistance_r18ActiveDurationR18Constraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.ActiveDuration_r18.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case IDC_TDM_Assistance_r18_ActiveDuration_r18_SubMilliSeconds_r18:
			ie.ActiveDuration_r18.SubMilliSeconds_r18 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(1, 31))
			if err != nil {
				return err
			}
			(*ie.ActiveDuration_r18.SubMilliSeconds_r18) = v
		case IDC_TDM_Assistance_r18_ActiveDuration_r18_MilliSeconds_r18:
			ie.ActiveDuration_r18.MilliSeconds_r18 = new(int64)
			v, err := d.DecodeEnumerated(iDCTDMAssistanceR18ActiveDurationR18MilliSecondsR18Constraints)
			if err != nil {
				return err
			}
			(*ie.ActiveDuration_r18.MilliSeconds_r18) = v
		}
	}

	return nil
}
