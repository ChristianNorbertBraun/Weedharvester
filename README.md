[![Build Status](https://travis-ci.org/ChristianNorbertBraun/Weedharvester.svg?branch=master)](https://travis-ci.org/ChristianNorbertBraun/Weedharvester)
# Weedharvester
An client API for SeaweedFS

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
