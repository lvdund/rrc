package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type UE_RadioPagingInfo_r17 struct {
	Pei_SubgroupingSupportBandList_r17 []FreqBandIndicatorNR `lb:1,ub:maxBands,optional`
}

func (ie *UE_RadioPagingInfo_r17) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{len(ie.Pei_SubgroupingSupportBandList_r17) > 0}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if len(ie.Pei_SubgroupingSupportBandList_r17) > 0 {
		tmp_Pei_SubgroupingSupportBandList_r17 := utils.NewSequence[*FreqBandIndicatorNR]([]*FreqBandIndicatorNR{}, aper.Constraint{Lb: 1, Ub: maxBands}, false)
		for _, i := range ie.Pei_SubgroupingSupportBandList_r17 {
			tmp_Pei_SubgroupingSupportBandList_r17.Value = append(tmp_Pei_SubgroupingSupportBandList_r17.Value, &i)
		}
		if err = tmp_Pei_SubgroupingSupportBandList_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Pei_SubgroupingSupportBandList_r17", err)
		}
	}
	return nil
}

func (ie *UE_RadioPagingInfo_r17) Decode(r *aper.AperReader) error {
	var err error
	var Pei_SubgroupingSupportBandList_r17Present bool
	if Pei_SubgroupingSupportBandList_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Pei_SubgroupingSupportBandList_r17Present {
		tmp_Pei_SubgroupingSupportBandList_r17 := utils.NewSequence[*FreqBandIndicatorNR]([]*FreqBandIndicatorNR{}, aper.Constraint{Lb: 1, Ub: maxBands}, false)
		fn_Pei_SubgroupingSupportBandList_r17 := func() *FreqBandIndicatorNR {
			return new(FreqBandIndicatorNR)
		}
		if err = tmp_Pei_SubgroupingSupportBandList_r17.Decode(r, fn_Pei_SubgroupingSupportBandList_r17); err != nil {
			return utils.WrapError("Decode Pei_SubgroupingSupportBandList_r17", err)
		}
		ie.Pei_SubgroupingSupportBandList_r17 = []FreqBandIndicatorNR{}
		for _, i := range tmp_Pei_SubgroupingSupportBandList_r17.Value {
			ie.Pei_SubgroupingSupportBandList_r17 = append(ie.Pei_SubgroupingSupportBandList_r17, *i)
		}
	}
	return nil
}
