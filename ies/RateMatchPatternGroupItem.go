package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	RateMatchPatternGroupItem_Choice_nothing uint64 = iota
	RateMatchPatternGroupItem_Choice_CellLevel
	RateMatchPatternGroupItem_Choice_BwpLevel
)

type RateMatchPatternGroupItem struct {
	Choice    uint64
	CellLevel *RateMatchPatternId
	BwpLevel  *RateMatchPatternId
}

func (ie *RateMatchPatternGroupItem) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case RateMatchPatternGroupItem_Choice_CellLevel:
		if err = ie.CellLevel.Encode(w); err != nil {
			err = utils.WrapError("Encode CellLevel", err)
		}
	case RateMatchPatternGroupItem_Choice_BwpLevel:
		if err = ie.BwpLevel.Encode(w); err != nil {
			err = utils.WrapError("Encode BwpLevel", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *RateMatchPatternGroupItem) Decode(r *aper.AperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case RateMatchPatternGroupItem_Choice_CellLevel:
		ie.CellLevel = new(RateMatchPatternId)
		if err = ie.CellLevel.Decode(r); err != nil {
			return utils.WrapError("Decode CellLevel", err)
		}
	case RateMatchPatternGroupItem_Choice_BwpLevel:
		ie.BwpLevel = new(RateMatchPatternId)
		if err = ie.BwpLevel.Decode(r); err != nil {
			return utils.WrapError("Decode BwpLevel", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
