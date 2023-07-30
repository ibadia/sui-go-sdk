// Copyright (c) BlockVision, Inc. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package sui

import (
	"context"

	"github.com/yasir7ca/sui-go-sdk/common/httpconn"
	"github.com/yasir7ca/sui-go-sdk/models"
)

type IReadObjectFromSuiAPI interface {
	SuiGetObject(ctx context.Context, req models.SuiGetObjectRequest) (models.SuiObjectData, error)
	SuiXGetOwnedObjects(ctx context.Context, req models.SuiXGetOwnedObjectsRequest) (models.PaginatedObjectsResponse, error)
	SuiMultiGetObjects(ctx context.Context, req models.SuiMultiGetObjectsRequest) ([]*models.SuiObjectResponse, error)
	SuiXGetDynamicField(ctx context.Context, req models.SuiXGetDynamicFieldRequest) (models.PaginatedDynamicFieldInfoResponse, error)
	SuiXGetDynamicFieldObject(ctx context.Context, req models.SuiXGetDynamicFieldObjectRequest) (models.SuiObjectResponse, error)
	SuiTryGetPastObject(ctx context.Context, req models.SuiTryGetPastObjectRequest) (models.PastObjectResponse, error)
	SuiGetLoadedChildObjects(ctx context.Context, req models.SuiGetLoadedChildObjectsRequest) (models.ChildObjectsResponse, error)
}

type suiReadObjectFromSuiImpl struct {
	conn *httpconn.HttpConn
}

// SuiGetObject implements the method `sui_getObject`, gets the object information for a specified object.
func (s *suiReadObjectFromSuiImpl) SuiGetObject(ctx context.Context, req models.SuiGetObjectRequest) (models.SuiObjectData, error) {
	var rsp models.SuiObjectData
	err := s.conn.CallContext(ctx, &rsp, httpconn.Operation{
		Method: "sui_getObject",
		Params: []interface{}{
			req.ObjectId,
			req.Options,
		},
	})
	if err != nil {
		return rsp, err
	}
	return rsp, nil
}

// SuiXGetOwnedObjects implements the method `suix_getOwnedObjects`, gets the list of objects owned by an address.
func (s *suiReadObjectFromSuiImpl) SuiXGetOwnedObjects(ctx context.Context, req models.SuiXGetOwnedObjectsRequest) (models.PaginatedObjectsResponse, error) {
	var rsp models.PaginatedObjectsResponse
	if err := validate.ValidateStruct(req); err != nil {
		return rsp, err
	}
	err := s.conn.CallContext(ctx, &rsp, httpconn.Operation{
		Method: "suix_getOwnedObjects",
		Params: []interface{}{
			req.Address,
			req.Query,
			req.Cursor,
			req.Limit,
		},
	})
	if err != nil {
		return rsp, err
	}
	return rsp, nil
}

// SuiMultiGetObjects implements the method `sui_multiGetObjects`, gets the object data for a list of objects.
func (s *suiReadObjectFromSuiImpl) SuiMultiGetObjects(ctx context.Context, req models.SuiMultiGetObjectsRequest) ([]*models.SuiObjectResponse, error) {
	var rsp []*models.SuiObjectResponse
	err := s.conn.CallContext(ctx, &rsp, httpconn.Operation{
		Method: "sui_multiGetObjects",
		Params: []interface{}{
			req.ObjectIds,
			req.Options,
		},
	})
	if err != nil {
		return rsp, err
	}
	return rsp, nil
}

// SuiXGetDynamicField implements the method `suix_getDynamicFields`, gets the list of dynamic field objects owned by an object.
func (s *suiReadObjectFromSuiImpl) SuiXGetDynamicField(ctx context.Context, req models.SuiXGetDynamicFieldRequest) (models.PaginatedDynamicFieldInfoResponse, error) {
	var rsp models.PaginatedDynamicFieldInfoResponse
	if err := validate.ValidateStruct(req); err != nil {
		return rsp, err
	}
	err := s.conn.CallContext(ctx, &rsp, httpconn.Operation{
		Method: "suix_getDynamicFields",
		Params: []interface{}{
			req.ObjectId,
			req.Cursor,
			req.Limit,
		},
	})
	if err != nil {
		return rsp, err
	}
	return rsp, nil
}

// SuiXGetDynamicFieldObject implements the method `suix_getDynamicFieldObject`, gets the dynamic field object information for a specified object.
func (s *suiReadObjectFromSuiImpl) SuiXGetDynamicFieldObject(ctx context.Context, req models.SuiXGetDynamicFieldObjectRequest) (models.SuiObjectResponse, error) {
	var rsp models.SuiObjectResponse
	err := s.conn.CallContext(ctx, &rsp, httpconn.Operation{
		Method: "suix_getDynamicFieldObject",
		Params: []interface{}{
			req.ObjectId,
			req.DynamicFieldName,
		},
	})
	if err != nil {
		return rsp, err
	}
	return rsp, nil
}

// SuiTryGetPastObject implements the method `sui_tryGetPastObject`, gets the object information for a specified version.
// There is no guarantee that objects with past versions can be retrieved by this API. The result may vary across nodes depending on their pruning policies.
func (s *suiReadObjectFromSuiImpl) SuiTryGetPastObject(ctx context.Context, req models.SuiTryGetPastObjectRequest) (models.PastObjectResponse, error) {
	var rsp models.PastObjectResponse
	err := s.conn.CallContext(ctx, &rsp, httpconn.Operation{
		Method: "sui_tryGetPastObject",
		Params: []interface{}{
			req.ObjectId,
			req.Version,
			req.Options,
		},
	})
	if err != nil {
		return rsp, err
	}
	return rsp, nil
}

// SuiGetLoadedChildObjects implements the method `sui_getLoadedChildObjects`, return the object information for a specified digest.
func (s *suiReadObjectFromSuiImpl) SuiGetLoadedChildObjects(ctx context.Context, req models.SuiGetLoadedChildObjectsRequest) (models.ChildObjectsResponse, error) {
	var rsp models.ChildObjectsResponse
	err := s.conn.CallContext(ctx, &rsp, httpconn.Operation{
		Method: "sui_getLoadedChildObjects",
		Params: []interface{}{
			req.Digest,
		},
	})
	if err != nil {
		return rsp, err
	}
	return rsp, nil
}
