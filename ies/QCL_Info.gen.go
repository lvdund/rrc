// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: QCL-Info (line 16016).

var qCLInfoConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "cell", Optional: true},
		{Name: "bwp-Id", Optional: true},
		{Name: "referenceSignal"},
		{Name: "qcl-Type"},
	},
}

var qCL_InfoReferenceSignalConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "csi-rs"},
		{Name: "ssb"},
	},
}

const (
	QCL_Info_ReferenceSignal_Csi_Rs = 0
	QCL_Info_ReferenceSignal_Ssb    = 1
)

const (
	QCL_Info_Qcl_Type_TypeA = 0
	QCL_Info_Qcl_Type_TypeB = 1
	QCL_Info_Qcl_Type_TypeC = 2
	QCL_Info_Qcl_Type_TypeD = 3
)

var qCLInfoQclTypeConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{QCL_Info_Qcl_Type_TypeA, QCL_Info_Qcl_Type_TypeB, QCL_Info_Qcl_Type_TypeC, QCL_Info_Qcl_Type_TypeD},
}

type QCL_Info struct {
	Cell            *ServCellIndex
	Bwp_Id          *BWP_Id
	ReferenceSignal struct {
		Choice int
		Csi_Rs *NZP_CSI_RS_ResourceId
		Ssb    *SSB_Index
	}
	Qcl_Type int64
}

func (ie *QCL_Info) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(qCLInfoConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Cell != nil, ie.Bwp_Id != nil}); err != nil {
		return err
	}

	// 3. cell: ref
	{
		if ie.Cell != nil {
			if err := ie.Cell.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. bwp-Id: ref
	{
		if ie.Bwp_Id != nil {
			if err := ie.Bwp_Id.Encode(e); err != nil {
				return err
			}
		}
	}

	// 5. referenceSignal: choice
	{
		choiceEnc := e.NewChoiceEncoder(qCL_InfoReferenceSignalConstraints)
		if err := choiceEnc.EncodeChoice(int64(ie.ReferenceSignal.Choice), false, nil); err != nil {
			return err
		}
		switch ie.ReferenceSignal.Choice {
		case QCL_Info_ReferenceSignal_Csi_Rs:
			if err := ie.ReferenceSignal.Csi_Rs.Encode(e); err != nil {
				return err
			}
		case QCL_Info_ReferenceSignal_Ssb:
			if err := ie.ReferenceSignal.Ssb.Encode(e); err != nil {
				return err
			}
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.ReferenceSignal.Choice), Constraint: "index out of range"}
		}
	}

	// 6. qcl-Type: enumerated
	{
		if err := e.EncodeEnumerated(ie.Qcl_Type, qCLInfoQclTypeConstraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *QCL_Info) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(qCLInfoConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. cell: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Cell = new(ServCellIndex)
			if err := ie.Cell.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. bwp-Id: ref
	{
		if seq.IsComponentPresent(1) {
			ie.Bwp_Id = new(BWP_Id)
			if err := ie.Bwp_Id.Decode(d); err != nil {
				return err
			}
		}
	}

	// 5. referenceSignal: choice
	{
		choiceDec := d.NewChoiceDecoder(qCL_InfoReferenceSignalConstraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.ReferenceSignal.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case QCL_Info_ReferenceSignal_Csi_Rs:
			ie.ReferenceSignal.Csi_Rs = new(NZP_CSI_RS_ResourceId)
			if err := ie.ReferenceSignal.Csi_Rs.Decode(d); err != nil {
				return err
			}
		case QCL_Info_ReferenceSignal_Ssb:
			ie.ReferenceSignal.Ssb = new(SSB_Index)
			if err := ie.ReferenceSignal.Ssb.Decode(d); err != nil {
				return err
			}
		}
	}

	// 6. qcl-Type: enumerated
	{
		v3, err := d.DecodeEnumerated(qCLInfoQclTypeConstraints)
		if err != nil {
			return err
		}
		ie.Qcl_Type = v3
	}

	return nil
}
