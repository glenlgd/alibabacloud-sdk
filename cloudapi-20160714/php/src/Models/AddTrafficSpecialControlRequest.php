<?php

// This file is auto-generated, don't edit it. Thanks.

namespace AlibabaCloud\SDK\CloudAPI\V20160714\Models;

use AlibabaCloud\Tea\Model;

class AddTrafficSpecialControlRequest extends Model
{
    /**
     * @description securityToken
     *
     * @var string
     */
    public $securityToken;

    /**
     * @description accessKeyId
     *
     * @var string
     */
    public $accessKeyId;

    /**
     * @description trafficControlId
     *
     * @var string
     */
    public $trafficControlId;

    /**
     * @description specialType
     *
     * @var string
     */
    public $specialType;

    /**
     * @description specialKey
     *
     * @var string
     */
    public $specialKey;

    /**
     * @description trafficValue
     *
     * @var int
     */
    public $trafficValue;
    protected $_name = [
        'securityToken'    => 'SecurityToken',
        'accessKeyId'      => 'AccessKeyId',
        'trafficControlId' => 'TrafficControlId',
        'specialType'      => 'SpecialType',
        'specialKey'       => 'SpecialKey',
        'trafficValue'     => 'TrafficValue',
    ];

    public function validate()
    {
        Model::validateRequired('trafficControlId', $this->trafficControlId, true);
        Model::validateRequired('specialType', $this->specialType, true);
        Model::validateRequired('specialKey', $this->specialKey, true);
        Model::validateRequired('trafficValue', $this->trafficValue, true);
    }

    public function toMap()
    {
        $res = [];
        if (null !== $this->securityToken) {
            $res['SecurityToken'] = $this->securityToken;
        }
        if (null !== $this->accessKeyId) {
            $res['AccessKeyId'] = $this->accessKeyId;
        }
        if (null !== $this->trafficControlId) {
            $res['TrafficControlId'] = $this->trafficControlId;
        }
        if (null !== $this->specialType) {
            $res['SpecialType'] = $this->specialType;
        }
        if (null !== $this->specialKey) {
            $res['SpecialKey'] = $this->specialKey;
        }
        if (null !== $this->trafficValue) {
            $res['TrafficValue'] = $this->trafficValue;
        }

        return $res;
    }

    /**
     * @param array $map
     *
     * @return AddTrafficSpecialControlRequest
     */
    public static function fromMap($map = [])
    {
        $model = new self();
        if (isset($map['SecurityToken'])) {
            $model->securityToken = $map['SecurityToken'];
        }
        if (isset($map['AccessKeyId'])) {
            $model->accessKeyId = $map['AccessKeyId'];
        }
        if (isset($map['TrafficControlId'])) {
            $model->trafficControlId = $map['TrafficControlId'];
        }
        if (isset($map['SpecialType'])) {
            $model->specialType = $map['SpecialType'];
        }
        if (isset($map['SpecialKey'])) {
            $model->specialKey = $map['SpecialKey'];
        }
        if (isset($map['TrafficValue'])) {
            $model->trafficValue = $map['TrafficValue'];
        }

        return $model;
    }
}
