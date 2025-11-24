package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SCGFailureInformation_IEs struct {
	FailureReportSCG     *FailureReportSCG                `optional`
	NonCriticalExtension *SCGFailureInformation_v1590_IEs `optional`
}

func (ie *SCGFailureInformation_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.FailureReportSCG != nil, ie.NonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.FailureReportSCG != nil {
		if err = ie.FailureReportSCG.Encode(w); err != nil {
			return utils.WrapError("Encode FailureReportSCG", err)
		}
	}
	if ie.NonCriticalExtension != nil {
		if err = ie.NonCriticalExtension.Encode(w); err != nil {
			return utils.WrapError("Encode NonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *SCGFailureInformation_IEs) Decode(r *uper.UperReader) error {
	var err error
	var FailureReportSCGPresent bool
	if FailureReportSCGPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var NonCriticalExtensionPresent bool
	if NonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if FailureReportSCGPresent {
		ie.FailureReportSCG = new(FailureReportSCG)
		if err = ie.FailureReportSCG.Decode(r); err != nil {
			return utils.WrapError("Decode FailureReportSCG", err)
		}
	}
	if NonCriticalExtensionPresent {
		ie.NonCriticalExtension = new(SCGFailureInformation_v1590_IEs)
		if err = ie.NonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode NonCriticalExtension", err)
		}
	}
	return nil
}
