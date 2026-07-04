// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: BandParametersSidelinkEUTRA-NR-r16 (line 17151).

var bandParametersSidelinkEUTRANRR16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "eutra"},
		{Name: "nr"},
	},
}

const (
	BandParametersSidelinkEUTRA_NR_r16_Eutra = 0
	BandParametersSidelinkEUTRA_NR_r16_Nr    = 1
)

var bandParametersSidelinkEUTRANRR16EutraConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "bandParametersSidelinkEUTRA1-r16", Optional: true},
		{Name: "bandParametersSidelinkEUTRA2-r16", Optional: true},
	},
}

type BandParametersSidelinkEUTRA_NR_r16 struct {
	Choice int
	Eutra  *struct {
		BandParametersSidelinkEUTRA1_r16 []byte
		BandParametersSidelinkEUTRA2_r16 []byte
	}
	Nr *struct{ BandParametersSidelinkNR_r16 BandParametersSidelink_r16 }
}

func (ie *BandParametersSidelinkEUTRA_NR_r16) Encode(e *per.Encoder) error {
	choiceEnc := e.NewChoiceEncoder(bandParametersSidelinkEUTRANRR16Constraints)
	isExt := false
	if err := choiceEnc.EncodeChoice(int64(ie.Choice), isExt, nil); err != nil {
		return err
	}
	switch ie.Choice {
	case BandParametersSidelinkEUTRA_NR_r16_Eutra:
		c := (*ie.Eutra)
		bandParametersSidelinkEUTRANRR16EutraSeq := e.NewSequenceEncoder(bandParametersSidelinkEUTRANRR16EutraConstraints)
		if err := bandParametersSidelinkEUTRANRR16EutraSeq.EncodePreamble([]bool{c.BandParametersSidelinkEUTRA1_r16 != nil, c.BandParametersSidelinkEUTRA2_r16 != nil}); err != nil {
			return err
		}
		if c.BandParametersSidelinkEUTRA1_r16 != nil {
			if err := e.EncodeOctetString(c.BandParametersSidelinkEUTRA1_r16, per.SizeConstraints{}); err != nil {
				return err
			}
		}
		if c.BandParametersSidelinkEUTRA2_r16 != nil {
			if err := e.EncodeOctetString(c.BandParametersSidelinkEUTRA2_r16, per.SizeConstraints{}); err != nil {
				return err
			}
		}
	case BandParametersSidelinkEUTRA_NR_r16_Nr:
		c := (*ie.Nr)
		if err := c.BandParametersSidelinkNR_r16.Encode(e); err != nil {
			return err
		}
	default:
		// Choice index out of root range.
		return &per.ConstraintViolationError{
			TypeName:   "BandParametersSidelinkEUTRA-NR-r16",
			Value:      int64(ie.Choice),
			Constraint: "choice index out of root range",
		}
	}
	return nil
}

func (ie *BandParametersSidelinkEUTRA_NR_r16) Decode(d *per.Decoder) error {
	choiceDec := d.NewChoiceDecoder(bandParametersSidelinkEUTRANRR16Constraints)
	index, isExt, _, err := choiceDec.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(index)
	if isExt {
		return &per.DecodeError{TypeName: "BandParametersSidelinkEUTRA-NR-r16", Reason: "extension alternative not supported", Position: 0}
	}
	switch index {
	case BandParametersSidelinkEUTRA_NR_r16_Eutra:
		ie.Eutra = &struct {
			BandParametersSidelinkEUTRA1_r16 []byte
			BandParametersSidelinkEUTRA2_r16 []byte
		}{}
		c := (*ie.Eutra)
		bandParametersSidelinkEUTRANRR16EutraSeq := d.NewSequenceDecoder(bandParametersSidelinkEUTRANRR16EutraConstraints)
		if err := bandParametersSidelinkEUTRANRR16EutraSeq.DecodePreamble(); err != nil {
			return err
		}
		if bandParametersSidelinkEUTRANRR16EutraSeq.IsComponentPresent(0) {
			v, err := d.DecodeOctetString(per.SizeConstraints{})
			if err != nil {
				return err
			}
			c.BandParametersSidelinkEUTRA1_r16 = v
		}
		if bandParametersSidelinkEUTRANRR16EutraSeq.IsComponentPresent(1) {
			v, err := d.DecodeOctetString(per.SizeConstraints{})
			if err != nil {
				return err
			}
			c.BandParametersSidelinkEUTRA2_r16 = v
		}
	case BandParametersSidelinkEUTRA_NR_r16_Nr:
		ie.Nr = &struct{ BandParametersSidelinkNR_r16 BandParametersSidelink_r16 }{}
		c := (*ie.Nr)
		{
			if err := c.BandParametersSidelinkNR_r16.Decode(d); err != nil {
				return err
			}
		}
	default:
		return &per.DecodeError{TypeName: "BandParametersSidelinkEUTRA-NR-r16", Reason: "choice index out of root range", Position: 0}
	}
	return nil
}
