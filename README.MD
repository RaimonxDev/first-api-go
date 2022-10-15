Generar un clave rsa 1024 bits
```shell openssl genrsa -out app.rsa 1024 ```
Nota: -out: indica donde se va a guardar la clave privada.
1024 : Es el tamaño en bits de la clave privada

Generar una clave publica a la clave privada

```shell openssl rsa -in app.rsa -pubout > app.rsa.pub ```
Syntax:
-in: indica la ruta de la clave privada
-pubout: indica cual va ser la salida y el nombre de la clave publica generada