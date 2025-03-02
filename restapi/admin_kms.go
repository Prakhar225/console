// This file is part of MinIO Console Server
// Copyright (c) 2022 MinIO, Inc.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.
//

package restapi

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/runtime/middleware"
	"github.com/minio/console/models"
	"github.com/minio/console/restapi/operations"
	kmsAPI "github.com/minio/console/restapi/operations/k_m_s"
	"github.com/minio/madmin-go/v2"
)

func registerKMSHandlers(api *operations.ConsoleAPI) {
	registerKMSStatusHandlers(api)
	registerKMSKeyHandlers(api)
	registerKMSPolicyHandlers(api)
	registerKMSIdentityHandlers(api)
}

func registerKMSStatusHandlers(api *operations.ConsoleAPI) {
	api.KmsKMSStatusHandler = kmsAPI.KMSStatusHandlerFunc(func(params kmsAPI.KMSStatusParams, session *models.Principal) middleware.Responder {
		resp, err := GetKMSStatusResponse(session, params)
		if err != nil {
			return kmsAPI.NewKMSStatusDefault(int(err.Code)).WithPayload(err)
		}
		return kmsAPI.NewKMSStatusOK().WithPayload(resp)
	})

	api.KmsKMSMetricsHandler = kmsAPI.KMSMetricsHandlerFunc(func(params kmsAPI.KMSMetricsParams, session *models.Principal) middleware.Responder {
		resp, err := GetKMSMetricsResponse(session, params)
		if err != nil {
			return kmsAPI.NewKMSMetricsDefault(int(err.Code)).WithPayload(err)
		}
		return kmsAPI.NewKMSMetricsOK().WithPayload(resp)
	})

	api.KmsKMSAPIsHandler = kmsAPI.KMSAPIsHandlerFunc(func(params kmsAPI.KMSAPIsParams, session *models.Principal) middleware.Responder {
		resp, err := GetKMSAPIsResponse(session, params)
		if err != nil {
			return kmsAPI.NewKMSAPIsDefault(int(err.Code)).WithPayload(err)
		}
		return kmsAPI.NewKMSAPIsOK().WithPayload(resp)
	})

	api.KmsKMSVersionHandler = kmsAPI.KMSVersionHandlerFunc(func(params kmsAPI.KMSVersionParams, session *models.Principal) middleware.Responder {
		resp, err := GetKMSVersionResponse(session, params)
		if err != nil {
			return kmsAPI.NewKMSVersionDefault(int(err.Code)).WithPayload(err)
		}
		return kmsAPI.NewKMSVersionOK().WithPayload(resp)
	})
}

func GetKMSStatusResponse(session *models.Principal, params kmsAPI.KMSStatusParams) (*models.KmsStatusResponse, *models.Error) {
	ctx, cancel := context.WithCancel(params.HTTPRequest.Context())
	defer cancel()
	mAdmin, err := NewMinioAdminClient(session)
	if err != nil {
		return nil, ErrorWithContext(ctx, err)
	}
	return kmsStatus(ctx, AdminClient{Client: mAdmin})
}

func kmsStatus(ctx context.Context, minioClient MinioAdmin) (*models.KmsStatusResponse, *models.Error) {
	st, err := minioClient.kmsStatus(ctx)
	if err != nil {
		return nil, ErrorWithContext(ctx, err)
	}
	return &models.KmsStatusResponse{
		DefaultKeyID: st.DefaultKeyID,
		Name:         st.Name,
		Endpoints:    parseStatusEndpoints(st.Endpoints),
	}, nil
}

func parseStatusEndpoints(endpoints map[string]madmin.ItemState) (kmsEndpoints []*models.KmsEndpoint) {
	for key, value := range endpoints {
		kmsEndpoints = append(kmsEndpoints, &models.KmsEndpoint{URL: key, Status: string(value)})
	}
	return kmsEndpoints
}

func GetKMSMetricsResponse(session *models.Principal, params kmsAPI.KMSMetricsParams) (*models.KmsMetricsResponse, *models.Error) {
	ctx, cancel := context.WithCancel(params.HTTPRequest.Context())
	defer cancel()
	mAdmin, err := NewMinioAdminClient(session)
	if err != nil {
		return nil, ErrorWithContext(ctx, err)
	}
	return kmsMetrics(ctx, AdminClient{Client: mAdmin})
}

func kmsMetrics(ctx context.Context, minioClient MinioAdmin) (*models.KmsMetricsResponse, *models.Error) {
	metrics, err := minioClient.kmsMetrics(ctx)
	if err != nil {
		return nil, ErrorWithContext(ctx, err)
	}
	return &models.KmsMetricsResponse{
		RequestOK:        &metrics.RequestOK,
		RequestErr:       &metrics.RequestErr,
		RequestFail:      &metrics.RequestFail,
		RequestActive:    &metrics.RequestActive,
		AuditEvents:      &metrics.AuditEvents,
		ErrorEvents:      &metrics.ErrorEvents,
		LatencyHistogram: nil,
		Uptime:           &metrics.UpTime,
		Cpus:             &metrics.CPUs,
		UsableCPUs:       &metrics.UsableCPUs,
		Threads:          &metrics.Threads,
		HeapAlloc:        &metrics.HeapAlloc,
		HeapObjects:      metrics.HeapObjects,
		StackAlloc:       &metrics.StackAlloc,
	}, nil
}

func GetKMSAPIsResponse(session *models.Principal, params kmsAPI.KMSAPIsParams) (*models.KmsAPIsResponse, *models.Error) {
	ctx, cancel := context.WithCancel(params.HTTPRequest.Context())
	defer cancel()
	mAdmin, err := NewMinioAdminClient(session)
	if err != nil {
		return nil, ErrorWithContext(ctx, err)
	}
	return kmsAPIs(ctx, AdminClient{Client: mAdmin})
}

func kmsAPIs(ctx context.Context, minioClient MinioAdmin) (*models.KmsAPIsResponse, *models.Error) {
	apis, err := minioClient.kmsAPIs(ctx)
	if err != nil {
		return nil, ErrorWithContext(ctx, err)
	}
	return &models.KmsAPIsResponse{
		Results: parseApis(apis),
	}, nil
}

func parseApis(apis []madmin.KMSAPI) (data []*models.KmsAPI) {
	for _, api := range apis {
		data = append(data, &models.KmsAPI{
			Method:  api.Method,
			Path:    api.Path,
			MaxBody: api.MaxBody,
			Timeout: api.Timeout,
		})
	}
	return data
}

func GetKMSVersionResponse(session *models.Principal, params kmsAPI.KMSVersionParams) (*models.KmsVersionResponse, *models.Error) {
	ctx, cancel := context.WithCancel(params.HTTPRequest.Context())
	defer cancel()
	mAdmin, err := NewMinioAdminClient(session)
	if err != nil {
		return nil, ErrorWithContext(ctx, err)
	}
	return kmsVersion(ctx, AdminClient{Client: mAdmin})
}

func kmsVersion(ctx context.Context, minioClient MinioAdmin) (*models.KmsVersionResponse, *models.Error) {
	version, err := minioClient.kmsVersion(ctx)
	if err != nil {
		return nil, ErrorWithContext(ctx, err)
	}
	return &models.KmsVersionResponse{
		Version: version.Version,
	}, nil
}

func registerKMSKeyHandlers(api *operations.ConsoleAPI) {
	api.KmsKMSCreateKeyHandler = kmsAPI.KMSCreateKeyHandlerFunc(func(params kmsAPI.KMSCreateKeyParams, session *models.Principal) middleware.Responder {
		err := GetKMSCreateKeyResponse(session, params)
		if err != nil {
			return kmsAPI.NewKMSCreateKeyDefault(int(err.Code)).WithPayload(err)
		}
		return kmsAPI.NewKMSCreateKeyCreated()
	})

	api.KmsKMSImportKeyHandler = kmsAPI.KMSImportKeyHandlerFunc(func(params kmsAPI.KMSImportKeyParams, session *models.Principal) middleware.Responder {
		err := GetKMSImportKeyResponse(session, params)
		if err != nil {
			return kmsAPI.NewKMSImportKeyDefault(int(err.Code)).WithPayload(err)
		}
		return kmsAPI.NewKMSImportKeyCreated()
	})

	api.KmsKMSListKeysHandler = kmsAPI.KMSListKeysHandlerFunc(func(params kmsAPI.KMSListKeysParams, session *models.Principal) middleware.Responder {
		resp, err := GetKMSListKeysResponse(session, params)
		if err != nil {
			return kmsAPI.NewKMSListKeysDefault(int(err.Code)).WithPayload(err)
		}
		return kmsAPI.NewKMSListKeysOK().WithPayload(resp)
	})

	api.KmsKMSKeyStatusHandler = kmsAPI.KMSKeyStatusHandlerFunc(func(params kmsAPI.KMSKeyStatusParams, session *models.Principal) middleware.Responder {
		resp, err := GetKMSKeyStatusResponse(session, params)
		if err != nil {
			return kmsAPI.NewKMSKeyStatusDefault(int(err.Code)).WithPayload(err)
		}
		return kmsAPI.NewKMSKeyStatusOK().WithPayload(resp)
	})

	api.KmsKMSDeleteKeyHandler = kmsAPI.KMSDeleteKeyHandlerFunc(func(params kmsAPI.KMSDeleteKeyParams, session *models.Principal) middleware.Responder {
		err := GetKMSDeleteKeyResponse(session, params)
		if err != nil {
			return kmsAPI.NewKMSDeleteKeyDefault(int(err.Code)).WithPayload(err)
		}
		return kmsAPI.NewKMSDeleteKeyOK()
	})
}

func GetKMSCreateKeyResponse(session *models.Principal, params kmsAPI.KMSCreateKeyParams) *models.Error {
	ctx, cancel := context.WithCancel(params.HTTPRequest.Context())
	defer cancel()
	mAdmin, err := NewMinioAdminClient(session)
	if err != nil {
		return ErrorWithContext(ctx, err)
	}
	return createKey(ctx, *params.Body.Key, AdminClient{Client: mAdmin})
}

func createKey(ctx context.Context, key string, minioClient MinioAdmin) *models.Error {
	if err := minioClient.createKey(ctx, key); err != nil {
		return ErrorWithContext(ctx, err)
	}
	return nil
}

func GetKMSImportKeyResponse(session *models.Principal, params kmsAPI.KMSImportKeyParams) *models.Error {
	ctx, cancel := context.WithCancel(params.HTTPRequest.Context())
	defer cancel()
	mAdmin, err := NewMinioAdminClient(session)
	if err != nil {
		return ErrorWithContext(ctx, err)
	}
	bytes, err := json.Marshal(params.Body)
	if err != nil {
		return ErrorWithContext(ctx, err)
	}
	return importKey(ctx, params.Name, bytes, AdminClient{Client: mAdmin})
}

func importKey(ctx context.Context, key string, bytes []byte, minioClient MinioAdmin) *models.Error {
	if err := minioClient.importKey(ctx, key, bytes); err != nil {
		return ErrorWithContext(ctx, err)
	}
	return nil
}

func GetKMSListKeysResponse(session *models.Principal, params kmsAPI.KMSListKeysParams) (*models.KmsListKeysResponse, *models.Error) {
	ctx, cancel := context.WithCancel(params.HTTPRequest.Context())
	defer cancel()
	mAdmin, err := NewMinioAdminClient(session)
	if err != nil {
		return nil, ErrorWithContext(ctx, err)
	}
	pattern := ""
	if params.Pattern != nil {
		pattern = *params.Pattern
	}
	return listKeys(ctx, pattern, AdminClient{Client: mAdmin})
}

func listKeys(ctx context.Context, pattern string, minioClient MinioAdmin) (*models.KmsListKeysResponse, *models.Error) {
	results, err := minioClient.listKeys(ctx, pattern)
	if err != nil {
		return nil, ErrorWithContext(ctx, err)
	}
	return &models.KmsListKeysResponse{Results: parseKeys(results)}, nil
}

func parseKeys(results []madmin.KMSKeyInfo) (data []*models.KmsKeyInfo) {
	for _, key := range results {
		data = append(data, &models.KmsKeyInfo{
			CreatedAt: key.CreatedAt,
			CreatedBy: key.CreatedBy,
			Name:      key.Name,
		})
	}
	return data
}

func GetKMSKeyStatusResponse(session *models.Principal, params kmsAPI.KMSKeyStatusParams) (*models.KmsKeyStatusResponse, *models.Error) {
	ctx, cancel := context.WithCancel(params.HTTPRequest.Context())
	defer cancel()
	mAdmin, err := NewMinioAdminClient(session)
	if err != nil {
		return nil, ErrorWithContext(ctx, err)
	}
	return keyStatus(ctx, params.Name, AdminClient{Client: mAdmin})
}

func keyStatus(ctx context.Context, key string, minioClient MinioAdmin) (*models.KmsKeyStatusResponse, *models.Error) {
	ks, err := minioClient.keyStatus(ctx, key)
	if err != nil {
		return nil, ErrorWithContext(ctx, err)
	}
	return &models.KmsKeyStatusResponse{
		KeyID:         ks.KeyID,
		EncryptionErr: ks.EncryptionErr,
		DecryptionErr: ks.DecryptionErr,
	}, nil
}

func GetKMSDeleteKeyResponse(session *models.Principal, params kmsAPI.KMSDeleteKeyParams) *models.Error {
	ctx, cancel := context.WithCancel(params.HTTPRequest.Context())
	defer cancel()
	mAdmin, err := NewMinioAdminClient(session)
	if err != nil {
		return ErrorWithContext(ctx, err)
	}
	return deleteKey(ctx, params.Name, AdminClient{Client: mAdmin})
}

func deleteKey(ctx context.Context, key string, minioClient MinioAdmin) *models.Error {
	if err := minioClient.deleteKey(ctx, key); err != nil {
		return ErrorWithContext(ctx, err)
	}
	return nil
}

func registerKMSPolicyHandlers(api *operations.ConsoleAPI) {
	api.KmsKMSSetPolicyHandler = kmsAPI.KMSSetPolicyHandlerFunc(func(params kmsAPI.KMSSetPolicyParams, session *models.Principal) middleware.Responder {
		err := GetKMSSetPolicyResponse(session, params)
		if err != nil {
			return kmsAPI.NewKMSSetPolicyDefault(int(err.Code)).WithPayload(err)
		}
		return kmsAPI.NewKMSSetPolicyOK()
	})

	api.KmsKMSAssignPolicyHandler = kmsAPI.KMSAssignPolicyHandlerFunc(func(params kmsAPI.KMSAssignPolicyParams, session *models.Principal) middleware.Responder {
		err := GetKMSAssignPolicyResponse(session, params)
		if err != nil {
			return kmsAPI.NewKMSAssignPolicyDefault(int(err.Code)).WithPayload(err)
		}
		return kmsAPI.NewKMSAssignPolicyOK()
	})

	api.KmsKMSDescribePolicyHandler = kmsAPI.KMSDescribePolicyHandlerFunc(func(params kmsAPI.KMSDescribePolicyParams, session *models.Principal) middleware.Responder {
		resp, err := GetKMSDescribePolicyResponse(session, params)
		if err != nil {
			return kmsAPI.NewKMSDescribePolicyDefault(int(err.Code)).WithPayload(err)
		}
		return kmsAPI.NewKMSDescribePolicyOK().WithPayload(resp)
	})

	api.KmsKMSGetPolicyHandler = kmsAPI.KMSGetPolicyHandlerFunc(func(params kmsAPI.KMSGetPolicyParams, session *models.Principal) middleware.Responder {
		resp, err := GetKMSGetPolicyResponse(session, params)
		if err != nil {
			return kmsAPI.NewKMSGetPolicyDefault(int(err.Code)).WithPayload(err)
		}
		return kmsAPI.NewKMSGetPolicyOK().WithPayload(resp)
	})

	api.KmsKMSListPoliciesHandler = kmsAPI.KMSListPoliciesHandlerFunc(func(params kmsAPI.KMSListPoliciesParams, session *models.Principal) middleware.Responder {
		resp, err := GetKMSListPoliciesResponse(session, params)
		if err != nil {
			return kmsAPI.NewKMSListPoliciesDefault(int(err.Code)).WithPayload(err)
		}
		return kmsAPI.NewKMSListPoliciesOK().WithPayload(resp)
	})

	api.KmsKMSDeletePolicyHandler = kmsAPI.KMSDeletePolicyHandlerFunc(func(params kmsAPI.KMSDeletePolicyParams, session *models.Principal) middleware.Responder {
		err := GetKMSDeletePolicyResponse(session, params)
		if err != nil {
			return kmsAPI.NewKMSDeletePolicyDefault(int(err.Code)).WithPayload(err)
		}
		return kmsAPI.NewKMSDeletePolicyOK()
	})
}

func GetKMSSetPolicyResponse(session *models.Principal, params kmsAPI.KMSSetPolicyParams) *models.Error {
	ctx, cancel := context.WithCancel(params.HTTPRequest.Context())
	defer cancel()
	mAdmin, err := NewMinioAdminClient(session)
	if err != nil {
		return ErrorWithContext(ctx, err)
	}
	bytes, err := json.Marshal(params.Body)
	if err != nil {
		return ErrorWithContext(ctx, err)
	}
	return setPolicy(ctx, *params.Body.Policy, bytes, AdminClient{Client: mAdmin})
}

func setPolicy(ctx context.Context, policy string, content []byte, minioClient MinioAdmin) *models.Error {
	if err := minioClient.setKMSPolicy(ctx, policy, content); err != nil {
		return ErrorWithContext(ctx, err)
	}
	return nil
}

func GetKMSAssignPolicyResponse(session *models.Principal, params kmsAPI.KMSAssignPolicyParams) *models.Error {
	ctx, cancel := context.WithCancel(params.HTTPRequest.Context())
	defer cancel()
	mAdmin, err := NewMinioAdminClient(session)
	if err != nil {
		return ErrorWithContext(ctx, err)
	}
	bytes, err := json.Marshal(params.Body)
	if err != nil {
		return ErrorWithContext(ctx, err)
	}
	return assignPolicy(ctx, params.Name, bytes, AdminClient{Client: mAdmin})
}

func assignPolicy(ctx context.Context, policy string, content []byte, minioClient MinioAdmin) *models.Error {
	if err := minioClient.assignPolicy(ctx, policy, content); err != nil {
		return ErrorWithContext(ctx, err)
	}
	return nil
}

func GetKMSDescribePolicyResponse(session *models.Principal, params kmsAPI.KMSDescribePolicyParams) (*models.KmsDescribePolicyResponse, *models.Error) {
	ctx, cancel := context.WithCancel(params.HTTPRequest.Context())
	defer cancel()
	mAdmin, err := NewMinioAdminClient(session)
	if err != nil {
		return nil, ErrorWithContext(ctx, err)
	}
	return describePolicy(ctx, params.Name, AdminClient{Client: mAdmin})
}

func describePolicy(ctx context.Context, policy string, minioClient MinioAdmin) (*models.KmsDescribePolicyResponse, *models.Error) {
	dp, err := minioClient.describePolicy(ctx, policy)
	if err != nil {
		return nil, ErrorWithContext(ctx, err)
	}
	return &models.KmsDescribePolicyResponse{
		Name:      dp.Name,
		CreatedAt: dp.CreatedAt,
		CreatedBy: dp.CreatedBy,
	}, nil
}

func GetKMSGetPolicyResponse(session *models.Principal, params kmsAPI.KMSGetPolicyParams) (*models.KmsGetPolicyResponse, *models.Error) {
	ctx, cancel := context.WithCancel(params.HTTPRequest.Context())
	defer cancel()
	mAdmin, err := NewMinioAdminClient(session)
	if err != nil {
		return nil, ErrorWithContext(ctx, err)
	}
	return getPolicy(ctx, params.Name, AdminClient{Client: mAdmin})
}

func getPolicy(ctx context.Context, policy string, minioClient MinioAdmin) (*models.KmsGetPolicyResponse, *models.Error) {
	p, err := minioClient.getKMSPolicy(ctx, policy)
	if err != nil {
		return nil, ErrorWithContext(ctx, err)
	}
	return &models.KmsGetPolicyResponse{
		Allow: p.Allow,
		Deny:  p.Deny,
	}, nil
}

func GetKMSListPoliciesResponse(session *models.Principal, params kmsAPI.KMSListPoliciesParams) (*models.KmsListPoliciesResponse, *models.Error) {
	ctx, cancel := context.WithCancel(params.HTTPRequest.Context())
	defer cancel()
	mAdmin, err := NewMinioAdminClient(session)
	if err != nil {
		return nil, ErrorWithContext(ctx, err)
	}
	pattern := ""
	if params.Pattern != nil {
		pattern = *params.Pattern
	}
	return listKMSPolicies(ctx, pattern, AdminClient{Client: mAdmin})
}

func listKMSPolicies(ctx context.Context, pattern string, minioClient MinioAdmin) (*models.KmsListPoliciesResponse, *models.Error) {
	results, err := minioClient.listKMSPolicies(ctx, pattern)
	if err != nil {
		return nil, ErrorWithContext(ctx, err)
	}
	return &models.KmsListPoliciesResponse{Results: parsePolicies(results)}, nil
}

func parsePolicies(results []madmin.KMSPolicyInfo) (data []*models.KmsPolicyInfo) {
	for _, policy := range results {
		data = append(data, &models.KmsPolicyInfo{
			CreatedAt: policy.CreatedAt,
			CreatedBy: policy.CreatedBy,
			Name:      policy.Name,
		})
	}
	return data
}

func GetKMSDeletePolicyResponse(session *models.Principal, params kmsAPI.KMSDeletePolicyParams) *models.Error {
	ctx, cancel := context.WithCancel(params.HTTPRequest.Context())
	defer cancel()
	mAdmin, err := NewMinioAdminClient(session)
	if err != nil {
		return ErrorWithContext(ctx, err)
	}
	return deletePolicy(ctx, params.Name, AdminClient{Client: mAdmin})
}

func deletePolicy(ctx context.Context, policy string, minioClient MinioAdmin) *models.Error {
	if err := minioClient.deletePolicy(ctx, policy); err != nil {
		return ErrorWithContext(ctx, err)
	}
	return nil
}

func registerKMSIdentityHandlers(api *operations.ConsoleAPI) {
	api.KmsKMSDescribeIdentityHandler = kmsAPI.KMSDescribeIdentityHandlerFunc(func(params kmsAPI.KMSDescribeIdentityParams, session *models.Principal) middleware.Responder {
		resp, err := GetKMSDescribeIdentityResponse(session, params)
		if err != nil {
			return kmsAPI.NewKMSDescribeIdentityDefault(int(err.Code)).WithPayload(err)
		}
		return kmsAPI.NewKMSDescribeIdentityOK().WithPayload(resp)
	})

	api.KmsKMSDescribeSelfIdentityHandler = kmsAPI.KMSDescribeSelfIdentityHandlerFunc(func(params kmsAPI.KMSDescribeSelfIdentityParams, session *models.Principal) middleware.Responder {
		resp, err := GetKMSDescribeSelfIdentityResponse(session, params)
		if err != nil {
			return kmsAPI.NewKMSDescribeSelfIdentityDefault(int(err.Code)).WithPayload(err)
		}
		return kmsAPI.NewKMSDescribeSelfIdentityOK().WithPayload(resp)
	})

	api.KmsKMSListIdentitiesHandler = kmsAPI.KMSListIdentitiesHandlerFunc(func(params kmsAPI.KMSListIdentitiesParams, session *models.Principal) middleware.Responder {
		resp, err := GetKMSListIdentitiesResponse(session, params)
		if err != nil {
			return kmsAPI.NewKMSListIdentitiesDefault(int(err.Code)).WithPayload(err)
		}
		return kmsAPI.NewKMSListIdentitiesOK().WithPayload(resp)
	})
	api.KmsKMSDeleteIdentityHandler = kmsAPI.KMSDeleteIdentityHandlerFunc(func(params kmsAPI.KMSDeleteIdentityParams, session *models.Principal) middleware.Responder {
		err := GetKMSDeleteIdentityResponse(session, params)
		if err != nil {
			return kmsAPI.NewKMSDeleteIdentityDefault(int(err.Code)).WithPayload(err)
		}
		return kmsAPI.NewKMSDeleteIdentityOK()
	})
}

func GetKMSDescribeIdentityResponse(session *models.Principal, params kmsAPI.KMSDescribeIdentityParams) (*models.KmsDescribeIdentityResponse, *models.Error) {
	ctx, cancel := context.WithCancel(params.HTTPRequest.Context())
	defer cancel()
	mAdmin, err := NewMinioAdminClient(session)
	if err != nil {
		return nil, ErrorWithContext(ctx, err)
	}
	return describeIdentity(ctx, params.Name, AdminClient{Client: mAdmin})
}

func describeIdentity(ctx context.Context, identity string, minioClient MinioAdmin) (*models.KmsDescribeIdentityResponse, *models.Error) {
	i, err := minioClient.describeIdentity(ctx, identity)
	if err != nil {
		return nil, ErrorWithContext(ctx, err)
	}
	return &models.KmsDescribeIdentityResponse{
		Policy:    i.Policy,
		Admin:     i.IsAdmin,
		Identity:  i.Identity,
		CreatedAt: i.CreatedAt,
		CreatedBy: i.CreatedBy,
	}, nil
}

func GetKMSDescribeSelfIdentityResponse(session *models.Principal, params kmsAPI.KMSDescribeSelfIdentityParams) (*models.KmsDescribeSelfIdentityResponse, *models.Error) {
	ctx, cancel := context.WithCancel(params.HTTPRequest.Context())
	defer cancel()
	mAdmin, err := NewMinioAdminClient(session)
	if err != nil {
		return nil, ErrorWithContext(ctx, err)
	}
	return describeSelfIdentity(ctx, AdminClient{Client: mAdmin})
}

func describeSelfIdentity(ctx context.Context, minioClient MinioAdmin) (*models.KmsDescribeSelfIdentityResponse, *models.Error) {
	i, err := minioClient.describeSelfIdentity(ctx)
	if err != nil {
		return nil, ErrorWithContext(ctx, err)
	}
	return &models.KmsDescribeSelfIdentityResponse{
		Policy: &models.KmsGetPolicyResponse{
			Allow: i.Policy.Allow,
			Deny:  i.Policy.Deny,
		},
		Identity:  i.Identity,
		Admin:     i.IsAdmin,
		CreatedAt: i.CreatedAt,
		CreatedBy: i.CreatedBy,
	}, nil
}

func GetKMSListIdentitiesResponse(session *models.Principal, params kmsAPI.KMSListIdentitiesParams) (*models.KmsListIdentitiesResponse, *models.Error) {
	ctx, cancel := context.WithCancel(params.HTTPRequest.Context())
	defer cancel()
	mAdmin, err := NewMinioAdminClient(session)
	if err != nil {
		return nil, ErrorWithContext(ctx, err)
	}
	pattern := ""
	if params.Pattern != nil {
		pattern = *params.Pattern
	}
	return listIdentities(ctx, pattern, AdminClient{Client: mAdmin})
}

func listIdentities(ctx context.Context, pattern string, minioClient MinioAdmin) (*models.KmsListIdentitiesResponse, *models.Error) {
	results, err := minioClient.listIdentities(ctx, pattern)
	if err != nil {
		return nil, ErrorWithContext(ctx, err)
	}
	return &models.KmsListIdentitiesResponse{Results: parseIdentities(results)}, nil
}

func parseIdentities(results []madmin.KMSIdentityInfo) (data []*models.KmsIdentityInfo) {
	for _, policy := range results {
		data = append(data, &models.KmsIdentityInfo{
			CreatedAt: policy.CreatedAt,
			CreatedBy: policy.CreatedBy,
			Identity:  policy.Identity,
			Error:     policy.Error,
			Policy:    policy.Policy,
		})
	}
	return data
}

func GetKMSDeleteIdentityResponse(session *models.Principal, params kmsAPI.KMSDeleteIdentityParams) *models.Error {
	ctx, cancel := context.WithCancel(params.HTTPRequest.Context())
	defer cancel()
	mAdmin, err := NewMinioAdminClient(session)
	if err != nil {
		return ErrorWithContext(ctx, err)
	}
	return deleteIdentity(ctx, params.Name, AdminClient{Client: mAdmin})
}

func deleteIdentity(ctx context.Context, identity string, minioClient MinioAdmin) *models.Error {
	if err := minioClient.deleteIdentity(ctx, identity); err != nil {
		return ErrorWithContext(ctx, err)
	}
	return nil
}
