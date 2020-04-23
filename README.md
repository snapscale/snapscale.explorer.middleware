# SnapScale-explorer-middleware

SnapScale-explorer-middleware is an snapscale-explorer gateway.

This software provides more statistical dimensions that eos api cannot provide.

It's a part of snapscale-explorer, such like a gateway and background.If you want to start a snapscale-explorer,you also need to start a snapscale-explorer-middleware.

## Disclaimer

SnapScale-explorer-middleware is neither launching nor operating any initial public blockchains based upon the SnapScale software. This release refers only to version 1.0 of our open source software. We caution those who wish to use blockchains built on SnapScale to carefully vet the companies and organizations launching blockchains based on SnapScale before disclosing any private keys to their derivative software.

## Before Everything

You need to start a snapscale node,and make sure snapscale using mongo save history copy.
Then you need to create collection and index.

```sh
db.createCollection("daily");
db.transactions.createIndex({"createdAt":1},{"background":true});
db.daily.ensureIndex({xid:1})
```

**Note: If your database is verybig,createIndex will cost a long time,just wait**

## Build

```sh
./build build
```

## Configuration

1. SNAPSCALE_EXPLORER_MID_API_BASE:       CHAIN API ENDPOINT
2. SNAPSCALE_EXPLORER_MID_HTTP_PORT:      Middle-ware http service port
3. SNAPSCALE_EXPLORER_MID_WS_PORT:        Middle-ware websocket service port
4. SNAPSCALE_EXPLORER_MID_MONGO_CONFIG:   Mongo address and certification
5. SNAPSCALE_EXPLORER_MID_LOG_PATH:       Middle-ware log path

```sh
docker run -itd \
  -p 8089:8089 \
  -p 8089:8089/udp \
  -p 8090:8090 \
  -e SNAPSCALE_EXPLORER_MID_API_BASE='http://192.168.1.201:30132/v1/'
  -e SNAPSCALE_EXPLORER_MID_HTTP_PORT='8090'
  -e SNAPSCALE_EXPLORER_MID_WS_PORT='8089'
  -e SNAPSCALE_EXPLORER_MID_MONGO_CONFIG='mongodb://username:password@192.168.1.1:30017/?authSource=admin'
  -e SNAPSCALE_EXPLORER_MID_LOG_PATH='/log'
  --name $NAME $NAME
```

## License

SnapScale is released under the open source [MIT](./LICENSE) license and is offered “AS IS” without warranty of any kind, express or implied. Any security provided by the SnapScale software depends in part on how it is used, configured, and deployed. SnapScale is built upon many third-party libraries such as WABT (Apache License) and WAVM (BSD 3-clause) which are also provided “AS IS” without warranty of any kind. Without limiting the generality of the foregoing, Block.one makes no representation or guarantee that SnapScale or any third-party libraries will perform as intended or will be free of errors, bugs or faulty code. Both may fail in large or small ways that could completely or partially limit functionality or compromise computer systems. If you use or implement SnapScale, you do so at your own risk. In no event will Block.one be liable to any party for any damages whatsoever, even if it had been advised of the possibility of damage.  

## Important

See [LICENSE](./LICENSE) for copyright and license terms.

All repositories and other materials are provided subject to the terms of this [IMPORTANT](./IMPORTANT.md) notice and you must familiarize yourself with its terms.  The notice contains important information, limitations and restrictions relating to our software, publications, trademarks, third-party resources, and forward-looking statements.  By accessing any of our repositories and other materials, you accept and agree to the terms of the notice.