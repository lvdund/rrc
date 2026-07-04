// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MBSInterestIndication-v1800 (line 654).

var mBSInterestIndicationV1800Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "mbs-NonServingInfoList-r18", Optional: true},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

type MBSInterestIndication_v1800 struct {
	Mbs_NonServingInfoList_r18 *MBS_NonServingInfoList_r18
	NonCriticalExtension       *struct{}
}

func (ie *MBSInterestIndication_v1800) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(mBSInterestIndicationV1800Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Mbs_NonServingInfoList_r18 != nil, ie.NonCriticalExtension != nil}); err != nil {
		return err
	}

	// 2. mbs-NonServingInfoList-r18: ref
	{
		if ie.Mbs_NonServingInfoList_r18 != nil {
			if err := ie.Mbs_NonServingInfoList_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. nonCriticalExtension: sequence
	{
		if ie.NonCriticalExtension != nil {
			// empty SEQUENCE {}: no content to encode.
		}
	}

	return nil
}

func (ie *MBSInterestIndication_v1800) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(mBSInterestIndicationV1800Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. mbs-NonServingInfoList-r18: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Mbs_NonServingInfoList_r18 = new(MBS_NonServingInfoList_r18)
			if err := ie.Mbs_NonServingInfoList_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. nonCriticalExtension: sequence
	{
		if seq.IsComponentPresent(1) {
			ie.NonCriticalExtension = &struct{}{}
		}
	}

	return nil
}
