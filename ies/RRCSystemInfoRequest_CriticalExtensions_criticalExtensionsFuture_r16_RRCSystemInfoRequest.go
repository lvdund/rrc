package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	RRCSystemInfoRequest_CriticalExtensions_criticalExtensionsFuture_r16_RRCSystemInfoRequest_Choice_nothing uint64 = iota
	RRCSystemInfoRequest_CriticalExtensions_criticalExtensionsFuture_r16_RRCSystemInfoRequest_Choice_RrcPosSystemInfoRequest_r16
	RRCSystemInfoRequest_CriticalExtensions_criticalExtensionsFuture_r16_RRCSystemInfoRequest_Choice_CriticalExtensionsFuture
)

type RRCSystemInfoRequest_CriticalExtensions_criticalExtensionsFuture_r16_RRCSystemInfoRequest struct {
	Choice                      uint64
	RrcPosSystemInfoRequest_r16 *RRC_PosSystemInfoRequest_r16_IEs
	CriticalExtensionsFuture    interface{} `madatory`
}

func (ie *RRCSystemInfoRequest_CriticalExtensions_criticalExtensionsFuture_r16_RRCSystemInfoRequest) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case RRCSystemInfoRequest_CriticalExtensions_criticalExtensionsFuture_r16_RRCSystemInfoRequest_Choice_RrcPosSystemInfoRequest_r16:
		if err = ie.RrcPosSystemInfoRequest_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode RrcPosSystemInfoRequest_r16", err)
		}
	case RRCSystemInfoRequest_CriticalExtensions_criticalExtensionsFuture_r16_RRCSystemInfoRequest_Choice_CriticalExtensionsFuture:
		// interface{} field of choice - nothing to encode
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *RRCSystemInfoRequest_CriticalExtensions_criticalExtensionsFuture_r16_RRCSystemInfoRequest) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case RRCSystemInfoRequest_CriticalExtensions_criticalExtensionsFuture_r16_RRCSystemInfoRequest_Choice_RrcPosSystemInfoRequest_r16:
		ie.RrcPosSystemInfoRequest_r16 = new(RRC_PosSystemInfoRequest_r16_IEs)
		if err = ie.RrcPosSystemInfoRequest_r16.Decode(r); err != nil {
			return utils.WrapError("Decode RrcPosSystemInfoRequest_r16", err)
		}
	case RRCSystemInfoRequest_CriticalExtensions_criticalExtensionsFuture_r16_RRCSystemInfoRequest_Choice_CriticalExtensionsFuture:
		// interface{} field of choice - nothing to decode
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
