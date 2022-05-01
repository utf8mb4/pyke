
# Overview

临时文件中转站, 在下载一次后, 一分钟后删除文件

运行镜像
```bash
docker run -d -p 3000:3000 registry.cn-hangzhou.aliyuncs.com/utf8mb4/pyke
```

查看用法
```bash
curl 127.0.0.1:3000
```

上传文件
```bash
curl --upload-file ./example http://127.0.0.1:3000
```