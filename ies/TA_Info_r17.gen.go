// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: TA-Info-r17 (line 10804).

var tAInfoR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "ta-Common-r17"},
		{Name: "ta-CommonDrift-r17", Optional: true},
		{Name: "ta-CommonDriftVariant-r17", Optional: true},
	},
}

var tAInfoR17TaCommonR17Constraints = per.Constrained(0, 66485757)

var tAInfoR17TaCommonDriftR17Constraints = per.Constrained(-257303, 257303)

var tAInfoR17TaCommonDriftVariantR17Constraints = per.Constrained(0, 28949)

type TA_Info_r17 struct {
	Ta_Common_r17             int64
	Ta_CommonDrift_r17        *int64
	Ta_CommonDriftVariant_r17 *int64
}

func (ie *TA_Info_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(tAInfoR17Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Ta_CommonDrift_r17 != nil, ie.Ta_CommonDriftVariant_r17 != nil}); err != nil {
		return err
	}

	// 2. ta-Common-r17: integer
	{
		if err := e.EncodeInteger(ie.Ta_Common_r17, tAInfoR17TaCommonR17Constraints); err != nil {
			return err
		}
	}

	// 3. ta-CommonDrift-r17: integer
	{
		if ie.Ta_CommonDrift_r17 != nil {
			if err := e.EncodeInteger(*ie.Ta_CommonDrift_r17, tAInfoR17TaCommonDriftR17Constraints); err != nil {
				return err
			}
		}
	}

	// 4. ta-CommonDriftVariant-r17: integer
	{
		if ie.Ta_CommonDriftVariant_r17 != nil {
			if err := e.EncodeInteger(*ie.Ta_CommonDriftVariant_r17, tAInfoR17TaCommonDriftVariantR17Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *TA_Info_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(tAInfoR17Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. ta-Common-r17: integer
	{
		v0, err := d.DecodeInteger(tAInfoR17TaCommonR17Constraints)
		if err != nil {
			return err
		}
		ie.Ta_Common_r17 = v0
	}

	// 3. ta-CommonDrift-r17: integer
	{
		if seq.IsComponentPresent(1) {
			v, err := d.DecodeInteger(tAInfoR17TaCommonDriftR17Constraints)
			if err != nil {
				return err
			}
			ie.Ta_CommonDrift_r17 = &v
		}
	}

	// 4. ta-CommonDriftVariant-r17: integer
	{
		if seq.IsComponentPresent(2) {
			v, err := d.DecodeInteger(tAInfoR17TaCommonDriftVariantR17Constraints)
			if err != nil {
				return err
			}
			ie.Ta_CommonDriftVariant_r17 = &v
		}
	}

	return nil
}
