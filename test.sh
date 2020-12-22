#!/usr/bin/env bash

# 上传图片
curl --location --request POST 'http://127.0.0.1:8080/file/upload' \
    --form 'up_file=@"/Users/sybs/aaa.png"' \
    --form 'remark="测试一下"'

# 获取图片
curl --location --request GET 'http://127.0.0.1:8080/file/c3b5e206-8821-45de-81be-309b64be5eaf'
