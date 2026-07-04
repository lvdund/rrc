// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: UE-NR-Capability-v1900 (line 25885).

var uENRCapabilityV1900Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "aiml-Parameters-r19", Optional: true},
		{Name: "ue-RadioPagingInfo-r19", Optional: true},
		{Name: "ntn-VSAT-AntennaTypeKuBand-r19", Optional: true},
		{Name: "ntn-VSAT-MobilityTypeKuBand-r19", Optional: true},
		{Name: "ntn-ERedCap-FR1-r19", Optional: true},
		{Name: "onDemandSIB1-r19", Optional: true},
		{Name: "onDemandPosSIB-ConnectedCtrlParam-r19", Optional: true},
		{Name: "ntn-Redirection-r19", Optional: true},
		{Name: "drx-PreferenceCellDTX-DRX-r19", Optional: true},
		{Name: "lpwus-SupportedBandsIdleInactiveOFDM-r19", Optional: true},
		{Name: "lpwus-SupportedBandsIdleInactiveOOK-r19", Optional: true},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

var uENRCapabilityV1900UeRadioPagingInfoR19Constraints = per.SizeConstraints{}

const (
	UE_NR_Capability_v1900_Ntn_VSAT_AntennaTypeKuBand_r19_Electronic = 0
	UE_NR_Capability_v1900_Ntn_VSAT_AntennaTypeKuBand_r19_Mechanical = 1
)

var uENRCapabilityV1900NtnVSATAntennaTypeKuBandR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UE_NR_Capability_v1900_Ntn_VSAT_AntennaTypeKuBand_r19_Electronic, UE_NR_Capability_v1900_Ntn_VSAT_AntennaTypeKuBand_r19_Mechanical},
}

const (
	UE_NR_Capability_v1900_Ntn_VSAT_MobilityTypeKuBand_r19_Fixed  = 0
	UE_NR_Capability_v1900_Ntn_VSAT_MobilityTypeKuBand_r19_Mobile = 1
)

var uENRCapabilityV1900NtnVSATMobilityTypeKuBandR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UE_NR_Capability_v1900_Ntn_VSAT_MobilityTypeKuBand_r19_Fixed, UE_NR_Capability_v1900_Ntn_VSAT_MobilityTypeKuBand_r19_Mobile},
}

const (
	UE_NR_Capability_v1900_Ntn_ERedCap_FR1_r19_Supported = 0
)

var uENRCapabilityV1900NtnERedCapFR1R19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UE_NR_Capability_v1900_Ntn_ERedCap_FR1_r19_Supported},
}

const (
	UE_NR_Capability_v1900_OnDemandSIB1_r19_Supported = 0
)

var uENRCapabilityV1900OnDemandSIB1R19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UE_NR_Capability_v1900_OnDemandSIB1_r19_Supported},
}

const (
	UE_NR_Capability_v1900_OnDemandPosSIB_ConnectedCtrlParam_r19_Supported = 0
)

var uENRCapabilityV1900OnDemandPosSIBConnectedCtrlParamR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UE_NR_Capability_v1900_OnDemandPosSIB_ConnectedCtrlParam_r19_Supported},
}

const (
	UE_NR_Capability_v1900_Ntn_Redirection_r19_Supported = 0
)

var uENRCapabilityV1900NtnRedirectionR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UE_NR_Capability_v1900_Ntn_Redirection_r19_Supported},
}

const (
	UE_NR_Capability_v1900_Drx_PreferenceCellDTX_DRX_r19_Supported = 0
)

var uENRCapabilityV1900DrxPreferenceCellDTXDRXR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UE_NR_Capability_v1900_Drx_PreferenceCellDTX_DRX_r19_Supported},
}

var uENRCapabilityV1900LpwusSupportedBandsIdleInactiveOFDMR19Constraints = per.SizeRange(1, common.MaxBands)

var uENRCapabilityV1900LpwusSupportedBandsIdleInactiveOOKR19Constraints = per.SizeRange(1, common.MaxBands)

type UE_NR_Capability_v1900 struct {
	Aiml_Parameters_r19                      *AIML_Parameters_r19
	Ue_RadioPagingInfo_r19                   []byte
	Ntn_VSAT_AntennaTypeKuBand_r19           *int64
	Ntn_VSAT_MobilityTypeKuBand_r19          *int64
	Ntn_ERedCap_FR1_r19                      *int64
	OnDemandSIB1_r19                         *int64
	OnDemandPosSIB_ConnectedCtrlParam_r19    *int64
	Ntn_Redirection_r19                      *int64
	Drx_PreferenceCellDTX_DRX_r19            *int64
	Lpwus_SupportedBandsIdleInactiveOFDM_r19 []FreqBandIndicatorNR
	Lpwus_SupportedBandsIdleInactiveOOK_r19  []FreqBandIndicatorNR
	NonCriticalExtension                     *UE_NR_Capability_v1920
}

func (ie *UE_NR_Capability_v1900) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(uENRCapabilityV1900Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Aiml_Parameters_r19 != nil, ie.Ue_RadioPagingInfo_r19 != nil, ie.Ntn_VSAT_AntennaTypeKuBand_r19 != nil, ie.Ntn_VSAT_MobilityTypeKuBand_r19 != nil, ie.Ntn_ERedCap_FR1_r19 != nil, ie.OnDemandSIB1_r19 != nil, ie.OnDemandPosSIB_ConnectedCtrlParam_r19 != nil, ie.Ntn_Redirection_r19 != nil, ie.Drx_PreferenceCellDTX_DRX_r19 != nil, ie.Lpwus_SupportedBandsIdleInactiveOFDM_r19 != nil, ie.Lpwus_SupportedBandsIdleInactiveOOK_r19 != nil, ie.NonCriticalExtension != nil}); err != nil {
		return err
	}

	// 2. aiml-Parameters-r19: ref
	{
		if ie.Aiml_Parameters_r19 != nil {
			if err := ie.Aiml_Parameters_r19.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. ue-RadioPagingInfo-r19: octet-string
	{
		if ie.Ue_RadioPagingInfo_r19 != nil {
			if err := e.EncodeOctetString(ie.Ue_RadioPagingInfo_r19, uENRCapabilityV1900UeRadioPagingInfoR19Constraints); err != nil {
				return err
			}
		}
	}

	// 4. ntn-VSAT-AntennaTypeKuBand-r19: enumerated
	{
		if ie.Ntn_VSAT_AntennaTypeKuBand_r19 != nil {
			if err := e.EncodeEnumerated(*ie.Ntn_VSAT_AntennaTypeKuBand_r19, uENRCapabilityV1900NtnVSATAntennaTypeKuBandR19Constraints); err != nil {
				return err
			}
		}
	}

	// 5. ntn-VSAT-MobilityTypeKuBand-r19: enumerated
	{
		if ie.Ntn_VSAT_MobilityTypeKuBand_r19 != nil {
			if err := e.EncodeEnumerated(*ie.Ntn_VSAT_MobilityTypeKuBand_r19, uENRCapabilityV1900NtnVSATMobilityTypeKuBandR19Constraints); err != nil {
				return err
			}
		}
	}

	// 6. ntn-ERedCap-FR1-r19: enumerated
	{
		if ie.Ntn_ERedCap_FR1_r19 != nil {
			if err := e.EncodeEnumerated(*ie.Ntn_ERedCap_FR1_r19, uENRCapabilityV1900NtnERedCapFR1R19Constraints); err != nil {
				return err
			}
		}
	}

	// 7. onDemandSIB1-r19: enumerated
	{
		if ie.OnDemandSIB1_r19 != nil {
			if err := e.EncodeEnumerated(*ie.OnDemandSIB1_r19, uENRCapabilityV1900OnDemandSIB1R19Constraints); err != nil {
				return err
			}
		}
	}

	// 8. onDemandPosSIB-ConnectedCtrlParam-r19: enumerated
	{
		if ie.OnDemandPosSIB_ConnectedCtrlParam_r19 != nil {
			if err := e.EncodeEnumerated(*ie.OnDemandPosSIB_ConnectedCtrlParam_r19, uENRCapabilityV1900OnDemandPosSIBConnectedCtrlParamR19Constraints); err != nil {
				return err
			}
		}
	}

	// 9. ntn-Redirection-r19: enumerated
	{
		if ie.Ntn_Redirection_r19 != nil {
			if err := e.EncodeEnumerated(*ie.Ntn_Redirection_r19, uENRCapabilityV1900NtnRedirectionR19Constraints); err != nil {
				return err
			}
		}
	}

	// 10. drx-PreferenceCellDTX-DRX-r19: enumerated
	{
		if ie.Drx_PreferenceCellDTX_DRX_r19 != nil {
			if err := e.EncodeEnumerated(*ie.Drx_PreferenceCellDTX_DRX_r19, uENRCapabilityV1900DrxPreferenceCellDTXDRXR19Constraints); err != nil {
				return err
			}
		}
	}

	// 11. lpwus-SupportedBandsIdleInactiveOFDM-r19: sequence-of
	{
		if ie.Lpwus_SupportedBandsIdleInactiveOFDM_r19 != nil {
			seqOf := e.NewSequenceOfEncoder(uENRCapabilityV1900LpwusSupportedBandsIdleInactiveOFDMR19Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Lpwus_SupportedBandsIdleInactiveOFDM_r19))); err != nil {
				return err
			}
			for i := range ie.Lpwus_SupportedBandsIdleInactiveOFDM_r19 {
				if err := ie.Lpwus_SupportedBandsIdleInactiveOFDM_r19[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 12. lpwus-SupportedBandsIdleInactiveOOK-r19: sequence-of
	{
		if ie.Lpwus_SupportedBandsIdleInactiveOOK_r19 != nil {
			seqOf := e.NewSequenceOfEncoder(uENRCapabilityV1900LpwusSupportedBandsIdleInactiveOOKR19Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Lpwus_SupportedBandsIdleInactiveOOK_r19))); err != nil {
				return err
			}
			for i := range ie.Lpwus_SupportedBandsIdleInactiveOOK_r19 {
				if err := ie.Lpwus_SupportedBandsIdleInactiveOOK_r19[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 13. nonCriticalExtension: ref
	{
		if ie.NonCriticalExtension != nil {
			if err := ie.NonCriticalExtension.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *UE_NR_Capability_v1900) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(uENRCapabilityV1900Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. aiml-Parameters-r19: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Aiml_Parameters_r19 = new(AIML_Parameters_r19)
			if err := ie.Aiml_Parameters_r19.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. ue-RadioPagingInfo-r19: octet-string
	{
		if seq.IsComponentPresent(1) {
			v, err := d.DecodeOctetString(uENRCapabilityV1900UeRadioPagingInfoR19Constraints)
			if err != nil {
				return err
			}
			ie.Ue_RadioPagingInfo_r19 = v
		}
	}

	// 4. ntn-VSAT-AntennaTypeKuBand-r19: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(uENRCapabilityV1900NtnVSATAntennaTypeKuBandR19Constraints)
			if err != nil {
				return err
			}
			ie.Ntn_VSAT_AntennaTypeKuBand_r19 = &idx
		}
	}

	// 5. ntn-VSAT-MobilityTypeKuBand-r19: enumerated
	{
		if seq.IsComponentPresent(3) {
			idx, err := d.DecodeEnumerated(uENRCapabilityV1900NtnVSATMobilityTypeKuBandR19Constraints)
			if err != nil {
				return err
			}
			ie.Ntn_VSAT_MobilityTypeKuBand_r19 = &idx
		}
	}

	// 6. ntn-ERedCap-FR1-r19: enumerated
	{
		if seq.IsComponentPresent(4) {
			idx, err := d.DecodeEnumerated(uENRCapabilityV1900NtnERedCapFR1R19Constraints)
			if err != nil {
				return err
			}
			ie.Ntn_ERedCap_FR1_r19 = &idx
		}
	}

	// 7. onDemandSIB1-r19: enumerated
	{
		if seq.IsComponentPresent(5) {
			idx, err := d.DecodeEnumerated(uENRCapabilityV1900OnDemandSIB1R19Constraints)
			if err != nil {
				return err
			}
			ie.OnDemandSIB1_r19 = &idx
		}
	}

	// 8. onDemandPosSIB-ConnectedCtrlParam-r19: enumerated
	{
		if seq.IsComponentPresent(6) {
			idx, err := d.DecodeEnumerated(uENRCapabilityV1900OnDemandPosSIBConnectedCtrlParamR19Constraints)
			if err != nil {
				return err
			}
			ie.OnDemandPosSIB_ConnectedCtrlParam_r19 = &idx
		}
	}

	// 9. ntn-Redirection-r19: enumerated
	{
		if seq.IsComponentPresent(7) {
			idx, err := d.DecodeEnumerated(uENRCapabilityV1900NtnRedirectionR19Constraints)
			if err != nil {
				return err
			}
			ie.Ntn_Redirection_r19 = &idx
		}
	}

	// 10. drx-PreferenceCellDTX-DRX-r19: enumerated
	{
		if seq.IsComponentPresent(8) {
			idx, err := d.DecodeEnumerated(uENRCapabilityV1900DrxPreferenceCellDTXDRXR19Constraints)
			if err != nil {
				return err
			}
			ie.Drx_PreferenceCellDTX_DRX_r19 = &idx
		}
	}

	// 11. lpwus-SupportedBandsIdleInactiveOFDM-r19: sequence-of
	{
		if seq.IsComponentPresent(9) {
			seqOf := d.NewSequenceOfDecoder(uENRCapabilityV1900LpwusSupportedBandsIdleInactiveOFDMR19Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Lpwus_SupportedBandsIdleInactiveOFDM_r19 = make([]FreqBandIndicatorNR, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Lpwus_SupportedBandsIdleInactiveOFDM_r19[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 12. lpwus-SupportedBandsIdleInactiveOOK-r19: sequence-of
	{
		if seq.IsComponentPresent(10) {
			seqOf := d.NewSequenceOfDecoder(uENRCapabilityV1900LpwusSupportedBandsIdleInactiveOOKR19Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Lpwus_SupportedBandsIdleInactiveOOK_r19 = make([]FreqBandIndicatorNR, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Lpwus_SupportedBandsIdleInactiveOOK_r19[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 13. nonCriticalExtension: ref
	{
		if seq.IsComponentPresent(11) {
			ie.NonCriticalExtension = new(UE_NR_Capability_v1920)
			if err := ie.NonCriticalExtension.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
