package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type UERadioPagingInformation_v1700_IEs struct {
	Ue_RadioPagingInfo_r17            *[]byte                                                               `optional`
	InactiveStatePO_Determination_r17 *UERadioPagingInformation_v1700_IEs_inactiveStatePO_Determination_r17 `optional`
	NumberOfRxRedCap_r17              *UERadioPagingInformation_v1700_IEs_numberOfRxRedCap_r17              `optional`
	HalfDuplexFDD_TypeA_RedCap_r17    []FreqBandIndicatorNR                                                 `lb:1,ub:maxBands,optional`
	NonCriticalExtension              interface{}                                                           `optional`
}

func (ie *UERadioPagingInformation_v1700_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.Ue_RadioPagingInfo_r17 != nil, ie.InactiveStatePO_Determination_r17 != nil, ie.NumberOfRxRedCap_r17 != nil, len(ie.HalfDuplexFDD_TypeA_RedCap_r17) > 0}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Ue_RadioPagingInfo_r17 != nil {
		if err = w.WriteOctetString(*ie.Ue_RadioPagingInfo_r17, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode Ue_RadioPagingInfo_r17", err)
		}
	}
	if ie.InactiveStatePO_Determination_r17 != nil {
		if err = ie.InactiveStatePO_Determination_r17.Encode(w); err != nil {
			return utils.WrapError("Encode InactiveStatePO_Determination_r17", err)
		}
	}
	if ie.NumberOfRxRedCap_r17 != nil {
		if err = ie.NumberOfRxRedCap_r17.Encode(w); err != nil {
			return utils.WrapError("Encode NumberOfRxRedCap_r17", err)
		}
	}
	if len(ie.HalfDuplexFDD_TypeA_RedCap_r17) > 0 {
		tmp_HalfDuplexFDD_TypeA_RedCap_r17 := utils.NewSequence[*FreqBandIndicatorNR]([]*FreqBandIndicatorNR{}, uper.Constraint{Lb: 1, Ub: maxBands}, false)
		for _, i := range ie.HalfDuplexFDD_TypeA_RedCap_r17 {
			tmp_HalfDuplexFDD_TypeA_RedCap_r17.Value = append(tmp_HalfDuplexFDD_TypeA_RedCap_r17.Value, &i)
		}
		if err = tmp_HalfDuplexFDD_TypeA_RedCap_r17.Encode(w); err != nil {
			return utils.WrapError("Encode HalfDuplexFDD_TypeA_RedCap_r17", err)
		}
	}
	return nil
}

func (ie *UERadioPagingInformation_v1700_IEs) Decode(r *uper.UperReader) error {
	var err error
	var Ue_RadioPagingInfo_r17Present bool
	if Ue_RadioPagingInfo_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var InactiveStatePO_Determination_r17Present bool
	if InactiveStatePO_Determination_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var NumberOfRxRedCap_r17Present bool
	if NumberOfRxRedCap_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var HalfDuplexFDD_TypeA_RedCap_r17Present bool
	if HalfDuplexFDD_TypeA_RedCap_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Ue_RadioPagingInfo_r17Present {
		var tmp_os_Ue_RadioPagingInfo_r17 []byte
		if tmp_os_Ue_RadioPagingInfo_r17, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode Ue_RadioPagingInfo_r17", err)
		}
		ie.Ue_RadioPagingInfo_r17 = &tmp_os_Ue_RadioPagingInfo_r17
	}
	if InactiveStatePO_Determination_r17Present {
		ie.InactiveStatePO_Determination_r17 = new(UERadioPagingInformation_v1700_IEs_inactiveStatePO_Determination_r17)
		if err = ie.InactiveStatePO_Determination_r17.Decode(r); err != nil {
			return utils.WrapError("Decode InactiveStatePO_Determination_r17", err)
		}
	}
	if NumberOfRxRedCap_r17Present {
		ie.NumberOfRxRedCap_r17 = new(UERadioPagingInformation_v1700_IEs_numberOfRxRedCap_r17)
		if err = ie.NumberOfRxRedCap_r17.Decode(r); err != nil {
			return utils.WrapError("Decode NumberOfRxRedCap_r17", err)
		}
	}
	if HalfDuplexFDD_TypeA_RedCap_r17Present {
		tmp_HalfDuplexFDD_TypeA_RedCap_r17 := utils.NewSequence[*FreqBandIndicatorNR]([]*FreqBandIndicatorNR{}, uper.Constraint{Lb: 1, Ub: maxBands}, false)
		fn_HalfDuplexFDD_TypeA_RedCap_r17 := func() *FreqBandIndicatorNR {
			return new(FreqBandIndicatorNR)
		}
		if err = tmp_HalfDuplexFDD_TypeA_RedCap_r17.Decode(r, fn_HalfDuplexFDD_TypeA_RedCap_r17); err != nil {
			return utils.WrapError("Decode HalfDuplexFDD_TypeA_RedCap_r17", err)
		}
		ie.HalfDuplexFDD_TypeA_RedCap_r17 = []FreqBandIndicatorNR{}
		for _, i := range tmp_HalfDuplexFDD_TypeA_RedCap_r17.Value {
			ie.HalfDuplexFDD_TypeA_RedCap_r17 = append(ie.HalfDuplexFDD_TypeA_RedCap_r17, *i)
		}
	}
	return nil
}
