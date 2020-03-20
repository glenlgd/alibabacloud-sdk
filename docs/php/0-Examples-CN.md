[← 首页](../../README.md) | 快速使用[(English)](0-Examples-EN.md)

***

## 要求

- 要使用 Alibaba Cloud SDK for PHP，您需要一个云账号以及一对 `Access Key ID` 和 `Access Key Secret`。 请在阿里云控制台中的[AccessKey管理页面](https://usercenter.console.aliyun.com/#/manage/ak)上创建和查看您的Access Key，或者联系您的系统管理员。
- 要使用 Alibaba Cloud SDK for PHP 访问某个产品的API，您需要事先在[阿里云控制台](https://home.console.aliyun.com/)中开通这个产品。
- Alibaba Cloud SDK for PHP 需要1.8以上的JDK。

## 通过 Composer 来管理项目依赖(推荐)

```bash
# require alibabacloud/imagesearch-20200212 for example
composer require alibabacloud/imagesearch-20200212
```

# 快速使用

以下这个代码示例向您展示了调用 Alibaba Cloud SDK for PHP 的3个主要步骤：

1. 创建`Config`实例并初始化。

2. 创建`Client`实例并初始化。

3. 创建`RuntimeOptions`实例并设置运行参数。

4. 创建 API 请求并设置参数。

5. 发起请求并处理应答或异常。

```php
namespace demo;

require __DIR__ . '/vendor/autoload.php';

use AlibabaCloud\SDK\ImageSearch\V20200212\ImageSearch;
use AlibabaCloud\SDK\ImageSearch\V20200212\ImageSearch\Config;
use AlibabaCloud\Tea\Utils\Utils\RuntimeOptions;

$config = new Config([
    'type'            => 'access_key',
    'accessKeyId'     => '<Access-Key-Id>',
    'accessKeySecret' => '<Access-Key-Secret>',
]);

$client = new ImageSearch($config);

$request          = new ImageSearch\SearchImageByNameRequest();
$request->picName = 'test';

$runtime                 = new RuntimeOptions();
$runtime->maxIdleConns   = 3;
$runtime->connectTimeout = 10000;
$runtime->readTimeout    = 10000;

try {
    $response = $client->searchImageByName($request, $runtime);
    var_dump($response);
} catch (TeaError $e) {
    var_dump($e->getErrorInfo());
}
```

***
[← 首页](../../README.md) | 快速使用[(English)](0-Examples-EN.md)