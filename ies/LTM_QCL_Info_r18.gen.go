// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: LTM-QCL-Info-r18 (line 5474).

var lTMQCLInfoR18Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "referenceSignal-r18"},
		{Name: "qcl-Type-r18"},
	},
}

var lTM_QCL_Info_r18ReferenceSignalR18Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "ssb-Index"},
		{Name: "csi-RS-Index"},
	},
}

const (
	LTM_QCL_Info_r18_ReferenceSignal_r18_Ssb_Index    = 0
	LTM_QCL_Info_r18_ReferenceSignal_r18_Csi_RS_Index = 1
)

const (
	LTM_QCL_Info_r18_Qcl_Type_r18_TypeA = 0
	LTM_QCL_Info_r18_Qcl_Type_r18_TypeB = 1
	LTM_QCL_Info_r18_Qcl_Type_r18_TypeC = 2
	LTM_QCL_Info_r18_Qcl_Type_r18_TypeD = 3
)

var lTMQCLInfoR18QclTypeR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{LTM_QCL_Info_r18_Qcl_Type_r18_TypeA, LTM_QCL_Info_r18_Qcl_Type_r18_TypeB, LTM_QCL_Info_r18_Qcl_Type_r18_TypeC, LTM_QCL_Info_r18_Qcl_Type_r18_TypeD},
}

type LTM_QCL_Info_r18 struct {
	ReferenceSignal_r18 struct {
		Choice       int
		Ssb_Index    *SSB_Index
		Csi_RS_Index *NZP_CSI_RS_ResourceId
	}
	Qcl_Type_r18 int64
}

func (ie *LTM_QCL_Info_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(lTMQCLInfoR18Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. referenceSignal-r18: choice
	{
		choiceEnc := e.NewChoiceEncoder(lTM_QCL_Info_r18ReferenceSignalR18Constraints)
		if err := choiceEnc.EncodeChoice(int64(ie.ReferenceSignal_r18.Choice), false, nil); err != nil {
			return err
		}
		switch ie.ReferenceSignal_r18.Choice {
		case LTM_QCL_Info_r18_ReferenceSignal_r18_Ssb_Index:
			if err := ie.ReferenceSignal_r18.Ssb_Index.Encode(e); err != nil {
				return err
			}
		case LTM_QCL_Info_r18_ReferenceSignal_r18_Csi_RS_Index:
			if err := ie.ReferenceSignal_r18.Csi_RS_Index.Encode(e); err != nil {
				return err
			}
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.ReferenceSignal_r18.Choice), Constraint: "index out of range"}
		}
	}

	// 3. qcl-Type-r18: enumerated
	{
		if err := e.EncodeEnumerated(ie.Qcl_Type_r18, lTMQCLInfoR18QclTypeR18Constraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *LTM_QCL_Info_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(lTMQCLInfoR18Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. referenceSignal-r18: choice
	{
		choiceDec := d.NewChoiceDecoder(lTM_QCL_Info_r18ReferenceSignalR18Constraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.ReferenceSignal_r18.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case LTM_QCL_Info_r18_ReferenceSignal_r18_Ssb_Index:
			ie.ReferenceSignal_r18.Ssb_Index = new(SSB_Index)
			if err := ie.ReferenceSignal_r18.Ssb_Index.Decode(d); err != nil {
				return err
			}
		case LTM_QCL_Info_r18_ReferenceSignal_r18_Csi_RS_Index:
			ie.ReferenceSignal_r18.Csi_RS_Index = new(NZP_CSI_RS_ResourceId)
			if err := ie.ReferenceSignal_r18.Csi_RS_Index.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. qcl-Type-r18: enumerated
	{
		v1, err := d.DecodeEnumerated(lTMQCLInfoR18QclTypeR18Constraints)
		if err != nil {
			return err
		}
		ie.Qcl_Type_r18 = v1
	}

	return nil
}
