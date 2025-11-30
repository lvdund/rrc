package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type GNSS_ID_r16 struct {
	Gnss_id_r16 GNSS_ID_r16_gnss_id_r16 `madatory`
}

func (ie *GNSS_ID_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.Gnss_id_r16.Encode(w); err != nil {
		return utils.WrapError("Encode Gnss_id_r16", err)
	}
	return nil
}

func (ie *GNSS_ID_r16) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.Gnss_id_r16.Decode(r); err != nil {
		return utils.WrapError("Decode Gnss_id_r16", err)
	}
	return nil
}
