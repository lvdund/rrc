package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type SBAS_ID_r16 struct {
	Sbas_id_r16 SBAS_ID_r16_sbas_id_r16 `madatory`
}

func (ie *SBAS_ID_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.Sbas_id_r16.Encode(w); err != nil {
		return utils.WrapError("Encode Sbas_id_r16", err)
	}
	return nil
}

func (ie *SBAS_ID_r16) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.Sbas_id_r16.Decode(r); err != nil {
		return utils.WrapError("Decode Sbas_id_r16", err)
	}
	return nil
}
