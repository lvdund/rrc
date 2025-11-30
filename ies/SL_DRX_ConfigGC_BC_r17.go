package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type SL_DRX_ConfigGC_BC_r17 struct {
	Sl_DRX_GC_BC_PerQoS_List_r17 []SL_DRX_GC_BC_QoS_r17 `lb:1,ub:maxSL_GC_BC_DRX_QoS_r17,optional`
	Sl_DRX_GC_generic_r17        *SL_DRX_GC_Generic_r17 `optional`
	Sl_DefaultDRX_GC_BC_r17      *SL_DRX_GC_BC_QoS_r17  `optional`
}

func (ie *SL_DRX_ConfigGC_BC_r17) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{len(ie.Sl_DRX_GC_BC_PerQoS_List_r17) > 0, ie.Sl_DRX_GC_generic_r17 != nil, ie.Sl_DefaultDRX_GC_BC_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if len(ie.Sl_DRX_GC_BC_PerQoS_List_r17) > 0 {
		tmp_Sl_DRX_GC_BC_PerQoS_List_r17 := utils.NewSequence[*SL_DRX_GC_BC_QoS_r17]([]*SL_DRX_GC_BC_QoS_r17{}, aper.Constraint{Lb: 1, Ub: maxSL_GC_BC_DRX_QoS_r17}, false)
		for _, i := range ie.Sl_DRX_GC_BC_PerQoS_List_r17 {
			tmp_Sl_DRX_GC_BC_PerQoS_List_r17.Value = append(tmp_Sl_DRX_GC_BC_PerQoS_List_r17.Value, &i)
		}
		if err = tmp_Sl_DRX_GC_BC_PerQoS_List_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_DRX_GC_BC_PerQoS_List_r17", err)
		}
	}
	if ie.Sl_DRX_GC_generic_r17 != nil {
		if err = ie.Sl_DRX_GC_generic_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_DRX_GC_generic_r17", err)
		}
	}
	if ie.Sl_DefaultDRX_GC_BC_r17 != nil {
		if err = ie.Sl_DefaultDRX_GC_BC_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_DefaultDRX_GC_BC_r17", err)
		}
	}
	return nil
}

func (ie *SL_DRX_ConfigGC_BC_r17) Decode(r *aper.AperReader) error {
	var err error
	var Sl_DRX_GC_BC_PerQoS_List_r17Present bool
	if Sl_DRX_GC_BC_PerQoS_List_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_DRX_GC_generic_r17Present bool
	if Sl_DRX_GC_generic_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_DefaultDRX_GC_BC_r17Present bool
	if Sl_DefaultDRX_GC_BC_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Sl_DRX_GC_BC_PerQoS_List_r17Present {
		tmp_Sl_DRX_GC_BC_PerQoS_List_r17 := utils.NewSequence[*SL_DRX_GC_BC_QoS_r17]([]*SL_DRX_GC_BC_QoS_r17{}, aper.Constraint{Lb: 1, Ub: maxSL_GC_BC_DRX_QoS_r17}, false)
		fn_Sl_DRX_GC_BC_PerQoS_List_r17 := func() *SL_DRX_GC_BC_QoS_r17 {
			return new(SL_DRX_GC_BC_QoS_r17)
		}
		if err = tmp_Sl_DRX_GC_BC_PerQoS_List_r17.Decode(r, fn_Sl_DRX_GC_BC_PerQoS_List_r17); err != nil {
			return utils.WrapError("Decode Sl_DRX_GC_BC_PerQoS_List_r17", err)
		}
		ie.Sl_DRX_GC_BC_PerQoS_List_r17 = []SL_DRX_GC_BC_QoS_r17{}
		for _, i := range tmp_Sl_DRX_GC_BC_PerQoS_List_r17.Value {
			ie.Sl_DRX_GC_BC_PerQoS_List_r17 = append(ie.Sl_DRX_GC_BC_PerQoS_List_r17, *i)
		}
	}
	if Sl_DRX_GC_generic_r17Present {
		ie.Sl_DRX_GC_generic_r17 = new(SL_DRX_GC_Generic_r17)
		if err = ie.Sl_DRX_GC_generic_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_DRX_GC_generic_r17", err)
		}
	}
	if Sl_DefaultDRX_GC_BC_r17Present {
		ie.Sl_DefaultDRX_GC_BC_r17 = new(SL_DRX_GC_BC_QoS_r17)
		if err = ie.Sl_DefaultDRX_GC_BC_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_DefaultDRX_GC_BC_r17", err)
		}
	}
	return nil
}
