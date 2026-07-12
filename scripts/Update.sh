#! /bin/bash

# Copyright (C) 2026 dfkdream
# 
# This file is part of ogage_go2.
# 
# ogage_go2 is free software: you can redistribute it and/or modify it under the terms of the GNU Affero General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.
# 
# ogage_go2 is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU Affero General Public License for more details.
# 
# You should have received a copy of the GNU Affero General Public License along with ogage_go2. If not, see https://www.gnu.org/licenses/.

exec &> /dev/stdout

BASEDIR=$(dirname "$0")
BACKUP="$HOME/.ogage_go2/ogage.bak"

if [ ! -e $BACKUP ]
then
    msgbox "ogage.bak not found. Check if ogage_go2 is installed. Aborting."
    exit 1
fi

sudo systemctl stop oga_events.service

sudo cp $BASEDIR/ogage /usr/local/bin/ogage

sudo systemctl start oga_events.service

retval=$?

if [ $retval -eq 0 ]
then
    msgbox "ogage_go2 is updated successfully. Restart the device if new ogage is not working."
else
    msgbox "Something went wrong. Error code $retval"
fi
