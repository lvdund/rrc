package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SI_RequestResources struct {
	Ra_PreambleStartIndex     int64  `lb:0,ub:63,madatory`
	Ra_AssociationPeriodIndex *int64 `lb:0,ub:15,optional`
	Ra_ssb_OccasionMaskIndex  *int64 `lb:0,ub:15,optional`
}

func (ie *SI_RequestResources) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.Ra_AssociationPeriodIndex != nil, ie.Ra_ssb_OccasionMaskIndex != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = w.WriteInteger(ie.Ra_PreambleStartIndex, &uper.Constraint{Lb: 0, Ub: 63}, false); err != nil {
		return utils.WrapError("WriteInteger Ra_PreambleStartIndex", err)
	}
	if ie.Ra_AssociationPeriodIndex != nil {
		if err = w.WriteInteger(*ie.Ra_AssociationPeriodIndex, &uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
			return utils.WrapError("Encode Ra_AssociationPeriodIndex", err)
		}
	}
	if ie.Ra_ssb_OccasionMaskIndex != nil {
		if err = w.WriteInteger(*ie.Ra_ssb_OccasionMaskIndex, &uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
			return utils.WrapError("Encode Ra_ssb_OccasionMaskIndex", err)
		}
	}
	return nil
}

func (ie *SI_RequestResources) Decode(r *uper.UperReader) error {
	var err error
	var Ra_AssociationPeriodIndexPresent bool
	if Ra_AssociationPeriodIndexPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Ra_ssb_OccasionMaskIndexPresent bool
	if Ra_ssb_OccasionMaskIndexPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var tmp_int_Ra_PreambleStartIndex int64
	if tmp_int_Ra_PreambleStartIndex, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 63}, false); err != nil {
		return utils.WrapError("ReadInteger Ra_PreambleStartIndex", err)
	}
	ie.Ra_PreambleStartIndex = tmp_int_Ra_PreambleStartIndex
	if Ra_AssociationPeriodIndexPresent {
		var tmp_int_Ra_AssociationPeriodIndex int64
		if tmp_int_Ra_AssociationPeriodIndex, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
			return utils.WrapError("Decode Ra_AssociationPeriodIndex", err)
		}
		ie.Ra_AssociationPeriodIndex = &tmp_int_Ra_AssociationPeriodIndex
	}
	if Ra_ssb_OccasionMaskIndexPresent {
		var tmp_int_Ra_ssb_OccasionMaskIndex int64
		if tmp_int_Ra_ssb_OccasionMaskIndex, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
			return utils.WrapError("Decode Ra_ssb_OccasionMaskIndex", err)
		}
		ie.Ra_ssb_OccasionMaskIndex = &tmp_int_Ra_ssb_OccasionMaskIndex
	}
	return nil
}
