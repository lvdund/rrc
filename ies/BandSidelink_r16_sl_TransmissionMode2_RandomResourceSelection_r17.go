package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type BandSidelink_r16_sl_TransmissionMode2_RandomResourceSelection_r17 struct {
	Harq_TxProcessModeTwoSidelink_r17   BandSidelink_r16_sl_TransmissionMode2_RandomResourceSelection_r17_harq_TxProcessModeTwoSidelink_r17    `madatory`
	Scs_CP_PatternTxSidelinkModeTwo_r17 *BandSidelink_r16_sl_TransmissionMode2_RandomResourceSelection_r17_scs_CP_PatternTxSidelinkModeTwo_r17 `optional`
	ExtendedCP_Mode2Random_r17          *BandSidelink_r16_sl_TransmissionMode2_RandomResourceSelection_r17_extendedCP_Mode2Random_r17          `optional`
	Dl_openLoopPC_Sidelink_r17          *BandSidelink_r16_sl_TransmissionMode2_RandomResourceSelection_r17_dl_openLoopPC_Sidelink_r17          `optional`
}

func (ie *BandSidelink_r16_sl_TransmissionMode2_RandomResourceSelection_r17) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Scs_CP_PatternTxSidelinkModeTwo_r17 != nil, ie.ExtendedCP_Mode2Random_r17 != nil, ie.Dl_openLoopPC_Sidelink_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.Harq_TxProcessModeTwoSidelink_r17.Encode(w); err != nil {
		return utils.WrapError("Encode Harq_TxProcessModeTwoSidelink_r17", err)
	}
	if ie.Scs_CP_PatternTxSidelinkModeTwo_r17 != nil {
		if err = ie.Scs_CP_PatternTxSidelinkModeTwo_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Scs_CP_PatternTxSidelinkModeTwo_r17", err)
		}
	}
	if ie.ExtendedCP_Mode2Random_r17 != nil {
		if err = ie.ExtendedCP_Mode2Random_r17.Encode(w); err != nil {
			return utils.WrapError("Encode ExtendedCP_Mode2Random_r17", err)
		}
	}
	if ie.Dl_openLoopPC_Sidelink_r17 != nil {
		if err = ie.Dl_openLoopPC_Sidelink_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Dl_openLoopPC_Sidelink_r17", err)
		}
	}
	return nil
}

func (ie *BandSidelink_r16_sl_TransmissionMode2_RandomResourceSelection_r17) Decode(r *aper.AperReader) error {
	var err error
	var Scs_CP_PatternTxSidelinkModeTwo_r17Present bool
	if Scs_CP_PatternTxSidelinkModeTwo_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ExtendedCP_Mode2Random_r17Present bool
	if ExtendedCP_Mode2Random_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Dl_openLoopPC_Sidelink_r17Present bool
	if Dl_openLoopPC_Sidelink_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.Harq_TxProcessModeTwoSidelink_r17.Decode(r); err != nil {
		return utils.WrapError("Decode Harq_TxProcessModeTwoSidelink_r17", err)
	}
	if Scs_CP_PatternTxSidelinkModeTwo_r17Present {
		ie.Scs_CP_PatternTxSidelinkModeTwo_r17 = new(BandSidelink_r16_sl_TransmissionMode2_RandomResourceSelection_r17_scs_CP_PatternTxSidelinkModeTwo_r17)
		if err = ie.Scs_CP_PatternTxSidelinkModeTwo_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Scs_CP_PatternTxSidelinkModeTwo_r17", err)
		}
	}
	if ExtendedCP_Mode2Random_r17Present {
		ie.ExtendedCP_Mode2Random_r17 = new(BandSidelink_r16_sl_TransmissionMode2_RandomResourceSelection_r17_extendedCP_Mode2Random_r17)
		if err = ie.ExtendedCP_Mode2Random_r17.Decode(r); err != nil {
			return utils.WrapError("Decode ExtendedCP_Mode2Random_r17", err)
		}
	}
	if Dl_openLoopPC_Sidelink_r17Present {
		ie.Dl_openLoopPC_Sidelink_r17 = new(BandSidelink_r16_sl_TransmissionMode2_RandomResourceSelection_r17_dl_openLoopPC_Sidelink_r17)
		if err = ie.Dl_openLoopPC_Sidelink_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Dl_openLoopPC_Sidelink_r17", err)
		}
	}
	return nil
}
