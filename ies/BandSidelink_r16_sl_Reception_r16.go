package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type BandSidelink_r16_sl_Reception_r16 struct {
	Harq_RxProcessSidelink_r16   BandSidelink_r16_sl_Reception_r16_harq_RxProcessSidelink_r16    `madatory`
	Pscch_RxSidelink_r16         BandSidelink_r16_sl_Reception_r16_pscch_RxSidelink_r16          `madatory`
	Scs_CP_PatternRxSidelink_r16 *BandSidelink_r16_sl_Reception_r16_scs_CP_PatternRxSidelink_r16 `optional`
	ExtendedCP_RxSidelink_r16    *BandSidelink_r16_sl_Reception_r16_extendedCP_RxSidelink_r16    `optional`
}

func (ie *BandSidelink_r16_sl_Reception_r16) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Scs_CP_PatternRxSidelink_r16 != nil, ie.ExtendedCP_RxSidelink_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.Harq_RxProcessSidelink_r16.Encode(w); err != nil {
		return utils.WrapError("Encode Harq_RxProcessSidelink_r16", err)
	}
	if err = ie.Pscch_RxSidelink_r16.Encode(w); err != nil {
		return utils.WrapError("Encode Pscch_RxSidelink_r16", err)
	}
	if ie.Scs_CP_PatternRxSidelink_r16 != nil {
		if err = ie.Scs_CP_PatternRxSidelink_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Scs_CP_PatternRxSidelink_r16", err)
		}
	}
	if ie.ExtendedCP_RxSidelink_r16 != nil {
		if err = ie.ExtendedCP_RxSidelink_r16.Encode(w); err != nil {
			return utils.WrapError("Encode ExtendedCP_RxSidelink_r16", err)
		}
	}
	return nil
}

func (ie *BandSidelink_r16_sl_Reception_r16) Decode(r *aper.AperReader) error {
	var err error
	var Scs_CP_PatternRxSidelink_r16Present bool
	if Scs_CP_PatternRxSidelink_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ExtendedCP_RxSidelink_r16Present bool
	if ExtendedCP_RxSidelink_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.Harq_RxProcessSidelink_r16.Decode(r); err != nil {
		return utils.WrapError("Decode Harq_RxProcessSidelink_r16", err)
	}
	if err = ie.Pscch_RxSidelink_r16.Decode(r); err != nil {
		return utils.WrapError("Decode Pscch_RxSidelink_r16", err)
	}
	if Scs_CP_PatternRxSidelink_r16Present {
		ie.Scs_CP_PatternRxSidelink_r16 = new(BandSidelink_r16_sl_Reception_r16_scs_CP_PatternRxSidelink_r16)
		if err = ie.Scs_CP_PatternRxSidelink_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Scs_CP_PatternRxSidelink_r16", err)
		}
	}
	if ExtendedCP_RxSidelink_r16Present {
		ie.ExtendedCP_RxSidelink_r16 = new(BandSidelink_r16_sl_Reception_r16_extendedCP_RxSidelink_r16)
		if err = ie.ExtendedCP_RxSidelink_r16.Decode(r); err != nil {
			return utils.WrapError("Decode ExtendedCP_RxSidelink_r16", err)
		}
	}
	return nil
}
