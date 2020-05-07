# SnapScale Explorer Middleware

SnapScale Explorer Middleware is a Snapscale Explorer gateway.

This software provides substantial statistical data and dimensions compared to EOS API.

It is part of Snapscale Explorer, for which it serves as a gateway. Deploying SnapScale Explorer Middleware is mandatory for anyone looking to run SnapScale Explorer.

## Prerequisites

Start a SnapScale node and make sure that SnapScale is using a copy of Mongo save history. Create collection and index.

```sh
db.createCollection("daily");
db.transactions.createIndex({"createdAt":1},{"background":true});
db.daily.ensureIndex({xid:1})
```

**Note: If your database is very large, createIndex may take a long time to execute, please be patient**

## Build

```sh
./build build
```

## Configuration

1.	SNAPSCALE_EXPLORER_MID_API_BASE: CHAIN API ENDPOINT
2.	SNAPSCALE_EXPLORER_MID_HTTP_PORT: Middle-ware http service port
3.	SNAPSCALE_EXPLORER_MID_WS_PORT: Middleware websocket service port
4.	SNAPSCALE_EXPLORER_MID_MONGO_CONFIG: Mongo address and certification
5.	SNAPSCALE_EXPLORER_MID_LOG_PATH: Middleware log path

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

SnapScale Explorer is released under the Apache 2.0 license and is offered “AS IS” without warranty of any kind, express or implied. Any security provided by the SnapScale software depends in part on how it is used, configured, and deployed. SnapScale is built upon many third-party libraries such as WABT (Apache License) and WAVM (BSD 3-clause) which are also provided “AS IS” without warranty of any kind.

## Important

See [LICENSE](./LICENSE) for copyright and license terms.

All repositories and other materials are provided subject to the terms of this [IMPORTANT](./IMPORTANT.md) notice and you must familiarize yourself with its terms.  The notice contains important information, limitations and restrictions relating to our software, publications, trademarks, third-party resources, and forward-looking statements.  By accessing any of our repositories and other materials, you accept and agree to the terms of the notice.
