![goandroid_logo](goandroid_logo.png)

goandroid
=========
[![Project Status](http://stillmaintained.com/kunaldawn/goandroid.png)](https://stillmaintained.com/kunaldawn/goandroid) [![Build Status](http://img.shields.io/travis/kunaldawn/goandroid.svg?style=flat-square)](https://travis-ci.org/kunaldawn/goandroid) [![Coverage Status](http://img.shields.io/coveralls/kunaldawn/goandroid.svg?style=flat-square)](https://coveralls.io/r/kunaldawn/goandroid) [![Issues](http://img.shields.io/github/issues/kunaldawn/goandroid.svg?style=flat-square)](https://github.com/kunaldawn/goandroid/issues) [![Documentation](http://img.shields.io/badge/go-Documentation-blue.svg?style=flat-square)](https://godoc.org/github.com/kunaldawn/goandroid) [![License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)](https://github.com/kunaldawn/goandroid/blob/master/LICENSE)

Introduction
------------
**"goandroid" is an Android automation library purely written in GO.**

**Project Status :** Under Development, not yet ready for v1.0 release.
**TODO List for v1.0 :**
- [ ] Implement remaning generic fearures that are marked for v1.0.
- [ ] Complete Documentation.
- [ ] Fix public API and make release for v1.0, after that only new API will be added.
- [ ] Write test code for emulator so that each test can be performed over Travis CI.
- [ ] Create examples.
- [ ] Provide documentation in README file.

Whether you are an android developer developing applications for your project and want to do some automation tasks on your application to reduce some manual human work over multiple devices/emulators, or an enthusiastic developer who want to do some cool automation taks on your android device, this library allows you to write automation code and do cool stuffs on your android device.

This project was inspired by [xiaocong/uiautomator](https://github.com/xiaocong/uiautomator) which is a python wrapper for [Android UI Automator](https://developer.android.com/tools/testing-support-library/index.html#UIAutomator) java library.

Please note that this is not an "Android UI Automation Test Framework", though it can be used as one, if you want for your purpose.

The main implementation difference between [xiaocong/uiautomator](https://github.com/xiaocong/uiautomator) and **goandroid** library is following:

|                               | goandroid     | uiautomator   |
| ----------------------------- | ------------- | ------------- |
| Uses any Java code ?          | **No**        | Yes           |
| Any APK installed on device ? | **No**        | Yes           |
| How does it work ?            | This library is purely based on top of **adb** tool, all features provied by this library uses only adb command to achive that. No android instrumentation or android ui automation java library is used as backend. This may limit features provided by this library, but  | This library installs an APK on device which starts a http server and listens for RPC calls. This http server backend is based on [android-uiautomator-server](https://github.com/xiaocong/android-uiautomator-server) and uses android ui automation framework for its internal implementation. Each call on python code invokes some method in java code via RPC using this http server running on the device. |


#### Any contributions / pull requests are always welcome.
