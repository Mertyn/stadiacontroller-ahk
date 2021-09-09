# Stadia Controller Autohotkey

Support for using the Stadia Controller as an Xbox 360 controller
on Windows with the addition of AutoHotkey scripting.

The Original [stadiacontroller](https://github.com/71/stadiacontroller) program was made by [Grégoire Geis (71)](https://github.com/71).

## Features
- All buttons are mapped to their Xbox 360 equivalents.
- Triggers are analog.
- The Assistant and Capture buttons can be remapped to run any AutoHotkey v2 code (see usage for more details).
- Vibrations are supported.
- Emulation via [ViGEm](https://vigem.org) (must be installed), which means that
everything just works. There won't be pesky Denuvo games that refuse to accept that input.

## Installation
1. Install [ViGEm](https://github.com/ViGEm/ViGEmBus/releases).
2. Download the [latest release](https://github.com/Mertyn/stadiacontroller-ahk/releases/latest).
3. Extract the zip into a directory.

## Usage
### Running the program
After installation of ViGEm run `stadiacontroller.exe` and the sound of a device being plugged in should be heard.

Now a connected Google Stadia controller can be used as an Xbox 360 controller.

### Remapping extra buttons
The Assistant and Capture buttons can be remapped in `stadiacontroller.ahk` using the AutoHotkey v2 scripting language. Any code inside the functions `Assistant`, `AssistantUp`, `Capture` and `CaptureUp` will be run according to the button pressed.

A simple button to keyboard key remap can be achieved as follows
```AutoHotkey
Capture() {
    Send "{PrintScreen}"
}
```
The script can only be reloaded by restarting `stadiacontroller.exe`. <br>
If any of the four functions are missing in the script they will just be ignored. <br>
In the default script there are console outputs defined according to the button presses.

Any arbitrary AutoHotkey v2 code defined in `stadiacontroller.ahk` should be able to run, so there are more possiblities than this simple button to keyboard remap. <br>
For more information see the [AutoHotkey v2 documentation](https://lexikos.github.io/v2/docs/AutoHotkey.htm). It is important to always reference the documentation for AutoHotkey version 2, as there are pretty big differences from version 1.

## Alternatives
- [Original stadiacontroller](https://github.com/71/stadiacontroller) has all the same functionality, but with Windows command support instead of AutoHotkey.
- [XOutput](https://github.com/csutorasa/XOutput) does not support vibrations,
analog triggers and additional buttons, but it has more features and is more stable overall.
