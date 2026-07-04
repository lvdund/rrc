// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: BWP-DownlinkCommon (line 5249).

var bWPDownlinkCommonConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "genericParameters"},
		{Name: "pdcch-ConfigCommon", Optional: true},
		{Name: "pdsch-ConfigCommon", Optional: true},
	},
}

var bWP_DownlinkCommonPdcchConfigCommonConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	BWP_DownlinkCommon_Pdcch_ConfigCommon_Release = 0
	BWP_DownlinkCommon_Pdcch_ConfigCommon_Setup   = 1
)

var bWP_DownlinkCommonPdschConfigCommonConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	BWP_DownlinkCommon_Pdsch_ConfigCommon_Release = 0
	BWP_DownlinkCommon_Pdsch_ConfigCommon_Setup   = 1
)

type BWP_DownlinkCommon struct {
	GenericParameters  BWP
	Pdcch_ConfigCommon *struct {
		Choice  int
		Release *struct{}
		Setup   *PDCCH_ConfigCommon
	}
	Pdsch_ConfigCommon *struct {
		Choice  int
		Release *struct{}
		Setup   *PDSCH_ConfigCommon
	}
}

func (ie *BWP_DownlinkCommon) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(bWPDownlinkCommonConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Pdcch_ConfigCommon != nil, ie.Pdsch_ConfigCommon != nil}); err != nil {
		return err
	}

	// 3. genericParameters: ref
	{
		if err := ie.GenericParameters.Encode(e); err != nil {
			return err
		}
	}

	// 4. pdcch-ConfigCommon: choice
	{
		if ie.Pdcch_ConfigCommon != nil {
			choiceEnc := e.NewChoiceEncoder(bWP_DownlinkCommonPdcchConfigCommonConstraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.Pdcch_ConfigCommon).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.Pdcch_ConfigCommon).Choice {
			case BWP_DownlinkCommon_Pdcch_ConfigCommon_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case BWP_DownlinkCommon_Pdcch_ConfigCommon_Setup:
				if err := (*ie.Pdcch_ConfigCommon).Setup.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.Pdcch_ConfigCommon).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 5. pdsch-ConfigCommon: choice
	{
		if ie.Pdsch_ConfigCommon != nil {
			choiceEnc := e.NewChoiceEncoder(bWP_DownlinkCommonPdschConfigCommonConstraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.Pdsch_ConfigCommon).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.Pdsch_ConfigCommon).Choice {
			case BWP_DownlinkCommon_Pdsch_ConfigCommon_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case BWP_DownlinkCommon_Pdsch_ConfigCommon_Setup:
				if err := (*ie.Pdsch_ConfigCommon).Setup.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.Pdsch_ConfigCommon).Choice), Constraint: "index out of range"}
			}
		}
	}

	return nil
}

func (ie *BWP_DownlinkCommon) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(bWPDownlinkCommonConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. genericParameters: ref
	{
		if err := ie.GenericParameters.Decode(d); err != nil {
			return err
		}
	}

	// 4. pdcch-ConfigCommon: choice
	{
		if seq.IsComponentPresent(1) {
			ie.Pdcch_ConfigCommon = &struct {
				Choice  int
				Release *struct{}
				Setup   *PDCCH_ConfigCommon
			}{}
			choiceDec := d.NewChoiceDecoder(bWP_DownlinkCommonPdcchConfigCommonConstraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Pdcch_ConfigCommon).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case BWP_DownlinkCommon_Pdcch_ConfigCommon_Release:
				(*ie.Pdcch_ConfigCommon).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case BWP_DownlinkCommon_Pdcch_ConfigCommon_Setup:
				(*ie.Pdcch_ConfigCommon).Setup = new(PDCCH_ConfigCommon)
				if err := (*ie.Pdcch_ConfigCommon).Setup.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 5. pdsch-ConfigCommon: choice
	{
		if seq.IsComponentPresent(2) {
			ie.Pdsch_ConfigCommon = &struct {
				Choice  int
				Release *struct{}
				Setup   *PDSCH_ConfigCommon
			}{}
			choiceDec := d.NewChoiceDecoder(bWP_DownlinkCommonPdschConfigCommonConstraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Pdsch_ConfigCommon).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case BWP_DownlinkCommon_Pdsch_ConfigCommon_Release:
				(*ie.Pdsch_ConfigCommon).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case BWP_DownlinkCommon_Pdsch_ConfigCommon_Setup:
				(*ie.Pdsch_ConfigCommon).Setup = new(PDSCH_ConfigCommon)
				if err := (*ie.Pdsch_ConfigCommon).Setup.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
