# go-jarvis
Just Another Ridiculous Very Inteligent System in Go

2 Developpment
==============

With gin (for live reload)

gin -p 3000 -a 8080 -i -- -http=8080

3 Setup
=======

3.1 pre-requisites
------------------

Jarvis use Golang stack with
- Open RDF database called cayley (https://cayley.io)
- SQL Lite engine (https://sqlite.org/index.html)

3.2 platforms
-------------

Any plateform can run Jarvis (it's golang based)
- windows
- linux, and also raspberry pi :)
- ios

3.3 raspberry setup
-------------------

Configuration is stored in /etc/jarvis/jarvis.conf (it will be used with go-jarvis-service)
```
jarvis.zway.password=**********
jarvis.slack.api=**********
```

A go-jarvis user
```
sudo useradd -m go-jarvis
```

And a local config.properties in go-jarvis home (/home/go-jarvis)

```
jarvis.slack.url=https://hooks.slack.com/services
jarvis.zway.url=http://192.168.1.111:8083
jarvis.rflink.comport=/dev/ttyACM0
```

And then some commands
```
sudo useradd -m go-jarvis
export GITHUB=$(curl -s https://github.com/yroffin/go-jarvis/releases/latest -s | cut -f2 -d\" | sed s:/tag/:/download/:)
sudo wget ${GITHUB}/go-jarvis-0.0.1-SNAPSHOT.armel -O /home/go-jarvis/go-jarvis
sudo chmod 755 /home/go-jarvis/go-jarvis
sudo chown go-jarvis:go-jarvis /home/go-jarvis/go-jarvis
sudo wget ${GITHUB}/go-jarvis-service -O /etc/init.d/go-jarvis-service
sudo chmod 755 /etc/init.d/go-jarvis-service
sudo update-rc.d go-jarvis-service defaults
sudo service go-jarvis-service restart
```

3.4 rflink setup
----------------

Just add go-jarvis in dialout group to access to /dev/ttyACM0

3.5 zway setup
---------------

go on https://z-wave.me :
- https://z-wave.me/z-way/download-z-way
- on stretch ... setup with http://ftp.nl.debian.org/debian/pool/main/o/openssl/libssl1.0.0_1.0.1t-1%2Bdeb8u8_armhf.deb



