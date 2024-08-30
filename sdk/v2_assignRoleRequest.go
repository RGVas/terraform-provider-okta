// DO NOT EDIT LOCAL SDK - USE v3 okta-sdk-golang FOR API CALLS THAT DO NOT EXIST IN LOCAL SDK
package sdk

type AssignRoleRequest struct {
	Type        string `json:"type,omitempty"`
	Role        string `json:"role,omitempty"`
	ResourceSet string `json:"resource-set,omitempty"`
}
