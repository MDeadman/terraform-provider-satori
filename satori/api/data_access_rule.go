package api

import "strconv"

const DataAccessPermissionApiPrefix = "/api/data-access-permission"

type DataAccessPermission struct {
	Id               *string                   `json:"id,omitempty"`
	ParentId         *string                   `json:"parentId,omitempty"`
	Suspended        *bool                     `json:"suspended,omitempty"`
	Identity         *DataAccessIdentity       `json:"identity,omitempty"`
	AccessLevel      string                    `json:"accessLevel"`
	TimeLimit        DataAccessTimeLimit       `json:"timeLimit"`
	UnusedTimeLimit  DataAccessUnusedTimeLimit `json:"unusedTimeLimit"`
	SecurityPolicies *[]string                 `json:"securityPolicyIds,omitempty"`
}

func (c *Client) CreateDataAccessPermission(parentId string, input *DataAccessPermission) (*DataAccessPermission, error) {
	output := DataAccessPermission{}
	params := make(map[string]string, 1)
	params["parentId"] = parentId
	return &output, c.postJsonWithParams(DataAccessPermissionApiPrefix, &params, input, &output)
}

func (c *Client) UpdateDataAccessPermission(id string, input *DataAccessPermission) (*DataAccessPermission, error) {
	output := DataAccessPermission{}
	return &output, c.putJson(DataAccessPermissionApiPrefix, id, input, &output)
}

func (c *Client) UpdateDataAccessPermissionSuspendedStatus(id string, suspend bool) (*DataAccessPermission, error) {
	output := DataAccessPermission{}
	params := map[string]string{"shouldSuspend": strconv.FormatBool(suspend)}
	return &output, c.putWithParams(DataAccessPermissionApiPrefix+"/suspend", id, &params, &output)
}

func (c *Client) GetDataAccessPermission(id string) (*DataAccessPermission, error) {
	output := DataAccessPermission{}
	return &output, c.getJson(DataAccessPermissionApiPrefix, id, &output)
}

func (c *Client) DeleteDataAccessPermission(id string) error {
	return c.delete(DataAccessPermissionApiPrefix, id)
}
