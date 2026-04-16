package packet

import (
	"fmt"
	"github.com/potlounge/multiversion/protocols/v486/types"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

// CraftingData is sent by the server to let the client know all crafting data that the server maintains. This
// includes shapeless crafting, crafting table recipes, furnace recipes etc. Each crafting station's recipes
// are included in it.
type CraftingData struct {
	// Recipes is a list of all recipes available on the server. It includes among others shapeless, shaped
	// and furnace recipes. The client will only be able to craft these recipes.
	Recipes []types.Recipe
	// PotionRecipes is a list of all potion mixing recipes which may be used in the brewing stand.
	PotionRecipes []protocol.PotionRecipe
	// PotionContainerChangeRecipes is a list of all recipes to convert a potion from one type to another,
	// such as from a drinkable potion to a splash potion, or from a splash potion to a lingering potion.
	PotionContainerChangeRecipes []protocol.PotionContainerChangeRecipe
	// MaterialReducers is a list of all material reducers which is used in education edition chemistry.
	MaterialReducers []protocol.MaterialReducer
	// ClearRecipes indicates if all recipes currently active on the client should be cleaned. Doing this
	// means that the client will have no recipes active by itself: Any CraftingData packets previously sent
	// will also be discarded, and only the recipes in this CraftingData packet will be used.
	ClearRecipes bool
}

// ID ...
func (*CraftingData) ID() uint32 {
	return packet.IDCraftingData
}

// Marshal ...
func (pk *CraftingData) Marshal( /*w *protocol.Writer*/ io protocol.IO) {
	if w, ok := io.(*protocol.Writer); ok {
		l := uint32(len(pk.Recipes))
		w.Varuint32(&l)
		for _, recipe := range pk.Recipes {
			var c int32
			switch recipe.(type) {
			case *types.ShapedRecipe:
				c = protocol.RecipeShaped
			default:
				w.UnknownEnumOption(fmt.Sprintf("%T", recipe), "crafting recipe type")
			}
			w.Varint32(&c)
			recipe.Marshal(w)
		}
		protocol.Slice(w, &pk.PotionRecipes)
		protocol.Slice(w, &pk.PotionContainerChangeRecipes)
		protocol.FuncSlice(w, &pk.MaterialReducers, w.MaterialReducer)
		w.Bool(&pk.ClearRecipes)
	}

	if r, ok := io.(*protocol.Reader); ok {
		var length uint32
		r.Varuint32(&length)
		pk.Recipes = make([]types.Recipe, length)
		for i := uint32(0); i < length; i++ {
			var recipeType int32
			r.Varint32(&recipeType)

			var recipe types.Recipe
			switch recipeType {
			case protocol.RecipeShaped:
				recipe = &types.ShapedRecipe{}
			default:
				r.UnknownEnumOption(recipeType, "crafting data recipe type")
				return
			}
			recipe.Unmarshal(r)
			pk.Recipes[i] = recipe
		}
		protocol.Slice(r, &pk.PotionRecipes)
		protocol.Slice(r, &pk.PotionContainerChangeRecipes)
		protocol.FuncSlice(r, &pk.MaterialReducers, r.MaterialReducer)
		r.Bool(&pk.ClearRecipes)
	}
}
