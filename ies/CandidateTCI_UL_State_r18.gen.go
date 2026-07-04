// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: CandidateTCI-UL-State-r18 (line 5486).

var candidateTCIULStateR18Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "tci-UL-StateId-r18"},
		{Name: "referenceSignal-r18"},
		{Name: "pathlossReferenceRS-Id-r18", Optional: true},
		{Name: "tag-Id-ptr-r18", Optional: true},
		{Name: "ul-PowerControl-r18", Optional: true},
	},
}

var candidateTCI_UL_State_r18ReferenceSignalR18Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "ssb-Index"},
		{Name: "csi-RS-Index"},
	},
}

const (
	CandidateTCI_UL_State_r18_ReferenceSignal_r18_Ssb_Index    = 0
	CandidateTCI_UL_State_r18_ReferenceSignal_r18_Csi_RS_Index = 1
)

const (
	CandidateTCI_UL_State_r18_Tag_Id_Ptr_r18_N0 = 0
	CandidateTCI_UL_State_r18_Tag_Id_Ptr_r18_N1 = 1
)

var candidateTCIULStateR18TagIdPtrR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CandidateTCI_UL_State_r18_Tag_Id_Ptr_r18_N0, CandidateTCI_UL_State_r18_Tag_Id_Ptr_r18_N1},
}

type CandidateTCI_UL_State_r18 struct {
	Tci_UL_StateId_r18  TCI_UL_StateId_r17
	ReferenceSignal_r18 struct {
		Choice       int
		Ssb_Index    *SSB_Index
		Csi_RS_Index *NZP_CSI_RS_ResourceId
	}
	PathlossReferenceRS_Id_r18 *PathlossReferenceRS_Id_r17
	Tag_Id_Ptr_r18             *int64
	Ul_PowerControl_r18        *Uplink_PowerControlId_r17
}

func (ie *CandidateTCI_UL_State_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(candidateTCIULStateR18Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.PathlossReferenceRS_Id_r18 != nil, ie.Tag_Id_Ptr_r18 != nil, ie.Ul_PowerControl_r18 != nil}); err != nil {
		return err
	}

	// 3. tci-UL-StateId-r18: ref
	{
		if err := ie.Tci_UL_StateId_r18.Encode(e); err != nil {
			return err
		}
	}

	// 4. referenceSignal-r18: choice
	{
		choiceEnc := e.NewChoiceEncoder(candidateTCI_UL_State_r18ReferenceSignalR18Constraints)
		if err := choiceEnc.EncodeChoice(int64(ie.ReferenceSignal_r18.Choice), false, nil); err != nil {
			return err
		}
		switch ie.ReferenceSignal_r18.Choice {
		case CandidateTCI_UL_State_r18_ReferenceSignal_r18_Ssb_Index:
			if err := ie.ReferenceSignal_r18.Ssb_Index.Encode(e); err != nil {
				return err
			}
		case CandidateTCI_UL_State_r18_ReferenceSignal_r18_Csi_RS_Index:
			if err := ie.ReferenceSignal_r18.Csi_RS_Index.Encode(e); err != nil {
				return err
			}
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.ReferenceSignal_r18.Choice), Constraint: "index out of range"}
		}
	}

	// 5. pathlossReferenceRS-Id-r18: ref
	{
		if ie.PathlossReferenceRS_Id_r18 != nil {
			if err := ie.PathlossReferenceRS_Id_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	// 6. tag-Id-ptr-r18: enumerated
	{
		if ie.Tag_Id_Ptr_r18 != nil {
			if err := e.EncodeEnumerated(*ie.Tag_Id_Ptr_r18, candidateTCIULStateR18TagIdPtrR18Constraints); err != nil {
				return err
			}
		}
	}

	// 7. ul-PowerControl-r18: ref
	{
		if ie.Ul_PowerControl_r18 != nil {
			if err := ie.Ul_PowerControl_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *CandidateTCI_UL_State_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(candidateTCIULStateR18Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. tci-UL-StateId-r18: ref
	{
		if err := ie.Tci_UL_StateId_r18.Decode(d); err != nil {
			return err
		}
	}

	// 4. referenceSignal-r18: choice
	{
		choiceDec := d.NewChoiceDecoder(candidateTCI_UL_State_r18ReferenceSignalR18Constraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.ReferenceSignal_r18.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case CandidateTCI_UL_State_r18_ReferenceSignal_r18_Ssb_Index:
			ie.ReferenceSignal_r18.Ssb_Index = new(SSB_Index)
			if err := ie.ReferenceSignal_r18.Ssb_Index.Decode(d); err != nil {
				return err
			}
		case CandidateTCI_UL_State_r18_ReferenceSignal_r18_Csi_RS_Index:
			ie.ReferenceSignal_r18.Csi_RS_Index = new(NZP_CSI_RS_ResourceId)
			if err := ie.ReferenceSignal_r18.Csi_RS_Index.Decode(d); err != nil {
				return err
			}
		}
	}

	// 5. pathlossReferenceRS-Id-r18: ref
	{
		if seq.IsComponentPresent(2) {
			ie.PathlossReferenceRS_Id_r18 = new(PathlossReferenceRS_Id_r17)
			if err := ie.PathlossReferenceRS_Id_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	// 6. tag-Id-ptr-r18: enumerated
	{
		if seq.IsComponentPresent(3) {
			idx, err := d.DecodeEnumerated(candidateTCIULStateR18TagIdPtrR18Constraints)
			if err != nil {
				return err
			}
			ie.Tag_Id_Ptr_r18 = &idx
		}
	}

	// 7. ul-PowerControl-r18: ref
	{
		if seq.IsComponentPresent(4) {
			ie.Ul_PowerControl_r18 = new(Uplink_PowerControlId_r17)
			if err := ie.Ul_PowerControl_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
