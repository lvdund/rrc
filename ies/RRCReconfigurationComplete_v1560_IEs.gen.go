// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: RRCReconfigurationComplete-v1560-IEs (line 1154).

var rRCReconfigurationCompleteV1560IEsConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "scg-Response", Optional: true},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

var rRCReconfigurationComplete_v1560_IEsScgResponseConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "nr-SCG-Response"},
		{Name: "eutra-SCG-Response"},
	},
}

const (
	RRCReconfigurationComplete_v1560_IEs_Scg_Response_Nr_SCG_Response    = 0
	RRCReconfigurationComplete_v1560_IEs_Scg_Response_Eutra_SCG_Response = 1
)

type RRCReconfigurationComplete_v1560_IEs struct {
	Scg_Response *struct {
		Choice             int
		Nr_SCG_Response    *[]byte
		Eutra_SCG_Response *[]byte
	}
	NonCriticalExtension *RRCReconfigurationComplete_v1610_IEs
}

func (ie *RRCReconfigurationComplete_v1560_IEs) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(rRCReconfigurationCompleteV1560IEsConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Scg_Response != nil, ie.NonCriticalExtension != nil}); err != nil {
		return err
	}

	// 2. scg-Response: choice
	{
		if ie.Scg_Response != nil {
			choiceEnc := e.NewChoiceEncoder(rRCReconfigurationComplete_v1560_IEsScgResponseConstraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.Scg_Response).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.Scg_Response).Choice {
			case RRCReconfigurationComplete_v1560_IEs_Scg_Response_Nr_SCG_Response:
				if err := e.EncodeOctetString((*(*ie.Scg_Response).Nr_SCG_Response), per.SizeConstraints{}); err != nil {
					return err
				}
			case RRCReconfigurationComplete_v1560_IEs_Scg_Response_Eutra_SCG_Response:
				if err := e.EncodeOctetString((*(*ie.Scg_Response).Eutra_SCG_Response), per.SizeConstraints{}); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.Scg_Response).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 3. nonCriticalExtension: ref
	{
		if ie.NonCriticalExtension != nil {
			if err := ie.NonCriticalExtension.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *RRCReconfigurationComplete_v1560_IEs) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(rRCReconfigurationCompleteV1560IEsConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. scg-Response: choice
	{
		if seq.IsComponentPresent(0) {
			ie.Scg_Response = &struct {
				Choice             int
				Nr_SCG_Response    *[]byte
				Eutra_SCG_Response *[]byte
			}{}
			choiceDec := d.NewChoiceDecoder(rRCReconfigurationComplete_v1560_IEsScgResponseConstraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Scg_Response).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case RRCReconfigurationComplete_v1560_IEs_Scg_Response_Nr_SCG_Response:
				(*ie.Scg_Response).Nr_SCG_Response = new([]byte)
				v, err := d.DecodeOctetString(per.SizeConstraints{})
				if err != nil {
					return err
				}
				(*(*ie.Scg_Response).Nr_SCG_Response) = v
			case RRCReconfigurationComplete_v1560_IEs_Scg_Response_Eutra_SCG_Response:
				(*ie.Scg_Response).Eutra_SCG_Response = new([]byte)
				v, err := d.DecodeOctetString(per.SizeConstraints{})
				if err != nil {
					return err
				}
				(*(*ie.Scg_Response).Eutra_SCG_Response) = v
			}
		}
	}

	// 3. nonCriticalExtension: ref
	{
		if seq.IsComponentPresent(1) {
			ie.NonCriticalExtension = new(RRCReconfigurationComplete_v1610_IEs)
			if err := ie.NonCriticalExtension.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
