<?php

// This file is auto-generated, don't edit it. Thanks.

namespace AlibabaCloud\SDK\ECI\V20180808\ECI;

use AlibabaCloud\Tea\Model;

class ExecContainerCommandRequest extends Model
{
    public $ownerId;

    public $resourceOwnerAccount;

    public $resourceOwnerId;

    public $ownerAccount;

    public $regionId;

    public $containerGroupId;

    public $containerName;

    public $command;

    public $TTY;
    protected $_name = [
        'ownerId'              => 'OwnerId',
        'resourceOwnerAccount' => 'ResourceOwnerAccount',
        'resourceOwnerId'      => 'ResourceOwnerId',
        'ownerAccount'         => 'OwnerAccount',
        'regionId'             => 'RegionId',
        'containerGroupId'     => 'ContainerGroupId',
        'containerName'        => 'ContainerName',
        'command'              => 'Command',
        'TTY'                  => 'TTY',
    ];
}