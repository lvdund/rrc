package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type SL_ConfigDedicatedNR_r16 struct {
	Sl_PHY_MAC_RLC_Config_r16          *SL_PHY_MAC_RLC_Config_r16         `optional`
	Sl_RadioBearerToReleaseList_r16    []SLRB_Uu_ConfigIndex_r16          `lb:1,ub:maxNrofSLRB_r16,optional`
	Sl_RadioBearerToAddModList_r16     []SL_RadioBearerConfig_r16         `lb:1,ub:maxNrofSLRB_r16,optional`
	Sl_MeasConfigInfoToReleaseList_r16 []SL_DestinationIndex_r16          `lb:1,ub:maxNrofSL_Dest_r16,optional`
	Sl_MeasConfigInfoToAddModList_r16  []SL_MeasConfigInfo_r16            `lb:1,ub:maxNrofSL_Dest_r16,optional`
	T400_r16                           *SL_ConfigDedicatedNR_r16_t400_r16 `optional`
	Sl_PHY_MAC_RLC_Config_v1700        *SL_PHY_MAC_RLC_Config_v1700       `optional,ext-1,setuprelease`
	Sl_DiscConfig_r17                  *SL_DiscConfig_r17                 `optional,ext-1,setuprelease`
}

func (ie *SL_ConfigDedicatedNR_r16) Encode(w *aper.AperWriter) error {
	var err error
	hasExtensions := ie.Sl_PHY_MAC_RLC_Config_v1700 != nil || ie.Sl_DiscConfig_r17 != nil
	preambleBits := []bool{hasExtensions, ie.Sl_PHY_MAC_RLC_Config_r16 != nil, len(ie.Sl_RadioBearerToReleaseList_r16) > 0, len(ie.Sl_RadioBearerToAddModList_r16) > 0, len(ie.Sl_MeasConfigInfoToReleaseList_r16) > 0, len(ie.Sl_MeasConfigInfoToAddModList_r16) > 0, ie.T400_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Sl_PHY_MAC_RLC_Config_r16 != nil {
		if err = ie.Sl_PHY_MAC_RLC_Config_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_PHY_MAC_RLC_Config_r16", err)
		}
	}
	if len(ie.Sl_RadioBearerToReleaseList_r16) > 0 {
		tmp_Sl_RadioBearerToReleaseList_r16 := utils.NewSequence[*SLRB_Uu_ConfigIndex_r16]([]*SLRB_Uu_ConfigIndex_r16{}, aper.Constraint{Lb: 1, Ub: maxNrofSLRB_r16}, false)
		for _, i := range ie.Sl_RadioBearerToReleaseList_r16 {
			tmp_Sl_RadioBearerToReleaseList_r16.Value = append(tmp_Sl_RadioBearerToReleaseList_r16.Value, &i)
		}
		if err = tmp_Sl_RadioBearerToReleaseList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_RadioBearerToReleaseList_r16", err)
		}
	}
	if len(ie.Sl_RadioBearerToAddModList_r16) > 0 {
		tmp_Sl_RadioBearerToAddModList_r16 := utils.NewSequence[*SL_RadioBearerConfig_r16]([]*SL_RadioBearerConfig_r16{}, aper.Constraint{Lb: 1, Ub: maxNrofSLRB_r16}, false)
		for _, i := range ie.Sl_RadioBearerToAddModList_r16 {
			tmp_Sl_RadioBearerToAddModList_r16.Value = append(tmp_Sl_RadioBearerToAddModList_r16.Value, &i)
		}
		if err = tmp_Sl_RadioBearerToAddModList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_RadioBearerToAddModList_r16", err)
		}
	}
	if len(ie.Sl_MeasConfigInfoToReleaseList_r16) > 0 {
		tmp_Sl_MeasConfigInfoToReleaseList_r16 := utils.NewSequence[*SL_DestinationIndex_r16]([]*SL_DestinationIndex_r16{}, aper.Constraint{Lb: 1, Ub: maxNrofSL_Dest_r16}, false)
		for _, i := range ie.Sl_MeasConfigInfoToReleaseList_r16 {
			tmp_Sl_MeasConfigInfoToReleaseList_r16.Value = append(tmp_Sl_MeasConfigInfoToReleaseList_r16.Value, &i)
		}
		if err = tmp_Sl_MeasConfigInfoToReleaseList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_MeasConfigInfoToReleaseList_r16", err)
		}
	}
	if len(ie.Sl_MeasConfigInfoToAddModList_r16) > 0 {
		tmp_Sl_MeasConfigInfoToAddModList_r16 := utils.NewSequence[*SL_MeasConfigInfo_r16]([]*SL_MeasConfigInfo_r16{}, aper.Constraint{Lb: 1, Ub: maxNrofSL_Dest_r16}, false)
		for _, i := range ie.Sl_MeasConfigInfoToAddModList_r16 {
			tmp_Sl_MeasConfigInfoToAddModList_r16.Value = append(tmp_Sl_MeasConfigInfoToAddModList_r16.Value, &i)
		}
		if err = tmp_Sl_MeasConfigInfoToAddModList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_MeasConfigInfoToAddModList_r16", err)
		}
	}
	if ie.T400_r16 != nil {
		if err = ie.T400_r16.Encode(w); err != nil {
			return utils.WrapError("Encode T400_r16", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 1 bits for 1 extension groups
		extBitmap := []bool{ie.Sl_PHY_MAC_RLC_Config_v1700 != nil || ie.Sl_DiscConfig_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap SL_ConfigDedicatedNR_r16", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := aper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.Sl_PHY_MAC_RLC_Config_v1700 != nil, ie.Sl_DiscConfig_r17 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode Sl_PHY_MAC_RLC_Config_v1700 optional
			if ie.Sl_PHY_MAC_RLC_Config_v1700 != nil {
				tmp_Sl_PHY_MAC_RLC_Config_v1700 := utils.SetupRelease[*SL_PHY_MAC_RLC_Config_v1700]{
					Setup: ie.Sl_PHY_MAC_RLC_Config_v1700,
				}
				if err = tmp_Sl_PHY_MAC_RLC_Config_v1700.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Sl_PHY_MAC_RLC_Config_v1700", err)
				}
			}
			// encode Sl_DiscConfig_r17 optional
			if ie.Sl_DiscConfig_r17 != nil {
				tmp_Sl_DiscConfig_r17 := utils.SetupRelease[*SL_DiscConfig_r17]{
					Setup: ie.Sl_DiscConfig_r17,
				}
				if err = tmp_Sl_DiscConfig_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Sl_DiscConfig_r17", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}
	}
	return nil
}

func (ie *SL_ConfigDedicatedNR_r16) Decode(r *aper.AperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_PHY_MAC_RLC_Config_r16Present bool
	if Sl_PHY_MAC_RLC_Config_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_RadioBearerToReleaseList_r16Present bool
	if Sl_RadioBearerToReleaseList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_RadioBearerToAddModList_r16Present bool
	if Sl_RadioBearerToAddModList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_MeasConfigInfoToReleaseList_r16Present bool
	if Sl_MeasConfigInfoToReleaseList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_MeasConfigInfoToAddModList_r16Present bool
	if Sl_MeasConfigInfoToAddModList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var T400_r16Present bool
	if T400_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Sl_PHY_MAC_RLC_Config_r16Present {
		ie.Sl_PHY_MAC_RLC_Config_r16 = new(SL_PHY_MAC_RLC_Config_r16)
		if err = ie.Sl_PHY_MAC_RLC_Config_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_PHY_MAC_RLC_Config_r16", err)
		}
	}
	if Sl_RadioBearerToReleaseList_r16Present {
		tmp_Sl_RadioBearerToReleaseList_r16 := utils.NewSequence[*SLRB_Uu_ConfigIndex_r16]([]*SLRB_Uu_ConfigIndex_r16{}, aper.Constraint{Lb: 1, Ub: maxNrofSLRB_r16}, false)
		fn_Sl_RadioBearerToReleaseList_r16 := func() *SLRB_Uu_ConfigIndex_r16 {
			return new(SLRB_Uu_ConfigIndex_r16)
		}
		if err = tmp_Sl_RadioBearerToReleaseList_r16.Decode(r, fn_Sl_RadioBearerToReleaseList_r16); err != nil {
			return utils.WrapError("Decode Sl_RadioBearerToReleaseList_r16", err)
		}
		ie.Sl_RadioBearerToReleaseList_r16 = []SLRB_Uu_ConfigIndex_r16{}
		for _, i := range tmp_Sl_RadioBearerToReleaseList_r16.Value {
			ie.Sl_RadioBearerToReleaseList_r16 = append(ie.Sl_RadioBearerToReleaseList_r16, *i)
		}
	}
	if Sl_RadioBearerToAddModList_r16Present {
		tmp_Sl_RadioBearerToAddModList_r16 := utils.NewSequence[*SL_RadioBearerConfig_r16]([]*SL_RadioBearerConfig_r16{}, aper.Constraint{Lb: 1, Ub: maxNrofSLRB_r16}, false)
		fn_Sl_RadioBearerToAddModList_r16 := func() *SL_RadioBearerConfig_r16 {
			return new(SL_RadioBearerConfig_r16)
		}
		if err = tmp_Sl_RadioBearerToAddModList_r16.Decode(r, fn_Sl_RadioBearerToAddModList_r16); err != nil {
			return utils.WrapError("Decode Sl_RadioBearerToAddModList_r16", err)
		}
		ie.Sl_RadioBearerToAddModList_r16 = []SL_RadioBearerConfig_r16{}
		for _, i := range tmp_Sl_RadioBearerToAddModList_r16.Value {
			ie.Sl_RadioBearerToAddModList_r16 = append(ie.Sl_RadioBearerToAddModList_r16, *i)
		}
	}
	if Sl_MeasConfigInfoToReleaseList_r16Present {
		tmp_Sl_MeasConfigInfoToReleaseList_r16 := utils.NewSequence[*SL_DestinationIndex_r16]([]*SL_DestinationIndex_r16{}, aper.Constraint{Lb: 1, Ub: maxNrofSL_Dest_r16}, false)
		fn_Sl_MeasConfigInfoToReleaseList_r16 := func() *SL_DestinationIndex_r16 {
			return new(SL_DestinationIndex_r16)
		}
		if err = tmp_Sl_MeasConfigInfoToReleaseList_r16.Decode(r, fn_Sl_MeasConfigInfoToReleaseList_r16); err != nil {
			return utils.WrapError("Decode Sl_MeasConfigInfoToReleaseList_r16", err)
		}
		ie.Sl_MeasConfigInfoToReleaseList_r16 = []SL_DestinationIndex_r16{}
		for _, i := range tmp_Sl_MeasConfigInfoToReleaseList_r16.Value {
			ie.Sl_MeasConfigInfoToReleaseList_r16 = append(ie.Sl_MeasConfigInfoToReleaseList_r16, *i)
		}
	}
	if Sl_MeasConfigInfoToAddModList_r16Present {
		tmp_Sl_MeasConfigInfoToAddModList_r16 := utils.NewSequence[*SL_MeasConfigInfo_r16]([]*SL_MeasConfigInfo_r16{}, aper.Constraint{Lb: 1, Ub: maxNrofSL_Dest_r16}, false)
		fn_Sl_MeasConfigInfoToAddModList_r16 := func() *SL_MeasConfigInfo_r16 {
			return new(SL_MeasConfigInfo_r16)
		}
		if err = tmp_Sl_MeasConfigInfoToAddModList_r16.Decode(r, fn_Sl_MeasConfigInfoToAddModList_r16); err != nil {
			return utils.WrapError("Decode Sl_MeasConfigInfoToAddModList_r16", err)
		}
		ie.Sl_MeasConfigInfoToAddModList_r16 = []SL_MeasConfigInfo_r16{}
		for _, i := range tmp_Sl_MeasConfigInfoToAddModList_r16.Value {
			ie.Sl_MeasConfigInfoToAddModList_r16 = append(ie.Sl_MeasConfigInfoToAddModList_r16, *i)
		}
	}
	if T400_r16Present {
		ie.T400_r16 = new(SL_ConfigDedicatedNR_r16_t400_r16)
		if err = ie.T400_r16.Decode(r); err != nil {
			return utils.WrapError("Decode T400_r16", err)
		}
	}

	if extensionBit {
		// Read extension bitmap: 1 bits for 1 extension groups
		extBitmap, err := r.ReadExtBitMap()
		if err != nil {
			return err
		}

		// decode extension group 1
		if len(extBitmap) > 0 && extBitmap[0] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := aper.NewReader(bytes.NewReader(extBytes))

			Sl_PHY_MAC_RLC_Config_v1700Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Sl_DiscConfig_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode Sl_PHY_MAC_RLC_Config_v1700 optional
			if Sl_PHY_MAC_RLC_Config_v1700Present {
				tmp_Sl_PHY_MAC_RLC_Config_v1700 := utils.SetupRelease[*SL_PHY_MAC_RLC_Config_v1700]{}
				if err = tmp_Sl_PHY_MAC_RLC_Config_v1700.Decode(extReader); err != nil {
					return utils.WrapError("Decode Sl_PHY_MAC_RLC_Config_v1700", err)
				}
				ie.Sl_PHY_MAC_RLC_Config_v1700 = tmp_Sl_PHY_MAC_RLC_Config_v1700.Setup
			}
			// decode Sl_DiscConfig_r17 optional
			if Sl_DiscConfig_r17Present {
				tmp_Sl_DiscConfig_r17 := utils.SetupRelease[*SL_DiscConfig_r17]{}
				if err = tmp_Sl_DiscConfig_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Sl_DiscConfig_r17", err)
				}
				ie.Sl_DiscConfig_r17 = tmp_Sl_DiscConfig_r17.Setup
			}
		}
	}
	return nil
}
