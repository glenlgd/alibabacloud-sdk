// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.opensearch20171225.models;

import com.aliyun.tea.*;

public class DescribeAppGroupResponseBody extends TeaModel {
    @NameInMap("requestId")
    @Validation(required = true)
    public String requestId;

    @NameInMap("result")
    @Validation(required = true)
    public DescribeAppGroupResponseBodyResult result;

    public static DescribeAppGroupResponseBody build(java.util.Map<String, ?> map) throws Exception {
        DescribeAppGroupResponseBody self = new DescribeAppGroupResponseBody();
        return TeaModel.build(map, self);
    }

    public static class DescribeAppGroupResponseBodyResultQuota extends TeaModel {
        @NameInMap("docSize")
        @Validation(required = true)
        public Integer docSize;

        @NameInMap("computeResource")
        @Validation(required = true)
        public Integer computeResource;

        @NameInMap("spec")
        @Validation(required = true)
        public String spec;

        public static DescribeAppGroupResponseBodyResultQuota build(java.util.Map<String, ?> map) throws Exception {
            DescribeAppGroupResponseBodyResultQuota self = new DescribeAppGroupResponseBodyResultQuota();
            return TeaModel.build(map, self);
        }

    }

    public static class DescribeAppGroupResponseBodyResult extends TeaModel {
        @NameInMap("id")
        @Validation(required = true)
        public String id;

        @NameInMap("name")
        @Validation(required = true)
        public String name;

        @NameInMap("currentVersion")
        @Validation(required = true)
        public String currentVersion;

        @NameInMap("switchedTime")
        @Validation(required = true)
        public Integer switchedTime;

        @NameInMap("chargingWay")
        @Validation(required = true)
        public Integer chargingWay;

        @NameInMap("type")
        @Validation(required = true)
        public String type;

        @NameInMap("projectId")
        @Validation(required = true)
        public String projectId;

        @NameInMap("chargeType")
        @Validation(required = true)
        public String chargeType;

        @NameInMap("expireOn")
        @Validation(required = true)
        public String expireOn;

        @NameInMap("instanceId")
        @Validation(required = true)
        public String instanceId;

        @NameInMap("commodityCode")
        @Validation(required = true)
        public String commodityCode;

        @NameInMap("processingOrderId")
        @Validation(required = true)
        public String processingOrderId;

        @NameInMap("firstRankAlgoDeploymentId")
        @Validation(required = true)
        public Integer firstRankAlgoDeploymentId;

        @NameInMap("secondRankAlgoDeploymentId")
        @Validation(required = true)
        public Integer secondRankAlgoDeploymentId;

        @NameInMap("pendingSecondRankAlgoDeploymentId")
        @Validation(required = true)
        public Integer pendingSecondRankAlgoDeploymentId;

        @NameInMap("description")
        @Validation(required = true)
        public String description;

        @NameInMap("produced")
        @Validation(required = true)
        public Integer produced;

        @NameInMap("lockedByExpiration")
        @Validation(required = true)
        public Integer lockedByExpiration;

        @NameInMap("hasPendingQuotaReviewTask")
        @Validation(required = true)
        public Integer hasPendingQuotaReviewTask;

        @NameInMap("created")
        @Validation(required = true)
        public Integer created;

        @NameInMap("updated")
        @Validation(required = true)
        public Integer updated;

        @NameInMap("status")
        @Validation(required = true)
        public String status;

        @NameInMap("lockMode")
        @Validation(required = true)
        public String lockMode;

        @NameInMap("quota")
        @Validation(required = true)
        public DescribeAppGroupResponseBodyResultQuota quota;

        public static DescribeAppGroupResponseBodyResult build(java.util.Map<String, ?> map) throws Exception {
            DescribeAppGroupResponseBodyResult self = new DescribeAppGroupResponseBodyResult();
            return TeaModel.build(map, self);
        }

    }

}
