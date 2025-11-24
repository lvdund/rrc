package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	BetaOffsetsCrossPriSelDCI_0_2_r17_Choice_nothing uint64 = iota
	BetaOffsetsCrossPriSelDCI_0_2_r17_Choice_DynamicDCI_0_2_r17
	BetaOffsetsCrossPriSelDCI_0_2_r17_Choice_SemiStaticDCI_0_2_r17
)

type BetaOffsetsCrossPriSelDCI_0_2_r17 struct {
	Choice                uint64
	DynamicDCI_0_2_r17    *BetaOffsetsCrossPriSelDCI_0_2_r17_dynamicDCI_0_2_r17
	SemiStaticDCI_0_2_r17 *BetaOffsetsCrossPri_r17
}

func (ie *BetaOffsetsCrossPriSelDCI_0_2_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case BetaOffsetsCrossPriSelDCI_0_2_r17_Choice_DynamicDCI_0_2_r17:
		if err = ie.DynamicDCI_0_2_r17.Encode(w); err != nil {
			err = utils.WrapError("Encode DynamicDCI_0_2_r17", err)
		}
	case BetaOffsetsCrossPriSelDCI_0_2_r17_Choice_SemiStaticDCI_0_2_r17:
		if err = ie.SemiStaticDCI_0_2_r17.Encode(w); err != nil {
			err = utils.WrapError("Encode SemiStaticDCI_0_2_r17", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *BetaOffsetsCrossPriSelDCI_0_2_r17) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case BetaOffsetsCrossPriSelDCI_0_2_r17_Choice_DynamicDCI_0_2_r17:
		ie.DynamicDCI_0_2_r17 = new(BetaOffsetsCrossPriSelDCI_0_2_r17_dynamicDCI_0_2_r17)
		if err = ie.DynamicDCI_0_2_r17.Decode(r); err != nil {
			return utils.WrapError("Decode DynamicDCI_0_2_r17", err)
		}
	case BetaOffsetsCrossPriSelDCI_0_2_r17_Choice_SemiStaticDCI_0_2_r17:
		ie.SemiStaticDCI_0_2_r17 = new(BetaOffsetsCrossPri_r17)
		if err = ie.SemiStaticDCI_0_2_r17.Decode(r); err != nil {
			return utils.WrapError("Decode SemiStaticDCI_0_2_r17", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
