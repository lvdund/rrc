// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SupportedBandwidth-v1700 (line 25418).

var supportedBandwidthV1700Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "fr1-r17"},
		{Name: "fr2-r17"},
	},
}

const (
	SupportedBandwidth_v1700_Fr1_r17 = 0
	SupportedBandwidth_v1700_Fr2_r17 = 1
)

const (
	SupportedBandwidth_v1700_Fr1_r17_Mhz5   = 0
	SupportedBandwidth_v1700_Fr1_r17_Mhz10  = 1
	SupportedBandwidth_v1700_Fr1_r17_Mhz15  = 2
	SupportedBandwidth_v1700_Fr1_r17_Mhz20  = 3
	SupportedBandwidth_v1700_Fr1_r17_Mhz25  = 4
	SupportedBandwidth_v1700_Fr1_r17_Mhz30  = 5
	SupportedBandwidth_v1700_Fr1_r17_Mhz35  = 6
	SupportedBandwidth_v1700_Fr1_r17_Mhz40  = 7
	SupportedBandwidth_v1700_Fr1_r17_Mhz45  = 8
	SupportedBandwidth_v1700_Fr1_r17_Mhz50  = 9
	SupportedBandwidth_v1700_Fr1_r17_Mhz60  = 10
	SupportedBandwidth_v1700_Fr1_r17_Mhz70  = 11
	SupportedBandwidth_v1700_Fr1_r17_Mhz80  = 12
	SupportedBandwidth_v1700_Fr1_r17_Mhz90  = 13
	SupportedBandwidth_v1700_Fr1_r17_Mhz100 = 14
)

var supportedBandwidthV1700Fr1R17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SupportedBandwidth_v1700_Fr1_r17_Mhz5, SupportedBandwidth_v1700_Fr1_r17_Mhz10, SupportedBandwidth_v1700_Fr1_r17_Mhz15, SupportedBandwidth_v1700_Fr1_r17_Mhz20, SupportedBandwidth_v1700_Fr1_r17_Mhz25, SupportedBandwidth_v1700_Fr1_r17_Mhz30, SupportedBandwidth_v1700_Fr1_r17_Mhz35, SupportedBandwidth_v1700_Fr1_r17_Mhz40, SupportedBandwidth_v1700_Fr1_r17_Mhz45, SupportedBandwidth_v1700_Fr1_r17_Mhz50, SupportedBandwidth_v1700_Fr1_r17_Mhz60, SupportedBandwidth_v1700_Fr1_r17_Mhz70, SupportedBandwidth_v1700_Fr1_r17_Mhz80, SupportedBandwidth_v1700_Fr1_r17_Mhz90, SupportedBandwidth_v1700_Fr1_r17_Mhz100},
}

const (
	SupportedBandwidth_v1700_Fr2_r17_Mhz50   = 0
	SupportedBandwidth_v1700_Fr2_r17_Mhz100  = 1
	SupportedBandwidth_v1700_Fr2_r17_Mhz200  = 2
	SupportedBandwidth_v1700_Fr2_r17_Mhz400  = 3
	SupportedBandwidth_v1700_Fr2_r17_Mhz800  = 4
	SupportedBandwidth_v1700_Fr2_r17_Mhz1600 = 5
	SupportedBandwidth_v1700_Fr2_r17_Mhz2000 = 6
)

var supportedBandwidthV1700Fr2R17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SupportedBandwidth_v1700_Fr2_r17_Mhz50, SupportedBandwidth_v1700_Fr2_r17_Mhz100, SupportedBandwidth_v1700_Fr2_r17_Mhz200, SupportedBandwidth_v1700_Fr2_r17_Mhz400, SupportedBandwidth_v1700_Fr2_r17_Mhz800, SupportedBandwidth_v1700_Fr2_r17_Mhz1600, SupportedBandwidth_v1700_Fr2_r17_Mhz2000},
}

type SupportedBandwidth_v1700 struct {
	Choice  int
	Fr1_r17 *int64
	Fr2_r17 *int64
}

func (ie *SupportedBandwidth_v1700) Encode(e *per.Encoder) error {
	choiceEnc := e.NewChoiceEncoder(supportedBandwidthV1700Constraints)
	isExt := false
	if err := choiceEnc.EncodeChoice(int64(ie.Choice), isExt, nil); err != nil {
		return err
	}
	switch ie.Choice {
	case SupportedBandwidth_v1700_Fr1_r17:
		if err := e.EncodeEnumerated((*ie.Fr1_r17), supportedBandwidthV1700Fr1R17Constraints); err != nil {
			return err
		}
	case SupportedBandwidth_v1700_Fr2_r17:
		if err := e.EncodeEnumerated((*ie.Fr2_r17), supportedBandwidthV1700Fr2R17Constraints); err != nil {
			return err
		}
	default:
		// Choice index out of root range.
		return &per.ConstraintViolationError{
			TypeName:   "SupportedBandwidth-v1700",
			Value:      int64(ie.Choice),
			Constraint: "choice index out of root range",
		}
	}
	return nil
}

func (ie *SupportedBandwidth_v1700) Decode(d *per.Decoder) error {
	choiceDec := d.NewChoiceDecoder(supportedBandwidthV1700Constraints)
	index, isExt, _, err := choiceDec.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(index)
	if isExt {
		return &per.DecodeError{TypeName: "SupportedBandwidth-v1700", Reason: "extension alternative not supported", Position: 0}
	}
	switch index {
	case SupportedBandwidth_v1700_Fr1_r17:
		ie.Fr1_r17 = new(int64)
		v, err := d.DecodeEnumerated(supportedBandwidthV1700Fr1R17Constraints)
		if err != nil {
			return err
		}
		(*ie.Fr1_r17) = v
	case SupportedBandwidth_v1700_Fr2_r17:
		ie.Fr2_r17 = new(int64)
		v, err := d.DecodeEnumerated(supportedBandwidthV1700Fr2R17Constraints)
		if err != nil {
			return err
		}
		(*ie.Fr2_r17) = v
	default:
		return &per.DecodeError{TypeName: "SupportedBandwidth-v1700", Reason: "choice index out of root range", Position: 0}
	}
	return nil
}
