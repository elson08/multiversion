package v419

import (
	"github.com/potlounge/multiversion/protocols/v419/types"
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
	oldData := data
	for index, value := range oldData {
		switch index {
		case 60:
			index = protocol.EntityDataKeyDataRadius
		case 61:
			index = protocol.EntityDataKeyDataWaiting
		case 62:
			index = protocol.EntityDataKeyDataParticle
		case 64:
			index = protocol.EntityDataKeyAttachFace
		case 66:
			index = protocol.EntityDataKeyAttachedPosition
		case 67:
			index = protocol.EntityDataKeyTradeTarget
		case 70:
			index = protocol.EntityDataKeyCommandName
		case 71:
			index = protocol.EntityDataKeyLastCommandOutput
		case 72:
			index = protocol.EntityDataKeyTrackCommandOutput
		case 73:
			index = protocol.EntityDataKeyControllingSeatIndex
		case 74:
			index = protocol.EntityDataKeyStrength
		case 75:
			index = protocol.EntityDataKeyStrengthMax
		case 77:
			index = protocol.EntityDataKeyDataLifetimeTicks
		case 78:
			index = protocol.EntityDataKeyPoseIndex
		case 79:
			index = protocol.EntityDataKeyDataTickOffset
		case 80:
			index = protocol.EntityDataKeyAlwaysShowNameTag
		case 81:
			index = protocol.EntityDataKeyColorTwoIndex
		case 83:
			index = protocol.EntityDataKeyScore
		case 84:
			index = protocol.EntityDataKeyBalloonAnchor
		case 85:
			index = protocol.EntityDataKeyPuffedState
		case 86:
			index = protocol.EntityDataKeyBubbleTime
		case 87:
			index = protocol.EntityDataKeyAgent
		case 90:
			index = protocol.EntityDataKeyEatingCounter
		case 91:
			index = protocol.EntityDataKeyFlagsTwo
		case 94:
			index = protocol.EntityDataKeyDataDuration
		case 95:
			index = protocol.EntityDataKeyDataSpawnTime
		case 96:
			index = protocol.EntityDataKeyDataChangeRate
		case 97:
			index = protocol.EntityDataKeyDataChangeOnPickup
		case 98:
			index = protocol.EntityDataKeyDataPickupCount
		case 99:
			index = protocol.EntityDataKeyInteractText
		case 100:
			index = protocol.EntityDataKeyTradeTier
		case 101:
			index = protocol.EntityDataKeyMaxTradeTier
		case 102:
			index = protocol.EntityDataKeyTradeExperience
		case 104:
			index = protocol.EntityDataKeySkinID
		case 105:
			index = protocol.EntityDataKeyCommandBlockTickDelay
		case 106:
			index = protocol.EntityDataKeyCommandBlockExecuteOnFirstTick
		case 107:
			index = protocol.EntityDataKeyAmbientSoundInterval
		case 108:
			index = protocol.EntityDataKeyAmbientSoundIntervalRange
		case 109:
			index = protocol.EntityDataKeyAmbientSoundEventName
		}
		data[index] = value
	}

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

func upgradeLevelEvent(eventType int32) int32 {
	return eventType
}

// TODO: add upgrade entity flags
