package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type RemoteUEInformationSidelink_r17 struct {
	CriticalExtensions RemoteUEInformationSidelink_r17_CriticalExtensions `madatory`
}

func (ie *RemoteUEInformationSidelink_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.CriticalExtensions.Encode(w); err != nil {
		return utils.WrapError("Encode CriticalExtensions", err)
	}
	return nil
}

func (ie *RemoteUEInformationSidelink_r17) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.CriticalExtensions.Decode(r); err != nil {
		return utils.WrapError("Decode CriticalExtensions", err)
	}
	return nil
}
