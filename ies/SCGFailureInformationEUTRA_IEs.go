package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SCGFailureInformationEUTRA_IEs struct {
	FailureReportSCG_EUTRA *FailureReportSCG_EUTRA               `optional`
	NonCriticalExtension   *SCGFailureInformationEUTRA_v1590_IEs `optional`
}

func (ie *SCGFailureInformationEUTRA_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.FailureReportSCG_EUTRA != nil, ie.NonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.FailureReportSCG_EUTRA != nil {
		if err = ie.FailureReportSCG_EUTRA.Encode(w); err != nil {
			return utils.WrapError("Encode FailureReportSCG_EUTRA", err)
		}
	}
	if ie.NonCriticalExtension != nil {
		if err = ie.NonCriticalExtension.Encode(w); err != nil {
			return utils.WrapError("Encode NonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *SCGFailureInformationEUTRA_IEs) Decode(r *uper.UperReader) error {
	var err error
	var FailureReportSCG_EUTRAPresent bool
	if FailureReportSCG_EUTRAPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var NonCriticalExtensionPresent bool
	if NonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if FailureReportSCG_EUTRAPresent {
		ie.FailureReportSCG_EUTRA = new(FailureReportSCG_EUTRA)
		if err = ie.FailureReportSCG_EUTRA.Decode(r); err != nil {
			return utils.WrapError("Decode FailureReportSCG_EUTRA", err)
		}
	}
	if NonCriticalExtensionPresent {
		ie.NonCriticalExtension = new(SCGFailureInformationEUTRA_v1590_IEs)
		if err = ie.NonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode NonCriticalExtension", err)
		}
	}
	return nil
}
