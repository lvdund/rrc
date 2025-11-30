package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type CSI_AssociatedReportConfigInfo struct {
	ReportConfigId                      CSI_ReportConfigId                                       `madatory`
	ResourcesForChannel                 *CSI_AssociatedReportConfigInfo_resourcesForChannel      `lb:1,ub:maxNrofAP_CSI_RS_ResourcesPerSet,optional`
	Csi_IM_ResourcesForInterference     *int64                                                   `lb:1,ub:maxNrofCSI_IM_ResourceSetsPerConfig,optional`
	Nzp_CSI_RS_ResourcesForInterference *int64                                                   `lb:1,ub:maxNrofNZP_CSI_RS_ResourceSetsPerConfig,optional`
	ResourcesForChannel2_r17            *CSI_AssociatedReportConfigInfo_resourcesForChannel2_r17 `lb:1,ub:maxNrofAP_CSI_RS_ResourcesPerSet,optional,ext-1`
	Csi_SSB_ResourceSetExt              *int64                                                   `lb:1,ub:maxNrofCSI_SSB_ResourceSetsPerConfigExt,optional,ext-1`
}

func (ie *CSI_AssociatedReportConfigInfo) Encode(w *aper.AperWriter) error {
	var err error
	hasExtensions := ie.ResourcesForChannel2_r17 != nil || ie.Csi_SSB_ResourceSetExt != nil
	preambleBits := []bool{hasExtensions, ie.ResourcesForChannel != nil, ie.Csi_IM_ResourcesForInterference != nil, ie.Nzp_CSI_RS_ResourcesForInterference != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.ReportConfigId.Encode(w); err != nil {
		return utils.WrapError("Encode ReportConfigId", err)
	}
	if ie.ResourcesForChannel != nil {
		if err = ie.ResourcesForChannel.Encode(w); err != nil {
			return utils.WrapError("Encode ResourcesForChannel", err)
		}
	}
	if ie.Csi_IM_ResourcesForInterference != nil {
		if err = w.WriteInteger(*ie.Csi_IM_ResourcesForInterference, &aper.Constraint{Lb: 1, Ub: maxNrofCSI_IM_ResourceSetsPerConfig}, false); err != nil {
			return utils.WrapError("Encode Csi_IM_ResourcesForInterference", err)
		}
	}
	if ie.Nzp_CSI_RS_ResourcesForInterference != nil {
		if err = w.WriteInteger(*ie.Nzp_CSI_RS_ResourcesForInterference, &aper.Constraint{Lb: 1, Ub: maxNrofNZP_CSI_RS_ResourceSetsPerConfig}, false); err != nil {
			return utils.WrapError("Encode Nzp_CSI_RS_ResourcesForInterference", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 1 bits for 1 extension groups
		extBitmap := []bool{ie.ResourcesForChannel2_r17 != nil || ie.Csi_SSB_ResourceSetExt != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap CSI_AssociatedReportConfigInfo", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := aper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.ResourcesForChannel2_r17 != nil, ie.Csi_SSB_ResourceSetExt != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode ResourcesForChannel2_r17 optional
			if ie.ResourcesForChannel2_r17 != nil {
				if err = ie.ResourcesForChannel2_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode ResourcesForChannel2_r17", err)
				}
			}
			// encode Csi_SSB_ResourceSetExt optional
			if ie.Csi_SSB_ResourceSetExt != nil {
				if err = extWriter.WriteInteger(*ie.Csi_SSB_ResourceSetExt, &aper.Constraint{Lb: 1, Ub: maxNrofCSI_SSB_ResourceSetsPerConfigExt}, false); err != nil {
					return utils.WrapError("Encode Csi_SSB_ResourceSetExt", err)
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

func (ie *CSI_AssociatedReportConfigInfo) Decode(r *aper.AperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var ResourcesForChannelPresent bool
	if ResourcesForChannelPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Csi_IM_ResourcesForInterferencePresent bool
	if Csi_IM_ResourcesForInterferencePresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Nzp_CSI_RS_ResourcesForInterferencePresent bool
	if Nzp_CSI_RS_ResourcesForInterferencePresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.ReportConfigId.Decode(r); err != nil {
		return utils.WrapError("Decode ReportConfigId", err)
	}
	if ResourcesForChannelPresent {
		ie.ResourcesForChannel = new(CSI_AssociatedReportConfigInfo_resourcesForChannel)
		if err = ie.ResourcesForChannel.Decode(r); err != nil {
			return utils.WrapError("Decode ResourcesForChannel", err)
		}
	}
	if Csi_IM_ResourcesForInterferencePresent {
		var tmp_int_Csi_IM_ResourcesForInterference int64
		if tmp_int_Csi_IM_ResourcesForInterference, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: maxNrofCSI_IM_ResourceSetsPerConfig}, false); err != nil {
			return utils.WrapError("Decode Csi_IM_ResourcesForInterference", err)
		}
		ie.Csi_IM_ResourcesForInterference = &tmp_int_Csi_IM_ResourcesForInterference
	}
	if Nzp_CSI_RS_ResourcesForInterferencePresent {
		var tmp_int_Nzp_CSI_RS_ResourcesForInterference int64
		if tmp_int_Nzp_CSI_RS_ResourcesForInterference, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: maxNrofNZP_CSI_RS_ResourceSetsPerConfig}, false); err != nil {
			return utils.WrapError("Decode Nzp_CSI_RS_ResourcesForInterference", err)
		}
		ie.Nzp_CSI_RS_ResourcesForInterference = &tmp_int_Nzp_CSI_RS_ResourcesForInterference
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

			ResourcesForChannel2_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Csi_SSB_ResourceSetExtPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode ResourcesForChannel2_r17 optional
			if ResourcesForChannel2_r17Present {
				ie.ResourcesForChannel2_r17 = new(CSI_AssociatedReportConfigInfo_resourcesForChannel2_r17)
				if err = ie.ResourcesForChannel2_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode ResourcesForChannel2_r17", err)
				}
			}
			// decode Csi_SSB_ResourceSetExt optional
			if Csi_SSB_ResourceSetExtPresent {
				var tmp_int_Csi_SSB_ResourceSetExt int64
				if tmp_int_Csi_SSB_ResourceSetExt, err = extReader.ReadInteger(&aper.Constraint{Lb: 1, Ub: maxNrofCSI_SSB_ResourceSetsPerConfigExt}, false); err != nil {
					return utils.WrapError("Decode Csi_SSB_ResourceSetExt", err)
				}
				ie.Csi_SSB_ResourceSetExt = &tmp_int_Csi_SSB_ResourceSetExt
			}
		}
	}
	return nil
}
