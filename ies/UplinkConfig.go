package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type UplinkConfig struct {
	InitialUplinkBWP                    *BWP_UplinkDedicated                              `optional`
	UplinkBWP_ToReleaseList             []BWP_Id                                          `lb:1,ub:maxNrofBWPs,optional`
	UplinkBWP_ToAddModList              []BWP_Uplink                                      `lb:1,ub:maxNrofBWPs,optional`
	FirstActiveUplinkBWP_Id             *BWP_Id                                           `optional`
	Pusch_ServingCellConfig             *PUSCH_ServingCellConfig                          `optional,setuprelease`
	CarrierSwitching                    *SRS_CarrierSwitching                             `optional,setuprelease`
	PowerBoostPi2BPSK                   *bool                                             `optional,ext-1`
	UplinkChannelBW_PerSCS_List         []SCS_SpecificCarrier                             `lb:1,ub:maxSCSs,optional,ext-1`
	EnablePL_RS_UpdateForPUSCH_SRS_r16  *UplinkConfig_enablePL_RS_UpdateForPUSCH_SRS_r16  `optional,ext-2`
	EnableDefaultBeamPL_ForPUSCH0_0_r16 *UplinkConfig_enableDefaultBeamPL_ForPUSCH0_0_r16 `optional,ext-2`
	EnableDefaultBeamPL_ForPUCCH_r16    *UplinkConfig_enableDefaultBeamPL_ForPUCCH_r16    `optional,ext-2`
	EnableDefaultBeamPL_ForSRS_r16      *UplinkConfig_enableDefaultBeamPL_ForSRS_r16      `optional,ext-2`
	UplinkTxSwitching_r16               *UplinkTxSwitching_r16                            `optional,ext-2,setuprelease`
	Mpr_PowerBoost_FR2_r16              *UplinkConfig_mpr_PowerBoost_FR2_r16              `optional,ext-2`
}

func (ie *UplinkConfig) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.PowerBoostPi2BPSK != nil || len(ie.UplinkChannelBW_PerSCS_List) > 0 || ie.EnablePL_RS_UpdateForPUSCH_SRS_r16 != nil || ie.EnableDefaultBeamPL_ForPUSCH0_0_r16 != nil || ie.EnableDefaultBeamPL_ForPUCCH_r16 != nil || ie.EnableDefaultBeamPL_ForSRS_r16 != nil || ie.UplinkTxSwitching_r16 != nil || ie.Mpr_PowerBoost_FR2_r16 != nil
	preambleBits := []bool{hasExtensions, ie.InitialUplinkBWP != nil, len(ie.UplinkBWP_ToReleaseList) > 0, len(ie.UplinkBWP_ToAddModList) > 0, ie.FirstActiveUplinkBWP_Id != nil, ie.Pusch_ServingCellConfig != nil, ie.CarrierSwitching != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.InitialUplinkBWP != nil {
		if err = ie.InitialUplinkBWP.Encode(w); err != nil {
			return utils.WrapError("Encode InitialUplinkBWP", err)
		}
	}
	if len(ie.UplinkBWP_ToReleaseList) > 0 {
		tmp_UplinkBWP_ToReleaseList := utils.NewSequence[*BWP_Id]([]*BWP_Id{}, uper.Constraint{Lb: 1, Ub: maxNrofBWPs}, false)
		for _, i := range ie.UplinkBWP_ToReleaseList {
			tmp_UplinkBWP_ToReleaseList.Value = append(tmp_UplinkBWP_ToReleaseList.Value, &i)
		}
		if err = tmp_UplinkBWP_ToReleaseList.Encode(w); err != nil {
			return utils.WrapError("Encode UplinkBWP_ToReleaseList", err)
		}
	}
	if len(ie.UplinkBWP_ToAddModList) > 0 {
		tmp_UplinkBWP_ToAddModList := utils.NewSequence[*BWP_Uplink]([]*BWP_Uplink{}, uper.Constraint{Lb: 1, Ub: maxNrofBWPs}, false)
		for _, i := range ie.UplinkBWP_ToAddModList {
			tmp_UplinkBWP_ToAddModList.Value = append(tmp_UplinkBWP_ToAddModList.Value, &i)
		}
		if err = tmp_UplinkBWP_ToAddModList.Encode(w); err != nil {
			return utils.WrapError("Encode UplinkBWP_ToAddModList", err)
		}
	}
	if ie.FirstActiveUplinkBWP_Id != nil {
		if err = ie.FirstActiveUplinkBWP_Id.Encode(w); err != nil {
			return utils.WrapError("Encode FirstActiveUplinkBWP_Id", err)
		}
	}
	if ie.Pusch_ServingCellConfig != nil {
		tmp_Pusch_ServingCellConfig := utils.SetupRelease[*PUSCH_ServingCellConfig]{
			Setup: ie.Pusch_ServingCellConfig,
		}
		if err = tmp_Pusch_ServingCellConfig.Encode(w); err != nil {
			return utils.WrapError("Encode Pusch_ServingCellConfig", err)
		}
	}
	if ie.CarrierSwitching != nil {
		tmp_CarrierSwitching := utils.SetupRelease[*SRS_CarrierSwitching]{
			Setup: ie.CarrierSwitching,
		}
		if err = tmp_CarrierSwitching.Encode(w); err != nil {
			return utils.WrapError("Encode CarrierSwitching", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 2 bits for 2 extension groups
		extBitmap := []bool{ie.PowerBoostPi2BPSK != nil || len(ie.UplinkChannelBW_PerSCS_List) > 0, ie.EnablePL_RS_UpdateForPUSCH_SRS_r16 != nil || ie.EnableDefaultBeamPL_ForPUSCH0_0_r16 != nil || ie.EnableDefaultBeamPL_ForPUCCH_r16 != nil || ie.EnableDefaultBeamPL_ForSRS_r16 != nil || ie.UplinkTxSwitching_r16 != nil || ie.Mpr_PowerBoost_FR2_r16 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap UplinkConfig", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.PowerBoostPi2BPSK != nil, len(ie.UplinkChannelBW_PerSCS_List) > 0}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode PowerBoostPi2BPSK optional
			if ie.PowerBoostPi2BPSK != nil {
				if err = extWriter.WriteBoolean(*ie.PowerBoostPi2BPSK); err != nil {
					return utils.WrapError("Encode PowerBoostPi2BPSK", err)
				}
			}
			// encode UplinkChannelBW_PerSCS_List optional
			if len(ie.UplinkChannelBW_PerSCS_List) > 0 {
				tmp_UplinkChannelBW_PerSCS_List := utils.NewSequence[*SCS_SpecificCarrier]([]*SCS_SpecificCarrier{}, uper.Constraint{Lb: 1, Ub: maxSCSs}, false)
				for _, i := range ie.UplinkChannelBW_PerSCS_List {
					tmp_UplinkChannelBW_PerSCS_List.Value = append(tmp_UplinkChannelBW_PerSCS_List.Value, &i)
				}
				if err = tmp_UplinkChannelBW_PerSCS_List.Encode(extWriter); err != nil {
					return utils.WrapError("Encode UplinkChannelBW_PerSCS_List", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}

		// encode extension group 2
		if extBitmap[1] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 2
			optionals_ext_2 := []bool{ie.EnablePL_RS_UpdateForPUSCH_SRS_r16 != nil, ie.EnableDefaultBeamPL_ForPUSCH0_0_r16 != nil, ie.EnableDefaultBeamPL_ForPUCCH_r16 != nil, ie.EnableDefaultBeamPL_ForSRS_r16 != nil, ie.UplinkTxSwitching_r16 != nil, ie.Mpr_PowerBoost_FR2_r16 != nil}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode EnablePL_RS_UpdateForPUSCH_SRS_r16 optional
			if ie.EnablePL_RS_UpdateForPUSCH_SRS_r16 != nil {
				if err = ie.EnablePL_RS_UpdateForPUSCH_SRS_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode EnablePL_RS_UpdateForPUSCH_SRS_r16", err)
				}
			}
			// encode EnableDefaultBeamPL_ForPUSCH0_0_r16 optional
			if ie.EnableDefaultBeamPL_ForPUSCH0_0_r16 != nil {
				if err = ie.EnableDefaultBeamPL_ForPUSCH0_0_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode EnableDefaultBeamPL_ForPUSCH0_0_r16", err)
				}
			}
			// encode EnableDefaultBeamPL_ForPUCCH_r16 optional
			if ie.EnableDefaultBeamPL_ForPUCCH_r16 != nil {
				if err = ie.EnableDefaultBeamPL_ForPUCCH_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode EnableDefaultBeamPL_ForPUCCH_r16", err)
				}
			}
			// encode EnableDefaultBeamPL_ForSRS_r16 optional
			if ie.EnableDefaultBeamPL_ForSRS_r16 != nil {
				if err = ie.EnableDefaultBeamPL_ForSRS_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode EnableDefaultBeamPL_ForSRS_r16", err)
				}
			}
			// encode UplinkTxSwitching_r16 optional
			if ie.UplinkTxSwitching_r16 != nil {
				tmp_UplinkTxSwitching_r16 := utils.SetupRelease[*UplinkTxSwitching_r16]{
					Setup: ie.UplinkTxSwitching_r16,
				}
				if err = tmp_UplinkTxSwitching_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode UplinkTxSwitching_r16", err)
				}
			}
			// encode Mpr_PowerBoost_FR2_r16 optional
			if ie.Mpr_PowerBoost_FR2_r16 != nil {
				if err = ie.Mpr_PowerBoost_FR2_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Mpr_PowerBoost_FR2_r16", err)
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

func (ie *UplinkConfig) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var InitialUplinkBWPPresent bool
	if InitialUplinkBWPPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var UplinkBWP_ToReleaseListPresent bool
	if UplinkBWP_ToReleaseListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var UplinkBWP_ToAddModListPresent bool
	if UplinkBWP_ToAddModListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var FirstActiveUplinkBWP_IdPresent bool
	if FirstActiveUplinkBWP_IdPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Pusch_ServingCellConfigPresent bool
	if Pusch_ServingCellConfigPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var CarrierSwitchingPresent bool
	if CarrierSwitchingPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if InitialUplinkBWPPresent {
		ie.InitialUplinkBWP = new(BWP_UplinkDedicated)
		if err = ie.InitialUplinkBWP.Decode(r); err != nil {
			return utils.WrapError("Decode InitialUplinkBWP", err)
		}
	}
	if UplinkBWP_ToReleaseListPresent {
		tmp_UplinkBWP_ToReleaseList := utils.NewSequence[*BWP_Id]([]*BWP_Id{}, uper.Constraint{Lb: 1, Ub: maxNrofBWPs}, false)
		fn_UplinkBWP_ToReleaseList := func() *BWP_Id {
			return new(BWP_Id)
		}
		if err = tmp_UplinkBWP_ToReleaseList.Decode(r, fn_UplinkBWP_ToReleaseList); err != nil {
			return utils.WrapError("Decode UplinkBWP_ToReleaseList", err)
		}
		ie.UplinkBWP_ToReleaseList = []BWP_Id{}
		for _, i := range tmp_UplinkBWP_ToReleaseList.Value {
			ie.UplinkBWP_ToReleaseList = append(ie.UplinkBWP_ToReleaseList, *i)
		}
	}
	if UplinkBWP_ToAddModListPresent {
		tmp_UplinkBWP_ToAddModList := utils.NewSequence[*BWP_Uplink]([]*BWP_Uplink{}, uper.Constraint{Lb: 1, Ub: maxNrofBWPs}, false)
		fn_UplinkBWP_ToAddModList := func() *BWP_Uplink {
			return new(BWP_Uplink)
		}
		if err = tmp_UplinkBWP_ToAddModList.Decode(r, fn_UplinkBWP_ToAddModList); err != nil {
			return utils.WrapError("Decode UplinkBWP_ToAddModList", err)
		}
		ie.UplinkBWP_ToAddModList = []BWP_Uplink{}
		for _, i := range tmp_UplinkBWP_ToAddModList.Value {
			ie.UplinkBWP_ToAddModList = append(ie.UplinkBWP_ToAddModList, *i)
		}
	}
	if FirstActiveUplinkBWP_IdPresent {
		ie.FirstActiveUplinkBWP_Id = new(BWP_Id)
		if err = ie.FirstActiveUplinkBWP_Id.Decode(r); err != nil {
			return utils.WrapError("Decode FirstActiveUplinkBWP_Id", err)
		}
	}
	if Pusch_ServingCellConfigPresent {
		tmp_Pusch_ServingCellConfig := utils.SetupRelease[*PUSCH_ServingCellConfig]{}
		if err = tmp_Pusch_ServingCellConfig.Decode(r); err != nil {
			return utils.WrapError("Decode Pusch_ServingCellConfig", err)
		}
		ie.Pusch_ServingCellConfig = tmp_Pusch_ServingCellConfig.Setup
	}
	if CarrierSwitchingPresent {
		tmp_CarrierSwitching := utils.SetupRelease[*SRS_CarrierSwitching]{}
		if err = tmp_CarrierSwitching.Decode(r); err != nil {
			return utils.WrapError("Decode CarrierSwitching", err)
		}
		ie.CarrierSwitching = tmp_CarrierSwitching.Setup
	}

	if extensionBit {
		// Read extension bitmap: 2 bits for 2 extension groups
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

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			PowerBoostPi2BPSKPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			UplinkChannelBW_PerSCS_ListPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode PowerBoostPi2BPSK optional
			if PowerBoostPi2BPSKPresent {
				var tmp_bool_PowerBoostPi2BPSK bool
				if tmp_bool_PowerBoostPi2BPSK, err = extReader.ReadBoolean(); err != nil {
					return utils.WrapError("Decode PowerBoostPi2BPSK", err)
				}
				ie.PowerBoostPi2BPSK = &tmp_bool_PowerBoostPi2BPSK
			}
			// decode UplinkChannelBW_PerSCS_List optional
			if UplinkChannelBW_PerSCS_ListPresent {
				tmp_UplinkChannelBW_PerSCS_List := utils.NewSequence[*SCS_SpecificCarrier]([]*SCS_SpecificCarrier{}, uper.Constraint{Lb: 1, Ub: maxSCSs}, false)
				fn_UplinkChannelBW_PerSCS_List := func() *SCS_SpecificCarrier {
					return new(SCS_SpecificCarrier)
				}
				if err = tmp_UplinkChannelBW_PerSCS_List.Decode(extReader, fn_UplinkChannelBW_PerSCS_List); err != nil {
					return utils.WrapError("Decode UplinkChannelBW_PerSCS_List", err)
				}
				ie.UplinkChannelBW_PerSCS_List = []SCS_SpecificCarrier{}
				for _, i := range tmp_UplinkChannelBW_PerSCS_List.Value {
					ie.UplinkChannelBW_PerSCS_List = append(ie.UplinkChannelBW_PerSCS_List, *i)
				}
			}
		}
		// decode extension group 2
		if len(extBitmap) > 1 && extBitmap[1] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			EnablePL_RS_UpdateForPUSCH_SRS_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			EnableDefaultBeamPL_ForPUSCH0_0_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			EnableDefaultBeamPL_ForPUCCH_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			EnableDefaultBeamPL_ForSRS_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			UplinkTxSwitching_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Mpr_PowerBoost_FR2_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode EnablePL_RS_UpdateForPUSCH_SRS_r16 optional
			if EnablePL_RS_UpdateForPUSCH_SRS_r16Present {
				ie.EnablePL_RS_UpdateForPUSCH_SRS_r16 = new(UplinkConfig_enablePL_RS_UpdateForPUSCH_SRS_r16)
				if err = ie.EnablePL_RS_UpdateForPUSCH_SRS_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode EnablePL_RS_UpdateForPUSCH_SRS_r16", err)
				}
			}
			// decode EnableDefaultBeamPL_ForPUSCH0_0_r16 optional
			if EnableDefaultBeamPL_ForPUSCH0_0_r16Present {
				ie.EnableDefaultBeamPL_ForPUSCH0_0_r16 = new(UplinkConfig_enableDefaultBeamPL_ForPUSCH0_0_r16)
				if err = ie.EnableDefaultBeamPL_ForPUSCH0_0_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode EnableDefaultBeamPL_ForPUSCH0_0_r16", err)
				}
			}
			// decode EnableDefaultBeamPL_ForPUCCH_r16 optional
			if EnableDefaultBeamPL_ForPUCCH_r16Present {
				ie.EnableDefaultBeamPL_ForPUCCH_r16 = new(UplinkConfig_enableDefaultBeamPL_ForPUCCH_r16)
				if err = ie.EnableDefaultBeamPL_ForPUCCH_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode EnableDefaultBeamPL_ForPUCCH_r16", err)
				}
			}
			// decode EnableDefaultBeamPL_ForSRS_r16 optional
			if EnableDefaultBeamPL_ForSRS_r16Present {
				ie.EnableDefaultBeamPL_ForSRS_r16 = new(UplinkConfig_enableDefaultBeamPL_ForSRS_r16)
				if err = ie.EnableDefaultBeamPL_ForSRS_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode EnableDefaultBeamPL_ForSRS_r16", err)
				}
			}
			// decode UplinkTxSwitching_r16 optional
			if UplinkTxSwitching_r16Present {
				tmp_UplinkTxSwitching_r16 := utils.SetupRelease[*UplinkTxSwitching_r16]{}
				if err = tmp_UplinkTxSwitching_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode UplinkTxSwitching_r16", err)
				}
				ie.UplinkTxSwitching_r16 = tmp_UplinkTxSwitching_r16.Setup
			}
			// decode Mpr_PowerBoost_FR2_r16 optional
			if Mpr_PowerBoost_FR2_r16Present {
				ie.Mpr_PowerBoost_FR2_r16 = new(UplinkConfig_mpr_PowerBoost_FR2_r16)
				if err = ie.Mpr_PowerBoost_FR2_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Mpr_PowerBoost_FR2_r16", err)
				}
			}
		}
	}
	return nil
}
