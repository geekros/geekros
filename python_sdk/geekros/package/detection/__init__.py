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

import random
import pyaudio
import pvporcupine
import numpy as np
from pvrecorder import PvRecorder

class Detection:

    def __init__(self):
        self.pyaudio = pyaudio.PyAudio()
        self.access_key = []
        self.porcupine = None
        self.recorder = None

    def random_access_key(self):
        access_key = [
            "PSMtKVzOysHScGy5" + "g2mgc2ClrX1Xf/PYYafb4o7kQ" + "QMUsZmrbBAR1Q==",
            "zTGV3YMeYwYREJkJ2" + "BcHym1/VUkH/Ei2yB45M64c+" + "KW1LiLGLjsj7Q==",
            "FB7jo4vQwScNqXa5J" + "2E+/nCXw3wag8lSs4oj6Rpf" + "nJCK6NZopYcxSg==",
            "cN572/y58OER" + "we04A7PoCohuFcdP21y1kZ" + "UgF5JYDZpivfqh63kpvw==",
            "HpeCultXwGTuGx7" + "3JSVdAOF1yJx1r5oV3HI5oMBa5m" + "CxriF4QlYYSQ==",
            "5fVf05t0LvVl0F2zbnuSUAA" + "BmRpi4HucWIMuAiCMxIMl" + "GpFZQOA6SQ==",
            "u/kod4bBbp87SvR+Hnf" + "TE8jalPhs9pO6BKHgURpX0hcZv" + "09roVGPZQ==",
            "kj/04yPb1D3qtG" + "9dC5Lxs8PxQidMWBwonaxdtd" + "bAbYeTz+1aI2GA7A==",
            "PkFORWS+ELPnVKmh+" + "oSdWz57LEFP5+/8TRn+" + "kUAWwTpHRbK14uJoWw==",
            "NjvbZkSQ/MQ8WGFNX" + "i71elnRsI1KaaaN2jUa5AfCWy" + "4enks4KfC0rQ==",
            "E+N7gJdK80YIia+XDsf" + "o4YQtQHciMnTG68Fd9x++" + "+84y8M+F7Z75pA==",
            "MN667eSAK/O+tM+" + "fXUeOaOhI3uap8rOuiTgOwKwm" + "7JUy3eiOZ6lnpg=="
        ]
        self.access_key = random.choice(access_key)
        return self.access_key

    def create_porcupine(self, keyword=None, language="en"):
        if keyword:
            keyword_paths = []
            sensitivities = []
            language_model = None
            if language == "zh":
                language_model = "/opt/geekros/model/voice/language/porcupine_params_zh.pv"
            for name, path in keyword:
                keyword_paths.append(path)
                sensitivities.append(0.5)
            self.porcupine = pvporcupine.create(
                access_key=self.random_access_key(),
                library_path=None,
                model_path=language_model,
                keyword_paths=keyword_paths,
                sensitivities=sensitivities
            )

    def start_recorder(self, device_index=-1):
        if self.porcupine:
            self.recorder = PvRecorder(
                frame_length=self.porcupine.frame_length,
                device_index=device_index
            )
            self.recorder.start()

    def stop(self):
        if self.porcupine:
            self.porcupine.delete()
        if self.recorder:
            self.recorder.delete()
