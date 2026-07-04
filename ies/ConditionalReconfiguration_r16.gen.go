// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: ConditionalReconfiguration-r16 (line 6503).

var conditionalReconfigurationR16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "attemptCondReconfig-r16", Optional: true},
		{Name: "condReconfigToRemoveList-r16", Optional: true},
		{Name: "condReconfigToAddModList-r16", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
	},
}

const (
	ConditionalReconfiguration_r16_AttemptCondReconfig_r16_True = 0
)

var conditionalReconfigurationR16AttemptCondReconfigR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{ConditionalReconfiguration_r16_AttemptCondReconfig_r16_True},
}

var conditionalReconfigurationR16ExtScpacReferenceConfigurationR18Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	ConditionalReconfiguration_r16_Ext_Scpac_ReferenceConfiguration_r18_Release = 0
	ConditionalReconfiguration_r16_Ext_Scpac_ReferenceConfiguration_r18_Setup   = 1
)

type ConditionalReconfiguration_r16 struct {
	AttemptCondReconfig_r16          *int64
	CondReconfigToRemoveList_r16     *CondReconfigToRemoveList_r16
	CondReconfigToAddModList_r16     *CondReconfigToAddModList_r16
	Scpac_ReferenceConfiguration_r18 *struct {
		Choice  int
		Release *struct{}
		Setup   *ReferenceConfiguration_r18
	}
	ServingSecurityCellSetId_r18 *SecurityCellSetId_r18
	Sk_CounterConfiguration_r18  *SK_CounterConfiguration_r18
}

func (ie *ConditionalReconfiguration_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(conditionalReconfigurationR16Constraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.Scpac_ReferenceConfiguration_r18 != nil || ie.ServingSecurityCellSetId_r18 != nil || ie.Sk_CounterConfiguration_r18 != nil
	hasExtensions := hasExtGroup0

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.AttemptCondReconfig_r16 != nil, ie.CondReconfigToRemoveList_r16 != nil, ie.CondReconfigToAddModList_r16 != nil}); err != nil {
		return err
	}

	// 3. attemptCondReconfig-r16: enumerated
	{
		if ie.AttemptCondReconfig_r16 != nil {
			if err := e.EncodeEnumerated(*ie.AttemptCondReconfig_r16, conditionalReconfigurationR16AttemptCondReconfigR16Constraints); err != nil {
				return err
			}
		}
	}

	// 4. condReconfigToRemoveList-r16: ref
	{
		if ie.CondReconfigToRemoveList_r16 != nil {
			if err := ie.CondReconfigToRemoveList_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 5. condReconfigToAddModList-r16: ref
	{
		if ie.CondReconfigToAddModList_r16 != nil {
			if err := ie.CondReconfigToAddModList_r16.Encode(e); err != nil {
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
					{Name: "scpac-ReferenceConfiguration-r18", Optional: true},
					{Name: "servingSecurityCellSetId-r18", Optional: true},
					{Name: "sk-CounterConfiguration-r18", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Scpac_ReferenceConfiguration_r18 != nil, ie.ServingSecurityCellSetId_r18 != nil, ie.Sk_CounterConfiguration_r18 != nil}); err != nil {
				return err
			}

			if ie.Scpac_ReferenceConfiguration_r18 != nil {
				choiceEnc := ex.NewChoiceEncoder(conditionalReconfigurationR16ExtScpacReferenceConfigurationR18Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.Scpac_ReferenceConfiguration_r18).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.Scpac_ReferenceConfiguration_r18).Choice {
				case ConditionalReconfiguration_r16_Ext_Scpac_ReferenceConfiguration_r18_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case ConditionalReconfiguration_r16_Ext_Scpac_ReferenceConfiguration_r18_Setup:
					if err := (*ie.Scpac_ReferenceConfiguration_r18).Setup.Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.ServingSecurityCellSetId_r18 != nil {
				if err := ie.ServingSecurityCellSetId_r18.Encode(ex); err != nil {
					return err
				}
			}

			if ie.Sk_CounterConfiguration_r18 != nil {
				if err := ie.Sk_CounterConfiguration_r18.Encode(ex); err != nil {
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

func (ie *ConditionalReconfiguration_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(conditionalReconfigurationR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. attemptCondReconfig-r16: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(conditionalReconfigurationR16AttemptCondReconfigR16Constraints)
			if err != nil {
				return err
			}
			ie.AttemptCondReconfig_r16 = &idx
		}
	}

	// 4. condReconfigToRemoveList-r16: ref
	{
		if seq.IsComponentPresent(1) {
			ie.CondReconfigToRemoveList_r16 = new(CondReconfigToRemoveList_r16)
			if err := ie.CondReconfigToRemoveList_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 5. condReconfigToAddModList-r16: ref
	{
		if seq.IsComponentPresent(2) {
			ie.CondReconfigToAddModList_r16 = new(CondReconfigToAddModList_r16)
			if err := ie.CondReconfigToAddModList_r16.Decode(d); err != nil {
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
				{Name: "scpac-ReferenceConfiguration-r18", Optional: true},
				{Name: "servingSecurityCellSetId-r18", Optional: true},
				{Name: "sk-CounterConfiguration-r18", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.Scpac_ReferenceConfiguration_r18 = &struct {
				Choice  int
				Release *struct{}
				Setup   *ReferenceConfiguration_r18
			}{}
			choiceDec := dx.NewChoiceDecoder(conditionalReconfigurationR16ExtScpacReferenceConfigurationR18Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Scpac_ReferenceConfiguration_r18).Choice = int(index)
			switch index {
			case ConditionalReconfiguration_r16_Ext_Scpac_ReferenceConfiguration_r18_Release:
				(*ie.Scpac_ReferenceConfiguration_r18).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case ConditionalReconfiguration_r16_Ext_Scpac_ReferenceConfiguration_r18_Setup:
				(*ie.Scpac_ReferenceConfiguration_r18).Setup = new(ReferenceConfiguration_r18)
				if err := (*ie.Scpac_ReferenceConfiguration_r18).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(1) {
			ie.ServingSecurityCellSetId_r18 = new(SecurityCellSetId_r18)
			if err := ie.ServingSecurityCellSetId_r18.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(2) {
			ie.Sk_CounterConfiguration_r18 = new(SK_CounterConfiguration_r18)
			if err := ie.Sk_CounterConfiguration_r18.Decode(dx); err != nil {
				return err
			}
		}
	}

	return nil
}
