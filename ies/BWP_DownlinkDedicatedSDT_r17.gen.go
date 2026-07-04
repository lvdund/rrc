// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: BWP-DownlinkDedicatedSDT-r17 (line 1405).

var bWPDownlinkDedicatedSDTR17Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "pdcch-Config-r17", Optional: true},
		{Name: "pdsch-Config-r17", Optional: true},
	},
}

var bWP_DownlinkDedicatedSDT_r17PdcchConfigR17Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	BWP_DownlinkDedicatedSDT_r17_Pdcch_Config_r17_Release = 0
	BWP_DownlinkDedicatedSDT_r17_Pdcch_Config_r17_Setup   = 1
)

var bWP_DownlinkDedicatedSDT_r17PdschConfigR17Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	BWP_DownlinkDedicatedSDT_r17_Pdsch_Config_r17_Release = 0
	BWP_DownlinkDedicatedSDT_r17_Pdsch_Config_r17_Setup   = 1
)

type BWP_DownlinkDedicatedSDT_r17 struct {
	Pdcch_Config_r17 *struct {
		Choice  int
		Release *struct{}
		Setup   *PDCCH_Config
	}
	Pdsch_Config_r17 *struct {
		Choice  int
		Release *struct{}
		Setup   *PDSCH_Config
	}
}

func (ie *BWP_DownlinkDedicatedSDT_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(bWPDownlinkDedicatedSDTR17Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Pdcch_Config_r17 != nil, ie.Pdsch_Config_r17 != nil}); err != nil {
		return err
	}

	// 3. pdcch-Config-r17: choice
	{
		if ie.Pdcch_Config_r17 != nil {
			choiceEnc := e.NewChoiceEncoder(bWP_DownlinkDedicatedSDT_r17PdcchConfigR17Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.Pdcch_Config_r17).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.Pdcch_Config_r17).Choice {
			case BWP_DownlinkDedicatedSDT_r17_Pdcch_Config_r17_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case BWP_DownlinkDedicatedSDT_r17_Pdcch_Config_r17_Setup:
				if err := (*ie.Pdcch_Config_r17).Setup.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.Pdcch_Config_r17).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 4. pdsch-Config-r17: choice
	{
		if ie.Pdsch_Config_r17 != nil {
			choiceEnc := e.NewChoiceEncoder(bWP_DownlinkDedicatedSDT_r17PdschConfigR17Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.Pdsch_Config_r17).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.Pdsch_Config_r17).Choice {
			case BWP_DownlinkDedicatedSDT_r17_Pdsch_Config_r17_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case BWP_DownlinkDedicatedSDT_r17_Pdsch_Config_r17_Setup:
				if err := (*ie.Pdsch_Config_r17).Setup.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.Pdsch_Config_r17).Choice), Constraint: "index out of range"}
			}
		}
	}

	return nil
}

func (ie *BWP_DownlinkDedicatedSDT_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(bWPDownlinkDedicatedSDTR17Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. pdcch-Config-r17: choice
	{
		if seq.IsComponentPresent(0) {
			ie.Pdcch_Config_r17 = &struct {
				Choice  int
				Release *struct{}
				Setup   *PDCCH_Config
			}{}
			choiceDec := d.NewChoiceDecoder(bWP_DownlinkDedicatedSDT_r17PdcchConfigR17Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Pdcch_Config_r17).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case BWP_DownlinkDedicatedSDT_r17_Pdcch_Config_r17_Release:
				(*ie.Pdcch_Config_r17).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case BWP_DownlinkDedicatedSDT_r17_Pdcch_Config_r17_Setup:
				(*ie.Pdcch_Config_r17).Setup = new(PDCCH_Config)
				if err := (*ie.Pdcch_Config_r17).Setup.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 4. pdsch-Config-r17: choice
	{
		if seq.IsComponentPresent(1) {
			ie.Pdsch_Config_r17 = &struct {
				Choice  int
				Release *struct{}
				Setup   *PDSCH_Config
			}{}
			choiceDec := d.NewChoiceDecoder(bWP_DownlinkDedicatedSDT_r17PdschConfigR17Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Pdsch_Config_r17).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case BWP_DownlinkDedicatedSDT_r17_Pdsch_Config_r17_Release:
				(*ie.Pdsch_Config_r17).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case BWP_DownlinkDedicatedSDT_r17_Pdsch_Config_r17_Setup:
				(*ie.Pdsch_Config_r17).Setup = new(PDSCH_Config)
				if err := (*ie.Pdsch_Config_r17).Setup.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
