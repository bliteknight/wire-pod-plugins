# Greetings Plugins
This plug will respond to "good morning", "morning", "good Afternoon", "afternoon", "noon", "good night" and "night"

It can be installed like so

```
wget -O ~/wire-pod/chipper/plugins/greetings.go https://raw.githubusercontent.com/bliteknight/wire-pod-plugins/main/greetings/greetings.go
/usr/local/go/bin/go build -buildmode=plugin -o ~/wire-pod/chipper/plugins/ ~/wire-pod/chipper/plugins/greetings.go 
```

After any plugin is installed you need to retart wire-pod service to pick them up.

Once installed you will get a random greeting from 4 responses for the three diff times of day: morning, afternoon and night

If you say a greeting from the wrong time of day, Vector will correct you.

There is a slim chance he will respond with a cheeky response 🤣 - so watch out
