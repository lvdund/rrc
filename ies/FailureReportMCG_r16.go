package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type FailureReportMCG_r16 struct {
	FailureType_r16                *FailureReportMCG_r16_failureType_r16 `optional`
	MeasResultFreqList_r16         *MeasResultList2NR                    `optional`
	MeasResultFreqListEUTRA_r16    *MeasResultList2EUTRA                 `optional`
	MeasResultSCG_r16              *[]byte                               `optional`
	MeasResultSCG_EUTRA_r16        *[]byte                               `optional`
	MeasResultFreqListUTRA_FDD_r16 *MeasResultList2UTRA                  `optional`
}

func (ie *FailureReportMCG_r16) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.FailureType_r16 != nil, ie.MeasResultFreqList_r16 != nil, ie.MeasResultFreqListEUTRA_r16 != nil, ie.MeasResultSCG_r16 != nil, ie.MeasResultSCG_EUTRA_r16 != nil, ie.MeasResultFreqListUTRA_FDD_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.FailureType_r16 != nil {
		if err = ie.FailureType_r16.Encode(w); err != nil {
			return utils.WrapError("Encode FailureType_r16", err)
		}
	}
	if ie.MeasResultFreqList_r16 != nil {
		if err = ie.MeasResultFreqList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode MeasResultFreqList_r16", err)
		}
	}
	if ie.MeasResultFreqListEUTRA_r16 != nil {
		if err = ie.MeasResultFreqListEUTRA_r16.Encode(w); err != nil {
			return utils.WrapError("Encode MeasResultFreqListEUTRA_r16", err)
		}
	}
	if ie.MeasResultSCG_r16 != nil {
		if err = w.WriteOctetString(*ie.MeasResultSCG_r16, nil, false); err != nil {
			return utils.WrapError("Encode MeasResultSCG_r16", err)
		}
	}
	if ie.MeasResultSCG_EUTRA_r16 != nil {
		if err = w.WriteOctetString(*ie.MeasResultSCG_EUTRA_r16, nil, false); err != nil {
			return utils.WrapError("Encode MeasResultSCG_EUTRA_r16", err)
		}
	}
	if ie.MeasResultFreqListUTRA_FDD_r16 != nil {
		if err = ie.MeasResultFreqListUTRA_FDD_r16.Encode(w); err != nil {
			return utils.WrapError("Encode MeasResultFreqListUTRA_FDD_r16", err)
		}
	}
	return nil
}

func (ie *FailureReportMCG_r16) Decode(r *aper.AperReader) error {
	var err error
	var FailureType_r16Present bool
	if FailureType_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var MeasResultFreqList_r16Present bool
	if MeasResultFreqList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var MeasResultFreqListEUTRA_r16Present bool
	if MeasResultFreqListEUTRA_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var MeasResultSCG_r16Present bool
	if MeasResultSCG_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var MeasResultSCG_EUTRA_r16Present bool
	if MeasResultSCG_EUTRA_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var MeasResultFreqListUTRA_FDD_r16Present bool
	if MeasResultFreqListUTRA_FDD_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if FailureType_r16Present {
		ie.FailureType_r16 = new(FailureReportMCG_r16_failureType_r16)
		if err = ie.FailureType_r16.Decode(r); err != nil {
			return utils.WrapError("Decode FailureType_r16", err)
		}
	}
	if MeasResultFreqList_r16Present {
		ie.MeasResultFreqList_r16 = new(MeasResultList2NR)
		if err = ie.MeasResultFreqList_r16.Decode(r); err != nil {
			return utils.WrapError("Decode MeasResultFreqList_r16", err)
		}
	}
	if MeasResultFreqListEUTRA_r16Present {
		ie.MeasResultFreqListEUTRA_r16 = new(MeasResultList2EUTRA)
		if err = ie.MeasResultFreqListEUTRA_r16.Decode(r); err != nil {
			return utils.WrapError("Decode MeasResultFreqListEUTRA_r16", err)
		}
	}
	if MeasResultSCG_r16Present {
		var tmp_os_MeasResultSCG_r16 []byte
		if tmp_os_MeasResultSCG_r16, err = r.ReadOctetString(nil, false); err != nil {
			return utils.WrapError("Decode MeasResultSCG_r16", err)
		}
		ie.MeasResultSCG_r16 = &tmp_os_MeasResultSCG_r16
	}
	if MeasResultSCG_EUTRA_r16Present {
		var tmp_os_MeasResultSCG_EUTRA_r16 []byte
		if tmp_os_MeasResultSCG_EUTRA_r16, err = r.ReadOctetString(nil, false); err != nil {
			return utils.WrapError("Decode MeasResultSCG_EUTRA_r16", err)
		}
		ie.MeasResultSCG_EUTRA_r16 = &tmp_os_MeasResultSCG_EUTRA_r16
	}
	if MeasResultFreqListUTRA_FDD_r16Present {
		ie.MeasResultFreqListUTRA_FDD_r16 = new(MeasResultList2UTRA)
		if err = ie.MeasResultFreqListUTRA_FDD_r16.Decode(r); err != nil {
			return utils.WrapError("Decode MeasResultFreqListUTRA_FDD_r16", err)
		}
	}
	return nil
}
