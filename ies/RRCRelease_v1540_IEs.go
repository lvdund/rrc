package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type RRCRelease_v1540_IEs struct {
	WaitTime             *RejectWaitTime       `optional`
	NonCriticalExtension *RRCRelease_v1610_IEs `optional`
}

func (ie *RRCRelease_v1540_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.WaitTime != nil, ie.NonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.WaitTime != nil {
		if err = ie.WaitTime.Encode(w); err != nil {
			return utils.WrapError("Encode WaitTime", err)
		}
	}
	if ie.NonCriticalExtension != nil {
		if err = ie.NonCriticalExtension.Encode(w); err != nil {
			return utils.WrapError("Encode NonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *RRCRelease_v1540_IEs) Decode(r *uper.UperReader) error {
	var err error
	var WaitTimePresent bool
	if WaitTimePresent, err = r.ReadBool(); err != nil {
		return err
	}
	var NonCriticalExtensionPresent bool
	if NonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if WaitTimePresent {
		ie.WaitTime = new(RejectWaitTime)
		if err = ie.WaitTime.Decode(r); err != nil {
			return utils.WrapError("Decode WaitTime", err)
		}
	}
	if NonCriticalExtensionPresent {
		ie.NonCriticalExtension = new(RRCRelease_v1610_IEs)
		if err = ie.NonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode NonCriticalExtension", err)
		}
	}
	return nil
}
