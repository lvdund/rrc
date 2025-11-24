package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type BandSidelink_r16_sl_TransmissionMode2_r16 struct {
	Harq_TxProcessModeTwoSidelink_r16   BandSidelink_r16_sl_TransmissionMode2_r16_harq_TxProcessModeTwoSidelink_r16    `madatory`
	Scs_CP_PatternTxSidelinkModeTwo_r16 *BandSidelink_r16_sl_TransmissionMode2_r16_scs_CP_PatternTxSidelinkModeTwo_r16 `optional`
	Dl_openLoopPC_Sidelink_r16          *BandSidelink_r16_sl_TransmissionMode2_r16_dl_openLoopPC_Sidelink_r16          `optional`
}

func (ie *BandSidelink_r16_sl_TransmissionMode2_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.Scs_CP_PatternTxSidelinkModeTwo_r16 != nil, ie.Dl_openLoopPC_Sidelink_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.Harq_TxProcessModeTwoSidelink_r16.Encode(w); err != nil {
		return utils.WrapError("Encode Harq_TxProcessModeTwoSidelink_r16", err)
	}
	if ie.Scs_CP_PatternTxSidelinkModeTwo_r16 != nil {
		if err = ie.Scs_CP_PatternTxSidelinkModeTwo_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Scs_CP_PatternTxSidelinkModeTwo_r16", err)
		}
	}
	if ie.Dl_openLoopPC_Sidelink_r16 != nil {
		if err = ie.Dl_openLoopPC_Sidelink_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Dl_openLoopPC_Sidelink_r16", err)
		}
	}
	return nil
}

func (ie *BandSidelink_r16_sl_TransmissionMode2_r16) Decode(r *uper.UperReader) error {
	var err error
	var Scs_CP_PatternTxSidelinkModeTwo_r16Present bool
	if Scs_CP_PatternTxSidelinkModeTwo_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Dl_openLoopPC_Sidelink_r16Present bool
	if Dl_openLoopPC_Sidelink_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.Harq_TxProcessModeTwoSidelink_r16.Decode(r); err != nil {
		return utils.WrapError("Decode Harq_TxProcessModeTwoSidelink_r16", err)
	}
	if Scs_CP_PatternTxSidelinkModeTwo_r16Present {
		ie.Scs_CP_PatternTxSidelinkModeTwo_r16 = new(BandSidelink_r16_sl_TransmissionMode2_r16_scs_CP_PatternTxSidelinkModeTwo_r16)
		if err = ie.Scs_CP_PatternTxSidelinkModeTwo_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Scs_CP_PatternTxSidelinkModeTwo_r16", err)
		}
	}
	if Dl_openLoopPC_Sidelink_r16Present {
		ie.Dl_openLoopPC_Sidelink_r16 = new(BandSidelink_r16_sl_TransmissionMode2_r16_dl_openLoopPC_Sidelink_r16)
		if err = ie.Dl_openLoopPC_Sidelink_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Dl_openLoopPC_Sidelink_r16", err)
		}
	}
	return nil
}
