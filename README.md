# Ve-Price-BCV

Este es un proyecto de scraping en Go utilizando el framework Colly para obtener información de precios desde el sitio web del Banco Central de Venezuela (BCV).

## Descripción

Este proyecto está diseñado para extraer información relacionada con los precios de monedas (dólar y euro) y la fecha de valor desde la página web del BCV. Utiliza el framework Colly para realizar el scraping de manera eficiente y fácil de mantener.

## Requisitos

- [Go](https://golang.org/) instalado en tu sistema.

## Instalación

1. Clona el repositorio:

  ```bash
    git clone https://github.com/DanMarqz/ve-price-bcv.git
  ```

2. Navega al directorio del proyecto:

  ```bash
    cd ve-price-bcv
  ```

3. Ejecuta el programa:

  ```bash
    go run main.go
  ```

## Ejemplo de respuesta:

  ```bash
    {"usd_price":" 35,95930000 ","eur_price":" 39,81665411 ","value_date":"Martes, 02 Enero  2024"}
  ```

## Configuración

El programa no requiere configuración adicional. Sin embargo, asegúrate de tener una conexión a Internet estable antes de ejecutarlo.

## Contribuciones

Las contribuciones son bienvenidas. Si deseas mejorar o ampliar la funcionalidad de este proyecto, siéntete libre de enviar un pull request.

## Licencia

Este proyecto está bajo la licencia MIT.  