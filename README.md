# homerun2m3u


This utility reads the channel list of a Homerun / DigitalSilicon settop box and converts it to a M3U playlist which can be used in tvheadend or VLC
Usage:

$ ./homerun2m3u [IP ADDRESS OF HD HOMERUN], e.g. ./homerun2m3u 10.0.1.6 >playlist.m3u

Will produce :

#EXTM3U
#EXTINF:-1,2.1,KATU-HD
http://10.0.1.6:5004/auto/v2.1

Import this file into tvheadend or VLC.

Written by Sascha Siekmann in 2018.
