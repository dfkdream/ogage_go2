#! /bin/bash

exec &> /dev/stdout

BASEDIR=$(dirname "$0")

sudo systemctl stop oga_events.service

sudo cp /usr/local/bin/ogage $BASEDIR/ogage.bak
sudo cp $BASEDIR/ogage /usr/local/bin/ogage

sudo systemctl start oga_events.service

msgbox "ogage_go2 is installed successfully. Restart the device if new ogage is not working."