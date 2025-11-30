package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type UE_CapabilityRequestFilterNR_v1710 struct {
	SidelinkRequest_r17  *UE_CapabilityRequestFilterNR_v1710_sidelinkRequest_r17 `optional`
	NonCriticalExtension interface{}                                             `optional`
}

func (ie *UE_CapabilityRequestFilterNR_v1710) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.SidelinkRequest_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.SidelinkRequest_r17 != nil {
		if err = ie.SidelinkRequest_r17.Encode(w); err != nil {
			return utils.WrapError("Encode SidelinkRequest_r17", err)
		}
	}
	return nil
}

func (ie *UE_CapabilityRequestFilterNR_v1710) Decode(r *aper.AperReader) error {
	var err error
	var SidelinkRequest_r17Present bool
	if SidelinkRequest_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if SidelinkRequest_r17Present {
		ie.SidelinkRequest_r17 = new(UE_CapabilityRequestFilterNR_v1710_sidelinkRequest_r17)
		if err = ie.SidelinkRequest_r17.Decode(r); err != nil {
			return utils.WrapError("Decode SidelinkRequest_r17", err)
		}
	}
	return nil
}
