package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type BandSidelink_r16_congestionControlSidelink_r16 struct {
	Cbr_ReportSidelink_r16       *BandSidelink_r16_congestionControlSidelink_r16_cbr_ReportSidelink_r16      `optional`
	Cbr_CR_TimeLimitSidelink_r16 BandSidelink_r16_congestionControlSidelink_r16_cbr_CR_TimeLimitSidelink_r16 `madatory`
}

func (ie *BandSidelink_r16_congestionControlSidelink_r16) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Cbr_ReportSidelink_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Cbr_ReportSidelink_r16 != nil {
		if err = ie.Cbr_ReportSidelink_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Cbr_ReportSidelink_r16", err)
		}
	}
	if err = ie.Cbr_CR_TimeLimitSidelink_r16.Encode(w); err != nil {
		return utils.WrapError("Encode Cbr_CR_TimeLimitSidelink_r16", err)
	}
	return nil
}

func (ie *BandSidelink_r16_congestionControlSidelink_r16) Decode(r *aper.AperReader) error {
	var err error
	var Cbr_ReportSidelink_r16Present bool
	if Cbr_ReportSidelink_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Cbr_ReportSidelink_r16Present {
		ie.Cbr_ReportSidelink_r16 = new(BandSidelink_r16_congestionControlSidelink_r16_cbr_ReportSidelink_r16)
		if err = ie.Cbr_ReportSidelink_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Cbr_ReportSidelink_r16", err)
		}
	}
	if err = ie.Cbr_CR_TimeLimitSidelink_r16.Decode(r); err != nil {
		return utils.WrapError("Decode Cbr_CR_TimeLimitSidelink_r16", err)
	}
	return nil
}
