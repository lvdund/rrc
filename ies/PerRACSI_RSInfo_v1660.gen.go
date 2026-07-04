// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: PerRACSI-RSInfo-v1660 (line 3187).

var perRACSIRSInfoV1660Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "csi-RS-Index-v1660", Optional: true},
	},
}

var perRACSIRSInfoV1660CsiRSIndexV1660Constraints = per.Constrained(1, 96)

type PerRACSI_RSInfo_v1660 struct {
	Csi_RS_Index_v1660 *int64
}

func (ie *PerRACSI_RSInfo_v1660) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(perRACSIRSInfoV1660Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Csi_RS_Index_v1660 != nil}); err != nil {
		return err
	}

	// 2. csi-RS-Index-v1660: integer
	{
		if ie.Csi_RS_Index_v1660 != nil {
			if err := e.EncodeInteger(*ie.Csi_RS_Index_v1660, perRACSIRSInfoV1660CsiRSIndexV1660Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *PerRACSI_RSInfo_v1660) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(perRACSIRSInfoV1660Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. csi-RS-Index-v1660: integer
	{
		if seq.IsComponentPresent(0) {
			v, err := d.DecodeInteger(perRACSIRSInfoV1660CsiRSIndexV1660Constraints)
			if err != nil {
				return err
			}
			ie.Csi_RS_Index_v1660 = &v
		}
	}

	return nil
}
