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

import inspect
from datetime import datetime

class Loging:

    def __init__(self):
        self.current_time = datetime.now().strftime("%Y-%m-%d %H:%M:%S.%f")[:-3]
        self.class_name = ""
        self.module_name = ""
        self.function_name = ""

    def _get_stack(self, caller_frame=None):
        if caller_frame:
            module = inspect.getmodule(caller_frame[0])
            self.module_name = module.__name__ if module else "<unknown>"
            self.function_name = caller_frame.function if caller_frame.function else "<unknown>"
            if "self" in caller_frame.frame.f_locals:
                class_name = caller_frame.frame.f_locals["self"].__class__.__name__
                self.function_name = f"{class_name}.{self.function_name}"
            self.current_time = datetime.now().strftime("%Y-%m-%d %H:%M:%S.%f")[:-3]

    def default(self, *log):
        log = " ".join(map(str, log))
        self._get_stack(inspect.stack()[1])
        print(f"[{self.current_time}][{self.module_name}.{self.function_name}] \033[97m{log}\033[0m")

    def ignore(self, *log):
        log = " ".join(map(str, log))
        self._get_stack(inspect.stack()[1])
        print(f"[{self.current_time}][{self.module_name}.{self.function_name}] \033[90m{log}\033[0m")

    def debug(self, *log):
        log = " ".join(map(str, log))
        self._get_stack(inspect.stack()[1])
        print(f"[{self.current_time}][{self.module_name}.{self.function_name}] \033[94m{log}\033[0m")

    def info(self, *log):
        log = " ".join(map(str, log))
        self._get_stack(inspect.stack()[1])
        print(f"[{self.current_time}][{self.module_name}.{self.function_name}] \033[96m{log}\033[0m")

    def success(self, *log):
        log = " ".join(map(str, log))
        self._get_stack(inspect.stack()[1])
        print(f"[{self.current_time}][{self.module_name}.{self.function_name}] \033[92m{log}\033[0m")

    def warning(self, *log):
        log = " ".join(map(str, log))
        self._get_stack(inspect.stack()[1])
        print(f"[{self.current_time}][{self.module_name}.{self.function_name}] \033[93m{log}\033[0m")

    def error(self, *log):
        log = " ".join(map(str, log))
        self._get_stack(inspect.stack()[1])
        print(f"[{self.current_time}][{self.module_name}.{self.function_name}] \033[91m{log}\033[0m")
