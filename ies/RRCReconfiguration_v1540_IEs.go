package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type RRCReconfiguration_v1540_IEs struct {
	OtherConfig_v1540    *OtherConfig_v1540            `optional`
	NonCriticalExtension *RRCReconfiguration_v1560_IEs `optional`
}

func (ie *RRCReconfiguration_v1540_IEs) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.OtherConfig_v1540 != nil, ie.NonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.OtherConfig_v1540 != nil {
		if err = ie.OtherConfig_v1540.Encode(w); err != nil {
			return utils.WrapError("Encode OtherConfig_v1540", err)
		}
	}
	if ie.NonCriticalExtension != nil {
		if err = ie.NonCriticalExtension.Encode(w); err != nil {
			return utils.WrapError("Encode NonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *RRCReconfiguration_v1540_IEs) Decode(r *aper.AperReader) error {
	var err error
	var OtherConfig_v1540Present bool
	if OtherConfig_v1540Present, err = r.ReadBool(); err != nil {
		return err
	}
	var NonCriticalExtensionPresent bool
	if NonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if OtherConfig_v1540Present {
		ie.OtherConfig_v1540 = new(OtherConfig_v1540)
		if err = ie.OtherConfig_v1540.Decode(r); err != nil {
			return utils.WrapError("Decode OtherConfig_v1540", err)
		}
	}
	if NonCriticalExtensionPresent {
		ie.NonCriticalExtension = new(RRCReconfiguration_v1560_IEs)
		if err = ie.NonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode NonCriticalExtension", err)
		}
	}
	return nil
}
