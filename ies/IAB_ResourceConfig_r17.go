package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type IAB_ResourceConfig_r17 struct {
	Iab_ResourceConfigID_r17      IAB_ResourceConfigID_r17                        `madatory`
	SlotList_r17                  []int64                                         `lb:1,ub:5120,e_lb:0,e_ub:5119,optional`
	PeriodicitySlotList_r17       *IAB_ResourceConfig_r17_periodicitySlotList_r17 `optional`
	SlotListSubcarrierSpacing_r17 *SubcarrierSpacing                              `optional`
}

func (ie *IAB_ResourceConfig_r17) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{len(ie.SlotList_r17) > 0, ie.PeriodicitySlotList_r17 != nil, ie.SlotListSubcarrierSpacing_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.Iab_ResourceConfigID_r17.Encode(w); err != nil {
		return utils.WrapError("Encode Iab_ResourceConfigID_r17", err)
	}
	if len(ie.SlotList_r17) > 0 {
		tmp_SlotList_r17 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, aper.Constraint{Lb: 1, Ub: 5120}, false)
		for _, i := range ie.SlotList_r17 {
			tmp_ie := utils.NewINTEGER(int64(i), aper.Constraint{Lb: 0, Ub: 5119}, false)
			tmp_SlotList_r17.Value = append(tmp_SlotList_r17.Value, &tmp_ie)
		}
		if err = tmp_SlotList_r17.Encode(w); err != nil {
			return utils.WrapError("Encode SlotList_r17", err)
		}
	}
	if ie.PeriodicitySlotList_r17 != nil {
		if err = ie.PeriodicitySlotList_r17.Encode(w); err != nil {
			return utils.WrapError("Encode PeriodicitySlotList_r17", err)
		}
	}
	if ie.SlotListSubcarrierSpacing_r17 != nil {
		if err = ie.SlotListSubcarrierSpacing_r17.Encode(w); err != nil {
			return utils.WrapError("Encode SlotListSubcarrierSpacing_r17", err)
		}
	}
	return nil
}

func (ie *IAB_ResourceConfig_r17) Decode(r *aper.AperReader) error {
	var err error
	var SlotList_r17Present bool
	if SlotList_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var PeriodicitySlotList_r17Present bool
	if PeriodicitySlotList_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var SlotListSubcarrierSpacing_r17Present bool
	if SlotListSubcarrierSpacing_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.Iab_ResourceConfigID_r17.Decode(r); err != nil {
		return utils.WrapError("Decode Iab_ResourceConfigID_r17", err)
	}
	if SlotList_r17Present {
		tmp_SlotList_r17 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, aper.Constraint{Lb: 1, Ub: 5120}, false)
		fn_SlotList_r17 := func() *utils.INTEGER {
			ie := utils.NewINTEGER(0, aper.Constraint{Lb: 0, Ub: 5119}, false)
			return &ie
		}
		if err = tmp_SlotList_r17.Decode(r, fn_SlotList_r17); err != nil {
			return utils.WrapError("Decode SlotList_r17", err)
		}
		ie.SlotList_r17 = []int64{}
		for _, i := range tmp_SlotList_r17.Value {
			ie.SlotList_r17 = append(ie.SlotList_r17, int64(i.Value))
		}
	}
	if PeriodicitySlotList_r17Present {
		ie.PeriodicitySlotList_r17 = new(IAB_ResourceConfig_r17_periodicitySlotList_r17)
		if err = ie.PeriodicitySlotList_r17.Decode(r); err != nil {
			return utils.WrapError("Decode PeriodicitySlotList_r17", err)
		}
	}
	if SlotListSubcarrierSpacing_r17Present {
		ie.SlotListSubcarrierSpacing_r17 = new(SubcarrierSpacing)
		if err = ie.SlotListSubcarrierSpacing_r17.Decode(r); err != nil {
			return utils.WrapError("Decode SlotListSubcarrierSpacing_r17", err)
		}
	}
	return nil
}
