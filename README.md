## RRC Library for Open RAN UE↔DU/CU Messaging

This Go module packages ASN.1-derived RRC message definitions and helpers so a User Equipment (UE) implementation can construct, encode, and decode signaling exchanged with the Distributed Unit (DU) and interpreted at the Centralized Unit (CU) inside an Open RAN architecture. It keeps the UE-side control-plane logic aligned with the DU/CU expectations by reusing the same IE layouts and UPER encoding rules.

## Usage

1. **Install the module**

   ```bash
   go get github.com/lvdund/rrc
   ```

2. **Import the packages you need** (core helpers plus the generated IE sets):

   ```go
   import (
       "github.com/lvdund/rrc"
       "github.com/lvdund/rrc/ies"
       "github.com/lvdund/asn1go/uper"
   )
   ```

3. **Wrap your IE in the `rrc.RRCMessage` interface** so it can be serialized and parsed:

   ```go
   type RRCSetupComplete struct {
       ies.RRCSetupComplete
   }

   func (m *RRCSetupComplete) Encode(w *uper.UperWriter) error {
       return m.RRCSetupComplete.Encode(w)
   }

   func (m *RRCSetupComplete) Decode(r *uper.UperReader) error {
       return m.RRCSetupComplete.Decode(r)
   }
   ```

4. **Call the helpers from UE signaling handlers when talking to the DU/CU:**

   ```go
   payload, err := rrc.Encode(&RRCSetupComplete{ /* set IE fields */ })
   if err != nil {
       // handle encoding error
   }

   msg, err := rrc.Decode(payload)
   if err != nil {
       // handle decoding error
   }
   ```

   `rrc.Encode` and `rrc.Decode` manage the UPER writer/reader plumbing so that UE-originated RRC PDUs remain spec-compliant when exchanged with the DU and later consumed by the CU.
