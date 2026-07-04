// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: RedirectedCarrierInfo (line 1262).

var redirectedCarrierInfoConstraints = per.ChoiceConstraints{
	Extensible: true,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "nr"},
		{Name: "eutra"},
	},
	ExtAlternatives: []per.AlternativeInfo{
		{Name: "nr-v1900"},
	},
}

const (
	RedirectedCarrierInfo_Nr    = 0
	RedirectedCarrierInfo_Eutra = 1
)

type RedirectedCarrierInfo struct {
	Choice int
	Nr     *CarrierInfoNR
	Eutra  *RedirectedCarrierInfo_EUTRA
}

func (ie *RedirectedCarrierInfo) Encode(e *per.Encoder) error {
	choiceEnc := e.NewChoiceEncoder(redirectedCarrierInfoConstraints)
	isExt := false
	if err := choiceEnc.EncodeChoice(int64(ie.Choice), isExt, nil); err != nil {
		return err
	}
	switch ie.Choice {
	case RedirectedCarrierInfo_Nr:
		if err := ie.Nr.Encode(e); err != nil {
			return err
		}
	case RedirectedCarrierInfo_Eutra:
		if err := ie.Eutra.Encode(e); err != nil {
			return err
		}
	default:
		// Extension alternative: not modeled; bail out.
		return &per.ConstraintViolationError{
			TypeName:   "RedirectedCarrierInfo",
			Value:      int64(ie.Choice),
			Constraint: "choice index out of root range",
		}
	}
	return nil
}

func (ie *RedirectedCarrierInfo) Decode(d *per.Decoder) error {
	choiceDec := d.NewChoiceDecoder(redirectedCarrierInfoConstraints)
	index, isExt, _, err := choiceDec.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(index)
	if isExt {
		return &per.DecodeError{TypeName: "RedirectedCarrierInfo", Reason: "extension alternative not supported", Position: 0}
	}
	switch index {
	case RedirectedCarrierInfo_Nr:
		ie.Nr = new(CarrierInfoNR)
		if err := ie.Nr.Decode(d); err != nil {
			return err
		}
	case RedirectedCarrierInfo_Eutra:
		ie.Eutra = new(RedirectedCarrierInfo_EUTRA)
		if err := ie.Eutra.Decode(d); err != nil {
			return err
		}
	default:
		return &per.DecodeError{TypeName: "RedirectedCarrierInfo", Reason: "choice index out of root range", Position: 0}
	}
	return nil
}
