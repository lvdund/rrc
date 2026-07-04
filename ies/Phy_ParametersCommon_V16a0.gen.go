// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: Phy-ParametersCommon-v16a0 (line 23075).

var phyParametersCommonV16a0Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "srs-PeriodicityAndOffsetExt-r16", Optional: true},
	},
}

const (
	Phy_ParametersCommon_V16a0_Srs_PeriodicityAndOffsetExt_r16_Supported = 0
)

var phyParametersCommonV16a0SrsPeriodicityAndOffsetExtR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersCommon_V16a0_Srs_PeriodicityAndOffsetExt_r16_Supported},
}

type Phy_ParametersCommon_V16a0 struct {
	Srs_PeriodicityAndOffsetExt_r16 *int64
}

func (ie *Phy_ParametersCommon_V16a0) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(phyParametersCommonV16a0Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Srs_PeriodicityAndOffsetExt_r16 != nil}); err != nil {
		return err
	}

	// 2. srs-PeriodicityAndOffsetExt-r16: enumerated
	{
		if ie.Srs_PeriodicityAndOffsetExt_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Srs_PeriodicityAndOffsetExt_r16, phyParametersCommonV16a0SrsPeriodicityAndOffsetExtR16Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *Phy_ParametersCommon_V16a0) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(phyParametersCommonV16a0Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. srs-PeriodicityAndOffsetExt-r16: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(phyParametersCommonV16a0SrsPeriodicityAndOffsetExtR16Constraints)
			if err != nil {
				return err
			}
			ie.Srs_PeriodicityAndOffsetExt_r16 = &idx
		}
	}

	return nil
}
