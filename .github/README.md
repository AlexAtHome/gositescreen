# a little script that takes a screenshot of a website

A downscaled version of [sitealert](https://github.com/AlexAtHome/sitealert) that only takes a screenshot of a website.

Requires Selenium and Geckodriver to work. Selenium should run in the background. On Arch Linux, you can do the following:
```sh
paru -S selenium-driver-standalone geckodriver
java -jar /usr/share/selenium-server/selenium-server-standalone.jar
```

For now, only Firefox screenshots are implemented.

<!-- TODO: Use headless browser -->
<!-- TODO: Add Chrome support -->
