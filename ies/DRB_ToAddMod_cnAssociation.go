package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	DRB_ToAddMod_cnAssociation_Choice_nothing uint64 = iota
	DRB_ToAddMod_cnAssociation_Choice_Eps_BearerIdentity
	DRB_ToAddMod_cnAssociation_Choice_Sdap_Config
)

type DRB_ToAddMod_cnAssociation struct {
	Choice             uint64
	Eps_BearerIdentity int64 `lb:0,ub:15,madatory`
	Sdap_Config        *SDAP_Config
}

func (ie *DRB_ToAddMod_cnAssociation) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case DRB_ToAddMod_cnAssociation_Choice_Eps_BearerIdentity:
		if err = w.WriteInteger(int64(ie.Eps_BearerIdentity), &aper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
			err = utils.WrapError("Encode Eps_BearerIdentity", err)
		}
	case DRB_ToAddMod_cnAssociation_Choice_Sdap_Config:
		if err = ie.Sdap_Config.Encode(w); err != nil {
			err = utils.WrapError("Encode Sdap_Config", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *DRB_ToAddMod_cnAssociation) Decode(r *aper.AperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case DRB_ToAddMod_cnAssociation_Choice_Eps_BearerIdentity:
		var tmp_int_Eps_BearerIdentity int64
		if tmp_int_Eps_BearerIdentity, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
			return utils.WrapError("Decode Eps_BearerIdentity", err)
		}
		ie.Eps_BearerIdentity = tmp_int_Eps_BearerIdentity
	case DRB_ToAddMod_cnAssociation_Choice_Sdap_Config:
		ie.Sdap_Config = new(SDAP_Config)
		if err = ie.Sdap_Config.Decode(r); err != nil {
			return utils.WrapError("Decode Sdap_Config", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
