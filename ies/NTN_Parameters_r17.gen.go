// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: NTN-Parameters-r17 (line 22677).

var nTNParametersR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "inactiveStateNTN-r17", Optional: true},
		{Name: "ra-SDT-NTN-r17", Optional: true},
		{Name: "srb-SDT-NTN-r17", Optional: true},
		{Name: "measAndMobParametersNTN-r17", Optional: true},
		{Name: "mac-ParametersNTN-r17", Optional: true},
		{Name: "phy-ParametersNTN-r17", Optional: true},
		{Name: "fdd-Add-UE-NR-CapabilitiesNTN-r17", Optional: true},
		{Name: "fr1-Add-UE-NR-CapabilitiesNTN-r17", Optional: true},
		{Name: "ue-BasedPerfMeas-ParametersNTN-r17", Optional: true},
		{Name: "son-ParametersNTN-r17", Optional: true},
	},
}

const (
	NTN_Parameters_r17_InactiveStateNTN_r17_Supported = 0
)

var nTNParametersR17InactiveStateNTNR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{NTN_Parameters_r17_InactiveStateNTN_r17_Supported},
}

const (
	NTN_Parameters_r17_Ra_SDT_NTN_r17_Supported = 0
)

var nTNParametersR17RaSDTNTNR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{NTN_Parameters_r17_Ra_SDT_NTN_r17_Supported},
}

const (
	NTN_Parameters_r17_Srb_SDT_NTN_r17_Supported = 0
)

var nTNParametersR17SrbSDTNTNR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{NTN_Parameters_r17_Srb_SDT_NTN_r17_Supported},
}

type NTN_Parameters_r17 struct {
	InactiveStateNTN_r17               *int64
	Ra_SDT_NTN_r17                     *int64
	Srb_SDT_NTN_r17                    *int64
	MeasAndMobParametersNTN_r17        *MeasAndMobParameters
	Mac_ParametersNTN_r17              *MAC_Parameters
	Phy_ParametersNTN_r17              *Phy_Parameters
	Fdd_Add_UE_NR_CapabilitiesNTN_r17  *UE_NR_CapabilityAddXDD_Mode
	Fr1_Add_UE_NR_CapabilitiesNTN_r17  *UE_NR_CapabilityAddFRX_Mode
	Ue_BasedPerfMeas_ParametersNTN_r17 *UE_BasedPerfMeas_Parameters_r16
	Son_ParametersNTN_r17              *SON_Parameters_r16
}

func (ie *NTN_Parameters_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(nTNParametersR17Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.InactiveStateNTN_r17 != nil, ie.Ra_SDT_NTN_r17 != nil, ie.Srb_SDT_NTN_r17 != nil, ie.MeasAndMobParametersNTN_r17 != nil, ie.Mac_ParametersNTN_r17 != nil, ie.Phy_ParametersNTN_r17 != nil, ie.Fdd_Add_UE_NR_CapabilitiesNTN_r17 != nil, ie.Fr1_Add_UE_NR_CapabilitiesNTN_r17 != nil, ie.Ue_BasedPerfMeas_ParametersNTN_r17 != nil, ie.Son_ParametersNTN_r17 != nil}); err != nil {
		return err
	}

	// 2. inactiveStateNTN-r17: enumerated
	{
		if ie.InactiveStateNTN_r17 != nil {
			if err := e.EncodeEnumerated(*ie.InactiveStateNTN_r17, nTNParametersR17InactiveStateNTNR17Constraints); err != nil {
				return err
			}
		}
	}

	// 3. ra-SDT-NTN-r17: enumerated
	{
		if ie.Ra_SDT_NTN_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Ra_SDT_NTN_r17, nTNParametersR17RaSDTNTNR17Constraints); err != nil {
				return err
			}
		}
	}

	// 4. srb-SDT-NTN-r17: enumerated
	{
		if ie.Srb_SDT_NTN_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Srb_SDT_NTN_r17, nTNParametersR17SrbSDTNTNR17Constraints); err != nil {
				return err
			}
		}
	}

	// 5. measAndMobParametersNTN-r17: ref
	{
		if ie.MeasAndMobParametersNTN_r17 != nil {
			if err := ie.MeasAndMobParametersNTN_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	// 6. mac-ParametersNTN-r17: ref
	{
		if ie.Mac_ParametersNTN_r17 != nil {
			if err := ie.Mac_ParametersNTN_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	// 7. phy-ParametersNTN-r17: ref
	{
		if ie.Phy_ParametersNTN_r17 != nil {
			if err := ie.Phy_ParametersNTN_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	// 8. fdd-Add-UE-NR-CapabilitiesNTN-r17: ref
	{
		if ie.Fdd_Add_UE_NR_CapabilitiesNTN_r17 != nil {
			if err := ie.Fdd_Add_UE_NR_CapabilitiesNTN_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	// 9. fr1-Add-UE-NR-CapabilitiesNTN-r17: ref
	{
		if ie.Fr1_Add_UE_NR_CapabilitiesNTN_r17 != nil {
			if err := ie.Fr1_Add_UE_NR_CapabilitiesNTN_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	// 10. ue-BasedPerfMeas-ParametersNTN-r17: ref
	{
		if ie.Ue_BasedPerfMeas_ParametersNTN_r17 != nil {
			if err := ie.Ue_BasedPerfMeas_ParametersNTN_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	// 11. son-ParametersNTN-r17: ref
	{
		if ie.Son_ParametersNTN_r17 != nil {
			if err := ie.Son_ParametersNTN_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *NTN_Parameters_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(nTNParametersR17Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. inactiveStateNTN-r17: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(nTNParametersR17InactiveStateNTNR17Constraints)
			if err != nil {
				return err
			}
			ie.InactiveStateNTN_r17 = &idx
		}
	}

	// 3. ra-SDT-NTN-r17: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(nTNParametersR17RaSDTNTNR17Constraints)
			if err != nil {
				return err
			}
			ie.Ra_SDT_NTN_r17 = &idx
		}
	}

	// 4. srb-SDT-NTN-r17: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(nTNParametersR17SrbSDTNTNR17Constraints)
			if err != nil {
				return err
			}
			ie.Srb_SDT_NTN_r17 = &idx
		}
	}

	// 5. measAndMobParametersNTN-r17: ref
	{
		if seq.IsComponentPresent(3) {
			ie.MeasAndMobParametersNTN_r17 = new(MeasAndMobParameters)
			if err := ie.MeasAndMobParametersNTN_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	// 6. mac-ParametersNTN-r17: ref
	{
		if seq.IsComponentPresent(4) {
			ie.Mac_ParametersNTN_r17 = new(MAC_Parameters)
			if err := ie.Mac_ParametersNTN_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	// 7. phy-ParametersNTN-r17: ref
	{
		if seq.IsComponentPresent(5) {
			ie.Phy_ParametersNTN_r17 = new(Phy_Parameters)
			if err := ie.Phy_ParametersNTN_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	// 8. fdd-Add-UE-NR-CapabilitiesNTN-r17: ref
	{
		if seq.IsComponentPresent(6) {
			ie.Fdd_Add_UE_NR_CapabilitiesNTN_r17 = new(UE_NR_CapabilityAddXDD_Mode)
			if err := ie.Fdd_Add_UE_NR_CapabilitiesNTN_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	// 9. fr1-Add-UE-NR-CapabilitiesNTN-r17: ref
	{
		if seq.IsComponentPresent(7) {
			ie.Fr1_Add_UE_NR_CapabilitiesNTN_r17 = new(UE_NR_CapabilityAddFRX_Mode)
			if err := ie.Fr1_Add_UE_NR_CapabilitiesNTN_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	// 10. ue-BasedPerfMeas-ParametersNTN-r17: ref
	{
		if seq.IsComponentPresent(8) {
			ie.Ue_BasedPerfMeas_ParametersNTN_r17 = new(UE_BasedPerfMeas_Parameters_r16)
			if err := ie.Ue_BasedPerfMeas_ParametersNTN_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	// 11. son-ParametersNTN-r17: ref
	{
		if seq.IsComponentPresent(9) {
			ie.Son_ParametersNTN_r17 = new(SON_Parameters_r16)
			if err := ie.Son_ParametersNTN_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
