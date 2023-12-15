1. websocket 原理及簡單實作

   1.1. http upgrade request

   1.2. socket connection with 6455 format

   1.3. ping-pong / readystate

   1.4. 簡單實作: upgrader / frame-codec

2. golang websocket 的選擇

   3.1. Fork 3.5K / Star 20.3k: [gorilla/websocket](https://github.com/gorilla/websocket)

   3.2. Fork 866 / Star 5.4k: [go-socketio](https://github.com/googollee/go-socket.io)

   3.4. Fork 13 / Star 104: [socket.io](https://github.com/zishang520/socket.io)

   3.3. Fork 386 / Star 3.4k:[melody](https://github.com/olahol/melody)

3. 開始做聊天室

   3.1. 核心功能: chat / room / history_message / users / 已讀 / @

   3.2. 技術選擇: gorilla / gin / mysql

   3.3. gorilla 跟 gin 的互動

   3.4. database schema

   3.5. pkg/caht, pkg/client, store

   3.6. handlers

   3.7. web ui 的部份

   3.8. demo

   3.9. 壓測 / profile

4. 其他相關分享:

   4.1. socketio / engineio

   4.2. websocket testsuit: [autobahn-testsuite](https://github.com/crossbario/autobahn-testsuite)

   4.3. 底層優化的 ws 實作: [gobwas/ws](https://github.com/gobwas/ws)

   4.4. bilibili 的彈幕架構 [bilibili/goim](https://github.com/Terry-Mao/goim)

5. references:
   working on
