// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MRDC-Parameters-v1790 (line 22619).

var mRDCParametersV1790Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "intraBandENDC-Support-v1790", Optional: true},
		{Name: "intraBandENDC-Support-UL-v1790", Optional: true},
	},
}

const (
	MRDC_Parameters_v1790_IntraBandENDC_Support_v1790_Non_Contiguous = 0
	MRDC_Parameters_v1790_IntraBandENDC_Support_v1790_Both           = 1
)

var mRDCParametersV1790IntraBandENDCSupportV1790Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MRDC_Parameters_v1790_IntraBandENDC_Support_v1790_Non_Contiguous, MRDC_Parameters_v1790_IntraBandENDC_Support_v1790_Both},
}

const (
	MRDC_Parameters_v1790_IntraBandENDC_Support_UL_v1790_Non_Contiguous = 0
	MRDC_Parameters_v1790_IntraBandENDC_Support_UL_v1790_Both           = 1
)

var mRDCParametersV1790IntraBandENDCSupportULV1790Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MRDC_Parameters_v1790_IntraBandENDC_Support_UL_v1790_Non_Contiguous, MRDC_Parameters_v1790_IntraBandENDC_Support_UL_v1790_Both},
}

type MRDC_Parameters_v1790 struct {
	IntraBandENDC_Support_v1790    *int64
	IntraBandENDC_Support_UL_v1790 *int64
}

func (ie *MRDC_Parameters_v1790) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(mRDCParametersV1790Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.IntraBandENDC_Support_v1790 != nil, ie.IntraBandENDC_Support_UL_v1790 != nil}); err != nil {
		return err
	}

	// 2. intraBandENDC-Support-v1790: enumerated
	{
		if ie.IntraBandENDC_Support_v1790 != nil {
			if err := e.EncodeEnumerated(*ie.IntraBandENDC_Support_v1790, mRDCParametersV1790IntraBandENDCSupportV1790Constraints); err != nil {
				return err
			}
		}
	}

	// 3. intraBandENDC-Support-UL-v1790: enumerated
	{
		if ie.IntraBandENDC_Support_UL_v1790 != nil {
			if err := e.EncodeEnumerated(*ie.IntraBandENDC_Support_UL_v1790, mRDCParametersV1790IntraBandENDCSupportULV1790Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *MRDC_Parameters_v1790) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(mRDCParametersV1790Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. intraBandENDC-Support-v1790: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(mRDCParametersV1790IntraBandENDCSupportV1790Constraints)
			if err != nil {
				return err
			}
			ie.IntraBandENDC_Support_v1790 = &idx
		}
	}

	// 3. intraBandENDC-Support-UL-v1790: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(mRDCParametersV1790IntraBandENDCSupportULV1790Constraints)
			if err != nil {
				return err
			}
			ie.IntraBandENDC_Support_UL_v1790 = &idx
		}
	}

	return nil
}
