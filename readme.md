# go rocksdb #

gorocksdb: `git clone https://github.com/facebook/rocksdb.git`

# 测试 #

1. 更新子模块

```
    git submodule init
    git submodule update
```

2. 依赖库安装

```
yum groupinstall -y "Development Tools"
yum install -y snappy-devel zlib-devel
```

3. 编译

```
cd rocksdb && make static_lib
```

4. go调用

```
go build main.go

./main -H
./main -O w -K test -V kk
./main -O r -K test
./main -O d -K test
```