// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: Tag2-r18 (line 14722).

var tag2R18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "tag2-Id-r18"},
		{Name: "tag2-flag-r18"},
		{Name: "n-TimingAdvanceOffset2-r18", Optional: true},
	},
}

const (
	Tag2_r18_N_TimingAdvanceOffset2_r18_N0     = 0
	Tag2_r18_N_TimingAdvanceOffset2_r18_N25600 = 1
	Tag2_r18_N_TimingAdvanceOffset2_r18_N39936 = 2
	Tag2_r18_N_TimingAdvanceOffset2_r18_Spare1 = 3
)

var tag2R18NTimingAdvanceOffset2R18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Tag2_r18_N_TimingAdvanceOffset2_r18_N0, Tag2_r18_N_TimingAdvanceOffset2_r18_N25600, Tag2_r18_N_TimingAdvanceOffset2_r18_N39936, Tag2_r18_N_TimingAdvanceOffset2_r18_Spare1},
}

type Tag2_r18 struct {
	Tag2_Id_r18                TAG_Id
	Tag2_Flag_r18              bool
	N_TimingAdvanceOffset2_r18 *int64
}

func (ie *Tag2_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(tag2R18Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.N_TimingAdvanceOffset2_r18 != nil}); err != nil {
		return err
	}

	// 2. tag2-Id-r18: ref
	{
		if err := ie.Tag2_Id_r18.Encode(e); err != nil {
			return err
		}
	}

	// 3. tag2-flag-r18: boolean
	{
		if err := e.EncodeBoolean(ie.Tag2_Flag_r18); err != nil {
			return err
		}
	}

	// 4. n-TimingAdvanceOffset2-r18: enumerated
	{
		if ie.N_TimingAdvanceOffset2_r18 != nil {
			if err := e.EncodeEnumerated(*ie.N_TimingAdvanceOffset2_r18, tag2R18NTimingAdvanceOffset2R18Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *Tag2_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(tag2R18Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. tag2-Id-r18: ref
	{
		if err := ie.Tag2_Id_r18.Decode(d); err != nil {
			return err
		}
	}

	// 3. tag2-flag-r18: boolean
	{
		v1, err := d.DecodeBoolean()
		if err != nil {
			return err
		}
		ie.Tag2_Flag_r18 = v1
	}

	// 4. n-TimingAdvanceOffset2-r18: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(tag2R18NTimingAdvanceOffset2R18Constraints)
			if err != nil {
				return err
			}
			ie.N_TimingAdvanceOffset2_r18 = &idx
		}
	}

	return nil
}
