// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.imagerecog.models;

import com.aliyun.tea.*;

public class TaggingImageAdvanceRequest extends TeaModel {
    @NameInMap("ImageURLObject")
    @Validation(required = true)
    public java.io.InputStream imageURLObject;

    @NameInMap("ImageType")
    public Integer imageType;

    public static TaggingImageAdvanceRequest build(java.util.Map<String, ?> map) throws Exception {
        TaggingImageAdvanceRequest self = new TaggingImageAdvanceRequest();
        return TeaModel.build(map, self);
    }

}