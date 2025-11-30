package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type SCellActivationRS_Config_r17 struct {
	ScellActivationRS_Id_r17 SCellActivationRS_ConfigId_r17 `madatory`
	ResourceSet_r17          NZP_CSI_RS_ResourceSetId       `madatory`
	GapBetweenBursts_r17     *int64                         `lb:2,ub:31,optional`
	Qcl_Info_r17             TCI_StateId                    `madatory`
}

func (ie *SCellActivationRS_Config_r17) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.GapBetweenBursts_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.ScellActivationRS_Id_r17.Encode(w); err != nil {
		return utils.WrapError("Encode ScellActivationRS_Id_r17", err)
	}
	if err = ie.ResourceSet_r17.Encode(w); err != nil {
		return utils.WrapError("Encode ResourceSet_r17", err)
	}
	if ie.GapBetweenBursts_r17 != nil {
		if err = w.WriteInteger(*ie.GapBetweenBursts_r17, &aper.Constraint{Lb: 2, Ub: 31}, false); err != nil {
			return utils.WrapError("Encode GapBetweenBursts_r17", err)
		}
	}
	if err = ie.Qcl_Info_r17.Encode(w); err != nil {
		return utils.WrapError("Encode Qcl_Info_r17", err)
	}
	return nil
}

func (ie *SCellActivationRS_Config_r17) Decode(r *aper.AperReader) error {
	var err error
	var GapBetweenBursts_r17Present bool
	if GapBetweenBursts_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.ScellActivationRS_Id_r17.Decode(r); err != nil {
		return utils.WrapError("Decode ScellActivationRS_Id_r17", err)
	}
	if err = ie.ResourceSet_r17.Decode(r); err != nil {
		return utils.WrapError("Decode ResourceSet_r17", err)
	}
	if GapBetweenBursts_r17Present {
		var tmp_int_GapBetweenBursts_r17 int64
		if tmp_int_GapBetweenBursts_r17, err = r.ReadInteger(&aper.Constraint{Lb: 2, Ub: 31}, false); err != nil {
			return utils.WrapError("Decode GapBetweenBursts_r17", err)
		}
		ie.GapBetweenBursts_r17 = &tmp_int_GapBetweenBursts_r17
	}
	if err = ie.Qcl_Info_r17.Decode(r); err != nil {
		return utils.WrapError("Decode Qcl_Info_r17", err)
	}
	return nil
}
