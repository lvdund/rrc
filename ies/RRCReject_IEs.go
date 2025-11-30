package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type RRCReject_IEs struct {
	WaitTime                 *RejectWaitTime `optional`
	LateNonCriticalExtension *[]byte         `optional`
	NonCriticalExtension     interface{}     `optional`
}

func (ie *RRCReject_IEs) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.WaitTime != nil, ie.LateNonCriticalExtension != nil}
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
	if ie.LateNonCriticalExtension != nil {
		if err = w.WriteOctetString(*ie.LateNonCriticalExtension, nil, false); err != nil {
			return utils.WrapError("Encode LateNonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *RRCReject_IEs) Decode(r *aper.AperReader) error {
	var err error
	var WaitTimePresent bool
	if WaitTimePresent, err = r.ReadBool(); err != nil {
		return err
	}
	var LateNonCriticalExtensionPresent bool
	if LateNonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if WaitTimePresent {
		ie.WaitTime = new(RejectWaitTime)
		if err = ie.WaitTime.Decode(r); err != nil {
			return utils.WrapError("Decode WaitTime", err)
		}
	}
	if LateNonCriticalExtensionPresent {
		var tmp_os_LateNonCriticalExtension []byte
		if tmp_os_LateNonCriticalExtension, err = r.ReadOctetString(nil, false); err != nil {
			return utils.WrapError("Decode LateNonCriticalExtension", err)
		}
		ie.LateNonCriticalExtension = &tmp_os_LateNonCriticalExtension
	}
	return nil
}
