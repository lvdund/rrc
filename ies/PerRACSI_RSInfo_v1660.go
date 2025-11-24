package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PerRACSI_RSInfo_v1660 struct {
	Csi_RS_Index_v1660 *int64 `lb:1,ub:96,optional`
}

func (ie *PerRACSI_RSInfo_v1660) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.Csi_RS_Index_v1660 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Csi_RS_Index_v1660 != nil {
		if err = w.WriteInteger(*ie.Csi_RS_Index_v1660, &uper.Constraint{Lb: 1, Ub: 96}, false); err != nil {
			return utils.WrapError("Encode Csi_RS_Index_v1660", err)
		}
	}
	return nil
}

func (ie *PerRACSI_RSInfo_v1660) Decode(r *uper.UperReader) error {
	var err error
	var Csi_RS_Index_v1660Present bool
	if Csi_RS_Index_v1660Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Csi_RS_Index_v1660Present {
		var tmp_int_Csi_RS_Index_v1660 int64
		if tmp_int_Csi_RS_Index_v1660, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 96}, false); err != nil {
			return utils.WrapError("Decode Csi_RS_Index_v1660", err)
		}
		ie.Csi_RS_Index_v1660 = &tmp_int_Csi_RS_Index_v1660
	}
	return nil
}
