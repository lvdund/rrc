package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type SL_ConfigDedicatedEUTRA_Info_r16 struct {
	Sl_ConfigDedicatedEUTRA_r16 *[]byte                  `optional`
	Sl_TimeOffsetEUTRA_List_r16 []SL_TimeOffsetEUTRA_r16 `lb:8,ub:8,optional`
}

func (ie *SL_ConfigDedicatedEUTRA_Info_r16) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Sl_ConfigDedicatedEUTRA_r16 != nil, len(ie.Sl_TimeOffsetEUTRA_List_r16) > 0}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Sl_ConfigDedicatedEUTRA_r16 != nil {
		if err = w.WriteOctetString(*ie.Sl_ConfigDedicatedEUTRA_r16, &aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode Sl_ConfigDedicatedEUTRA_r16", err)
		}
	}
	if len(ie.Sl_TimeOffsetEUTRA_List_r16) > 0 {
		tmp_Sl_TimeOffsetEUTRA_List_r16 := utils.NewSequence[*SL_TimeOffsetEUTRA_r16]([]*SL_TimeOffsetEUTRA_r16{}, aper.Constraint{Lb: 8, Ub: 8}, false)
		for _, i := range ie.Sl_TimeOffsetEUTRA_List_r16 {
			tmp_Sl_TimeOffsetEUTRA_List_r16.Value = append(tmp_Sl_TimeOffsetEUTRA_List_r16.Value, &i)
		}
		if err = tmp_Sl_TimeOffsetEUTRA_List_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_TimeOffsetEUTRA_List_r16", err)
		}
	}
	return nil
}

func (ie *SL_ConfigDedicatedEUTRA_Info_r16) Decode(r *aper.AperReader) error {
	var err error
	var Sl_ConfigDedicatedEUTRA_r16Present bool
	if Sl_ConfigDedicatedEUTRA_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_TimeOffsetEUTRA_List_r16Present bool
	if Sl_TimeOffsetEUTRA_List_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Sl_ConfigDedicatedEUTRA_r16Present {
		var tmp_os_Sl_ConfigDedicatedEUTRA_r16 []byte
		if tmp_os_Sl_ConfigDedicatedEUTRA_r16, err = r.ReadOctetString(&aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode Sl_ConfigDedicatedEUTRA_r16", err)
		}
		ie.Sl_ConfigDedicatedEUTRA_r16 = &tmp_os_Sl_ConfigDedicatedEUTRA_r16
	}
	if Sl_TimeOffsetEUTRA_List_r16Present {
		tmp_Sl_TimeOffsetEUTRA_List_r16 := utils.NewSequence[*SL_TimeOffsetEUTRA_r16]([]*SL_TimeOffsetEUTRA_r16{}, aper.Constraint{Lb: 8, Ub: 8}, false)
		fn_Sl_TimeOffsetEUTRA_List_r16 := func() *SL_TimeOffsetEUTRA_r16 {
			return new(SL_TimeOffsetEUTRA_r16)
		}
		if err = tmp_Sl_TimeOffsetEUTRA_List_r16.Decode(r, fn_Sl_TimeOffsetEUTRA_List_r16); err != nil {
			return utils.WrapError("Decode Sl_TimeOffsetEUTRA_List_r16", err)
		}
		ie.Sl_TimeOffsetEUTRA_List_r16 = []SL_TimeOffsetEUTRA_r16{}
		for _, i := range tmp_Sl_TimeOffsetEUTRA_List_r16.Value {
			ie.Sl_TimeOffsetEUTRA_List_r16 = append(ie.Sl_TimeOffsetEUTRA_List_r16, *i)
		}
	}
	return nil
}
