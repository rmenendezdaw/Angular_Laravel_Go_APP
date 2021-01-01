# Play It

## Práctica Proyecto Servidor

### ¿Qué es Golang?

Go es un lenguaje de programación cncurrente y compilado, desarrollado por los ingnieros de Google. Vio la luz en el año 2009, esto hace a Go un lenguaje relativamente nuevo, pero que esto no nos engañe, Go es un lenguaje maduro, con el cual se han desarrollado miles de proyectos alrededor del mundo.

Go es un lenguaje tipado estático, una vez declaremos el tipo de dato de nuestra variable, la variable no podrá cambiar de tipo en todo el programa, eso sí, en Go no es necesario definir el tipo.

Go esta pensado para ser implementado en tareas las cuales no tengan mucha interacción con los usuarios, es más un lenguaje de sistemas que de aplicaciones. Un áreas muy recurrente en donde podemos implementar Go es del lado servidor, ya sea que nuestro programa funciones como backend, como un microservicios o se encuentre realizando tareas un poco más complejas, tal vez, procesando datos.

#### Microservicios 

La arquitectura de microservicios es un patrón de diseño que estructura las aplicaciones como una colección de servicios. En un mundo impulsado por la computación en la nube, los microservicios se han convertido en uno de los patros más populares de diseño de arquitectura. Las ventajas de utilizar microservicios son:

- Son fáciles de mantener y de poner a prueba.
- Poseen bajo acoplamiento.
- Se despliegan independientemente.
- Se organizan con base a las capacidades del negocio.
- Son propiedad de un grupo pequeño de personas.

### ¿Que és Laravel?

Laravel es uno de los frameworks de código abierto más fáciles de asimilar para PHP. Como framework resulta bastante moderno y ofrece muchas utilidades potentes a los desarrolladores, que permiten agilizar el desarrollo de las aplicaciones web.

Laravel pone énfasis en la calidad del código, la facilidad de mantenimiento y escalabilidad, lo que permite realizar proyectos desde pequeños a grandes o muy grandes. Además permite y facilita el trabajo en equipo y promueve las mejores prácticas.

Este trabaja con una arquitectura de carpetas avanzada, de modo que promueve la separación de los archivos con un orden correcto y definido, que guiará a todos los integrantes del equipo de trabajo y será un estándar a lo largo de los distintos proyectos. Por supuesto, dispone también de una arquitectura de clases también muy adecuada, que promueve la separación del código por responsabilidades. Su estilo arquitectónico es MVC.

### Refactorización de GO

Lo primero que tuvimos que hacer fue modificar el archivo docker-compose.

Establecimos el direcotrio de trabajo a "/go/src/goApp"
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