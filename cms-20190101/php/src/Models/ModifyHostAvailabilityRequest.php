<?php

// This file is auto-generated, don't edit it. Thanks.

namespace AlibabaCloud\SDK\Cms\V20190101\Models;

use AlibabaCloud\SDK\Cms\V20190101\Models\ModifyHostAvailabilityRequest\alertConfig;
use AlibabaCloud\SDK\Cms\V20190101\Models\ModifyHostAvailabilityRequest\alertConfigEscalationList;
use AlibabaCloud\SDK\Cms\V20190101\Models\ModifyHostAvailabilityRequest\taskOption;
use AlibabaCloud\Tea\Model;

class ModifyHostAvailabilityRequest extends Model
{
    /**
     * @description groupId
     *
     * @var int
     */
    public $groupId;

    /**
     * @description id
     *
     * @var int
     */
    public $id;

    /**
     * @description taskName
     *
     * @var string
     */
    public $taskName;

    /**
     * @description taskScope
     *
     * @var string
     */
    public $taskScope;

    /**
     * @description TaskOption
     *
     * @var taskOption
     */
    public $taskOption;

    /**
     * @description AlertConfig
     *
     * @var alertConfig
     */
    public $alertConfig;

    /**
     * @description escalationListJson
     *
     * @var array
     */
    public $alertConfigEscalationList;

    /**
     * @description instanceListJson
     *
     * @var array
     */
    public $instanceList;
    protected $_name = [
        'groupId'                   => 'GroupId',
        'id'                        => 'Id',
        'taskName'                  => 'TaskName',
        'taskScope'                 => 'TaskScope',
        'taskOption'                => 'TaskOption',
        'alertConfig'               => 'AlertConfig',
        'alertConfigEscalationList' => 'AlertConfigEscalationList',
        'instanceList'              => 'InstanceList',
    ];

    public function validate()
    {
        Model::validateRequired('groupId', $this->groupId, true);
        Model::validateRequired('id', $this->id, true);
        Model::validateRequired('taskName', $this->taskName, true);
        Model::validateRequired('taskOption', $this->taskOption, true);
        Model::validateRequired('alertConfig', $this->alertConfig, true);
        Model::validateRequired('alertConfigEscalationList', $this->alertConfigEscalationList, true);
    }

    public function toMap()
    {
        $res = [];
        if (null !== $this->groupId) {
            $res['GroupId'] = $this->groupId;
        }
        if (null !== $this->id) {
            $res['Id'] = $this->id;
        }
        if (null !== $this->taskName) {
            $res['TaskName'] = $this->taskName;
        }
        if (null !== $this->taskScope) {
            $res['TaskScope'] = $this->taskScope;
        }
        if (null !== $this->taskOption) {
            $res['TaskOption'] = null !== $this->taskOption ? $this->taskOption->toMap() : null;
        }
        if (null !== $this->alertConfig) {
            $res['AlertConfig'] = null !== $this->alertConfig ? $this->alertConfig->toMap() : null;
        }
        if (null !== $this->alertConfigEscalationList) {
            $res['AlertConfigEscalationList'] = [];
            if (null !== $this->alertConfigEscalationList && \is_array($this->alertConfigEscalationList)) {
                $n = 0;
                foreach ($this->alertConfigEscalationList as $item) {
                    $res['AlertConfigEscalationList'][$n++] = null !== $item ? $item->toMap() : $item;
                }
            }
        }
        if (null !== $this->instanceList) {
            $res['InstanceList'] = [];
            if (null !== $this->instanceList) {
                $res['InstanceList'] = $this->instanceList;
            }
        }

        return $res;
    }

    /**
     * @param array $map
     *
     * @return ModifyHostAvailabilityRequest
     */
    public static function fromMap($map = [])
    {
        $model = new self();
        if (isset($map['GroupId'])) {
            $model->groupId = $map['GroupId'];
        }
        if (isset($map['Id'])) {
            $model->id = $map['Id'];
        }
        if (isset($map['TaskName'])) {
            $model->taskName = $map['TaskName'];
        }
        if (isset($map['TaskScope'])) {
            $model->taskScope = $map['TaskScope'];
        }
        if (isset($map['TaskOption'])) {
            $model->taskOption = taskOption::fromMap($map['TaskOption']);
        }
        if (isset($map['AlertConfig'])) {
            $model->alertConfig = alertConfig::fromMap($map['AlertConfig']);
        }
        if (isset($map['AlertConfigEscalationList'])) {
            if (!empty($map['AlertConfigEscalationList'])) {
                $model->alertConfigEscalationList = [];
                $n                                = 0;
                foreach ($map['AlertConfigEscalationList'] as $item) {
                    $model->alertConfigEscalationList[$n++] = null !== $item ? alertConfigEscalationList::fromMap($item) : $item;
                }
            }
        }
        if (isset($map['InstanceList'])) {
            if (!empty($map['InstanceList'])) {
                $model->instanceList = [];
                $model->instanceList = $map['InstanceList'];
            }
        }

        return $model;
    }
}
