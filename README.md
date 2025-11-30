## RRC Library for Open RAN UE↔DU/CU Messaging

This Go module packages ASN.1-derived RRC message definitions and helpers so a User Equipment (UE) implementation can construct, encode, and decode signaling exchanged with the Distributed Unit (DU) and interpreted at the Centralized Unit (CU) inside an Open RAN architecture. It keeps the UE-side control-plane logic aligned with the DU/CU expectations by reusing the same IE layouts and APER encoding rules.

**Note:** Currently uses APER (Aligned Packed Encoding Rules). Migration to UPER (Unaligned Packed Encoding Rules) is planned for the future.

## Installation

```bash
go get github.com/lvdund/rrc
```

## Quick Start

### Encoding RRC Messages

All RRC messages must be wrapped in a container type (UL-CCCH, UL-DCCH, DL-CCCH, DL-DCCH, BCCH-BCH, BCCH-DL-SCH, or PCCH). Here's an example of encoding an RRC Setup Request:

```go
package main

import (
    "fmt"
    "github.com/lvdund/asn1go/aper"
    "github.com/lvdund/rrc"
    "github.com/lvdund/rrc/ies"
)

func main() {
    // Create an RRC Setup Request message
    rrcSetupRequest := ies.RRCSetupRequest{
        RrcSetupRequest: ies.RRCSetupRequest_IEs{
            Ue_Identity: ies.InitialUE_Identity{
                Choice: ies.InitialUE_Identity_Choice_Ng_5G_S_TMSI_Part1,
                Ng_5G_S_TMSI_Part1: aper.BitString{
                    Bytes:   []byte{0x07, 0xFF, 0xFF, 0xFF, 0xFF},
                    NumBits: 39,
                },
            },
            EstablishmentCause: ies.EstablishmentCause{
                Value: ies.EstablishmentCause_Enum_mo_Signalling,
            },
            Spare: aper.BitString{
                Bytes:   []byte{1},
                NumBits: 1,
            },
        },
    }

    // Wrap in UL-CCCH container (used before RRC connection exists)
    ulccchMessage := &ies.UL_CCCH_Message{
        Message: ies.UL_CCCH_MessageType{
            Choice: ies.UL_CCCH_MessageType_Choice_C1,
            C1: &ies.UL_CCCH_MessageType_C1{
                Choice:          ies.UL_CCCH_MessageType_C1_Choice_RrcSetupRequest,
                RrcSetupRequest: &rrcSetupRequest,
            },
        },
    }

    // Encode the message
    encoded, err := rrc.Encode(ulccchMessage)
    if err != nil {
        fmt.Printf("Failed to encode: %v\n", err)
        return
    }

    fmt.Printf("Encoded message: %x\n", encoded)
}
```

### Decoding RRC Messages

The library provides three decoding methods depending on your use case:

#### 1. Decode When You Know the Container Type

If you know which container type the message uses (e.g., from the channel context), use `Decode`:

```go
// Decode a UL-CCCH message
decoded := &ies.UL_CCCH_Message{}
if err := rrc.Decode(encodedBytes, decoded); err != nil {
    fmt.Printf("Failed to decode: %v\n", err)
    return
}

// Access the message
if decoded.Message.C1 != nil {
    rrcMsg := decoded.Message.C1.RrcSetupRequest
    fmt.Printf("Decoded RRC Setup Request\n")
}
```

#### 2. Decode When You Know the Channel Type

If you know the channel type but want a more convenient API, use `DecodeWithChannel`:

```go
msg, err := rrc.DecodeWithChannel(encodedBytes, rrc.MessageContainerTypeUL_CCCH)
if err != nil {
    fmt.Printf("Failed to decode: %v\n", err)
    return
}

// Type assert to access specific message
if ulccch, ok := msg.(*ies.UL_CCCH_Message); ok {
    // Access UL-CCCH specific fields
    fmt.Printf("Decoded UL-CCCH message\n")
}
```

#### 3. Decode When Container Type is Unknown

If you receive bytes but don't know which container type it is, use `DecodeAny`:

```go
decoded, err := rrc.DecodeAny(encodedBytes)
if err != nil {
    fmt.Printf("Failed to decode: %v\n", err)
    return
}

fmt.Printf("Detected container type: %s\n", decoded.Type)

// Type assert to access the specific message
switch msg := decoded.Message.(type) {
case *ies.UL_CCCH_Message:
    fmt.Printf("UL-CCCH message decoded\n")
    // Access UL-CCCH fields
case *ies.UL_DCCH_Message:
    fmt.Printf("UL-DCCH message decoded\n")
    // Access UL-DCCH fields
case *ies.DL_CCCH_Message:
    fmt.Printf("DL-CCCH message decoded\n")
    // Access DL-CCCH fields
case *ies.DL_DCCH_Message:
    fmt.Printf("DL-DCCH message decoded\n")
    // Access DL-DCCH fields
case *ies.BCCH_BCH_Message:
    fmt.Printf("BCCH-BCH message decoded\n")
    // Access BCCH-BCH fields
case *ies.BCCH_DL_SCH_Message:
    fmt.Printf("BCCH-DL-SCH message decoded\n")
    // Access BCCH-DL-SCH fields
case *ies.PCCH_Message:
    fmt.Printf("PCCH message decoded\n")
    // Access PCCH fields
}
```

## RRC Message Container Types

RRC messages are organized into 7 container types based on direction and channel:

| Container Type | Direction | Channel | Usage Example |
|----------------|-----------|---------|---------------|
| **UL-CCCH** | Uplink | Common Control | `RRCSetupRequest`, `RRCReestablishmentRequest`... |
| **UL-DCCH** | Uplink | Dedicated Control | `RRCSetupComplete`, `RRCReconfigurationComplete`... |
| **DL-CCCH** | Downlink | Common Control | `RRCSetup`, `RRCReestablishment`... |
| **DL-DCCH** | Downlink | Dedicated Control | `RRCReconfiguration`, `SecurityModeCommand`... |
| **BCCH-BCH** | Broadcast | Master Information | `MIB` (Master Information Block)... |
| **BCCH-DL-SCH** | Broadcast | System Information | `SystemInformationBlockType1` (SIB1)... |
| **PCCH** | Broadcast | Paging | `Paging` records... |

## Complete Example: RRC Setup Complete

```go
package main

import (
    "fmt"
    "github.com/lvdund/asn1go/aper"
    "github.com/lvdund/rrc"
    "github.com/lvdund/rrc/ies"
)

func main() {
    // Create RRC Setup Complete message
    rrcSetupComplete := ies.RRCSetupComplete{
        Rrc_TransactionIdentifier: ies.RRC_TransactionIdentifier{Value: 0},
        CriticalExtensions: ies.RRCSetupComplete_CriticalExtensions{
            Choice: ies.RRCSetupComplete_CriticalExtensions_Choice_RrcSetupComplete,
            RrcSetupComplete: &ies.RRCSetupComplete_IEs{
                SelectedPLMN_Identity: 1,
                DedicatedNAS_Message: ies.DedicatedNAS_Message{
                    Value: []byte{0x7e, 0x00, 0x41, 0x79, 0x00, 0x0d},
                },
            },
        },
    }

    // Wrap in UL-DCCH container (used after RRC connection is set up)
    uldcchMessage := &ies.UL_DCCH_Message{
        Message: ies.UL_DCCH_MessageType{
            Choice: ies.UL_DCCH_MessageType_Choice_C1,
            C1: &ies.UL_DCCH_MessageType_C1{
                Choice:           ies.UL_DCCH_MessageType_C1_Choice_RrcSetupComplete,
                RrcSetupComplete: &rrcSetupComplete,
            },
        },
    }

    // Encode
    encoded, err := rrc.Encode(uldcchMessage)
    if err != nil {
        fmt.Printf("Failed to encode: %v\n", err)
        return
    }

    // Decode using DecodeAny (when container type is unknown)
    decoded, err := rrc.DecodeAny(encoded)
    if err != nil {
        fmt.Printf("Failed to decode: %v\n", err)
        return
    }

    fmt.Printf("Detected container: %s\n", decoded.Type)

    // Access the decoded message
    if uldcch, ok := decoded.Message.(*ies.UL_DCCH_Message); ok {
        if uldcch.Message.C1 != nil && uldcch.Message.C1.RrcSetupComplete != nil {
            setupComplete := uldcch.Message.C1.RrcSetupComplete.CriticalExtensions.RrcSetupComplete
            fmt.Printf("Selected PLMN Identity: %d\n", setupComplete.SelectedPLMN_Identity)
            fmt.Printf("NAS Message: %x\n", setupComplete.DedicatedNAS_Message.Value)
        }
    }
}
```

## API Reference

### Encoding

- **`rrc.Encode(msg RRCMessage) ([]byte, error)`** - Encodes an RRC message container into bytes

### Decoding

- **`rrc.Decode(buf []byte, msg RRCMessage) error`** - Decodes bytes into a known message type
- **`rrc.DecodeAny(buf []byte) (*DecodedMessage, error)`** - Automatically detects and decodes the container type
- **`rrc.DecodeWithChannel(buf []byte, containerType MessageContainerType) (RRCMessage, error)`** - Decodes with a known channel type

### Types

- **`MessageContainerType`** - Enum representing container types (UL-CCCH, UL-DCCH, DL-CCCH, DL-DCCH, BCCH-BCH, BCCH-DL-SCH, PCCH)
- **`DecodedMessage`** - Wraps a decoded message with its detected container type

## Notes

- All RRC messages must be wrapped in one of the 7 container types
- The library uses APER encoding (will migrate to UPER in the future)
- `DecodeAny` tries all container types sequentially, so it's less efficient than `Decode` or `DecodeWithChannel` when you know the type
- Use `DecodeAny` when receiving bytes from an unknown source, otherwise prefer the more specific decode methods for better performance
