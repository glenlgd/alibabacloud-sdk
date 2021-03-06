// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.pts20190810.models;

import com.aliyun.tea.*;

public class RemoveReportsResponse extends TeaModel {
    @NameInMap("Message")
    @Validation(required = true)
    public String message;

    @NameInMap("Code")
    @Validation(required = true)
    public String code;

    @NameInMap("Success")
    @Validation(required = true)
    public Boolean success;

    @NameInMap("RequestId")
    @Validation(required = true)
    public String requestId;

    @NameInMap("HttpStatusCode")
    @Validation(required = true)
    public Integer httpStatusCode;

    public static RemoveReportsResponse build(java.util.Map<String, ?> map) throws Exception {
        RemoveReportsResponse self = new RemoveReportsResponse();
        return TeaModel.build(map, self);
    }

}
