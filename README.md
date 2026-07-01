# Dijkstra en Go

Migración del proyecto "Algoritmo de Dijkstra en Java" (Spring Boot + Thymeleaf) a Go.

## Estructura

- cmd/server        -> punto de entrada de la aplicación
- internal/model     -> Nodo, Arista, Grafo
- internal/algoritmo -> implementación de Dijkstra
- internal/handler   -> controladores HTTP (equivalente a @Controller)
- web/templates      -> plantillas HTML
- web/static         -> CSS / JS

## Ejecutar (una vez implementado)

    go run ./cmd/server
