// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SupportedAggBandwidth-r17 (line 25401).

var supportedAggBandwidthR17Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "fr1-r17"},
		{Name: "fr2-r17"},
	},
}

const (
	SupportedAggBandwidth_r17_Fr1_r17 = 0
	SupportedAggBandwidth_r17_Fr2_r17 = 1
)

const (
	SupportedAggBandwidth_r17_Fr1_r17_Mhz20  = 0
	SupportedAggBandwidth_r17_Fr1_r17_Mhz30  = 1
	SupportedAggBandwidth_r17_Fr1_r17_Mhz35  = 2
	SupportedAggBandwidth_r17_Fr1_r17_Mhz40  = 3
	SupportedAggBandwidth_r17_Fr1_r17_Mhz50  = 4
	SupportedAggBandwidth_r17_Fr1_r17_Mhz60  = 5
	SupportedAggBandwidth_r17_Fr1_r17_Mhz70  = 6
	SupportedAggBandwidth_r17_Fr1_r17_Mhz80  = 7
	SupportedAggBandwidth_r17_Fr1_r17_Mhz90  = 8
	SupportedAggBandwidth_r17_Fr1_r17_Mhz100 = 9
	SupportedAggBandwidth_r17_Fr1_r17_Mhz110 = 10
	SupportedAggBandwidth_r17_Fr1_r17_Mhz120 = 11
	SupportedAggBandwidth_r17_Fr1_r17_Mhz130 = 12
	SupportedAggBandwidth_r17_Fr1_r17_Mhz140 = 13
	SupportedAggBandwidth_r17_Fr1_r17_Mhz150 = 14
	SupportedAggBandwidth_r17_Fr1_r17_Mhz160 = 15
	SupportedAggBandwidth_r17_Fr1_r17_Mhz180 = 16
	SupportedAggBandwidth_r17_Fr1_r17_Mhz200 = 17
	SupportedAggBandwidth_r17_Fr1_r17_Mhz220 = 18
	SupportedAggBandwidth_r17_Fr1_r17_Mhz230 = 19
	SupportedAggBandwidth_r17_Fr1_r17_Mhz250 = 20
	SupportedAggBandwidth_r17_Fr1_r17_Mhz280 = 21
	SupportedAggBandwidth_r17_Fr1_r17_Mhz290 = 22
	SupportedAggBandwidth_r17_Fr1_r17_Mhz300 = 23
	SupportedAggBandwidth_r17_Fr1_r17_Mhz350 = 24
	SupportedAggBandwidth_r17_Fr1_r17_Mhz400 = 25
	SupportedAggBandwidth_r17_Fr1_r17_Mhz450 = 26
	SupportedAggBandwidth_r17_Fr1_r17_Mhz500 = 27
	SupportedAggBandwidth_r17_Fr1_r17_Mhz600 = 28
	SupportedAggBandwidth_r17_Fr1_r17_Mhz700 = 29
	SupportedAggBandwidth_r17_Fr1_r17_Mhz800 = 30
	SupportedAggBandwidth_r17_Fr1_r17_Spare1 = 31
)

var supportedAggBandwidthR17Fr1R17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SupportedAggBandwidth_r17_Fr1_r17_Mhz20, SupportedAggBandwidth_r17_Fr1_r17_Mhz30, SupportedAggBandwidth_r17_Fr1_r17_Mhz35, SupportedAggBandwidth_r17_Fr1_r17_Mhz40, SupportedAggBandwidth_r17_Fr1_r17_Mhz50, SupportedAggBandwidth_r17_Fr1_r17_Mhz60, SupportedAggBandwidth_r17_Fr1_r17_Mhz70, SupportedAggBandwidth_r17_Fr1_r17_Mhz80, SupportedAggBandwidth_r17_Fr1_r17_Mhz90, SupportedAggBandwidth_r17_Fr1_r17_Mhz100, SupportedAggBandwidth_r17_Fr1_r17_Mhz110, SupportedAggBandwidth_r17_Fr1_r17_Mhz120, SupportedAggBandwidth_r17_Fr1_r17_Mhz130, SupportedAggBandwidth_r17_Fr1_r17_Mhz140, SupportedAggBandwidth_r17_Fr1_r17_Mhz150, SupportedAggBandwidth_r17_Fr1_r17_Mhz160, SupportedAggBandwidth_r17_Fr1_r17_Mhz180, SupportedAggBandwidth_r17_Fr1_r17_Mhz200, SupportedAggBandwidth_r17_Fr1_r17_Mhz220, SupportedAggBandwidth_r17_Fr1_r17_Mhz230, SupportedAggBandwidth_r17_Fr1_r17_Mhz250, SupportedAggBandwidth_r17_Fr1_r17_Mhz280, SupportedAggBandwidth_r17_Fr1_r17_Mhz290, SupportedAggBandwidth_r17_Fr1_r17_Mhz300, SupportedAggBandwidth_r17_Fr1_r17_Mhz350, SupportedAggBandwidth_r17_Fr1_r17_Mhz400, SupportedAggBandwidth_r17_Fr1_r17_Mhz450, SupportedAggBandwidth_r17_Fr1_r17_Mhz500, SupportedAggBandwidth_r17_Fr1_r17_Mhz600, SupportedAggBandwidth_r17_Fr1_r17_Mhz700, SupportedAggBandwidth_r17_Fr1_r17_Mhz800, SupportedAggBandwidth_r17_Fr1_r17_Spare1},
}

const (
	SupportedAggBandwidth_r17_Fr2_r17_Mhz200  = 0
	SupportedAggBandwidth_r17_Fr2_r17_Mhz300  = 1
	SupportedAggBandwidth_r17_Fr2_r17_Mhz400  = 2
	SupportedAggBandwidth_r17_Fr2_r17_Mhz500  = 3
	SupportedAggBandwidth_r17_Fr2_r17_Mhz600  = 4
	SupportedAggBandwidth_r17_Fr2_r17_Mhz700  = 5
	SupportedAggBandwidth_r17_Fr2_r17_Mhz800  = 6
	SupportedAggBandwidth_r17_Fr2_r17_Mhz900  = 7
	SupportedAggBandwidth_r17_Fr2_r17_Mhz1000 = 8
	SupportedAggBandwidth_r17_Fr2_r17_Mhz1100 = 9
	SupportedAggBandwidth_r17_Fr2_r17_Mhz1200 = 10
	SupportedAggBandwidth_r17_Fr2_r17_Mhz1300 = 11
	SupportedAggBandwidth_r17_Fr2_r17_Mhz1400 = 12
	SupportedAggBandwidth_r17_Fr2_r17_Mhz1500 = 13
	SupportedAggBandwidth_r17_Fr2_r17_Mhz1600 = 14
	SupportedAggBandwidth_r17_Fr2_r17_Mhz1700 = 15
	SupportedAggBandwidth_r17_Fr2_r17_Mhz1800 = 16
	SupportedAggBandwidth_r17_Fr2_r17_Mhz1900 = 17
	SupportedAggBandwidth_r17_Fr2_r17_Mhz2000 = 18
	SupportedAggBandwidth_r17_Fr2_r17_Mhz2100 = 19
	SupportedAggBandwidth_r17_Fr2_r17_Mhz2200 = 20
	SupportedAggBandwidth_r17_Fr2_r17_Mhz2300 = 21
	SupportedAggBandwidth_r17_Fr2_r17_Mhz2400 = 22
	SupportedAggBandwidth_r17_Fr2_r17_Spare9  = 23
	SupportedAggBandwidth_r17_Fr2_r17_Spare8  = 24
	SupportedAggBandwidth_r17_Fr2_r17_Spare7  = 25
	SupportedAggBandwidth_r17_Fr2_r17_Spare6  = 26
	SupportedAggBandwidth_r17_Fr2_r17_Spare5  = 27
	SupportedAggBandwidth_r17_Fr2_r17_Spare4  = 28
	SupportedAggBandwidth_r17_Fr2_r17_Spare3  = 29
	SupportedAggBandwidth_r17_Fr2_r17_Spare2  = 30
	SupportedAggBandwidth_r17_Fr2_r17_Spare1  = 31
)

var supportedAggBandwidthR17Fr2R17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SupportedAggBandwidth_r17_Fr2_r17_Mhz200, SupportedAggBandwidth_r17_Fr2_r17_Mhz300, SupportedAggBandwidth_r17_Fr2_r17_Mhz400, SupportedAggBandwidth_r17_Fr2_r17_Mhz500, SupportedAggBandwidth_r17_Fr2_r17_Mhz600, SupportedAggBandwidth_r17_Fr2_r17_Mhz700, SupportedAggBandwidth_r17_Fr2_r17_Mhz800, SupportedAggBandwidth_r17_Fr2_r17_Mhz900, SupportedAggBandwidth_r17_Fr2_r17_Mhz1000, SupportedAggBandwidth_r17_Fr2_r17_Mhz1100, SupportedAggBandwidth_r17_Fr2_r17_Mhz1200, SupportedAggBandwidth_r17_Fr2_r17_Mhz1300, SupportedAggBandwidth_r17_Fr2_r17_Mhz1400, SupportedAggBandwidth_r17_Fr2_r17_Mhz1500, SupportedAggBandwidth_r17_Fr2_r17_Mhz1600, SupportedAggBandwidth_r17_Fr2_r17_Mhz1700, SupportedAggBandwidth_r17_Fr2_r17_Mhz1800, SupportedAggBandwidth_r17_Fr2_r17_Mhz1900, SupportedAggBandwidth_r17_Fr2_r17_Mhz2000, SupportedAggBandwidth_r17_Fr2_r17_Mhz2100, SupportedAggBandwidth_r17_Fr2_r17_Mhz2200, SupportedAggBandwidth_r17_Fr2_r17_Mhz2300, SupportedAggBandwidth_r17_Fr2_r17_Mhz2400, SupportedAggBandwidth_r17_Fr2_r17_Spare9, SupportedAggBandwidth_r17_Fr2_r17_Spare8, SupportedAggBandwidth_r17_Fr2_r17_Spare7, SupportedAggBandwidth_r17_Fr2_r17_Spare6, SupportedAggBandwidth_r17_Fr2_r17_Spare5, SupportedAggBandwidth_r17_Fr2_r17_Spare4, SupportedAggBandwidth_r17_Fr2_r17_Spare3, SupportedAggBandwidth_r17_Fr2_r17_Spare2, SupportedAggBandwidth_r17_Fr2_r17_Spare1},
}

type SupportedAggBandwidth_r17 struct {
	Choice  int
	Fr1_r17 *int64
	Fr2_r17 *int64
}

func (ie *SupportedAggBandwidth_r17) Encode(e *per.Encoder) error {
	choiceEnc := e.NewChoiceEncoder(supportedAggBandwidthR17Constraints)
	isExt := false
	if err := choiceEnc.EncodeChoice(int64(ie.Choice), isExt, nil); err != nil {
		return err
	}
	switch ie.Choice {
	case SupportedAggBandwidth_r17_Fr1_r17:
		if err := e.EncodeEnumerated((*ie.Fr1_r17), supportedAggBandwidthR17Fr1R17Constraints); err != nil {
			return err
		}
	case SupportedAggBandwidth_r17_Fr2_r17:
		if err := e.EncodeEnumerated((*ie.Fr2_r17), supportedAggBandwidthR17Fr2R17Constraints); err != nil {
			return err
		}
	default:
		// Choice index out of root range.
		return &per.ConstraintViolationError{
			TypeName:   "SupportedAggBandwidth-r17",
			Value:      int64(ie.Choice),
			Constraint: "choice index out of root range",
		}
	}
	return nil
}

func (ie *SupportedAggBandwidth_r17) Decode(d *per.Decoder) error {
	choiceDec := d.NewChoiceDecoder(supportedAggBandwidthR17Constraints)
	index, isExt, _, err := choiceDec.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(index)
	if isExt {
		return &per.DecodeError{TypeName: "SupportedAggBandwidth-r17", Reason: "extension alternative not supported", Position: 0}
	}
	switch index {
	case SupportedAggBandwidth_r17_Fr1_r17:
		ie.Fr1_r17 = new(int64)
		v, err := d.DecodeEnumerated(supportedAggBandwidthR17Fr1R17Constraints)
		if err != nil {
			return err
		}
		(*ie.Fr1_r17) = v
	case SupportedAggBandwidth_r17_Fr2_r17:
		ie.Fr2_r17 = new(int64)
		v, err := d.DecodeEnumerated(supportedAggBandwidthR17Fr2R17Constraints)
		if err != nil {
			return err
		}
		(*ie.Fr2_r17) = v
	default:
		return &per.DecodeError{TypeName: "SupportedAggBandwidth-r17", Reason: "choice index out of root range", Position: 0}
	}
	return nil
}
