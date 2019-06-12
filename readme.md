# go rocksdb #

# rocksdb 安装&编译 #

1. `git clone https://github.com/facebook/rocksdb.git`

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