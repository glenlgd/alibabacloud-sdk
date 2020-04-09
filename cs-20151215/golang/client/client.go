// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/tea"
	endpointutil "github.com/aliyun/endpoint-util/golang/service"
	roa "github.com/aliyun/tea-roa/golang/client"
	util "github.com/aliyun/tea-util/golang/service"
)

type ResumeComponentUpgradeRequest struct {
	Headers map[string]string `json:"headers" xml:"headers"`
}

func (s ResumeComponentUpgradeRequest) String() string {
	return tea.Prettify(s)
}

func (s ResumeComponentUpgradeRequest) GoString() string {
	return s.String()
}

func (s *ResumeComponentUpgradeRequest) SetHeaders(v map[string]string) *ResumeComponentUpgradeRequest {
	s.Headers = v
	return s
}

type ResumeComponentUpgradeResponseBody struct {
	RequestId *string `json:"RequestId" xml:"RequestId" require:"true"`
}

func (s ResumeComponentUpgradeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ResumeComponentUpgradeResponseBody) GoString() string {
	return s.String()
}

func (s *ResumeComponentUpgradeResponseBody) SetRequestId(v string) *ResumeComponentUpgradeResponseBody {
	s.RequestId = &v
	return s
}

type ResumeComponentUpgradeResponse struct {
	Headers map[string]string                   `json:"headers" xml:"headers" require:"true"`
	Body    *ResumeComponentUpgradeResponseBody `json:"body" xml:"body" require:"true"`
}

func (s ResumeComponentUpgradeResponse) String() string {
	return tea.Prettify(s)
}

func (s ResumeComponentUpgradeResponse) GoString() string {
	return s.String()
}

func (s *ResumeComponentUpgradeResponse) SetHeaders(v map[string]string) *ResumeComponentUpgradeResponse {
	s.Headers = v
	return s
}

func (s *ResumeComponentUpgradeResponse) SetBody(v *ResumeComponentUpgradeResponseBody) *ResumeComponentUpgradeResponse {
	s.Body = v
	return s
}

type PauseComponentUpgradeRequest struct {
	Headers map[string]string `json:"headers" xml:"headers"`
}

func (s PauseComponentUpgradeRequest) String() string {
	return tea.Prettify(s)
}

func (s PauseComponentUpgradeRequest) GoString() string {
	return s.String()
}

func (s *PauseComponentUpgradeRequest) SetHeaders(v map[string]string) *PauseComponentUpgradeRequest {
	s.Headers = v
	return s
}

type PauseComponentUpgradeResponseBody struct {
	RequestId *string `json:"RequestId" xml:"RequestId" require:"true"`
}

func (s PauseComponentUpgradeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PauseComponentUpgradeResponseBody) GoString() string {
	return s.String()
}

func (s *PauseComponentUpgradeResponseBody) SetRequestId(v string) *PauseComponentUpgradeResponseBody {
	s.RequestId = &v
	return s
}

type PauseComponentUpgradeResponse struct {
	Headers map[string]string                  `json:"headers" xml:"headers" require:"true"`
	Body    *PauseComponentUpgradeResponseBody `json:"body" xml:"body" require:"true"`
}

func (s PauseComponentUpgradeResponse) String() string {
	return tea.Prettify(s)
}

func (s PauseComponentUpgradeResponse) GoString() string {
	return s.String()
}

func (s *PauseComponentUpgradeResponse) SetHeaders(v map[string]string) *PauseComponentUpgradeResponse {
	s.Headers = v
	return s
}

func (s *PauseComponentUpgradeResponse) SetBody(v *PauseComponentUpgradeResponseBody) *PauseComponentUpgradeResponse {
	s.Body = v
	return s
}

type CancelComponentUpgradeRequest struct {
	Headers map[string]string `json:"headers" xml:"headers"`
}

func (s CancelComponentUpgradeRequest) String() string {
	return tea.Prettify(s)
}

func (s CancelComponentUpgradeRequest) GoString() string {
	return s.String()
}

func (s *CancelComponentUpgradeRequest) SetHeaders(v map[string]string) *CancelComponentUpgradeRequest {
	s.Headers = v
	return s
}

type CancelComponentUpgradeResponseBody struct {
	RequestId *string `json:"RequestId" xml:"RequestId" require:"true"`
}

func (s CancelComponentUpgradeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CancelComponentUpgradeResponseBody) GoString() string {
	return s.String()
}

func (s *CancelComponentUpgradeResponseBody) SetRequestId(v string) *CancelComponentUpgradeResponseBody {
	s.RequestId = &v
	return s
}

type CancelComponentUpgradeResponse struct {
	Headers map[string]string                   `json:"headers" xml:"headers" require:"true"`
	Body    *CancelComponentUpgradeResponseBody `json:"body" xml:"body" require:"true"`
}

func (s CancelComponentUpgradeResponse) String() string {
	return tea.Prettify(s)
}

func (s CancelComponentUpgradeResponse) GoString() string {
	return s.String()
}

func (s *CancelComponentUpgradeResponse) SetHeaders(v map[string]string) *CancelComponentUpgradeResponse {
	s.Headers = v
	return s
}

func (s *CancelComponentUpgradeResponse) SetBody(v *CancelComponentUpgradeResponseBody) *CancelComponentUpgradeResponse {
	s.Body = v
	return s
}

type CancelWorkflowRequest struct {
	Headers map[string]string `json:"headers" xml:"headers"`
}

func (s CancelWorkflowRequest) String() string {
	return tea.Prettify(s)
}

func (s CancelWorkflowRequest) GoString() string {
	return s.String()
}

func (s *CancelWorkflowRequest) SetHeaders(v map[string]string) *CancelWorkflowRequest {
	s.Headers = v
	return s
}

type CancelWorkflowResponse struct {
	Headers map[string]string `json:"headers" xml:"headers" require:"true"`
}

func (s CancelWorkflowResponse) String() string {
	return tea.Prettify(s)
}

func (s CancelWorkflowResponse) GoString() string {
	return s.String()
}

func (s *CancelWorkflowResponse) SetHeaders(v map[string]string) *CancelWorkflowResponse {
	s.Headers = v
	return s
}

type DescirbeWorkflowRequest struct {
	Headers map[string]string `json:"headers" xml:"headers"`
}

func (s DescirbeWorkflowRequest) String() string {
	return tea.Prettify(s)
}

func (s DescirbeWorkflowRequest) GoString() string {
	return s.String()
}

func (s *DescirbeWorkflowRequest) SetHeaders(v map[string]string) *DescirbeWorkflowRequest {
	s.Headers = v
	return s
}

type DescirbeWorkflowResponse struct {
	Headers map[string]string `json:"headers" xml:"headers" require:"true"`
}

func (s DescirbeWorkflowResponse) String() string {
	return tea.Prettify(s)
}

func (s DescirbeWorkflowResponse) GoString() string {
	return s.String()
}

func (s *DescirbeWorkflowResponse) SetHeaders(v map[string]string) *DescirbeWorkflowResponse {
	s.Headers = v
	return s
}

type RemoveWorkflowRequest struct {
	Headers map[string]string `json:"headers" xml:"headers"`
}

func (s RemoveWorkflowRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveWorkflowRequest) GoString() string {
	return s.String()
}

func (s *RemoveWorkflowRequest) SetHeaders(v map[string]string) *RemoveWorkflowRequest {
	s.Headers = v
	return s
}

type RemoveWorkflowResponse struct {
	Headers map[string]string `json:"headers" xml:"headers" require:"true"`
}

func (s RemoveWorkflowResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveWorkflowResponse) GoString() string {
	return s.String()
}

func (s *RemoveWorkflowResponse) SetHeaders(v map[string]string) *RemoveWorkflowResponse {
	s.Headers = v
	return s
}

type DescribeWorkflowsRequest struct {
	Headers map[string]string `json:"headers" xml:"headers"`
}

func (s DescribeWorkflowsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeWorkflowsRequest) GoString() string {
	return s.String()
}

func (s *DescribeWorkflowsRequest) SetHeaders(v map[string]string) *DescribeWorkflowsRequest {
	s.Headers = v
	return s
}

type DescribeWorkflowsResponse struct {
	Headers map[string]string `json:"headers" xml:"headers" require:"true"`
}

func (s DescribeWorkflowsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeWorkflowsResponse) GoString() string {
	return s.String()
}

func (s *DescribeWorkflowsResponse) SetHeaders(v map[string]string) *DescribeWorkflowsResponse {
	s.Headers = v
	return s
}

type StartWorkflowRequest struct {
	Headers map[string]string `json:"headers" xml:"headers"`
}

func (s StartWorkflowRequest) String() string {
	return tea.Prettify(s)
}

func (s StartWorkflowRequest) GoString() string {
	return s.String()
}

func (s *StartWorkflowRequest) SetHeaders(v map[string]string) *StartWorkflowRequest {
	s.Headers = v
	return s
}

type StartWorkflowResponse struct {
	Headers map[string]string `json:"headers" xml:"headers" require:"true"`
}

func (s StartWorkflowResponse) String() string {
	return tea.Prettify(s)
}

func (s StartWorkflowResponse) GoString() string {
	return s.String()
}

func (s *StartWorkflowResponse) SetHeaders(v map[string]string) *StartWorkflowResponse {
	s.Headers = v
	return s
}

type CreateServiceMeshRequest struct {
	Headers map[string]string `json:"headers" xml:"headers"`
}

func (s CreateServiceMeshRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceMeshRequest) GoString() string {
	return s.String()
}

func (s *CreateServiceMeshRequest) SetHeaders(v map[string]string) *CreateServiceMeshRequest {
	s.Headers = v
	return s
}

type CreateServiceMeshResponseBody struct {
	RequestId *string `json:"RequestId" xml:"RequestId" require:"true"`
}

func (s CreateServiceMeshResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceMeshResponseBody) GoString() string {
	return s.String()
}

func (s *CreateServiceMeshResponseBody) SetRequestId(v string) *CreateServiceMeshResponseBody {
	s.RequestId = &v
	return s
}

type CreateServiceMeshResponse struct {
	Headers map[string]string              `json:"headers" xml:"headers" require:"true"`
	Body    *CreateServiceMeshResponseBody `json:"body" xml:"body" require:"true"`
}

func (s CreateServiceMeshResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceMeshResponse) GoString() string {
	return s.String()
}

func (s *CreateServiceMeshResponse) SetHeaders(v map[string]string) *CreateServiceMeshResponse {
	s.Headers = v
	return s
}

func (s *CreateServiceMeshResponse) SetBody(v *CreateServiceMeshResponseBody) *CreateServiceMeshResponse {
	s.Body = v
	return s
}

type ServiceMeshAddClusterRequest struct {
	Headers map[string]string `json:"headers" xml:"headers"`
}

func (s ServiceMeshAddClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s ServiceMeshAddClusterRequest) GoString() string {
	return s.String()
}

func (s *ServiceMeshAddClusterRequest) SetHeaders(v map[string]string) *ServiceMeshAddClusterRequest {
	s.Headers = v
	return s
}

type ServiceMeshAddClusterResponseBody struct {
	RequestId *string `json:"RequestId" xml:"RequestId" require:"true"`
}

func (s ServiceMeshAddClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ServiceMeshAddClusterResponseBody) GoString() string {
	return s.String()
}

func (s *ServiceMeshAddClusterResponseBody) SetRequestId(v string) *ServiceMeshAddClusterResponseBody {
	s.RequestId = &v
	return s
}

type ServiceMeshAddClusterResponse struct {
	Headers map[string]string                  `json:"headers" xml:"headers" require:"true"`
	Body    *ServiceMeshAddClusterResponseBody `json:"body" xml:"body" require:"true"`
}

func (s ServiceMeshAddClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s ServiceMeshAddClusterResponse) GoString() string {
	return s.String()
}

func (s *ServiceMeshAddClusterResponse) SetHeaders(v map[string]string) *ServiceMeshAddClusterResponse {
	s.Headers = v
	return s
}

func (s *ServiceMeshAddClusterResponse) SetBody(v *ServiceMeshAddClusterResponseBody) *ServiceMeshAddClusterResponse {
	s.Body = v
	return s
}

type ServiceMeshRemoveClusterRequest struct {
	Headers map[string]string `json:"headers" xml:"headers"`
}

func (s ServiceMeshRemoveClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s ServiceMeshRemoveClusterRequest) GoString() string {
	return s.String()
}

func (s *ServiceMeshRemoveClusterRequest) SetHeaders(v map[string]string) *ServiceMeshRemoveClusterRequest {
	s.Headers = v
	return s
}

type ServiceMeshRemoveClusterResponseBody struct {
	RequestId *string `json:"RequestId" xml:"RequestId" require:"true"`
}

func (s ServiceMeshRemoveClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ServiceMeshRemoveClusterResponseBody) GoString() string {
	return s.String()
}

func (s *ServiceMeshRemoveClusterResponseBody) SetRequestId(v string) *ServiceMeshRemoveClusterResponseBody {
	s.RequestId = &v
	return s
}

type ServiceMeshRemoveClusterResponse struct {
	Headers map[string]string                     `json:"headers" xml:"headers" require:"true"`
	Body    *ServiceMeshRemoveClusterResponseBody `json:"body" xml:"body" require:"true"`
}

func (s ServiceMeshRemoveClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s ServiceMeshRemoveClusterResponse) GoString() string {
	return s.String()
}

func (s *ServiceMeshRemoveClusterResponse) SetHeaders(v map[string]string) *ServiceMeshRemoveClusterResponse {
	s.Headers = v
	return s
}

func (s *ServiceMeshRemoveClusterResponse) SetBody(v *ServiceMeshRemoveClusterResponseBody) *ServiceMeshRemoveClusterResponse {
	s.Body = v
	return s
}

type UpdateServiceMeshRequest struct {
	Headers map[string]string `json:"headers" xml:"headers"`
}

func (s UpdateServiceMeshRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceMeshRequest) GoString() string {
	return s.String()
}

func (s *UpdateServiceMeshRequest) SetHeaders(v map[string]string) *UpdateServiceMeshRequest {
	s.Headers = v
	return s
}

type UpdateServiceMeshResponseBody struct {
	RequestId *string `json:"RequestId" xml:"RequestId" require:"true"`
}

func (s UpdateServiceMeshResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceMeshResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateServiceMeshResponseBody) SetRequestId(v string) *UpdateServiceMeshResponseBody {
	s.RequestId = &v
	return s
}

type UpdateServiceMeshResponse struct {
	Headers map[string]string              `json:"headers" xml:"headers" require:"true"`
	Body    *UpdateServiceMeshResponseBody `json:"body" xml:"body" require:"true"`
}

func (s UpdateServiceMeshResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceMeshResponse) GoString() string {
	return s.String()
}

func (s *UpdateServiceMeshResponse) SetHeaders(v map[string]string) *UpdateServiceMeshResponse {
	s.Headers = v
	return s
}

func (s *UpdateServiceMeshResponse) SetBody(v *UpdateServiceMeshResponseBody) *UpdateServiceMeshResponse {
	s.Body = v
	return s
}

type ServiceMeshApiServerRequest struct {
	Headers map[string]string `json:"headers" xml:"headers"`
}

func (s ServiceMeshApiServerRequest) String() string {
	return tea.Prettify(s)
}

func (s ServiceMeshApiServerRequest) GoString() string {
	return s.String()
}

func (s *ServiceMeshApiServerRequest) SetHeaders(v map[string]string) *ServiceMeshApiServerRequest {
	s.Headers = v
	return s
}

type ServiceMeshApiServerResponse struct {
	Headers map[string]string `json:"headers" xml:"headers" require:"true"`
}

func (s ServiceMeshApiServerResponse) String() string {
	return tea.Prettify(s)
}

func (s ServiceMeshApiServerResponse) GoString() string {
	return s.String()
}

func (s *ServiceMeshApiServerResponse) SetHeaders(v map[string]string) *ServiceMeshApiServerResponse {
	s.Headers = v
	return s
}

type UnInstallClusterAddonsBody struct {
	Name *string `json:"name" xml:"name" require:"true"`
}

func (s UnInstallClusterAddonsBody) String() string {
	return tea.Prettify(s)
}

func (s UnInstallClusterAddonsBody) GoString() string {
	return s.String()
}

func (s *UnInstallClusterAddonsBody) SetName(v string) *UnInstallClusterAddonsBody {
	s.Name = &v
	return s
}

type UnInstallClusterAddonsRequest struct {
	Headers map[string]string           `json:"headers" xml:"headers"`
	Body    *UnInstallClusterAddonsBody `json:"body" xml:"body" require:"true"`
}

func (s UnInstallClusterAddonsRequest) String() string {
	return tea.Prettify(s)
}

func (s UnInstallClusterAddonsRequest) GoString() string {
	return s.String()
}

func (s *UnInstallClusterAddonsRequest) SetHeaders(v map[string]string) *UnInstallClusterAddonsRequest {
	s.Headers = v
	return s
}

func (s *UnInstallClusterAddonsRequest) SetBody(v *UnInstallClusterAddonsBody) *UnInstallClusterAddonsRequest {
	s.Body = v
	return s
}

type UnInstallClusterAddonsResponseBody struct {
	RequestId *string `json:"RequestId" xml:"RequestId" require:"true"`
}

func (s UnInstallClusterAddonsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UnInstallClusterAddonsResponseBody) GoString() string {
	return s.String()
}

func (s *UnInstallClusterAddonsResponseBody) SetRequestId(v string) *UnInstallClusterAddonsResponseBody {
	s.RequestId = &v
	return s
}

type UnInstallClusterAddonsResponse struct {
	Headers map[string]string                   `json:"headers" xml:"headers" require:"true"`
	Body    *UnInstallClusterAddonsResponseBody `json:"body" xml:"body" require:"true"`
}

func (s UnInstallClusterAddonsResponse) String() string {
	return tea.Prettify(s)
}

func (s UnInstallClusterAddonsResponse) GoString() string {
	return s.String()
}

func (s *UnInstallClusterAddonsResponse) SetHeaders(v map[string]string) *UnInstallClusterAddonsResponse {
	s.Headers = v
	return s
}

func (s *UnInstallClusterAddonsResponse) SetBody(v *UnInstallClusterAddonsResponseBody) *UnInstallClusterAddonsResponse {
	s.Body = v
	return s
}

type DescribeAddonsQuery struct {
	Region *string `json:"region" xml:"region" require:"true"`
}

func (s DescribeAddonsQuery) String() string {
	return tea.Prettify(s)
}

func (s DescribeAddonsQuery) GoString() string {
	return s.String()
}

func (s *DescribeAddonsQuery) SetRegion(v string) *DescribeAddonsQuery {
	s.Region = &v
	return s
}

type DescribeAddonsRequest struct {
	Query   *DescribeAddonsQuery `json:"query" xml:"query" require:"true"`
	Headers map[string]string    `json:"headers" xml:"headers"`
}

func (s DescribeAddonsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAddonsRequest) GoString() string {
	return s.String()
}

func (s *DescribeAddonsRequest) SetQuery(v *DescribeAddonsQuery) *DescribeAddonsRequest {
	s.Query = v
	return s
}

func (s *DescribeAddonsRequest) SetHeaders(v map[string]string) *DescribeAddonsRequest {
	s.Headers = v
	return s
}

type DescribeAddonsResponseBody struct {
	ComponentGroups    []*DescribeAddonsResponseBodyComponentGroups  `json:"ComponentGroups" xml:"ComponentGroups" require:"true" type:"Repeated"`
	StandardComponents *DescribeAddonsResponseBodyStandardComponents `json:"StandardComponents" xml:"StandardComponents" require:"true" type:"Struct"`
}

func (s DescribeAddonsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAddonsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAddonsResponseBody) SetComponentGroups(v []*DescribeAddonsResponseBodyComponentGroups) *DescribeAddonsResponseBody {
	s.ComponentGroups = v
	return s
}

func (s *DescribeAddonsResponseBody) SetStandardComponents(v *DescribeAddonsResponseBodyStandardComponents) *DescribeAddonsResponseBody {
	s.StandardComponents = v
	return s
}

type DescribeAddonsResponseBodyComponentGroups struct {
	GroupName *string                                           `json:"group_name" xml:"group_name" require:"true"`
	Items     []*DescribeAddonsResponseBodyComponentGroupsItems `json:"items" xml:"items" require:"true" type:"Repeated"`
	Default   []*string                                         `json:"default" xml:"default" require:"true" type:"Repeated"`
}

func (s DescribeAddonsResponseBodyComponentGroups) String() string {
	return tea.Prettify(s)
}

func (s DescribeAddonsResponseBodyComponentGroups) GoString() string {
	return s.String()
}

func (s *DescribeAddonsResponseBodyComponentGroups) SetGroupName(v string) *DescribeAddonsResponseBodyComponentGroups {
	s.GroupName = &v
	return s
}

func (s *DescribeAddonsResponseBodyComponentGroups) SetItems(v []*DescribeAddonsResponseBodyComponentGroupsItems) *DescribeAddonsResponseBodyComponentGroups {
	s.Items = v
	return s
}

func (s *DescribeAddonsResponseBodyComponentGroups) SetDefault(v []*string) *DescribeAddonsResponseBodyComponentGroups {
	s.Default = v
	return s
}

type DescribeAddonsResponseBodyComponentGroupsItems struct {
	Name     *string `json:"name" xml:"name" require:"true"`
	Config   *string `json:"config" xml:"config" require:"true"`
	Required *string `json:"required" xml:"required" require:"true"`
	Disabled *bool   `json:"disabled" xml:"disabled" require:"true"`
	Version  *string `json:"version" xml:"version" require:"true"`
}

func (s DescribeAddonsResponseBodyComponentGroupsItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeAddonsResponseBodyComponentGroupsItems) GoString() string {
	return s.String()
}

func (s *DescribeAddonsResponseBodyComponentGroupsItems) SetName(v string) *DescribeAddonsResponseBodyComponentGroupsItems {
	s.Name = &v
	return s
}

func (s *DescribeAddonsResponseBodyComponentGroupsItems) SetConfig(v string) *DescribeAddonsResponseBodyComponentGroupsItems {
	s.Config = &v
	return s
}

func (s *DescribeAddonsResponseBodyComponentGroupsItems) SetRequired(v string) *DescribeAddonsResponseBodyComponentGroupsItems {
	s.Required = &v
	return s
}

func (s *DescribeAddonsResponseBodyComponentGroupsItems) SetDisabled(v bool) *DescribeAddonsResponseBodyComponentGroupsItems {
	s.Disabled = &v
	return s
}

func (s *DescribeAddonsResponseBodyComponentGroupsItems) SetVersion(v string) *DescribeAddonsResponseBodyComponentGroupsItems {
	s.Version = &v
	return s
}

type DescribeAddonsResponseBodyStandardComponents struct {
	AddonName *DescribeAddonsResponseBodyStandardComponentsAddonName `json:"addon_name" xml:"addon_name" require:"true" type:"Struct"`
}

func (s DescribeAddonsResponseBodyStandardComponents) String() string {
	return tea.Prettify(s)
}

func (s DescribeAddonsResponseBodyStandardComponents) GoString() string {
	return s.String()
}

func (s *DescribeAddonsResponseBodyStandardComponents) SetAddonName(v *DescribeAddonsResponseBodyStandardComponentsAddonName) *DescribeAddonsResponseBodyStandardComponents {
	s.AddonName = v
	return s
}

type DescribeAddonsResponseBodyStandardComponentsAddonName struct {
	Name     *string `json:"name" xml:"name" require:"true"`
	Config   *string `json:"config" xml:"config" require:"true"`
	Required *string `json:"required" xml:"required" require:"true"`
	Disabled *bool   `json:"disabled" xml:"disabled" require:"true"`
	Version  *string `json:"version" xml:"version" require:"true"`
}

func (s DescribeAddonsResponseBodyStandardComponentsAddonName) String() string {
	return tea.Prettify(s)
}

func (s DescribeAddonsResponseBodyStandardComponentsAddonName) GoString() string {
	return s.String()
}

func (s *DescribeAddonsResponseBodyStandardComponentsAddonName) SetName(v string) *DescribeAddonsResponseBodyStandardComponentsAddonName {
	s.Name = &v
	return s
}

func (s *DescribeAddonsResponseBodyStandardComponentsAddonName) SetConfig(v string) *DescribeAddonsResponseBodyStandardComponentsAddonName {
	s.Config = &v
	return s
}

func (s *DescribeAddonsResponseBodyStandardComponentsAddonName) SetRequired(v string) *DescribeAddonsResponseBodyStandardComponentsAddonName {
	s.Required = &v
	return s
}

func (s *DescribeAddonsResponseBodyStandardComponentsAddonName) SetDisabled(v bool) *DescribeAddonsResponseBodyStandardComponentsAddonName {
	s.Disabled = &v
	return s
}

func (s *DescribeAddonsResponseBodyStandardComponentsAddonName) SetVersion(v string) *DescribeAddonsResponseBodyStandardComponentsAddonName {
	s.Version = &v
	return s
}

type DescribeAddonsResponse struct {
	Headers map[string]string           `json:"headers" xml:"headers" require:"true"`
	Body    *DescribeAddonsResponseBody `json:"body" xml:"body" require:"true"`
}

func (s DescribeAddonsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAddonsResponse) GoString() string {
	return s.String()
}

func (s *DescribeAddonsResponse) SetHeaders(v map[string]string) *DescribeAddonsResponse {
	s.Headers = v
	return s
}

func (s *DescribeAddonsResponse) SetBody(v *DescribeAddonsResponseBody) *DescribeAddonsResponse {
	s.Body = v
	return s
}

type UpdateK8sClusterUserConfigExpireRequest struct {
	Headers map[string]string `json:"headers" xml:"headers"`
}

func (s UpdateK8sClusterUserConfigExpireRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateK8sClusterUserConfigExpireRequest) GoString() string {
	return s.String()
}

func (s *UpdateK8sClusterUserConfigExpireRequest) SetHeaders(v map[string]string) *UpdateK8sClusterUserConfigExpireRequest {
	s.Headers = v
	return s
}

type UpdateK8sClusterUserConfigExpireResponseBody struct {
	RequestId *string `json:"RequestId" xml:"RequestId" require:"true"`
}

func (s UpdateK8sClusterUserConfigExpireResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateK8sClusterUserConfigExpireResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateK8sClusterUserConfigExpireResponseBody) SetRequestId(v string) *UpdateK8sClusterUserConfigExpireResponseBody {
	s.RequestId = &v
	return s
}

type UpdateK8sClusterUserConfigExpireResponse struct {
	Headers map[string]string                             `json:"headers" xml:"headers" require:"true"`
	Body    *UpdateK8sClusterUserConfigExpireResponseBody `json:"body" xml:"body" require:"true"`
}

func (s UpdateK8sClusterUserConfigExpireResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateK8sClusterUserConfigExpireResponse) GoString() string {
	return s.String()
}

func (s *UpdateK8sClusterUserConfigExpireResponse) SetHeaders(v map[string]string) *UpdateK8sClusterUserConfigExpireResponse {
	s.Headers = v
	return s
}

func (s *UpdateK8sClusterUserConfigExpireResponse) SetBody(v *UpdateK8sClusterUserConfigExpireResponseBody) *UpdateK8sClusterUserConfigExpireResponse {
	s.Body = v
	return s
}

type DescribeClusterV2UserKubeconfigQuery struct {
	PrivateIpAddress *bool `json:"PrivateIpAddress" xml:"PrivateIpAddress" require:"true"`
}

func (s DescribeClusterV2UserKubeconfigQuery) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterV2UserKubeconfigQuery) GoString() string {
	return s.String()
}

func (s *DescribeClusterV2UserKubeconfigQuery) SetPrivateIpAddress(v bool) *DescribeClusterV2UserKubeconfigQuery {
	s.PrivateIpAddress = &v
	return s
}

type DescribeClusterV2UserKubeconfigRequest struct {
	Query   *DescribeClusterV2UserKubeconfigQuery `json:"query" xml:"query" require:"true"`
	Headers map[string]string                     `json:"headers" xml:"headers"`
}

func (s DescribeClusterV2UserKubeconfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterV2UserKubeconfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeClusterV2UserKubeconfigRequest) SetQuery(v *DescribeClusterV2UserKubeconfigQuery) *DescribeClusterV2UserKubeconfigRequest {
	s.Query = v
	return s
}

func (s *DescribeClusterV2UserKubeconfigRequest) SetHeaders(v map[string]string) *DescribeClusterV2UserKubeconfigRequest {
	s.Headers = v
	return s
}

type DescribeClusterV2UserKubeconfigResponseBody struct {
	RequestId *string `json:"RequestId" xml:"RequestId" require:"true"`
}

func (s DescribeClusterV2UserKubeconfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterV2UserKubeconfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeClusterV2UserKubeconfigResponseBody) SetRequestId(v string) *DescribeClusterV2UserKubeconfigResponseBody {
	s.RequestId = &v
	return s
}

type DescribeClusterV2UserKubeconfigResponse struct {
	Headers map[string]string                            `json:"headers" xml:"headers" require:"true"`
	Body    *DescribeClusterV2UserKubeconfigResponseBody `json:"body" xml:"body" require:"true"`
}

func (s DescribeClusterV2UserKubeconfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterV2UserKubeconfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeClusterV2UserKubeconfigResponse) SetHeaders(v map[string]string) *DescribeClusterV2UserKubeconfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeClusterV2UserKubeconfigResponse) SetBody(v *DescribeClusterV2UserKubeconfigResponseBody) *DescribeClusterV2UserKubeconfigResponse {
	s.Body = v
	return s
}

type ModifyClusterTagsBody struct {
	Key   *string `json:"key" xml:"key" require:"true"`
	Value *string `json:"value" xml:"value" require:"true"`
}

func (s ModifyClusterTagsBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyClusterTagsBody) GoString() string {
	return s.String()
}

func (s *ModifyClusterTagsBody) SetKey(v string) *ModifyClusterTagsBody {
	s.Key = &v
	return s
}

func (s *ModifyClusterTagsBody) SetValue(v string) *ModifyClusterTagsBody {
	s.Value = &v
	return s
}

type ModifyClusterTagsRequest struct {
	Headers map[string]string      `json:"headers" xml:"headers"`
	Body    *ModifyClusterTagsBody `json:"body" xml:"body" require:"true"`
}

func (s ModifyClusterTagsRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyClusterTagsRequest) GoString() string {
	return s.String()
}

func (s *ModifyClusterTagsRequest) SetHeaders(v map[string]string) *ModifyClusterTagsRequest {
	s.Headers = v
	return s
}

func (s *ModifyClusterTagsRequest) SetBody(v *ModifyClusterTagsBody) *ModifyClusterTagsRequest {
	s.Body = v
	return s
}

type ModifyClusterTagsResponseBody struct {
	RequestId *string `json:"requestId" xml:"requestId" require:"true"`
}

func (s ModifyClusterTagsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyClusterTagsResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyClusterTagsResponseBody) SetRequestId(v string) *ModifyClusterTagsResponseBody {
	s.RequestId = &v
	return s
}

type ModifyClusterTagsResponse struct {
	Headers map[string]string              `json:"headers" xml:"headers" require:"true"`
	Body    *ModifyClusterTagsResponseBody `json:"body" xml:"body" require:"true"`
}

func (s ModifyClusterTagsResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyClusterTagsResponse) GoString() string {
	return s.String()
}

func (s *ModifyClusterTagsResponse) SetHeaders(v map[string]string) *ModifyClusterTagsResponse {
	s.Headers = v
	return s
}

func (s *ModifyClusterTagsResponse) SetBody(v *ModifyClusterTagsResponseBody) *ModifyClusterTagsResponse {
	s.Body = v
	return s
}

type ScaleOutClusterBody struct {
	Count                    *int                                  `json:"count" xml:"count" require:"true"`
	WorkerInstanceChargeType *string                               `json:"worker_instance_charge_type" xml:"worker_instance_charge_type" require:"true"`
	WorkerPeriod             *int                                  `json:"worker_period" xml:"worker_period" require:"true"`
	WorkerPeriodUnit         *string                               `json:"worker_period_unit" xml:"worker_period_unit" require:"true"`
	WorkerAutoRenew          *bool                                 `json:"worker_auto_renew" xml:"worker_auto_renew" require:"true"`
	WorkerAutoRenewPeriod    *int                                  `json:"worker_auto_renew_period" xml:"worker_auto_renew_period" require:"true"`
	WorkerSystemDiskCategory *string                               `json:"worker_system_disk_category" xml:"worker_system_disk_category" require:"true"`
	WorkerSystemDiskSize     *int                                  `json:"worker_system_disk_size" xml:"worker_system_disk_size" require:"true"`
	WorkerDataDisk           *bool                                 `json:"worker_data_disk" xml:"worker_data_disk" require:"true"`
	KeyPair                  *string                               `json:"key_pair" xml:"key_pair" require:"true"`
	LoginPassword            *string                               `json:"login_password" xml:"login_password" require:"true"`
	CloudMonitorFlags        *bool                                 `json:"cloud_monitor_flags" xml:"cloud_monitor_flags" require:"true"`
	CpuPolicy                *string                               `json:"cpu_policy" xml:"cpu_policy" require:"true"`
	DisableRollback          *bool                                 `json:"disable_rollback" xml:"disable_rollback" require:"true"`
	WorkerDataDisks          []*ScaleOutClusterBodyWorkerDataDisks `json:"worker_data_disks" xml:"worker_data_disks" require:"true" type:"Repeated"`
	Tags                     []*ScaleOutClusterBodyTags            `json:"tags" xml:"tags" require:"true" type:"Repeated"`
	Taints                   []*ScaleOutClusterBodyTaints          `json:"taints" xml:"taints" require:"true" type:"Repeated"`
	VswitchIds               []*string                             `json:"vswitch_ids" xml:"vswitch_ids" require:"true" type:"Repeated"`
	WorkerInstanceTypes      []*string                             `json:"worker_instance_types" xml:"worker_instance_types" require:"true" type:"Repeated"`
}

func (s ScaleOutClusterBody) String() string {
	return tea.Prettify(s)
}

func (s ScaleOutClusterBody) GoString() string {
	return s.String()
}

func (s *ScaleOutClusterBody) SetCount(v int) *ScaleOutClusterBody {
	s.Count = &v
	return s
}

func (s *ScaleOutClusterBody) SetWorkerInstanceChargeType(v string) *ScaleOutClusterBody {
	s.WorkerInstanceChargeType = &v
	return s
}

func (s *ScaleOutClusterBody) SetWorkerPeriod(v int) *ScaleOutClusterBody {
	s.WorkerPeriod = &v
	return s
}

func (s *ScaleOutClusterBody) SetWorkerPeriodUnit(v string) *ScaleOutClusterBody {
	s.WorkerPeriodUnit = &v
	return s
}

func (s *ScaleOutClusterBody) SetWorkerAutoRenew(v bool) *ScaleOutClusterBody {
	s.WorkerAutoRenew = &v
	return s
}

func (s *ScaleOutClusterBody) SetWorkerAutoRenewPeriod(v int) *ScaleOutClusterBody {
	s.WorkerAutoRenewPeriod = &v
	return s
}

func (s *ScaleOutClusterBody) SetWorkerSystemDiskCategory(v string) *ScaleOutClusterBody {
	s.WorkerSystemDiskCategory = &v
	return s
}

func (s *ScaleOutClusterBody) SetWorkerSystemDiskSize(v int) *ScaleOutClusterBody {
	s.WorkerSystemDiskSize = &v
	return s
}

func (s *ScaleOutClusterBody) SetWorkerDataDisk(v bool) *ScaleOutClusterBody {
	s.WorkerDataDisk = &v
	return s
}

func (s *ScaleOutClusterBody) SetKeyPair(v string) *ScaleOutClusterBody {
	s.KeyPair = &v
	return s
}

func (s *ScaleOutClusterBody) SetLoginPassword(v string) *ScaleOutClusterBody {
	s.LoginPassword = &v
	return s
}

func (s *ScaleOutClusterBody) SetCloudMonitorFlags(v bool) *ScaleOutClusterBody {
	s.CloudMonitorFlags = &v
	return s
}

func (s *ScaleOutClusterBody) SetCpuPolicy(v string) *ScaleOutClusterBody {
	s.CpuPolicy = &v
	return s
}

func (s *ScaleOutClusterBody) SetDisableRollback(v bool) *ScaleOutClusterBody {
	s.DisableRollback = &v
	return s
}

func (s *ScaleOutClusterBody) SetWorkerDataDisks(v []*ScaleOutClusterBodyWorkerDataDisks) *ScaleOutClusterBody {
	s.WorkerDataDisks = v
	return s
}

func (s *ScaleOutClusterBody) SetTags(v []*ScaleOutClusterBodyTags) *ScaleOutClusterBody {
	s.Tags = v
	return s
}

func (s *ScaleOutClusterBody) SetTaints(v []*ScaleOutClusterBodyTaints) *ScaleOutClusterBody {
	s.Taints = v
	return s
}

func (s *ScaleOutClusterBody) SetVswitchIds(v []*string) *ScaleOutClusterBody {
	s.VswitchIds = v
	return s
}

func (s *ScaleOutClusterBody) SetWorkerInstanceTypes(v []*string) *ScaleOutClusterBody {
	s.WorkerInstanceTypes = v
	return s
}

type ScaleOutClusterBodyWorkerDataDisks struct {
	Category  *string `json:"category" xml:"category" require:"true"`
	Size      *string `json:"size" xml:"size" require:"true"`
	Encrypted *string `json:"encrypted" xml:"encrypted" require:"true"`
}

func (s ScaleOutClusterBodyWorkerDataDisks) String() string {
	return tea.Prettify(s)
}

func (s ScaleOutClusterBodyWorkerDataDisks) GoString() string {
	return s.String()
}

func (s *ScaleOutClusterBodyWorkerDataDisks) SetCategory(v string) *ScaleOutClusterBodyWorkerDataDisks {
	s.Category = &v
	return s
}

func (s *ScaleOutClusterBodyWorkerDataDisks) SetSize(v string) *ScaleOutClusterBodyWorkerDataDisks {
	s.Size = &v
	return s
}

func (s *ScaleOutClusterBodyWorkerDataDisks) SetEncrypted(v string) *ScaleOutClusterBodyWorkerDataDisks {
	s.Encrypted = &v
	return s
}

type ScaleOutClusterBodyTags struct {
	Key   *string `json:"key" xml:"key" require:"true"`
	Value *string `json:"value" xml:"value" require:"true"`
}

func (s ScaleOutClusterBodyTags) String() string {
	return tea.Prettify(s)
}

func (s ScaleOutClusterBodyTags) GoString() string {
	return s.String()
}

func (s *ScaleOutClusterBodyTags) SetKey(v string) *ScaleOutClusterBodyTags {
	s.Key = &v
	return s
}

func (s *ScaleOutClusterBodyTags) SetValue(v string) *ScaleOutClusterBodyTags {
	s.Value = &v
	return s
}

type ScaleOutClusterBodyTaints struct {
	Key    *string `json:"key" xml:"key" require:"true"`
	Value  *string `json:"value" xml:"value" require:"true"`
	Effect *string `json:"effect" xml:"effect" require:"true"`
}

func (s ScaleOutClusterBodyTaints) String() string {
	return tea.Prettify(s)
}

func (s ScaleOutClusterBodyTaints) GoString() string {
	return s.String()
}

func (s *ScaleOutClusterBodyTaints) SetKey(v string) *ScaleOutClusterBodyTaints {
	s.Key = &v
	return s
}

func (s *ScaleOutClusterBodyTaints) SetValue(v string) *ScaleOutClusterBodyTaints {
	s.Value = &v
	return s
}

func (s *ScaleOutClusterBodyTaints) SetEffect(v string) *ScaleOutClusterBodyTaints {
	s.Effect = &v
	return s
}

type ScaleOutClusterRequest struct {
	Headers map[string]string    `json:"headers" xml:"headers"`
	Body    *ScaleOutClusterBody `json:"body" xml:"body" require:"true"`
}

func (s ScaleOutClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s ScaleOutClusterRequest) GoString() string {
	return s.String()
}

func (s *ScaleOutClusterRequest) SetHeaders(v map[string]string) *ScaleOutClusterRequest {
	s.Headers = v
	return s
}

func (s *ScaleOutClusterRequest) SetBody(v *ScaleOutClusterBody) *ScaleOutClusterRequest {
	s.Body = v
	return s
}

type ScaleOutClusterResponseBody struct {
	ClusterId  *string `json:"cluster_id" xml:"cluster_id" require:"true"`
	RequestId  *string `json:"request_id" xml:"request_id" require:"true"`
	TaskId     *string `json:"task_id" xml:"task_id" require:"true"`
	InstanceId *string `json:"instanceId" xml:"instanceId" require:"true"`
}

func (s ScaleOutClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ScaleOutClusterResponseBody) GoString() string {
	return s.String()
}

func (s *ScaleOutClusterResponseBody) SetClusterId(v string) *ScaleOutClusterResponseBody {
	s.ClusterId = &v
	return s
}

func (s *ScaleOutClusterResponseBody) SetRequestId(v string) *ScaleOutClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *ScaleOutClusterResponseBody) SetTaskId(v string) *ScaleOutClusterResponseBody {
	s.TaskId = &v
	return s
}

func (s *ScaleOutClusterResponseBody) SetInstanceId(v string) *ScaleOutClusterResponseBody {
	s.InstanceId = &v
	return s
}

type ScaleOutClusterResponse struct {
	Headers map[string]string            `json:"headers" xml:"headers" require:"true"`
	Body    *ScaleOutClusterResponseBody `json:"body" xml:"body" require:"true"`
}

func (s ScaleOutClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s ScaleOutClusterResponse) GoString() string {
	return s.String()
}

func (s *ScaleOutClusterResponse) SetHeaders(v map[string]string) *ScaleOutClusterResponse {
	s.Headers = v
	return s
}

func (s *ScaleOutClusterResponse) SetBody(v *ScaleOutClusterResponseBody) *ScaleOutClusterResponse {
	s.Body = v
	return s
}

type DescribeClusterAddonsVersionRequest struct {
	Headers map[string]string `json:"headers" xml:"headers"`
}

func (s DescribeClusterAddonsVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterAddonsVersionRequest) GoString() string {
	return s.String()
}

func (s *DescribeClusterAddonsVersionRequest) SetHeaders(v map[string]string) *DescribeClusterAddonsVersionRequest {
	s.Headers = v
	return s
}

type DescribeClusterAddonsVersionResponseBody struct {
	Template       *string `json:"template" xml:"template" require:"true"`
	ReadyToUpgrade *string `json:"ready_to_upgrade" xml:"ready_to_upgrade" require:"true"`
	NextVersion    *string `json:"next_version" xml:"next_version" require:"true"`
	CanUpgrade     *bool   `json:"can_upgrade" xml:"can_upgrade" require:"true"`
	ComponentName  *string `json:"component_name" xml:"component_name" require:"true"`
	Version        *string `json:"version" xml:"version" require:"true"`
	Changed        *string `json:"changed" xml:"changed" require:"true"`
	Message        *string `json:"message" xml:"message" require:"true"`
	Policy         *string `json:"policy" xml:"policy" require:"true"`
	Force          *bool   `json:"force" xml:"force" require:"true"`
	Required       *bool   `json:"required" xml:"required" require:"true"`
	Exist          *bool   `json:"exist" xml:"exist" require:"true"`
}

func (s DescribeClusterAddonsVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterAddonsVersionResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeClusterAddonsVersionResponseBody) SetTemplate(v string) *DescribeClusterAddonsVersionResponseBody {
	s.Template = &v
	return s
}

func (s *DescribeClusterAddonsVersionResponseBody) SetReadyToUpgrade(v string) *DescribeClusterAddonsVersionResponseBody {
	s.ReadyToUpgrade = &v
	return s
}

func (s *DescribeClusterAddonsVersionResponseBody) SetNextVersion(v string) *DescribeClusterAddonsVersionResponseBody {
	s.NextVersion = &v
	return s
}

func (s *DescribeClusterAddonsVersionResponseBody) SetCanUpgrade(v bool) *DescribeClusterAddonsVersionResponseBody {
	s.CanUpgrade = &v
	return s
}

func (s *DescribeClusterAddonsVersionResponseBody) SetComponentName(v string) *DescribeClusterAddonsVersionResponseBody {
	s.ComponentName = &v
	return s
}

func (s *DescribeClusterAddonsVersionResponseBody) SetVersion(v string) *DescribeClusterAddonsVersionResponseBody {
	s.Version = &v
	return s
}

func (s *DescribeClusterAddonsVersionResponseBody) SetChanged(v string) *DescribeClusterAddonsVersionResponseBody {
	s.Changed = &v
	return s
}

func (s *DescribeClusterAddonsVersionResponseBody) SetMessage(v string) *DescribeClusterAddonsVersionResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeClusterAddonsVersionResponseBody) SetPolicy(v string) *DescribeClusterAddonsVersionResponseBody {
	s.Policy = &v
	return s
}

func (s *DescribeClusterAddonsVersionResponseBody) SetForce(v bool) *DescribeClusterAddonsVersionResponseBody {
	s.Force = &v
	return s
}

func (s *DescribeClusterAddonsVersionResponseBody) SetRequired(v bool) *DescribeClusterAddonsVersionResponseBody {
	s.Required = &v
	return s
}

func (s *DescribeClusterAddonsVersionResponseBody) SetExist(v bool) *DescribeClusterAddonsVersionResponseBody {
	s.Exist = &v
	return s
}

type DescribeClusterAddonsVersionResponse struct {
	Headers map[string]string                         `json:"headers" xml:"headers" require:"true"`
	Body    *DescribeClusterAddonsVersionResponseBody `json:"body" xml:"body" require:"true"`
}

func (s DescribeClusterAddonsVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterAddonsVersionResponse) GoString() string {
	return s.String()
}

func (s *DescribeClusterAddonsVersionResponse) SetHeaders(v map[string]string) *DescribeClusterAddonsVersionResponse {
	s.Headers = v
	return s
}

func (s *DescribeClusterAddonsVersionResponse) SetBody(v *DescribeClusterAddonsVersionResponseBody) *DescribeClusterAddonsVersionResponse {
	s.Body = v
	return s
}

type DescribeClusterAddonUpgradeStatusRequest struct {
	Headers map[string]string `json:"headers" xml:"headers"`
}

func (s DescribeClusterAddonUpgradeStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterAddonUpgradeStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeClusterAddonUpgradeStatusRequest) SetHeaders(v map[string]string) *DescribeClusterAddonUpgradeStatusRequest {
	s.Headers = v
	return s
}

type DescribeClusterAddonUpgradeStatusResponseBody struct {
	ComponentId *DescribeClusterAddonUpgradeStatusResponseBodyComponentId `json:"ComponentId" xml:"ComponentId" require:"true" type:"Struct"`
}

func (s DescribeClusterAddonUpgradeStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterAddonUpgradeStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeClusterAddonUpgradeStatusResponseBody) SetComponentId(v *DescribeClusterAddonUpgradeStatusResponseBodyComponentId) *DescribeClusterAddonUpgradeStatusResponseBody {
	s.ComponentId = v
	return s
}

type DescribeClusterAddonUpgradeStatusResponseBodyComponentId struct {
	Template   *string                                                            `json:"template" xml:"template" require:"true"`
	CanUpgrade *bool                                                              `json:"can_upgrade" xml:"can_upgrade" require:"true"`
	Changed    *string                                                            `json:"changed" xml:"changed" require:"true"`
	AddonInfo  *DescribeClusterAddonUpgradeStatusResponseBodyComponentIdAddonInfo `json:"addon_info" xml:"addon_info" require:"true" type:"Struct"`
	Tasks      *DescribeClusterAddonUpgradeStatusResponseBodyComponentIdTasks     `json:"tasks" xml:"tasks" require:"true" type:"Struct"`
}

func (s DescribeClusterAddonUpgradeStatusResponseBodyComponentId) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterAddonUpgradeStatusResponseBodyComponentId) GoString() string {
	return s.String()
}

func (s *DescribeClusterAddonUpgradeStatusResponseBodyComponentId) SetTemplate(v string) *DescribeClusterAddonUpgradeStatusResponseBodyComponentId {
	s.Template = &v
	return s
}

func (s *DescribeClusterAddonUpgradeStatusResponseBodyComponentId) SetCanUpgrade(v bool) *DescribeClusterAddonUpgradeStatusResponseBodyComponentId {
	s.CanUpgrade = &v
	return s
}

func (s *DescribeClusterAddonUpgradeStatusResponseBodyComponentId) SetChanged(v string) *DescribeClusterAddonUpgradeStatusResponseBodyComponentId {
	s.Changed = &v
	return s
}

func (s *DescribeClusterAddonUpgradeStatusResponseBodyComponentId) SetAddonInfo(v *DescribeClusterAddonUpgradeStatusResponseBodyComponentIdAddonInfo) *DescribeClusterAddonUpgradeStatusResponseBodyComponentId {
	s.AddonInfo = v
	return s
}

func (s *DescribeClusterAddonUpgradeStatusResponseBodyComponentId) SetTasks(v *DescribeClusterAddonUpgradeStatusResponseBodyComponentIdTasks) *DescribeClusterAddonUpgradeStatusResponseBodyComponentId {
	s.Tasks = v
	return s
}

type DescribeClusterAddonUpgradeStatusResponseBodyComponentIdAddonInfo struct {
	ComponentName  *string `json:"component_name" xml:"component_name" require:"true"`
	ReadyToUpgrade *string `json:"ready_to_upgrade" xml:"ready_to_upgrade" require:"true"`
	Message        *string `json:"message" xml:"message" require:"true"`
	Version        *string `json:"version" xml:"version" require:"true"`
	Yaml           *string `json:"yaml" xml:"yaml" require:"true"`
}

func (s DescribeClusterAddonUpgradeStatusResponseBodyComponentIdAddonInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterAddonUpgradeStatusResponseBodyComponentIdAddonInfo) GoString() string {
	return s.String()
}

func (s *DescribeClusterAddonUpgradeStatusResponseBodyComponentIdAddonInfo) SetComponentName(v string) *DescribeClusterAddonUpgradeStatusResponseBodyComponentIdAddonInfo {
	s.ComponentName = &v
	return s
}

func (s *DescribeClusterAddonUpgradeStatusResponseBodyComponentIdAddonInfo) SetReadyToUpgrade(v string) *DescribeClusterAddonUpgradeStatusResponseBodyComponentIdAddonInfo {
	s.ReadyToUpgrade = &v
	return s
}

func (s *DescribeClusterAddonUpgradeStatusResponseBodyComponentIdAddonInfo) SetMessage(v string) *DescribeClusterAddonUpgradeStatusResponseBodyComponentIdAddonInfo {
	s.Message = &v
	return s
}

func (s *DescribeClusterAddonUpgradeStatusResponseBodyComponentIdAddonInfo) SetVersion(v string) *DescribeClusterAddonUpgradeStatusResponseBodyComponentIdAddonInfo {
	s.Version = &v
	return s
}

func (s *DescribeClusterAddonUpgradeStatusResponseBodyComponentIdAddonInfo) SetYaml(v string) *DescribeClusterAddonUpgradeStatusResponseBodyComponentIdAddonInfo {
	s.Yaml = &v
	return s
}

type DescribeClusterAddonUpgradeStatusResponseBodyComponentIdTasks struct {
	FinishedAt *string `json:"finished_at" xml:"finished_at" require:"true"`
	MasterUrl  *string `json:"master_url" xml:"master_url" require:"true"`
	CreatedAt  *string `json:"created_at" xml:"created_at" require:"true"`
	Status     *string `json:"status" xml:"status" require:"true"`
}

func (s DescribeClusterAddonUpgradeStatusResponseBodyComponentIdTasks) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterAddonUpgradeStatusResponseBodyComponentIdTasks) GoString() string {
	return s.String()
}

func (s *DescribeClusterAddonUpgradeStatusResponseBodyComponentIdTasks) SetFinishedAt(v string) *DescribeClusterAddonUpgradeStatusResponseBodyComponentIdTasks {
	s.FinishedAt = &v
	return s
}

func (s *DescribeClusterAddonUpgradeStatusResponseBodyComponentIdTasks) SetMasterUrl(v string) *DescribeClusterAddonUpgradeStatusResponseBodyComponentIdTasks {
	s.MasterUrl = &v
	return s
}

func (s *DescribeClusterAddonUpgradeStatusResponseBodyComponentIdTasks) SetCreatedAt(v string) *DescribeClusterAddonUpgradeStatusResponseBodyComponentIdTasks {
	s.CreatedAt = &v
	return s
}

func (s *DescribeClusterAddonUpgradeStatusResponseBodyComponentIdTasks) SetStatus(v string) *DescribeClusterAddonUpgradeStatusResponseBodyComponentIdTasks {
	s.Status = &v
	return s
}

type DescribeClusterAddonUpgradeStatusResponse struct {
	Headers map[string]string                              `json:"headers" xml:"headers" require:"true"`
	Body    *DescribeClusterAddonUpgradeStatusResponseBody `json:"body" xml:"body" require:"true"`
}

func (s DescribeClusterAddonUpgradeStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterAddonUpgradeStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeClusterAddonUpgradeStatusResponse) SetHeaders(v map[string]string) *DescribeClusterAddonUpgradeStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeClusterAddonUpgradeStatusResponse) SetBody(v *DescribeClusterAddonUpgradeStatusResponseBody) *DescribeClusterAddonUpgradeStatusResponse {
	s.Body = v
	return s
}

type DescribeEdgeClusterAttachScriptsQuery struct {
	NamePrefix *string `json:"NamePrefix" xml:"NamePrefix" require:"true"`
}

func (s DescribeEdgeClusterAttachScriptsQuery) String() string {
	return tea.Prettify(s)
}

func (s DescribeEdgeClusterAttachScriptsQuery) GoString() string {
	return s.String()
}

func (s *DescribeEdgeClusterAttachScriptsQuery) SetNamePrefix(v string) *DescribeEdgeClusterAttachScriptsQuery {
	s.NamePrefix = &v
	return s
}

type DescribeEdgeClusterAttachScriptsBody struct {
	Options *DescribeEdgeClusterAttachScriptsBodyOptions `json:"options" xml:"options" require:"true" type:"Struct"`
}

func (s DescribeEdgeClusterAttachScriptsBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeEdgeClusterAttachScriptsBody) GoString() string {
	return s.String()
}

func (s *DescribeEdgeClusterAttachScriptsBody) SetOptions(v *DescribeEdgeClusterAttachScriptsBodyOptions) *DescribeEdgeClusterAttachScriptsBody {
	s.Options = v
	return s
}

type DescribeEdgeClusterAttachScriptsBodyOptions struct {
	EnableIptables   *bool   `json:"enableIptables" xml:"enableIptables" require:"true"`
	FlannelIface     *string `json:"flannelIface" xml:"flannelIface" require:"true"`
	GpuVersion       *string `json:"gpuVersion" xml:"gpuVersion" require:"true"`
	ManageRuntime    *bool   `json:"manageRuntime" xml:"manageRuntime" require:"true"`
	NodeName         *string `json:"nodeName" xml:"nodeName" require:"true"`
	NodeNamePrefix   *string `json:"nodeNamePrefix" xml:"nodeNamePrefix" require:"true"`
	NodeNameStrategy *bool   `json:"nodeNameStrategy" xml:"nodeNameStrategy" require:"true"`
}

func (s DescribeEdgeClusterAttachScriptsBodyOptions) String() string {
	return tea.Prettify(s)
}

func (s DescribeEdgeClusterAttachScriptsBodyOptions) GoString() string {
	return s.String()
}

func (s *DescribeEdgeClusterAttachScriptsBodyOptions) SetEnableIptables(v bool) *DescribeEdgeClusterAttachScriptsBodyOptions {
	s.EnableIptables = &v
	return s
}

func (s *DescribeEdgeClusterAttachScriptsBodyOptions) SetFlannelIface(v string) *DescribeEdgeClusterAttachScriptsBodyOptions {
	s.FlannelIface = &v
	return s
}

func (s *DescribeEdgeClusterAttachScriptsBodyOptions) SetGpuVersion(v string) *DescribeEdgeClusterAttachScriptsBodyOptions {
	s.GpuVersion = &v
	return s
}

func (s *DescribeEdgeClusterAttachScriptsBodyOptions) SetManageRuntime(v bool) *DescribeEdgeClusterAttachScriptsBodyOptions {
	s.ManageRuntime = &v
	return s
}

func (s *DescribeEdgeClusterAttachScriptsBodyOptions) SetNodeName(v string) *DescribeEdgeClusterAttachScriptsBodyOptions {
	s.NodeName = &v
	return s
}

func (s *DescribeEdgeClusterAttachScriptsBodyOptions) SetNodeNamePrefix(v string) *DescribeEdgeClusterAttachScriptsBodyOptions {
	s.NodeNamePrefix = &v
	return s
}

func (s *DescribeEdgeClusterAttachScriptsBodyOptions) SetNodeNameStrategy(v bool) *DescribeEdgeClusterAttachScriptsBodyOptions {
	s.NodeNameStrategy = &v
	return s
}

type DescribeEdgeClusterAttachScriptsRequest struct {
	Query   *DescribeEdgeClusterAttachScriptsQuery `json:"query" xml:"query" require:"true"`
	Headers map[string]string                      `json:"headers" xml:"headers"`
	Body    *DescribeEdgeClusterAttachScriptsBody  `json:"body" xml:"body" require:"true"`
}

func (s DescribeEdgeClusterAttachScriptsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeEdgeClusterAttachScriptsRequest) GoString() string {
	return s.String()
}

func (s *DescribeEdgeClusterAttachScriptsRequest) SetQuery(v *DescribeEdgeClusterAttachScriptsQuery) *DescribeEdgeClusterAttachScriptsRequest {
	s.Query = v
	return s
}

func (s *DescribeEdgeClusterAttachScriptsRequest) SetHeaders(v map[string]string) *DescribeEdgeClusterAttachScriptsRequest {
	s.Headers = v
	return s
}

func (s *DescribeEdgeClusterAttachScriptsRequest) SetBody(v *DescribeEdgeClusterAttachScriptsBody) *DescribeEdgeClusterAttachScriptsRequest {
	s.Body = v
	return s
}

type DescribeEdgeClusterAttachScriptsResponseBody struct {
	RequestId *string `json:"RequestId" xml:"RequestId" require:"true"`
	Config    *string `json:"config" xml:"config" require:"true"`
}

func (s DescribeEdgeClusterAttachScriptsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeEdgeClusterAttachScriptsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEdgeClusterAttachScriptsResponseBody) SetRequestId(v string) *DescribeEdgeClusterAttachScriptsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEdgeClusterAttachScriptsResponseBody) SetConfig(v string) *DescribeEdgeClusterAttachScriptsResponseBody {
	s.Config = &v
	return s
}

type DescribeEdgeClusterAttachScriptsResponse struct {
	Headers map[string]string                             `json:"headers" xml:"headers" require:"true"`
	Body    *DescribeEdgeClusterAttachScriptsResponseBody `json:"body" xml:"body" require:"true"`
}

func (s DescribeEdgeClusterAttachScriptsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeEdgeClusterAttachScriptsResponse) GoString() string {
	return s.String()
}

func (s *DescribeEdgeClusterAttachScriptsResponse) SetHeaders(v map[string]string) *DescribeEdgeClusterAttachScriptsResponse {
	s.Headers = v
	return s
}

func (s *DescribeEdgeClusterAttachScriptsResponse) SetBody(v *DescribeEdgeClusterAttachScriptsResponseBody) *DescribeEdgeClusterAttachScriptsResponse {
	s.Body = v
	return s
}

type DeleteClusterNodesBody struct {
	ReleaseNode *string   `json:"release_node" xml:"release_node" require:"true"`
	Nodes       []*string `json:"nodes" xml:"nodes" require:"true" type:"Repeated"`
}

func (s DeleteClusterNodesBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteClusterNodesBody) GoString() string {
	return s.String()
}

func (s *DeleteClusterNodesBody) SetReleaseNode(v string) *DeleteClusterNodesBody {
	s.ReleaseNode = &v
	return s
}

func (s *DeleteClusterNodesBody) SetNodes(v []*string) *DeleteClusterNodesBody {
	s.Nodes = v
	return s
}

type DeleteClusterNodesRequest struct {
	Headers map[string]string       `json:"headers" xml:"headers"`
	Body    *DeleteClusterNodesBody `json:"body" xml:"body" require:"true"`
}

func (s DeleteClusterNodesRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteClusterNodesRequest) GoString() string {
	return s.String()
}

func (s *DeleteClusterNodesRequest) SetHeaders(v map[string]string) *DeleteClusterNodesRequest {
	s.Headers = v
	return s
}

func (s *DeleteClusterNodesRequest) SetBody(v *DeleteClusterNodesBody) *DeleteClusterNodesRequest {
	s.Body = v
	return s
}

type DeleteClusterNodesResponseBody struct {
	RequestId *string `json:"RequestId" xml:"RequestId" require:"true"`
}

func (s DeleteClusterNodesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteClusterNodesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteClusterNodesResponseBody) SetRequestId(v string) *DeleteClusterNodesResponseBody {
	s.RequestId = &v
	return s
}

type DeleteClusterNodesResponse struct {
	Headers map[string]string               `json:"headers" xml:"headers" require:"true"`
	Body    *DeleteClusterNodesResponseBody `json:"body" xml:"body" require:"true"`
}

func (s DeleteClusterNodesResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteClusterNodesResponse) GoString() string {
	return s.String()
}

func (s *DeleteClusterNodesResponse) SetHeaders(v map[string]string) *DeleteClusterNodesResponse {
	s.Headers = v
	return s
}

func (s *DeleteClusterNodesResponse) SetBody(v *DeleteClusterNodesResponseBody) *DeleteClusterNodesResponse {
	s.Body = v
	return s
}

type DescribeClusterUserKubeconfigQuery struct {
	PrivateIpAddress *bool `json:"PrivateIpAddress" xml:"PrivateIpAddress" require:"true"`
}

func (s DescribeClusterUserKubeconfigQuery) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterUserKubeconfigQuery) GoString() string {
	return s.String()
}

func (s *DescribeClusterUserKubeconfigQuery) SetPrivateIpAddress(v bool) *DescribeClusterUserKubeconfigQuery {
	s.PrivateIpAddress = &v
	return s
}

type DescribeClusterUserKubeconfigRequest struct {
	Query   *DescribeClusterUserKubeconfigQuery `json:"query" xml:"query" require:"true"`
	Headers map[string]string                   `json:"headers" xml:"headers"`
}

func (s DescribeClusterUserKubeconfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterUserKubeconfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeClusterUserKubeconfigRequest) SetQuery(v *DescribeClusterUserKubeconfigQuery) *DescribeClusterUserKubeconfigRequest {
	s.Query = v
	return s
}

func (s *DescribeClusterUserKubeconfigRequest) SetHeaders(v map[string]string) *DescribeClusterUserKubeconfigRequest {
	s.Headers = v
	return s
}

type DescribeClusterUserKubeconfigResponseBody struct {
	RequestId *string `json:"RequestId" xml:"RequestId" require:"true"`
	Config    *string `json:"config" xml:"config" require:"true"`
}

func (s DescribeClusterUserKubeconfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterUserKubeconfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeClusterUserKubeconfigResponseBody) SetRequestId(v string) *DescribeClusterUserKubeconfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeClusterUserKubeconfigResponseBody) SetConfig(v string) *DescribeClusterUserKubeconfigResponseBody {
	s.Config = &v
	return s
}

type DescribeClusterUserKubeconfigResponse struct {
	Headers map[string]string                          `json:"headers" xml:"headers" require:"true"`
	Body    *DescribeClusterUserKubeconfigResponseBody `json:"body" xml:"body" require:"true"`
}

func (s DescribeClusterUserKubeconfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterUserKubeconfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeClusterUserKubeconfigResponse) SetHeaders(v map[string]string) *DescribeClusterUserKubeconfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeClusterUserKubeconfigResponse) SetBody(v *DescribeClusterUserKubeconfigResponseBody) *DescribeClusterUserKubeconfigResponse {
	s.Body = v
	return s
}

type AttachInstancesBody struct {
	FormatDisk       *bool                      `json:"format_disk" xml:"format_disk" require:"true"`
	KeepInstanceName *bool                      `json:"keep_instance_name" xml:"keep_instance_name" require:"true"`
	CpuPolicy        *string                    `json:"cpu_policy" xml:"cpu_policy" require:"true"`
	KeyPair          *string                    `json:"key_pair" xml:"key_pair" require:"true"`
	Password         *string                    `json:"password" xml:"password" require:"true"`
	Tags             []*AttachInstancesBodyTags `json:"tags" xml:"tags" require:"true" type:"Repeated"`
	Instances        []*string                  `json:"instances" xml:"instances" require:"true" type:"Repeated"`
}

func (s AttachInstancesBody) String() string {
	return tea.Prettify(s)
}

func (s AttachInstancesBody) GoString() string {
	return s.String()
}

func (s *AttachInstancesBody) SetFormatDisk(v bool) *AttachInstancesBody {
	s.FormatDisk = &v
	return s
}

func (s *AttachInstancesBody) SetKeepInstanceName(v bool) *AttachInstancesBody {
	s.KeepInstanceName = &v
	return s
}

func (s *AttachInstancesBody) SetCpuPolicy(v string) *AttachInstancesBody {
	s.CpuPolicy = &v
	return s
}

func (s *AttachInstancesBody) SetKeyPair(v string) *AttachInstancesBody {
	s.KeyPair = &v
	return s
}

func (s *AttachInstancesBody) SetPassword(v string) *AttachInstancesBody {
	s.Password = &v
	return s
}

func (s *AttachInstancesBody) SetTags(v []*AttachInstancesBodyTags) *AttachInstancesBody {
	s.Tags = v
	return s
}

func (s *AttachInstancesBody) SetInstances(v []*string) *AttachInstancesBody {
	s.Instances = v
	return s
}

type AttachInstancesBodyTags struct {
	Key   *string `json:"key" xml:"key" require:"true"`
	Value *string `json:"value" xml:"value" require:"true"`
}

func (s AttachInstancesBodyTags) String() string {
	return tea.Prettify(s)
}

func (s AttachInstancesBodyTags) GoString() string {
	return s.String()
}

func (s *AttachInstancesBodyTags) SetKey(v string) *AttachInstancesBodyTags {
	s.Key = &v
	return s
}

func (s *AttachInstancesBodyTags) SetValue(v string) *AttachInstancesBodyTags {
	s.Value = &v
	return s
}

type AttachInstancesRequest struct {
	Headers map[string]string    `json:"headers" xml:"headers"`
	Body    *AttachInstancesBody `json:"body" xml:"body" require:"true"`
}

func (s AttachInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s AttachInstancesRequest) GoString() string {
	return s.String()
}

func (s *AttachInstancesRequest) SetHeaders(v map[string]string) *AttachInstancesRequest {
	s.Headers = v
	return s
}

func (s *AttachInstancesRequest) SetBody(v *AttachInstancesBody) *AttachInstancesRequest {
	s.Body = v
	return s
}

type AttachInstancesResponseBody struct {
	TaskId *string                          `json:"task_id" xml:"task_id" require:"true"`
	List   *AttachInstancesResponseBodyList `json:"list" xml:"list" require:"true" type:"Struct"`
}

func (s AttachInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AttachInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *AttachInstancesResponseBody) SetTaskId(v string) *AttachInstancesResponseBody {
	s.TaskId = &v
	return s
}

func (s *AttachInstancesResponseBody) SetList(v *AttachInstancesResponseBodyList) *AttachInstancesResponseBody {
	s.List = v
	return s
}

type AttachInstancesResponseBodyList struct {
	List []*AttachInstancesResponseBodyListList `json:"list" xml:"list" require:"true" type:"Repeated"`
}

func (s AttachInstancesResponseBodyList) String() string {
	return tea.Prettify(s)
}

func (s AttachInstancesResponseBodyList) GoString() string {
	return s.String()
}

func (s *AttachInstancesResponseBodyList) SetList(v []*AttachInstancesResponseBodyListList) *AttachInstancesResponseBodyList {
	s.List = v
	return s
}

type AttachInstancesResponseBodyListList struct {
	Code       *string `json:"code" xml:"code" require:"true"`
	InstanceId *string `json:"instanceId" xml:"instanceId" require:"true"`
	Message    *string `json:"message" xml:"message" require:"true"`
}

func (s AttachInstancesResponseBodyListList) String() string {
	return tea.Prettify(s)
}

func (s AttachInstancesResponseBodyListList) GoString() string {
	return s.String()
}

func (s *AttachInstancesResponseBodyListList) SetCode(v string) *AttachInstancesResponseBodyListList {
	s.Code = &v
	return s
}

func (s *AttachInstancesResponseBodyListList) SetInstanceId(v string) *AttachInstancesResponseBodyListList {
	s.InstanceId = &v
	return s
}

func (s *AttachInstancesResponseBodyListList) SetMessage(v string) *AttachInstancesResponseBodyListList {
	s.Message = &v
	return s
}

type AttachInstancesResponse struct {
	Headers map[string]string            `json:"headers" xml:"headers" require:"true"`
	Body    *AttachInstancesResponseBody `json:"body" xml:"body" require:"true"`
}

func (s AttachInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s AttachInstancesResponse) GoString() string {
	return s.String()
}

func (s *AttachInstancesResponse) SetHeaders(v map[string]string) *AttachInstancesResponse {
	s.Headers = v
	return s
}

func (s *AttachInstancesResponse) SetBody(v *AttachInstancesResponseBody) *AttachInstancesResponse {
	s.Body = v
	return s
}

type CreateClusterBody struct {
	ClusterType              *string                             `json:"cluster_type" xml:"cluster_type" require:"true"`
	Name                     *string                             `json:"name" xml:"name" require:"true"`
	RegionId                 *string                             `json:"region_id" xml:"region_id" require:"true"`
	KubernetesVersion        *string                             `json:"kubernetes_version" xml:"kubernetes_version" require:"true"`
	SnatEntry                *bool                               `json:"snat_entry" xml:"snat_entry" require:"true"`
	EndpointPublicAccess     *bool                               `json:"endpoint_public_access" xml:"endpoint_public_access" require:"true"`
	SshFlags                 *bool                               `json:"ssh_flags" xml:"ssh_flags" require:"true"`
	CloudMonitorFlags        *bool                               `json:"cloud_monitor_flags" xml:"cloud_monitor_flags" require:"true"`
	DeletionProtection       *bool                               `json:"deletion_protection" xml:"deletion_protection" require:"true"`
	NodeCidrMask             *string                             `json:"node_cidr_mask" xml:"node_cidr_mask" require:"true"`
	ProxyMode                *string                             `json:"proxy_mode" xml:"proxy_mode" require:"true"`
	OsType                   *string                             `json:"os_type" xml:"os_type" require:"true"`
	Platform                 *string                             `json:"platform" xml:"platform" require:"true"`
	NodePortRange            *string                             `json:"node_port_range" xml:"node_port_range" require:"true"`
	KeyPair                  *string                             `json:"key_pair" xml:"key_pair" require:"true"`
	LoginPassword            *string                             `json:"login_password" xml:"login_password" require:"true"`
	MasterInstanceChargeType *string                             `json:"master_instance_charge_type" xml:"master_instance_charge_type" require:"true"`
	WorkerInstanceChargeType *string                             `json:"worker_instance_charge_type" xml:"worker_instance_charge_type" require:"true"`
	MasterPeriod             *int                                `json:"master_period" xml:"master_period" require:"true"`
	WorkerPeriod             *int                                `json:"worker_period" xml:"worker_period" require:"true"`
	MasterPeriodUnit         *string                             `json:"master_period_unit" xml:"master_period_unit" require:"true"`
	WorkerPeriodUnit         *string                             `json:"worker_period_unit" xml:"worker_period_unit" require:"true"`
	MasterAutoRenew          *bool                               `json:"master_auto_renew" xml:"master_auto_renew" require:"true"`
	MasterAutoRenewPeriod    *int                                `json:"master_auto_renew_period" xml:"master_auto_renew_period" require:"true"`
	WorkerAutoRenew          *bool                               `json:"worker_auto_renew" xml:"worker_auto_renew" require:"true"`
	WorkerAutoRenewPeriod    *int                                `json:"worker_auto_renew_period" xml:"worker_auto_renew_period" require:"true"`
	CpuPolicy                *string                             `json:"cpu_policy" xml:"cpu_policy" require:"true"`
	MasterCount              *int                                `json:"master_count" xml:"master_count" require:"true"`
	MasterSystemDiskCategory *string                             `json:"master_system_disk_category" xml:"master_system_disk_category" require:"true"`
	MasterSystemDiskSize     *int                                `json:"master_system_disk_size" xml:"master_system_disk_size" require:"true"`
	Runtime                  map[string]interface{}              `json:"runtime" xml:"runtime" require:"true"`
	NumOfNodes               *int                                `json:"num_of_nodes" xml:"num_of_nodes" require:"true"`
	WorkerSystemDiskCategory *string                             `json:"worker_system_disk_category" xml:"worker_system_disk_category" require:"true"`
	WorkerSystemDiskSize     *int                                `json:"worker_system_disk_size" xml:"worker_system_disk_size" require:"true"`
	WorkerDataDisk           *bool                               `json:"worker_data_disk" xml:"worker_data_disk" require:"true"`
	Vpcid                    *string                             `json:"vpcid" xml:"vpcid" require:"true"`
	SecurityGroupId          *string                             `json:"security_group_id" xml:"security_group_id" require:"true"`
	ContainerCidr            *string                             `json:"container_cidr" xml:"container_cidr" require:"true"`
	ServiceCidr              *string                             `json:"service_cidr" xml:"service_cidr" require:"true"`
	DisableRollback          *bool                               `json:"disable_rollback" xml:"disable_rollback" require:"true"`
	TimeoutMins              *int                                `json:"timeout_mins" xml:"timeout_mins" require:"true"`
	Tags                     []*CreateClusterBodyTags            `json:"tags" xml:"tags" require:"true" type:"Repeated"`
	Addons                   []*CreateClusterBodyAddons          `json:"addons" xml:"addons" require:"true" type:"Repeated"`
	Taints                   []*CreateClusterBodyTaints          `json:"taints" xml:"taints" require:"true" type:"Repeated"`
	WorkerDataDisks          []*CreateClusterBodyWorkerDataDisks `json:"worker_data_disks" xml:"worker_data_disks" require:"true" type:"Repeated"`
	MasterVswitchIds         []*string                           `json:"master_vswitch_ids" xml:"master_vswitch_ids" require:"true" type:"Repeated"`
	MasterInstanceTypes      []*string                           `json:"master_instance_types" xml:"master_instance_types" require:"true" type:"Repeated"`
	WorkerInstanceTypes      []*string                           `json:"worker_instance_types" xml:"worker_instance_types" require:"true" type:"Repeated"`
	WorkerVswitchIds         []*string                           `json:"worker_vswitch_ids" xml:"worker_vswitch_ids" require:"true" type:"Repeated"`
}

func (s CreateClusterBody) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterBody) GoString() string {
	return s.String()
}

func (s *CreateClusterBody) SetClusterType(v string) *CreateClusterBody {
	s.ClusterType = &v
	return s
}

func (s *CreateClusterBody) SetName(v string) *CreateClusterBody {
	s.Name = &v
	return s
}

func (s *CreateClusterBody) SetRegionId(v string) *CreateClusterBody {
	s.RegionId = &v
	return s
}

func (s *CreateClusterBody) SetKubernetesVersion(v string) *CreateClusterBody {
	s.KubernetesVersion = &v
	return s
}

func (s *CreateClusterBody) SetSnatEntry(v bool) *CreateClusterBody {
	s.SnatEntry = &v
	return s
}

func (s *CreateClusterBody) SetEndpointPublicAccess(v bool) *CreateClusterBody {
	s.EndpointPublicAccess = &v
	return s
}

func (s *CreateClusterBody) SetSshFlags(v bool) *CreateClusterBody {
	s.SshFlags = &v
	return s
}

func (s *CreateClusterBody) SetCloudMonitorFlags(v bool) *CreateClusterBody {
	s.CloudMonitorFlags = &v
	return s
}

func (s *CreateClusterBody) SetDeletionProtection(v bool) *CreateClusterBody {
	s.DeletionProtection = &v
	return s
}

func (s *CreateClusterBody) SetNodeCidrMask(v string) *CreateClusterBody {
	s.NodeCidrMask = &v
	return s
}

func (s *CreateClusterBody) SetProxyMode(v string) *CreateClusterBody {
	s.ProxyMode = &v
	return s
}

func (s *CreateClusterBody) SetOsType(v string) *CreateClusterBody {
	s.OsType = &v
	return s
}

func (s *CreateClusterBody) SetPlatform(v string) *CreateClusterBody {
	s.Platform = &v
	return s
}

func (s *CreateClusterBody) SetNodePortRange(v string) *CreateClusterBody {
	s.NodePortRange = &v
	return s
}

func (s *CreateClusterBody) SetKeyPair(v string) *CreateClusterBody {
	s.KeyPair = &v
	return s
}

func (s *CreateClusterBody) SetLoginPassword(v string) *CreateClusterBody {
	s.LoginPassword = &v
	return s
}

func (s *CreateClusterBody) SetMasterInstanceChargeType(v string) *CreateClusterBody {
	s.MasterInstanceChargeType = &v
	return s
}

func (s *CreateClusterBody) SetWorkerInstanceChargeType(v string) *CreateClusterBody {
	s.WorkerInstanceChargeType = &v
	return s
}

func (s *CreateClusterBody) SetMasterPeriod(v int) *CreateClusterBody {
	s.MasterPeriod = &v
	return s
}

func (s *CreateClusterBody) SetWorkerPeriod(v int) *CreateClusterBody {
	s.WorkerPeriod = &v
	return s
}

func (s *CreateClusterBody) SetMasterPeriodUnit(v string) *CreateClusterBody {
	s.MasterPeriodUnit = &v
	return s
}

func (s *CreateClusterBody) SetWorkerPeriodUnit(v string) *CreateClusterBody {
	s.WorkerPeriodUnit = &v
	return s
}

func (s *CreateClusterBody) SetMasterAutoRenew(v bool) *CreateClusterBody {
	s.MasterAutoRenew = &v
	return s
}

func (s *CreateClusterBody) SetMasterAutoRenewPeriod(v int) *CreateClusterBody {
	s.MasterAutoRenewPeriod = &v
	return s
}

func (s *CreateClusterBody) SetWorkerAutoRenew(v bool) *CreateClusterBody {
	s.WorkerAutoRenew = &v
	return s
}

func (s *CreateClusterBody) SetWorkerAutoRenewPeriod(v int) *CreateClusterBody {
	s.WorkerAutoRenewPeriod = &v
	return s
}

func (s *CreateClusterBody) SetCpuPolicy(v string) *CreateClusterBody {
	s.CpuPolicy = &v
	return s
}

func (s *CreateClusterBody) SetMasterCount(v int) *CreateClusterBody {
	s.MasterCount = &v
	return s
}

func (s *CreateClusterBody) SetMasterSystemDiskCategory(v string) *CreateClusterBody {
	s.MasterSystemDiskCategory = &v
	return s
}

func (s *CreateClusterBody) SetMasterSystemDiskSize(v int) *CreateClusterBody {
	s.MasterSystemDiskSize = &v
	return s
}

func (s *CreateClusterBody) SetRuntime(v map[string]interface{}) *CreateClusterBody {
	s.Runtime = v
	return s
}

func (s *CreateClusterBody) SetNumOfNodes(v int) *CreateClusterBody {
	s.NumOfNodes = &v
	return s
}

func (s *CreateClusterBody) SetWorkerSystemDiskCategory(v string) *CreateClusterBody {
	s.WorkerSystemDiskCategory = &v
	return s
}

func (s *CreateClusterBody) SetWorkerSystemDiskSize(v int) *CreateClusterBody {
	s.WorkerSystemDiskSize = &v
	return s
}

func (s *CreateClusterBody) SetWorkerDataDisk(v bool) *CreateClusterBody {
	s.WorkerDataDisk = &v
	return s
}

func (s *CreateClusterBody) SetVpcid(v string) *CreateClusterBody {
	s.Vpcid = &v
	return s
}

func (s *CreateClusterBody) SetSecurityGroupId(v string) *CreateClusterBody {
	s.SecurityGroupId = &v
	return s
}

func (s *CreateClusterBody) SetContainerCidr(v string) *CreateClusterBody {
	s.ContainerCidr = &v
	return s
}

func (s *CreateClusterBody) SetServiceCidr(v string) *CreateClusterBody {
	s.ServiceCidr = &v
	return s
}

func (s *CreateClusterBody) SetDisableRollback(v bool) *CreateClusterBody {
	s.DisableRollback = &v
	return s
}

func (s *CreateClusterBody) SetTimeoutMins(v int) *CreateClusterBody {
	s.TimeoutMins = &v
	return s
}

func (s *CreateClusterBody) SetTags(v []*CreateClusterBodyTags) *CreateClusterBody {
	s.Tags = v
	return s
}

func (s *CreateClusterBody) SetAddons(v []*CreateClusterBodyAddons) *CreateClusterBody {
	s.Addons = v
	return s
}

func (s *CreateClusterBody) SetTaints(v []*CreateClusterBodyTaints) *CreateClusterBody {
	s.Taints = v
	return s
}

func (s *CreateClusterBody) SetWorkerDataDisks(v []*CreateClusterBodyWorkerDataDisks) *CreateClusterBody {
	s.WorkerDataDisks = v
	return s
}

func (s *CreateClusterBody) SetMasterVswitchIds(v []*string) *CreateClusterBody {
	s.MasterVswitchIds = v
	return s
}

func (s *CreateClusterBody) SetMasterInstanceTypes(v []*string) *CreateClusterBody {
	s.MasterInstanceTypes = v
	return s
}

func (s *CreateClusterBody) SetWorkerInstanceTypes(v []*string) *CreateClusterBody {
	s.WorkerInstanceTypes = v
	return s
}

func (s *CreateClusterBody) SetWorkerVswitchIds(v []*string) *CreateClusterBody {
	s.WorkerVswitchIds = v
	return s
}

type CreateClusterBodyTags struct {
	Key   *string `json:"key" xml:"key" require:"true"`
	Value *string `json:"value" xml:"value" require:"true"`
}

func (s CreateClusterBodyTags) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterBodyTags) GoString() string {
	return s.String()
}

func (s *CreateClusterBodyTags) SetKey(v string) *CreateClusterBodyTags {
	s.Key = &v
	return s
}

func (s *CreateClusterBodyTags) SetValue(v string) *CreateClusterBodyTags {
	s.Value = &v
	return s
}

type CreateClusterBodyAddons struct {
	Name   *string `json:"name" xml:"name" require:"true"`
	Config *string `json:"config" xml:"config" require:"true"`
}

func (s CreateClusterBodyAddons) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterBodyAddons) GoString() string {
	return s.String()
}

func (s *CreateClusterBodyAddons) SetName(v string) *CreateClusterBodyAddons {
	s.Name = &v
	return s
}

func (s *CreateClusterBodyAddons) SetConfig(v string) *CreateClusterBodyAddons {
	s.Config = &v
	return s
}

type CreateClusterBodyTaints struct {
	Key    *string `json:"key" xml:"key" require:"true"`
	Value  *string `json:"value" xml:"value" require:"true"`
	Effect *string `json:"effect" xml:"effect" require:"true"`
}

func (s CreateClusterBodyTaints) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterBodyTaints) GoString() string {
	return s.String()
}

func (s *CreateClusterBodyTaints) SetKey(v string) *CreateClusterBodyTaints {
	s.Key = &v
	return s
}

func (s *CreateClusterBodyTaints) SetValue(v string) *CreateClusterBodyTaints {
	s.Value = &v
	return s
}

func (s *CreateClusterBodyTaints) SetEffect(v string) *CreateClusterBodyTaints {
	s.Effect = &v
	return s
}

type CreateClusterBodyWorkerDataDisks struct {
	Category  *string `json:"category" xml:"category" require:"true"`
	Size      *string `json:"size" xml:"size" require:"true"`
	Encrypted *string `json:"encrypted" xml:"encrypted" require:"true"`
}

func (s CreateClusterBodyWorkerDataDisks) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterBodyWorkerDataDisks) GoString() string {
	return s.String()
}

func (s *CreateClusterBodyWorkerDataDisks) SetCategory(v string) *CreateClusterBodyWorkerDataDisks {
	s.Category = &v
	return s
}

func (s *CreateClusterBodyWorkerDataDisks) SetSize(v string) *CreateClusterBodyWorkerDataDisks {
	s.Size = &v
	return s
}

func (s *CreateClusterBodyWorkerDataDisks) SetEncrypted(v string) *CreateClusterBodyWorkerDataDisks {
	s.Encrypted = &v
	return s
}

type CreateClusterRequest struct {
	Headers map[string]string  `json:"headers" xml:"headers"`
	Body    *CreateClusterBody `json:"body" xml:"body" require:"true"`
}

func (s CreateClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterRequest) GoString() string {
	return s.String()
}

func (s *CreateClusterRequest) SetHeaders(v map[string]string) *CreateClusterRequest {
	s.Headers = v
	return s
}

func (s *CreateClusterRequest) SetBody(v *CreateClusterBody) *CreateClusterRequest {
	s.Body = v
	return s
}

type CreateClusterResponseBody struct {
	ClusterId *string `json:"cluster_id" xml:"cluster_id" require:"true"`
	RequestId *string `json:"request_id" xml:"request_id" require:"true"`
	TaskId    *string `json:"task_id" xml:"task_id" require:"true"`
}

func (s CreateClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterResponseBody) GoString() string {
	return s.String()
}

func (s *CreateClusterResponseBody) SetClusterId(v string) *CreateClusterResponseBody {
	s.ClusterId = &v
	return s
}

func (s *CreateClusterResponseBody) SetRequestId(v string) *CreateClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateClusterResponseBody) SetTaskId(v string) *CreateClusterResponseBody {
	s.TaskId = &v
	return s
}

type CreateClusterResponse struct {
	Headers map[string]string          `json:"headers" xml:"headers" require:"true"`
	Body    *CreateClusterResponseBody `json:"body" xml:"body" require:"true"`
}

func (s CreateClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterResponse) GoString() string {
	return s.String()
}

func (s *CreateClusterResponse) SetHeaders(v map[string]string) *CreateClusterResponse {
	s.Headers = v
	return s
}

func (s *CreateClusterResponse) SetBody(v *CreateClusterResponseBody) *CreateClusterResponse {
	s.Body = v
	return s
}

type ScaleClusterRequest struct {
	Headers map[string]string `json:"headers" xml:"headers"`
}

func (s ScaleClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s ScaleClusterRequest) GoString() string {
	return s.String()
}

func (s *ScaleClusterRequest) SetHeaders(v map[string]string) *ScaleClusterRequest {
	s.Headers = v
	return s
}

type ScaleClusterResponseBody struct {
	RequestId *string `json:"RequestId" xml:"RequestId" require:"true"`
}

func (s ScaleClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ScaleClusterResponseBody) GoString() string {
	return s.String()
}

func (s *ScaleClusterResponseBody) SetRequestId(v string) *ScaleClusterResponseBody {
	s.RequestId = &v
	return s
}

type ScaleClusterResponse struct {
	Headers map[string]string         `json:"headers" xml:"headers" require:"true"`
	Body    *ScaleClusterResponseBody `json:"body" xml:"body" require:"true"`
}

func (s ScaleClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s ScaleClusterResponse) GoString() string {
	return s.String()
}

func (s *ScaleClusterResponse) SetHeaders(v map[string]string) *ScaleClusterResponse {
	s.Headers = v
	return s
}

func (s *ScaleClusterResponse) SetBody(v *ScaleClusterResponseBody) *ScaleClusterResponse {
	s.Body = v
	return s
}

type DescribeClustersQuery struct {
	Name        *string `json:"name" xml:"name" require:"true"`
	ClusterType *string `json:"clusterType" xml:"clusterType" require:"true"`
}

func (s DescribeClustersQuery) String() string {
	return tea.Prettify(s)
}

func (s DescribeClustersQuery) GoString() string {
	return s.String()
}

func (s *DescribeClustersQuery) SetName(v string) *DescribeClustersQuery {
	s.Name = &v
	return s
}

func (s *DescribeClustersQuery) SetClusterType(v string) *DescribeClustersQuery {
	s.ClusterType = &v
	return s
}

type DescribeClustersRequest struct {
	Query   *DescribeClustersQuery `json:"query" xml:"query" require:"true"`
	Headers map[string]string      `json:"headers" xml:"headers"`
}

func (s DescribeClustersRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeClustersRequest) GoString() string {
	return s.String()
}

func (s *DescribeClustersRequest) SetQuery(v *DescribeClustersQuery) *DescribeClustersRequest {
	s.Query = v
	return s
}

func (s *DescribeClustersRequest) SetHeaders(v map[string]string) *DescribeClustersRequest {
	s.Headers = v
	return s
}

type DescribeClustersResponseBody struct {
	Clusters []*DescribeClustersResponseBodyClusters `json:"clusters" xml:"clusters" require:"true" type:"Repeated"`
}

func (s DescribeClustersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeClustersResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeClustersResponseBody) SetClusters(v []*DescribeClustersResponseBodyClusters) *DescribeClustersResponseBody {
	s.Clusters = v
	return s
}

type DescribeClustersResponseBodyClusters struct {
	Name                   *string                                                `json:"name" xml:"name" require:"true"`
	ClusterId              *string                                                `json:"cluster_id" xml:"cluster_id" require:"true"`
	Size                   *int                                                   `json:"size" xml:"size" require:"true"`
	RegionId               *string                                                `json:"region_id" xml:"region_id" require:"true"`
	State                  *string                                                `json:"state" xml:"state" require:"true"`
	ClusterType            *string                                                `json:"cluster_type" xml:"cluster_type" require:"true"`
	Created                *string                                                `json:"created" xml:"created" require:"true"`
	Updated                *string                                                `json:"updated" xml:"updated" require:"true"`
	InitVersion            *string                                                `json:"init_version" xml:"init_version" require:"true"`
	CurrentVersion         *string                                                `json:"current_version" xml:"current_version" require:"true"`
	MetaData               *string                                                `json:"meta_data" xml:"meta_data" require:"true"`
	ResourceGroupId        *string                                                `json:"resource_group_id" xml:"resource_group_id" require:"true"`
	InstanceType           *string                                                `json:"instance_type" xml:"instance_type" require:"true"`
	VpcId                  *string                                                `json:"vpc_id" xml:"vpc_id" require:"true"`
	VswitchId              *string                                                `json:"vswitch_id" xml:"vswitch_id" require:"true"`
	VswitchCidr            *string                                                `json:"vswitch_cidr" xml:"vswitch_cidr" require:"true"`
	DataDiskSize           *int                                                   `json:"data_disk_size" xml:"data_disk_size" require:"true"`
	DataDiskCategory       *string                                                `json:"data_disk_category" xml:"data_disk_category" require:"true"`
	SecurityGroupId        *string                                                `json:"security_group_id" xml:"security_group_id" require:"true"`
	ZoneId                 *string                                                `json:"zone_id" xml:"zone_id" require:"true"`
	NetworkMode            *string                                                `json:"network_mode" xml:"network_mode" require:"true"`
	SubnetCidr             *string                                                `json:"subnet_cidr" xml:"subnet_cidr" require:"true"`
	MasterUrl              *string                                                `json:"master_url" xml:"master_url" require:"true"`
	ExternalLoadbalancerId *string                                                `json:"external_loadbalancer_id" xml:"external_loadbalancer_id" require:"true"`
	Port                   *int                                                   `json:"port" xml:"port" require:"true"`
	NodeStatus             *string                                                `json:"node_status" xml:"node_status" require:"true"`
	ClusterHealthy         *string                                                `json:"cluster_healthy" xml:"cluster_healthy" require:"true"`
	DockerVersion          *string                                                `json:"docker_version" xml:"docker_version" require:"true"`
	SwarmMode              *bool                                                  `json:"swarm_mode" xml:"swarm_mode" require:"true"`
	GwBridge               *string                                                `json:"gw_bridge" xml:"gw_bridge" require:"true"`
	PrivateZone            *bool                                                  `json:"private_zone" xml:"private_zone" require:"true"`
	Profile                *string                                                `json:"profile" xml:"profile" require:"true"`
	DeletionProtection     *bool                                                  `json:"deletion_protection" xml:"deletion_protection" require:"true"`
	Capabilities           *string                                                `json:"capabilities" xml:"capabilities" require:"true"`
	EnabledMigration       *bool                                                  `json:"enabled_migration" xml:"enabled_migration" require:"true"`
	NeedUpdateAgent        *bool                                                  `json:"need_update_agent" xml:"need_update_agent" require:"true"`
	Tags                   []*DescribeClustersResponseBodyClustersTags            `json:"tags" xml:"tags" require:"true" type:"Repeated"`
	Outputs                []*DescribeClustersResponseBodyClustersOutputs         `json:"outputs" xml:"outputs" require:"true" type:"Repeated"`
	UpgradeComponents      *DescribeClustersResponseBodyClustersUpgradeComponents `json:"upgrade_components" xml:"upgrade_components" require:"true" type:"Struct"`
	Parameters             *DescribeClustersResponseBodyClustersParameters        `json:"parameters" xml:"parameters" require:"true" type:"Struct"`
}

func (s DescribeClustersResponseBodyClusters) String() string {
	return tea.Prettify(s)
}

func (s DescribeClustersResponseBodyClusters) GoString() string {
	return s.String()
}

func (s *DescribeClustersResponseBodyClusters) SetName(v string) *DescribeClustersResponseBodyClusters {
	s.Name = &v
	return s
}

func (s *DescribeClustersResponseBodyClusters) SetClusterId(v string) *DescribeClustersResponseBodyClusters {
	s.ClusterId = &v
	return s
}

func (s *DescribeClustersResponseBodyClusters) SetSize(v int) *DescribeClustersResponseBodyClusters {
	s.Size = &v
	return s
}

func (s *DescribeClustersResponseBodyClusters) SetRegionId(v string) *DescribeClustersResponseBodyClusters {
	s.RegionId = &v
	return s
}

func (s *DescribeClustersResponseBodyClusters) SetState(v string) *DescribeClustersResponseBodyClusters {
	s.State = &v
	return s
}

func (s *DescribeClustersResponseBodyClusters) SetClusterType(v string) *DescribeClustersResponseBodyClusters {
	s.ClusterType = &v
	return s
}

func (s *DescribeClustersResponseBodyClusters) SetCreated(v string) *DescribeClustersResponseBodyClusters {
	s.Created = &v
	return s
}

func (s *DescribeClustersResponseBodyClusters) SetUpdated(v string) *DescribeClustersResponseBodyClusters {
	s.Updated = &v
	return s
}

func (s *DescribeClustersResponseBodyClusters) SetInitVersion(v string) *DescribeClustersResponseBodyClusters {
	s.InitVersion = &v
	return s
}

func (s *DescribeClustersResponseBodyClusters) SetCurrentVersion(v string) *DescribeClustersResponseBodyClusters {
	s.CurrentVersion = &v
	return s
}

func (s *DescribeClustersResponseBodyClusters) SetMetaData(v string) *DescribeClustersResponseBodyClusters {
	s.MetaData = &v
	return s
}

func (s *DescribeClustersResponseBodyClusters) SetResourceGroupId(v string) *DescribeClustersResponseBodyClusters {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeClustersResponseBodyClusters) SetInstanceType(v string) *DescribeClustersResponseBodyClusters {
	s.InstanceType = &v
	return s
}

func (s *DescribeClustersResponseBodyClusters) SetVpcId(v string) *DescribeClustersResponseBodyClusters {
	s.VpcId = &v
	return s
}

func (s *DescribeClustersResponseBodyClusters) SetVswitchId(v string) *DescribeClustersResponseBodyClusters {
	s.VswitchId = &v
	return s
}

func (s *DescribeClustersResponseBodyClusters) SetVswitchCidr(v string) *DescribeClustersResponseBodyClusters {
	s.VswitchCidr = &v
	return s
}

func (s *DescribeClustersResponseBodyClusters) SetDataDiskSize(v int) *DescribeClustersResponseBodyClusters {
	s.DataDiskSize = &v
	return s
}

func (s *DescribeClustersResponseBodyClusters) SetDataDiskCategory(v string) *DescribeClustersResponseBodyClusters {
	s.DataDiskCategory = &v
	return s
}

func (s *DescribeClustersResponseBodyClusters) SetSecurityGroupId(v string) *DescribeClustersResponseBodyClusters {
	s.SecurityGroupId = &v
	return s
}

func (s *DescribeClustersResponseBodyClusters) SetZoneId(v string) *DescribeClustersResponseBodyClusters {
	s.ZoneId = &v
	return s
}

func (s *DescribeClustersResponseBodyClusters) SetNetworkMode(v string) *DescribeClustersResponseBodyClusters {
	s.NetworkMode = &v
	return s
}

func (s *DescribeClustersResponseBodyClusters) SetSubnetCidr(v string) *DescribeClustersResponseBodyClusters {
	s.SubnetCidr = &v
	return s
}

func (s *DescribeClustersResponseBodyClusters) SetMasterUrl(v string) *DescribeClustersResponseBodyClusters {
	s.MasterUrl = &v
	return s
}

func (s *DescribeClustersResponseBodyClusters) SetExternalLoadbalancerId(v string) *DescribeClustersResponseBodyClusters {
	s.ExternalLoadbalancerId = &v
	return s
}

func (s *DescribeClustersResponseBodyClusters) SetPort(v int) *DescribeClustersResponseBodyClusters {
	s.Port = &v
	return s
}

func (s *DescribeClustersResponseBodyClusters) SetNodeStatus(v string) *DescribeClustersResponseBodyClusters {
	s.NodeStatus = &v
	return s
}

func (s *DescribeClustersResponseBodyClusters) SetClusterHealthy(v string) *DescribeClustersResponseBodyClusters {
	s.ClusterHealthy = &v
	return s
}

func (s *DescribeClustersResponseBodyClusters) SetDockerVersion(v string) *DescribeClustersResponseBodyClusters {
	s.DockerVersion = &v
	return s
}

func (s *DescribeClustersResponseBodyClusters) SetSwarmMode(v bool) *DescribeClustersResponseBodyClusters {
	s.SwarmMode = &v
	return s
}

func (s *DescribeClustersResponseBodyClusters) SetGwBridge(v string) *DescribeClustersResponseBodyClusters {
	s.GwBridge = &v
	return s
}

func (s *DescribeClustersResponseBodyClusters) SetPrivateZone(v bool) *DescribeClustersResponseBodyClusters {
	s.PrivateZone = &v
	return s
}

func (s *DescribeClustersResponseBodyClusters) SetProfile(v string) *DescribeClustersResponseBodyClusters {
	s.Profile = &v
	return s
}

func (s *DescribeClustersResponseBodyClusters) SetDeletionProtection(v bool) *DescribeClustersResponseBodyClusters {
	s.DeletionProtection = &v
	return s
}

func (s *DescribeClustersResponseBodyClusters) SetCapabilities(v string) *DescribeClustersResponseBodyClusters {
	s.Capabilities = &v
	return s
}

func (s *DescribeClustersResponseBodyClusters) SetEnabledMigration(v bool) *DescribeClustersResponseBodyClusters {
	s.EnabledMigration = &v
	return s
}

func (s *DescribeClustersResponseBodyClusters) SetNeedUpdateAgent(v bool) *DescribeClustersResponseBodyClusters {
	s.NeedUpdateAgent = &v
	return s
}

func (s *DescribeClustersResponseBodyClusters) SetTags(v []*DescribeClustersResponseBodyClustersTags) *DescribeClustersResponseBodyClusters {
	s.Tags = v
	return s
}

func (s *DescribeClustersResponseBodyClusters) SetOutputs(v []*DescribeClustersResponseBodyClustersOutputs) *DescribeClustersResponseBodyClusters {
	s.Outputs = v
	return s
}

func (s *DescribeClustersResponseBodyClusters) SetUpgradeComponents(v *DescribeClustersResponseBodyClustersUpgradeComponents) *DescribeClustersResponseBodyClusters {
	s.UpgradeComponents = v
	return s
}

func (s *DescribeClustersResponseBodyClusters) SetParameters(v *DescribeClustersResponseBodyClustersParameters) *DescribeClustersResponseBodyClusters {
	s.Parameters = v
	return s
}

type DescribeClustersResponseBodyClustersTags struct {
	Key   *string `json:"key" xml:"key" require:"true"`
	Value *string `json:"value" xml:"value" require:"true"`
}

func (s DescribeClustersResponseBodyClustersTags) String() string {
	return tea.Prettify(s)
}

func (s DescribeClustersResponseBodyClustersTags) GoString() string {
	return s.String()
}

func (s *DescribeClustersResponseBodyClustersTags) SetKey(v string) *DescribeClustersResponseBodyClustersTags {
	s.Key = &v
	return s
}

func (s *DescribeClustersResponseBodyClustersTags) SetValue(v string) *DescribeClustersResponseBodyClustersTags {
	s.Value = &v
	return s
}

type DescribeClustersResponseBodyClustersOutputs struct {
	Description *string `json:"Description" xml:"Description" require:"true"`
	OutputKey   *string `json:"OutputKey" xml:"OutputKey" require:"true"`
	OutputValue *string `json:"OutputValue" xml:"OutputValue" require:"true"`
}

func (s DescribeClustersResponseBodyClustersOutputs) String() string {
	return tea.Prettify(s)
}

func (s DescribeClustersResponseBodyClustersOutputs) GoString() string {
	return s.String()
}

func (s *DescribeClustersResponseBodyClustersOutputs) SetDescription(v string) *DescribeClustersResponseBodyClustersOutputs {
	s.Description = &v
	return s
}

func (s *DescribeClustersResponseBodyClustersOutputs) SetOutputKey(v string) *DescribeClustersResponseBodyClustersOutputs {
	s.OutputKey = &v
	return s
}

func (s *DescribeClustersResponseBodyClustersOutputs) SetOutputValue(v string) *DescribeClustersResponseBodyClustersOutputs {
	s.OutputValue = &v
	return s
}

type DescribeClustersResponseBodyClustersUpgradeComponents struct {
	Kubernetes *DescribeClustersResponseBodyClustersUpgradeComponentsKubernetes `json:"Kubernetes" xml:"Kubernetes" require:"true" type:"Struct"`
}

func (s DescribeClustersResponseBodyClustersUpgradeComponents) String() string {
	return tea.Prettify(s)
}

func (s DescribeClustersResponseBodyClustersUpgradeComponents) GoString() string {
	return s.String()
}

func (s *DescribeClustersResponseBodyClustersUpgradeComponents) SetKubernetes(v *DescribeClustersResponseBodyClustersUpgradeComponentsKubernetes) *DescribeClustersResponseBodyClustersUpgradeComponents {
	s.Kubernetes = v
	return s
}

type DescribeClustersResponseBodyClustersUpgradeComponentsKubernetes struct {
	ComponentName  *string `json:"component_name" xml:"component_name" require:"true"`
	Version        *string `json:"version" xml:"version" require:"true"`
	NextVersion    *string `json:"next_version" xml:"next_version" require:"true"`
	Changed        *string `json:"changed" xml:"changed" require:"true"`
	CanUpgrade     *bool   `json:"can_upgrade" xml:"can_upgrade" require:"true"`
	Force          *bool   `json:"force" xml:"force" require:"true"`
	Policy         *string `json:"policy" xml:"policy" require:"true"`
	ExtraVars      *string `json:"ExtraVars" xml:"ExtraVars" require:"true"`
	ReadyToUpgrade *string `json:"ready_to_upgrade" xml:"ready_to_upgrade" require:"true"`
	Message        *string `json:"message" xml:"message" require:"true"`
	Exist          *bool   `json:"exist" xml:"exist" require:"true"`
	Required       *bool   `json:"required" xml:"required" require:"true"`
	Template       *string `json:"template" xml:"template" require:"true"`
}

func (s DescribeClustersResponseBodyClustersUpgradeComponentsKubernetes) String() string {
	return tea.Prettify(s)
}

func (s DescribeClustersResponseBodyClustersUpgradeComponentsKubernetes) GoString() string {
	return s.String()
}

func (s *DescribeClustersResponseBodyClustersUpgradeComponentsKubernetes) SetComponentName(v string) *DescribeClustersResponseBodyClustersUpgradeComponentsKubernetes {
	s.ComponentName = &v
	return s
}

func (s *DescribeClustersResponseBodyClustersUpgradeComponentsKubernetes) SetVersion(v string) *DescribeClustersResponseBodyClustersUpgradeComponentsKubernetes {
	s.Version = &v
	return s
}

func (s *DescribeClustersResponseBodyClustersUpgradeComponentsKubernetes) SetNextVersion(v string) *DescribeClustersResponseBodyClustersUpgradeComponentsKubernetes {
	s.NextVersion = &v
	return s
}

func (s *DescribeClustersResponseBodyClustersUpgradeComponentsKubernetes) SetChanged(v string) *DescribeClustersResponseBodyClustersUpgradeComponentsKubernetes {
	s.Changed = &v
	return s
}

func (s *DescribeClustersResponseBodyClustersUpgradeComponentsKubernetes) SetCanUpgrade(v bool) *DescribeClustersResponseBodyClustersUpgradeComponentsKubernetes {
	s.CanUpgrade = &v
	return s
}

func (s *DescribeClustersResponseBodyClustersUpgradeComponentsKubernetes) SetForce(v bool) *DescribeClustersResponseBodyClustersUpgradeComponentsKubernetes {
	s.Force = &v
	return s
}

func (s *DescribeClustersResponseBodyClustersUpgradeComponentsKubernetes) SetPolicy(v string) *DescribeClustersResponseBodyClustersUpgradeComponentsKubernetes {
	s.Policy = &v
	return s
}

func (s *DescribeClustersResponseBodyClustersUpgradeComponentsKubernetes) SetExtraVars(v string) *DescribeClustersResponseBodyClustersUpgradeComponentsKubernetes {
	s.ExtraVars = &v
	return s
}

func (s *DescribeClustersResponseBodyClustersUpgradeComponentsKubernetes) SetReadyToUpgrade(v string) *DescribeClustersResponseBodyClustersUpgradeComponentsKubernetes {
	s.ReadyToUpgrade = &v
	return s
}

func (s *DescribeClustersResponseBodyClustersUpgradeComponentsKubernetes) SetMessage(v string) *DescribeClustersResponseBodyClustersUpgradeComponentsKubernetes {
	s.Message = &v
	return s
}

func (s *DescribeClustersResponseBodyClustersUpgradeComponentsKubernetes) SetExist(v bool) *DescribeClustersResponseBodyClustersUpgradeComponentsKubernetes {
	s.Exist = &v
	return s
}

func (s *DescribeClustersResponseBodyClustersUpgradeComponentsKubernetes) SetRequired(v bool) *DescribeClustersResponseBodyClustersUpgradeComponentsKubernetes {
	s.Required = &v
	return s
}

func (s *DescribeClustersResponseBodyClustersUpgradeComponentsKubernetes) SetTemplate(v string) *DescribeClustersResponseBodyClustersUpgradeComponentsKubernetes {
	s.Template = &v
	return s
}

type DescribeClustersResponseBodyClustersParameters struct {
	AdjustmentType           *string `json:"AdjustmentType" xml:"AdjustmentType" require:"true"`
	AuditFlags               *string `json:"AuditFlags" xml:"AuditFlags" require:"true"`
	BetaVersion              *string `json:"BetaVersion" xml:"BetaVersion" require:"true"`
	CA                       *string `json:"CA" xml:"CA" require:"true"`
	ClientCA                 *string `json:"ClientCA" xml:"ClientCA" require:"true"`
	CloudMonitorFlags        *string `json:"CloudMonitorFlags" xml:"CloudMonitorFlags" require:"true"`
	CloudMonitorVersion      *string `json:"CloudMonitorVersion" xml:"CloudMonitorVersion" require:"true"`
	ContainerCIDR            *string `json:"ContainerCIDR" xml:"ContainerCIDR" require:"true"`
	DisableAddons            *string `json:"DisableAddons" xml:"DisableAddons" require:"true"`
	DockerVersion            *string `json:"DockerVersion" xml:"DockerVersion" require:"true"`
	ESSDeletionProtection    *string `json:"ESSDeletionProtection" xml:"ESSDeletionProtection" require:"true"`
	Eip                      *string `json:"Eip" xml:"Eip" require:"true"`
	EipAddress               *string `json:"EipAddress" xml:"EipAddress" require:"true"`
	ElasticSearchHost        *string `json:"ElasticSearchHost" xml:"ElasticSearchHost" require:"true"`
	ElasticSearchPass        *string `json:"ElasticSearchPass" xml:"ElasticSearchPass" require:"true"`
	ElasticSearchPort        *string `json:"ElasticSearchPort" xml:"ElasticSearchPort" require:"true"`
	ElasticSearchUser        *string `json:"ElasticSearchUser" xml:"ElasticSearchUser" require:"true"`
	EtcdVersion              *string `json:"EtcdVersion" xml:"EtcdVersion" require:"true"`
	ExecuteVersion           *string `json:"ExecuteVersion" xml:"ExecuteVersion" require:"true"`
	GPUFlags                 *string `json:"GPUFlags" xml:"GPUFlags" require:"true"`
	HealthCheckType          *string `json:"HealthCheckType" xml:"HealthCheckType" require:"true"`
	ImageId                  *string `json:"ImageId" xml:"ImageId" require:"true"`
	K8SMasterPolicyDocument  *string `json:"K8SMasterPolicyDocument" xml:"K8SMasterPolicyDocument" require:"true"`
	K8sWorkerPolicyDocument  *string `json:"K8sWorkerPolicyDocument" xml:"K8sWorkerPolicyDocument" require:"true"`
	Key                      *string `json:"Key" xml:"Key" require:"true"`
	KeyPair                  *string `json:"KeyPair" xml:"KeyPair" require:"true"`
	KubernetesVersion        *string `json:"KubernetesVersion" xml:"KubernetesVersion" require:"true"`
	LoggingType              *string `json:"LoggingType" xml:"LoggingType" require:"true"`
	MasterAmounts            *string `json:"MasterAmounts" xml:"MasterAmounts" require:"true"`
	MasterAutoRenew          *string `json:"MasterAutoRenew" xml:"MasterAutoRenew" require:"true"`
	MasterAutoRenewPeriod    *string `json:"MasterAutoRenewPeriod" xml:"MasterAutoRenewPeriod" require:"true"`
	MasterCount              *string `json:"MasterCount" xml:"MasterCount" require:"true"`
	MasterDataDisk           *string `json:"MasterDataDisk" xml:"MasterDataDisk" require:"true"`
	MasterDataDisks          *string `json:"MasterDataDisks" xml:"MasterDataDisks" require:"true"`
	MasterDeletionProtection *string `json:"MasterDeletionProtection" xml:"MasterDeletionProtection" require:"true"`
	MasterDeploymentSetId    *string `json:"MasterDeploymentSetId" xml:"MasterDeploymentSetId" require:"true"`
	MasterHpcClusterId       *string `json:"MasterHpcClusterId" xml:"MasterHpcClusterId" require:"true"`
	MasterImageId            *string `json:"MasterImageId" xml:"MasterImageId" require:"true"`
	MasterInstanceChargeType *string `json:"MasterInstanceChargeType" xml:"MasterInstanceChargeType" require:"true"`
	MasterInstanceTypes      *string `json:"MasterInstanceTypes" xml:"MasterInstanceTypes" require:"true"`
	MasterKeyPair            *string `json:"MasterKeyPair" xml:"MasterKeyPair" require:"true"`
	MasterLoginPassword      *string `json:"MasterLoginPassword" xml:"MasterLoginPassword" require:"true"`
	MasterPeriod             *string `json:"MasterPeriod" xml:"MasterPeriod" require:"true"`
	MasterPeriodUnit         *string `json:"MasterPeriodUnit" xml:"MasterPeriodUnit" require:"true"`
	MasterSlbSShHealthCheck  *string `json:"MasterSlbSShHealthCheck" xml:"MasterSlbSShHealthCheck" require:"true"`
	MasterSnapshotPolicyId   *string `json:"MasterSnapshotPolicyId" xml:"MasterSnapshotPolicyId" require:"true"`
	MasterSystemDiskCategory *string `json:"MasterSystemDiskCategory" xml:"MasterSystemDiskCategory" require:"true"`
	MasterSystemDiskSize     *string `json:"MasterSystemDiskSize" xml:"MasterSystemDiskSize" require:"true"`
	MasterVSwitchIds         *string `json:"MasterVSwitchIds" xml:"MasterVSwitchIds" require:"true"`
	NatGateway               *string `json:"NatGateway" xml:"NatGateway" require:"true"`
	NatGatewayId             *string `json:"NatGatewayId" xml:"NatGatewayId" require:"true"`
	Network                  *string `json:"Network" xml:"Network" require:"true"`
	NodeCIDRMask             *string `json:"NodeCIDRMask" xml:"NodeCIDRMask" require:"true"`
	NodeNameMode             *string `json:"NodeNameMode" xml:"NodeNameMode" require:"true"`
	NumOfNodes               *string `json:"NumOfNodes" xml:"NumOfNodes" require:"true"`
	Password                 *string `json:"Password" xml:"Password" require:"true"`
	PodVswitchIds            *string `json:"PodVswitchIds" xml:"PodVswitchIds" require:"true"`
	ProtectedInstances       *string `json:"ProtectedInstances" xml:"ProtectedInstances" require:"true"`
	ProxyMode                *string `json:"ProxyMode" xml:"ProxyMode" require:"true"`
	PublicSLB                *string `json:"PublicSLB" xml:"PublicSLB" require:"true"`
	RemoveInstanceIds        *string `json:"RemoveInstanceIds" xml:"RemoveInstanceIds" require:"true"`
	ResourceGroupId          *string `json:"ResourceGroupId" xml:"ResourceGroupId" require:"true"`
	SLBDeletionProtection    *string `json:"SLBDeletionProtection" xml:"SLBDeletionProtection" require:"true"`
	SLSProjectName           *string `json:"SLSProjectName" xml:"SLSProjectName" require:"true"`
	SNatEntry                *string `json:"SNatEntry" xml:"SNatEntry" require:"true"`
	SSHFlags                 *string `json:"SSHFlags" xml:"SSHFlags" require:"true"`
	SecurityGroupId          *string `json:"SecurityGroupId" xml:"SecurityGroupId" require:"true"`
	ServiceCIDR              *string `json:"ServiceCIDR" xml:"ServiceCIDR" require:"true"`
	SetUpArgs                *string `json:"SetUpArgs" xml:"SetUpArgs" require:"true"`
	SnatTableId              *string `json:"SnatTableId" xml:"SnatTableId" require:"true"`
	Tags                     *string `json:"Tags" xml:"Tags" require:"true"`
	UserCA                   *string `json:"UserCA" xml:"UserCA" require:"true"`
	UserData                 *string `json:"UserData" xml:"UserData" require:"true"`
	VpcId                    *string `json:"VpcId" xml:"VpcId" require:"true"`
	WillReplace              *string `json:"WillReplace" xml:"WillReplace" require:"true"`
	WorkerAutoRenew          *string `json:"WorkerAutoRenew" xml:"WorkerAutoRenew" require:"true"`
	WorkerAutoRenewPeriod    *string `json:"WorkerAutoRenewPeriod" xml:"WorkerAutoRenewPeriod" require:"true"`
	WorkerDataDisk           *string `json:"WorkerDataDisk" xml:"WorkerDataDisk" require:"true"`
	WorkerDataDisks          *string `json:"WorkerDataDisks" xml:"WorkerDataDisks" require:"true"`
	WorkerDeletionProtection *string `json:"WorkerDeletionProtection" xml:"WorkerDeletionProtection" require:"true"`
	WorkerDeploymentSetId    *string `json:"WorkerDeploymentSetId" xml:"WorkerDeploymentSetId" require:"true"`
	WorkerHpcClusterId       *string `json:"WorkerHpcClusterId" xml:"WorkerHpcClusterId" require:"true"`
	WorkerImageId            *string `json:"WorkerImageId" xml:"WorkerImageId" require:"true"`
	WorkerInstanceChargeType *string `json:"WorkerInstanceChargeType" xml:"WorkerInstanceChargeType" require:"true"`
	WorkerInstanceTypes      *string `json:"WorkerInstanceTypes" xml:"WorkerInstanceTypes" require:"true"`
	WorkerKeyPair            *string `json:"WorkerKeyPair" xml:"WorkerKeyPair" require:"true"`
	WorkerLoginPassword      *string `json:"WorkerLoginPassword" xml:"WorkerLoginPassword" require:"true"`
	WorkerPeriod             *string `json:"WorkerPeriod" xml:"WorkerPeriod" require:"true"`
	WorkerPeriodUnit         *string `json:"WorkerPeriodUnit" xml:"WorkerPeriodUnit" require:"true"`
	WorkerSnapshotPolicyId   *string `json:"WorkerSnapshotPolicyId" xml:"WorkerSnapshotPolicyId" require:"true"`
	WorkerSystemDiskCategory *string `json:"WorkerSystemDiskCategory" xml:"WorkerSystemDiskCategory" require:"true"`
	WorkerSystemDiskSize     *string `json:"WorkerSystemDiskSize" xml:"WorkerSystemDiskSize" require:"true"`
	WorkerVSwitchIds         *string `json:"WorkerVSwitchIds" xml:"WorkerVSwitchIds" require:"true"`
	ZoneId                   *string `json:"ZoneId" xml:"ZoneId" require:"true"`
}

func (s DescribeClustersResponseBodyClustersParameters) String() string {
	return tea.Prettify(s)
}

func (s DescribeClustersResponseBodyClustersParameters) GoString() string {
	return s.String()
}

func (s *DescribeClustersResponseBodyClustersParameters) SetAdjustmentType(v string) *DescribeClustersResponseBodyClustersParameters {
	s.AdjustmentType = &v
	return s
}

func (s *DescribeClustersResponseBodyClustersParameters) SetAuditFlags(v string) *DescribeClustersResponseBodyClustersParameters {
	s.AuditFlags = &v
	return s
}

func (s *DescribeClustersResponseBodyClustersParameters) SetBetaVersion(v string) *DescribeClustersResponseBodyClustersParameters {
	s.BetaVersion = &v
	return s
}

func (s *DescribeClustersResponseBodyClustersParameters) SetCA(v string) *DescribeClustersResponseBodyClustersParameters {
	s.CA = &v
	return s
}

func (s *DescribeClustersResponseBodyClustersParameters) SetClientCA(v string) *DescribeClustersResponseBodyClustersParameters {
	s.ClientCA = &v
	return s
}

func (s *DescribeClustersResponseBodyClustersParameters) SetCloudMonitorFlags(v string) *DescribeClustersResponseBodyClustersParameters {
	s.CloudMonitorFlags = &v
	return s
}

func (s *DescribeClustersResponseBodyClustersParameters) SetCloudMonitorVersion(v string) *DescribeClustersResponseBodyClustersParameters {
	s.CloudMonitorVersion = &v
	return s
}

func (s *DescribeClustersResponseBodyClustersParameters) SetContainerCIDR(v string) *DescribeClustersResponseBodyClustersParameters {
	s.ContainerCIDR = &v
	return s
}

func (s *DescribeClustersResponseBodyClustersParameters) SetDisableAddons(v string) *DescribeClustersResponseBodyClustersParameters {
	s.DisableAddons = &v
	return s
}

func (s *DescribeClustersResponseBodyClustersParameters) SetDockerVersion(v string) *DescribeClustersResponseBodyClustersParameters {
	s.DockerVersion = &v
	return s
}

func (s *DescribeClustersResponseBodyClustersParameters) SetESSDeletionProtection(v string) *DescribeClustersResponseBodyClustersParameters {
	s.ESSDeletionProtection = &v
	return s
}

func (s *DescribeClustersResponseBodyClustersParameters) SetEip(v string) *DescribeClustersResponseBodyClustersParameters {
	s.Eip = &v
	return s
}

func (s *DescribeClustersResponseBodyClustersParameters) SetEipAddress(v string) *DescribeClustersResponseBodyClustersParameters {
	s.EipAddress = &v
	return s
}

func (s *DescribeClustersResponseBodyClustersParameters) SetElasticSearchHost(v string) *DescribeClustersResponseBodyClustersParameters {
	s.ElasticSearchHost = &v
	return s
}

func (s *DescribeClustersResponseBodyClustersParameters) SetElasticSearchPass(v string) *DescribeClustersResponseBodyClustersParameters {
	s.ElasticSearchPass = &v
	return s
}

func (s *DescribeClustersResponseBodyClustersParameters) SetElasticSearchPort(v string) *DescribeClustersResponseBodyClustersParameters {
	s.ElasticSearchPort = &v
	return s
}

func (s *DescribeClustersResponseBodyClustersParameters) SetElasticSearchUser(v string) *DescribeClustersResponseBodyClustersParameters {
	s.ElasticSearchUser = &v
	return s
}

func (s *DescribeClustersResponseBodyClustersParameters) SetEtcdVersion(v string) *DescribeClustersResponseBodyClustersParameters {
	s.EtcdVersion = &v
	return s
}

func (s *DescribeClustersResponseBodyClustersParameters) SetExecuteVersion(v string) *DescribeClustersResponseBodyClustersParameters {
	s.ExecuteVersion = &v
	return s
}

func (s *DescribeClustersResponseBodyClustersParameters) SetGPUFlags(v string) *DescribeClustersResponseBodyClustersParameters {
	s.GPUFlags = &v
	return s
}

func (s *DescribeClustersResponseBodyClustersParameters) SetHealthCheckType(v string) *DescribeClustersResponseBodyClustersParameters {
	s.HealthCheckType = &v
	return s
}

func (s *DescribeClustersResponseBodyClustersParameters) SetImageId(v string) *DescribeClustersResponseBodyClustersParameters {
	s.ImageId = &v
	return s
}

func (s *DescribeClustersResponseBodyClustersParameters) SetK8SMasterPolicyDocument(v string) *DescribeClustersResponseBodyClustersParameters {
	s.K8SMasterPolicyDocument = &v
	return s
}

func (s *DescribeClustersResponseBodyClustersParameters) SetK8sWorkerPolicyDocument(v string) *DescribeClustersResponseBodyClustersParameters {
	s.K8sWorkerPolicyDocument = &v
	return s
}

func (s *DescribeClustersResponseBodyClustersParameters) SetKey(v string) *DescribeClustersResponseBodyClustersParameters {
	s.Key = &v
	return s
}

func (s *DescribeClustersResponseBodyClustersParameters) SetKeyPair(v string) *DescribeClustersResponseBodyClustersParameters {
	s.KeyPair = &v
	return s
}

func (s *DescribeClustersResponseBodyClustersParameters) SetKubernetesVersion(v string) *DescribeClustersResponseBodyClustersParameters {
	s.KubernetesVersion = &v
	return s
}

func (s *DescribeClustersResponseBodyClustersParameters) SetLoggingType(v string) *DescribeClustersResponseBodyClustersParameters {
	s.LoggingType = &v
	return s
}

func (s *DescribeClustersResponseBodyClustersParameters) SetMasterAmounts(v string) *DescribeClustersResponseBodyClustersParameters {
	s.MasterAmounts = &v
	return s
}

func (s *DescribeClustersResponseBodyClustersParameters) SetMasterAutoRenew(v string) *DescribeClustersResponseBodyClustersParameters {
	s.MasterAutoRenew = &v
	return s
}

func (s *DescribeClustersResponseBodyClustersParameters) SetMasterAutoRenewPeriod(v string) *DescribeClustersResponseBodyClustersParameters {
	s.MasterAutoRenewPeriod = &v
	return s
}

func (s *DescribeClustersResponseBodyClustersParameters) SetMasterCount(v string) *DescribeClustersResponseBodyClustersParameters {
	s.MasterCount = &v
	return s
}

func (s *DescribeClustersResponseBodyClustersParameters) SetMasterDataDisk(v string) *DescribeClustersResponseBodyClustersParameters {
	s.MasterDataDisk = &v
	return s
}

func (s *DescribeClustersResponseBodyClustersParameters) SetMasterDataDisks(v string) *DescribeClustersResponseBodyClustersParameters {
	s.MasterDataDisks = &v
	return s
}

func (s *DescribeClustersResponseBodyClustersParameters) SetMasterDeletionProtection(v string) *DescribeClustersResponseBodyClustersParameters {
	s.MasterDeletionProtection = &v
	return s
}

func (s *DescribeClustersResponseBodyClustersParameters) SetMasterDeploymentSetId(v string) *DescribeClustersResponseBodyClustersParameters {
	s.MasterDeploymentSetId = &v
	return s
}

func (s *DescribeClustersResponseBodyClustersParameters) SetMasterHpcClusterId(v string) *DescribeClustersResponseBodyClustersParameters {
	s.MasterHpcClusterId = &v
	return s
}

func (s *DescribeClustersResponseBodyClustersParameters) SetMasterImageId(v string) *DescribeClustersResponseBodyClustersParameters {
	s.MasterImageId = &v
	return s
}

func (s *DescribeClustersResponseBodyClustersParameters) SetMasterInstanceChargeType(v string) *DescribeClustersResponseBodyClustersParameters {
	s.MasterInstanceChargeType = &v
	return s
}

func (s *DescribeClustersResponseBodyClustersParameters) SetMasterInstanceTypes(v string) *DescribeClustersResponseBodyClustersParameters {
	s.MasterInstanceTypes = &v
	return s
}

func (s *DescribeClustersResponseBodyClustersParameters) SetMasterKeyPair(v string) *DescribeClustersResponseBodyClustersParameters {
	s.MasterKeyPair = &v
	return s
}

func (s *DescribeClustersResponseBodyClustersParameters) SetMasterLoginPassword(v string) *DescribeClustersResponseBodyClustersParameters {
	s.MasterLoginPassword = &v
	return s
}

func (s *DescribeClustersResponseBodyClustersParameters) SetMasterPeriod(v string) *DescribeClustersResponseBodyClustersParameters {
	s.MasterPeriod = &v
	return s
}

func (s *DescribeClustersResponseBodyClustersParameters) SetMasterPeriodUnit(v string) *DescribeClustersResponseBodyClustersParameters {
	s.MasterPeriodUnit = &v
	return s
}

func (s *DescribeClustersResponseBodyClustersParameters) SetMasterSlbSShHealthCheck(v string) *DescribeClustersResponseBodyClustersParameters {
	s.MasterSlbSShHealthCheck = &v
	return s
}

func (s *DescribeClustersResponseBodyClustersParameters) SetMasterSnapshotPolicyId(v string) *DescribeClustersResponseBodyClustersParameters {
	s.MasterSnapshotPolicyId = &v
	return s
}

func (s *DescribeClustersResponseBodyClustersParameters) SetMasterSystemDiskCategory(v string) *DescribeClustersResponseBodyClustersParameters {
	s.MasterSystemDiskCategory = &v
	return s
}

func (s *DescribeClustersResponseBodyClustersParameters) SetMasterSystemDiskSize(v string) *DescribeClustersResponseBodyClustersParameters {
	s.MasterSystemDiskSize = &v
	return s
}

func (s *DescribeClustersResponseBodyClustersParameters) SetMasterVSwitchIds(v string) *DescribeClustersResponseBodyClustersParameters {
	s.MasterVSwitchIds = &v
	return s
}

func (s *DescribeClustersResponseBodyClustersParameters) SetNatGateway(v string) *DescribeClustersResponseBodyClustersParameters {
	s.NatGateway = &v
	return s
}

func (s *DescribeClustersResponseBodyClustersParameters) SetNatGatewayId(v string) *DescribeClustersResponseBodyClustersParameters {
	s.NatGatewayId = &v
	return s
}

func (s *DescribeClustersResponseBodyClustersParameters) SetNetwork(v string) *DescribeClustersResponseBodyClustersParameters {
	s.Network = &v
	return s
}

func (s *DescribeClustersResponseBodyClustersParameters) SetNodeCIDRMask(v string) *DescribeClustersResponseBodyClustersParameters {
	s.NodeCIDRMask = &v
	return s
}

func (s *DescribeClustersResponseBodyClustersParameters) SetNodeNameMode(v string) *DescribeClustersResponseBodyClustersParameters {
	s.NodeNameMode = &v
	return s
}

func (s *DescribeClustersResponseBodyClustersParameters) SetNumOfNodes(v string) *DescribeClustersResponseBodyClustersParameters {
	s.NumOfNodes = &v
	return s
}

func (s *DescribeClustersResponseBodyClustersParameters) SetPassword(v string) *DescribeClustersResponseBodyClustersParameters {
	s.Password = &v
	return s
}

func (s *DescribeClustersResponseBodyClustersParameters) SetPodVswitchIds(v string) *DescribeClustersResponseBodyClustersParameters {
	s.PodVswitchIds = &v
	return s
}

func (s *DescribeClustersResponseBodyClustersParameters) SetProtectedInstances(v string) *DescribeClustersResponseBodyClustersParameters {
	s.ProtectedInstances = &v
	return s
}

func (s *DescribeClustersResponseBodyClustersParameters) SetProxyMode(v string) *DescribeClustersResponseBodyClustersParameters {
	s.ProxyMode = &v
	return s
}

func (s *DescribeClustersResponseBodyClustersParameters) SetPublicSLB(v string) *DescribeClustersResponseBodyClustersParameters {
	s.PublicSLB = &v
	return s
}

func (s *DescribeClustersResponseBodyClustersParameters) SetRemoveInstanceIds(v string) *DescribeClustersResponseBodyClustersParameters {
	s.RemoveInstanceIds = &v
	return s
}

func (s *DescribeClustersResponseBodyClustersParameters) SetResourceGroupId(v string) *DescribeClustersResponseBodyClustersParameters {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeClustersResponseBodyClustersParameters) SetSLBDeletionProtection(v string) *DescribeClustersResponseBodyClustersParameters {
	s.SLBDeletionProtection = &v
	return s
}

func (s *DescribeClustersResponseBodyClustersParameters) SetSLSProjectName(v string) *DescribeClustersResponseBodyClustersParameters {
	s.SLSProjectName = &v
	return s
}

func (s *DescribeClustersResponseBodyClustersParameters) SetSNatEntry(v string) *DescribeClustersResponseBodyClustersParameters {
	s.SNatEntry = &v
	return s
}

func (s *DescribeClustersResponseBodyClustersParameters) SetSSHFlags(v string) *DescribeClustersResponseBodyClustersParameters {
	s.SSHFlags = &v
	return s
}

func (s *DescribeClustersResponseBodyClustersParameters) SetSecurityGroupId(v string) *DescribeClustersResponseBodyClustersParameters {
	s.SecurityGroupId = &v
	return s
}

func (s *DescribeClustersResponseBodyClustersParameters) SetServiceCIDR(v string) *DescribeClustersResponseBodyClustersParameters {
	s.ServiceCIDR = &v
	return s
}

func (s *DescribeClustersResponseBodyClustersParameters) SetSetUpArgs(v string) *DescribeClustersResponseBodyClustersParameters {
	s.SetUpArgs = &v
	return s
}

func (s *DescribeClustersResponseBodyClustersParameters) SetSnatTableId(v string) *DescribeClustersResponseBodyClustersParameters {
	s.SnatTableId = &v
	return s
}

func (s *DescribeClustersResponseBodyClustersParameters) SetTags(v string) *DescribeClustersResponseBodyClustersParameters {
	s.Tags = &v
	return s
}

func (s *DescribeClustersResponseBodyClustersParameters) SetUserCA(v string) *DescribeClustersResponseBodyClustersParameters {
	s.UserCA = &v
	return s
}

func (s *DescribeClustersResponseBodyClustersParameters) SetUserData(v string) *DescribeClustersResponseBodyClustersParameters {
	s.UserData = &v
	return s
}

func (s *DescribeClustersResponseBodyClustersParameters) SetVpcId(v string) *DescribeClustersResponseBodyClustersParameters {
	s.VpcId = &v
	return s
}

func (s *DescribeClustersResponseBodyClustersParameters) SetWillReplace(v string) *DescribeClustersResponseBodyClustersParameters {
	s.WillReplace = &v
	return s
}

func (s *DescribeClustersResponseBodyClustersParameters) SetWorkerAutoRenew(v string) *DescribeClustersResponseBodyClustersParameters {
	s.WorkerAutoRenew = &v
	return s
}

func (s *DescribeClustersResponseBodyClustersParameters) SetWorkerAutoRenewPeriod(v string) *DescribeClustersResponseBodyClustersParameters {
	s.WorkerAutoRenewPeriod = &v
	return s
}

func (s *DescribeClustersResponseBodyClustersParameters) SetWorkerDataDisk(v string) *DescribeClustersResponseBodyClustersParameters {
	s.WorkerDataDisk = &v
	return s
}

func (s *DescribeClustersResponseBodyClustersParameters) SetWorkerDataDisks(v string) *DescribeClustersResponseBodyClustersParameters {
	s.WorkerDataDisks = &v
	return s
}

func (s *DescribeClustersResponseBodyClustersParameters) SetWorkerDeletionProtection(v string) *DescribeClustersResponseBodyClustersParameters {
	s.WorkerDeletionProtection = &v
	return s
}

func (s *DescribeClustersResponseBodyClustersParameters) SetWorkerDeploymentSetId(v string) *DescribeClustersResponseBodyClustersParameters {
	s.WorkerDeploymentSetId = &v
	return s
}

func (s *DescribeClustersResponseBodyClustersParameters) SetWorkerHpcClusterId(v string) *DescribeClustersResponseBodyClustersParameters {
	s.WorkerHpcClusterId = &v
	return s
}

func (s *DescribeClustersResponseBodyClustersParameters) SetWorkerImageId(v string) *DescribeClustersResponseBodyClustersParameters {
	s.WorkerImageId = &v
	return s
}

func (s *DescribeClustersResponseBodyClustersParameters) SetWorkerInstanceChargeType(v string) *DescribeClustersResponseBodyClustersParameters {
	s.WorkerInstanceChargeType = &v
	return s
}

func (s *DescribeClustersResponseBodyClustersParameters) SetWorkerInstanceTypes(v string) *DescribeClustersResponseBodyClustersParameters {
	s.WorkerInstanceTypes = &v
	return s
}

func (s *DescribeClustersResponseBodyClustersParameters) SetWorkerKeyPair(v string) *DescribeClustersResponseBodyClustersParameters {
	s.WorkerKeyPair = &v
	return s
}

func (s *DescribeClustersResponseBodyClustersParameters) SetWorkerLoginPassword(v string) *DescribeClustersResponseBodyClustersParameters {
	s.WorkerLoginPassword = &v
	return s
}

func (s *DescribeClustersResponseBodyClustersParameters) SetWorkerPeriod(v string) *DescribeClustersResponseBodyClustersParameters {
	s.WorkerPeriod = &v
	return s
}

func (s *DescribeClustersResponseBodyClustersParameters) SetWorkerPeriodUnit(v string) *DescribeClustersResponseBodyClustersParameters {
	s.WorkerPeriodUnit = &v
	return s
}

func (s *DescribeClustersResponseBodyClustersParameters) SetWorkerSnapshotPolicyId(v string) *DescribeClustersResponseBodyClustersParameters {
	s.WorkerSnapshotPolicyId = &v
	return s
}

func (s *DescribeClustersResponseBodyClustersParameters) SetWorkerSystemDiskCategory(v string) *DescribeClustersResponseBodyClustersParameters {
	s.WorkerSystemDiskCategory = &v
	return s
}

func (s *DescribeClustersResponseBodyClustersParameters) SetWorkerSystemDiskSize(v string) *DescribeClustersResponseBodyClustersParameters {
	s.WorkerSystemDiskSize = &v
	return s
}

func (s *DescribeClustersResponseBodyClustersParameters) SetWorkerVSwitchIds(v string) *DescribeClustersResponseBodyClustersParameters {
	s.WorkerVSwitchIds = &v
	return s
}

func (s *DescribeClustersResponseBodyClustersParameters) SetZoneId(v string) *DescribeClustersResponseBodyClustersParameters {
	s.ZoneId = &v
	return s
}

type DescribeClustersResponse struct {
	Headers map[string]string             `json:"headers" xml:"headers" require:"true"`
	Body    *DescribeClustersResponseBody `json:"body" xml:"body" require:"true"`
}

func (s DescribeClustersResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeClustersResponse) GoString() string {
	return s.String()
}

func (s *DescribeClustersResponse) SetHeaders(v map[string]string) *DescribeClustersResponse {
	s.Headers = v
	return s
}

func (s *DescribeClustersResponse) SetBody(v *DescribeClustersResponseBody) *DescribeClustersResponse {
	s.Body = v
	return s
}

type DescribeClusterDetailRequest struct {
	Headers map[string]string `json:"headers" xml:"headers"`
}

func (s DescribeClusterDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeClusterDetailRequest) SetHeaders(v map[string]string) *DescribeClusterDetailRequest {
	s.Headers = v
	return s
}

type DescribeClusterDetailResponseBody struct {
	Name                   *string                                             `json:"name" xml:"name" require:"true"`
	ClusterId              *string                                             `json:"cluster_id" xml:"cluster_id" require:"true"`
	Size                   *int                                                `json:"size" xml:"size" require:"true"`
	RegionId               *string                                             `json:"region_id" xml:"region_id" require:"true"`
	State                  *string                                             `json:"state" xml:"state" require:"true"`
	ClusterType            *string                                             `json:"cluster_type" xml:"cluster_type" require:"true"`
	Created                *string                                             `json:"created" xml:"created" require:"true"`
	Updated                *string                                             `json:"updated" xml:"updated" require:"true"`
	InitVersion            *string                                             `json:"init_version" xml:"init_version" require:"true"`
	CurrentVersion         *string                                             `json:"current_version" xml:"current_version" require:"true"`
	MetaData               *string                                             `json:"meta_data" xml:"meta_data" require:"true"`
	ResourceGroupId        *string                                             `json:"resource_group_id" xml:"resource_group_id" require:"true"`
	InstanceType           *string                                             `json:"instance_type" xml:"instance_type" require:"true"`
	VpcId                  *string                                             `json:"vpc_id" xml:"vpc_id" require:"true"`
	VswitchId              *string                                             `json:"vswitch_id" xml:"vswitch_id" require:"true"`
	VswitchCidr            *string                                             `json:"vswitch_cidr" xml:"vswitch_cidr" require:"true"`
	DataDiskSize           *int                                                `json:"data_disk_size" xml:"data_disk_size" require:"true"`
	DataDiskCategory       *string                                             `json:"data_disk_category" xml:"data_disk_category" require:"true"`
	SecurityGroupId        *string                                             `json:"security_group_id" xml:"security_group_id" require:"true"`
	ZoneId                 *string                                             `json:"zone_id" xml:"zone_id" require:"true"`
	NetworkMode            *string                                             `json:"network_mode" xml:"network_mode" require:"true"`
	SubnetCidr             *string                                             `json:"subnet_cidr" xml:"subnet_cidr" require:"true"`
	MasterUrl              *string                                             `json:"master_url" xml:"master_url" require:"true"`
	ExternalLoadbalancerId *string                                             `json:"external_loadbalancer_id" xml:"external_loadbalancer_id" require:"true"`
	Port                   *int                                                `json:"port" xml:"port" require:"true"`
	NodeStatus             *string                                             `json:"node_status" xml:"node_status" require:"true"`
	ClusterHealthy         *string                                             `json:"cluster_healthy" xml:"cluster_healthy" require:"true"`
	DockerVersion          *string                                             `json:"docker_version" xml:"docker_version" require:"true"`
	SwarmMode              *bool                                               `json:"swarm_mode" xml:"swarm_mode" require:"true"`
	GwBridge               *string                                             `json:"gw_bridge" xml:"gw_bridge" require:"true"`
	PrivateZone            *bool                                               `json:"private_zone" xml:"private_zone" require:"true"`
	Profile                *string                                             `json:"profile" xml:"profile" require:"true"`
	DeletionProtection     *bool                                               `json:"deletion_protection" xml:"deletion_protection" require:"true"`
	Capabilities           *string                                             `json:"capabilities" xml:"capabilities" require:"true"`
	EnabledMigration       *bool                                               `json:"enabled_migration" xml:"enabled_migration" require:"true"`
	NeedUpdateAgent        *bool                                               `json:"need_update_agent" xml:"need_update_agent" require:"true"`
	Tags                   []*DescribeClusterDetailResponseBodyTags            `json:"tags" xml:"tags" require:"true" type:"Repeated"`
	Outputs                []*DescribeClusterDetailResponseBodyOutputs         `json:"outputs" xml:"outputs" require:"true" type:"Repeated"`
	UpgradeComponents      *DescribeClusterDetailResponseBodyUpgradeComponents `json:"upgrade_components" xml:"upgrade_components" require:"true" type:"Struct"`
	Parameters             *DescribeClusterDetailResponseBodyParameters        `json:"parameters" xml:"parameters" require:"true" type:"Struct"`
}

func (s DescribeClusterDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeClusterDetailResponseBody) SetName(v string) *DescribeClusterDetailResponseBody {
	s.Name = &v
	return s
}

func (s *DescribeClusterDetailResponseBody) SetClusterId(v string) *DescribeClusterDetailResponseBody {
	s.ClusterId = &v
	return s
}

func (s *DescribeClusterDetailResponseBody) SetSize(v int) *DescribeClusterDetailResponseBody {
	s.Size = &v
	return s
}

func (s *DescribeClusterDetailResponseBody) SetRegionId(v string) *DescribeClusterDetailResponseBody {
	s.RegionId = &v
	return s
}

func (s *DescribeClusterDetailResponseBody) SetState(v string) *DescribeClusterDetailResponseBody {
	s.State = &v
	return s
}

func (s *DescribeClusterDetailResponseBody) SetClusterType(v string) *DescribeClusterDetailResponseBody {
	s.ClusterType = &v
	return s
}

func (s *DescribeClusterDetailResponseBody) SetCreated(v string) *DescribeClusterDetailResponseBody {
	s.Created = &v
	return s
}

func (s *DescribeClusterDetailResponseBody) SetUpdated(v string) *DescribeClusterDetailResponseBody {
	s.Updated = &v
	return s
}

func (s *DescribeClusterDetailResponseBody) SetInitVersion(v string) *DescribeClusterDetailResponseBody {
	s.InitVersion = &v
	return s
}

func (s *DescribeClusterDetailResponseBody) SetCurrentVersion(v string) *DescribeClusterDetailResponseBody {
	s.CurrentVersion = &v
	return s
}

func (s *DescribeClusterDetailResponseBody) SetMetaData(v string) *DescribeClusterDetailResponseBody {
	s.MetaData = &v
	return s
}

func (s *DescribeClusterDetailResponseBody) SetResourceGroupId(v string) *DescribeClusterDetailResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeClusterDetailResponseBody) SetInstanceType(v string) *DescribeClusterDetailResponseBody {
	s.InstanceType = &v
	return s
}

func (s *DescribeClusterDetailResponseBody) SetVpcId(v string) *DescribeClusterDetailResponseBody {
	s.VpcId = &v
	return s
}

func (s *DescribeClusterDetailResponseBody) SetVswitchId(v string) *DescribeClusterDetailResponseBody {
	s.VswitchId = &v
	return s
}

func (s *DescribeClusterDetailResponseBody) SetVswitchCidr(v string) *DescribeClusterDetailResponseBody {
	s.VswitchCidr = &v
	return s
}

func (s *DescribeClusterDetailResponseBody) SetDataDiskSize(v int) *DescribeClusterDetailResponseBody {
	s.DataDiskSize = &v
	return s
}

func (s *DescribeClusterDetailResponseBody) SetDataDiskCategory(v string) *DescribeClusterDetailResponseBody {
	s.DataDiskCategory = &v
	return s
}

func (s *DescribeClusterDetailResponseBody) SetSecurityGroupId(v string) *DescribeClusterDetailResponseBody {
	s.SecurityGroupId = &v
	return s
}

func (s *DescribeClusterDetailResponseBody) SetZoneId(v string) *DescribeClusterDetailResponseBody {
	s.ZoneId = &v
	return s
}

func (s *DescribeClusterDetailResponseBody) SetNetworkMode(v string) *DescribeClusterDetailResponseBody {
	s.NetworkMode = &v
	return s
}

func (s *DescribeClusterDetailResponseBody) SetSubnetCidr(v string) *DescribeClusterDetailResponseBody {
	s.SubnetCidr = &v
	return s
}

func (s *DescribeClusterDetailResponseBody) SetMasterUrl(v string) *DescribeClusterDetailResponseBody {
	s.MasterUrl = &v
	return s
}

func (s *DescribeClusterDetailResponseBody) SetExternalLoadbalancerId(v string) *DescribeClusterDetailResponseBody {
	s.ExternalLoadbalancerId = &v
	return s
}

func (s *DescribeClusterDetailResponseBody) SetPort(v int) *DescribeClusterDetailResponseBody {
	s.Port = &v
	return s
}

func (s *DescribeClusterDetailResponseBody) SetNodeStatus(v string) *DescribeClusterDetailResponseBody {
	s.NodeStatus = &v
	return s
}

func (s *DescribeClusterDetailResponseBody) SetClusterHealthy(v string) *DescribeClusterDetailResponseBody {
	s.ClusterHealthy = &v
	return s
}

func (s *DescribeClusterDetailResponseBody) SetDockerVersion(v string) *DescribeClusterDetailResponseBody {
	s.DockerVersion = &v
	return s
}

func (s *DescribeClusterDetailResponseBody) SetSwarmMode(v bool) *DescribeClusterDetailResponseBody {
	s.SwarmMode = &v
	return s
}

func (s *DescribeClusterDetailResponseBody) SetGwBridge(v string) *DescribeClusterDetailResponseBody {
	s.GwBridge = &v
	return s
}

func (s *DescribeClusterDetailResponseBody) SetPrivateZone(v bool) *DescribeClusterDetailResponseBody {
	s.PrivateZone = &v
	return s
}

func (s *DescribeClusterDetailResponseBody) SetProfile(v string) *DescribeClusterDetailResponseBody {
	s.Profile = &v
	return s
}

func (s *DescribeClusterDetailResponseBody) SetDeletionProtection(v bool) *DescribeClusterDetailResponseBody {
	s.DeletionProtection = &v
	return s
}

func (s *DescribeClusterDetailResponseBody) SetCapabilities(v string) *DescribeClusterDetailResponseBody {
	s.Capabilities = &v
	return s
}

func (s *DescribeClusterDetailResponseBody) SetEnabledMigration(v bool) *DescribeClusterDetailResponseBody {
	s.EnabledMigration = &v
	return s
}

func (s *DescribeClusterDetailResponseBody) SetNeedUpdateAgent(v bool) *DescribeClusterDetailResponseBody {
	s.NeedUpdateAgent = &v
	return s
}

func (s *DescribeClusterDetailResponseBody) SetTags(v []*DescribeClusterDetailResponseBodyTags) *DescribeClusterDetailResponseBody {
	s.Tags = v
	return s
}

func (s *DescribeClusterDetailResponseBody) SetOutputs(v []*DescribeClusterDetailResponseBodyOutputs) *DescribeClusterDetailResponseBody {
	s.Outputs = v
	return s
}

func (s *DescribeClusterDetailResponseBody) SetUpgradeComponents(v *DescribeClusterDetailResponseBodyUpgradeComponents) *DescribeClusterDetailResponseBody {
	s.UpgradeComponents = v
	return s
}

func (s *DescribeClusterDetailResponseBody) SetParameters(v *DescribeClusterDetailResponseBodyParameters) *DescribeClusterDetailResponseBody {
	s.Parameters = v
	return s
}

type DescribeClusterDetailResponseBodyTags struct {
	Key   *string `json:"key" xml:"key" require:"true"`
	Value *string `json:"value" xml:"value" require:"true"`
}

func (s DescribeClusterDetailResponseBodyTags) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterDetailResponseBodyTags) GoString() string {
	return s.String()
}

func (s *DescribeClusterDetailResponseBodyTags) SetKey(v string) *DescribeClusterDetailResponseBodyTags {
	s.Key = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyTags) SetValue(v string) *DescribeClusterDetailResponseBodyTags {
	s.Value = &v
	return s
}

type DescribeClusterDetailResponseBodyOutputs struct {
	Description *string `json:"Description" xml:"Description" require:"true"`
	OutputKey   *string `json:"OutputKey" xml:"OutputKey" require:"true"`
	OutputValue *string `json:"OutputValue" xml:"OutputValue" require:"true"`
}

func (s DescribeClusterDetailResponseBodyOutputs) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterDetailResponseBodyOutputs) GoString() string {
	return s.String()
}

func (s *DescribeClusterDetailResponseBodyOutputs) SetDescription(v string) *DescribeClusterDetailResponseBodyOutputs {
	s.Description = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyOutputs) SetOutputKey(v string) *DescribeClusterDetailResponseBodyOutputs {
	s.OutputKey = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyOutputs) SetOutputValue(v string) *DescribeClusterDetailResponseBodyOutputs {
	s.OutputValue = &v
	return s
}

type DescribeClusterDetailResponseBodyUpgradeComponents struct {
	Kubernetes *DescribeClusterDetailResponseBodyUpgradeComponentsKubernetes `json:"Kubernetes" xml:"Kubernetes" require:"true" type:"Struct"`
}

func (s DescribeClusterDetailResponseBodyUpgradeComponents) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterDetailResponseBodyUpgradeComponents) GoString() string {
	return s.String()
}

func (s *DescribeClusterDetailResponseBodyUpgradeComponents) SetKubernetes(v *DescribeClusterDetailResponseBodyUpgradeComponentsKubernetes) *DescribeClusterDetailResponseBodyUpgradeComponents {
	s.Kubernetes = v
	return s
}

type DescribeClusterDetailResponseBodyUpgradeComponentsKubernetes struct {
	ComponentName  *string `json:"component_name" xml:"component_name" require:"true"`
	Version        *string `json:"version" xml:"version" require:"true"`
	NextVersion    *string `json:"next_version" xml:"next_version" require:"true"`
	Changed        *string `json:"changed" xml:"changed" require:"true"`
	CanUpgrade     *bool   `json:"can_upgrade" xml:"can_upgrade" require:"true"`
	Force          *bool   `json:"force" xml:"force" require:"true"`
	Policy         *string `json:"policy" xml:"policy" require:"true"`
	ExtraVars      *string `json:"ExtraVars" xml:"ExtraVars" require:"true"`
	ReadyToUpgrade *string `json:"ready_to_upgrade" xml:"ready_to_upgrade" require:"true"`
	Message        *string `json:"message" xml:"message" require:"true"`
	Exist          *bool   `json:"exist" xml:"exist" require:"true"`
	Required       *bool   `json:"required" xml:"required" require:"true"`
	Template       *string `json:"template" xml:"template" require:"true"`
}

func (s DescribeClusterDetailResponseBodyUpgradeComponentsKubernetes) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterDetailResponseBodyUpgradeComponentsKubernetes) GoString() string {
	return s.String()
}

func (s *DescribeClusterDetailResponseBodyUpgradeComponentsKubernetes) SetComponentName(v string) *DescribeClusterDetailResponseBodyUpgradeComponentsKubernetes {
	s.ComponentName = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyUpgradeComponentsKubernetes) SetVersion(v string) *DescribeClusterDetailResponseBodyUpgradeComponentsKubernetes {
	s.Version = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyUpgradeComponentsKubernetes) SetNextVersion(v string) *DescribeClusterDetailResponseBodyUpgradeComponentsKubernetes {
	s.NextVersion = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyUpgradeComponentsKubernetes) SetChanged(v string) *DescribeClusterDetailResponseBodyUpgradeComponentsKubernetes {
	s.Changed = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyUpgradeComponentsKubernetes) SetCanUpgrade(v bool) *DescribeClusterDetailResponseBodyUpgradeComponentsKubernetes {
	s.CanUpgrade = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyUpgradeComponentsKubernetes) SetForce(v bool) *DescribeClusterDetailResponseBodyUpgradeComponentsKubernetes {
	s.Force = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyUpgradeComponentsKubernetes) SetPolicy(v string) *DescribeClusterDetailResponseBodyUpgradeComponentsKubernetes {
	s.Policy = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyUpgradeComponentsKubernetes) SetExtraVars(v string) *DescribeClusterDetailResponseBodyUpgradeComponentsKubernetes {
	s.ExtraVars = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyUpgradeComponentsKubernetes) SetReadyToUpgrade(v string) *DescribeClusterDetailResponseBodyUpgradeComponentsKubernetes {
	s.ReadyToUpgrade = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyUpgradeComponentsKubernetes) SetMessage(v string) *DescribeClusterDetailResponseBodyUpgradeComponentsKubernetes {
	s.Message = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyUpgradeComponentsKubernetes) SetExist(v bool) *DescribeClusterDetailResponseBodyUpgradeComponentsKubernetes {
	s.Exist = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyUpgradeComponentsKubernetes) SetRequired(v bool) *DescribeClusterDetailResponseBodyUpgradeComponentsKubernetes {
	s.Required = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyUpgradeComponentsKubernetes) SetTemplate(v string) *DescribeClusterDetailResponseBodyUpgradeComponentsKubernetes {
	s.Template = &v
	return s
}

type DescribeClusterDetailResponseBodyParameters struct {
	AdjustmentType           *string `json:"AdjustmentType" xml:"AdjustmentType" require:"true"`
	AuditFlags               *string `json:"AuditFlags" xml:"AuditFlags" require:"true"`
	BetaVersion              *string `json:"BetaVersion" xml:"BetaVersion" require:"true"`
	CA                       *string `json:"CA" xml:"CA" require:"true"`
	ClientCA                 *string `json:"ClientCA" xml:"ClientCA" require:"true"`
	CloudMonitorFlags        *string `json:"CloudMonitorFlags" xml:"CloudMonitorFlags" require:"true"`
	CloudMonitorVersion      *string `json:"CloudMonitorVersion" xml:"CloudMonitorVersion" require:"true"`
	ContainerCIDR            *string `json:"ContainerCIDR" xml:"ContainerCIDR" require:"true"`
	DisableAddons            *string `json:"DisableAddons" xml:"DisableAddons" require:"true"`
	DockerVersion            *string `json:"DockerVersion" xml:"DockerVersion" require:"true"`
	ESSDeletionProtection    *string `json:"ESSDeletionProtection" xml:"ESSDeletionProtection" require:"true"`
	Eip                      *string `json:"Eip" xml:"Eip" require:"true"`
	EipAddress               *string `json:"EipAddress" xml:"EipAddress" require:"true"`
	ElasticSearchHost        *string `json:"ElasticSearchHost" xml:"ElasticSearchHost" require:"true"`
	ElasticSearchPass        *string `json:"ElasticSearchPass" xml:"ElasticSearchPass" require:"true"`
	ElasticSearchPort        *string `json:"ElasticSearchPort" xml:"ElasticSearchPort" require:"true"`
	ElasticSearchUser        *string `json:"ElasticSearchUser" xml:"ElasticSearchUser" require:"true"`
	EtcdVersion              *string `json:"EtcdVersion" xml:"EtcdVersion" require:"true"`
	ExecuteVersion           *string `json:"ExecuteVersion" xml:"ExecuteVersion" require:"true"`
	GPUFlags                 *string `json:"GPUFlags" xml:"GPUFlags" require:"true"`
	HealthCheckType          *string `json:"HealthCheckType" xml:"HealthCheckType" require:"true"`
	ImageId                  *string `json:"ImageId" xml:"ImageId" require:"true"`
	K8SMasterPolicyDocument  *string `json:"K8SMasterPolicyDocument" xml:"K8SMasterPolicyDocument" require:"true"`
	K8sWorkerPolicyDocument  *string `json:"K8sWorkerPolicyDocument" xml:"K8sWorkerPolicyDocument" require:"true"`
	Key                      *string `json:"Key" xml:"Key" require:"true"`
	KeyPair                  *string `json:"KeyPair" xml:"KeyPair" require:"true"`
	KubernetesVersion        *string `json:"KubernetesVersion" xml:"KubernetesVersion" require:"true"`
	LoggingType              *string `json:"LoggingType" xml:"LoggingType" require:"true"`
	MasterAmounts            *string `json:"MasterAmounts" xml:"MasterAmounts" require:"true"`
	MasterAutoRenew          *string `json:"MasterAutoRenew" xml:"MasterAutoRenew" require:"true"`
	MasterAutoRenewPeriod    *string `json:"MasterAutoRenewPeriod" xml:"MasterAutoRenewPeriod" require:"true"`
	MasterCount              *string `json:"MasterCount" xml:"MasterCount" require:"true"`
	MasterDataDisk           *string `json:"MasterDataDisk" xml:"MasterDataDisk" require:"true"`
	MasterDataDisks          *string `json:"MasterDataDisks" xml:"MasterDataDisks" require:"true"`
	MasterDeletionProtection *string `json:"MasterDeletionProtection" xml:"MasterDeletionProtection" require:"true"`
	MasterDeploymentSetId    *string `json:"MasterDeploymentSetId" xml:"MasterDeploymentSetId" require:"true"`
	MasterHpcClusterId       *string `json:"MasterHpcClusterId" xml:"MasterHpcClusterId" require:"true"`
	MasterImageId            *string `json:"MasterImageId" xml:"MasterImageId" require:"true"`
	MasterInstanceChargeType *string `json:"MasterInstanceChargeType" xml:"MasterInstanceChargeType" require:"true"`
	MasterInstanceTypes      *string `json:"MasterInstanceTypes" xml:"MasterInstanceTypes" require:"true"`
	MasterKeyPair            *string `json:"MasterKeyPair" xml:"MasterKeyPair" require:"true"`
	MasterLoginPassword      *string `json:"MasterLoginPassword" xml:"MasterLoginPassword" require:"true"`
	MasterPeriod             *string `json:"MasterPeriod" xml:"MasterPeriod" require:"true"`
	MasterPeriodUnit         *string `json:"MasterPeriodUnit" xml:"MasterPeriodUnit" require:"true"`
	MasterSlbSShHealthCheck  *string `json:"MasterSlbSShHealthCheck" xml:"MasterSlbSShHealthCheck" require:"true"`
	MasterSnapshotPolicyId   *string `json:"MasterSnapshotPolicyId" xml:"MasterSnapshotPolicyId" require:"true"`
	MasterSystemDiskCategory *string `json:"MasterSystemDiskCategory" xml:"MasterSystemDiskCategory" require:"true"`
	MasterSystemDiskSize     *string `json:"MasterSystemDiskSize" xml:"MasterSystemDiskSize" require:"true"`
	MasterVSwitchIds         *string `json:"MasterVSwitchIds" xml:"MasterVSwitchIds" require:"true"`
	NatGateway               *string `json:"NatGateway" xml:"NatGateway" require:"true"`
	NatGatewayId             *string `json:"NatGatewayId" xml:"NatGatewayId" require:"true"`
	Network                  *string `json:"Network" xml:"Network" require:"true"`
	NodeCIDRMask             *string `json:"NodeCIDRMask" xml:"NodeCIDRMask" require:"true"`
	NodeNameMode             *string `json:"NodeNameMode" xml:"NodeNameMode" require:"true"`
	NumOfNodes               *string `json:"NumOfNodes" xml:"NumOfNodes" require:"true"`
	Password                 *string `json:"Password" xml:"Password" require:"true"`
	PodVswitchIds            *string `json:"PodVswitchIds" xml:"PodVswitchIds" require:"true"`
	ProtectedInstances       *string `json:"ProtectedInstances" xml:"ProtectedInstances" require:"true"`
	ProxyMode                *string `json:"ProxyMode" xml:"ProxyMode" require:"true"`
	PublicSLB                *string `json:"PublicSLB" xml:"PublicSLB" require:"true"`
	RemoveInstanceIds        *string `json:"RemoveInstanceIds" xml:"RemoveInstanceIds" require:"true"`
	ResourceGroupId          *string `json:"ResourceGroupId" xml:"ResourceGroupId" require:"true"`
	SLBDeletionProtection    *string `json:"SLBDeletionProtection" xml:"SLBDeletionProtection" require:"true"`
	SLSProjectName           *string `json:"SLSProjectName" xml:"SLSProjectName" require:"true"`
	SNatEntry                *string `json:"SNatEntry" xml:"SNatEntry" require:"true"`
	SSHFlags                 *string `json:"SSHFlags" xml:"SSHFlags" require:"true"`
	SecurityGroupId          *string `json:"SecurityGroupId" xml:"SecurityGroupId" require:"true"`
	ServiceCIDR              *string `json:"ServiceCIDR" xml:"ServiceCIDR" require:"true"`
	SetUpArgs                *string `json:"SetUpArgs" xml:"SetUpArgs" require:"true"`
	SnatTableId              *string `json:"SnatTableId" xml:"SnatTableId" require:"true"`
	Tags                     *string `json:"Tags" xml:"Tags" require:"true"`
	UserCA                   *string `json:"UserCA" xml:"UserCA" require:"true"`
	UserData                 *string `json:"UserData" xml:"UserData" require:"true"`
	VpcId                    *string `json:"VpcId" xml:"VpcId" require:"true"`
	WillReplace              *string `json:"WillReplace" xml:"WillReplace" require:"true"`
	WorkerAutoRenew          *string `json:"WorkerAutoRenew" xml:"WorkerAutoRenew" require:"true"`
	WorkerAutoRenewPeriod    *string `json:"WorkerAutoRenewPeriod" xml:"WorkerAutoRenewPeriod" require:"true"`
	WorkerDataDisk           *string `json:"WorkerDataDisk" xml:"WorkerDataDisk" require:"true"`
	WorkerDataDisks          *string `json:"WorkerDataDisks" xml:"WorkerDataDisks" require:"true"`
	WorkerDeletionProtection *string `json:"WorkerDeletionProtection" xml:"WorkerDeletionProtection" require:"true"`
	WorkerDeploymentSetId    *string `json:"WorkerDeploymentSetId" xml:"WorkerDeploymentSetId" require:"true"`
	WorkerHpcClusterId       *string `json:"WorkerHpcClusterId" xml:"WorkerHpcClusterId" require:"true"`
	WorkerImageId            *string `json:"WorkerImageId" xml:"WorkerImageId" require:"true"`
	WorkerInstanceChargeType *string `json:"WorkerInstanceChargeType" xml:"WorkerInstanceChargeType" require:"true"`
	WorkerInstanceTypes      *string `json:"WorkerInstanceTypes" xml:"WorkerInstanceTypes" require:"true"`
	WorkerKeyPair            *string `json:"WorkerKeyPair" xml:"WorkerKeyPair" require:"true"`
	WorkerLoginPassword      *string `json:"WorkerLoginPassword" xml:"WorkerLoginPassword" require:"true"`
	WorkerPeriod             *string `json:"WorkerPeriod" xml:"WorkerPeriod" require:"true"`
	WorkerPeriodUnit         *string `json:"WorkerPeriodUnit" xml:"WorkerPeriodUnit" require:"true"`
	WorkerSnapshotPolicyId   *string `json:"WorkerSnapshotPolicyId" xml:"WorkerSnapshotPolicyId" require:"true"`
	WorkerSystemDiskCategory *string `json:"WorkerSystemDiskCategory" xml:"WorkerSystemDiskCategory" require:"true"`
	WorkerSystemDiskSize     *string `json:"WorkerSystemDiskSize" xml:"WorkerSystemDiskSize" require:"true"`
	WorkerVSwitchIds         *string `json:"WorkerVSwitchIds" xml:"WorkerVSwitchIds" require:"true"`
	ZoneId                   *string `json:"ZoneId" xml:"ZoneId" require:"true"`
}

func (s DescribeClusterDetailResponseBodyParameters) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterDetailResponseBodyParameters) GoString() string {
	return s.String()
}

func (s *DescribeClusterDetailResponseBodyParameters) SetAdjustmentType(v string) *DescribeClusterDetailResponseBodyParameters {
	s.AdjustmentType = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyParameters) SetAuditFlags(v string) *DescribeClusterDetailResponseBodyParameters {
	s.AuditFlags = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyParameters) SetBetaVersion(v string) *DescribeClusterDetailResponseBodyParameters {
	s.BetaVersion = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyParameters) SetCA(v string) *DescribeClusterDetailResponseBodyParameters {
	s.CA = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyParameters) SetClientCA(v string) *DescribeClusterDetailResponseBodyParameters {
	s.ClientCA = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyParameters) SetCloudMonitorFlags(v string) *DescribeClusterDetailResponseBodyParameters {
	s.CloudMonitorFlags = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyParameters) SetCloudMonitorVersion(v string) *DescribeClusterDetailResponseBodyParameters {
	s.CloudMonitorVersion = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyParameters) SetContainerCIDR(v string) *DescribeClusterDetailResponseBodyParameters {
	s.ContainerCIDR = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyParameters) SetDisableAddons(v string) *DescribeClusterDetailResponseBodyParameters {
	s.DisableAddons = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyParameters) SetDockerVersion(v string) *DescribeClusterDetailResponseBodyParameters {
	s.DockerVersion = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyParameters) SetESSDeletionProtection(v string) *DescribeClusterDetailResponseBodyParameters {
	s.ESSDeletionProtection = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyParameters) SetEip(v string) *DescribeClusterDetailResponseBodyParameters {
	s.Eip = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyParameters) SetEipAddress(v string) *DescribeClusterDetailResponseBodyParameters {
	s.EipAddress = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyParameters) SetElasticSearchHost(v string) *DescribeClusterDetailResponseBodyParameters {
	s.ElasticSearchHost = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyParameters) SetElasticSearchPass(v string) *DescribeClusterDetailResponseBodyParameters {
	s.ElasticSearchPass = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyParameters) SetElasticSearchPort(v string) *DescribeClusterDetailResponseBodyParameters {
	s.ElasticSearchPort = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyParameters) SetElasticSearchUser(v string) *DescribeClusterDetailResponseBodyParameters {
	s.ElasticSearchUser = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyParameters) SetEtcdVersion(v string) *DescribeClusterDetailResponseBodyParameters {
	s.EtcdVersion = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyParameters) SetExecuteVersion(v string) *DescribeClusterDetailResponseBodyParameters {
	s.ExecuteVersion = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyParameters) SetGPUFlags(v string) *DescribeClusterDetailResponseBodyParameters {
	s.GPUFlags = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyParameters) SetHealthCheckType(v string) *DescribeClusterDetailResponseBodyParameters {
	s.HealthCheckType = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyParameters) SetImageId(v string) *DescribeClusterDetailResponseBodyParameters {
	s.ImageId = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyParameters) SetK8SMasterPolicyDocument(v string) *DescribeClusterDetailResponseBodyParameters {
	s.K8SMasterPolicyDocument = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyParameters) SetK8sWorkerPolicyDocument(v string) *DescribeClusterDetailResponseBodyParameters {
	s.K8sWorkerPolicyDocument = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyParameters) SetKey(v string) *DescribeClusterDetailResponseBodyParameters {
	s.Key = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyParameters) SetKeyPair(v string) *DescribeClusterDetailResponseBodyParameters {
	s.KeyPair = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyParameters) SetKubernetesVersion(v string) *DescribeClusterDetailResponseBodyParameters {
	s.KubernetesVersion = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyParameters) SetLoggingType(v string) *DescribeClusterDetailResponseBodyParameters {
	s.LoggingType = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyParameters) SetMasterAmounts(v string) *DescribeClusterDetailResponseBodyParameters {
	s.MasterAmounts = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyParameters) SetMasterAutoRenew(v string) *DescribeClusterDetailResponseBodyParameters {
	s.MasterAutoRenew = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyParameters) SetMasterAutoRenewPeriod(v string) *DescribeClusterDetailResponseBodyParameters {
	s.MasterAutoRenewPeriod = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyParameters) SetMasterCount(v string) *DescribeClusterDetailResponseBodyParameters {
	s.MasterCount = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyParameters) SetMasterDataDisk(v string) *DescribeClusterDetailResponseBodyParameters {
	s.MasterDataDisk = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyParameters) SetMasterDataDisks(v string) *DescribeClusterDetailResponseBodyParameters {
	s.MasterDataDisks = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyParameters) SetMasterDeletionProtection(v string) *DescribeClusterDetailResponseBodyParameters {
	s.MasterDeletionProtection = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyParameters) SetMasterDeploymentSetId(v string) *DescribeClusterDetailResponseBodyParameters {
	s.MasterDeploymentSetId = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyParameters) SetMasterHpcClusterId(v string) *DescribeClusterDetailResponseBodyParameters {
	s.MasterHpcClusterId = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyParameters) SetMasterImageId(v string) *DescribeClusterDetailResponseBodyParameters {
	s.MasterImageId = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyParameters) SetMasterInstanceChargeType(v string) *DescribeClusterDetailResponseBodyParameters {
	s.MasterInstanceChargeType = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyParameters) SetMasterInstanceTypes(v string) *DescribeClusterDetailResponseBodyParameters {
	s.MasterInstanceTypes = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyParameters) SetMasterKeyPair(v string) *DescribeClusterDetailResponseBodyParameters {
	s.MasterKeyPair = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyParameters) SetMasterLoginPassword(v string) *DescribeClusterDetailResponseBodyParameters {
	s.MasterLoginPassword = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyParameters) SetMasterPeriod(v string) *DescribeClusterDetailResponseBodyParameters {
	s.MasterPeriod = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyParameters) SetMasterPeriodUnit(v string) *DescribeClusterDetailResponseBodyParameters {
	s.MasterPeriodUnit = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyParameters) SetMasterSlbSShHealthCheck(v string) *DescribeClusterDetailResponseBodyParameters {
	s.MasterSlbSShHealthCheck = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyParameters) SetMasterSnapshotPolicyId(v string) *DescribeClusterDetailResponseBodyParameters {
	s.MasterSnapshotPolicyId = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyParameters) SetMasterSystemDiskCategory(v string) *DescribeClusterDetailResponseBodyParameters {
	s.MasterSystemDiskCategory = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyParameters) SetMasterSystemDiskSize(v string) *DescribeClusterDetailResponseBodyParameters {
	s.MasterSystemDiskSize = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyParameters) SetMasterVSwitchIds(v string) *DescribeClusterDetailResponseBodyParameters {
	s.MasterVSwitchIds = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyParameters) SetNatGateway(v string) *DescribeClusterDetailResponseBodyParameters {
	s.NatGateway = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyParameters) SetNatGatewayId(v string) *DescribeClusterDetailResponseBodyParameters {
	s.NatGatewayId = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyParameters) SetNetwork(v string) *DescribeClusterDetailResponseBodyParameters {
	s.Network = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyParameters) SetNodeCIDRMask(v string) *DescribeClusterDetailResponseBodyParameters {
	s.NodeCIDRMask = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyParameters) SetNodeNameMode(v string) *DescribeClusterDetailResponseBodyParameters {
	s.NodeNameMode = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyParameters) SetNumOfNodes(v string) *DescribeClusterDetailResponseBodyParameters {
	s.NumOfNodes = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyParameters) SetPassword(v string) *DescribeClusterDetailResponseBodyParameters {
	s.Password = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyParameters) SetPodVswitchIds(v string) *DescribeClusterDetailResponseBodyParameters {
	s.PodVswitchIds = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyParameters) SetProtectedInstances(v string) *DescribeClusterDetailResponseBodyParameters {
	s.ProtectedInstances = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyParameters) SetProxyMode(v string) *DescribeClusterDetailResponseBodyParameters {
	s.ProxyMode = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyParameters) SetPublicSLB(v string) *DescribeClusterDetailResponseBodyParameters {
	s.PublicSLB = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyParameters) SetRemoveInstanceIds(v string) *DescribeClusterDetailResponseBodyParameters {
	s.RemoveInstanceIds = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyParameters) SetResourceGroupId(v string) *DescribeClusterDetailResponseBodyParameters {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyParameters) SetSLBDeletionProtection(v string) *DescribeClusterDetailResponseBodyParameters {
	s.SLBDeletionProtection = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyParameters) SetSLSProjectName(v string) *DescribeClusterDetailResponseBodyParameters {
	s.SLSProjectName = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyParameters) SetSNatEntry(v string) *DescribeClusterDetailResponseBodyParameters {
	s.SNatEntry = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyParameters) SetSSHFlags(v string) *DescribeClusterDetailResponseBodyParameters {
	s.SSHFlags = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyParameters) SetSecurityGroupId(v string) *DescribeClusterDetailResponseBodyParameters {
	s.SecurityGroupId = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyParameters) SetServiceCIDR(v string) *DescribeClusterDetailResponseBodyParameters {
	s.ServiceCIDR = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyParameters) SetSetUpArgs(v string) *DescribeClusterDetailResponseBodyParameters {
	s.SetUpArgs = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyParameters) SetSnatTableId(v string) *DescribeClusterDetailResponseBodyParameters {
	s.SnatTableId = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyParameters) SetTags(v string) *DescribeClusterDetailResponseBodyParameters {
	s.Tags = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyParameters) SetUserCA(v string) *DescribeClusterDetailResponseBodyParameters {
	s.UserCA = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyParameters) SetUserData(v string) *DescribeClusterDetailResponseBodyParameters {
	s.UserData = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyParameters) SetVpcId(v string) *DescribeClusterDetailResponseBodyParameters {
	s.VpcId = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyParameters) SetWillReplace(v string) *DescribeClusterDetailResponseBodyParameters {
	s.WillReplace = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyParameters) SetWorkerAutoRenew(v string) *DescribeClusterDetailResponseBodyParameters {
	s.WorkerAutoRenew = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyParameters) SetWorkerAutoRenewPeriod(v string) *DescribeClusterDetailResponseBodyParameters {
	s.WorkerAutoRenewPeriod = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyParameters) SetWorkerDataDisk(v string) *DescribeClusterDetailResponseBodyParameters {
	s.WorkerDataDisk = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyParameters) SetWorkerDataDisks(v string) *DescribeClusterDetailResponseBodyParameters {
	s.WorkerDataDisks = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyParameters) SetWorkerDeletionProtection(v string) *DescribeClusterDetailResponseBodyParameters {
	s.WorkerDeletionProtection = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyParameters) SetWorkerDeploymentSetId(v string) *DescribeClusterDetailResponseBodyParameters {
	s.WorkerDeploymentSetId = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyParameters) SetWorkerHpcClusterId(v string) *DescribeClusterDetailResponseBodyParameters {
	s.WorkerHpcClusterId = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyParameters) SetWorkerImageId(v string) *DescribeClusterDetailResponseBodyParameters {
	s.WorkerImageId = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyParameters) SetWorkerInstanceChargeType(v string) *DescribeClusterDetailResponseBodyParameters {
	s.WorkerInstanceChargeType = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyParameters) SetWorkerInstanceTypes(v string) *DescribeClusterDetailResponseBodyParameters {
	s.WorkerInstanceTypes = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyParameters) SetWorkerKeyPair(v string) *DescribeClusterDetailResponseBodyParameters {
	s.WorkerKeyPair = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyParameters) SetWorkerLoginPassword(v string) *DescribeClusterDetailResponseBodyParameters {
	s.WorkerLoginPassword = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyParameters) SetWorkerPeriod(v string) *DescribeClusterDetailResponseBodyParameters {
	s.WorkerPeriod = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyParameters) SetWorkerPeriodUnit(v string) *DescribeClusterDetailResponseBodyParameters {
	s.WorkerPeriodUnit = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyParameters) SetWorkerSnapshotPolicyId(v string) *DescribeClusterDetailResponseBodyParameters {
	s.WorkerSnapshotPolicyId = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyParameters) SetWorkerSystemDiskCategory(v string) *DescribeClusterDetailResponseBodyParameters {
	s.WorkerSystemDiskCategory = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyParameters) SetWorkerSystemDiskSize(v string) *DescribeClusterDetailResponseBodyParameters {
	s.WorkerSystemDiskSize = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyParameters) SetWorkerVSwitchIds(v string) *DescribeClusterDetailResponseBodyParameters {
	s.WorkerVSwitchIds = &v
	return s
}

func (s *DescribeClusterDetailResponseBodyParameters) SetZoneId(v string) *DescribeClusterDetailResponseBodyParameters {
	s.ZoneId = &v
	return s
}

type DescribeClusterDetailResponse struct {
	Headers map[string]string                  `json:"headers" xml:"headers" require:"true"`
	Body    *DescribeClusterDetailResponseBody `json:"body" xml:"body" require:"true"`
}

func (s DescribeClusterDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeClusterDetailResponse) SetHeaders(v map[string]string) *DescribeClusterDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeClusterDetailResponse) SetBody(v *DescribeClusterDetailResponseBody) *DescribeClusterDetailResponse {
	s.Body = v
	return s
}

type DeleteClusterRequest struct {
	Headers map[string]string `json:"headers" xml:"headers"`
}

func (s DeleteClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteClusterRequest) GoString() string {
	return s.String()
}

func (s *DeleteClusterRequest) SetHeaders(v map[string]string) *DeleteClusterRequest {
	s.Headers = v
	return s
}

type DeleteClusterResponseBody struct {
	RequestId *string `json:"RequestId" xml:"RequestId" require:"true"`
}

func (s DeleteClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteClusterResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteClusterResponseBody) SetRequestId(v string) *DeleteClusterResponseBody {
	s.RequestId = &v
	return s
}

type DeleteClusterResponse struct {
	Headers map[string]string          `json:"headers" xml:"headers" require:"true"`
	Body    *DeleteClusterResponseBody `json:"body" xml:"body" require:"true"`
}

func (s DeleteClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteClusterResponse) GoString() string {
	return s.String()
}

func (s *DeleteClusterResponse) SetHeaders(v map[string]string) *DeleteClusterResponse {
	s.Headers = v
	return s
}

func (s *DeleteClusterResponse) SetBody(v *DeleteClusterResponseBody) *DeleteClusterResponse {
	s.Body = v
	return s
}

type Client struct {
	roa.Client
}

func NewClient(config *roa.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *roa.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = ""
	_err = client.CheckConfig(config)
	if _err != nil {
		return
	}
	client.EndpointHost, _err = client.GetEndpoint(client.ProductId, client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.EndpointHost)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) ResumeComponentUpgrade(clusterid string, componentid string, request *ResumeComponentUpgradeRequest, runtime *util.RuntimeOptions) (_result *ResumeComponentUpgradeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return
	}
	_result = &ResumeComponentUpgradeResponse{}
	_body, _err := client.DoRequest("2015-12-15", "HTTPS", "POST", "AK", "/clusters/"+tea.ToString(clusterid)+"/components/"+tea.ToString(componentid)+"/resume", nil, request.Headers, nil, runtime)
	if _err != nil {
		return nil, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PauseComponentUpgrade(clusterid string, componentid string, request *PauseComponentUpgradeRequest, runtime *util.RuntimeOptions) (_result *PauseComponentUpgradeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return
	}
	_result = &PauseComponentUpgradeResponse{}
	_body, _err := client.DoRequest("2015-12-15", "HTTPS", "POST", "AK", "/clusters/"+tea.ToString(clusterid)+"/components/"+tea.ToString(componentid)+"/pause", nil, request.Headers, nil, runtime)
	if _err != nil {
		return nil, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CancelComponentUpgrade(clusterid string, componentid string, request *CancelComponentUpgradeRequest, runtime *util.RuntimeOptions) (_result *CancelComponentUpgradeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return
	}
	_result = &CancelComponentUpgradeResponse{}
	_body, _err := client.DoRequest("2015-12-15", "HTTPS", "POST", "AK", "/clusters/"+tea.ToString(clusterid)+"/components/"+tea.ToString(componentid)+"/cancel", nil, request.Headers, nil, runtime)
	if _err != nil {
		return nil, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CancelWorkflow(workflowName string, request *CancelWorkflowRequest, runtime *util.RuntimeOptions) (_result *CancelWorkflowResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return
	}
	_result = &CancelWorkflowResponse{}
	_body, _err := client.DoRequest("2015-12-15", "HTTPS", "PUT", "AK", "/gs/workflow/"+tea.ToString(workflowName), nil, request.Headers, nil, runtime)
	if _err != nil {
		return nil, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescirbeWorkflow(workflowName string, request *DescirbeWorkflowRequest, runtime *util.RuntimeOptions) (_result *DescirbeWorkflowResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return
	}
	_result = &DescirbeWorkflowResponse{}
	_body, _err := client.DoRequest("2015-12-15", "HTTPS", "GET", "AK", "/gs/workflow/"+tea.ToString(workflowName), nil, request.Headers, nil, runtime)
	if _err != nil {
		return nil, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemoveWorkflow(workflowName string, request *RemoveWorkflowRequest, runtime *util.RuntimeOptions) (_result *RemoveWorkflowResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return
	}
	_result = &RemoveWorkflowResponse{}
	_body, _err := client.DoRequest("2015-12-15", "HTTPS", "DELETE", "AK", "/gs/workflow/"+tea.ToString(workflowName), nil, request.Headers, nil, runtime)
	if _err != nil {
		return nil, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeWorkflows(request *DescribeWorkflowsRequest, runtime *util.RuntimeOptions) (_result *DescribeWorkflowsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return
	}
	_result = &DescribeWorkflowsResponse{}
	_body, _err := client.DoRequest("2015-12-15", "HTTPS", "GET", "AK", "/gs/workflows", nil, request.Headers, nil, runtime)
	if _err != nil {
		return nil, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartWorkflow(request *StartWorkflowRequest, runtime *util.RuntimeOptions) (_result *StartWorkflowResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return
	}
	_result = &StartWorkflowResponse{}
	_body, _err := client.DoRequest("2015-12-15", "HTTPS", "POST", "AK", "/gs/workflow", nil, request.Headers, nil, runtime)
	if _err != nil {
		return nil, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateServiceMesh(request *CreateServiceMeshRequest, runtime *util.RuntimeOptions) (_result *CreateServiceMeshResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return
	}
	_result = &CreateServiceMeshResponse{}
	_body, _err := client.DoRequest("2015-12-15", "HTTPS", "POST", "AK", "/servicemesh", nil, request.Headers, nil, runtime)
	if _err != nil {
		return nil, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ServiceMeshAddCluster(serviceMeshId string, request *ServiceMeshAddClusterRequest, runtime *util.RuntimeOptions) (_result *ServiceMeshAddClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return
	}
	_result = &ServiceMeshAddClusterResponse{}
	_body, _err := client.DoRequest("2015-12-15", "HTTPS", "PUT", "AK", "/servicemesh/"+tea.ToString(serviceMeshId)+"/add/clusters", nil, request.Headers, nil, runtime)
	if _err != nil {
		return nil, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ServiceMeshRemoveCluster(serviceMeshId string, request *ServiceMeshRemoveClusterRequest, runtime *util.RuntimeOptions) (_result *ServiceMeshRemoveClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return
	}
	_result = &ServiceMeshRemoveClusterResponse{}
	_body, _err := client.DoRequest("2015-12-15", "HTTPS", "PUT", "AK", "/servicemesh/"+tea.ToString(serviceMeshId)+"/remove/clusters", nil, request.Headers, nil, runtime)
	if _err != nil {
		return nil, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateServiceMesh(serviceMeshId string, request *UpdateServiceMeshRequest, runtime *util.RuntimeOptions) (_result *UpdateServiceMeshResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return
	}
	_result = &UpdateServiceMeshResponse{}
	_body, _err := client.DoRequest("2015-12-15", "HTTPS", "PUT", "AK", "/servicemesh/"+tea.ToString(serviceMeshId), nil, request.Headers, nil, runtime)
	if _err != nil {
		return nil, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ServiceMeshApiServer(serviceMeshId string, request *ServiceMeshApiServerRequest, runtime *util.RuntimeOptions) (_result *ServiceMeshApiServerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return
	}
	_result = &ServiceMeshApiServerResponse{}
	_body, _err := client.DoRequest("2015-12-15", "HTTPS", "POST", "AK", "/servicemesh/"+tea.ToString(serviceMeshId)+"/api_proxy", nil, request.Headers, nil, runtime)
	if _err != nil {
		return nil, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UnInstallClusterAddons(clusterId string, request *UnInstallClusterAddonsRequest, runtime *util.RuntimeOptions) (_result *UnInstallClusterAddonsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return
	}
	_result = &UnInstallClusterAddonsResponse{}
	_body, _err := client.DoRequest("2015-12-15", "HTTPS", "POST", "AK", "/clusters/"+tea.ToString(clusterId)+"/components/uninstall", nil, request.Headers, tea.ToMap(request.Body), runtime)
	if _err != nil {
		return nil, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAddons(request *DescribeAddonsRequest, runtime *util.RuntimeOptions) (_result *DescribeAddonsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return
	}
	_result = &DescribeAddonsResponse{}
	_body, _err := client.DoRequest("2015-12-15", "HTTPS", "GET", "AK", "/clusters/components/metadata", util.StringifyMapValue(tea.ToMap(request.Query)), request.Headers, nil, runtime)
	if _err != nil {
		return nil, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateK8sClusterUserConfigExpire(clusterId string, request *UpdateK8sClusterUserConfigExpireRequest, runtime *util.RuntimeOptions) (_result *UpdateK8sClusterUserConfigExpireResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return
	}
	_result = &UpdateK8sClusterUserConfigExpireResponse{}
	_body, _err := client.DoRequest("2015-12-15", "HTTPS", "POST", "AK", "/k8s/"+tea.ToString(clusterId)+"/user_config/expire", nil, request.Headers, nil, runtime)
	if _err != nil {
		return nil, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeClusterV2UserKubeconfig(clusterId string, request *DescribeClusterV2UserKubeconfigRequest, runtime *util.RuntimeOptions) (_result *DescribeClusterV2UserKubeconfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return
	}
	_result = &DescribeClusterV2UserKubeconfigResponse{}
	_body, _err := client.DoRequest("2015-12-15", "HTTPS", "GET", "AK", "/api/v2/k8s/"+tea.ToString(clusterId)+"/user_config", util.StringifyMapValue(tea.ToMap(request.Query)), request.Headers, nil, runtime)
	if _err != nil {
		return nil, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyClusterTags(clusterId string, request *ModifyClusterTagsRequest, runtime *util.RuntimeOptions) (_result *ModifyClusterTagsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return
	}
	_result = &ModifyClusterTagsResponse{}
	_body, _err := client.DoRequest("2015-12-15", "HTTPS", "POST", "AK", "/clusters/"+tea.ToString(clusterId)+"/tags", nil, request.Headers, tea.ToMap(request.Body), runtime)
	if _err != nil {
		return nil, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ScaleOutCluster(clusterId string, request *ScaleOutClusterRequest, runtime *util.RuntimeOptions) (_result *ScaleOutClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return
	}
	_result = &ScaleOutClusterResponse{}
	_body, _err := client.DoRequest("2015-12-15", "HTTPS", "POST", "AK", "/api/v2/clusters/"+tea.ToString(clusterId), nil, request.Headers, tea.ToMap(request.Body), runtime)
	if _err != nil {
		return nil, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeClusterAddonsVersion(clusterId string, request *DescribeClusterAddonsVersionRequest, runtime *util.RuntimeOptions) (_result *DescribeClusterAddonsVersionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return
	}
	_result = &DescribeClusterAddonsVersionResponse{}
	_body, _err := client.DoRequest("2015-12-15", "HTTPS", "GET", "AK", "/clusters/"+tea.ToString(clusterId)+"/components/version", nil, request.Headers, nil, runtime)
	if _err != nil {
		return nil, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeClusterAddonUpgradeStatus(clusterId string, componentId string, request *DescribeClusterAddonUpgradeStatusRequest, runtime *util.RuntimeOptions) (_result *DescribeClusterAddonUpgradeStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return
	}
	_result = &DescribeClusterAddonUpgradeStatusResponse{}
	_body, _err := client.DoRequest("2015-12-15", "HTTPS", "GET", "AK", "/clusters/"+tea.ToString(clusterId)+"/components/"+tea.ToString(componentId)+"/upgradestatus", nil, request.Headers, nil, runtime)
	if _err != nil {
		return nil, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeEdgeClusterAttachScripts(clusterId string, request *DescribeEdgeClusterAttachScriptsRequest, runtime *util.RuntimeOptions) (_result *DescribeEdgeClusterAttachScriptsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return
	}
	_result = &DescribeEdgeClusterAttachScriptsResponse{}
	_body, _err := client.DoRequest("2015-12-15", "HTTPS", "GET", "AK", "/clusters/"+tea.ToString(clusterId)+"/attachscript", util.StringifyMapValue(tea.ToMap(request.Query)), request.Headers, nil, runtime)
	if _err != nil {
		return nil, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteClusterNodes(clusterId string, request *DeleteClusterNodesRequest, runtime *util.RuntimeOptions) (_result *DeleteClusterNodesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return
	}
	_result = &DeleteClusterNodesResponse{}
	_body, _err := client.DoRequest("2015-12-15", "HTTPS", "POST", "AK", "/clusters/"+tea.ToString(clusterId)+"/nodes", nil, request.Headers, tea.ToMap(request.Body), runtime)
	if _err != nil {
		return nil, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeClusterUserKubeconfig(clusterId string, request *DescribeClusterUserKubeconfigRequest, runtime *util.RuntimeOptions) (_result *DescribeClusterUserKubeconfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return
	}
	_result = &DescribeClusterUserKubeconfigResponse{}
	_body, _err := client.DoRequest("2015-12-15", "HTTPS", "GET", "AK", "/k8s/"+tea.ToString(clusterId)+"/user_config", util.StringifyMapValue(tea.ToMap(request.Query)), request.Headers, nil, runtime)
	if _err != nil {
		return nil, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AttachInstances(clusterId string, request *AttachInstancesRequest, runtime *util.RuntimeOptions) (_result *AttachInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return
	}
	_result = &AttachInstancesResponse{}
	_body, _err := client.DoRequest("2015-12-15", "HTTPS", "POST", "AK", "/clusters/"+tea.ToString(clusterId)+"/attach", nil, request.Headers, tea.ToMap(request.Body), runtime)
	if _err != nil {
		return nil, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateCluster(request *CreateClusterRequest, runtime *util.RuntimeOptions) (_result *CreateClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return
	}
	_result = &CreateClusterResponse{}
	_body, _err := client.DoRequest("2015-12-15", "HTTPS", "POST", "AK", "/clusters", nil, request.Headers, tea.ToMap(request.Body), runtime)
	if _err != nil {
		return nil, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ScaleCluster(clusterId string, request *ScaleClusterRequest, runtime *util.RuntimeOptions) (_result *ScaleClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return
	}
	_result = &ScaleClusterResponse{}
	_body, _err := client.DoRequest("2015-12-15", "HTTPS", "PUT", "AK", "/clusters/"+tea.ToString(clusterId), nil, request.Headers, nil, runtime)
	if _err != nil {
		return nil, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeClusters(request *DescribeClustersRequest, runtime *util.RuntimeOptions) (_result *DescribeClustersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return
	}
	_result = &DescribeClustersResponse{}
	_body, _err := client.DoRequest("2015-12-15", "HTTPS", "GET", "AK", "/clusters", util.StringifyMapValue(tea.ToMap(request.Query)), request.Headers, nil, runtime)
	if _err != nil {
		return nil, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeClusterDetail(clusterId string, request *DescribeClusterDetailRequest, runtime *util.RuntimeOptions) (_result *DescribeClusterDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return
	}
	_result = &DescribeClusterDetailResponse{}
	_body, _err := client.DoRequest("2015-12-15", "HTTPS", "GET", "AK", "/clusters/"+tea.ToString(clusterId), nil, request.Headers, nil, runtime)
	if _err != nil {
		return nil, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteCluster(clusterId string, request *DeleteClusterRequest, runtime *util.RuntimeOptions) (_result *DeleteClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return
	}
	_result = &DeleteClusterResponse{}
	_body, _err := client.DoRequest("2015-12-15", "HTTPS", "DELETE", "AK", "/clusters/"+tea.ToString(clusterId), nil, request.Headers, nil, runtime)
	if _err != nil {
		return nil, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetEndpoint(productId string, regionId string, endpointRule string, network string, suffix string, endpointMap map[string]string, endpoint string) (_result string, _err error) {
	if !util.Empty(endpoint) {
		_result = endpoint
		return _result, _err
	}

	if !util.Empty(endpointMap[regionId]) {
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return "", _err
	}
	_result = _body
	return _result, _err
}