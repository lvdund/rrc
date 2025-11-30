package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	BandSidelinkEUTRA_r16_gnb_ScheduledMode3SidelinkEUTRA_r16_gnb_ScheduledMode3DelaySidelinkEUTRA_r16_Enum_ms0       aper.Enumerated = 0
	BandSidelinkEUTRA_r16_gnb_ScheduledMode3SidelinkEUTRA_r16_gnb_ScheduledMode3DelaySidelinkEUTRA_r16_Enum_ms0dot25  aper.Enumerated = 1
	BandSidelinkEUTRA_r16_gnb_ScheduledMode3SidelinkEUTRA_r16_gnb_ScheduledMode3DelaySidelinkEUTRA_r16_Enum_ms0dot5   aper.Enumerated = 2
	BandSidelinkEUTRA_r16_gnb_ScheduledMode3SidelinkEUTRA_r16_gnb_ScheduledMode3DelaySidelinkEUTRA_r16_Enum_ms0dot625 aper.Enumerated = 3
	BandSidelinkEUTRA_r16_gnb_ScheduledMode3SidelinkEUTRA_r16_gnb_ScheduledMode3DelaySidelinkEUTRA_r16_Enum_ms0dot75  aper.Enumerated = 4
	BandSidelinkEUTRA_r16_gnb_ScheduledMode3SidelinkEUTRA_r16_gnb_ScheduledMode3DelaySidelinkEUTRA_r16_Enum_ms1       aper.Enumerated = 5
	BandSidelinkEUTRA_r16_gnb_ScheduledMode3SidelinkEUTRA_r16_gnb_ScheduledMode3DelaySidelinkEUTRA_r16_Enum_ms1dot25  aper.Enumerated = 6
	BandSidelinkEUTRA_r16_gnb_ScheduledMode3SidelinkEUTRA_r16_gnb_ScheduledMode3DelaySidelinkEUTRA_r16_Enum_ms1dot5   aper.Enumerated = 7
	BandSidelinkEUTRA_r16_gnb_ScheduledMode3SidelinkEUTRA_r16_gnb_ScheduledMode3DelaySidelinkEUTRA_r16_Enum_ms1dot75  aper.Enumerated = 8
	BandSidelinkEUTRA_r16_gnb_ScheduledMode3SidelinkEUTRA_r16_gnb_ScheduledMode3DelaySidelinkEUTRA_r16_Enum_ms2       aper.Enumerated = 9
	BandSidelinkEUTRA_r16_gnb_ScheduledMode3SidelinkEUTRA_r16_gnb_ScheduledMode3DelaySidelinkEUTRA_r16_Enum_ms2dot5   aper.Enumerated = 10
	BandSidelinkEUTRA_r16_gnb_ScheduledMode3SidelinkEUTRA_r16_gnb_ScheduledMode3DelaySidelinkEUTRA_r16_Enum_ms3       aper.Enumerated = 11
	BandSidelinkEUTRA_r16_gnb_ScheduledMode3SidelinkEUTRA_r16_gnb_ScheduledMode3DelaySidelinkEUTRA_r16_Enum_ms4       aper.Enumerated = 12
	BandSidelinkEUTRA_r16_gnb_ScheduledMode3SidelinkEUTRA_r16_gnb_ScheduledMode3DelaySidelinkEUTRA_r16_Enum_ms5       aper.Enumerated = 13
	BandSidelinkEUTRA_r16_gnb_ScheduledMode3SidelinkEUTRA_r16_gnb_ScheduledMode3DelaySidelinkEUTRA_r16_Enum_ms6       aper.Enumerated = 14
	BandSidelinkEUTRA_r16_gnb_ScheduledMode3SidelinkEUTRA_r16_gnb_ScheduledMode3DelaySidelinkEUTRA_r16_Enum_ms8       aper.Enumerated = 15
	BandSidelinkEUTRA_r16_gnb_ScheduledMode3SidelinkEUTRA_r16_gnb_ScheduledMode3DelaySidelinkEUTRA_r16_Enum_ms10      aper.Enumerated = 16
	BandSidelinkEUTRA_r16_gnb_ScheduledMode3SidelinkEUTRA_r16_gnb_ScheduledMode3DelaySidelinkEUTRA_r16_Enum_ms20      aper.Enumerated = 17
)

type BandSidelinkEUTRA_r16_gnb_ScheduledMode3SidelinkEUTRA_r16_gnb_ScheduledMode3DelaySidelinkEUTRA_r16 struct {
	Value aper.Enumerated `lb:0,ub:17,madatory`
}

func (ie *BandSidelinkEUTRA_r16_gnb_ScheduledMode3SidelinkEUTRA_r16_gnb_ScheduledMode3DelaySidelinkEUTRA_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 17}, false); err != nil {
		return utils.WrapError("Encode BandSidelinkEUTRA_r16_gnb_ScheduledMode3SidelinkEUTRA_r16_gnb_ScheduledMode3DelaySidelinkEUTRA_r16", err)
	}
	return nil
}

func (ie *BandSidelinkEUTRA_r16_gnb_ScheduledMode3SidelinkEUTRA_r16_gnb_ScheduledMode3DelaySidelinkEUTRA_r16) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 17}, false); err != nil {
		return utils.WrapError("Decode BandSidelinkEUTRA_r16_gnb_ScheduledMode3SidelinkEUTRA_r16_gnb_ScheduledMode3DelaySidelinkEUTRA_r16", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
