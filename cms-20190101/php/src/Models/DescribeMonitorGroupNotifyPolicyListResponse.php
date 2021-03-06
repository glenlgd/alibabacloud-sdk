<?php

// This file is auto-generated, don't edit it. Thanks.

namespace AlibabaCloud\SDK\Cms\V20190101\Models;

use AlibabaCloud\SDK\Cms\V20190101\Models\DescribeMonitorGroupNotifyPolicyListResponse\notifyPolicyList;
use AlibabaCloud\Tea\Model;

class DescribeMonitorGroupNotifyPolicyListResponse extends Model
{
    /**
     * @description code
     *
     * @var string
     */
    public $code;

    /**
     * @description msg
     *
     * @var string
     */
    public $message;

    /**
     * @description success
     *
     * @var string
     */
    public $success;

    /**
     * @description traceId
     *
     * @var string
     */
    public $requestId;

    /**
     * @description total
     *
     * @var int
     */
    public $total;

    /**
     * @description datapoints
     *
     * @var notifyPolicyList
     */
    public $notifyPolicyList;
    protected $_name = [
        'code'             => 'Code',
        'message'          => 'Message',
        'success'          => 'Success',
        'requestId'        => 'RequestId',
        'total'            => 'Total',
        'notifyPolicyList' => 'NotifyPolicyList',
    ];

    public function validate()
    {
        Model::validateRequired('code', $this->code, true);
        Model::validateRequired('message', $this->message, true);
        Model::validateRequired('success', $this->success, true);
        Model::validateRequired('requestId', $this->requestId, true);
        Model::validateRequired('total', $this->total, true);
        Model::validateRequired('notifyPolicyList', $this->notifyPolicyList, true);
    }

    public function toMap()
    {
        $res = [];
        if (null !== $this->code) {
            $res['Code'] = $this->code;
        }
        if (null !== $this->message) {
            $res['Message'] = $this->message;
        }
        if (null !== $this->success) {
            $res['Success'] = $this->success;
        }
        if (null !== $this->requestId) {
            $res['RequestId'] = $this->requestId;
        }
        if (null !== $this->total) {
            $res['Total'] = $this->total;
        }
        if (null !== $this->notifyPolicyList) {
            $res['NotifyPolicyList'] = null !== $this->notifyPolicyList ? $this->notifyPolicyList->toMap() : null;
        }

        return $res;
    }

    /**
     * @param array $map
     *
     * @return DescribeMonitorGroupNotifyPolicyListResponse
     */
    public static function fromMap($map = [])
    {
        $model = new self();
        if (isset($map['Code'])) {
            $model->code = $map['Code'];
        }
        if (isset($map['Message'])) {
            $model->message = $map['Message'];
        }
        if (isset($map['Success'])) {
            $model->success = $map['Success'];
        }
        if (isset($map['RequestId'])) {
            $model->requestId = $map['RequestId'];
        }
        if (isset($map['Total'])) {
            $model->total = $map['Total'];
        }
        if (isset($map['NotifyPolicyList'])) {
            $model->notifyPolicyList = notifyPolicyList::fromMap($map['NotifyPolicyList']);
        }

        return $model;
    }
}
