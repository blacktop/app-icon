# app-icon

[![Go](https://github.com/blacktop/app-icon/actions/workflows/go.yml/badge.svg)](https://github.com/blacktop/app-icon/actions/workflows/go.yml) [![Downloads](https://img.shields.io/github/downloads/blacktop/app-icon/total.svg)](https://github.com/blacktop/app-icon/releases)  [![Github All Releases](https://img.shields.io/github/release/blacktop/app-icon.svg)](https://github.com/blacktop/app-icon/releases) [![License](http://img.shields.io/:license-mit-blue.svg)](http://doge.mit-license.org)

> Generate macOS App Icons

## Install

```bash
brew install blacktop/tap/app-icon
```

## Usage

Start with a **sRGB** `1024x1024` **PNG** file and run the following command:

```bash
❯ app-icon Icon.png
```
```bash
❯ ll

total 1472
drwxr-xr-x  26 blacktop  staff   832B Sep  7 13:08 .
drwxr-xr-x  67 blacktop  staff   2.1K Sep  7 11:47 ..
-rw-r--r--   1 blacktop  staff   270K Sep  7 13:08 Icon.png
-rw-r--r--@  1 blacktop  staff    12K Sep  7 13:08 Icon_128x128.png
-rw-r--r--@  1 blacktop  staff    28K Sep  7 13:08 Icon_128x128@2x.png
-rw-r--r--@  1 blacktop  staff   887B Sep  7 13:08 Icon_16x16.png
-rw-r--r--@  1 blacktop  staff   2.4K Sep  7 13:08 Icon_16x16@2x.png
-rw-r--r--@  1 blacktop  staff    28K Sep  7 13:08 Icon_256x256.png
-rw-r--r--@  1 blacktop  staff    73K Sep  7 13:08 Icon_256x256@2x.png
-rw-r--r--@  1 blacktop  staff   2.4K Sep  7 13:08 Icon_32x32.png
-rw-r--r--@  1 blacktop  staff   5.3K Sep  7 13:08 Icon_32x32@2x.png
-rw-r--r--@  1 blacktop  staff    73K Sep  7 13:08 Icon_512x512.png
-rw-r--r--@  1 blacktop  staff   177K Sep  7 13:08 Icon_512x512@2x.png
```

## License

MIT Copyright (c) 2023 **blacktop**
