// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: CLI-ResourceConfig-r16 (line 9342).

var cLIResourceConfigR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "srs-ResourceConfig-r16", Optional: true},
		{Name: "rssi-ResourceConfig-r16", Optional: true},
	},
}

var cLI_ResourceConfig_r16SrsResourceConfigR16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	CLI_ResourceConfig_r16_Srs_ResourceConfig_r16_Release = 0
	CLI_ResourceConfig_r16_Srs_ResourceConfig_r16_Setup   = 1
)

var cLI_ResourceConfig_r16RssiResourceConfigR16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	CLI_ResourceConfig_r16_Rssi_ResourceConfig_r16_Release = 0
	CLI_ResourceConfig_r16_Rssi_ResourceConfig_r16_Setup   = 1
)

type CLI_ResourceConfig_r16 struct {
	Srs_ResourceConfig_r16 *struct {
		Choice  int
		Release *struct{}
		Setup   *SRS_ResourceListConfigCLI_r16
	}
	Rssi_ResourceConfig_r16 *struct {
		Choice  int
		Release *struct{}
		Setup   *RSSI_ResourceListConfigCLI_r16
	}
}

func (ie *CLI_ResourceConfig_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(cLIResourceConfigR16Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Srs_ResourceConfig_r16 != nil, ie.Rssi_ResourceConfig_r16 != nil}); err != nil {
		return err
	}

	// 2. srs-ResourceConfig-r16: choice
	{
		if ie.Srs_ResourceConfig_r16 != nil {
			choiceEnc := e.NewChoiceEncoder(cLI_ResourceConfig_r16SrsResourceConfigR16Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.Srs_ResourceConfig_r16).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.Srs_ResourceConfig_r16).Choice {
			case CLI_ResourceConfig_r16_Srs_ResourceConfig_r16_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case CLI_ResourceConfig_r16_Srs_ResourceConfig_r16_Setup:
				if err := (*ie.Srs_ResourceConfig_r16).Setup.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.Srs_ResourceConfig_r16).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 3. rssi-ResourceConfig-r16: choice
	{
		if ie.Rssi_ResourceConfig_r16 != nil {
			choiceEnc := e.NewChoiceEncoder(cLI_ResourceConfig_r16RssiResourceConfigR16Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.Rssi_ResourceConfig_r16).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.Rssi_ResourceConfig_r16).Choice {
			case CLI_ResourceConfig_r16_Rssi_ResourceConfig_r16_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case CLI_ResourceConfig_r16_Rssi_ResourceConfig_r16_Setup:
				if err := (*ie.Rssi_ResourceConfig_r16).Setup.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.Rssi_ResourceConfig_r16).Choice), Constraint: "index out of range"}
			}
		}
	}

	return nil
}

func (ie *CLI_ResourceConfig_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(cLIResourceConfigR16Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. srs-ResourceConfig-r16: choice
	{
		if seq.IsComponentPresent(0) {
			ie.Srs_ResourceConfig_r16 = &struct {
				Choice  int
				Release *struct{}
				Setup   *SRS_ResourceListConfigCLI_r16
			}{}
			choiceDec := d.NewChoiceDecoder(cLI_ResourceConfig_r16SrsResourceConfigR16Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Srs_ResourceConfig_r16).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case CLI_ResourceConfig_r16_Srs_ResourceConfig_r16_Release:
				(*ie.Srs_ResourceConfig_r16).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case CLI_ResourceConfig_r16_Srs_ResourceConfig_r16_Setup:
				(*ie.Srs_ResourceConfig_r16).Setup = new(SRS_ResourceListConfigCLI_r16)
				if err := (*ie.Srs_ResourceConfig_r16).Setup.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 3. rssi-ResourceConfig-r16: choice
	{
		if seq.IsComponentPresent(1) {
			ie.Rssi_ResourceConfig_r16 = &struct {
				Choice  int
				Release *struct{}
				Setup   *RSSI_ResourceListConfigCLI_r16
			}{}
			choiceDec := d.NewChoiceDecoder(cLI_ResourceConfig_r16RssiResourceConfigR16Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Rssi_ResourceConfig_r16).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case CLI_ResourceConfig_r16_Rssi_ResourceConfig_r16_Release:
				(*ie.Rssi_ResourceConfig_r16).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case CLI_ResourceConfig_r16_Rssi_ResourceConfig_r16_Setup:
				(*ie.Rssi_ResourceConfig_r16).Setup = new(RSSI_ResourceListConfigCLI_r16)
				if err := (*ie.Rssi_ResourceConfig_r16).Setup.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
