FROM swaggerapi/swagger-ui

COPY ./api/openapi.yml /openapi.yml

ENV SWAGGER_JSON=/openapi.yml