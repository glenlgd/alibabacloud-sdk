// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.vod20170321.models;

import com.aliyun.tea.*;

public class AttachAppPolicyToIdentityRequest extends TeaModel {
    @NameInMap("OwnerId")
    public Long ownerId;

    @NameInMap("ResourceOwnerAccount")
    public String resourceOwnerAccount;

    @NameInMap("ResourceOwnerId")
    public Long resourceOwnerId;

    @NameInMap("IdentityType")
    @Validation(required = true)
    public String identityType;

    @NameInMap("IdentityName")
    @Validation(required = true)
    public String identityName;

    @NameInMap("AppId")
    public String appId;

    @NameInMap("PolicyNames")
    @Validation(required = true)
    public String policyNames;

    public static AttachAppPolicyToIdentityRequest build(java.util.Map<String, ?> map) throws Exception {
        AttachAppPolicyToIdentityRequest self = new AttachAppPolicyToIdentityRequest();
        return TeaModel.build(map, self);
    }

}
