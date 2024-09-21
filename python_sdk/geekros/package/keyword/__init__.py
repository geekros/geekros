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

import os
from ...utils import Utils

class Keyword:

    def __init__(self):
        self.language = "en"
        self.path = "/opt/geekros/model/voice/keyword/"
        self.list = []

    def set_language(self, language="en"):
        self.language = language

    def on_init(self):
        self.list = []
        list_string = ""
        for root, dirs, files in os.walk(self.path + self.language):
            for file in files:
                if file.endswith(".ppn"):
                    abs_path = os.path.join(root, file)
                    formatted_name = os.path.splitext(file)[0]
                    list_string += formatted_name + " "
                    self.list.append((formatted_name, abs_path))
        Utils().log.warning("Current language", self.language, "Supported wake words", list_string.rstrip(" "))

    def get_by_index(self, index=None):
        if index is None:
            return None, None
        if index > len(self.list) or index < 0:
            return None, None
        return self.list[index]

