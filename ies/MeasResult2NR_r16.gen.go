// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MeasResult2NR-r16 (line 3409).

var measResult2NRR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "ssbFrequency-r16", Optional: true},
		{Name: "refFreqCSI-RS-r16", Optional: true},
		{Name: "measResultList-r16"},
	},
}

type MeasResult2NR_r16 struct {
	SsbFrequency_r16   *ARFCN_ValueNR
	RefFreqCSI_RS_r16  *ARFCN_ValueNR
	MeasResultList_r16 MeasResultListNR
}

func (ie *MeasResult2NR_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(measResult2NRR16Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.SsbFrequency_r16 != nil, ie.RefFreqCSI_RS_r16 != nil}); err != nil {
		return err
	}

	// 2. ssbFrequency-r16: ref
	{
		if ie.SsbFrequency_r16 != nil {
			if err := ie.SsbFrequency_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. refFreqCSI-RS-r16: ref
	{
		if ie.RefFreqCSI_RS_r16 != nil {
			if err := ie.RefFreqCSI_RS_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. measResultList-r16: ref
	{
		if err := ie.MeasResultList_r16.Encode(e); err != nil {
			return err
		}
	}

	return nil
}

func (ie *MeasResult2NR_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(measResult2NRR16Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. ssbFrequency-r16: ref
	{
		if seq.IsComponentPresent(0) {
			ie.SsbFrequency_r16 = new(ARFCN_ValueNR)
			if err := ie.SsbFrequency_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. refFreqCSI-RS-r16: ref
	{
		if seq.IsComponentPresent(1) {
			ie.RefFreqCSI_RS_r16 = new(ARFCN_ValueNR)
			if err := ie.RefFreqCSI_RS_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. measResultList-r16: ref
	{
		if err := ie.MeasResultList_r16.Decode(d); err != nil {
			return err
		}
	}

	return nil
}
