package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type DL_PRS_Info_r16 struct {
	Dl_PRS_ID_r16            int64  `lb:0,ub:255,madatory`
	Dl_PRS_ResourceSetId_r16 int64  `lb:0,ub:7,madatory`
	Dl_PRS_ResourceId_r16    *int64 `lb:0,ub:63,optional`
}

func (ie *DL_PRS_Info_r16) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Dl_PRS_ResourceId_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = w.WriteInteger(ie.Dl_PRS_ID_r16, &aper.Constraint{Lb: 0, Ub: 255}, false); err != nil {
		return utils.WrapError("WriteInteger Dl_PRS_ID_r16", err)
	}
	if err = w.WriteInteger(ie.Dl_PRS_ResourceSetId_r16, &aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("WriteInteger Dl_PRS_ResourceSetId_r16", err)
	}
	if ie.Dl_PRS_ResourceId_r16 != nil {
		if err = w.WriteInteger(*ie.Dl_PRS_ResourceId_r16, &aper.Constraint{Lb: 0, Ub: 63}, false); err != nil {
			return utils.WrapError("Encode Dl_PRS_ResourceId_r16", err)
		}
	}
	return nil
}

func (ie *DL_PRS_Info_r16) Decode(r *aper.AperReader) error {
	var err error
	var Dl_PRS_ResourceId_r16Present bool
	if Dl_PRS_ResourceId_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var tmp_int_Dl_PRS_ID_r16 int64
	if tmp_int_Dl_PRS_ID_r16, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 255}, false); err != nil {
		return utils.WrapError("ReadInteger Dl_PRS_ID_r16", err)
	}
	ie.Dl_PRS_ID_r16 = tmp_int_Dl_PRS_ID_r16
	var tmp_int_Dl_PRS_ResourceSetId_r16 int64
	if tmp_int_Dl_PRS_ResourceSetId_r16, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("ReadInteger Dl_PRS_ResourceSetId_r16", err)
	}
	ie.Dl_PRS_ResourceSetId_r16 = tmp_int_Dl_PRS_ResourceSetId_r16
	if Dl_PRS_ResourceId_r16Present {
		var tmp_int_Dl_PRS_ResourceId_r16 int64
		if tmp_int_Dl_PRS_ResourceId_r16, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 63}, false); err != nil {
			return utils.WrapError("Decode Dl_PRS_ResourceId_r16", err)
		}
		ie.Dl_PRS_ResourceId_r16 = &tmp_int_Dl_PRS_ResourceId_r16
	}
	return nil
}
