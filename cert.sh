#!/bin/bash

# 证书信息
COUNTRY="CN"
STATE="Guangdong"
CITY="Shenzhen"
ORGANIZATION="My Company"
ORGANIZATIONAL_UNIT="IT"
COMMON_NAME="www.example.com"
EMAIL="admin@example.com"

# 文件名
KEY_FILE="server.key"
CSR_FILE="server.csr"
CERT_FILE="server.crt"

# 生成私钥
openssl genrsa -out $KEY_FILE 2048

# 生成证书签名请求 (CSR)
openssl req -new -key $KEY_FILE -out $CSR_FILE -subj "/C=$COUNTRY/ST=$STATE/L=$CITY/O=$ORGANIZATION/OU=$ORGANIZATIONAL_UNIT/CN=$COMMON_NAME/emailAddress=$EMAIL"

# 生成自签名证书
openssl x509 -req -days 365 -in $CSR_FILE -signkey $KEY_FILE -out $CERT_FILE

# 清理不再需要的 CSR 文件
rm $CSR_FILE

echo "自签名证书和私钥已生成："
echo "私钥文件: $KEY_FILE"
echo "证书文件: $CERT_FILE"
