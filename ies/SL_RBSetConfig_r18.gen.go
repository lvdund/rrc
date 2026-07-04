// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SL-RBSetConfig-r18 (line 27794).

var sLRBSetConfigR18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-RBSetIndex-r18"},
		{Name: "sl-NumOfSSSBRepetition-r18", Optional: true},
		{Name: "sl-GapBetweenSSSBRepetition-r18", Optional: true},
	},
}

var sLRBSetConfigR18SlRBSetIndexR18Constraints = per.Constrained(0, 4)

var sLRBSetConfigR18SlNumOfSSSBRepetitionR18Constraints = per.Constrained(2, 9)

var sLRBSetConfigR18SlGapBetweenSSSBRepetitionR18Constraints = per.Constrained(1, 84)

type SL_RBSetConfig_r18 struct {
	Sl_RBSetIndex_r18               int64
	Sl_NumOfSSSBRepetition_r18      *int64
	Sl_GapBetweenSSSBRepetition_r18 *int64
}

func (ie *SL_RBSetConfig_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sLRBSetConfigR18Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Sl_NumOfSSSBRepetition_r18 != nil, ie.Sl_GapBetweenSSSBRepetition_r18 != nil}); err != nil {
		return err
	}

	// 2. sl-RBSetIndex-r18: integer
	{
		if err := e.EncodeInteger(ie.Sl_RBSetIndex_r18, sLRBSetConfigR18SlRBSetIndexR18Constraints); err != nil {
			return err
		}
	}

	// 3. sl-NumOfSSSBRepetition-r18: integer
	{
		if ie.Sl_NumOfSSSBRepetition_r18 != nil {
			if err := e.EncodeInteger(*ie.Sl_NumOfSSSBRepetition_r18, sLRBSetConfigR18SlNumOfSSSBRepetitionR18Constraints); err != nil {
				return err
			}
		}
	}

	// 4. sl-GapBetweenSSSBRepetition-r18: integer
	{
		if ie.Sl_GapBetweenSSSBRepetition_r18 != nil {
			if err := e.EncodeInteger(*ie.Sl_GapBetweenSSSBRepetition_r18, sLRBSetConfigR18SlGapBetweenSSSBRepetitionR18Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *SL_RBSetConfig_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sLRBSetConfigR18Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. sl-RBSetIndex-r18: integer
	{
		v0, err := d.DecodeInteger(sLRBSetConfigR18SlRBSetIndexR18Constraints)
		if err != nil {
			return err
		}
		ie.Sl_RBSetIndex_r18 = v0
	}

	// 3. sl-NumOfSSSBRepetition-r18: integer
	{
		if seq.IsComponentPresent(1) {
			v, err := d.DecodeInteger(sLRBSetConfigR18SlNumOfSSSBRepetitionR18Constraints)
			if err != nil {
				return err
			}
			ie.Sl_NumOfSSSBRepetition_r18 = &v
		}
	}

	// 4. sl-GapBetweenSSSBRepetition-r18: integer
	{
		if seq.IsComponentPresent(2) {
			v, err := d.DecodeInteger(sLRBSetConfigR18SlGapBetweenSSSBRepetitionR18Constraints)
			if err != nil {
				return err
			}
			ie.Sl_GapBetweenSSSBRepetition_r18 = &v
		}
	}

	return nil
}
