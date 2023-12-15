## 檔案功能

```
├─ server.go  : 主程式
├─ upgrade.go : 升級連線
├─ frame.go   : 按照 RFC6455 格式 encode/decode
├─ client.go  : 封裝過後的 websocket 連線
├─ httpy.go   : 處理 http 封包內容的小函式
```

## 執行

server

```
go run .
```

client

```
打開網頁 index.html
```
