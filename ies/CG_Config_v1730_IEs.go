package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type CG_Config_v1730_IEs struct {
	Fr1_Carriers_SCG_r17 *int64      `lb:1,ub:32,optional`
	Fr2_Carriers_SCG_r17 *int64      `lb:1,ub:32,optional`
	NonCriticalExtension interface{} `optional`
}

func (ie *CG_Config_v1730_IEs) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Fr1_Carriers_SCG_r17 != nil, ie.Fr2_Carriers_SCG_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Fr1_Carriers_SCG_r17 != nil {
		if err = w.WriteInteger(*ie.Fr1_Carriers_SCG_r17, &aper.Constraint{Lb: 1, Ub: 32}, false); err != nil {
			return utils.WrapError("Encode Fr1_Carriers_SCG_r17", err)
		}
	}
	if ie.Fr2_Carriers_SCG_r17 != nil {
		if err = w.WriteInteger(*ie.Fr2_Carriers_SCG_r17, &aper.Constraint{Lb: 1, Ub: 32}, false); err != nil {
			return utils.WrapError("Encode Fr2_Carriers_SCG_r17", err)
		}
	}
	return nil
}

func (ie *CG_Config_v1730_IEs) Decode(r *aper.AperReader) error {
	var err error
	var Fr1_Carriers_SCG_r17Present bool
	if Fr1_Carriers_SCG_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Fr2_Carriers_SCG_r17Present bool
	if Fr2_Carriers_SCG_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Fr1_Carriers_SCG_r17Present {
		var tmp_int_Fr1_Carriers_SCG_r17 int64
		if tmp_int_Fr1_Carriers_SCG_r17, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 32}, false); err != nil {
			return utils.WrapError("Decode Fr1_Carriers_SCG_r17", err)
		}
		ie.Fr1_Carriers_SCG_r17 = &tmp_int_Fr1_Carriers_SCG_r17
	}
	if Fr2_Carriers_SCG_r17Present {
		var tmp_int_Fr2_Carriers_SCG_r17 int64
		if tmp_int_Fr2_Carriers_SCG_r17, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 32}, false); err != nil {
			return utils.WrapError("Decode Fr2_Carriers_SCG_r17", err)
		}
		ie.Fr2_Carriers_SCG_r17 = &tmp_int_Fr2_Carriers_SCG_r17
	}
	return nil
}
