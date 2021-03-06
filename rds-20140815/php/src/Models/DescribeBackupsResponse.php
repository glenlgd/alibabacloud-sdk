<?php

// This file is auto-generated, don't edit it. Thanks.

namespace AlibabaCloud\SDK\Rds\V20140815\Models;

use AlibabaCloud\SDK\Rds\V20140815\Models\DescribeBackupsResponse\items;
use AlibabaCloud\Tea\Model;

class DescribeBackupsResponse extends Model
{
    /**
     * @description requestId
     *
     * @var string
     */
    public $requestId;

    /**
     * @description data.totalRecords
     *
     * @var string
     */
    public $totalRecordCount;

    /**
     * @description data.pageNumbers
     *
     * @var string
     */
    public $pageNumber;

    /**
     * @description data.backupsetNumbers
     *
     * @var string
     */
    public $pageRecordCount;

    /**
     * @description data.totalBackupSize
     *
     * @var int
     */
    public $totalBackupSize;

    /**
     * @description data.totalEcsSnapshotSize
     *
     * @var int
     */
    public $totalEcsSnapshotSize;

    /**
     * @description data.backupsetItems
     *
     * @var items
     */
    public $items;
    protected $_name = [
        'requestId'            => 'RequestId',
        'totalRecordCount'     => 'TotalRecordCount',
        'pageNumber'           => 'PageNumber',
        'pageRecordCount'      => 'PageRecordCount',
        'totalBackupSize'      => 'TotalBackupSize',
        'totalEcsSnapshotSize' => 'TotalEcsSnapshotSize',
        'items'                => 'Items',
    ];

    public function validate()
    {
        Model::validateRequired('requestId', $this->requestId, true);
        Model::validateRequired('totalRecordCount', $this->totalRecordCount, true);
        Model::validateRequired('pageNumber', $this->pageNumber, true);
        Model::validateRequired('pageRecordCount', $this->pageRecordCount, true);
        Model::validateRequired('totalBackupSize', $this->totalBackupSize, true);
        Model::validateRequired('totalEcsSnapshotSize', $this->totalEcsSnapshotSize, true);
        Model::validateRequired('items', $this->items, true);
    }

    public function toMap()
    {
        $res = [];
        if (null !== $this->requestId) {
            $res['RequestId'] = $this->requestId;
        }
        if (null !== $this->totalRecordCount) {
            $res['TotalRecordCount'] = $this->totalRecordCount;
        }
        if (null !== $this->pageNumber) {
            $res['PageNumber'] = $this->pageNumber;
        }
        if (null !== $this->pageRecordCount) {
            $res['PageRecordCount'] = $this->pageRecordCount;
        }
        if (null !== $this->totalBackupSize) {
            $res['TotalBackupSize'] = $this->totalBackupSize;
        }
        if (null !== $this->totalEcsSnapshotSize) {
            $res['TotalEcsSnapshotSize'] = $this->totalEcsSnapshotSize;
        }
        if (null !== $this->items) {
            $res['Items'] = null !== $this->items ? $this->items->toMap() : null;
        }

        return $res;
    }

    /**
     * @param array $map
     *
     * @return DescribeBackupsResponse
     */
    public static function fromMap($map = [])
    {
        $model = new self();
        if (isset($map['RequestId'])) {
            $model->requestId = $map['RequestId'];
        }
        if (isset($map['TotalRecordCount'])) {
            $model->totalRecordCount = $map['TotalRecordCount'];
        }
        if (isset($map['PageNumber'])) {
            $model->pageNumber = $map['PageNumber'];
        }
        if (isset($map['PageRecordCount'])) {
            $model->pageRecordCount = $map['PageRecordCount'];
        }
        if (isset($map['TotalBackupSize'])) {
            $model->totalBackupSize = $map['TotalBackupSize'];
        }
        if (isset($map['TotalEcsSnapshotSize'])) {
            $model->totalEcsSnapshotSize = $map['TotalEcsSnapshotSize'];
        }
        if (isset($map['Items'])) {
            $model->items = items::fromMap($map['Items']);
        }

        return $model;
    }
}
