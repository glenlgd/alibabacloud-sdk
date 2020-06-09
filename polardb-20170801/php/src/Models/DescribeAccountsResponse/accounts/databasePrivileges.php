<?php

// This file is auto-generated, don't edit it. Thanks.
namespace AlibabaCloud\SDK\Polardb\V20170801\Models\DescribeAccountsResponse\accounts;

use AlibabaCloud\Tea\Model;

class databasePrivileges extends Model {
    protected $_name = [
        'DBName' => 'DBName',
        'accountPrivilege' => 'AccountPrivilege',
    ];
    public function validate() {
        Model::validateRequired('DBName', $this->DBName, true);
        Model::validateRequired('accountPrivilege', $this->accountPrivilege, true);
    }
    public function toMap() {
        $res = [];
        $res['DBName'] = $this->DBName;
        $res['AccountPrivilege'] = $this->accountPrivilege;
        return $res;
    }
    /**
     * @param array $map
     * @return databasePrivileges
     */
    public static function fromMap($map = []) {
        $model = new self();
        if(isset($map['DBName'])){
            $model->DBName = $map['DBName'];
        }
        if(isset($map['AccountPrivilege'])){
            $model->accountPrivilege = $map['AccountPrivilege'];
        }
        return $model;
    }
    /**
     * @description dbName
     * @var string
     */
    public $DBName;

    /**
     * @description accountPrivilege
     * @var string
     */
    public $accountPrivilege;

}
