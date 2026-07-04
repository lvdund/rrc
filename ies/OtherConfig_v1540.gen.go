// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: OtherConfig-v1540 (line 26309).

var otherConfigV1540Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "overheatingAssistanceConfig", Optional: true},
	},
}

var otherConfig_v1540OverheatingAssistanceConfigConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	OtherConfig_v1540_OverheatingAssistanceConfig_Release = 0
	OtherConfig_v1540_OverheatingAssistanceConfig_Setup   = 1
)

type OtherConfig_v1540 struct {
	OverheatingAssistanceConfig *struct {
		Choice  int
		Release *struct{}
		Setup   *OverheatingAssistanceConfig
	}
}

func (ie *OtherConfig_v1540) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(otherConfigV1540Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.OverheatingAssistanceConfig != nil}); err != nil {
		return err
	}

	// 3. overheatingAssistanceConfig: choice
	{
		if ie.OverheatingAssistanceConfig != nil {
			choiceEnc := e.NewChoiceEncoder(otherConfig_v1540OverheatingAssistanceConfigConstraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.OverheatingAssistanceConfig).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.OverheatingAssistanceConfig).Choice {
			case OtherConfig_v1540_OverheatingAssistanceConfig_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case OtherConfig_v1540_OverheatingAssistanceConfig_Setup:
				if err := (*ie.OverheatingAssistanceConfig).Setup.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.OverheatingAssistanceConfig).Choice), Constraint: "index out of range"}
			}
		}
	}

	return nil
}

func (ie *OtherConfig_v1540) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(otherConfigV1540Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. overheatingAssistanceConfig: choice
	{
		if seq.IsComponentPresent(0) {
			ie.OverheatingAssistanceConfig = &struct {
				Choice  int
				Release *struct{}
				Setup   *OverheatingAssistanceConfig
			}{}
			choiceDec := d.NewChoiceDecoder(otherConfig_v1540OverheatingAssistanceConfigConstraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.OverheatingAssistanceConfig).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case OtherConfig_v1540_OverheatingAssistanceConfig_Release:
				(*ie.OverheatingAssistanceConfig).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case OtherConfig_v1540_OverheatingAssistanceConfig_Setup:
				(*ie.OverheatingAssistanceConfig).Setup = new(OverheatingAssistanceConfig)
				if err := (*ie.OverheatingAssistanceConfig).Setup.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
