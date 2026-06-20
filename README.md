# go_nixos_template
Go development environment template for NixOS.

## How to use
```sh
nix develop
```
### Initialize project
```sh
./init.sh
```
This script will help you populate README.md, Makefile and Go modules.

After initialization, this script and `README.template.md` will be deleted.

### Enable direnv
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

This file is part of go_nixos_template.

go_nixos_template is free software: you can redistribute it and/or modify it under the terms of the GNU Affero General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

go_nixos_template is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License along with go_nixos_template. If not, see https://www.gnu.org/licenses/.
