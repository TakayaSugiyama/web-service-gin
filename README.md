<img width="1181" alt="Screen Shot 2022-04-03 at 0 54 04" src="https://user-images.githubusercontent.com/44254887/161390987-ccb1ead5-c15e-4249-8889-706dc53dca02.png">


### 環境構築

Docker と[golang-migrate](https://github.com/golang-migrate/migrate)を使用しているので入ってない場合は brew を使って入れる。

```shell
migrate -v
v4.14.1
docker -v
Docker version 20.10.8, build 3967b7d
```

Docker のネットワークを作成する。

```
make createnetwork
```

DB を構築する。

```
make rundb
```

ちょっと待ってから、golang-migrate でマイグレーションを実行する

```
make migrateup
```

Go(gin)のサーバーを構築する。

```
make build
make run
```

### 参考

- gin tutorial
  https://go.dev/doc/tutorial/web-service-gin

- Docker
  https://docs.docker.com/language/golang/build-images/
