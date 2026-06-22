# ogage_go2
Global hotkey listener for Odroid Go Advance

Drop-in replacement of [ArkOS](https://github.com/christianhaitian/arkos)' `ogage` hotkey listener.

## New features
### Long press
* Long press the D-Pad buttons to adjust volume/brightness quickly.
* Long press power button to shutdown (or suspend, depending on system configuration).

### Work in progress
* Configurable key combinations

## References
* [https://github.com/christianhaitian/ogage](https://github.com/christianhaitian/ogage)

## How to build
```sh
make
```

## Installation
Copy the `ogage` file into your sdcard's `/usr/local/bin/` directory.

You might need root privilege.

I recommend you to backup the original `ogage`.

## How to setup development environment
```sh
nix develop
```

### With direnv
```sh
direnv allow
```

### Update `flake.lock`
```sh
nix flake update
```

### Format `flake.nix`
```sh
nix fmt flake.nix
```

# License
Copyright (C) 2026 dfkdream

This file is part of ogage_go2.

ogage_go2 is free software: you can redistribute it and/or modify it under the terms of the GNU Affero General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

ogage_go2 is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License along with ogage_go2. If not, see https://www.gnu.org/licenses/.
