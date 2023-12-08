package duit_core

import (
	"encoding/json"
	"errors"
)

type UiBuilder struct {
	root *DuitElementModel
}

// Build generates a JSON representation of the UiBuilder.
//
// It returns a byte slice and an error.
func (builder *UiBuilder) Build() ([]byte, error) {
	json, err := json.Marshal(*builder.root)

	if err != nil {
		return nil, errors.New("Failed to build JSON: " + err.Error())
	}

	return json, nil
}

func (builder *UiBuilder) CreateRoot() *DuitElementModel {
	builder.root = &DuitElementModel{
		ElementType: Column,
	}
	return builder.root
}

// CreateRootOfExactType creates a root element of the specified type and sets it as the root of the UiBuilder.
//
// Parameters:
// - elType: The type of the root element.
// - attributes: The attributes of the root element.
// - id: The ID of the root element.
//
// Returns:
// - *DuitElementModel: The root element created.
func (builder *UiBuilder) CreateRootOfExactType(elType DuitElementType, attributes interface{}, id string) *DuitElementModel {
	switch elType {
	case Column:
		builder.root = new(DuitElementModel).CreateElement(elType, id, "", attributes, nil, false, 2)
	case Row:
		builder.root = new(DuitElementModel).CreateElement(elType, id, "", attributes, nil, false, 2)
	case Stack:
		builder.root = new(DuitElementModel).CreateElement(elType, id, "", attributes, nil, false, 2)
	case Center:
		builder.root = new(DuitElementModel).CreateElement(elType, id, "", attributes, nil, false, 1)
	case ColoredBox:
		builder.root = new(DuitElementModel).CreateElement(elType, id, "", attributes, nil, false, 1)
	case Container:
		builder.root = new(DuitElementModel).CreateElement(elType, id, "", attributes, nil, false, 1)
	case DecoratedBox:
		builder.root = new(DuitElementModel).CreateElement(elType, id, "", attributes, nil, false, 1)
	default:
		panic("Invalid element type")
	}

	return builder.root
}
