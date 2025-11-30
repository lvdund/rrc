package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type UL_PDCP_ExcessDelayResult_r17 struct {
	Drb_Id_r17      DRB_Identity `madatory`
	ExcessDelay_r17 int64        `lb:0,ub:31,madatory`
}

func (ie *UL_PDCP_ExcessDelayResult_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.Drb_Id_r17.Encode(w); err != nil {
		return utils.WrapError("Encode Drb_Id_r17", err)
	}
	if err = w.WriteInteger(ie.ExcessDelay_r17, &aper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
		return utils.WrapError("WriteInteger ExcessDelay_r17", err)
	}
	return nil
}

func (ie *UL_PDCP_ExcessDelayResult_r17) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.Drb_Id_r17.Decode(r); err != nil {
		return utils.WrapError("Decode Drb_Id_r17", err)
	}
	var tmp_int_ExcessDelay_r17 int64
	if tmp_int_ExcessDelay_r17, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
		return utils.WrapError("ReadInteger ExcessDelay_r17", err)
	}
	ie.ExcessDelay_r17 = tmp_int_ExcessDelay_r17
	return nil
}
