@echo off
title clean steamCMD
echo cleaning steamCMD...
cd tools
rmdir /S /Q steamcmd
mkdir steamcmd
copy steamcmd.exe steamcmd\steamcmd.exe
