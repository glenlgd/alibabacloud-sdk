<?php

// This file is auto-generated, don't edit it. Thanks.

namespace AlibabaCloud\SDK\Objectdet\V20191230\Models\DetectObjectResponse;

use AlibabaCloud\SDK\Objectdet\V20191230\Models\DetectObjectResponse\data\elements;
use AlibabaCloud\Tea\Model;

class data extends Model
{
    /**
     * @description width
     *
     * @var int
     */
    public $width;

    /**
     * @description height
     *
     * @var int
     */
    public $height;

    /**
     * @description elements
     *
     * @var array
     */
    public $elements;
    protected $_name = [
        'width'    => 'Width',
        'height'   => 'Height',
        'elements' => 'Elements',
    ];

    public function validate()
    {
        Model::validateRequired('width', $this->width, true);
        Model::validateRequired('height', $this->height, true);
        Model::validateRequired('elements', $this->elements, true);
    }

    public function toMap()
    {
        $res = [];
        if (null !== $this->width) {
            $res['Width'] = $this->width;
        }
        if (null !== $this->height) {
            $res['Height'] = $this->height;
        }
        if (null !== $this->elements) {
            $res['Elements'] = [];
            if (null !== $this->elements && \is_array($this->elements)) {
                $n = 0;
                foreach ($this->elements as $item) {
                    $res['Elements'][$n++] = null !== $item ? $item->toMap() : $item;
                }
            }
        }

        return $res;
    }

    /**
     * @param array $map
     *
     * @return data
     */
    public static function fromMap($map = [])
    {
        $model = new self();
        if (isset($map['Width'])) {
            $model->width = $map['Width'];
        }
        if (isset($map['Height'])) {
            $model->height = $map['Height'];
        }
        if (isset($map['Elements'])) {
            if (!empty($map['Elements'])) {
                $model->elements = [];
                $n               = 0;
                foreach ($map['Elements'] as $item) {
                    $model->elements[$n++] = null !== $item ? elements::fromMap($item) : $item;
                }
            }
        }

        return $model;
    }
}
