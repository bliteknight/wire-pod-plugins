# Greetings Plugins
This plug will respond to "what day is it", "date today", "date", "what's today's date"

It can be installed like so

```
wget -O ~/wire-pod/chipper/plugins/whatdate.go https://raw.githubusercontent.com/bliteknight/wire-pod-plugins/main/whatdate/whatdate.go
/usr/local/go/bin/go build -buildmode=plugin -o ~/wire-pod/chipper/plugins/ ~/wire-pod/chipper/plugins/whatdate.go 
```

After any plugin is installed you need to retart wire-pod service to pick them up.
