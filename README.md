# Flutter Screenshot
![example workflow name](https://github.com/borosr/flutter-screenshot/workflows/Main/badge.svg)
[![Coverage Status](https://coveralls.io/repos/github/borosr/flutter-screenshot/badge.svg)](https://coveralls.io/github/borosr/flutter-screenshot)

A tool which helps in multi device screenshot creation for Flutter projects.

**Currently Android is not supported, but it planned in the nearest future!**

## Requirements
1. Go 1.15 or later, to install this project
2. Installed Xcode and simulator
3. (Optional) Installed Android emulator and avdmanager
4. Flutter Integration tests
5. Extend Integration tests with screenshot calls

## Usage
### Install
1. `go get -u -t github.com/borosr/flutter-screenshot`
2. Create a configuration file in your Flutter project's root and name it `screenshots.yaml`, config example below
3. Just call `flutter-screenshot` in you Flutter project's directory
4. (Optional) Use `--verbose` after the command, to see more log messages

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
  # android: # not supported yet
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
