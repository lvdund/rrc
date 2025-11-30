package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type RLF_Report_r16_eutra_RLF_Report_r16 struct {
	FailedPCellId_EUTRA               CGI_InfoEUTRALogging `madatory`
	MeasResult_RLF_Report_EUTRA_r16   []byte               `madatory`
	MeasResult_RLF_Report_EUTRA_v1690 *[]byte              `optional`
}

func (ie *RLF_Report_r16_eutra_RLF_Report_r16) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.MeasResult_RLF_Report_EUTRA_v1690 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.FailedPCellId_EUTRA.Encode(w); err != nil {
		return utils.WrapError("Encode FailedPCellId_EUTRA", err)
	}
	if err = w.WriteOctetString(ie.MeasResult_RLF_Report_EUTRA_r16, &aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("WriteOctetString MeasResult_RLF_Report_EUTRA_r16", err)
	}
	if ie.MeasResult_RLF_Report_EUTRA_v1690 != nil {
		if err = w.WriteOctetString(*ie.MeasResult_RLF_Report_EUTRA_v1690, &aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode MeasResult_RLF_Report_EUTRA_v1690", err)
		}
	}
	return nil
}

func (ie *RLF_Report_r16_eutra_RLF_Report_r16) Decode(r *aper.AperReader) error {
	var err error
	var MeasResult_RLF_Report_EUTRA_v1690Present bool
	if MeasResult_RLF_Report_EUTRA_v1690Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.FailedPCellId_EUTRA.Decode(r); err != nil {
		return utils.WrapError("Decode FailedPCellId_EUTRA", err)
	}
	var tmp_os_MeasResult_RLF_Report_EUTRA_r16 []byte
	if tmp_os_MeasResult_RLF_Report_EUTRA_r16, err = r.ReadOctetString(&aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("ReadOctetString MeasResult_RLF_Report_EUTRA_r16", err)
	}
	ie.MeasResult_RLF_Report_EUTRA_r16 = tmp_os_MeasResult_RLF_Report_EUTRA_r16
	if MeasResult_RLF_Report_EUTRA_v1690Present {
		var tmp_os_MeasResult_RLF_Report_EUTRA_v1690 []byte
		if tmp_os_MeasResult_RLF_Report_EUTRA_v1690, err = r.ReadOctetString(&aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode MeasResult_RLF_Report_EUTRA_v1690", err)
		}
		ie.MeasResult_RLF_Report_EUTRA_v1690 = &tmp_os_MeasResult_RLF_Report_EUTRA_v1690
	}
	return nil
}
