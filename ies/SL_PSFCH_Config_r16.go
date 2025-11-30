package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type SL_PSFCH_Config_r16 struct {
	Sl_PSFCH_Period_r16                *SL_PSFCH_Config_r16_sl_PSFCH_Period_r16                `optional`
	Sl_PSFCH_RB_Set_r16                *aper.BitString                                         `lb:10,ub:275,optional`
	Sl_NumMuxCS_Pair_r16               *SL_PSFCH_Config_r16_sl_NumMuxCS_Pair_r16               `optional`
	Sl_MinTimeGapPSFCH_r16             *SL_PSFCH_Config_r16_sl_MinTimeGapPSFCH_r16             `optional`
	Sl_PSFCH_HopID_r16                 *int64                                                  `lb:0,ub:1023,optional`
	Sl_PSFCH_CandidateResourceType_r16 *SL_PSFCH_Config_r16_sl_PSFCH_CandidateResourceType_r16 `optional`
}

func (ie *SL_PSFCH_Config_r16) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Sl_PSFCH_Period_r16 != nil, ie.Sl_PSFCH_RB_Set_r16 != nil, ie.Sl_NumMuxCS_Pair_r16 != nil, ie.Sl_MinTimeGapPSFCH_r16 != nil, ie.Sl_PSFCH_HopID_r16 != nil, ie.Sl_PSFCH_CandidateResourceType_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Sl_PSFCH_Period_r16 != nil {
		if err = ie.Sl_PSFCH_Period_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_PSFCH_Period_r16", err)
		}
	}
	if ie.Sl_PSFCH_RB_Set_r16 != nil {
		if err = w.WriteBitString(ie.Sl_PSFCH_RB_Set_r16.Bytes, uint(ie.Sl_PSFCH_RB_Set_r16.NumBits), &aper.Constraint{Lb: 10, Ub: 275}, false); err != nil {
			return utils.WrapError("Encode Sl_PSFCH_RB_Set_r16", err)
		}
	}
	if ie.Sl_NumMuxCS_Pair_r16 != nil {
		if err = ie.Sl_NumMuxCS_Pair_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_NumMuxCS_Pair_r16", err)
		}
	}
	if ie.Sl_MinTimeGapPSFCH_r16 != nil {
		if err = ie.Sl_MinTimeGapPSFCH_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_MinTimeGapPSFCH_r16", err)
		}
	}
	if ie.Sl_PSFCH_HopID_r16 != nil {
		if err = w.WriteInteger(*ie.Sl_PSFCH_HopID_r16, &aper.Constraint{Lb: 0, Ub: 1023}, false); err != nil {
			return utils.WrapError("Encode Sl_PSFCH_HopID_r16", err)
		}
	}
	if ie.Sl_PSFCH_CandidateResourceType_r16 != nil {
		if err = ie.Sl_PSFCH_CandidateResourceType_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_PSFCH_CandidateResourceType_r16", err)
		}
	}
	return nil
}

func (ie *SL_PSFCH_Config_r16) Decode(r *aper.AperReader) error {
	var err error
	var Sl_PSFCH_Period_r16Present bool
	if Sl_PSFCH_Period_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_PSFCH_RB_Set_r16Present bool
	if Sl_PSFCH_RB_Set_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_NumMuxCS_Pair_r16Present bool
	if Sl_NumMuxCS_Pair_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_MinTimeGapPSFCH_r16Present bool
	if Sl_MinTimeGapPSFCH_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_PSFCH_HopID_r16Present bool
	if Sl_PSFCH_HopID_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_PSFCH_CandidateResourceType_r16Present bool
	if Sl_PSFCH_CandidateResourceType_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Sl_PSFCH_Period_r16Present {
		ie.Sl_PSFCH_Period_r16 = new(SL_PSFCH_Config_r16_sl_PSFCH_Period_r16)
		if err = ie.Sl_PSFCH_Period_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_PSFCH_Period_r16", err)
		}
	}
	if Sl_PSFCH_RB_Set_r16Present {
		var tmp_bs_Sl_PSFCH_RB_Set_r16 []byte
		var n_Sl_PSFCH_RB_Set_r16 uint
		if tmp_bs_Sl_PSFCH_RB_Set_r16, n_Sl_PSFCH_RB_Set_r16, err = r.ReadBitString(&aper.Constraint{Lb: 10, Ub: 275}, false); err != nil {
			return utils.WrapError("Decode Sl_PSFCH_RB_Set_r16", err)
		}
		tmp_bitstring := aper.BitString{
			Bytes:   tmp_bs_Sl_PSFCH_RB_Set_r16,
			NumBits: uint64(n_Sl_PSFCH_RB_Set_r16),
		}
		ie.Sl_PSFCH_RB_Set_r16 = &tmp_bitstring
	}
	if Sl_NumMuxCS_Pair_r16Present {
		ie.Sl_NumMuxCS_Pair_r16 = new(SL_PSFCH_Config_r16_sl_NumMuxCS_Pair_r16)
		if err = ie.Sl_NumMuxCS_Pair_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_NumMuxCS_Pair_r16", err)
		}
	}
	if Sl_MinTimeGapPSFCH_r16Present {
		ie.Sl_MinTimeGapPSFCH_r16 = new(SL_PSFCH_Config_r16_sl_MinTimeGapPSFCH_r16)
		if err = ie.Sl_MinTimeGapPSFCH_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_MinTimeGapPSFCH_r16", err)
		}
	}
	if Sl_PSFCH_HopID_r16Present {
		var tmp_int_Sl_PSFCH_HopID_r16 int64
		if tmp_int_Sl_PSFCH_HopID_r16, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 1023}, false); err != nil {
			return utils.WrapError("Decode Sl_PSFCH_HopID_r16", err)
		}
		ie.Sl_PSFCH_HopID_r16 = &tmp_int_Sl_PSFCH_HopID_r16
	}
	if Sl_PSFCH_CandidateResourceType_r16Present {
		ie.Sl_PSFCH_CandidateResourceType_r16 = new(SL_PSFCH_Config_r16_sl_PSFCH_CandidateResourceType_r16)
		if err = ie.Sl_PSFCH_CandidateResourceType_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_PSFCH_CandidateResourceType_r16", err)
		}
	}
	return nil
}
