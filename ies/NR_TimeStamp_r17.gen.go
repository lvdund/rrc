// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: NR-TimeStamp-r17 (line 3601).

var nRTimeStampR17Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "nr-SFN-r17"},
		{Name: "nr-Slot-r17"},
	},
}

var nRTimeStampR17NrSFNR17Constraints = per.Constrained(0, 1023)

var nR_TimeStamp_r17NrSlotR17Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "scs15-r17"},
		{Name: "scs30-r17"},
		{Name: "scs60-r17"},
		{Name: "scs120-r17"},
	},
}

const (
	NR_TimeStamp_r17_Nr_Slot_r17_Scs15_r17  = 0
	NR_TimeStamp_r17_Nr_Slot_r17_Scs30_r17  = 1
	NR_TimeStamp_r17_Nr_Slot_r17_Scs60_r17  = 2
	NR_TimeStamp_r17_Nr_Slot_r17_Scs120_r17 = 3
)

type NR_TimeStamp_r17 struct {
	Nr_SFN_r17  int64
	Nr_Slot_r17 struct {
		Choice     int
		Scs15_r17  *int64
		Scs30_r17  *int64
		Scs60_r17  *int64
		Scs120_r17 *int64
	}
}

func (ie *NR_TimeStamp_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(nRTimeStampR17Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. nr-SFN-r17: integer
	{
		if err := e.EncodeInteger(ie.Nr_SFN_r17, nRTimeStampR17NrSFNR17Constraints); err != nil {
			return err
		}
	}

	// 3. nr-Slot-r17: choice
	{
		choiceEnc := e.NewChoiceEncoder(nR_TimeStamp_r17NrSlotR17Constraints)
		if err := choiceEnc.EncodeChoice(int64(ie.Nr_Slot_r17.Choice), false, nil); err != nil {
			return err
		}
		switch ie.Nr_Slot_r17.Choice {
		case NR_TimeStamp_r17_Nr_Slot_r17_Scs15_r17:
			if err := e.EncodeInteger((*ie.Nr_Slot_r17.Scs15_r17), per.Constrained(0, 9)); err != nil {
				return err
			}
		case NR_TimeStamp_r17_Nr_Slot_r17_Scs30_r17:
			if err := e.EncodeInteger((*ie.Nr_Slot_r17.Scs30_r17), per.Constrained(0, 19)); err != nil {
				return err
			}
		case NR_TimeStamp_r17_Nr_Slot_r17_Scs60_r17:
			if err := e.EncodeInteger((*ie.Nr_Slot_r17.Scs60_r17), per.Constrained(0, 39)); err != nil {
				return err
			}
		case NR_TimeStamp_r17_Nr_Slot_r17_Scs120_r17:
			if err := e.EncodeInteger((*ie.Nr_Slot_r17.Scs120_r17), per.Constrained(0, 79)); err != nil {
				return err
			}
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.Nr_Slot_r17.Choice), Constraint: "index out of range"}
		}
	}

	return nil
}

func (ie *NR_TimeStamp_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(nRTimeStampR17Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. nr-SFN-r17: integer
	{
		v0, err := d.DecodeInteger(nRTimeStampR17NrSFNR17Constraints)
		if err != nil {
			return err
		}
		ie.Nr_SFN_r17 = v0
	}

	// 3. nr-Slot-r17: choice
	{
		choiceDec := d.NewChoiceDecoder(nR_TimeStamp_r17NrSlotR17Constraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.Nr_Slot_r17.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case NR_TimeStamp_r17_Nr_Slot_r17_Scs15_r17:
			ie.Nr_Slot_r17.Scs15_r17 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 9))
			if err != nil {
				return err
			}
			(*ie.Nr_Slot_r17.Scs15_r17) = v
		case NR_TimeStamp_r17_Nr_Slot_r17_Scs30_r17:
			ie.Nr_Slot_r17.Scs30_r17 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 19))
			if err != nil {
				return err
			}
			(*ie.Nr_Slot_r17.Scs30_r17) = v
		case NR_TimeStamp_r17_Nr_Slot_r17_Scs60_r17:
			ie.Nr_Slot_r17.Scs60_r17 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 39))
			if err != nil {
				return err
			}
			(*ie.Nr_Slot_r17.Scs60_r17) = v
		case NR_TimeStamp_r17_Nr_Slot_r17_Scs120_r17:
			ie.Nr_Slot_r17.Scs120_r17 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 79))
			if err != nil {
				return err
			}
			(*ie.Nr_Slot_r17.Scs120_r17) = v
		}
	}

	return nil
}
