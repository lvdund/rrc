package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type PLMN_RAN_AreaCell struct {
	Plmn_Identity *PLMN_Identity `optional`
	Ran_AreaCells []CellIdentity `lb:1,ub:32,madatory`
}

func (ie *PLMN_RAN_AreaCell) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Plmn_Identity != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Plmn_Identity != nil {
		if err = ie.Plmn_Identity.Encode(w); err != nil {
			return utils.WrapError("Encode Plmn_Identity", err)
		}
	}
	tmp_Ran_AreaCells := utils.NewSequence[*CellIdentity]([]*CellIdentity{}, aper.Constraint{Lb: 1, Ub: 32}, false)
	for _, i := range ie.Ran_AreaCells {
		tmp_Ran_AreaCells.Value = append(tmp_Ran_AreaCells.Value, &i)
	}
	if err = tmp_Ran_AreaCells.Encode(w); err != nil {
		return utils.WrapError("Encode Ran_AreaCells", err)
	}
	return nil
}

func (ie *PLMN_RAN_AreaCell) Decode(r *aper.AperReader) error {
	var err error
	var Plmn_IdentityPresent bool
	if Plmn_IdentityPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if Plmn_IdentityPresent {
		ie.Plmn_Identity = new(PLMN_Identity)
		if err = ie.Plmn_Identity.Decode(r); err != nil {
			return utils.WrapError("Decode Plmn_Identity", err)
		}
	}
	tmp_Ran_AreaCells := utils.NewSequence[*CellIdentity]([]*CellIdentity{}, aper.Constraint{Lb: 1, Ub: 32}, false)
	fn_Ran_AreaCells := func() *CellIdentity {
		return new(CellIdentity)
	}
	if err = tmp_Ran_AreaCells.Decode(r, fn_Ran_AreaCells); err != nil {
		return utils.WrapError("Decode Ran_AreaCells", err)
	}
	ie.Ran_AreaCells = []CellIdentity{}
	for _, i := range tmp_Ran_AreaCells.Value {
		ie.Ran_AreaCells = append(ie.Ran_AreaCells, *i)
	}
	return nil
}
