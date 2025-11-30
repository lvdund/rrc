package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type MCGFailureInformation_r16_IEs struct {
	FailureReportMCG_r16     *FailureReportMCG_r16 `optional`
	LateNonCriticalExtension *[]byte               `optional`
	NonCriticalExtension     interface{}           `optional`
}

func (ie *MCGFailureInformation_r16_IEs) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.FailureReportMCG_r16 != nil, ie.LateNonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.FailureReportMCG_r16 != nil {
		if err = ie.FailureReportMCG_r16.Encode(w); err != nil {
			return utils.WrapError("Encode FailureReportMCG_r16", err)
		}
	}
	if ie.LateNonCriticalExtension != nil {
		if err = w.WriteOctetString(*ie.LateNonCriticalExtension, nil, false); err != nil {
			return utils.WrapError("Encode LateNonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *MCGFailureInformation_r16_IEs) Decode(r *aper.AperReader) error {
	var err error
	var FailureReportMCG_r16Present bool
	if FailureReportMCG_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var LateNonCriticalExtensionPresent bool
	if LateNonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if FailureReportMCG_r16Present {
		ie.FailureReportMCG_r16 = new(FailureReportMCG_r16)
		if err = ie.FailureReportMCG_r16.Decode(r); err != nil {
			return utils.WrapError("Decode FailureReportMCG_r16", err)
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
