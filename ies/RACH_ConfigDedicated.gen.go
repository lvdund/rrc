// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: RACH-ConfigDedicated (line 12920).

var rACHConfigDedicatedConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "cfra", Optional: true},
		{Name: "ra-Prioritization", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
	},
}

type RACH_ConfigDedicated struct {
	Cfra                         *CFRA
	Ra_Prioritization            *RA_Prioritization
	Ra_PrioritizationTwoStep_r16 *RA_Prioritization
	Cfra_TwoStep_r16             *CFRA_TwoStep_r16
}

func (ie *RACH_ConfigDedicated) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(rACHConfigDedicatedConstraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.Ra_PrioritizationTwoStep_r16 != nil || ie.Cfra_TwoStep_r16 != nil
	hasExtensions := hasExtGroup0

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Cfra != nil, ie.Ra_Prioritization != nil}); err != nil {
		return err
	}

	// 3. cfra: ref
	{
		if ie.Cfra != nil {
			if err := ie.Cfra.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. ra-Prioritization: ref
	{
		if ie.Ra_Prioritization != nil {
			if err := ie.Ra_Prioritization.Encode(e); err != nil {
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
					{Name: "ra-PrioritizationTwoStep-r16", Optional: true},
					{Name: "cfra-TwoStep-r16", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Ra_PrioritizationTwoStep_r16 != nil, ie.Cfra_TwoStep_r16 != nil}); err != nil {
				return err
			}

			if ie.Ra_PrioritizationTwoStep_r16 != nil {
				if err := ie.Ra_PrioritizationTwoStep_r16.Encode(ex); err != nil {
					return err
				}
			}

			if ie.Cfra_TwoStep_r16 != nil {
				if err := ie.Cfra_TwoStep_r16.Encode(ex); err != nil {
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

func (ie *RACH_ConfigDedicated) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(rACHConfigDedicatedConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. cfra: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Cfra = new(CFRA)
			if err := ie.Cfra.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. ra-Prioritization: ref
	{
		if seq.IsComponentPresent(1) {
			ie.Ra_Prioritization = new(RA_Prioritization)
			if err := ie.Ra_Prioritization.Decode(d); err != nil {
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
				{Name: "ra-PrioritizationTwoStep-r16", Optional: true},
				{Name: "cfra-TwoStep-r16", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.Ra_PrioritizationTwoStep_r16 = new(RA_Prioritization)
			if err := ie.Ra_PrioritizationTwoStep_r16.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(1) {
			ie.Cfra_TwoStep_r16 = new(CFRA_TwoStep_r16)
			if err := ie.Cfra_TwoStep_r16.Decode(dx); err != nil {
				return err
			}
		}
	}

	return nil
}
