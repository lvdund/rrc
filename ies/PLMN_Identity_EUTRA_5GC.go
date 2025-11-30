package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	PLMN_Identity_EUTRA_5GC_Choice_nothing uint64 = iota
	PLMN_Identity_EUTRA_5GC_Choice_Plmn_Identity_EUTRA_5GC
	PLMN_Identity_EUTRA_5GC_Choice_Plmn_index
)

type PLMN_Identity_EUTRA_5GC struct {
	Choice                  uint64
	Plmn_Identity_EUTRA_5GC *PLMN_Identity
	Plmn_index              int64 `lb:1,ub:maxPLMN,madatory`
}

func (ie *PLMN_Identity_EUTRA_5GC) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case PLMN_Identity_EUTRA_5GC_Choice_Plmn_Identity_EUTRA_5GC:
		if err = ie.Plmn_Identity_EUTRA_5GC.Encode(w); err != nil {
			err = utils.WrapError("Encode Plmn_Identity_EUTRA_5GC", err)
		}
	case PLMN_Identity_EUTRA_5GC_Choice_Plmn_index:
		if err = w.WriteInteger(int64(ie.Plmn_index), &aper.Constraint{Lb: 1, Ub: maxPLMN}, false); err != nil {
			err = utils.WrapError("Encode Plmn_index", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *PLMN_Identity_EUTRA_5GC) Decode(r *aper.AperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case PLMN_Identity_EUTRA_5GC_Choice_Plmn_Identity_EUTRA_5GC:
		ie.Plmn_Identity_EUTRA_5GC = new(PLMN_Identity)
		if err = ie.Plmn_Identity_EUTRA_5GC.Decode(r); err != nil {
			return utils.WrapError("Decode Plmn_Identity_EUTRA_5GC", err)
		}
	case PLMN_Identity_EUTRA_5GC_Choice_Plmn_index:
		var tmp_int_Plmn_index int64
		if tmp_int_Plmn_index, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: maxPLMN}, false); err != nil {
			return utils.WrapError("Decode Plmn_index", err)
		}
		ie.Plmn_index = tmp_int_Plmn_index
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
