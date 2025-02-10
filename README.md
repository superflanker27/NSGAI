
# NSGAI


NSGAI (non-steam gmod addon installer), is a simple tool that let's you install gmod addons from the steam workshop without steam




## HOW DOES IT WORK?

It's pretty simple, it uses the steamCMD, logs-in anonymously and downloads the workshop item, then it will use the 'gmad' tool to "unpack" it and drop it in your addons folder, it also cleans up after itself unlike the offical steam workshop

## Installation

How to install

```bash
  git clone https://github.com/superflanker27/NSGAI
  cd NSGAI-master
  go run main.go
```
Or you can compile it:
```bash
  git clone https://github.com/superflanker27/NSGAI
  cd NSGAI-master
  go build
  project.exe
```
You can directly download the compiled version from the releases if you wish.
    
## Usage

- First, find the gmodlocation.cfg file and edit it, place the actual directory to gmod like this:

```javascript
C:\Users\X\Downloads\gmod\GarrysMod
```
- Select 1 to download addon

- Enter the steam workshop link, example:
```javascript
https://steamcommunity.com/sharedfiles/filedetails/?id=3421081011
```

- It will ask you for the name of the gma file, go to *tools\steamcmd\steamapps\workshop\content\4000* inside there should be a folder named the id of whatever addon you just installed, in that folder you have the gma file, it's recommended to re-name this file to something else so you can have an easier time uninstalling the addon later and to avoid messing with other addons. So type the name of the gma file **without the '.gma'**

- Done, just open your game now.

If it doesn't work then use the second option to fix the steamCMD, alternatively you can run the "cleansteamcmd.bat" file.

if it still doesn't work you can try running it from cmd so you can read the error messages, *they are not detailed*, not even remotely close to that but worth a try.

you can also go to *"tools\steamcmd\steamapps\workshop\content"* and take a look in the folders there, to see if there is any folder & file related to the addon, normally only a empty folder of the addon should be left

if you want to delete addon, just go into gmod and find the addons folder, then find the folder of the addon you want to delete and well, delete it.

## Contact

Need help? or want to contact me for any other reasons, my discord is: **kroot1**

no, im not giving you my telegram XP

