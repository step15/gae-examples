# gae-examples
Example apps for Google App Engine
Google App Engine用の例文。

実際にdeployするのでしたら、app.yamlのapplication:の行を自分のapplication名に変えるのを忘れないでください！

# 使い方

Mac OS XまたはWindowsなら、App Engine SDKのGUIで好き言語のAppをimportをして、「Run」のボタンを押せます。
LinuxだったらGUIはないので、以下のようなコマンドを使ってください。

## Go言語

実際のAppの中身がfoo.goに入れました（package名やファイル名が何でもいいですが）。
### 試しで実行

```sh
goapp serve gotest
```

### App engineにdeploy

```sh
goapp deploy gotest
```

### go testについて

Go言語のunit testが分かりやすいアピールをしたかったので、foo_test.goも入れておきました。
実行するのに、これでできます。
```sh
go test gotest/*.go
```

## PythonとPHP
### 試しで実行

```sh
dev_appserver.py pytest
```

### App engineにdeploy

```sh
appcfg update pytest
```

### PHPの注意点

PHPだと自動的にPHPを見つからない場合があるので、

```sh
dev_appserver.py --php_executable_path=/usr/bin/php-cgi phptest
```

のように実行する必要があるかもしれません。詳しくはhttps://cloud.google.com/appengine/docs/php/gettingstarted/helloworld も参考してください。
