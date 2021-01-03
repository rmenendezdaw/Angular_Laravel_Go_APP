# Play It

## Práctica Proyecto Servidor

### ¿Qué es Golang?

Go es un lenguaje de programación concurrente y compilado, desarrollado por los ingenieros de Google. Vio la luz en el año 2009, esto hace a Go un lenguaje relativamente nuevo, pero que esto no nos engañe, Go es un lenguaje maduro, con el cual se han desarrollado miles de proyectos alrededor del mundo.

Go es un lenguaje tipado estático, una vez declaremos el tipo de dato de nuestra variable, la variable no podrá cambiar de tipo en todo el programa, eso sí, en Go no es necesario definir el tipo.

Go esta pensado para ser implementado en tareas las cuales no tengan mucha interacción con los usuarios, es más un lenguaje de sistemas que de aplicaciones. Un área muy recurrente en donde podemos implementar Go es del lado servidor, ya sea que nuestro programa funciona como backend, como un microservicio o se encuentre realizando tareas un poco más complejas, tal vez, procesando datos.

#### Microservicios 

La arquitectura de microservicios es un patrón de diseño que estructura las aplicaciones como una colección de servicios. En un mundo impulsado por la computación en la nube, los microservicios se han convertido en uno de los patrones más populares de diseño de arquitectura. Las ventajas de utilizar microservicios son:

- Son fáciles de mantener y de poner a prueba.
- Poseen bajo acoplamiento.
- Se despliegan independientemente.
- Se organizan en base a las capacidades del negocio.
- Son propiedad de un grupo pequeño de personas.

### ¿Que és Laravel?

Laravel es uno de los frameworks de código abierto más fáciles de asimilar para PHP. Como framework resulta bastante moderno y ofrece muchas utilidades potentes a los desarrolladores, que permiten agilizar el desarrollo de las aplicaciones web.

Laravel pone énfasis en la calidad del código, la facilidad de mantenimiento y escalabilidad, lo que permite realizar proyectos desde pequeños a grandes o muy grandes. Además permite y facilita el trabajo en equipo y promueve las mejores prácticas.

Este trabaja con una arquitectura de carpetas avanzada, de modo que promueve la separación de los archivos con un orden correcto y definido, que guiará a todos los integrantes del equipo de trabajo y será un estándar a lo largo de los distintos proyectos. Por supuesto, dispone también de una arquitectura de clases también muy adecuada, que promueve la separación del código por responsabilidades. Su estilo arquitectónico es MVC.

### Traefik

Traefik es un equilibrador de carga y proxy inverso HTTP moderno que facilita la implementación de microservicios. 
Traefik se integra con los componentes de su infraestructura existente (Docker, Swarm mode, Kubernetes, Marathon, Consul, Etcd, Rancher, Amazon ECS, ...) y se configura automáticamente y dinámicamente. 
Traefik es un Edge Router de código abierto que hace que la publicación de sus servicios sea una experiencia divertida y fácil.
Recibe solicitudes en nombre de su sistema y descubre qué componentes son responsables de manejarlas.

Traefik es compatible de forma nativa con todas las principales tecnologías de clúster, como Kubernetes, Docker, Docker Swarm, AWS, Mesos, Marathon, etc. 
Puede manejar muchos al mismo tiempo.

Para poder utilizarlo tendremos que configurar varios archivos.

### Refactorización de GO

Lo primero que tuvimos que hacer fue modificar el archivo docker-compose.

Establecemos el directorio de trabajo a "/go/src/goApp"
Asignamos el volumen a "./backend/go:/go/src/goApp"

La nueva parte de Go quedaria de la siguiente forma:

Imagen Docker-compose

Una vez tengamos los cambios realizados, actualizaremos todos los imports en Go con la nueva ruta. 
Pasará de:

github.com/rmenendezdaw/Angular_Laravel_Go_APP/backend/go/users

A

goApp/users

### Creación de microservicios en GO

Para realizar la refactorización en GO para que pase a utilizar microservicios empezaremos creando la estructura de directorios.
Dentro de GO creamos el directorio songs, y dentro de este creamos los directorios:

| Directorio | Descripción |
| ------------- | ------------- |
| common | Aqui dentro se localizaran los ficheros compartidos asi como utilidades. |
| controllers | Incluirá los controladores necesarios para llevar a cabo las distintas acciones de los endpoints.| 
| data | Incluirá las funciones necesarias para obtener la información solicitada. |
| models | Incluirá los modelos utilizados por el microservicio. |
| routers | Aquí estarán la definición de rutas y endpoints. |

Una vez tengamos la arquitectura de directorios habrá que ir añadiendo los archivos y modificar los packages e imports.

#### Microservicio Songs

 La estructura de directorios tendrá la siguiente forma:

![estructura-songs](images/estructura_songs.png)

 También tenemos que configurar el microservicio en el fichero docker-compose:

![docker-songs](images/docker-songs.png)

Dónde:
 - Partiremos de una imágen de Go llamada "golang v1.15".
 - El nombre del contenedor será go_songs_container.
 - Su directorio de trabajo será /go/src/goSongs.
 - Utilizará el puerto 3001.
 - Dispondrá de los comandos para iniciar Go.
 - Dependerá de los servicios de MYSQL y de Redis.
 - Trabajará en la red común "servidor_network".

#### Microservicio users

 La estructura de directorios tendrá la siguiente forma:

![estructura-users](images/estructura-users.png)

 También tenemos que configurar el microservicio en el fichero docker-compose:

![docker-users](images/docker-users.png)

Dónde:
 - Partiremos de la imágen de Go "golang v1.15".
 - El nombre del contenedor será go_users_container.
 - Su directorio de trabajo será /go/src/goUsers.
 - Utilizará el puerto 3000.
 - Dispondrá de los comandos para iniciar Go.
 - Dependerá de los servicios de MYSQL y de Redis.
 - Trabajará en la red común "servidor_network".


### Méctricas entre Traefik, Prometheus y Grafana

#### Traefik
Empezamos por el archivo "traefik.ylm" en el cual configuraremos los entrypoints:

![traefik.yml](images/captura3.png)

Seguimos con el docker-compose:

![traefik-compose](images/captura7.png)

Dónde:
- Partirá de una imágen de traefik v2.3.
- El nombre del contenedor será "traefik_container".
- Utilizará los puertos 80:80 y 8080:8080.
- Tendrá dos volumenes.
- Trabajará en la red común.


En los microservicios que queramos utilizar Traefik debemos utilizar labels:

![traefik-compose](images/docker-songs.png)

#### Prometheus

Cada servicio tiene su propio archivo de configuración:

![prometheus.yml](images/captura2.png)

Y tiene que estar configurado en el archivo docker-compose:

![prometheus-docker](images/prometheus-docker.png)

Dónde:
- Partirá de una imágen de prometheus v2.20.1.
- El nombre del contenedor será "prometheus_container".
- Tendrá un volumen.
- Expondrá el puerto 9090 del host y el puerto 9090 del contenedor.
- Dependerá del servicio de Traefik.
- Compartirá la red común.

Si vamos al navegador y escribimos "http://localhost:9090" nos aparecerá lo siguiente:
![prometheus-working](images/prometheuswork.png)


#### Grafana 

Debemos integrar el archivo de configuración de grafana:

![grafana.yml](images/captura1.png)

También debemos de configurara el servicio en el docker-compose:

![grafana-docker](images/grafana-docker.png)

Dónde:
- Dependerá del servicio Prometheus.
- Expondrá los puertos "3500:3000".
- El nombre del contenedor será "grafana_container".
- Partirá de la imágen de grafana v7.1.5.
- Tendrá variables de entorno que necesita grafana para poder funcionar.
- Trabajará en la red común.
- Dispondrá de dos volúmenes. 

Ahora vamos al navegador y ponemos la ruta "http://localhost:3500":

![grafana-working](images/grafanawork.png)

Para poder visualizar las métricas de Prometheus tendremos que ir a la configuración de Grafana y añadir:
![grafana-config](images/captura4.png)

Una vez añadido vamos a crear el Dashboard:

![grafana-config](images/captura5.png)

Le pondremos las métricas de traefic para las request de los microservicios:

![grafana-config](images/captura6.png)
