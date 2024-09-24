package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = DeviceManagementIntentAssignment{}

type DeviceManagementIntentAssignment struct {
	// The assignment target
	Target DeviceAndAppManagementAssignmentTarget `json:"target"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s DeviceManagementIntentAssignment) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = DeviceManagementIntentAssignment{}

func (s DeviceManagementIntentAssignment) MarshalJSON() ([]byte, error) {
	type wrapper DeviceManagementIntentAssignment
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DeviceManagementIntentAssignment: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DeviceManagementIntentAssignment: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.deviceManagementIntentAssignment"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DeviceManagementIntentAssignment: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &DeviceManagementIntentAssignment{}

func (s *DeviceManagementIntentAssignment) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Id        *string `json:"id,omitempty"`
		ODataId   *string `json:"@odata.id,omitempty"`
		ODataType *string `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling DeviceManagementIntentAssignment into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["target"]; ok {
		impl, err := UnmarshalDeviceAndAppManagementAssignmentTargetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Target' for 'DeviceManagementIntentAssignment': %+v", err)
		}
		s.Target = impl
	}

	return nil
}