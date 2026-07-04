// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MeasResult3NR-r19 (line 3532).

var measResult3NRR19Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "ssbFrequency-r19", Optional: true},
		{Name: "l1-MeasResultList-r19", Optional: true},
	},
}

type MeasResult3NR_r19 struct {
	SsbFrequency_r19      *ARFCN_ValueNR
	L1_MeasResultList_r19 *L1_MeasResultList_r19
}

func (ie *MeasResult3NR_r19) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(measResult3NRR19Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.SsbFrequency_r19 != nil, ie.L1_MeasResultList_r19 != nil}); err != nil {
		return err
	}

	// 3. ssbFrequency-r19: ref
	{
		if ie.SsbFrequency_r19 != nil {
			if err := ie.SsbFrequency_r19.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. l1-MeasResultList-r19: ref
	{
		if ie.L1_MeasResultList_r19 != nil {
			if err := ie.L1_MeasResultList_r19.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *MeasResult3NR_r19) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(measResult3NRR19Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. ssbFrequency-r19: ref
	{
		if seq.IsComponentPresent(0) {
			ie.SsbFrequency_r19 = new(ARFCN_ValueNR)
			if err := ie.SsbFrequency_r19.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. l1-MeasResultList-r19: ref
	{
		if seq.IsComponentPresent(1) {
			ie.L1_MeasResultList_r19 = new(L1_MeasResultList_r19)
			if err := ie.L1_MeasResultList_r19.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
