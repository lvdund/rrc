package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CG_ConfigInfo_v1610_IEs struct {
	Drx_InfoMCG2                   *DRX_Info2                                     `optional`
	AlignedDRX_Indication          *CG_ConfigInfo_v1610_IEs_alignedDRX_Indication `optional`
	ScgFailureInfo_r16             *CG_ConfigInfo_v1610_IEs_scgFailureInfo_r16    `optional`
	Dummy1                         *CG_ConfigInfo_v1610_IEs_dummy1                `optional`
	SidelinkUEInformationNR_r16    *[]byte                                        `optional`
	SidelinkUEInformationEUTRA_r16 *[]byte                                        `optional`
	NonCriticalExtension           *CG_ConfigInfo_v1620_IEs                       `optional`
}

func (ie *CG_ConfigInfo_v1610_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.Drx_InfoMCG2 != nil, ie.AlignedDRX_Indication != nil, ie.ScgFailureInfo_r16 != nil, ie.Dummy1 != nil, ie.SidelinkUEInformationNR_r16 != nil, ie.SidelinkUEInformationEUTRA_r16 != nil, ie.NonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Drx_InfoMCG2 != nil {
		if err = ie.Drx_InfoMCG2.Encode(w); err != nil {
			return utils.WrapError("Encode Drx_InfoMCG2", err)
		}
	}
	if ie.AlignedDRX_Indication != nil {
		if err = ie.AlignedDRX_Indication.Encode(w); err != nil {
			return utils.WrapError("Encode AlignedDRX_Indication", err)
		}
	}
	if ie.ScgFailureInfo_r16 != nil {
		if err = ie.ScgFailureInfo_r16.Encode(w); err != nil {
			return utils.WrapError("Encode ScgFailureInfo_r16", err)
		}
	}
	if ie.Dummy1 != nil {
		if err = ie.Dummy1.Encode(w); err != nil {
			return utils.WrapError("Encode Dummy1", err)
		}
	}
	if ie.SidelinkUEInformationNR_r16 != nil {
		if err = w.WriteOctetString(*ie.SidelinkUEInformationNR_r16, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode SidelinkUEInformationNR_r16", err)
		}
	}
	if ie.SidelinkUEInformationEUTRA_r16 != nil {
		if err = w.WriteOctetString(*ie.SidelinkUEInformationEUTRA_r16, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode SidelinkUEInformationEUTRA_r16", err)
		}
	}
	if ie.NonCriticalExtension != nil {
		if err = ie.NonCriticalExtension.Encode(w); err != nil {
			return utils.WrapError("Encode NonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *CG_ConfigInfo_v1610_IEs) Decode(r *uper.UperReader) error {
	var err error
	var Drx_InfoMCG2Present bool
	if Drx_InfoMCG2Present, err = r.ReadBool(); err != nil {
		return err
	}
	var AlignedDRX_IndicationPresent bool
	if AlignedDRX_IndicationPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var ScgFailureInfo_r16Present bool
	if ScgFailureInfo_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Dummy1Present bool
	if Dummy1Present, err = r.ReadBool(); err != nil {
		return err
	}
	var SidelinkUEInformationNR_r16Present bool
	if SidelinkUEInformationNR_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var SidelinkUEInformationEUTRA_r16Present bool
	if SidelinkUEInformationEUTRA_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var NonCriticalExtensionPresent bool
	if NonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if Drx_InfoMCG2Present {
		ie.Drx_InfoMCG2 = new(DRX_Info2)
		if err = ie.Drx_InfoMCG2.Decode(r); err != nil {
			return utils.WrapError("Decode Drx_InfoMCG2", err)
		}
	}
	if AlignedDRX_IndicationPresent {
		ie.AlignedDRX_Indication = new(CG_ConfigInfo_v1610_IEs_alignedDRX_Indication)
		if err = ie.AlignedDRX_Indication.Decode(r); err != nil {
			return utils.WrapError("Decode AlignedDRX_Indication", err)
		}
	}
	if ScgFailureInfo_r16Present {
		ie.ScgFailureInfo_r16 = new(CG_ConfigInfo_v1610_IEs_scgFailureInfo_r16)
		if err = ie.ScgFailureInfo_r16.Decode(r); err != nil {
			return utils.WrapError("Decode ScgFailureInfo_r16", err)
		}
	}
	if Dummy1Present {
		ie.Dummy1 = new(CG_ConfigInfo_v1610_IEs_dummy1)
		if err = ie.Dummy1.Decode(r); err != nil {
			return utils.WrapError("Decode Dummy1", err)
		}
	}
	if SidelinkUEInformationNR_r16Present {
		var tmp_os_SidelinkUEInformationNR_r16 []byte
		if tmp_os_SidelinkUEInformationNR_r16, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode SidelinkUEInformationNR_r16", err)
		}
		ie.SidelinkUEInformationNR_r16 = &tmp_os_SidelinkUEInformationNR_r16
	}
	if SidelinkUEInformationEUTRA_r16Present {
		var tmp_os_SidelinkUEInformationEUTRA_r16 []byte
		if tmp_os_SidelinkUEInformationEUTRA_r16, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode SidelinkUEInformationEUTRA_r16", err)
		}
		ie.SidelinkUEInformationEUTRA_r16 = &tmp_os_SidelinkUEInformationEUTRA_r16
	}
	if NonCriticalExtensionPresent {
		ie.NonCriticalExtension = new(CG_ConfigInfo_v1620_IEs)
		if err = ie.NonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode NonCriticalExtension", err)
		}
	}
	return nil
}
