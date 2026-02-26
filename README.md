# Alist Syncer
支持Alist上多存储间的数据同步

## 配置示例
src_dir和dst_dir均为alist中访问的路径
```json
{
  "endpoint": "http://127.0.0.1:5244",
  "username": "username",
  "password": "password",
  "src_dir": "/src",
  "dst_dir": "/remote"
}
```

## 运行
```bash
go build -o ./alist-syncer .

./alist-syncer -config config.json
```
