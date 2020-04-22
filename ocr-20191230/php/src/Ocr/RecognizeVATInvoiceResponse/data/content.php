<?php

// This file is auto-generated, don't edit it. Thanks.

namespace AlibabaCloud\SDK\Ocr\V20191230\Ocr\RecognizeVATInvoiceResponse\data;

use AlibabaCloud\Tea\Model;

class content extends Model
{
    /**
     * @description invoiceCode
     *
     * @var string
     */
    public $invoiceCode;

    /**
     * @description invoiceNo
     *
     * @var string
     */
    public $invoiceNo;

    /**
     * @description invoiceDate
     *
     * @var string
     */
    public $invoiceDate;

    /**
     * @description antiFakeCode
     *
     * @var string
     */
    public $antiFakeCode;

    /**
     * @description payerName
     *
     * @var string
     */
    public $payerName;

    /**
     * @description payerRegisterNo
     *
     * @var string
     */
    public $payerRegisterNo;

    /**
     * @description payerAddress
     *
     * @var string
     */
    public $payerAddress;

    /**
     * @description payerBankName
     *
     * @var string
     */
    public $payerBankName;

    /**
     * @description withoutTaxAmount
     *
     * @var string
     */
    public $withoutTaxAmount;

    /**
     * @description taxAmount
     *
     * @var string
     */
    public $taxAmount;

    /**
     * @description sumAmount
     *
     * @var string
     */
    public $sumAmount;

    /**
     * @description invoiceAmount
     *
     * @var string
     */
    public $invoiceAmount;

    /**
     * @description payeeName
     *
     * @var string
     */
    public $payeeName;

    /**
     * @description payeeRegisterNo
     *
     * @var string
     */
    public $payeeRegisterNo;

    /**
     * @description payeeAddress
     *
     * @var string
     */
    public $payeeAddress;

    /**
     * @description payeeBankName
     *
     * @var string
     */
    public $payeeBankName;

    /**
     * @description payee
     *
     * @var string
     */
    public $payee;

    /**
     * @description checker
     *
     * @var string
     */
    public $checker;

    /**
     * @description clerk
     *
     * @var string
     */
    public $clerk;
    protected $_name = [
        'invoiceCode'      => 'InvoiceCode',
        'invoiceNo'        => 'InvoiceNo',
        'invoiceDate'      => 'InvoiceDate',
        'antiFakeCode'     => 'AntiFakeCode',
        'payerName'        => 'PayerName',
        'payerRegisterNo'  => 'PayerRegisterNo',
        'payerAddress'     => 'PayerAddress',
        'payerBankName'    => 'PayerBankName',
        'withoutTaxAmount' => 'WithoutTaxAmount',
        'taxAmount'        => 'TaxAmount',
        'sumAmount'        => 'SumAmount',
        'invoiceAmount'    => 'InvoiceAmount',
        'payeeName'        => 'PayeeName',
        'payeeRegisterNo'  => 'PayeeRegisterNo',
        'payeeAddress'     => 'PayeeAddress',
        'payeeBankName'    => 'PayeeBankName',
        'payee'            => 'Payee',
        'checker'          => 'Checker',
        'clerk'            => 'Clerk',
    ];

    public function validate()
    {
        Model::validateRequired('invoiceCode', $this->invoiceCode, true);
        Model::validateRequired('invoiceNo', $this->invoiceNo, true);
        Model::validateRequired('invoiceDate', $this->invoiceDate, true);
        Model::validateRequired('antiFakeCode', $this->antiFakeCode, true);
        Model::validateRequired('payerName', $this->payerName, true);
        Model::validateRequired('payerRegisterNo', $this->payerRegisterNo, true);
        Model::validateRequired('payerAddress', $this->payerAddress, true);
        Model::validateRequired('payerBankName', $this->payerBankName, true);
        Model::validateRequired('withoutTaxAmount', $this->withoutTaxAmount, true);
        Model::validateRequired('taxAmount', $this->taxAmount, true);
        Model::validateRequired('sumAmount', $this->sumAmount, true);
        Model::validateRequired('invoiceAmount', $this->invoiceAmount, true);
        Model::validateRequired('payeeName', $this->payeeName, true);
        Model::validateRequired('payeeRegisterNo', $this->payeeRegisterNo, true);
        Model::validateRequired('payeeAddress', $this->payeeAddress, true);
        Model::validateRequired('payeeBankName', $this->payeeBankName, true);
        Model::validateRequired('payee', $this->payee, true);
        Model::validateRequired('checker', $this->checker, true);
        Model::validateRequired('clerk', $this->clerk, true);
    }

    public function toMap()
    {
        $res                     = [];
        $res['InvoiceCode']      = $this->invoiceCode;
        $res['InvoiceNo']        = $this->invoiceNo;
        $res['InvoiceDate']      = $this->invoiceDate;
        $res['AntiFakeCode']     = $this->antiFakeCode;
        $res['PayerName']        = $this->payerName;
        $res['PayerRegisterNo']  = $this->payerRegisterNo;
        $res['PayerAddress']     = $this->payerAddress;
        $res['PayerBankName']    = $this->payerBankName;
        $res['WithoutTaxAmount'] = $this->withoutTaxAmount;
        $res['TaxAmount']        = $this->taxAmount;
        $res['SumAmount']        = $this->sumAmount;
        $res['InvoiceAmount']    = $this->invoiceAmount;
        $res['PayeeName']        = $this->payeeName;
        $res['PayeeRegisterNo']  = $this->payeeRegisterNo;
        $res['PayeeAddress']     = $this->payeeAddress;
        $res['PayeeBankName']    = $this->payeeBankName;
        $res['Payee']            = $this->payee;
        $res['Checker']          = $this->checker;
        $res['Clerk']            = $this->clerk;

        return $res;
    }

    /**
     * @param array $map
     *
     * @return content
     */
    public static function fromMap($map = [])
    {
        $model = new self();
        if (isset($map['InvoiceCode'])) {
            $model->invoiceCode = $map['InvoiceCode'];
        }
        if (isset($map['InvoiceNo'])) {
            $model->invoiceNo = $map['InvoiceNo'];
        }
        if (isset($map['InvoiceDate'])) {
            $model->invoiceDate = $map['InvoiceDate'];
        }
        if (isset($map['AntiFakeCode'])) {
            $model->antiFakeCode = $map['AntiFakeCode'];
        }
        if (isset($map['PayerName'])) {
            $model->payerName = $map['PayerName'];
        }
        if (isset($map['PayerRegisterNo'])) {
            $model->payerRegisterNo = $map['PayerRegisterNo'];
        }
        if (isset($map['PayerAddress'])) {
            $model->payerAddress = $map['PayerAddress'];
        }
        if (isset($map['PayerBankName'])) {
            $model->payerBankName = $map['PayerBankName'];
        }
        if (isset($map['WithoutTaxAmount'])) {
            $model->withoutTaxAmount = $map['WithoutTaxAmount'];
        }
        if (isset($map['TaxAmount'])) {
            $model->taxAmount = $map['TaxAmount'];
        }
        if (isset($map['SumAmount'])) {
            $model->sumAmount = $map['SumAmount'];
        }
        if (isset($map['InvoiceAmount'])) {
            $model->invoiceAmount = $map['InvoiceAmount'];
        }
        if (isset($map['PayeeName'])) {
            $model->payeeName = $map['PayeeName'];
        }
        if (isset($map['PayeeRegisterNo'])) {
            $model->payeeRegisterNo = $map['PayeeRegisterNo'];
        }
        if (isset($map['PayeeAddress'])) {
            $model->payeeAddress = $map['PayeeAddress'];
        }
        if (isset($map['PayeeBankName'])) {
            $model->payeeBankName = $map['PayeeBankName'];
        }
        if (isset($map['Payee'])) {
            $model->payee = $map['Payee'];
        }
        if (isset($map['Checker'])) {
            $model->checker = $map['Checker'];
        }
        if (isset($map['Clerk'])) {
            $model->clerk = $map['Clerk'];
        }

        return $model;
    }
}
