// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: IDC-AssistanceConfig-v1800 (line 26376).

var iDCAssistanceConfigV1800Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "idc-FDM-AssistanceConfig-r18", Optional: true},
		{Name: "idc-TDM-AssistanceConfig-r18", Optional: true},
	},
}

var iDC_AssistanceConfig_v1800IdcFDMAssistanceConfigR18Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	IDC_AssistanceConfig_v1800_Idc_FDM_AssistanceConfig_r18_Release = 0
	IDC_AssistanceConfig_v1800_Idc_FDM_AssistanceConfig_r18_Setup   = 1
)

const (
	IDC_AssistanceConfig_v1800_Idc_TDM_AssistanceConfig_r18_Setup = 0
)

var iDCAssistanceConfigV1800IdcTDMAssistanceConfigR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{IDC_AssistanceConfig_v1800_Idc_TDM_AssistanceConfig_r18_Setup},
}

type IDC_AssistanceConfig_v1800 struct {
	Idc_FDM_AssistanceConfig_r18 *struct {
		Choice  int
		Release *struct{}
		Setup   *IDC_FDM_AssistanceConfig_r18
	}
	Idc_TDM_AssistanceConfig_r18 *int64
}

func (ie *IDC_AssistanceConfig_v1800) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(iDCAssistanceConfigV1800Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Idc_FDM_AssistanceConfig_r18 != nil, ie.Idc_TDM_AssistanceConfig_r18 != nil}); err != nil {
		return err
	}

	// 2. idc-FDM-AssistanceConfig-r18: choice
	{
		if ie.Idc_FDM_AssistanceConfig_r18 != nil {
			choiceEnc := e.NewChoiceEncoder(iDC_AssistanceConfig_v1800IdcFDMAssistanceConfigR18Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.Idc_FDM_AssistanceConfig_r18).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.Idc_FDM_AssistanceConfig_r18).Choice {
			case IDC_AssistanceConfig_v1800_Idc_FDM_AssistanceConfig_r18_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case IDC_AssistanceConfig_v1800_Idc_FDM_AssistanceConfig_r18_Setup:
				if err := (*ie.Idc_FDM_AssistanceConfig_r18).Setup.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.Idc_FDM_AssistanceConfig_r18).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 3. idc-TDM-AssistanceConfig-r18: enumerated
	{
		if ie.Idc_TDM_AssistanceConfig_r18 != nil {
			if err := e.EncodeEnumerated(*ie.Idc_TDM_AssistanceConfig_r18, iDCAssistanceConfigV1800IdcTDMAssistanceConfigR18Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *IDC_AssistanceConfig_v1800) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(iDCAssistanceConfigV1800Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. idc-FDM-AssistanceConfig-r18: choice
	{
		if seq.IsComponentPresent(0) {
			ie.Idc_FDM_AssistanceConfig_r18 = &struct {
				Choice  int
				Release *struct{}
				Setup   *IDC_FDM_AssistanceConfig_r18
			}{}
			choiceDec := d.NewChoiceDecoder(iDC_AssistanceConfig_v1800IdcFDMAssistanceConfigR18Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Idc_FDM_AssistanceConfig_r18).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case IDC_AssistanceConfig_v1800_Idc_FDM_AssistanceConfig_r18_Release:
				(*ie.Idc_FDM_AssistanceConfig_r18).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case IDC_AssistanceConfig_v1800_Idc_FDM_AssistanceConfig_r18_Setup:
				(*ie.Idc_FDM_AssistanceConfig_r18).Setup = new(IDC_FDM_AssistanceConfig_r18)
				if err := (*ie.Idc_FDM_AssistanceConfig_r18).Setup.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 3. idc-TDM-AssistanceConfig-r18: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(iDCAssistanceConfigV1800IdcTDMAssistanceConfigR18Constraints)
			if err != nil {
				return err
			}
			ie.Idc_TDM_AssistanceConfig_r18 = &idx
		}
	}

	return nil
}
