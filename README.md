[![Build Status](https://travis-ci.org/ChristianNorbertBraun/Weedharvester.svg?branch=master)](https://travis-ci.org/ChristianNorbertBraun/Weedharvester)
# Weedharvester
An client API for SeaweedFS

## Getting Started
First Start SeaweedFS. You will need to make an entry in your `/etc/hosts` with
docker and the ip adress of your docker container.

This lib was tested for SeaweedFS version 0.70 linux
There is no option to set replication within the lib but you can 
always set the default replication of seaweedFS.

To start SeaweedFS use 

```
make start-seaweed
```

To check if everything is working correctly run tests with
```
go test
```
If your using a different url for master and filer just hand the base
url in at testing.

```
go test --master http://docker:9333 --filer http://docker:8888
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
