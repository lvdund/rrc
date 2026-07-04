// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: BandCombination-v1770 (line 16774).

var bandCombinationV1770Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "bandList-v1770"},
		{Name: "mrdc-Parameters-v1770", Optional: true},
		{Name: "ca-ParametersNR-v1770", Optional: true},
	},
}

var bandCombinationV1770BandListV1770Constraints = per.SizeRange(1, common.MaxSimultaneousBands)

type BandCombination_v1770 struct {
	BandList_v1770        []BandParameters_v1770
	Mrdc_Parameters_v1770 *MRDC_Parameters_v1770
	Ca_ParametersNR_v1770 *CA_ParametersNR_v1770
}

func (ie *BandCombination_v1770) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(bandCombinationV1770Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Mrdc_Parameters_v1770 != nil, ie.Ca_ParametersNR_v1770 != nil}); err != nil {
		return err
	}

	// 2. bandList-v1770: sequence-of
	{
		seqOf := e.NewSequenceOfEncoder(bandCombinationV1770BandListV1770Constraints)
		if err := seqOf.EncodeLength(int64(len(ie.BandList_v1770))); err != nil {
			return err
		}
		for i := range ie.BandList_v1770 {
			if err := ie.BandList_v1770[i].Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. mrdc-Parameters-v1770: ref
	{
		if ie.Mrdc_Parameters_v1770 != nil {
			if err := ie.Mrdc_Parameters_v1770.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. ca-ParametersNR-v1770: ref
	{
		if ie.Ca_ParametersNR_v1770 != nil {
			if err := ie.Ca_ParametersNR_v1770.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *BandCombination_v1770) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(bandCombinationV1770Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. bandList-v1770: sequence-of
	{
		seqOf := d.NewSequenceOfDecoder(bandCombinationV1770BandListV1770Constraints)
		n, err := seqOf.DecodeLength()
		if err != nil {
			return err
		}
		ie.BandList_v1770 = make([]BandParameters_v1770, n)
		for i := int64(0); i < n; i++ {
			if err := ie.BandList_v1770[i].Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. mrdc-Parameters-v1770: ref
	{
		if seq.IsComponentPresent(1) {
			ie.Mrdc_Parameters_v1770 = new(MRDC_Parameters_v1770)
			if err := ie.Mrdc_Parameters_v1770.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. ca-ParametersNR-v1770: ref
	{
		if seq.IsComponentPresent(2) {
			ie.Ca_ParametersNR_v1770 = new(CA_ParametersNR_v1770)
			if err := ie.Ca_ParametersNR_v1770.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
