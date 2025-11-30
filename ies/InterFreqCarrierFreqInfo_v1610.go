package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type InterFreqCarrierFreqInfo_v1610 struct {
	InterFreqNeighCellList_v1610 *InterFreqNeighCellList_v1610      `optional`
	Smtc2_LP_r16                 *SSB_MTC2_LP_r16                   `optional`
	InterFreqAllowedCellList_r16 *InterFreqAllowedCellList_r16      `optional`
	Ssb_PositionQCL_Common_r16   *SSB_PositionQCL_Relation_r16      `optional`
	InterFreqCAG_CellList_r16    []InterFreqCAG_CellListPerPLMN_r16 `lb:1,ub:maxPLMN,optional`
}

func (ie *InterFreqCarrierFreqInfo_v1610) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.InterFreqNeighCellList_v1610 != nil, ie.Smtc2_LP_r16 != nil, ie.InterFreqAllowedCellList_r16 != nil, ie.Ssb_PositionQCL_Common_r16 != nil, len(ie.InterFreqCAG_CellList_r16) > 0}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.InterFreqNeighCellList_v1610 != nil {
		if err = ie.InterFreqNeighCellList_v1610.Encode(w); err != nil {
			return utils.WrapError("Encode InterFreqNeighCellList_v1610", err)
		}
	}
	if ie.Smtc2_LP_r16 != nil {
		if err = ie.Smtc2_LP_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Smtc2_LP_r16", err)
		}
	}
	if ie.InterFreqAllowedCellList_r16 != nil {
		if err = ie.InterFreqAllowedCellList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode InterFreqAllowedCellList_r16", err)
		}
	}
	if ie.Ssb_PositionQCL_Common_r16 != nil {
		if err = ie.Ssb_PositionQCL_Common_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Ssb_PositionQCL_Common_r16", err)
		}
	}
	if len(ie.InterFreqCAG_CellList_r16) > 0 {
		tmp_InterFreqCAG_CellList_r16 := utils.NewSequence[*InterFreqCAG_CellListPerPLMN_r16]([]*InterFreqCAG_CellListPerPLMN_r16{}, aper.Constraint{Lb: 1, Ub: maxPLMN}, false)
		for _, i := range ie.InterFreqCAG_CellList_r16 {
			tmp_InterFreqCAG_CellList_r16.Value = append(tmp_InterFreqCAG_CellList_r16.Value, &i)
		}
		if err = tmp_InterFreqCAG_CellList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode InterFreqCAG_CellList_r16", err)
		}
	}
	return nil
}

func (ie *InterFreqCarrierFreqInfo_v1610) Decode(r *aper.AperReader) error {
	var err error
	var InterFreqNeighCellList_v1610Present bool
	if InterFreqNeighCellList_v1610Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Smtc2_LP_r16Present bool
	if Smtc2_LP_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var InterFreqAllowedCellList_r16Present bool
	if InterFreqAllowedCellList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Ssb_PositionQCL_Common_r16Present bool
	if Ssb_PositionQCL_Common_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var InterFreqCAG_CellList_r16Present bool
	if InterFreqCAG_CellList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if InterFreqNeighCellList_v1610Present {
		ie.InterFreqNeighCellList_v1610 = new(InterFreqNeighCellList_v1610)
		if err = ie.InterFreqNeighCellList_v1610.Decode(r); err != nil {
			return utils.WrapError("Decode InterFreqNeighCellList_v1610", err)
		}
	}
	if Smtc2_LP_r16Present {
		ie.Smtc2_LP_r16 = new(SSB_MTC2_LP_r16)
		if err = ie.Smtc2_LP_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Smtc2_LP_r16", err)
		}
	}
	if InterFreqAllowedCellList_r16Present {
		ie.InterFreqAllowedCellList_r16 = new(InterFreqAllowedCellList_r16)
		if err = ie.InterFreqAllowedCellList_r16.Decode(r); err != nil {
			return utils.WrapError("Decode InterFreqAllowedCellList_r16", err)
		}
	}
	if Ssb_PositionQCL_Common_r16Present {
		ie.Ssb_PositionQCL_Common_r16 = new(SSB_PositionQCL_Relation_r16)
		if err = ie.Ssb_PositionQCL_Common_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Ssb_PositionQCL_Common_r16", err)
		}
	}
	if InterFreqCAG_CellList_r16Present {
		tmp_InterFreqCAG_CellList_r16 := utils.NewSequence[*InterFreqCAG_CellListPerPLMN_r16]([]*InterFreqCAG_CellListPerPLMN_r16{}, aper.Constraint{Lb: 1, Ub: maxPLMN}, false)
		fn_InterFreqCAG_CellList_r16 := func() *InterFreqCAG_CellListPerPLMN_r16 {
			return new(InterFreqCAG_CellListPerPLMN_r16)
		}
		if err = tmp_InterFreqCAG_CellList_r16.Decode(r, fn_InterFreqCAG_CellList_r16); err != nil {
			return utils.WrapError("Decode InterFreqCAG_CellList_r16", err)
		}
		ie.InterFreqCAG_CellList_r16 = []InterFreqCAG_CellListPerPLMN_r16{}
		for _, i := range tmp_InterFreqCAG_CellList_r16.Value {
			ie.InterFreqCAG_CellList_r16 = append(ie.InterFreqCAG_CellList_r16, *i)
		}
	}
	return nil
}
