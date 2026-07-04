// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: BandParameters (line 17028).

var bandParametersConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "eutra"},
		{Name: "nr"},
	},
}

const (
	BandParameters_Eutra = 0
	BandParameters_Nr    = 1
)

var bandParametersEutraConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "bandEUTRA"},
		{Name: "ca-BandwidthClassDL-EUTRA", Optional: true},
		{Name: "ca-BandwidthClassUL-EUTRA", Optional: true},
	},
}

var bandParametersNrConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "bandNR"},
		{Name: "ca-BandwidthClassDL-NR", Optional: true},
		{Name: "ca-BandwidthClassUL-NR", Optional: true},
	},
}

type BandParameters struct {
	Choice int
	Eutra  *struct {
		BandEUTRA                 FreqBandIndicatorEUTRA
		Ca_BandwidthClassDL_EUTRA *CA_BandwidthClassEUTRA
		Ca_BandwidthClassUL_EUTRA *CA_BandwidthClassEUTRA
	}
	Nr *struct {
		BandNR                 FreqBandIndicatorNR
		Ca_BandwidthClassDL_NR *CA_BandwidthClassNR
		Ca_BandwidthClassUL_NR *CA_BandwidthClassNR
	}
}

func (ie *BandParameters) Encode(e *per.Encoder) error {
	choiceEnc := e.NewChoiceEncoder(bandParametersConstraints)
	isExt := false
	if err := choiceEnc.EncodeChoice(int64(ie.Choice), isExt, nil); err != nil {
		return err
	}
	switch ie.Choice {
	case BandParameters_Eutra:
		c := (*ie.Eutra)
		bandParametersEutraSeq := e.NewSequenceEncoder(bandParametersEutraConstraints)
		if err := bandParametersEutraSeq.EncodePreamble([]bool{c.Ca_BandwidthClassDL_EUTRA != nil, c.Ca_BandwidthClassUL_EUTRA != nil}); err != nil {
			return err
		}
		if err := c.BandEUTRA.Encode(e); err != nil {
			return err
		}
		if c.Ca_BandwidthClassDL_EUTRA != nil {
			if err := c.Ca_BandwidthClassDL_EUTRA.Encode(e); err != nil {
				return err
			}
		}
		if c.Ca_BandwidthClassUL_EUTRA != nil {
			if err := c.Ca_BandwidthClassUL_EUTRA.Encode(e); err != nil {
				return err
			}
		}
	case BandParameters_Nr:
		c := (*ie.Nr)
		bandParametersNrSeq := e.NewSequenceEncoder(bandParametersNrConstraints)
		if err := bandParametersNrSeq.EncodePreamble([]bool{c.Ca_BandwidthClassDL_NR != nil, c.Ca_BandwidthClassUL_NR != nil}); err != nil {
			return err
		}
		if err := c.BandNR.Encode(e); err != nil {
			return err
		}
		if c.Ca_BandwidthClassDL_NR != nil {
			if err := c.Ca_BandwidthClassDL_NR.Encode(e); err != nil {
				return err
			}
		}
		if c.Ca_BandwidthClassUL_NR != nil {
			if err := c.Ca_BandwidthClassUL_NR.Encode(e); err != nil {
				return err
			}
		}
	default:
		// Choice index out of root range.
		return &per.ConstraintViolationError{
			TypeName:   "BandParameters",
			Value:      int64(ie.Choice),
			Constraint: "choice index out of root range",
		}
	}
	return nil
}

func (ie *BandParameters) Decode(d *per.Decoder) error {
	choiceDec := d.NewChoiceDecoder(bandParametersConstraints)
	index, isExt, _, err := choiceDec.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(index)
	if isExt {
		return &per.DecodeError{TypeName: "BandParameters", Reason: "extension alternative not supported", Position: 0}
	}
	switch index {
	case BandParameters_Eutra:
		ie.Eutra = &struct {
			BandEUTRA                 FreqBandIndicatorEUTRA
			Ca_BandwidthClassDL_EUTRA *CA_BandwidthClassEUTRA
			Ca_BandwidthClassUL_EUTRA *CA_BandwidthClassEUTRA
		}{}
		c := (*ie.Eutra)
		bandParametersEutraSeq := d.NewSequenceDecoder(bandParametersEutraConstraints)
		if err := bandParametersEutraSeq.DecodePreamble(); err != nil {
			return err
		}
		{
			if err := c.BandEUTRA.Decode(d); err != nil {
				return err
			}
		}
		if bandParametersEutraSeq.IsComponentPresent(1) {
			c.Ca_BandwidthClassDL_EUTRA = new(CA_BandwidthClassEUTRA)
			if err := (*c.Ca_BandwidthClassDL_EUTRA).Decode(d); err != nil {
				return err
			}
		}
		if bandParametersEutraSeq.IsComponentPresent(2) {
			c.Ca_BandwidthClassUL_EUTRA = new(CA_BandwidthClassEUTRA)
			if err := (*c.Ca_BandwidthClassUL_EUTRA).Decode(d); err != nil {
				return err
			}
		}
	case BandParameters_Nr:
		ie.Nr = &struct {
			BandNR                 FreqBandIndicatorNR
			Ca_BandwidthClassDL_NR *CA_BandwidthClassNR
			Ca_BandwidthClassUL_NR *CA_BandwidthClassNR
		}{}
		c := (*ie.Nr)
		bandParametersNrSeq := d.NewSequenceDecoder(bandParametersNrConstraints)
		if err := bandParametersNrSeq.DecodePreamble(); err != nil {
			return err
		}
		{
			if err := c.BandNR.Decode(d); err != nil {
				return err
			}
		}
		if bandParametersNrSeq.IsComponentPresent(1) {
			c.Ca_BandwidthClassDL_NR = new(CA_BandwidthClassNR)
			if err := (*c.Ca_BandwidthClassDL_NR).Decode(d); err != nil {
				return err
			}
		}
		if bandParametersNrSeq.IsComponentPresent(2) {
			c.Ca_BandwidthClassUL_NR = new(CA_BandwidthClassNR)
			if err := (*c.Ca_BandwidthClassUL_NR).Decode(d); err != nil {
				return err
			}
		}
	default:
		return &per.DecodeError{TypeName: "BandParameters", Reason: "choice index out of root range", Position: 0}
	}
	return nil
}
