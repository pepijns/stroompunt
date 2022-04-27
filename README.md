# Stroompunt

Stroompunt enables you to navigate to the next and the previous slide during presentations from a remote device. This can be convenient when you want to control your presentation with, for example, your phone.

## Requirements

Machine running Stroompunt:

- Windows, Linux or Mac

Mobile device to be used as remote control:

- Access to a web browser
- Connected to the same network as the machine running Stroompunt

## Usage

1. Download the latest version of Stroompunt from [the releases page](https://github.com/pepijns/stroompunt/releases). Download the .exe version for Windows or the .linux version for Linux. If you are on Mac you have to build the executable yourself (see Building below).
2. Run Stroompunt by executing the executable you have downloaded. On Linux sudo permissions are required.
3. In the web browser on your mobile device navigate to the URL that is indicated in the output of Stroompunt on your screen.
4. Start your presentation in full screen mode.
5. You are good to go!

## Building

Go 1.16 or later is required when compiling Stroompunt yourself.

    go get
    go build

Or compile for Windows, Linux and Mac all at the same time by executing `build.sh`.

Note that Stroompunt makes use of the library micmonay/keybd_event to simulate key presses. This library depends on Apple's frameworks. This means it is not possible to cross-compile from another OS to MacOS.
