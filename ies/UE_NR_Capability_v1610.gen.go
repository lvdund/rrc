// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: UE-NR-Capability-v1610 (line 25721).

var uENRCapabilityV1610Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "inDeviceCoexInd-r16", Optional: true},
		{Name: "dl-DedicatedMessageSegmentation-r16", Optional: true},
		{Name: "nrdc-Parameters-v1610", Optional: true},
		{Name: "powSav-Parameters-r16", Optional: true},
		{Name: "fr1-Add-UE-NR-Capabilities-v1610", Optional: true},
		{Name: "fr2-Add-UE-NR-Capabilities-v1610", Optional: true},
		{Name: "bh-RLF-Indication-r16", Optional: true},
		{Name: "directSN-AdditionFirstRRC-IAB-r16", Optional: true},
		{Name: "bap-Parameters-r16", Optional: true},
		{Name: "referenceTimeProvision-r16", Optional: true},
		{Name: "sidelinkParameters-r16", Optional: true},
		{Name: "highSpeedParameters-r16", Optional: true},
		{Name: "mac-Parameters-v1610", Optional: true},
		{Name: "mcgRLF-RecoveryViaSCG-r16", Optional: true},
		{Name: "resumeWithStoredMCG-SCells-r16", Optional: true},
		{Name: "resumeWithStoredSCG-r16", Optional: true},
		{Name: "resumeWithSCG-Config-r16", Optional: true},
		{Name: "ue-BasedPerfMeas-Parameters-r16", Optional: true},
		{Name: "son-Parameters-r16", Optional: true},
		{Name: "onDemandSIB-Connected-r16", Optional: true},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

const (
	UE_NR_Capability_v1610_InDeviceCoexInd_r16_Supported = 0
)

var uENRCapabilityV1610InDeviceCoexIndR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UE_NR_Capability_v1610_InDeviceCoexInd_r16_Supported},
}

const (
	UE_NR_Capability_v1610_Dl_DedicatedMessageSegmentation_r16_Supported = 0
)

var uENRCapabilityV1610DlDedicatedMessageSegmentationR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UE_NR_Capability_v1610_Dl_DedicatedMessageSegmentation_r16_Supported},
}

const (
	UE_NR_Capability_v1610_Bh_RLF_Indication_r16_Supported = 0
)

var uENRCapabilityV1610BhRLFIndicationR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UE_NR_Capability_v1610_Bh_RLF_Indication_r16_Supported},
}

const (
	UE_NR_Capability_v1610_DirectSN_AdditionFirstRRC_IAB_r16_Supported = 0
)

var uENRCapabilityV1610DirectSNAdditionFirstRRCIABR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UE_NR_Capability_v1610_DirectSN_AdditionFirstRRC_IAB_r16_Supported},
}

const (
	UE_NR_Capability_v1610_ReferenceTimeProvision_r16_Supported = 0
)

var uENRCapabilityV1610ReferenceTimeProvisionR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UE_NR_Capability_v1610_ReferenceTimeProvision_r16_Supported},
}

const (
	UE_NR_Capability_v1610_McgRLF_RecoveryViaSCG_r16_Supported = 0
)

var uENRCapabilityV1610McgRLFRecoveryViaSCGR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UE_NR_Capability_v1610_McgRLF_RecoveryViaSCG_r16_Supported},
}

const (
	UE_NR_Capability_v1610_ResumeWithStoredMCG_SCells_r16_Supported = 0
)

var uENRCapabilityV1610ResumeWithStoredMCGSCellsR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UE_NR_Capability_v1610_ResumeWithStoredMCG_SCells_r16_Supported},
}

const (
	UE_NR_Capability_v1610_ResumeWithStoredSCG_r16_Supported = 0
)

var uENRCapabilityV1610ResumeWithStoredSCGR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UE_NR_Capability_v1610_ResumeWithStoredSCG_r16_Supported},
}

const (
	UE_NR_Capability_v1610_ResumeWithSCG_Config_r16_Supported = 0
)

var uENRCapabilityV1610ResumeWithSCGConfigR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UE_NR_Capability_v1610_ResumeWithSCG_Config_r16_Supported},
}

const (
	UE_NR_Capability_v1610_OnDemandSIB_Connected_r16_Supported = 0
)

var uENRCapabilityV1610OnDemandSIBConnectedR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UE_NR_Capability_v1610_OnDemandSIB_Connected_r16_Supported},
}

type UE_NR_Capability_v1610 struct {
	InDeviceCoexInd_r16                 *int64
	Dl_DedicatedMessageSegmentation_r16 *int64
	Nrdc_Parameters_v1610               *NRDC_Parameters_v1610
	PowSav_Parameters_r16               *PowSav_Parameters_r16
	Fr1_Add_UE_NR_Capabilities_v1610    *UE_NR_CapabilityAddFRX_Mode_v1610
	Fr2_Add_UE_NR_Capabilities_v1610    *UE_NR_CapabilityAddFRX_Mode_v1610
	Bh_RLF_Indication_r16               *int64
	DirectSN_AdditionFirstRRC_IAB_r16   *int64
	Bap_Parameters_r16                  *BAP_Parameters_r16
	ReferenceTimeProvision_r16          *int64
	SidelinkParameters_r16              *SidelinkParameters_r16
	HighSpeedParameters_r16             *HighSpeedParameters_r16
	Mac_Parameters_v1610                *MAC_Parameters_v1610
	McgRLF_RecoveryViaSCG_r16           *int64
	ResumeWithStoredMCG_SCells_r16      *int64
	ResumeWithStoredSCG_r16             *int64
	ResumeWithSCG_Config_r16            *int64
	Ue_BasedPerfMeas_Parameters_r16     *UE_BasedPerfMeas_Parameters_r16
	Son_Parameters_r16                  *SON_Parameters_r16
	OnDemandSIB_Connected_r16           *int64
	NonCriticalExtension                *UE_NR_Capability_v1640
}

func (ie *UE_NR_Capability_v1610) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(uENRCapabilityV1610Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.InDeviceCoexInd_r16 != nil, ie.Dl_DedicatedMessageSegmentation_r16 != nil, ie.Nrdc_Parameters_v1610 != nil, ie.PowSav_Parameters_r16 != nil, ie.Fr1_Add_UE_NR_Capabilities_v1610 != nil, ie.Fr2_Add_UE_NR_Capabilities_v1610 != nil, ie.Bh_RLF_Indication_r16 != nil, ie.DirectSN_AdditionFirstRRC_IAB_r16 != nil, ie.Bap_Parameters_r16 != nil, ie.ReferenceTimeProvision_r16 != nil, ie.SidelinkParameters_r16 != nil, ie.HighSpeedParameters_r16 != nil, ie.Mac_Parameters_v1610 != nil, ie.McgRLF_RecoveryViaSCG_r16 != nil, ie.ResumeWithStoredMCG_SCells_r16 != nil, ie.ResumeWithStoredSCG_r16 != nil, ie.ResumeWithSCG_Config_r16 != nil, ie.Ue_BasedPerfMeas_Parameters_r16 != nil, ie.Son_Parameters_r16 != nil, ie.OnDemandSIB_Connected_r16 != nil, ie.NonCriticalExtension != nil}); err != nil {
		return err
	}

	// 2. inDeviceCoexInd-r16: enumerated
	{
		if ie.InDeviceCoexInd_r16 != nil {
			if err := e.EncodeEnumerated(*ie.InDeviceCoexInd_r16, uENRCapabilityV1610InDeviceCoexIndR16Constraints); err != nil {
				return err
			}
		}
	}

	// 3. dl-DedicatedMessageSegmentation-r16: enumerated
	{
		if ie.Dl_DedicatedMessageSegmentation_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Dl_DedicatedMessageSegmentation_r16, uENRCapabilityV1610DlDedicatedMessageSegmentationR16Constraints); err != nil {
				return err
			}
		}
	}

	// 4. nrdc-Parameters-v1610: ref
	{
		if ie.Nrdc_Parameters_v1610 != nil {
			if err := ie.Nrdc_Parameters_v1610.Encode(e); err != nil {
				return err
			}
		}
	}

	// 5. powSav-Parameters-r16: ref
	{
		if ie.PowSav_Parameters_r16 != nil {
			if err := ie.PowSav_Parameters_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 6. fr1-Add-UE-NR-Capabilities-v1610: ref
	{
		if ie.Fr1_Add_UE_NR_Capabilities_v1610 != nil {
			if err := ie.Fr1_Add_UE_NR_Capabilities_v1610.Encode(e); err != nil {
				return err
			}
		}
	}

	// 7. fr2-Add-UE-NR-Capabilities-v1610: ref
	{
		if ie.Fr2_Add_UE_NR_Capabilities_v1610 != nil {
			if err := ie.Fr2_Add_UE_NR_Capabilities_v1610.Encode(e); err != nil {
				return err
			}
		}
	}

	// 8. bh-RLF-Indication-r16: enumerated
	{
		if ie.Bh_RLF_Indication_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Bh_RLF_Indication_r16, uENRCapabilityV1610BhRLFIndicationR16Constraints); err != nil {
				return err
			}
		}
	}

	// 9. directSN-AdditionFirstRRC-IAB-r16: enumerated
	{
		if ie.DirectSN_AdditionFirstRRC_IAB_r16 != nil {
			if err := e.EncodeEnumerated(*ie.DirectSN_AdditionFirstRRC_IAB_r16, uENRCapabilityV1610DirectSNAdditionFirstRRCIABR16Constraints); err != nil {
				return err
			}
		}
	}

	// 10. bap-Parameters-r16: ref
	{
		if ie.Bap_Parameters_r16 != nil {
			if err := ie.Bap_Parameters_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 11. referenceTimeProvision-r16: enumerated
	{
		if ie.ReferenceTimeProvision_r16 != nil {
			if err := e.EncodeEnumerated(*ie.ReferenceTimeProvision_r16, uENRCapabilityV1610ReferenceTimeProvisionR16Constraints); err != nil {
				return err
			}
		}
	}

	// 12. sidelinkParameters-r16: ref
	{
		if ie.SidelinkParameters_r16 != nil {
			if err := ie.SidelinkParameters_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 13. highSpeedParameters-r16: ref
	{
		if ie.HighSpeedParameters_r16 != nil {
			if err := ie.HighSpeedParameters_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 14. mac-Parameters-v1610: ref
	{
		if ie.Mac_Parameters_v1610 != nil {
			if err := ie.Mac_Parameters_v1610.Encode(e); err != nil {
				return err
			}
		}
	}

	// 15. mcgRLF-RecoveryViaSCG-r16: enumerated
	{
		if ie.McgRLF_RecoveryViaSCG_r16 != nil {
			if err := e.EncodeEnumerated(*ie.McgRLF_RecoveryViaSCG_r16, uENRCapabilityV1610McgRLFRecoveryViaSCGR16Constraints); err != nil {
				return err
			}
		}
	}

	// 16. resumeWithStoredMCG-SCells-r16: enumerated
	{
		if ie.ResumeWithStoredMCG_SCells_r16 != nil {
			if err := e.EncodeEnumerated(*ie.ResumeWithStoredMCG_SCells_r16, uENRCapabilityV1610ResumeWithStoredMCGSCellsR16Constraints); err != nil {
				return err
			}
		}
	}

	// 17. resumeWithStoredSCG-r16: enumerated
	{
		if ie.ResumeWithStoredSCG_r16 != nil {
			if err := e.EncodeEnumerated(*ie.ResumeWithStoredSCG_r16, uENRCapabilityV1610ResumeWithStoredSCGR16Constraints); err != nil {
				return err
			}
		}
	}

	// 18. resumeWithSCG-Config-r16: enumerated
	{
		if ie.ResumeWithSCG_Config_r16 != nil {
			if err := e.EncodeEnumerated(*ie.ResumeWithSCG_Config_r16, uENRCapabilityV1610ResumeWithSCGConfigR16Constraints); err != nil {
				return err
			}
		}
	}

	// 19. ue-BasedPerfMeas-Parameters-r16: ref
	{
		if ie.Ue_BasedPerfMeas_Parameters_r16 != nil {
			if err := ie.Ue_BasedPerfMeas_Parameters_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 20. son-Parameters-r16: ref
	{
		if ie.Son_Parameters_r16 != nil {
			if err := ie.Son_Parameters_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 21. onDemandSIB-Connected-r16: enumerated
	{
		if ie.OnDemandSIB_Connected_r16 != nil {
			if err := e.EncodeEnumerated(*ie.OnDemandSIB_Connected_r16, uENRCapabilityV1610OnDemandSIBConnectedR16Constraints); err != nil {
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

func (ie *UE_NR_Capability_v1610) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(uENRCapabilityV1610Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. inDeviceCoexInd-r16: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(uENRCapabilityV1610InDeviceCoexIndR16Constraints)
			if err != nil {
				return err
			}
			ie.InDeviceCoexInd_r16 = &idx
		}
	}

	// 3. dl-DedicatedMessageSegmentation-r16: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(uENRCapabilityV1610DlDedicatedMessageSegmentationR16Constraints)
			if err != nil {
				return err
			}
			ie.Dl_DedicatedMessageSegmentation_r16 = &idx
		}
	}

	// 4. nrdc-Parameters-v1610: ref
	{
		if seq.IsComponentPresent(2) {
			ie.Nrdc_Parameters_v1610 = new(NRDC_Parameters_v1610)
			if err := ie.Nrdc_Parameters_v1610.Decode(d); err != nil {
				return err
			}
		}
	}

	// 5. powSav-Parameters-r16: ref
	{
		if seq.IsComponentPresent(3) {
			ie.PowSav_Parameters_r16 = new(PowSav_Parameters_r16)
			if err := ie.PowSav_Parameters_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 6. fr1-Add-UE-NR-Capabilities-v1610: ref
	{
		if seq.IsComponentPresent(4) {
			ie.Fr1_Add_UE_NR_Capabilities_v1610 = new(UE_NR_CapabilityAddFRX_Mode_v1610)
			if err := ie.Fr1_Add_UE_NR_Capabilities_v1610.Decode(d); err != nil {
				return err
			}
		}
	}

	// 7. fr2-Add-UE-NR-Capabilities-v1610: ref
	{
		if seq.IsComponentPresent(5) {
			ie.Fr2_Add_UE_NR_Capabilities_v1610 = new(UE_NR_CapabilityAddFRX_Mode_v1610)
			if err := ie.Fr2_Add_UE_NR_Capabilities_v1610.Decode(d); err != nil {
				return err
			}
		}
	}

	// 8. bh-RLF-Indication-r16: enumerated
	{
		if seq.IsComponentPresent(6) {
			idx, err := d.DecodeEnumerated(uENRCapabilityV1610BhRLFIndicationR16Constraints)
			if err != nil {
				return err
			}
			ie.Bh_RLF_Indication_r16 = &idx
		}
	}

	// 9. directSN-AdditionFirstRRC-IAB-r16: enumerated
	{
		if seq.IsComponentPresent(7) {
			idx, err := d.DecodeEnumerated(uENRCapabilityV1610DirectSNAdditionFirstRRCIABR16Constraints)
			if err != nil {
				return err
			}
			ie.DirectSN_AdditionFirstRRC_IAB_r16 = &idx
		}
	}

	// 10. bap-Parameters-r16: ref
	{
		if seq.IsComponentPresent(8) {
			ie.Bap_Parameters_r16 = new(BAP_Parameters_r16)
			if err := ie.Bap_Parameters_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 11. referenceTimeProvision-r16: enumerated
	{
		if seq.IsComponentPresent(9) {
			idx, err := d.DecodeEnumerated(uENRCapabilityV1610ReferenceTimeProvisionR16Constraints)
			if err != nil {
				return err
			}
			ie.ReferenceTimeProvision_r16 = &idx
		}
	}

	// 12. sidelinkParameters-r16: ref
	{
		if seq.IsComponentPresent(10) {
			ie.SidelinkParameters_r16 = new(SidelinkParameters_r16)
			if err := ie.SidelinkParameters_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 13. highSpeedParameters-r16: ref
	{
		if seq.IsComponentPresent(11) {
			ie.HighSpeedParameters_r16 = new(HighSpeedParameters_r16)
			if err := ie.HighSpeedParameters_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 14. mac-Parameters-v1610: ref
	{
		if seq.IsComponentPresent(12) {
			ie.Mac_Parameters_v1610 = new(MAC_Parameters_v1610)
			if err := ie.Mac_Parameters_v1610.Decode(d); err != nil {
				return err
			}
		}
	}

	// 15. mcgRLF-RecoveryViaSCG-r16: enumerated
	{
		if seq.IsComponentPresent(13) {
			idx, err := d.DecodeEnumerated(uENRCapabilityV1610McgRLFRecoveryViaSCGR16Constraints)
			if err != nil {
				return err
			}
			ie.McgRLF_RecoveryViaSCG_r16 = &idx
		}
	}

	// 16. resumeWithStoredMCG-SCells-r16: enumerated
	{
		if seq.IsComponentPresent(14) {
			idx, err := d.DecodeEnumerated(uENRCapabilityV1610ResumeWithStoredMCGSCellsR16Constraints)
			if err != nil {
				return err
			}
			ie.ResumeWithStoredMCG_SCells_r16 = &idx
		}
	}

	// 17. resumeWithStoredSCG-r16: enumerated
	{
		if seq.IsComponentPresent(15) {
			idx, err := d.DecodeEnumerated(uENRCapabilityV1610ResumeWithStoredSCGR16Constraints)
			if err != nil {
				return err
			}
			ie.ResumeWithStoredSCG_r16 = &idx
		}
	}

	// 18. resumeWithSCG-Config-r16: enumerated
	{
		if seq.IsComponentPresent(16) {
			idx, err := d.DecodeEnumerated(uENRCapabilityV1610ResumeWithSCGConfigR16Constraints)
			if err != nil {
				return err
			}
			ie.ResumeWithSCG_Config_r16 = &idx
		}
	}

	// 19. ue-BasedPerfMeas-Parameters-r16: ref
	{
		if seq.IsComponentPresent(17) {
			ie.Ue_BasedPerfMeas_Parameters_r16 = new(UE_BasedPerfMeas_Parameters_r16)
			if err := ie.Ue_BasedPerfMeas_Parameters_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 20. son-Parameters-r16: ref
	{
		if seq.IsComponentPresent(18) {
			ie.Son_Parameters_r16 = new(SON_Parameters_r16)
			if err := ie.Son_Parameters_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 21. onDemandSIB-Connected-r16: enumerated
	{
		if seq.IsComponentPresent(19) {
			idx, err := d.DecodeEnumerated(uENRCapabilityV1610OnDemandSIBConnectedR16Constraints)
			if err != nil {
				return err
			}
			ie.OnDemandSIB_Connected_r16 = &idx
		}
	}

	// 22. nonCriticalExtension: ref
	{
		if seq.IsComponentPresent(20) {
			ie.NonCriticalExtension = new(UE_NR_Capability_v1640)
			if err := ie.NonCriticalExtension.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
