package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type BandSidelink_r16_sl_TransmissionMode1_r16 struct {
	Harq_TxProcessModeOneSidelink_r16   BandSidelink_r16_sl_TransmissionMode1_r16_harq_TxProcessModeOneSidelink_r16   `madatory`
	Scs_CP_PatternTxSidelinkModeOne_r16 BandSidelink_r16_sl_TransmissionMode1_r16_scs_CP_PatternTxSidelinkModeOne_r16 `madatory`
	ExtendedCP_TxSidelink_r16           *BandSidelink_r16_sl_TransmissionMode1_r16_extendedCP_TxSidelink_r16          `optional`
	Harq_ReportOnPUCCH_r16              *BandSidelink_r16_sl_TransmissionMode1_r16_harq_ReportOnPUCCH_r16             `optional`
}

func (ie *BandSidelink_r16_sl_TransmissionMode1_r16) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.ExtendedCP_TxSidelink_r16 != nil, ie.Harq_ReportOnPUCCH_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.Harq_TxProcessModeOneSidelink_r16.Encode(w); err != nil {
		return utils.WrapError("Encode Harq_TxProcessModeOneSidelink_r16", err)
	}
	if err = ie.Scs_CP_PatternTxSidelinkModeOne_r16.Encode(w); err != nil {
		return utils.WrapError("Encode Scs_CP_PatternTxSidelinkModeOne_r16", err)
	}
	if ie.ExtendedCP_TxSidelink_r16 != nil {
		if err = ie.ExtendedCP_TxSidelink_r16.Encode(w); err != nil {
			return utils.WrapError("Encode ExtendedCP_TxSidelink_r16", err)
		}
	}
	if ie.Harq_ReportOnPUCCH_r16 != nil {
		if err = ie.Harq_ReportOnPUCCH_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Harq_ReportOnPUCCH_r16", err)
		}
	}
	return nil
}

func (ie *BandSidelink_r16_sl_TransmissionMode1_r16) Decode(r *aper.AperReader) error {
	var err error
	var ExtendedCP_TxSidelink_r16Present bool
	if ExtendedCP_TxSidelink_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Harq_ReportOnPUCCH_r16Present bool
	if Harq_ReportOnPUCCH_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.Harq_TxProcessModeOneSidelink_r16.Decode(r); err != nil {
		return utils.WrapError("Decode Harq_TxProcessModeOneSidelink_r16", err)
	}
	if err = ie.Scs_CP_PatternTxSidelinkModeOne_r16.Decode(r); err != nil {
		return utils.WrapError("Decode Scs_CP_PatternTxSidelinkModeOne_r16", err)
	}
	if ExtendedCP_TxSidelink_r16Present {
		ie.ExtendedCP_TxSidelink_r16 = new(BandSidelink_r16_sl_TransmissionMode1_r16_extendedCP_TxSidelink_r16)
		if err = ie.ExtendedCP_TxSidelink_r16.Decode(r); err != nil {
			return utils.WrapError("Decode ExtendedCP_TxSidelink_r16", err)
		}
	}
	if Harq_ReportOnPUCCH_r16Present {
		ie.Harq_ReportOnPUCCH_r16 = new(BandSidelink_r16_sl_TransmissionMode1_r16_harq_ReportOnPUCCH_r16)
		if err = ie.Harq_ReportOnPUCCH_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Harq_ReportOnPUCCH_r16", err)
		}
	}
	return nil
}
