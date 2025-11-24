package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	CG_UCI_OnPUSCH_Choice_nothing uint64 = iota
	CG_UCI_OnPUSCH_Choice_Dynamic
	CG_UCI_OnPUSCH_Choice_SemiStatic
)

type CG_UCI_OnPUSCH struct {
	Choice     uint64
	Dynamic    []BetaOffsets `lb:1,ub:4,madatory`
	SemiStatic *BetaOffsets
}

func (ie *CG_UCI_OnPUSCH) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case CG_UCI_OnPUSCH_Choice_Dynamic:
		tmp := utils.NewSequence[*BetaOffsets]([]*BetaOffsets{}, uper.Constraint{Lb: 1, Ub: 4}, false)
		for _, i := range ie.Dynamic {
			tmp.Value = append(tmp.Value, &i)
		}
		if err = tmp.Encode(w); err != nil {
			err = utils.WrapError("Encode Dynamic", err)
		}
	case CG_UCI_OnPUSCH_Choice_SemiStatic:
		if err = ie.SemiStatic.Encode(w); err != nil {
			err = utils.WrapError("Encode SemiStatic", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *CG_UCI_OnPUSCH) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case CG_UCI_OnPUSCH_Choice_Dynamic:
		tmp := utils.NewSequence[*BetaOffsets]([]*BetaOffsets{}, uper.Constraint{Lb: 1, Ub: 4}, false)
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
	case CG_UCI_OnPUSCH_Choice_SemiStatic:
		ie.SemiStatic = new(BetaOffsets)
		if err = ie.SemiStatic.Decode(r); err != nil {
			return utils.WrapError("Decode SemiStatic", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
