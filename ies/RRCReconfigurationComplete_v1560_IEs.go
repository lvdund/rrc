package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type RRCReconfigurationComplete_v1560_IEs struct {
	Scg_Response         *RRCReconfigurationComplete_v1560_IEs_scg_Response `optional`
	NonCriticalExtension *RRCReconfigurationComplete_v1610_IEs              `optional`
}

func (ie *RRCReconfigurationComplete_v1560_IEs) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Scg_Response != nil, ie.NonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Scg_Response != nil {
		if err = ie.Scg_Response.Encode(w); err != nil {
			return utils.WrapError("Encode Scg_Response", err)
		}
	}
	if ie.NonCriticalExtension != nil {
		if err = ie.NonCriticalExtension.Encode(w); err != nil {
			return utils.WrapError("Encode NonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *RRCReconfigurationComplete_v1560_IEs) Decode(r *aper.AperReader) error {
	var err error
	var Scg_ResponsePresent bool
	if Scg_ResponsePresent, err = r.ReadBool(); err != nil {
		return err
	}
	var NonCriticalExtensionPresent bool
	if NonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if Scg_ResponsePresent {
		ie.Scg_Response = new(RRCReconfigurationComplete_v1560_IEs_scg_Response)
		if err = ie.Scg_Response.Decode(r); err != nil {
			return utils.WrapError("Decode Scg_Response", err)
		}
	}
	if NonCriticalExtensionPresent {
		ie.NonCriticalExtension = new(RRCReconfigurationComplete_v1610_IEs)
		if err = ie.NonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode NonCriticalExtension", err)
		}
	}
	return nil
}
