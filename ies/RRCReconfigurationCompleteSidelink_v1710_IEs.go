package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type RRCReconfigurationCompleteSidelink_v1710_IEs struct {
	Dummy                RRCReconfigurationCompleteSidelink_v1710_IEs_dummy `madatory`
	NonCriticalExtension *RRCReconfigurationCompleteSidelink_v1720_IEs      `optional`
}

func (ie *RRCReconfigurationCompleteSidelink_v1710_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.NonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.Dummy.Encode(w); err != nil {
		return utils.WrapError("Encode Dummy", err)
	}
	if ie.NonCriticalExtension != nil {
		if err = ie.NonCriticalExtension.Encode(w); err != nil {
			return utils.WrapError("Encode NonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *RRCReconfigurationCompleteSidelink_v1710_IEs) Decode(r *uper.UperReader) error {
	var err error
	var NonCriticalExtensionPresent bool
	if NonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.Dummy.Decode(r); err != nil {
		return utils.WrapError("Decode Dummy", err)
	}
	if NonCriticalExtensionPresent {
		ie.NonCriticalExtension = new(RRCReconfigurationCompleteSidelink_v1720_IEs)
		if err = ie.NonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode NonCriticalExtension", err)
		}
	}
	return nil
}
