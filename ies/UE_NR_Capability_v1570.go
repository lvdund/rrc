package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type UE_NR_Capability_v1570 struct {
	Nrdc_Parameters_v1570 *NRDC_Parameters_v1570  `optional`
	NonCriticalExtension  *UE_NR_Capability_v1610 `optional`
}

func (ie *UE_NR_Capability_v1570) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Nrdc_Parameters_v1570 != nil, ie.NonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Nrdc_Parameters_v1570 != nil {
		if err = ie.Nrdc_Parameters_v1570.Encode(w); err != nil {
			return utils.WrapError("Encode Nrdc_Parameters_v1570", err)
		}
	}
	if ie.NonCriticalExtension != nil {
		if err = ie.NonCriticalExtension.Encode(w); err != nil {
			return utils.WrapError("Encode NonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *UE_NR_Capability_v1570) Decode(r *aper.AperReader) error {
	var err error
	var Nrdc_Parameters_v1570Present bool
	if Nrdc_Parameters_v1570Present, err = r.ReadBool(); err != nil {
		return err
	}
	var NonCriticalExtensionPresent bool
	if NonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if Nrdc_Parameters_v1570Present {
		ie.Nrdc_Parameters_v1570 = new(NRDC_Parameters_v1570)
		if err = ie.Nrdc_Parameters_v1570.Decode(r); err != nil {
			return utils.WrapError("Decode Nrdc_Parameters_v1570", err)
		}
	}
	if NonCriticalExtensionPresent {
		ie.NonCriticalExtension = new(UE_NR_Capability_v1610)
		if err = ie.NonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode NonCriticalExtension", err)
		}
	}
	return nil
}
