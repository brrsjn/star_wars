### Tarea de start_wars

Wena cabros, les cuento que instalé esta libreria para que hagamos el seteo de variables de entorno, esto para agilizar la
puesta en produccion de lo que son los servidores y en que direccion estos estan escuchando a diferencia del local

https://github.com/joho/godotenv

En el github hay un ejemplo


Por si sale este error al ejecutar el make gen:

jbarros@jbarros-wherex:~/Universidad/star_wars$ make gen
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative pb/*.proto
protoc-gen-go: program not found or is not executable
--go_out: protoc-gen-go: Plugin failed with status code 1.
make: *** [makefile:2: gen] Error 1


Ejecutar por consola lo siguiente:

export PATH="$PATH:$(go env GOPATH)/bin"
