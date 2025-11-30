package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type MRB_PDCP_ConfigBroadcast_r17_headerCompression_r17_rohc_profiles_r17 struct {
	Profile0x0000_r17 bool `madatory`
	Profile0x0001_r17 bool `madatory`
	Profile0x0002_r17 bool `madatory`
}

func (ie *MRB_PDCP_ConfigBroadcast_r17_headerCompression_r17_rohc_profiles_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteBoolean(ie.Profile0x0000_r17); err != nil {
		return utils.WrapError("WriteBoolean Profile0x0000_r17", err)
	}
	if err = w.WriteBoolean(ie.Profile0x0001_r17); err != nil {
		return utils.WrapError("WriteBoolean Profile0x0001_r17", err)
	}
	if err = w.WriteBoolean(ie.Profile0x0002_r17); err != nil {
		return utils.WrapError("WriteBoolean Profile0x0002_r17", err)
	}
	return nil
}

func (ie *MRB_PDCP_ConfigBroadcast_r17_headerCompression_r17_rohc_profiles_r17) Decode(r *aper.AperReader) error {
	var err error
	var tmp_bool_Profile0x0000_r17 bool
	if tmp_bool_Profile0x0000_r17, err = r.ReadBoolean(); err != nil {
		return utils.WrapError("ReadBoolean Profile0x0000_r17", err)
	}
	ie.Profile0x0000_r17 = tmp_bool_Profile0x0000_r17
	var tmp_bool_Profile0x0001_r17 bool
	if tmp_bool_Profile0x0001_r17, err = r.ReadBoolean(); err != nil {
		return utils.WrapError("ReadBoolean Profile0x0001_r17", err)
	}
	ie.Profile0x0001_r17 = tmp_bool_Profile0x0001_r17
	var tmp_bool_Profile0x0002_r17 bool
	if tmp_bool_Profile0x0002_r17, err = r.ReadBoolean(); err != nil {
		return utils.WrapError("ReadBoolean Profile0x0002_r17", err)
	}
	ie.Profile0x0002_r17 = tmp_bool_Profile0x0002_r17
	return nil
}
