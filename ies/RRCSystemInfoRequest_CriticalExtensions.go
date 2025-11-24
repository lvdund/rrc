package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	RRCSystemInfoRequest_CriticalExtensions_Choice_nothing uint64 = iota
	RRCSystemInfoRequest_CriticalExtensions_Choice_RrcSystemInfoRequest
	RRCSystemInfoRequest_CriticalExtensions_Choice_CriticalExtensionsFuture_r16_RRCSystemInfoRequest
)

type RRCSystemInfoRequest_CriticalExtensions struct {
	Choice                                            uint64
	RrcSystemInfoRequest                              *RRCSystemInfoRequest_IEs
	CriticalExtensionsFuture_r16_RRCSystemInfoRequest *RRCSystemInfoRequest_CriticalExtensions_criticalExtensionsFuture_r16_RRCSystemInfoRequest
}

func (ie *RRCSystemInfoRequest_CriticalExtensions) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case RRCSystemInfoRequest_CriticalExtensions_Choice_RrcSystemInfoRequest:
		if err = ie.RrcSystemInfoRequest.Encode(w); err != nil {
			err = utils.WrapError("Encode RrcSystemInfoRequest", err)
		}
	case RRCSystemInfoRequest_CriticalExtensions_Choice_CriticalExtensionsFuture_r16_RRCSystemInfoRequest:
		if err = ie.CriticalExtensionsFuture_r16_RRCSystemInfoRequest.Encode(w); err != nil {
			err = utils.WrapError("Encode CriticalExtensionsFuture_r16_RRCSystemInfoRequest", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *RRCSystemInfoRequest_CriticalExtensions) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case RRCSystemInfoRequest_CriticalExtensions_Choice_RrcSystemInfoRequest:
		ie.RrcSystemInfoRequest = new(RRCSystemInfoRequest_IEs)
		if err = ie.RrcSystemInfoRequest.Decode(r); err != nil {
			return utils.WrapError("Decode RrcSystemInfoRequest", err)
		}
	case RRCSystemInfoRequest_CriticalExtensions_Choice_CriticalExtensionsFuture_r16_RRCSystemInfoRequest:
		ie.CriticalExtensionsFuture_r16_RRCSystemInfoRequest = new(RRCSystemInfoRequest_CriticalExtensions_criticalExtensionsFuture_r16_RRCSystemInfoRequest)
		if err = ie.CriticalExtensionsFuture_r16_RRCSystemInfoRequest.Decode(r); err != nil {
			return utils.WrapError("Decode CriticalExtensionsFuture_r16_RRCSystemInfoRequest", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
