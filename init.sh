#!/bin/sh

echo '
Copyright (C) 2026 dfkdream

This file is part of go_nixos_template.

go_nixos_template is free software: you can redistribute it and/or modify it under the terms of the GNU Affero General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

go_nixos_template is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License along with go_nixos_template. If not, see https://www.gnu.org/licenses/.
'

echo "Welcome to go_nixos_template initialization script!"

echo "This script will help you populate README.md, Makefile and Go modules."
echo

read -p "Enter title of your program: " program_title
sed -i "s/<your_application>/$program_title/" README.template.md
sed -i "s/<your_application>/$program_title/" Makefile
read -p "Enter your name: " name
sed -i "s/<your_name>/$name/" README.template.md

sed -i "s/<year>/$(date +%Y)/" README.template.md

echo

mv README.template.md README.md

mv cmd/your_application cmd/$program_title

go mod init $program_title

rm init.sh

echo
echo "All done! Enjoy programming!"
