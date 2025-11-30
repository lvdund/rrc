package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type PDCP_Config_drb_headerCompression_rohc_profiles struct {
	Profile0x0001 bool `madatory`
	Profile0x0002 bool `madatory`
	Profile0x0003 bool `madatory`
	Profile0x0004 bool `madatory`
	Profile0x0006 bool `madatory`
	Profile0x0101 bool `madatory`
	Profile0x0102 bool `madatory`
	Profile0x0103 bool `madatory`
	Profile0x0104 bool `madatory`
}

func (ie *PDCP_Config_drb_headerCompression_rohc_profiles) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteBoolean(ie.Profile0x0001); err != nil {
		return utils.WrapError("WriteBoolean Profile0x0001", err)
	}
	if err = w.WriteBoolean(ie.Profile0x0002); err != nil {
		return utils.WrapError("WriteBoolean Profile0x0002", err)
	}
	if err = w.WriteBoolean(ie.Profile0x0003); err != nil {
		return utils.WrapError("WriteBoolean Profile0x0003", err)
	}
	if err = w.WriteBoolean(ie.Profile0x0004); err != nil {
		return utils.WrapError("WriteBoolean Profile0x0004", err)
	}
	if err = w.WriteBoolean(ie.Profile0x0006); err != nil {
		return utils.WrapError("WriteBoolean Profile0x0006", err)
	}
	if err = w.WriteBoolean(ie.Profile0x0101); err != nil {
		return utils.WrapError("WriteBoolean Profile0x0101", err)
	}
	if err = w.WriteBoolean(ie.Profile0x0102); err != nil {
		return utils.WrapError("WriteBoolean Profile0x0102", err)
	}
	if err = w.WriteBoolean(ie.Profile0x0103); err != nil {
		return utils.WrapError("WriteBoolean Profile0x0103", err)
	}
	if err = w.WriteBoolean(ie.Profile0x0104); err != nil {
		return utils.WrapError("WriteBoolean Profile0x0104", err)
	}
	return nil
}

func (ie *PDCP_Config_drb_headerCompression_rohc_profiles) Decode(r *aper.AperReader) error {
	var err error
	var tmp_bool_Profile0x0001 bool
	if tmp_bool_Profile0x0001, err = r.ReadBoolean(); err != nil {
		return utils.WrapError("ReadBoolean Profile0x0001", err)
	}
	ie.Profile0x0001 = tmp_bool_Profile0x0001
	var tmp_bool_Profile0x0002 bool
	if tmp_bool_Profile0x0002, err = r.ReadBoolean(); err != nil {
		return utils.WrapError("ReadBoolean Profile0x0002", err)
	}
	ie.Profile0x0002 = tmp_bool_Profile0x0002
	var tmp_bool_Profile0x0003 bool
	if tmp_bool_Profile0x0003, err = r.ReadBoolean(); err != nil {
		return utils.WrapError("ReadBoolean Profile0x0003", err)
	}
	ie.Profile0x0003 = tmp_bool_Profile0x0003
	var tmp_bool_Profile0x0004 bool
	if tmp_bool_Profile0x0004, err = r.ReadBoolean(); err != nil {
		return utils.WrapError("ReadBoolean Profile0x0004", err)
	}
	ie.Profile0x0004 = tmp_bool_Profile0x0004
	var tmp_bool_Profile0x0006 bool
	if tmp_bool_Profile0x0006, err = r.ReadBoolean(); err != nil {
		return utils.WrapError("ReadBoolean Profile0x0006", err)
	}
	ie.Profile0x0006 = tmp_bool_Profile0x0006
	var tmp_bool_Profile0x0101 bool
	if tmp_bool_Profile0x0101, err = r.ReadBoolean(); err != nil {
		return utils.WrapError("ReadBoolean Profile0x0101", err)
	}
	ie.Profile0x0101 = tmp_bool_Profile0x0101
	var tmp_bool_Profile0x0102 bool
	if tmp_bool_Profile0x0102, err = r.ReadBoolean(); err != nil {
		return utils.WrapError("ReadBoolean Profile0x0102", err)
	}
	ie.Profile0x0102 = tmp_bool_Profile0x0102
	var tmp_bool_Profile0x0103 bool
	if tmp_bool_Profile0x0103, err = r.ReadBoolean(); err != nil {
		return utils.WrapError("ReadBoolean Profile0x0103", err)
	}
	ie.Profile0x0103 = tmp_bool_Profile0x0103
	var tmp_bool_Profile0x0104 bool
	if tmp_bool_Profile0x0104, err = r.ReadBoolean(); err != nil {
		return utils.WrapError("ReadBoolean Profile0x0104", err)
	}
	ie.Profile0x0104 = tmp_bool_Profile0x0104
	return nil
}
