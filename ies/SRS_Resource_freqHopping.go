package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SRS_Resource_freqHopping struct {
	C_SRS int64 `lb:0,ub:63,madatory`
	B_SRS int64 `lb:0,ub:3,madatory`
	B_hop int64 `lb:0,ub:3,madatory`
}

func (ie *SRS_Resource_freqHopping) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(ie.C_SRS, &uper.Constraint{Lb: 0, Ub: 63}, false); err != nil {
		return utils.WrapError("WriteInteger C_SRS", err)
	}
	if err = w.WriteInteger(ie.B_SRS, &uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("WriteInteger B_SRS", err)
	}
	if err = w.WriteInteger(ie.B_hop, &uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("WriteInteger B_hop", err)
	}
	return nil
}

func (ie *SRS_Resource_freqHopping) Decode(r *uper.UperReader) error {
	var err error
	var tmp_int_C_SRS int64
	if tmp_int_C_SRS, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 63}, false); err != nil {
		return utils.WrapError("ReadInteger C_SRS", err)
	}
	ie.C_SRS = tmp_int_C_SRS
	var tmp_int_B_SRS int64
	if tmp_int_B_SRS, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("ReadInteger B_SRS", err)
	}
	ie.B_SRS = tmp_int_B_SRS
	var tmp_int_B_hop int64
	if tmp_int_B_hop, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("ReadInteger B_hop", err)
	}
	ie.B_hop = tmp_int_B_hop
	return nil
}
