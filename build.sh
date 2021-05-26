#!/bin/bash
fyne package -icon ./images/icon.png -tags=no_native_menus
fyne-cross windows -icon ./images/icon.png -tags no_native_menus -env GOPROXY=https://goproxy.cn,direct
fyne-cross linux -icon ./images/icon.png -tags no_native_menus -env GOPROXY=https://goproxy.cn,direct