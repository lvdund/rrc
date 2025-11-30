package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	UCI_OnPUSCH_betaOffsets_Choice_nothing uint64 = iota
	UCI_OnPUSCH_betaOffsets_Choice_Dynamic
	UCI_OnPUSCH_betaOffsets_Choice_SemiStatic
)

type UCI_OnPUSCH_betaOffsets struct {
	Choice     uint64
	Dynamic    []BetaOffsets `lb:4,ub:4,madatory`
	SemiStatic *BetaOffsets
}

func (ie *UCI_OnPUSCH_betaOffsets) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case UCI_OnPUSCH_betaOffsets_Choice_Dynamic:
		tmp := utils.NewSequence[*BetaOffsets]([]*BetaOffsets{}, aper.Constraint{Lb: 4, Ub: 4}, false)
		for _, i := range ie.Dynamic {
			tmp.Value = append(tmp.Value, &i)
		}
		if err = tmp.Encode(w); err != nil {
			err = utils.WrapError("Encode Dynamic", err)
		}
	case UCI_OnPUSCH_betaOffsets_Choice_SemiStatic:
		if err = ie.SemiStatic.Encode(w); err != nil {
			err = utils.WrapError("Encode SemiStatic", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *UCI_OnPUSCH_betaOffsets) Decode(r *aper.AperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case UCI_OnPUSCH_betaOffsets_Choice_Dynamic:
		tmp := utils.NewSequence[*BetaOffsets]([]*BetaOffsets{}, aper.Constraint{Lb: 4, Ub: 4}, false)
		fn := func() *BetaOffsets {
			return new(BetaOffsets)
		}
		if err = tmp.Decode(r, fn); err != nil {
			return utils.WrapError("Decode Dynamic", err)
		}
		ie.Dynamic = []BetaOffsets{}
		for _, i := range tmp.Value {
			ie.Dynamic = append(ie.Dynamic, *i)
		}
	case UCI_OnPUSCH_betaOffsets_Choice_SemiStatic:
		ie.SemiStatic = new(BetaOffsets)
		if err = ie.SemiStatic.Decode(r); err != nil {
			return utils.WrapError("Decode SemiStatic", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
