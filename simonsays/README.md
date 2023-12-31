# Greetings Plugins
This plug will respond to "simon says ...." or "repeat..."

So if you were to say "simon says I am a goat", vector would respond with "I am a goat"

It can be installed like so

```
wget -O ~/wire-pod/chipper/plugins/simonsays.go https://raw.githubusercontent.com/bliteknight/wire-pod-plugins/main/simonsays/simonsays.go
/usr/local/go/bin/go build -buildmode=plugin -o ~/wire-pod/chipper/plugins/ ~/wire-pod/chipper/plugins/simonsays.go 
```

After any plugin is installed you need to retart wire-pod service to pick them up.
