// This file is auto-generated, don't edit it. Thanks.
package client

import (
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	rpc "github.com/alibabacloud-go/tea-rpc/client"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type DescribeDBClusterAuditLogCollectorRequest struct {
	AccessKeyId          *string `json:"AccessKeyId" xml:"AccessKeyId"`
	OwnerId              *int64  `json:"OwnerId" xml:"OwnerId"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount" xml:"ResourceOwnerAccount"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId" xml:"ResourceOwnerId"`
	DBClusterId          *string `json:"DBClusterId" xml:"DBClusterId" require:"true"`
	OwnerAccount         *string `json:"OwnerAccount" xml:"OwnerAccount"`
}

func (s DescribeDBClusterAuditLogCollectorRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterAuditLogCollectorRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterAuditLogCollectorRequest) SetAccessKeyId(v string) *DescribeDBClusterAuditLogCollectorRequest {
	s.AccessKeyId = &v
	return s
}

func (s *DescribeDBClusterAuditLogCollectorRequest) SetOwnerId(v int64) *DescribeDBClusterAuditLogCollectorRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDBClusterAuditLogCollectorRequest) SetResourceOwnerAccount(v string) *DescribeDBClusterAuditLogCollectorRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDBClusterAuditLogCollectorRequest) SetResourceOwnerId(v int64) *DescribeDBClusterAuditLogCollectorRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDBClusterAuditLogCollectorRequest) SetDBClusterId(v string) *DescribeDBClusterAuditLogCollectorRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDBClusterAuditLogCollectorRequest) SetOwnerAccount(v string) *DescribeDBClusterAuditLogCollectorRequest {
	s.OwnerAccount = &v
	return s
}

type DescribeDBClusterAuditLogCollectorResponse struct {
	CollectorStatus *string `json:"CollectorStatus" xml:"CollectorStatus" require:"true"`
	RequestId       *string `json:"RequestId" xml:"RequestId" require:"true"`
}

func (s DescribeDBClusterAuditLogCollectorResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterAuditLogCollectorResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterAuditLogCollectorResponse) SetCollectorStatus(v string) *DescribeDBClusterAuditLogCollectorResponse {
	s.CollectorStatus = &v
	return s
}

func (s *DescribeDBClusterAuditLogCollectorResponse) SetRequestId(v string) *DescribeDBClusterAuditLogCollectorResponse {
	s.RequestId = &v
	return s
}

type ModifyDBClusterAuditLogCollectorRequest struct {
	AccessKeyId          *string `json:"AccessKeyId" xml:"AccessKeyId"`
	OwnerId              *int64  `json:"OwnerId" xml:"OwnerId"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount" xml:"ResourceOwnerAccount"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId" xml:"ResourceOwnerId"`
	DBClusterId          *string `json:"DBClusterId" xml:"DBClusterId" require:"true"`
	CollectorStatus      *string `json:"CollectorStatus" xml:"CollectorStatus" require:"true"`
	OwnerAccount         *string `json:"OwnerAccount" xml:"OwnerAccount"`
}

func (s ModifyDBClusterAuditLogCollectorRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBClusterAuditLogCollectorRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterAuditLogCollectorRequest) SetAccessKeyId(v string) *ModifyDBClusterAuditLogCollectorRequest {
	s.AccessKeyId = &v
	return s
}

func (s *ModifyDBClusterAuditLogCollectorRequest) SetOwnerId(v int64) *ModifyDBClusterAuditLogCollectorRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDBClusterAuditLogCollectorRequest) SetResourceOwnerAccount(v string) *ModifyDBClusterAuditLogCollectorRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDBClusterAuditLogCollectorRequest) SetResourceOwnerId(v int64) *ModifyDBClusterAuditLogCollectorRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDBClusterAuditLogCollectorRequest) SetDBClusterId(v string) *ModifyDBClusterAuditLogCollectorRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyDBClusterAuditLogCollectorRequest) SetCollectorStatus(v string) *ModifyDBClusterAuditLogCollectorRequest {
	s.CollectorStatus = &v
	return s
}

func (s *ModifyDBClusterAuditLogCollectorRequest) SetOwnerAccount(v string) *ModifyDBClusterAuditLogCollectorRequest {
	s.OwnerAccount = &v
	return s
}

type ModifyDBClusterAuditLogCollectorResponse struct {
	RequestId *string `json:"RequestId" xml:"RequestId" require:"true"`
}

func (s ModifyDBClusterAuditLogCollectorResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBClusterAuditLogCollectorResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterAuditLogCollectorResponse) SetRequestId(v string) *ModifyDBClusterAuditLogCollectorResponse {
	s.RequestId = &v
	return s
}

type DescribeDetachedBackupsRequest struct {
	AccessKeyId          *string `json:"AccessKeyId" xml:"AccessKeyId"`
	OwnerId              *int64  `json:"OwnerId" xml:"OwnerId"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount" xml:"ResourceOwnerAccount"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId" xml:"ResourceOwnerId"`
	OwnerAccount         *string `json:"OwnerAccount" xml:"OwnerAccount"`
	DBClusterId          *string `json:"DBClusterId" xml:"DBClusterId" require:"true"`
	BackupId             *string `json:"BackupId" xml:"BackupId"`
	BackupStatus         *string `json:"BackupStatus" xml:"BackupStatus"`
	BackupMode           *string `json:"BackupMode" xml:"BackupMode"`
	StartTime            *string `json:"StartTime" xml:"StartTime" require:"true"`
	EndTime              *string `json:"EndTime" xml:"EndTime" require:"true"`
	PageSize             *int    `json:"PageSize" xml:"PageSize"`
	PageNumber           *int    `json:"PageNumber" xml:"PageNumber"`
}

func (s DescribeDetachedBackupsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDetachedBackupsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDetachedBackupsRequest) SetAccessKeyId(v string) *DescribeDetachedBackupsRequest {
	s.AccessKeyId = &v
	return s
}

func (s *DescribeDetachedBackupsRequest) SetOwnerId(v int64) *DescribeDetachedBackupsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDetachedBackupsRequest) SetResourceOwnerAccount(v string) *DescribeDetachedBackupsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDetachedBackupsRequest) SetResourceOwnerId(v int64) *DescribeDetachedBackupsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDetachedBackupsRequest) SetOwnerAccount(v string) *DescribeDetachedBackupsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDetachedBackupsRequest) SetDBClusterId(v string) *DescribeDetachedBackupsRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDetachedBackupsRequest) SetBackupId(v string) *DescribeDetachedBackupsRequest {
	s.BackupId = &v
	return s
}

func (s *DescribeDetachedBackupsRequest) SetBackupStatus(v string) *DescribeDetachedBackupsRequest {
	s.BackupStatus = &v
	return s
}

func (s *DescribeDetachedBackupsRequest) SetBackupMode(v string) *DescribeDetachedBackupsRequest {
	s.BackupMode = &v
	return s
}

func (s *DescribeDetachedBackupsRequest) SetStartTime(v string) *DescribeDetachedBackupsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDetachedBackupsRequest) SetEndTime(v string) *DescribeDetachedBackupsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDetachedBackupsRequest) SetPageSize(v int) *DescribeDetachedBackupsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDetachedBackupsRequest) SetPageNumber(v int) *DescribeDetachedBackupsRequest {
	s.PageNumber = &v
	return s
}

type DescribeDetachedBackupsResponse struct {
	RequestId        *string                               `json:"RequestId" xml:"RequestId" require:"true"`
	TotalRecordCount *string                               `json:"TotalRecordCount" xml:"TotalRecordCount" require:"true"`
	PageNumber       *string                               `json:"PageNumber" xml:"PageNumber" require:"true"`
	PageRecordCount  *string                               `json:"PageRecordCount" xml:"PageRecordCount" require:"true"`
	Items            *DescribeDetachedBackupsResponseItems `json:"Items" xml:"Items" require:"true" type:"Struct"`
}

func (s DescribeDetachedBackupsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDetachedBackupsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDetachedBackupsResponse) SetRequestId(v string) *DescribeDetachedBackupsResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeDetachedBackupsResponse) SetTotalRecordCount(v string) *DescribeDetachedBackupsResponse {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeDetachedBackupsResponse) SetPageNumber(v string) *DescribeDetachedBackupsResponse {
	s.PageNumber = &v
	return s
}

func (s *DescribeDetachedBackupsResponse) SetPageRecordCount(v string) *DescribeDetachedBackupsResponse {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeDetachedBackupsResponse) SetItems(v *DescribeDetachedBackupsResponseItems) *DescribeDetachedBackupsResponse {
	s.Items = v
	return s
}

type DescribeDetachedBackupsResponseItems struct {
	Backup []*DescribeDetachedBackupsResponseItemsBackup `json:"Backup" xml:"Backup" require:"true" type:"Repeated"`
}

func (s DescribeDetachedBackupsResponseItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeDetachedBackupsResponseItems) GoString() string {
	return s.String()
}

func (s *DescribeDetachedBackupsResponseItems) SetBackup(v []*DescribeDetachedBackupsResponseItemsBackup) *DescribeDetachedBackupsResponseItems {
	s.Backup = v
	return s
}

type DescribeDetachedBackupsResponseItemsBackup struct {
	BackupId        *string `json:"BackupId" xml:"BackupId" require:"true"`
	DBClusterId     *string `json:"DBClusterId" xml:"DBClusterId" require:"true"`
	BackupStatus    *string `json:"BackupStatus" xml:"BackupStatus" require:"true"`
	BackupStartTime *string `json:"BackupStartTime" xml:"BackupStartTime" require:"true"`
	BackupEndTime   *string `json:"BackupEndTime" xml:"BackupEndTime" require:"true"`
	BackupType      *string `json:"BackupType" xml:"BackupType" require:"true"`
	BackupMode      *string `json:"BackupMode" xml:"BackupMode" require:"true"`
	BackupMethod    *string `json:"BackupMethod" xml:"BackupMethod" require:"true"`
	StoreStatus     *string `json:"StoreStatus" xml:"StoreStatus" require:"true"`
	BackupSetSize   *string `json:"BackupSetSize" xml:"BackupSetSize" require:"true"`
	ConsistentTime  *string `json:"ConsistentTime" xml:"ConsistentTime" require:"true"`
	BackupsLevel    *string `json:"BackupsLevel" xml:"BackupsLevel" require:"true"`
	IsAvail         *string `json:"IsAvail" xml:"IsAvail" require:"true"`
}

func (s DescribeDetachedBackupsResponseItemsBackup) String() string {
	return tea.Prettify(s)
}

func (s DescribeDetachedBackupsResponseItemsBackup) GoString() string {
	return s.String()
}

func (s *DescribeDetachedBackupsResponseItemsBackup) SetBackupId(v string) *DescribeDetachedBackupsResponseItemsBackup {
	s.BackupId = &v
	return s
}

func (s *DescribeDetachedBackupsResponseItemsBackup) SetDBClusterId(v string) *DescribeDetachedBackupsResponseItemsBackup {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDetachedBackupsResponseItemsBackup) SetBackupStatus(v string) *DescribeDetachedBackupsResponseItemsBackup {
	s.BackupStatus = &v
	return s
}

func (s *DescribeDetachedBackupsResponseItemsBackup) SetBackupStartTime(v string) *DescribeDetachedBackupsResponseItemsBackup {
	s.BackupStartTime = &v
	return s
}

func (s *DescribeDetachedBackupsResponseItemsBackup) SetBackupEndTime(v string) *DescribeDetachedBackupsResponseItemsBackup {
	s.BackupEndTime = &v
	return s
}

func (s *DescribeDetachedBackupsResponseItemsBackup) SetBackupType(v string) *DescribeDetachedBackupsResponseItemsBackup {
	s.BackupType = &v
	return s
}

func (s *DescribeDetachedBackupsResponseItemsBackup) SetBackupMode(v string) *DescribeDetachedBackupsResponseItemsBackup {
	s.BackupMode = &v
	return s
}

func (s *DescribeDetachedBackupsResponseItemsBackup) SetBackupMethod(v string) *DescribeDetachedBackupsResponseItemsBackup {
	s.BackupMethod = &v
	return s
}

func (s *DescribeDetachedBackupsResponseItemsBackup) SetStoreStatus(v string) *DescribeDetachedBackupsResponseItemsBackup {
	s.StoreStatus = &v
	return s
}

func (s *DescribeDetachedBackupsResponseItemsBackup) SetBackupSetSize(v string) *DescribeDetachedBackupsResponseItemsBackup {
	s.BackupSetSize = &v
	return s
}

func (s *DescribeDetachedBackupsResponseItemsBackup) SetConsistentTime(v string) *DescribeDetachedBackupsResponseItemsBackup {
	s.ConsistentTime = &v
	return s
}

func (s *DescribeDetachedBackupsResponseItemsBackup) SetBackupsLevel(v string) *DescribeDetachedBackupsResponseItemsBackup {
	s.BackupsLevel = &v
	return s
}

func (s *DescribeDetachedBackupsResponseItemsBackup) SetIsAvail(v string) *DescribeDetachedBackupsResponseItemsBackup {
	s.IsAvail = &v
	return s
}

type DescribeDBClustersWithBackupsRequest struct {
	AccessKeyId          *string `json:"AccessKeyId" xml:"AccessKeyId"`
	OwnerId              *int64  `json:"OwnerId" xml:"OwnerId"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount" xml:"ResourceOwnerAccount"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId" xml:"ResourceOwnerId"`
	OwnerAccount         *string `json:"OwnerAccount" xml:"OwnerAccount"`
	RegionId             *string `json:"RegionId" xml:"RegionId" require:"true"`
	DBClusterIds         *string `json:"DBClusterIds" xml:"DBClusterIds"`
	DBClusterDescription *string `json:"DBClusterDescription" xml:"DBClusterDescription"`
	DBType               *string `json:"DBType" xml:"DBType"`
	IsDeleted            *int    `json:"IsDeleted" xml:"IsDeleted"`
	PageSize             *int    `json:"PageSize" xml:"PageSize"`
	PageNumber           *int    `json:"PageNumber" xml:"PageNumber"`
	DBVersion            *string `json:"DBVersion" xml:"DBVersion"`
}

func (s DescribeDBClustersWithBackupsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClustersWithBackupsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBClustersWithBackupsRequest) SetAccessKeyId(v string) *DescribeDBClustersWithBackupsRequest {
	s.AccessKeyId = &v
	return s
}

func (s *DescribeDBClustersWithBackupsRequest) SetOwnerId(v int64) *DescribeDBClustersWithBackupsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDBClustersWithBackupsRequest) SetResourceOwnerAccount(v string) *DescribeDBClustersWithBackupsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDBClustersWithBackupsRequest) SetResourceOwnerId(v int64) *DescribeDBClustersWithBackupsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDBClustersWithBackupsRequest) SetOwnerAccount(v string) *DescribeDBClustersWithBackupsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDBClustersWithBackupsRequest) SetRegionId(v string) *DescribeDBClustersWithBackupsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDBClustersWithBackupsRequest) SetDBClusterIds(v string) *DescribeDBClustersWithBackupsRequest {
	s.DBClusterIds = &v
	return s
}

func (s *DescribeDBClustersWithBackupsRequest) SetDBClusterDescription(v string) *DescribeDBClustersWithBackupsRequest {
	s.DBClusterDescription = &v
	return s
}

func (s *DescribeDBClustersWithBackupsRequest) SetDBType(v string) *DescribeDBClustersWithBackupsRequest {
	s.DBType = &v
	return s
}

func (s *DescribeDBClustersWithBackupsRequest) SetIsDeleted(v int) *DescribeDBClustersWithBackupsRequest {
	s.IsDeleted = &v
	return s
}

func (s *DescribeDBClustersWithBackupsRequest) SetPageSize(v int) *DescribeDBClustersWithBackupsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDBClustersWithBackupsRequest) SetPageNumber(v int) *DescribeDBClustersWithBackupsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDBClustersWithBackupsRequest) SetDBVersion(v string) *DescribeDBClustersWithBackupsRequest {
	s.DBVersion = &v
	return s
}

type DescribeDBClustersWithBackupsResponse struct {
	RequestId        *string                                     `json:"RequestId" xml:"RequestId" require:"true"`
	PageNumber       *int                                        `json:"PageNumber" xml:"PageNumber" require:"true"`
	TotalRecordCount *int                                        `json:"TotalRecordCount" xml:"TotalRecordCount" require:"true"`
	PageRecordCount  *int                                        `json:"PageRecordCount" xml:"PageRecordCount" require:"true"`
	Items            *DescribeDBClustersWithBackupsResponseItems `json:"Items" xml:"Items" require:"true" type:"Struct"`
}

func (s DescribeDBClustersWithBackupsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClustersWithBackupsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBClustersWithBackupsResponse) SetRequestId(v string) *DescribeDBClustersWithBackupsResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeDBClustersWithBackupsResponse) SetPageNumber(v int) *DescribeDBClustersWithBackupsResponse {
	s.PageNumber = &v
	return s
}

func (s *DescribeDBClustersWithBackupsResponse) SetTotalRecordCount(v int) *DescribeDBClustersWithBackupsResponse {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeDBClustersWithBackupsResponse) SetPageRecordCount(v int) *DescribeDBClustersWithBackupsResponse {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeDBClustersWithBackupsResponse) SetItems(v *DescribeDBClustersWithBackupsResponseItems) *DescribeDBClustersWithBackupsResponse {
	s.Items = v
	return s
}

type DescribeDBClustersWithBackupsResponseItems struct {
	DBCluster []*DescribeDBClustersWithBackupsResponseItemsDBCluster `json:"DBCluster" xml:"DBCluster" require:"true" type:"Repeated"`
}

func (s DescribeDBClustersWithBackupsResponseItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClustersWithBackupsResponseItems) GoString() string {
	return s.String()
}

func (s *DescribeDBClustersWithBackupsResponseItems) SetDBCluster(v []*DescribeDBClustersWithBackupsResponseItemsDBCluster) *DescribeDBClustersWithBackupsResponseItems {
	s.DBCluster = v
	return s
}

type DescribeDBClustersWithBackupsResponseItemsDBCluster struct {
	DBClusterId          *string `json:"DBClusterId" xml:"DBClusterId" require:"true"`
	DBClusterDescription *string `json:"DBClusterDescription" xml:"DBClusterDescription" require:"true"`
	PayType              *string `json:"PayType" xml:"PayType" require:"true"`
	DBClusterNetworkType *string `json:"DBClusterNetworkType" xml:"DBClusterNetworkType" require:"true"`
	RegionId             *string `json:"RegionId" xml:"RegionId" require:"true"`
	ZoneId               *string `json:"ZoneId" xml:"ZoneId" require:"true"`
	ExpireTime           *string `json:"ExpireTime" xml:"ExpireTime" require:"true"`
	Expired              *string `json:"Expired" xml:"Expired" require:"true"`
	DBClusterStatus      *string `json:"DBClusterStatus" xml:"DBClusterStatus" require:"true"`
	Engine               *string `json:"Engine" xml:"Engine" require:"true"`
	DBType               *string `json:"DBType" xml:"DBType" require:"true"`
	DBVersion            *string `json:"DBVersion" xml:"DBVersion" require:"true"`
	LockMode             *string `json:"LockMode" xml:"LockMode" require:"true"`
	DeletionLock         *int    `json:"DeletionLock" xml:"DeletionLock" require:"true"`
	CreateTime           *string `json:"CreateTime" xml:"CreateTime" require:"true"`
	VpcId                *string `json:"VpcId" xml:"VpcId" require:"true"`
	IsDeleted            *int    `json:"IsDeleted" xml:"IsDeleted" require:"true"`
	DeletedTime          *string `json:"DeletedTime" xml:"DeletedTime" require:"true"`
	DBNodeClass          *string `json:"DBNodeClass" xml:"DBNodeClass" require:"true"`
}

func (s DescribeDBClustersWithBackupsResponseItemsDBCluster) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClustersWithBackupsResponseItemsDBCluster) GoString() string {
	return s.String()
}

func (s *DescribeDBClustersWithBackupsResponseItemsDBCluster) SetDBClusterId(v string) *DescribeDBClustersWithBackupsResponseItemsDBCluster {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDBClustersWithBackupsResponseItemsDBCluster) SetDBClusterDescription(v string) *DescribeDBClustersWithBackupsResponseItemsDBCluster {
	s.DBClusterDescription = &v
	return s
}

func (s *DescribeDBClustersWithBackupsResponseItemsDBCluster) SetPayType(v string) *DescribeDBClustersWithBackupsResponseItemsDBCluster {
	s.PayType = &v
	return s
}

func (s *DescribeDBClustersWithBackupsResponseItemsDBCluster) SetDBClusterNetworkType(v string) *DescribeDBClustersWithBackupsResponseItemsDBCluster {
	s.DBClusterNetworkType = &v
	return s
}

func (s *DescribeDBClustersWithBackupsResponseItemsDBCluster) SetRegionId(v string) *DescribeDBClustersWithBackupsResponseItemsDBCluster {
	s.RegionId = &v
	return s
}

func (s *DescribeDBClustersWithBackupsResponseItemsDBCluster) SetZoneId(v string) *DescribeDBClustersWithBackupsResponseItemsDBCluster {
	s.ZoneId = &v
	return s
}

func (s *DescribeDBClustersWithBackupsResponseItemsDBCluster) SetExpireTime(v string) *DescribeDBClustersWithBackupsResponseItemsDBCluster {
	s.ExpireTime = &v
	return s
}

func (s *DescribeDBClustersWithBackupsResponseItemsDBCluster) SetExpired(v string) *DescribeDBClustersWithBackupsResponseItemsDBCluster {
	s.Expired = &v
	return s
}

func (s *DescribeDBClustersWithBackupsResponseItemsDBCluster) SetDBClusterStatus(v string) *DescribeDBClustersWithBackupsResponseItemsDBCluster {
	s.DBClusterStatus = &v
	return s
}

func (s *DescribeDBClustersWithBackupsResponseItemsDBCluster) SetEngine(v string) *DescribeDBClustersWithBackupsResponseItemsDBCluster {
	s.Engine = &v
	return s
}

func (s *DescribeDBClustersWithBackupsResponseItemsDBCluster) SetDBType(v string) *DescribeDBClustersWithBackupsResponseItemsDBCluster {
	s.DBType = &v
	return s
}

func (s *DescribeDBClustersWithBackupsResponseItemsDBCluster) SetDBVersion(v string) *DescribeDBClustersWithBackupsResponseItemsDBCluster {
	s.DBVersion = &v
	return s
}

func (s *DescribeDBClustersWithBackupsResponseItemsDBCluster) SetLockMode(v string) *DescribeDBClustersWithBackupsResponseItemsDBCluster {
	s.LockMode = &v
	return s
}

func (s *DescribeDBClustersWithBackupsResponseItemsDBCluster) SetDeletionLock(v int) *DescribeDBClustersWithBackupsResponseItemsDBCluster {
	s.DeletionLock = &v
	return s
}

func (s *DescribeDBClustersWithBackupsResponseItemsDBCluster) SetCreateTime(v string) *DescribeDBClustersWithBackupsResponseItemsDBCluster {
	s.CreateTime = &v
	return s
}

func (s *DescribeDBClustersWithBackupsResponseItemsDBCluster) SetVpcId(v string) *DescribeDBClustersWithBackupsResponseItemsDBCluster {
	s.VpcId = &v
	return s
}

func (s *DescribeDBClustersWithBackupsResponseItemsDBCluster) SetIsDeleted(v int) *DescribeDBClustersWithBackupsResponseItemsDBCluster {
	s.IsDeleted = &v
	return s
}

func (s *DescribeDBClustersWithBackupsResponseItemsDBCluster) SetDeletedTime(v string) *DescribeDBClustersWithBackupsResponseItemsDBCluster {
	s.DeletedTime = &v
	return s
}

func (s *DescribeDBClustersWithBackupsResponseItemsDBCluster) SetDBNodeClass(v string) *DescribeDBClustersWithBackupsResponseItemsDBCluster {
	s.DBNodeClass = &v
	return s
}

type DescribeLogBackupPolicyRequest struct {
	AccessKeyId          *string `json:"AccessKeyId" xml:"AccessKeyId"`
	OwnerId              *int64  `json:"OwnerId" xml:"OwnerId"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount" xml:"ResourceOwnerAccount"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId" xml:"ResourceOwnerId"`
	OwnerAccount         *string `json:"OwnerAccount" xml:"OwnerAccount"`
	DBClusterId          *string `json:"DBClusterId" xml:"DBClusterId" require:"true"`
}

func (s DescribeLogBackupPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeLogBackupPolicyRequest) GoString() string {
	return s.String()
}

func (s *DescribeLogBackupPolicyRequest) SetAccessKeyId(v string) *DescribeLogBackupPolicyRequest {
	s.AccessKeyId = &v
	return s
}

func (s *DescribeLogBackupPolicyRequest) SetOwnerId(v int64) *DescribeLogBackupPolicyRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLogBackupPolicyRequest) SetResourceOwnerAccount(v string) *DescribeLogBackupPolicyRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeLogBackupPolicyRequest) SetResourceOwnerId(v int64) *DescribeLogBackupPolicyRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeLogBackupPolicyRequest) SetOwnerAccount(v string) *DescribeLogBackupPolicyRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeLogBackupPolicyRequest) SetDBClusterId(v string) *DescribeLogBackupPolicyRequest {
	s.DBClusterId = &v
	return s
}

type DescribeLogBackupPolicyResponse struct {
	RequestId                *string `json:"RequestId" xml:"RequestId" require:"true"`
	EnableBackupLog          *int    `json:"EnableBackupLog" xml:"EnableBackupLog" require:"true"`
	LogBackupRetentionPeriod *int    `json:"LogBackupRetentionPeriod" xml:"LogBackupRetentionPeriod" require:"true"`
}

func (s DescribeLogBackupPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeLogBackupPolicyResponse) GoString() string {
	return s.String()
}

func (s *DescribeLogBackupPolicyResponse) SetRequestId(v string) *DescribeLogBackupPolicyResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeLogBackupPolicyResponse) SetEnableBackupLog(v int) *DescribeLogBackupPolicyResponse {
	s.EnableBackupLog = &v
	return s
}

func (s *DescribeLogBackupPolicyResponse) SetLogBackupRetentionPeriod(v int) *DescribeLogBackupPolicyResponse {
	s.LogBackupRetentionPeriod = &v
	return s
}

type ModifyLogBackupPolicyRequest struct {
	AccessKeyId              *string `json:"AccessKeyId" xml:"AccessKeyId"`
	OwnerId                  *int64  `json:"OwnerId" xml:"OwnerId"`
	ResourceOwnerAccount     *string `json:"ResourceOwnerAccount" xml:"ResourceOwnerAccount"`
	ResourceOwnerId          *int64  `json:"ResourceOwnerId" xml:"ResourceOwnerId"`
	OwnerAccount             *string `json:"OwnerAccount" xml:"OwnerAccount"`
	DBClusterId              *string `json:"DBClusterId" xml:"DBClusterId" require:"true"`
	LogBackupRetentionPeriod *string `json:"LogBackupRetentionPeriod" xml:"LogBackupRetentionPeriod"`
}

func (s ModifyLogBackupPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyLogBackupPolicyRequest) GoString() string {
	return s.String()
}

func (s *ModifyLogBackupPolicyRequest) SetAccessKeyId(v string) *ModifyLogBackupPolicyRequest {
	s.AccessKeyId = &v
	return s
}

func (s *ModifyLogBackupPolicyRequest) SetOwnerId(v int64) *ModifyLogBackupPolicyRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyLogBackupPolicyRequest) SetResourceOwnerAccount(v string) *ModifyLogBackupPolicyRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyLogBackupPolicyRequest) SetResourceOwnerId(v int64) *ModifyLogBackupPolicyRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyLogBackupPolicyRequest) SetOwnerAccount(v string) *ModifyLogBackupPolicyRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyLogBackupPolicyRequest) SetDBClusterId(v string) *ModifyLogBackupPolicyRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyLogBackupPolicyRequest) SetLogBackupRetentionPeriod(v string) *ModifyLogBackupPolicyRequest {
	s.LogBackupRetentionPeriod = &v
	return s
}

type ModifyLogBackupPolicyResponse struct {
	RequestId *string `json:"RequestId" xml:"RequestId" require:"true"`
}

func (s ModifyLogBackupPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyLogBackupPolicyResponse) GoString() string {
	return s.String()
}

func (s *ModifyLogBackupPolicyResponse) SetRequestId(v string) *ModifyLogBackupPolicyResponse {
	s.RequestId = &v
	return s
}

type ModifyDBClusterMonitorRequest struct {
	AccessKeyId          *string `json:"AccessKeyId" xml:"AccessKeyId"`
	OwnerId              *int64  `json:"OwnerId" xml:"OwnerId"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount" xml:"ResourceOwnerAccount"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId" xml:"ResourceOwnerId"`
	DBClusterId          *string `json:"DBClusterId" xml:"DBClusterId" require:"true"`
	Period               *string `json:"Period" xml:"Period" require:"true"`
	OwnerAccount         *string `json:"OwnerAccount" xml:"OwnerAccount"`
}

func (s ModifyDBClusterMonitorRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBClusterMonitorRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterMonitorRequest) SetAccessKeyId(v string) *ModifyDBClusterMonitorRequest {
	s.AccessKeyId = &v
	return s
}

func (s *ModifyDBClusterMonitorRequest) SetOwnerId(v int64) *ModifyDBClusterMonitorRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDBClusterMonitorRequest) SetResourceOwnerAccount(v string) *ModifyDBClusterMonitorRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDBClusterMonitorRequest) SetResourceOwnerId(v int64) *ModifyDBClusterMonitorRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDBClusterMonitorRequest) SetDBClusterId(v string) *ModifyDBClusterMonitorRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyDBClusterMonitorRequest) SetPeriod(v string) *ModifyDBClusterMonitorRequest {
	s.Period = &v
	return s
}

func (s *ModifyDBClusterMonitorRequest) SetOwnerAccount(v string) *ModifyDBClusterMonitorRequest {
	s.OwnerAccount = &v
	return s
}

type ModifyDBClusterMonitorResponse struct {
	RequestId *string `json:"RequestId" xml:"RequestId" require:"true"`
}

func (s ModifyDBClusterMonitorResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBClusterMonitorResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterMonitorResponse) SetRequestId(v string) *ModifyDBClusterMonitorResponse {
	s.RequestId = &v
	return s
}

type DescribeDBClusterMonitorRequest struct {
	AccessKeyId          *string `json:"AccessKeyId" xml:"AccessKeyId"`
	OwnerId              *int64  `json:"OwnerId" xml:"OwnerId"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount" xml:"ResourceOwnerAccount"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId" xml:"ResourceOwnerId"`
	DBClusterId          *string `json:"DBClusterId" xml:"DBClusterId" require:"true"`
	OwnerAccount         *string `json:"OwnerAccount" xml:"OwnerAccount"`
}

func (s DescribeDBClusterMonitorRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterMonitorRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterMonitorRequest) SetAccessKeyId(v string) *DescribeDBClusterMonitorRequest {
	s.AccessKeyId = &v
	return s
}

func (s *DescribeDBClusterMonitorRequest) SetOwnerId(v int64) *DescribeDBClusterMonitorRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDBClusterMonitorRequest) SetResourceOwnerAccount(v string) *DescribeDBClusterMonitorRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDBClusterMonitorRequest) SetResourceOwnerId(v int64) *DescribeDBClusterMonitorRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDBClusterMonitorRequest) SetDBClusterId(v string) *DescribeDBClusterMonitorRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDBClusterMonitorRequest) SetOwnerAccount(v string) *DescribeDBClusterMonitorRequest {
	s.OwnerAccount = &v
	return s
}

type DescribeDBClusterMonitorResponse struct {
	Period    *string `json:"Period" xml:"Period" require:"true"`
	RequestId *string `json:"RequestId" xml:"RequestId" require:"true"`
}

func (s DescribeDBClusterMonitorResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterMonitorResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterMonitorResponse) SetPeriod(v string) *DescribeDBClusterMonitorResponse {
	s.Period = &v
	return s
}

func (s *DescribeDBClusterMonitorResponse) SetRequestId(v string) *DescribeDBClusterMonitorResponse {
	s.RequestId = &v
	return s
}

type DescribeDBClusterAvailableResourcesRequest struct {
	AccessKeyId          *string `json:"AccessKeyId" xml:"AccessKeyId"`
	OwnerId              *int64  `json:"OwnerId" xml:"OwnerId"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount" xml:"ResourceOwnerAccount"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId" xml:"ResourceOwnerId"`
	OwnerAccount         *string `json:"OwnerAccount" xml:"OwnerAccount"`
	PayType              *string `json:"PayType" xml:"PayType" require:"true"`
	DBType               *string `json:"DBType" xml:"DBType"`
	DBVersion            *string `json:"DBVersion" xml:"DBVersion"`
	DBNodeClass          *string `json:"DBNodeClass" xml:"DBNodeClass"`
	RegionId             *string `json:"RegionId" xml:"RegionId"`
	ZoneId               *string `json:"ZoneId" xml:"ZoneId"`
}

func (s DescribeDBClusterAvailableResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterAvailableResourcesRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterAvailableResourcesRequest) SetAccessKeyId(v string) *DescribeDBClusterAvailableResourcesRequest {
	s.AccessKeyId = &v
	return s
}

func (s *DescribeDBClusterAvailableResourcesRequest) SetOwnerId(v int64) *DescribeDBClusterAvailableResourcesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDBClusterAvailableResourcesRequest) SetResourceOwnerAccount(v string) *DescribeDBClusterAvailableResourcesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDBClusterAvailableResourcesRequest) SetResourceOwnerId(v int64) *DescribeDBClusterAvailableResourcesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDBClusterAvailableResourcesRequest) SetOwnerAccount(v string) *DescribeDBClusterAvailableResourcesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDBClusterAvailableResourcesRequest) SetPayType(v string) *DescribeDBClusterAvailableResourcesRequest {
	s.PayType = &v
	return s
}

func (s *DescribeDBClusterAvailableResourcesRequest) SetDBType(v string) *DescribeDBClusterAvailableResourcesRequest {
	s.DBType = &v
	return s
}

func (s *DescribeDBClusterAvailableResourcesRequest) SetDBVersion(v string) *DescribeDBClusterAvailableResourcesRequest {
	s.DBVersion = &v
	return s
}

func (s *DescribeDBClusterAvailableResourcesRequest) SetDBNodeClass(v string) *DescribeDBClusterAvailableResourcesRequest {
	s.DBNodeClass = &v
	return s
}

func (s *DescribeDBClusterAvailableResourcesRequest) SetRegionId(v string) *DescribeDBClusterAvailableResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDBClusterAvailableResourcesRequest) SetZoneId(v string) *DescribeDBClusterAvailableResourcesRequest {
	s.ZoneId = &v
	return s
}

type DescribeDBClusterAvailableResourcesResponse struct {
	RequestId      *string                                                      `json:"RequestId" xml:"RequestId" require:"true"`
	AvailableZones []*DescribeDBClusterAvailableResourcesResponseAvailableZones `json:"AvailableZones" xml:"AvailableZones" require:"true" type:"Repeated"`
}

func (s DescribeDBClusterAvailableResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterAvailableResourcesResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterAvailableResourcesResponse) SetRequestId(v string) *DescribeDBClusterAvailableResourcesResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeDBClusterAvailableResourcesResponse) SetAvailableZones(v []*DescribeDBClusterAvailableResourcesResponseAvailableZones) *DescribeDBClusterAvailableResourcesResponse {
	s.AvailableZones = v
	return s
}

type DescribeDBClusterAvailableResourcesResponseAvailableZones struct {
	RegionId         *string                                                                      `json:"RegionId" xml:"RegionId" require:"true"`
	ZoneId           *string                                                                      `json:"ZoneId" xml:"ZoneId" require:"true"`
	SupportedEngines []*DescribeDBClusterAvailableResourcesResponseAvailableZonesSupportedEngines `json:"SupportedEngines" xml:"SupportedEngines" require:"true" type:"Repeated"`
}

func (s DescribeDBClusterAvailableResourcesResponseAvailableZones) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterAvailableResourcesResponseAvailableZones) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterAvailableResourcesResponseAvailableZones) SetRegionId(v string) *DescribeDBClusterAvailableResourcesResponseAvailableZones {
	s.RegionId = &v
	return s
}

func (s *DescribeDBClusterAvailableResourcesResponseAvailableZones) SetZoneId(v string) *DescribeDBClusterAvailableResourcesResponseAvailableZones {
	s.ZoneId = &v
	return s
}

func (s *DescribeDBClusterAvailableResourcesResponseAvailableZones) SetSupportedEngines(v []*DescribeDBClusterAvailableResourcesResponseAvailableZonesSupportedEngines) *DescribeDBClusterAvailableResourcesResponseAvailableZones {
	s.SupportedEngines = v
	return s
}

type DescribeDBClusterAvailableResourcesResponseAvailableZonesSupportedEngines struct {
	Engine             *string                                                                                        `json:"Engine" xml:"Engine" require:"true"`
	AvailableResources []*DescribeDBClusterAvailableResourcesResponseAvailableZonesSupportedEnginesAvailableResources `json:"AvailableResources" xml:"AvailableResources" require:"true" type:"Repeated"`
}

func (s DescribeDBClusterAvailableResourcesResponseAvailableZonesSupportedEngines) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterAvailableResourcesResponseAvailableZonesSupportedEngines) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterAvailableResourcesResponseAvailableZonesSupportedEngines) SetEngine(v string) *DescribeDBClusterAvailableResourcesResponseAvailableZonesSupportedEngines {
	s.Engine = &v
	return s
}

func (s *DescribeDBClusterAvailableResourcesResponseAvailableZonesSupportedEngines) SetAvailableResources(v []*DescribeDBClusterAvailableResourcesResponseAvailableZonesSupportedEnginesAvailableResources) *DescribeDBClusterAvailableResourcesResponseAvailableZonesSupportedEngines {
	s.AvailableResources = v
	return s
}

type DescribeDBClusterAvailableResourcesResponseAvailableZonesSupportedEnginesAvailableResources struct {
	DBNodeClass *string `json:"DBNodeClass" xml:"DBNodeClass" require:"true"`
	Category    *string `json:"Category" xml:"Category" require:"true"`
}

func (s DescribeDBClusterAvailableResourcesResponseAvailableZonesSupportedEnginesAvailableResources) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterAvailableResourcesResponseAvailableZonesSupportedEnginesAvailableResources) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterAvailableResourcesResponseAvailableZonesSupportedEnginesAvailableResources) SetDBNodeClass(v string) *DescribeDBClusterAvailableResourcesResponseAvailableZonesSupportedEnginesAvailableResources {
	s.DBNodeClass = &v
	return s
}

func (s *DescribeDBClusterAvailableResourcesResponseAvailableZonesSupportedEnginesAvailableResources) SetCategory(v string) *DescribeDBClusterAvailableResourcesResponseAvailableZonesSupportedEnginesAvailableResources {
	s.Category = &v
	return s
}

type DescribeBackupLogsRequest struct {
	AccessKeyId          *string `json:"AccessKeyId" xml:"AccessKeyId"`
	OwnerId              *int64  `json:"OwnerId" xml:"OwnerId"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount" xml:"ResourceOwnerAccount"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId" xml:"ResourceOwnerId"`
	OwnerAccount         *string `json:"OwnerAccount" xml:"OwnerAccount"`
	DBClusterId          *string `json:"DBClusterId" xml:"DBClusterId" require:"true"`
	StartTime            *string `json:"StartTime" xml:"StartTime" require:"true"`
	EndTime              *string `json:"EndTime" xml:"EndTime" require:"true"`
	PageSize             *int    `json:"PageSize" xml:"PageSize"`
	PageNumber           *int    `json:"PageNumber" xml:"PageNumber"`
}

func (s DescribeBackupLogsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupLogsRequest) GoString() string {
	return s.String()
}

func (s *DescribeBackupLogsRequest) SetAccessKeyId(v string) *DescribeBackupLogsRequest {
	s.AccessKeyId = &v
	return s
}

func (s *DescribeBackupLogsRequest) SetOwnerId(v int64) *DescribeBackupLogsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeBackupLogsRequest) SetResourceOwnerAccount(v string) *DescribeBackupLogsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeBackupLogsRequest) SetResourceOwnerId(v int64) *DescribeBackupLogsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeBackupLogsRequest) SetOwnerAccount(v string) *DescribeBackupLogsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeBackupLogsRequest) SetDBClusterId(v string) *DescribeBackupLogsRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeBackupLogsRequest) SetStartTime(v string) *DescribeBackupLogsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeBackupLogsRequest) SetEndTime(v string) *DescribeBackupLogsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeBackupLogsRequest) SetPageSize(v int) *DescribeBackupLogsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeBackupLogsRequest) SetPageNumber(v int) *DescribeBackupLogsRequest {
	s.PageNumber = &v
	return s
}

type DescribeBackupLogsResponse struct {
	RequestId        *string                          `json:"RequestId" xml:"RequestId" require:"true"`
	TotalRecordCount *string                          `json:"TotalRecordCount" xml:"TotalRecordCount" require:"true"`
	PageNumber       *string                          `json:"PageNumber" xml:"PageNumber" require:"true"`
	PageRecordCount  *string                          `json:"PageRecordCount" xml:"PageRecordCount" require:"true"`
	Items            *DescribeBackupLogsResponseItems `json:"Items" xml:"Items" require:"true" type:"Struct"`
}

func (s DescribeBackupLogsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupLogsResponse) GoString() string {
	return s.String()
}

func (s *DescribeBackupLogsResponse) SetRequestId(v string) *DescribeBackupLogsResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeBackupLogsResponse) SetTotalRecordCount(v string) *DescribeBackupLogsResponse {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeBackupLogsResponse) SetPageNumber(v string) *DescribeBackupLogsResponse {
	s.PageNumber = &v
	return s
}

func (s *DescribeBackupLogsResponse) SetPageRecordCount(v string) *DescribeBackupLogsResponse {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeBackupLogsResponse) SetItems(v *DescribeBackupLogsResponseItems) *DescribeBackupLogsResponse {
	s.Items = v
	return s
}

type DescribeBackupLogsResponseItems struct {
	BackupLog []*DescribeBackupLogsResponseItemsBackupLog `json:"BackupLog" xml:"BackupLog" require:"true" type:"Repeated"`
}

func (s DescribeBackupLogsResponseItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupLogsResponseItems) GoString() string {
	return s.String()
}

func (s *DescribeBackupLogsResponseItems) SetBackupLog(v []*DescribeBackupLogsResponseItemsBackupLog) *DescribeBackupLogsResponseItems {
	s.BackupLog = v
	return s
}

type DescribeBackupLogsResponseItemsBackupLog struct {
	BackupLogId          *string `json:"BackupLogId" xml:"BackupLogId" require:"true"`
	BackupLogName        *string `json:"BackupLogName" xml:"BackupLogName" require:"true"`
	BackupLogStartTime   *string `json:"BackupLogStartTime" xml:"BackupLogStartTime" require:"true"`
	BackupLogEndTime     *string `json:"BackupLogEndTime" xml:"BackupLogEndTime" require:"true"`
	BackupLogSize        *string `json:"BackupLogSize" xml:"BackupLogSize" require:"true"`
	DownloadLink         *string `json:"DownloadLink" xml:"DownloadLink" require:"true"`
	IntranetDownloadLink *string `json:"IntranetDownloadLink" xml:"IntranetDownloadLink" require:"true"`
	LinkExpiredTime      *string `json:"LinkExpiredTime" xml:"LinkExpiredTime" require:"true"`
}

func (s DescribeBackupLogsResponseItemsBackupLog) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupLogsResponseItemsBackupLog) GoString() string {
	return s.String()
}

func (s *DescribeBackupLogsResponseItemsBackupLog) SetBackupLogId(v string) *DescribeBackupLogsResponseItemsBackupLog {
	s.BackupLogId = &v
	return s
}

func (s *DescribeBackupLogsResponseItemsBackupLog) SetBackupLogName(v string) *DescribeBackupLogsResponseItemsBackupLog {
	s.BackupLogName = &v
	return s
}

func (s *DescribeBackupLogsResponseItemsBackupLog) SetBackupLogStartTime(v string) *DescribeBackupLogsResponseItemsBackupLog {
	s.BackupLogStartTime = &v
	return s
}

func (s *DescribeBackupLogsResponseItemsBackupLog) SetBackupLogEndTime(v string) *DescribeBackupLogsResponseItemsBackupLog {
	s.BackupLogEndTime = &v
	return s
}

func (s *DescribeBackupLogsResponseItemsBackupLog) SetBackupLogSize(v string) *DescribeBackupLogsResponseItemsBackupLog {
	s.BackupLogSize = &v
	return s
}

func (s *DescribeBackupLogsResponseItemsBackupLog) SetDownloadLink(v string) *DescribeBackupLogsResponseItemsBackupLog {
	s.DownloadLink = &v
	return s
}

func (s *DescribeBackupLogsResponseItemsBackupLog) SetIntranetDownloadLink(v string) *DescribeBackupLogsResponseItemsBackupLog {
	s.IntranetDownloadLink = &v
	return s
}

func (s *DescribeBackupLogsResponseItemsBackupLog) SetLinkExpiredTime(v string) *DescribeBackupLogsResponseItemsBackupLog {
	s.LinkExpiredTime = &v
	return s
}

type ModifyDBClusterSSLRequest struct {
	AccessKeyId          *string `json:"AccessKeyId" xml:"AccessKeyId"`
	OwnerId              *int64  `json:"OwnerId" xml:"OwnerId"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount" xml:"ResourceOwnerAccount"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId" xml:"ResourceOwnerId"`
	OwnerAccount         *string `json:"OwnerAccount" xml:"OwnerAccount"`
	DBClusterId          *string `json:"DBClusterId" xml:"DBClusterId" require:"true"`
	SSLEnabled           *string `json:"SSLEnabled" xml:"SSLEnabled" require:"true"`
	DBEndpointId         *string `json:"DBEndpointId" xml:"DBEndpointId"`
	NetType              *string `json:"NetType" xml:"NetType"`
}

func (s ModifyDBClusterSSLRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBClusterSSLRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterSSLRequest) SetAccessKeyId(v string) *ModifyDBClusterSSLRequest {
	s.AccessKeyId = &v
	return s
}

func (s *ModifyDBClusterSSLRequest) SetOwnerId(v int64) *ModifyDBClusterSSLRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDBClusterSSLRequest) SetResourceOwnerAccount(v string) *ModifyDBClusterSSLRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDBClusterSSLRequest) SetResourceOwnerId(v int64) *ModifyDBClusterSSLRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDBClusterSSLRequest) SetOwnerAccount(v string) *ModifyDBClusterSSLRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyDBClusterSSLRequest) SetDBClusterId(v string) *ModifyDBClusterSSLRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyDBClusterSSLRequest) SetSSLEnabled(v string) *ModifyDBClusterSSLRequest {
	s.SSLEnabled = &v
	return s
}

func (s *ModifyDBClusterSSLRequest) SetDBEndpointId(v string) *ModifyDBClusterSSLRequest {
	s.DBEndpointId = &v
	return s
}

func (s *ModifyDBClusterSSLRequest) SetNetType(v string) *ModifyDBClusterSSLRequest {
	s.NetType = &v
	return s
}

type ModifyDBClusterSSLResponse struct {
	RequestId *string `json:"RequestId" xml:"RequestId" require:"true"`
}

func (s ModifyDBClusterSSLResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBClusterSSLResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterSSLResponse) SetRequestId(v string) *ModifyDBClusterSSLResponse {
	s.RequestId = &v
	return s
}

type DescribeDBClusterSSLRequest struct {
	AccessKeyId          *string `json:"AccessKeyId" xml:"AccessKeyId"`
	OwnerId              *int64  `json:"OwnerId" xml:"OwnerId"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount" xml:"ResourceOwnerAccount"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId" xml:"ResourceOwnerId"`
	OwnerAccount         *string `json:"OwnerAccount" xml:"OwnerAccount"`
	DBClusterId          *string `json:"DBClusterId" xml:"DBClusterId" require:"true"`
}

func (s DescribeDBClusterSSLRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterSSLRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterSSLRequest) SetAccessKeyId(v string) *DescribeDBClusterSSLRequest {
	s.AccessKeyId = &v
	return s
}

func (s *DescribeDBClusterSSLRequest) SetOwnerId(v int64) *DescribeDBClusterSSLRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDBClusterSSLRequest) SetResourceOwnerAccount(v string) *DescribeDBClusterSSLRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDBClusterSSLRequest) SetResourceOwnerId(v int64) *DescribeDBClusterSSLRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDBClusterSSLRequest) SetOwnerAccount(v string) *DescribeDBClusterSSLRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDBClusterSSLRequest) SetDBClusterId(v string) *DescribeDBClusterSSLRequest {
	s.DBClusterId = &v
	return s
}

type DescribeDBClusterSSLResponse struct {
	RequestId *string                              `json:"RequestId" xml:"RequestId" require:"true"`
	Items     []*DescribeDBClusterSSLResponseItems `json:"Items" xml:"Items" require:"true" type:"Repeated"`
}

func (s DescribeDBClusterSSLResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterSSLResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterSSLResponse) SetRequestId(v string) *DescribeDBClusterSSLResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeDBClusterSSLResponse) SetItems(v []*DescribeDBClusterSSLResponseItems) *DescribeDBClusterSSLResponse {
	s.Items = v
	return s
}

type DescribeDBClusterSSLResponseItems struct {
	DBEndpointId        *string `json:"DBEndpointId" xml:"DBEndpointId" require:"true"`
	SSLEnabled          *string `json:"SSLEnabled" xml:"SSLEnabled" require:"true"`
	SSLConnectionString *string `json:"SSLConnectionString" xml:"SSLConnectionString" require:"true"`
	SSLExpireTime       *string `json:"SSLExpireTime" xml:"SSLExpireTime" require:"true"`
}

func (s DescribeDBClusterSSLResponseItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterSSLResponseItems) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterSSLResponseItems) SetDBEndpointId(v string) *DescribeDBClusterSSLResponseItems {
	s.DBEndpointId = &v
	return s
}

func (s *DescribeDBClusterSSLResponseItems) SetSSLEnabled(v string) *DescribeDBClusterSSLResponseItems {
	s.SSLEnabled = &v
	return s
}

func (s *DescribeDBClusterSSLResponseItems) SetSSLConnectionString(v string) *DescribeDBClusterSSLResponseItems {
	s.SSLConnectionString = &v
	return s
}

func (s *DescribeDBClusterSSLResponseItems) SetSSLExpireTime(v string) *DescribeDBClusterSSLResponseItems {
	s.SSLExpireTime = &v
	return s
}

type DescribeDBClusterMigrationRequest struct {
	AccessKeyId          *string `json:"AccessKeyId" xml:"AccessKeyId"`
	OwnerId              *int64  `json:"OwnerId" xml:"OwnerId"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount" xml:"ResourceOwnerAccount"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId" xml:"ResourceOwnerId"`
	OwnerAccount         *string `json:"OwnerAccount" xml:"OwnerAccount"`
	DBClusterId          *string `json:"DBClusterId" xml:"DBClusterId" require:"true"`
}

func (s DescribeDBClusterMigrationRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterMigrationRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterMigrationRequest) SetAccessKeyId(v string) *DescribeDBClusterMigrationRequest {
	s.AccessKeyId = &v
	return s
}

func (s *DescribeDBClusterMigrationRequest) SetOwnerId(v int64) *DescribeDBClusterMigrationRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDBClusterMigrationRequest) SetResourceOwnerAccount(v string) *DescribeDBClusterMigrationRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDBClusterMigrationRequest) SetResourceOwnerId(v int64) *DescribeDBClusterMigrationRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDBClusterMigrationRequest) SetOwnerAccount(v string) *DescribeDBClusterMigrationRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDBClusterMigrationRequest) SetDBClusterId(v string) *DescribeDBClusterMigrationRequest {
	s.DBClusterId = &v
	return s
}

type DescribeDBClusterMigrationResponse struct {
	RequestId              *string                                                    `json:"RequestId" xml:"RequestId" require:"true"`
	DBClusterId            *string                                                    `json:"DBClusterId" xml:"DBClusterId" require:"true"`
	SourceRDSDBInstanceId  *string                                                    `json:"SourceRDSDBInstanceId" xml:"SourceRDSDBInstanceId" require:"true"`
	MigrationStatus        *string                                                    `json:"MigrationStatus" xml:"MigrationStatus" require:"true"`
	Topologies             *string                                                    `json:"Topologies" xml:"Topologies" require:"true"`
	DelayedSeconds         *int                                                       `json:"DelayedSeconds" xml:"DelayedSeconds" require:"true"`
	ExpiredTime            *string                                                    `json:"ExpiredTime" xml:"ExpiredTime" require:"true"`
	RdsReadWriteMode       *string                                                    `json:"RdsReadWriteMode" xml:"RdsReadWriteMode" require:"true"`
	DBClusterReadWriteMode *string                                                    `json:"DBClusterReadWriteMode" xml:"DBClusterReadWriteMode" require:"true"`
	Comment                *string                                                    `json:"Comment" xml:"Comment" require:"true"`
	DBClusterEndpointList  []*DescribeDBClusterMigrationResponseDBClusterEndpointList `json:"DBClusterEndpointList" xml:"DBClusterEndpointList" require:"true" type:"Repeated"`
	RdsEndpointList        []*DescribeDBClusterMigrationResponseRdsEndpointList       `json:"RdsEndpointList" xml:"RdsEndpointList" require:"true" type:"Repeated"`
}

func (s DescribeDBClusterMigrationResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterMigrationResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterMigrationResponse) SetRequestId(v string) *DescribeDBClusterMigrationResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeDBClusterMigrationResponse) SetDBClusterId(v string) *DescribeDBClusterMigrationResponse {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDBClusterMigrationResponse) SetSourceRDSDBInstanceId(v string) *DescribeDBClusterMigrationResponse {
	s.SourceRDSDBInstanceId = &v
	return s
}

func (s *DescribeDBClusterMigrationResponse) SetMigrationStatus(v string) *DescribeDBClusterMigrationResponse {
	s.MigrationStatus = &v
	return s
}

func (s *DescribeDBClusterMigrationResponse) SetTopologies(v string) *DescribeDBClusterMigrationResponse {
	s.Topologies = &v
	return s
}

func (s *DescribeDBClusterMigrationResponse) SetDelayedSeconds(v int) *DescribeDBClusterMigrationResponse {
	s.DelayedSeconds = &v
	return s
}

func (s *DescribeDBClusterMigrationResponse) SetExpiredTime(v string) *DescribeDBClusterMigrationResponse {
	s.ExpiredTime = &v
	return s
}

func (s *DescribeDBClusterMigrationResponse) SetRdsReadWriteMode(v string) *DescribeDBClusterMigrationResponse {
	s.RdsReadWriteMode = &v
	return s
}

func (s *DescribeDBClusterMigrationResponse) SetDBClusterReadWriteMode(v string) *DescribeDBClusterMigrationResponse {
	s.DBClusterReadWriteMode = &v
	return s
}

func (s *DescribeDBClusterMigrationResponse) SetComment(v string) *DescribeDBClusterMigrationResponse {
	s.Comment = &v
	return s
}

func (s *DescribeDBClusterMigrationResponse) SetDBClusterEndpointList(v []*DescribeDBClusterMigrationResponseDBClusterEndpointList) *DescribeDBClusterMigrationResponse {
	s.DBClusterEndpointList = v
	return s
}

func (s *DescribeDBClusterMigrationResponse) SetRdsEndpointList(v []*DescribeDBClusterMigrationResponseRdsEndpointList) *DescribeDBClusterMigrationResponse {
	s.RdsEndpointList = v
	return s
}

type DescribeDBClusterMigrationResponseDBClusterEndpointList struct {
	DBEndpointId *string                                                                `json:"DBEndpointId" xml:"DBEndpointId" require:"true"`
	EndpointType *string                                                                `json:"EndpointType" xml:"EndpointType" require:"true"`
	AddressItems []*DescribeDBClusterMigrationResponseDBClusterEndpointListAddressItems `json:"AddressItems" xml:"AddressItems" require:"true" type:"Repeated"`
}

func (s DescribeDBClusterMigrationResponseDBClusterEndpointList) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterMigrationResponseDBClusterEndpointList) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterMigrationResponseDBClusterEndpointList) SetDBEndpointId(v string) *DescribeDBClusterMigrationResponseDBClusterEndpointList {
	s.DBEndpointId = &v
	return s
}

func (s *DescribeDBClusterMigrationResponseDBClusterEndpointList) SetEndpointType(v string) *DescribeDBClusterMigrationResponseDBClusterEndpointList {
	s.EndpointType = &v
	return s
}

func (s *DescribeDBClusterMigrationResponseDBClusterEndpointList) SetAddressItems(v []*DescribeDBClusterMigrationResponseDBClusterEndpointListAddressItems) *DescribeDBClusterMigrationResponseDBClusterEndpointList {
	s.AddressItems = v
	return s
}

type DescribeDBClusterMigrationResponseDBClusterEndpointListAddressItems struct {
	ConnectionString *string `json:"ConnectionString" xml:"ConnectionString" require:"true"`
	IPAddress        *string `json:"IPAddress" xml:"IPAddress" require:"true"`
	NetType          *string `json:"NetType" xml:"NetType" require:"true"`
	Port             *string `json:"Port" xml:"Port" require:"true"`
	VPCId            *string `json:"VPCId" xml:"VPCId" require:"true"`
	VSwitchId        *string `json:"VSwitchId" xml:"VSwitchId" require:"true"`
}

func (s DescribeDBClusterMigrationResponseDBClusterEndpointListAddressItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterMigrationResponseDBClusterEndpointListAddressItems) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterMigrationResponseDBClusterEndpointListAddressItems) SetConnectionString(v string) *DescribeDBClusterMigrationResponseDBClusterEndpointListAddressItems {
	s.ConnectionString = &v
	return s
}

func (s *DescribeDBClusterMigrationResponseDBClusterEndpointListAddressItems) SetIPAddress(v string) *DescribeDBClusterMigrationResponseDBClusterEndpointListAddressItems {
	s.IPAddress = &v
	return s
}

func (s *DescribeDBClusterMigrationResponseDBClusterEndpointListAddressItems) SetNetType(v string) *DescribeDBClusterMigrationResponseDBClusterEndpointListAddressItems {
	s.NetType = &v
	return s
}

func (s *DescribeDBClusterMigrationResponseDBClusterEndpointListAddressItems) SetPort(v string) *DescribeDBClusterMigrationResponseDBClusterEndpointListAddressItems {
	s.Port = &v
	return s
}

func (s *DescribeDBClusterMigrationResponseDBClusterEndpointListAddressItems) SetVPCId(v string) *DescribeDBClusterMigrationResponseDBClusterEndpointListAddressItems {
	s.VPCId = &v
	return s
}

func (s *DescribeDBClusterMigrationResponseDBClusterEndpointListAddressItems) SetVSwitchId(v string) *DescribeDBClusterMigrationResponseDBClusterEndpointListAddressItems {
	s.VSwitchId = &v
	return s
}

type DescribeDBClusterMigrationResponseRdsEndpointList struct {
	DBEndpointId *string                                                          `json:"DBEndpointId" xml:"DBEndpointId" require:"true"`
	EndpointType *string                                                          `json:"EndpointType" xml:"EndpointType" require:"true"`
	AddressItems []*DescribeDBClusterMigrationResponseRdsEndpointListAddressItems `json:"AddressItems" xml:"AddressItems" require:"true" type:"Repeated"`
}

func (s DescribeDBClusterMigrationResponseRdsEndpointList) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterMigrationResponseRdsEndpointList) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterMigrationResponseRdsEndpointList) SetDBEndpointId(v string) *DescribeDBClusterMigrationResponseRdsEndpointList {
	s.DBEndpointId = &v
	return s
}

func (s *DescribeDBClusterMigrationResponseRdsEndpointList) SetEndpointType(v string) *DescribeDBClusterMigrationResponseRdsEndpointList {
	s.EndpointType = &v
	return s
}

func (s *DescribeDBClusterMigrationResponseRdsEndpointList) SetAddressItems(v []*DescribeDBClusterMigrationResponseRdsEndpointListAddressItems) *DescribeDBClusterMigrationResponseRdsEndpointList {
	s.AddressItems = v
	return s
}

type DescribeDBClusterMigrationResponseRdsEndpointListAddressItems struct {
	ConnectionString *string `json:"ConnectionString" xml:"ConnectionString" require:"true"`
	IPAddress        *string `json:"IPAddress" xml:"IPAddress" require:"true"`
	NetType          *string `json:"NetType" xml:"NetType" require:"true"`
	Port             *string `json:"Port" xml:"Port" require:"true"`
	VPCId            *string `json:"VPCId" xml:"VPCId" require:"true"`
	VSwitchId        *string `json:"VSwitchId" xml:"VSwitchId" require:"true"`
}

func (s DescribeDBClusterMigrationResponseRdsEndpointListAddressItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterMigrationResponseRdsEndpointListAddressItems) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterMigrationResponseRdsEndpointListAddressItems) SetConnectionString(v string) *DescribeDBClusterMigrationResponseRdsEndpointListAddressItems {
	s.ConnectionString = &v
	return s
}

func (s *DescribeDBClusterMigrationResponseRdsEndpointListAddressItems) SetIPAddress(v string) *DescribeDBClusterMigrationResponseRdsEndpointListAddressItems {
	s.IPAddress = &v
	return s
}

func (s *DescribeDBClusterMigrationResponseRdsEndpointListAddressItems) SetNetType(v string) *DescribeDBClusterMigrationResponseRdsEndpointListAddressItems {
	s.NetType = &v
	return s
}

func (s *DescribeDBClusterMigrationResponseRdsEndpointListAddressItems) SetPort(v string) *DescribeDBClusterMigrationResponseRdsEndpointListAddressItems {
	s.Port = &v
	return s
}

func (s *DescribeDBClusterMigrationResponseRdsEndpointListAddressItems) SetVPCId(v string) *DescribeDBClusterMigrationResponseRdsEndpointListAddressItems {
	s.VPCId = &v
	return s
}

func (s *DescribeDBClusterMigrationResponseRdsEndpointListAddressItems) SetVSwitchId(v string) *DescribeDBClusterMigrationResponseRdsEndpointListAddressItems {
	s.VSwitchId = &v
	return s
}

type CloseDBClusterMigrationRequest struct {
	AccessKeyId          *string `json:"AccessKeyId" xml:"AccessKeyId"`
	OwnerId              *int64  `json:"OwnerId" xml:"OwnerId"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount" xml:"ResourceOwnerAccount"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId" xml:"ResourceOwnerId"`
	OwnerAccount         *string `json:"OwnerAccount" xml:"OwnerAccount"`
	DBClusterId          *string `json:"DBClusterId" xml:"DBClusterId" require:"true"`
	ContinueEnableBinlog *bool   `json:"ContinueEnableBinlog" xml:"ContinueEnableBinlog"`
}

func (s CloseDBClusterMigrationRequest) String() string {
	return tea.Prettify(s)
}

func (s CloseDBClusterMigrationRequest) GoString() string {
	return s.String()
}

func (s *CloseDBClusterMigrationRequest) SetAccessKeyId(v string) *CloseDBClusterMigrationRequest {
	s.AccessKeyId = &v
	return s
}

func (s *CloseDBClusterMigrationRequest) SetOwnerId(v int64) *CloseDBClusterMigrationRequest {
	s.OwnerId = &v
	return s
}

func (s *CloseDBClusterMigrationRequest) SetResourceOwnerAccount(v string) *CloseDBClusterMigrationRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CloseDBClusterMigrationRequest) SetResourceOwnerId(v int64) *CloseDBClusterMigrationRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CloseDBClusterMigrationRequest) SetOwnerAccount(v string) *CloseDBClusterMigrationRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CloseDBClusterMigrationRequest) SetDBClusterId(v string) *CloseDBClusterMigrationRequest {
	s.DBClusterId = &v
	return s
}

func (s *CloseDBClusterMigrationRequest) SetContinueEnableBinlog(v bool) *CloseDBClusterMigrationRequest {
	s.ContinueEnableBinlog = &v
	return s
}

type CloseDBClusterMigrationResponse struct {
	RequestId *string `json:"RequestId" xml:"RequestId" require:"true"`
}

func (s CloseDBClusterMigrationResponse) String() string {
	return tea.Prettify(s)
}

func (s CloseDBClusterMigrationResponse) GoString() string {
	return s.String()
}

func (s *CloseDBClusterMigrationResponse) SetRequestId(v string) *CloseDBClusterMigrationResponse {
	s.RequestId = &v
	return s
}

type ModifyDBClusterMigrationRequest struct {
	SecurityToken         *string `json:"SecurityToken" xml:"SecurityToken"`
	AccessKeyId           *string `json:"AccessKeyId" xml:"AccessKeyId"`
	OwnerId               *int64  `json:"OwnerId" xml:"OwnerId"`
	ResourceOwnerAccount  *string `json:"ResourceOwnerAccount" xml:"ResourceOwnerAccount"`
	ResourceOwnerId       *int64  `json:"ResourceOwnerId" xml:"ResourceOwnerId"`
	OwnerAccount          *string `json:"OwnerAccount" xml:"OwnerAccount"`
	DBClusterId           *string `json:"DBClusterId" xml:"DBClusterId" require:"true"`
	SourceRDSDBInstanceId *string `json:"SourceRDSDBInstanceId" xml:"SourceRDSDBInstanceId" require:"true"`
	NewMasterInstanceId   *string `json:"NewMasterInstanceId" xml:"NewMasterInstanceId" require:"true"`
	SwapConnectionString  *string `json:"SwapConnectionString" xml:"SwapConnectionString"`
}

func (s ModifyDBClusterMigrationRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBClusterMigrationRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterMigrationRequest) SetSecurityToken(v string) *ModifyDBClusterMigrationRequest {
	s.SecurityToken = &v
	return s
}

func (s *ModifyDBClusterMigrationRequest) SetAccessKeyId(v string) *ModifyDBClusterMigrationRequest {
	s.AccessKeyId = &v
	return s
}

func (s *ModifyDBClusterMigrationRequest) SetOwnerId(v int64) *ModifyDBClusterMigrationRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDBClusterMigrationRequest) SetResourceOwnerAccount(v string) *ModifyDBClusterMigrationRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDBClusterMigrationRequest) SetResourceOwnerId(v int64) *ModifyDBClusterMigrationRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDBClusterMigrationRequest) SetOwnerAccount(v string) *ModifyDBClusterMigrationRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyDBClusterMigrationRequest) SetDBClusterId(v string) *ModifyDBClusterMigrationRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyDBClusterMigrationRequest) SetSourceRDSDBInstanceId(v string) *ModifyDBClusterMigrationRequest {
	s.SourceRDSDBInstanceId = &v
	return s
}

func (s *ModifyDBClusterMigrationRequest) SetNewMasterInstanceId(v string) *ModifyDBClusterMigrationRequest {
	s.NewMasterInstanceId = &v
	return s
}

func (s *ModifyDBClusterMigrationRequest) SetSwapConnectionString(v string) *ModifyDBClusterMigrationRequest {
	s.SwapConnectionString = &v
	return s
}

type ModifyDBClusterMigrationResponse struct {
	RequestId *string `json:"RequestId" xml:"RequestId" require:"true"`
}

func (s ModifyDBClusterMigrationResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBClusterMigrationResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterMigrationResponse) SetRequestId(v string) *ModifyDBClusterMigrationResponse {
	s.RequestId = &v
	return s
}

type ModifyAutoRenewAttributeRequest struct {
	AccessKeyId          *string `json:"AccessKeyId" xml:"AccessKeyId"`
	OwnerId              *int64  `json:"OwnerId" xml:"OwnerId"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount" xml:"ResourceOwnerAccount"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId" xml:"ResourceOwnerId"`
	OwnerAccount         *string `json:"OwnerAccount" xml:"OwnerAccount"`
	DBClusterIds         *string `json:"DBClusterIds" xml:"DBClusterIds" require:"true"`
	RegionId             *string `json:"RegionId" xml:"RegionId" require:"true"`
	RenewalStatus        *string `json:"RenewalStatus" xml:"RenewalStatus"`
	Duration             *string `json:"Duration" xml:"Duration"`
	PeriodUnit           *string `json:"PeriodUnit" xml:"PeriodUnit"`
	ResourceGroupId      *string `json:"ResourceGroupId" xml:"ResourceGroupId"`
}

func (s ModifyAutoRenewAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyAutoRenewAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyAutoRenewAttributeRequest) SetAccessKeyId(v string) *ModifyAutoRenewAttributeRequest {
	s.AccessKeyId = &v
	return s
}

func (s *ModifyAutoRenewAttributeRequest) SetOwnerId(v int64) *ModifyAutoRenewAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyAutoRenewAttributeRequest) SetResourceOwnerAccount(v string) *ModifyAutoRenewAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyAutoRenewAttributeRequest) SetResourceOwnerId(v int64) *ModifyAutoRenewAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyAutoRenewAttributeRequest) SetOwnerAccount(v string) *ModifyAutoRenewAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyAutoRenewAttributeRequest) SetDBClusterIds(v string) *ModifyAutoRenewAttributeRequest {
	s.DBClusterIds = &v
	return s
}

func (s *ModifyAutoRenewAttributeRequest) SetRegionId(v string) *ModifyAutoRenewAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyAutoRenewAttributeRequest) SetRenewalStatus(v string) *ModifyAutoRenewAttributeRequest {
	s.RenewalStatus = &v
	return s
}

func (s *ModifyAutoRenewAttributeRequest) SetDuration(v string) *ModifyAutoRenewAttributeRequest {
	s.Duration = &v
	return s
}

func (s *ModifyAutoRenewAttributeRequest) SetPeriodUnit(v string) *ModifyAutoRenewAttributeRequest {
	s.PeriodUnit = &v
	return s
}

func (s *ModifyAutoRenewAttributeRequest) SetResourceGroupId(v string) *ModifyAutoRenewAttributeRequest {
	s.ResourceGroupId = &v
	return s
}

type ModifyAutoRenewAttributeResponse struct {
	RequestId *string `json:"RequestId" xml:"RequestId" require:"true"`
}

func (s ModifyAutoRenewAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyAutoRenewAttributeResponse) GoString() string {
	return s.String()
}

func (s *ModifyAutoRenewAttributeResponse) SetRequestId(v string) *ModifyAutoRenewAttributeResponse {
	s.RequestId = &v
	return s
}

type ModifyDBNodeClassRequest struct {
	AccessKeyId          *string `json:"AccessKeyId" xml:"AccessKeyId"`
	OwnerId              *int64  `json:"OwnerId" xml:"OwnerId"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount" xml:"ResourceOwnerAccount"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId" xml:"ResourceOwnerId"`
	OwnerAccount         *string `json:"OwnerAccount" xml:"OwnerAccount"`
	DBClusterId          *string `json:"DBClusterId" xml:"DBClusterId" require:"true"`
	ModifyType           *string `json:"ModifyType" xml:"ModifyType" require:"true"`
	DBNodeTargetClass    *string `json:"DBNodeTargetClass" xml:"DBNodeTargetClass" require:"true"`
	ClientToken          *string `json:"ClientToken" xml:"ClientToken"`
}

func (s ModifyDBNodeClassRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBNodeClassRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBNodeClassRequest) SetAccessKeyId(v string) *ModifyDBNodeClassRequest {
	s.AccessKeyId = &v
	return s
}

func (s *ModifyDBNodeClassRequest) SetOwnerId(v int64) *ModifyDBNodeClassRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDBNodeClassRequest) SetResourceOwnerAccount(v string) *ModifyDBNodeClassRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDBNodeClassRequest) SetResourceOwnerId(v int64) *ModifyDBNodeClassRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDBNodeClassRequest) SetOwnerAccount(v string) *ModifyDBNodeClassRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyDBNodeClassRequest) SetDBClusterId(v string) *ModifyDBNodeClassRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyDBNodeClassRequest) SetModifyType(v string) *ModifyDBNodeClassRequest {
	s.ModifyType = &v
	return s
}

func (s *ModifyDBNodeClassRequest) SetDBNodeTargetClass(v string) *ModifyDBNodeClassRequest {
	s.DBNodeTargetClass = &v
	return s
}

func (s *ModifyDBNodeClassRequest) SetClientToken(v string) *ModifyDBNodeClassRequest {
	s.ClientToken = &v
	return s
}

type ModifyDBNodeClassResponse struct {
	RequestId   *string `json:"RequestId" xml:"RequestId" require:"true"`
	DBClusterId *string `json:"DBClusterId" xml:"DBClusterId" require:"true"`
	OrderId     *string `json:"OrderId" xml:"OrderId" require:"true"`
}

func (s ModifyDBNodeClassResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBNodeClassResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBNodeClassResponse) SetRequestId(v string) *ModifyDBNodeClassResponse {
	s.RequestId = &v
	return s
}

func (s *ModifyDBNodeClassResponse) SetDBClusterId(v string) *ModifyDBNodeClassResponse {
	s.DBClusterId = &v
	return s
}

func (s *ModifyDBNodeClassResponse) SetOrderId(v string) *ModifyDBNodeClassResponse {
	s.OrderId = &v
	return s
}

type DescribeAutoRenewAttributeRequest struct {
	AccessKeyId          *string `json:"AccessKeyId" xml:"AccessKeyId"`
	OwnerId              *int64  `json:"OwnerId" xml:"OwnerId"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount" xml:"ResourceOwnerAccount"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId" xml:"ResourceOwnerId"`
	OwnerAccount         *string `json:"OwnerAccount" xml:"OwnerAccount"`
	RegionId             *string `json:"RegionId" xml:"RegionId" require:"true"`
	DBClusterIds         *string `json:"DBClusterIds" xml:"DBClusterIds"`
	PageSize             *int    `json:"PageSize" xml:"PageSize"`
	PageNumber           *int    `json:"PageNumber" xml:"PageNumber"`
	ResourceGroupId      *string `json:"ResourceGroupId" xml:"ResourceGroupId"`
}

func (s DescribeAutoRenewAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAutoRenewAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeAutoRenewAttributeRequest) SetAccessKeyId(v string) *DescribeAutoRenewAttributeRequest {
	s.AccessKeyId = &v
	return s
}

func (s *DescribeAutoRenewAttributeRequest) SetOwnerId(v int64) *DescribeAutoRenewAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeAutoRenewAttributeRequest) SetResourceOwnerAccount(v string) *DescribeAutoRenewAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeAutoRenewAttributeRequest) SetResourceOwnerId(v int64) *DescribeAutoRenewAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeAutoRenewAttributeRequest) SetOwnerAccount(v string) *DescribeAutoRenewAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeAutoRenewAttributeRequest) SetRegionId(v string) *DescribeAutoRenewAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAutoRenewAttributeRequest) SetDBClusterIds(v string) *DescribeAutoRenewAttributeRequest {
	s.DBClusterIds = &v
	return s
}

func (s *DescribeAutoRenewAttributeRequest) SetPageSize(v int) *DescribeAutoRenewAttributeRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAutoRenewAttributeRequest) SetPageNumber(v int) *DescribeAutoRenewAttributeRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeAutoRenewAttributeRequest) SetResourceGroupId(v string) *DescribeAutoRenewAttributeRequest {
	s.ResourceGroupId = &v
	return s
}

type DescribeAutoRenewAttributeResponse struct {
	RequestId        *string                                  `json:"RequestId" xml:"RequestId" require:"true"`
	PageNumber       *int                                     `json:"PageNumber" xml:"PageNumber" require:"true"`
	TotalRecordCount *int                                     `json:"TotalRecordCount" xml:"TotalRecordCount" require:"true"`
	PageRecordCount  *int                                     `json:"PageRecordCount" xml:"PageRecordCount" require:"true"`
	Items            *DescribeAutoRenewAttributeResponseItems `json:"Items" xml:"Items" require:"true" type:"Struct"`
}

func (s DescribeAutoRenewAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAutoRenewAttributeResponse) GoString() string {
	return s.String()
}

func (s *DescribeAutoRenewAttributeResponse) SetRequestId(v string) *DescribeAutoRenewAttributeResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeAutoRenewAttributeResponse) SetPageNumber(v int) *DescribeAutoRenewAttributeResponse {
	s.PageNumber = &v
	return s
}

func (s *DescribeAutoRenewAttributeResponse) SetTotalRecordCount(v int) *DescribeAutoRenewAttributeResponse {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeAutoRenewAttributeResponse) SetPageRecordCount(v int) *DescribeAutoRenewAttributeResponse {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeAutoRenewAttributeResponse) SetItems(v *DescribeAutoRenewAttributeResponseItems) *DescribeAutoRenewAttributeResponse {
	s.Items = v
	return s
}

type DescribeAutoRenewAttributeResponseItems struct {
	AutoRenewAttribute []*DescribeAutoRenewAttributeResponseItemsAutoRenewAttribute `json:"AutoRenewAttribute" xml:"AutoRenewAttribute" require:"true" type:"Repeated"`
}

func (s DescribeAutoRenewAttributeResponseItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeAutoRenewAttributeResponseItems) GoString() string {
	return s.String()
}

func (s *DescribeAutoRenewAttributeResponseItems) SetAutoRenewAttribute(v []*DescribeAutoRenewAttributeResponseItemsAutoRenewAttribute) *DescribeAutoRenewAttributeResponseItems {
	s.AutoRenewAttribute = v
	return s
}

type DescribeAutoRenewAttributeResponseItemsAutoRenewAttribute struct {
	DBClusterId      *string `json:"DBClusterId" xml:"DBClusterId" require:"true"`
	RegionId         *string `json:"RegionId" xml:"RegionId" require:"true"`
	AutoRenewEnabled *bool   `json:"AutoRenewEnabled" xml:"AutoRenewEnabled" require:"true"`
	Duration         *int    `json:"Duration" xml:"Duration" require:"true"`
	PeriodUnit       *string `json:"PeriodUnit" xml:"PeriodUnit" require:"true"`
	RenewalStatus    *string `json:"RenewalStatus" xml:"RenewalStatus" require:"true"`
}

func (s DescribeAutoRenewAttributeResponseItemsAutoRenewAttribute) String() string {
	return tea.Prettify(s)
}

func (s DescribeAutoRenewAttributeResponseItemsAutoRenewAttribute) GoString() string {
	return s.String()
}

func (s *DescribeAutoRenewAttributeResponseItemsAutoRenewAttribute) SetDBClusterId(v string) *DescribeAutoRenewAttributeResponseItemsAutoRenewAttribute {
	s.DBClusterId = &v
	return s
}

func (s *DescribeAutoRenewAttributeResponseItemsAutoRenewAttribute) SetRegionId(v string) *DescribeAutoRenewAttributeResponseItemsAutoRenewAttribute {
	s.RegionId = &v
	return s
}

func (s *DescribeAutoRenewAttributeResponseItemsAutoRenewAttribute) SetAutoRenewEnabled(v bool) *DescribeAutoRenewAttributeResponseItemsAutoRenewAttribute {
	s.AutoRenewEnabled = &v
	return s
}

func (s *DescribeAutoRenewAttributeResponseItemsAutoRenewAttribute) SetDuration(v int) *DescribeAutoRenewAttributeResponseItemsAutoRenewAttribute {
	s.Duration = &v
	return s
}

func (s *DescribeAutoRenewAttributeResponseItemsAutoRenewAttribute) SetPeriodUnit(v string) *DescribeAutoRenewAttributeResponseItemsAutoRenewAttribute {
	s.PeriodUnit = &v
	return s
}

func (s *DescribeAutoRenewAttributeResponseItemsAutoRenewAttribute) SetRenewalStatus(v string) *DescribeAutoRenewAttributeResponseItemsAutoRenewAttribute {
	s.RenewalStatus = &v
	return s
}

type CreateDBNodesRequest struct {
	AccessKeyId          *string                       `json:"AccessKeyId" xml:"AccessKeyId"`
	OwnerId              *int64                        `json:"OwnerId" xml:"OwnerId"`
	ResourceOwnerAccount *string                       `json:"ResourceOwnerAccount" xml:"ResourceOwnerAccount"`
	ResourceOwnerId      *int64                        `json:"ResourceOwnerId" xml:"ResourceOwnerId"`
	OwnerAccount         *string                       `json:"OwnerAccount" xml:"OwnerAccount"`
	DBClusterId          *string                       `json:"DBClusterId" xml:"DBClusterId" require:"true"`
	DBNode               []*CreateDBNodesRequestDBNode `json:"DBNode" xml:"DBNode" require:"true" type:"Repeated"`
	ClientToken          *string                       `json:"ClientToken" xml:"ClientToken"`
}

func (s CreateDBNodesRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDBNodesRequest) GoString() string {
	return s.String()
}

func (s *CreateDBNodesRequest) SetAccessKeyId(v string) *CreateDBNodesRequest {
	s.AccessKeyId = &v
	return s
}

func (s *CreateDBNodesRequest) SetOwnerId(v int64) *CreateDBNodesRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateDBNodesRequest) SetResourceOwnerAccount(v string) *CreateDBNodesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateDBNodesRequest) SetResourceOwnerId(v int64) *CreateDBNodesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateDBNodesRequest) SetOwnerAccount(v string) *CreateDBNodesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateDBNodesRequest) SetDBClusterId(v string) *CreateDBNodesRequest {
	s.DBClusterId = &v
	return s
}

func (s *CreateDBNodesRequest) SetDBNode(v []*CreateDBNodesRequestDBNode) *CreateDBNodesRequest {
	s.DBNode = v
	return s
}

func (s *CreateDBNodesRequest) SetClientToken(v string) *CreateDBNodesRequest {
	s.ClientToken = &v
	return s
}

type CreateDBNodesRequestDBNode struct {
	TargetClass *string `json:"TargetClass" xml:"TargetClass" require:"true"`
	ZoneId      *string `json:"ZoneId" xml:"ZoneId" require:"true"`
}

func (s CreateDBNodesRequestDBNode) String() string {
	return tea.Prettify(s)
}

func (s CreateDBNodesRequestDBNode) GoString() string {
	return s.String()
}

func (s *CreateDBNodesRequestDBNode) SetTargetClass(v string) *CreateDBNodesRequestDBNode {
	s.TargetClass = &v
	return s
}

func (s *CreateDBNodesRequestDBNode) SetZoneId(v string) *CreateDBNodesRequestDBNode {
	s.ZoneId = &v
	return s
}

type CreateDBNodesResponse struct {
	RequestId   *string                         `json:"RequestId" xml:"RequestId" require:"true"`
	DBClusterId *string                         `json:"DBClusterId" xml:"DBClusterId" require:"true"`
	OrderId     *string                         `json:"OrderId" xml:"OrderId" require:"true"`
	DBNodeIds   *CreateDBNodesResponseDBNodeIds `json:"DBNodeIds" xml:"DBNodeIds" require:"true" type:"Struct"`
}

func (s CreateDBNodesResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDBNodesResponse) GoString() string {
	return s.String()
}

func (s *CreateDBNodesResponse) SetRequestId(v string) *CreateDBNodesResponse {
	s.RequestId = &v
	return s
}

func (s *CreateDBNodesResponse) SetDBClusterId(v string) *CreateDBNodesResponse {
	s.DBClusterId = &v
	return s
}

func (s *CreateDBNodesResponse) SetOrderId(v string) *CreateDBNodesResponse {
	s.OrderId = &v
	return s
}

func (s *CreateDBNodesResponse) SetDBNodeIds(v *CreateDBNodesResponseDBNodeIds) *CreateDBNodesResponse {
	s.DBNodeIds = v
	return s
}

type CreateDBNodesResponseDBNodeIds struct {
	DBNodeId []*string `json:"DBNodeId" xml:"DBNodeId" require:"true" type:"Repeated"`
}

func (s CreateDBNodesResponseDBNodeIds) String() string {
	return tea.Prettify(s)
}

func (s CreateDBNodesResponseDBNodeIds) GoString() string {
	return s.String()
}

func (s *CreateDBNodesResponseDBNodeIds) SetDBNodeId(v []*string) *CreateDBNodesResponseDBNodeIds {
	s.DBNodeId = v
	return s
}

type DeleteDBNodesRequest struct {
	AccessKeyId          *string   `json:"AccessKeyId" xml:"AccessKeyId"`
	OwnerId              *int64    `json:"OwnerId" xml:"OwnerId"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount" xml:"ResourceOwnerAccount"`
	ResourceOwnerId      *int64    `json:"ResourceOwnerId" xml:"ResourceOwnerId"`
	OwnerAccount         *string   `json:"OwnerAccount" xml:"OwnerAccount"`
	DBClusterId          *string   `json:"DBClusterId" xml:"DBClusterId" require:"true"`
	DBNodeId             []*string `json:"DBNodeId" xml:"DBNodeId" require:"true" type:"Repeated"`
	ClientToken          *string   `json:"ClientToken" xml:"ClientToken"`
}

func (s DeleteDBNodesRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDBNodesRequest) GoString() string {
	return s.String()
}

func (s *DeleteDBNodesRequest) SetAccessKeyId(v string) *DeleteDBNodesRequest {
	s.AccessKeyId = &v
	return s
}

func (s *DeleteDBNodesRequest) SetOwnerId(v int64) *DeleteDBNodesRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteDBNodesRequest) SetResourceOwnerAccount(v string) *DeleteDBNodesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteDBNodesRequest) SetResourceOwnerId(v int64) *DeleteDBNodesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteDBNodesRequest) SetOwnerAccount(v string) *DeleteDBNodesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteDBNodesRequest) SetDBClusterId(v string) *DeleteDBNodesRequest {
	s.DBClusterId = &v
	return s
}

func (s *DeleteDBNodesRequest) SetDBNodeId(v []*string) *DeleteDBNodesRequest {
	s.DBNodeId = v
	return s
}

func (s *DeleteDBNodesRequest) SetClientToken(v string) *DeleteDBNodesRequest {
	s.ClientToken = &v
	return s
}

type DeleteDBNodesResponse struct {
	RequestId   *string `json:"RequestId" xml:"RequestId" require:"true"`
	DBClusterId *string `json:"DBClusterId" xml:"DBClusterId" require:"true"`
	OrderId     *string `json:"OrderId" xml:"OrderId" require:"true"`
}

func (s DeleteDBNodesResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDBNodesResponse) GoString() string {
	return s.String()
}

func (s *DeleteDBNodesResponse) SetRequestId(v string) *DeleteDBNodesResponse {
	s.RequestId = &v
	return s
}

func (s *DeleteDBNodesResponse) SetDBClusterId(v string) *DeleteDBNodesResponse {
	s.DBClusterId = &v
	return s
}

func (s *DeleteDBNodesResponse) SetOrderId(v string) *DeleteDBNodesResponse {
	s.OrderId = &v
	return s
}

type UntagResourcesRequest struct {
	AccessKeyId          *string   `json:"AccessKeyId" xml:"AccessKeyId"`
	OwnerId              *int64    `json:"OwnerId" xml:"OwnerId"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount" xml:"ResourceOwnerAccount"`
	ResourceOwnerId      *int64    `json:"ResourceOwnerId" xml:"ResourceOwnerId"`
	OwnerAccount         *string   `json:"OwnerAccount" xml:"OwnerAccount"`
	RegionId             *string   `json:"RegionId" xml:"RegionId" require:"true"`
	ResourceType         *string   `json:"ResourceType" xml:"ResourceType" require:"true"`
	ResourceId           []*string `json:"ResourceId" xml:"ResourceId" require:"true" type:"Repeated"`
	TagKey               []*string `json:"TagKey" xml:"TagKey" type:"Repeated"`
	All                  *bool     `json:"All" xml:"All"`
}

func (s UntagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesRequest) GoString() string {
	return s.String()
}

func (s *UntagResourcesRequest) SetAccessKeyId(v string) *UntagResourcesRequest {
	s.AccessKeyId = &v
	return s
}

func (s *UntagResourcesRequest) SetOwnerId(v int64) *UntagResourcesRequest {
	s.OwnerId = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceOwnerAccount(v string) *UntagResourcesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceOwnerId(v int64) *UntagResourcesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UntagResourcesRequest) SetOwnerAccount(v string) *UntagResourcesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UntagResourcesRequest) SetRegionId(v string) *UntagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceType(v string) *UntagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceId(v []*string) *UntagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *UntagResourcesRequest) SetTagKey(v []*string) *UntagResourcesRequest {
	s.TagKey = v
	return s
}

func (s *UntagResourcesRequest) SetAll(v bool) *UntagResourcesRequest {
	s.All = &v
	return s
}

type UntagResourcesResponse struct {
	RequestId *string `json:"RequestId" xml:"RequestId" require:"true"`
}

func (s UntagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesResponse) GoString() string {
	return s.String()
}

func (s *UntagResourcesResponse) SetRequestId(v string) *UntagResourcesResponse {
	s.RequestId = &v
	return s
}

type TagResourcesRequest struct {
	AccessKeyId          *string                   `json:"AccessKeyId" xml:"AccessKeyId"`
	OwnerId              *int64                    `json:"OwnerId" xml:"OwnerId"`
	ResourceOwnerAccount *string                   `json:"ResourceOwnerAccount" xml:"ResourceOwnerAccount"`
	ResourceOwnerId      *int64                    `json:"ResourceOwnerId" xml:"ResourceOwnerId"`
	OwnerAccount         *string                   `json:"OwnerAccount" xml:"OwnerAccount"`
	RegionId             *string                   `json:"RegionId" xml:"RegionId" require:"true"`
	ResourceType         *string                   `json:"ResourceType" xml:"ResourceType" require:"true"`
	ResourceId           []*string                 `json:"ResourceId" xml:"ResourceId" require:"true" type:"Repeated"`
	Tag                  []*TagResourcesRequestTag `json:"Tag" xml:"Tag" require:"true" type:"Repeated"`
}

func (s TagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesRequest) GoString() string {
	return s.String()
}

func (s *TagResourcesRequest) SetAccessKeyId(v string) *TagResourcesRequest {
	s.AccessKeyId = &v
	return s
}

func (s *TagResourcesRequest) SetOwnerId(v int64) *TagResourcesRequest {
	s.OwnerId = &v
	return s
}

func (s *TagResourcesRequest) SetResourceOwnerAccount(v string) *TagResourcesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *TagResourcesRequest) SetResourceOwnerId(v int64) *TagResourcesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *TagResourcesRequest) SetOwnerAccount(v string) *TagResourcesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *TagResourcesRequest) SetRegionId(v string) *TagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *TagResourcesRequest) SetResourceType(v string) *TagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *TagResourcesRequest) SetResourceId(v []*string) *TagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *TagResourcesRequest) SetTag(v []*TagResourcesRequestTag) *TagResourcesRequest {
	s.Tag = v
	return s
}

type TagResourcesRequestTag struct {
	Key   *string `json:"Key" xml:"Key" require:"true"`
	Value *string `json:"Value" xml:"Value" require:"true"`
}

func (s TagResourcesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesRequestTag) GoString() string {
	return s.String()
}

func (s *TagResourcesRequestTag) SetKey(v string) *TagResourcesRequestTag {
	s.Key = &v
	return s
}

func (s *TagResourcesRequestTag) SetValue(v string) *TagResourcesRequestTag {
	s.Value = &v
	return s
}

type TagResourcesResponse struct {
	RequestId *string `json:"RequestId" xml:"RequestId" require:"true"`
}

func (s TagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesResponse) GoString() string {
	return s.String()
}

func (s *TagResourcesResponse) SetRequestId(v string) *TagResourcesResponse {
	s.RequestId = &v
	return s
}

type ListTagResourcesRequest struct {
	AccessKeyId          *string                       `json:"AccessKeyId" xml:"AccessKeyId"`
	OwnerId              *int64                        `json:"OwnerId" xml:"OwnerId"`
	ResourceOwnerAccount *string                       `json:"ResourceOwnerAccount" xml:"ResourceOwnerAccount"`
	ResourceOwnerId      *int64                        `json:"ResourceOwnerId" xml:"ResourceOwnerId"`
	OwnerAccount         *string                       `json:"OwnerAccount" xml:"OwnerAccount"`
	RegionId             *string                       `json:"RegionId" xml:"RegionId" require:"true"`
	ResourceType         *string                       `json:"ResourceType" xml:"ResourceType" require:"true"`
	ResourceId           []*string                     `json:"ResourceId" xml:"ResourceId" type:"Repeated"`
	Tag                  []*ListTagResourcesRequestTag `json:"Tag" xml:"Tag" type:"Repeated"`
	NextToken            *string                       `json:"NextToken" xml:"NextToken"`
}

func (s ListTagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequest) SetAccessKeyId(v string) *ListTagResourcesRequest {
	s.AccessKeyId = &v
	return s
}

func (s *ListTagResourcesRequest) SetOwnerId(v int64) *ListTagResourcesRequest {
	s.OwnerId = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceOwnerAccount(v string) *ListTagResourcesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceOwnerId(v int64) *ListTagResourcesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListTagResourcesRequest) SetOwnerAccount(v string) *ListTagResourcesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListTagResourcesRequest) SetRegionId(v string) *ListTagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceType(v string) *ListTagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceId(v []*string) *ListTagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *ListTagResourcesRequest) SetTag(v []*ListTagResourcesRequestTag) *ListTagResourcesRequest {
	s.Tag = v
	return s
}

func (s *ListTagResourcesRequest) SetNextToken(v string) *ListTagResourcesRequest {
	s.NextToken = &v
	return s
}

type ListTagResourcesRequestTag struct {
	Key   *string `json:"Key" xml:"Key" require:"true"`
	Value *string `json:"Value" xml:"Value" require:"true"`
}

func (s ListTagResourcesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesRequestTag) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequestTag) SetKey(v string) *ListTagResourcesRequestTag {
	s.Key = &v
	return s
}

func (s *ListTagResourcesRequestTag) SetValue(v string) *ListTagResourcesRequestTag {
	s.Value = &v
	return s
}

type ListTagResourcesResponse struct {
	RequestId    *string                               `json:"RequestId" xml:"RequestId" require:"true"`
	NextToken    *string                               `json:"NextToken" xml:"NextToken" require:"true"`
	TagResources *ListTagResourcesResponseTagResources `json:"TagResources" xml:"TagResources" require:"true" type:"Struct"`
}

func (s ListTagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponse) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponse) SetRequestId(v string) *ListTagResourcesResponse {
	s.RequestId = &v
	return s
}

func (s *ListTagResourcesResponse) SetNextToken(v string) *ListTagResourcesResponse {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesResponse) SetTagResources(v *ListTagResourcesResponseTagResources) *ListTagResourcesResponse {
	s.TagResources = v
	return s
}

type ListTagResourcesResponseTagResources struct {
	TagResource []*ListTagResourcesResponseTagResourcesTagResource `json:"TagResource" xml:"TagResource" require:"true" type:"Repeated"`
}

func (s ListTagResourcesResponseTagResources) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseTagResources) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseTagResources) SetTagResource(v []*ListTagResourcesResponseTagResourcesTagResource) *ListTagResourcesResponseTagResources {
	s.TagResource = v
	return s
}

type ListTagResourcesResponseTagResourcesTagResource struct {
	TagKey       *string `json:"TagKey" xml:"TagKey" require:"true"`
	TagValue     *string `json:"TagValue" xml:"TagValue" require:"true"`
	ResourceType *string `json:"ResourceType" xml:"ResourceType" require:"true"`
	ResourceId   *string `json:"ResourceId" xml:"ResourceId" require:"true"`
}

func (s ListTagResourcesResponseTagResourcesTagResource) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseTagResourcesTagResource) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseTagResourcesTagResource) SetTagKey(v string) *ListTagResourcesResponseTagResourcesTagResource {
	s.TagKey = &v
	return s
}

func (s *ListTagResourcesResponseTagResourcesTagResource) SetTagValue(v string) *ListTagResourcesResponseTagResourcesTagResource {
	s.TagValue = &v
	return s
}

func (s *ListTagResourcesResponseTagResourcesTagResource) SetResourceType(v string) *ListTagResourcesResponseTagResourcesTagResource {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesResponseTagResourcesTagResource) SetResourceId(v string) *ListTagResourcesResponseTagResourcesTagResource {
	s.ResourceId = &v
	return s
}

type ModifyDBEndpointAddressRequest struct {
	AccessKeyId              *string `json:"AccessKeyId" xml:"AccessKeyId"`
	OwnerId                  *int64  `json:"OwnerId" xml:"OwnerId"`
	ResourceOwnerAccount     *string `json:"ResourceOwnerAccount" xml:"ResourceOwnerAccount"`
	ResourceOwnerId          *int64  `json:"ResourceOwnerId" xml:"ResourceOwnerId"`
	OwnerAccount             *string `json:"OwnerAccount" xml:"OwnerAccount"`
	DBClusterId              *string `json:"DBClusterId" xml:"DBClusterId" require:"true"`
	DBEndpointId             *string `json:"DBEndpointId" xml:"DBEndpointId" require:"true"`
	NetType                  *string `json:"NetType" xml:"NetType" require:"true"`
	ConnectionStringPrefix   *string `json:"ConnectionStringPrefix" xml:"ConnectionStringPrefix"`
	PrivateZoneAddressPrefix *string `json:"PrivateZoneAddressPrefix" xml:"PrivateZoneAddressPrefix"`
	PrivateZoneName          *string `json:"PrivateZoneName" xml:"PrivateZoneName"`
}

func (s ModifyDBEndpointAddressRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBEndpointAddressRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBEndpointAddressRequest) SetAccessKeyId(v string) *ModifyDBEndpointAddressRequest {
	s.AccessKeyId = &v
	return s
}

func (s *ModifyDBEndpointAddressRequest) SetOwnerId(v int64) *ModifyDBEndpointAddressRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDBEndpointAddressRequest) SetResourceOwnerAccount(v string) *ModifyDBEndpointAddressRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDBEndpointAddressRequest) SetResourceOwnerId(v int64) *ModifyDBEndpointAddressRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDBEndpointAddressRequest) SetOwnerAccount(v string) *ModifyDBEndpointAddressRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyDBEndpointAddressRequest) SetDBClusterId(v string) *ModifyDBEndpointAddressRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyDBEndpointAddressRequest) SetDBEndpointId(v string) *ModifyDBEndpointAddressRequest {
	s.DBEndpointId = &v
	return s
}

func (s *ModifyDBEndpointAddressRequest) SetNetType(v string) *ModifyDBEndpointAddressRequest {
	s.NetType = &v
	return s
}

func (s *ModifyDBEndpointAddressRequest) SetConnectionStringPrefix(v string) *ModifyDBEndpointAddressRequest {
	s.ConnectionStringPrefix = &v
	return s
}

func (s *ModifyDBEndpointAddressRequest) SetPrivateZoneAddressPrefix(v string) *ModifyDBEndpointAddressRequest {
	s.PrivateZoneAddressPrefix = &v
	return s
}

func (s *ModifyDBEndpointAddressRequest) SetPrivateZoneName(v string) *ModifyDBEndpointAddressRequest {
	s.PrivateZoneName = &v
	return s
}

type ModifyDBEndpointAddressResponse struct {
	RequestId *string `json:"RequestId" xml:"RequestId" require:"true"`
}

func (s ModifyDBEndpointAddressResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBEndpointAddressResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBEndpointAddressResponse) SetRequestId(v string) *ModifyDBEndpointAddressResponse {
	s.RequestId = &v
	return s
}

type ModifyDBDescriptionRequest struct {
	AccessKeyId          *string `json:"AccessKeyId" xml:"AccessKeyId"`
	OwnerId              *int64  `json:"OwnerId" xml:"OwnerId"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount" xml:"ResourceOwnerAccount"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId" xml:"ResourceOwnerId"`
	OwnerAccount         *string `json:"OwnerAccount" xml:"OwnerAccount"`
	DBClusterId          *string `json:"DBClusterId" xml:"DBClusterId" require:"true"`
	DBName               *string `json:"DBName" xml:"DBName" require:"true"`
	DBDescription        *string `json:"DBDescription" xml:"DBDescription" require:"true"`
}

func (s ModifyDBDescriptionRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBDescriptionRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBDescriptionRequest) SetAccessKeyId(v string) *ModifyDBDescriptionRequest {
	s.AccessKeyId = &v
	return s
}

func (s *ModifyDBDescriptionRequest) SetOwnerId(v int64) *ModifyDBDescriptionRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDBDescriptionRequest) SetResourceOwnerAccount(v string) *ModifyDBDescriptionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDBDescriptionRequest) SetResourceOwnerId(v int64) *ModifyDBDescriptionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDBDescriptionRequest) SetOwnerAccount(v string) *ModifyDBDescriptionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyDBDescriptionRequest) SetDBClusterId(v string) *ModifyDBDescriptionRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyDBDescriptionRequest) SetDBName(v string) *ModifyDBDescriptionRequest {
	s.DBName = &v
	return s
}

func (s *ModifyDBDescriptionRequest) SetDBDescription(v string) *ModifyDBDescriptionRequest {
	s.DBDescription = &v
	return s
}

type ModifyDBDescriptionResponse struct {
	RequestId *string `json:"RequestId" xml:"RequestId" require:"true"`
}

func (s ModifyDBDescriptionResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBDescriptionResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBDescriptionResponse) SetRequestId(v string) *ModifyDBDescriptionResponse {
	s.RequestId = &v
	return s
}

type ModifyDBClusterParametersRequest struct {
	AccessKeyId          *string `json:"AccessKeyId" xml:"AccessKeyId"`
	OwnerId              *int64  `json:"OwnerId" xml:"OwnerId"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount" xml:"ResourceOwnerAccount"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId" xml:"ResourceOwnerId"`
	OwnerAccount         *string `json:"OwnerAccount" xml:"OwnerAccount"`
	DBClusterId          *string `json:"DBClusterId" xml:"DBClusterId" require:"true"`
	Parameters           *string `json:"Parameters" xml:"Parameters" require:"true"`
	EffectiveTime        *string `json:"EffectiveTime" xml:"EffectiveTime"`
}

func (s ModifyDBClusterParametersRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBClusterParametersRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterParametersRequest) SetAccessKeyId(v string) *ModifyDBClusterParametersRequest {
	s.AccessKeyId = &v
	return s
}

func (s *ModifyDBClusterParametersRequest) SetOwnerId(v int64) *ModifyDBClusterParametersRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDBClusterParametersRequest) SetResourceOwnerAccount(v string) *ModifyDBClusterParametersRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDBClusterParametersRequest) SetResourceOwnerId(v int64) *ModifyDBClusterParametersRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDBClusterParametersRequest) SetOwnerAccount(v string) *ModifyDBClusterParametersRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyDBClusterParametersRequest) SetDBClusterId(v string) *ModifyDBClusterParametersRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyDBClusterParametersRequest) SetParameters(v string) *ModifyDBClusterParametersRequest {
	s.Parameters = &v
	return s
}

func (s *ModifyDBClusterParametersRequest) SetEffectiveTime(v string) *ModifyDBClusterParametersRequest {
	s.EffectiveTime = &v
	return s
}

type ModifyDBClusterParametersResponse struct {
	RequestId *string `json:"RequestId" xml:"RequestId" require:"true"`
}

func (s ModifyDBClusterParametersResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBClusterParametersResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterParametersResponse) SetRequestId(v string) *ModifyDBClusterParametersResponse {
	s.RequestId = &v
	return s
}

type ModifyDBClusterEndpointRequest struct {
	AccessKeyId          *string `json:"AccessKeyId" xml:"AccessKeyId"`
	OwnerId              *int64  `json:"OwnerId" xml:"OwnerId"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount" xml:"ResourceOwnerAccount"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId" xml:"ResourceOwnerId"`
	OwnerAccount         *string `json:"OwnerAccount" xml:"OwnerAccount"`
	DBClusterId          *string `json:"DBClusterId" xml:"DBClusterId" require:"true"`
	DBEndpointId         *string `json:"DBEndpointId" xml:"DBEndpointId" require:"true"`
	Nodes                *string `json:"Nodes" xml:"Nodes"`
	ReadWriteMode        *string `json:"ReadWriteMode" xml:"ReadWriteMode"`
	AutoAddNewNodes      *string `json:"AutoAddNewNodes" xml:"AutoAddNewNodes"`
	EndpointConfig       *string `json:"EndpointConfig" xml:"EndpointConfig"`
}

func (s ModifyDBClusterEndpointRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBClusterEndpointRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterEndpointRequest) SetAccessKeyId(v string) *ModifyDBClusterEndpointRequest {
	s.AccessKeyId = &v
	return s
}

func (s *ModifyDBClusterEndpointRequest) SetOwnerId(v int64) *ModifyDBClusterEndpointRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDBClusterEndpointRequest) SetResourceOwnerAccount(v string) *ModifyDBClusterEndpointRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDBClusterEndpointRequest) SetResourceOwnerId(v int64) *ModifyDBClusterEndpointRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDBClusterEndpointRequest) SetOwnerAccount(v string) *ModifyDBClusterEndpointRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyDBClusterEndpointRequest) SetDBClusterId(v string) *ModifyDBClusterEndpointRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyDBClusterEndpointRequest) SetDBEndpointId(v string) *ModifyDBClusterEndpointRequest {
	s.DBEndpointId = &v
	return s
}

func (s *ModifyDBClusterEndpointRequest) SetNodes(v string) *ModifyDBClusterEndpointRequest {
	s.Nodes = &v
	return s
}

func (s *ModifyDBClusterEndpointRequest) SetReadWriteMode(v string) *ModifyDBClusterEndpointRequest {
	s.ReadWriteMode = &v
	return s
}

func (s *ModifyDBClusterEndpointRequest) SetAutoAddNewNodes(v string) *ModifyDBClusterEndpointRequest {
	s.AutoAddNewNodes = &v
	return s
}

func (s *ModifyDBClusterEndpointRequest) SetEndpointConfig(v string) *ModifyDBClusterEndpointRequest {
	s.EndpointConfig = &v
	return s
}

type ModifyDBClusterEndpointResponse struct {
	RequestId *string `json:"RequestId" xml:"RequestId" require:"true"`
}

func (s ModifyDBClusterEndpointResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBClusterEndpointResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterEndpointResponse) SetRequestId(v string) *ModifyDBClusterEndpointResponse {
	s.RequestId = &v
	return s
}

type ModifyAccountPasswordRequest struct {
	AccessKeyId          *string `json:"AccessKeyId" xml:"AccessKeyId"`
	OwnerId              *int64  `json:"OwnerId" xml:"OwnerId"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount" xml:"ResourceOwnerAccount"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId" xml:"ResourceOwnerId"`
	OwnerAccount         *string `json:"OwnerAccount" xml:"OwnerAccount"`
	DBClusterId          *string `json:"DBClusterId" xml:"DBClusterId" require:"true"`
	AccountName          *string `json:"AccountName" xml:"AccountName" require:"true"`
	NewAccountPassword   *string `json:"NewAccountPassword" xml:"NewAccountPassword" require:"true"`
}

func (s ModifyAccountPasswordRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyAccountPasswordRequest) GoString() string {
	return s.String()
}

func (s *ModifyAccountPasswordRequest) SetAccessKeyId(v string) *ModifyAccountPasswordRequest {
	s.AccessKeyId = &v
	return s
}

func (s *ModifyAccountPasswordRequest) SetOwnerId(v int64) *ModifyAccountPasswordRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyAccountPasswordRequest) SetResourceOwnerAccount(v string) *ModifyAccountPasswordRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyAccountPasswordRequest) SetResourceOwnerId(v int64) *ModifyAccountPasswordRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyAccountPasswordRequest) SetOwnerAccount(v string) *ModifyAccountPasswordRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyAccountPasswordRequest) SetDBClusterId(v string) *ModifyAccountPasswordRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyAccountPasswordRequest) SetAccountName(v string) *ModifyAccountPasswordRequest {
	s.AccountName = &v
	return s
}

func (s *ModifyAccountPasswordRequest) SetNewAccountPassword(v string) *ModifyAccountPasswordRequest {
	s.NewAccountPassword = &v
	return s
}

type ModifyAccountPasswordResponse struct {
	RequestId *string `json:"RequestId" xml:"RequestId" require:"true"`
}

func (s ModifyAccountPasswordResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyAccountPasswordResponse) GoString() string {
	return s.String()
}

func (s *ModifyAccountPasswordResponse) SetRequestId(v string) *ModifyAccountPasswordResponse {
	s.RequestId = &v
	return s
}

type DescribeDBClusterPerformanceRequest struct {
	AccessKeyId *string `json:"AccessKeyId" xml:"AccessKeyId"`
	DBClusterId *string `json:"DBClusterId" xml:"DBClusterId" require:"true"`
	Key         *string `json:"Key" xml:"Key" require:"true"`
	StartTime   *string `json:"StartTime" xml:"StartTime" require:"true"`
	EndTime     *string `json:"EndTime" xml:"EndTime" require:"true"`
}

func (s DescribeDBClusterPerformanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterPerformanceRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterPerformanceRequest) SetAccessKeyId(v string) *DescribeDBClusterPerformanceRequest {
	s.AccessKeyId = &v
	return s
}

func (s *DescribeDBClusterPerformanceRequest) SetDBClusterId(v string) *DescribeDBClusterPerformanceRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDBClusterPerformanceRequest) SetKey(v string) *DescribeDBClusterPerformanceRequest {
	s.Key = &v
	return s
}

func (s *DescribeDBClusterPerformanceRequest) SetStartTime(v string) *DescribeDBClusterPerformanceRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDBClusterPerformanceRequest) SetEndTime(v string) *DescribeDBClusterPerformanceRequest {
	s.EndTime = &v
	return s
}

type DescribeDBClusterPerformanceResponse struct {
	RequestId       *string                                              `json:"RequestId" xml:"RequestId" require:"true"`
	DBClusterId     *string                                              `json:"DBClusterId" xml:"DBClusterId" require:"true"`
	Engine          *string                                              `json:"Engine" xml:"Engine" require:"true"`
	DBType          *string                                              `json:"DBType" xml:"DBType" require:"true"`
	DBVersion       *string                                              `json:"DBVersion" xml:"DBVersion" require:"true"`
	StartTime       *string                                              `json:"StartTime" xml:"StartTime" require:"true"`
	EndTime         *string                                              `json:"EndTime" xml:"EndTime" require:"true"`
	PerformanceKeys *DescribeDBClusterPerformanceResponsePerformanceKeys `json:"PerformanceKeys" xml:"PerformanceKeys" require:"true" type:"Struct"`
}

func (s DescribeDBClusterPerformanceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterPerformanceResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterPerformanceResponse) SetRequestId(v string) *DescribeDBClusterPerformanceResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeDBClusterPerformanceResponse) SetDBClusterId(v string) *DescribeDBClusterPerformanceResponse {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDBClusterPerformanceResponse) SetEngine(v string) *DescribeDBClusterPerformanceResponse {
	s.Engine = &v
	return s
}

func (s *DescribeDBClusterPerformanceResponse) SetDBType(v string) *DescribeDBClusterPerformanceResponse {
	s.DBType = &v
	return s
}

func (s *DescribeDBClusterPerformanceResponse) SetDBVersion(v string) *DescribeDBClusterPerformanceResponse {
	s.DBVersion = &v
	return s
}

func (s *DescribeDBClusterPerformanceResponse) SetStartTime(v string) *DescribeDBClusterPerformanceResponse {
	s.StartTime = &v
	return s
}

func (s *DescribeDBClusterPerformanceResponse) SetEndTime(v string) *DescribeDBClusterPerformanceResponse {
	s.EndTime = &v
	return s
}

func (s *DescribeDBClusterPerformanceResponse) SetPerformanceKeys(v *DescribeDBClusterPerformanceResponsePerformanceKeys) *DescribeDBClusterPerformanceResponse {
	s.PerformanceKeys = v
	return s
}

type DescribeDBClusterPerformanceResponsePerformanceKeys struct {
	PerformanceItem []*DescribeDBClusterPerformanceResponsePerformanceKeysPerformanceItem `json:"PerformanceItem" xml:"PerformanceItem" require:"true" type:"Repeated"`
}

func (s DescribeDBClusterPerformanceResponsePerformanceKeys) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterPerformanceResponsePerformanceKeys) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterPerformanceResponsePerformanceKeys) SetPerformanceItem(v []*DescribeDBClusterPerformanceResponsePerformanceKeysPerformanceItem) *DescribeDBClusterPerformanceResponsePerformanceKeys {
	s.PerformanceItem = v
	return s
}

type DescribeDBClusterPerformanceResponsePerformanceKeysPerformanceItem struct {
	DBNodeId    *string                                                                   `json:"DBNodeId" xml:"DBNodeId" require:"true"`
	Measurement *string                                                                   `json:"Measurement" xml:"Measurement" require:"true"`
	MetricName  *string                                                                   `json:"MetricName" xml:"MetricName" require:"true"`
	Points      *DescribeDBClusterPerformanceResponsePerformanceKeysPerformanceItemPoints `json:"Points" xml:"Points" require:"true" type:"Struct"`
}

func (s DescribeDBClusterPerformanceResponsePerformanceKeysPerformanceItem) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterPerformanceResponsePerformanceKeysPerformanceItem) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterPerformanceResponsePerformanceKeysPerformanceItem) SetDBNodeId(v string) *DescribeDBClusterPerformanceResponsePerformanceKeysPerformanceItem {
	s.DBNodeId = &v
	return s
}

func (s *DescribeDBClusterPerformanceResponsePerformanceKeysPerformanceItem) SetMeasurement(v string) *DescribeDBClusterPerformanceResponsePerformanceKeysPerformanceItem {
	s.Measurement = &v
	return s
}

func (s *DescribeDBClusterPerformanceResponsePerformanceKeysPerformanceItem) SetMetricName(v string) *DescribeDBClusterPerformanceResponsePerformanceKeysPerformanceItem {
	s.MetricName = &v
	return s
}

func (s *DescribeDBClusterPerformanceResponsePerformanceKeysPerformanceItem) SetPoints(v *DescribeDBClusterPerformanceResponsePerformanceKeysPerformanceItemPoints) *DescribeDBClusterPerformanceResponsePerformanceKeysPerformanceItem {
	s.Points = v
	return s
}

type DescribeDBClusterPerformanceResponsePerformanceKeysPerformanceItemPoints struct {
	PerformanceItemValue []*DescribeDBClusterPerformanceResponsePerformanceKeysPerformanceItemPointsPerformanceItemValue `json:"PerformanceItemValue" xml:"PerformanceItemValue" require:"true" type:"Repeated"`
}

func (s DescribeDBClusterPerformanceResponsePerformanceKeysPerformanceItemPoints) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterPerformanceResponsePerformanceKeysPerformanceItemPoints) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterPerformanceResponsePerformanceKeysPerformanceItemPoints) SetPerformanceItemValue(v []*DescribeDBClusterPerformanceResponsePerformanceKeysPerformanceItemPointsPerformanceItemValue) *DescribeDBClusterPerformanceResponsePerformanceKeysPerformanceItemPoints {
	s.PerformanceItemValue = v
	return s
}

type DescribeDBClusterPerformanceResponsePerformanceKeysPerformanceItemPointsPerformanceItemValue struct {
	Value     *string `json:"Value" xml:"Value" require:"true"`
	Timestamp *int64  `json:"Timestamp" xml:"Timestamp" require:"true"`
}

func (s DescribeDBClusterPerformanceResponsePerformanceKeysPerformanceItemPointsPerformanceItemValue) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterPerformanceResponsePerformanceKeysPerformanceItemPointsPerformanceItemValue) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterPerformanceResponsePerformanceKeysPerformanceItemPointsPerformanceItemValue) SetValue(v string) *DescribeDBClusterPerformanceResponsePerformanceKeysPerformanceItemPointsPerformanceItemValue {
	s.Value = &v
	return s
}

func (s *DescribeDBClusterPerformanceResponsePerformanceKeysPerformanceItemPointsPerformanceItemValue) SetTimestamp(v int64) *DescribeDBClusterPerformanceResponsePerformanceKeysPerformanceItemPointsPerformanceItemValue {
	s.Timestamp = &v
	return s
}

type DescribeDBClusterParametersRequest struct {
	AccessKeyId          *string `json:"AccessKeyId" xml:"AccessKeyId"`
	OwnerId              *int64  `json:"OwnerId" xml:"OwnerId"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount" xml:"ResourceOwnerAccount"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId" xml:"ResourceOwnerId"`
	OwnerAccount         *string `json:"OwnerAccount" xml:"OwnerAccount"`
	DBClusterId          *string `json:"DBClusterId" xml:"DBClusterId" require:"true"`
}

func (s DescribeDBClusterParametersRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterParametersRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterParametersRequest) SetAccessKeyId(v string) *DescribeDBClusterParametersRequest {
	s.AccessKeyId = &v
	return s
}

func (s *DescribeDBClusterParametersRequest) SetOwnerId(v int64) *DescribeDBClusterParametersRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDBClusterParametersRequest) SetResourceOwnerAccount(v string) *DescribeDBClusterParametersRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDBClusterParametersRequest) SetResourceOwnerId(v int64) *DescribeDBClusterParametersRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDBClusterParametersRequest) SetOwnerAccount(v string) *DescribeDBClusterParametersRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDBClusterParametersRequest) SetDBClusterId(v string) *DescribeDBClusterParametersRequest {
	s.DBClusterId = &v
	return s
}

type DescribeDBClusterParametersResponse struct {
	RequestId         *string                                               `json:"RequestId" xml:"RequestId" require:"true"`
	Engine            *string                                               `json:"Engine" xml:"Engine" require:"true"`
	DBType            *string                                               `json:"DBType" xml:"DBType" require:"true"`
	DBVersion         *string                                               `json:"DBVersion" xml:"DBVersion" require:"true"`
	RunningParameters *DescribeDBClusterParametersResponseRunningParameters `json:"RunningParameters" xml:"RunningParameters" require:"true" type:"Struct"`
}

func (s DescribeDBClusterParametersResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterParametersResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterParametersResponse) SetRequestId(v string) *DescribeDBClusterParametersResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeDBClusterParametersResponse) SetEngine(v string) *DescribeDBClusterParametersResponse {
	s.Engine = &v
	return s
}

func (s *DescribeDBClusterParametersResponse) SetDBType(v string) *DescribeDBClusterParametersResponse {
	s.DBType = &v
	return s
}

func (s *DescribeDBClusterParametersResponse) SetDBVersion(v string) *DescribeDBClusterParametersResponse {
	s.DBVersion = &v
	return s
}

func (s *DescribeDBClusterParametersResponse) SetRunningParameters(v *DescribeDBClusterParametersResponseRunningParameters) *DescribeDBClusterParametersResponse {
	s.RunningParameters = v
	return s
}

type DescribeDBClusterParametersResponseRunningParameters struct {
	Parameter []*DescribeDBClusterParametersResponseRunningParametersParameter `json:"Parameter" xml:"Parameter" require:"true" type:"Repeated"`
}

func (s DescribeDBClusterParametersResponseRunningParameters) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterParametersResponseRunningParameters) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterParametersResponseRunningParameters) SetParameter(v []*DescribeDBClusterParametersResponseRunningParametersParameter) *DescribeDBClusterParametersResponseRunningParameters {
	s.Parameter = v
	return s
}

type DescribeDBClusterParametersResponseRunningParametersParameter struct {
	ParameterName         *string `json:"ParameterName" xml:"ParameterName" require:"true"`
	DataType              *string `json:"DataType" xml:"DataType" require:"true"`
	DefaultParameterValue *string `json:"DefaultParameterValue" xml:"DefaultParameterValue" require:"true"`
	ParameterValue        *string `json:"ParameterValue" xml:"ParameterValue" require:"true"`
	IsModifiable          *bool   `json:"IsModifiable" xml:"IsModifiable" require:"true"`
	ForceRestart          *bool   `json:"ForceRestart" xml:"ForceRestart" require:"true"`
	ParameterStatus       *string `json:"ParameterStatus" xml:"ParameterStatus" require:"true"`
	CheckingCode          *string `json:"CheckingCode" xml:"CheckingCode" require:"true"`
	ParameterDescription  *string `json:"ParameterDescription" xml:"ParameterDescription" require:"true"`
}

func (s DescribeDBClusterParametersResponseRunningParametersParameter) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterParametersResponseRunningParametersParameter) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterParametersResponseRunningParametersParameter) SetParameterName(v string) *DescribeDBClusterParametersResponseRunningParametersParameter {
	s.ParameterName = &v
	return s
}

func (s *DescribeDBClusterParametersResponseRunningParametersParameter) SetDataType(v string) *DescribeDBClusterParametersResponseRunningParametersParameter {
	s.DataType = &v
	return s
}

func (s *DescribeDBClusterParametersResponseRunningParametersParameter) SetDefaultParameterValue(v string) *DescribeDBClusterParametersResponseRunningParametersParameter {
	s.DefaultParameterValue = &v
	return s
}

func (s *DescribeDBClusterParametersResponseRunningParametersParameter) SetParameterValue(v string) *DescribeDBClusterParametersResponseRunningParametersParameter {
	s.ParameterValue = &v
	return s
}

func (s *DescribeDBClusterParametersResponseRunningParametersParameter) SetIsModifiable(v bool) *DescribeDBClusterParametersResponseRunningParametersParameter {
	s.IsModifiable = &v
	return s
}

func (s *DescribeDBClusterParametersResponseRunningParametersParameter) SetForceRestart(v bool) *DescribeDBClusterParametersResponseRunningParametersParameter {
	s.ForceRestart = &v
	return s
}

func (s *DescribeDBClusterParametersResponseRunningParametersParameter) SetParameterStatus(v string) *DescribeDBClusterParametersResponseRunningParametersParameter {
	s.ParameterStatus = &v
	return s
}

func (s *DescribeDBClusterParametersResponseRunningParametersParameter) SetCheckingCode(v string) *DescribeDBClusterParametersResponseRunningParametersParameter {
	s.CheckingCode = &v
	return s
}

func (s *DescribeDBClusterParametersResponseRunningParametersParameter) SetParameterDescription(v string) *DescribeDBClusterParametersResponseRunningParametersParameter {
	s.ParameterDescription = &v
	return s
}

type DescribeDBClusterEndpointsRequest struct {
	AccessKeyId          *string `json:"AccessKeyId" xml:"AccessKeyId"`
	OwnerId              *int64  `json:"OwnerId" xml:"OwnerId"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount" xml:"ResourceOwnerAccount"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId" xml:"ResourceOwnerId"`
	OwnerAccount         *string `json:"OwnerAccount" xml:"OwnerAccount"`
	DBClusterId          *string `json:"DBClusterId" xml:"DBClusterId" require:"true"`
	DBEndpointId         *string `json:"DBEndpointId" xml:"DBEndpointId"`
}

func (s DescribeDBClusterEndpointsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterEndpointsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterEndpointsRequest) SetAccessKeyId(v string) *DescribeDBClusterEndpointsRequest {
	s.AccessKeyId = &v
	return s
}

func (s *DescribeDBClusterEndpointsRequest) SetOwnerId(v int64) *DescribeDBClusterEndpointsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDBClusterEndpointsRequest) SetResourceOwnerAccount(v string) *DescribeDBClusterEndpointsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDBClusterEndpointsRequest) SetResourceOwnerId(v int64) *DescribeDBClusterEndpointsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDBClusterEndpointsRequest) SetOwnerAccount(v string) *DescribeDBClusterEndpointsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDBClusterEndpointsRequest) SetDBClusterId(v string) *DescribeDBClusterEndpointsRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDBClusterEndpointsRequest) SetDBEndpointId(v string) *DescribeDBClusterEndpointsRequest {
	s.DBEndpointId = &v
	return s
}

type DescribeDBClusterEndpointsResponse struct {
	RequestId *string                                    `json:"RequestId" xml:"RequestId" require:"true"`
	Items     []*DescribeDBClusterEndpointsResponseItems `json:"Items" xml:"Items" require:"true" type:"Repeated"`
}

func (s DescribeDBClusterEndpointsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterEndpointsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterEndpointsResponse) SetRequestId(v string) *DescribeDBClusterEndpointsResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeDBClusterEndpointsResponse) SetItems(v []*DescribeDBClusterEndpointsResponseItems) *DescribeDBClusterEndpointsResponse {
	s.Items = v
	return s
}

type DescribeDBClusterEndpointsResponseItems struct {
	DBEndpointId    *string                                                `json:"DBEndpointId" xml:"DBEndpointId" require:"true"`
	EndpointType    *string                                                `json:"EndpointType" xml:"EndpointType" require:"true"`
	Nodes           *string                                                `json:"Nodes" xml:"Nodes" require:"true"`
	ReadWriteMode   *string                                                `json:"ReadWriteMode" xml:"ReadWriteMode" require:"true"`
	AutoAddNewNodes *string                                                `json:"AutoAddNewNodes" xml:"AutoAddNewNodes" require:"true"`
	EndpointConfig  *string                                                `json:"EndpointConfig" xml:"EndpointConfig" require:"true"`
	AddressItems    []*DescribeDBClusterEndpointsResponseItemsAddressItems `json:"AddressItems" xml:"AddressItems" require:"true" type:"Repeated"`
}

func (s DescribeDBClusterEndpointsResponseItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterEndpointsResponseItems) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterEndpointsResponseItems) SetDBEndpointId(v string) *DescribeDBClusterEndpointsResponseItems {
	s.DBEndpointId = &v
	return s
}

func (s *DescribeDBClusterEndpointsResponseItems) SetEndpointType(v string) *DescribeDBClusterEndpointsResponseItems {
	s.EndpointType = &v
	return s
}

func (s *DescribeDBClusterEndpointsResponseItems) SetNodes(v string) *DescribeDBClusterEndpointsResponseItems {
	s.Nodes = &v
	return s
}

func (s *DescribeDBClusterEndpointsResponseItems) SetReadWriteMode(v string) *DescribeDBClusterEndpointsResponseItems {
	s.ReadWriteMode = &v
	return s
}

func (s *DescribeDBClusterEndpointsResponseItems) SetAutoAddNewNodes(v string) *DescribeDBClusterEndpointsResponseItems {
	s.AutoAddNewNodes = &v
	return s
}

func (s *DescribeDBClusterEndpointsResponseItems) SetEndpointConfig(v string) *DescribeDBClusterEndpointsResponseItems {
	s.EndpointConfig = &v
	return s
}

func (s *DescribeDBClusterEndpointsResponseItems) SetAddressItems(v []*DescribeDBClusterEndpointsResponseItemsAddressItems) *DescribeDBClusterEndpointsResponseItems {
	s.AddressItems = v
	return s
}

type DescribeDBClusterEndpointsResponseItemsAddressItems struct {
	ConnectionString            *string `json:"ConnectionString" xml:"ConnectionString" require:"true"`
	IPAddress                   *string `json:"IPAddress" xml:"IPAddress" require:"true"`
	NetType                     *string `json:"NetType" xml:"NetType" require:"true"`
	Port                        *string `json:"Port" xml:"Port" require:"true"`
	VPCId                       *string `json:"VPCId" xml:"VPCId" require:"true"`
	VSwitchId                   *string `json:"VSwitchId" xml:"VSwitchId" require:"true"`
	VpcInstanceId               *string `json:"VpcInstanceId" xml:"VpcInstanceId" require:"true"`
	PrivateZoneConnectionString *string `json:"PrivateZoneConnectionString" xml:"PrivateZoneConnectionString" require:"true"`
}

func (s DescribeDBClusterEndpointsResponseItemsAddressItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterEndpointsResponseItemsAddressItems) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterEndpointsResponseItemsAddressItems) SetConnectionString(v string) *DescribeDBClusterEndpointsResponseItemsAddressItems {
	s.ConnectionString = &v
	return s
}

func (s *DescribeDBClusterEndpointsResponseItemsAddressItems) SetIPAddress(v string) *DescribeDBClusterEndpointsResponseItemsAddressItems {
	s.IPAddress = &v
	return s
}

func (s *DescribeDBClusterEndpointsResponseItemsAddressItems) SetNetType(v string) *DescribeDBClusterEndpointsResponseItemsAddressItems {
	s.NetType = &v
	return s
}

func (s *DescribeDBClusterEndpointsResponseItemsAddressItems) SetPort(v string) *DescribeDBClusterEndpointsResponseItemsAddressItems {
	s.Port = &v
	return s
}

func (s *DescribeDBClusterEndpointsResponseItemsAddressItems) SetVPCId(v string) *DescribeDBClusterEndpointsResponseItemsAddressItems {
	s.VPCId = &v
	return s
}

func (s *DescribeDBClusterEndpointsResponseItemsAddressItems) SetVSwitchId(v string) *DescribeDBClusterEndpointsResponseItemsAddressItems {
	s.VSwitchId = &v
	return s
}

func (s *DescribeDBClusterEndpointsResponseItemsAddressItems) SetVpcInstanceId(v string) *DescribeDBClusterEndpointsResponseItemsAddressItems {
	s.VpcInstanceId = &v
	return s
}

func (s *DescribeDBClusterEndpointsResponseItemsAddressItems) SetPrivateZoneConnectionString(v string) *DescribeDBClusterEndpointsResponseItemsAddressItems {
	s.PrivateZoneConnectionString = &v
	return s
}

type DeleteDBEndpointAddressRequest struct {
	AccessKeyId          *string `json:"AccessKeyId" xml:"AccessKeyId"`
	OwnerId              *int64  `json:"OwnerId" xml:"OwnerId"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount" xml:"ResourceOwnerAccount"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId" xml:"ResourceOwnerId"`
	OwnerAccount         *string `json:"OwnerAccount" xml:"OwnerAccount"`
	DBClusterId          *string `json:"DBClusterId" xml:"DBClusterId" require:"true"`
	DBEndpointId         *string `json:"DBEndpointId" xml:"DBEndpointId" require:"true"`
	NetType              *string `json:"NetType" xml:"NetType" require:"true"`
}

func (s DeleteDBEndpointAddressRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDBEndpointAddressRequest) GoString() string {
	return s.String()
}

func (s *DeleteDBEndpointAddressRequest) SetAccessKeyId(v string) *DeleteDBEndpointAddressRequest {
	s.AccessKeyId = &v
	return s
}

func (s *DeleteDBEndpointAddressRequest) SetOwnerId(v int64) *DeleteDBEndpointAddressRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteDBEndpointAddressRequest) SetResourceOwnerAccount(v string) *DeleteDBEndpointAddressRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteDBEndpointAddressRequest) SetResourceOwnerId(v int64) *DeleteDBEndpointAddressRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteDBEndpointAddressRequest) SetOwnerAccount(v string) *DeleteDBEndpointAddressRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteDBEndpointAddressRequest) SetDBClusterId(v string) *DeleteDBEndpointAddressRequest {
	s.DBClusterId = &v
	return s
}

func (s *DeleteDBEndpointAddressRequest) SetDBEndpointId(v string) *DeleteDBEndpointAddressRequest {
	s.DBEndpointId = &v
	return s
}

func (s *DeleteDBEndpointAddressRequest) SetNetType(v string) *DeleteDBEndpointAddressRequest {
	s.NetType = &v
	return s
}

type DeleteDBEndpointAddressResponse struct {
	RequestId *string `json:"RequestId" xml:"RequestId" require:"true"`
}

func (s DeleteDBEndpointAddressResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDBEndpointAddressResponse) GoString() string {
	return s.String()
}

func (s *DeleteDBEndpointAddressResponse) SetRequestId(v string) *DeleteDBEndpointAddressResponse {
	s.RequestId = &v
	return s
}

type DeleteDBClusterEndpointRequest struct {
	AccessKeyId          *string `json:"AccessKeyId" xml:"AccessKeyId"`
	OwnerId              *int64  `json:"OwnerId" xml:"OwnerId"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount" xml:"ResourceOwnerAccount"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId" xml:"ResourceOwnerId"`
	OwnerAccount         *string `json:"OwnerAccount" xml:"OwnerAccount"`
	DBClusterId          *string `json:"DBClusterId" xml:"DBClusterId" require:"true"`
	DBEndpointId         *string `json:"DBEndpointId" xml:"DBEndpointId" require:"true"`
}

func (s DeleteDBClusterEndpointRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDBClusterEndpointRequest) GoString() string {
	return s.String()
}

func (s *DeleteDBClusterEndpointRequest) SetAccessKeyId(v string) *DeleteDBClusterEndpointRequest {
	s.AccessKeyId = &v
	return s
}

func (s *DeleteDBClusterEndpointRequest) SetOwnerId(v int64) *DeleteDBClusterEndpointRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteDBClusterEndpointRequest) SetResourceOwnerAccount(v string) *DeleteDBClusterEndpointRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteDBClusterEndpointRequest) SetResourceOwnerId(v int64) *DeleteDBClusterEndpointRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteDBClusterEndpointRequest) SetOwnerAccount(v string) *DeleteDBClusterEndpointRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteDBClusterEndpointRequest) SetDBClusterId(v string) *DeleteDBClusterEndpointRequest {
	s.DBClusterId = &v
	return s
}

func (s *DeleteDBClusterEndpointRequest) SetDBEndpointId(v string) *DeleteDBClusterEndpointRequest {
	s.DBEndpointId = &v
	return s
}

type DeleteDBClusterEndpointResponse struct {
	RequestId *string `json:"RequestId" xml:"RequestId" require:"true"`
}

func (s DeleteDBClusterEndpointResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDBClusterEndpointResponse) GoString() string {
	return s.String()
}

func (s *DeleteDBClusterEndpointResponse) SetRequestId(v string) *DeleteDBClusterEndpointResponse {
	s.RequestId = &v
	return s
}

type CreateDBEndpointAddressRequest struct {
	AccessKeyId            *string `json:"AccessKeyId" xml:"AccessKeyId"`
	OwnerId                *int64  `json:"OwnerId" xml:"OwnerId"`
	ResourceOwnerAccount   *string `json:"ResourceOwnerAccount" xml:"ResourceOwnerAccount"`
	ResourceOwnerId        *int64  `json:"ResourceOwnerId" xml:"ResourceOwnerId"`
	OwnerAccount           *string `json:"OwnerAccount" xml:"OwnerAccount"`
	DBClusterId            *string `json:"DBClusterId" xml:"DBClusterId" require:"true"`
	DBEndpointId           *string `json:"DBEndpointId" xml:"DBEndpointId" require:"true"`
	ConnectionStringPrefix *string `json:"ConnectionStringPrefix" xml:"ConnectionStringPrefix"`
	NetType                *string `json:"NetType" xml:"NetType" require:"true"`
}

func (s CreateDBEndpointAddressRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDBEndpointAddressRequest) GoString() string {
	return s.String()
}

func (s *CreateDBEndpointAddressRequest) SetAccessKeyId(v string) *CreateDBEndpointAddressRequest {
	s.AccessKeyId = &v
	return s
}

func (s *CreateDBEndpointAddressRequest) SetOwnerId(v int64) *CreateDBEndpointAddressRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateDBEndpointAddressRequest) SetResourceOwnerAccount(v string) *CreateDBEndpointAddressRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateDBEndpointAddressRequest) SetResourceOwnerId(v int64) *CreateDBEndpointAddressRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateDBEndpointAddressRequest) SetOwnerAccount(v string) *CreateDBEndpointAddressRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateDBEndpointAddressRequest) SetDBClusterId(v string) *CreateDBEndpointAddressRequest {
	s.DBClusterId = &v
	return s
}

func (s *CreateDBEndpointAddressRequest) SetDBEndpointId(v string) *CreateDBEndpointAddressRequest {
	s.DBEndpointId = &v
	return s
}

func (s *CreateDBEndpointAddressRequest) SetConnectionStringPrefix(v string) *CreateDBEndpointAddressRequest {
	s.ConnectionStringPrefix = &v
	return s
}

func (s *CreateDBEndpointAddressRequest) SetNetType(v string) *CreateDBEndpointAddressRequest {
	s.NetType = &v
	return s
}

type CreateDBEndpointAddressResponse struct {
	RequestId *string `json:"RequestId" xml:"RequestId" require:"true"`
}

func (s CreateDBEndpointAddressResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDBEndpointAddressResponse) GoString() string {
	return s.String()
}

func (s *CreateDBEndpointAddressResponse) SetRequestId(v string) *CreateDBEndpointAddressResponse {
	s.RequestId = &v
	return s
}

type CreateDBClusterEndpointRequest struct {
	AccessKeyId          *string `json:"AccessKeyId" xml:"AccessKeyId"`
	OwnerId              *int64  `json:"OwnerId" xml:"OwnerId"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount" xml:"ResourceOwnerAccount"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId" xml:"ResourceOwnerId"`
	OwnerAccount         *string `json:"OwnerAccount" xml:"OwnerAccount"`
	DBClusterId          *string `json:"DBClusterId" xml:"DBClusterId" require:"true"`
	EndpointType         *string `json:"EndpointType" xml:"EndpointType" require:"true"`
	Nodes                *string `json:"Nodes" xml:"Nodes"`
	ReadWriteMode        *string `json:"ReadWriteMode" xml:"ReadWriteMode"`
	AutoAddNewNodes      *string `json:"AutoAddNewNodes" xml:"AutoAddNewNodes"`
	EndpointConfig       *string `json:"EndpointConfig" xml:"EndpointConfig"`
	ClientToken          *string `json:"ClientToken" xml:"ClientToken"`
}

func (s CreateDBClusterEndpointRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDBClusterEndpointRequest) GoString() string {
	return s.String()
}

func (s *CreateDBClusterEndpointRequest) SetAccessKeyId(v string) *CreateDBClusterEndpointRequest {
	s.AccessKeyId = &v
	return s
}

func (s *CreateDBClusterEndpointRequest) SetOwnerId(v int64) *CreateDBClusterEndpointRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateDBClusterEndpointRequest) SetResourceOwnerAccount(v string) *CreateDBClusterEndpointRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateDBClusterEndpointRequest) SetResourceOwnerId(v int64) *CreateDBClusterEndpointRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateDBClusterEndpointRequest) SetOwnerAccount(v string) *CreateDBClusterEndpointRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateDBClusterEndpointRequest) SetDBClusterId(v string) *CreateDBClusterEndpointRequest {
	s.DBClusterId = &v
	return s
}

func (s *CreateDBClusterEndpointRequest) SetEndpointType(v string) *CreateDBClusterEndpointRequest {
	s.EndpointType = &v
	return s
}

func (s *CreateDBClusterEndpointRequest) SetNodes(v string) *CreateDBClusterEndpointRequest {
	s.Nodes = &v
	return s
}

func (s *CreateDBClusterEndpointRequest) SetReadWriteMode(v string) *CreateDBClusterEndpointRequest {
	s.ReadWriteMode = &v
	return s
}

func (s *CreateDBClusterEndpointRequest) SetAutoAddNewNodes(v string) *CreateDBClusterEndpointRequest {
	s.AutoAddNewNodes = &v
	return s
}

func (s *CreateDBClusterEndpointRequest) SetEndpointConfig(v string) *CreateDBClusterEndpointRequest {
	s.EndpointConfig = &v
	return s
}

func (s *CreateDBClusterEndpointRequest) SetClientToken(v string) *CreateDBClusterEndpointRequest {
	s.ClientToken = &v
	return s
}

type CreateDBClusterEndpointResponse struct {
	RequestId *string `json:"RequestId" xml:"RequestId" require:"true"`
}

func (s CreateDBClusterEndpointResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDBClusterEndpointResponse) GoString() string {
	return s.String()
}

func (s *CreateDBClusterEndpointResponse) SetRequestId(v string) *CreateDBClusterEndpointResponse {
	s.RequestId = &v
	return s
}

type RestartDBNodeRequest struct {
	AccessKeyId          *string `json:"AccessKeyId" xml:"AccessKeyId"`
	OwnerId              *int64  `json:"OwnerId" xml:"OwnerId"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount" xml:"ResourceOwnerAccount"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId" xml:"ResourceOwnerId"`
	OwnerAccount         *string `json:"OwnerAccount" xml:"OwnerAccount"`
	DBNodeId             *string `json:"DBNodeId" xml:"DBNodeId" require:"true"`
}

func (s RestartDBNodeRequest) String() string {
	return tea.Prettify(s)
}

func (s RestartDBNodeRequest) GoString() string {
	return s.String()
}

func (s *RestartDBNodeRequest) SetAccessKeyId(v string) *RestartDBNodeRequest {
	s.AccessKeyId = &v
	return s
}

func (s *RestartDBNodeRequest) SetOwnerId(v int64) *RestartDBNodeRequest {
	s.OwnerId = &v
	return s
}

func (s *RestartDBNodeRequest) SetResourceOwnerAccount(v string) *RestartDBNodeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *RestartDBNodeRequest) SetResourceOwnerId(v int64) *RestartDBNodeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *RestartDBNodeRequest) SetOwnerAccount(v string) *RestartDBNodeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *RestartDBNodeRequest) SetDBNodeId(v string) *RestartDBNodeRequest {
	s.DBNodeId = &v
	return s
}

type RestartDBNodeResponse struct {
	RequestId *string `json:"RequestId" xml:"RequestId" require:"true"`
}

func (s RestartDBNodeResponse) String() string {
	return tea.Prettify(s)
}

func (s RestartDBNodeResponse) GoString() string {
	return s.String()
}

func (s *RestartDBNodeResponse) SetRequestId(v string) *RestartDBNodeResponse {
	s.RequestId = &v
	return s
}

type DescribeDBNodePerformanceRequest struct {
	AccessKeyId *string `json:"AccessKeyId" xml:"AccessKeyId"`
	DBNodeId    *string `json:"DBNodeId" xml:"DBNodeId" require:"true"`
	Key         *string `json:"Key" xml:"Key" require:"true"`
	StartTime   *string `json:"StartTime" xml:"StartTime" require:"true"`
	EndTime     *string `json:"EndTime" xml:"EndTime" require:"true"`
	DBClusterId *string `json:"DBClusterId" xml:"DBClusterId"`
}

func (s DescribeDBNodePerformanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBNodePerformanceRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBNodePerformanceRequest) SetAccessKeyId(v string) *DescribeDBNodePerformanceRequest {
	s.AccessKeyId = &v
	return s
}

func (s *DescribeDBNodePerformanceRequest) SetDBNodeId(v string) *DescribeDBNodePerformanceRequest {
	s.DBNodeId = &v
	return s
}

func (s *DescribeDBNodePerformanceRequest) SetKey(v string) *DescribeDBNodePerformanceRequest {
	s.Key = &v
	return s
}

func (s *DescribeDBNodePerformanceRequest) SetStartTime(v string) *DescribeDBNodePerformanceRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDBNodePerformanceRequest) SetEndTime(v string) *DescribeDBNodePerformanceRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDBNodePerformanceRequest) SetDBClusterId(v string) *DescribeDBNodePerformanceRequest {
	s.DBClusterId = &v
	return s
}

type DescribeDBNodePerformanceResponse struct {
	RequestId       *string                                           `json:"RequestId" xml:"RequestId" require:"true"`
	DBNodeId        *string                                           `json:"DBNodeId" xml:"DBNodeId" require:"true"`
	Engine          *string                                           `json:"Engine" xml:"Engine" require:"true"`
	DBType          *string                                           `json:"DBType" xml:"DBType" require:"true"`
	DBVersion       *string                                           `json:"DBVersion" xml:"DBVersion" require:"true"`
	StartTime       *string                                           `json:"StartTime" xml:"StartTime" require:"true"`
	EndTime         *string                                           `json:"EndTime" xml:"EndTime" require:"true"`
	PerformanceKeys *DescribeDBNodePerformanceResponsePerformanceKeys `json:"PerformanceKeys" xml:"PerformanceKeys" require:"true" type:"Struct"`
}

func (s DescribeDBNodePerformanceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBNodePerformanceResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBNodePerformanceResponse) SetRequestId(v string) *DescribeDBNodePerformanceResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeDBNodePerformanceResponse) SetDBNodeId(v string) *DescribeDBNodePerformanceResponse {
	s.DBNodeId = &v
	return s
}

func (s *DescribeDBNodePerformanceResponse) SetEngine(v string) *DescribeDBNodePerformanceResponse {
	s.Engine = &v
	return s
}

func (s *DescribeDBNodePerformanceResponse) SetDBType(v string) *DescribeDBNodePerformanceResponse {
	s.DBType = &v
	return s
}

func (s *DescribeDBNodePerformanceResponse) SetDBVersion(v string) *DescribeDBNodePerformanceResponse {
	s.DBVersion = &v
	return s
}

func (s *DescribeDBNodePerformanceResponse) SetStartTime(v string) *DescribeDBNodePerformanceResponse {
	s.StartTime = &v
	return s
}

func (s *DescribeDBNodePerformanceResponse) SetEndTime(v string) *DescribeDBNodePerformanceResponse {
	s.EndTime = &v
	return s
}

func (s *DescribeDBNodePerformanceResponse) SetPerformanceKeys(v *DescribeDBNodePerformanceResponsePerformanceKeys) *DescribeDBNodePerformanceResponse {
	s.PerformanceKeys = v
	return s
}

type DescribeDBNodePerformanceResponsePerformanceKeys struct {
	PerformanceItem []*DescribeDBNodePerformanceResponsePerformanceKeysPerformanceItem `json:"PerformanceItem" xml:"PerformanceItem" require:"true" type:"Repeated"`
}

func (s DescribeDBNodePerformanceResponsePerformanceKeys) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBNodePerformanceResponsePerformanceKeys) GoString() string {
	return s.String()
}

func (s *DescribeDBNodePerformanceResponsePerformanceKeys) SetPerformanceItem(v []*DescribeDBNodePerformanceResponsePerformanceKeysPerformanceItem) *DescribeDBNodePerformanceResponsePerformanceKeys {
	s.PerformanceItem = v
	return s
}

type DescribeDBNodePerformanceResponsePerformanceKeysPerformanceItem struct {
	MetricName  *string                                                                `json:"MetricName" xml:"MetricName" require:"true"`
	Measurement *string                                                                `json:"Measurement" xml:"Measurement" require:"true"`
	Points      *DescribeDBNodePerformanceResponsePerformanceKeysPerformanceItemPoints `json:"Points" xml:"Points" require:"true" type:"Struct"`
}

func (s DescribeDBNodePerformanceResponsePerformanceKeysPerformanceItem) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBNodePerformanceResponsePerformanceKeysPerformanceItem) GoString() string {
	return s.String()
}

func (s *DescribeDBNodePerformanceResponsePerformanceKeysPerformanceItem) SetMetricName(v string) *DescribeDBNodePerformanceResponsePerformanceKeysPerformanceItem {
	s.MetricName = &v
	return s
}

func (s *DescribeDBNodePerformanceResponsePerformanceKeysPerformanceItem) SetMeasurement(v string) *DescribeDBNodePerformanceResponsePerformanceKeysPerformanceItem {
	s.Measurement = &v
	return s
}

func (s *DescribeDBNodePerformanceResponsePerformanceKeysPerformanceItem) SetPoints(v *DescribeDBNodePerformanceResponsePerformanceKeysPerformanceItemPoints) *DescribeDBNodePerformanceResponsePerformanceKeysPerformanceItem {
	s.Points = v
	return s
}

type DescribeDBNodePerformanceResponsePerformanceKeysPerformanceItemPoints struct {
	PerformanceItemValue []*DescribeDBNodePerformanceResponsePerformanceKeysPerformanceItemPointsPerformanceItemValue `json:"PerformanceItemValue" xml:"PerformanceItemValue" require:"true" type:"Repeated"`
}

func (s DescribeDBNodePerformanceResponsePerformanceKeysPerformanceItemPoints) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBNodePerformanceResponsePerformanceKeysPerformanceItemPoints) GoString() string {
	return s.String()
}

func (s *DescribeDBNodePerformanceResponsePerformanceKeysPerformanceItemPoints) SetPerformanceItemValue(v []*DescribeDBNodePerformanceResponsePerformanceKeysPerformanceItemPointsPerformanceItemValue) *DescribeDBNodePerformanceResponsePerformanceKeysPerformanceItemPoints {
	s.PerformanceItemValue = v
	return s
}

type DescribeDBNodePerformanceResponsePerformanceKeysPerformanceItemPointsPerformanceItemValue struct {
	Value     *string `json:"Value" xml:"Value" require:"true"`
	Timestamp *int64  `json:"Timestamp" xml:"Timestamp" require:"true"`
}

func (s DescribeDBNodePerformanceResponsePerformanceKeysPerformanceItemPointsPerformanceItemValue) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBNodePerformanceResponsePerformanceKeysPerformanceItemPointsPerformanceItemValue) GoString() string {
	return s.String()
}

func (s *DescribeDBNodePerformanceResponsePerformanceKeysPerformanceItemPointsPerformanceItemValue) SetValue(v string) *DescribeDBNodePerformanceResponsePerformanceKeysPerformanceItemPointsPerformanceItemValue {
	s.Value = &v
	return s
}

func (s *DescribeDBNodePerformanceResponsePerformanceKeysPerformanceItemPointsPerformanceItemValue) SetTimestamp(v int64) *DescribeDBNodePerformanceResponsePerformanceKeysPerformanceItemPointsPerformanceItemValue {
	s.Timestamp = &v
	return s
}

type ModifyDBClusterAccessWhitelistRequest struct {
	AccessKeyId               *string `json:"AccessKeyId" xml:"AccessKeyId"`
	OwnerId                   *int64  `json:"OwnerId" xml:"OwnerId"`
	ResourceOwnerAccount      *string `json:"ResourceOwnerAccount" xml:"ResourceOwnerAccount"`
	ResourceOwnerId           *int64  `json:"ResourceOwnerId" xml:"ResourceOwnerId"`
	OwnerAccount              *string `json:"OwnerAccount" xml:"OwnerAccount"`
	DBClusterId               *string `json:"DBClusterId" xml:"DBClusterId" require:"true"`
	SecurityIps               *string `json:"SecurityIps" xml:"SecurityIps"`
	DBClusterIPArrayName      *string `json:"DBClusterIPArrayName" xml:"DBClusterIPArrayName"`
	DBClusterIPArrayAttribute *string `json:"DBClusterIPArrayAttribute" xml:"DBClusterIPArrayAttribute"`
	WhiteListType             *string `json:"WhiteListType" xml:"WhiteListType"`
	SecurityGroupIds          *string `json:"SecurityGroupIds" xml:"SecurityGroupIds"`
	ModifyMode                *string `json:"ModifyMode" xml:"ModifyMode"`
}

func (s ModifyDBClusterAccessWhitelistRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBClusterAccessWhitelistRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterAccessWhitelistRequest) SetAccessKeyId(v string) *ModifyDBClusterAccessWhitelistRequest {
	s.AccessKeyId = &v
	return s
}

func (s *ModifyDBClusterAccessWhitelistRequest) SetOwnerId(v int64) *ModifyDBClusterAccessWhitelistRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDBClusterAccessWhitelistRequest) SetResourceOwnerAccount(v string) *ModifyDBClusterAccessWhitelistRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDBClusterAccessWhitelistRequest) SetResourceOwnerId(v int64) *ModifyDBClusterAccessWhitelistRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDBClusterAccessWhitelistRequest) SetOwnerAccount(v string) *ModifyDBClusterAccessWhitelistRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyDBClusterAccessWhitelistRequest) SetDBClusterId(v string) *ModifyDBClusterAccessWhitelistRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyDBClusterAccessWhitelistRequest) SetSecurityIps(v string) *ModifyDBClusterAccessWhitelistRequest {
	s.SecurityIps = &v
	return s
}

func (s *ModifyDBClusterAccessWhitelistRequest) SetDBClusterIPArrayName(v string) *ModifyDBClusterAccessWhitelistRequest {
	s.DBClusterIPArrayName = &v
	return s
}

func (s *ModifyDBClusterAccessWhitelistRequest) SetDBClusterIPArrayAttribute(v string) *ModifyDBClusterAccessWhitelistRequest {
	s.DBClusterIPArrayAttribute = &v
	return s
}

func (s *ModifyDBClusterAccessWhitelistRequest) SetWhiteListType(v string) *ModifyDBClusterAccessWhitelistRequest {
	s.WhiteListType = &v
	return s
}

func (s *ModifyDBClusterAccessWhitelistRequest) SetSecurityGroupIds(v string) *ModifyDBClusterAccessWhitelistRequest {
	s.SecurityGroupIds = &v
	return s
}

func (s *ModifyDBClusterAccessWhitelistRequest) SetModifyMode(v string) *ModifyDBClusterAccessWhitelistRequest {
	s.ModifyMode = &v
	return s
}

type ModifyDBClusterAccessWhitelistResponse struct {
	RequestId *string `json:"RequestId" xml:"RequestId" require:"true"`
}

func (s ModifyDBClusterAccessWhitelistResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBClusterAccessWhitelistResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterAccessWhitelistResponse) SetRequestId(v string) *ModifyDBClusterAccessWhitelistResponse {
	s.RequestId = &v
	return s
}

type DescribeDBClusterAccessWhitelistRequest struct {
	AccessKeyId          *string `json:"AccessKeyId" xml:"AccessKeyId"`
	OwnerId              *int64  `json:"OwnerId" xml:"OwnerId"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount" xml:"ResourceOwnerAccount"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId" xml:"ResourceOwnerId"`
	OwnerAccount         *string `json:"OwnerAccount" xml:"OwnerAccount"`
	DBClusterId          *string `json:"DBClusterId" xml:"DBClusterId" require:"true"`
}

func (s DescribeDBClusterAccessWhitelistRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterAccessWhitelistRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterAccessWhitelistRequest) SetAccessKeyId(v string) *DescribeDBClusterAccessWhitelistRequest {
	s.AccessKeyId = &v
	return s
}

func (s *DescribeDBClusterAccessWhitelistRequest) SetOwnerId(v int64) *DescribeDBClusterAccessWhitelistRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDBClusterAccessWhitelistRequest) SetResourceOwnerAccount(v string) *DescribeDBClusterAccessWhitelistRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDBClusterAccessWhitelistRequest) SetResourceOwnerId(v int64) *DescribeDBClusterAccessWhitelistRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDBClusterAccessWhitelistRequest) SetOwnerAccount(v string) *DescribeDBClusterAccessWhitelistRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDBClusterAccessWhitelistRequest) SetDBClusterId(v string) *DescribeDBClusterAccessWhitelistRequest {
	s.DBClusterId = &v
	return s
}

type DescribeDBClusterAccessWhitelistResponse struct {
	RequestId               *string                                                          `json:"RequestId" xml:"RequestId" require:"true"`
	Items                   *DescribeDBClusterAccessWhitelistResponseItems                   `json:"Items" xml:"Items" require:"true" type:"Struct"`
	DBClusterSecurityGroups *DescribeDBClusterAccessWhitelistResponseDBClusterSecurityGroups `json:"DBClusterSecurityGroups" xml:"DBClusterSecurityGroups" require:"true" type:"Struct"`
}

func (s DescribeDBClusterAccessWhitelistResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterAccessWhitelistResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterAccessWhitelistResponse) SetRequestId(v string) *DescribeDBClusterAccessWhitelistResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeDBClusterAccessWhitelistResponse) SetItems(v *DescribeDBClusterAccessWhitelistResponseItems) *DescribeDBClusterAccessWhitelistResponse {
	s.Items = v
	return s
}

func (s *DescribeDBClusterAccessWhitelistResponse) SetDBClusterSecurityGroups(v *DescribeDBClusterAccessWhitelistResponseDBClusterSecurityGroups) *DescribeDBClusterAccessWhitelistResponse {
	s.DBClusterSecurityGroups = v
	return s
}

type DescribeDBClusterAccessWhitelistResponseItems struct {
	DBClusterIPArray []*DescribeDBClusterAccessWhitelistResponseItemsDBClusterIPArray `json:"DBClusterIPArray" xml:"DBClusterIPArray" require:"true" type:"Repeated"`
}

func (s DescribeDBClusterAccessWhitelistResponseItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterAccessWhitelistResponseItems) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterAccessWhitelistResponseItems) SetDBClusterIPArray(v []*DescribeDBClusterAccessWhitelistResponseItemsDBClusterIPArray) *DescribeDBClusterAccessWhitelistResponseItems {
	s.DBClusterIPArray = v
	return s
}

type DescribeDBClusterAccessWhitelistResponseItemsDBClusterIPArray struct {
	DBClusterIPArrayName      *string `json:"DBClusterIPArrayName" xml:"DBClusterIPArrayName" require:"true"`
	DBClusterIPArrayAttribute *string `json:"DBClusterIPArrayAttribute" xml:"DBClusterIPArrayAttribute" require:"true"`
	SecurityIps               *string `json:"SecurityIps" xml:"SecurityIps" require:"true"`
}

func (s DescribeDBClusterAccessWhitelistResponseItemsDBClusterIPArray) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterAccessWhitelistResponseItemsDBClusterIPArray) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterAccessWhitelistResponseItemsDBClusterIPArray) SetDBClusterIPArrayName(v string) *DescribeDBClusterAccessWhitelistResponseItemsDBClusterIPArray {
	s.DBClusterIPArrayName = &v
	return s
}

func (s *DescribeDBClusterAccessWhitelistResponseItemsDBClusterIPArray) SetDBClusterIPArrayAttribute(v string) *DescribeDBClusterAccessWhitelistResponseItemsDBClusterIPArray {
	s.DBClusterIPArrayAttribute = &v
	return s
}

func (s *DescribeDBClusterAccessWhitelistResponseItemsDBClusterIPArray) SetSecurityIps(v string) *DescribeDBClusterAccessWhitelistResponseItemsDBClusterIPArray {
	s.SecurityIps = &v
	return s
}

type DescribeDBClusterAccessWhitelistResponseDBClusterSecurityGroups struct {
	DBClusterSecurityGroup []*DescribeDBClusterAccessWhitelistResponseDBClusterSecurityGroupsDBClusterSecurityGroup `json:"DBClusterSecurityGroup" xml:"DBClusterSecurityGroup" require:"true" type:"Repeated"`
}

func (s DescribeDBClusterAccessWhitelistResponseDBClusterSecurityGroups) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterAccessWhitelistResponseDBClusterSecurityGroups) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterAccessWhitelistResponseDBClusterSecurityGroups) SetDBClusterSecurityGroup(v []*DescribeDBClusterAccessWhitelistResponseDBClusterSecurityGroupsDBClusterSecurityGroup) *DescribeDBClusterAccessWhitelistResponseDBClusterSecurityGroups {
	s.DBClusterSecurityGroup = v
	return s
}

type DescribeDBClusterAccessWhitelistResponseDBClusterSecurityGroupsDBClusterSecurityGroup struct {
	SecurityGroupId   *string `json:"SecurityGroupId" xml:"SecurityGroupId" require:"true"`
	SecurityGroupName *string `json:"SecurityGroupName" xml:"SecurityGroupName" require:"true"`
}

func (s DescribeDBClusterAccessWhitelistResponseDBClusterSecurityGroupsDBClusterSecurityGroup) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterAccessWhitelistResponseDBClusterSecurityGroupsDBClusterSecurityGroup) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterAccessWhitelistResponseDBClusterSecurityGroupsDBClusterSecurityGroup) SetSecurityGroupId(v string) *DescribeDBClusterAccessWhitelistResponseDBClusterSecurityGroupsDBClusterSecurityGroup {
	s.SecurityGroupId = &v
	return s
}

func (s *DescribeDBClusterAccessWhitelistResponseDBClusterSecurityGroupsDBClusterSecurityGroup) SetSecurityGroupName(v string) *DescribeDBClusterAccessWhitelistResponseDBClusterSecurityGroupsDBClusterSecurityGroup {
	s.SecurityGroupName = &v
	return s
}

type ModifyDBClusterMaintainTimeRequest struct {
	AccessKeyId          *string `json:"AccessKeyId" xml:"AccessKeyId"`
	OwnerId              *int64  `json:"OwnerId" xml:"OwnerId"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount" xml:"ResourceOwnerAccount"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId" xml:"ResourceOwnerId"`
	OwnerAccount         *string `json:"OwnerAccount" xml:"OwnerAccount"`
	DBClusterId          *string `json:"DBClusterId" xml:"DBClusterId" require:"true"`
	MaintainTime         *string `json:"MaintainTime" xml:"MaintainTime" require:"true"`
}

func (s ModifyDBClusterMaintainTimeRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBClusterMaintainTimeRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterMaintainTimeRequest) SetAccessKeyId(v string) *ModifyDBClusterMaintainTimeRequest {
	s.AccessKeyId = &v
	return s
}

func (s *ModifyDBClusterMaintainTimeRequest) SetOwnerId(v int64) *ModifyDBClusterMaintainTimeRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDBClusterMaintainTimeRequest) SetResourceOwnerAccount(v string) *ModifyDBClusterMaintainTimeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDBClusterMaintainTimeRequest) SetResourceOwnerId(v int64) *ModifyDBClusterMaintainTimeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDBClusterMaintainTimeRequest) SetOwnerAccount(v string) *ModifyDBClusterMaintainTimeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyDBClusterMaintainTimeRequest) SetDBClusterId(v string) *ModifyDBClusterMaintainTimeRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyDBClusterMaintainTimeRequest) SetMaintainTime(v string) *ModifyDBClusterMaintainTimeRequest {
	s.MaintainTime = &v
	return s
}

type ModifyDBClusterMaintainTimeResponse struct {
	RequestId *string `json:"RequestId" xml:"RequestId" require:"true"`
}

func (s ModifyDBClusterMaintainTimeResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBClusterMaintainTimeResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterMaintainTimeResponse) SetRequestId(v string) *ModifyDBClusterMaintainTimeResponse {
	s.RequestId = &v
	return s
}

type RevokeAccountPrivilegeRequest struct {
	AccessKeyId          *string `json:"AccessKeyId" xml:"AccessKeyId"`
	OwnerId              *int64  `json:"OwnerId" xml:"OwnerId"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount" xml:"ResourceOwnerAccount"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId" xml:"ResourceOwnerId"`
	OwnerAccount         *string `json:"OwnerAccount" xml:"OwnerAccount"`
	DBClusterId          *string `json:"DBClusterId" xml:"DBClusterId" require:"true"`
	AccountName          *string `json:"AccountName" xml:"AccountName" require:"true"`
	DBName               *string `json:"DBName" xml:"DBName" require:"true"`
}

func (s RevokeAccountPrivilegeRequest) String() string {
	return tea.Prettify(s)
}

func (s RevokeAccountPrivilegeRequest) GoString() string {
	return s.String()
}

func (s *RevokeAccountPrivilegeRequest) SetAccessKeyId(v string) *RevokeAccountPrivilegeRequest {
	s.AccessKeyId = &v
	return s
}

func (s *RevokeAccountPrivilegeRequest) SetOwnerId(v int64) *RevokeAccountPrivilegeRequest {
	s.OwnerId = &v
	return s
}

func (s *RevokeAccountPrivilegeRequest) SetResourceOwnerAccount(v string) *RevokeAccountPrivilegeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *RevokeAccountPrivilegeRequest) SetResourceOwnerId(v int64) *RevokeAccountPrivilegeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *RevokeAccountPrivilegeRequest) SetOwnerAccount(v string) *RevokeAccountPrivilegeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *RevokeAccountPrivilegeRequest) SetDBClusterId(v string) *RevokeAccountPrivilegeRequest {
	s.DBClusterId = &v
	return s
}

func (s *RevokeAccountPrivilegeRequest) SetAccountName(v string) *RevokeAccountPrivilegeRequest {
	s.AccountName = &v
	return s
}

func (s *RevokeAccountPrivilegeRequest) SetDBName(v string) *RevokeAccountPrivilegeRequest {
	s.DBName = &v
	return s
}

type RevokeAccountPrivilegeResponse struct {
	RequestId *string `json:"RequestId" xml:"RequestId" require:"true"`
}

func (s RevokeAccountPrivilegeResponse) String() string {
	return tea.Prettify(s)
}

func (s RevokeAccountPrivilegeResponse) GoString() string {
	return s.String()
}

func (s *RevokeAccountPrivilegeResponse) SetRequestId(v string) *RevokeAccountPrivilegeResponse {
	s.RequestId = &v
	return s
}

type ResetAccountRequest struct {
	AccessKeyId          *string `json:"AccessKeyId" xml:"AccessKeyId"`
	OwnerId              *int64  `json:"OwnerId" xml:"OwnerId"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount" xml:"ResourceOwnerAccount"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId" xml:"ResourceOwnerId"`
	OwnerAccount         *string `json:"OwnerAccount" xml:"OwnerAccount"`
	DBClusterId          *string `json:"DBClusterId" xml:"DBClusterId" require:"true"`
	AccountName          *string `json:"AccountName" xml:"AccountName" require:"true"`
	AccountPassword      *string `json:"AccountPassword" xml:"AccountPassword" require:"true"`
}

func (s ResetAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s ResetAccountRequest) GoString() string {
	return s.String()
}

func (s *ResetAccountRequest) SetAccessKeyId(v string) *ResetAccountRequest {
	s.AccessKeyId = &v
	return s
}

func (s *ResetAccountRequest) SetOwnerId(v int64) *ResetAccountRequest {
	s.OwnerId = &v
	return s
}

func (s *ResetAccountRequest) SetResourceOwnerAccount(v string) *ResetAccountRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ResetAccountRequest) SetResourceOwnerId(v int64) *ResetAccountRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ResetAccountRequest) SetOwnerAccount(v string) *ResetAccountRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ResetAccountRequest) SetDBClusterId(v string) *ResetAccountRequest {
	s.DBClusterId = &v
	return s
}

func (s *ResetAccountRequest) SetAccountName(v string) *ResetAccountRequest {
	s.AccountName = &v
	return s
}

func (s *ResetAccountRequest) SetAccountPassword(v string) *ResetAccountRequest {
	s.AccountPassword = &v
	return s
}

type ResetAccountResponse struct {
	RequestId *string `json:"RequestId" xml:"RequestId" require:"true"`
}

func (s ResetAccountResponse) String() string {
	return tea.Prettify(s)
}

func (s ResetAccountResponse) GoString() string {
	return s.String()
}

func (s *ResetAccountResponse) SetRequestId(v string) *ResetAccountResponse {
	s.RequestId = &v
	return s
}

type GrantAccountPrivilegeRequest struct {
	AccessKeyId          *string `json:"AccessKeyId" xml:"AccessKeyId"`
	OwnerId              *int64  `json:"OwnerId" xml:"OwnerId"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount" xml:"ResourceOwnerAccount"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId" xml:"ResourceOwnerId"`
	OwnerAccount         *string `json:"OwnerAccount" xml:"OwnerAccount"`
	DBClusterId          *string `json:"DBClusterId" xml:"DBClusterId" require:"true"`
	AccountName          *string `json:"AccountName" xml:"AccountName" require:"true"`
	DBName               *string `json:"DBName" xml:"DBName" require:"true"`
	AccountPrivilege     *string `json:"AccountPrivilege" xml:"AccountPrivilege" require:"true"`
}

func (s GrantAccountPrivilegeRequest) String() string {
	return tea.Prettify(s)
}

func (s GrantAccountPrivilegeRequest) GoString() string {
	return s.String()
}

func (s *GrantAccountPrivilegeRequest) SetAccessKeyId(v string) *GrantAccountPrivilegeRequest {
	s.AccessKeyId = &v
	return s
}

func (s *GrantAccountPrivilegeRequest) SetOwnerId(v int64) *GrantAccountPrivilegeRequest {
	s.OwnerId = &v
	return s
}

func (s *GrantAccountPrivilegeRequest) SetResourceOwnerAccount(v string) *GrantAccountPrivilegeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GrantAccountPrivilegeRequest) SetResourceOwnerId(v int64) *GrantAccountPrivilegeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GrantAccountPrivilegeRequest) SetOwnerAccount(v string) *GrantAccountPrivilegeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *GrantAccountPrivilegeRequest) SetDBClusterId(v string) *GrantAccountPrivilegeRequest {
	s.DBClusterId = &v
	return s
}

func (s *GrantAccountPrivilegeRequest) SetAccountName(v string) *GrantAccountPrivilegeRequest {
	s.AccountName = &v
	return s
}

func (s *GrantAccountPrivilegeRequest) SetDBName(v string) *GrantAccountPrivilegeRequest {
	s.DBName = &v
	return s
}

func (s *GrantAccountPrivilegeRequest) SetAccountPrivilege(v string) *GrantAccountPrivilegeRequest {
	s.AccountPrivilege = &v
	return s
}

type GrantAccountPrivilegeResponse struct {
	RequestId *string `json:"RequestId" xml:"RequestId" require:"true"`
}

func (s GrantAccountPrivilegeResponse) String() string {
	return tea.Prettify(s)
}

func (s GrantAccountPrivilegeResponse) GoString() string {
	return s.String()
}

func (s *GrantAccountPrivilegeResponse) SetRequestId(v string) *GrantAccountPrivilegeResponse {
	s.RequestId = &v
	return s
}

type DescribeDatabasesRequest struct {
	AccessKeyId          *string `json:"AccessKeyId" xml:"AccessKeyId"`
	OwnerId              *int64  `json:"OwnerId" xml:"OwnerId"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount" xml:"ResourceOwnerAccount"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId" xml:"ResourceOwnerId"`
	OwnerAccount         *string `json:"OwnerAccount" xml:"OwnerAccount"`
	DBClusterId          *string `json:"DBClusterId" xml:"DBClusterId" require:"true"`
	DBName               *string `json:"DBName" xml:"DBName"`
}

func (s DescribeDatabasesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDatabasesRequest) GoString() string {
	return s.String()
}

func (s *DescribeDatabasesRequest) SetAccessKeyId(v string) *DescribeDatabasesRequest {
	s.AccessKeyId = &v
	return s
}

func (s *DescribeDatabasesRequest) SetOwnerId(v int64) *DescribeDatabasesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDatabasesRequest) SetResourceOwnerAccount(v string) *DescribeDatabasesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDatabasesRequest) SetResourceOwnerId(v int64) *DescribeDatabasesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDatabasesRequest) SetOwnerAccount(v string) *DescribeDatabasesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDatabasesRequest) SetDBClusterId(v string) *DescribeDatabasesRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDatabasesRequest) SetDBName(v string) *DescribeDatabasesRequest {
	s.DBName = &v
	return s
}

type DescribeDatabasesResponse struct {
	RequestId *string                             `json:"RequestId" xml:"RequestId" require:"true"`
	Databases *DescribeDatabasesResponseDatabases `json:"Databases" xml:"Databases" require:"true" type:"Struct"`
}

func (s DescribeDatabasesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDatabasesResponse) GoString() string {
	return s.String()
}

func (s *DescribeDatabasesResponse) SetRequestId(v string) *DescribeDatabasesResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeDatabasesResponse) SetDatabases(v *DescribeDatabasesResponseDatabases) *DescribeDatabasesResponse {
	s.Databases = v
	return s
}

type DescribeDatabasesResponseDatabases struct {
	Database []*DescribeDatabasesResponseDatabasesDatabase `json:"Database" xml:"Database" require:"true" type:"Repeated"`
}

func (s DescribeDatabasesResponseDatabases) String() string {
	return tea.Prettify(s)
}

func (s DescribeDatabasesResponseDatabases) GoString() string {
	return s.String()
}

func (s *DescribeDatabasesResponseDatabases) SetDatabase(v []*DescribeDatabasesResponseDatabasesDatabase) *DescribeDatabasesResponseDatabases {
	s.Database = v
	return s
}

type DescribeDatabasesResponseDatabasesDatabase struct {
	DBName           *string                                             `json:"DBName" xml:"DBName" require:"true"`
	DBStatus         *string                                             `json:"DBStatus" xml:"DBStatus" require:"true"`
	DBDescription    *string                                             `json:"DBDescription" xml:"DBDescription" require:"true"`
	CharacterSetName *string                                             `json:"CharacterSetName" xml:"CharacterSetName" require:"true"`
	Engine           *string                                             `json:"Engine" xml:"Engine" require:"true"`
	Accounts         *DescribeDatabasesResponseDatabasesDatabaseAccounts `json:"Accounts" xml:"Accounts" require:"true" type:"Struct"`
}

func (s DescribeDatabasesResponseDatabasesDatabase) String() string {
	return tea.Prettify(s)
}

func (s DescribeDatabasesResponseDatabasesDatabase) GoString() string {
	return s.String()
}

func (s *DescribeDatabasesResponseDatabasesDatabase) SetDBName(v string) *DescribeDatabasesResponseDatabasesDatabase {
	s.DBName = &v
	return s
}

func (s *DescribeDatabasesResponseDatabasesDatabase) SetDBStatus(v string) *DescribeDatabasesResponseDatabasesDatabase {
	s.DBStatus = &v
	return s
}

func (s *DescribeDatabasesResponseDatabasesDatabase) SetDBDescription(v string) *DescribeDatabasesResponseDatabasesDatabase {
	s.DBDescription = &v
	return s
}

func (s *DescribeDatabasesResponseDatabasesDatabase) SetCharacterSetName(v string) *DescribeDatabasesResponseDatabasesDatabase {
	s.CharacterSetName = &v
	return s
}

func (s *DescribeDatabasesResponseDatabasesDatabase) SetEngine(v string) *DescribeDatabasesResponseDatabasesDatabase {
	s.Engine = &v
	return s
}

func (s *DescribeDatabasesResponseDatabasesDatabase) SetAccounts(v *DescribeDatabasesResponseDatabasesDatabaseAccounts) *DescribeDatabasesResponseDatabasesDatabase {
	s.Accounts = v
	return s
}

type DescribeDatabasesResponseDatabasesDatabaseAccounts struct {
	Account []*DescribeDatabasesResponseDatabasesDatabaseAccountsAccount `json:"Account" xml:"Account" require:"true" type:"Repeated"`
}

func (s DescribeDatabasesResponseDatabasesDatabaseAccounts) String() string {
	return tea.Prettify(s)
}

func (s DescribeDatabasesResponseDatabasesDatabaseAccounts) GoString() string {
	return s.String()
}

func (s *DescribeDatabasesResponseDatabasesDatabaseAccounts) SetAccount(v []*DescribeDatabasesResponseDatabasesDatabaseAccountsAccount) *DescribeDatabasesResponseDatabasesDatabaseAccounts {
	s.Account = v
	return s
}

type DescribeDatabasesResponseDatabasesDatabaseAccountsAccount struct {
	AccountName      *string `json:"AccountName" xml:"AccountName" require:"true"`
	AccountStatus    *string `json:"AccountStatus" xml:"AccountStatus" require:"true"`
	AccountPrivilege *string `json:"AccountPrivilege" xml:"AccountPrivilege" require:"true"`
	PrivilegeStatus  *string `json:"PrivilegeStatus" xml:"PrivilegeStatus" require:"true"`
}

func (s DescribeDatabasesResponseDatabasesDatabaseAccountsAccount) String() string {
	return tea.Prettify(s)
}

func (s DescribeDatabasesResponseDatabasesDatabaseAccountsAccount) GoString() string {
	return s.String()
}

func (s *DescribeDatabasesResponseDatabasesDatabaseAccountsAccount) SetAccountName(v string) *DescribeDatabasesResponseDatabasesDatabaseAccountsAccount {
	s.AccountName = &v
	return s
}

func (s *DescribeDatabasesResponseDatabasesDatabaseAccountsAccount) SetAccountStatus(v string) *DescribeDatabasesResponseDatabasesDatabaseAccountsAccount {
	s.AccountStatus = &v
	return s
}

func (s *DescribeDatabasesResponseDatabasesDatabaseAccountsAccount) SetAccountPrivilege(v string) *DescribeDatabasesResponseDatabasesDatabaseAccountsAccount {
	s.AccountPrivilege = &v
	return s
}

func (s *DescribeDatabasesResponseDatabasesDatabaseAccountsAccount) SetPrivilegeStatus(v string) *DescribeDatabasesResponseDatabasesDatabaseAccountsAccount {
	s.PrivilegeStatus = &v
	return s
}

type DeleteDatabaseRequest struct {
	AccessKeyId          *string `json:"AccessKeyId" xml:"AccessKeyId"`
	OwnerId              *int64  `json:"OwnerId" xml:"OwnerId"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount" xml:"ResourceOwnerAccount"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId" xml:"ResourceOwnerId"`
	OwnerAccount         *string `json:"OwnerAccount" xml:"OwnerAccount"`
	DBClusterId          *string `json:"DBClusterId" xml:"DBClusterId" require:"true"`
	DBName               *string `json:"DBName" xml:"DBName" require:"true"`
}

func (s DeleteDatabaseRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDatabaseRequest) GoString() string {
	return s.String()
}

func (s *DeleteDatabaseRequest) SetAccessKeyId(v string) *DeleteDatabaseRequest {
	s.AccessKeyId = &v
	return s
}

func (s *DeleteDatabaseRequest) SetOwnerId(v int64) *DeleteDatabaseRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteDatabaseRequest) SetResourceOwnerAccount(v string) *DeleteDatabaseRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteDatabaseRequest) SetResourceOwnerId(v int64) *DeleteDatabaseRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteDatabaseRequest) SetOwnerAccount(v string) *DeleteDatabaseRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteDatabaseRequest) SetDBClusterId(v string) *DeleteDatabaseRequest {
	s.DBClusterId = &v
	return s
}

func (s *DeleteDatabaseRequest) SetDBName(v string) *DeleteDatabaseRequest {
	s.DBName = &v
	return s
}

type DeleteDatabaseResponse struct {
	RequestId *string `json:"RequestId" xml:"RequestId" require:"true"`
}

func (s DeleteDatabaseResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDatabaseResponse) GoString() string {
	return s.String()
}

func (s *DeleteDatabaseResponse) SetRequestId(v string) *DeleteDatabaseResponse {
	s.RequestId = &v
	return s
}

type DeleteAccountRequest struct {
	AccessKeyId          *string `json:"AccessKeyId" xml:"AccessKeyId"`
	OwnerId              *int64  `json:"OwnerId" xml:"OwnerId"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount" xml:"ResourceOwnerAccount"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId" xml:"ResourceOwnerId"`
	OwnerAccount         *string `json:"OwnerAccount" xml:"OwnerAccount"`
	DBClusterId          *string `json:"DBClusterId" xml:"DBClusterId" require:"true"`
	AccountName          *string `json:"AccountName" xml:"AccountName" require:"true"`
}

func (s DeleteAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAccountRequest) GoString() string {
	return s.String()
}

func (s *DeleteAccountRequest) SetAccessKeyId(v string) *DeleteAccountRequest {
	s.AccessKeyId = &v
	return s
}

func (s *DeleteAccountRequest) SetOwnerId(v int64) *DeleteAccountRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteAccountRequest) SetResourceOwnerAccount(v string) *DeleteAccountRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteAccountRequest) SetResourceOwnerId(v int64) *DeleteAccountRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteAccountRequest) SetOwnerAccount(v string) *DeleteAccountRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteAccountRequest) SetDBClusterId(v string) *DeleteAccountRequest {
	s.DBClusterId = &v
	return s
}

func (s *DeleteAccountRequest) SetAccountName(v string) *DeleteAccountRequest {
	s.AccountName = &v
	return s
}

type DeleteAccountResponse struct {
	RequestId *string `json:"RequestId" xml:"RequestId" require:"true"`
}

func (s DeleteAccountResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAccountResponse) GoString() string {
	return s.String()
}

func (s *DeleteAccountResponse) SetRequestId(v string) *DeleteAccountResponse {
	s.RequestId = &v
	return s
}

type CreateDatabaseRequest struct {
	AccessKeyId          *string `json:"AccessKeyId" xml:"AccessKeyId"`
	OwnerId              *int64  `json:"OwnerId" xml:"OwnerId"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount" xml:"ResourceOwnerAccount"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId" xml:"ResourceOwnerId"`
	OwnerAccount         *string `json:"OwnerAccount" xml:"OwnerAccount"`
	DBClusterId          *string `json:"DBClusterId" xml:"DBClusterId" require:"true"`
	DBName               *string `json:"DBName" xml:"DBName" require:"true"`
	CharacterSetName     *string `json:"CharacterSetName" xml:"CharacterSetName" require:"true"`
	DBDescription        *string `json:"DBDescription" xml:"DBDescription"`
	AccountName          *string `json:"AccountName" xml:"AccountName"`
	AccountPrivilege     *string `json:"AccountPrivilege" xml:"AccountPrivilege"`
	Collate              *string `json:"Collate" xml:"Collate"`
	Ctype                *string `json:"Ctype" xml:"Ctype"`
}

func (s CreateDatabaseRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDatabaseRequest) GoString() string {
	return s.String()
}

func (s *CreateDatabaseRequest) SetAccessKeyId(v string) *CreateDatabaseRequest {
	s.AccessKeyId = &v
	return s
}

func (s *CreateDatabaseRequest) SetOwnerId(v int64) *CreateDatabaseRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateDatabaseRequest) SetResourceOwnerAccount(v string) *CreateDatabaseRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateDatabaseRequest) SetResourceOwnerId(v int64) *CreateDatabaseRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateDatabaseRequest) SetOwnerAccount(v string) *CreateDatabaseRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateDatabaseRequest) SetDBClusterId(v string) *CreateDatabaseRequest {
	s.DBClusterId = &v
	return s
}

func (s *CreateDatabaseRequest) SetDBName(v string) *CreateDatabaseRequest {
	s.DBName = &v
	return s
}

func (s *CreateDatabaseRequest) SetCharacterSetName(v string) *CreateDatabaseRequest {
	s.CharacterSetName = &v
	return s
}

func (s *CreateDatabaseRequest) SetDBDescription(v string) *CreateDatabaseRequest {
	s.DBDescription = &v
	return s
}

func (s *CreateDatabaseRequest) SetAccountName(v string) *CreateDatabaseRequest {
	s.AccountName = &v
	return s
}

func (s *CreateDatabaseRequest) SetAccountPrivilege(v string) *CreateDatabaseRequest {
	s.AccountPrivilege = &v
	return s
}

func (s *CreateDatabaseRequest) SetCollate(v string) *CreateDatabaseRequest {
	s.Collate = &v
	return s
}

func (s *CreateDatabaseRequest) SetCtype(v string) *CreateDatabaseRequest {
	s.Ctype = &v
	return s
}

type CreateDatabaseResponse struct {
	RequestId *string `json:"RequestId" xml:"RequestId" require:"true"`
}

func (s CreateDatabaseResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDatabaseResponse) GoString() string {
	return s.String()
}

func (s *CreateDatabaseResponse) SetRequestId(v string) *CreateDatabaseResponse {
	s.RequestId = &v
	return s
}

type DeleteBackupRequest struct {
	AccessKeyId          *string `json:"AccessKeyId" xml:"AccessKeyId"`
	OwnerId              *int64  `json:"OwnerId" xml:"OwnerId"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount" xml:"ResourceOwnerAccount"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId" xml:"ResourceOwnerId"`
	OwnerAccount         *string `json:"OwnerAccount" xml:"OwnerAccount"`
	DBClusterId          *string `json:"DBClusterId" xml:"DBClusterId" require:"true"`
	BackupId             *string `json:"BackupId" xml:"BackupId" require:"true"`
}

func (s DeleteBackupRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteBackupRequest) GoString() string {
	return s.String()
}

func (s *DeleteBackupRequest) SetAccessKeyId(v string) *DeleteBackupRequest {
	s.AccessKeyId = &v
	return s
}

func (s *DeleteBackupRequest) SetOwnerId(v int64) *DeleteBackupRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteBackupRequest) SetResourceOwnerAccount(v string) *DeleteBackupRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteBackupRequest) SetResourceOwnerId(v int64) *DeleteBackupRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteBackupRequest) SetOwnerAccount(v string) *DeleteBackupRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteBackupRequest) SetDBClusterId(v string) *DeleteBackupRequest {
	s.DBClusterId = &v
	return s
}

func (s *DeleteBackupRequest) SetBackupId(v string) *DeleteBackupRequest {
	s.BackupId = &v
	return s
}

type DeleteBackupResponse struct {
	RequestId *string `json:"RequestId" xml:"RequestId" require:"true"`
}

func (s DeleteBackupResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteBackupResponse) GoString() string {
	return s.String()
}

func (s *DeleteBackupResponse) SetRequestId(v string) *DeleteBackupResponse {
	s.RequestId = &v
	return s
}

type DescribeSlowLogsRequest struct {
	AccessKeyId          *string `json:"AccessKeyId" xml:"AccessKeyId"`
	OwnerId              *int64  `json:"OwnerId" xml:"OwnerId"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount" xml:"ResourceOwnerAccount"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId" xml:"ResourceOwnerId"`
	OwnerAccount         *string `json:"OwnerAccount" xml:"OwnerAccount"`
	RegionId             *string `json:"RegionId" xml:"RegionId" require:"true"`
	DBClusterId          *string `json:"DBClusterId" xml:"DBClusterId" require:"true"`
	StartTime            *string `json:"StartTime" xml:"StartTime" require:"true"`
	EndTime              *string `json:"EndTime" xml:"EndTime" require:"true"`
	DBName               *string `json:"DBName" xml:"DBName"`
	PageSize             *int    `json:"PageSize" xml:"PageSize"`
	PageNumber           *int    `json:"PageNumber" xml:"PageNumber"`
}

func (s DescribeSlowLogsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlowLogsRequest) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogsRequest) SetAccessKeyId(v string) *DescribeSlowLogsRequest {
	s.AccessKeyId = &v
	return s
}

func (s *DescribeSlowLogsRequest) SetOwnerId(v int64) *DescribeSlowLogsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeSlowLogsRequest) SetResourceOwnerAccount(v string) *DescribeSlowLogsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeSlowLogsRequest) SetResourceOwnerId(v int64) *DescribeSlowLogsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeSlowLogsRequest) SetOwnerAccount(v string) *DescribeSlowLogsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeSlowLogsRequest) SetRegionId(v string) *DescribeSlowLogsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSlowLogsRequest) SetDBClusterId(v string) *DescribeSlowLogsRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeSlowLogsRequest) SetStartTime(v string) *DescribeSlowLogsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeSlowLogsRequest) SetEndTime(v string) *DescribeSlowLogsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeSlowLogsRequest) SetDBName(v string) *DescribeSlowLogsRequest {
	s.DBName = &v
	return s
}

func (s *DescribeSlowLogsRequest) SetPageSize(v int) *DescribeSlowLogsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSlowLogsRequest) SetPageNumber(v int) *DescribeSlowLogsRequest {
	s.PageNumber = &v
	return s
}

type DescribeSlowLogsResponse struct {
	RequestId        *string                        `json:"RequestId" xml:"RequestId" require:"true"`
	DBClusterId      *string                        `json:"DBClusterId" xml:"DBClusterId" require:"true"`
	StartTime        *string                        `json:"StartTime" xml:"StartTime" require:"true"`
	EndTime          *string                        `json:"EndTime" xml:"EndTime" require:"true"`
	Engine           *string                        `json:"Engine" xml:"Engine" require:"true"`
	TotalRecordCount *int                           `json:"TotalRecordCount" xml:"TotalRecordCount" require:"true"`
	PageNumber       *int                           `json:"PageNumber" xml:"PageNumber" require:"true"`
	PageRecordCount  *int                           `json:"PageRecordCount" xml:"PageRecordCount" require:"true"`
	Items            *DescribeSlowLogsResponseItems `json:"Items" xml:"Items" require:"true" type:"Struct"`
}

func (s DescribeSlowLogsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlowLogsResponse) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogsResponse) SetRequestId(v string) *DescribeSlowLogsResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeSlowLogsResponse) SetDBClusterId(v string) *DescribeSlowLogsResponse {
	s.DBClusterId = &v
	return s
}

func (s *DescribeSlowLogsResponse) SetStartTime(v string) *DescribeSlowLogsResponse {
	s.StartTime = &v
	return s
}

func (s *DescribeSlowLogsResponse) SetEndTime(v string) *DescribeSlowLogsResponse {
	s.EndTime = &v
	return s
}

func (s *DescribeSlowLogsResponse) SetEngine(v string) *DescribeSlowLogsResponse {
	s.Engine = &v
	return s
}

func (s *DescribeSlowLogsResponse) SetTotalRecordCount(v int) *DescribeSlowLogsResponse {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeSlowLogsResponse) SetPageNumber(v int) *DescribeSlowLogsResponse {
	s.PageNumber = &v
	return s
}

func (s *DescribeSlowLogsResponse) SetPageRecordCount(v int) *DescribeSlowLogsResponse {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeSlowLogsResponse) SetItems(v *DescribeSlowLogsResponseItems) *DescribeSlowLogsResponse {
	s.Items = v
	return s
}

type DescribeSlowLogsResponseItems struct {
	SQLSlowLog []*DescribeSlowLogsResponseItemsSQLSlowLog `json:"SQLSlowLog" xml:"SQLSlowLog" require:"true" type:"Repeated"`
}

func (s DescribeSlowLogsResponseItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlowLogsResponseItems) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogsResponseItems) SetSQLSlowLog(v []*DescribeSlowLogsResponseItemsSQLSlowLog) *DescribeSlowLogsResponseItems {
	s.SQLSlowLog = v
	return s
}

type DescribeSlowLogsResponseItemsSQLSlowLog struct {
	DBNodeId             *string `json:"DBNodeId" xml:"DBNodeId" require:"true"`
	ParseMaxRowCount     *int64  `json:"ParseMaxRowCount" xml:"ParseMaxRowCount" require:"true"`
	TotalLockTimes       *int64  `json:"TotalLockTimes" xml:"TotalLockTimes" require:"true"`
	DBName               *string `json:"DBName" xml:"DBName" require:"true"`
	MaxExecutionTime     *int64  `json:"MaxExecutionTime" xml:"MaxExecutionTime" require:"true"`
	SQLHASH              *string `json:"SQLHASH" xml:"SQLHASH" require:"true"`
	SQLText              *string `json:"SQLText" xml:"SQLText" require:"true"`
	CreateTime           *string `json:"CreateTime" xml:"CreateTime" require:"true"`
	TotalExecutionTimes  *int64  `json:"TotalExecutionTimes" xml:"TotalExecutionTimes" require:"true"`
	ReturnTotalRowCounts *int64  `json:"ReturnTotalRowCounts" xml:"ReturnTotalRowCounts" require:"true"`
	TotalExecutionCounts *int64  `json:"TotalExecutionCounts" xml:"TotalExecutionCounts" require:"true"`
	MaxLockTime          *int64  `json:"MaxLockTime" xml:"MaxLockTime" require:"true"`
	ReturnMaxRowCount    *int64  `json:"ReturnMaxRowCount" xml:"ReturnMaxRowCount" require:"true"`
	ParseTotalRowCounts  *int64  `json:"ParseTotalRowCounts" xml:"ParseTotalRowCounts" require:"true"`
}

func (s DescribeSlowLogsResponseItemsSQLSlowLog) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlowLogsResponseItemsSQLSlowLog) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogsResponseItemsSQLSlowLog) SetDBNodeId(v string) *DescribeSlowLogsResponseItemsSQLSlowLog {
	s.DBNodeId = &v
	return s
}

func (s *DescribeSlowLogsResponseItemsSQLSlowLog) SetParseMaxRowCount(v int64) *DescribeSlowLogsResponseItemsSQLSlowLog {
	s.ParseMaxRowCount = &v
	return s
}

func (s *DescribeSlowLogsResponseItemsSQLSlowLog) SetTotalLockTimes(v int64) *DescribeSlowLogsResponseItemsSQLSlowLog {
	s.TotalLockTimes = &v
	return s
}

func (s *DescribeSlowLogsResponseItemsSQLSlowLog) SetDBName(v string) *DescribeSlowLogsResponseItemsSQLSlowLog {
	s.DBName = &v
	return s
}

func (s *DescribeSlowLogsResponseItemsSQLSlowLog) SetMaxExecutionTime(v int64) *DescribeSlowLogsResponseItemsSQLSlowLog {
	s.MaxExecutionTime = &v
	return s
}

func (s *DescribeSlowLogsResponseItemsSQLSlowLog) SetSQLHASH(v string) *DescribeSlowLogsResponseItemsSQLSlowLog {
	s.SQLHASH = &v
	return s
}

func (s *DescribeSlowLogsResponseItemsSQLSlowLog) SetSQLText(v string) *DescribeSlowLogsResponseItemsSQLSlowLog {
	s.SQLText = &v
	return s
}

func (s *DescribeSlowLogsResponseItemsSQLSlowLog) SetCreateTime(v string) *DescribeSlowLogsResponseItemsSQLSlowLog {
	s.CreateTime = &v
	return s
}

func (s *DescribeSlowLogsResponseItemsSQLSlowLog) SetTotalExecutionTimes(v int64) *DescribeSlowLogsResponseItemsSQLSlowLog {
	s.TotalExecutionTimes = &v
	return s
}

func (s *DescribeSlowLogsResponseItemsSQLSlowLog) SetReturnTotalRowCounts(v int64) *DescribeSlowLogsResponseItemsSQLSlowLog {
	s.ReturnTotalRowCounts = &v
	return s
}

func (s *DescribeSlowLogsResponseItemsSQLSlowLog) SetTotalExecutionCounts(v int64) *DescribeSlowLogsResponseItemsSQLSlowLog {
	s.TotalExecutionCounts = &v
	return s
}

func (s *DescribeSlowLogsResponseItemsSQLSlowLog) SetMaxLockTime(v int64) *DescribeSlowLogsResponseItemsSQLSlowLog {
	s.MaxLockTime = &v
	return s
}

func (s *DescribeSlowLogsResponseItemsSQLSlowLog) SetReturnMaxRowCount(v int64) *DescribeSlowLogsResponseItemsSQLSlowLog {
	s.ReturnMaxRowCount = &v
	return s
}

func (s *DescribeSlowLogsResponseItemsSQLSlowLog) SetParseTotalRowCounts(v int64) *DescribeSlowLogsResponseItemsSQLSlowLog {
	s.ParseTotalRowCounts = &v
	return s
}

type DescribeSlowLogRecordsRequest struct {
	AccessKeyId          *string `json:"AccessKeyId" xml:"AccessKeyId"`
	OwnerId              *int64  `json:"OwnerId" xml:"OwnerId"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount" xml:"ResourceOwnerAccount"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId" xml:"ResourceOwnerId"`
	OwnerAccount         *string `json:"OwnerAccount" xml:"OwnerAccount"`
	RegionId             *string `json:"RegionId" xml:"RegionId" require:"true"`
	DBClusterId          *string `json:"DBClusterId" xml:"DBClusterId" require:"true"`
	StartTime            *string `json:"StartTime" xml:"StartTime" require:"true"`
	EndTime              *string `json:"EndTime" xml:"EndTime" require:"true"`
	DBName               *string `json:"DBName" xml:"DBName"`
	PageSize             *int    `json:"PageSize" xml:"PageSize"`
	PageNumber           *int    `json:"PageNumber" xml:"PageNumber"`
	SQLHASH              *string `json:"SQLHASH" xml:"SQLHASH"`
}

func (s DescribeSlowLogRecordsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlowLogRecordsRequest) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogRecordsRequest) SetAccessKeyId(v string) *DescribeSlowLogRecordsRequest {
	s.AccessKeyId = &v
	return s
}

func (s *DescribeSlowLogRecordsRequest) SetOwnerId(v int64) *DescribeSlowLogRecordsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeSlowLogRecordsRequest) SetResourceOwnerAccount(v string) *DescribeSlowLogRecordsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeSlowLogRecordsRequest) SetResourceOwnerId(v int64) *DescribeSlowLogRecordsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeSlowLogRecordsRequest) SetOwnerAccount(v string) *DescribeSlowLogRecordsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeSlowLogRecordsRequest) SetRegionId(v string) *DescribeSlowLogRecordsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSlowLogRecordsRequest) SetDBClusterId(v string) *DescribeSlowLogRecordsRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeSlowLogRecordsRequest) SetStartTime(v string) *DescribeSlowLogRecordsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeSlowLogRecordsRequest) SetEndTime(v string) *DescribeSlowLogRecordsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeSlowLogRecordsRequest) SetDBName(v string) *DescribeSlowLogRecordsRequest {
	s.DBName = &v
	return s
}

func (s *DescribeSlowLogRecordsRequest) SetPageSize(v int) *DescribeSlowLogRecordsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSlowLogRecordsRequest) SetPageNumber(v int) *DescribeSlowLogRecordsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeSlowLogRecordsRequest) SetSQLHASH(v string) *DescribeSlowLogRecordsRequest {
	s.SQLHASH = &v
	return s
}

type DescribeSlowLogRecordsResponse struct {
	RequestId        *string                              `json:"RequestId" xml:"RequestId" require:"true"`
	DBClusterId      *string                              `json:"DBClusterId" xml:"DBClusterId" require:"true"`
	Engine           *string                              `json:"Engine" xml:"Engine" require:"true"`
	TotalRecordCount *int                                 `json:"TotalRecordCount" xml:"TotalRecordCount" require:"true"`
	PageNumber       *int                                 `json:"PageNumber" xml:"PageNumber" require:"true"`
	PageRecordCount  *int                                 `json:"PageRecordCount" xml:"PageRecordCount" require:"true"`
	Items            *DescribeSlowLogRecordsResponseItems `json:"Items" xml:"Items" require:"true" type:"Struct"`
}

func (s DescribeSlowLogRecordsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlowLogRecordsResponse) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogRecordsResponse) SetRequestId(v string) *DescribeSlowLogRecordsResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeSlowLogRecordsResponse) SetDBClusterId(v string) *DescribeSlowLogRecordsResponse {
	s.DBClusterId = &v
	return s
}

func (s *DescribeSlowLogRecordsResponse) SetEngine(v string) *DescribeSlowLogRecordsResponse {
	s.Engine = &v
	return s
}

func (s *DescribeSlowLogRecordsResponse) SetTotalRecordCount(v int) *DescribeSlowLogRecordsResponse {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeSlowLogRecordsResponse) SetPageNumber(v int) *DescribeSlowLogRecordsResponse {
	s.PageNumber = &v
	return s
}

func (s *DescribeSlowLogRecordsResponse) SetPageRecordCount(v int) *DescribeSlowLogRecordsResponse {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeSlowLogRecordsResponse) SetItems(v *DescribeSlowLogRecordsResponseItems) *DescribeSlowLogRecordsResponse {
	s.Items = v
	return s
}

type DescribeSlowLogRecordsResponseItems struct {
	SQLSlowRecord []*DescribeSlowLogRecordsResponseItemsSQLSlowRecord `json:"SQLSlowRecord" xml:"SQLSlowRecord" require:"true" type:"Repeated"`
}

func (s DescribeSlowLogRecordsResponseItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlowLogRecordsResponseItems) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogRecordsResponseItems) SetSQLSlowRecord(v []*DescribeSlowLogRecordsResponseItemsSQLSlowRecord) *DescribeSlowLogRecordsResponseItems {
	s.SQLSlowRecord = v
	return s
}

type DescribeSlowLogRecordsResponseItemsSQLSlowRecord struct {
	HostAddress        *string `json:"HostAddress" xml:"HostAddress" require:"true"`
	DBName             *string `json:"DBName" xml:"DBName" require:"true"`
	SQLText            *string `json:"SQLText" xml:"SQLText" require:"true"`
	QueryTimes         *int64  `json:"QueryTimes" xml:"QueryTimes" require:"true"`
	LockTimes          *int64  `json:"LockTimes" xml:"LockTimes" require:"true"`
	ParseRowCounts     *int64  `json:"ParseRowCounts" xml:"ParseRowCounts" require:"true"`
	ReturnRowCounts    *int64  `json:"ReturnRowCounts" xml:"ReturnRowCounts" require:"true"`
	ExecutionStartTime *string `json:"ExecutionStartTime" xml:"ExecutionStartTime" require:"true"`
	DBNodeId           *string `json:"DBNodeId" xml:"DBNodeId" require:"true"`
}

func (s DescribeSlowLogRecordsResponseItemsSQLSlowRecord) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlowLogRecordsResponseItemsSQLSlowRecord) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogRecordsResponseItemsSQLSlowRecord) SetHostAddress(v string) *DescribeSlowLogRecordsResponseItemsSQLSlowRecord {
	s.HostAddress = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseItemsSQLSlowRecord) SetDBName(v string) *DescribeSlowLogRecordsResponseItemsSQLSlowRecord {
	s.DBName = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseItemsSQLSlowRecord) SetSQLText(v string) *DescribeSlowLogRecordsResponseItemsSQLSlowRecord {
	s.SQLText = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseItemsSQLSlowRecord) SetQueryTimes(v int64) *DescribeSlowLogRecordsResponseItemsSQLSlowRecord {
	s.QueryTimes = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseItemsSQLSlowRecord) SetLockTimes(v int64) *DescribeSlowLogRecordsResponseItemsSQLSlowRecord {
	s.LockTimes = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseItemsSQLSlowRecord) SetParseRowCounts(v int64) *DescribeSlowLogRecordsResponseItemsSQLSlowRecord {
	s.ParseRowCounts = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseItemsSQLSlowRecord) SetReturnRowCounts(v int64) *DescribeSlowLogRecordsResponseItemsSQLSlowRecord {
	s.ReturnRowCounts = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseItemsSQLSlowRecord) SetExecutionStartTime(v string) *DescribeSlowLogRecordsResponseItemsSQLSlowRecord {
	s.ExecutionStartTime = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseItemsSQLSlowRecord) SetDBNodeId(v string) *DescribeSlowLogRecordsResponseItemsSQLSlowRecord {
	s.DBNodeId = &v
	return s
}

type FailoverDBClusterRequest struct {
	AccessKeyId          *string `json:"AccessKeyId" xml:"AccessKeyId"`
	OwnerId              *int64  `json:"OwnerId" xml:"OwnerId"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount" xml:"ResourceOwnerAccount"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId" xml:"ResourceOwnerId"`
	OwnerAccount         *string `json:"OwnerAccount" xml:"OwnerAccount"`
	DBClusterId          *string `json:"DBClusterId" xml:"DBClusterId" require:"true"`
	TargetDBNodeId       *string `json:"TargetDBNodeId" xml:"TargetDBNodeId"`
	ClientToken          *string `json:"ClientToken" xml:"ClientToken"`
}

func (s FailoverDBClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s FailoverDBClusterRequest) GoString() string {
	return s.String()
}

func (s *FailoverDBClusterRequest) SetAccessKeyId(v string) *FailoverDBClusterRequest {
	s.AccessKeyId = &v
	return s
}

func (s *FailoverDBClusterRequest) SetOwnerId(v int64) *FailoverDBClusterRequest {
	s.OwnerId = &v
	return s
}

func (s *FailoverDBClusterRequest) SetResourceOwnerAccount(v string) *FailoverDBClusterRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *FailoverDBClusterRequest) SetResourceOwnerId(v int64) *FailoverDBClusterRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *FailoverDBClusterRequest) SetOwnerAccount(v string) *FailoverDBClusterRequest {
	s.OwnerAccount = &v
	return s
}

func (s *FailoverDBClusterRequest) SetDBClusterId(v string) *FailoverDBClusterRequest {
	s.DBClusterId = &v
	return s
}

func (s *FailoverDBClusterRequest) SetTargetDBNodeId(v string) *FailoverDBClusterRequest {
	s.TargetDBNodeId = &v
	return s
}

func (s *FailoverDBClusterRequest) SetClientToken(v string) *FailoverDBClusterRequest {
	s.ClientToken = &v
	return s
}

type FailoverDBClusterResponse struct {
	RequestId *string `json:"RequestId" xml:"RequestId" require:"true"`
}

func (s FailoverDBClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s FailoverDBClusterResponse) GoString() string {
	return s.String()
}

func (s *FailoverDBClusterResponse) SetRequestId(v string) *FailoverDBClusterResponse {
	s.RequestId = &v
	return s
}

type ModifyDBClusterDescriptionRequest struct {
	AccessKeyId          *string `json:"AccessKeyId" xml:"AccessKeyId"`
	OwnerId              *int64  `json:"OwnerId" xml:"OwnerId"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount" xml:"ResourceOwnerAccount"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId" xml:"ResourceOwnerId"`
	OwnerAccount         *string `json:"OwnerAccount" xml:"OwnerAccount"`
	DBClusterId          *string `json:"DBClusterId" xml:"DBClusterId" require:"true"`
	DBClusterDescription *string `json:"DBClusterDescription" xml:"DBClusterDescription" require:"true"`
}

func (s ModifyDBClusterDescriptionRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBClusterDescriptionRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterDescriptionRequest) SetAccessKeyId(v string) *ModifyDBClusterDescriptionRequest {
	s.AccessKeyId = &v
	return s
}

func (s *ModifyDBClusterDescriptionRequest) SetOwnerId(v int64) *ModifyDBClusterDescriptionRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDBClusterDescriptionRequest) SetResourceOwnerAccount(v string) *ModifyDBClusterDescriptionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDBClusterDescriptionRequest) SetResourceOwnerId(v int64) *ModifyDBClusterDescriptionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDBClusterDescriptionRequest) SetOwnerAccount(v string) *ModifyDBClusterDescriptionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyDBClusterDescriptionRequest) SetDBClusterId(v string) *ModifyDBClusterDescriptionRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyDBClusterDescriptionRequest) SetDBClusterDescription(v string) *ModifyDBClusterDescriptionRequest {
	s.DBClusterDescription = &v
	return s
}

type ModifyDBClusterDescriptionResponse struct {
	RequestId *string `json:"RequestId" xml:"RequestId" require:"true"`
}

func (s ModifyDBClusterDescriptionResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBClusterDescriptionResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterDescriptionResponse) SetRequestId(v string) *ModifyDBClusterDescriptionResponse {
	s.RequestId = &v
	return s
}

type ModifyAccountDescriptionRequest struct {
	AccessKeyId          *string `json:"AccessKeyId" xml:"AccessKeyId"`
	OwnerId              *int64  `json:"OwnerId" xml:"OwnerId"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount" xml:"ResourceOwnerAccount"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId" xml:"ResourceOwnerId"`
	OwnerAccount         *string `json:"OwnerAccount" xml:"OwnerAccount"`
	DBClusterId          *string `json:"DBClusterId" xml:"DBClusterId" require:"true"`
	AccountName          *string `json:"AccountName" xml:"AccountName" require:"true"`
	AccountDescription   *string `json:"AccountDescription" xml:"AccountDescription" require:"true"`
}

func (s ModifyAccountDescriptionRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyAccountDescriptionRequest) GoString() string {
	return s.String()
}

func (s *ModifyAccountDescriptionRequest) SetAccessKeyId(v string) *ModifyAccountDescriptionRequest {
	s.AccessKeyId = &v
	return s
}

func (s *ModifyAccountDescriptionRequest) SetOwnerId(v int64) *ModifyAccountDescriptionRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyAccountDescriptionRequest) SetResourceOwnerAccount(v string) *ModifyAccountDescriptionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyAccountDescriptionRequest) SetResourceOwnerId(v int64) *ModifyAccountDescriptionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyAccountDescriptionRequest) SetOwnerAccount(v string) *ModifyAccountDescriptionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyAccountDescriptionRequest) SetDBClusterId(v string) *ModifyAccountDescriptionRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyAccountDescriptionRequest) SetAccountName(v string) *ModifyAccountDescriptionRequest {
	s.AccountName = &v
	return s
}

func (s *ModifyAccountDescriptionRequest) SetAccountDescription(v string) *ModifyAccountDescriptionRequest {
	s.AccountDescription = &v
	return s
}

type ModifyAccountDescriptionResponse struct {
	RequestId *string `json:"RequestId" xml:"RequestId" require:"true"`
}

func (s ModifyAccountDescriptionResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyAccountDescriptionResponse) GoString() string {
	return s.String()
}

func (s *ModifyAccountDescriptionResponse) SetRequestId(v string) *ModifyAccountDescriptionResponse {
	s.RequestId = &v
	return s
}

type DescribeRegionsRequest struct {
	AccessKeyId          *string `json:"AccessKeyId" xml:"AccessKeyId"`
	OwnerId              *int64  `json:"OwnerId" xml:"OwnerId"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount" xml:"ResourceOwnerAccount"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId" xml:"ResourceOwnerId"`
	OwnerAccount         *string `json:"OwnerAccount" xml:"OwnerAccount"`
}

func (s DescribeRegionsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeRegionsRequest) SetAccessKeyId(v string) *DescribeRegionsRequest {
	s.AccessKeyId = &v
	return s
}

func (s *DescribeRegionsRequest) SetOwnerId(v int64) *DescribeRegionsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeRegionsRequest) SetResourceOwnerAccount(v string) *DescribeRegionsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeRegionsRequest) SetResourceOwnerId(v int64) *DescribeRegionsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeRegionsRequest) SetOwnerAccount(v string) *DescribeRegionsRequest {
	s.OwnerAccount = &v
	return s
}

type DescribeRegionsResponse struct {
	RequestId *string                         `json:"RequestId" xml:"RequestId" require:"true"`
	Regions   *DescribeRegionsResponseRegions `json:"Regions" xml:"Regions" require:"true" type:"Struct"`
}

func (s DescribeRegionsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponse) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponse) SetRequestId(v string) *DescribeRegionsResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeRegionsResponse) SetRegions(v *DescribeRegionsResponseRegions) *DescribeRegionsResponse {
	s.Regions = v
	return s
}

type DescribeRegionsResponseRegions struct {
	Region []*DescribeRegionsResponseRegionsRegion `json:"Region" xml:"Region" require:"true" type:"Repeated"`
}

func (s DescribeRegionsResponseRegions) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseRegions) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseRegions) SetRegion(v []*DescribeRegionsResponseRegionsRegion) *DescribeRegionsResponseRegions {
	s.Region = v
	return s
}

type DescribeRegionsResponseRegionsRegion struct {
	RegionId *string                                    `json:"RegionId" xml:"RegionId" require:"true"`
	Zones    *DescribeRegionsResponseRegionsRegionZones `json:"Zones" xml:"Zones" require:"true" type:"Struct"`
}

func (s DescribeRegionsResponseRegionsRegion) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseRegionsRegion) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseRegionsRegion) SetRegionId(v string) *DescribeRegionsResponseRegionsRegion {
	s.RegionId = &v
	return s
}

func (s *DescribeRegionsResponseRegionsRegion) SetZones(v *DescribeRegionsResponseRegionsRegionZones) *DescribeRegionsResponseRegionsRegion {
	s.Zones = v
	return s
}

type DescribeRegionsResponseRegionsRegionZones struct {
	Zone []*DescribeRegionsResponseRegionsRegionZonesZone `json:"Zone" xml:"Zone" require:"true" type:"Repeated"`
}

func (s DescribeRegionsResponseRegionsRegionZones) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseRegionsRegionZones) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseRegionsRegionZones) SetZone(v []*DescribeRegionsResponseRegionsRegionZonesZone) *DescribeRegionsResponseRegionsRegionZones {
	s.Zone = v
	return s
}

type DescribeRegionsResponseRegionsRegionZonesZone struct {
	ZoneId     *string `json:"ZoneId" xml:"ZoneId" require:"true"`
	VpcEnabled *bool   `json:"VpcEnabled" xml:"VpcEnabled" require:"true"`
}

func (s DescribeRegionsResponseRegionsRegionZonesZone) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseRegionsRegionZonesZone) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseRegionsRegionZonesZone) SetZoneId(v string) *DescribeRegionsResponseRegionsRegionZonesZone {
	s.ZoneId = &v
	return s
}

func (s *DescribeRegionsResponseRegionsRegionZonesZone) SetVpcEnabled(v bool) *DescribeRegionsResponseRegionsRegionZonesZone {
	s.VpcEnabled = &v
	return s
}

type DescribeAccountsRequest struct {
	AccessKeyId          *string `json:"AccessKeyId" xml:"AccessKeyId"`
	OwnerId              *int64  `json:"OwnerId" xml:"OwnerId"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount" xml:"ResourceOwnerAccount"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId" xml:"ResourceOwnerId"`
	OwnerAccount         *string `json:"OwnerAccount" xml:"OwnerAccount"`
	DBClusterId          *string `json:"DBClusterId" xml:"DBClusterId" require:"true"`
	AccountName          *string `json:"AccountName" xml:"AccountName"`
}

func (s DescribeAccountsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccountsRequest) GoString() string {
	return s.String()
}

func (s *DescribeAccountsRequest) SetAccessKeyId(v string) *DescribeAccountsRequest {
	s.AccessKeyId = &v
	return s
}

func (s *DescribeAccountsRequest) SetOwnerId(v int64) *DescribeAccountsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeAccountsRequest) SetResourceOwnerAccount(v string) *DescribeAccountsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeAccountsRequest) SetResourceOwnerId(v int64) *DescribeAccountsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeAccountsRequest) SetOwnerAccount(v string) *DescribeAccountsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeAccountsRequest) SetDBClusterId(v string) *DescribeAccountsRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeAccountsRequest) SetAccountName(v string) *DescribeAccountsRequest {
	s.AccountName = &v
	return s
}

type DescribeAccountsResponse struct {
	RequestId *string                             `json:"RequestId" xml:"RequestId" require:"true"`
	Accounts  []*DescribeAccountsResponseAccounts `json:"Accounts" xml:"Accounts" require:"true" type:"Repeated"`
}

func (s DescribeAccountsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccountsResponse) GoString() string {
	return s.String()
}

func (s *DescribeAccountsResponse) SetRequestId(v string) *DescribeAccountsResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeAccountsResponse) SetAccounts(v []*DescribeAccountsResponseAccounts) *DescribeAccountsResponse {
	s.Accounts = v
	return s
}

type DescribeAccountsResponseAccounts struct {
	AccountName              *string                                               `json:"AccountName" xml:"AccountName" require:"true"`
	AccountStatus            *string                                               `json:"AccountStatus" xml:"AccountStatus" require:"true"`
	AccountDescription       *string                                               `json:"AccountDescription" xml:"AccountDescription" require:"true"`
	AccountType              *string                                               `json:"AccountType" xml:"AccountType" require:"true"`
	AccountLockState         *string                                               `json:"AccountLockState" xml:"AccountLockState" require:"true"`
	PrivilegeExceeded        *string                                               `json:"PrivilegeExceeded" xml:"PrivilegeExceeded" require:"true"`
	AccountPasswordValidTime *string                                               `json:"AccountPasswordValidTime" xml:"AccountPasswordValidTime" require:"true"`
	DatabasePrivileges       []*DescribeAccountsResponseAccountsDatabasePrivileges `json:"DatabasePrivileges" xml:"DatabasePrivileges" require:"true" type:"Repeated"`
}

func (s DescribeAccountsResponseAccounts) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccountsResponseAccounts) GoString() string {
	return s.String()
}

func (s *DescribeAccountsResponseAccounts) SetAccountName(v string) *DescribeAccountsResponseAccounts {
	s.AccountName = &v
	return s
}

func (s *DescribeAccountsResponseAccounts) SetAccountStatus(v string) *DescribeAccountsResponseAccounts {
	s.AccountStatus = &v
	return s
}

func (s *DescribeAccountsResponseAccounts) SetAccountDescription(v string) *DescribeAccountsResponseAccounts {
	s.AccountDescription = &v
	return s
}

func (s *DescribeAccountsResponseAccounts) SetAccountType(v string) *DescribeAccountsResponseAccounts {
	s.AccountType = &v
	return s
}

func (s *DescribeAccountsResponseAccounts) SetAccountLockState(v string) *DescribeAccountsResponseAccounts {
	s.AccountLockState = &v
	return s
}

func (s *DescribeAccountsResponseAccounts) SetPrivilegeExceeded(v string) *DescribeAccountsResponseAccounts {
	s.PrivilegeExceeded = &v
	return s
}

func (s *DescribeAccountsResponseAccounts) SetAccountPasswordValidTime(v string) *DescribeAccountsResponseAccounts {
	s.AccountPasswordValidTime = &v
	return s
}

func (s *DescribeAccountsResponseAccounts) SetDatabasePrivileges(v []*DescribeAccountsResponseAccountsDatabasePrivileges) *DescribeAccountsResponseAccounts {
	s.DatabasePrivileges = v
	return s
}

type DescribeAccountsResponseAccountsDatabasePrivileges struct {
	DBName           *string `json:"DBName" xml:"DBName" require:"true"`
	AccountPrivilege *string `json:"AccountPrivilege" xml:"AccountPrivilege" require:"true"`
}

func (s DescribeAccountsResponseAccountsDatabasePrivileges) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccountsResponseAccountsDatabasePrivileges) GoString() string {
	return s.String()
}

func (s *DescribeAccountsResponseAccountsDatabasePrivileges) SetDBName(v string) *DescribeAccountsResponseAccountsDatabasePrivileges {
	s.DBName = &v
	return s
}

func (s *DescribeAccountsResponseAccountsDatabasePrivileges) SetAccountPrivilege(v string) *DescribeAccountsResponseAccountsDatabasePrivileges {
	s.AccountPrivilege = &v
	return s
}

type DeleteDBClusterRequest struct {
	AccessKeyId                            *string `json:"AccessKeyId" xml:"AccessKeyId"`
	OwnerId                                *int64  `json:"OwnerId" xml:"OwnerId"`
	ResourceOwnerAccount                   *string `json:"ResourceOwnerAccount" xml:"ResourceOwnerAccount"`
	ResourceOwnerId                        *int64  `json:"ResourceOwnerId" xml:"ResourceOwnerId"`
	OwnerAccount                           *string `json:"OwnerAccount" xml:"OwnerAccount"`
	DBClusterId                            *string `json:"DBClusterId" xml:"DBClusterId" require:"true"`
	BackupRetentionPolicyOnClusterDeletion *string `json:"BackupRetentionPolicyOnClusterDeletion" xml:"BackupRetentionPolicyOnClusterDeletion"`
}

func (s DeleteDBClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDBClusterRequest) GoString() string {
	return s.String()
}

func (s *DeleteDBClusterRequest) SetAccessKeyId(v string) *DeleteDBClusterRequest {
	s.AccessKeyId = &v
	return s
}

func (s *DeleteDBClusterRequest) SetOwnerId(v int64) *DeleteDBClusterRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteDBClusterRequest) SetResourceOwnerAccount(v string) *DeleteDBClusterRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteDBClusterRequest) SetResourceOwnerId(v int64) *DeleteDBClusterRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteDBClusterRequest) SetOwnerAccount(v string) *DeleteDBClusterRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteDBClusterRequest) SetDBClusterId(v string) *DeleteDBClusterRequest {
	s.DBClusterId = &v
	return s
}

func (s *DeleteDBClusterRequest) SetBackupRetentionPolicyOnClusterDeletion(v string) *DeleteDBClusterRequest {
	s.BackupRetentionPolicyOnClusterDeletion = &v
	return s
}

type DeleteDBClusterResponse struct {
	RequestId *string `json:"RequestId" xml:"RequestId" require:"true"`
}

func (s DeleteDBClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDBClusterResponse) GoString() string {
	return s.String()
}

func (s *DeleteDBClusterResponse) SetRequestId(v string) *DeleteDBClusterResponse {
	s.RequestId = &v
	return s
}

type CreateDBClusterRequest struct {
	AccessKeyId          *string `json:"AccessKeyId" xml:"AccessKeyId"`
	OwnerId              *int64  `json:"OwnerId" xml:"OwnerId"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount" xml:"ResourceOwnerAccount"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId" xml:"ResourceOwnerId"`
	OwnerAccount         *string `json:"OwnerAccount" xml:"OwnerAccount"`
	ZoneId               *string `json:"ZoneId" xml:"ZoneId"`
	DBType               *string `json:"DBType" xml:"DBType" require:"true"`
	DBVersion            *string `json:"DBVersion" xml:"DBVersion" require:"true"`
	DBNodeClass          *string `json:"DBNodeClass" xml:"DBNodeClass" require:"true"`
	ClusterNetworkType   *string `json:"ClusterNetworkType" xml:"ClusterNetworkType"`
	DBClusterDescription *string `json:"DBClusterDescription" xml:"DBClusterDescription"`
	PayType              *string `json:"PayType" xml:"PayType" require:"true"`
	AutoRenew            *bool   `json:"AutoRenew" xml:"AutoRenew"`
	Period               *string `json:"Period" xml:"Period"`
	UsedTime             *string `json:"UsedTime" xml:"UsedTime"`
	VPCId                *string `json:"VPCId" xml:"VPCId"`
	VSwitchId            *string `json:"VSwitchId" xml:"VSwitchId"`
	CreationOption       *string `json:"CreationOption" xml:"CreationOption"`
	SourceResourceId     *string `json:"SourceResourceId" xml:"SourceResourceId"`
	CloneDataPoint       *string `json:"CloneDataPoint" xml:"CloneDataPoint"`
	ClientToken          *string `json:"ClientToken" xml:"ClientToken"`
	ResourceGroupId      *string `json:"ResourceGroupId" xml:"ResourceGroupId"`
	SecurityIPList       *string `json:"SecurityIPList" xml:"SecurityIPList"`
	TDEStatus            *bool   `json:"TDEStatus" xml:"TDEStatus"`
	GDNId                *string `json:"GDNId" xml:"GDNId"`
	CreationCategory     *string `json:"CreationCategory" xml:"CreationCategory"`
}

func (s CreateDBClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDBClusterRequest) GoString() string {
	return s.String()
}

func (s *CreateDBClusterRequest) SetAccessKeyId(v string) *CreateDBClusterRequest {
	s.AccessKeyId = &v
	return s
}

func (s *CreateDBClusterRequest) SetOwnerId(v int64) *CreateDBClusterRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateDBClusterRequest) SetResourceOwnerAccount(v string) *CreateDBClusterRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateDBClusterRequest) SetResourceOwnerId(v int64) *CreateDBClusterRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateDBClusterRequest) SetOwnerAccount(v string) *CreateDBClusterRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateDBClusterRequest) SetZoneId(v string) *CreateDBClusterRequest {
	s.ZoneId = &v
	return s
}

func (s *CreateDBClusterRequest) SetDBType(v string) *CreateDBClusterRequest {
	s.DBType = &v
	return s
}

func (s *CreateDBClusterRequest) SetDBVersion(v string) *CreateDBClusterRequest {
	s.DBVersion = &v
	return s
}

func (s *CreateDBClusterRequest) SetDBNodeClass(v string) *CreateDBClusterRequest {
	s.DBNodeClass = &v
	return s
}

func (s *CreateDBClusterRequest) SetClusterNetworkType(v string) *CreateDBClusterRequest {
	s.ClusterNetworkType = &v
	return s
}

func (s *CreateDBClusterRequest) SetDBClusterDescription(v string) *CreateDBClusterRequest {
	s.DBClusterDescription = &v
	return s
}

func (s *CreateDBClusterRequest) SetPayType(v string) *CreateDBClusterRequest {
	s.PayType = &v
	return s
}

func (s *CreateDBClusterRequest) SetAutoRenew(v bool) *CreateDBClusterRequest {
	s.AutoRenew = &v
	return s
}

func (s *CreateDBClusterRequest) SetPeriod(v string) *CreateDBClusterRequest {
	s.Period = &v
	return s
}

func (s *CreateDBClusterRequest) SetUsedTime(v string) *CreateDBClusterRequest {
	s.UsedTime = &v
	return s
}

func (s *CreateDBClusterRequest) SetVPCId(v string) *CreateDBClusterRequest {
	s.VPCId = &v
	return s
}

func (s *CreateDBClusterRequest) SetVSwitchId(v string) *CreateDBClusterRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateDBClusterRequest) SetCreationOption(v string) *CreateDBClusterRequest {
	s.CreationOption = &v
	return s
}

func (s *CreateDBClusterRequest) SetSourceResourceId(v string) *CreateDBClusterRequest {
	s.SourceResourceId = &v
	return s
}

func (s *CreateDBClusterRequest) SetCloneDataPoint(v string) *CreateDBClusterRequest {
	s.CloneDataPoint = &v
	return s
}

func (s *CreateDBClusterRequest) SetClientToken(v string) *CreateDBClusterRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateDBClusterRequest) SetResourceGroupId(v string) *CreateDBClusterRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateDBClusterRequest) SetSecurityIPList(v string) *CreateDBClusterRequest {
	s.SecurityIPList = &v
	return s
}

func (s *CreateDBClusterRequest) SetTDEStatus(v bool) *CreateDBClusterRequest {
	s.TDEStatus = &v
	return s
}

func (s *CreateDBClusterRequest) SetGDNId(v string) *CreateDBClusterRequest {
	s.GDNId = &v
	return s
}

func (s *CreateDBClusterRequest) SetCreationCategory(v string) *CreateDBClusterRequest {
	s.CreationCategory = &v
	return s
}

type CreateDBClusterResponse struct {
	RequestId       *string `json:"RequestId" xml:"RequestId" require:"true"`
	DBClusterId     *string `json:"DBClusterId" xml:"DBClusterId" require:"true"`
	OrderId         *string `json:"OrderId" xml:"OrderId" require:"true"`
	ResourceGroupId *string `json:"ResourceGroupId" xml:"ResourceGroupId" require:"true"`
}

func (s CreateDBClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDBClusterResponse) GoString() string {
	return s.String()
}

func (s *CreateDBClusterResponse) SetRequestId(v string) *CreateDBClusterResponse {
	s.RequestId = &v
	return s
}

func (s *CreateDBClusterResponse) SetDBClusterId(v string) *CreateDBClusterResponse {
	s.DBClusterId = &v
	return s
}

func (s *CreateDBClusterResponse) SetOrderId(v string) *CreateDBClusterResponse {
	s.OrderId = &v
	return s
}

func (s *CreateDBClusterResponse) SetResourceGroupId(v string) *CreateDBClusterResponse {
	s.ResourceGroupId = &v
	return s
}

type CreateBackupRequest struct {
	AccessKeyId          *string `json:"AccessKeyId" xml:"AccessKeyId"`
	OwnerId              *int64  `json:"OwnerId" xml:"OwnerId"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount" xml:"ResourceOwnerAccount"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId" xml:"ResourceOwnerId"`
	OwnerAccount         *string `json:"OwnerAccount" xml:"OwnerAccount"`
	DBClusterId          *string `json:"DBClusterId" xml:"DBClusterId" require:"true"`
}

func (s CreateBackupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateBackupRequest) GoString() string {
	return s.String()
}

func (s *CreateBackupRequest) SetAccessKeyId(v string) *CreateBackupRequest {
	s.AccessKeyId = &v
	return s
}

func (s *CreateBackupRequest) SetOwnerId(v int64) *CreateBackupRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateBackupRequest) SetResourceOwnerAccount(v string) *CreateBackupRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateBackupRequest) SetResourceOwnerId(v int64) *CreateBackupRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateBackupRequest) SetOwnerAccount(v string) *CreateBackupRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateBackupRequest) SetDBClusterId(v string) *CreateBackupRequest {
	s.DBClusterId = &v
	return s
}

type CreateBackupResponse struct {
	RequestId   *string `json:"RequestId" xml:"RequestId" require:"true"`
	BackupJobId *string `json:"BackupJobId" xml:"BackupJobId" require:"true"`
}

func (s CreateBackupResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateBackupResponse) GoString() string {
	return s.String()
}

func (s *CreateBackupResponse) SetRequestId(v string) *CreateBackupResponse {
	s.RequestId = &v
	return s
}

func (s *CreateBackupResponse) SetBackupJobId(v string) *CreateBackupResponse {
	s.BackupJobId = &v
	return s
}

type CreateAccountRequest struct {
	AccessKeyId          *string `json:"AccessKeyId" xml:"AccessKeyId"`
	OwnerId              *int64  `json:"OwnerId" xml:"OwnerId"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount" xml:"ResourceOwnerAccount"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId" xml:"ResourceOwnerId"`
	OwnerAccount         *string `json:"OwnerAccount" xml:"OwnerAccount"`
	DBClusterId          *string `json:"DBClusterId" xml:"DBClusterId" require:"true"`
	AccountName          *string `json:"AccountName" xml:"AccountName" require:"true"`
	AccountPassword      *string `json:"AccountPassword" xml:"AccountPassword" require:"true"`
	AccountType          *string `json:"AccountType" xml:"AccountType"`
	AccountDescription   *string `json:"AccountDescription" xml:"AccountDescription"`
	DBName               *string `json:"DBName" xml:"DBName"`
	AccountPrivilege     *string `json:"AccountPrivilege" xml:"AccountPrivilege"`
}

func (s CreateAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAccountRequest) GoString() string {
	return s.String()
}

func (s *CreateAccountRequest) SetAccessKeyId(v string) *CreateAccountRequest {
	s.AccessKeyId = &v
	return s
}

func (s *CreateAccountRequest) SetOwnerId(v int64) *CreateAccountRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateAccountRequest) SetResourceOwnerAccount(v string) *CreateAccountRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateAccountRequest) SetResourceOwnerId(v int64) *CreateAccountRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateAccountRequest) SetOwnerAccount(v string) *CreateAccountRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateAccountRequest) SetDBClusterId(v string) *CreateAccountRequest {
	s.DBClusterId = &v
	return s
}

func (s *CreateAccountRequest) SetAccountName(v string) *CreateAccountRequest {
	s.AccountName = &v
	return s
}

func (s *CreateAccountRequest) SetAccountPassword(v string) *CreateAccountRequest {
	s.AccountPassword = &v
	return s
}

func (s *CreateAccountRequest) SetAccountType(v string) *CreateAccountRequest {
	s.AccountType = &v
	return s
}

func (s *CreateAccountRequest) SetAccountDescription(v string) *CreateAccountRequest {
	s.AccountDescription = &v
	return s
}

func (s *CreateAccountRequest) SetDBName(v string) *CreateAccountRequest {
	s.DBName = &v
	return s
}

func (s *CreateAccountRequest) SetAccountPrivilege(v string) *CreateAccountRequest {
	s.AccountPrivilege = &v
	return s
}

type CreateAccountResponse struct {
	RequestId *string `json:"RequestId" xml:"RequestId" require:"true"`
}

func (s CreateAccountResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAccountResponse) GoString() string {
	return s.String()
}

func (s *CreateAccountResponse) SetRequestId(v string) *CreateAccountResponse {
	s.RequestId = &v
	return s
}

type DescribeBackupsRequest struct {
	AccessKeyId          *string `json:"AccessKeyId" xml:"AccessKeyId"`
	OwnerId              *int64  `json:"OwnerId" xml:"OwnerId"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount" xml:"ResourceOwnerAccount"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId" xml:"ResourceOwnerId"`
	OwnerAccount         *string `json:"OwnerAccount" xml:"OwnerAccount"`
	DBClusterId          *string `json:"DBClusterId" xml:"DBClusterId" require:"true"`
	BackupId             *string `json:"BackupId" xml:"BackupId"`
	BackupStatus         *string `json:"BackupStatus" xml:"BackupStatus"`
	BackupMode           *string `json:"BackupMode" xml:"BackupMode"`
	StartTime            *string `json:"StartTime" xml:"StartTime" require:"true"`
	EndTime              *string `json:"EndTime" xml:"EndTime" require:"true"`
	PageSize             *int    `json:"PageSize" xml:"PageSize"`
	PageNumber           *int    `json:"PageNumber" xml:"PageNumber"`
}

func (s DescribeBackupsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupsRequest) GoString() string {
	return s.String()
}

func (s *DescribeBackupsRequest) SetAccessKeyId(v string) *DescribeBackupsRequest {
	s.AccessKeyId = &v
	return s
}

func (s *DescribeBackupsRequest) SetOwnerId(v int64) *DescribeBackupsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeBackupsRequest) SetResourceOwnerAccount(v string) *DescribeBackupsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeBackupsRequest) SetResourceOwnerId(v int64) *DescribeBackupsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeBackupsRequest) SetOwnerAccount(v string) *DescribeBackupsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeBackupsRequest) SetDBClusterId(v string) *DescribeBackupsRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeBackupsRequest) SetBackupId(v string) *DescribeBackupsRequest {
	s.BackupId = &v
	return s
}

func (s *DescribeBackupsRequest) SetBackupStatus(v string) *DescribeBackupsRequest {
	s.BackupStatus = &v
	return s
}

func (s *DescribeBackupsRequest) SetBackupMode(v string) *DescribeBackupsRequest {
	s.BackupMode = &v
	return s
}

func (s *DescribeBackupsRequest) SetStartTime(v string) *DescribeBackupsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeBackupsRequest) SetEndTime(v string) *DescribeBackupsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeBackupsRequest) SetPageSize(v int) *DescribeBackupsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeBackupsRequest) SetPageNumber(v int) *DescribeBackupsRequest {
	s.PageNumber = &v
	return s
}

type DescribeBackupsResponse struct {
	RequestId        *string                       `json:"RequestId" xml:"RequestId" require:"true"`
	TotalRecordCount *string                       `json:"TotalRecordCount" xml:"TotalRecordCount" require:"true"`
	PageNumber       *string                       `json:"PageNumber" xml:"PageNumber" require:"true"`
	PageRecordCount  *string                       `json:"PageRecordCount" xml:"PageRecordCount" require:"true"`
	Items            *DescribeBackupsResponseItems `json:"Items" xml:"Items" require:"true" type:"Struct"`
}

func (s DescribeBackupsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupsResponse) GoString() string {
	return s.String()
}

func (s *DescribeBackupsResponse) SetRequestId(v string) *DescribeBackupsResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeBackupsResponse) SetTotalRecordCount(v string) *DescribeBackupsResponse {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeBackupsResponse) SetPageNumber(v string) *DescribeBackupsResponse {
	s.PageNumber = &v
	return s
}

func (s *DescribeBackupsResponse) SetPageRecordCount(v string) *DescribeBackupsResponse {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeBackupsResponse) SetItems(v *DescribeBackupsResponseItems) *DescribeBackupsResponse {
	s.Items = v
	return s
}

type DescribeBackupsResponseItems struct {
	Backup []*DescribeBackupsResponseItemsBackup `json:"Backup" xml:"Backup" require:"true" type:"Repeated"`
}

func (s DescribeBackupsResponseItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupsResponseItems) GoString() string {
	return s.String()
}

func (s *DescribeBackupsResponseItems) SetBackup(v []*DescribeBackupsResponseItemsBackup) *DescribeBackupsResponseItems {
	s.Backup = v
	return s
}

type DescribeBackupsResponseItemsBackup struct {
	BackupId        *string `json:"BackupId" xml:"BackupId" require:"true"`
	DBClusterId     *string `json:"DBClusterId" xml:"DBClusterId" require:"true"`
	BackupStatus    *string `json:"BackupStatus" xml:"BackupStatus" require:"true"`
	BackupStartTime *string `json:"BackupStartTime" xml:"BackupStartTime" require:"true"`
	BackupEndTime   *string `json:"BackupEndTime" xml:"BackupEndTime" require:"true"`
	BackupType      *string `json:"BackupType" xml:"BackupType" require:"true"`
	BackupMode      *string `json:"BackupMode" xml:"BackupMode" require:"true"`
	BackupMethod    *string `json:"BackupMethod" xml:"BackupMethod" require:"true"`
	StoreStatus     *string `json:"StoreStatus" xml:"StoreStatus" require:"true"`
	BackupSetSize   *string `json:"BackupSetSize" xml:"BackupSetSize" require:"true"`
	ConsistentTime  *string `json:"ConsistentTime" xml:"ConsistentTime" require:"true"`
	BackupsLevel    *string `json:"BackupsLevel" xml:"BackupsLevel" require:"true"`
	IsAvail         *string `json:"IsAvail" xml:"IsAvail" require:"true"`
}

func (s DescribeBackupsResponseItemsBackup) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupsResponseItemsBackup) GoString() string {
	return s.String()
}

func (s *DescribeBackupsResponseItemsBackup) SetBackupId(v string) *DescribeBackupsResponseItemsBackup {
	s.BackupId = &v
	return s
}

func (s *DescribeBackupsResponseItemsBackup) SetDBClusterId(v string) *DescribeBackupsResponseItemsBackup {
	s.DBClusterId = &v
	return s
}

func (s *DescribeBackupsResponseItemsBackup) SetBackupStatus(v string) *DescribeBackupsResponseItemsBackup {
	s.BackupStatus = &v
	return s
}

func (s *DescribeBackupsResponseItemsBackup) SetBackupStartTime(v string) *DescribeBackupsResponseItemsBackup {
	s.BackupStartTime = &v
	return s
}

func (s *DescribeBackupsResponseItemsBackup) SetBackupEndTime(v string) *DescribeBackupsResponseItemsBackup {
	s.BackupEndTime = &v
	return s
}

func (s *DescribeBackupsResponseItemsBackup) SetBackupType(v string) *DescribeBackupsResponseItemsBackup {
	s.BackupType = &v
	return s
}

func (s *DescribeBackupsResponseItemsBackup) SetBackupMode(v string) *DescribeBackupsResponseItemsBackup {
	s.BackupMode = &v
	return s
}

func (s *DescribeBackupsResponseItemsBackup) SetBackupMethod(v string) *DescribeBackupsResponseItemsBackup {
	s.BackupMethod = &v
	return s
}

func (s *DescribeBackupsResponseItemsBackup) SetStoreStatus(v string) *DescribeBackupsResponseItemsBackup {
	s.StoreStatus = &v
	return s
}

func (s *DescribeBackupsResponseItemsBackup) SetBackupSetSize(v string) *DescribeBackupsResponseItemsBackup {
	s.BackupSetSize = &v
	return s
}

func (s *DescribeBackupsResponseItemsBackup) SetConsistentTime(v string) *DescribeBackupsResponseItemsBackup {
	s.ConsistentTime = &v
	return s
}

func (s *DescribeBackupsResponseItemsBackup) SetBackupsLevel(v string) *DescribeBackupsResponseItemsBackup {
	s.BackupsLevel = &v
	return s
}

func (s *DescribeBackupsResponseItemsBackup) SetIsAvail(v string) *DescribeBackupsResponseItemsBackup {
	s.IsAvail = &v
	return s
}

type ModifyBackupPolicyRequest struct {
	AccessKeyId                            *string `json:"AccessKeyId" xml:"AccessKeyId"`
	OwnerId                                *int64  `json:"OwnerId" xml:"OwnerId"`
	ResourceOwnerAccount                   *string `json:"ResourceOwnerAccount" xml:"ResourceOwnerAccount"`
	ResourceOwnerId                        *int64  `json:"ResourceOwnerId" xml:"ResourceOwnerId"`
	OwnerAccount                           *string `json:"OwnerAccount" xml:"OwnerAccount"`
	DBClusterId                            *string `json:"DBClusterId" xml:"DBClusterId" require:"true"`
	PreferredBackupTime                    *string `json:"PreferredBackupTime" xml:"PreferredBackupTime"`
	PreferredBackupPeriod                  *string `json:"PreferredBackupPeriod" xml:"PreferredBackupPeriod"`
	DataLevel1BackupRetentionPeriod        *string `json:"DataLevel1BackupRetentionPeriod" xml:"DataLevel1BackupRetentionPeriod"`
	DataLevel2BackupRetentionPeriod        *string `json:"DataLevel2BackupRetentionPeriod" xml:"DataLevel2BackupRetentionPeriod"`
	BackupRetentionPolicyOnClusterDeletion *string `json:"BackupRetentionPolicyOnClusterDeletion" xml:"BackupRetentionPolicyOnClusterDeletion"`
}

func (s ModifyBackupPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyBackupPolicyRequest) GoString() string {
	return s.String()
}

func (s *ModifyBackupPolicyRequest) SetAccessKeyId(v string) *ModifyBackupPolicyRequest {
	s.AccessKeyId = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetOwnerId(v int64) *ModifyBackupPolicyRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetResourceOwnerAccount(v string) *ModifyBackupPolicyRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetResourceOwnerId(v int64) *ModifyBackupPolicyRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetOwnerAccount(v string) *ModifyBackupPolicyRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetDBClusterId(v string) *ModifyBackupPolicyRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetPreferredBackupTime(v string) *ModifyBackupPolicyRequest {
	s.PreferredBackupTime = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetPreferredBackupPeriod(v string) *ModifyBackupPolicyRequest {
	s.PreferredBackupPeriod = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetDataLevel1BackupRetentionPeriod(v string) *ModifyBackupPolicyRequest {
	s.DataLevel1BackupRetentionPeriod = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetDataLevel2BackupRetentionPeriod(v string) *ModifyBackupPolicyRequest {
	s.DataLevel2BackupRetentionPeriod = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetBackupRetentionPolicyOnClusterDeletion(v string) *ModifyBackupPolicyRequest {
	s.BackupRetentionPolicyOnClusterDeletion = &v
	return s
}

type ModifyBackupPolicyResponse struct {
	RequestId *string `json:"RequestId" xml:"RequestId" require:"true"`
}

func (s ModifyBackupPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyBackupPolicyResponse) GoString() string {
	return s.String()
}

func (s *ModifyBackupPolicyResponse) SetRequestId(v string) *ModifyBackupPolicyResponse {
	s.RequestId = &v
	return s
}

type DescribeBackupPolicyRequest struct {
	AccessKeyId          *string `json:"AccessKeyId" xml:"AccessKeyId"`
	OwnerId              *int64  `json:"OwnerId" xml:"OwnerId"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount" xml:"ResourceOwnerAccount"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId" xml:"ResourceOwnerId"`
	OwnerAccount         *string `json:"OwnerAccount" xml:"OwnerAccount"`
	DBClusterId          *string `json:"DBClusterId" xml:"DBClusterId" require:"true"`
}

func (s DescribeBackupPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupPolicyRequest) GoString() string {
	return s.String()
}

func (s *DescribeBackupPolicyRequest) SetAccessKeyId(v string) *DescribeBackupPolicyRequest {
	s.AccessKeyId = &v
	return s
}

func (s *DescribeBackupPolicyRequest) SetOwnerId(v int64) *DescribeBackupPolicyRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeBackupPolicyRequest) SetResourceOwnerAccount(v string) *DescribeBackupPolicyRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeBackupPolicyRequest) SetResourceOwnerId(v int64) *DescribeBackupPolicyRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeBackupPolicyRequest) SetOwnerAccount(v string) *DescribeBackupPolicyRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeBackupPolicyRequest) SetDBClusterId(v string) *DescribeBackupPolicyRequest {
	s.DBClusterId = &v
	return s
}

type DescribeBackupPolicyResponse struct {
	RequestId                              *string `json:"RequestId" xml:"RequestId" require:"true"`
	BackupRetentionPeriod                  *int    `json:"BackupRetentionPeriod" xml:"BackupRetentionPeriod" require:"true"`
	PreferredNextBackupTime                *string `json:"PreferredNextBackupTime" xml:"PreferredNextBackupTime" require:"true"`
	PreferredBackupTime                    *string `json:"PreferredBackupTime" xml:"PreferredBackupTime" require:"true"`
	PreferredBackupPeriod                  *string `json:"PreferredBackupPeriod" xml:"PreferredBackupPeriod" require:"true"`
	DataLevel1BackupRetentionPeriod        *string `json:"DataLevel1BackupRetentionPeriod" xml:"DataLevel1BackupRetentionPeriod" require:"true"`
	DataLevel2BackupRetentionPeriod        *string `json:"DataLevel2BackupRetentionPeriod" xml:"DataLevel2BackupRetentionPeriod" require:"true"`
	BackupRetentionPolicyOnClusterDeletion *string `json:"BackupRetentionPolicyOnClusterDeletion" xml:"BackupRetentionPolicyOnClusterDeletion" require:"true"`
}

func (s DescribeBackupPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupPolicyResponse) GoString() string {
	return s.String()
}

func (s *DescribeBackupPolicyResponse) SetRequestId(v string) *DescribeBackupPolicyResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeBackupPolicyResponse) SetBackupRetentionPeriod(v int) *DescribeBackupPolicyResponse {
	s.BackupRetentionPeriod = &v
	return s
}

func (s *DescribeBackupPolicyResponse) SetPreferredNextBackupTime(v string) *DescribeBackupPolicyResponse {
	s.PreferredNextBackupTime = &v
	return s
}

func (s *DescribeBackupPolicyResponse) SetPreferredBackupTime(v string) *DescribeBackupPolicyResponse {
	s.PreferredBackupTime = &v
	return s
}

func (s *DescribeBackupPolicyResponse) SetPreferredBackupPeriod(v string) *DescribeBackupPolicyResponse {
	s.PreferredBackupPeriod = &v
	return s
}

func (s *DescribeBackupPolicyResponse) SetDataLevel1BackupRetentionPeriod(v string) *DescribeBackupPolicyResponse {
	s.DataLevel1BackupRetentionPeriod = &v
	return s
}

func (s *DescribeBackupPolicyResponse) SetDataLevel2BackupRetentionPeriod(v string) *DescribeBackupPolicyResponse {
	s.DataLevel2BackupRetentionPeriod = &v
	return s
}

func (s *DescribeBackupPolicyResponse) SetBackupRetentionPolicyOnClusterDeletion(v string) *DescribeBackupPolicyResponse {
	s.BackupRetentionPolicyOnClusterDeletion = &v
	return s
}

type DescribeDBClustersRequest struct {
	AccessKeyId          *string                         `json:"AccessKeyId" xml:"AccessKeyId"`
	OwnerId              *int64                          `json:"OwnerId" xml:"OwnerId"`
	ResourceOwnerAccount *string                         `json:"ResourceOwnerAccount" xml:"ResourceOwnerAccount"`
	ResourceOwnerId      *int64                          `json:"ResourceOwnerId" xml:"ResourceOwnerId"`
	OwnerAccount         *string                         `json:"OwnerAccount" xml:"OwnerAccount"`
	RegionId             *string                         `json:"RegionId" xml:"RegionId" require:"true"`
	DBClusterIds         *string                         `json:"DBClusterIds" xml:"DBClusterIds"`
	DBClusterDescription *string                         `json:"DBClusterDescription" xml:"DBClusterDescription"`
	DBClusterStatus      *string                         `json:"DBClusterStatus" xml:"DBClusterStatus"`
	DBType               *string                         `json:"DBType" xml:"DBType"`
	Tag                  []*DescribeDBClustersRequestTag `json:"Tag" xml:"Tag" type:"Repeated"`
	PageSize             *int                            `json:"PageSize" xml:"PageSize"`
	PageNumber           *int                            `json:"PageNumber" xml:"PageNumber"`
	ResourceGroupId      *string                         `json:"ResourceGroupId" xml:"ResourceGroupId"`
}

func (s DescribeDBClustersRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClustersRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBClustersRequest) SetAccessKeyId(v string) *DescribeDBClustersRequest {
	s.AccessKeyId = &v
	return s
}

func (s *DescribeDBClustersRequest) SetOwnerId(v int64) *DescribeDBClustersRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDBClustersRequest) SetResourceOwnerAccount(v string) *DescribeDBClustersRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDBClustersRequest) SetResourceOwnerId(v int64) *DescribeDBClustersRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDBClustersRequest) SetOwnerAccount(v string) *DescribeDBClustersRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDBClustersRequest) SetRegionId(v string) *DescribeDBClustersRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDBClustersRequest) SetDBClusterIds(v string) *DescribeDBClustersRequest {
	s.DBClusterIds = &v
	return s
}

func (s *DescribeDBClustersRequest) SetDBClusterDescription(v string) *DescribeDBClustersRequest {
	s.DBClusterDescription = &v
	return s
}

func (s *DescribeDBClustersRequest) SetDBClusterStatus(v string) *DescribeDBClustersRequest {
	s.DBClusterStatus = &v
	return s
}

func (s *DescribeDBClustersRequest) SetDBType(v string) *DescribeDBClustersRequest {
	s.DBType = &v
	return s
}

func (s *DescribeDBClustersRequest) SetTag(v []*DescribeDBClustersRequestTag) *DescribeDBClustersRequest {
	s.Tag = v
	return s
}

func (s *DescribeDBClustersRequest) SetPageSize(v int) *DescribeDBClustersRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDBClustersRequest) SetPageNumber(v int) *DescribeDBClustersRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDBClustersRequest) SetResourceGroupId(v string) *DescribeDBClustersRequest {
	s.ResourceGroupId = &v
	return s
}

type DescribeDBClustersRequestTag struct {
	Key   *string `json:"Key" xml:"Key" require:"true"`
	Value *string `json:"Value" xml:"Value" require:"true"`
}

func (s DescribeDBClustersRequestTag) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClustersRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeDBClustersRequestTag) SetKey(v string) *DescribeDBClustersRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeDBClustersRequestTag) SetValue(v string) *DescribeDBClustersRequestTag {
	s.Value = &v
	return s
}

type DescribeDBClustersResponse struct {
	RequestId        *string                          `json:"RequestId" xml:"RequestId" require:"true"`
	PageNumber       *int                             `json:"PageNumber" xml:"PageNumber" require:"true"`
	TotalRecordCount *int                             `json:"TotalRecordCount" xml:"TotalRecordCount" require:"true"`
	PageRecordCount  *int                             `json:"PageRecordCount" xml:"PageRecordCount" require:"true"`
	Items            *DescribeDBClustersResponseItems `json:"Items" xml:"Items" require:"true" type:"Struct"`
}

func (s DescribeDBClustersResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClustersResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBClustersResponse) SetRequestId(v string) *DescribeDBClustersResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeDBClustersResponse) SetPageNumber(v int) *DescribeDBClustersResponse {
	s.PageNumber = &v
	return s
}

func (s *DescribeDBClustersResponse) SetTotalRecordCount(v int) *DescribeDBClustersResponse {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeDBClustersResponse) SetPageRecordCount(v int) *DescribeDBClustersResponse {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeDBClustersResponse) SetItems(v *DescribeDBClustersResponseItems) *DescribeDBClustersResponse {
	s.Items = v
	return s
}

type DescribeDBClustersResponseItems struct {
	DBCluster []*DescribeDBClustersResponseItemsDBCluster `json:"DBCluster" xml:"DBCluster" require:"true" type:"Repeated"`
}

func (s DescribeDBClustersResponseItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClustersResponseItems) GoString() string {
	return s.String()
}

func (s *DescribeDBClustersResponseItems) SetDBCluster(v []*DescribeDBClustersResponseItemsDBCluster) *DescribeDBClustersResponseItems {
	s.DBCluster = v
	return s
}

type DescribeDBClustersResponseItemsDBCluster struct {
	DBClusterId          *string                                          `json:"DBClusterId" xml:"DBClusterId" require:"true"`
	DBClusterDescription *string                                          `json:"DBClusterDescription" xml:"DBClusterDescription" require:"true"`
	PayType              *string                                          `json:"PayType" xml:"PayType" require:"true"`
	DBClusterNetworkType *string                                          `json:"DBClusterNetworkType" xml:"DBClusterNetworkType" require:"true"`
	RegionId             *string                                          `json:"RegionId" xml:"RegionId" require:"true"`
	ZoneId               *string                                          `json:"ZoneId" xml:"ZoneId" require:"true"`
	ExpireTime           *string                                          `json:"ExpireTime" xml:"ExpireTime" require:"true"`
	Expired              *string                                          `json:"Expired" xml:"Expired" require:"true"`
	DBClusterStatus      *string                                          `json:"DBClusterStatus" xml:"DBClusterStatus" require:"true"`
	Engine               *string                                          `json:"Engine" xml:"Engine" require:"true"`
	DBType               *string                                          `json:"DBType" xml:"DBType" require:"true"`
	DBVersion            *string                                          `json:"DBVersion" xml:"DBVersion" require:"true"`
	LockMode             *string                                          `json:"LockMode" xml:"LockMode" require:"true"`
	DeletionLock         *int                                             `json:"DeletionLock" xml:"DeletionLock" require:"true"`
	CreateTime           *string                                          `json:"CreateTime" xml:"CreateTime" require:"true"`
	VpcId                *string                                          `json:"VpcId" xml:"VpcId" require:"true"`
	DBNodeNumber         *int                                             `json:"DBNodeNumber" xml:"DBNodeNumber" require:"true"`
	DBNodeClass          *string                                          `json:"DBNodeClass" xml:"DBNodeClass" require:"true"`
	StorageUsed          *int64                                           `json:"StorageUsed" xml:"StorageUsed" require:"true"`
	ResourceGroupId      *string                                          `json:"ResourceGroupId" xml:"ResourceGroupId" require:"true"`
	DBNodes              *DescribeDBClustersResponseItemsDBClusterDBNodes `json:"DBNodes" xml:"DBNodes" require:"true" type:"Struct"`
	Tags                 *DescribeDBClustersResponseItemsDBClusterTags    `json:"Tags" xml:"Tags" require:"true" type:"Struct"`
}

func (s DescribeDBClustersResponseItemsDBCluster) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClustersResponseItemsDBCluster) GoString() string {
	return s.String()
}

func (s *DescribeDBClustersResponseItemsDBCluster) SetDBClusterId(v string) *DescribeDBClustersResponseItemsDBCluster {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDBClustersResponseItemsDBCluster) SetDBClusterDescription(v string) *DescribeDBClustersResponseItemsDBCluster {
	s.DBClusterDescription = &v
	return s
}

func (s *DescribeDBClustersResponseItemsDBCluster) SetPayType(v string) *DescribeDBClustersResponseItemsDBCluster {
	s.PayType = &v
	return s
}

func (s *DescribeDBClustersResponseItemsDBCluster) SetDBClusterNetworkType(v string) *DescribeDBClustersResponseItemsDBCluster {
	s.DBClusterNetworkType = &v
	return s
}

func (s *DescribeDBClustersResponseItemsDBCluster) SetRegionId(v string) *DescribeDBClustersResponseItemsDBCluster {
	s.RegionId = &v
	return s
}

func (s *DescribeDBClustersResponseItemsDBCluster) SetZoneId(v string) *DescribeDBClustersResponseItemsDBCluster {
	s.ZoneId = &v
	return s
}

func (s *DescribeDBClustersResponseItemsDBCluster) SetExpireTime(v string) *DescribeDBClustersResponseItemsDBCluster {
	s.ExpireTime = &v
	return s
}

func (s *DescribeDBClustersResponseItemsDBCluster) SetExpired(v string) *DescribeDBClustersResponseItemsDBCluster {
	s.Expired = &v
	return s
}

func (s *DescribeDBClustersResponseItemsDBCluster) SetDBClusterStatus(v string) *DescribeDBClustersResponseItemsDBCluster {
	s.DBClusterStatus = &v
	return s
}

func (s *DescribeDBClustersResponseItemsDBCluster) SetEngine(v string) *DescribeDBClustersResponseItemsDBCluster {
	s.Engine = &v
	return s
}

func (s *DescribeDBClustersResponseItemsDBCluster) SetDBType(v string) *DescribeDBClustersResponseItemsDBCluster {
	s.DBType = &v
	return s
}

func (s *DescribeDBClustersResponseItemsDBCluster) SetDBVersion(v string) *DescribeDBClustersResponseItemsDBCluster {
	s.DBVersion = &v
	return s
}

func (s *DescribeDBClustersResponseItemsDBCluster) SetLockMode(v string) *DescribeDBClustersResponseItemsDBCluster {
	s.LockMode = &v
	return s
}

func (s *DescribeDBClustersResponseItemsDBCluster) SetDeletionLock(v int) *DescribeDBClustersResponseItemsDBCluster {
	s.DeletionLock = &v
	return s
}

func (s *DescribeDBClustersResponseItemsDBCluster) SetCreateTime(v string) *DescribeDBClustersResponseItemsDBCluster {
	s.CreateTime = &v
	return s
}

func (s *DescribeDBClustersResponseItemsDBCluster) SetVpcId(v string) *DescribeDBClustersResponseItemsDBCluster {
	s.VpcId = &v
	return s
}

func (s *DescribeDBClustersResponseItemsDBCluster) SetDBNodeNumber(v int) *DescribeDBClustersResponseItemsDBCluster {
	s.DBNodeNumber = &v
	return s
}

func (s *DescribeDBClustersResponseItemsDBCluster) SetDBNodeClass(v string) *DescribeDBClustersResponseItemsDBCluster {
	s.DBNodeClass = &v
	return s
}

func (s *DescribeDBClustersResponseItemsDBCluster) SetStorageUsed(v int64) *DescribeDBClustersResponseItemsDBCluster {
	s.StorageUsed = &v
	return s
}

func (s *DescribeDBClustersResponseItemsDBCluster) SetResourceGroupId(v string) *DescribeDBClustersResponseItemsDBCluster {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDBClustersResponseItemsDBCluster) SetDBNodes(v *DescribeDBClustersResponseItemsDBClusterDBNodes) *DescribeDBClustersResponseItemsDBCluster {
	s.DBNodes = v
	return s
}

func (s *DescribeDBClustersResponseItemsDBCluster) SetTags(v *DescribeDBClustersResponseItemsDBClusterTags) *DescribeDBClustersResponseItemsDBCluster {
	s.Tags = v
	return s
}

type DescribeDBClustersResponseItemsDBClusterDBNodes struct {
	DBNode []*DescribeDBClustersResponseItemsDBClusterDBNodesDBNode `json:"DBNode" xml:"DBNode" require:"true" type:"Repeated"`
}

func (s DescribeDBClustersResponseItemsDBClusterDBNodes) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClustersResponseItemsDBClusterDBNodes) GoString() string {
	return s.String()
}

func (s *DescribeDBClustersResponseItemsDBClusterDBNodes) SetDBNode(v []*DescribeDBClustersResponseItemsDBClusterDBNodesDBNode) *DescribeDBClustersResponseItemsDBClusterDBNodes {
	s.DBNode = v
	return s
}

type DescribeDBClustersResponseItemsDBClusterDBNodesDBNode struct {
	DBNodeId    *string `json:"DBNodeId" xml:"DBNodeId" require:"true"`
	DBNodeClass *string `json:"DBNodeClass" xml:"DBNodeClass" require:"true"`
	DBNodeRole  *string `json:"DBNodeRole" xml:"DBNodeRole" require:"true"`
	RegionId    *string `json:"RegionId" xml:"RegionId" require:"true"`
	ZoneId      *string `json:"ZoneId" xml:"ZoneId" require:"true"`
}

func (s DescribeDBClustersResponseItemsDBClusterDBNodesDBNode) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClustersResponseItemsDBClusterDBNodesDBNode) GoString() string {
	return s.String()
}

func (s *DescribeDBClustersResponseItemsDBClusterDBNodesDBNode) SetDBNodeId(v string) *DescribeDBClustersResponseItemsDBClusterDBNodesDBNode {
	s.DBNodeId = &v
	return s
}

func (s *DescribeDBClustersResponseItemsDBClusterDBNodesDBNode) SetDBNodeClass(v string) *DescribeDBClustersResponseItemsDBClusterDBNodesDBNode {
	s.DBNodeClass = &v
	return s
}

func (s *DescribeDBClustersResponseItemsDBClusterDBNodesDBNode) SetDBNodeRole(v string) *DescribeDBClustersResponseItemsDBClusterDBNodesDBNode {
	s.DBNodeRole = &v
	return s
}

func (s *DescribeDBClustersResponseItemsDBClusterDBNodesDBNode) SetRegionId(v string) *DescribeDBClustersResponseItemsDBClusterDBNodesDBNode {
	s.RegionId = &v
	return s
}

func (s *DescribeDBClustersResponseItemsDBClusterDBNodesDBNode) SetZoneId(v string) *DescribeDBClustersResponseItemsDBClusterDBNodesDBNode {
	s.ZoneId = &v
	return s
}

type DescribeDBClustersResponseItemsDBClusterTags struct {
	Tag []*DescribeDBClustersResponseItemsDBClusterTagsTag `json:"Tag" xml:"Tag" require:"true" type:"Repeated"`
}

func (s DescribeDBClustersResponseItemsDBClusterTags) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClustersResponseItemsDBClusterTags) GoString() string {
	return s.String()
}

func (s *DescribeDBClustersResponseItemsDBClusterTags) SetTag(v []*DescribeDBClustersResponseItemsDBClusterTagsTag) *DescribeDBClustersResponseItemsDBClusterTags {
	s.Tag = v
	return s
}

type DescribeDBClustersResponseItemsDBClusterTagsTag struct {
	Key   *string `json:"Key" xml:"Key" require:"true"`
	Value *string `json:"Value" xml:"Value" require:"true"`
}

func (s DescribeDBClustersResponseItemsDBClusterTagsTag) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClustersResponseItemsDBClusterTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeDBClustersResponseItemsDBClusterTagsTag) SetKey(v string) *DescribeDBClustersResponseItemsDBClusterTagsTag {
	s.Key = &v
	return s
}

func (s *DescribeDBClustersResponseItemsDBClusterTagsTag) SetValue(v string) *DescribeDBClustersResponseItemsDBClusterTagsTag {
	s.Value = &v
	return s
}

type DescribeDBClusterAttributeRequest struct {
	AccessKeyId          *string `json:"AccessKeyId" xml:"AccessKeyId"`
	OwnerId              *int64  `json:"OwnerId" xml:"OwnerId"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount" xml:"ResourceOwnerAccount"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId" xml:"ResourceOwnerId"`
	OwnerAccount         *string `json:"OwnerAccount" xml:"OwnerAccount"`
	DBClusterId          *string `json:"DBClusterId" xml:"DBClusterId" require:"true"`
}

func (s DescribeDBClusterAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterAttributeRequest) SetAccessKeyId(v string) *DescribeDBClusterAttributeRequest {
	s.AccessKeyId = &v
	return s
}

func (s *DescribeDBClusterAttributeRequest) SetOwnerId(v int64) *DescribeDBClusterAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDBClusterAttributeRequest) SetResourceOwnerAccount(v string) *DescribeDBClusterAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDBClusterAttributeRequest) SetResourceOwnerId(v int64) *DescribeDBClusterAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDBClusterAttributeRequest) SetOwnerAccount(v string) *DescribeDBClusterAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDBClusterAttributeRequest) SetDBClusterId(v string) *DescribeDBClusterAttributeRequest {
	s.DBClusterId = &v
	return s
}

type DescribeDBClusterAttributeResponse struct {
	RequestId                 *string                                      `json:"RequestId" xml:"RequestId" require:"true"`
	RegionId                  *string                                      `json:"RegionId" xml:"RegionId" require:"true"`
	DBClusterNetworkType      *string                                      `json:"DBClusterNetworkType" xml:"DBClusterNetworkType" require:"true"`
	VPCId                     *string                                      `json:"VPCId" xml:"VPCId" require:"true"`
	VSwitchId                 *string                                      `json:"VSwitchId" xml:"VSwitchId" require:"true"`
	PayType                   *string                                      `json:"PayType" xml:"PayType" require:"true"`
	DBClusterId               *string                                      `json:"DBClusterId" xml:"DBClusterId" require:"true"`
	DBClusterStatus           *string                                      `json:"DBClusterStatus" xml:"DBClusterStatus" require:"true"`
	DBClusterDescription      *string                                      `json:"DBClusterDescription" xml:"DBClusterDescription" require:"true"`
	Engine                    *string                                      `json:"Engine" xml:"Engine" require:"true"`
	DBType                    *string                                      `json:"DBType" xml:"DBType" require:"true"`
	DBVersion                 *string                                      `json:"DBVersion" xml:"DBVersion" require:"true"`
	LockMode                  *string                                      `json:"LockMode" xml:"LockMode" require:"true"`
	DeletionLock              *int                                         `json:"DeletionLock" xml:"DeletionLock" require:"true"`
	CreationTime              *string                                      `json:"CreationTime" xml:"CreationTime" require:"true"`
	ExpireTime                *string                                      `json:"ExpireTime" xml:"ExpireTime" require:"true"`
	Expired                   *string                                      `json:"Expired" xml:"Expired" require:"true"`
	MaintainTime              *string                                      `json:"MaintainTime" xml:"MaintainTime" require:"true"`
	StorageUsed               *int64                                       `json:"StorageUsed" xml:"StorageUsed" require:"true"`
	StorageMax                *int64                                       `json:"StorageMax" xml:"StorageMax" require:"true"`
	ZoneIds                   *string                                      `json:"ZoneIds" xml:"ZoneIds" require:"true"`
	SQLSize                   *int64                                       `json:"SQLSize" xml:"SQLSize" require:"true"`
	IsLatestVersion           *bool                                        `json:"IsLatestVersion" xml:"IsLatestVersion" require:"true"`
	ResourceGroupId           *string                                      `json:"ResourceGroupId" xml:"ResourceGroupId" require:"true"`
	DataLevel1BackupChainSize *int64                                       `json:"DataLevel1BackupChainSize" xml:"DataLevel1BackupChainSize" require:"true"`
	Tags                      []*DescribeDBClusterAttributeResponseTags    `json:"Tags" xml:"Tags" require:"true" type:"Repeated"`
	DBNodes                   []*DescribeDBClusterAttributeResponseDBNodes `json:"DBNodes" xml:"DBNodes" require:"true" type:"Repeated"`
}

func (s DescribeDBClusterAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterAttributeResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterAttributeResponse) SetRequestId(v string) *DescribeDBClusterAttributeResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeDBClusterAttributeResponse) SetRegionId(v string) *DescribeDBClusterAttributeResponse {
	s.RegionId = &v
	return s
}

func (s *DescribeDBClusterAttributeResponse) SetDBClusterNetworkType(v string) *DescribeDBClusterAttributeResponse {
	s.DBClusterNetworkType = &v
	return s
}

func (s *DescribeDBClusterAttributeResponse) SetVPCId(v string) *DescribeDBClusterAttributeResponse {
	s.VPCId = &v
	return s
}

func (s *DescribeDBClusterAttributeResponse) SetVSwitchId(v string) *DescribeDBClusterAttributeResponse {
	s.VSwitchId = &v
	return s
}

func (s *DescribeDBClusterAttributeResponse) SetPayType(v string) *DescribeDBClusterAttributeResponse {
	s.PayType = &v
	return s
}

func (s *DescribeDBClusterAttributeResponse) SetDBClusterId(v string) *DescribeDBClusterAttributeResponse {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDBClusterAttributeResponse) SetDBClusterStatus(v string) *DescribeDBClusterAttributeResponse {
	s.DBClusterStatus = &v
	return s
}

func (s *DescribeDBClusterAttributeResponse) SetDBClusterDescription(v string) *DescribeDBClusterAttributeResponse {
	s.DBClusterDescription = &v
	return s
}

func (s *DescribeDBClusterAttributeResponse) SetEngine(v string) *DescribeDBClusterAttributeResponse {
	s.Engine = &v
	return s
}

func (s *DescribeDBClusterAttributeResponse) SetDBType(v string) *DescribeDBClusterAttributeResponse {
	s.DBType = &v
	return s
}

func (s *DescribeDBClusterAttributeResponse) SetDBVersion(v string) *DescribeDBClusterAttributeResponse {
	s.DBVersion = &v
	return s
}

func (s *DescribeDBClusterAttributeResponse) SetLockMode(v string) *DescribeDBClusterAttributeResponse {
	s.LockMode = &v
	return s
}

func (s *DescribeDBClusterAttributeResponse) SetDeletionLock(v int) *DescribeDBClusterAttributeResponse {
	s.DeletionLock = &v
	return s
}

func (s *DescribeDBClusterAttributeResponse) SetCreationTime(v string) *DescribeDBClusterAttributeResponse {
	s.CreationTime = &v
	return s
}

func (s *DescribeDBClusterAttributeResponse) SetExpireTime(v string) *DescribeDBClusterAttributeResponse {
	s.ExpireTime = &v
	return s
}

func (s *DescribeDBClusterAttributeResponse) SetExpired(v string) *DescribeDBClusterAttributeResponse {
	s.Expired = &v
	return s
}

func (s *DescribeDBClusterAttributeResponse) SetMaintainTime(v string) *DescribeDBClusterAttributeResponse {
	s.MaintainTime = &v
	return s
}

func (s *DescribeDBClusterAttributeResponse) SetStorageUsed(v int64) *DescribeDBClusterAttributeResponse {
	s.StorageUsed = &v
	return s
}

func (s *DescribeDBClusterAttributeResponse) SetStorageMax(v int64) *DescribeDBClusterAttributeResponse {
	s.StorageMax = &v
	return s
}

func (s *DescribeDBClusterAttributeResponse) SetZoneIds(v string) *DescribeDBClusterAttributeResponse {
	s.ZoneIds = &v
	return s
}

func (s *DescribeDBClusterAttributeResponse) SetSQLSize(v int64) *DescribeDBClusterAttributeResponse {
	s.SQLSize = &v
	return s
}

func (s *DescribeDBClusterAttributeResponse) SetIsLatestVersion(v bool) *DescribeDBClusterAttributeResponse {
	s.IsLatestVersion = &v
	return s
}

func (s *DescribeDBClusterAttributeResponse) SetResourceGroupId(v string) *DescribeDBClusterAttributeResponse {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDBClusterAttributeResponse) SetDataLevel1BackupChainSize(v int64) *DescribeDBClusterAttributeResponse {
	s.DataLevel1BackupChainSize = &v
	return s
}

func (s *DescribeDBClusterAttributeResponse) SetTags(v []*DescribeDBClusterAttributeResponseTags) *DescribeDBClusterAttributeResponse {
	s.Tags = v
	return s
}

func (s *DescribeDBClusterAttributeResponse) SetDBNodes(v []*DescribeDBClusterAttributeResponseDBNodes) *DescribeDBClusterAttributeResponse {
	s.DBNodes = v
	return s
}

type DescribeDBClusterAttributeResponseTags struct {
	Key   *string `json:"Key" xml:"Key" require:"true"`
	Value *string `json:"Value" xml:"Value" require:"true"`
}

func (s DescribeDBClusterAttributeResponseTags) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterAttributeResponseTags) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterAttributeResponseTags) SetKey(v string) *DescribeDBClusterAttributeResponseTags {
	s.Key = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseTags) SetValue(v string) *DescribeDBClusterAttributeResponseTags {
	s.Value = &v
	return s
}

type DescribeDBClusterAttributeResponseDBNodes struct {
	DBNodeId         *string `json:"DBNodeId" xml:"DBNodeId" require:"true"`
	ZoneId           *string `json:"ZoneId" xml:"ZoneId" require:"true"`
	DBNodeStatus     *string `json:"DBNodeStatus" xml:"DBNodeStatus" require:"true"`
	CreationTime     *string `json:"CreationTime" xml:"CreationTime" require:"true"`
	DBNodeClass      *string `json:"DBNodeClass" xml:"DBNodeClass" require:"true"`
	DBNodeRole       *string `json:"DBNodeRole" xml:"DBNodeRole" require:"true"`
	MaxIOPS          *int    `json:"MaxIOPS" xml:"MaxIOPS" require:"true"`
	MaxConnections   *int    `json:"MaxConnections" xml:"MaxConnections" require:"true"`
	FailoverPriority *int    `json:"FailoverPriority" xml:"FailoverPriority" require:"true"`
}

func (s DescribeDBClusterAttributeResponseDBNodes) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterAttributeResponseDBNodes) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterAttributeResponseDBNodes) SetDBNodeId(v string) *DescribeDBClusterAttributeResponseDBNodes {
	s.DBNodeId = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseDBNodes) SetZoneId(v string) *DescribeDBClusterAttributeResponseDBNodes {
	s.ZoneId = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseDBNodes) SetDBNodeStatus(v string) *DescribeDBClusterAttributeResponseDBNodes {
	s.DBNodeStatus = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseDBNodes) SetCreationTime(v string) *DescribeDBClusterAttributeResponseDBNodes {
	s.CreationTime = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseDBNodes) SetDBNodeClass(v string) *DescribeDBClusterAttributeResponseDBNodes {
	s.DBNodeClass = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseDBNodes) SetDBNodeRole(v string) *DescribeDBClusterAttributeResponseDBNodes {
	s.DBNodeRole = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseDBNodes) SetMaxIOPS(v int) *DescribeDBClusterAttributeResponseDBNodes {
	s.MaxIOPS = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseDBNodes) SetMaxConnections(v int) *DescribeDBClusterAttributeResponseDBNodes {
	s.MaxConnections = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseDBNodes) SetFailoverPriority(v int) *DescribeDBClusterAttributeResponseDBNodes {
	s.FailoverPriority = &v
	return s
}

type Client struct {
	rpc.Client
}

func NewClient(config *rpc.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *rpc.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("regional")
	client.EndpointMap = map[string]*string{
		"cn-qingdao":            tea.String("polardb.aliyuncs.com"),
		"cn-beijing":            tea.String("polardb.aliyuncs.com"),
		"cn-hangzhou":           tea.String("polardb.aliyuncs.com"),
		"cn-shanghai":           tea.String("polardb.aliyuncs.com"),
		"cn-shenzhen":           tea.String("polardb.aliyuncs.com"),
		"cn-hongkong":           tea.String("polardb.aliyuncs.com"),
		"ap-southeast-1":        tea.String("polardb.aliyuncs.com"),
		"us-west-1":             tea.String("polardb.aliyuncs.com"),
		"cn-hangzhou-finance":   tea.String("polardb.aliyuncs.com"),
		"cn-shanghai-finance-1": tea.String("polardb.aliyuncs.com"),
		"us-east-1":             tea.String("polardb.ap-northeast-1.aliyuncs.com"),
		"cn-shenzhen-finance-1": tea.String("polardb.aliyuncs.com"),
		"cn-north-2-gov-1":      tea.String("polardb.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("polardb"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) DescribeDBClusterAuditLogCollectorWithOptions(request *DescribeDBClusterAuditLogCollectorRequest, runtime *util.RuntimeOptions) (_result *DescribeDBClusterAuditLogCollectorResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeDBClusterAuditLogCollectorResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeDBClusterAuditLogCollector"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-08-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDBClusterAuditLogCollector(request *DescribeDBClusterAuditLogCollectorRequest) (_result *DescribeDBClusterAuditLogCollectorResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDBClusterAuditLogCollectorResponse{}
	_body, _err := client.DescribeDBClusterAuditLogCollectorWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDBClusterAuditLogCollectorWithOptions(request *ModifyDBClusterAuditLogCollectorRequest, runtime *util.RuntimeOptions) (_result *ModifyDBClusterAuditLogCollectorResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ModifyDBClusterAuditLogCollectorResponse{}
	_body, _err := client.DoRequest(tea.String("ModifyDBClusterAuditLogCollector"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-08-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDBClusterAuditLogCollector(request *ModifyDBClusterAuditLogCollectorRequest) (_result *ModifyDBClusterAuditLogCollectorResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDBClusterAuditLogCollectorResponse{}
	_body, _err := client.ModifyDBClusterAuditLogCollectorWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDetachedBackupsWithOptions(request *DescribeDetachedBackupsRequest, runtime *util.RuntimeOptions) (_result *DescribeDetachedBackupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeDetachedBackupsResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeDetachedBackups"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-08-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDetachedBackups(request *DescribeDetachedBackupsRequest) (_result *DescribeDetachedBackupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDetachedBackupsResponse{}
	_body, _err := client.DescribeDetachedBackupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDBClustersWithBackupsWithOptions(request *DescribeDBClustersWithBackupsRequest, runtime *util.RuntimeOptions) (_result *DescribeDBClustersWithBackupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeDBClustersWithBackupsResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeDBClustersWithBackups"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-08-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDBClustersWithBackups(request *DescribeDBClustersWithBackupsRequest) (_result *DescribeDBClustersWithBackupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDBClustersWithBackupsResponse{}
	_body, _err := client.DescribeDBClustersWithBackupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeLogBackupPolicyWithOptions(request *DescribeLogBackupPolicyRequest, runtime *util.RuntimeOptions) (_result *DescribeLogBackupPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeLogBackupPolicyResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeLogBackupPolicy"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-08-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeLogBackupPolicy(request *DescribeLogBackupPolicyRequest) (_result *DescribeLogBackupPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeLogBackupPolicyResponse{}
	_body, _err := client.DescribeLogBackupPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyLogBackupPolicyWithOptions(request *ModifyLogBackupPolicyRequest, runtime *util.RuntimeOptions) (_result *ModifyLogBackupPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ModifyLogBackupPolicyResponse{}
	_body, _err := client.DoRequest(tea.String("ModifyLogBackupPolicy"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-08-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyLogBackupPolicy(request *ModifyLogBackupPolicyRequest) (_result *ModifyLogBackupPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyLogBackupPolicyResponse{}
	_body, _err := client.ModifyLogBackupPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDBClusterMonitorWithOptions(request *ModifyDBClusterMonitorRequest, runtime *util.RuntimeOptions) (_result *ModifyDBClusterMonitorResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ModifyDBClusterMonitorResponse{}
	_body, _err := client.DoRequest(tea.String("ModifyDBClusterMonitor"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-08-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDBClusterMonitor(request *ModifyDBClusterMonitorRequest) (_result *ModifyDBClusterMonitorResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDBClusterMonitorResponse{}
	_body, _err := client.ModifyDBClusterMonitorWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDBClusterMonitorWithOptions(request *DescribeDBClusterMonitorRequest, runtime *util.RuntimeOptions) (_result *DescribeDBClusterMonitorResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeDBClusterMonitorResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeDBClusterMonitor"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-08-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDBClusterMonitor(request *DescribeDBClusterMonitorRequest) (_result *DescribeDBClusterMonitorResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDBClusterMonitorResponse{}
	_body, _err := client.DescribeDBClusterMonitorWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDBClusterAvailableResourcesWithOptions(request *DescribeDBClusterAvailableResourcesRequest, runtime *util.RuntimeOptions) (_result *DescribeDBClusterAvailableResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeDBClusterAvailableResourcesResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeDBClusterAvailableResources"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-08-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDBClusterAvailableResources(request *DescribeDBClusterAvailableResourcesRequest) (_result *DescribeDBClusterAvailableResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDBClusterAvailableResourcesResponse{}
	_body, _err := client.DescribeDBClusterAvailableResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeBackupLogsWithOptions(request *DescribeBackupLogsRequest, runtime *util.RuntimeOptions) (_result *DescribeBackupLogsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeBackupLogsResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeBackupLogs"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-08-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeBackupLogs(request *DescribeBackupLogsRequest) (_result *DescribeBackupLogsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeBackupLogsResponse{}
	_body, _err := client.DescribeBackupLogsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDBClusterSSLWithOptions(request *ModifyDBClusterSSLRequest, runtime *util.RuntimeOptions) (_result *ModifyDBClusterSSLResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ModifyDBClusterSSLResponse{}
	_body, _err := client.DoRequest(tea.String("ModifyDBClusterSSL"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-08-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDBClusterSSL(request *ModifyDBClusterSSLRequest) (_result *ModifyDBClusterSSLResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDBClusterSSLResponse{}
	_body, _err := client.ModifyDBClusterSSLWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDBClusterSSLWithOptions(request *DescribeDBClusterSSLRequest, runtime *util.RuntimeOptions) (_result *DescribeDBClusterSSLResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeDBClusterSSLResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeDBClusterSSL"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-08-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDBClusterSSL(request *DescribeDBClusterSSLRequest) (_result *DescribeDBClusterSSLResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDBClusterSSLResponse{}
	_body, _err := client.DescribeDBClusterSSLWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDBClusterMigrationWithOptions(request *DescribeDBClusterMigrationRequest, runtime *util.RuntimeOptions) (_result *DescribeDBClusterMigrationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeDBClusterMigrationResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeDBClusterMigration"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-08-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDBClusterMigration(request *DescribeDBClusterMigrationRequest) (_result *DescribeDBClusterMigrationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDBClusterMigrationResponse{}
	_body, _err := client.DescribeDBClusterMigrationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CloseDBClusterMigrationWithOptions(request *CloseDBClusterMigrationRequest, runtime *util.RuntimeOptions) (_result *CloseDBClusterMigrationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CloseDBClusterMigrationResponse{}
	_body, _err := client.DoRequest(tea.String("CloseDBClusterMigration"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-08-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CloseDBClusterMigration(request *CloseDBClusterMigrationRequest) (_result *CloseDBClusterMigrationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CloseDBClusterMigrationResponse{}
	_body, _err := client.CloseDBClusterMigrationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDBClusterMigrationWithOptions(request *ModifyDBClusterMigrationRequest, runtime *util.RuntimeOptions) (_result *ModifyDBClusterMigrationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ModifyDBClusterMigrationResponse{}
	_body, _err := client.DoRequest(tea.String("ModifyDBClusterMigration"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-08-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDBClusterMigration(request *ModifyDBClusterMigrationRequest) (_result *ModifyDBClusterMigrationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDBClusterMigrationResponse{}
	_body, _err := client.ModifyDBClusterMigrationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyAutoRenewAttributeWithOptions(request *ModifyAutoRenewAttributeRequest, runtime *util.RuntimeOptions) (_result *ModifyAutoRenewAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ModifyAutoRenewAttributeResponse{}
	_body, _err := client.DoRequest(tea.String("ModifyAutoRenewAttribute"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-08-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyAutoRenewAttribute(request *ModifyAutoRenewAttributeRequest) (_result *ModifyAutoRenewAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyAutoRenewAttributeResponse{}
	_body, _err := client.ModifyAutoRenewAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDBNodeClassWithOptions(request *ModifyDBNodeClassRequest, runtime *util.RuntimeOptions) (_result *ModifyDBNodeClassResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ModifyDBNodeClassResponse{}
	_body, _err := client.DoRequest(tea.String("ModifyDBNodeClass"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-08-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDBNodeClass(request *ModifyDBNodeClassRequest) (_result *ModifyDBNodeClassResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDBNodeClassResponse{}
	_body, _err := client.ModifyDBNodeClassWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAutoRenewAttributeWithOptions(request *DescribeAutoRenewAttributeRequest, runtime *util.RuntimeOptions) (_result *DescribeAutoRenewAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeAutoRenewAttributeResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeAutoRenewAttribute"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-08-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAutoRenewAttribute(request *DescribeAutoRenewAttributeRequest) (_result *DescribeAutoRenewAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAutoRenewAttributeResponse{}
	_body, _err := client.DescribeAutoRenewAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateDBNodesWithOptions(request *CreateDBNodesRequest, runtime *util.RuntimeOptions) (_result *CreateDBNodesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CreateDBNodesResponse{}
	_body, _err := client.DoRequest(tea.String("CreateDBNodes"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-08-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDBNodes(request *CreateDBNodesRequest) (_result *CreateDBNodesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDBNodesResponse{}
	_body, _err := client.CreateDBNodesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteDBNodesWithOptions(request *DeleteDBNodesRequest, runtime *util.RuntimeOptions) (_result *DeleteDBNodesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DeleteDBNodesResponse{}
	_body, _err := client.DoRequest(tea.String("DeleteDBNodes"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-08-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteDBNodes(request *DeleteDBNodesRequest) (_result *DeleteDBNodesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDBNodesResponse{}
	_body, _err := client.DeleteDBNodesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UntagResourcesWithOptions(request *UntagResourcesRequest, runtime *util.RuntimeOptions) (_result *UntagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &UntagResourcesResponse{}
	_body, _err := client.DoRequest(tea.String("UntagResources"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-08-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UntagResources(request *UntagResourcesRequest) (_result *UntagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UntagResourcesResponse{}
	_body, _err := client.UntagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) TagResourcesWithOptions(request *TagResourcesRequest, runtime *util.RuntimeOptions) (_result *TagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &TagResourcesResponse{}
	_body, _err := client.DoRequest(tea.String("TagResources"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-08-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TagResources(request *TagResourcesRequest) (_result *TagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TagResourcesResponse{}
	_body, _err := client.TagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListTagResourcesWithOptions(request *ListTagResourcesRequest, runtime *util.RuntimeOptions) (_result *ListTagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.DoRequest(tea.String("ListTagResources"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-08-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTagResources(request *ListTagResourcesRequest) (_result *ListTagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.ListTagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDBEndpointAddressWithOptions(request *ModifyDBEndpointAddressRequest, runtime *util.RuntimeOptions) (_result *ModifyDBEndpointAddressResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ModifyDBEndpointAddressResponse{}
	_body, _err := client.DoRequest(tea.String("ModifyDBEndpointAddress"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-08-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDBEndpointAddress(request *ModifyDBEndpointAddressRequest) (_result *ModifyDBEndpointAddressResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDBEndpointAddressResponse{}
	_body, _err := client.ModifyDBEndpointAddressWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDBDescriptionWithOptions(request *ModifyDBDescriptionRequest, runtime *util.RuntimeOptions) (_result *ModifyDBDescriptionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ModifyDBDescriptionResponse{}
	_body, _err := client.DoRequest(tea.String("ModifyDBDescription"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-08-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDBDescription(request *ModifyDBDescriptionRequest) (_result *ModifyDBDescriptionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDBDescriptionResponse{}
	_body, _err := client.ModifyDBDescriptionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDBClusterParametersWithOptions(request *ModifyDBClusterParametersRequest, runtime *util.RuntimeOptions) (_result *ModifyDBClusterParametersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ModifyDBClusterParametersResponse{}
	_body, _err := client.DoRequest(tea.String("ModifyDBClusterParameters"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-08-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDBClusterParameters(request *ModifyDBClusterParametersRequest) (_result *ModifyDBClusterParametersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDBClusterParametersResponse{}
	_body, _err := client.ModifyDBClusterParametersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDBClusterEndpointWithOptions(request *ModifyDBClusterEndpointRequest, runtime *util.RuntimeOptions) (_result *ModifyDBClusterEndpointResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ModifyDBClusterEndpointResponse{}
	_body, _err := client.DoRequest(tea.String("ModifyDBClusterEndpoint"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-08-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDBClusterEndpoint(request *ModifyDBClusterEndpointRequest) (_result *ModifyDBClusterEndpointResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDBClusterEndpointResponse{}
	_body, _err := client.ModifyDBClusterEndpointWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyAccountPasswordWithOptions(request *ModifyAccountPasswordRequest, runtime *util.RuntimeOptions) (_result *ModifyAccountPasswordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ModifyAccountPasswordResponse{}
	_body, _err := client.DoRequest(tea.String("ModifyAccountPassword"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-08-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyAccountPassword(request *ModifyAccountPasswordRequest) (_result *ModifyAccountPasswordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyAccountPasswordResponse{}
	_body, _err := client.ModifyAccountPasswordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDBClusterPerformanceWithOptions(request *DescribeDBClusterPerformanceRequest, runtime *util.RuntimeOptions) (_result *DescribeDBClusterPerformanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeDBClusterPerformanceResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeDBClusterPerformance"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-08-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDBClusterPerformance(request *DescribeDBClusterPerformanceRequest) (_result *DescribeDBClusterPerformanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDBClusterPerformanceResponse{}
	_body, _err := client.DescribeDBClusterPerformanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDBClusterParametersWithOptions(request *DescribeDBClusterParametersRequest, runtime *util.RuntimeOptions) (_result *DescribeDBClusterParametersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeDBClusterParametersResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeDBClusterParameters"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-08-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDBClusterParameters(request *DescribeDBClusterParametersRequest) (_result *DescribeDBClusterParametersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDBClusterParametersResponse{}
	_body, _err := client.DescribeDBClusterParametersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDBClusterEndpointsWithOptions(request *DescribeDBClusterEndpointsRequest, runtime *util.RuntimeOptions) (_result *DescribeDBClusterEndpointsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeDBClusterEndpointsResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeDBClusterEndpoints"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-08-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDBClusterEndpoints(request *DescribeDBClusterEndpointsRequest) (_result *DescribeDBClusterEndpointsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDBClusterEndpointsResponse{}
	_body, _err := client.DescribeDBClusterEndpointsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteDBEndpointAddressWithOptions(request *DeleteDBEndpointAddressRequest, runtime *util.RuntimeOptions) (_result *DeleteDBEndpointAddressResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DeleteDBEndpointAddressResponse{}
	_body, _err := client.DoRequest(tea.String("DeleteDBEndpointAddress"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-08-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteDBEndpointAddress(request *DeleteDBEndpointAddressRequest) (_result *DeleteDBEndpointAddressResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDBEndpointAddressResponse{}
	_body, _err := client.DeleteDBEndpointAddressWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteDBClusterEndpointWithOptions(request *DeleteDBClusterEndpointRequest, runtime *util.RuntimeOptions) (_result *DeleteDBClusterEndpointResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DeleteDBClusterEndpointResponse{}
	_body, _err := client.DoRequest(tea.String("DeleteDBClusterEndpoint"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-08-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteDBClusterEndpoint(request *DeleteDBClusterEndpointRequest) (_result *DeleteDBClusterEndpointResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDBClusterEndpointResponse{}
	_body, _err := client.DeleteDBClusterEndpointWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateDBEndpointAddressWithOptions(request *CreateDBEndpointAddressRequest, runtime *util.RuntimeOptions) (_result *CreateDBEndpointAddressResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CreateDBEndpointAddressResponse{}
	_body, _err := client.DoRequest(tea.String("CreateDBEndpointAddress"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-08-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDBEndpointAddress(request *CreateDBEndpointAddressRequest) (_result *CreateDBEndpointAddressResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDBEndpointAddressResponse{}
	_body, _err := client.CreateDBEndpointAddressWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateDBClusterEndpointWithOptions(request *CreateDBClusterEndpointRequest, runtime *util.RuntimeOptions) (_result *CreateDBClusterEndpointResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CreateDBClusterEndpointResponse{}
	_body, _err := client.DoRequest(tea.String("CreateDBClusterEndpoint"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-08-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDBClusterEndpoint(request *CreateDBClusterEndpointRequest) (_result *CreateDBClusterEndpointResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDBClusterEndpointResponse{}
	_body, _err := client.CreateDBClusterEndpointWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RestartDBNodeWithOptions(request *RestartDBNodeRequest, runtime *util.RuntimeOptions) (_result *RestartDBNodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &RestartDBNodeResponse{}
	_body, _err := client.DoRequest(tea.String("RestartDBNode"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-08-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RestartDBNode(request *RestartDBNodeRequest) (_result *RestartDBNodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RestartDBNodeResponse{}
	_body, _err := client.RestartDBNodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDBNodePerformanceWithOptions(request *DescribeDBNodePerformanceRequest, runtime *util.RuntimeOptions) (_result *DescribeDBNodePerformanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeDBNodePerformanceResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeDBNodePerformance"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-08-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDBNodePerformance(request *DescribeDBNodePerformanceRequest) (_result *DescribeDBNodePerformanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDBNodePerformanceResponse{}
	_body, _err := client.DescribeDBNodePerformanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDBClusterAccessWhitelistWithOptions(request *ModifyDBClusterAccessWhitelistRequest, runtime *util.RuntimeOptions) (_result *ModifyDBClusterAccessWhitelistResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ModifyDBClusterAccessWhitelistResponse{}
	_body, _err := client.DoRequest(tea.String("ModifyDBClusterAccessWhitelist"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-08-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDBClusterAccessWhitelist(request *ModifyDBClusterAccessWhitelistRequest) (_result *ModifyDBClusterAccessWhitelistResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDBClusterAccessWhitelistResponse{}
	_body, _err := client.ModifyDBClusterAccessWhitelistWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDBClusterAccessWhitelistWithOptions(request *DescribeDBClusterAccessWhitelistRequest, runtime *util.RuntimeOptions) (_result *DescribeDBClusterAccessWhitelistResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeDBClusterAccessWhitelistResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeDBClusterAccessWhitelist"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-08-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDBClusterAccessWhitelist(request *DescribeDBClusterAccessWhitelistRequest) (_result *DescribeDBClusterAccessWhitelistResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDBClusterAccessWhitelistResponse{}
	_body, _err := client.DescribeDBClusterAccessWhitelistWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDBClusterMaintainTimeWithOptions(request *ModifyDBClusterMaintainTimeRequest, runtime *util.RuntimeOptions) (_result *ModifyDBClusterMaintainTimeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ModifyDBClusterMaintainTimeResponse{}
	_body, _err := client.DoRequest(tea.String("ModifyDBClusterMaintainTime"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-08-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDBClusterMaintainTime(request *ModifyDBClusterMaintainTimeRequest) (_result *ModifyDBClusterMaintainTimeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDBClusterMaintainTimeResponse{}
	_body, _err := client.ModifyDBClusterMaintainTimeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RevokeAccountPrivilegeWithOptions(request *RevokeAccountPrivilegeRequest, runtime *util.RuntimeOptions) (_result *RevokeAccountPrivilegeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &RevokeAccountPrivilegeResponse{}
	_body, _err := client.DoRequest(tea.String("RevokeAccountPrivilege"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-08-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RevokeAccountPrivilege(request *RevokeAccountPrivilegeRequest) (_result *RevokeAccountPrivilegeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RevokeAccountPrivilegeResponse{}
	_body, _err := client.RevokeAccountPrivilegeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ResetAccountWithOptions(request *ResetAccountRequest, runtime *util.RuntimeOptions) (_result *ResetAccountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ResetAccountResponse{}
	_body, _err := client.DoRequest(tea.String("ResetAccount"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-08-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ResetAccount(request *ResetAccountRequest) (_result *ResetAccountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ResetAccountResponse{}
	_body, _err := client.ResetAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GrantAccountPrivilegeWithOptions(request *GrantAccountPrivilegeRequest, runtime *util.RuntimeOptions) (_result *GrantAccountPrivilegeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GrantAccountPrivilegeResponse{}
	_body, _err := client.DoRequest(tea.String("GrantAccountPrivilege"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-08-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GrantAccountPrivilege(request *GrantAccountPrivilegeRequest) (_result *GrantAccountPrivilegeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GrantAccountPrivilegeResponse{}
	_body, _err := client.GrantAccountPrivilegeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDatabasesWithOptions(request *DescribeDatabasesRequest, runtime *util.RuntimeOptions) (_result *DescribeDatabasesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeDatabasesResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeDatabases"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-08-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDatabases(request *DescribeDatabasesRequest) (_result *DescribeDatabasesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDatabasesResponse{}
	_body, _err := client.DescribeDatabasesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteDatabaseWithOptions(request *DeleteDatabaseRequest, runtime *util.RuntimeOptions) (_result *DeleteDatabaseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DeleteDatabaseResponse{}
	_body, _err := client.DoRequest(tea.String("DeleteDatabase"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-08-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteDatabase(request *DeleteDatabaseRequest) (_result *DeleteDatabaseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDatabaseResponse{}
	_body, _err := client.DeleteDatabaseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteAccountWithOptions(request *DeleteAccountRequest, runtime *util.RuntimeOptions) (_result *DeleteAccountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DeleteAccountResponse{}
	_body, _err := client.DoRequest(tea.String("DeleteAccount"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-08-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteAccount(request *DeleteAccountRequest) (_result *DeleteAccountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteAccountResponse{}
	_body, _err := client.DeleteAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateDatabaseWithOptions(request *CreateDatabaseRequest, runtime *util.RuntimeOptions) (_result *CreateDatabaseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CreateDatabaseResponse{}
	_body, _err := client.DoRequest(tea.String("CreateDatabase"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-08-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDatabase(request *CreateDatabaseRequest) (_result *CreateDatabaseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDatabaseResponse{}
	_body, _err := client.CreateDatabaseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteBackupWithOptions(request *DeleteBackupRequest, runtime *util.RuntimeOptions) (_result *DeleteBackupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DeleteBackupResponse{}
	_body, _err := client.DoRequest(tea.String("DeleteBackup"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-08-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteBackup(request *DeleteBackupRequest) (_result *DeleteBackupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteBackupResponse{}
	_body, _err := client.DeleteBackupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSlowLogsWithOptions(request *DescribeSlowLogsRequest, runtime *util.RuntimeOptions) (_result *DescribeSlowLogsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeSlowLogsResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeSlowLogs"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-08-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSlowLogs(request *DescribeSlowLogsRequest) (_result *DescribeSlowLogsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSlowLogsResponse{}
	_body, _err := client.DescribeSlowLogsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSlowLogRecordsWithOptions(request *DescribeSlowLogRecordsRequest, runtime *util.RuntimeOptions) (_result *DescribeSlowLogRecordsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeSlowLogRecordsResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeSlowLogRecords"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-08-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSlowLogRecords(request *DescribeSlowLogRecordsRequest) (_result *DescribeSlowLogRecordsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSlowLogRecordsResponse{}
	_body, _err := client.DescribeSlowLogRecordsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) FailoverDBClusterWithOptions(request *FailoverDBClusterRequest, runtime *util.RuntimeOptions) (_result *FailoverDBClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &FailoverDBClusterResponse{}
	_body, _err := client.DoRequest(tea.String("FailoverDBCluster"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-08-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) FailoverDBCluster(request *FailoverDBClusterRequest) (_result *FailoverDBClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &FailoverDBClusterResponse{}
	_body, _err := client.FailoverDBClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDBClusterDescriptionWithOptions(request *ModifyDBClusterDescriptionRequest, runtime *util.RuntimeOptions) (_result *ModifyDBClusterDescriptionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ModifyDBClusterDescriptionResponse{}
	_body, _err := client.DoRequest(tea.String("ModifyDBClusterDescription"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-08-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDBClusterDescription(request *ModifyDBClusterDescriptionRequest) (_result *ModifyDBClusterDescriptionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDBClusterDescriptionResponse{}
	_body, _err := client.ModifyDBClusterDescriptionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyAccountDescriptionWithOptions(request *ModifyAccountDescriptionRequest, runtime *util.RuntimeOptions) (_result *ModifyAccountDescriptionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ModifyAccountDescriptionResponse{}
	_body, _err := client.DoRequest(tea.String("ModifyAccountDescription"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-08-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyAccountDescription(request *ModifyAccountDescriptionRequest) (_result *ModifyAccountDescriptionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyAccountDescriptionResponse{}
	_body, _err := client.ModifyAccountDescriptionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRegionsWithOptions(request *DescribeRegionsRequest, runtime *util.RuntimeOptions) (_result *DescribeRegionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeRegions"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-08-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRegions(request *DescribeRegionsRequest) (_result *DescribeRegionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.DescribeRegionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAccountsWithOptions(request *DescribeAccountsRequest, runtime *util.RuntimeOptions) (_result *DescribeAccountsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeAccountsResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeAccounts"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-08-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAccounts(request *DescribeAccountsRequest) (_result *DescribeAccountsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAccountsResponse{}
	_body, _err := client.DescribeAccountsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteDBClusterWithOptions(request *DeleteDBClusterRequest, runtime *util.RuntimeOptions) (_result *DeleteDBClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DeleteDBClusterResponse{}
	_body, _err := client.DoRequest(tea.String("DeleteDBCluster"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-08-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteDBCluster(request *DeleteDBClusterRequest) (_result *DeleteDBClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDBClusterResponse{}
	_body, _err := client.DeleteDBClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateDBClusterWithOptions(request *CreateDBClusterRequest, runtime *util.RuntimeOptions) (_result *CreateDBClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CreateDBClusterResponse{}
	_body, _err := client.DoRequest(tea.String("CreateDBCluster"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-08-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDBCluster(request *CreateDBClusterRequest) (_result *CreateDBClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDBClusterResponse{}
	_body, _err := client.CreateDBClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateBackupWithOptions(request *CreateBackupRequest, runtime *util.RuntimeOptions) (_result *CreateBackupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CreateBackupResponse{}
	_body, _err := client.DoRequest(tea.String("CreateBackup"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-08-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateBackup(request *CreateBackupRequest) (_result *CreateBackupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateBackupResponse{}
	_body, _err := client.CreateBackupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateAccountWithOptions(request *CreateAccountRequest, runtime *util.RuntimeOptions) (_result *CreateAccountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CreateAccountResponse{}
	_body, _err := client.DoRequest(tea.String("CreateAccount"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-08-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateAccount(request *CreateAccountRequest) (_result *CreateAccountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateAccountResponse{}
	_body, _err := client.CreateAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeBackupsWithOptions(request *DescribeBackupsRequest, runtime *util.RuntimeOptions) (_result *DescribeBackupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeBackupsResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeBackups"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-08-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeBackups(request *DescribeBackupsRequest) (_result *DescribeBackupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeBackupsResponse{}
	_body, _err := client.DescribeBackupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyBackupPolicyWithOptions(request *ModifyBackupPolicyRequest, runtime *util.RuntimeOptions) (_result *ModifyBackupPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ModifyBackupPolicyResponse{}
	_body, _err := client.DoRequest(tea.String("ModifyBackupPolicy"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-08-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyBackupPolicy(request *ModifyBackupPolicyRequest) (_result *ModifyBackupPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyBackupPolicyResponse{}
	_body, _err := client.ModifyBackupPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeBackupPolicyWithOptions(request *DescribeBackupPolicyRequest, runtime *util.RuntimeOptions) (_result *DescribeBackupPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeBackupPolicyResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeBackupPolicy"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-08-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeBackupPolicy(request *DescribeBackupPolicyRequest) (_result *DescribeBackupPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeBackupPolicyResponse{}
	_body, _err := client.DescribeBackupPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDBClustersWithOptions(request *DescribeDBClustersRequest, runtime *util.RuntimeOptions) (_result *DescribeDBClustersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeDBClustersResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeDBClusters"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-08-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDBClusters(request *DescribeDBClustersRequest) (_result *DescribeDBClustersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDBClustersResponse{}
	_body, _err := client.DescribeDBClustersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDBClusterAttributeWithOptions(request *DescribeDBClusterAttributeRequest, runtime *util.RuntimeOptions) (_result *DescribeDBClusterAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeDBClusterAttributeResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeDBClusterAttribute"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-08-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDBClusterAttribute(request *DescribeDBClusterAttributeRequest) (_result *DescribeDBClusterAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDBClusterAttributeResponse{}
	_body, _err := client.DescribeDBClusterAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
