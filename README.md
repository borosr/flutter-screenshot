# Flutter Screenshot
![example workflow name](https://github.com/borosr/flutter-screenshot/workflows/Main/badge.svg)
[![Coverage Status](https://coveralls.io/repos/github/borosr/flutter-screenshot/badge.svg?branch=main)](https://coveralls.io/github/borosr/flutter-screenshot?branch=main)
[![Go Report Card](https://goreportcard.com/badge/github.com/borosr/flutter-screenshot)](https://goreportcard.com/report/github.com/borosr/flutter-screenshot)
[![codebeat badge](https://codebeat.co/badges/8d09cfcb-78b4-4ede-8b7a-3e520d3ba95f)](https://codebeat.co/projects/github-com-borosr-flutter-screenshot-main)

A tool which helps in multi device screenshot creation for Flutter projects.

## Requirements
1. Go 1.16 or later, to install this project
2. Installed Xcode and simulator
3. Installed Android emulator, avdmanager and adb (e.x.: [Command Line Tools](https://developer.android.com/studio#downloads))
4. Flutter Integration tests
5. Extend Integration tests with screenshot calls

## Usage
### Setup
#### Install with go modules
`go get -u -t github.com/borosr/flutter-screenshot`
#### Install from release
Download one of the supported versions from [here](https://github.com/borosr/flutter-screenshot/releases). 

1. Create a configuration file in your Flutter project's root and name it `screenshots.yaml`, config example below
2. Just call `flutter-screenshot` in you Flutter project's directory
3. (Optional) Use `--verbose` after the command, to see more log messages

#### Note
- The flutter-screenshot will set the `EMU_DEVICE` environment variable before every `flutter drive` execution,
  the value will be the momentary device name from the configuration

### Configuration
#### Example
```yaml
# screenshots.yaml
command: flutter drive --target=test_driver/app.dart
devices:
  ios:
    - name: iPhone X
      mode: both # can be both, light, dark, the default value is light
    - name: iPad Pro (12.9-inch) (4th generation)
      mode: dark
  android:
    - name: Pixel_API_30
```

### In Flutter project
You should create a helper method for your UI tests, something like:
```dart
final now = new DateTime.now();
makeScreenshot(FlutterDriver driver, String filename) async {
    final imageInPixels = await driver.screenshot();
    new File('screenshots/${Platform.environment['EMU_DEVICE']}_${now.toString()}/$filename.png')
        .writeAsBytes(imageInPixels);
}
```
Then call this `makeScreenshot` method with the current FlutterDriver object and
a (test level) unique filename. The result will appear in a screenshots directory
in you Flutter project root and will contains subdirectories with the configured 
device names which will contains the captured images.


# [Licence](LICENSE)
