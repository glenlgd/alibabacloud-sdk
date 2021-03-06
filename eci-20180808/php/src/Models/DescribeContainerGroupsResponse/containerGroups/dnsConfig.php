<?php

// This file is auto-generated, don't edit it. Thanks.

namespace AlibabaCloud\SDK\Eci\V20180808\Models\DescribeContainerGroupsResponse\containerGroups;

use AlibabaCloud\SDK\Eci\V20180808\Models\DescribeContainerGroupsResponse\containerGroups\dnsConfig\options;
use AlibabaCloud\Tea\Model;

class dnsConfig extends Model
{
    /**
     * @description dnsConfigOptions
     *
     * @var array
     */
    public $options;

    /**
     * @description dnsConfigNameServers
     *
     * @var array
     */
    public $nameServers;

    /**
     * @description dnsConfigSearches
     *
     * @var array
     */
    public $searches;
    protected $_name = [
        'options'     => 'Options',
        'nameServers' => 'NameServers',
        'searches'    => 'Searches',
    ];

    public function validate()
    {
        Model::validateRequired('options', $this->options, true);
        Model::validateRequired('nameServers', $this->nameServers, true);
        Model::validateRequired('searches', $this->searches, true);
    }

    public function toMap()
    {
        $res = [];
        if (null !== $this->options) {
            $res['Options'] = [];
            if (null !== $this->options && \is_array($this->options)) {
                $n = 0;
                foreach ($this->options as $item) {
                    $res['Options'][$n++] = null !== $item ? $item->toMap() : $item;
                }
            }
        }
        if (null !== $this->nameServers) {
            $res['NameServers'] = [];
            if (null !== $this->nameServers) {
                $res['NameServers'] = $this->nameServers;
            }
        }
        if (null !== $this->searches) {
            $res['Searches'] = [];
            if (null !== $this->searches) {
                $res['Searches'] = $this->searches;
            }
        }

        return $res;
    }

    /**
     * @param array $map
     *
     * @return dnsConfig
     */
    public static function fromMap($map = [])
    {
        $model = new self();
        if (isset($map['Options'])) {
            if (!empty($map['Options'])) {
                $model->options = [];
                $n              = 0;
                foreach ($map['Options'] as $item) {
                    $model->options[$n++] = null !== $item ? options::fromMap($item) : $item;
                }
            }
        }
        if (isset($map['NameServers'])) {
            if (!empty($map['NameServers'])) {
                $model->nameServers = [];
                $model->nameServers = $map['NameServers'];
            }
        }
        if (isset($map['Searches'])) {
            if (!empty($map['Searches'])) {
                $model->searches = [];
                $model->searches = $map['Searches'];
            }
        }

        return $model;
    }
}
