// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: BandCombination-v1560 (line 16668).

var bandCombinationV1560Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "ne-DC-BC", Optional: true},
		{Name: "ca-ParametersNRDC", Optional: true},
		{Name: "ca-ParametersEUTRA-v1560", Optional: true},
		{Name: "ca-ParametersNR-v1560", Optional: true},
	},
}

const (
	BandCombination_v1560_Ne_DC_BC_Supported = 0
)

var bandCombinationV1560NeDCBCConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandCombination_v1560_Ne_DC_BC_Supported},
}

type BandCombination_v1560 struct {
	Ne_DC_BC                 *int64
	Ca_ParametersNRDC        *CA_ParametersNRDC
	Ca_ParametersEUTRA_v1560 *CA_ParametersEUTRA_v1560
	Ca_ParametersNR_v1560    *CA_ParametersNR_v1560
}

func (ie *BandCombination_v1560) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(bandCombinationV1560Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Ne_DC_BC != nil, ie.Ca_ParametersNRDC != nil, ie.Ca_ParametersEUTRA_v1560 != nil, ie.Ca_ParametersNR_v1560 != nil}); err != nil {
		return err
	}

	// 2. ne-DC-BC: enumerated
	{
		if ie.Ne_DC_BC != nil {
			if err := e.EncodeEnumerated(*ie.Ne_DC_BC, bandCombinationV1560NeDCBCConstraints); err != nil {
				return err
			}
		}
	}

	// 3. ca-ParametersNRDC: ref
	{
		if ie.Ca_ParametersNRDC != nil {
			if err := ie.Ca_ParametersNRDC.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. ca-ParametersEUTRA-v1560: ref
	{
		if ie.Ca_ParametersEUTRA_v1560 != nil {
			if err := ie.Ca_ParametersEUTRA_v1560.Encode(e); err != nil {
				return err
			}
		}
	}

	// 5. ca-ParametersNR-v1560: ref
	{
		if ie.Ca_ParametersNR_v1560 != nil {
			if err := ie.Ca_ParametersNR_v1560.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *BandCombination_v1560) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(bandCombinationV1560Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. ne-DC-BC: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(bandCombinationV1560NeDCBCConstraints)
			if err != nil {
				return err
			}
			ie.Ne_DC_BC = &idx
		}
	}

	// 3. ca-ParametersNRDC: ref
	{
		if seq.IsComponentPresent(1) {
			ie.Ca_ParametersNRDC = new(CA_ParametersNRDC)
			if err := ie.Ca_ParametersNRDC.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. ca-ParametersEUTRA-v1560: ref
	{
		if seq.IsComponentPresent(2) {
			ie.Ca_ParametersEUTRA_v1560 = new(CA_ParametersEUTRA_v1560)
			if err := ie.Ca_ParametersEUTRA_v1560.Decode(d); err != nil {
				return err
			}
		}
	}

	// 5. ca-ParametersNR-v1560: ref
	{
		if seq.IsComponentPresent(3) {
			ie.Ca_ParametersNR_v1560 = new(CA_ParametersNR_v1560)
			if err := ie.Ca_ParametersNR_v1560.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
