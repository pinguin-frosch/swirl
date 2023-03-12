# Swirl
Cambia fácilmente entre temas y fondos de las aplicaciones que quieras.
Personalmente lo uso para neovim, tmux, alacritty y vscode. Aunque también
tengo configuraciones para kde, obs, onlyoffice y google chrome.

https://user-images.githubusercontent.com/22999877/224569790-4eaefcec-d32c-4ca5-86df-aa46882bb8ff.mp4

## Funcionamiento
El programa lee un archivo de configuración para determinar qué comandos
ejecutar para cada aplicación. Como referencia pueden ver mi 
[configuración](https://github.com/pinguin-frosch/dotfiles/blob/main/swirl/.config/swirl/config.json)
actual. Se pueden usar %variables% y serán reemplazadas al tiempo de ejecución
con las variables de la aplicación o la configuración global.

## Limitaciones
Algunas aplicaciones no tienen recarga en tiempo real, así que habrá que
reiniciarlas para ver los cambios, de las que uso son: Google Chrome,
Only Office y Obs. Para neovim tuve que user `nvim --listen` para poder
controlarlo de forma remota. Más que eso, los comandos se deben crear
por el usuario, cada sistema es muy diferente para poder tener configuraciones
por defecto.

## Planes
A futuro me gustaría agregar la opción de cambiar las fuentes también, ya que
a veces me gusta cambiarla y debo configurar todas las aplicaciones manualmente.
En el fondo, voy a hacer una interfaz general para unificar todos los scripts
que tenía sueltos por todas partes para estas cosas.

## Créditos
Fue un proyecto bastante emocionante, el nombre lo tomé de una canción
de [Waterflame](https://www.youtube.com/watch?v=UZ3AbQbWl0I) del mismo nombre.
Es muy alegre, así me sentí al trabajar en esta aplicación :D
