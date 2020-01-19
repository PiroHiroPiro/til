# blockchain

[Learn Blockchains by Building One](https://hackernoon.com/learn-blockchains-by-building-one-117428612f46?gi=4b95e4449c3a)
[ブロックチェーンを作ることで学ぶ 〜ブロックチェーンがどのように動いているのか学ぶ最速の方法は作ってみることだ〜](https://qiita.com/hidehiro98/items/841ece65d896aeaa8a2a)

## Requirement

- [Docker](https://www.docker.com/)
  - docker-compose

## Usage

Run blockchain2.py:

```console
$ docker-compose run python bash
> python blockchain2.py
```

Run cURL:

```console
$ curl -X POST -H "Content-Type: application/json" -d '{
 "sender": "d4ee26eee15148ee92c6cd394edd974e",
 "recipient": "someone-other-address",
 "amount": 5
}' "http://localhost:5000/transactions/new"
```

## Install

Clone repository:

```console
$ git clone https://github.com/PiroHiroPiro/blockchain.git
$ cd blockchain
```

Build image:

```console
$ docker-compose build
```
