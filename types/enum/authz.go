// Copyright 2022 Harness Inc. All rights reserved.
// Use of this source code is governed by the Polyform Free Trial License
// that can be found in the LICENSE.md file for this repository.

package enum

// Represents the different types of resources that can be guarded with permissions.
type ResourceType string

const (
	ResourceTypeSpace ResourceType = "SPACE"
	ResourceTypeRepo  ResourceType = "REPOSITORY"
	//   ResourceType_Branch ResourceType = "BRANCH"
)

// Represents the available permissions
type Permission string

const (
	// ----- SPACE -----
	PermissionSpaceCreate Permission = "space_create"
	PermissionSpaceView   Permission = "space_view"
	PermissionSpaceEdit   Permission = "space_edit"
	PermissionSpaceDelete Permission = "space_delete"

	// ----- REPOSITORY -----
	PermissionRepoCreate Permission = "repository_create"
	PermissionRepoView   Permission = "repository_view"
	PermissionRepoEdit   Permission = "repository_edit"
	PermissionRepoDelete Permission = "repository_delete"

	// ----- BRANCH -----
	// PermissionBranchCreate Permission = "branch_create"
	// PermissionBranchView   Permission = "branch_view"
	// PermissionBranchEdit   Permission = "branch_edit"
	// PermissionBranchDelete Permission = "branch_delete"
)

// Represents the type of the entity requesting permission
type PrincipalType string

const (
	// Represents actions executed by a loged-in user
	PrincipalTypeUser PrincipalType = "USER"

	// Represents actions executed by an entity with an api key
	PrincipalTypeApiKey PrincipalType = "API_KEY"
)
