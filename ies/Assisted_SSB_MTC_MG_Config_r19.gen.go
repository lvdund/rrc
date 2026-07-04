// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: Assisted-SSB-MTC-MG-Config-r19 (line 26371).

var assistedSSBMTCMGConfigR19Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "closestLocsToReport-r19"},
		{Name: "refLocList-r19", Optional: true},
	},
}

var assistedSSBMTCMGConfigR19ClosestLocsToReportR19Constraints = per.Constrained(1, 4)

var assistedSSBMTCMGConfigR19RefLocListR19Constraints = per.SizeRange(1, 7)

type Assisted_SSB_MTC_MG_Config_r19 struct {
	ClosestLocsToReport_r19 int64
	RefLocList_r19          []ReferenceLocation_r17
}

func (ie *Assisted_SSB_MTC_MG_Config_r19) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(assistedSSBMTCMGConfigR19Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.RefLocList_r19 != nil}); err != nil {
		return err
	}

	// 2. closestLocsToReport-r19: integer
	{
		if err := e.EncodeInteger(ie.ClosestLocsToReport_r19, assistedSSBMTCMGConfigR19ClosestLocsToReportR19Constraints); err != nil {
			return err
		}
	}

	// 3. refLocList-r19: sequence-of
	{
		if ie.RefLocList_r19 != nil {
			seqOf := e.NewSequenceOfEncoder(assistedSSBMTCMGConfigR19RefLocListR19Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.RefLocList_r19))); err != nil {
				return err
			}
			for i := range ie.RefLocList_r19 {
				if err := ie.RefLocList_r19[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (ie *Assisted_SSB_MTC_MG_Config_r19) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(assistedSSBMTCMGConfigR19Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. closestLocsToReport-r19: integer
	{
		v0, err := d.DecodeInteger(assistedSSBMTCMGConfigR19ClosestLocsToReportR19Constraints)
		if err != nil {
			return err
		}
		ie.ClosestLocsToReport_r19 = v0
	}

	// 3. refLocList-r19: sequence-of
	{
		if seq.IsComponentPresent(1) {
			seqOf := d.NewSequenceOfDecoder(assistedSSBMTCMGConfigR19RefLocListR19Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.RefLocList_r19 = make([]ReferenceLocation_r17, n)
			for i := int64(0); i < n; i++ {
				if err := ie.RefLocList_r19[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
