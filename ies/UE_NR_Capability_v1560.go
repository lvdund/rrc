package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type UE_NR_Capability_v1560 struct {
	Nrdc_Parameters      *NRDC_Parameters        `optional`
	ReceivedFilters      *[]byte                 `optional`
	NonCriticalExtension *UE_NR_Capability_v1570 `optional`
}

func (ie *UE_NR_Capability_v1560) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Nrdc_Parameters != nil, ie.ReceivedFilters != nil, ie.NonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Nrdc_Parameters != nil {
		if err = ie.Nrdc_Parameters.Encode(w); err != nil {
			return utils.WrapError("Encode Nrdc_Parameters", err)
		}
	}
	if ie.ReceivedFilters != nil {
		if err = w.WriteOctetString(*ie.ReceivedFilters, &aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode ReceivedFilters", err)
		}
	}
	if ie.NonCriticalExtension != nil {
		if err = ie.NonCriticalExtension.Encode(w); err != nil {
			return utils.WrapError("Encode NonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *UE_NR_Capability_v1560) Decode(r *aper.AperReader) error {
	var err error
	var Nrdc_ParametersPresent bool
	if Nrdc_ParametersPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var ReceivedFiltersPresent bool
	if ReceivedFiltersPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var NonCriticalExtensionPresent bool
	if NonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if Nrdc_ParametersPresent {
		ie.Nrdc_Parameters = new(NRDC_Parameters)
		if err = ie.Nrdc_Parameters.Decode(r); err != nil {
			return utils.WrapError("Decode Nrdc_Parameters", err)
		}
	}
	if ReceivedFiltersPresent {
		var tmp_os_ReceivedFilters []byte
		if tmp_os_ReceivedFilters, err = r.ReadOctetString(&aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode ReceivedFilters", err)
		}
		ie.ReceivedFilters = &tmp_os_ReceivedFilters
	}
	if NonCriticalExtensionPresent {
		ie.NonCriticalExtension = new(UE_NR_Capability_v1570)
		if err = ie.NonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode NonCriticalExtension", err)
		}
	}
	return nil
}
