<?php

// This file is auto-generated, don't edit it. Thanks.
namespace AlibabaCloud\SDK\Polardb\V20170801\Models;

use AlibabaCloud\Tea\Model;

class DescribeDBNodePerformanceRequest extends Model {
    protected $_name = [
        'accessKeyId' => 'AccessKeyId',
        'DBNodeId' => 'DBNodeId',
        'key' => 'Key',
        'startTime' => 'StartTime',
        'endTime' => 'EndTime',
        'DBClusterId' => 'DBClusterId',
    ];
    public function validate() {
        Model::validateRequired('DBNodeId', $this->DBNodeId, true);
        Model::validateRequired('key', $this->key, true);
        Model::validateRequired('startTime', $this->startTime, true);
        Model::validateRequired('endTime', $this->endTime, true);
    }
    public function toMap() {
        $res = [];
        $res['AccessKeyId'] = $this->accessKeyId;
        $res['DBNodeId'] = $this->DBNodeId;
        $res['Key'] = $this->key;
        $res['StartTime'] = $this->startTime;
        $res['EndTime'] = $this->endTime;
        $res['DBClusterId'] = $this->DBClusterId;
        return $res;
    }
    /**
     * @param array $map
     * @return DescribeDBNodePerformanceRequest
     */
    public static function fromMap($map = []) {
        $model = new self();
        if(isset($map['AccessKeyId'])){
            $model->accessKeyId = $map['AccessKeyId'];
        }
        if(isset($map['DBNodeId'])){
            $model->DBNodeId = $map['DBNodeId'];
        }
        if(isset($map['Key'])){
            $model->key = $map['Key'];
        }
        if(isset($map['StartTime'])){
            $model->startTime = $map['StartTime'];
        }
        if(isset($map['EndTime'])){
            $model->endTime = $map['EndTime'];
        }
        if(isset($map['DBClusterId'])){
            $model->DBClusterId = $map['DBClusterId'];
        }
        return $model;
    }
    /**
     * @description appKey
     * @var string
     */
    public $accessKeyId;

    /**
     * @description dbInstanceId
     * @var string
     */
    public $DBNodeId;

    /**
     * @description key
     * @var string
     */
    public $key;

    /**
     * @description startTime
     * @var string
     */
    public $startTime;

    /**
     * @description endTime
     * @var string
     */
    public $endTime;

    /**
     * @description dbClusterId
     * @var string
     */
    public $DBClusterId;

}
