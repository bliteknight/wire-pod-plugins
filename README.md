# Wire-Pod Plugins
Custom plugins for wire-pod. You can install specific plugins by downloading the plugin and compiling it on your wire-pod server

Like so
```
wget -O ~/wire-pod/chipper/plugins/whatdate.go https://raw.githubusercontent.com/bliteknight/wire-pod-plugins/main/whatdate/whatdate.go
/usr/local/go/bin/go build -buildmode=plugin -o ~/wire-pod/chipper/plugins/ ~/wire-pod/chipper/plugins/whatdate.go 
```

After any plugin in installed you need to retart wire-pod to pick them up.

All Plugins are only compatible with wire-pod as of 01-18-2024

If you install a plug-in and it does not work, you may need to update your wire-pod installation.


# Install All Plugins
This will execute the bash script to install all the plugins

```
wget -O - https://raw.githubusercontent.com/bliteknight/wire-pod-plugins/main/install_all.sh | bash
```

# Contributing plugins
If you want to add a plugin to this repository, all you need to do is:
* Fork this repo, create a folder for your plugin
* create README detailing what you plugin does, the utterances and actions
* create a pull request and I'll merge your plugin
