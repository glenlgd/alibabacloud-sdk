<?php

// This file is auto-generated, don't edit it. Thanks.

namespace AlibabaCloud\SDK\Cms\V20190101\Models\CreateMetricRuleTemplateRequest\alertTemplates\escalations;

use AlibabaCloud\Tea\Model;

class critical extends Model
{
    /**
     * @description criticalStatistics
     *
     * @var string
     */
    public $statistics;

    /**
     * @description criticalComparisonOperator
     *
     * @var string
     */
    public $comparisonOperator;

    /**
     * @description criticalThreshold
     *
     * @var string
     */
    public $threshold;

    /**
     * @description criticalTimes
     *
     * @var int
     */
    public $times;
    protected $_name = [
        'statistics'         => 'Statistics',
        'comparisonOperator' => 'ComparisonOperator',
        'threshold'          => 'Threshold',
        'times'              => 'Times',
    ];

    public function validate()
    {
        Model::validateRequired('statistics', $this->statistics, true);
        Model::validateRequired('comparisonOperator', $this->comparisonOperator, true);
        Model::validateRequired('threshold', $this->threshold, true);
        Model::validateRequired('times', $this->times, true);
    }

    public function toMap()
    {
        $res = [];
        if (null !== $this->statistics) {
            $res['Statistics'] = $this->statistics;
        }
        if (null !== $this->comparisonOperator) {
            $res['ComparisonOperator'] = $this->comparisonOperator;
        }
        if (null !== $this->threshold) {
            $res['Threshold'] = $this->threshold;
        }
        if (null !== $this->times) {
            $res['Times'] = $this->times;
        }

        return $res;
    }

    /**
     * @param array $map
     *
     * @return critical
     */
    public static function fromMap($map = [])
    {
        $model = new self();
        if (isset($map['Statistics'])) {
            $model->statistics = $map['Statistics'];
        }
        if (isset($map['ComparisonOperator'])) {
            $model->comparisonOperator = $map['ComparisonOperator'];
        }
        if (isset($map['Threshold'])) {
            $model->threshold = $map['Threshold'];
        }
        if (isset($map['Times'])) {
            $model->times = $map['Times'];
        }

        return $model;
    }
}
