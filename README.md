# Simulador de File System en Go

Este proyecto implementa un simulador básico de un sistema de archivos similar al de Linux utilizando el lenguaje de programación Go. Los usuarios pueden ejecutar comandos como `cd`, `touch`, `mkdir`, `ls`, y `pwd` para interactuar con el sistema de archivos virtual.

## Características

- **Navegación por directorios**: Usa `cd` para cambiar entre directorios, incluido el comando `..` para regresar al directorio anterior.
- **Creación de archivos**: Usa `touch` para crear archivos en el directorio actual.
- **Listar contenido**: Usa `ls` para ver el contenido (carpetas y archivos) del directorio actual.
- **Creación de carpetas**: Usa `mkdir` para crear nuevas carpetas.
- **Mostrar ruta actual**: Usa `pwd` para mostrar la ruta completa donde el usuario se encuentra.
- **Salir**: Usa `exit` para salir del simulador.

## Requisitos

- [Go](https://golang.org/doc/install) (versión 1.13 o superior)

## Instalación

1. Clona este repositorio en tu máquina local:
   ```bash
   git clone https://github.com/tu_usuario/file-system-simulator.git
   ```
