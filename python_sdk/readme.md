# Python SDK

⚡ Python SDK For Real-Time Robot-Human Interaction Applications ⚡

## License

[![License:Apache2.0](https://img.shields.io/badge/License-Apache2.0-yellow.svg)](https://opensource.org/licenses/Apache2.0)

## Basic Requirements

Your device should meet the following basic requirements.

```shell
Distributor ID: Ubuntu
Description:    Ubuntu 22.04.4 LTS
Release:        22.04
Codename:       jammy

Python          3.10.12
```

## Install

```shell
# Dependencies that need to be installed manually.
pip install chromadb colcon-common-extensions colorlog empy hidapi logging numpy==1.26.4 openai pvporcupine==3.0.3 pvrecorder==1.2.2 pyaudio pymongo pyserial pyusb requests soundfile tornado wave websocket-client==0.48.0 webrtcvad
```

```shell
pip install geekros
# or
pip3 install geekros
# or
python3 -m pip install geekros
```

## Build

> Build and publish Python packages to pypi.

1、Install the required dependencies.

```shell
python -m pip install twine setuptools wheel
```

2、Build software package.

```shell
python setup.py sdist bdist_wheel
```

3、Upload software package to PyPI.

```shell
twine upload dist/*
```