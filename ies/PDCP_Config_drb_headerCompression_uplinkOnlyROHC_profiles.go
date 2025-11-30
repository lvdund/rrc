package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type PDCP_Config_drb_headerCompression_uplinkOnlyROHC_profiles struct {
	Profile0x0006 bool `madatory`
}

func (ie *PDCP_Config_drb_headerCompression_uplinkOnlyROHC_profiles) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteBoolean(ie.Profile0x0006); err != nil {
		return utils.WrapError("WriteBoolean Profile0x0006", err)
	}
	return nil
}

func (ie *PDCP_Config_drb_headerCompression_uplinkOnlyROHC_profiles) Decode(r *aper.AperReader) error {
	var err error
	var tmp_bool_Profile0x0006 bool
	if tmp_bool_Profile0x0006, err = r.ReadBoolean(); err != nil {
		return utils.WrapError("ReadBoolean Profile0x0006", err)
	}
	ie.Profile0x0006 = tmp_bool_Profile0x0006
	return nil
}
