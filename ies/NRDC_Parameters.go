package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type NRDC_Parameters struct {
	MeasAndMobParametersNRDC     *MeasAndMobParametersMRDC       `optional`
	GeneralParametersNRDC        *GeneralParametersMRDC_XDD_Diff `optional`
	Fdd_Add_UE_NRDC_Capabilities *UE_MRDC_CapabilityAddXDD_Mode  `optional`
	Tdd_Add_UE_NRDC_Capabilities *UE_MRDC_CapabilityAddXDD_Mode  `optional`
	Fr1_Add_UE_NRDC_Capabilities *UE_MRDC_CapabilityAddFRX_Mode  `optional`
	Fr2_Add_UE_NRDC_Capabilities *UE_MRDC_CapabilityAddFRX_Mode  `optional`
	Dummy2                       *[]byte                         `optional`
	Dummy                        interface{}                     `optional`
}

func (ie *NRDC_Parameters) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.MeasAndMobParametersNRDC != nil, ie.GeneralParametersNRDC != nil, ie.Fdd_Add_UE_NRDC_Capabilities != nil, ie.Tdd_Add_UE_NRDC_Capabilities != nil, ie.Fr1_Add_UE_NRDC_Capabilities != nil, ie.Fr2_Add_UE_NRDC_Capabilities != nil, ie.Dummy2 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.MeasAndMobParametersNRDC != nil {
		if err = ie.MeasAndMobParametersNRDC.Encode(w); err != nil {
			return utils.WrapError("Encode MeasAndMobParametersNRDC", err)
		}
	}
	if ie.GeneralParametersNRDC != nil {
		if err = ie.GeneralParametersNRDC.Encode(w); err != nil {
			return utils.WrapError("Encode GeneralParametersNRDC", err)
		}
	}
	if ie.Fdd_Add_UE_NRDC_Capabilities != nil {
		if err = ie.Fdd_Add_UE_NRDC_Capabilities.Encode(w); err != nil {
			return utils.WrapError("Encode Fdd_Add_UE_NRDC_Capabilities", err)
		}
	}
	if ie.Tdd_Add_UE_NRDC_Capabilities != nil {
		if err = ie.Tdd_Add_UE_NRDC_Capabilities.Encode(w); err != nil {
			return utils.WrapError("Encode Tdd_Add_UE_NRDC_Capabilities", err)
		}
	}
	if ie.Fr1_Add_UE_NRDC_Capabilities != nil {
		if err = ie.Fr1_Add_UE_NRDC_Capabilities.Encode(w); err != nil {
			return utils.WrapError("Encode Fr1_Add_UE_NRDC_Capabilities", err)
		}
	}
	if ie.Fr2_Add_UE_NRDC_Capabilities != nil {
		if err = ie.Fr2_Add_UE_NRDC_Capabilities.Encode(w); err != nil {
			return utils.WrapError("Encode Fr2_Add_UE_NRDC_Capabilities", err)
		}
	}
	if ie.Dummy2 != nil {
		if err = w.WriteOctetString(*ie.Dummy2, &aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode Dummy2", err)
		}
	}
	return nil
}

func (ie *NRDC_Parameters) Decode(r *aper.AperReader) error {
	var err error
	var MeasAndMobParametersNRDCPresent bool
	if MeasAndMobParametersNRDCPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var GeneralParametersNRDCPresent bool
	if GeneralParametersNRDCPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Fdd_Add_UE_NRDC_CapabilitiesPresent bool
	if Fdd_Add_UE_NRDC_CapabilitiesPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Tdd_Add_UE_NRDC_CapabilitiesPresent bool
	if Tdd_Add_UE_NRDC_CapabilitiesPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Fr1_Add_UE_NRDC_CapabilitiesPresent bool
	if Fr1_Add_UE_NRDC_CapabilitiesPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Fr2_Add_UE_NRDC_CapabilitiesPresent bool
	if Fr2_Add_UE_NRDC_CapabilitiesPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Dummy2Present bool
	if Dummy2Present, err = r.ReadBool(); err != nil {
		return err
	}
	if MeasAndMobParametersNRDCPresent {
		ie.MeasAndMobParametersNRDC = new(MeasAndMobParametersMRDC)
		if err = ie.MeasAndMobParametersNRDC.Decode(r); err != nil {
			return utils.WrapError("Decode MeasAndMobParametersNRDC", err)
		}
	}
	if GeneralParametersNRDCPresent {
		ie.GeneralParametersNRDC = new(GeneralParametersMRDC_XDD_Diff)
		if err = ie.GeneralParametersNRDC.Decode(r); err != nil {
			return utils.WrapError("Decode GeneralParametersNRDC", err)
		}
	}
	if Fdd_Add_UE_NRDC_CapabilitiesPresent {
		ie.Fdd_Add_UE_NRDC_Capabilities = new(UE_MRDC_CapabilityAddXDD_Mode)
		if err = ie.Fdd_Add_UE_NRDC_Capabilities.Decode(r); err != nil {
			return utils.WrapError("Decode Fdd_Add_UE_NRDC_Capabilities", err)
		}
	}
	if Tdd_Add_UE_NRDC_CapabilitiesPresent {
		ie.Tdd_Add_UE_NRDC_Capabilities = new(UE_MRDC_CapabilityAddXDD_Mode)
		if err = ie.Tdd_Add_UE_NRDC_Capabilities.Decode(r); err != nil {
			return utils.WrapError("Decode Tdd_Add_UE_NRDC_Capabilities", err)
		}
	}
	if Fr1_Add_UE_NRDC_CapabilitiesPresent {
		ie.Fr1_Add_UE_NRDC_Capabilities = new(UE_MRDC_CapabilityAddFRX_Mode)
		if err = ie.Fr1_Add_UE_NRDC_Capabilities.Decode(r); err != nil {
			return utils.WrapError("Decode Fr1_Add_UE_NRDC_Capabilities", err)
		}
	}
	if Fr2_Add_UE_NRDC_CapabilitiesPresent {
		ie.Fr2_Add_UE_NRDC_Capabilities = new(UE_MRDC_CapabilityAddFRX_Mode)
		if err = ie.Fr2_Add_UE_NRDC_Capabilities.Decode(r); err != nil {
			return utils.WrapError("Decode Fr2_Add_UE_NRDC_Capabilities", err)
		}
	}
	if Dummy2Present {
		var tmp_os_Dummy2 []byte
		if tmp_os_Dummy2, err = r.ReadOctetString(&aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode Dummy2", err)
		}
		ie.Dummy2 = &tmp_os_Dummy2
	}
	return nil
}
