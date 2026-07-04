// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: BWP-UplinkDedicatedSDT-r17 (line 1411).

var bWPUplinkDedicatedSDTR17Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "pusch-Config-r17", Optional: true},
		{Name: "configuredGrantConfigToAddModList-r17", Optional: true},
		{Name: "configuredGrantConfigToReleaseList-r17", Optional: true},
	},
}

var bWP_UplinkDedicatedSDT_r17PuschConfigR17Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	BWP_UplinkDedicatedSDT_r17_Pusch_Config_r17_Release = 0
	BWP_UplinkDedicatedSDT_r17_Pusch_Config_r17_Setup   = 1
)

type BWP_UplinkDedicatedSDT_r17 struct {
	Pusch_Config_r17 *struct {
		Choice  int
		Release *struct{}
		Setup   *PUSCH_Config
	}
	ConfiguredGrantConfigToAddModList_r17  *ConfiguredGrantConfigToAddModList_r16
	ConfiguredGrantConfigToReleaseList_r17 *ConfiguredGrantConfigToReleaseList_r16
}

func (ie *BWP_UplinkDedicatedSDT_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(bWPUplinkDedicatedSDTR17Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Pusch_Config_r17 != nil, ie.ConfiguredGrantConfigToAddModList_r17 != nil, ie.ConfiguredGrantConfigToReleaseList_r17 != nil}); err != nil {
		return err
	}

	// 3. pusch-Config-r17: choice
	{
		if ie.Pusch_Config_r17 != nil {
			choiceEnc := e.NewChoiceEncoder(bWP_UplinkDedicatedSDT_r17PuschConfigR17Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.Pusch_Config_r17).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.Pusch_Config_r17).Choice {
			case BWP_UplinkDedicatedSDT_r17_Pusch_Config_r17_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case BWP_UplinkDedicatedSDT_r17_Pusch_Config_r17_Setup:
				if err := (*ie.Pusch_Config_r17).Setup.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.Pusch_Config_r17).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 4. configuredGrantConfigToAddModList-r17: ref
	{
		if ie.ConfiguredGrantConfigToAddModList_r17 != nil {
			if err := ie.ConfiguredGrantConfigToAddModList_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	// 5. configuredGrantConfigToReleaseList-r17: ref
	{
		if ie.ConfiguredGrantConfigToReleaseList_r17 != nil {
			if err := ie.ConfiguredGrantConfigToReleaseList_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *BWP_UplinkDedicatedSDT_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(bWPUplinkDedicatedSDTR17Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. pusch-Config-r17: choice
	{
		if seq.IsComponentPresent(0) {
			ie.Pusch_Config_r17 = &struct {
				Choice  int
				Release *struct{}
				Setup   *PUSCH_Config
			}{}
			choiceDec := d.NewChoiceDecoder(bWP_UplinkDedicatedSDT_r17PuschConfigR17Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Pusch_Config_r17).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case BWP_UplinkDedicatedSDT_r17_Pusch_Config_r17_Release:
				(*ie.Pusch_Config_r17).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case BWP_UplinkDedicatedSDT_r17_Pusch_Config_r17_Setup:
				(*ie.Pusch_Config_r17).Setup = new(PUSCH_Config)
				if err := (*ie.Pusch_Config_r17).Setup.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 4. configuredGrantConfigToAddModList-r17: ref
	{
		if seq.IsComponentPresent(1) {
			ie.ConfiguredGrantConfigToAddModList_r17 = new(ConfiguredGrantConfigToAddModList_r16)
			if err := ie.ConfiguredGrantConfigToAddModList_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	// 5. configuredGrantConfigToReleaseList-r17: ref
	{
		if seq.IsComponentPresent(2) {
			ie.ConfiguredGrantConfigToReleaseList_r17 = new(ConfiguredGrantConfigToReleaseList_r16)
			if err := ie.ConfiguredGrantConfigToReleaseList_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
