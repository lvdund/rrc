// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: AdditionalRACH-Config-r17 (line 5373).

var additionalRACHConfigR17Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "rach-ConfigCommon-r17", Optional: true},
		{Name: "msgA-ConfigCommon-r17", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
	},
}

type AdditionalRACH_Config_r17 struct {
	Rach_ConfigCommon_r17    *RACH_ConfigCommon
	MsgA_ConfigCommon_r17    *MsgA_ConfigCommon_r16
	Sbfd_RACH_DualConfig_r19 *SBFD_RACH_DualConfig_r19
}

func (ie *AdditionalRACH_Config_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(additionalRACHConfigR17Constraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.Sbfd_RACH_DualConfig_r19 != nil
	hasExtensions := hasExtGroup0

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Rach_ConfigCommon_r17 != nil, ie.MsgA_ConfigCommon_r17 != nil}); err != nil {
		return err
	}

	// 3. rach-ConfigCommon-r17: ref
	{
		if ie.Rach_ConfigCommon_r17 != nil {
			if err := ie.Rach_ConfigCommon_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. msgA-ConfigCommon-r17: ref
	{
		if ie.MsgA_ConfigCommon_r17 != nil {
			if err := ie.MsgA_ConfigCommon_r17.Encode(e); err != nil {
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
					{Name: "sbfd-RACH-DualConfig-r19", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Sbfd_RACH_DualConfig_r19 != nil}); err != nil {
				return err
			}

			if ie.Sbfd_RACH_DualConfig_r19 != nil {
				if err := ie.Sbfd_RACH_DualConfig_r19.Encode(ex); err != nil {
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

func (ie *AdditionalRACH_Config_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(additionalRACHConfigR17Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. rach-ConfigCommon-r17: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Rach_ConfigCommon_r17 = new(RACH_ConfigCommon)
			if err := ie.Rach_ConfigCommon_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. msgA-ConfigCommon-r17: ref
	{
		if seq.IsComponentPresent(1) {
			ie.MsgA_ConfigCommon_r17 = new(MsgA_ConfigCommon_r16)
			if err := ie.MsgA_ConfigCommon_r17.Decode(d); err != nil {
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
				{Name: "sbfd-RACH-DualConfig-r19", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.Sbfd_RACH_DualConfig_r19 = new(SBFD_RACH_DualConfig_r19)
			if err := ie.Sbfd_RACH_DualConfig_r19.Decode(dx); err != nil {
				return err
			}
		}
	}

	return nil
}
