// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.polardb20170801.models;

import com.aliyun.tea.*;

public class CreateAccountResponse extends TeaModel {
    @NameInMap("RequestId")
    @Validation(required = true)
    public String requestId;

    public static CreateAccountResponse build(java.util.Map<String, ?> map) throws Exception {
        CreateAccountResponse self = new CreateAccountResponse();
        return TeaModel.build(map, self);
    }

}
