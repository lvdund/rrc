// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MeasTriggerQuantityCLI-r16 (line 13907).

var measTriggerQuantityCLIR16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "srs-RSRP-r16"},
		{Name: "cli-RSSI-r16"},
	},
}

const (
	MeasTriggerQuantityCLI_r16_Srs_RSRP_r16 = 0
	MeasTriggerQuantityCLI_r16_Cli_RSSI_r16 = 1
)

type MeasTriggerQuantityCLI_r16 struct {
	Choice       int
	Srs_RSRP_r16 *SRS_RSRP_Range_r16
	Cli_RSSI_r16 *CLI_RSSI_Range_r16
}

func (ie *MeasTriggerQuantityCLI_r16) Encode(e *per.Encoder) error {
	choiceEnc := e.NewChoiceEncoder(measTriggerQuantityCLIR16Constraints)
	isExt := false
	if err := choiceEnc.EncodeChoice(int64(ie.Choice), isExt, nil); err != nil {
		return err
	}
	switch ie.Choice {
	case MeasTriggerQuantityCLI_r16_Srs_RSRP_r16:
		if err := ie.Srs_RSRP_r16.Encode(e); err != nil {
			return err
		}
	case MeasTriggerQuantityCLI_r16_Cli_RSSI_r16:
		if err := ie.Cli_RSSI_r16.Encode(e); err != nil {
			return err
		}
	default:
		// Choice index out of root range.
		return &per.ConstraintViolationError{
			TypeName:   "MeasTriggerQuantityCLI-r16",
			Value:      int64(ie.Choice),
			Constraint: "choice index out of root range",
		}
	}
	return nil
}

func (ie *MeasTriggerQuantityCLI_r16) Decode(d *per.Decoder) error {
	choiceDec := d.NewChoiceDecoder(measTriggerQuantityCLIR16Constraints)
	index, isExt, _, err := choiceDec.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(index)
	if isExt {
		return &per.DecodeError{TypeName: "MeasTriggerQuantityCLI-r16", Reason: "extension alternative not supported", Position: 0}
	}
	switch index {
	case MeasTriggerQuantityCLI_r16_Srs_RSRP_r16:
		ie.Srs_RSRP_r16 = new(SRS_RSRP_Range_r16)
		if err := ie.Srs_RSRP_r16.Decode(d); err != nil {
			return err
		}
	case MeasTriggerQuantityCLI_r16_Cli_RSSI_r16:
		ie.Cli_RSSI_r16 = new(CLI_RSSI_Range_r16)
		if err := ie.Cli_RSSI_r16.Decode(d); err != nil {
			return err
		}
	default:
		return &per.DecodeError{TypeName: "MeasTriggerQuantityCLI-r16", Reason: "choice index out of root range", Position: 0}
	}
	return nil
}
