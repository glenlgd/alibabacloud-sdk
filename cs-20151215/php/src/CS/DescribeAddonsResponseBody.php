<?php

// This file is auto-generated, don't edit it. Thanks.

namespace AlibabaCloud\SDK\CS\V20151215\CS;

use AlibabaCloud\SDK\CS\V20151215\CS\DescribeAddonsResponseBody\standardComponents;
use AlibabaCloud\Tea\Model;

class DescribeAddonsResponseBody extends Model
{
    /**
     * @description ComponentGroups
     *
     * @var array
     */
    public $componentGroups;

    /**
     * @description StandardComponents
     *
     * @var DescribeAddonsResponseBody.standardComponents
     */
    public $standardComponents;
    protected $_name = [
        'componentGroups'    => 'ComponentGroups',
        'standardComponents' => 'StandardComponents',
    ];

    public function validate()
    {
        Model::validateRequired('componentGroups', $this->componentGroups, true);
        Model::validateRequired('standardComponents', $this->standardComponents, true);
    }

    public function toMap()
    {
        $res                    = [];
        $res['ComponentGroups'] = [];
        if (null !== $this->componentGroups && \is_array($this->componentGroups)) {
            $n = 0;
            foreach ($this->componentGroups as $item) {
                $res['ComponentGroups'][$n++] = null !== $item ? $item->toMap() : $item;
            }
        }
        $res['StandardComponents'] = null !== $this->standardComponents ? $this->standardComponents->toMap() : null;

        return $res;
    }

    /**
     * @param array $map
     *
     * @return DescribeAddonsResponseBody
     */
    public static function fromMap($map = [])
    {
        $model = new self();
        if (isset($map['ComponentGroups'])) {
            if (!empty($map['ComponentGroups'])) {
                $model->componentGroups = [];
                $n                      = 0;
                foreach ($map['ComponentGroups'] as $item) {
                    $model->componentGroups[$n++] = null !== $item ? DescribeAddonsResponseBody\componentGroups::fromMap($item) : $item;
                }
            }
        }
        if (isset($map['StandardComponents'])) {
            $model->standardComponents = DescribeAddonsResponseBody\standardComponents::fromMap($map['StandardComponents']);
        }

        return $model;
    }
}
