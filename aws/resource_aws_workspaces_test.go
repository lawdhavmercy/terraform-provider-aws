package aws

import (
	"testing"
)

// These tests need to be serialized, because they all rely on the IAM Role `workspaces_DefaultRole`.
func TestAccAwsWorkspaces(t *testing.T) {
	testCases := map[string]map[string]func(t *testing.T){
		"Directory": {
			"basic":      testAccAwsWorkspacesDirectory_basic,
			"disappears": testAccAwsWorkspacesDirectory_disappears,
			"subnetIds":  testAccAwsWorkspacesDirectory_subnetIds,
			"tags":       testAccAwsWorkspacesDirectory_tags,
		},
		"Workspace": {
			"basic":          testAccAwsWorkspacesWorkspace_basic,
			"tags":           testAccAwsWorkspacesWorkspace_Tags,
			"properties":     testAccAwsWorkspacesWorkspace_WorkspaceProperties,
			"rootVolumeSize": testAccAwsWorkspacesWorkspace_validateRootVolumeSize,
			"userVolumeSize": testAccAwsWorkspacesWorkspace_validateUserVolumeSize,
		},
	}

	for group, m := range testCases {
		m := m
		t.Run(group, func(t *testing.T) {
			for name, tc := range m {
				tc := tc
				t.Run(name, func(t *testing.T) {
					tc(t)
				})
			}
		})
	}
}
