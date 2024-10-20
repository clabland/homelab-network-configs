# homelab-dialup

2600 and 9111 config files used in these videos:

https://youtu.be/AEiYyMwW8gY
https://youtu.be/xuBMaAi5wZU

The 2600 is configured as a two-line dial-up "ISP," largely inspired by The Serial Port's guide: https://serialport.org/blog/cisco-router-modem-isp-in-a-box/

The 2911 is essentially an ATA, calling one or the other of the two configured FSX ports actually in turn dials out on others, in my case to the 2600. You should be able to use the 2600 config with an ATA.

You will very likely need to adjust IP addresses for your own network.

[2600-and-9111-dialup](./homelab-dialup.jpg)
