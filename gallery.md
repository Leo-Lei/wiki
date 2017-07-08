title: Gallery Example
date: 2017-07-08 20:30:00
photos:
- http://ohaq3i4w3.bkt.clouddn.com/car-31.jpg
- http://ohaq3i4w3.bkt.clouddn.com/car-10.jpg
- http://ohaq3i4w3.bkt.clouddn.com/car-18.jpg
- http://ohaq3i4w3.bkt.clouddn.com/car-26.jpg
- http://ohaq3i4w3.bkt.clouddn.com/car-27.jpg
- http://ohaq3i4w3.bkt.clouddn.com/car-18.jpg
- http://ohaq3i4w3.bkt.clouddn.com/car-4.jpg
- http://ohaq3i4w3.bkt.clouddn.com/car-9.jpg    
tags:
---
To add a photo gallery placeholder to your post, just add the following lines to your front-matter:
```yml
photos:
- https://www.hdwallpapers.net/previews/starry-night-over-corsica-986.jpg
- https://www.hdwallpapers.net/previews/water-plant-close-up-979.jpg
- ...
```
<!-- more -->
and enable `lightgallery` plugin in your `_config.yml`:
```diff
# Plugins
plugins:
+   lightgallery: true # options: true, false
    google_analytics: # enter the tracking ID for your Google Analytics
    ...
```

You can also add photos between text to create another gallery like this:

![Lion](https://www.hdwallpapers.net/previews/lion-453.jpg)

or this:

![Red Panda](https://www.hdwallpapers.net/previews/red-panda-523.jpg)

Finally, you can also use Justified Gallery to display you photos in a grid:

```diff
# Plugins
plugins:
+   justifiedgallery: true # options: true, false
    google_analytics: # enter the tracking ID for your Google Analytics
    ...
```

<div class="justified-gallery">
![Lake Prags - Italy Wallpaper](https://www.hdwallpapers.net/previews/lake-prags-italy-1053.jpg)
![MacOS Sierra Wallpaper](https://www.hdwallpapers.net/previews/macos-sierra-1021.jpg)
![Cloudy Blue Sky Wallpaper](https://www.hdwallpapers.net/previews/cloudy-blue-sky-1048.jpg)
![California Road Wallpaper](https://www.hdwallpapers.net/previews/california-road-1016.jpg)
</div>
