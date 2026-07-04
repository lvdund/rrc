// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: NRDC-Parameters (line 22645).

var nRDCParametersConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "measAndMobParametersNRDC", Optional: true},
		{Name: "generalParametersNRDC", Optional: true},
		{Name: "fdd-Add-UE-NRDC-Capabilities", Optional: true},
		{Name: "tdd-Add-UE-NRDC-Capabilities", Optional: true},
		{Name: "fr1-Add-UE-NRDC-Capabilities", Optional: true},
		{Name: "fr2-Add-UE-NRDC-Capabilities", Optional: true},
		{Name: "dummy2", Optional: true},
		{Name: "dummy", Optional: true},
	},
}

var nRDCParametersDummy2Constraints = per.SizeConstraints{}

type NRDC_Parameters struct {
	MeasAndMobParametersNRDC     *MeasAndMobParametersMRDC
	GeneralParametersNRDC        *GeneralParametersMRDC_XDD_Diff
	Fdd_Add_UE_NRDC_Capabilities *UE_MRDC_CapabilityAddXDD_Mode
	Tdd_Add_UE_NRDC_Capabilities *UE_MRDC_CapabilityAddXDD_Mode
	Fr1_Add_UE_NRDC_Capabilities *UE_MRDC_CapabilityAddFRX_Mode
	Fr2_Add_UE_NRDC_Capabilities *UE_MRDC_CapabilityAddFRX_Mode
	Dummy2                       []byte
	Dummy                        *struct{}
}

func (ie *NRDC_Parameters) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(nRDCParametersConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.MeasAndMobParametersNRDC != nil, ie.GeneralParametersNRDC != nil, ie.Fdd_Add_UE_NRDC_Capabilities != nil, ie.Tdd_Add_UE_NRDC_Capabilities != nil, ie.Fr1_Add_UE_NRDC_Capabilities != nil, ie.Fr2_Add_UE_NRDC_Capabilities != nil, ie.Dummy2 != nil, ie.Dummy != nil}); err != nil {
		return err
	}

	// 2. measAndMobParametersNRDC: ref
	{
		if ie.MeasAndMobParametersNRDC != nil {
			if err := ie.MeasAndMobParametersNRDC.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. generalParametersNRDC: ref
	{
		if ie.GeneralParametersNRDC != nil {
			if err := ie.GeneralParametersNRDC.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. fdd-Add-UE-NRDC-Capabilities: ref
	{
		if ie.Fdd_Add_UE_NRDC_Capabilities != nil {
			if err := ie.Fdd_Add_UE_NRDC_Capabilities.Encode(e); err != nil {
				return err
			}
		}
	}

	// 5. tdd-Add-UE-NRDC-Capabilities: ref
	{
		if ie.Tdd_Add_UE_NRDC_Capabilities != nil {
			if err := ie.Tdd_Add_UE_NRDC_Capabilities.Encode(e); err != nil {
				return err
			}
		}
	}

	// 6. fr1-Add-UE-NRDC-Capabilities: ref
	{
		if ie.Fr1_Add_UE_NRDC_Capabilities != nil {
			if err := ie.Fr1_Add_UE_NRDC_Capabilities.Encode(e); err != nil {
				return err
			}
		}
	}

	// 7. fr2-Add-UE-NRDC-Capabilities: ref
	{
		if ie.Fr2_Add_UE_NRDC_Capabilities != nil {
			if err := ie.Fr2_Add_UE_NRDC_Capabilities.Encode(e); err != nil {
				return err
			}
		}
	}

	// 8. dummy2: octet-string
	{
		if ie.Dummy2 != nil {
			if err := e.EncodeOctetString(ie.Dummy2, nRDCParametersDummy2Constraints); err != nil {
				return err
			}
		}
	}

	// 9. dummy: sequence
	{
		if ie.Dummy != nil {
			// empty SEQUENCE {}: no content to encode.
		}
	}

	return nil
}

func (ie *NRDC_Parameters) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(nRDCParametersConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. measAndMobParametersNRDC: ref
	{
		if seq.IsComponentPresent(0) {
			ie.MeasAndMobParametersNRDC = new(MeasAndMobParametersMRDC)
			if err := ie.MeasAndMobParametersNRDC.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. generalParametersNRDC: ref
	{
		if seq.IsComponentPresent(1) {
			ie.GeneralParametersNRDC = new(GeneralParametersMRDC_XDD_Diff)
			if err := ie.GeneralParametersNRDC.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. fdd-Add-UE-NRDC-Capabilities: ref
	{
		if seq.IsComponentPresent(2) {
			ie.Fdd_Add_UE_NRDC_Capabilities = new(UE_MRDC_CapabilityAddXDD_Mode)
			if err := ie.Fdd_Add_UE_NRDC_Capabilities.Decode(d); err != nil {
				return err
			}
		}
	}

	// 5. tdd-Add-UE-NRDC-Capabilities: ref
	{
		if seq.IsComponentPresent(3) {
			ie.Tdd_Add_UE_NRDC_Capabilities = new(UE_MRDC_CapabilityAddXDD_Mode)
			if err := ie.Tdd_Add_UE_NRDC_Capabilities.Decode(d); err != nil {
				return err
			}
		}
	}

	// 6. fr1-Add-UE-NRDC-Capabilities: ref
	{
		if seq.IsComponentPresent(4) {
			ie.Fr1_Add_UE_NRDC_Capabilities = new(UE_MRDC_CapabilityAddFRX_Mode)
			if err := ie.Fr1_Add_UE_NRDC_Capabilities.Decode(d); err != nil {
				return err
			}
		}
	}

	// 7. fr2-Add-UE-NRDC-Capabilities: ref
	{
		if seq.IsComponentPresent(5) {
			ie.Fr2_Add_UE_NRDC_Capabilities = new(UE_MRDC_CapabilityAddFRX_Mode)
			if err := ie.Fr2_Add_UE_NRDC_Capabilities.Decode(d); err != nil {
				return err
			}
		}
	}

	// 8. dummy2: octet-string
	{
		if seq.IsComponentPresent(6) {
			v, err := d.DecodeOctetString(nRDCParametersDummy2Constraints)
			if err != nil {
				return err
			}
			ie.Dummy2 = v
		}
	}

	// 9. dummy: sequence
	{
		if seq.IsComponentPresent(7) {
			ie.Dummy = &struct{}{}
		}
	}

	return nil
}
