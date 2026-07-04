// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: UE-NR-Capability-v1540 (line 25670).

var uENRCapabilityV1540Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "sdap-Parameters", Optional: true},
		{Name: "overheatingInd", Optional: true},
		{Name: "ims-Parameters", Optional: true},
		{Name: "fr1-Add-UE-NR-Capabilities-v1540", Optional: true},
		{Name: "fr2-Add-UE-NR-Capabilities-v1540", Optional: true},
		{Name: "fr1-fr2-Add-UE-NR-Capabilities", Optional: true},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

const (
	UE_NR_Capability_v1540_OverheatingInd_Supported = 0
)

var uENRCapabilityV1540OverheatingIndConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UE_NR_Capability_v1540_OverheatingInd_Supported},
}

type UE_NR_Capability_v1540 struct {
	Sdap_Parameters                  *SDAP_Parameters
	OverheatingInd                   *int64
	Ims_Parameters                   *IMS_Parameters
	Fr1_Add_UE_NR_Capabilities_v1540 *UE_NR_CapabilityAddFRX_Mode_v1540
	Fr2_Add_UE_NR_Capabilities_v1540 *UE_NR_CapabilityAddFRX_Mode_v1540
	Fr1_Fr2_Add_UE_NR_Capabilities   *UE_NR_CapabilityAddFRX_Mode
	NonCriticalExtension             *UE_NR_Capability_v1550
}

func (ie *UE_NR_Capability_v1540) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(uENRCapabilityV1540Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Sdap_Parameters != nil, ie.OverheatingInd != nil, ie.Ims_Parameters != nil, ie.Fr1_Add_UE_NR_Capabilities_v1540 != nil, ie.Fr2_Add_UE_NR_Capabilities_v1540 != nil, ie.Fr1_Fr2_Add_UE_NR_Capabilities != nil, ie.NonCriticalExtension != nil}); err != nil {
		return err
	}

	// 2. sdap-Parameters: ref
	{
		if ie.Sdap_Parameters != nil {
			if err := ie.Sdap_Parameters.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. overheatingInd: enumerated
	{
		if ie.OverheatingInd != nil {
			if err := e.EncodeEnumerated(*ie.OverheatingInd, uENRCapabilityV1540OverheatingIndConstraints); err != nil {
				return err
			}
		}
	}

	// 4. ims-Parameters: ref
	{
		if ie.Ims_Parameters != nil {
			if err := ie.Ims_Parameters.Encode(e); err != nil {
				return err
			}
		}
	}

	// 5. fr1-Add-UE-NR-Capabilities-v1540: ref
	{
		if ie.Fr1_Add_UE_NR_Capabilities_v1540 != nil {
			if err := ie.Fr1_Add_UE_NR_Capabilities_v1540.Encode(e); err != nil {
				return err
			}
		}
	}

	// 6. fr2-Add-UE-NR-Capabilities-v1540: ref
	{
		if ie.Fr2_Add_UE_NR_Capabilities_v1540 != nil {
			if err := ie.Fr2_Add_UE_NR_Capabilities_v1540.Encode(e); err != nil {
				return err
			}
		}
	}

	// 7. fr1-fr2-Add-UE-NR-Capabilities: ref
	{
		if ie.Fr1_Fr2_Add_UE_NR_Capabilities != nil {
			if err := ie.Fr1_Fr2_Add_UE_NR_Capabilities.Encode(e); err != nil {
				return err
			}
		}
	}

	// 8. nonCriticalExtension: ref
	{
		if ie.NonCriticalExtension != nil {
			if err := ie.NonCriticalExtension.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *UE_NR_Capability_v1540) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(uENRCapabilityV1540Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. sdap-Parameters: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Sdap_Parameters = new(SDAP_Parameters)
			if err := ie.Sdap_Parameters.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. overheatingInd: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(uENRCapabilityV1540OverheatingIndConstraints)
			if err != nil {
				return err
			}
			ie.OverheatingInd = &idx
		}
	}

	// 4. ims-Parameters: ref
	{
		if seq.IsComponentPresent(2) {
			ie.Ims_Parameters = new(IMS_Parameters)
			if err := ie.Ims_Parameters.Decode(d); err != nil {
				return err
			}
		}
	}

	// 5. fr1-Add-UE-NR-Capabilities-v1540: ref
	{
		if seq.IsComponentPresent(3) {
			ie.Fr1_Add_UE_NR_Capabilities_v1540 = new(UE_NR_CapabilityAddFRX_Mode_v1540)
			if err := ie.Fr1_Add_UE_NR_Capabilities_v1540.Decode(d); err != nil {
				return err
			}
		}
	}

	// 6. fr2-Add-UE-NR-Capabilities-v1540: ref
	{
		if seq.IsComponentPresent(4) {
			ie.Fr2_Add_UE_NR_Capabilities_v1540 = new(UE_NR_CapabilityAddFRX_Mode_v1540)
			if err := ie.Fr2_Add_UE_NR_Capabilities_v1540.Decode(d); err != nil {
				return err
			}
		}
	}

	// 7. fr1-fr2-Add-UE-NR-Capabilities: ref
	{
		if seq.IsComponentPresent(5) {
			ie.Fr1_Fr2_Add_UE_NR_Capabilities = new(UE_NR_CapabilityAddFRX_Mode)
			if err := ie.Fr1_Fr2_Add_UE_NR_Capabilities.Decode(d); err != nil {
				return err
			}
		}
	}

	// 8. nonCriticalExtension: ref
	{
		if seq.IsComponentPresent(6) {
			ie.NonCriticalExtension = new(UE_NR_Capability_v1550)
			if err := ie.NonCriticalExtension.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
