package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	RateMatchPattern_patternType_Choice_nothing uint64 = iota
	RateMatchPattern_patternType_Choice_Bitmaps
	RateMatchPattern_patternType_Choice_ControlResourceSet
)

type RateMatchPattern_patternType struct {
	Choice             uint64
	Bitmaps            *RateMatchPattern_patternType_bitmaps
	ControlResourceSet *ControlResourceSetId
}

func (ie *RateMatchPattern_patternType) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case RateMatchPattern_patternType_Choice_Bitmaps:
		if err = ie.Bitmaps.Encode(w); err != nil {
			err = utils.WrapError("Encode Bitmaps", err)
		}
	case RateMatchPattern_patternType_Choice_ControlResourceSet:
		if err = ie.ControlResourceSet.Encode(w); err != nil {
			err = utils.WrapError("Encode ControlResourceSet", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *RateMatchPattern_patternType) Decode(r *aper.AperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case RateMatchPattern_patternType_Choice_Bitmaps:
		ie.Bitmaps = new(RateMatchPattern_patternType_bitmaps)
		if err = ie.Bitmaps.Decode(r); err != nil {
			return utils.WrapError("Decode Bitmaps", err)
		}
	case RateMatchPattern_patternType_Choice_ControlResourceSet:
		ie.ControlResourceSet = new(ControlResourceSetId)
		if err = ie.ControlResourceSet.Decode(r); err != nil {
			return utils.WrapError("Decode ControlResourceSet", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
