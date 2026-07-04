// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: RAN-AreaConfig (line 1364).

var rANAreaConfigConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "trackingAreaCode"},
		{Name: "ran-AreaCodeList", Optional: true},
	},
}

var rANAreaConfigRanAreaCodeListConstraints = per.SizeRange(1, 32)

type RAN_AreaConfig struct {
	TrackingAreaCode TrackingAreaCode
	Ran_AreaCodeList []RAN_AreaCode
}

func (ie *RAN_AreaConfig) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(rANAreaConfigConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Ran_AreaCodeList != nil}); err != nil {
		return err
	}

	// 2. trackingAreaCode: ref
	{
		if err := ie.TrackingAreaCode.Encode(e); err != nil {
			return err
		}
	}

	// 3. ran-AreaCodeList: sequence-of
	{
		if ie.Ran_AreaCodeList != nil {
			seqOf := e.NewSequenceOfEncoder(rANAreaConfigRanAreaCodeListConstraints)
			if err := seqOf.EncodeLength(int64(len(ie.Ran_AreaCodeList))); err != nil {
				return err
			}
			for i := range ie.Ran_AreaCodeList {
				if err := ie.Ran_AreaCodeList[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (ie *RAN_AreaConfig) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(rANAreaConfigConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. trackingAreaCode: ref
	{
		if err := ie.TrackingAreaCode.Decode(d); err != nil {
			return err
		}
	}

	// 3. ran-AreaCodeList: sequence-of
	{
		if seq.IsComponentPresent(1) {
			seqOf := d.NewSequenceOfDecoder(rANAreaConfigRanAreaCodeListConstraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Ran_AreaCodeList = make([]RAN_AreaCode, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Ran_AreaCodeList[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
