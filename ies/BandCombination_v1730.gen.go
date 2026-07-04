// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: BandCombination-v1730 (line 16759).

var bandCombinationV1730Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "ca-ParametersNR-v1730", Optional: true},
		{Name: "ca-ParametersNRDC-v1730", Optional: true},
		{Name: "bandList-v1730", Optional: true},
	},
}

var bandCombinationV1730BandListV1730Constraints = per.SizeRange(1, common.MaxSimultaneousBands)

type BandCombination_v1730 struct {
	Ca_ParametersNR_v1730   *CA_ParametersNR_v1730
	Ca_ParametersNRDC_v1730 *CA_ParametersNRDC_v1730
	BandList_v1730          []BandParameters_v1730
}

func (ie *BandCombination_v1730) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(bandCombinationV1730Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Ca_ParametersNR_v1730 != nil, ie.Ca_ParametersNRDC_v1730 != nil, ie.BandList_v1730 != nil}); err != nil {
		return err
	}

	// 2. ca-ParametersNR-v1730: ref
	{
		if ie.Ca_ParametersNR_v1730 != nil {
			if err := ie.Ca_ParametersNR_v1730.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. ca-ParametersNRDC-v1730: ref
	{
		if ie.Ca_ParametersNRDC_v1730 != nil {
			if err := ie.Ca_ParametersNRDC_v1730.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. bandList-v1730: sequence-of
	{
		if ie.BandList_v1730 != nil {
			seqOf := e.NewSequenceOfEncoder(bandCombinationV1730BandListV1730Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.BandList_v1730))); err != nil {
				return err
			}
			for i := range ie.BandList_v1730 {
				if err := ie.BandList_v1730[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (ie *BandCombination_v1730) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(bandCombinationV1730Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. ca-ParametersNR-v1730: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Ca_ParametersNR_v1730 = new(CA_ParametersNR_v1730)
			if err := ie.Ca_ParametersNR_v1730.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. ca-ParametersNRDC-v1730: ref
	{
		if seq.IsComponentPresent(1) {
			ie.Ca_ParametersNRDC_v1730 = new(CA_ParametersNRDC_v1730)
			if err := ie.Ca_ParametersNRDC_v1730.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. bandList-v1730: sequence-of
	{
		if seq.IsComponentPresent(2) {
			seqOf := d.NewSequenceOfDecoder(bandCombinationV1730BandListV1730Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.BandList_v1730 = make([]BandParameters_v1730, n)
			for i := int64(0); i < n; i++ {
				if err := ie.BandList_v1730[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
