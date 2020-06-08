<?php

// This file is auto-generated, don't edit it. Thanks.
namespace AlibabaCloud\SDK\Rds\V20140815\Models;

use AlibabaCloud\Tea\Model;

class DescribeActionEventPolicyResponse extends Model {
    protected $_name = [
        'requestId' => 'RequestId',
        'regionId' => 'RegionId',
        'enableEventLog' => 'EnableEventLog',
    ];
    public function validate() {
        Model::validateRequired('requestId', $this->requestId, true);
        Model::validateRequired('regionId', $this->regionId, true);
        Model::validateRequired('enableEventLog', $this->enableEventLog, true);
    }
    public function toMap() {
        $res = [];
        $res['RequestId'] = $this->requestId;
        $res['RegionId'] = $this->regionId;
        $res['EnableEventLog'] = $this->enableEventLog;
        return $res;
    }
    /**
     * @param array $map
     * @return DescribeActionEventPolicyResponse
     */
    public static function fromMap($map = []) {
        $model = new self();
        if(isset($map['RequestId'])){
            $model->requestId = $map['RequestId'];
        }
        if(isset($map['RegionId'])){
            $model->regionId = $map['RegionId'];
        }
        if(isset($map['EnableEventLog'])){
            $model->enableEventLog = $map['EnableEventLog'];
        }
        return $model;
    }
    /**
     * @description requestId
     * @var string
     */
    public $requestId;

    /**
     * @description data.regionId
     * @var string
     */
    public $regionId;

    /**
     * @description data.enableEventLog
     * @var string
     */
    public $enableEventLog;

}