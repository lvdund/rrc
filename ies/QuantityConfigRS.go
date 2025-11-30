package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type QuantityConfigRS struct {
	Ssb_FilterConfig    FilterConfig `madatory`
	Csi_RS_FilterConfig FilterConfig `madatory`
}

func (ie *QuantityConfigRS) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.Ssb_FilterConfig.Encode(w); err != nil {
		return utils.WrapError("Encode Ssb_FilterConfig", err)
	}
	if err = ie.Csi_RS_FilterConfig.Encode(w); err != nil {
		return utils.WrapError("Encode Csi_RS_FilterConfig", err)
	}
	return nil
}

func (ie *QuantityConfigRS) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.Ssb_FilterConfig.Decode(r); err != nil {
		return utils.WrapError("Decode Ssb_FilterConfig", err)
	}
	if err = ie.Csi_RS_FilterConfig.Decode(r); err != nil {
		return utils.WrapError("Decode Csi_RS_FilterConfig", err)
	}
	return nil
}
