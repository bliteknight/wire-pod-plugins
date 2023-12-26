# Wire-Pod Plugins
Custom plugins for wire-pod. You can install specific plugins by downloading the plugin and compiling it on your wire-pod server

Like so
```
wget -O ~/wire-pod/chipper/plugins/whatdate.go https://raw.githubusercontent.com/bliteknight/wire-pod-plugins/main/whatdate/whatdate.go
/usr/local/go/bin/go build -buildmode=plugin -o ~/wire-pod/chipper/plugins/ ~/wire-pod/chipper/plugins/whatdate.go 
```

After any plugin in installed you need to retart wire-pod to pick them up.

# Contributing plugins
If you want to add a plugin to this repository, all you need to do is:
* Fork this repo, create a folder for your plugin
* create README detailing what you plugin does, the utterances and actions
* create a pull request and I'll merge your plugin
