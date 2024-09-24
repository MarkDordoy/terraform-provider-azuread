package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeDriveIdRootPermissionId{}

// MeDriveIdRootPermissionId is a struct representing the Resource ID for a Me Drive Id Root Permission
type MeDriveIdRootPermissionId struct {
	DriveId      string
	PermissionId string
}

// NewMeDriveIdRootPermissionID returns a new MeDriveIdRootPermissionId struct
func NewMeDriveIdRootPermissionID(driveId string, permissionId string) MeDriveIdRootPermissionId {
	return MeDriveIdRootPermissionId{
		DriveId:      driveId,
		PermissionId: permissionId,
	}
}

// ParseMeDriveIdRootPermissionID parses 'input' into a MeDriveIdRootPermissionId
func ParseMeDriveIdRootPermissionID(input string) (*MeDriveIdRootPermissionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeDriveIdRootPermissionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeDriveIdRootPermissionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeDriveIdRootPermissionIDInsensitively parses 'input' case-insensitively into a MeDriveIdRootPermissionId
// note: this method should only be used for API response data and not user input
func ParseMeDriveIdRootPermissionIDInsensitively(input string) (*MeDriveIdRootPermissionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeDriveIdRootPermissionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeDriveIdRootPermissionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeDriveIdRootPermissionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DriveId, ok = input.Parsed["driveId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "driveId", input)
	}

	if id.PermissionId, ok = input.Parsed["permissionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "permissionId", input)
	}

	return nil
}

// ValidateMeDriveIdRootPermissionID checks that 'input' can be parsed as a Me Drive Id Root Permission ID
func ValidateMeDriveIdRootPermissionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeDriveIdRootPermissionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Drive Id Root Permission ID
func (id MeDriveIdRootPermissionId) ID() string {
	fmtString := "/me/drives/%s/root/permissions/%s"
	return fmt.Sprintf(fmtString, id.DriveId, id.PermissionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Drive Id Root Permission ID
func (id MeDriveIdRootPermissionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("drives", "drives", "drives"),
		resourceids.UserSpecifiedSegment("driveId", "driveId"),
		resourceids.StaticSegment("root", "root", "root"),
		resourceids.StaticSegment("permissions", "permissions", "permissions"),
		resourceids.UserSpecifiedSegment("permissionId", "permissionId"),
	}
}

// String returns a human-readable description of this Me Drive Id Root Permission ID
func (id MeDriveIdRootPermissionId) String() string {
	components := []string{
		fmt.Sprintf("Drive: %q", id.DriveId),
		fmt.Sprintf("Permission: %q", id.PermissionId),
	}
	return fmt.Sprintf("Me Drive Id Root Permission (%s)", strings.Join(components, "\n"))
}