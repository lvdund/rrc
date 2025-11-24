package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type DLInformationTransfer_v1610_IEs struct {
	ReferenceTimeInfo_r16 *ReferenceTimeInfo_r16           `optional`
	NonCriticalExtension  *DLInformationTransfer_v1700_IEs `optional`
}

func (ie *DLInformationTransfer_v1610_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.ReferenceTimeInfo_r16 != nil, ie.NonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.ReferenceTimeInfo_r16 != nil {
		if err = ie.ReferenceTimeInfo_r16.Encode(w); err != nil {
			return utils.WrapError("Encode ReferenceTimeInfo_r16", err)
		}
	}
	if ie.NonCriticalExtension != nil {
		if err = ie.NonCriticalExtension.Encode(w); err != nil {
			return utils.WrapError("Encode NonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *DLInformationTransfer_v1610_IEs) Decode(r *uper.UperReader) error {
	var err error
	var ReferenceTimeInfo_r16Present bool
	if ReferenceTimeInfo_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var NonCriticalExtensionPresent bool
	if NonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if ReferenceTimeInfo_r16Present {
		ie.ReferenceTimeInfo_r16 = new(ReferenceTimeInfo_r16)
		if err = ie.ReferenceTimeInfo_r16.Decode(r); err != nil {
			return utils.WrapError("Decode ReferenceTimeInfo_r16", err)
		}
	}
	if NonCriticalExtensionPresent {
		ie.NonCriticalExtension = new(DLInformationTransfer_v1700_IEs)
		if err = ie.NonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode NonCriticalExtension", err)
		}
	}
	return nil
}
