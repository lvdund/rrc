package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type UplinkConfigCommonSIB_v1700 struct {
	InitialUplinkBWP_RedCap_r17 *BWP_UplinkCommon `optional`
}

func (ie *UplinkConfigCommonSIB_v1700) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.InitialUplinkBWP_RedCap_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.InitialUplinkBWP_RedCap_r17 != nil {
		if err = ie.InitialUplinkBWP_RedCap_r17.Encode(w); err != nil {
			return utils.WrapError("Encode InitialUplinkBWP_RedCap_r17", err)
		}
	}
	return nil
}

func (ie *UplinkConfigCommonSIB_v1700) Decode(r *aper.AperReader) error {
	var err error
	var InitialUplinkBWP_RedCap_r17Present bool
	if InitialUplinkBWP_RedCap_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if InitialUplinkBWP_RedCap_r17Present {
		ie.InitialUplinkBWP_RedCap_r17 = new(BWP_UplinkCommon)
		if err = ie.InitialUplinkBWP_RedCap_r17.Decode(r); err != nil {
			return utils.WrapError("Decode InitialUplinkBWP_RedCap_r17", err)
		}
	}
	return nil
}
