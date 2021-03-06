<?php

// This file is auto-generated, don't edit it. Thanks.

namespace AlibabaCloud\SDK\Facebody\V20191230\Models;

use AlibabaCloud\Tea\Model;

class AddFaceEntityRequest extends Model
{
    /**
     * @description dbName
     *
     * @var string
     */
    public $dbName;

    /**
     * @description entityId
     *
     * @var string
     */
    public $entityId;

    /**
     * @description labels
     *
     * @var string
     */
    public $labels;
    protected $_name = [
        'dbName'   => 'DbName',
        'entityId' => 'EntityId',
        'labels'   => 'Labels',
    ];

    public function validate()
    {
        Model::validateRequired('dbName', $this->dbName, true);
        Model::validateRequired('entityId', $this->entityId, true);
    }

    public function toMap()
    {
        $res = [];
        if (null !== $this->dbName) {
            $res['DbName'] = $this->dbName;
        }
        if (null !== $this->entityId) {
            $res['EntityId'] = $this->entityId;
        }
        if (null !== $this->labels) {
            $res['Labels'] = $this->labels;
        }

        return $res;
    }

    /**
     * @param array $map
     *
     * @return AddFaceEntityRequest
     */
    public static function fromMap($map = [])
    {
        $model = new self();
        if (isset($map['DbName'])) {
            $model->dbName = $map['DbName'];
        }
        if (isset($map['EntityId'])) {
            $model->entityId = $map['EntityId'];
        }
        if (isset($map['Labels'])) {
            $model->labels = $map['Labels'];
        }

        return $model;
    }
}
