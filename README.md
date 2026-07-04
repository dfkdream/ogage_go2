# ogage_go2
Global hotkey listener for Odroid Go Advance

Drop-in replacement of [ArkOS](https://github.com/christianhaitian/arkos)' `ogage` hotkey listener.

See this project on GitHub: [https://github.com/dfkdream/ogage_go2](https://github.com/dfkdream/ogage_go2)

See this project on GitLab: [https://gitlab.com/dfkdream/ogage_go2](https://gitlab.com/dfkdream/ogage_go2)

## New features
### Long press
* Long press the D-Pad buttons to adjust volume/brightness quickly.
* Long press power button to shutdown (or suspend, depending on system configuration).

### Configurable key combinations
Almost every key combinations are configurable. See [Configurations](#configurations) for more.

### Configuration auto-reloading
Update the configuration and see the changes on the fly. No need to reboot the device.

### Key combination blocking (Experimental)

> [!WARNING]
> This feature is experimental, and disabled by default. 

Block key combination inputs, preventing other programs from misinterpreting them.

For example, while a game is on, you can change brightness or volume while not moving the in-game cursor.

## How to install
1. Download this file: [ogage_go2.zip](https://github.com/dfkdream/ogage_go2/releases/latest/download/ogage_go2.zip)
2. Unzip it into your `EASYROMS/tools` directory.
3. Open configuration/options and move into `Tools > ogage_go2`.
4. Run `Install`.

## Configurations
After first run, you can find default configurations in `/etc/ogage/config.yml`.

Or have a look at [/internal/config/config.yml](/internal/config/config.yml).


## References
* [https://github.com/christianhaitian/ogage](https://github.com/christianhaitian/ogage)

## How to build
```sh
make
```

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

## If you want to..
### Add new command
TL;DR: See [/internal/config/codes.go](/internal/config/codes.go), [/cmd/ogage/commands.go](/cmd/ogage/commands.go) and [/cmd/ogage/handlers.go](/cmd/ogage/handlers.go).

1. Open [/internal/config/codes.go](/internal/config/codes.go).
2. Scroll to the bottom.
3. Add new event constant. e.g. `SOMETHING_NEW = "SOMETHING_NEW"`
4. Open [/cmd/ogage/commands.go](/cmd/ogage/commands.go).
5. Write command function. e.g. `func somethingNew() {...}`
6. Open [/cmd/ogage/handlers.go](/cmd/ogage/handlers.go).
7. Find appropreate processor function and add binding code.
    * See `joypadPressProcessor` and `joypadReleaseProcessor`. Most of the commands are processed there.

### Add new button
See [/internal/evdev/codes.go](/internal/evdev/codes.go) and [/internal/config/codes.go](/internal/config/codes.go).

# License
Copyright (C) 2026 dfkdream

This file is part of ogage_go2.

ogage_go2 is free software: you can redistribute it and/or modify it under the terms of the GNU Affero General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

ogage_go2 is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License along with ogage_go2. If not, see https://www.gnu.org/licenses/.
