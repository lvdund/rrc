package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type PropDelayDiffReportConfig_r17 struct {
	ThreshPropDelayDiff_r17 *PropDelayDiffReportConfig_r17_threshPropDelayDiff_r17 `optional`
	NeighCellInfoList_r17   []NeighbourCellInfo_r17                                `lb:1,ub:maxCellNTN_r17,optional`
}

func (ie *PropDelayDiffReportConfig_r17) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.ThreshPropDelayDiff_r17 != nil, len(ie.NeighCellInfoList_r17) > 0}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.ThreshPropDelayDiff_r17 != nil {
		if err = ie.ThreshPropDelayDiff_r17.Encode(w); err != nil {
			return utils.WrapError("Encode ThreshPropDelayDiff_r17", err)
		}
	}
	if len(ie.NeighCellInfoList_r17) > 0 {
		tmp_NeighCellInfoList_r17 := utils.NewSequence[*NeighbourCellInfo_r17]([]*NeighbourCellInfo_r17{}, aper.Constraint{Lb: 1, Ub: maxCellNTN_r17}, false)
		for _, i := range ie.NeighCellInfoList_r17 {
			tmp_NeighCellInfoList_r17.Value = append(tmp_NeighCellInfoList_r17.Value, &i)
		}
		if err = tmp_NeighCellInfoList_r17.Encode(w); err != nil {
			return utils.WrapError("Encode NeighCellInfoList_r17", err)
		}
	}
	return nil
}

func (ie *PropDelayDiffReportConfig_r17) Decode(r *aper.AperReader) error {
	var err error
	var ThreshPropDelayDiff_r17Present bool
	if ThreshPropDelayDiff_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var NeighCellInfoList_r17Present bool
	if NeighCellInfoList_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if ThreshPropDelayDiff_r17Present {
		ie.ThreshPropDelayDiff_r17 = new(PropDelayDiffReportConfig_r17_threshPropDelayDiff_r17)
		if err = ie.ThreshPropDelayDiff_r17.Decode(r); err != nil {
			return utils.WrapError("Decode ThreshPropDelayDiff_r17", err)
		}
	}
	if NeighCellInfoList_r17Present {
		tmp_NeighCellInfoList_r17 := utils.NewSequence[*NeighbourCellInfo_r17]([]*NeighbourCellInfo_r17{}, aper.Constraint{Lb: 1, Ub: maxCellNTN_r17}, false)
		fn_NeighCellInfoList_r17 := func() *NeighbourCellInfo_r17 {
			return new(NeighbourCellInfo_r17)
		}
		if err = tmp_NeighCellInfoList_r17.Decode(r, fn_NeighCellInfoList_r17); err != nil {
			return utils.WrapError("Decode NeighCellInfoList_r17", err)
		}
		ie.NeighCellInfoList_r17 = []NeighbourCellInfo_r17{}
		for _, i := range tmp_NeighCellInfoList_r17.Value {
			ie.NeighCellInfoList_r17 = append(ie.NeighCellInfoList_r17, *i)
		}
	}
	return nil
}
