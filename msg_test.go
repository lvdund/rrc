package rrc

import (
	"testing"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/ies"
)

func TestDecodeAny_DL_CCCH(t *testing.T) {
	rrcmsg := ies.DL_CCCH_Message{
		Message: ies.DL_CCCH_MessageType{
			Choice: ies.DL_CCCH_MessageType_Choice_C1,
			C1: &ies.DL_CCCH_MessageType_C1{
				Choice: ies.DL_CCCH_MessageType_C1_Choice_RrcSetup,
				RrcSetup: &ies.RRCSetup{
					Rrc_TransactionIdentifier: ies.RRC_TransactionIdentifier{
						Value: 0,
					},
					CriticalExtensions: ies.RRCSetup_CriticalExtensions{
						Choice: ies.RRCSetup_CriticalExtensions_Choice_RrcSetup,
						RrcSetup: &ies.RRCSetup_IEs{
							RadioBearerConfig: ies.RadioBearerConfig{
								Srb_ToAddModList: &ies.SRB_ToAddModList{
									Value: []ies.SRB_ToAddMod{ies.SRB_ToAddMod{
										Srb_Identity: ies.SRB_Identity{
											Value: 1,
										},
									}},
								},
							},
							MasterCellGroup: []byte{0, 0},
						},
					},
				},
			},
		},
	}

	encoded, err := Encode(&rrcmsg)
	if err != nil {
		t.Fatalf("Failed to encode DL-CCCH message: %v", err)
	}

	decoded, err := DecodeAny(encoded)
	if err != nil {
		t.Fatalf("Failed to decode: %v", err)
	}

	if decoded.Type != MessageContainerTypeDL_CCCH {
		t.Errorf("Expected DL-CCCH, got %s", decoded.Type)
	}

	if _, ok := decoded.Message.(*ies.DL_CCCH_Message); !ok {
		t.Errorf("Expected *ies.DL_CCCH_Message, got %T", decoded.Message)
	}
}

func TestDecodeAny_UL_DCCH(t *testing.T) {
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
		t.Fatalf("Failed to encode UL-DCCH message: %v", err)
	}

	decoded, err := DecodeAny(encoded)
	if err != nil {
		t.Fatalf("Failed to decode: %v", err)
	}

	if decoded.Type != MessageContainerTypeUL_DCCH {
		t.Errorf("Expected UL-DCCH, got %s", decoded.Type)
	}

	if _, ok := decoded.Message.(*ies.UL_DCCH_Message); !ok {
		t.Errorf("Expected *ies.UL_DCCH_Message, got %T", decoded.Message)
	}
}

func TestDecodeAny_UL_CCCH(t *testing.T) {
	ulccchMessage := &ies.UL_CCCH_Message{
		Message: ies.UL_CCCH_MessageType{
			Choice: ies.UL_CCCH_MessageType_Choice_C1,
			C1: &ies.UL_CCCH_MessageType_C1{
				Choice: ies.UL_CCCH_MessageType_C1_Choice_RrcSetupRequest,
				RrcSetupRequest: &ies.RRCSetupRequest{
					RrcSetupRequest: ies.RRCSetupRequest_IEs{
						Ue_Identity: ies.InitialUE_Identity{
							Choice: ies.InitialUE_Identity_Choice_RandomValue,
							RandomValue: aper.BitString{
								Bytes:   []byte{0x12, 0x34, 0x56, 0x78, 0x9A},
								NumBits: 39,
							},
						},
						EstablishmentCause: ies.EstablishmentCause{
							Value: ies.EstablishmentCause_Enum_mo_Data,
						},
						Spare: aper.BitString{
							Bytes:   []byte{1},
							NumBits: 1,
						},
					},
				},
			},
		},
	}

	encoded, err := Encode(ulccchMessage)
	if err != nil {
		t.Fatalf("Failed to encode UL-CCCH message: %v", err)
	}

	decoded, err := DecodeAny(encoded)
	if err != nil {
		t.Fatalf("Failed to decode: %v", err)
	}

	if decoded.Type != MessageContainerTypeUL_CCCH {
		t.Errorf("Expected UL-CCCH, got %s", decoded.Type)
	}

	if _, ok := decoded.Message.(*ies.UL_CCCH_Message); !ok {
		t.Errorf("Expected *ies.UL_CCCH_Message, got %T", decoded.Message)
	}
}

func TestDecodeAny_DL_DCCH(t *testing.T) {
	dldcchMessage := &ies.DL_DCCH_Message{
		Message: ies.DL_DCCH_MessageType{
			Choice: ies.DL_DCCH_MessageType_Choice_C1,
			C1: &ies.DL_DCCH_MessageType_C1{
				Choice: ies.DL_DCCH_MessageType_C1_Choice_RrcRelease,
				RrcRelease: &ies.RRCRelease{
					Rrc_TransactionIdentifier: ies.RRC_TransactionIdentifier{Value: 0},
					CriticalExtensions: ies.RRCRelease_CriticalExtensions{
						Choice:     ies.RRCRelease_CriticalExtensions_Choice_RrcRelease,
						RrcRelease: &ies.RRCRelease_IEs{},
					},
				},
			},
		},
	}

	encoded, err := Encode(dldcchMessage)
	if err != nil {
		t.Fatalf("Failed to encode DL-DCCH message: %v", err)
	}

	decoded, err := DecodeAny(encoded)
	if err != nil {
		t.Fatalf("Failed to decode: %v", err)
	}

	if decoded.Type != MessageContainerTypeDL_DCCH {
		t.Errorf("Expected DL-DCCH, got %s", decoded.Type)
	}

	if _, ok := decoded.Message.(*ies.DL_DCCH_Message); !ok {
		t.Errorf("Expected *ies.DL_DCCH_Message, got %T", decoded.Message)
	}
}

func TestDecodeAny_EmptyBuffer(t *testing.T) {
	_, err := DecodeAny([]byte{})
	if err == nil {
		t.Error("Expected error for empty buffer, got nil")
	}
}

func TestDecodeAny_InvalidBuffer(t *testing.T) {
	invalidBytes := []byte{0xFF, 0xFF, 0xFF, 0xFF}
	_, err := DecodeAny(invalidBytes)
	if err == nil {
		t.Error("Expected error for invalid buffer, got nil")
	}
}

func TestDecodeAny_RoundTrip(t *testing.T) {
	tests := []struct {
		name    string
		message RRCMessage
		want    MessageContainerType
	}{
		{
			name: "DL-CCCH RRCSetup",
			message: &ies.DL_CCCH_Message{
				Message: ies.DL_CCCH_MessageType{
					Choice: ies.DL_CCCH_MessageType_Choice_C1,
					C1: &ies.DL_CCCH_MessageType_C1{
						Choice: ies.DL_CCCH_MessageType_C1_Choice_RrcSetup,
						RrcSetup: &ies.RRCSetup{
							Rrc_TransactionIdentifier: ies.RRC_TransactionIdentifier{Value: 0},
							CriticalExtensions: ies.RRCSetup_CriticalExtensions{
								Choice: ies.RRCSetup_CriticalExtensions_Choice_RrcSetup,
								RrcSetup: &ies.RRCSetup_IEs{
									MasterCellGroup: []byte{0, 0},
								},
							},
						},
					},
				},
			},
			want: MessageContainerTypeDL_CCCH,
		},
		{
			name: "UL-DCCH RRCSetupComplete",
			message: &ies.UL_DCCH_Message{
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
			},
			want: MessageContainerTypeUL_DCCH,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			encoded, err := Encode(tt.message)
			if err != nil {
				t.Fatalf("Failed to encode: %v", err)
			}

			decoded, err := DecodeAny(encoded)
			if err != nil {
				t.Fatalf("Failed to decode: %v", err)
			}

			if decoded.Type != tt.want {
				t.Errorf("Expected %s, got %s", tt.want, decoded.Type)
			}

			// Verify round-trip: re-encode and compare
			reencoded, err := Encode(decoded.Message)
			if err != nil {
				t.Fatalf("Failed to re-encode: %v", err)
			}

			if len(reencoded) != len(encoded) {
				t.Errorf("Re-encoded length mismatch: got %d, want %d", len(reencoded), len(encoded))
			}
		})
	}
}
