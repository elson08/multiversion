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
	LAVA                    = 9
	LARGE_SMOKE             = 10
	REDSTONE                = 11
	RISING_RED_DUST         = 12
	ITEM_BREAK              = 13
	SNOWBALL_POOF           = 14
	HUGE_EXPLODE            = 15
	HUGE_EXPLODE_SEED       = 16
	MOB_FLAME               = 17
	HEART                   = 18
	TERRAIN                 = 19
	SUSPENDED_TOWN          = 20
	TOWN_AURA               = 20
	PORTAL                  = 21
	SPLASH                  = 23
	WATER_SPLASH            = 23
	WATER_SPLASH_MANUAL     = 24
	WATER_WAKE              = 25
	DRIP_WATER              = 26
	DRIP_LAVA               = 27
	DRIP_HONEY              = 28
	FALLING_DUST            = 29
	DUST                    = 29
	MOB_SPELL               = 30
	MOB_SPELL_AMBIENT       = 31
	MOB_SPELL_INSTANTANEOUS = 32
	INK                     = 33
	SLIME                   = 34
	RAIN_SPLASH             = 35
	VILLAGER_ANGRY          = 36
	VILLAGER_HAPPY          = 37
	ENCHANTMENT_TABLE       = 38
	TRACKING_EMITTER        = 39
	NOTE                    = 40
	WITCH_SPELL             = 41
	CARROT                  = 42
	MOB_APPEARANCE          = 43
	END_ROD                 = 44
	DRAGONS_BREATH          = 45
	SPIT                    = 46
	TOTEM                   = 47
	FOOD                    = 48
	FIREWORKS_STARTER       = 49
	FIREWORKS_SPARK         = 50
	FIREWORKS_OVERLAY       = 51
	BALLOON_GAS             = 52
	COLORED_FLAME           = 53
	SPARKLER                = 54
	CONDUIT                 = 55
	BUBBLE_COLUMN_UP        = 56
	BUBBLE_COLUMN_DOWN      = 57
	SNEEZE                  = 58
	SHULKER_BULLET          = 59
	BLEACH                  = 60
	DRAGON_DESTROY_BLOCK    = 61
	MYCELIUM_DUST           = 62
	FALLING_RED_DUST        = 63
	CAMPFIRE_SMOKE          = 64
	TALL_CAMPFIRE_SMOKE     = 65
	DRAGON_BREATH_FIRE      = 66
	DRAGON_BREATH_TRAIL     = 67
)

// new -> old (DowngradeParticleId input -> old constant output)
var newToOld = map[int32]int32{
	1:  BUBBLE,
	2:  BUBBLE_MANUAL,
	3:  CRITICAL,
	4:  BLOCK_FORCE_FIELD,
	5:  SMOKE,
	6:  EXPLODE,
	7:  EVAPORATION,
	8:  FLAME,
	10: LAVA,
	11: LARGE_SMOKE,
	12: REDSTONE,
	13: RISING_RED_DUST,
	14: ITEM_BREAK,
	15: SNOWBALL_POOF,
	16: HUGE_EXPLODE,
	17: HUGE_EXPLODE_SEED,
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
}

// old -> new (UpgradeParticleId input -> new id output)
var oldToNew = map[int32]int32{
	BUBBLE:                  1,
	BUBBLE_MANUAL:           2,
	CRITICAL:                3,
	BLOCK_FORCE_FIELD:       4,
	SMOKE:                   5,
	EXPLODE:                 6,
	EVAPORATION:             7,
	FLAME:                   8,
	LAVA:                    10,
	LARGE_SMOKE:             11,
	REDSTONE:                12,
	RISING_RED_DUST:         13,
	ITEM_BREAK:              14,
	SNOWBALL_POOF:           15,
	HUGE_EXPLODE:            16,
	HUGE_EXPLODE_SEED:       17,
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
}

func DowngradeParticleId(id int32) int32 {
	if oldID, ok := newToOld[id]; ok {
		return oldID
	}

	return id
}

func UpgradeParticleId(id int32) int32 {
	if newID, ok := oldToNew[id]; ok {
		return newID
	}

	return id
}
