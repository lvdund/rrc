// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: BandCombination-v1540 (line 16660).

var bandCombinationV1540Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "bandList-v1540"},
		{Name: "ca-ParametersNR-v1540", Optional: true},
	},
}

var bandCombinationV1540BandListV1540Constraints = per.SizeRange(1, common.MaxSimultaneousBands)

type BandCombination_v1540 struct {
	BandList_v1540        []BandParameters_v1540
	Ca_ParametersNR_v1540 *CA_ParametersNR_v1540
}

func (ie *BandCombination_v1540) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(bandCombinationV1540Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Ca_ParametersNR_v1540 != nil}); err != nil {
		return err
	}

	// 2. bandList-v1540: sequence-of
	{
		seqOf := e.NewSequenceOfEncoder(bandCombinationV1540BandListV1540Constraints)
		if err := seqOf.EncodeLength(int64(len(ie.BandList_v1540))); err != nil {
			return err
		}
		for i := range ie.BandList_v1540 {
			if err := ie.BandList_v1540[i].Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. ca-ParametersNR-v1540: ref
	{
		if ie.Ca_ParametersNR_v1540 != nil {
			if err := ie.Ca_ParametersNR_v1540.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *BandCombination_v1540) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(bandCombinationV1540Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. bandList-v1540: sequence-of
	{
		seqOf := d.NewSequenceOfDecoder(bandCombinationV1540BandListV1540Constraints)
		n, err := seqOf.DecodeLength()
		if err != nil {
			return err
		}
		ie.BandList_v1540 = make([]BandParameters_v1540, n)
		for i := int64(0); i < n; i++ {
			if err := ie.BandList_v1540[i].Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. ca-ParametersNR-v1540: ref
	{
		if seq.IsComponentPresent(1) {
			ie.Ca_ParametersNR_v1540 = new(CA_ParametersNR_v1540)
			if err := ie.Ca_ParametersNR_v1540.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
