// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SupportedBandwidth (line 25413).

var supportedBandwidthConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "fr1"},
		{Name: "fr2"},
	},
}

const (
	SupportedBandwidth_Fr1 = 0
	SupportedBandwidth_Fr2 = 1
)

const (
	SupportedBandwidth_Fr1_Mhz5   = 0
	SupportedBandwidth_Fr1_Mhz10  = 1
	SupportedBandwidth_Fr1_Mhz15  = 2
	SupportedBandwidth_Fr1_Mhz20  = 3
	SupportedBandwidth_Fr1_Mhz25  = 4
	SupportedBandwidth_Fr1_Mhz30  = 5
	SupportedBandwidth_Fr1_Mhz40  = 6
	SupportedBandwidth_Fr1_Mhz50  = 7
	SupportedBandwidth_Fr1_Mhz60  = 8
	SupportedBandwidth_Fr1_Mhz80  = 9
	SupportedBandwidth_Fr1_Mhz100 = 10
)

var supportedBandwidthFr1Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SupportedBandwidth_Fr1_Mhz5, SupportedBandwidth_Fr1_Mhz10, SupportedBandwidth_Fr1_Mhz15, SupportedBandwidth_Fr1_Mhz20, SupportedBandwidth_Fr1_Mhz25, SupportedBandwidth_Fr1_Mhz30, SupportedBandwidth_Fr1_Mhz40, SupportedBandwidth_Fr1_Mhz50, SupportedBandwidth_Fr1_Mhz60, SupportedBandwidth_Fr1_Mhz80, SupportedBandwidth_Fr1_Mhz100},
}

const (
	SupportedBandwidth_Fr2_Mhz50  = 0
	SupportedBandwidth_Fr2_Mhz100 = 1
	SupportedBandwidth_Fr2_Mhz200 = 2
	SupportedBandwidth_Fr2_Mhz400 = 3
)

var supportedBandwidthFr2Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SupportedBandwidth_Fr2_Mhz50, SupportedBandwidth_Fr2_Mhz100, SupportedBandwidth_Fr2_Mhz200, SupportedBandwidth_Fr2_Mhz400},
}

type SupportedBandwidth struct {
	Choice int
	Fr1    *int64
	Fr2    *int64
}

func (ie *SupportedBandwidth) Encode(e *per.Encoder) error {
	choiceEnc := e.NewChoiceEncoder(supportedBandwidthConstraints)
	isExt := false
	if err := choiceEnc.EncodeChoice(int64(ie.Choice), isExt, nil); err != nil {
		return err
	}
	switch ie.Choice {
	case SupportedBandwidth_Fr1:
		if err := e.EncodeEnumerated((*ie.Fr1), supportedBandwidthFr1Constraints); err != nil {
			return err
		}
	case SupportedBandwidth_Fr2:
		if err := e.EncodeEnumerated((*ie.Fr2), supportedBandwidthFr2Constraints); err != nil {
			return err
		}
	default:
		// Choice index out of root range.
		return &per.ConstraintViolationError{
			TypeName:   "SupportedBandwidth",
			Value:      int64(ie.Choice),
			Constraint: "choice index out of root range",
		}
	}
	return nil
}

func (ie *SupportedBandwidth) Decode(d *per.Decoder) error {
	choiceDec := d.NewChoiceDecoder(supportedBandwidthConstraints)
	index, isExt, _, err := choiceDec.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(index)
	if isExt {
		return &per.DecodeError{TypeName: "SupportedBandwidth", Reason: "extension alternative not supported", Position: 0}
	}
	switch index {
	case SupportedBandwidth_Fr1:
		ie.Fr1 = new(int64)
		v, err := d.DecodeEnumerated(supportedBandwidthFr1Constraints)
		if err != nil {
			return err
		}
		(*ie.Fr1) = v
	case SupportedBandwidth_Fr2:
		ie.Fr2 = new(int64)
		v, err := d.DecodeEnumerated(supportedBandwidthFr2Constraints)
		if err != nil {
			return err
		}
		(*ie.Fr2) = v
	default:
		return &per.DecodeError{TypeName: "SupportedBandwidth", Reason: "choice index out of root range", Position: 0}
	}
	return nil
}
