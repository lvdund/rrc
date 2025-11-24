package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type RRCRelease_v1650_IEs struct {
	MpsPriorityIndication_r16 *RRCRelease_v1650_IEs_mpsPriorityIndication_r16 `optional`
	NonCriticalExtension      *RRCRelease_v1710_IEs                           `optional`
}

func (ie *RRCRelease_v1650_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.MpsPriorityIndication_r16 != nil, ie.NonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.MpsPriorityIndication_r16 != nil {
		if err = ie.MpsPriorityIndication_r16.Encode(w); err != nil {
			return utils.WrapError("Encode MpsPriorityIndication_r16", err)
		}
	}
	if ie.NonCriticalExtension != nil {
		if err = ie.NonCriticalExtension.Encode(w); err != nil {
			return utils.WrapError("Encode NonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *RRCRelease_v1650_IEs) Decode(r *uper.UperReader) error {
	var err error
	var MpsPriorityIndication_r16Present bool
	if MpsPriorityIndication_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var NonCriticalExtensionPresent bool
	if NonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if MpsPriorityIndication_r16Present {
		ie.MpsPriorityIndication_r16 = new(RRCRelease_v1650_IEs_mpsPriorityIndication_r16)
		if err = ie.MpsPriorityIndication_r16.Decode(r); err != nil {
			return utils.WrapError("Decode MpsPriorityIndication_r16", err)
		}
	}
	if NonCriticalExtensionPresent {
		ie.NonCriticalExtension = new(RRCRelease_v1710_IEs)
		if err = ie.NonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode NonCriticalExtension", err)
		}
	}
	return nil
}
