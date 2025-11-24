package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SIB1_v1700_IEs struct {
	Hsdn_Cell_r17                  *SIB1_v1700_IEs_hsdn_Cell_r17                  `optional`
	Uac_BarringInfo_v1700          *SIB1_v1700_IEs_uac_BarringInfo_v1700          `optional`
	Sdt_ConfigCommon_r17           *SDT_ConfigCommonSIB_r17                       `optional`
	RedCap_ConfigCommon_r17        *RedCap_ConfigCommonSIB_r17                    `optional`
	FeaturePriorities_r17          *FeaturePriorities_r17                         `optional`
	Si_SchedulingInfo_v1700        *SI_SchedulingInfo_v1700                       `optional`
	HyperSFN_r17                   *uper.BitString                                `lb:10,ub:10,optional`
	EDRX_AllowedIdle_r17           *SIB1_v1700_IEs_eDRX_AllowedIdle_r17           `optional`
	EDRX_AllowedInactive_r17       *SIB1_v1700_IEs_eDRX_AllowedInactive_r17       `optional`
	IntraFreqReselectionRedCap_r17 *SIB1_v1700_IEs_intraFreqReselectionRedCap_r17 `optional`
	CellBarredNTN_r17              *SIB1_v1700_IEs_cellBarredNTN_r17              `optional`
	NonCriticalExtension           interface{}                                    `optional`
}

func (ie *SIB1_v1700_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.Hsdn_Cell_r17 != nil, ie.Uac_BarringInfo_v1700 != nil, ie.Sdt_ConfigCommon_r17 != nil, ie.RedCap_ConfigCommon_r17 != nil, ie.FeaturePriorities_r17 != nil, ie.Si_SchedulingInfo_v1700 != nil, ie.HyperSFN_r17 != nil, ie.EDRX_AllowedIdle_r17 != nil, ie.EDRX_AllowedInactive_r17 != nil, ie.IntraFreqReselectionRedCap_r17 != nil, ie.CellBarredNTN_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Hsdn_Cell_r17 != nil {
		if err = ie.Hsdn_Cell_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Hsdn_Cell_r17", err)
		}
	}
	if ie.Uac_BarringInfo_v1700 != nil {
		if err = ie.Uac_BarringInfo_v1700.Encode(w); err != nil {
			return utils.WrapError("Encode Uac_BarringInfo_v1700", err)
		}
	}
	if ie.Sdt_ConfigCommon_r17 != nil {
		if err = ie.Sdt_ConfigCommon_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Sdt_ConfigCommon_r17", err)
		}
	}
	if ie.RedCap_ConfigCommon_r17 != nil {
		if err = ie.RedCap_ConfigCommon_r17.Encode(w); err != nil {
			return utils.WrapError("Encode RedCap_ConfigCommon_r17", err)
		}
	}
	if ie.FeaturePriorities_r17 != nil {
		if err = ie.FeaturePriorities_r17.Encode(w); err != nil {
			return utils.WrapError("Encode FeaturePriorities_r17", err)
		}
	}
	if ie.Si_SchedulingInfo_v1700 != nil {
		if err = ie.Si_SchedulingInfo_v1700.Encode(w); err != nil {
			return utils.WrapError("Encode Si_SchedulingInfo_v1700", err)
		}
	}
	if ie.HyperSFN_r17 != nil {
		if err = w.WriteBitString(ie.HyperSFN_r17.Bytes, uint(ie.HyperSFN_r17.NumBits), &uper.Constraint{Lb: 10, Ub: 10}, false); err != nil {
			return utils.WrapError("Encode HyperSFN_r17", err)
		}
	}
	if ie.EDRX_AllowedIdle_r17 != nil {
		if err = ie.EDRX_AllowedIdle_r17.Encode(w); err != nil {
			return utils.WrapError("Encode EDRX_AllowedIdle_r17", err)
		}
	}
	if ie.EDRX_AllowedInactive_r17 != nil {
		if err = ie.EDRX_AllowedInactive_r17.Encode(w); err != nil {
			return utils.WrapError("Encode EDRX_AllowedInactive_r17", err)
		}
	}
	if ie.IntraFreqReselectionRedCap_r17 != nil {
		if err = ie.IntraFreqReselectionRedCap_r17.Encode(w); err != nil {
			return utils.WrapError("Encode IntraFreqReselectionRedCap_r17", err)
		}
	}
	if ie.CellBarredNTN_r17 != nil {
		if err = ie.CellBarredNTN_r17.Encode(w); err != nil {
			return utils.WrapError("Encode CellBarredNTN_r17", err)
		}
	}
	return nil
}

func (ie *SIB1_v1700_IEs) Decode(r *uper.UperReader) error {
	var err error
	var Hsdn_Cell_r17Present bool
	if Hsdn_Cell_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Uac_BarringInfo_v1700Present bool
	if Uac_BarringInfo_v1700Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sdt_ConfigCommon_r17Present bool
	if Sdt_ConfigCommon_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var RedCap_ConfigCommon_r17Present bool
	if RedCap_ConfigCommon_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var FeaturePriorities_r17Present bool
	if FeaturePriorities_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Si_SchedulingInfo_v1700Present bool
	if Si_SchedulingInfo_v1700Present, err = r.ReadBool(); err != nil {
		return err
	}
	var HyperSFN_r17Present bool
	if HyperSFN_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var EDRX_AllowedIdle_r17Present bool
	if EDRX_AllowedIdle_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var EDRX_AllowedInactive_r17Present bool
	if EDRX_AllowedInactive_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var IntraFreqReselectionRedCap_r17Present bool
	if IntraFreqReselectionRedCap_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var CellBarredNTN_r17Present bool
	if CellBarredNTN_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Hsdn_Cell_r17Present {
		ie.Hsdn_Cell_r17 = new(SIB1_v1700_IEs_hsdn_Cell_r17)
		if err = ie.Hsdn_Cell_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Hsdn_Cell_r17", err)
		}
	}
	if Uac_BarringInfo_v1700Present {
		ie.Uac_BarringInfo_v1700 = new(SIB1_v1700_IEs_uac_BarringInfo_v1700)
		if err = ie.Uac_BarringInfo_v1700.Decode(r); err != nil {
			return utils.WrapError("Decode Uac_BarringInfo_v1700", err)
		}
	}
	if Sdt_ConfigCommon_r17Present {
		ie.Sdt_ConfigCommon_r17 = new(SDT_ConfigCommonSIB_r17)
		if err = ie.Sdt_ConfigCommon_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Sdt_ConfigCommon_r17", err)
		}
	}
	if RedCap_ConfigCommon_r17Present {
		ie.RedCap_ConfigCommon_r17 = new(RedCap_ConfigCommonSIB_r17)
		if err = ie.RedCap_ConfigCommon_r17.Decode(r); err != nil {
			return utils.WrapError("Decode RedCap_ConfigCommon_r17", err)
		}
	}
	if FeaturePriorities_r17Present {
		ie.FeaturePriorities_r17 = new(FeaturePriorities_r17)
		if err = ie.FeaturePriorities_r17.Decode(r); err != nil {
			return utils.WrapError("Decode FeaturePriorities_r17", err)
		}
	}
	if Si_SchedulingInfo_v1700Present {
		ie.Si_SchedulingInfo_v1700 = new(SI_SchedulingInfo_v1700)
		if err = ie.Si_SchedulingInfo_v1700.Decode(r); err != nil {
			return utils.WrapError("Decode Si_SchedulingInfo_v1700", err)
		}
	}
	if HyperSFN_r17Present {
		var tmp_bs_HyperSFN_r17 []byte
		var n_HyperSFN_r17 uint
		if tmp_bs_HyperSFN_r17, n_HyperSFN_r17, err = r.ReadBitString(&uper.Constraint{Lb: 10, Ub: 10}, false); err != nil {
			return utils.WrapError("Decode HyperSFN_r17", err)
		}
		tmp_bitstring := uper.BitString{
			Bytes:   tmp_bs_HyperSFN_r17,
			NumBits: uint64(n_HyperSFN_r17),
		}
		ie.HyperSFN_r17 = &tmp_bitstring
	}
	if EDRX_AllowedIdle_r17Present {
		ie.EDRX_AllowedIdle_r17 = new(SIB1_v1700_IEs_eDRX_AllowedIdle_r17)
		if err = ie.EDRX_AllowedIdle_r17.Decode(r); err != nil {
			return utils.WrapError("Decode EDRX_AllowedIdle_r17", err)
		}
	}
	if EDRX_AllowedInactive_r17Present {
		ie.EDRX_AllowedInactive_r17 = new(SIB1_v1700_IEs_eDRX_AllowedInactive_r17)
		if err = ie.EDRX_AllowedInactive_r17.Decode(r); err != nil {
			return utils.WrapError("Decode EDRX_AllowedInactive_r17", err)
		}
	}
	if IntraFreqReselectionRedCap_r17Present {
		ie.IntraFreqReselectionRedCap_r17 = new(SIB1_v1700_IEs_intraFreqReselectionRedCap_r17)
		if err = ie.IntraFreqReselectionRedCap_r17.Decode(r); err != nil {
			return utils.WrapError("Decode IntraFreqReselectionRedCap_r17", err)
		}
	}
	if CellBarredNTN_r17Present {
		ie.CellBarredNTN_r17 = new(SIB1_v1700_IEs_cellBarredNTN_r17)
		if err = ie.CellBarredNTN_r17.Decode(r); err != nil {
			return utils.WrapError("Decode CellBarredNTN_r17", err)
		}
	}
	return nil
}
