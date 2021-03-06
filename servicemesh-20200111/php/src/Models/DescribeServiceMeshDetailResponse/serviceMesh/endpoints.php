<?php

// This file is auto-generated, don't edit it. Thanks.

namespace AlibabaCloud\SDK\Servicemesh\V20200111\Models\DescribeServiceMeshDetailResponse\serviceMesh;

use AlibabaCloud\Tea\Model;

class endpoints extends Model
{
    /**
     * @description intranet_api_server_endpoint
     *
     * @var string
     */
    public $intranetApiServerEndpoint;

    /**
     * @description intranet_pilot_endpoint
     *
     * @var string
     */
    public $intranetPilotEndpoint;

    /**
     * @description public_api_server_endpoint
     *
     * @var string
     */
    public $publicApiServerEndpoint;

    /**
     * @description public_pilot_endpoint
     *
     * @var string
     */
    public $publicPilotEndpoint;
    protected $_name = [
        'intranetApiServerEndpoint' => 'IntranetApiServerEndpoint',
        'intranetPilotEndpoint'     => 'IntranetPilotEndpoint',
        'publicApiServerEndpoint'   => 'PublicApiServerEndpoint',
        'publicPilotEndpoint'       => 'PublicPilotEndpoint',
    ];

    public function validate()
    {
        Model::validateRequired('intranetApiServerEndpoint', $this->intranetApiServerEndpoint, true);
        Model::validateRequired('intranetPilotEndpoint', $this->intranetPilotEndpoint, true);
        Model::validateRequired('publicApiServerEndpoint', $this->publicApiServerEndpoint, true);
        Model::validateRequired('publicPilotEndpoint', $this->publicPilotEndpoint, true);
    }

    public function toMap()
    {
        $res = [];
        if (null !== $this->intranetApiServerEndpoint) {
            $res['IntranetApiServerEndpoint'] = $this->intranetApiServerEndpoint;
        }
        if (null !== $this->intranetPilotEndpoint) {
            $res['IntranetPilotEndpoint'] = $this->intranetPilotEndpoint;
        }
        if (null !== $this->publicApiServerEndpoint) {
            $res['PublicApiServerEndpoint'] = $this->publicApiServerEndpoint;
        }
        if (null !== $this->publicPilotEndpoint) {
            $res['PublicPilotEndpoint'] = $this->publicPilotEndpoint;
        }

        return $res;
    }

    /**
     * @param array $map
     *
     * @return endpoints
     */
    public static function fromMap($map = [])
    {
        $model = new self();
        if (isset($map['IntranetApiServerEndpoint'])) {
            $model->intranetApiServerEndpoint = $map['IntranetApiServerEndpoint'];
        }
        if (isset($map['IntranetPilotEndpoint'])) {
            $model->intranetPilotEndpoint = $map['IntranetPilotEndpoint'];
        }
        if (isset($map['PublicApiServerEndpoint'])) {
            $model->publicApiServerEndpoint = $map['PublicApiServerEndpoint'];
        }
        if (isset($map['PublicPilotEndpoint'])) {
            $model->publicPilotEndpoint = $map['PublicPilotEndpoint'];
        }

        return $model;
    }
}
