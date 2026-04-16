package v486

import (
	"github.com/potlounge/multiversion/mapping"
	"github.com/potlounge/multiversion/protocols/v486/types"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"math"
)

// downgradeBlockActorData downgrades a block actor from latest version to legacy version.
func downgradeBlockActorData(data map[string]any) map[string]any {
	switch data["id"] {
	case "Sign":
		delete(data, "BackText")
		frontRaw, ok := data["FrontText"]
		if !ok {
			frontRaw = map[string]any{"Text": ""}
		}
		front, ok := frontRaw.(map[string]any)
		if !ok {
			front = map[string]any{"Text": ""}
		}
		textRaw, ok := front["Text"]
		if !ok {
			textRaw = ""
		}
		text, ok := textRaw.(string)
		if !ok {
			text = ""
		}
		data["Text"] = text
	}

	return data
}

// downgradeEntityMetadata downgrades entity metadata from latest version to legacy version.
func downgradeEntityMetadata(data map[uint32]any) map[uint32]any {
	var flag1, flag2 int64
	if v, ok := data[protocol.EntityDataKeyFlags]; ok {
		flag1 = v.(int64)
	}
	if v, ok := data[protocol.EntityDataKeyFlagsTwo]; ok {
		flag2 = v.(int64)
	}
	if flag1 == 0 && flag2 == 0 {
		return data
	}

	newFlag1 := flag1 & ^(^0 << (protocol.EntityDataFlagDash - 1))
	lastHalf := flag1 & (^0 << protocol.EntityDataFlagDash)
	lastHalf >>= 1
	lastHalf &= math.MaxInt64

	newFlag1 |= lastHalf

	if flag2 != 0 {
		newFlag1 ^= (flag2 & 1) << 63
		flag2 >>= 1
		flag2 &= math.MaxInt64

		data[protocol.EntityDataKeyFlagsTwo] = flag2
	}

	data[protocol.EntityDataKeyFlags] = newFlag1
	return data
}

func downgradeCraftingDescription(descriptor protocol.ItemDescriptor, m mapping.Item) protocol.ItemDescriptor {
	var networkId int32
	var metadata int32
	switch descriptor := descriptor.(type) {
	case *protocol.DefaultItemDescriptor:
		networkId = int32(descriptor.NetworkID)
		metadata = int32(descriptor.MetadataValue)
	case *protocol.DeferredItemDescriptor:
		if rid, ok := m.ItemNameToRuntimeID(descriptor.Name); ok {
			networkId = rid
			metadata = int32(descriptor.MetadataValue)
		}
	case *protocol.ItemTagItemDescriptor:
		/// ?????
	case *protocol.ComplexAliasItemDescriptor:
		/// ?????
	}
	return &types.DefaultItemDescriptor{
		NetworkID:     networkId,
		MetadataValue: metadata,
	}
}

// TODO: add downgrade entity flags
