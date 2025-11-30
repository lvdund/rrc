package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type BandSidelinkEUTRA_r16 struct {
	FreqBandSidelinkEUTRA_r16           FreqBandIndicatorEUTRA                                     `madatory`
	Gnb_ScheduledMode3SidelinkEUTRA_r16 *BandSidelinkEUTRA_r16_gnb_ScheduledMode3SidelinkEUTRA_r16 `optional`
	Gnb_ScheduledMode4SidelinkEUTRA_r16 *BandSidelinkEUTRA_r16_gnb_ScheduledMode4SidelinkEUTRA_r16 `optional`
}

func (ie *BandSidelinkEUTRA_r16) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Gnb_ScheduledMode3SidelinkEUTRA_r16 != nil, ie.Gnb_ScheduledMode4SidelinkEUTRA_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.FreqBandSidelinkEUTRA_r16.Encode(w); err != nil {
		return utils.WrapError("Encode FreqBandSidelinkEUTRA_r16", err)
	}
	if ie.Gnb_ScheduledMode3SidelinkEUTRA_r16 != nil {
		if err = ie.Gnb_ScheduledMode3SidelinkEUTRA_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Gnb_ScheduledMode3SidelinkEUTRA_r16", err)
		}
	}
	if ie.Gnb_ScheduledMode4SidelinkEUTRA_r16 != nil {
		if err = ie.Gnb_ScheduledMode4SidelinkEUTRA_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Gnb_ScheduledMode4SidelinkEUTRA_r16", err)
		}
	}
	return nil
}

func (ie *BandSidelinkEUTRA_r16) Decode(r *aper.AperReader) error {
	var err error
	var Gnb_ScheduledMode3SidelinkEUTRA_r16Present bool
	if Gnb_ScheduledMode3SidelinkEUTRA_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Gnb_ScheduledMode4SidelinkEUTRA_r16Present bool
	if Gnb_ScheduledMode4SidelinkEUTRA_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.FreqBandSidelinkEUTRA_r16.Decode(r); err != nil {
		return utils.WrapError("Decode FreqBandSidelinkEUTRA_r16", err)
	}
	if Gnb_ScheduledMode3SidelinkEUTRA_r16Present {
		ie.Gnb_ScheduledMode3SidelinkEUTRA_r16 = new(BandSidelinkEUTRA_r16_gnb_ScheduledMode3SidelinkEUTRA_r16)
		if err = ie.Gnb_ScheduledMode3SidelinkEUTRA_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Gnb_ScheduledMode3SidelinkEUTRA_r16", err)
		}
	}
	if Gnb_ScheduledMode4SidelinkEUTRA_r16Present {
		ie.Gnb_ScheduledMode4SidelinkEUTRA_r16 = new(BandSidelinkEUTRA_r16_gnb_ScheduledMode4SidelinkEUTRA_r16)
		if err = ie.Gnb_ScheduledMode4SidelinkEUTRA_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Gnb_ScheduledMode4SidelinkEUTRA_r16", err)
		}
	}
	return nil
}
