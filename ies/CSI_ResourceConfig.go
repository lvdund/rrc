package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type CSI_ResourceConfig struct {
	Csi_ResourceConfigId           CSI_ResourceConfigId                       `madatory`
	Csi_RS_ResourceSetList         *CSI_ResourceConfig_csi_RS_ResourceSetList `lb:1,ub:maxNrofNZP_CSI_RS_ResourceSetsPerConfig,optional`
	Bwp_Id                         BWP_Id                                     `madatory`
	ResourceType                   CSI_ResourceConfig_resourceType            `madatory`
	Csi_SSB_ResourceSetListExt_r17 *CSI_SSB_ResourceSetId                     `optional,ext-1`
}

func (ie *CSI_ResourceConfig) Encode(w *aper.AperWriter) error {
	var err error
	hasExtensions := ie.Csi_SSB_ResourceSetListExt_r17 != nil
	preambleBits := []bool{hasExtensions, ie.Csi_RS_ResourceSetList != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.Csi_ResourceConfigId.Encode(w); err != nil {
		return utils.WrapError("Encode Csi_ResourceConfigId", err)
	}
	if ie.Csi_RS_ResourceSetList != nil {
		if err = ie.Csi_RS_ResourceSetList.Encode(w); err != nil {
			return utils.WrapError("Encode Csi_RS_ResourceSetList", err)
		}
	}
	if err = ie.Bwp_Id.Encode(w); err != nil {
		return utils.WrapError("Encode Bwp_Id", err)
	}
	if err = ie.ResourceType.Encode(w); err != nil {
		return utils.WrapError("Encode ResourceType", err)
	}
	if hasExtensions {
		// Extension bitmap: 1 bits for 1 extension groups
		extBitmap := []bool{ie.Csi_SSB_ResourceSetListExt_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap CSI_ResourceConfig", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := aper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.Csi_SSB_ResourceSetListExt_r17 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode Csi_SSB_ResourceSetListExt_r17 optional
			if ie.Csi_SSB_ResourceSetListExt_r17 != nil {
				if err = ie.Csi_SSB_ResourceSetListExt_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Csi_SSB_ResourceSetListExt_r17", err)
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

func (ie *CSI_ResourceConfig) Decode(r *aper.AperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var Csi_RS_ResourceSetListPresent bool
	if Csi_RS_ResourceSetListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.Csi_ResourceConfigId.Decode(r); err != nil {
		return utils.WrapError("Decode Csi_ResourceConfigId", err)
	}
	if Csi_RS_ResourceSetListPresent {
		ie.Csi_RS_ResourceSetList = new(CSI_ResourceConfig_csi_RS_ResourceSetList)
		if err = ie.Csi_RS_ResourceSetList.Decode(r); err != nil {
			return utils.WrapError("Decode Csi_RS_ResourceSetList", err)
		}
	}
	if err = ie.Bwp_Id.Decode(r); err != nil {
		return utils.WrapError("Decode Bwp_Id", err)
	}
	if err = ie.ResourceType.Decode(r); err != nil {
		return utils.WrapError("Decode ResourceType", err)
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

			Csi_SSB_ResourceSetListExt_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode Csi_SSB_ResourceSetListExt_r17 optional
			if Csi_SSB_ResourceSetListExt_r17Present {
				ie.Csi_SSB_ResourceSetListExt_r17 = new(CSI_SSB_ResourceSetId)
				if err = ie.Csi_SSB_ResourceSetListExt_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Csi_SSB_ResourceSetListExt_r17", err)
				}
			}
		}
	}
	return nil
}
