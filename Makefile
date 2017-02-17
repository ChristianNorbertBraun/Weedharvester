start-seaweed:
	 docker run -p 8080:8080 -p 9333:9333 -p 8888:8888 chrislusf/seaweedfs  server -master.port=9333 -volume.port=8080 -filer=true -volume.publicUrl http://docker:8080
start-seaweed-deamon:
	 docker run -d -p 8080:8080 -p 9333:9333 -p 8888:8888 chrislusf/seaweedfs  server -master.port=9333 -volume.port=8080 -filer=true -volume.publicUrl http://docker:8080

