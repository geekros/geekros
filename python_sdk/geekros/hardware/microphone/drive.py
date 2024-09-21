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

import struct
import usb.core
import usb.util
from ...utils import Utils

class Drive:

    TIMEOUT = 200

    def __init__(self):
        self.device = None
        self.parameter = None

    def find(self, vid=0x2886, pid=0x0018):
        self.device = usb.core.find(idVendor=vid, idProduct=pid)
        if not self.device:
            Utils().log.warning("No device found")
            self.device = None
        Utils().log.debug("Found a device")
        return self.device

    def read(self, parameter):
        if self.device:
            parameter_id = parameter[0]
            cmd = 0x80 | parameter[1]
            if parameter[2] == 'int':
                cmd |= 0x40
            length = 8
            response = self.device.ctrl_transfer(usb.util.CTRL_IN | usb.util.CTRL_TYPE_VENDOR | usb.util.CTRL_RECIPIENT_DEVICE, 0, cmd, parameter_id, length, self.TIMEOUT)
            response = struct.unpack(b'ii', response.tobytes())
            if parameter[2] == 'int':
                result = response[0]
            else:
                result = response[0] * (2.**response[1])
            return result

    def write(self, parameter, value):
        if self.device:
            if parameter[5] == 'ro':
                Utils().log.warning("Is read-only")
                return
            parameter_id = parameter[0]
            if parameter[2] == 'int':
                payload = struct.pack(b'iii', parameter[1], int(value), 1)
            else:
                payload = struct.pack(b'ifi', parameter[1], float(value), 0)
            self.device.ctrl_transfer(usb.util.CTRL_OUT | usb.util.CTRL_TYPE_VENDOR | usb.util.CTRL_RECIPIENT_DEVICE, 0, 0, parameter_id, payload, self.TIMEOUT)

    def control_listen(self, direction=None):
        self.control(2)

    def control_wakeup(self, direction=0):
        position = int((direction + 15) / (360 / 12)) % 12
        pixels = [0, 0, 0, 24 * 8] * 12
        pixels[position * 4 + 2] = 24 * 8
        self.control_show(pixels)

    def control_trace(self):
        self.control(0)

    def control_think(self):
        self.control(4)

    def control_speak(self):
        self.control(3)

    def control_spin(self):
        self.control(5)

    def control_show(self, data):
        self.control(6, data)

    def control_color(self, rgb=None, r=0, g=0, b=0):
        if rgb:
            self.control(1, [(rgb >> 16) & 0xFF, (rgb >> 8) & 0xFF, rgb & 0xFF, 0])
        else:
            self.control(1, [r, g, b, 0])

    def control_brightness(self, brightness):
        self.control(0x20, [brightness])

    def control_color_palette(self, a, b):
        self.control(0x21, [(a >> 16) & 0xFF, (a >> 8) & 0xFF, a & 0xFF, 0, (b >> 16) & 0xFF, (b >> 8) & 0xFF, b & 0xFF, 0])

    def control_vad_led(self, state):
        self.control(0x22, [state])

    def control_volume(self, volume):
        self.control(0x23, [volume])

    def change_pattern(self, pattern=None):
        if self.device:
            pass
        Utils().log.warning("Not support to change pattern")

    def control_off(self):
        self.control(1, [(0 >> 16) & 0xFF, (0 >> 8) & 0xFF, 0 & 0xFF, 0])

    def control(self, cmd, data=None):
        if data is None:
            data = [0]
        self.device.ctrl_transfer(usb.util.CTRL_OUT | usb.util.CTRL_TYPE_VENDOR | usb.util.CTRL_RECIPIENT_DEVICE, 0, cmd, 0x1C, data, self.TIMEOUT)

    def set_vad_threshold(self, value):
        self.write(self.id19_offset39(), value)

    def is_voice(self):
        return self.read(self.id19_offset32())

    def direction(self):
        return self.read(self.id20_offset0())

    def version(self):
        return self.device.ctrl_transfer(usb.util.CTRL_IN | usb.util.CTRL_TYPE_VENDOR | usb.util.CTRL_RECIPIENT_DEVICE, 0, 0x80, 0, 1, self.TIMEOUT)[0]

    def close(self):
        if self.device:
            self.control_off()
            usb.util.dispose_resources(self.device)

    def id18_offset7(self):
        self.parameter = (18, 7, 'int', 1, 0, 'rw', 'Adaptive Echo Canceler updates inhibit.', '0 = Adaptation enabled', '1 = Freeze adaptation, filter only')
        return self.parameter

    def id18_offset19(self):
        self.parameter = (18, 19, 'float', 16, 0.25, 'rw', 'Limit on norm of AEC filter coefficients')
        return self.parameter

    def id18_offset25(self):
        self.parameter = (18, 25, 'int', 1, 0, 'ro', 'AEC Path Change Detection.', '0 = false (no path change detected)', '1 = true (path change detected)')
        return self.parameter

    def id18_offset26(self):
        self.parameter = (18, 26, 'float', 0.9, 0.25, 'ro', 'Current RT60 estimate in seconds')
        return self.parameter

    def id18_offset27(self):
        self.parameter = (18, 27, 'int', 3, 0, 'rw', 'High-pass Filter on microphone signals.', '0 = OFF', '1 = ON - 70 Hz cut-off', '2 = ON - 125 Hz cut-off', '3 = ON - 180 Hz cut-off')
        return self.parameter

    def id18_offset28(self):
        self.parameter = (18, 28, 'int', 1, 0, 'rw', 'RT60 Estimation for AES. 0 = OFF 1 = ON')
        return self.parameter

    def id18_offset30(self):
        self.parameter = (18, 30, 'float', 1, 1e-09, 'rw', 'Threshold for signal detection in AEC [-inf .. 0] dBov (Default: -80dBov = 10log10(1x10-8))')
        return self.parameter

    def id18_offset31(self):
        self.parameter = (18, 31, 'int', 1, 0, 'ro', 'AEC far-end silence detection status. ', '0 = false (signal detected) ', '1 = true (silence detected)')
        return self.parameter

    def id19_offset0(self):
        self.parameter = (19, 0, 'int', 1, 0, 'rw', 'Automatic Gain Control. ', '0 = OFF ', '1 = ON')
        return self.parameter

    def id19_offset1(self):
        self.parameter = (19, 1, 'float', 1000, 1, 'rw', 'Maximum AGC gain factor. ', '[0 .. 60] dB (default 30dB = 20log10(31.6))')
        return self.parameter

    def id19_offset2(self):
        self.parameter = (19, 2, 'float', 0.99, 1e-08, 'rw', 'Target power level of the output signal. ', '[-inf .. 0] dBov (default: -23dBov = 10log10(0.005))')
        return self.parameter

    def id19_offset3(self):
        self.parameter = (19, 3, 'float', 1000, 1, 'rw', 'Current AGC gain factor. ', '[0 .. 60] dB (default: 0.0dB = 20log10(1.0))')
        return self.parameter

    def id19_offset04(self):
        self.parameter = (19, 4, 'float', 1, 0.1, 'rw', 'Ramps-up / down time-constant in seconds.')
        return self.parameter

    def id19_offset5(self):
        self.parameter = (19, 5, 'int', 1, 0, 'rw', 'Comfort Noise Insertion.', '0 = OFF', '1 = ON')
        return self.parameter

    def id19_offset6(self):
        self.parameter = (19, 6, 'int', 1, 0, 'rw', 'Adaptive beamformer updates.', '0 = Adaptation enabled', '1 = Freeze adaptation, filter only')
        return self.parameter

    def id19_offset8(self):
        self.parameter = (19, 8, 'int', 1, 0, 'rw', 'Stationary noise suppression.', '0 = OFF', '1 = ON')
        return self.parameter

    def id19_offset9(self):
        self.parameter = (19, 9, 'float', 3, 0, 'rw', 'Over-subtraction factor of stationary noise. min .. max attenuation')
        return self.parameter

    def id19_offset10(self):
        self.parameter = (19, 10, 'float', 1, 0, 'rw', 'Gain-floor for stationary noise suppression.', '[-inf .. 0] dB (default: -16dB = 20log10(0.15))')
        return self.parameter

    def id19_offset11(self):
        self.parameter = (19, 11, 'int', 1, 0, 'rw', 'Non-stationary noise suppression.', '0 = OFF', '1 = ON')
        return self.parameter

    def id19_offset12(self):
        self.parameter = (19, 12, 'float', 3, 0, 'rw', 'Over-subtraction factor of non- stationary noise. min .. max attenuation')
        return self.parameter

    def id19_offset13(self):
        self.parameter = (19, 13, 'float', 1, 0, 'rw', 'Gain-floor for non-stationary noise suppression.', '[-inf .. 0] dB (default: -10dB = 20log10(0.3))')
        return self.parameter

    def id19_offset14(self):
        self.parameter = (19, 14, 'int', 1, 0, 'rw', 'Echo suppression.', '0 = OFF', '1 = ON')
        return self.parameter

    def id19_offset15(self):
        self.parameter = (19, 15, 'float', 3, 0, 'rw', 'Over-subtraction factor of echo (direct and early components). min .. max attenuation')
        return self.parameter

    def id19_offset16(self):
        self.parameter = (19, 16, 'float', 3, 0, 'rw', 'Over-subtraction factor of echo (tail components). min .. max attenuation')
        return self.parameter

    def id19_offset17(self):
        self.parameter = (19, 17, 'float', 5, 0, 'rw', 'Over-subtraction factor of non-linear echo. min .. max attenuation')
        return self.parameter

    def id19_offset18(self):
        self.parameter = (19, 18, 'int', 1, 0, 'rw', 'Non-Linear echo attenuation.', '0 = OFF', '1 = ON')
        return self.parameter

    def id19_offset20(self):
        self.parameter = (19, 20, 'int', 2, 0, 'rw', 'Non-Linear AEC training mode.', '0 = OFF', '1 = ON - phase 1', '2 = ON - phase 2')
        return self.parameter

    def id19_offset22(self):
        self.parameter = (19, 22, 'int', 1, 0, 'ro', 'Speech detection status.', '0 = false (no speech detected)', '1 = true (speech detected)')
        return self.parameter

    def id19_offset23(self):
        self.parameter = (19, 23, 'int', 1, 0, 'ro', 'FSB Update Decision.', '0 = false (FSB was not updated)', '1 = true (FSB was updated)')
        return self.parameter

    def id19_offset24(self):
        self.parameter = (19, 24, 'int', 1, 0, 'ro', 'FSB Path Change Detection.', '0 = false (no path change detected)', '1 = true (path change detected)')
        return self.parameter

    def id19_offset29(self):
        self.parameter = (19, 29, 'int', 1, 0, 'rw', 'Transient echo suppression.', '0 = OFF', '1 = ON')
        return self.parameter

    def id19_offset32(self):
        self.parameter = (19, 32, 'int', 1, 0, 'ro', 'VAD voice activity status.', '0 = false (no voice activity)', '1 = true (voice activity)')
        return self.parameter

    def id19_offset33(self):
        self.parameter = (19, 33, 'int', 1, 0, 'rw', 'Stationary noise suppression for ASR.', '0 = OFF', '1 = ON')
        return self.parameter

    def id19_offset34(self):
        self.parameter = (19, 34, 'int', 1, 0, 'rw', 'Non-stationary noise suppression for ASR.', '0 = OFF', '1 = ON')
        return self.parameter

    def id19_offset35(self):
        self.parameter = (19, 35, 'float', 3, 0, 'rw', 'Over-subtraction factor of stationary noise for ASR. ', '[0.0 .. 3.0] (default: 1.0)')
        return self.parameter

    def id19_offset36(self):
        self.parameter = (19, 36, 'float', 3, 0, 'rw', 'Over-subtraction factor of non-stationary noise for ASR. ', '[0.0 .. 3.0] (default: 1.1)')
        return self.parameter

    def id19_offset37(self):
        self.parameter = (19, 37, 'float', 1, 0, 'rw', 'Gain-floor for stationary noise suppression for ASR.', '[-inf .. 0] dB (default: -16dB = 20log10(0.15))')
        return self.parameter

    def id19_offset38(self):
        self.parameter = (19, 38, 'float', 1, 0, 'rw', 'Gain-floor for non-stationary noise suppression for ASR.', '[-inf .. 0] dB (default: -10dB = 20log10(0.3))')
        return self.parameter

    def id19_offset39(self):
        self.parameter = (19, 39, 'float', 1000, 0, 'rw', 'Set the threshold for voice activity detection.', '[-inf .. 60] dB (default: 3.5dB 20log10(1.5))')
        return self.parameter

    def id20_offset0(self):
        self.parameter = (21, 0, 'int', 359, 0, 'ro', 'DOA angle. Current value. Orientation depends on build configuration.')
        return self.parameter