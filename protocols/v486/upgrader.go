package v486

import (
	"github.com/potlounge/multiversion/protocols/v486/types"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// upgradeBlockActorData upgrades a block actor from a legacy version to the latest version.
func upgradeBlockActorData(data map[string]any) map[string]any {
	switch data["id"] {
	case "Sign":
		textRaw, ok := data["Text"]
		if !ok {
			textRaw = ""
		}
		text, ok := textRaw.(string)
		if !ok {
			text = ""
		}
		data["FrontText"] = map[string]any{
			"IgnoreLighting":    uint8(0),
			"PersistFormatting": uint8(1),
			"SignTextColor":     int32(-16777216),
			"Text":              text,
		}
		data["BackText"] = map[string]any{
			"IgnoreLighting":    uint8(0),
			"PersistFormatting": uint8(1),
			"SignTextColor":     int32(-16777216),
			"Text":              "",
		}
	}
	return data
}

// upgradeEntityMetadata upgrades entity metadata from legacy version to latest version.
func upgradeEntityMetadata(data map[uint32]any) map[uint32]any {
	var flag1, flag2 int64
	if v, ok := data[protocol.EntityDataKeyFlags]; ok {
		flag1 = v.(int64)
	}
	if v, ok := data[protocol.EntityDataKeyFlagsTwo]; ok {
		flag2 = v.(int64)
	}

	flag2 <<= 1
	flag2 |= (flag1 >> 63) & 1

	newFlag1 := flag1 & ^(^0 << (protocol.EntityDataFlagDash - 1))
	lastHalf := flag1 & (^0 << (protocol.EntityDataFlagDash - 1))
	lastHalf <<= 1
	newFlag1 |= lastHalf

	data[protocol.EntityDataKeyFlagsTwo] = flag2
	data[protocol.EntityDataKeyFlags] = newFlag1
	return data
}

func upgradeCraftingDescription(descriptor *types.DefaultItemDescriptor) protocol.ItemDescriptor {
	return &protocol.DefaultItemDescriptor{
		NetworkID:     int16(descriptor.NetworkID),
		MetadataValue: int16(descriptor.MetadataValue),
	}
}

// TODO: add upgrade entity flags
