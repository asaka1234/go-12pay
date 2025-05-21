Deposit
==============
1. 生成一个支付qrcode二维码
2. 用户扫描支付
3. 回调通知merchant (回调接口是merchant商户申请时即指定的)

Withdraw
==============
用户发起即可, 银行收到后会直接执行

Comment
===============
both support deposit && withdrawl


鉴权
==============
请求时在header里放了partner_code 和 authorization . 通过这个进行鉴权 (并没有对请求body做签名)


回调地址
==============
是提前让12pay设置好的.（无法api中动态修改）