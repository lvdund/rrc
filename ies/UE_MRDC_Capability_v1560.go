package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type UE_MRDC_Capability_v1560 struct {
	ReceivedFilters                    *[]byte                              `optional`
	MeasAndMobParametersMRDC_v1560     *MeasAndMobParametersMRDC_v1560      `optional`
	Fdd_Add_UE_MRDC_Capabilities_v1560 *UE_MRDC_CapabilityAddXDD_Mode_v1560 `optional`
	Tdd_Add_UE_MRDC_Capabilities_v1560 *UE_MRDC_CapabilityAddXDD_Mode_v1560 `optional`
	NonCriticalExtension               *UE_MRDC_Capability_v1610            `optional`
}

func (ie *UE_MRDC_Capability_v1560) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.ReceivedFilters != nil, ie.MeasAndMobParametersMRDC_v1560 != nil, ie.Fdd_Add_UE_MRDC_Capabilities_v1560 != nil, ie.Tdd_Add_UE_MRDC_Capabilities_v1560 != nil, ie.NonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.ReceivedFilters != nil {
		if err = w.WriteOctetString(*ie.ReceivedFilters, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode ReceivedFilters", err)
		}
	}
	if ie.MeasAndMobParametersMRDC_v1560 != nil {
		if err = ie.MeasAndMobParametersMRDC_v1560.Encode(w); err != nil {
			return utils.WrapError("Encode MeasAndMobParametersMRDC_v1560", err)
		}
	}
	if ie.Fdd_Add_UE_MRDC_Capabilities_v1560 != nil {
		if err = ie.Fdd_Add_UE_MRDC_Capabilities_v1560.Encode(w); err != nil {
			return utils.WrapError("Encode Fdd_Add_UE_MRDC_Capabilities_v1560", err)
		}
	}
	if ie.Tdd_Add_UE_MRDC_Capabilities_v1560 != nil {
		if err = ie.Tdd_Add_UE_MRDC_Capabilities_v1560.Encode(w); err != nil {
			return utils.WrapError("Encode Tdd_Add_UE_MRDC_Capabilities_v1560", err)
		}
	}
	if ie.NonCriticalExtension != nil {
		if err = ie.NonCriticalExtension.Encode(w); err != nil {
			return utils.WrapError("Encode NonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *UE_MRDC_Capability_v1560) Decode(r *uper.UperReader) error {
	var err error
	var ReceivedFiltersPresent bool
	if ReceivedFiltersPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var MeasAndMobParametersMRDC_v1560Present bool
	if MeasAndMobParametersMRDC_v1560Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Fdd_Add_UE_MRDC_Capabilities_v1560Present bool
	if Fdd_Add_UE_MRDC_Capabilities_v1560Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Tdd_Add_UE_MRDC_Capabilities_v1560Present bool
	if Tdd_Add_UE_MRDC_Capabilities_v1560Present, err = r.ReadBool(); err != nil {
		return err
	}
	var NonCriticalExtensionPresent bool
	if NonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if ReceivedFiltersPresent {
		var tmp_os_ReceivedFilters []byte
		if tmp_os_ReceivedFilters, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode ReceivedFilters", err)
		}
		ie.ReceivedFilters = &tmp_os_ReceivedFilters
	}
	if MeasAndMobParametersMRDC_v1560Present {
		ie.MeasAndMobParametersMRDC_v1560 = new(MeasAndMobParametersMRDC_v1560)
		if err = ie.MeasAndMobParametersMRDC_v1560.Decode(r); err != nil {
			return utils.WrapError("Decode MeasAndMobParametersMRDC_v1560", err)
		}
	}
	if Fdd_Add_UE_MRDC_Capabilities_v1560Present {
		ie.Fdd_Add_UE_MRDC_Capabilities_v1560 = new(UE_MRDC_CapabilityAddXDD_Mode_v1560)
		if err = ie.Fdd_Add_UE_MRDC_Capabilities_v1560.Decode(r); err != nil {
			return utils.WrapError("Decode Fdd_Add_UE_MRDC_Capabilities_v1560", err)
		}
	}
	if Tdd_Add_UE_MRDC_Capabilities_v1560Present {
		ie.Tdd_Add_UE_MRDC_Capabilities_v1560 = new(UE_MRDC_CapabilityAddXDD_Mode_v1560)
		if err = ie.Tdd_Add_UE_MRDC_Capabilities_v1560.Decode(r); err != nil {
			return utils.WrapError("Decode Tdd_Add_UE_MRDC_Capabilities_v1560", err)
		}
	}
	if NonCriticalExtensionPresent {
		ie.NonCriticalExtension = new(UE_MRDC_Capability_v1610)
		if err = ie.NonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode NonCriticalExtension", err)
		}
	}
	return nil
}
