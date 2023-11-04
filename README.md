# Swirl
Define categorías de scripts que se ejecutan a voluntad, pueden ser diferentes
por cada aplicación para controlar tu sistema como desees.  Personalmente lo
que más uso es el cambio de tema y fondo, aunque tengo scripts para cambiar
entre distribución de teclado, la interfaz y atajos de teclado.

https://github.com/pinguin-frosch/swirl/assets/22999877/aa9debc9-eeb5-4b5d-ad1c-66b98f72e8be

## Funcionamiento
El programa lee un archivo de configuración para determinar qué comandos
ejecutar para cada aplicación. Como referencia pueden ver mi 
[configuración](https://github.com/pinguin-frosch/dotfiles/blob/main/swirl/config.json)
actual. Se pueden usar %variables% y serán reemplazadas al tiempo de ejecución
con las variables de la aplicación o la configuración global.

## Limitaciones
Algunas aplicaciones no tienen recarga en tiempo real, así que habrá que
reiniciarlas para ver los cambios. Más que eso, los comandos se deben crear por
el usuario, cada sistema es muy diferente para poder tener configuraciones por
defecto.

## Planes
Me gustaría agregar una opción `--help` que lea la configuración actual y en
base a eso indique qué cosas se pueden realizar. Quizás agregar una opción para
que los comandos se ejecuten en un sistema operativo en particular. Ejecutar
los scripts por aplicación de forma paralela también está pendiente.

## Créditos
Fue un proyecto bastante emocionante, el nombre lo tomé de una canción
de [Waterflame](https://www.youtube.com/watch?v=UZ3AbQbWl0I) del mismo nombre.
Es muy alegre, así me sentí al trabajar en esta aplicación :D
