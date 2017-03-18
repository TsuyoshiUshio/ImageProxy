# ImageProxy

Image Proxy Server written by go lang. This proxy server convert jpeg image into png image. 
Now I just support both png and jpeg. Not for gif or other formats.

A png image is not converted. Only jpeg will be encoded.

# Usage

```
http://{server url}:{port}/convertedimage/{url with url encoding}

example:
http://localhost:8090/convertedimage/https%3A%2F%2Fc.s-microsoft.com%2Fja-jp%2FCMSImages%2Fspk-ushio.jpg%3Fversion%3Df6328834-736f-4b59-2f85-7b198346ef4d
```
This repo is created only for learning pupose. I wrote this code for converting img into png format. 
Currently, Bot framework thumbnail supports png only. :)
