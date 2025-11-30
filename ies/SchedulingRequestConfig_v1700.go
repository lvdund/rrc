package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type SchedulingRequestConfig_v1700 struct {
	SchedulingRequestToAddModListExt_v1700 []SchedulingRequestToAddModExt_v1700 `lb:1,ub:maxNrofSR_ConfigPerCellGroup,optional`
}

func (ie *SchedulingRequestConfig_v1700) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{len(ie.SchedulingRequestToAddModListExt_v1700) > 0}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if len(ie.SchedulingRequestToAddModListExt_v1700) > 0 {
		tmp_SchedulingRequestToAddModListExt_v1700 := utils.NewSequence[*SchedulingRequestToAddModExt_v1700]([]*SchedulingRequestToAddModExt_v1700{}, aper.Constraint{Lb: 1, Ub: maxNrofSR_ConfigPerCellGroup}, false)
		for _, i := range ie.SchedulingRequestToAddModListExt_v1700 {
			tmp_SchedulingRequestToAddModListExt_v1700.Value = append(tmp_SchedulingRequestToAddModListExt_v1700.Value, &i)
		}
		if err = tmp_SchedulingRequestToAddModListExt_v1700.Encode(w); err != nil {
			return utils.WrapError("Encode SchedulingRequestToAddModListExt_v1700", err)
		}
	}
	return nil
}

func (ie *SchedulingRequestConfig_v1700) Decode(r *aper.AperReader) error {
	var err error
	var SchedulingRequestToAddModListExt_v1700Present bool
	if SchedulingRequestToAddModListExt_v1700Present, err = r.ReadBool(); err != nil {
		return err
	}
	if SchedulingRequestToAddModListExt_v1700Present {
		tmp_SchedulingRequestToAddModListExt_v1700 := utils.NewSequence[*SchedulingRequestToAddModExt_v1700]([]*SchedulingRequestToAddModExt_v1700{}, aper.Constraint{Lb: 1, Ub: maxNrofSR_ConfigPerCellGroup}, false)
		fn_SchedulingRequestToAddModListExt_v1700 := func() *SchedulingRequestToAddModExt_v1700 {
			return new(SchedulingRequestToAddModExt_v1700)
		}
		if err = tmp_SchedulingRequestToAddModListExt_v1700.Decode(r, fn_SchedulingRequestToAddModListExt_v1700); err != nil {
			return utils.WrapError("Decode SchedulingRequestToAddModListExt_v1700", err)
		}
		ie.SchedulingRequestToAddModListExt_v1700 = []SchedulingRequestToAddModExt_v1700{}
		for _, i := range tmp_SchedulingRequestToAddModListExt_v1700.Value {
			ie.SchedulingRequestToAddModListExt_v1700 = append(ie.SchedulingRequestToAddModListExt_v1700, *i)
		}
	}
	return nil
}
