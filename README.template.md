# <your_application>
Enter your description here.

## How to build
```sh
go build -o . ./...
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

# License
Copyright (C) <year> <your_name>

This file is part of <your_application>.

<your_application> is free software: you can redistribute it and/or modify it under the terms of the GNU Affero General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

<your_application> is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License along with <your_application>. If not, see https://www.gnu.org/licenses/.
