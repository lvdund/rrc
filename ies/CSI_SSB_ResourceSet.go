package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type CSI_SSB_ResourceSet struct {
	Csi_SSB_ResourceSetId        CSI_SSB_ResourceSetId           `madatory`
	Csi_SSB_ResourceList         []SSB_Index                     `lb:1,ub:maxNrofCSI_SSB_ResourcePerSet,madatory`
	ServingAdditionalPCIList_r17 []ServingAdditionalPCIIndex_r17 `lb:1,ub:maxNrofCSI_SSB_ResourcePerSet,optional,ext-1`
}

func (ie *CSI_SSB_ResourceSet) Encode(w *aper.AperWriter) error {
	var err error
	hasExtensions := len(ie.ServingAdditionalPCIList_r17) > 0
	preambleBits := []bool{hasExtensions}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.Csi_SSB_ResourceSetId.Encode(w); err != nil {
		return utils.WrapError("Encode Csi_SSB_ResourceSetId", err)
	}
	tmp_Csi_SSB_ResourceList := utils.NewSequence[*SSB_Index]([]*SSB_Index{}, aper.Constraint{Lb: 1, Ub: maxNrofCSI_SSB_ResourcePerSet}, false)
	for _, i := range ie.Csi_SSB_ResourceList {
		tmp_Csi_SSB_ResourceList.Value = append(tmp_Csi_SSB_ResourceList.Value, &i)
	}
	if err = tmp_Csi_SSB_ResourceList.Encode(w); err != nil {
		return utils.WrapError("Encode Csi_SSB_ResourceList", err)
	}
	if hasExtensions {
		// Extension bitmap: 1 bits for 1 extension groups
		extBitmap := []bool{len(ie.ServingAdditionalPCIList_r17) > 0}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap CSI_SSB_ResourceSet", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := aper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{len(ie.ServingAdditionalPCIList_r17) > 0}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode ServingAdditionalPCIList_r17 optional
			if len(ie.ServingAdditionalPCIList_r17) > 0 {
				tmp_ServingAdditionalPCIList_r17 := utils.NewSequence[*ServingAdditionalPCIIndex_r17]([]*ServingAdditionalPCIIndex_r17{}, aper.Constraint{Lb: 1, Ub: maxNrofCSI_SSB_ResourcePerSet}, false)
				for _, i := range ie.ServingAdditionalPCIList_r17 {
					tmp_ServingAdditionalPCIList_r17.Value = append(tmp_ServingAdditionalPCIList_r17.Value, &i)
				}
				if err = tmp_ServingAdditionalPCIList_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode ServingAdditionalPCIList_r17", err)
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

func (ie *CSI_SSB_ResourceSet) Decode(r *aper.AperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.Csi_SSB_ResourceSetId.Decode(r); err != nil {
		return utils.WrapError("Decode Csi_SSB_ResourceSetId", err)
	}
	tmp_Csi_SSB_ResourceList := utils.NewSequence[*SSB_Index]([]*SSB_Index{}, aper.Constraint{Lb: 1, Ub: maxNrofCSI_SSB_ResourcePerSet}, false)
	fn_Csi_SSB_ResourceList := func() *SSB_Index {
		return new(SSB_Index)
	}
	if err = tmp_Csi_SSB_ResourceList.Decode(r, fn_Csi_SSB_ResourceList); err != nil {
		return utils.WrapError("Decode Csi_SSB_ResourceList", err)
	}
	ie.Csi_SSB_ResourceList = []SSB_Index{}
	for _, i := range tmp_Csi_SSB_ResourceList.Value {
		ie.Csi_SSB_ResourceList = append(ie.Csi_SSB_ResourceList, *i)
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

			ServingAdditionalPCIList_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode ServingAdditionalPCIList_r17 optional
			if ServingAdditionalPCIList_r17Present {
				tmp_ServingAdditionalPCIList_r17 := utils.NewSequence[*ServingAdditionalPCIIndex_r17]([]*ServingAdditionalPCIIndex_r17{}, aper.Constraint{Lb: 1, Ub: maxNrofCSI_SSB_ResourcePerSet}, false)
				fn_ServingAdditionalPCIList_r17 := func() *ServingAdditionalPCIIndex_r17 {
					return new(ServingAdditionalPCIIndex_r17)
				}
				if err = tmp_ServingAdditionalPCIList_r17.Decode(extReader, fn_ServingAdditionalPCIList_r17); err != nil {
					return utils.WrapError("Decode ServingAdditionalPCIList_r17", err)
				}
				ie.ServingAdditionalPCIList_r17 = []ServingAdditionalPCIIndex_r17{}
				for _, i := range tmp_ServingAdditionalPCIList_r17.Value {
					ie.ServingAdditionalPCIList_r17 = append(ie.ServingAdditionalPCIList_r17, *i)
				}
			}
		}
	}
	return nil
}
