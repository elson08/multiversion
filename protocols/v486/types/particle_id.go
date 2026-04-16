package types

const (
	BUBBLE                  = 1
	BUBBLE_MANUAL           = 2
	CRITICAL                = 3
	BLOCK_FORCE_FIELD       = 4
	SMOKE                   = 5
	EXPLODE                 = 6
	EVAPORATION             = 7
	FLAME                   = 8
	CANDLE_FLAME            = 9
	LAVA                    = 10
	LARGE_SMOKE             = 11
	REDSTONE                = 12
	RISING_RED_DUST         = 13
	ITEM_BREAK              = 14
	SNOWBALL_POOF           = 15
	HUGE_EXPLODE            = 16
	HUGE_EXPLODE_SEED       = 17
	MOB_FLAME               = 18
	HEART                   = 19
	TERRAIN                 = 20
	SUSPENDED_TOWN          = 21
	PORTAL                  = 22
	SPLASH                  = 24
	WATER_SPLASH            = 24
	WATER_SPLASH_MANUAL     = 25
	WATER_WAKE              = 26
	DRIP_WATER              = 27
	DRIP_LAVA               = 28
	DRIP_HONEY              = 29
	STALACTITE_DRIP_WATER   = 30
	STALACTITE_DRIP_LAVA    = 31
	FALLING_DUST            = 32
	MOB_SPELL               = 33
	MOB_SPELL_AMBIENT       = 34
	MOB_SPELL_INSTANTANEOUS = 35
	INK                     = 36
	SLIME                   = 37
	RAIN_SPLASH             = 38
	VILLAGER_ANGRY          = 39
	VILLAGER_HAPPY          = 40
	ENCHANTMENT_TABLE       = 41
	TRACKING_EMITTER        = 42
	NOTE                    = 43
	WITCH_SPELL             = 44
	CARROT                  = 45
	MOB_APPEARANCE          = 46
	END_ROD                 = 47
	DRAGONS_BREATH          = 48
	SPIT                    = 49
	TOTEM                   = 50
	FOOD                    = 51
	FIREWORKS_STARTER       = 52
	FIREWORKS_SPARK         = 53
	FIREWORKS_OVERLAY       = 54
	BALLOON_GAS             = 55
	COLORED_FLAME           = 56
	SPARKLER                = 57
	CONDUIT                 = 58
	BUBBLE_COLUMN_UP        = 59
	BUBBLE_COLUMN_DOWN      = 60
	SNEEZE                  = 61
	SHULKER_BULLET          = 62
	BLEACH                  = 63
	DRAGON_DESTROY_BLOCK    = 64
	MYCELIUM_DUST           = 65
	FALLING_RED_DUST        = 66
	CAMPFIRE_SMOKE          = 67
	TALL_CAMPFIRE_SMOKE     = 68
	DRAGON_BREATH_FIRE      = 69
	DRAGON_BREATH_TRAIL     = 70
	BLUE_FLAME              = 71
	SOUL                    = 72
	OBSIDIAN_TEAR           = 73
	PORTAL_REVERSE          = 74
	SNOWFLAKE               = 75
	VIBRATION_SIGNAL        = 76
	SCULK_SENSOR_REDSTONE   = 77
	SPORE_BLOSSOM_SHOWER    = 78
	SPORE_BLOSSOM_AMBIENT   = 79
	WAX                     = 80
	ELECTRIC_SPARK          = 81
)

var newToOld = map[int32]int32{
	19: MOB_FLAME,
	20: HEART,
	21: TERRAIN,
	22: SUSPENDED_TOWN,
	23: PORTAL,
	25: SPLASH,
	26: WATER_SPLASH_MANUAL,
	27: WATER_WAKE,
	28: DRIP_WATER,
	29: DRIP_LAVA,
	30: DRIP_HONEY,
	31: STALACTITE_DRIP_WATER,
	32: STALACTITE_DRIP_LAVA,
	33: FALLING_DUST,
	34: MOB_SPELL,
	35: MOB_SPELL_AMBIENT,
	36: MOB_SPELL_INSTANTANEOUS,
	37: INK,
	38: SLIME,
	39: RAIN_SPLASH,
	40: VILLAGER_ANGRY,
	41: VILLAGER_HAPPY,
	42: ENCHANTMENT_TABLE,
	43: TRACKING_EMITTER,
	44: NOTE,
	45: WITCH_SPELL,
	46: CARROT,
	47: MOB_APPEARANCE,
	48: END_ROD,
	49: DRAGONS_BREATH,
	50: SPIT,
	51: TOTEM,
	52: FOOD,
	53: FIREWORKS_STARTER,
	54: FIREWORKS_SPARK,
	55: FIREWORKS_OVERLAY,
	56: BALLOON_GAS,
	57: COLORED_FLAME,
	58: SPARKLER,
	59: CONDUIT,
	60: BUBBLE_COLUMN_UP,
	61: BUBBLE_COLUMN_DOWN,
	62: SNEEZE,
	63: SHULKER_BULLET,
	64: BLEACH,
	65: DRAGON_DESTROY_BLOCK,
	66: MYCELIUM_DUST,
	67: FALLING_RED_DUST,
	68: CAMPFIRE_SMOKE,
	69: TALL_CAMPFIRE_SMOKE,
	70: DRAGON_BREATH_FIRE,
	71: DRAGON_BREATH_TRAIL,
	72: BLUE_FLAME,
	73: SOUL,
	74: OBSIDIAN_TEAR,
	75: PORTAL_REVERSE,
	76: SNOWFLAKE,
	77: VIBRATION_SIGNAL,
	78: SCULK_SENSOR_REDSTONE,
	79: SPORE_BLOSSOM_SHOWER,
	80: SPORE_BLOSSOM_AMBIENT,
	81: WAX,
	82: ELECTRIC_SPARK,
}

var oldToNew = map[int32]int32{
	MOB_FLAME:               19,
	HEART:                   20,
	TERRAIN:                 21,
	SUSPENDED_TOWN:          22,
	PORTAL:                  23,
	SPLASH:                  25,
	WATER_SPLASH_MANUAL:     26,
	WATER_WAKE:              27,
	DRIP_WATER:              28,
	DRIP_LAVA:               29,
	DRIP_HONEY:              30,
	STALACTITE_DRIP_WATER:   31,
	STALACTITE_DRIP_LAVA:    32,
	FALLING_DUST:            33,
	MOB_SPELL:               34,
	MOB_SPELL_AMBIENT:       35,
	MOB_SPELL_INSTANTANEOUS: 36,
	INK:                     37,
	SLIME:                   38,
	RAIN_SPLASH:             39,
	VILLAGER_ANGRY:          40,
	VILLAGER_HAPPY:          41,
	ENCHANTMENT_TABLE:       42,
	TRACKING_EMITTER:        43,
	NOTE:                    44,
	WITCH_SPELL:             45,
	CARROT:                  46,
	MOB_APPEARANCE:          47,
	END_ROD:                 48,
	DRAGONS_BREATH:          49,
	SPIT:                    50,
	TOTEM:                   51,
	FOOD:                    52,
	FIREWORKS_STARTER:       53,
	FIREWORKS_SPARK:         54,
	FIREWORKS_OVERLAY:       55,
	BALLOON_GAS:             56,
	COLORED_FLAME:           57,
	SPARKLER:                58,
	CONDUIT:                 59,
	BUBBLE_COLUMN_UP:        60,
	BUBBLE_COLUMN_DOWN:      61,
	SNEEZE:                  62,
	SHULKER_BULLET:          63,
	BLEACH:                  64,
	DRAGON_DESTROY_BLOCK:    65,
	MYCELIUM_DUST:           66,
	FALLING_RED_DUST:        67,
	CAMPFIRE_SMOKE:          68,
	TALL_CAMPFIRE_SMOKE:     69,
	DRAGON_BREATH_FIRE:      70,
	DRAGON_BREATH_TRAIL:     71,
	BLUE_FLAME:              72,
	SOUL:                    73,
	OBSIDIAN_TEAR:           74,
	PORTAL_REVERSE:          75,
	SNOWFLAKE:               76,
	VIBRATION_SIGNAL:        77,
	SCULK_SENSOR_REDSTONE:   78,
	SPORE_BLOSSOM_SHOWER:    79,
	SPORE_BLOSSOM_AMBIENT:   80,
	WAX:                     81,
	ELECTRIC_SPARK:          82,
}

func UpgradeParticleId(id int32) int32 {
	if newID, ok := oldToNew[id]; ok {
		return newID
	}

	return id
}

func DowngradeParticleId(id int32) int32 {
	if oldID, ok := newToOld[id]; ok {
		return oldID
	}

	return id
}
