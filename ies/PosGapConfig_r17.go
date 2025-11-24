package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PosGapConfig_r17 struct {
	MeasPosPreConfigGapId_r17 MeasPosPreConfigGapId_r17    `madatory`
	GapOffset_r17             int64                        `lb:0,ub:159,madatory`
	Mgl_r17                   PosGapConfig_r17_mgl_r17     `madatory`
	Mgrp_r17                  PosGapConfig_r17_mgrp_r17    `madatory`
	Mgta_r17                  PosGapConfig_r17_mgta_r17    `madatory`
	GapType_r17               PosGapConfig_r17_gapType_r17 `madatory`
}

func (ie *PosGapConfig_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.MeasPosPreConfigGapId_r17.Encode(w); err != nil {
		return utils.WrapError("Encode MeasPosPreConfigGapId_r17", err)
	}
	if err = w.WriteInteger(ie.GapOffset_r17, &uper.Constraint{Lb: 0, Ub: 159}, false); err != nil {
		return utils.WrapError("WriteInteger GapOffset_r17", err)
	}
	if err = ie.Mgl_r17.Encode(w); err != nil {
		return utils.WrapError("Encode Mgl_r17", err)
	}
	if err = ie.Mgrp_r17.Encode(w); err != nil {
		return utils.WrapError("Encode Mgrp_r17", err)
	}
	if err = ie.Mgta_r17.Encode(w); err != nil {
		return utils.WrapError("Encode Mgta_r17", err)
	}
	if err = ie.GapType_r17.Encode(w); err != nil {
		return utils.WrapError("Encode GapType_r17", err)
	}
	return nil
}

func (ie *PosGapConfig_r17) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.MeasPosPreConfigGapId_r17.Decode(r); err != nil {
		return utils.WrapError("Decode MeasPosPreConfigGapId_r17", err)
	}
	var tmp_int_GapOffset_r17 int64
	if tmp_int_GapOffset_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 159}, false); err != nil {
		return utils.WrapError("ReadInteger GapOffset_r17", err)
	}
	ie.GapOffset_r17 = tmp_int_GapOffset_r17
	if err = ie.Mgl_r17.Decode(r); err != nil {
		return utils.WrapError("Decode Mgl_r17", err)
	}
	if err = ie.Mgrp_r17.Decode(r); err != nil {
		return utils.WrapError("Decode Mgrp_r17", err)
	}
	if err = ie.Mgta_r17.Decode(r); err != nil {
		return utils.WrapError("Decode Mgta_r17", err)
	}
	if err = ie.GapType_r17.Decode(r); err != nil {
		return utils.WrapError("Decode GapType_r17", err)
	}
	return nil
}
