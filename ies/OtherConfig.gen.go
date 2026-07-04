// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: OtherConfig (line 26300).

var otherConfigConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "delayBudgetReportingConfig", Optional: true},
	},
}

var otherConfigDelayBudgetReportingConfigConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	OtherConfig_DelayBudgetReportingConfig_Release = 0
	OtherConfig_DelayBudgetReportingConfig_Setup   = 1
)

const (
	OtherConfig_DelayBudgetReportingConfig_Setup_DelayBudgetReportingProhibitTimer_S0     = 0
	OtherConfig_DelayBudgetReportingConfig_Setup_DelayBudgetReportingProhibitTimer_S0dot4 = 1
	OtherConfig_DelayBudgetReportingConfig_Setup_DelayBudgetReportingProhibitTimer_S0dot8 = 2
	OtherConfig_DelayBudgetReportingConfig_Setup_DelayBudgetReportingProhibitTimer_S1dot6 = 3
	OtherConfig_DelayBudgetReportingConfig_Setup_DelayBudgetReportingProhibitTimer_S3     = 4
	OtherConfig_DelayBudgetReportingConfig_Setup_DelayBudgetReportingProhibitTimer_S6     = 5
	OtherConfig_DelayBudgetReportingConfig_Setup_DelayBudgetReportingProhibitTimer_S12    = 6
	OtherConfig_DelayBudgetReportingConfig_Setup_DelayBudgetReportingProhibitTimer_S30    = 7
)

var otherConfigDelayBudgetReportingConfigSetupDelayBudgetReportingProhibitTimerConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{OtherConfig_DelayBudgetReportingConfig_Setup_DelayBudgetReportingProhibitTimer_S0, OtherConfig_DelayBudgetReportingConfig_Setup_DelayBudgetReportingProhibitTimer_S0dot4, OtherConfig_DelayBudgetReportingConfig_Setup_DelayBudgetReportingProhibitTimer_S0dot8, OtherConfig_DelayBudgetReportingConfig_Setup_DelayBudgetReportingProhibitTimer_S1dot6, OtherConfig_DelayBudgetReportingConfig_Setup_DelayBudgetReportingProhibitTimer_S3, OtherConfig_DelayBudgetReportingConfig_Setup_DelayBudgetReportingProhibitTimer_S6, OtherConfig_DelayBudgetReportingConfig_Setup_DelayBudgetReportingProhibitTimer_S12, OtherConfig_DelayBudgetReportingConfig_Setup_DelayBudgetReportingProhibitTimer_S30},
}

type OtherConfig struct {
	DelayBudgetReportingConfig *struct {
		Choice  int
		Release *struct{}
		Setup   *struct{ DelayBudgetReportingProhibitTimer int64 }
	}
}

func (ie *OtherConfig) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(otherConfigConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.DelayBudgetReportingConfig != nil}); err != nil {
		return err
	}

	// 2. delayBudgetReportingConfig: choice
	{
		if ie.DelayBudgetReportingConfig != nil {
			choiceEnc := e.NewChoiceEncoder(otherConfigDelayBudgetReportingConfigConstraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.DelayBudgetReportingConfig).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.DelayBudgetReportingConfig).Choice {
			case OtherConfig_DelayBudgetReportingConfig_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case OtherConfig_DelayBudgetReportingConfig_Setup:
				c := (*(*ie.DelayBudgetReportingConfig).Setup)
				if err := e.EncodeEnumerated(c.DelayBudgetReportingProhibitTimer, otherConfigDelayBudgetReportingConfigSetupDelayBudgetReportingProhibitTimerConstraints); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.DelayBudgetReportingConfig).Choice), Constraint: "index out of range"}
			}
		}
	}

	return nil
}

func (ie *OtherConfig) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(otherConfigConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. delayBudgetReportingConfig: choice
	{
		if seq.IsComponentPresent(0) {
			ie.DelayBudgetReportingConfig = &struct {
				Choice  int
				Release *struct{}
				Setup   *struct{ DelayBudgetReportingProhibitTimer int64 }
			}{}
			choiceDec := d.NewChoiceDecoder(otherConfigDelayBudgetReportingConfigConstraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.DelayBudgetReportingConfig).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case OtherConfig_DelayBudgetReportingConfig_Release:
				(*ie.DelayBudgetReportingConfig).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case OtherConfig_DelayBudgetReportingConfig_Setup:
				(*ie.DelayBudgetReportingConfig).Setup = &struct{ DelayBudgetReportingProhibitTimer int64 }{}
				c := (*(*ie.DelayBudgetReportingConfig).Setup)
				{
					v, err := d.DecodeEnumerated(otherConfigDelayBudgetReportingConfigSetupDelayBudgetReportingProhibitTimerConstraints)
					if err != nil {
						return err
					}
					c.DelayBudgetReportingProhibitTimer = v
				}
			}
		}
	}

	return nil
}
