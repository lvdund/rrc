// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: PLMN-RAN-AreaCell (line 1352).

var pLMNRANAreaCellConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "plmn-Identity", Optional: true},
		{Name: "ran-AreaCells"},
	},
}

var pLMNRANAreaCellRanAreaCellsConstraints = per.SizeRange(1, 32)

type PLMN_RAN_AreaCell struct {
	Plmn_Identity *PLMN_Identity
	Ran_AreaCells []CellIdentity
}

func (ie *PLMN_RAN_AreaCell) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(pLMNRANAreaCellConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Plmn_Identity != nil}); err != nil {
		return err
	}

	// 2. plmn-Identity: ref
	{
		if ie.Plmn_Identity != nil {
			if err := ie.Plmn_Identity.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. ran-AreaCells: sequence-of
	{
		seqOf := e.NewSequenceOfEncoder(pLMNRANAreaCellRanAreaCellsConstraints)
		if err := seqOf.EncodeLength(int64(len(ie.Ran_AreaCells))); err != nil {
			return err
		}
		for i := range ie.Ran_AreaCells {
			if err := ie.Ran_AreaCells[i].Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *PLMN_RAN_AreaCell) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(pLMNRANAreaCellConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. plmn-Identity: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Plmn_Identity = new(PLMN_Identity)
			if err := ie.Plmn_Identity.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. ran-AreaCells: sequence-of
	{
		seqOf := d.NewSequenceOfDecoder(pLMNRANAreaCellRanAreaCellsConstraints)
		n, err := seqOf.DecodeLength()
		if err != nil {
			return err
		}
		ie.Ran_AreaCells = make([]CellIdentity, n)
		for i := int64(0); i < n; i++ {
			if err := ie.Ran_AreaCells[i].Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
