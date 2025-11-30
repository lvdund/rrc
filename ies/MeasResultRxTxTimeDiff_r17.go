package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type MeasResultRxTxTimeDiff_r17 struct {
	RxTxTimeDiff_ue_r17 *RxTxTimeDiff_r17 `optional`
}

func (ie *MeasResultRxTxTimeDiff_r17) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.RxTxTimeDiff_ue_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.RxTxTimeDiff_ue_r17 != nil {
		if err = ie.RxTxTimeDiff_ue_r17.Encode(w); err != nil {
			return utils.WrapError("Encode RxTxTimeDiff_ue_r17", err)
		}
	}
	return nil
}

func (ie *MeasResultRxTxTimeDiff_r17) Decode(r *aper.AperReader) error {
	var err error
	var RxTxTimeDiff_ue_r17Present bool
	if RxTxTimeDiff_ue_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if RxTxTimeDiff_ue_r17Present {
		ie.RxTxTimeDiff_ue_r17 = new(RxTxTimeDiff_r17)
		if err = ie.RxTxTimeDiff_ue_r17.Decode(r); err != nil {
			return utils.WrapError("Decode RxTxTimeDiff_ue_r17", err)
		}
	}
	return nil
}
