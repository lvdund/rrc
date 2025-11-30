package rrc

import (
	"fmt"
	"testing"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/ies"
)

func TestRRCSetupRequest(t *testing.T) {
	rrcSetupRequest := ies.RRCSetupRequest{
		RrcSetupRequest: ies.RRCSetupRequest_IEs{
			Ue_Identity: ies.InitialUE_Identity{
				Choice: ies.InitialUE_Identity_Choice_Ng_5G_S_TMSI_Part1,
				Ng_5G_S_TMSI_Part1: aper.BitString{
					Bytes:   []byte{0x07, 0xFF, 0xFF, 0xFF, 0xFF},
					NumBits: 39,
				},
			},
			EstablishmentCause: ies.EstablishmentCause{
				Value: ies.EstablishmentCause_Enum_mo_Signalling,
			},
			Spare: aper.BitString{
				Bytes:   []byte{1},
				NumBits: 1,
			},
		},
	}

	ulccchMessage := &ies.UL_CCCH_Message{
		Message: ies.UL_CCCH_MessageType{
			Choice: ies.UL_CCCH_MessageType_Choice_C1,
			C1: &ies.UL_CCCH_MessageType_C1{
				Choice:          ies.UL_CCCH_MessageType_C1_Choice_RrcSetupRequest,
				RrcSetupRequest: &rrcSetupRequest,
			},
		},
	}

	encoded, err := Encode(ulccchMessage)
	if err != nil {
		fmt.Printf("Failed to encode UL-CCCH-Message: %v", err)
		return
	}
	fmt.Printf("Encoded UL-CCCH-Message: %d\n", encoded)

	decoded := &ies.UL_CCCH_Message{}
	if err := Decode(encoded, decoded); err != nil {
		fmt.Printf("Failed to decode UL-CCCH-Message: %v", err)
		return
	}
	rrcmsg := decoded.Message.C1.RrcSetupRequest
	fmt.Printf("Decoded UL-CCCH-Message: %+v", rrcmsg)
}

func TestRRCSetupComplete(t *testing.T) {
	rrcSetupComplete := ies.RRCSetupComplete{
		Rrc_TransactionIdentifier: ies.RRC_TransactionIdentifier{
			Value: 0,
		},
		CriticalExtensions: ies.RRCSetupComplete_CriticalExtensions{
			Choice: ies.RRCSetupComplete_CriticalExtensions_Choice_RrcSetupComplete,
			RrcSetupComplete: &ies.RRCSetupComplete_IEs{
				SelectedPLMN_Identity: 10,
				RegisteredAMF: &ies.RegisteredAMF{
					Amf_Identifier: ies.AMF_Identifier{Value: aper.BitString{
						Bytes:   []byte{0xFF, 0xFF, 0xFF},
						NumBits: 24,
					}},
				},
				Guami_Type: &ies.RRCSetupComplete_IEs_guami_Type{
					Value: ies.RRCSetupComplete_IEs_guami_Type_Enum_mapped,
				},
				S_NSSAI_List: []ies.S_NSSAI{
					{
						Choice: ies.S_NSSAI_Choice_Sst,
						Sst: aper.BitString{
							Bytes:   []byte{15},
							NumBits: 8,
						},
					},
				},
				DedicatedNAS_Message: ies.DedicatedNAS_Message{
					Value: []byte{0x7e, 0x00, 0x41, 0x79, 0x00, 0x0d},
				},
				Ng_5G_S_TMSI_Value: &ies.RRCSetupComplete_IEs_ng_5G_S_TMSI_Value{
					Choice: ies.RRCSetupComplete_IEs_ng_5G_S_TMSI_Value_Choice_Ng_5G_S_TMSI,
					Ng_5G_S_TMSI: &ies.NG_5G_S_TMSI{
						Value: aper.BitString{
							Bytes:   []byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF},
							NumBits: 48,
						},
					},
				},
				LateNonCriticalExtension: &[]byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0A, 0x0B, 0x0C, 0x0D, 0x0E, 0x0F},
				NonCriticalExtension: &ies.RRCSetupComplete_v1610_IEs{
					MobilityState_r16: &ies.RRCSetupComplete_v1610_IEs_mobilityState_r16{Value: 3},
				},
			},
		},
	}

	uldcchMessage := &ies.UL_DCCH_Message{
		Message: ies.UL_DCCH_MessageType{
			Choice: ies.UL_DCCH_MessageType_Choice_C1,
			C1: &ies.UL_DCCH_MessageType_C1{
				Choice:           ies.UL_DCCH_MessageType_C1_Choice_RrcSetupComplete,
				RrcSetupComplete: &rrcSetupComplete,
			},
		},
	}

	encoded, err := Encode(uldcchMessage)
	if err != nil {
		fmt.Printf("Failed to encode UL-DCCH-Message: %v", err)
		return
	}

	fmt.Printf("Encoded UL-DCCH-Message==============\n %0.8b\n===============================\n\n", encoded)

	// 00000100 00010000 01001000 11000011 11111111 11111111 11111100 00000000 00000000 00000000 00011001 11111000 00000001 00000101 11100100 00000000 00110100
	// nas: 0111111 00000000 01000001 01111001 00000000 00001101

	decoded := &ies.UL_DCCH_Message{}
	if err := Decode(encoded, decoded); err != nil {
		fmt.Printf("Failed to decode UL-DCCH-Message: %v", err)
		return
	}
	rrcmsg := decoded.Message.C1.RrcSetupComplete.CriticalExtensions.RrcSetupComplete

	fmt.Printf("================================\n")
	fmt.Printf("SelectedPLMN_Identity:\t\t%v\n", rrcmsg.SelectedPLMN_Identity)
	fmt.Printf("RegisteredAMF:\t\t\t %+v\n", rrcmsg.RegisteredAMF.Amf_Identifier.Value)
	fmt.Printf("Guami_Type:\t\t\t %+v\n", rrcmsg.Guami_Type)
	fmt.Printf("S_NSSAI_List:\t\t\t %+v\n", rrcmsg.S_NSSAI_List)
	fmt.Printf("DedicatedNAS_Message:\t\t %+v\n", rrcmsg.DedicatedNAS_Message.Value)
	fmt.Printf("Ng_5G_S_TMSI_Value:\t\t %0.8b\n", rrcmsg.Ng_5G_S_TMSI_Value.Ng_5G_S_TMSI.Value)
	fmt.Printf("LateNonCriticalExtension:\t %+v\n", rrcmsg.LateNonCriticalExtension)
	fmt.Printf("NonCriticalExtension -> MobilityState_r16:\t\t %+v\n", rrcmsg.NonCriticalExtension.MobilityState_r16)
}

// TestDecodeAny demonstrates how to decode RRC messages when you don't know the container type
func TestDecodeAny(t *testing.T) {
	// Example: Encode a UL-DCCH message
	uldcchMessage := &ies.UL_DCCH_Message{
		Message: ies.UL_DCCH_MessageType{
			Choice: ies.UL_DCCH_MessageType_Choice_C1,
			C1: &ies.UL_DCCH_MessageType_C1{
				Choice: ies.UL_DCCH_MessageType_C1_Choice_RrcSetupComplete,
				RrcSetupComplete: &ies.RRCSetupComplete{
					Rrc_TransactionIdentifier: ies.RRC_TransactionIdentifier{Value: 0},
					CriticalExtensions: ies.RRCSetupComplete_CriticalExtensions{
						Choice: ies.RRCSetupComplete_CriticalExtensions_Choice_RrcSetupComplete,
						RrcSetupComplete: &ies.RRCSetupComplete_IEs{
							SelectedPLMN_Identity: 1,
							DedicatedNAS_Message: ies.DedicatedNAS_Message{
								Value: []byte{0x7e, 0x00},
							},
						},
					},
				},
			},
		},
	}

	encoded, err := Encode(uldcchMessage)
	if err != nil {
		t.Fatalf("Failed to encode: %v", err)
	}

	// Now decode without knowing the container type
	decoded, err := DecodeAny(encoded)
	if err != nil {
		t.Fatalf("Failed to decode: %v", err)
	}

	fmt.Printf("Detected container type: %s\n", decoded.Type)

	// Type assert to access the specific message
	switch msg := decoded.Message.(type) {
	case *ies.UL_DCCH_Message:
		fmt.Printf("Successfully decoded UL-DCCH message\n")
		if msg.Message.C1 != nil && msg.Message.C1.RrcSetupComplete != nil {
			fmt.Printf("RRC Setup Complete message found\n")
		}
	case *ies.UL_CCCH_Message:
		fmt.Printf("Decoded as UL-CCCH message\n")
	case *ies.DL_DCCH_Message:
		fmt.Printf("Decoded as DL-DCCH message\n")
	case *ies.DL_CCCH_Message:
		fmt.Printf("Decoded as DL-CCCH message\n")
	default:
		fmt.Printf("Decoded as unknown message type\n")
	}
}

func Test_RRCReestablishmentRequest(t *testing.T) {
	rrcReestablishmentRequest := ies.RRCReestablishmentRequest{
		RrcReestablishmentRequest: ies.RRCReestablishmentRequest_IEs{
			Ue_Identity: ies.ReestabUE_Identity{
				C_RNTI:     ies.RNTI_Value{Value: 123},
				PhysCellId: ies.PhysCellId{Value: 234},
				ShortMAC_I: ies.ShortMAC_I{Value: aper.BitString{Bytes: []byte{0xFF, 0xFF}, NumBits: 16}},
			},
			ReestablishmentCause: ies.ReestablishmentCause{Value: 2},
			Spare:                aper.BitString{Bytes: []byte{1}, NumBits: 1},
		},
	}
	ulccchMessage := &ies.UL_CCCH_Message{
		Message: ies.UL_CCCH_MessageType{
			Choice: ies.UL_CCCH_MessageType_Choice_C1,
			C1: &ies.UL_CCCH_MessageType_C1{
				Choice:                    ies.UL_CCCH_MessageType_C1_Choice_RrcReestablishmentRequest,
				RrcReestablishmentRequest: &rrcReestablishmentRequest,
			},
		},
	}

	encoded, err := Encode(ulccchMessage)
	if err != nil {
		fmt.Printf("Failed to encode UL-CCCH-Message: %v", err)
		return
	}
	fmt.Printf("Encoded UL-CCCH-Message: %d\n", encoded)

	decoded := &ies.UL_CCCH_Message{}
	if err := Decode(encoded, decoded); err != nil {
		fmt.Printf("Failed to decode UL-CCCH-Message: %v", err)
		return
	}
	rrcmsg := decoded.Message.C1.RrcReestablishmentRequest.RrcReestablishmentRequest
	fmt.Printf("================================\n")
	fmt.Printf("Ue_Identity:\t\t\t %+v\n", rrcmsg.Ue_Identity.C_RNTI.Value)
	fmt.Printf("PhysCellId:\t\t\t %+v\n", rrcmsg.Ue_Identity.PhysCellId.Value)
	fmt.Printf("ShortMAC_I:\t\t\t %+v\n", rrcmsg.Ue_Identity.ShortMAC_I.Value)
	fmt.Printf("ReestablishmentCause:\t\t %+v\n", rrcmsg.ReestablishmentCause.Value)
	fmt.Printf("Spare:\t\t\t %+v\n", rrcmsg.Spare.Bytes)
}

func Test_UeAssistanceInformation(t *testing.T) {
	msg := ies.UEAssistanceInformation{
		CriticalExtensions: ies.UEAssistanceInformation_CriticalExtensions{
			Choice: 1,
			UeAssistanceInformation: &ies.UEAssistanceInformation_IEs{
				DelayBudgetReport: &ies.DelayBudgetReport{
					Choice: ies.DelayBudgetReport_Choice_Type1,
					Type1: &ies.DelayBudgetReport_type1{
						Value: 10,
					},
				},
				LateNonCriticalExtension: &[]byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0A, 0x0B, 0x0C, 0x0D, 0x0E, 0x0F},
				NonCriticalExtension: &ies.UEAssistanceInformation_v1540_IEs{
					OverheatingAssistance: &ies.OverheatingAssistance{
						ReducedMaxCCs: &ies.ReducedMaxCCs_r16{
							ReducedCCsDL_r16: 11,
							ReducedCCsUL_r16: 12,
						},
						ReducedMaxBW_FR1: &ies.ReducedMaxBW_FRx_r16{
							ReducedBW_DL_r16: ies.ReducedAggregatedBandwidth{Value: 8},
							ReducedBW_UL_r16: ies.ReducedAggregatedBandwidth{Value: 9},
						},
						ReducedMaxBW_FR2: &ies.ReducedMaxBW_FRx_r16{
							ReducedBW_DL_r16: ies.ReducedAggregatedBandwidth{Value: 10},
							ReducedBW_UL_r16: ies.ReducedAggregatedBandwidth{Value: 11},
						},
						ReducedMaxMIMO_LayersFR1: &ies.OverheatingAssistance_reducedMaxMIMO_LayersFR1{
							ReducedMIMO_LayersFR1_DL: ies.MIMO_LayersDL{Value: 1},
							ReducedMIMO_LayersFR1_UL: ies.MIMO_LayersUL{Value: 2},
						},
						ReducedMaxMIMO_LayersFR2: &ies.OverheatingAssistance_reducedMaxMIMO_LayersFR2{
							ReducedMIMO_LayersFR2_DL: ies.MIMO_LayersDL{Value: 1},
							ReducedMIMO_LayersFR2_UL: ies.MIMO_LayersUL{Value: 2},
						},
					},
					NonCriticalExtension: &ies.UEAssistanceInformation_v1610_IEs{
						Idc_Assistance_r16: &ies.IDC_Assistance_r16{
							AffectedCarrierFreqList_r16: &ies.AffectedCarrierFreqList_r16{
								Value: []ies.AffectedCarrierFreq_r16{
									{
										CarrierFreq_r16:           ies.ARFCN_ValueNR{Value: 123},
										InterferenceDirection_r16: ies.AffectedCarrierFreq_r16_interferenceDirection_r16{Value: 3},
									},
								},
							},
							// AffectedCarrierFreqCombList_r16: &ies.AffectedCarrierFreqCombList_r16{
							// 	Value: []ies.AffectedCarrierFreqComb_r16{
							// 		ies.AffectedCarrierFreqComb_r16{
							// 			AffectedCarrierFreqComb_r16: []ies.ARFCN_ValueNR{
							// 				{Value: 123},
							// 				{Value: 124},
							// 			},
							// 			VictimSystemType_r16: ies.VictimSystemType_r16{
							// 				Gps_r16:     &ies.VictimSystemType_r16_gps_r16{Value: 0},
							// 				Glonass_r16: &ies.VictimSystemType_r16_glonass_r16{Value: 0},
							// 			},
							// 		},
							// 	},
							// },
						},
					},
				},
			},
		},
	}
	encoded, err := Encode(&msg)
	if err != nil {
		fmt.Printf("Failed to encode UEAssistanceInformation: %v", err)
		return
	}

	fmt.Println("ENCODE\n", encoded)

	decoded := &ies.UEAssistanceInformation{}
	if err := Decode(encoded, decoded); err != nil {
		fmt.Printf("Failed to decode UEAssistanceInformation: %v", err)
		return
	}

	d := decoded.CriticalExtensions.UeAssistanceInformation
	// Print ALL fields of decoded
	fmt.Printf("================================\n")

	noncri := decoded.CriticalExtensions.UeAssistanceInformation.NonCriticalExtension
	overheat := noncri.OverheatingAssistance
	sub_noncri := noncri.NonCriticalExtension
	// r16 := sub_noncri.Idc_Assistance_r16.AffectedCarrierFreqCombList_r16.Value[0]

	fmt.Println(decoded.CriticalExtensions.Choice)
	fmt.Println("type1", d.DelayBudgetReport.Type1)
	fmt.Println("lateNonCriticalExtension", d.LateNonCriticalExtension)
	fmt.Println("reducedMaxCCs", overheat.ReducedMaxCCs)
	fmt.Println("reducedMaxBW_FR1", overheat.ReducedMaxBW_FR1)
	fmt.Println("reducedMaxBW_FR2", overheat.ReducedMaxBW_FR2)
	fmt.Println("reducedMaxMIMO_LayersFR1", overheat.ReducedMaxMIMO_LayersFR1)
	fmt.Println("reducedMaxMIMO_LayersFR2", overheat.ReducedMaxMIMO_LayersFR2)
	fmt.Println("affectedCarrierFreqList_r16", sub_noncri.Idc_Assistance_r16.AffectedCarrierFreqList_r16.Value[0].CarrierFreq_r16)
	fmt.Println("affectedCarrierFreqList_r16", sub_noncri.Idc_Assistance_r16.AffectedCarrierFreqList_r16.Value[0].InterferenceDirection_r16)
	// fmt.Println("affectedCarrierFreqCombList_r16[0]", r16.AffectedCarrierFreqComb_r16)
	// fmt.Println("affectedCarrierFreqCombList_r16[0]", r16.VictimSystemType_r16.Gps_r16, r16.VictimSystemType_r16.Glonass_r16)
	fmt.Printf("================================\n")
}
