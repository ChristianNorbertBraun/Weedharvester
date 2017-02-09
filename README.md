[![Build Status](https://travis-ci.org/ChristianNorbertBraun/Weedharvester.svg?branch=master)](https://travis-ci.org/ChristianNorbertBraun/Weedharvester)
# Weedharvester
An client API for SeaweedFS

## Getting Started
First Start SeaweedFS. You will need to make an entry in your `/etc/hosts` with
docker and the ip adress of your docker container.

This lib was tested for SeaweedFS version 0.70 linux

To start SeaweedFS use 

```
make start-seaweed
```

To check if everything is working correctly run tests with
```
go test
```

## Client

```
client := NewClient("http://masterurl.com")

fid, err := client.Create(reader)
reader := client.Read(fid)

```

## Filer

```
filer := NewFiler("http://filerurl.com")

err := filer.Create(reader, filename, path)
reader := filer.Read(filename, path)
```
