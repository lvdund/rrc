// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SL-QoS-Profile-r16 (line 27749).

var sLQoSProfileR16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-PQI-r16", Optional: true},
		{Name: "sl-GFBR-r16", Optional: true},
		{Name: "sl-MFBR-r16", Optional: true},
		{Name: "sl-Range-r16", Optional: true},
	},
}

var sLQoSProfileR16SlGFBRR16Constraints = per.Constrained(0, 4000000000)

var sLQoSProfileR16SlMFBRR16Constraints = per.Constrained(0, 4000000000)

var sLQoSProfileR16SlRangeR16Constraints = per.Constrained(1, 1000)

type SL_QoS_Profile_r16 struct {
	Sl_PQI_r16   *SL_PQI_r16
	Sl_GFBR_r16  *int64
	Sl_MFBR_r16  *int64
	Sl_Range_r16 *int64
}

func (ie *SL_QoS_Profile_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sLQoSProfileR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Sl_PQI_r16 != nil, ie.Sl_GFBR_r16 != nil, ie.Sl_MFBR_r16 != nil, ie.Sl_Range_r16 != nil}); err != nil {
		return err
	}

	// 3. sl-PQI-r16: ref
	{
		if ie.Sl_PQI_r16 != nil {
			if err := ie.Sl_PQI_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. sl-GFBR-r16: integer
	{
		if ie.Sl_GFBR_r16 != nil {
			if err := e.EncodeInteger(*ie.Sl_GFBR_r16, sLQoSProfileR16SlGFBRR16Constraints); err != nil {
				return err
			}
		}
	}

	// 5. sl-MFBR-r16: integer
	{
		if ie.Sl_MFBR_r16 != nil {
			if err := e.EncodeInteger(*ie.Sl_MFBR_r16, sLQoSProfileR16SlMFBRR16Constraints); err != nil {
				return err
			}
		}
	}

	// 6. sl-Range-r16: integer
	{
		if ie.Sl_Range_r16 != nil {
			if err := e.EncodeInteger(*ie.Sl_Range_r16, sLQoSProfileR16SlRangeR16Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *SL_QoS_Profile_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sLQoSProfileR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. sl-PQI-r16: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Sl_PQI_r16 = new(SL_PQI_r16)
			if err := ie.Sl_PQI_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. sl-GFBR-r16: integer
	{
		if seq.IsComponentPresent(1) {
			v, err := d.DecodeInteger(sLQoSProfileR16SlGFBRR16Constraints)
			if err != nil {
				return err
			}
			ie.Sl_GFBR_r16 = &v
		}
	}

	// 5. sl-MFBR-r16: integer
	{
		if seq.IsComponentPresent(2) {
			v, err := d.DecodeInteger(sLQoSProfileR16SlMFBRR16Constraints)
			if err != nil {
				return err
			}
			ie.Sl_MFBR_r16 = &v
		}
	}

	// 6. sl-Range-r16: integer
	{
		if seq.IsComponentPresent(3) {
			v, err := d.DecodeInteger(sLQoSProfileR16SlRangeR16Constraints)
			if err != nil {
				return err
			}
			ie.Sl_Range_r16 = &v
		}
	}

	return nil
}
