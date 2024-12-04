# Greetings Plugins
This plug will respond to "what time is it", "what's the time", "what time is", "time"

It can be installed like so

```
wget -O ~/wire-pod/chipper/plugins/time.go https://raw.githubusercontent.com/bliteknight/wire-pod-plugins/main/time/time.go
/usr/local/go/bin/go build -buildmode=plugin -o ~/wire-pod/chipper/plugins/ ~/wire-pod/chipper/plugins/time.go 
```

After any plugin is installed you need to retart wire-pod service to pick them up.
