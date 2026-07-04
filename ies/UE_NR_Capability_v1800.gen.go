// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: UE-NR-Capability-v1800 (line 25849).

var uENRCapabilityV1800Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "airToGroundNetwork-r18", Optional: true},
		{Name: "eRedCapParameters-r18", Optional: true},
		{Name: "ncr-Parameters-r18", Optional: true},
		{Name: "softSatelliteSwitchResyncNTN-r18", Optional: true},
		{Name: "hardSatelliteSwitchResyncNTN-r18", Optional: true},
		{Name: "mt-SDT-r18", Optional: true},
		{Name: "mt-SDT-NTN-r18", Optional: true},
		{Name: "inDeviceCoexIndAutonomousDenial-r18", Optional: true},
		{Name: "inDeviceCoexIndFDM-r18", Optional: true},
		{Name: "inDeviceCoexIndTDM-r18", Optional: true},
		{Name: "musim-GapPriorityPreference-r18", Optional: true},
		{Name: "musim-CapabilityRestriction-r18", Optional: true},
		{Name: "dummy", Optional: true},
		{Name: "ra-InsteadCG-SDT-r18", Optional: true},
		{Name: "resumeAfterSDT-Release-r18", Optional: true},
		{Name: "ul-TrafficInfo-r18", Optional: true},
		{Name: "aerialParameters-r18", Optional: true},
		{Name: "ntn-VSAT-AntennaType-r18", Optional: true},
		{Name: "ntn-VSAT-MobilityType-r18", Optional: true},
		{Name: "ntn-Parameters-v1820", Optional: true},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

const (
	UE_NR_Capability_v1800_AirToGroundNetwork_r18_Supported = 0
)

var uENRCapabilityV1800AirToGroundNetworkR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UE_NR_Capability_v1800_AirToGroundNetwork_r18_Supported},
}

const (
	UE_NR_Capability_v1800_SoftSatelliteSwitchResyncNTN_r18_Supported = 0
)

var uENRCapabilityV1800SoftSatelliteSwitchResyncNTNR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UE_NR_Capability_v1800_SoftSatelliteSwitchResyncNTN_r18_Supported},
}

const (
	UE_NR_Capability_v1800_HardSatelliteSwitchResyncNTN_r18_Supported = 0
)

var uENRCapabilityV1800HardSatelliteSwitchResyncNTNR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UE_NR_Capability_v1800_HardSatelliteSwitchResyncNTN_r18_Supported},
}

const (
	UE_NR_Capability_v1800_Mt_SDT_r18_Supported = 0
)

var uENRCapabilityV1800MtSDTR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UE_NR_Capability_v1800_Mt_SDT_r18_Supported},
}

const (
	UE_NR_Capability_v1800_Mt_SDT_NTN_r18_Supported = 0
)

var uENRCapabilityV1800MtSDTNTNR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UE_NR_Capability_v1800_Mt_SDT_NTN_r18_Supported},
}

const (
	UE_NR_Capability_v1800_InDeviceCoexIndAutonomousDenial_r18_Supported = 0
)

var uENRCapabilityV1800InDeviceCoexIndAutonomousDenialR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UE_NR_Capability_v1800_InDeviceCoexIndAutonomousDenial_r18_Supported},
}

const (
	UE_NR_Capability_v1800_InDeviceCoexIndFDM_r18_Supported = 0
)

var uENRCapabilityV1800InDeviceCoexIndFDMR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UE_NR_Capability_v1800_InDeviceCoexIndFDM_r18_Supported},
}

const (
	UE_NR_Capability_v1800_InDeviceCoexIndTDM_r18_Supported = 0
)

var uENRCapabilityV1800InDeviceCoexIndTDMR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UE_NR_Capability_v1800_InDeviceCoexIndTDM_r18_Supported},
}

const (
	UE_NR_Capability_v1800_Musim_GapPriorityPreference_r18_Supported = 0
)

var uENRCapabilityV1800MusimGapPriorityPreferenceR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UE_NR_Capability_v1800_Musim_GapPriorityPreference_r18_Supported},
}

const (
	UE_NR_Capability_v1800_Musim_CapabilityRestriction_r18_Supported = 0
)

var uENRCapabilityV1800MusimCapabilityRestrictionR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UE_NR_Capability_v1800_Musim_CapabilityRestriction_r18_Supported},
}

const (
	UE_NR_Capability_v1800_Dummy_Supported = 0
)

var uENRCapabilityV1800DummyConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UE_NR_Capability_v1800_Dummy_Supported},
}

const (
	UE_NR_Capability_v1800_Ra_InsteadCG_SDT_r18_Supported = 0
)

var uENRCapabilityV1800RaInsteadCGSDTR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UE_NR_Capability_v1800_Ra_InsteadCG_SDT_r18_Supported},
}

const (
	UE_NR_Capability_v1800_ResumeAfterSDT_Release_r18_Supported = 0
)

var uENRCapabilityV1800ResumeAfterSDTReleaseR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UE_NR_Capability_v1800_ResumeAfterSDT_Release_r18_Supported},
}

const (
	UE_NR_Capability_v1800_Ul_TrafficInfo_r18_Supported = 0
)

var uENRCapabilityV1800UlTrafficInfoR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UE_NR_Capability_v1800_Ul_TrafficInfo_r18_Supported},
}

const (
	UE_NR_Capability_v1800_Ntn_VSAT_AntennaType_r18_Electronic = 0
	UE_NR_Capability_v1800_Ntn_VSAT_AntennaType_r18_Mechanical = 1
)

var uENRCapabilityV1800NtnVSATAntennaTypeR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UE_NR_Capability_v1800_Ntn_VSAT_AntennaType_r18_Electronic, UE_NR_Capability_v1800_Ntn_VSAT_AntennaType_r18_Mechanical},
}

const (
	UE_NR_Capability_v1800_Ntn_VSAT_MobilityType_r18_Fixed  = 0
	UE_NR_Capability_v1800_Ntn_VSAT_MobilityType_r18_Mobile = 1
)

var uENRCapabilityV1800NtnVSATMobilityTypeR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UE_NR_Capability_v1800_Ntn_VSAT_MobilityType_r18_Fixed, UE_NR_Capability_v1800_Ntn_VSAT_MobilityType_r18_Mobile},
}

type UE_NR_Capability_v1800 struct {
	AirToGroundNetwork_r18              *int64
	ERedCapParameters_r18               *ERedCapParameters_r18
	Ncr_Parameters_r18                  *NCR_Parameters_r18
	SoftSatelliteSwitchResyncNTN_r18    *int64
	HardSatelliteSwitchResyncNTN_r18    *int64
	Mt_SDT_r18                          *int64
	Mt_SDT_NTN_r18                      *int64
	InDeviceCoexIndAutonomousDenial_r18 *int64
	InDeviceCoexIndFDM_r18              *int64
	InDeviceCoexIndTDM_r18              *int64
	Musim_GapPriorityPreference_r18     *int64
	Musim_CapabilityRestriction_r18     *int64
	Dummy                               *int64
	Ra_InsteadCG_SDT_r18                *int64
	ResumeAfterSDT_Release_r18          *int64
	Ul_TrafficInfo_r18                  *int64
	AerialParameters_r18                *AerialParameters_r18
	Ntn_VSAT_AntennaType_r18            *int64
	Ntn_VSAT_MobilityType_r18           *int64
	Ntn_Parameters_v1820                *NTN_Parameters_v1820
	NonCriticalExtension                *UE_NR_Capability_v1830
}

func (ie *UE_NR_Capability_v1800) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(uENRCapabilityV1800Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.AirToGroundNetwork_r18 != nil, ie.ERedCapParameters_r18 != nil, ie.Ncr_Parameters_r18 != nil, ie.SoftSatelliteSwitchResyncNTN_r18 != nil, ie.HardSatelliteSwitchResyncNTN_r18 != nil, ie.Mt_SDT_r18 != nil, ie.Mt_SDT_NTN_r18 != nil, ie.InDeviceCoexIndAutonomousDenial_r18 != nil, ie.InDeviceCoexIndFDM_r18 != nil, ie.InDeviceCoexIndTDM_r18 != nil, ie.Musim_GapPriorityPreference_r18 != nil, ie.Musim_CapabilityRestriction_r18 != nil, ie.Dummy != nil, ie.Ra_InsteadCG_SDT_r18 != nil, ie.ResumeAfterSDT_Release_r18 != nil, ie.Ul_TrafficInfo_r18 != nil, ie.AerialParameters_r18 != nil, ie.Ntn_VSAT_AntennaType_r18 != nil, ie.Ntn_VSAT_MobilityType_r18 != nil, ie.Ntn_Parameters_v1820 != nil, ie.NonCriticalExtension != nil}); err != nil {
		return err
	}

	// 2. airToGroundNetwork-r18: enumerated
	{
		if ie.AirToGroundNetwork_r18 != nil {
			if err := e.EncodeEnumerated(*ie.AirToGroundNetwork_r18, uENRCapabilityV1800AirToGroundNetworkR18Constraints); err != nil {
				return err
			}
		}
	}

	// 3. eRedCapParameters-r18: ref
	{
		if ie.ERedCapParameters_r18 != nil {
			if err := ie.ERedCapParameters_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. ncr-Parameters-r18: ref
	{
		if ie.Ncr_Parameters_r18 != nil {
			if err := ie.Ncr_Parameters_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	// 5. softSatelliteSwitchResyncNTN-r18: enumerated
	{
		if ie.SoftSatelliteSwitchResyncNTN_r18 != nil {
			if err := e.EncodeEnumerated(*ie.SoftSatelliteSwitchResyncNTN_r18, uENRCapabilityV1800SoftSatelliteSwitchResyncNTNR18Constraints); err != nil {
				return err
			}
		}
	}

	// 6. hardSatelliteSwitchResyncNTN-r18: enumerated
	{
		if ie.HardSatelliteSwitchResyncNTN_r18 != nil {
			if err := e.EncodeEnumerated(*ie.HardSatelliteSwitchResyncNTN_r18, uENRCapabilityV1800HardSatelliteSwitchResyncNTNR18Constraints); err != nil {
				return err
			}
		}
	}

	// 7. mt-SDT-r18: enumerated
	{
		if ie.Mt_SDT_r18 != nil {
			if err := e.EncodeEnumerated(*ie.Mt_SDT_r18, uENRCapabilityV1800MtSDTR18Constraints); err != nil {
				return err
			}
		}
	}

	// 8. mt-SDT-NTN-r18: enumerated
	{
		if ie.Mt_SDT_NTN_r18 != nil {
			if err := e.EncodeEnumerated(*ie.Mt_SDT_NTN_r18, uENRCapabilityV1800MtSDTNTNR18Constraints); err != nil {
				return err
			}
		}
	}

	// 9. inDeviceCoexIndAutonomousDenial-r18: enumerated
	{
		if ie.InDeviceCoexIndAutonomousDenial_r18 != nil {
			if err := e.EncodeEnumerated(*ie.InDeviceCoexIndAutonomousDenial_r18, uENRCapabilityV1800InDeviceCoexIndAutonomousDenialR18Constraints); err != nil {
				return err
			}
		}
	}

	// 10. inDeviceCoexIndFDM-r18: enumerated
	{
		if ie.InDeviceCoexIndFDM_r18 != nil {
			if err := e.EncodeEnumerated(*ie.InDeviceCoexIndFDM_r18, uENRCapabilityV1800InDeviceCoexIndFDMR18Constraints); err != nil {
				return err
			}
		}
	}

	// 11. inDeviceCoexIndTDM-r18: enumerated
	{
		if ie.InDeviceCoexIndTDM_r18 != nil {
			if err := e.EncodeEnumerated(*ie.InDeviceCoexIndTDM_r18, uENRCapabilityV1800InDeviceCoexIndTDMR18Constraints); err != nil {
				return err
			}
		}
	}

	// 12. musim-GapPriorityPreference-r18: enumerated
	{
		if ie.Musim_GapPriorityPreference_r18 != nil {
			if err := e.EncodeEnumerated(*ie.Musim_GapPriorityPreference_r18, uENRCapabilityV1800MusimGapPriorityPreferenceR18Constraints); err != nil {
				return err
			}
		}
	}

	// 13. musim-CapabilityRestriction-r18: enumerated
	{
		if ie.Musim_CapabilityRestriction_r18 != nil {
			if err := e.EncodeEnumerated(*ie.Musim_CapabilityRestriction_r18, uENRCapabilityV1800MusimCapabilityRestrictionR18Constraints); err != nil {
				return err
			}
		}
	}

	// 14. dummy: enumerated
	{
		if ie.Dummy != nil {
			if err := e.EncodeEnumerated(*ie.Dummy, uENRCapabilityV1800DummyConstraints); err != nil {
				return err
			}
		}
	}

	// 15. ra-InsteadCG-SDT-r18: enumerated
	{
		if ie.Ra_InsteadCG_SDT_r18 != nil {
			if err := e.EncodeEnumerated(*ie.Ra_InsteadCG_SDT_r18, uENRCapabilityV1800RaInsteadCGSDTR18Constraints); err != nil {
				return err
			}
		}
	}

	// 16. resumeAfterSDT-Release-r18: enumerated
	{
		if ie.ResumeAfterSDT_Release_r18 != nil {
			if err := e.EncodeEnumerated(*ie.ResumeAfterSDT_Release_r18, uENRCapabilityV1800ResumeAfterSDTReleaseR18Constraints); err != nil {
				return err
			}
		}
	}

	// 17. ul-TrafficInfo-r18: enumerated
	{
		if ie.Ul_TrafficInfo_r18 != nil {
			if err := e.EncodeEnumerated(*ie.Ul_TrafficInfo_r18, uENRCapabilityV1800UlTrafficInfoR18Constraints); err != nil {
				return err
			}
		}
	}

	// 18. aerialParameters-r18: ref
	{
		if ie.AerialParameters_r18 != nil {
			if err := ie.AerialParameters_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	// 19. ntn-VSAT-AntennaType-r18: enumerated
	{
		if ie.Ntn_VSAT_AntennaType_r18 != nil {
			if err := e.EncodeEnumerated(*ie.Ntn_VSAT_AntennaType_r18, uENRCapabilityV1800NtnVSATAntennaTypeR18Constraints); err != nil {
				return err
			}
		}
	}

	// 20. ntn-VSAT-MobilityType-r18: enumerated
	{
		if ie.Ntn_VSAT_MobilityType_r18 != nil {
			if err := e.EncodeEnumerated(*ie.Ntn_VSAT_MobilityType_r18, uENRCapabilityV1800NtnVSATMobilityTypeR18Constraints); err != nil {
				return err
			}
		}
	}

	// 21. ntn-Parameters-v1820: ref
	{
		if ie.Ntn_Parameters_v1820 != nil {
			if err := ie.Ntn_Parameters_v1820.Encode(e); err != nil {
				return err
			}
		}
	}

	// 22. nonCriticalExtension: ref
	{
		if ie.NonCriticalExtension != nil {
			if err := ie.NonCriticalExtension.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *UE_NR_Capability_v1800) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(uENRCapabilityV1800Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. airToGroundNetwork-r18: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(uENRCapabilityV1800AirToGroundNetworkR18Constraints)
			if err != nil {
				return err
			}
			ie.AirToGroundNetwork_r18 = &idx
		}
	}

	// 3. eRedCapParameters-r18: ref
	{
		if seq.IsComponentPresent(1) {
			ie.ERedCapParameters_r18 = new(ERedCapParameters_r18)
			if err := ie.ERedCapParameters_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. ncr-Parameters-r18: ref
	{
		if seq.IsComponentPresent(2) {
			ie.Ncr_Parameters_r18 = new(NCR_Parameters_r18)
			if err := ie.Ncr_Parameters_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	// 5. softSatelliteSwitchResyncNTN-r18: enumerated
	{
		if seq.IsComponentPresent(3) {
			idx, err := d.DecodeEnumerated(uENRCapabilityV1800SoftSatelliteSwitchResyncNTNR18Constraints)
			if err != nil {
				return err
			}
			ie.SoftSatelliteSwitchResyncNTN_r18 = &idx
		}
	}

	// 6. hardSatelliteSwitchResyncNTN-r18: enumerated
	{
		if seq.IsComponentPresent(4) {
			idx, err := d.DecodeEnumerated(uENRCapabilityV1800HardSatelliteSwitchResyncNTNR18Constraints)
			if err != nil {
				return err
			}
			ie.HardSatelliteSwitchResyncNTN_r18 = &idx
		}
	}

	// 7. mt-SDT-r18: enumerated
	{
		if seq.IsComponentPresent(5) {
			idx, err := d.DecodeEnumerated(uENRCapabilityV1800MtSDTR18Constraints)
			if err != nil {
				return err
			}
			ie.Mt_SDT_r18 = &idx
		}
	}

	// 8. mt-SDT-NTN-r18: enumerated
	{
		if seq.IsComponentPresent(6) {
			idx, err := d.DecodeEnumerated(uENRCapabilityV1800MtSDTNTNR18Constraints)
			if err != nil {
				return err
			}
			ie.Mt_SDT_NTN_r18 = &idx
		}
	}

	// 9. inDeviceCoexIndAutonomousDenial-r18: enumerated
	{
		if seq.IsComponentPresent(7) {
			idx, err := d.DecodeEnumerated(uENRCapabilityV1800InDeviceCoexIndAutonomousDenialR18Constraints)
			if err != nil {
				return err
			}
			ie.InDeviceCoexIndAutonomousDenial_r18 = &idx
		}
	}

	// 10. inDeviceCoexIndFDM-r18: enumerated
	{
		if seq.IsComponentPresent(8) {
			idx, err := d.DecodeEnumerated(uENRCapabilityV1800InDeviceCoexIndFDMR18Constraints)
			if err != nil {
				return err
			}
			ie.InDeviceCoexIndFDM_r18 = &idx
		}
	}

	// 11. inDeviceCoexIndTDM-r18: enumerated
	{
		if seq.IsComponentPresent(9) {
			idx, err := d.DecodeEnumerated(uENRCapabilityV1800InDeviceCoexIndTDMR18Constraints)
			if err != nil {
				return err
			}
			ie.InDeviceCoexIndTDM_r18 = &idx
		}
	}

	// 12. musim-GapPriorityPreference-r18: enumerated
	{
		if seq.IsComponentPresent(10) {
			idx, err := d.DecodeEnumerated(uENRCapabilityV1800MusimGapPriorityPreferenceR18Constraints)
			if err != nil {
				return err
			}
			ie.Musim_GapPriorityPreference_r18 = &idx
		}
	}

	// 13. musim-CapabilityRestriction-r18: enumerated
	{
		if seq.IsComponentPresent(11) {
			idx, err := d.DecodeEnumerated(uENRCapabilityV1800MusimCapabilityRestrictionR18Constraints)
			if err != nil {
				return err
			}
			ie.Musim_CapabilityRestriction_r18 = &idx
		}
	}

	// 14. dummy: enumerated
	{
		if seq.IsComponentPresent(12) {
			idx, err := d.DecodeEnumerated(uENRCapabilityV1800DummyConstraints)
			if err != nil {
				return err
			}
			ie.Dummy = &idx
		}
	}

	// 15. ra-InsteadCG-SDT-r18: enumerated
	{
		if seq.IsComponentPresent(13) {
			idx, err := d.DecodeEnumerated(uENRCapabilityV1800RaInsteadCGSDTR18Constraints)
			if err != nil {
				return err
			}
			ie.Ra_InsteadCG_SDT_r18 = &idx
		}
	}

	// 16. resumeAfterSDT-Release-r18: enumerated
	{
		if seq.IsComponentPresent(14) {
			idx, err := d.DecodeEnumerated(uENRCapabilityV1800ResumeAfterSDTReleaseR18Constraints)
			if err != nil {
				return err
			}
			ie.ResumeAfterSDT_Release_r18 = &idx
		}
	}

	// 17. ul-TrafficInfo-r18: enumerated
	{
		if seq.IsComponentPresent(15) {
			idx, err := d.DecodeEnumerated(uENRCapabilityV1800UlTrafficInfoR18Constraints)
			if err != nil {
				return err
			}
			ie.Ul_TrafficInfo_r18 = &idx
		}
	}

	// 18. aerialParameters-r18: ref
	{
		if seq.IsComponentPresent(16) {
			ie.AerialParameters_r18 = new(AerialParameters_r18)
			if err := ie.AerialParameters_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	// 19. ntn-VSAT-AntennaType-r18: enumerated
	{
		if seq.IsComponentPresent(17) {
			idx, err := d.DecodeEnumerated(uENRCapabilityV1800NtnVSATAntennaTypeR18Constraints)
			if err != nil {
				return err
			}
			ie.Ntn_VSAT_AntennaType_r18 = &idx
		}
	}

	// 20. ntn-VSAT-MobilityType-r18: enumerated
	{
		if seq.IsComponentPresent(18) {
			idx, err := d.DecodeEnumerated(uENRCapabilityV1800NtnVSATMobilityTypeR18Constraints)
			if err != nil {
				return err
			}
			ie.Ntn_VSAT_MobilityType_r18 = &idx
		}
	}

	// 21. ntn-Parameters-v1820: ref
	{
		if seq.IsComponentPresent(19) {
			ie.Ntn_Parameters_v1820 = new(NTN_Parameters_v1820)
			if err := ie.Ntn_Parameters_v1820.Decode(d); err != nil {
				return err
			}
		}
	}

	// 22. nonCriticalExtension: ref
	{
		if seq.IsComponentPresent(20) {
			ie.NonCriticalExtension = new(UE_NR_Capability_v1830)
			if err := ie.NonCriticalExtension.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
