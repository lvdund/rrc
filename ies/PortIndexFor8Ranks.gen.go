// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: PortIndexFor8Ranks (line 7269).

var portIndexFor8RanksConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "portIndex8"},
		{Name: "portIndex4"},
		{Name: "portIndex2"},
		{Name: "portIndex1"},
	},
}

const (
	PortIndexFor8Ranks_PortIndex8 = 0
	PortIndexFor8Ranks_PortIndex4 = 1
	PortIndexFor8Ranks_PortIndex2 = 2
	PortIndexFor8Ranks_PortIndex1 = 3
)

var portIndexFor8RanksPortIndex8Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "rank1-8", Optional: true},
		{Name: "rank2-8", Optional: true},
		{Name: "rank3-8", Optional: true},
		{Name: "rank4-8", Optional: true},
		{Name: "rank5-8", Optional: true},
		{Name: "rank6-8", Optional: true},
		{Name: "rank7-8", Optional: true},
		{Name: "rank8-8", Optional: true},
	},
}

var portIndexFor8RanksPortIndex8Rank28Constraints = per.FixedSize(2)

var portIndexFor8RanksPortIndex8Rank38Constraints = per.FixedSize(3)

var portIndexFor8RanksPortIndex8Rank48Constraints = per.FixedSize(4)

var portIndexFor8RanksPortIndex8Rank58Constraints = per.FixedSize(5)

var portIndexFor8RanksPortIndex8Rank68Constraints = per.FixedSize(6)

var portIndexFor8RanksPortIndex8Rank78Constraints = per.FixedSize(7)

var portIndexFor8RanksPortIndex8Rank88Constraints = per.FixedSize(8)

var portIndexFor8RanksPortIndex4Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "rank1-4", Optional: true},
		{Name: "rank2-4", Optional: true},
		{Name: "rank3-4", Optional: true},
		{Name: "rank4-4", Optional: true},
	},
}

var portIndexFor8RanksPortIndex4Rank24Constraints = per.FixedSize(2)

var portIndexFor8RanksPortIndex4Rank34Constraints = per.FixedSize(3)

var portIndexFor8RanksPortIndex4Rank44Constraints = per.FixedSize(4)

var portIndexFor8RanksPortIndex2Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "rank1-2", Optional: true},
		{Name: "rank2-2", Optional: true},
	},
}

var portIndexFor8RanksPortIndex2Rank22Constraints = per.FixedSize(2)

type PortIndexFor8Ranks struct {
	Choice     int
	PortIndex8 *struct {
		Rank1_8 *PortIndex8
		Rank2_8 []PortIndex8
		Rank3_8 []PortIndex8
		Rank4_8 []PortIndex8
		Rank5_8 []PortIndex8
		Rank6_8 []PortIndex8
		Rank7_8 []PortIndex8
		Rank8_8 []PortIndex8
	}
	PortIndex4 *struct {
		Rank1_4 *PortIndex4
		Rank2_4 []PortIndex4
		Rank3_4 []PortIndex4
		Rank4_4 []PortIndex4
	}
	PortIndex2 *struct {
		Rank1_2 *PortIndex2
		Rank2_2 []PortIndex2
	}
	PortIndex1 *struct{}
}

func (ie *PortIndexFor8Ranks) Encode(e *per.Encoder) error {
	choiceEnc := e.NewChoiceEncoder(portIndexFor8RanksConstraints)
	isExt := false
	if err := choiceEnc.EncodeChoice(int64(ie.Choice), isExt, nil); err != nil {
		return err
	}
	switch ie.Choice {
	case PortIndexFor8Ranks_PortIndex8:
		c := (*ie.PortIndex8)
		portIndexFor8RanksPortIndex8Seq := e.NewSequenceEncoder(portIndexFor8RanksPortIndex8Constraints)
		if err := portIndexFor8RanksPortIndex8Seq.EncodePreamble([]bool{c.Rank1_8 != nil, c.Rank2_8 != nil, c.Rank3_8 != nil, c.Rank4_8 != nil, c.Rank5_8 != nil, c.Rank6_8 != nil, c.Rank7_8 != nil, c.Rank8_8 != nil}); err != nil {
			return err
		}
		if c.Rank1_8 != nil {
			if err := c.Rank1_8.Encode(e); err != nil {
				return err
			}
		}
		if c.Rank2_8 != nil {
			seqOf := e.NewSequenceOfEncoder(portIndexFor8RanksPortIndex8Rank28Constraints)
			if err := seqOf.EncodeLength(int64(len(c.Rank2_8))); err != nil {
				return err
			}
			for i := range c.Rank2_8 {
				if err := c.Rank2_8[i].Encode(e); err != nil {
					return err
				}
			}
		}
		if c.Rank3_8 != nil {
			seqOf := e.NewSequenceOfEncoder(portIndexFor8RanksPortIndex8Rank38Constraints)
			if err := seqOf.EncodeLength(int64(len(c.Rank3_8))); err != nil {
				return err
			}
			for i := range c.Rank3_8 {
				if err := c.Rank3_8[i].Encode(e); err != nil {
					return err
				}
			}
		}
		if c.Rank4_8 != nil {
			seqOf := e.NewSequenceOfEncoder(portIndexFor8RanksPortIndex8Rank48Constraints)
			if err := seqOf.EncodeLength(int64(len(c.Rank4_8))); err != nil {
				return err
			}
			for i := range c.Rank4_8 {
				if err := c.Rank4_8[i].Encode(e); err != nil {
					return err
				}
			}
		}
		if c.Rank5_8 != nil {
			seqOf := e.NewSequenceOfEncoder(portIndexFor8RanksPortIndex8Rank58Constraints)
			if err := seqOf.EncodeLength(int64(len(c.Rank5_8))); err != nil {
				return err
			}
			for i := range c.Rank5_8 {
				if err := c.Rank5_8[i].Encode(e); err != nil {
					return err
				}
			}
		}
		if c.Rank6_8 != nil {
			seqOf := e.NewSequenceOfEncoder(portIndexFor8RanksPortIndex8Rank68Constraints)
			if err := seqOf.EncodeLength(int64(len(c.Rank6_8))); err != nil {
				return err
			}
			for i := range c.Rank6_8 {
				if err := c.Rank6_8[i].Encode(e); err != nil {
					return err
				}
			}
		}
		if c.Rank7_8 != nil {
			seqOf := e.NewSequenceOfEncoder(portIndexFor8RanksPortIndex8Rank78Constraints)
			if err := seqOf.EncodeLength(int64(len(c.Rank7_8))); err != nil {
				return err
			}
			for i := range c.Rank7_8 {
				if err := c.Rank7_8[i].Encode(e); err != nil {
					return err
				}
			}
		}
		if c.Rank8_8 != nil {
			seqOf := e.NewSequenceOfEncoder(portIndexFor8RanksPortIndex8Rank88Constraints)
			if err := seqOf.EncodeLength(int64(len(c.Rank8_8))); err != nil {
				return err
			}
			for i := range c.Rank8_8 {
				if err := c.Rank8_8[i].Encode(e); err != nil {
					return err
				}
			}
		}
	case PortIndexFor8Ranks_PortIndex4:
		c := (*ie.PortIndex4)
		portIndexFor8RanksPortIndex4Seq := e.NewSequenceEncoder(portIndexFor8RanksPortIndex4Constraints)
		if err := portIndexFor8RanksPortIndex4Seq.EncodePreamble([]bool{c.Rank1_4 != nil, c.Rank2_4 != nil, c.Rank3_4 != nil, c.Rank4_4 != nil}); err != nil {
			return err
		}
		if c.Rank1_4 != nil {
			if err := c.Rank1_4.Encode(e); err != nil {
				return err
			}
		}
		if c.Rank2_4 != nil {
			seqOf := e.NewSequenceOfEncoder(portIndexFor8RanksPortIndex4Rank24Constraints)
			if err := seqOf.EncodeLength(int64(len(c.Rank2_4))); err != nil {
				return err
			}
			for i := range c.Rank2_4 {
				if err := c.Rank2_4[i].Encode(e); err != nil {
					return err
				}
			}
		}
		if c.Rank3_4 != nil {
			seqOf := e.NewSequenceOfEncoder(portIndexFor8RanksPortIndex4Rank34Constraints)
			if err := seqOf.EncodeLength(int64(len(c.Rank3_4))); err != nil {
				return err
			}
			for i := range c.Rank3_4 {
				if err := c.Rank3_4[i].Encode(e); err != nil {
					return err
				}
			}
		}
		if c.Rank4_4 != nil {
			seqOf := e.NewSequenceOfEncoder(portIndexFor8RanksPortIndex4Rank44Constraints)
			if err := seqOf.EncodeLength(int64(len(c.Rank4_4))); err != nil {
				return err
			}
			for i := range c.Rank4_4 {
				if err := c.Rank4_4[i].Encode(e); err != nil {
					return err
				}
			}
		}
	case PortIndexFor8Ranks_PortIndex2:
		c := (*ie.PortIndex2)
		portIndexFor8RanksPortIndex2Seq := e.NewSequenceEncoder(portIndexFor8RanksPortIndex2Constraints)
		if err := portIndexFor8RanksPortIndex2Seq.EncodePreamble([]bool{c.Rank1_2 != nil, c.Rank2_2 != nil}); err != nil {
			return err
		}
		if c.Rank1_2 != nil {
			if err := c.Rank1_2.Encode(e); err != nil {
				return err
			}
		}
		if c.Rank2_2 != nil {
			seqOf := e.NewSequenceOfEncoder(portIndexFor8RanksPortIndex2Rank22Constraints)
			if err := seqOf.EncodeLength(int64(len(c.Rank2_2))); err != nil {
				return err
			}
			for i := range c.Rank2_2 {
				if err := c.Rank2_2[i].Encode(e); err != nil {
					return err
				}
			}
		}
	case PortIndexFor8Ranks_PortIndex1:
		if err := e.EncodeNull(); err != nil {
			return err
		}
	default:
		// Choice index out of root range.
		return &per.ConstraintViolationError{
			TypeName:   "PortIndexFor8Ranks",
			Value:      int64(ie.Choice),
			Constraint: "choice index out of root range",
		}
	}
	return nil
}

func (ie *PortIndexFor8Ranks) Decode(d *per.Decoder) error {
	choiceDec := d.NewChoiceDecoder(portIndexFor8RanksConstraints)
	index, isExt, _, err := choiceDec.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(index)
	if isExt {
		return &per.DecodeError{TypeName: "PortIndexFor8Ranks", Reason: "extension alternative not supported", Position: 0}
	}
	switch index {
	case PortIndexFor8Ranks_PortIndex8:
		ie.PortIndex8 = &struct {
			Rank1_8 *PortIndex8
			Rank2_8 []PortIndex8
			Rank3_8 []PortIndex8
			Rank4_8 []PortIndex8
			Rank5_8 []PortIndex8
			Rank6_8 []PortIndex8
			Rank7_8 []PortIndex8
			Rank8_8 []PortIndex8
		}{}
		c := (*ie.PortIndex8)
		portIndexFor8RanksPortIndex8Seq := d.NewSequenceDecoder(portIndexFor8RanksPortIndex8Constraints)
		if err := portIndexFor8RanksPortIndex8Seq.DecodePreamble(); err != nil {
			return err
		}
		if portIndexFor8RanksPortIndex8Seq.IsComponentPresent(0) {
			c.Rank1_8 = new(PortIndex8)
			if err := (*c.Rank1_8).Decode(d); err != nil {
				return err
			}
		}
		if portIndexFor8RanksPortIndex8Seq.IsComponentPresent(1) {
			seqOf := d.NewSequenceOfDecoder(portIndexFor8RanksPortIndex8Rank28Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			c.Rank2_8 = make([]PortIndex8, n)
			for i := int64(0); i < n; i++ {
				if err := c.Rank2_8[i].Decode(d); err != nil {
					return err
				}
			}
		}
		if portIndexFor8RanksPortIndex8Seq.IsComponentPresent(2) {
			seqOf := d.NewSequenceOfDecoder(portIndexFor8RanksPortIndex8Rank38Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			c.Rank3_8 = make([]PortIndex8, n)
			for i := int64(0); i < n; i++ {
				if err := c.Rank3_8[i].Decode(d); err != nil {
					return err
				}
			}
		}
		if portIndexFor8RanksPortIndex8Seq.IsComponentPresent(3) {
			seqOf := d.NewSequenceOfDecoder(portIndexFor8RanksPortIndex8Rank48Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			c.Rank4_8 = make([]PortIndex8, n)
			for i := int64(0); i < n; i++ {
				if err := c.Rank4_8[i].Decode(d); err != nil {
					return err
				}
			}
		}
		if portIndexFor8RanksPortIndex8Seq.IsComponentPresent(4) {
			seqOf := d.NewSequenceOfDecoder(portIndexFor8RanksPortIndex8Rank58Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			c.Rank5_8 = make([]PortIndex8, n)
			for i := int64(0); i < n; i++ {
				if err := c.Rank5_8[i].Decode(d); err != nil {
					return err
				}
			}
		}
		if portIndexFor8RanksPortIndex8Seq.IsComponentPresent(5) {
			seqOf := d.NewSequenceOfDecoder(portIndexFor8RanksPortIndex8Rank68Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			c.Rank6_8 = make([]PortIndex8, n)
			for i := int64(0); i < n; i++ {
				if err := c.Rank6_8[i].Decode(d); err != nil {
					return err
				}
			}
		}
		if portIndexFor8RanksPortIndex8Seq.IsComponentPresent(6) {
			seqOf := d.NewSequenceOfDecoder(portIndexFor8RanksPortIndex8Rank78Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			c.Rank7_8 = make([]PortIndex8, n)
			for i := int64(0); i < n; i++ {
				if err := c.Rank7_8[i].Decode(d); err != nil {
					return err
				}
			}
		}
		if portIndexFor8RanksPortIndex8Seq.IsComponentPresent(7) {
			seqOf := d.NewSequenceOfDecoder(portIndexFor8RanksPortIndex8Rank88Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			c.Rank8_8 = make([]PortIndex8, n)
			for i := int64(0); i < n; i++ {
				if err := c.Rank8_8[i].Decode(d); err != nil {
					return err
				}
			}
		}
	case PortIndexFor8Ranks_PortIndex4:
		ie.PortIndex4 = &struct {
			Rank1_4 *PortIndex4
			Rank2_4 []PortIndex4
			Rank3_4 []PortIndex4
			Rank4_4 []PortIndex4
		}{}
		c := (*ie.PortIndex4)
		portIndexFor8RanksPortIndex4Seq := d.NewSequenceDecoder(portIndexFor8RanksPortIndex4Constraints)
		if err := portIndexFor8RanksPortIndex4Seq.DecodePreamble(); err != nil {
			return err
		}
		if portIndexFor8RanksPortIndex4Seq.IsComponentPresent(0) {
			c.Rank1_4 = new(PortIndex4)
			if err := (*c.Rank1_4).Decode(d); err != nil {
				return err
			}
		}
		if portIndexFor8RanksPortIndex4Seq.IsComponentPresent(1) {
			seqOf := d.NewSequenceOfDecoder(portIndexFor8RanksPortIndex4Rank24Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			c.Rank2_4 = make([]PortIndex4, n)
			for i := int64(0); i < n; i++ {
				if err := c.Rank2_4[i].Decode(d); err != nil {
					return err
				}
			}
		}
		if portIndexFor8RanksPortIndex4Seq.IsComponentPresent(2) {
			seqOf := d.NewSequenceOfDecoder(portIndexFor8RanksPortIndex4Rank34Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			c.Rank3_4 = make([]PortIndex4, n)
			for i := int64(0); i < n; i++ {
				if err := c.Rank3_4[i].Decode(d); err != nil {
					return err
				}
			}
		}
		if portIndexFor8RanksPortIndex4Seq.IsComponentPresent(3) {
			seqOf := d.NewSequenceOfDecoder(portIndexFor8RanksPortIndex4Rank44Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			c.Rank4_4 = make([]PortIndex4, n)
			for i := int64(0); i < n; i++ {
				if err := c.Rank4_4[i].Decode(d); err != nil {
					return err
				}
			}
		}
	case PortIndexFor8Ranks_PortIndex2:
		ie.PortIndex2 = &struct {
			Rank1_2 *PortIndex2
			Rank2_2 []PortIndex2
		}{}
		c := (*ie.PortIndex2)
		portIndexFor8RanksPortIndex2Seq := d.NewSequenceDecoder(portIndexFor8RanksPortIndex2Constraints)
		if err := portIndexFor8RanksPortIndex2Seq.DecodePreamble(); err != nil {
			return err
		}
		if portIndexFor8RanksPortIndex2Seq.IsComponentPresent(0) {
			c.Rank1_2 = new(PortIndex2)
			if err := (*c.Rank1_2).Decode(d); err != nil {
				return err
			}
		}
		if portIndexFor8RanksPortIndex2Seq.IsComponentPresent(1) {
			seqOf := d.NewSequenceOfDecoder(portIndexFor8RanksPortIndex2Rank22Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			c.Rank2_2 = make([]PortIndex2, n)
			for i := int64(0); i < n; i++ {
				if err := c.Rank2_2[i].Decode(d); err != nil {
					return err
				}
			}
		}
	case PortIndexFor8Ranks_PortIndex1:
		ie.PortIndex1 = &struct{}{}
		if err := d.DecodeNull(); err != nil {
			return err
		}
	default:
		return &per.DecodeError{TypeName: "PortIndexFor8Ranks", Reason: "choice index out of root range", Position: 0}
	}
	return nil
}
