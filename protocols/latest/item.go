package latest

import (
	_ "embed"

	"github.com/potlounge/multiversion/mapping"
)

const ItemVersion = 241

var (
	//go:embed required_item_list.json
	requiredItemList []byte
)

func NewItemMapping() mapping.Item {
	return mapping.NewItemMapping(requiredItemList, ItemVersion)
}
