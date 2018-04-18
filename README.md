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

Any plateform can run Jarvis
- windows
- linux, and also raspberry pi :)
- ios

3.3 raspberry setup
-------------------

Configuration is stored in /etc/jarvis/jarvis.conf (it will be used with go-jarvis-service)

    pi@raspberrypi:~ $ sudo useradd -m -b /home/go-jarvis go-jarvis
    pi@raspberrypi:~ $ export GITHUB=$(curl -s https://github.com/yroffin/go-jarvis/releases/latest -s | cut -f2 -d\" | sed s:/tag/:/download/:)
    pi@raspberrypi:~ $ sudo wget ${GITHUB}/go-jarvis-0.0.1-SNAPSHOT.armel -O /home/go-jarvis/go-jarvis
    pi@raspberrypi:~ $ sudo chmod 755 /home/go-jarvis/go-jarvis
    pi@raspberrypi:~ $ sudo chown go-jarvis:go-jarvis /home/go-jarvis/go-jarvis
    pi@raspberrypi:~ $ sudo wget ${GITHUB}/go-jarvis-service -O /etc/init.d/go-jarvis-service
    pi@raspberrypi:~ $ sudo chmod 755 /etc/init.d/go-jarvis-service
    pi@raspberrypi:~ $ sudo update-rc.d go-jarvis-service defaults
    pi@raspberrypi:~ $ sudo service go-jarvis-service restart
