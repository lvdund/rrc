// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MeasAndMobParametersMRDC-Common-v1730 (line 21431).

var measAndMobParametersMRDCCommonV1730Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "independentGapConfig-maxCC-r17"},
	},
}

var measAndMobParametersMRDCCommonV1730IndependentGapConfigMaxCCR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "fr1-Only-r17", Optional: true},
		{Name: "fr2-Only-r17", Optional: true},
		{Name: "fr1-AndFR2-r17", Optional: true},
	},
}

type MeasAndMobParametersMRDC_Common_v1730 struct {
	IndependentGapConfig_MaxCC_r17 struct {
		Fr1_Only_r17   *int64
		Fr2_Only_r17   *int64
		Fr1_AndFR2_r17 *int64
	}
}

func (ie *MeasAndMobParametersMRDC_Common_v1730) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(measAndMobParametersMRDCCommonV1730Constraints)
	_ = seq

	// 1. independentGapConfig-maxCC-r17: sequence
	{
		{
			c := &ie.IndependentGapConfig_MaxCC_r17
			measAndMobParametersMRDCCommonV1730IndependentGapConfigMaxCCR17Seq := e.NewSequenceEncoder(measAndMobParametersMRDCCommonV1730IndependentGapConfigMaxCCR17Constraints)
			if err := measAndMobParametersMRDCCommonV1730IndependentGapConfigMaxCCR17Seq.EncodePreamble([]bool{c.Fr1_Only_r17 != nil, c.Fr2_Only_r17 != nil, c.Fr1_AndFR2_r17 != nil}); err != nil {
				return err
			}
			if c.Fr1_Only_r17 != nil {
				if err := e.EncodeInteger((*c.Fr1_Only_r17), per.Constrained(1, 32)); err != nil {
					return err
				}
			}
			if c.Fr2_Only_r17 != nil {
				if err := e.EncodeInteger((*c.Fr2_Only_r17), per.Constrained(1, 32)); err != nil {
					return err
				}
			}
			if c.Fr1_AndFR2_r17 != nil {
				if err := e.EncodeInteger((*c.Fr1_AndFR2_r17), per.Constrained(1, 32)); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (ie *MeasAndMobParametersMRDC_Common_v1730) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(measAndMobParametersMRDCCommonV1730Constraints)
	_ = seq

	// 1. independentGapConfig-maxCC-r17: sequence
	{
		{
			c := &ie.IndependentGapConfig_MaxCC_r17
			measAndMobParametersMRDCCommonV1730IndependentGapConfigMaxCCR17Seq := d.NewSequenceDecoder(measAndMobParametersMRDCCommonV1730IndependentGapConfigMaxCCR17Constraints)
			if err := measAndMobParametersMRDCCommonV1730IndependentGapConfigMaxCCR17Seq.DecodePreamble(); err != nil {
				return err
			}
			if measAndMobParametersMRDCCommonV1730IndependentGapConfigMaxCCR17Seq.IsComponentPresent(0) {
				c.Fr1_Only_r17 = new(int64)
				v, err := d.DecodeInteger(per.Constrained(1, 32))
				if err != nil {
					return err
				}
				(*c.Fr1_Only_r17) = v
			}
			if measAndMobParametersMRDCCommonV1730IndependentGapConfigMaxCCR17Seq.IsComponentPresent(1) {
				c.Fr2_Only_r17 = new(int64)
				v, err := d.DecodeInteger(per.Constrained(1, 32))
				if err != nil {
					return err
				}
				(*c.Fr2_Only_r17) = v
			}
			if measAndMobParametersMRDCCommonV1730IndependentGapConfigMaxCCR17Seq.IsComponentPresent(2) {
				c.Fr1_AndFR2_r17 = new(int64)
				v, err := d.DecodeInteger(per.Constrained(1, 32))
				if err != nil {
					return err
				}
				(*c.Fr1_AndFR2_r17) = v
			}
		}
	}

	return nil
}
