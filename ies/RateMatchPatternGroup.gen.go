// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: RateMatchPatternGroup (line 11438).

var rateMatchPatternGroupSizeConstraints = per.SizeRange(1, common.MaxNrofRateMatchPatternsPerGroup)

var rateMatchPatternGroupRateMatchPatternGroupElemConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "cellLevel"},
		{Name: "bwpLevel"},
	},
}

const (
	RateMatchPatternGroup_RateMatchPatternGroup_Elem_CellLevel = 0
	RateMatchPatternGroup_RateMatchPatternGroup_Elem_BwpLevel  = 1
)

type RateMatchPatternGroup []struct {
	Choice    int
	CellLevel *RateMatchPatternId
	BwpLevel  *RateMatchPatternId
}

func (ie *RateMatchPatternGroup) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(rateMatchPatternGroupSizeConstraints)
	if err := seqOf.EncodeLength(int64(len(*ie))); err != nil {
		return err
	}
	for i := range *ie {
		choiceEnc := e.NewChoiceEncoder(rateMatchPatternGroupRateMatchPatternGroupElemConstraints)
		if err := choiceEnc.EncodeChoice(int64((*ie)[i].Choice), false, nil); err != nil {
			return err
		}
		switch (*ie)[i].Choice {
		case RateMatchPatternGroup_RateMatchPatternGroup_Elem_CellLevel:
			if err := (*ie)[i].CellLevel.Encode(e); err != nil {
				return err
			}
		case RateMatchPatternGroup_RateMatchPatternGroup_Elem_BwpLevel:
			if err := (*ie)[i].BwpLevel.Encode(e); err != nil {
				return err
			}
		}
	}
	return nil
}

func (ie *RateMatchPatternGroup) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(rateMatchPatternGroupSizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(RateMatchPatternGroup, n)
	for i := int64(0); i < n; i++ {
		choiceDec := d.NewChoiceDecoder(rateMatchPatternGroupRateMatchPatternGroupElemConstraints)
		index, _, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		(*ie)[i].Choice = int(index)
		switch index {
		case RateMatchPatternGroup_RateMatchPatternGroup_Elem_CellLevel:
			(*ie)[i].CellLevel = new(RateMatchPatternId)
			if err := (*ie)[i].CellLevel.Decode(d); err != nil {
				return err
			}
		case RateMatchPatternGroup_RateMatchPatternGroup_Elem_BwpLevel:
			(*ie)[i].BwpLevel = new(RateMatchPatternId)
			if err := (*ie)[i].BwpLevel.Decode(d); err != nil {
				return err
			}
		}
	}
	return nil
}
