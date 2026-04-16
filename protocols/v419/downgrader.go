package v419

import (
	"github.com/potlounge/multiversion/mapping"
	"github.com/potlounge/multiversion/protocols/v419/types"
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
	oldData := data
	for index, value := range oldData {
		switch index {
		case protocol.EntityDataKeyDataRadius:
			index = 60
		case protocol.EntityDataKeyDataWaiting:
			index = 61
		case protocol.EntityDataKeyDataParticle:
			index = 62
		case protocol.EntityDataKeyAttachFace:
			index = 64
		case protocol.EntityDataKeyAttachedPosition:
			index = 66
		case protocol.EntityDataKeyTradeTarget:
			index = 67
		case protocol.EntityDataKeyCommandName:
			index = 70
		case protocol.EntityDataKeyLastCommandOutput:
			index = 71
		case protocol.EntityDataKeyTrackCommandOutput:
			index = 72
		case protocol.EntityDataKeyControllingSeatIndex:
			index = 73
		case protocol.EntityDataKeyStrength:
			index = 74
		case protocol.EntityDataKeyStrengthMax:
			index = 75
		case protocol.EntityDataKeyDataLifetimeTicks:
			index = 77
		case protocol.EntityDataKeyPoseIndex:
			index = 78
		case protocol.EntityDataKeyDataTickOffset:
			index = 79
		case protocol.EntityDataKeyAlwaysShowNameTag:
			index = 80
		case protocol.EntityDataKeyColorTwoIndex:
			index = 81
		case protocol.EntityDataKeyScore:
			index = 83
		case protocol.EntityDataKeyBalloonAnchor:
			index = 84
		case protocol.EntityDataKeyPuffedState:
			index = 85
		case protocol.EntityDataKeyBubbleTime:
			index = 86
		case protocol.EntityDataKeyAgent:
			index = 87
		case protocol.EntityDataKeyEatingCounter:
			index = 90
		case protocol.EntityDataKeyFlagsTwo:
			index = 91
		case protocol.EntityDataKeyDataDuration:
			index = 94
		case protocol.EntityDataKeyDataSpawnTime:
			index = 95
		case protocol.EntityDataKeyDataChangeRate:
			index = 96
		case protocol.EntityDataKeyDataChangeOnPickup:
			index = 97
		case protocol.EntityDataKeyDataPickupCount:
			index = 98
		case protocol.EntityDataKeyInteractText:
			index = 99
		case protocol.EntityDataKeyTradeTier:
			index = 100
		case protocol.EntityDataKeyMaxTradeTier:
			index = 101
		case protocol.EntityDataKeyTradeExperience:
			index = 102
		case protocol.EntityDataKeySkinID:
			index = 104
		case protocol.EntityDataKeyCommandBlockTickDelay:
			index = 105
		case protocol.EntityDataKeyCommandBlockExecuteOnFirstTick:
			index = 106
		case protocol.EntityDataKeyAmbientSoundInterval:
			index = 107
		case protocol.EntityDataKeyAmbientSoundIntervalRange:
			index = 108
		case protocol.EntityDataKeyAmbientSoundEventName:
			index = 109
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
