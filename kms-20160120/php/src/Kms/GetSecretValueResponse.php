<?php

// This file is auto-generated, don't edit it. Thanks.

namespace AlibabaCloud\SDK\Kms\V20160120\Kms;

use AlibabaCloud\Tea\Model;

class GetSecretValueResponse extends Model
{
    /**
     * @description RequestId
     *
     * @var string
     */
    public $requestId;
    /**
     * @description SecretName
     *
     * @var string
     */
    public $secretName;
    /**
     * @description VersionId
     *
     * @var string
     */
    public $versionId;
    /**
     * @description CreateTime
     *
     * @var string
     */
    public $createTime;
    /**
     * @description SecretData
     *
     * @var string
     */
    public $secretData;
    /**
     * @description SecretDataType
     *
     * @var string
     */
    public $secretDataType;
    /**
     * @description VersionStages
     *
     * @var array
     */
    public $versionStages;
    protected $_name = [
        'requestId'      => 'RequestId',
        'secretName'     => 'SecretName',
        'versionId'      => 'VersionId',
        'createTime'     => 'CreateTime',
        'secretData'     => 'SecretData',
        'secretDataType' => 'SecretDataType',
        'versionStages'  => 'VersionStages',
    ];

    public function validate()
    {
        Model::validateRequired('requestId', $this->requestId, true);
        Model::validateRequired('secretName', $this->secretName, true);
        Model::validateRequired('versionId', $this->versionId, true);
        Model::validateRequired('createTime', $this->createTime, true);
        Model::validateRequired('secretData', $this->secretData, true);
        Model::validateRequired('secretDataType', $this->secretDataType, true);
        Model::validateRequired('versionStages', $this->versionStages, true);
    }

    public function toMap()
    {
        $res                   = [];
        $res['RequestId']      = $this->requestId;
        $res['SecretName']     = $this->secretName;
        $res['VersionId']      = $this->versionId;
        $res['CreateTime']     = $this->createTime;
        $res['SecretData']     = $this->secretData;
        $res['SecretDataType'] = $this->secretDataType;
        $res['VersionStages']  = [];
        if (null !== $this->versionStages && \is_array($this->versionStages)) {
            $n = 0;
            foreach ($this->versionStages as $item) {
                $res['VersionStages'][$n++] = null !== $item ? $item->toMap() : $item;
            }
        }

        return $res;
    }

    /**
     * @param array $map
     *
     * @return GetSecretValueResponse
     */
    public static function fromMap($map = [])
    {
        $model = new self();
        if (isset($map['RequestId'])) {
            $model->requestId = $map['RequestId'];
        }
        if (isset($map['SecretName'])) {
            $model->secretName = $map['SecretName'];
        }
        if (isset($map['VersionId'])) {
            $model->versionId = $map['VersionId'];
        }
        if (isset($map['CreateTime'])) {
            $model->createTime = $map['CreateTime'];
        }
        if (isset($map['SecretData'])) {
            $model->secretData = $map['SecretData'];
        }
        if (isset($map['SecretDataType'])) {
            $model->secretDataType = $map['SecretDataType'];
        }
        if (isset($map['VersionStages'])) {
            if (!empty($map['VersionStages'])) {
                $model->versionStages = [];
                $n                    = 0;
                foreach ($map['VersionStages'] as $item) {
                    $model->versionStages[$n++] = null !== $item ? GetSecretValueResponse\versionStages::fromMap($item) : $item;
                }
            }
        }

        return $model;
    }
}
