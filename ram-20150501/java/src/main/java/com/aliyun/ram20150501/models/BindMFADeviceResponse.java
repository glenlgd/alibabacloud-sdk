// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.ram20150501.models;

import com.aliyun.tea.*;

public class BindMFADeviceResponse extends TeaModel {
    @NameInMap("RequestId")
    @Validation(required = true)
    public String requestId;

    public static BindMFADeviceResponse build(java.util.Map<String, ?> map) throws Exception {
        BindMFADeviceResponse self = new BindMFADeviceResponse();
        return TeaModel.build(map, self);
    }

}
