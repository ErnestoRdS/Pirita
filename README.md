# 🚖 Pirita
Proyecto para manejar taxis, ubers o vehículos particulares, hecho con Svelte y Go.

## Approach 🚀
Para lograr los objetivos del proyecto, se utilizará una arquitectura de microservicios, lo que permitirá una mayor escalabilidad y flexibilidad. Svelte se utilizará para la interfaz de usuario y Go se utilizará para la lógica del servidor. 🏗️🕸️

## Outline 📑
El proyecto constará de los siguientes componentes principales:

- Componente de control de conductores
  - Lista de conductores 🚕📝
  - Formulario para agregar/editar conductores 🚕📝📝
  - Rastreador de ubicación de conductores 🚕📍🚨
- Componente de control de pasajeros
  - Registro de usuarios 🙋‍♀️📝
  - Formulario de reserva de viajes 🙋‍♀️🚗📝
  - Sistema de calificación de conductores 🙋‍♀️🚗🌟
- Componente de pagos
  - Integración con pasarelas de pago 💳💻
  - Historial de transacciones 💰📝
  

## Arquitectura 🏗️
La arquitectura del sistema estará basada en microservicios y constará de los siguientes componentes principales:

- API Gateway: Se encargará de manejar las solicitudes de los usuarios y direccionarlas al microservicio correspondiente.
- Microservicio de control de conductores: Se encargará de la gestión de conductores y su información.
- Microservicio de control de pasajeros: Se encargará de la gestión de usuarios y sus viajes.
- Microservicio de pagos: Se encargará de la gestión de pagos y transacciones.

Con esta arquitectura, el sistema será fácilmente escalable y flexible, lo que permitirá agregar nuevos componentes en el futuro si es necesario.

## Licencia 📜

Este proyecto se encuentra bajo la licencia GPL-3 para individuos y empresas pequeñas. Las empresas medianas y superiores deben usarlo bajo la licencia AGPL-3 o comprar un permiso para obtener una copia con licencia GPL-3 que puedan usar internamente sin hacer público el código, pero conservando sus libertades. 💼
