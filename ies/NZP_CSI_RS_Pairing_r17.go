package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type NZP_CSI_RS_Pairing_r17 struct {
	Nzp_CSI_RS_ResourceId1_r17 int64 `lb:1,ub:7,madatory`
	Nzp_CSI_RS_ResourceId2_r17 int64 `lb:1,ub:7,madatory`
}

func (ie *NZP_CSI_RS_Pairing_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(ie.Nzp_CSI_RS_ResourceId1_r17, &uper.Constraint{Lb: 1, Ub: 7}, false); err != nil {
		return utils.WrapError("WriteInteger Nzp_CSI_RS_ResourceId1_r17", err)
	}
	if err = w.WriteInteger(ie.Nzp_CSI_RS_ResourceId2_r17, &uper.Constraint{Lb: 1, Ub: 7}, false); err != nil {
		return utils.WrapError("WriteInteger Nzp_CSI_RS_ResourceId2_r17", err)
	}
	return nil
}

func (ie *NZP_CSI_RS_Pairing_r17) Decode(r *uper.UperReader) error {
	var err error
	var tmp_int_Nzp_CSI_RS_ResourceId1_r17 int64
	if tmp_int_Nzp_CSI_RS_ResourceId1_r17, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 7}, false); err != nil {
		return utils.WrapError("ReadInteger Nzp_CSI_RS_ResourceId1_r17", err)
	}
	ie.Nzp_CSI_RS_ResourceId1_r17 = tmp_int_Nzp_CSI_RS_ResourceId1_r17
	var tmp_int_Nzp_CSI_RS_ResourceId2_r17 int64
	if tmp_int_Nzp_CSI_RS_ResourceId2_r17, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 7}, false); err != nil {
		return utils.WrapError("ReadInteger Nzp_CSI_RS_ResourceId2_r17", err)
	}
	ie.Nzp_CSI_RS_ResourceId2_r17 = tmp_int_Nzp_CSI_RS_ResourceId2_r17
	return nil
}
