#!/usr/bin/env python
# -*- coding: utf-8 -*-

import os
import sys
import time
sys.path.insert(0, os.path.abspath(os.path.join(os.path.dirname(__file__), "..")))
from geekros import Framework

def on_start(sdk):
    sdk.utils.log.ignore("on_start")
    if sdk.hardware.microphone.device is not None:
        sdk.Package.keyword.on_init()
        sdk.Package.detection.create_porcupine(sdk.Package.keyword.list, sdk.Package.keyword.language)
        sdk.Package.detection.start_recorder(-1)
        sdk.hardware.microphone.drive.control_listen()
    while not sdk.quit_event.is_set():
        if sdk.Package.detection.recorder and len(sdk.Package.keyword.list) > 0:
            pcm = sdk.Package.detection.recorder.read()
            result = sdk.Package.detection.porcupine.process(pcm)
            if result >= 0:
                direction = sdk.hardware.microphone.drive.direction()
                sdk.hardware.microphone.drive.control_think()
                (name, path) = sdk.Package.keyword.get_by_index(result)
                sdk.utils.log.success("Detected:%s direction:%d" % (name, direction))
                time.sleep(3)
                sdk.hardware.microphone.drive.control_listen()

if __name__ == "__main__":
    framework = Framework()
    on_start(framework)
