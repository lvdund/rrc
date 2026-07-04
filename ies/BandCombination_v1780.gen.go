// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: BandCombination-v1780 (line 16780).

var bandCombinationV1780Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "ca-ParametersNR-v1780", Optional: true},
		{Name: "ca-ParametersNRDC-v1780", Optional: true},
		{Name: "bandList-v1780", Optional: true},
		{Name: "mrdc-Parameters-v1780", Optional: true},
	},
}

var bandCombinationV1780BandListV1780Constraints = per.SizeRange(1, common.MaxSimultaneousBands)

type BandCombination_v1780 struct {
	Ca_ParametersNR_v1780   *CA_ParametersNR_v1780
	Ca_ParametersNRDC_v1780 *CA_ParametersNRDC_v1780
	BandList_v1780          []BandParameters_v1780
	Mrdc_Parameters_v1780   *MRDC_Parameters_v1770
}

func (ie *BandCombination_v1780) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(bandCombinationV1780Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Ca_ParametersNR_v1780 != nil, ie.Ca_ParametersNRDC_v1780 != nil, ie.BandList_v1780 != nil, ie.Mrdc_Parameters_v1780 != nil}); err != nil {
		return err
	}

	// 2. ca-ParametersNR-v1780: ref
	{
		if ie.Ca_ParametersNR_v1780 != nil {
			if err := ie.Ca_ParametersNR_v1780.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. ca-ParametersNRDC-v1780: ref
	{
		if ie.Ca_ParametersNRDC_v1780 != nil {
			if err := ie.Ca_ParametersNRDC_v1780.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. bandList-v1780: sequence-of
	{
		if ie.BandList_v1780 != nil {
			seqOf := e.NewSequenceOfEncoder(bandCombinationV1780BandListV1780Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.BandList_v1780))); err != nil {
				return err
			}
			for i := range ie.BandList_v1780 {
				if err := ie.BandList_v1780[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 5. mrdc-Parameters-v1780: ref
	{
		if ie.Mrdc_Parameters_v1780 != nil {
			if err := ie.Mrdc_Parameters_v1780.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *BandCombination_v1780) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(bandCombinationV1780Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. ca-ParametersNR-v1780: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Ca_ParametersNR_v1780 = new(CA_ParametersNR_v1780)
			if err := ie.Ca_ParametersNR_v1780.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. ca-ParametersNRDC-v1780: ref
	{
		if seq.IsComponentPresent(1) {
			ie.Ca_ParametersNRDC_v1780 = new(CA_ParametersNRDC_v1780)
			if err := ie.Ca_ParametersNRDC_v1780.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. bandList-v1780: sequence-of
	{
		if seq.IsComponentPresent(2) {
			seqOf := d.NewSequenceOfDecoder(bandCombinationV1780BandListV1780Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.BandList_v1780 = make([]BandParameters_v1780, n)
			for i := int64(0); i < n; i++ {
				if err := ie.BandList_v1780[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 5. mrdc-Parameters-v1780: ref
	{
		if seq.IsComponentPresent(3) {
			ie.Mrdc_Parameters_v1780 = new(MRDC_Parameters_v1770)
			if err := ie.Mrdc_Parameters_v1780.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
