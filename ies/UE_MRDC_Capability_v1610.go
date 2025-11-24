package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type UE_MRDC_Capability_v1610 struct {
	MeasAndMobParametersMRDC_v1610 *MeasAndMobParametersMRDC_v1610 `optional`
	GeneralParametersMRDC_v1610    *GeneralParametersMRDC_v1610    `optional`
	Pdcp_ParametersMRDC_v1610      *PDCP_ParametersMRDC_v1610      `optional`
	NonCriticalExtension           *UE_MRDC_Capability_v1700       `optional`
}

func (ie *UE_MRDC_Capability_v1610) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.MeasAndMobParametersMRDC_v1610 != nil, ie.GeneralParametersMRDC_v1610 != nil, ie.Pdcp_ParametersMRDC_v1610 != nil, ie.NonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.MeasAndMobParametersMRDC_v1610 != nil {
		if err = ie.MeasAndMobParametersMRDC_v1610.Encode(w); err != nil {
			return utils.WrapError("Encode MeasAndMobParametersMRDC_v1610", err)
		}
	}
	if ie.GeneralParametersMRDC_v1610 != nil {
		if err = ie.GeneralParametersMRDC_v1610.Encode(w); err != nil {
			return utils.WrapError("Encode GeneralParametersMRDC_v1610", err)
		}
	}
	if ie.Pdcp_ParametersMRDC_v1610 != nil {
		if err = ie.Pdcp_ParametersMRDC_v1610.Encode(w); err != nil {
			return utils.WrapError("Encode Pdcp_ParametersMRDC_v1610", err)
		}
	}
	if ie.NonCriticalExtension != nil {
		if err = ie.NonCriticalExtension.Encode(w); err != nil {
			return utils.WrapError("Encode NonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *UE_MRDC_Capability_v1610) Decode(r *uper.UperReader) error {
	var err error
	var MeasAndMobParametersMRDC_v1610Present bool
	if MeasAndMobParametersMRDC_v1610Present, err = r.ReadBool(); err != nil {
		return err
	}
	var GeneralParametersMRDC_v1610Present bool
	if GeneralParametersMRDC_v1610Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Pdcp_ParametersMRDC_v1610Present bool
	if Pdcp_ParametersMRDC_v1610Present, err = r.ReadBool(); err != nil {
		return err
	}
	var NonCriticalExtensionPresent bool
	if NonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if MeasAndMobParametersMRDC_v1610Present {
		ie.MeasAndMobParametersMRDC_v1610 = new(MeasAndMobParametersMRDC_v1610)
		if err = ie.MeasAndMobParametersMRDC_v1610.Decode(r); err != nil {
			return utils.WrapError("Decode MeasAndMobParametersMRDC_v1610", err)
		}
	}
	if GeneralParametersMRDC_v1610Present {
		ie.GeneralParametersMRDC_v1610 = new(GeneralParametersMRDC_v1610)
		if err = ie.GeneralParametersMRDC_v1610.Decode(r); err != nil {
			return utils.WrapError("Decode GeneralParametersMRDC_v1610", err)
		}
	}
	if Pdcp_ParametersMRDC_v1610Present {
		ie.Pdcp_ParametersMRDC_v1610 = new(PDCP_ParametersMRDC_v1610)
		if err = ie.Pdcp_ParametersMRDC_v1610.Decode(r); err != nil {
			return utils.WrapError("Decode Pdcp_ParametersMRDC_v1610", err)
		}
	}
	if NonCriticalExtensionPresent {
		ie.NonCriticalExtension = new(UE_MRDC_Capability_v1700)
		if err = ie.NonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode NonCriticalExtension", err)
		}
	}
	return nil
}
