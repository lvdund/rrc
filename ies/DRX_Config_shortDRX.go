package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type DRX_Config_shortDRX struct {
	Drx_ShortCycle      DRX_Config_shortDRX_drx_ShortCycle `madatory`
	Drx_ShortCycleTimer int64                              `lb:1,ub:16,madatory`
}

func (ie *DRX_Config_shortDRX) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.Drx_ShortCycle.Encode(w); err != nil {
		return utils.WrapError("Encode Drx_ShortCycle", err)
	}
	if err = w.WriteInteger(ie.Drx_ShortCycleTimer, &aper.Constraint{Lb: 1, Ub: 16}, false); err != nil {
		return utils.WrapError("WriteInteger Drx_ShortCycleTimer", err)
	}
	return nil
}

func (ie *DRX_Config_shortDRX) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.Drx_ShortCycle.Decode(r); err != nil {
		return utils.WrapError("Decode Drx_ShortCycle", err)
	}
	var tmp_int_Drx_ShortCycleTimer int64
	if tmp_int_Drx_ShortCycleTimer, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 16}, false); err != nil {
		return utils.WrapError("ReadInteger Drx_ShortCycleTimer", err)
	}
	ie.Drx_ShortCycleTimer = tmp_int_Drx_ShortCycleTimer
	return nil
}
