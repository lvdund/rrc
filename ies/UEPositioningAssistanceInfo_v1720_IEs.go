package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type UEPositioningAssistanceInfo_v1720_IEs struct {
	Ue_TxTEG_TimingErrorMarginValue_r17 *UEPositioningAssistanceInfo_v1720_IEs_ue_TxTEG_TimingErrorMarginValue_r17 `optional`
	NonCriticalExtension                interface{}                                                                `optional`
}

func (ie *UEPositioningAssistanceInfo_v1720_IEs) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Ue_TxTEG_TimingErrorMarginValue_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Ue_TxTEG_TimingErrorMarginValue_r17 != nil {
		if err = ie.Ue_TxTEG_TimingErrorMarginValue_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Ue_TxTEG_TimingErrorMarginValue_r17", err)
		}
	}
	return nil
}

func (ie *UEPositioningAssistanceInfo_v1720_IEs) Decode(r *aper.AperReader) error {
	var err error
	var Ue_TxTEG_TimingErrorMarginValue_r17Present bool
	if Ue_TxTEG_TimingErrorMarginValue_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Ue_TxTEG_TimingErrorMarginValue_r17Present {
		ie.Ue_TxTEG_TimingErrorMarginValue_r17 = new(UEPositioningAssistanceInfo_v1720_IEs_ue_TxTEG_TimingErrorMarginValue_r17)
		if err = ie.Ue_TxTEG_TimingErrorMarginValue_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Ue_TxTEG_TimingErrorMarginValue_r17", err)
		}
	}
	return nil
}
