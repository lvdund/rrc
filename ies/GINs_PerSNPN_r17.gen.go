// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: GINs-PerSNPN-r17 (line 4502).

var gINsPerSNPNR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "supportedGINs-r17", Optional: true},
	},
}

var gINsPerSNPNR17SupportedGINsR17Constraints = per.SizeRange(1, common.MaxGIN_r17)

type GINs_PerSNPN_r17 struct {
	SupportedGINs_r17 *per.BitString
}

func (ie *GINs_PerSNPN_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(gINsPerSNPNR17Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.SupportedGINs_r17 != nil}); err != nil {
		return err
	}

	// 2. supportedGINs-r17: bit-string
	{
		if ie.SupportedGINs_r17 != nil {
			if err := e.EncodeBitString(*ie.SupportedGINs_r17, gINsPerSNPNR17SupportedGINsR17Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *GINs_PerSNPN_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(gINsPerSNPNR17Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. supportedGINs-r17: bit-string
	{
		if seq.IsComponentPresent(0) {
			v, err := d.DecodeBitString(gINsPerSNPNR17SupportedGINsR17Constraints)
			if err != nil {
				return err
			}
			ie.SupportedGINs_r17 = &v
		}
	}

	return nil
}
