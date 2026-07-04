// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: CandidateTCI-State-r18 (line 5464).

var candidateTCIStateR18Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "tci-StateId-r18"},
		{Name: "qcl-Type1-r18"},
		{Name: "qcl-Type2-r18", Optional: true},
		{Name: "pathlossReferenceRS-Id-r18", Optional: true},
		{Name: "tag-Id-ptr-r18", Optional: true},
		{Name: "ul-PowerControl-r18", Optional: true},
	},
}

const (
	CandidateTCI_State_r18_Tag_Id_Ptr_r18_N0 = 0
	CandidateTCI_State_r18_Tag_Id_Ptr_r18_N1 = 1
)

var candidateTCIStateR18TagIdPtrR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CandidateTCI_State_r18_Tag_Id_Ptr_r18_N0, CandidateTCI_State_r18_Tag_Id_Ptr_r18_N1},
}

type CandidateTCI_State_r18 struct {
	Tci_StateId_r18            TCI_StateId
	Qcl_Type1_r18              LTM_QCL_Info_r18
	Qcl_Type2_r18              *LTM_QCL_Info_r18
	PathlossReferenceRS_Id_r18 *PathlossReferenceRS_Id_r17
	Tag_Id_Ptr_r18             *int64
	Ul_PowerControl_r18        *Uplink_PowerControlId_r17
}

func (ie *CandidateTCI_State_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(candidateTCIStateR18Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Qcl_Type2_r18 != nil, ie.PathlossReferenceRS_Id_r18 != nil, ie.Tag_Id_Ptr_r18 != nil, ie.Ul_PowerControl_r18 != nil}); err != nil {
		return err
	}

	// 3. tci-StateId-r18: ref
	{
		if err := ie.Tci_StateId_r18.Encode(e); err != nil {
			return err
		}
	}

	// 4. qcl-Type1-r18: ref
	{
		if err := ie.Qcl_Type1_r18.Encode(e); err != nil {
			return err
		}
	}

	// 5. qcl-Type2-r18: ref
	{
		if ie.Qcl_Type2_r18 != nil {
			if err := ie.Qcl_Type2_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	// 6. pathlossReferenceRS-Id-r18: ref
	{
		if ie.PathlossReferenceRS_Id_r18 != nil {
			if err := ie.PathlossReferenceRS_Id_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	// 7. tag-Id-ptr-r18: enumerated
	{
		if ie.Tag_Id_Ptr_r18 != nil {
			if err := e.EncodeEnumerated(*ie.Tag_Id_Ptr_r18, candidateTCIStateR18TagIdPtrR18Constraints); err != nil {
				return err
			}
		}
	}

	// 8. ul-PowerControl-r18: ref
	{
		if ie.Ul_PowerControl_r18 != nil {
			if err := ie.Ul_PowerControl_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *CandidateTCI_State_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(candidateTCIStateR18Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. tci-StateId-r18: ref
	{
		if err := ie.Tci_StateId_r18.Decode(d); err != nil {
			return err
		}
	}

	// 4. qcl-Type1-r18: ref
	{
		if err := ie.Qcl_Type1_r18.Decode(d); err != nil {
			return err
		}
	}

	// 5. qcl-Type2-r18: ref
	{
		if seq.IsComponentPresent(2) {
			ie.Qcl_Type2_r18 = new(LTM_QCL_Info_r18)
			if err := ie.Qcl_Type2_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	// 6. pathlossReferenceRS-Id-r18: ref
	{
		if seq.IsComponentPresent(3) {
			ie.PathlossReferenceRS_Id_r18 = new(PathlossReferenceRS_Id_r17)
			if err := ie.PathlossReferenceRS_Id_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	// 7. tag-Id-ptr-r18: enumerated
	{
		if seq.IsComponentPresent(4) {
			idx, err := d.DecodeEnumerated(candidateTCIStateR18TagIdPtrR18Constraints)
			if err != nil {
				return err
			}
			ie.Tag_Id_Ptr_r18 = &idx
		}
	}

	// 8. ul-PowerControl-r18: ref
	{
		if seq.IsComponentPresent(5) {
			ie.Ul_PowerControl_r18 = new(Uplink_PowerControlId_r17)
			if err := ie.Ul_PowerControl_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
