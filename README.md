# task-management-api
タスク管理を行うAPIです。

## サインアップ
例えば以下のようにPOSTすることでユーザーを作成することができます。
```
curl -X POST "http://192.168.11.20:8080/signup" -d "username=USER" -d "password=PASSW0RD"
```

## ログイン
作成したユーザーを利用してログインすることができます。
```
curl -X POST "http://192.168.11.20:8080/login" -d "username=USER" -d "password=PASSW0RD"
```

上記のようなリクエストを発行すると、以下のようなレスポンスが返ってきます。
```
{"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyaWQiOjYsImFkbWluIjpmYWxzZSwiZXhwIjoxNjg0NDcxNTU2fQ.GSxYvjM9W8MF5nEQxAPiI3bvOL1-3u-_Zuz94nLeUok"}
```
こちらのトークンを Authorizationヘッダに加えることでタスク管理を行うことができます。
