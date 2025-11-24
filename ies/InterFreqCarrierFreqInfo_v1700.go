package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type InterFreqCarrierFreqInfo_v1700 struct {
	InterFreqNeighHSDN_CellList_r17 *InterFreqNeighHSDN_CellList_r17                           `optional`
	HighSpeedMeasInterFreq_r17      *InterFreqCarrierFreqInfo_v1700_highSpeedMeasInterFreq_r17 `optional`
	RedCapAccessAllowed_r17         *InterFreqCarrierFreqInfo_v1700_redCapAccessAllowed_r17    `optional`
	Ssb_PositionQCL_Common_r17      *SSB_PositionQCL_Relation_r17                              `optional`
	InterFreqNeighCellList_v1710    *InterFreqNeighCellList_v1710                              `optional`
}

func (ie *InterFreqCarrierFreqInfo_v1700) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.InterFreqNeighHSDN_CellList_r17 != nil, ie.HighSpeedMeasInterFreq_r17 != nil, ie.RedCapAccessAllowed_r17 != nil, ie.Ssb_PositionQCL_Common_r17 != nil, ie.InterFreqNeighCellList_v1710 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.InterFreqNeighHSDN_CellList_r17 != nil {
		if err = ie.InterFreqNeighHSDN_CellList_r17.Encode(w); err != nil {
			return utils.WrapError("Encode InterFreqNeighHSDN_CellList_r17", err)
		}
	}
	if ie.HighSpeedMeasInterFreq_r17 != nil {
		if err = ie.HighSpeedMeasInterFreq_r17.Encode(w); err != nil {
			return utils.WrapError("Encode HighSpeedMeasInterFreq_r17", err)
		}
	}
	if ie.RedCapAccessAllowed_r17 != nil {
		if err = ie.RedCapAccessAllowed_r17.Encode(w); err != nil {
			return utils.WrapError("Encode RedCapAccessAllowed_r17", err)
		}
	}
	if ie.Ssb_PositionQCL_Common_r17 != nil {
		if err = ie.Ssb_PositionQCL_Common_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Ssb_PositionQCL_Common_r17", err)
		}
	}
	if ie.InterFreqNeighCellList_v1710 != nil {
		if err = ie.InterFreqNeighCellList_v1710.Encode(w); err != nil {
			return utils.WrapError("Encode InterFreqNeighCellList_v1710", err)
		}
	}
	return nil
}

func (ie *InterFreqCarrierFreqInfo_v1700) Decode(r *uper.UperReader) error {
	var err error
	var InterFreqNeighHSDN_CellList_r17Present bool
	if InterFreqNeighHSDN_CellList_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var HighSpeedMeasInterFreq_r17Present bool
	if HighSpeedMeasInterFreq_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var RedCapAccessAllowed_r17Present bool
	if RedCapAccessAllowed_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Ssb_PositionQCL_Common_r17Present bool
	if Ssb_PositionQCL_Common_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var InterFreqNeighCellList_v1710Present bool
	if InterFreqNeighCellList_v1710Present, err = r.ReadBool(); err != nil {
		return err
	}
	if InterFreqNeighHSDN_CellList_r17Present {
		ie.InterFreqNeighHSDN_CellList_r17 = new(InterFreqNeighHSDN_CellList_r17)
		if err = ie.InterFreqNeighHSDN_CellList_r17.Decode(r); err != nil {
			return utils.WrapError("Decode InterFreqNeighHSDN_CellList_r17", err)
		}
	}
	if HighSpeedMeasInterFreq_r17Present {
		ie.HighSpeedMeasInterFreq_r17 = new(InterFreqCarrierFreqInfo_v1700_highSpeedMeasInterFreq_r17)
		if err = ie.HighSpeedMeasInterFreq_r17.Decode(r); err != nil {
			return utils.WrapError("Decode HighSpeedMeasInterFreq_r17", err)
		}
	}
	if RedCapAccessAllowed_r17Present {
		ie.RedCapAccessAllowed_r17 = new(InterFreqCarrierFreqInfo_v1700_redCapAccessAllowed_r17)
		if err = ie.RedCapAccessAllowed_r17.Decode(r); err != nil {
			return utils.WrapError("Decode RedCapAccessAllowed_r17", err)
		}
	}
	if Ssb_PositionQCL_Common_r17Present {
		ie.Ssb_PositionQCL_Common_r17 = new(SSB_PositionQCL_Relation_r17)
		if err = ie.Ssb_PositionQCL_Common_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Ssb_PositionQCL_Common_r17", err)
		}
	}
	if InterFreqNeighCellList_v1710Present {
		ie.InterFreqNeighCellList_v1710 = new(InterFreqNeighCellList_v1710)
		if err = ie.InterFreqNeighCellList_v1710.Decode(r); err != nil {
			return utils.WrapError("Decode InterFreqNeighCellList_v1710", err)
		}
	}
	return nil
}
