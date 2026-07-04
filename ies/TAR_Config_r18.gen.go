// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: TAR-Config-r18 (line 15973).

var tARConfigR18Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "offsetThresholdTA-r18", Optional: true},
		{Name: "timingAdvanceSR-r18", Optional: true},
	},
}

var tARConfigR18OffsetThresholdTAR18Constraints = per.Constrained(1, 56)

const (
	TAR_Config_r18_TimingAdvanceSR_r18_Enabled = 0
)

var tARConfigR18TimingAdvanceSRR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{TAR_Config_r18_TimingAdvanceSR_r18_Enabled},
}

type TAR_Config_r18 struct {
	OffsetThresholdTA_r18 *int64
	TimingAdvanceSR_r18   *int64
}

func (ie *TAR_Config_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(tARConfigR18Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.OffsetThresholdTA_r18 != nil, ie.TimingAdvanceSR_r18 != nil}); err != nil {
		return err
	}

	// 3. offsetThresholdTA-r18: integer
	{
		if ie.OffsetThresholdTA_r18 != nil {
			if err := e.EncodeInteger(*ie.OffsetThresholdTA_r18, tARConfigR18OffsetThresholdTAR18Constraints); err != nil {
				return err
			}
		}
	}

	// 4. timingAdvanceSR-r18: enumerated
	{
		if ie.TimingAdvanceSR_r18 != nil {
			if err := e.EncodeEnumerated(*ie.TimingAdvanceSR_r18, tARConfigR18TimingAdvanceSRR18Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *TAR_Config_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(tARConfigR18Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. offsetThresholdTA-r18: integer
	{
		if seq.IsComponentPresent(0) {
			v, err := d.DecodeInteger(tARConfigR18OffsetThresholdTAR18Constraints)
			if err != nil {
				return err
			}
			ie.OffsetThresholdTA_r18 = &v
		}
	}

	// 4. timingAdvanceSR-r18: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(tARConfigR18TimingAdvanceSRR18Constraints)
			if err != nil {
				return err
			}
			ie.TimingAdvanceSR_r18 = &idx
		}
	}

	return nil
}
