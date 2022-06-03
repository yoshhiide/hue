# Hue
Hue Bridgeを経由してライトをOn/Offするコマンドライン

# environments
```
HUE_BRIDGE_URL=http://192.168.0.0
HUE_BRIDGE_TOKEN=token
```

# install
`go install`

# usage
ライトの番号、ライトの状態を引数に指定します
```
# 5番目のライトをon
hue 5 on

# 3番目のライトをoff
hue 3
```

- offの場合は第二引数を指定しません。
- onの場合は第二引数になんらかを指定します。
