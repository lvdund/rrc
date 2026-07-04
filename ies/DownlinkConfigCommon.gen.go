// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: DownlinkConfigCommon (line 7817).

var downlinkConfigCommonConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "frequencyInfoDL", Optional: true},
		{Name: "initialDownlinkBWP", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
	},
}

type DownlinkConfigCommon struct {
	FrequencyInfoDL               *FrequencyInfoDL
	InitialDownlinkBWP            *BWP_DownlinkCommon
	InitialDownlinkBWP_RedCap_r17 *BWP_DownlinkCommon
}

func (ie *DownlinkConfigCommon) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(downlinkConfigCommonConstraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.InitialDownlinkBWP_RedCap_r17 != nil
	hasExtensions := hasExtGroup0

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.FrequencyInfoDL != nil, ie.InitialDownlinkBWP != nil}); err != nil {
		return err
	}

	// 3. frequencyInfoDL: ref
	{
		if ie.FrequencyInfoDL != nil {
			if err := ie.FrequencyInfoDL.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. initialDownlinkBWP: ref
	{
		if ie.InitialDownlinkBWP != nil {
			if err := ie.InitialDownlinkBWP.Encode(e); err != nil {
				return err
			}
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
					{Name: "initialDownlinkBWP-RedCap-r17", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.InitialDownlinkBWP_RedCap_r17 != nil}); err != nil {
				return err
			}

			if ie.InitialDownlinkBWP_RedCap_r17 != nil {
				if err := ie.InitialDownlinkBWP_RedCap_r17.Encode(ex); err != nil {
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

func (ie *DownlinkConfigCommon) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(downlinkConfigCommonConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. frequencyInfoDL: ref
	{
		if seq.IsComponentPresent(0) {
			ie.FrequencyInfoDL = new(FrequencyInfoDL)
			if err := ie.FrequencyInfoDL.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. initialDownlinkBWP: ref
	{
		if seq.IsComponentPresent(1) {
			ie.InitialDownlinkBWP = new(BWP_DownlinkCommon)
			if err := ie.InitialDownlinkBWP.Decode(d); err != nil {
				return err
			}
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
				{Name: "initialDownlinkBWP-RedCap-r17", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.InitialDownlinkBWP_RedCap_r17 = new(BWP_DownlinkCommon)
			if err := ie.InitialDownlinkBWP_RedCap_r17.Decode(dx); err != nil {
				return err
			}
		}
	}

	return nil
}
