// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.vod20170321.models;

import com.aliyun.tea.*;

public class AddVodTemplateResponse extends TeaModel {
    @NameInMap("RequestId")
    @Validation(required = true)
    public String requestId;

    @NameInMap("VodTemplateId")
    @Validation(required = true)
    public String vodTemplateId;

    public static AddVodTemplateResponse build(java.util.Map<String, ?> map) throws Exception {
        AddVodTemplateResponse self = new AddVodTemplateResponse();
        return TeaModel.build(map, self);
    }

}
