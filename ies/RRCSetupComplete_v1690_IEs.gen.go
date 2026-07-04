// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: RRCSetupComplete-v1690-IEs (line 1739).

var rRCSetupCompleteV1690IEsConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "ul-RRC-Segmentation-r16", Optional: true},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

const (
	RRCSetupComplete_v1690_IEs_Ul_RRC_Segmentation_r16_True = 0
)

var rRCSetupCompleteV1690IEsUlRRCSegmentationR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RRCSetupComplete_v1690_IEs_Ul_RRC_Segmentation_r16_True},
}

type RRCSetupComplete_v1690_IEs struct {
	Ul_RRC_Segmentation_r16 *int64
	NonCriticalExtension    *RRCSetupComplete_v1700_IEs
}

func (ie *RRCSetupComplete_v1690_IEs) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(rRCSetupCompleteV1690IEsConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Ul_RRC_Segmentation_r16 != nil, ie.NonCriticalExtension != nil}); err != nil {
		return err
	}

	// 2. ul-RRC-Segmentation-r16: enumerated
	{
		if ie.Ul_RRC_Segmentation_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Ul_RRC_Segmentation_r16, rRCSetupCompleteV1690IEsUlRRCSegmentationR16Constraints); err != nil {
				return err
			}
		}
	}

	// 3. nonCriticalExtension: ref
	{
		if ie.NonCriticalExtension != nil {
			if err := ie.NonCriticalExtension.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *RRCSetupComplete_v1690_IEs) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(rRCSetupCompleteV1690IEsConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. ul-RRC-Segmentation-r16: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(rRCSetupCompleteV1690IEsUlRRCSegmentationR16Constraints)
			if err != nil {
				return err
			}
			ie.Ul_RRC_Segmentation_r16 = &idx
		}
	}

	// 3. nonCriticalExtension: ref
	{
		if seq.IsComponentPresent(1) {
			ie.NonCriticalExtension = new(RRCSetupComplete_v1700_IEs)
			if err := ie.NonCriticalExtension.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
