#!/bin/bash
# Copyright 2025 GEEKROS, Inc.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -e

lsb_release -a

architecture=$(dpkg --print-architecture)

ubuntu_code=$(lsb_release -c -s)

on_init(){

    sudo apt update -y

    sudo apt install -y vim dpkg-dev gpg curl wget git gcc make cmake gcc-arm-none-eabi libusb-1.0-0-dev openssl portaudio19-dev

    if [ ! -d "/usr/local/go/bin/" ]; then
        golang_version="1.23.10"
        sudo wget -q https://golang.google.cn/dl/go"${golang_version}".linux-"${architecture}".tar.gz && sudo tar -C /usr/local -xzf go"${golang_version}".linux-"${architecture}".tar.gz
        touch /etc/profile.d/geekros-golang.sh
        sudo sh -c 'echo "export PATH=$PATH:/usr/local/go/bin" >> /etc/profile.d/geekros-golang.sh'
        source /etc/profile.d/geekros-golang.sh
        sudo rm -rf go"${golang_version}".linux-"${architecture}".tar.gz
    fi

    # if [ ! -f "/usr/local/bin/xmake" ]; then
    #     git clone --recursive git@github.com:xmake-io/xmake.git
    #     cd ./xmake && git checkout tags/v2.9.4 && ./configure && make && sudo make install PREFIX=/usr/local
    #     cd ../ && sudo rm -rf xmake
    # fi

    if [ ! -f "/usr/local/bin/st-info" ]; then
        git clone git@github.com:stlink-org/stlink.git
        cd ./stlink && git checkout tags/v1.8.0 && make release && sudo make install && sudo ldconfig
        cd ../ && sudo rm -rf stlink
    fi

    exit 0
}

on_update(){
    
    sudo systemctl stop geekros.service > /dev/null 2>&1
    sudo systemctl disable geekros.service > /dev/null 2>&1

    /usr/local/go/bin/go env -w GOSUMDB=off
    /usr/local/go/bin/go env -w GOPATH=/tmp/golang
    /usr/local/go/bin/go env -w GOMODCACHE=/tmp/golang/pkg/mod
    /usr/local/go/bin/go env -w GOTOOLCHAIN=local
    export GO111MODULE=on && export GOPROXY=https://goproxy.io
    cd ../cmd && /usr/local/go/bin/go mod tidy && /usr/local/go/bin/go build -o ../release/main main.go
    sudo cp ../release/main /opt/geekros/release/
    sudo rm -rf ../release/main && cd ../.tools
}

on_publish(){

    version=$(grep '^const Number' ../pkg/version/version.go | sed -E 's/const Number = "(.*)"/\1/')

    datetime=$(date +%Y%m%d%H%M%S)

    depends="libopus-dev,libopusfile-dev,pkg-config,redis-server"

    if [ ! -d "debian" ]; then
        mkdir -p debian
        sudo cp -r ubuntu/* ./debian/
        sudo chmod +x debian/DEBIAN/*
        find ./debian -type f -name ".gitkeep" -exec rm -f {} +
    else
        sudo rm -rf debian && sudo rm -rf ./*.deb && mkdir -p debian
        sudo cp -r ubuntu/* ./debian/
        sudo chmod +x debian/DEBIAN/*
        find ./debian -type f -name ".gitkeep" -exec rm -f {} +
    fi

    sudo chmod +x ./debian/etc/update-motd.d/*

    sudo cp -r /usr/local/lib/libstlink* ./debian/usr/local/lib/
    sudo cp -r /usr/local/bin/st-* ./debian/usr/local/bin/
    sudo cp -r /usr/local/bin/st-* ./debian/usr/local/bin/
    sudo cp /etc/modprobe.d/stlink_v1.conf ./debian/etc/modprobe.d/
    sudo cp -r /lib/udev/rules.d/49-stlinkv* ./debian/lib/udev/rules.d/
    sudo cp -r /usr/local/share/stlink/* ./debian/usr/local/share/stlink/
    sudo cp -r /usr/local/include/stlink/* ./debian/usr/local/include/stlink/
    sudo cp -r /usr/local/share/man/man1/st-* ./debian/usr/local/share/man/man1/

    /usr/local/go/bin/go env -w GOSUMDB=off
    /usr/local/go/bin/go env -w GOPATH=/tmp/golang
    /usr/local/go/bin/go env -w GOMODCACHE=/tmp/golang/pkg/mod
    /usr/local/go/bin/go env -w GOTOOLCHAIN=local
    export GO111MODULE=on && export GOPROXY=https://goproxy.io
    cd ../cmd && /usr/local/go/bin/go mod tidy && /usr/local/go/bin/go build -o ../release/main main.go
    sudo cp ../release/main ../.tools/debian/opt/geekros/release/
    sudo rm -rf ../release/main && cd ../.tools

    sudo touch debian/DEBIAN/conffiles && sudo chmod +x debian/DEBIAN/conffiles
    sudo touch debian/DEBIAN/control && sudo chmod +x debian/DEBIAN/control

    sudo sh -c "echo 'Package: geekros' >> debian/DEBIAN/control"
    sudo sh -c "echo 'Version: $version-$datetime' >> debian/DEBIAN/control"
    sudo sh -c "echo 'Maintainer: GEEKROS <admin@geekros.com>' >> debian/DEBIAN/control"
    sudo sh -c "echo 'Homepage: https://www.geekros.com' >> debian/DEBIAN/control"
    sudo sh -c "echo 'Architecture: $architecture' >> debian/DEBIAN/control"
    sudo sh -c "echo 'Installed-Size: 1048576' >> debian/DEBIAN/control"
    sudo sh -c "echo 'Section: utils' >> debian/DEBIAN/control"
    sudo sh -c "echo 'Depends: $depends' >> debian/DEBIAN/control"
    sudo sh -c "echo 'Recommends:' >> debian/DEBIAN/control"
    sudo sh -c "echo 'Suggests:' >> debian/DEBIAN/control"
    sudo sh -c "echo 'Description: Real-Time Robot-Human Interaction Applications Development Platform' >> debian/DEBIAN/control"

    sudo dpkg --build debian && dpkg-name debian.deb

    echo "sudo scp geekros_*.deb root@ip:/data/wwwroot/mirrors.com/ubuntu/pool/main/$ubuntu_code/ && sudo rm -rf *.deb && sudo rm -rf debian"
}

case "$1" in
    init)
        on_init
        ;;
    update)
        on_update
        ;;
    publish)
        on_publish
        ;;
    *)
        echo "Error: Unknown command '$1'"
        exit 1
        ;;
esac