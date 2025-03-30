1. Connect to the websocket server 
`wscat -c http://localhost:8080/track`
- do it from 2-3 terminal sessions

<details><summary>Troubleshooting wscat</summary>

```md

if error like: 
node:internal/modules/cjs/loader:1147
  throw err;
  ^

Error: Cannot find module 'https-proxy-agent'
Require stack:
- /usr/share/nodejs/wscat/bin/wscat
    at Module._resolveFilename (node:internal/modules/cjs/loader:1144:15)
    at Module._load (node:internal/modules/cjs/loader:985:27)
    at Module.require (node:internal/modules/cjs/loader:1235:19)
    at require (node:internal/modules/helpers:176:18)
    at Object.<anonymous> (/usr/share/nodejs/wscat/bin/wscat:7:25)
    at Module._compile (node:internal/modules/cjs/loader:1376:14)
    at Module._extensions..js (node:internal/modules/cjs/loader:1435:10)
    at Module.load (node:internal/modules/cjs/loader:1207:32)
    at Module._load (node:internal/modules/cjs/loader:1023:12)
    at Function.executeUserEntryPoint [as runMain] (node:internal/modules/run_main:135:12) {
  code: 'MODULE_NOT_FOUND',
  requireStack: [ '/usr/share/nodejs/wscat/bin/wscat' ]
}

Node.js v20.10.0

- simply go to /usr/share/nodejs/wscat and do sudo npm install 

```

</details>

2. Push a location update to the Hub 
```sh
curl -X POST -H "Content-Type: application/json" -d '{"latitude": 40.334, "longitude": -74.98}' http://localhost:8080/push 
```

all ws client get this update in real time