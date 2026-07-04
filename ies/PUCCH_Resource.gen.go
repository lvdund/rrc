// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: PUCCH-Resource (line 12072).

var pUCCHResourceConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "pucch-ResourceId"},
		{Name: "startingPRB"},
		{Name: "intraSlotFrequencyHopping", Optional: true},
		{Name: "secondHopPRB", Optional: true},
		{Name: "format"},
	},
}

const (
	PUCCH_Resource_IntraSlotFrequencyHopping_Enabled = 0
)

var pUCCHResourceIntraSlotFrequencyHoppingConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PUCCH_Resource_IntraSlotFrequencyHopping_Enabled},
}

var pUCCH_ResourceFormatConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "format0"},
		{Name: "format1"},
		{Name: "format2"},
		{Name: "format3"},
		{Name: "format4"},
	},
}

const (
	PUCCH_Resource_Format_Format0 = 0
	PUCCH_Resource_Format_Format1 = 1
	PUCCH_Resource_Format_Format2 = 2
	PUCCH_Resource_Format_Format3 = 3
	PUCCH_Resource_Format_Format4 = 4
)

type PUCCH_Resource struct {
	Pucch_ResourceId          PUCCH_ResourceId
	StartingPRB               PRB_Id
	IntraSlotFrequencyHopping *int64
	SecondHopPRB              *PRB_Id
	Format                    struct {
		Choice  int
		Format0 *PUCCH_Format0
		Format1 *PUCCH_Format1
		Format2 *PUCCH_Format2
		Format3 *PUCCH_Format3
		Format4 *PUCCH_Format4
	}
}

func (ie *PUCCH_Resource) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(pUCCHResourceConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.IntraSlotFrequencyHopping != nil, ie.SecondHopPRB != nil}); err != nil {
		return err
	}

	// 2. pucch-ResourceId: ref
	{
		if err := ie.Pucch_ResourceId.Encode(e); err != nil {
			return err
		}
	}

	// 3. startingPRB: ref
	{
		if err := ie.StartingPRB.Encode(e); err != nil {
			return err
		}
	}

	// 4. intraSlotFrequencyHopping: enumerated
	{
		if ie.IntraSlotFrequencyHopping != nil {
			if err := e.EncodeEnumerated(*ie.IntraSlotFrequencyHopping, pUCCHResourceIntraSlotFrequencyHoppingConstraints); err != nil {
				return err
			}
		}
	}

	// 5. secondHopPRB: ref
	{
		if ie.SecondHopPRB != nil {
			if err := ie.SecondHopPRB.Encode(e); err != nil {
				return err
			}
		}
	}

	// 6. format: choice
	{
		choiceEnc := e.NewChoiceEncoder(pUCCH_ResourceFormatConstraints)
		if err := choiceEnc.EncodeChoice(int64(ie.Format.Choice), false, nil); err != nil {
			return err
		}
		switch ie.Format.Choice {
		case PUCCH_Resource_Format_Format0:
			if err := ie.Format.Format0.Encode(e); err != nil {
				return err
			}
		case PUCCH_Resource_Format_Format1:
			if err := ie.Format.Format1.Encode(e); err != nil {
				return err
			}
		case PUCCH_Resource_Format_Format2:
			if err := ie.Format.Format2.Encode(e); err != nil {
				return err
			}
		case PUCCH_Resource_Format_Format3:
			if err := ie.Format.Format3.Encode(e); err != nil {
				return err
			}
		case PUCCH_Resource_Format_Format4:
			if err := ie.Format.Format4.Encode(e); err != nil {
				return err
			}
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.Format.Choice), Constraint: "index out of range"}
		}
	}

	return nil
}

func (ie *PUCCH_Resource) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(pUCCHResourceConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. pucch-ResourceId: ref
	{
		if err := ie.Pucch_ResourceId.Decode(d); err != nil {
			return err
		}
	}

	// 3. startingPRB: ref
	{
		if err := ie.StartingPRB.Decode(d); err != nil {
			return err
		}
	}

	// 4. intraSlotFrequencyHopping: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(pUCCHResourceIntraSlotFrequencyHoppingConstraints)
			if err != nil {
				return err
			}
			ie.IntraSlotFrequencyHopping = &idx
		}
	}

	// 5. secondHopPRB: ref
	{
		if seq.IsComponentPresent(3) {
			ie.SecondHopPRB = new(PRB_Id)
			if err := ie.SecondHopPRB.Decode(d); err != nil {
				return err
			}
		}
	}

	// 6. format: choice
	{
		choiceDec := d.NewChoiceDecoder(pUCCH_ResourceFormatConstraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.Format.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case PUCCH_Resource_Format_Format0:
			ie.Format.Format0 = new(PUCCH_Format0)
			if err := ie.Format.Format0.Decode(d); err != nil {
				return err
			}
		case PUCCH_Resource_Format_Format1:
			ie.Format.Format1 = new(PUCCH_Format1)
			if err := ie.Format.Format1.Decode(d); err != nil {
				return err
			}
		case PUCCH_Resource_Format_Format2:
			ie.Format.Format2 = new(PUCCH_Format2)
			if err := ie.Format.Format2.Decode(d); err != nil {
				return err
			}
		case PUCCH_Resource_Format_Format3:
			ie.Format.Format3 = new(PUCCH_Format3)
			if err := ie.Format.Format3.Decode(d); err != nil {
				return err
			}
		case PUCCH_Resource_Format_Format4:
			ie.Format.Format4 = new(PUCCH_Format4)
			if err := ie.Format.Format4.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
