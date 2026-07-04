// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: P0-PUSCH-Set-r16 (line 12651).

var p0PUSCHSetR16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "p0-PUSCH-SetId-r16"},
		{Name: "p0-List-r16", Optional: true},
	},
}

var p0PUSCHSetR16P0ListR16Constraints = per.SizeRange(1, common.MaxNrofP0_PUSCH_Set_r16)

type P0_PUSCH_Set_r16 struct {
	P0_PUSCH_SetId_r16 P0_PUSCH_SetId_r16
	P0_List_r16        []P0_PUSCH_r16
}

func (ie *P0_PUSCH_Set_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(p0PUSCHSetR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.P0_List_r16 != nil}); err != nil {
		return err
	}

	// 3. p0-PUSCH-SetId-r16: ref
	{
		if err := ie.P0_PUSCH_SetId_r16.Encode(e); err != nil {
			return err
		}
	}

	// 4. p0-List-r16: sequence-of
	{
		if ie.P0_List_r16 != nil {
			seqOf := e.NewSequenceOfEncoder(p0PUSCHSetR16P0ListR16Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.P0_List_r16))); err != nil {
				return err
			}
			for i := range ie.P0_List_r16 {
				if err := ie.P0_List_r16[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (ie *P0_PUSCH_Set_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(p0PUSCHSetR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. p0-PUSCH-SetId-r16: ref
	{
		if err := ie.P0_PUSCH_SetId_r16.Decode(d); err != nil {
			return err
		}
	}

	// 4. p0-List-r16: sequence-of
	{
		if seq.IsComponentPresent(1) {
			seqOf := d.NewSequenceOfDecoder(p0PUSCHSetR16P0ListR16Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.P0_List_r16 = make([]P0_PUSCH_r16, n)
			for i := int64(0); i < n; i++ {
				if err := ie.P0_List_r16[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
