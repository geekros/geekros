# Copyright 2024 GEEKROS, Inc.
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

import sys
import signal
from threading import Thread, Event
from .utils import Utils
from .hardware import Hardware
from .package import Package

class Framework:

    def __init__(self):
        signal.signal(signal.SIGINT, self.sigint_handler)
        signal.signal(signal.SIGTERM, self.sigint_handler)
        self.quit_event = Event()
        self.quit_thread = Thread(name="quit_thread_task", target=self.quit_task)
        self.quit_thread.daemon = True
        self.quit_thread.start()
        self.utils = Utils()
        self.hardware = Hardware()
        self.Package = Package()

    def quit_task(self):
        while not self.quit_event.is_set():
            pass

    def sigint_handler(self, signum, frame):
        self.hardware.microphone.close()
        self.Package.detection.stop()
        self.quit_event.set()
        sys.exit()
        


