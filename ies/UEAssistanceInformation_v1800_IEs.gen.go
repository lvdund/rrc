// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: UEAssistanceInformation-v1800-IEs (line 2475).

var uEAssistanceInformationV1800IEsConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "idc-FDM-Assistance-r18", Optional: true},
		{Name: "idc-TDM-Assistance-r18", Optional: true},
		{Name: "multiRx-PreferenceFR2-r18", Optional: true},
		{Name: "musim-Assistance-v1800", Optional: true},
		{Name: "flightPathInfoAvailable-r18", Optional: true},
		{Name: "ul-TrafficInfo-r18", Optional: true},
		{Name: "n3c-RelayUE-InfoList-r18", Optional: true},
		{Name: "sl-PRS-UE-AssistanceInformationNR-r18", Optional: true},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

const (
	UEAssistanceInformation_v1800_IEs_MultiRx_PreferenceFR2_r18_Single   = 0
	UEAssistanceInformation_v1800_IEs_MultiRx_PreferenceFR2_r18_Multiple = 1
)

var uEAssistanceInformationV1800IEsMultiRxPreferenceFR2R18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UEAssistanceInformation_v1800_IEs_MultiRx_PreferenceFR2_r18_Single, UEAssistanceInformation_v1800_IEs_MultiRx_PreferenceFR2_r18_Multiple},
}

const (
	UEAssistanceInformation_v1800_IEs_FlightPathInfoAvailable_r18_True = 0
)

var uEAssistanceInformationV1800IEsFlightPathInfoAvailableR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UEAssistanceInformation_v1800_IEs_FlightPathInfoAvailable_r18_True},
}

var uEAssistanceInformationV1800IEsN3cRelayUEInfoListR18Constraints = per.SizeRange(0, 8)

type UEAssistanceInformation_v1800_IEs struct {
	Idc_FDM_Assistance_r18                *IDC_FDM_Assistance_r18
	Idc_TDM_Assistance_r18                *IDC_TDM_Assistance_r18
	MultiRx_PreferenceFR2_r18             *int64
	Musim_Assistance_v1800                *MUSIM_Assistance_v1800
	FlightPathInfoAvailable_r18           *int64
	Ul_TrafficInfo_r18                    *UL_TrafficInfo_r18
	N3c_RelayUE_InfoList_r18              []N3C_RelayUE_Info_r18
	Sl_PRS_UE_AssistanceInformationNR_r18 *SL_PRS_UE_AssistanceInformationNR_r18
	NonCriticalExtension                  *UEAssistanceInformation_v1900_IEs
}

func (ie *UEAssistanceInformation_v1800_IEs) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(uEAssistanceInformationV1800IEsConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Idc_FDM_Assistance_r18 != nil, ie.Idc_TDM_Assistance_r18 != nil, ie.MultiRx_PreferenceFR2_r18 != nil, ie.Musim_Assistance_v1800 != nil, ie.FlightPathInfoAvailable_r18 != nil, ie.Ul_TrafficInfo_r18 != nil, ie.N3c_RelayUE_InfoList_r18 != nil, ie.Sl_PRS_UE_AssistanceInformationNR_r18 != nil, ie.NonCriticalExtension != nil}); err != nil {
		return err
	}

	// 2. idc-FDM-Assistance-r18: ref
	{
		if ie.Idc_FDM_Assistance_r18 != nil {
			if err := ie.Idc_FDM_Assistance_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. idc-TDM-Assistance-r18: ref
	{
		if ie.Idc_TDM_Assistance_r18 != nil {
			if err := ie.Idc_TDM_Assistance_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. multiRx-PreferenceFR2-r18: enumerated
	{
		if ie.MultiRx_PreferenceFR2_r18 != nil {
			if err := e.EncodeEnumerated(*ie.MultiRx_PreferenceFR2_r18, uEAssistanceInformationV1800IEsMultiRxPreferenceFR2R18Constraints); err != nil {
				return err
			}
		}
	}

	// 5. musim-Assistance-v1800: ref
	{
		if ie.Musim_Assistance_v1800 != nil {
			if err := ie.Musim_Assistance_v1800.Encode(e); err != nil {
				return err
			}
		}
	}

	// 6. flightPathInfoAvailable-r18: enumerated
	{
		if ie.FlightPathInfoAvailable_r18 != nil {
			if err := e.EncodeEnumerated(*ie.FlightPathInfoAvailable_r18, uEAssistanceInformationV1800IEsFlightPathInfoAvailableR18Constraints); err != nil {
				return err
			}
		}
	}

	// 7. ul-TrafficInfo-r18: ref
	{
		if ie.Ul_TrafficInfo_r18 != nil {
			if err := ie.Ul_TrafficInfo_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	// 8. n3c-RelayUE-InfoList-r18: sequence-of
	{
		if ie.N3c_RelayUE_InfoList_r18 != nil {
			seqOf := e.NewSequenceOfEncoder(uEAssistanceInformationV1800IEsN3cRelayUEInfoListR18Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.N3c_RelayUE_InfoList_r18))); err != nil {
				return err
			}
			for i := range ie.N3c_RelayUE_InfoList_r18 {
				if err := ie.N3c_RelayUE_InfoList_r18[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 9. sl-PRS-UE-AssistanceInformationNR-r18: ref
	{
		if ie.Sl_PRS_UE_AssistanceInformationNR_r18 != nil {
			if err := ie.Sl_PRS_UE_AssistanceInformationNR_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	// 10. nonCriticalExtension: ref
	{
		if ie.NonCriticalExtension != nil {
			if err := ie.NonCriticalExtension.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *UEAssistanceInformation_v1800_IEs) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(uEAssistanceInformationV1800IEsConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. idc-FDM-Assistance-r18: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Idc_FDM_Assistance_r18 = new(IDC_FDM_Assistance_r18)
			if err := ie.Idc_FDM_Assistance_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. idc-TDM-Assistance-r18: ref
	{
		if seq.IsComponentPresent(1) {
			ie.Idc_TDM_Assistance_r18 = new(IDC_TDM_Assistance_r18)
			if err := ie.Idc_TDM_Assistance_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. multiRx-PreferenceFR2-r18: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(uEAssistanceInformationV1800IEsMultiRxPreferenceFR2R18Constraints)
			if err != nil {
				return err
			}
			ie.MultiRx_PreferenceFR2_r18 = &idx
		}
	}

	// 5. musim-Assistance-v1800: ref
	{
		if seq.IsComponentPresent(3) {
			ie.Musim_Assistance_v1800 = new(MUSIM_Assistance_v1800)
			if err := ie.Musim_Assistance_v1800.Decode(d); err != nil {
				return err
			}
		}
	}

	// 6. flightPathInfoAvailable-r18: enumerated
	{
		if seq.IsComponentPresent(4) {
			idx, err := d.DecodeEnumerated(uEAssistanceInformationV1800IEsFlightPathInfoAvailableR18Constraints)
			if err != nil {
				return err
			}
			ie.FlightPathInfoAvailable_r18 = &idx
		}
	}

	// 7. ul-TrafficInfo-r18: ref
	{
		if seq.IsComponentPresent(5) {
			ie.Ul_TrafficInfo_r18 = new(UL_TrafficInfo_r18)
			if err := ie.Ul_TrafficInfo_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	// 8. n3c-RelayUE-InfoList-r18: sequence-of
	{
		if seq.IsComponentPresent(6) {
			seqOf := d.NewSequenceOfDecoder(uEAssistanceInformationV1800IEsN3cRelayUEInfoListR18Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.N3c_RelayUE_InfoList_r18 = make([]N3C_RelayUE_Info_r18, n)
			for i := int64(0); i < n; i++ {
				if err := ie.N3c_RelayUE_InfoList_r18[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 9. sl-PRS-UE-AssistanceInformationNR-r18: ref
	{
		if seq.IsComponentPresent(7) {
			ie.Sl_PRS_UE_AssistanceInformationNR_r18 = new(SL_PRS_UE_AssistanceInformationNR_r18)
			if err := ie.Sl_PRS_UE_AssistanceInformationNR_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	// 10. nonCriticalExtension: ref
	{
		if seq.IsComponentPresent(8) {
			ie.NonCriticalExtension = new(UEAssistanceInformation_v1900_IEs)
			if err := ie.NonCriticalExtension.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
