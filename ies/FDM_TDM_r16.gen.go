// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: FDM-TDM-r16 (line 13352).

var fDMTDMR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "repetitionScheme-r16"},
		{Name: "startingSymbolOffsetK-r16", Optional: true},
	},
}

const (
	FDM_TDM_r16_RepetitionScheme_r16_FdmSchemeA = 0
	FDM_TDM_r16_RepetitionScheme_r16_FdmSchemeB = 1
	FDM_TDM_r16_RepetitionScheme_r16_TdmSchemeA = 2
)

var fDMTDMR16RepetitionSchemeR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FDM_TDM_r16_RepetitionScheme_r16_FdmSchemeA, FDM_TDM_r16_RepetitionScheme_r16_FdmSchemeB, FDM_TDM_r16_RepetitionScheme_r16_TdmSchemeA},
}

var fDMTDMR16StartingSymbolOffsetKR16Constraints = per.Constrained(0, 7)

type FDM_TDM_r16 struct {
	RepetitionScheme_r16      int64
	StartingSymbolOffsetK_r16 *int64
}

func (ie *FDM_TDM_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(fDMTDMR16Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.StartingSymbolOffsetK_r16 != nil}); err != nil {
		return err
	}

	// 2. repetitionScheme-r16: enumerated
	{
		if err := e.EncodeEnumerated(ie.RepetitionScheme_r16, fDMTDMR16RepetitionSchemeR16Constraints); err != nil {
			return err
		}
	}

	// 3. startingSymbolOffsetK-r16: integer
	{
		if ie.StartingSymbolOffsetK_r16 != nil {
			if err := e.EncodeInteger(*ie.StartingSymbolOffsetK_r16, fDMTDMR16StartingSymbolOffsetKR16Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *FDM_TDM_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(fDMTDMR16Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. repetitionScheme-r16: enumerated
	{
		v0, err := d.DecodeEnumerated(fDMTDMR16RepetitionSchemeR16Constraints)
		if err != nil {
			return err
		}
		ie.RepetitionScheme_r16 = v0
	}

	// 3. startingSymbolOffsetK-r16: integer
	{
		if seq.IsComponentPresent(1) {
			v, err := d.DecodeInteger(fDMTDMR16StartingSymbolOffsetKR16Constraints)
			if err != nil {
				return err
			}
			ie.StartingSymbolOffsetK_r16 = &v
		}
	}

	return nil
}
