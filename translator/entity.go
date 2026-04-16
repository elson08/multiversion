package translator

import (
	"fmt"
	"github.com/sandertv/gophertunnel/minecraft/nbt"
)

var CustomEntities = []string{
	"custom:icespikes",
	"custom:cottoncandy",
	"custom:glowingaqua",
	"custom:portal",
	"custom:razor",
	"custom:dragon",
	"custom:rainbow",
	"custom:bloody",
	"custom:autumn",
}

func TranslateEntityIdentifiers(data []byte) ([]byte, error) {
	var entityIdentifiers map[string]any
	if err := nbt.UnmarshalEncoding(data, &entityIdentifiers, nbt.NetworkLittleEndian); err != nil {
		return nil, fmt.Errorf("failed to unmarshal entity identifiers: %w", err)
	}

	if idList, ok := entityIdentifiers["idlist"].([]any); ok {
		for _, id := range CustomEntities {
			idList = append(idList, map[string]any{"id": id})
		}

		entityIdentifiers["idlist"] = idList
	}

	encoded, err := nbt.MarshalEncoding(entityIdentifiers, nbt.NetworkLittleEndian)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal entity identifiers: %w", err)
	}

	return encoded, nil
}
