// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: PUSCH-MutingResources-r19 (line 12558).

var pUSCHMutingResourcesR19Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "symbolPos-r19", Optional: true},
		{Name: "combOffset-r19", Optional: true},
	},
}

var pUSCHMutingResourcesR19SymbolPosR19Constraints = per.SizeRange(1, 2)

var pUSCHMutingResourcesR19CombOffsetR19Constraints = per.Constrained(0, 1)

type PUSCH_MutingResources_r19 struct {
	SymbolPos_r19  []int64
	CombOffset_r19 *int64
}

func (ie *PUSCH_MutingResources_r19) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(pUSCHMutingResourcesR19Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.SymbolPos_r19 != nil, ie.CombOffset_r19 != nil}); err != nil {
		return err
	}

	// 3. symbolPos-r19: sequence-of
	{
		if ie.SymbolPos_r19 != nil {
			seqOf := e.NewSequenceOfEncoder(pUSCHMutingResourcesR19SymbolPosR19Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.SymbolPos_r19))); err != nil {
				return err
			}
			for i := range ie.SymbolPos_r19 {
				if err := e.EncodeInteger(ie.SymbolPos_r19[i], per.Constrained(0, 13)); err != nil {
					return err
				}
			}
		}
	}

	// 4. combOffset-r19: integer
	{
		if ie.CombOffset_r19 != nil {
			if err := e.EncodeInteger(*ie.CombOffset_r19, pUSCHMutingResourcesR19CombOffsetR19Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *PUSCH_MutingResources_r19) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(pUSCHMutingResourcesR19Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. symbolPos-r19: sequence-of
	{
		if seq.IsComponentPresent(0) {
			seqOf := d.NewSequenceOfDecoder(pUSCHMutingResourcesR19SymbolPosR19Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.SymbolPos_r19 = make([]int64, n)
			for i := int64(0); i < n; i++ {
				v, err := d.DecodeInteger(per.Constrained(0, 13))
				if err != nil {
					return err
				}
				ie.SymbolPos_r19[i] = v
			}
		}
	}

	// 4. combOffset-r19: integer
	{
		if seq.IsComponentPresent(1) {
			v, err := d.DecodeInteger(pUSCHMutingResourcesR19CombOffsetR19Constraints)
			if err != nil {
				return err
			}
			ie.CombOffset_r19 = &v
		}
	}

	return nil
}
