package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type SL_SyncConfig_r16 struct {
	Sl_SyncRefMinHyst_r16      *SL_SyncConfig_r16_sl_SyncRefMinHyst_r16  `optional`
	Sl_SyncRefDiffHyst_r16     *SL_SyncConfig_r16_sl_SyncRefDiffHyst_r16 `optional`
	Sl_filterCoefficient_r16   *FilterCoefficient                        `optional`
	Sl_SSB_TimeAllocation1_r16 *SL_SSB_TimeAllocation_r16                `optional`
	Sl_SSB_TimeAllocation2_r16 *SL_SSB_TimeAllocation_r16                `optional`
	Sl_SSB_TimeAllocation3_r16 *SL_SSB_TimeAllocation_r16                `optional`
	Sl_SSID_r16                *int64                                    `lb:0,ub:671,optional`
	TxParameters_r16           *SL_SyncConfig_r16_txParameters_r16       `lb:2,ub:2,optional`
	Gnss_Sync_r16              *SL_SyncConfig_r16_gnss_Sync_r16          `optional`
}

func (ie *SL_SyncConfig_r16) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Sl_SyncRefMinHyst_r16 != nil, ie.Sl_SyncRefDiffHyst_r16 != nil, ie.Sl_filterCoefficient_r16 != nil, ie.Sl_SSB_TimeAllocation1_r16 != nil, ie.Sl_SSB_TimeAllocation2_r16 != nil, ie.Sl_SSB_TimeAllocation3_r16 != nil, ie.Sl_SSID_r16 != nil, ie.TxParameters_r16 != nil, ie.Gnss_Sync_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Sl_SyncRefMinHyst_r16 != nil {
		if err = ie.Sl_SyncRefMinHyst_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_SyncRefMinHyst_r16", err)
		}
	}
	if ie.Sl_SyncRefDiffHyst_r16 != nil {
		if err = ie.Sl_SyncRefDiffHyst_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_SyncRefDiffHyst_r16", err)
		}
	}
	if ie.Sl_filterCoefficient_r16 != nil {
		if err = ie.Sl_filterCoefficient_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_filterCoefficient_r16", err)
		}
	}
	if ie.Sl_SSB_TimeAllocation1_r16 != nil {
		if err = ie.Sl_SSB_TimeAllocation1_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_SSB_TimeAllocation1_r16", err)
		}
	}
	if ie.Sl_SSB_TimeAllocation2_r16 != nil {
		if err = ie.Sl_SSB_TimeAllocation2_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_SSB_TimeAllocation2_r16", err)
		}
	}
	if ie.Sl_SSB_TimeAllocation3_r16 != nil {
		if err = ie.Sl_SSB_TimeAllocation3_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_SSB_TimeAllocation3_r16", err)
		}
	}
	if ie.Sl_SSID_r16 != nil {
		if err = w.WriteInteger(*ie.Sl_SSID_r16, &aper.Constraint{Lb: 0, Ub: 671}, false); err != nil {
			return utils.WrapError("Encode Sl_SSID_r16", err)
		}
	}
	if ie.TxParameters_r16 != nil {
		if err = ie.TxParameters_r16.Encode(w); err != nil {
			return utils.WrapError("Encode TxParameters_r16", err)
		}
	}
	if ie.Gnss_Sync_r16 != nil {
		if err = ie.Gnss_Sync_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Gnss_Sync_r16", err)
		}
	}
	return nil
}

func (ie *SL_SyncConfig_r16) Decode(r *aper.AperReader) error {
	var err error
	var Sl_SyncRefMinHyst_r16Present bool
	if Sl_SyncRefMinHyst_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_SyncRefDiffHyst_r16Present bool
	if Sl_SyncRefDiffHyst_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_filterCoefficient_r16Present bool
	if Sl_filterCoefficient_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_SSB_TimeAllocation1_r16Present bool
	if Sl_SSB_TimeAllocation1_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_SSB_TimeAllocation2_r16Present bool
	if Sl_SSB_TimeAllocation2_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_SSB_TimeAllocation3_r16Present bool
	if Sl_SSB_TimeAllocation3_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_SSID_r16Present bool
	if Sl_SSID_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var TxParameters_r16Present bool
	if TxParameters_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Gnss_Sync_r16Present bool
	if Gnss_Sync_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Sl_SyncRefMinHyst_r16Present {
		ie.Sl_SyncRefMinHyst_r16 = new(SL_SyncConfig_r16_sl_SyncRefMinHyst_r16)
		if err = ie.Sl_SyncRefMinHyst_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_SyncRefMinHyst_r16", err)
		}
	}
	if Sl_SyncRefDiffHyst_r16Present {
		ie.Sl_SyncRefDiffHyst_r16 = new(SL_SyncConfig_r16_sl_SyncRefDiffHyst_r16)
		if err = ie.Sl_SyncRefDiffHyst_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_SyncRefDiffHyst_r16", err)
		}
	}
	if Sl_filterCoefficient_r16Present {
		ie.Sl_filterCoefficient_r16 = new(FilterCoefficient)
		if err = ie.Sl_filterCoefficient_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_filterCoefficient_r16", err)
		}
	}
	if Sl_SSB_TimeAllocation1_r16Present {
		ie.Sl_SSB_TimeAllocation1_r16 = new(SL_SSB_TimeAllocation_r16)
		if err = ie.Sl_SSB_TimeAllocation1_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_SSB_TimeAllocation1_r16", err)
		}
	}
	if Sl_SSB_TimeAllocation2_r16Present {
		ie.Sl_SSB_TimeAllocation2_r16 = new(SL_SSB_TimeAllocation_r16)
		if err = ie.Sl_SSB_TimeAllocation2_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_SSB_TimeAllocation2_r16", err)
		}
	}
	if Sl_SSB_TimeAllocation3_r16Present {
		ie.Sl_SSB_TimeAllocation3_r16 = new(SL_SSB_TimeAllocation_r16)
		if err = ie.Sl_SSB_TimeAllocation3_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_SSB_TimeAllocation3_r16", err)
		}
	}
	if Sl_SSID_r16Present {
		var tmp_int_Sl_SSID_r16 int64
		if tmp_int_Sl_SSID_r16, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 671}, false); err != nil {
			return utils.WrapError("Decode Sl_SSID_r16", err)
		}
		ie.Sl_SSID_r16 = &tmp_int_Sl_SSID_r16
	}
	if TxParameters_r16Present {
		ie.TxParameters_r16 = new(SL_SyncConfig_r16_txParameters_r16)
		if err = ie.TxParameters_r16.Decode(r); err != nil {
			return utils.WrapError("Decode TxParameters_r16", err)
		}
	}
	if Gnss_Sync_r16Present {
		ie.Gnss_Sync_r16 = new(SL_SyncConfig_r16_gnss_Sync_r16)
		if err = ie.Gnss_Sync_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Gnss_Sync_r16", err)
		}
	}
	return nil
}
