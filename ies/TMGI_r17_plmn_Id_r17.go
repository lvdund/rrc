package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	TMGI_r17_plmn_Id_r17_Choice_nothing uint64 = iota
	TMGI_r17_plmn_Id_r17_Choice_Plmn_Index
	TMGI_r17_plmn_Id_r17_Choice_ExplicitValue
)

type TMGI_r17_plmn_Id_r17 struct {
	Choice        uint64
	Plmn_Index    int64 `lb:1,ub:maxPLMN,madatory`
	ExplicitValue *PLMN_Identity
}

func (ie *TMGI_r17_plmn_Id_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case TMGI_r17_plmn_Id_r17_Choice_Plmn_Index:
		if err = w.WriteInteger(int64(ie.Plmn_Index), &aper.Constraint{Lb: 1, Ub: maxPLMN}, false); err != nil {
			err = utils.WrapError("Encode Plmn_Index", err)
		}
	case TMGI_r17_plmn_Id_r17_Choice_ExplicitValue:
		if err = ie.ExplicitValue.Encode(w); err != nil {
			err = utils.WrapError("Encode ExplicitValue", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *TMGI_r17_plmn_Id_r17) Decode(r *aper.AperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case TMGI_r17_plmn_Id_r17_Choice_Plmn_Index:
		var tmp_int_Plmn_Index int64
		if tmp_int_Plmn_Index, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: maxPLMN}, false); err != nil {
			return utils.WrapError("Decode Plmn_Index", err)
		}
		ie.Plmn_Index = tmp_int_Plmn_Index
	case TMGI_r17_plmn_Id_r17_Choice_ExplicitValue:
		ie.ExplicitValue = new(PLMN_Identity)
		if err = ie.ExplicitValue.Decode(r); err != nil {
			return utils.WrapError("Decode ExplicitValue", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
