package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	GNSS_ID_r16_gnss_id_r16_Enum_gps     aper.Enumerated = 0
	GNSS_ID_r16_gnss_id_r16_Enum_sbas    aper.Enumerated = 1
	GNSS_ID_r16_gnss_id_r16_Enum_qzss    aper.Enumerated = 2
	GNSS_ID_r16_gnss_id_r16_Enum_galileo aper.Enumerated = 3
	GNSS_ID_r16_gnss_id_r16_Enum_glonass aper.Enumerated = 4
	GNSS_ID_r16_gnss_id_r16_Enum_bds     aper.Enumerated = 5
)

type GNSS_ID_r16_gnss_id_r16 struct {
	Value aper.Enumerated `lb:0,ub:5,madatory`
}

func (ie *GNSS_ID_r16_gnss_id_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 5}, false); err != nil {
		return utils.WrapError("Encode GNSS_ID_r16_gnss_id_r16", err)
	}
	return nil
}

func (ie *GNSS_ID_r16_gnss_id_r16) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 5}, false); err != nil {
		return utils.WrapError("Decode GNSS_ID_r16_gnss_id_r16", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
