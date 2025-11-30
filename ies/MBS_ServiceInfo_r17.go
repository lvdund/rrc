package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type MBS_ServiceInfo_r17 struct {
	Tmgi_r17 TMGI_r17 `madatory`
}

func (ie *MBS_ServiceInfo_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.Tmgi_r17.Encode(w); err != nil {
		return utils.WrapError("Encode Tmgi_r17", err)
	}
	return nil
}

func (ie *MBS_ServiceInfo_r17) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.Tmgi_r17.Decode(r); err != nil {
		return utils.WrapError("Decode Tmgi_r17", err)
	}
	return nil
}
