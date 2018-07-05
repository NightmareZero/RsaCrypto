# RsaCrypto

简单的Rsa加密解密小工具


```bash
# 转换生成的密钥格式到java的pkcs8

openssl pkcs8 -topk8 -inform PEM -in private.key -outform pem -nocrypt -out pkcs8.pem

# 生成公钥

openssl rsa -in pkcs8.pem -pubout -out pkcs8_pub.pem

```
