// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: DL-PRS-Info-r16 (line 15646).

var dLPRSInfoR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "dl-PRS-ID-r16"},
		{Name: "dl-PRS-ResourceSetId-r16"},
		{Name: "dl-PRS-ResourceId-r16", Optional: true},
	},
}

var dLPRSInfoR16DlPRSIDR16Constraints = per.Constrained(0, 255)

var dLPRSInfoR16DlPRSResourceSetIdR16Constraints = per.Constrained(0, 7)

var dLPRSInfoR16DlPRSResourceIdR16Constraints = per.Constrained(0, 63)

type DL_PRS_Info_r16 struct {
	Dl_PRS_ID_r16            int64
	Dl_PRS_ResourceSetId_r16 int64
	Dl_PRS_ResourceId_r16    *int64
}

func (ie *DL_PRS_Info_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(dLPRSInfoR16Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Dl_PRS_ResourceId_r16 != nil}); err != nil {
		return err
	}

	// 2. dl-PRS-ID-r16: integer
	{
		if err := e.EncodeInteger(ie.Dl_PRS_ID_r16, dLPRSInfoR16DlPRSIDR16Constraints); err != nil {
			return err
		}
	}

	// 3. dl-PRS-ResourceSetId-r16: integer
	{
		if err := e.EncodeInteger(ie.Dl_PRS_ResourceSetId_r16, dLPRSInfoR16DlPRSResourceSetIdR16Constraints); err != nil {
			return err
		}
	}

	// 4. dl-PRS-ResourceId-r16: integer
	{
		if ie.Dl_PRS_ResourceId_r16 != nil {
			if err := e.EncodeInteger(*ie.Dl_PRS_ResourceId_r16, dLPRSInfoR16DlPRSResourceIdR16Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *DL_PRS_Info_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(dLPRSInfoR16Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. dl-PRS-ID-r16: integer
	{
		v0, err := d.DecodeInteger(dLPRSInfoR16DlPRSIDR16Constraints)
		if err != nil {
			return err
		}
		ie.Dl_PRS_ID_r16 = v0
	}

	// 3. dl-PRS-ResourceSetId-r16: integer
	{
		v1, err := d.DecodeInteger(dLPRSInfoR16DlPRSResourceSetIdR16Constraints)
		if err != nil {
			return err
		}
		ie.Dl_PRS_ResourceSetId_r16 = v1
	}

	// 4. dl-PRS-ResourceId-r16: integer
	{
		if seq.IsComponentPresent(2) {
			v, err := d.DecodeInteger(dLPRSInfoR16DlPRSResourceIdR16Constraints)
			if err != nil {
				return err
			}
			ie.Dl_PRS_ResourceId_r16 = &v
		}
	}

	return nil
}
