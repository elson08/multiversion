package v419

import (
	_ "embed"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/samber/lo"
	"github.com/sandertv/gophertunnel/minecraft"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"

	"github.com/potlounge/multiversion/internal"
	"github.com/potlounge/multiversion/internal/chunk"
	"github.com/potlounge/multiversion/mapping"
	"github.com/potlounge/multiversion/protocols/latest"
	legacypacket "github.com/potlounge/multiversion/protocols/v419/packet"
	"github.com/potlounge/multiversion/protocols/v419/types"
	"github.com/potlounge/multiversion/translator"
)

var (
	//go:embed data/required_item_list.json
	requiredItemList []byte
	//go:embed data/block_states.nbt
	blockStateData []byte
	//go:embed data/entity_identifiers.nbt
	entityIdentifierData []byte
)

func init() {
	var err error
	entityIdentifierData, err = translator.TranslateEntityIdentifiers(entityIdentifierData)
	if err != nil {
		panic(err)
	}
}

const (
	// ItemVersion is the version of items of the game which use for downgrading and upgrading.
	ItemVersion = 21
	// BlockVersion is the version of blocks (states) of the game. This version is composed
	// of 4 bytes indicating a version, interpreted as a big endian int. The current version represents
	// 1.12.0.1 {1, 12, 0, 1}.
	BlockVersion int32 = (1 << 24) | (12 << 16) | (0 << 8) | 1
)

type Protocol struct {
	itemMapping     mapping.Item
	blockMapping    mapping.Block
	itemTranslator  translator.ItemTranslator
	blockTranslator translator.BlockTranslator
}

func New() *Protocol {
	itemMapping := mapping.NewItemMapping(requiredItemList, ItemVersion)
	blockMapping := mapping.NewBlockMapping(blockStateData).WithBlockActorRemapper(downgradeBlockActorData, upgradeBlockActorData)
	latestBlockMapping := latest.NewBlockMapping()
	return &Protocol{itemMapping: itemMapping, blockMapping: blockMapping,
		itemTranslator:  translator.NewItemTranslator(itemMapping, latest.NewItemMapping(), blockMapping, latestBlockMapping),
		blockTranslator: translator.NewBlockTranslator(blockMapping, latestBlockMapping, chunk.NewNetworkPersistentEncoding(blockMapping, BlockVersion), chunk.NewBlockPaletteEncoding(blockMapping, BlockVersion), true)}
}

func (p Protocol) ID() int32 {
	return 419
}

func (p Protocol) Ver() string {
	return "1.16.100"
}

func (Protocol) Packets(_ bool) packet.Pool {
	pool := packet.NewClientPool()
	for k, v := range packet.NewServerPool() {
		pool[k] = v
	}

	pool[packet.IDActorPickRequest] = func() packet.Packet { return &legacypacket.ActorPickRequest{} }
	pool[packet.IDAddActor] = func() packet.Packet { return &legacypacket.AddActor{} }
	pool[packet.IDAddPlayer] = func() packet.Packet { return &legacypacket.AddPlayer{} }
	pool[packet.IDAddVolumeEntity] = func() packet.Packet { return &legacypacket.AddVolumeEntity{} }
	pool[packet.IDAdventureSettings] = func() packet.Packet { return &legacypacket.AdventureSettings{} }
	pool[packet.IDAnimate] = func() packet.Packet { return &legacypacket.Animate{} }
	pool[packet.IDAnimateEntity] = func() packet.Packet { return &legacypacket.AnimateEntity{} }
	pool[packet.IDAvailableCommands] = func() packet.Packet { return &legacypacket.AvailableCommands{} }
	pool[packet.IDBiomeDefinitionList] = func() packet.Packet { return &legacypacket.BiomeDefinitionList{} }
	pool[packet.IDBossEvent] = func() packet.Packet { return &legacypacket.BossEvent{} }
	pool[packet.IDCameraInstruction] = func() packet.Packet { return &legacypacket.CameraInstruction{} }
	pool[packet.IDCameraShake] = func() packet.Packet { return &legacypacket.CameraShake{} }
	pool[packet.IDChangeDimension] = func() packet.Packet { return &legacypacket.ChangeDimension{} }
	pool[packet.IDCommandRequest] = func() packet.Packet { return &legacypacket.CommandRequest{} }
	pool[packet.IDContainerClose] = func() packet.Packet { return &legacypacket.ContainerClose{} }
	pool[packet.IDCorrectPlayerMovePrediction] = func() packet.Packet { return &legacypacket.CorrectPlayerMovePrediction{} }
	pool[packet.IDCraftingData] = func() packet.Packet { return &legacypacket.CraftingData{} }
	pool[packet.IDCreativeContent] = func() packet.Packet { return &legacypacket.CreativeContent{} }
	pool[packet.IDDisconnect] = func() packet.Packet { return &legacypacket.Disconnect{} }
	pool[packet.IDEmote] = func() packet.Packet { return &legacypacket.Emote{} }
	pool[packet.IDHurtArmour] = func() packet.Packet { return &legacypacket.HurtArmour{} }
	pool[packet.IDInventoryContent] = func() packet.Packet { return &legacypacket.InventoryContent{} }
	pool[packet.IDInventorySlot] = func() packet.Packet { return &legacypacket.InventorySlot{} }
	pool[packet.IDInventoryTransaction] = func() packet.Packet { return &legacypacket.InventoryTransaction{} }
	pool[packet.IDItemStackRequest] = func() packet.Packet { return &legacypacket.ItemStackRequest{} }
	pool[packet.IDItemStackResponse] = func() packet.Packet { return &legacypacket.ItemStackResponse{} }
	pool[packet.IDLevelChunk] = func() packet.Packet { return &legacypacket.LevelChunk{} }
	pool[packet.IDLevelSoundEvent] = func() packet.Packet { return &legacypacket.LevelSoundEvent{} }
	pool[packet.IDMobArmourEquipment] = func() packet.Packet { return &legacypacket.MobArmourEquipment{} }
	pool[packet.IDMobEffect] = func() packet.Packet { return &legacypacket.MobEffect{} }
	pool[packet.IDModalFormResponse] = func() packet.Packet { return &legacypacket.ModalFormResponse{} }
	pool[packet.IDNetworkChunkPublisherUpdate] = func() packet.Packet { return &legacypacket.NetworkChunkPublisherUpdate{} }
	pool[packet.IDNetworkSettings] = func() packet.Packet { return &legacypacket.NetworkSettings{} }
	pool[packet.IDPlayerAction] = func() packet.Packet { return &legacypacket.PlayerAction{} }
	pool[packet.IDPlayerArmourDamage] = func() packet.Packet { return &legacypacket.PlayerArmourDamage{} }
	pool[packet.IDPlayerAuthInput] = func() packet.Packet { return &legacypacket.PlayerAuthInput{} }
	pool[packet.IDPlayerList] = func() packet.Packet { return &legacypacket.PlayerList{} }
	pool[packet.IDPlayerSkin] = func() packet.Packet { return &legacypacket.PlayerSkin{} }
	pool[packet.IDRemoveVolumeEntity] = func() packet.Packet { return &legacypacket.RemoveVolumeEntity{} }
	pool[packet.IDRequestChunkRadius] = func() packet.Packet { return &legacypacket.RequestChunkRadius{} }
	pool[packet.IDResourcePackStack] = func() packet.Packet { return &legacypacket.ResourcePackStack{} }
	pool[packet.IDResourcePacksInfo] = func() packet.Packet { return &legacypacket.ResourcePacksInfo{} }
	pool[packet.IDSetActorLink] = func() packet.Packet { return &legacypacket.SetActorLink{} }
	pool[packet.IDSetActorData] = func() packet.Packet { return &legacypacket.SetActorData{} }
	pool[packet.IDSetActorMotion] = func() packet.Packet { return &legacypacket.SetActorMotion{} }
	pool[packet.IDSetTitle] = func() packet.Packet { return &legacypacket.SetTitle{} }
	pool[packet.IDSpawnParticleEffect] = func() packet.Packet { return &legacypacket.SpawnParticleEffect{} }
	pool[packet.IDStartGame] = func() packet.Packet { return &legacypacket.StartGame{} }
	pool[packet.IDStopSound] = func() packet.Packet { return &legacypacket.StopSound{} }
	pool[packet.IDStructureBlockUpdate] = func() packet.Packet { return &legacypacket.StructureBlockUpdate{} }
	pool[packet.IDStructureTemplateDataRequest] = func() packet.Packet { return &legacypacket.StructureTemplateDataRequest{} }
	pool[packet.IDText] = func() packet.Packet { return &legacypacket.Text{} }
	pool[packet.IDTransfer] = func() packet.Packet { return &legacypacket.Transfer{} }
	pool[packet.IDUpdateAttributes] = func() packet.Packet { return &legacypacket.UpdateAttributes{} }
	pool[packet.IDUpdatePlayerGameType] = func() packet.Packet { return &legacypacket.UpdatePlayerGameType{} }
	pool[packet.IDInteract] = func() packet.Packet { return &legacypacket.Interact{} }
	pool[packet.IDContainerOpen] = func() packet.Packet { return &legacypacket.ContainerOpen{} }
	pool[packet.IDUpdateBlock] = func() packet.Packet { return &legacypacket.UpdateBlock{} }
	pool[packet.IDBlockActorData] = func() packet.Packet { return &legacypacket.BlockActorData{} }
	pool[packet.IDAnvilDamage] = func() packet.Packet { return &legacypacket.AnvilDamage{} }
	pool[packet.IDBlockEvent] = func() packet.Packet { return &legacypacket.BlockEvent{} }
	pool[packet.IDUpdateBlockSynced] = func() packet.Packet { return &legacypacket.UpdateBlockSynced{} }

	pool[23] = func() packet.Packet { return &legacypacket.TickSync{} }
	pool[71] = func() packet.Packet { return &legacypacket.ItemFrameDropItem{} }

	return pool
}

func (Protocol) Encryption(key [32]byte) packet.Encryption {
	return newCFBEncryption(key[:])
}

func (Protocol) NewReader(r minecraft.ByteReader, shieldID int32, enableLimits bool) protocol.IO {
	return NewReader(protocol.NewReader(r, shieldID, enableLimits))
}

func (Protocol) NewWriter(w minecraft.ByteWriter, shieldID int32) protocol.IO {
	return NewWriter(protocol.NewWriter(w, shieldID))
}

func (p Protocol) ConvertToLatest(pk packet.Packet, conn *minecraft.Conn) []packet.Packet {
	// fmt.Printf("1.16.100 -> Latest: %T\n", pk)

	var newPks []packet.Packet
	switch pk := pk.(type) {
	case *legacypacket.Interact:
		interactPk := &packet.Interact{
			ActionType:            pk.ActionType,
			TargetEntityRuntimeID: pk.TargetEntityRuntimeID,
		}
		if pk.Position != (mgl32.Vec3{}) {
			interactPk.Position = protocol.Option(pk.Position)
		}
		newPks = append(newPks, interactPk)
	case *legacypacket.Transfer:
		newPks = append(newPks, &packet.Transfer{
			Address:     pk.Address,
			Port:        pk.Port,
			ReloadWorld: true,
		})
	case *legacypacket.ContainerOpen:
		newPks = append(newPks, &packet.ContainerOpen{
			WindowID:                pk.WindowID,
			ContainerType:           pk.ContainerType,
			ContainerPosition:       pk.ContainerPosition,
			ContainerEntityUniqueID: pk.ContainerEntityUniqueID,
		})
	case *legacypacket.UpdateBlock:
		newPks = append(newPks, &packet.UpdateBlock{
			Position:          pk.Position,
			NewBlockRuntimeID: pk.NewBlockRuntimeID,
			Flags:             pk.Flags,
			Layer:             pk.Layer,
		})
	case *legacypacket.BlockActorData:
		newPks = append(newPks, &packet.BlockActorData{
			Position: pk.Position,
			NBTData:  pk.NBTData,
		})
	case *legacypacket.AnvilDamage:
		newPks = append(newPks, &packet.AnvilDamage{
			Damage:        pk.Damage,
			AnvilPosition: pk.AnvilPosition,
		})
	case *legacypacket.BlockEvent:
		newPks = append(newPks, &packet.BlockEvent{
			Position:  pk.Position,
			EventType: pk.EventType,
			EventData: pk.EventData,
		})
	case *legacypacket.UpdateBlockSynced:
		newPks = append(newPks, &packet.UpdateBlockSynced{
			Position:          pk.Position,
			NewBlockRuntimeID: pk.NewBlockRuntimeID,
			Flags:             pk.Flags,
			Layer:             pk.Layer,
			EntityUniqueID:    pk.EntityUniqueID,
			TransitionType:    pk.TransitionType,
		})
	case *legacypacket.LevelSoundEvent:
		newPks = append(newPks, &packet.LevelSoundEvent{
			SoundType:             pk.SoundType,
			Position:              pk.Position,
			ExtraData:             pk.ExtraData,
			EntityType:            pk.EntityType,
			BabyMob:               pk.BabyMob,
			DisableRelativeVolume: pk.DisableRelativeVolume,
			EntityUniqueID:        conn.GameData().EntityUniqueID,
		})
	case *legacypacket.PlayerArmourDamage:
		var list []protocol.PlayerArmourDamageEntry
		if pk.Bitset&(1<<packet.PlayerArmourDamageFlagHelmet) != 0 {
			list = append(list, protocol.PlayerArmourDamageEntry{
				ArmourSlot: 0,
				Damage:     int16(pk.HelmetDamage),
			})
		}

		if pk.Bitset&(1<<packet.PlayerArmourDamageFlagChestplate) != 0 {
			list = append(list, protocol.PlayerArmourDamageEntry{
				ArmourSlot: 1,
				Damage:     int16(pk.ChestplateDamage),
			})
		}

		if pk.Bitset&(1<<packet.PlayerArmourDamageFlagLeggings) != 0 {
			list = append(list, protocol.PlayerArmourDamageEntry{
				ArmourSlot: 2,
				Damage:     int16(pk.LeggingsDamage),
			})
		}

		if pk.Bitset&(1<<packet.PlayerArmourDamageFlagBoots) != 0 {
			list = append(list, protocol.PlayerArmourDamageEntry{
				ArmourSlot: 3,
				Damage:     int16(pk.BootsDamage),
			})
		}

		newPks = append(newPks, &packet.PlayerArmourDamage{
			List: list,
		})
	case *legacypacket.MobArmourEquipment:
		newPks = append(newPks, &packet.MobArmourEquipment{
			EntityRuntimeID: pk.EntityRuntimeID,
			Helmet:          pk.Helmet,
			Chestplate:      pk.Chestplate,
			Leggings:        pk.Leggings,
			Boots:           pk.Boots,
			Body:            pk.Chestplate, // horray!!! hope this doesn't cause issues in the future...
		})
	case *legacypacket.ActorPickRequest:
		newPks = append(newPks, &packet.ActorPickRequest{
			EntityUniqueID: pk.EntityUniqueID,
			HotBarSlot:     pk.HotBarSlot,
		})
	case *packet.ClientCacheStatus:
		pk.Enabled = false
		newPks = append(newPks, pk)
	case *legacypacket.CommandRequest:
		newPks = append(newPks, &packet.CommandRequest{
			CommandLine:   pk.CommandLine,
			CommandOrigin: protocol.CommandOrigin(pk.CommandOrigin),
			Internal:      pk.Internal,
		})
	case *legacypacket.CommandOutput:
		newPks = append(newPks, &packet.CommandOutput{
			CommandOrigin:  protocol.CommandOrigin(pk.CommandOrigin),
			OutputType:     pk.OutputType,
			SuccessCount:   pk.SuccessCount,
			OutputMessages: pk.OutputMessages,
			DataSet:        protocol.Option(pk.DataSet),
		})
	case *legacypacket.ContainerClose:
		newPks = append(newPks, &packet.ContainerClose{
			WindowID:   pk.WindowID,
			ServerSide: pk.ServerSide,
		})
	case *legacypacket.Disconnect:
		newPks = append(newPks, &packet.Disconnect{
			HideDisconnectionScreen: pk.HideDisconnectionScreen,
			Message:                 pk.Message,
		})
	case *legacypacket.Emote:
		newPks = append(newPks, &packet.Emote{
			EntityRuntimeID: pk.EntityRuntimeID,
			EmoteID:         pk.EmoteID,
			XUID:            conn.IdentityData().XUID,
			PlatformID:      conn.ClientData().PlatformOnlineID,
			Flags:           pk.Flags,
		})
	case *legacypacket.InventoryTransaction:
		transactionData := pk.TransactionData
		if useItemData, ok := transactionData.(*types.UseItemTransactionData); ok {
			transactionData = &protocol.UseItemTransactionData{
				LegacyRequestID:    useItemData.LegacyRequestID,
				LegacySetItemSlots: useItemData.LegacySetItemSlots,
				Actions:            useItemData.Actions,
				ActionType:         useItemData.ActionType,
				BlockPosition:      useItemData.BlockPosition,
				BlockFace:          useItemData.BlockFace,
				HotBarSlot:         useItemData.HotBarSlot,
				HeldItem:           useItemData.HeldItem,
				Position:           useItemData.Position,
				ClickedPosition:    useItemData.ClickedPosition,
				ClientPrediction:   1,
				BlockRuntimeID:     useItemData.BlockRuntimeID,
			}

			if useItemData.ActionType == protocol.UseItemActionBreakBlock {
				newPks = append(newPks, &packet.PlayerAction{
					EntityRuntimeID: conn.GameData().EntityRuntimeID,
					ActionType:      protocol.PlayerActionPredictDestroyBlock,
					BlockPosition:   useItemData.BlockPosition,
					BlockFace:       useItemData.BlockFace,
				})
			}
		}
		newPks = append(newPks, &packet.InventoryTransaction{
			LegacyRequestID: pk.LegacyRequestID,
			LegacySetItemSlots: lo.Map(pk.LegacySetItemSlots, func(item protocol.LegacySetItemSlot, _ int) protocol.LegacySetItemSlot {
				if item.ContainerID >= 21 { // RECIPE_BOOK
					item.ContainerID += 1
				}
				return item
			}),
			Actions:         pk.Actions,
			TransactionData: transactionData,
		})
	case *legacypacket.ItemStackRequest:
		newPks = append(newPks, &packet.ItemStackRequest{
			Requests: lo.Map(pk.Requests, func(item types.ItemStackRequest, _ int) protocol.ItemStackRequest {
				return protocol.ItemStackRequest{
					RequestID: item.RequestID,
					Actions: lo.Map(item.Actions, func(item protocol.StackRequestAction, _ int) protocol.StackRequestAction {
						switch action := item.(type) {
						case *types.TakeStackRequestAction:
							return &action.TakeStackRequestAction
						case *types.PlaceStackRequestAction:
							return &action.PlaceStackRequestAction
						case *types.SwapStackRequestAction:
							return &action.SwapStackRequestAction
						case *types.DropStackRequestAction:
							return &action.DropStackRequestAction
						case *types.DestroyStackRequestAction:
							return &action.DestroyStackRequestAction
						case *types.ConsumeStackRequestAction:
							return &action.DestroyStackRequestAction
						case *types.PlaceInContainerStackRequestAction:
							return &action.PlaceInContainerStackRequestAction
						case *types.TakeOutContainerStackRequestAction:
							return &action.TakeOutContainerStackRequestAction
						case *types.AutoCraftRecipeStackRequestAction:
							return &action.AutoCraftRecipeStackRequestAction
						}
						return item
					}),
				}
			}),
		})
	case *legacypacket.SetActorLink:
		newEntityLink := protocol.EntityLink{
			RiddenEntityUniqueID:   pk.EntityLink.RiddenEntityUniqueID,
			RiderEntityUniqueID:    pk.EntityLink.RiderEntityUniqueID,
			Type:                   pk.EntityLink.Type,
			Immediate:              pk.EntityLink.Immediate,
			RiderInitiated:         pk.EntityLink.RiderInitiated,
			VehicleAngularVelocity: 0,
		}
		newPks = append(newPks, &packet.SetActorLink{
			EntityLink: newEntityLink,
		})
	case *legacypacket.ModalFormResponse:
		responseData := protocol.Optional[[]byte]{}
		cancelReason := protocol.Optional[uint8]{}
		if string(pk.ResponseData) == "null" {
			var cancelReasonType uint8 = packet.ModalFormCancelReasonUserClosed
			cancelReason = protocol.Option(cancelReasonType)
		} else {
			responseData = protocol.Option(pk.ResponseData)
		}
		newPks = append(newPks, &packet.ModalFormResponse{
			FormID:       pk.FormID,
			ResponseData: responseData,
			CancelReason: cancelReason,
		})
	case *legacypacket.PlayerAction:
		newPks = append(newPks, &packet.PlayerAction{
			EntityRuntimeID: pk.EntityRuntimeID,
			ActionType:      pk.ActionType,
			BlockPosition:   pk.BlockPosition,
			ResultPosition:  pk.BlockPosition,
			BlockFace:       pk.BlockFace,
		})
	case *legacypacket.PlayerAuthInput:
		newPks = append(newPks, &packet.PlayerAuthInput{
			Pitch:         pk.Pitch,
			Yaw:           pk.Yaw,
			Position:      pk.Position,
			MoveVector:    pk.MoveVector,
			HeadYaw:       pk.HeadYaw,
			InputData:     internal.Bitset(pk.InputData, packet.PlayerAuthInputBitsetSize),
			InputMode:     pk.InputMode,
			PlayMode:      pk.PlayMode,
			InteractPitch: pk.GazeDirection.X(),
			InteractYaw:   pk.GazeDirection.Y(),
			Tick:          pk.Tick,
			Delta:         pk.Delta,
			ItemInteractionData: func(data protocol.UseItemTransactionData) protocol.UseItemTransactionData {
				data.LegacySetItemSlots = lo.Map(data.LegacySetItemSlots, func(item protocol.LegacySetItemSlot, _ int) protocol.LegacySetItemSlot {
					if item.ContainerID >= 21 { // RECIPE_BOOK
						item.ContainerID += 1
					}
					return item
				})
				data.TriggerType = 0
				data.ClientPrediction = 1

				return data
			}(pk.ItemInteractionData),
			ItemStackRequest: protocol.ItemStackRequest{
				RequestID: pk.ItemStackRequest.RequestID,
				Actions: lo.Map(pk.ItemStackRequest.Actions, func(item protocol.StackRequestAction, _ int) protocol.StackRequestAction {
					switch action := item.(type) {
					case *types.TakeStackRequestAction:
						return &action.TakeStackRequestAction
					case *types.PlaceStackRequestAction:
						return &action.PlaceStackRequestAction
					case *types.SwapStackRequestAction:
						return &action.SwapStackRequestAction
					case *types.DropStackRequestAction:
						return &action.DropStackRequestAction
					case *types.DestroyStackRequestAction:
						return &action.DestroyStackRequestAction
					case *types.ConsumeStackRequestAction:
						return &action.DestroyStackRequestAction
					case *types.PlaceInContainerStackRequestAction:
						return &action.PlaceInContainerStackRequestAction
					case *types.TakeOutContainerStackRequestAction:
						return &action.TakeOutContainerStackRequestAction
					case *types.AutoCraftRecipeStackRequestAction:
						return &action.AutoCraftRecipeStackRequestAction
					}
					return item
				}),
				FilterStrings: pk.ItemStackRequest.FilterStrings,
			},
			BlockActions:       pk.BlockActions,
			AnalogueMoveVector: pk.MoveVector,
		})
	case *legacypacket.PlayerSkin:
		newPks = append(newPks, &packet.PlayerSkin{
			UUID:        pk.UUID,
			Skin:        pk.Skin.Skin,
			NewSkinName: pk.NewSkinName,
			OldSkinName: pk.OldSkinName,
		})
	case *legacypacket.RequestChunkRadius:
		newPks = append(newPks, &packet.RequestChunkRadius{
			ChunkRadius:    pk.ChunkRadius,
			MaxChunkRadius: uint8(pk.ChunkRadius),
		})
	case *legacypacket.ChangeDimension:
		newPks = append(newPks, &packet.ChangeDimension{
			Dimension: pk.Dimension,
			Position:  pk.Position,
			Respawn:   pk.Respawn,
		})
	case *legacypacket.Text:
		newPks = append(newPks, &packet.Text{
			TextType:         pk.TextType,
			NeedsTranslation: pk.NeedsTranslation,
			SourceName:       pk.SourceName,
			Message:          pk.Message,
			Parameters:       pk.Parameters,
			XUID:             pk.XUID,
			PlatformChatID:   pk.PlatformChatID,
		})
	case *packet.BlockActorData:
		pk.NBTData = upgradeBlockActorData(pk.NBTData)
		newPks = append(newPks, pk)
	case *legacypacket.Animate:
		newPks = append(newPks, &packet.Animate{
			ActionType:      uint8(pk.ActionType),
			EntityRuntimeID: pk.EntityRuntimeID,
			Data:            pk.BoatRowingTime,
		})
	case *packet.LevelEvent:
		if (pk.EventType & packet.LevelEventParticleLegacyEvent) != 0 {
			oldID := pk.EventType &^ packet.LevelEventParticleLegacyEvent
			pk.EventType = packet.LevelEventParticleLegacyEvent | types.UpgradeParticleId(oldID)
		}
	case *legacypacket.AdventureSettings:
		handleFlag := func(flags uint32, secondFlag bool) uint32 {
			layerMapping := map[uint32]uint32{
				packet.AdventureFlagAllowFlight:  protocol.AbilityMayFly,
				packet.AdventureFlagNoClip:       protocol.AbilityNoClip,
				packet.AdventureFlagWorldBuilder: protocol.AbilityWorldBuilder,
				packet.AdventureFlagFlying:       protocol.AbilityFlying,
				packet.AdventureFlagMuted:        protocol.AbilityMuted,
			}
			if secondFlag {
				layerMapping = map[uint32]uint32{
					packet.ActionPermissionMine:             protocol.AbilityMine,
					packet.ActionPermissionDoorsAndSwitches: protocol.AbilityDoorsAndSwitches,
					packet.ActionPermissionOpenContainers:   protocol.AbilityOpenContainers,
					packet.ActionPermissionAttackPlayers:    protocol.AbilityAttackPlayers,
					packet.ActionPermissionAttackMobs:       protocol.AbilityAttackMobs,
					packet.ActionPermissionOperator:         protocol.AbilityOperatorCommands,
					packet.ActionPermissionBuild:            protocol.AbilityBuild,
				}
			}

			out := uint32(0)
			for flag, mapped := range layerMapping {
				if (flags & flag) != 0 {
					out |= mapped
				}
			}
			return out
		}

		_ = handleFlag
	case *legacypacket.TickSync: // no longer exist on 1.21
	case *legacypacket.ItemFrameDropItem:
	case *packet.AdventureSettings:
	default:
		newPks = append(newPks, pk)
	}

	/*
		if pk.ID() != packet.IDNetworkStackLatency && pk.ID() != packet.IDPlayerAuthInput {
			fmt.Printf("C->S %d %T\n", pk.ID(), pk)
		}
	*/

	return p.blockTranslator.UpgradeBlockPackets(p.itemTranslator.UpgradeItemPackets(newPks, conn), conn)
}

func (p Protocol) ConvertFromLatest(pk packet.Packet, conn *minecraft.Conn) (result []packet.Packet) {
	/*
		if pk.ID() != packet.IDNetworkStackLatency && pk.ID() != packet.IDPlayerAuthInput && pk.ID() != packet.IDSetScore && pk.ID() != packet.IDSetActorData {
			fmt.Printf("Latest -> 1.16.100: %T\n", pk)
		}
	*/

	result = p.blockTranslator.DowngradeBlockPackets(p.itemTranslator.DowngradeItemPackets([]packet.Packet{pk}, conn), conn)

	for i, pk := range result {
		switch pk := pk.(type) {
		case *packet.Interact:
			var position mgl32.Vec3
			if v, ok := pk.Position.Value(); ok {
				position = v
			}

			result[i] = &legacypacket.Interact{
				ActionType:            pk.ActionType,
				TargetEntityRuntimeID: pk.TargetEntityRuntimeID,
				Position:              position,
			}
		case *packet.CommandOutput:
			var dataset string
			if v, ok := pk.DataSet.Value(); ok {
				dataset = v
			}

			result[i] = &legacypacket.CommandOutput{
				CommandOrigin:  types.CommandOrigin(pk.CommandOrigin),
				OutputType:     pk.OutputType,
				SuccessCount:   pk.SuccessCount,
				OutputMessages: pk.OutputMessages,
				DataSet:        dataset,
			}
		case *packet.Transfer:
			result[i] = &legacypacket.Transfer{
				Address: pk.Address,
				Port:    pk.Port,
			}
		case *packet.LevelSoundEvent:
			result[i] = &legacypacket.LevelSoundEvent{
				SoundType:             pk.SoundType,
				Position:              pk.Position,
				ExtraData:             pk.ExtraData,
				EntityType:            pk.EntityType,
				BabyMob:               pk.BabyMob,
				DisableRelativeVolume: pk.DisableRelativeVolume,
			}
		case *packet.BossEvent:
			result[i] = &legacypacket.BossEvent{
				BossEntityUniqueID: pk.BossEntityUniqueID,
				EventType:          pk.EventType,
				PlayerUniqueID:     pk.PlayerUniqueID,
				BossBarTitle:       pk.BossBarTitle,
				HealthPercentage:   pk.HealthPercentage,
				ScreenDarkening:    pk.ScreenDarkening,
				Colour:             pk.Colour,
				Overlay:            pk.Overlay,
			}
		case *packet.ContainerOpen:
			result[i] = &legacypacket.ContainerOpen{
				WindowID:                pk.WindowID,
				ContainerType:           pk.ContainerType,
				ContainerPosition:       pk.ContainerPosition,
				ContainerEntityUniqueID: pk.ContainerEntityUniqueID,
			}
		case *packet.UpdateBlock:
			result[i] = &legacypacket.UpdateBlock{
				Position:          pk.Position,
				NewBlockRuntimeID: pk.NewBlockRuntimeID,
				Flags:             pk.Flags,
				Layer:             pk.Layer,
			}
		case *packet.BlockActorData:
			result[i] = &legacypacket.BlockActorData{
				Position: pk.Position,
				NBTData:  pk.NBTData,
			}
		case *packet.AnvilDamage:
			result[i] = &legacypacket.AnvilDamage{
				Damage:        pk.Damage,
				AnvilPosition: pk.AnvilPosition,
			}
		case *packet.BlockEvent:
			result[i] = &legacypacket.BlockEvent{
				Position:  pk.Position,
				EventType: pk.EventType,
				EventData: pk.EventData,
			}
		case *packet.UpdateBlockSynced:
			result[i] = &legacypacket.UpdateBlockSynced{
				Position:          pk.Position,
				NewBlockRuntimeID: pk.NewBlockRuntimeID,
				Flags:             pk.Flags,
				Layer:             pk.Layer,
				EntityUniqueID:    pk.EntityUniqueID,
				TransitionType:    pk.TransitionType,
			}
		case *packet.PlayerArmourDamage:
			var bitset uint8
			var helmetDamage, chestplateDamage, leggingsDamage, bootsDamage int16

			for _, entry := range pk.List {
				switch entry.ArmourSlot {
				case 0:
					bitset |= 1 << packet.PlayerArmourDamageFlagHelmet
					helmetDamage = entry.Damage
				case 1:
					bitset |= 1 << packet.PlayerArmourDamageFlagChestplate
					chestplateDamage = entry.Damage
				case 2:
					bitset |= 1 << packet.PlayerArmourDamageFlagLeggings
					leggingsDamage = entry.Damage
				case 3:
					bitset |= 1 << packet.PlayerArmourDamageFlagBoots
					bootsDamage = entry.Damage
				}
			}
			result[i] = &legacypacket.PlayerArmourDamage{
				Bitset:           bitset,
				HelmetDamage:     int32(helmetDamage),
				ChestplateDamage: int32(chestplateDamage),
				LeggingsDamage:   int32(leggingsDamage),
				BootsDamage:      int32(bootsDamage),
			}
		case *packet.StopSound:
			result[i] = &legacypacket.StopSound{
				SoundName: pk.SoundName,
				StopAll:   pk.StopAll,
			}
		case *packet.ChangeDimension:
			result[i] = &legacypacket.ChangeDimension{
				Dimension: pk.Dimension,
				Position:  pk.Position,
				Respawn:   pk.Respawn,
			}
		case *packet.MobArmourEquipment:
			result[i] = &legacypacket.MobArmourEquipment{
				EntityRuntimeID: pk.EntityRuntimeID,
				Helmet:          pk.Helmet,
				Chestplate:      pk.Chestplate,
				Leggings:        pk.Leggings,
				Boots:           pk.Boots,
			}
		case *packet.AddActor:
			links := make([]types.EntityLink, len(pk.EntityLinks))
			for i, link := range pk.EntityLinks {
				links[i] = types.EntityLink{
					RiddenEntityUniqueID: link.RiddenEntityUniqueID,
					RiderEntityUniqueID:  link.RiderEntityUniqueID,
					Type:                 link.Type,
					Immediate:            link.Immediate,
					RiderInitiated:       link.RiderInitiated,
				}
			}

			result[i] = &legacypacket.AddActor{
				EntityMetadata:  downgradeEntityMetadata(pk.EntityMetadata),
				EntityRuntimeID: pk.EntityRuntimeID,
				EntityType:      pk.EntityType,
				EntityUniqueID:  pk.EntityUniqueID,
				HeadYaw:         pk.HeadYaw,
				Pitch:           pk.Pitch,
				Position:        pk.Position,
				Velocity:        pk.Velocity,
				Yaw:             pk.Yaw,
				Attributes:      pk.Attributes,
				EntityLinks:     links,
			}
		case *packet.AddItemActor:
			result[i] = &packet.AddItemActor{
				EntityMetadata:  downgradeEntityMetadata(pk.EntityMetadata),
				EntityRuntimeID: pk.EntityRuntimeID,
				EntityUniqueID:  pk.EntityUniqueID,
				FromFishing:     pk.FromFishing,
				Item:            pk.Item,
				Position:        pk.Position,
				Velocity:        pk.Velocity,
			}
		case *packet.AddPlayer:
			result[i] = &legacypacket.AddPlayer{
				UUID:            pk.UUID,
				Username:        pk.Username,
				EntityUniqueID:  pk.AbilityData.EntityUniqueID,
				EntityRuntimeID: pk.EntityRuntimeID,
				PlatformChatID:  pk.PlatformChatID,
				Position:        pk.Position,
				Velocity:        pk.Velocity,
				Pitch:           pk.Pitch,
				Yaw:             pk.Yaw,
				HeadYaw:         pk.HeadYaw,
				HeldItem:        pk.HeldItem,
				EntityMetadata:  downgradeEntityMetadata(pk.EntityMetadata),
				AdventureSettings: packet.AdventureSettings{
					CommandPermissionLevel: uint32(pk.AbilityData.CommandPermissions),
					PermissionLevel:        uint32(pk.AbilityData.PlayerPermissions),
					PlayerUniqueID:         pk.AbilityData.EntityUniqueID,
				},
				DeviceID:    pk.DeviceID,
				EntityLinks: pk.EntityLinks,
			}
		case *packet.AddVolumeEntity:
			result[i] = &legacypacket.AddVolumeEntity{
				EntityRuntimeID:    uint64(pk.EntityRuntimeID),
				EntityMetadata:     pk.EntityMetadata,
				EncodingIdentifier: pk.EncodingIdentifier,
				InstanceIdentifier: pk.InstanceIdentifier,
				EngineVersion:      pk.EngineVersion,
			}
		case *packet.AnimateEntity:
			result[i] = &legacypacket.AnimateEntity{
				Animation:        pk.Animation,
				NextState:        pk.NextState,
				StopCondition:    pk.StopCondition,
				Controller:       pk.Controller,
				BlendOutTime:     pk.BlendOutTime,
				EntityRuntimeIDs: pk.EntityRuntimeIDs,
			}
		case *packet.AvailableActorIdentifiers:
			result[i] = &packet.AvailableActorIdentifiers{
				SerialisedEntityIdentifiers: entityIdentifierData,
			}
		case *packet.AvailableCommands:
			for ind1, command := range pk.Commands {
				for ind2, overload := range command.Overloads {
					for ind3, parameter := range overload.Parameters {
						parameterType := uint32(parameter.Type) | protocol.CommandArgValid
						switch parameter.Type | protocol.CommandArgValid {
						case protocol.CommandArgTypeCompareOperator:
							parameterType = protocol.CommandArgTypeOperator
						case protocol.CommandArgTypeTarget:
							parameterType = 7
						case protocol.CommandArgTypeWildcardTarget:
							parameterType = 8
						case protocol.CommandArgTypeFilepath:
							parameterType = 16
						case protocol.CommandArgTypeString:
							parameterType = 32
						case protocol.CommandArgTypeBlockPosition:
							fallthrough
						case protocol.CommandArgTypePosition:
							parameterType = 40
						case protocol.CommandArgTypeMessage:
							parameterType = 44
						case protocol.CommandArgTypeRawText:
							parameterType = 46
						case protocol.CommandArgTypeJSON:
							parameterType = 50
						case protocol.CommandArgTypeCommand:
							parameterType = 63
						}
						parameter.Type = parameterType | protocol.CommandArgValid
						pk.Commands[ind1].Overloads[ind2].Parameters[ind3] = parameter
					}
				}
			}
			result[i] = &legacypacket.AvailableCommands{
				Commands: lo.Map(pk.Commands, func(c protocol.Command, _ int) types.Command {
					return types.Command{
						Name:            c.Name,
						Description:     c.Description,
						Flags:           byte(c.Flags),
						PermissionLevel: c.PermissionLevel,
						Overloads: lo.Map(c.Overloads, func(o protocol.CommandOverload, _ int) types.CommandOverload {
							return types.CommandOverload{Parameters: lo.Map(o.Parameters, func(p protocol.CommandParameter, _ int) types.CommandParameter {
								return types.CommandParameter{
									Name:                p.Name,
									Type:                types.DowngradeParamType(p.Type),
									Optional:            p.Optional,
									CollapseEnumOptions: true,
								}
							})}
						}),
					}
				}),
			}
		case *packet.CameraShake:
			result[i] = &legacypacket.CameraShake{
				Intensity: pk.Intensity,
				Duration:  pk.Duration,
				Type:      pk.Type,
			}
		case *packet.CraftingData:
			result[i] = &legacypacket.CraftingData{
				ClearRecipes: true,
			}
		case *packet.CreativeContent:
			result[i] = &legacypacket.CreativeContent{}
		case *packet.ContainerClose:
			result[i] = &legacypacket.ContainerClose{
				WindowID:   pk.WindowID,
				ServerSide: pk.ServerSide,
			}
		case *packet.Disconnect:
			result[i] = &legacypacket.Disconnect{
				HideDisconnectionScreen: pk.HideDisconnectionScreen,
				Message:                 pk.Message,
			}
		case *packet.Emote:
			result[i] = &legacypacket.Emote{
				EntityRuntimeID: pk.EntityRuntimeID,
				EmoteID:         pk.EmoteID,
				Flags:           pk.Flags,
			}
		case *packet.HurtArmour:
			result[i] = &legacypacket.HurtArmour{
				Cause:  pk.Cause,
				Damage: pk.Damage,
			}
		case *packet.InventoryContent:
			result[i] = &legacypacket.InventoryContent{
				WindowID: pk.WindowID,
				Content:  pk.Content,
			}
		case *packet.InventorySlot:
			result[i] = &legacypacket.InventorySlot{
				WindowID: pk.WindowID,
				Slot:     pk.Slot,
				NewItem:  pk.NewItem,
			}
		case *packet.InventoryTransaction:
			result[i] = &legacypacket.InventoryTransaction{
				LegacyRequestID: pk.LegacyRequestID,
				LegacySetItemSlots: lo.Map(pk.LegacySetItemSlots, func(item protocol.LegacySetItemSlot, _ int) protocol.LegacySetItemSlot {
					if item.ContainerID > 21 { // RECIPE_BOOK
						item.ContainerID -= 1
					}
					return item
				}),
				Actions:         pk.Actions,
				TransactionData: pk.TransactionData,
			}
		case *packet.ItemStackResponse:
			result[i] = &legacypacket.ItemStackResponse{
				Responses: lo.Map(pk.Responses, func(response protocol.ItemStackResponse, _ int) types.ItemStackResponse {
					return types.ItemStackResponse{
						Status:    response.Status,
						RequestID: response.RequestID,
						ContainerInfo: lo.Map(response.ContainerInfo, func(info protocol.StackResponseContainerInfo, _ int) types.StackResponseContainerInfo {
							if info.Container.ContainerID > 21 { // RECIPE_BOOK
								info.Container.ContainerID -= 1
							}
							return types.StackResponseContainerInfo{
								ContainerID: info.Container.ContainerID,
								SlotInfo: lo.Map(info.SlotInfo, func(slot protocol.StackResponseSlotInfo, _ int) types.StackResponseSlotInfo {
									return types.StackResponseSlotInfo{
										StackResponseSlotInfo: slot,
									}
								}),
							}
						}),
					}
				}),
			}
		case *packet.LevelChunk:
			result[i] = &legacypacket.LevelChunk{
				Position:      pk.Position,
				SubChunkCount: pk.SubChunkCount,
				CacheEnabled:  pk.CacheEnabled,
				BlobHashes:    pk.BlobHashes,
				RawPayload:    pk.RawPayload,
			}
		case *packet.MobEffect:
			result[i] = &legacypacket.MobEffect{
				EntityRuntimeID: pk.EntityRuntimeID,
				Operation:       pk.Operation,
				EffectType:      pk.EffectType,
				Amplifier:       pk.Amplifier,
				Particles:       pk.Particles,
				Duration:        pk.Duration,
			}
		case *packet.NetworkChunkPublisherUpdate:
			result[i] = &legacypacket.NetworkChunkPublisherUpdate{
				Position: pk.Position,
				Radius:   pk.Radius,
			}
		case *packet.NetworkSettings:
			result[i] = &legacypacket.NetworkSettings{
				CompressionThreshold: pk.CompressionThreshold,
			}
		case *packet.PlayerList:
			result[i] = &legacypacket.PlayerList{
				ActionType: pk.ActionType,
				Entries: lo.Map(pk.Entries, func(item protocol.PlayerListEntry, _ int) types.PlayerListEntry {
					return types.PlayerListEntry{PlayerListEntry: item}
				}),
			}
		case *packet.PlayerSkin:
			result[i] = &legacypacket.PlayerSkin{
				UUID:        pk.UUID,
				Skin:        types.Skin{Skin: pk.Skin},
				NewSkinName: pk.NewSkinName,
				OldSkinName: pk.OldSkinName,
			}
		case *packet.RemoveVolumeEntity:
			result[i] = &legacypacket.RemoveVolumeEntity{
				EntityRuntimeID: uint64(pk.EntityRuntimeID),
			}
		case *packet.ResourcePacksInfo:
			result[i] = &legacypacket.ResourcePacksInfo{
				TexturePackRequired: pk.TexturePackRequired,
				HasScripts:          pk.HasScripts,
				BehaviourPacks:      []types.BehaviourPackInfo{},
				TexturePacks: lo.Map(pk.TexturePacks, func(pack protocol.TexturePackInfo, _ int) types.TexturePackInfo {
					return types.TexturePackInfo{
						UUID:            pack.UUID.String(),
						Version:         pack.Version,
						Size:            pack.Size,
						ContentKey:      pack.ContentKey,
						SubPackName:     pack.SubPackName,
						ContentIdentity: pack.ContentIdentity,
						HasScripts:      pack.HasScripts,
					}
				}),
			}
		case *packet.ResourcePackStack:
			result[i] = &legacypacket.ResourcePackStack{
				TexturePackRequired:          pk.TexturePackRequired,
				TexturePacks:                 pk.TexturePacks,
				BaseGameVersion:              pk.BaseGameVersion,
				Experiments:                  pk.Experiments,
				ExperimentsPreviouslyToggled: pk.ExperimentsPreviouslyToggled,
			}
		case *packet.SetActorData:
			result[i] = &legacypacket.SetActorData{
				EntityRuntimeID: pk.EntityRuntimeID,
				EntityMetadata:  downgradeEntityMetadata(pk.EntityMetadata),
				Tick:            pk.Tick,
			}
		case *packet.SetActorMotion:
			result[i] = &legacypacket.SetActorMotion{
				EntityRuntimeID: pk.EntityRuntimeID,
				Velocity:        pk.Velocity,
			}
		case *packet.SetTitle:
			result[i] = &legacypacket.SetTitle{
				ActionType:      pk.ActionType,
				Text:            pk.Text,
				FadeInDuration:  pk.FadeInDuration,
				RemainDuration:  pk.RemainDuration,
				FadeOutDuration: pk.FadeOutDuration,
			}
		case *packet.SpawnParticleEffect:
			result[i] = &legacypacket.SpawnParticleEffect{
				Dimension:      pk.Dimension,
				EntityUniqueID: pk.EntityUniqueID,
				Position:       pk.Position,
				ParticleName:   pk.ParticleName,
			}
		case *packet.Animate:
			result[i] = &legacypacket.Animate{
				EntityRuntimeID: pk.EntityRuntimeID,
				ActionType:      int32(pk.ActionType),
				BoatRowingTime:  pk.Data,
			}
		case *packet.StartGame:
			enabled := pk.ForceExperimentalGameplay
			result[i] = &legacypacket.StartGame{
				EntityUniqueID:                  pk.EntityUniqueID,
				EntityRuntimeID:                 pk.EntityRuntimeID,
				PlayerGameMode:                  pk.PlayerGameMode,
				PlayerPosition:                  pk.PlayerPosition,
				Pitch:                           pk.Pitch,
				Yaw:                             pk.Yaw,
				WorldSeed:                       int32(pk.WorldSeed),
				SpawnBiomeType:                  pk.SpawnBiomeType,
				UserDefinedBiomeName:            pk.UserDefinedBiomeName,
				Dimension:                       pk.Dimension,
				Generator:                       pk.Generator,
				WorldGameMode:                   pk.WorldGameMode,
				Difficulty:                      pk.Difficulty,
				WorldSpawn:                      pk.WorldSpawn,
				AchievementsDisabled:            pk.AchievementsDisabled,
				DayCycleLockTime:                pk.DayCycleLockTime,
				EducationEditionOffer:           pk.EducationEditionOffer,
				EducationFeaturesEnabled:        pk.EducationFeaturesEnabled,
				EducationProductID:              pk.EducationProductID,
				RainLevel:                       pk.RainLevel,
				LightningLevel:                  pk.LightningLevel,
				ConfirmedPlatformLockedContent:  pk.ConfirmedPlatformLockedContent,
				MultiPlayerGame:                 pk.MultiPlayerGame,
				LANBroadcastEnabled:             pk.LANBroadcastEnabled,
				XBLBroadcastMode:                pk.XBLBroadcastMode,
				PlatformBroadcastMode:           pk.PlatformBroadcastMode,
				CommandsEnabled:                 pk.CommandsEnabled,
				TexturePackRequired:             pk.TexturePackRequired,
				GameRules:                       pk.GameRules,
				Experiments:                     pk.Experiments,
				ExperimentsPreviouslyToggled:    pk.ExperimentsPreviouslyToggled,
				BonusChestEnabled:               pk.BonusChestEnabled,
				StartWithMapEnabled:             pk.StartWithMapEnabled,
				PlayerPermissions:               pk.PlayerPermissions,
				ServerChunkTickRadius:           pk.ServerChunkTickRadius,
				HasLockedBehaviourPack:          pk.HasLockedBehaviourPack,
				HasLockedTexturePack:            pk.HasLockedTexturePack,
				FromLockedWorldTemplate:         pk.FromLockedWorldTemplate,
				MSAGamerTagsOnly:                pk.MSAGamerTagsOnly,
				FromWorldTemplate:               pk.FromWorldTemplate,
				WorldTemplateSettingsLocked:     pk.WorldTemplateSettingsLocked,
				OnlySpawnV1Villagers:            pk.OnlySpawnV1Villagers,
				BaseGameVersion:                 pk.BaseGameVersion,
				LimitedWorldWidth:               pk.LimitedWorldWidth,
				LimitedWorldDepth:               pk.LimitedWorldDepth,
				NewNether:                       pk.NewNether,
				ForceExperimentalGameplay:       enabled,
				LevelID:                         pk.LevelID,
				WorldName:                       pk.WorldName,
				TemplateContentIdentity:         pk.TemplateContentIdentity,
				Trial:                           pk.Trial,
				ServerAuthoritativeMovementMode: uint32(2),
				Time:                            pk.Time,
				EnchantmentSeed:                 pk.EnchantmentSeed,
				Blocks:                          pk.Blocks,
				Items:                           pk.Items,
				MultiPlayerCorrelationID:        pk.MultiPlayerCorrelationID,
				ServerAuthoritativeInventory:    pk.ServerAuthoritativeInventory,
			}
		case *packet.Text:
			result[i] = &legacypacket.Text{
				TextType:         pk.TextType,
				NeedsTranslation: pk.NeedsTranslation,
				SourceName:       pk.SourceName,
				Message:          pk.Message,
				Parameters:       pk.Parameters,
				XUID:             pk.XUID,
				PlatformChatID:   pk.PlatformChatID,
			}
		case *packet.UpdateAttributes:
			result[i] = &legacypacket.UpdateAttributes{
				EntityRuntimeID: pk.EntityRuntimeID,
				Attributes: lo.Map(pk.Attributes, func(item protocol.Attribute, _ int) types.Attribute {
					return types.Attribute{Attribute: item}
				}),
				Tick: pk.Tick,
			}
		case *packet.BiomeDefinitionList:
			result[i] = &legacypacket.BiomeDefinitionList{
				SerialisedBiomeDefinitions: internal.BiomeDecompressed,
			}
		case *packet.UpdateAbilities:
			if len(pk.AbilityData.Layers) == 0 || pk.AbilityData.EntityUniqueID != conn.GameData().EntityUniqueID {
				// We need at least one layer.
				return nil
			}

			base, flags, perms := pk.AbilityData.Layers[0].Values, uint32(0), uint32(0)
			if base&protocol.AbilityMayFly != 0 {
				flags |= legacypacket.AdventureFlagAllowFlight
				if base&protocol.AbilityFlying != 0 {
					flags |= legacypacket.AdventureFlagFlying
				}
			}
			if base&protocol.AbilityNoClip != 0 {
				flags |= legacypacket.AdventureFlagNoClip
			}

			if base&protocol.AbilityBuild != 0 && base&protocol.AbilityMine != 0 {
				flags |= legacypacket.AdventureFlagWorldBuilder
			} else {
				flags |= legacypacket.AdventureFlagWorldImmutable
			}
			if base&protocol.AbilityBuild != 0 {
				perms |= legacypacket.ActionPermissionBuildAndMine
			}
			if base&protocol.AbilityMine != 0 {
				perms |= legacypacket.ActionPermissionBuildAndMine
			}

			if base&protocol.AbilityDoorsAndSwitches != 0 {
				perms |= legacypacket.ActionPermissionDoorsAndSwitched
			}
			if base&protocol.AbilityOpenContainers != 0 {
				perms |= legacypacket.ActionPermissionOpenContainers
			}
			if base&protocol.AbilityAttackPlayers != 0 {
				perms |= legacypacket.ActionPermissionAttackPlayers
			}
			if base&protocol.AbilityAttackMobs != 0 {
				perms |= legacypacket.ActionPermissionAttackMobs
			}

			result[i] = &legacypacket.AdventureSettings{
				Flags:                  flags,
				ActionPermissions:      perms,
				PlayerUniqueID:         pk.AbilityData.EntityUniqueID,
				CommandPermissionLevel: uint32(pk.AbilityData.CommandPermissions),
				PermissionLevel:        uint32(pk.AbilityData.PlayerPermissions),
			}
		case *packet.UpdatePlayerGameType:
			result[i] = &legacypacket.UpdatePlayerGameType{
				GameType:       pk.GameType,
				PlayerUniqueID: pk.PlayerUniqueID,
			}
		case *packet.PlayerAction:
			result[i] = &legacypacket.PlayerAction{
				EntityRuntimeID: pk.EntityRuntimeID,
				ActionType:      pk.ActionType,
				BlockPosition:   pk.BlockPosition,
				BlockFace:       pk.BlockFace,
			}
		case *packet.SetActorLink:
			result[i] = &legacypacket.SetActorLink{
				EntityLink: types.EntityLink{
					RiddenEntityUniqueID: pk.EntityLink.RiddenEntityUniqueID,
					RiderEntityUniqueID:  pk.EntityLink.RiderEntityUniqueID,
					Type:                 pk.EntityLink.Type,
					Immediate:            pk.EntityLink.Immediate,
					RiderInitiated:       pk.EntityLink.RiderInitiated,
				},
			}
		case *packet.ItemRegistry:
			result[i] = &legacypacket.ItemComponent{}
		case *packet.CorrectPlayerMovePrediction:
			result[i] = &legacypacket.CorrectPlayerMovePrediction{
				Position: pk.Position,
				Delta:    pk.Delta,
				OnGround: pk.OnGround,
				Tick:     pk.Tick,
			}
		case *packet.LevelEvent:
			if (pk.EventType & packet.LevelEventParticleLegacyEvent) != 0 {
				newID := pk.EventType &^ packet.LevelEventParticleLegacyEvent
				pk.EventType = packet.LevelEventParticleLegacyEvent | types.DowngradeParticleId(newID)
			}
		}
	}

	//fmt.Printf("S->C %d %T\n", pk.ID(), pk)
	return result
}
