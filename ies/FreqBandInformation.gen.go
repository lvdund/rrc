// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: FreqBandInformation (line 20732).

var freqBandInformationConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "bandInformationEUTRA"},
		{Name: "bandInformationNR"},
	},
}

const (
	FreqBandInformation_BandInformationEUTRA = 0
	FreqBandInformation_BandInformationNR    = 1
)

type FreqBandInformation struct {
	Choice               int
	BandInformationEUTRA *FreqBandInformationEUTRA
	BandInformationNR    *FreqBandInformationNR
}

func (ie *FreqBandInformation) Encode(e *per.Encoder) error {
	choiceEnc := e.NewChoiceEncoder(freqBandInformationConstraints)
	isExt := false
	if err := choiceEnc.EncodeChoice(int64(ie.Choice), isExt, nil); err != nil {
		return err
	}
	switch ie.Choice {
	case FreqBandInformation_BandInformationEUTRA:
		if err := ie.BandInformationEUTRA.Encode(e); err != nil {
			return err
		}
	case FreqBandInformation_BandInformationNR:
		if err := ie.BandInformationNR.Encode(e); err != nil {
			return err
		}
	default:
		// Choice index out of root range.
		return &per.ConstraintViolationError{
			TypeName:   "FreqBandInformation",
			Value:      int64(ie.Choice),
			Constraint: "choice index out of root range",
		}
	}
	return nil
}

func (ie *FreqBandInformation) Decode(d *per.Decoder) error {
	choiceDec := d.NewChoiceDecoder(freqBandInformationConstraints)
	index, isExt, _, err := choiceDec.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(index)
	if isExt {
		return &per.DecodeError{TypeName: "FreqBandInformation", Reason: "extension alternative not supported", Position: 0}
	}
	switch index {
	case FreqBandInformation_BandInformationEUTRA:
		ie.BandInformationEUTRA = new(FreqBandInformationEUTRA)
		if err := ie.BandInformationEUTRA.Decode(d); err != nil {
			return err
		}
	case FreqBandInformation_BandInformationNR:
		ie.BandInformationNR = new(FreqBandInformationNR)
		if err := ie.BandInformationNR.Decode(d); err != nil {
			return err
		}
	default:
		return &per.DecodeError{TypeName: "FreqBandInformation", Reason: "choice index out of root range", Position: 0}
	}
	return nil
}
