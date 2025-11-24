package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	UL_CCCH_MessageType_C1_Choice_nothing uint64 = iota
	UL_CCCH_MessageType_C1_Choice_RrcSetupRequest
	UL_CCCH_MessageType_C1_Choice_RrcResumeRequest
	UL_CCCH_MessageType_C1_Choice_RrcReestablishmentRequest
	UL_CCCH_MessageType_C1_Choice_RrcSystemInfoRequest
)

type UL_CCCH_MessageType_C1 struct {
	Choice                    uint64
	RrcSetupRequest           *RRCSetupRequest
	RrcResumeRequest          *RRCResumeRequest
	RrcReestablishmentRequest *RRCReestablishmentRequest
	RrcSystemInfoRequest      *RRCSystemInfoRequest
}

func (ie *UL_CCCH_MessageType_C1) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 4, false); err != nil {
		return err
	}
	switch ie.Choice {
	case UL_CCCH_MessageType_C1_Choice_RrcSetupRequest:
		if err = ie.RrcSetupRequest.Encode(w); err != nil {
			err = utils.WrapError("Encode RrcSetupRequest", err)
		}
	case UL_CCCH_MessageType_C1_Choice_RrcResumeRequest:
		if err = ie.RrcResumeRequest.Encode(w); err != nil {
			err = utils.WrapError("Encode RrcResumeRequest", err)
		}
	case UL_CCCH_MessageType_C1_Choice_RrcReestablishmentRequest:
		if err = ie.RrcReestablishmentRequest.Encode(w); err != nil {
			err = utils.WrapError("Encode RrcReestablishmentRequest", err)
		}
	case UL_CCCH_MessageType_C1_Choice_RrcSystemInfoRequest:
		if err = ie.RrcSystemInfoRequest.Encode(w); err != nil {
			err = utils.WrapError("Encode RrcSystemInfoRequest", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *UL_CCCH_MessageType_C1) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(4, false); err != nil {
		return err
	}
	switch ie.Choice {
	case UL_CCCH_MessageType_C1_Choice_RrcSetupRequest:
		ie.RrcSetupRequest = new(RRCSetupRequest)
		if err = ie.RrcSetupRequest.Decode(r); err != nil {
			return utils.WrapError("Decode RrcSetupRequest", err)
		}
	case UL_CCCH_MessageType_C1_Choice_RrcResumeRequest:
		ie.RrcResumeRequest = new(RRCResumeRequest)
		if err = ie.RrcResumeRequest.Decode(r); err != nil {
			return utils.WrapError("Decode RrcResumeRequest", err)
		}
	case UL_CCCH_MessageType_C1_Choice_RrcReestablishmentRequest:
		ie.RrcReestablishmentRequest = new(RRCReestablishmentRequest)
		if err = ie.RrcReestablishmentRequest.Decode(r); err != nil {
			return utils.WrapError("Decode RrcReestablishmentRequest", err)
		}
	case UL_CCCH_MessageType_C1_Choice_RrcSystemInfoRequest:
		ie.RrcSystemInfoRequest = new(RRCSystemInfoRequest)
		if err = ie.RrcSystemInfoRequest.Decode(r); err != nil {
			return utils.WrapError("Decode RrcSystemInfoRequest", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
