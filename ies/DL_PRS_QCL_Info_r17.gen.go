// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: DL-PRS-QCL-Info-r17 (line 10721).

var dLPRSQCLInfoR17Constraints = per.ChoiceConstraints{
	Extensible: true,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "ssb-r17"},
		{Name: "dl-PRS-r17"},
	},
}

const (
	DL_PRS_QCL_Info_r17_Ssb_r17    = 0
	DL_PRS_QCL_Info_r17_Dl_PRS_r17 = 1
)

var dLPRSQCLInfoR17SsbR17Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "ssb-Index-r17"},
		{Name: "rs-Type-r17"},
	},
}

const (
	DL_PRS_QCL_Info_r17_Ssb_r17_Rs_Type_r17_TypeC            = 0
	DL_PRS_QCL_Info_r17_Ssb_r17_Rs_Type_r17_TypeD            = 1
	DL_PRS_QCL_Info_r17_Ssb_r17_Rs_Type_r17_TypeC_Plus_TypeD = 2
)

var dLPRSQCLInfoR17SsbR17RsTypeR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{DL_PRS_QCL_Info_r17_Ssb_r17_Rs_Type_r17_TypeC, DL_PRS_QCL_Info_r17_Ssb_r17_Rs_Type_r17_TypeD, DL_PRS_QCL_Info_r17_Ssb_r17_Rs_Type_r17_TypeC_Plus_TypeD},
}

var dLPRSQCLInfoR17DlPRSR17Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "qcl-DL-PRS-ResourceID-r17"},
	},
}

type DL_PRS_QCL_Info_r17 struct {
	Choice  int
	Ssb_r17 *struct {
		Ssb_Index_r17 int64
		Rs_Type_r17   int64
	}
	Dl_PRS_r17 *struct{ Qcl_DL_PRS_ResourceID_r17 NR_DL_PRS_ResourceID_r17 }
}

func (ie *DL_PRS_QCL_Info_r17) Encode(e *per.Encoder) error {
	choiceEnc := e.NewChoiceEncoder(dLPRSQCLInfoR17Constraints)
	isExt := false
	if err := choiceEnc.EncodeChoice(int64(ie.Choice), isExt, nil); err != nil {
		return err
	}
	switch ie.Choice {
	case DL_PRS_QCL_Info_r17_Ssb_r17:
		c := (*ie.Ssb_r17)
		dLPRSQCLInfoR17SsbR17Seq := e.NewSequenceEncoder(dLPRSQCLInfoR17SsbR17Constraints)
		if err := dLPRSQCLInfoR17SsbR17Seq.EncodeExtensionBit(false); err != nil {
			return err
		}
		if err := e.EncodeInteger(c.Ssb_Index_r17, per.Constrained(0, 63)); err != nil {
			return err
		}
		if err := e.EncodeEnumerated(c.Rs_Type_r17, dLPRSQCLInfoR17SsbR17RsTypeR17Constraints); err != nil {
			return err
		}
	case DL_PRS_QCL_Info_r17_Dl_PRS_r17:
		c := (*ie.Dl_PRS_r17)
		dLPRSQCLInfoR17DlPRSR17Seq := e.NewSequenceEncoder(dLPRSQCLInfoR17DlPRSR17Constraints)
		if err := dLPRSQCLInfoR17DlPRSR17Seq.EncodeExtensionBit(false); err != nil {
			return err
		}
		if err := c.Qcl_DL_PRS_ResourceID_r17.Encode(e); err != nil {
			return err
		}
	default:
		// Extension alternative: not modeled; bail out.
		return &per.ConstraintViolationError{
			TypeName:   "DL-PRS-QCL-Info-r17",
			Value:      int64(ie.Choice),
			Constraint: "choice index out of root range",
		}
	}
	return nil
}

func (ie *DL_PRS_QCL_Info_r17) Decode(d *per.Decoder) error {
	choiceDec := d.NewChoiceDecoder(dLPRSQCLInfoR17Constraints)
	index, isExt, _, err := choiceDec.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(index)
	if isExt {
		return &per.DecodeError{TypeName: "DL-PRS-QCL-Info-r17", Reason: "extension alternative not supported", Position: 0}
	}
	switch index {
	case DL_PRS_QCL_Info_r17_Ssb_r17:
		ie.Ssb_r17 = &struct {
			Ssb_Index_r17 int64
			Rs_Type_r17   int64
		}{}
		c := (*ie.Ssb_r17)
		dLPRSQCLInfoR17SsbR17Seq := d.NewSequenceDecoder(dLPRSQCLInfoR17SsbR17Constraints)
		if err := dLPRSQCLInfoR17SsbR17Seq.DecodeExtensionBit(); err != nil {
			return err
		}
		{
			v, err := d.DecodeInteger(per.Constrained(0, 63))
			if err != nil {
				return err
			}
			c.Ssb_Index_r17 = v
		}
		{
			v, err := d.DecodeEnumerated(dLPRSQCLInfoR17SsbR17RsTypeR17Constraints)
			if err != nil {
				return err
			}
			c.Rs_Type_r17 = v
		}
	case DL_PRS_QCL_Info_r17_Dl_PRS_r17:
		ie.Dl_PRS_r17 = &struct{ Qcl_DL_PRS_ResourceID_r17 NR_DL_PRS_ResourceID_r17 }{}
		c := (*ie.Dl_PRS_r17)
		dLPRSQCLInfoR17DlPRSR17Seq := d.NewSequenceDecoder(dLPRSQCLInfoR17DlPRSR17Constraints)
		if err := dLPRSQCLInfoR17DlPRSR17Seq.DecodeExtensionBit(); err != nil {
			return err
		}
		{
			if err := c.Qcl_DL_PRS_ResourceID_r17.Decode(d); err != nil {
				return err
			}
		}
	default:
		return &per.DecodeError{TypeName: "DL-PRS-QCL-Info-r17", Reason: "choice index out of root range", Position: 0}
	}
	return nil
}
