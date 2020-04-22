<?php

// This file is auto-generated, don't edit it. Thanks.

namespace AlibabaCloud\SDK\Imageenhan\V20190930\Imageenhan;

use AlibabaCloud\SDK\Imageenhan\V20190930\Imageenhan\ExtendImageStyleResponse\data;
use AlibabaCloud\Tea\Model;

class ExtendImageStyleResponse extends Model
{
    /**
     * @description requestId
     *
     * @var string
     */
    public $requestId;

    /**
     * @description data
     *
     * @var ExtendImageStyleResponse.data
     */
    public $data;
    protected $_name = [
        'requestId' => 'RequestId',
        'data'      => 'Data',
    ];

    public function validate()
    {
        Model::validateRequired('requestId', $this->requestId, true);
        Model::validateRequired('data', $this->data, true);
    }

    public function toMap()
    {
        $res              = [];
        $res['RequestId'] = $this->requestId;
        $res['Data']      = null !== $this->data ? $this->data->toMap() : null;

        return $res;
    }

    /**
     * @param array $map
     *
     * @return ExtendImageStyleResponse
     */
    public static function fromMap($map = [])
    {
        $model = new self();
        if (isset($map['RequestId'])) {
            $model->requestId = $map['RequestId'];
        }
        if (isset($map['Data'])) {
            $model->data = ExtendImageStyleResponse\data::fromMap($map['Data']);
        }

        return $model;
    }
}
