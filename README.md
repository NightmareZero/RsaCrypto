# RsaCrypto

简单的Rsa加密解密小工具


转换生成的密钥格式到java的pkcs8
```bash
openssl pkcs8 -topk8 -inform PEM -in private.key -outform pem -nocrypt -out pkcs8.pem
```
