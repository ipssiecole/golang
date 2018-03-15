# Go API

les fichiers du répertoire golang sont à placé dans
le répertoire src.

## Installation des dépendances

### Install UUID
```sh
go get -u github.com/satori/go.uuid
```

### Installer le driver pour MongoDB
```sh
go get -u github.com/globalsign/mgo
```

### Installer le router Gorilla/mux
```sh
go get -u github.com/gorilla/mux
```

```sh
openssl genrsa -out ./server/ssl/server.key 2048
```

Le subject du certificat:
- C == Country
- ST == Statte
- L == Location
- O == Organisation
- CN == Common Name

```sh
openssl req \
        -new \
        -x509 \
        -sha256 \
        -days 365 \
        -subj "/C=FR/ST=IDF/L=Paris/O=ipssi/CN=localhost" \
        -key ./server/ssl/server.key \
        -out ./server/ssl/server.cert
```

La même chose mais en ligne 

```sh
openssl req -new -x509 -sha256 -days 365 -subj "/C=FR/ST=IDF/L=Paris/O=ipssi/CN=localhost" -key ./server/ssl/server.key -out ./server/ssl/server.cert
```

```sh
docker run --name mongodb -p 27017:27017 -d mongo
```
