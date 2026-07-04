// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SimulSRS-ForAntennaSwitching-r16 (line 18309).

var simulSRSForAntennaSwitchingR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "supportSRS-xTyR-xLessThanY-r16", Optional: true},
		{Name: "supportSRS-xTyR-xEqualToY-r16", Optional: true},
		{Name: "supportSRS-AntennaSwitching-r16", Optional: true},
	},
}

const (
	SimulSRS_ForAntennaSwitching_r16_SupportSRS_XTyR_XLessThanY_r16_Supported = 0
)

var simulSRSForAntennaSwitchingR16SupportSRSXTyRXLessThanYR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SimulSRS_ForAntennaSwitching_r16_SupportSRS_XTyR_XLessThanY_r16_Supported},
}

const (
	SimulSRS_ForAntennaSwitching_r16_SupportSRS_XTyR_XEqualToY_r16_Supported = 0
)

var simulSRSForAntennaSwitchingR16SupportSRSXTyRXEqualToYR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SimulSRS_ForAntennaSwitching_r16_SupportSRS_XTyR_XEqualToY_r16_Supported},
}

const (
	SimulSRS_ForAntennaSwitching_r16_SupportSRS_AntennaSwitching_r16_Supported = 0
)

var simulSRSForAntennaSwitchingR16SupportSRSAntennaSwitchingR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SimulSRS_ForAntennaSwitching_r16_SupportSRS_AntennaSwitching_r16_Supported},
}

type SimulSRS_ForAntennaSwitching_r16 struct {
	SupportSRS_XTyR_XLessThanY_r16  *int64
	SupportSRS_XTyR_XEqualToY_r16   *int64
	SupportSRS_AntennaSwitching_r16 *int64
}

func (ie *SimulSRS_ForAntennaSwitching_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(simulSRSForAntennaSwitchingR16Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.SupportSRS_XTyR_XLessThanY_r16 != nil, ie.SupportSRS_XTyR_XEqualToY_r16 != nil, ie.SupportSRS_AntennaSwitching_r16 != nil}); err != nil {
		return err
	}

	// 2. supportSRS-xTyR-xLessThanY-r16: enumerated
	{
		if ie.SupportSRS_XTyR_XLessThanY_r16 != nil {
			if err := e.EncodeEnumerated(*ie.SupportSRS_XTyR_XLessThanY_r16, simulSRSForAntennaSwitchingR16SupportSRSXTyRXLessThanYR16Constraints); err != nil {
				return err
			}
		}
	}

	// 3. supportSRS-xTyR-xEqualToY-r16: enumerated
	{
		if ie.SupportSRS_XTyR_XEqualToY_r16 != nil {
			if err := e.EncodeEnumerated(*ie.SupportSRS_XTyR_XEqualToY_r16, simulSRSForAntennaSwitchingR16SupportSRSXTyRXEqualToYR16Constraints); err != nil {
				return err
			}
		}
	}

	// 4. supportSRS-AntennaSwitching-r16: enumerated
	{
		if ie.SupportSRS_AntennaSwitching_r16 != nil {
			if err := e.EncodeEnumerated(*ie.SupportSRS_AntennaSwitching_r16, simulSRSForAntennaSwitchingR16SupportSRSAntennaSwitchingR16Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *SimulSRS_ForAntennaSwitching_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(simulSRSForAntennaSwitchingR16Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. supportSRS-xTyR-xLessThanY-r16: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(simulSRSForAntennaSwitchingR16SupportSRSXTyRXLessThanYR16Constraints)
			if err != nil {
				return err
			}
			ie.SupportSRS_XTyR_XLessThanY_r16 = &idx
		}
	}

	// 3. supportSRS-xTyR-xEqualToY-r16: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(simulSRSForAntennaSwitchingR16SupportSRSXTyRXEqualToYR16Constraints)
			if err != nil {
				return err
			}
			ie.SupportSRS_XTyR_XEqualToY_r16 = &idx
		}
	}

	// 4. supportSRS-AntennaSwitching-r16: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(simulSRSForAntennaSwitchingR16SupportSRSAntennaSwitchingR16Constraints)
			if err != nil {
				return err
			}
			ie.SupportSRS_AntennaSwitching_r16 = &idx
		}
	}

	return nil
}
