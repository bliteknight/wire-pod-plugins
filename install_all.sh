echo "##################################"
echo "###  Installing all plugins"
echo "##################################"

wget -O ~/wire-pod/chipper/plugins/greetings.go https://raw.githubusercontent.com/bliteknight/wire-pod-plugins/main/greetings/greetings.go
/usr/local/go/bin/go build -buildmode=plugin -o ~/wire-pod/chipper/plugins/ ~/wire-pod/chipper/plugins/greetings.go 

wget -O ~/wire-pod/chipper/plugins/simonsays.go https://raw.githubusercontent.com/bliteknight/wire-pod-plugins/main/simonsays/simonsays.go
/usr/local/go/bin/go build -buildmode=plugin -o ~/wire-pod/chipper/plugins/ ~/wire-pod/chipper/plugins/simonsays.go 

wget -O ~/wire-pod/chipper/plugins/whatdate.go https://raw.githubusercontent.com/bliteknight/wire-pod-plugins/main/whatdate/whatdate.go
/usr/local/go/bin/go build -buildmode=plugin -o ~/wire-pod/chipper/plugins/ ~/wire-pod/chipper/plugins/whatdate.go 

sudo systemctl restart wire-pod