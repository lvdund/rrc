// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MeasResultSCG-Failure (line 9991).

var measResultSCGFailureConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "measResultPerMOList"},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
	},
}

type MeasResultSCG_Failure struct {
	MeasResultPerMOList MeasResultList2NR
	LocationInfo_r16    *LocationInfo_r16
}

func (ie *MeasResultSCG_Failure) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(measResultSCGFailureConstraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.LocationInfo_r16 != nil
	hasExtensions := hasExtGroup0

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. measResultPerMOList: ref
	{
		if err := ie.MeasResultPerMOList.Encode(e); err != nil {
			return err
		}
	}

	if hasExtensions {
		extPresence := []bool{hasExtGroup0}
		var extValues [][]byte

		// Extension group 0:
		if hasExtGroup0 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "locationInfo-r16", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.LocationInfo_r16 != nil}); err != nil {
				return err
			}

			if ie.LocationInfo_r16 != nil {
				if err := ie.LocationInfo_r16.Encode(ex); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		if err := seq.EncodeExtensionAdditions(extPresence, extValues); err != nil {
			return err
		}
	}

	return nil
}

func (ie *MeasResultSCG_Failure) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(measResultSCGFailureConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. measResultPerMOList: ref
	{
		if err := ie.MeasResultPerMOList.Decode(d); err != nil {
			return err
		}
	}

	// Extension additions as open types.
	extValues, err := seq.DecodeExtensionAdditions()
	if err != nil {
		return err
	}
	extValueIndex := 0

	// Extension group 0:
	if seq.IsExtensionPresent(0) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "locationInfo-r16", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.LocationInfo_r16 = new(LocationInfo_r16)
			if err := ie.LocationInfo_r16.Decode(dx); err != nil {
				return err
			}
		}
	}

	return nil
}
