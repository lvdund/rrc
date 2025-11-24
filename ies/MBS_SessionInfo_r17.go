package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MBS_SessionInfo_r17 struct {
	Mbs_SessionId_r17               TMGI_r17                         `madatory`
	G_RNTI_r17                      RNTI_Value                       `madatory`
	Mrb_ListBroadcast_r17           MRB_ListBroadcast_r17            `madatory`
	Mtch_SchedulingInfo_r17         *DRX_ConfigPTM_Index_r17         `optional`
	Mtch_NeighbourCell_r17          *uper.BitString                  `lb:maxNeighCellMBS_r17,ub:maxNeighCellMBS_r17,optional`
	Pdsch_ConfigIndex_r17           *PDSCH_ConfigIndex_r17           `optional`
	Mtch_SSB_MappingWindowIndex_r17 *MTCH_SSB_MappingWindowIndex_r17 `optional`
}

func (ie *MBS_SessionInfo_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.Mtch_SchedulingInfo_r17 != nil, ie.Mtch_NeighbourCell_r17 != nil, ie.Pdsch_ConfigIndex_r17 != nil, ie.Mtch_SSB_MappingWindowIndex_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.Mbs_SessionId_r17.Encode(w); err != nil {
		return utils.WrapError("Encode Mbs_SessionId_r17", err)
	}
	if err = ie.G_RNTI_r17.Encode(w); err != nil {
		return utils.WrapError("Encode G_RNTI_r17", err)
	}
	if err = ie.Mrb_ListBroadcast_r17.Encode(w); err != nil {
		return utils.WrapError("Encode Mrb_ListBroadcast_r17", err)
	}
	if ie.Mtch_SchedulingInfo_r17 != nil {
		if err = ie.Mtch_SchedulingInfo_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Mtch_SchedulingInfo_r17", err)
		}
	}
	if ie.Mtch_NeighbourCell_r17 != nil {
		if err = w.WriteBitString(ie.Mtch_NeighbourCell_r17.Bytes, uint(ie.Mtch_NeighbourCell_r17.NumBits), &uper.Constraint{Lb: maxNeighCellMBS_r17, Ub: maxNeighCellMBS_r17}, false); err != nil {
			return utils.WrapError("Encode Mtch_NeighbourCell_r17", err)
		}
	}
	if ie.Pdsch_ConfigIndex_r17 != nil {
		if err = ie.Pdsch_ConfigIndex_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Pdsch_ConfigIndex_r17", err)
		}
	}
	if ie.Mtch_SSB_MappingWindowIndex_r17 != nil {
		if err = ie.Mtch_SSB_MappingWindowIndex_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Mtch_SSB_MappingWindowIndex_r17", err)
		}
	}
	return nil
}

func (ie *MBS_SessionInfo_r17) Decode(r *uper.UperReader) error {
	var err error
	var Mtch_SchedulingInfo_r17Present bool
	if Mtch_SchedulingInfo_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Mtch_NeighbourCell_r17Present bool
	if Mtch_NeighbourCell_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Pdsch_ConfigIndex_r17Present bool
	if Pdsch_ConfigIndex_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Mtch_SSB_MappingWindowIndex_r17Present bool
	if Mtch_SSB_MappingWindowIndex_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.Mbs_SessionId_r17.Decode(r); err != nil {
		return utils.WrapError("Decode Mbs_SessionId_r17", err)
	}
	if err = ie.G_RNTI_r17.Decode(r); err != nil {
		return utils.WrapError("Decode G_RNTI_r17", err)
	}
	if err = ie.Mrb_ListBroadcast_r17.Decode(r); err != nil {
		return utils.WrapError("Decode Mrb_ListBroadcast_r17", err)
	}
	if Mtch_SchedulingInfo_r17Present {
		ie.Mtch_SchedulingInfo_r17 = new(DRX_ConfigPTM_Index_r17)
		if err = ie.Mtch_SchedulingInfo_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Mtch_SchedulingInfo_r17", err)
		}
	}
	if Mtch_NeighbourCell_r17Present {
		var tmp_bs_Mtch_NeighbourCell_r17 []byte
		var n_Mtch_NeighbourCell_r17 uint
		if tmp_bs_Mtch_NeighbourCell_r17, n_Mtch_NeighbourCell_r17, err = r.ReadBitString(&uper.Constraint{Lb: maxNeighCellMBS_r17, Ub: maxNeighCellMBS_r17}, false); err != nil {
			return utils.WrapError("Decode Mtch_NeighbourCell_r17", err)
		}
		tmp_bitstring := uper.BitString{
			Bytes:   tmp_bs_Mtch_NeighbourCell_r17,
			NumBits: uint64(n_Mtch_NeighbourCell_r17),
		}
		ie.Mtch_NeighbourCell_r17 = &tmp_bitstring
	}
	if Pdsch_ConfigIndex_r17Present {
		ie.Pdsch_ConfigIndex_r17 = new(PDSCH_ConfigIndex_r17)
		if err = ie.Pdsch_ConfigIndex_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Pdsch_ConfigIndex_r17", err)
		}
	}
	if Mtch_SSB_MappingWindowIndex_r17Present {
		ie.Mtch_SSB_MappingWindowIndex_r17 = new(MTCH_SSB_MappingWindowIndex_r17)
		if err = ie.Mtch_SSB_MappingWindowIndex_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Mtch_SSB_MappingWindowIndex_r17", err)
		}
	}
	return nil
}
