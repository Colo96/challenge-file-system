# Simulador de File System en Go

Este proyecto implementa un simulador básico de un sistema de archivos similar al de Linux utilizando el lenguaje de programación Go. Los usuarios pueden ejecutar comandos como `cd`, `touch`, `mkdir`, `ls`, y `pwd` para interactuar con el sistema de archivos virtual.

## Características

- **Navegación por directorios**: Usa `cd` para cambiar entre directorios, incluido el comando `..` para regresar al directorio anterior.
- **Creación de archivos**: Usa `touch` para crear archivos en el directorio actual.
- **Listar contenido**: Usa `ls` para ver el contenido (carpetas y archivos) del directorio actual.
- **Creación de carpetas**: Usa `mkdir` para crear nuevas carpetas.
- **Mostrar ruta actual**: Usa `pwd` para mostrar la ruta completa donde el usuario se encuentra.
- **Salir**: Usa `exit` para salir del simulador.

## Definiciones Adicionales y Suposiciones

- **Simulación en Memoria**: Este programa simula un sistema de archivos en memoria utilizando estructuras de datos (`Directory`, `File`), sin realizar interacciones reales con el sistema de archivos del host.
- **Comandos Simplificados**: Los comandos como `ls()` solo listan los archivos y carpetas en el directorio actual, sin explorar propiedades avanzadas de los archivos o carpetas.
- **Nombres de Archivos y Directorios**: No se aplican restricciones avanzadas de nombres de archivos o carpetas que podrían estar presentes en sistemas de archivos reales.
- **Sin Manejo Avanzado de Errores**: Esta versión del simulador no maneja casos como la creación de archivos duplicados o la verificación de permisos.

## Requisitos

- [Go](https://golang.org/doc/install) (versión 1.13 o superior)

## Instalación

1. Clona este repositorio en tu máquina local:
   ```bash
   git clone https://github.com/tu_usuario/file-system-simulator.git
   ```
